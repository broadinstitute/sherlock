package models

import "gorm.io/gorm"

type PagerdutyIntegration struct {
	gorm.Model
	PagerdutyID string
	Name        *string
	Key         *string
	Type        *string
}

func (p *PagerdutyIntegration) TableName() string {
	return "pagerduty_integrations"
}

func (p *PagerdutyIntegration) errorIfForbidden(tx *gorm.DB) error {
	if user, err := GetCurrentUserForDB(tx); err != nil {
		return err
	} else if err = user.Suitability().SuitableOrError(); err != nil {
		return err
	} else {
		return nil
	}
}

// BeforeCreate checks permissions
func (p *PagerdutyIntegration) BeforeCreate(tx *gorm.DB) error {
	return p.errorIfForbidden(tx)
}

// BeforeUpdate checks permissions
func (p *PagerdutyIntegration) BeforeUpdate(tx *gorm.DB) error {
	return p.errorIfForbidden(tx)
}

// BeforeDelete checks permissions
func (p *PagerdutyIntegration) BeforeDelete(tx *gorm.DB) error {
	return p.errorIfForbidden(tx)
}
