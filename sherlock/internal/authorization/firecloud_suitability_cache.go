package authorization

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/utils"
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
		return fmt.Errorf("failed to authenticate to Google Workspace: %v", err)
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
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	log.Debug().Msgf("AUTH | firecloud suitability cache updated, now contains %d total (primary+recovery) emails", len(newCache))
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
