package firecloud_account_manager

import (
	"cloud.google.com/go/auth/credentials/impersonate"
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/knadh/koanf"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"
	"gorm.io/gorm"
	"time"
)

var managers []firecloudAccountManager

// Init will authenticate and validate the firecloudAccountManager configurations (if there are any).
// This should be called before RunManagersHourly. RunManagersHourly should not run if this function
// returns an error.
func Init(ctx context.Context, db *gorm.DB) error {
	rawConfigs := config.Config.Slices("firecloudAccountManager")

	if len(rawConfigs) > 0 {
		for index, rawConfig := range rawConfigs {
			var manager firecloudAccountManager
			if err := rawConfig.UnmarshalWithConf("", &manager, koanf.UnmarshalConf{Tag: "firecloud_account_manager"}); err != nil {
				return fmt.Errorf("error parsing firecloudAccountManager[%d]: %w", index, err)
			}
			manager.indexPlusOneForLocking = index + 1
			manager.dbForLocking = db

			adminServiceOptions := []option.ClientOption{option.WithScopes(admin.AdminDirectoryUserScope)}
			if manager.ImpersonateAccount != "" {
				manager.NeverAffectEmails = append(manager.NeverAffectEmails, manager.ImpersonateAccount)
				credentials, err := impersonate.NewCredentials(&impersonate.CredentialsOptions{
					TargetPrincipal: models.SelfUser.Email,
					Scopes:          []string{admin.AdminDirectoryUserScope, "https://www.googleapis.com/auth/cloud-platform"},
					Subject:         manager.ImpersonateAccount,
				})
				if err != nil {
					return fmt.Errorf("failed to create impersonated credentials for %s for firecloudAccountManager[%d] (%s): %w",
						manager.ImpersonateAccount, index, manager.Domain, err)
				}
				adminServiceOptions = append(adminServiceOptions, option.WithAuthCredentials(credentials))
			}
			adminService, err := admin.NewService(ctx, adminServiceOptions...)
			if err != nil {
				return fmt.Errorf("failed to create admin service: %w", err)
			}
			manager.workspaceClient = &realWorkspaceClient{adminService: adminService}

			if err = manager.validate(); err != nil {
				return fmt.Errorf("error validating firecloudAccountManager[%d]: %w", index, err)
			}

			managers = append(managers, manager)
		}
	}
	return nil
}

// RunManagersHourly runs the suspendAccounts function for each firecloudAccountManager every hour.
// When this function is first called, it won't know the last hour that the managers ran, so it will
// run immediately. It's better for us to run the managers more often than every hour than less often.
func RunManagersHourly(ctx context.Context) {
	var lastRanAtHour *int
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Minute):
			if lastRanAtHour == nil || *lastRanAtHour != time.Now().Hour() {
				currentHour := time.Now().Hour()
				lastRanAtHour = &currentHour
				for _, manager := range managers {
					if manager.Enable {
						_, _ = manager.suspendAccounts(ctx)
					}
				}
			}
		}
	}
}
