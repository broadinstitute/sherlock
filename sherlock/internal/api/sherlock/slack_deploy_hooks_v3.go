package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"gorm.io/gorm"
)

type SlackDeployHookV3 struct {
	CommonFields
	DeployHookTriggerConfigV3
	SlackDeployHookFields
}

type SlackDeployHookFields struct {
	SlackChannel  *string `json:"slackChannel,omitempty" form:"slackChannel"`
	MentionPeople *bool   `json:"mentionPeople,omitempty" form:"mentionPeople"`
}

func (s SlackDeployHookV3) toModel(db *gorm.DB) (models.SlackDeployHook, error) {
	ret := models.SlackDeployHook{
		Model:         s.toGormModel(),
		SlackChannel:  s.SlackChannel,
		MentionPeople: s.MentionPeople,
	}
	var err error
	ret.Trigger, err = s.DeployHookTriggerConfigV3.toModel(db)
	return ret, err
}

func slackDeployHookFromModel(model models.SlackDeployHook) SlackDeployHookV3 {
	return SlackDeployHookV3{
		CommonFields:              commonFieldsFromGormModel(model.Model),
		DeployHookTriggerConfigV3: deployHookTriggerConfigFromModel(model.Trigger),
		SlackDeployHookFields: SlackDeployHookFields{
			SlackChannel:  model.SlackChannel,
			MentionPeople: model.MentionPeople,
		},
	}
}
