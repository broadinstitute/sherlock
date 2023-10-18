package models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"gorm.io/gorm"
)

type Cluster struct {
	gorm.Model
	CiIdentifier      *CiIdentifier `gorm:"polymorphic:Resource; polymorphicValue:cluster"`
	Name              string
	Provider          string
	GoogleProject     string
	AzureSubscription string
	Location          string
	// Mutable
	Base                *string
	Address             *string
	RequiresSuitability *bool
	HelmfileRef         *string
}

func (c *Cluster) TableName() string {
	return "clusters"
}

func (c *Cluster) GetCiIdentifier() CiIdentifier {
	if c.CiIdentifier != nil {
		return *c.CiIdentifier
	} else {
		return CiIdentifier{ResourceType: "cluster", ResourceID: c.ID}
	}
}

func (c *Cluster) ErrorIfForbidden(tx *gorm.DB) error {
	user, err := GetCurrentUserForDB(tx)
	if err != nil {
		return err
	}
	if c.RequiresSuitability == nil || *c.RequiresSuitability {
		if err = user.Suitability().SuitableOrError(); err != nil {
			return fmt.Errorf("(%s) suitability required: %w", errors.Forbidden, err)
		}
	}
	return nil
}

func (c *Cluster) BeforeUpdate(tx *gorm.DB) error {
	// Updates could potentially set suitability lower, so we check before updates too
	return c.ErrorIfForbidden(tx)
}

func (c *Cluster) AfterSave(tx *gorm.DB) error {
	return c.ErrorIfForbidden(tx)
}

func (c *Cluster) AfterDelete(tx *gorm.DB) error {
	return c.ErrorIfForbidden(tx)
}
