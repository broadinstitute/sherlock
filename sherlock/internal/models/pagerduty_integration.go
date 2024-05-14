package models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"gorm.io/gorm"
)

type PagerdutyIntegration struct {
	gorm.Model
	PagerdutyID string
	Name        *string
	Key         *string
	Type        *string
}

func (p *PagerdutyIntegration) errorIfForbidden(tx *gorm.DB) error {
	if user, err := GetCurrentUserForDB(tx); err != nil {
		return err
	} else if err = user.ErrIfNotSuitable(); err != nil {
		return fmt.Errorf("(%s) suitability required: %w", errors.Forbidden, err)
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

// BeforeDelete checks permissions and that no chart releases or environments reference this integration
func (p *PagerdutyIntegration) BeforeDelete(tx *gorm.DB) error {
	var chartReleases []ChartRelease
	if err := tx.Where(&ChartRelease{PagerdutyIntegrationID: &p.ID}).Select("name").Find(&chartReleases).Error; err != nil {
		return fmt.Errorf("wasn't able to check for chart releases that use this integration: %w", err)
	} else if len(chartReleases) > 0 {
		return fmt.Errorf("(%s) the following chart release uses this integration: %v", errors.BadRequest, utils.Map(chartReleases, func(cr ChartRelease) string {
			return cr.Name
		}))
	}
	var environments []Environment
	if err := tx.Where(&Environment{PagerdutyIntegrationID: &p.ID}).Select("name").Find(&environments).Error; err != nil {
		return fmt.Errorf("wasn't able to check for environments that use this integration: %w", err)
	} else if len(environments) > 0 {
		return fmt.Errorf("(%s) the following environment uses this integration: %v", errors.BadRequest, utils.Map(environments, func(e Environment) string {
			return e.Name
		}))
	}
	return p.errorIfForbidden(tx)
}
