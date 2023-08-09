package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"gorm.io/gorm"
)

type SlackDeployHookV3 struct {
	commonFields
	deployHookTriggerConfigV3
	slackDeployHookFields
}

type slackDeployHookFields struct {
	SlackChannel *string `json:"slackChannel,omitempty" form:"slackChannel"`
}

func (s SlackDeployHookV3) toModel(db *gorm.DB) (models.SlackDeployHook, error) {
	ret := models.SlackDeployHook{
		Model:        s.commonFields.toGormModel(),
		SlackChannel: s.SlackChannel,
	}
	var err error
	ret.Trigger, err = s.deployHookTriggerConfigV3.toModel(db)
	return ret, err
}

func slackDeployHookFromModel(model models.SlackDeployHook) SlackDeployHookV3 {
	return SlackDeployHookV3{
		commonFields:              commonFieldsFromGormModel(model.Model),
		deployHookTriggerConfigV3: deployHookTriggerConfigFromModel(model.Trigger),
		slackDeployHookFields:     slackDeployHookFields{SlackChannel: model.SlackChannel},
	}
}
