package models

import (
	"gorm.io/gorm"
	"time"
)

type RoleAssignmentFields struct {
	Suspended *bool
	ExpiresAt *time.Time
}

type RoleAssignment struct {
	Role   *Role
	RoleID uint `gorm:"primaryKey"`
	User   *User
	UserID uint `gorm:"primaryKey"`

	RoleAssignmentFields
}

type RoleAssignmentOperation struct {
	gorm.Model

	Role     *Role
	RoleID   uint
	User     *User
	UserID   uint
	Author   *User
	AuthorID uint

	// Operation must be one of "create", "update", or "delete".
	Operation string
	// From contains the state of the RoleAssignment before the Operation (required for "create" and "update" operations).
	From RoleAssignmentFields `gorm:"embedded;embeddedPrefix:from_"`
	// To contains the state of the RoleAssignment after the Operation (required for "delete" and "update" operations).
	To RoleAssignmentFields `gorm:"embedded;embeddedPrefix:to_"`
}
