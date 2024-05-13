package models

import (
	"fmt"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type RoleFields struct {
	Name *string

	// SuspendNonSuitableUsers instructs that any RoleAssignment between this Role and a User who is not suitable
	// should be marked as suspended (and any of this Role's grants should be suspended accordingly).
	SuspendNonSuitableUsers *bool

	// CanBeGlassBrokenByRole indicates that any User with a RoleAssignment to the given Role can temporarily give
	// themselves a RoleAssignment to this Role, which will be removed after DefaultGlassBreakDuration. Such a
	// RoleAssignment can last longer if the User either re-glass-breaks (refreshing the duration) or if a
	// super-admin modifies the RoleAssignment's expiry directly.
	CanBeGlassBrokenByRole   *Role
	CanBeGlassBrokenByRoleID *uint
	// DefaultGlassBreakDuration is represented by a nullable int64, as this is the underlying type for a
	// time.Duration. The raw value is in nanoseconds.
	DefaultGlassBreakDuration *int64

	// GrantsSherlockSuperAdmin indicates that any User with a RoleAssignment to this Role gets extra API access
	// to Sherlock itself. This super-admin access confers full control over Role and RoleAssignment records.
	GrantsSherlockSuperAdmin *bool

	// GrantsDevFirecloudGroup, when not null, indicates that a User with an unsuspended RoleAssignment to this
	// Role should have their Firecloud account (if they have one) added to this group.
	GrantsDevFirecloudGroup *string
	// GrantsDevAzureGroup, when not null, indicates that a User with an unsuspended RoleAssignment to this Role
	// should have their Azure account (if they have one) added to this group.
	GrantsDevAzureGroup *string
}

type Role struct {
	gorm.Model

	// Assignments lists User records who have this Role. A RoleAssignment can potentially be suspended,
	// which indicates that the User should not presently have any access commensurate with the Role.
	//
	// The behavior of a suspended RoleAssignment is dependent on the grant. For a group, it might mean
	// that the User should be removed from that group as if they weren't a member of the Role at all.
	// For an account, it might mean that the User should still have an account but that it should be
	// suspended so they can't access it.
	Assignments []*RoleAssignment

	RoleFields
}

type RoleOperation struct {
	gorm.Model

	Role     *Role
	RoleID   uint
	Author   *User
	AuthorID uint

	// Operation must be one of "create", "update", or "delete".
	Operation string
	// From contains the state of the Role before the Operation (required for "create" and "update" operations).
	From RoleFields `gorm:"embedded;embeddedPrefix:from_"`
	// To contains the state of the Role after the Operation (required for "delete" and "update" operations).
	To RoleFields `gorm:"embedded;embeddedPrefix:to_"`
}

// ReadRoleScope should be used in place of `db.Preload(clause.Associations)` for reading Role records, as it
// properly loads the User records opposite the many-to-many RoleAssignment relationship.
//
//nolint:unused
//goland:noinspection GoUnusedExportedFunction,GoUnnecessarilyExportedIdentifiers
func ReadRoleScope(db *gorm.DB) *gorm.DB {
	return db.
		Preload("Assignments").
		Preload("Assignments.User").
		Preload("CanBeGlassBrokenByRole")
}

func (r *Role) BeforeCreate(tx *gorm.DB) error {
	if user, err := GetCurrentUserForDB(tx); err != nil {
		return err
	} else if err = user.ErrIfNotSuperAdmin(); err != nil {
		return err
	}
	return nil
}

func (r *Role) AfterCreate(tx *gorm.DB) error {
	if user, err := GetCurrentUserForDB(tx); err != nil {
		return err
	} else if err = tx.Create(&RoleOperation{
		RoleID:    r.ID,
		AuthorID:  user.ID,
		Operation: "create",
		To:        r.RoleFields,
	}).Error; err != nil {
		return fmt.Errorf("failed to create RoleOperation: %w", err)
	}
	return nil
}

func (r *Role) BeforeUpdate(tx *gorm.DB) error {
	var current Role
	// We want to accurately represent the full state of the new fields, so we do a dance with copying the current
	// fields into this value and then copying the new ones over top, ignoring zero values. This approximates how
	// the database will be updated without us needing to accumulate the true before and after state somehow.
	var newFields RoleFields
	if user, err := GetCurrentUserForDB(tx); err != nil {
		return err
	} else if err = user.ErrIfNotSuperAdmin(); err != nil {
		return err
	} else if err = tx.First(&current, r.ID).Error; err != nil {
		return fmt.Errorf("failed to find current Role: %w", err)
	} else if err = copier.CopyWithOption(&newFields, current.RoleFields, copier.Option{IgnoreEmpty: true, DeepCopy: true}); err != nil {
		return fmt.Errorf("failed to make copy of current RoleFields: %w", err)
	} else if err = copier.CopyWithOption(&newFields, r.RoleFields, copier.Option{IgnoreEmpty: true, DeepCopy: true}); err != nil {
		return fmt.Errorf("failed to copy new RoleFields over current RoleFields: %w", err)
	} else if err = tx.Create(&RoleOperation{
		RoleID:    r.ID,
		AuthorID:  user.ID,
		Operation: "update",
		From:      current.RoleFields,
		To:        newFields,
	}).Error; err != nil {
		return fmt.Errorf("failed to create RoleOperation: %w", err)
	}
	return nil
}

func (r *Role) BeforeDelete(tx *gorm.DB) error {
	var current Role
	if user, err := GetCurrentUserForDB(tx); err != nil {
		return err
	} else if err = user.ErrIfNotSuperAdmin(); err != nil {
		return err
	} else if err = tx.First(&current, r.ID).Error; err != nil {
		return fmt.Errorf("failed to find current Role: %w", err)
	} else if err = tx.Create(&RoleOperation{
		RoleID:    r.ID,
		AuthorID:  user.ID,
		Operation: "delete",
		From:      current.RoleFields,
	}).Error; err != nil {
		return fmt.Errorf("failed to create RoleOperation: %w", err)
	}
	return nil
}