package firecloud_account_manager

import (
	"context"
	"fmt"
	"time"

	"github.com/broadinstitute/sherlock/sherlock/internal/clients/google_workspace"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/knadh/koanf"
	"gorm.io/gorm"
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

			var err error
			if manager.ImpersonateAccount != "" {
				manager.NeverAffectEmails = append(manager.NeverAffectEmails, manager.ImpersonateAccount)
				manager.workspaceClient, err = google_workspace.InitializeRealWorkspaceClient(ctx, manager.ImpersonateAccount)
			} else {
				manager.workspaceClient, err = google_workspace.InitializeRealWorkspaceClient(ctx)
			}
			if err != nil {
				return fmt.Errorf("error initializing firecloudAccountManager[%d] (%s): %w", index, manager.Domain, err)
			}

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
