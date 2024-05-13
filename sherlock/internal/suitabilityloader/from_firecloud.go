package suitabilityloader

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/rs/zerolog/log"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"
)

func fromFirecloud(ctx context.Context) ([]models.Suitability, error) {
	adminService, err := admin.NewService(ctx, option.WithScopes(admin.AdminDirectoryUserReadonlyScope, admin.AdminDirectoryGroupMemberReadonlyScope))
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate to Google Workspace: %w", err)
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
		return nil, err
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
		return nil, err
	}

	resultSet := make(map[string]models.Suitability)
	err = adminService.Users.List().Domain(config.Config.MustString("auth.firecloud.domain")).Pages(ctx, func(workspaceUsers *admin.Users) error {
		if workspaceUsers == nil {
			return fmt.Errorf("cacheFirecloudSuitability got a nil user page from Google")
		} else {
			for _, workspaceUser := range workspaceUsers.Users {
				if workspaceUser == nil {
					return fmt.Errorf("cacheFirecloudSuitability got a nil user from Google")
				} else {
					suitable, description := parseFirecloudUser(workspaceUser, fcAdminsGroupEmails, firecloudProjectOwnersGroupEmails)
					if workspaceUser.PrimaryEmail != "" {
						resultSet[workspaceUser.PrimaryEmail] = models.Suitability{
							Email:       &workspaceUser.PrimaryEmail,
							Suitable:    &suitable,
							Description: &description,
						}
					}
					if workspaceUser.RecoveryEmail != "" {
						resultSet[workspaceUser.RecoveryEmail] = models.Suitability{
							Email:       &workspaceUser.RecoveryEmail,
							Suitable:    &suitable,
							Description: &description,
						}
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
									resultSet[parsedEmail.Address] = models.Suitability{
										Email:       &parsedEmail.Address,
										Suitable:    &suitable,
										Description: &description,
									}
								}
							}
						}
					}
				}
			}
		}
		return nil
	})
	result := make([]models.Suitability, 0, len(resultSet))
	for _, suitability := range resultSet {
		result = append(result, suitability)
	}
	return result, err
}

func parseFirecloudUser(workspaceUser *admin.User, fcAdminsGroupEmails []string, firecloudProjectOwnersGroupEmails []string) (suitable bool, description string) {
	if workspaceUser.PrimaryEmail == "" {
		return false, "firecloud user doesn't appear to have a primary email? something's amiss, marking as not suitable"
	} else if !workspaceUser.AgreedToTerms {
		return false, fmt.Sprintf("firecloud user hasn't accepted Google Workspace terms (suggesting they've never logged in; they'll need to wait %d minutes after first login for Sherlock to pick it up)",
			config.Config.MustInt("auth.updateIntervalMinutes"))
	} else if !workspaceUser.IsEnrolledIn2Sv {
		return false, "firecloud user hasn't enrolled in two-factor authentication"
	} else if workspaceUser.Suspended {
		return false, fmt.Sprintf("firecloud user is suspended, probably due to inactivity (reach out to #dsp-devops-champions for help; they'll need to wait %d minutes after reactivation for Sherlock to pick it up)",
			config.Config.MustInt("auth.updateIntervalMinutes"))
	} else if workspaceUser.Archived {
		return false, "firecloud user is archived"
	} else if !utils.Contains(fcAdminsGroupEmails, workspaceUser.PrimaryEmail) {
		return false, fmt.Sprintf("firecloud user isn't in fc-admins group (reach out to #dsp-devops-champions for help; they'll need to wait %d minutes after being added for Sherlock to pick it up)",
			config.Config.MustInt("auth.updateIntervalMinutes"))
	} else if !utils.Contains(firecloudProjectOwnersGroupEmails, workspaceUser.PrimaryEmail) {
		return false, fmt.Sprintf("firecloud user isn't in firecloud-project-owners group (reach out to #dsp-devops-champions for help; they'll need to wait %d minutes after being added for Sherlock to pick it up)",
			config.Config.MustInt("auth.updateIntervalMinutes"))
	} else {
		return true, "firecloud user is suitable"
	}
}
