package auth

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth/auth_models"
	"github.com/broadinstitute/sherlock/internal/config"
	"github.com/rs/zerolog/log"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"
	"time"
)

// cachedFirecloudAccounts associates firecloud.org email addresses to FirecloudAccount info
var cachedFirecloudAccounts map[string]*auth_models.FirecloudAccount
var lastCacheTime time.Time

func CacheFirecloudAccounts(ctx context.Context) error {
	adminService, err := admin.NewService(ctx, option.WithScopes(admin.AdminDirectoryUserReadonlyScope, admin.AdminDirectoryGroupMemberReadonlyScope))
	if err != nil {
		return fmt.Errorf("failed to authenticate to Google Workspace: %v", err)
	}

	newCache := make(map[string]*auth_models.FirecloudAccount)
	err = adminService.Users.List().Domain(config.Config.MustString("auth.firecloud.domain")).Pages(ctx, func(workspaceUsers *admin.Users) error {
		if workspaceUsers == nil {
			log.Warn().Msg("CacheFirecloudAccounts got a nil user page from Google?")
		} else {
			for _, workspaceUser := range workspaceUsers.Users {
				if workspaceUser == nil {
					log.Warn().Msg("CacheFirecloudAccounts got a nil user from Google?")
				} else {
					newCache[workspaceUser.PrimaryEmail] = &auth_models.FirecloudAccount{
						Email:               workspaceUser.PrimaryEmail,
						AcceptedGoogleTerms: workspaceUser.AgreedToTerms,
						EnrolledIn2fa:       workspaceUser.IsEnrolledIn2Sv,
						Suspended:           workspaceUser.Suspended,
						Archived:            workspaceUser.Archived,
						SuspensionReason:    workspaceUser.SuspensionReason,
						Groups:              &auth_models.FirecloudGroupMembership{},
					}
				}
			}
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to update users from Google Workspace: %v", err)
	}

	err = adminService.Members.List(config.Config.MustString("auth.firecloud.groups.fcAdmins")).Pages(ctx, func(members *admin.Members) error {
		if members == nil {
			log.Warn().Msgf("CacheFirecloudAccounts got a nil %s member page from Google?", config.Config.MustString("auth.firecloud.groups.fcAdmins"))
		} else {
			for _, member := range members.Members {
				if member == nil {
					log.Warn().Msgf("CacheFirecloudAccounts got a nil %s member from Google?", config.Config.MustString("auth.firecloud.groups.fcAdmins"))
				} else if fcAccount, exists := newCache[member.Email]; exists {
					fcAccount.Groups.FcAdmins = true
				}
			}
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to update %s members from Google Workspace: %v", config.Config.MustString("auth.firecloud.groups.fcAdmins"), err)
	}

	err = adminService.Members.List(config.Config.MustString("auth.firecloud.groups.firecloudProjectOwners")).Pages(ctx, func(members *admin.Members) error {
		if members == nil {
			log.Warn().Msgf("CacheFirecloudAccounts got a nil %s member page from Google?", config.Config.MustString("auth.firecloud.groups.firecloudProjectOwners"))
		} else {
			for _, member := range members.Members {
				if member == nil {
					log.Warn().Msgf("CacheFirecloudAccounts got a nil %s member from Google?", config.Config.MustString("auth.firecloud.groups.firecloudProjectOwners"))
				} else if fcAccount, exists := newCache[member.Email]; exists {
					fcAccount.Groups.FirecloudProjectOwners = true
				}
			}
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to update %s members from Google Workspace: %v", config.Config.MustString("auth.firecloud.groups.firecloudProjectOwners"), err)
	}

	log.Debug().Msgf("AUTH | firecloud account cache updated, now contains %d accounts", len(newCache))
	cachedFirecloudAccounts = newCache
	lastCacheTime = time.Now()
	return nil
}

func KeepFirecloudCacheUpdated(ctx context.Context, interval time.Duration) {
	for {
		time.Sleep(interval)
		if err := CacheFirecloudAccounts(ctx); err != nil {
			log.Warn().Err(err).Msgf("failed to update suitability cache, now %s stale", time.Since(lastCacheTime).String())
		}
	}
}
