package models

import (
	"fmt"

	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ServiceAlert struct {
	gorm.Model
	Title           *string
	AlertMessage    *string
	Link            *string
	Severity        *string
	OnEnvironmentID *uint
	OnEnvironment   *Environment
	CreatedBy       *string
	UpdatedBy       *string
	DeletedBy       *string
	Uuid            *uuid.UUID
}

func (s *ServiceAlert) errorIfForbidden(tx *gorm.DB) error {
	if s.OnEnvironmentID == nil {
		return fmt.Errorf("(%s) service alert wasn't properly loaded, wasn't able to check permissions", errors.InternalServerError)
	}
	if s.OnEnvironmentID != nil {
		var environment Environment
		if err := tx.Take(&environment, *s.OnEnvironmentID).Error; err != nil {
			return fmt.Errorf("(%s) failed to read service alert's environment to evaluate permissions: %w", errors.InternalServerError, err)
		}
		if err := environment.errorIfForbidden(tx); err != nil {
			return fmt.Errorf("forbidden based on service alert's environment: %w", err)
		}
	}
	return nil
}

// BeforeCreate checks permissions
func (s *ServiceAlert) BeforeCreate(tx *gorm.DB) error {
	if err := s.errorIfForbidden(tx); err != nil {
		return err
	}
	return nil
}

// BeforeUpdate checks permissions
func (s *ServiceAlert) BeforeUpdate(tx *gorm.DB) error {
	return s.errorIfForbidden(tx)
}

// BeforeDelete checks permissions and propagates deletions
func (s *ServiceAlert) BeforeDelete(tx *gorm.DB) error {
	if err := s.errorIfForbidden(tx); err != nil {
		return err
	}
	return nil
}
