package models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
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

func (d *DeployHookTriggerConfig) TableName() string {
	return "v2_deploy_hook_trigger_configs"
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
	user, err := GetCurrentUserForDB(tx)
	if err != nil {
		return err
	}
	var requiresSuitability bool
	var clusterIDToCheck, environmentIDToCheck *uint
	if d.OnChartReleaseID != nil {
		var chartRelease ChartRelease
		if err = tx.First(&chartRelease, *d.OnChartReleaseID).Error; err != nil {
			return err
		}
		clusterIDToCheck = chartRelease.ClusterID
		environmentIDToCheck = chartRelease.EnvironmentID
	} else if d.OnEnvironmentID != nil {
		environmentIDToCheck = d.OnEnvironmentID
	}
	if clusterIDToCheck != nil {
		var cluster Cluster
		if err = tx.First(&cluster, *clusterIDToCheck).Error; err != nil {
			return err
		}
		if cluster.RequiresSuitability != nil {
			requiresSuitability = requiresSuitability || *cluster.RequiresSuitability
		}
	}
	if environmentIDToCheck != nil {
		var environment Environment
		if err = tx.First(&environment, *environmentIDToCheck).Error; err != nil {
			return err
		}
		if environment.RequiresSuitability != nil {
			requiresSuitability = requiresSuitability || *environment.RequiresSuitability
		}
	}
	if requiresSuitability {
		if err = user.Suitability().SuitableOrError(); err != nil {
			return fmt.Errorf("(%s) suitability required: %w", errors.Forbidden, err)
		}
	}
	return nil
}

func (d *DeployHookTriggerConfig) AfterSave(tx *gorm.DB) error {
	return d.ErrorIfForbidden(tx)
}

func (d *DeployHookTriggerConfig) AfterDelete(tx *gorm.DB) error {
	return d.ErrorIfForbidden(tx)
}
