package authorization

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/rs/zerolog/log"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"
	"time"
)

var cachedFirecloudSuitability map[string]*Suitability
var lastFirecloudCacheTime time.Time

func KeepFirecloudCacheUpdated(ctx context.Context) {
	interval := time.Duration(config.Config.MustInt("auth.updateIntervalMinutes")) * time.Minute
	for {
		time.Sleep(interval)
		if err := CacheFirecloudSuitability(ctx); err != nil {
			log.Warn().Err(err).Msgf("failed to update suitability cache, now %s stale", time.Since(lastFirecloudCacheTime).String())
		}
	}
}

func CacheFirecloudSuitability(ctx context.Context) error {
	adminService, err := admin.NewService(ctx, option.WithScopes(admin.AdminDirectoryUserReadonlyScope, admin.AdminDirectoryGroupMemberReadonlyScope))
	if err != nil {
		return fmt.Errorf("failed to authenticate to Google Workspace: %w", err)
	}

	var fcAdminsGroupEmails []string
	err = adminService.Members.List(config.Config.MustString("auth.firecloud.groups.fcAdmins")).Pages(ctx, func(members *admin.Members) error {
		if members == nil {
			return fmt.Errorf("cacheFirecloudSuitability got a nil %s member page from Google", config.Config.MustString("auth.firecloud.groups.fcAdmins"))
		} else {
			for _, member := range members.Members {
				if member == nil {
					return fmt.Errorf("cacheFirecloudSuitability got a nil %s member from Google", config.Config.MustString("auth.firecloud.groups.fcAdmins"))
				} else {
					fcAdminsGroupEmails = append(fcAdminsGroupEmails, member.Email)
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	var firecloudProjectOwnersGroupEmails []string
	err = adminService.Members.List(config.Config.MustString("auth.firecloud.groups.firecloudProjectOwners")).Pages(ctx, func(members *admin.Members) error {
		if members == nil {
			return fmt.Errorf("cacheFirecloudSuitability got a nil %s member page from Google", config.Config.MustString("auth.firecloud.groups.fcAdmins"))
		} else {
			for _, member := range members.Members {
				if member == nil {
					return fmt.Errorf("cacheFirecloudSuitability got a nil %s member from Google", config.Config.MustString("auth.firecloud.groups.fcAdmins"))
				} else {
					firecloudProjectOwnersGroupEmails = append(firecloudProjectOwnersGroupEmails, member.Email)
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	newCache := make(map[string]*Suitability)
	err = adminService.Users.List().Domain(config.Config.MustString("auth.firecloud.domain")).Pages(ctx, func(workspaceUsers *admin.Users) error {
		if workspaceUsers == nil {
			return fmt.Errorf("cacheFirecloudSuitability got a nil user page from Google")
		} else {
			for _, workspaceUser := range workspaceUsers.Users {
				if workspaceUser == nil {
					return fmt.Errorf("cacheFirecloudSuitability got a nil user from Google")
				} else {
					suitability := parseSuitability(workspaceUser, fcAdminsGroupEmails, firecloudProjectOwnersGroupEmails)
					if workspaceUser.PrimaryEmail != "" {
						newCache[workspaceUser.PrimaryEmail] = suitability
					}
					if workspaceUser.RecoveryEmail != "" {
						newCache[workspaceUser.RecoveryEmail] = suitability
					}

					// Secondary emails on the user's account aren't `admin.User.RecoveryEmail`, they're under
					// `admin.User.Emails`.
					//
					// Google doesn't bother typing the `admin.User.Emails` field; it's just `interface{}`.
					// Because Go is impressively bad at handling JSON, we can't easily get from `interface{}` to
					// the `[]admin.UserEmail` type we want, despite what Google's own engineers say
					// (https://github.com/googleapis/google-api-go-client/issues/325). GoLand can open up a scratch
					// file in sherlock's context with its dependencies if you want to see the panic for yourself.
					//
					// We could probably use the MapStructure package here but the rest of Sherlock doesn't use it.
					// Instead, we do the dumb-but-correct thing and serialize it back to JSON and parse back to what
					// we want.
					//
					// In theory, this madness will be somewhat short-lived, because Sherlock will become the source
					// of truth and will be more concerned with pushing info to Google Workspace than reading from it.
					if emailsJson, emailsParseErr := json.Marshal(workspaceUser.Emails); emailsParseErr != nil {
						log.Debug().Err(err).Msgf("AUTH | wasn't able to marshal %s's `emails` field back to JSON: %v", workspaceUser.PrimaryEmail, err)
					} else {
						var parsedEmails []admin.UserEmail
						if emailsParseErr = json.Unmarshal(emailsJson, &parsedEmails); emailsParseErr != nil {
							log.Debug().Err(err).Msgf("AUTH | wasn't able to unmarshal %s's `emails` field to %T: %v", workspaceUser.PrimaryEmail, parsedEmails, err)
						} else {
							for _, parsedEmail := range parsedEmails {
								if len(parsedEmail.Address) == 0 {
									log.Debug().Msgf("AUTH | one of %s's `emails` had an empty address", workspaceUser.PrimaryEmail)
								} else if parsedEmail.Address != workspaceUser.PrimaryEmail && parsedEmail.Address != workspaceUser.RecoveryEmail {
									// Only bother with the assignment if it wasn't an email we would've already recorded.
									newCache[parsedEmail.Address] = suitability
								}
							}
						}
					}
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	log.Info().Msgf("AUTH | firecloud suitability cache updated, now contains %d total (primary+recovery) emails", len(newCache))
	cachedFirecloudSuitability = newCache
	lastFirecloudCacheTime = time.Now()
	return nil
}

func parseSuitability(workspaceUser *admin.User, fcAdminsGroupEmails []string, firecloudProjectOwnersGroupEmails []string) *Suitability {
	if workspaceUser.PrimaryEmail == "" {
		return &Suitability{
			suitable:    false,
			description: "firecloud user doesn't appear to have a primary email? something's amiss, marking as not suitable",
			source:      FIRECLOUD,
		}
	} else if !workspaceUser.AgreedToTerms {
		return &Suitability{
			suitable: false,
			description: fmt.Sprintf("firecloud user hasn't accepted Google Workspace terms (suggesting they've never logged in; they'll need to wait %d minutes after first login for Sherlock to pick it up)",
				config.Config.MustInt("auth.updateIntervalMinutes")),
			source: FIRECLOUD,
		}
	} else if !workspaceUser.IsEnrolledIn2Sv {
		return &Suitability{
			suitable:    false,
			description: "firecloud user hasn't enrolled in two-factor authentication",
			source:      FIRECLOUD,
		}
	} else if workspaceUser.Suspended {
		return &Suitability{
			suitable: false,
			description: fmt.Sprintf("firecloud user is suspended, probably due to inactivity (reach out to #dsp-devops-champions for help; they'll need to wait %d minutes after reactivation for Sherlock to pick it up)",
				config.Config.MustInt("auth.updateIntervalMinutes")),
			source: FIRECLOUD,
		}
	} else if workspaceUser.Archived {
		return &Suitability{
			suitable:    false,
			description: "firecloud user is archived",
			source:      FIRECLOUD,
		}
	} else if !utils.Contains(fcAdminsGroupEmails, workspaceUser.PrimaryEmail) {
		return &Suitability{
			suitable: false,
			description: fmt.Sprintf("firecloud user isn't in fc-admins group (reach out to #dsp-devops-champions for help; they'll need to wait %d minutes after being added for Sherlock to pick it up)",
				config.Config.MustInt("auth.updateIntervalMinutes")),
			source: FIRECLOUD,
		}
	} else if !utils.Contains(firecloudProjectOwnersGroupEmails, workspaceUser.PrimaryEmail) {
		return &Suitability{
			suitable: false,
			description: fmt.Sprintf("firecloud user isn't in firecloud-project-owners group (reach out to #dsp-devops-champions for help; they'll need to wait %d minutes after being added for Sherlock to pick it up)",
				config.Config.MustInt("auth.updateIntervalMinutes")),
			source: FIRECLOUD,
		}
	} else {
		return &Suitability{
			suitable:    true,
			description: "firecloud user is suitable",
			source:      FIRECLOUD,
		}
	}
}
