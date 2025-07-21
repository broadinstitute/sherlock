package models

import (
	"fmt"

	"gorm.io/gorm"
)

type DeployHookTriggerConfig struct {
	gorm.Model
	HookID           uint
	HookType         string
	OnEnvironmentID  *uint
	OnEnvironment    *Environment
	OnChartReleaseID *uint
	OnChartRelease   *ChartRelease
	OnFailure        *bool
	OnSuccess        *bool
}

// ErrorIfForbidden looks at the environment or chart release's environment/cluster to determine
// if this entry requires suitability, and then errors if the caller lacks it. This is a little
// bit inefficient because we ignore potentially already-loaded associations on
// DeployHookTriggerConfig and query them anew, but that's protection against them being loaded
// improperly and that error making its way into the authorization system.
//
// It's important that this function be called only from Gorm's "After" hooks, because that's
// the only way we're guaranteed to have the other data we need in the DeployHookTriggerConfig.
// If those hooks return an error the whole transaction will be rolled back, so it'll still do
// its job.
func (d *DeployHookTriggerConfig) ErrorIfForbidden(tx *gorm.DB) error {
	if d.OnChartReleaseID != nil {
		var chartRelease ChartRelease
		if err := tx.Take(&chartRelease, *d.OnChartReleaseID).Error; err != nil {
			return fmt.Errorf("error querying suitability on chart release: %w", err)
		}
		return chartRelease.errorIfForbidden(tx)
	} else if d.OnEnvironmentID != nil {
		var environment Environment
		if err := tx.Take(&environment, *d.OnEnvironmentID).Error; err != nil {
			return fmt.Errorf("error querying suitability on environment: %w", err)
		}
		return environment.errorIfForbidden(tx)
	}
	return nil
}

func (d *DeployHookTriggerConfig) AfterSave(tx *gorm.DB) error {
	return d.ErrorIfForbidden(tx)
}

func (d *DeployHookTriggerConfig) AfterDelete(tx *gorm.DB) error {
	return d.ErrorIfForbidden(tx)
}

func (d *DeployHookTriggerConfig) SlackBeehiveLink() string {
	if d.OnEnvironment != nil {
		return d.OnEnvironment.SlackBeehiveLink()
	} else if d.OnChartRelease != nil {
		return d.OnChartRelease.SlackBeehiveLink()
	} else {
		return fmt.Sprintf("(orphaned deploy hook trigger config %d)", d.ID)
	}
}

func (d *DeployHookTriggerConfig) ArgoCdUrl() (string, bool) {
	if d.OnEnvironment != nil {
		return d.OnEnvironment.ArgoCdUrl()
	} else if d.OnChartRelease != nil {
		return d.OnChartRelease.ArgoCdUrl()
	} else {
		return "", false
	}
}
