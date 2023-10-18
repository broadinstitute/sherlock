package models

import "gorm.io/gorm"

type PagerdutyIntegration struct {
	gorm.Model
	PagerdutyID string
	Name        *string
	Key         *string
	Type        *string
}

func (p PagerdutyIntegration) TableName() string {
	return "pagerduty_integrations"
}
