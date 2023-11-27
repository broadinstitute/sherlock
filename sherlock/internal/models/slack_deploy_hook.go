package models

import (
	"gorm.io/gorm"
)

type SlackDeployHook struct {
	gorm.Model
	Trigger       DeployHookTriggerConfig `gorm:"polymorphic:Hook;polymorphicValue:slack"`
	SlackChannel  *string
	MentionPeople *bool
	Beta          *bool
}

func (s *SlackDeployHook) AfterSave(tx *gorm.DB) error {
	return s.Trigger.ErrorIfForbidden(tx)
}

func (s *SlackDeployHook) AfterDelete(tx *gorm.DB) error {
	return s.Trigger.ErrorIfForbidden(tx)
}
