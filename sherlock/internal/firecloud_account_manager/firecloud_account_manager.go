package firecloud_account_manager

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type firecloudAccountManager struct {
	// The index of this firecloudAccountManager's configuration in the config file, plus one.
	// This is used for locking, so that two firecloudAccountManagers with the same config
	// don't run at once. We essentially lock the entry in the config file using a PostgreSQL
	// advisory lock.
	indexPlusOneForLocking int
	// A reference to the DB, to use only for locking. This system isn't concerned with any
	// of the actual data in the DB.
	dbForLocking *gorm.DB

	Domain                string        `koanf:"domain"`
	Enable                bool          `koanf:"enable"`
	DryRun                bool          `koanf:"dryRun"`
	OnlyAffectEmails      []string      `koanf:"onlyAffectEmails"`
	NeverAffectEmails     []string      `koanf:"neverAffectEmails"`
	NewAccountGracePeriod time.Duration `koanf:"newAccountGracePeriod"`
	InactivityThreshold   time.Duration `koanf:"inactivityThreshold"`
	ImpersonateAccount    string        `koanf:"impersonateAccount"`

	// The workspace client to use for this firecloudAccountManager. This exists to allow
	// mocking.
	workspaceClient mockableWorkspaceClient
}

func (m *firecloudAccountManager) validate() error {
	if m.indexPlusOneForLocking == 0 {
		return fmt.Errorf("indexPlusOneForLocking must be set -- it should never be zero")
	}
	if m.dbForLocking == nil {
		return fmt.Errorf("dbForLocking must be set")
	}
	if m.Domain == "" {
		return fmt.Errorf("domain must be set")
	}
	if m.NewAccountGracePeriod.Seconds() == 0 {
		return fmt.Errorf("newAccountGracePeriod must be set")
	}
	if m.InactivityThreshold.Seconds() == 0 {
		return fmt.Errorf("inactivityThreshold must be set")
	}
	if m.workspaceClient == nil {
		return fmt.Errorf("workspaceClient must be set")
	}
	return nil
}
