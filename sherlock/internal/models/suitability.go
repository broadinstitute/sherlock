package models

import (
	"time"

	"gorm.io/gorm"
)

// Suitability represents a User's completion of compliance requirements,
// like training or background checks. The term comes from being "suitable"
// for production access.
//
// See Role for the where this is used: a Role can be marked as requiring
// each assigned User to be suitable. A non-suitable User would have their
// RoleAssignment suspended.
//
// Suitability *could* be a field on User, but we split it out to its own
// table for three reasons:
//
//  1. We expect this type to become much more complex in the future.
//     Our friends in InfoSec have indicated they'd like to move more of
//     their manual compliance tracking into Sherlock, to automate checks
//     like "if the last training for X was completed more than Y days ago,
//     the user is not suitable."
//  2. As part of the first point and also Sherlock simply becoming the
//     source of truth for DevOps's permission grants to other engineers,
//     Sherlock needs to understand the suitability of people who haven't
//     yet actually ever connected to the DevOps platform. Tracking this
//     in a separate table means we can have records for people who aren't
//     a User yet.
//  3. The concepts of "user" and "suitability" have different access
//     control. A User record is only editable by itself (changing name,
//     linking GitHub account, etc.), while Suitability is only editable
//     by a Sherlock super-admin. We can represent this in Gorm hooks
//     easily if the model types are different.
//
// We key this type on the User's Email, as that correlates to how humans
// track this information. Sherlock has defenses against email spoofing
// on the User table, when it's able to actually observe indicators of that.
// Entries in this table could theoretically predate the email actually
// existing (so things like Google Subject ID wouldn't exist).
type Suitability struct {
	CreatedAt time.Time
	UpdatedAt time.Time

	Email       *string `gorm:"primaryKey"`
	Suitable    *bool
	Description *string
}

func (s *Suitability) BeforeUpdate(tx *gorm.DB) error {
	if user, err := GetCurrentUserForDB(tx); err != nil {
		return err
	} else {
		return user.ErrIfNotSuperAdmin()
	}
}

func (s *Suitability) BeforeCreate(tx *gorm.DB) error {
	if user, err := GetCurrentUserForDB(tx); err != nil {
		return err
	} else {
		return user.ErrIfNotSuperAdmin()
	}
}

func (s *Suitability) BeforeDelete(tx *gorm.DB) error {
	if user, err := GetCurrentUserForDB(tx); err != nil {
		return err
	} else {
		return user.ErrIfNotSuperAdmin()
	}
}
