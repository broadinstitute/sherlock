package models

import (
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
	RequiredRole        *Role
	RequiredRoleID      *uint
	HelmfileRef         *string
}

func (c *Cluster) GetCiIdentifier() CiIdentifier {
	if c.CiIdentifier != nil {
		return *c.CiIdentifier
	} else {
		return CiIdentifier{ResourceType: "cluster", ResourceID: c.ID}
	}
}

func (c *Cluster) errorIfForbidden(tx *gorm.DB) error {
	user, err := GetCurrentUserForDB(tx)
	if err != nil {
		return err
	}
	if err = user.ErrIfNotActiveInRole(tx, c.RequiredRoleID); err != nil {
		return err
	}
	return nil
}

// BeforeCreate checks permissions
func (c *Cluster) BeforeCreate(tx *gorm.DB) error {
	return c.errorIfForbidden(tx)
}

// BeforeUpdate checks permissions
func (c *Cluster) BeforeUpdate(tx *gorm.DB) error {
	return c.errorIfForbidden(tx)
}

// AfterUpdate checks permissions. This is necessary because mutations can change permissions.
func (c *Cluster) AfterUpdate(tx *gorm.DB) error {
	return c.errorIfForbidden(tx)
}

// BeforeDelete checks permissions
func (c *Cluster) BeforeDelete(tx *gorm.DB) error {
	return c.errorIfForbidden(tx)
}
