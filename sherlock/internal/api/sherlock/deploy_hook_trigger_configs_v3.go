package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"gorm.io/gorm"
)

// DeployHookTriggerConfigV3 isn't an exported field because this data type isn't intended
// to be directly exposed in the API. Instead, SlackDeployHookV3 and GithubActionsDeployHookV3
// will embed this struct.
type DeployHookTriggerConfigV3 struct {
	OnEnvironment  *string `json:"onEnvironment,omitempty" form:"onEnvironment"`
	OnChartRelease *string `json:"onChartRelease,omitempty" form:"onChartRelease"`
	deployHookTriggerConfigV3EditableFields
}

type deployHookTriggerConfigV3EditableFields struct {
	OnFailure *bool `json:"onFailure,omitempty" form:"onFailure"`
	OnSuccess *bool `json:"onSuccess,omitempty" form:"onSuccess"`
}

func (d DeployHookTriggerConfigV3) toModel(db *gorm.DB) (models.DeployHookTriggerConfig, error) {
	ret := models.DeployHookTriggerConfig{
		OnFailure: d.OnFailure,
		OnSuccess: d.OnSuccess,
	}
	if d.OnEnvironment != nil {
		if environmentID, err := v2models.InternalEnvironmentStore.ResolveSelector(db, *d.OnEnvironment); err != nil {
			return models.DeployHookTriggerConfig{}, err
		} else {
			ret.OnEnvironmentID = &environmentID
		}
	}
	if d.OnChartRelease != nil {
		if chartReleaseID, err := v2models.InternalChartReleaseStore.ResolveSelector(db, *d.OnChartRelease); err != nil {
			return models.DeployHookTriggerConfig{}, err
		} else {
			ret.OnChartReleaseID = &chartReleaseID
		}
	}
	return ret, nil
}

func deployHookTriggerConfigFromModel(model models.DeployHookTriggerConfig) DeployHookTriggerConfigV3 {
	var onEnvironment *string
	if model.OnEnvironment != nil && model.OnEnvironment.Name != "" {
		onEnvironment = &model.OnEnvironment.Name
	} else if model.OnEnvironmentID != nil {
		s := utils.UintToString(*model.OnEnvironmentID)
		onEnvironment = &s
	}
	var onChartRelease *string
	if model.OnChartRelease != nil && model.OnChartRelease.Name != "" {
		onChartRelease = &model.OnChartRelease.Name
	} else if model.OnChartReleaseID != nil {
		s := utils.UintToString(*model.OnChartReleaseID)
		onChartRelease = &s
	}
	return DeployHookTriggerConfigV3{
		OnEnvironment:  onEnvironment,
		OnChartRelease: onChartRelease,
		deployHookTriggerConfigV3EditableFields: deployHookTriggerConfigV3EditableFields{
			OnFailure: model.OnFailure,
			OnSuccess: model.OnSuccess,
		},
	}
}
