package auth

import (
	"context"
	"github.com/rs/zerolog/log"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"
	"strings"
	"time"
)

const suitabilityGroup = "fc-admins@firecloud.org"

var cachedSuitableUsers map[string]*suitabilityInfo
var lastCacheTime time.Time

func CacheSuitableUsers(ctx context.Context) error {
	adminService, err := admin.NewService(ctx, option.WithScopes(admin.AdminDirectoryUserReadonlyScope, admin.AdminDirectoryGroupMemberReadonlyScope))
	if err != nil {
		return err
	}
	suitableUsers := map[string]*suitabilityInfo{}
	err = adminService.Members.List(suitabilityGroup).Pages(ctx, func(members *admin.Members) error {
		for _, member := range members.Members {
			if member.Type == "MEMBER" {
				user, err := adminService.Users.Get(member.Email).Do()
				if err != nil {
					return err
				}
				suitableUsers[user.PrimaryEmail] = &suitabilityInfo{
					acceptedWorkspaceTos: user.AgreedToTerms,
					enrolledIn2fa:        user.IsEnrolledIn2Sv,
					suspended:            user.Suspended,
					archived:             user.Archived,
					suspensionReason:     user.SuspensionReason,
				}
			}
		}
		return nil
	})
	cachedSuitableUsers = suitableUsers
	lastCacheTime = time.Now()
	return nil
}

func KeepCacheUpdated(ctx context.Context, interval time.Duration) {
	for {
		time.Sleep(interval)
		err := CacheSuitableUsers(ctx)
		if err != nil {
			log.Warn().Err(err).Msgf("failed to update suitability cache, suitability cache now %s stale", time.Now().Sub(lastCacheTime).String())
		}
	}
}

func getUserSuitabilityInfo(email string) *suitabilityInfo {
	if strings.HasSuffix(email, "@broadinstitute.org") {
		email = strings.TrimSuffix(email, "@broadinstitute.org") + "@firecloud.org"
	}
	return cachedSuitableUsers[email]
}
