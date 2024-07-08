package models

import (
	"database/sql"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/models/advisory_locks"
	"github.com/jinzhu/copier"
	"github.com/sanity-io/litter"
	"gorm.io/gorm"
	"time"
)

type RoleFields struct {
	Name *string

	// SuspendNonSuitableUsers instructs that any RoleAssignment between this Role and a User who is not suitable
	// should be marked as suspended (and any of this Role's grants should be suspended accordingly).
	//
	// When this field is true, RoleAssignmentFields.Suspended becomes effectively a computed field that can't
	// be set directly from the API.
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
	// Role should have their dev Firecloud account (if they have one) added to this group.
	GrantsDevFirecloudGroup *string
	// GrantsQaFirecloudGroup, when not null, indicates that a User with an unsuspended RoleAssignment to this
	// Role should have their qa Firecloud account (if they have one) added to this group.
	GrantsQaFirecloudGroup *string
	// GrantsProdFirecloudGroup, when not null, indicates that a User with an unsuspended RoleAssignment to this
	// Role should have their prod Firecloud account (if they have one) added to this group.
	GrantsProdFirecloudGroup *string

	// GrantsDevAzureGroup, when not null, indicates that a User with an unsuspended RoleAssignment to this Role
	// should have their Azure account (if they have one) added to this group.
	GrantsDevAzureGroup *string
	// GrantsProdAzureGroup, when not null, indicates that a User with an unsuspended RoleAssignment to this Role
	// should have their Azure account (if they have one) added to this group.
	GrantsProdAzureGroup *string
}

type Role struct {
	gorm.Model

	// PropagatedAt stores the last time that this Role's grants were propagated to cloud providers. See
	// the role_propagation package for more information.
	PropagatedAt sql.NullTime

	// Assignments lists User records who have this Role. A RoleAssignment can potentially be suspended,
	// which indicates that the User should not presently have any access commensurate with the Role.
	//
	// The behavior of a suspended RoleAssignment is dependent on the grant. For a group, it might mean
	// that the User should be removed from that group as if they weren't a member of the Role at all.
	// For an account, it might mean that the User should still have an account but that it should be
	// suspended so they can't access it.
	Assignments []*RoleAssignment

	RoleFields

	// previousFields is an unexported field ignored by Gorm. It exists so that the BeforeUpdate hook can
	// copy and store the current state so the AfterUpdate hook can correctly journal the change into a
	// RoleOperation record.
	//
	// (The struct tag to have Gorm ignore it is theoretically unnecessary because it's unexported, but
	// it's included for clarity.)
	previousFields RoleFields `gorm:"-:all"`
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

func (r *Role) AssignmentsMap() map[uint]RoleAssignment {
	roleAssignments := make(map[uint]RoleAssignment)
	for _, ra := range r.Assignments {
		if ra != nil && ra.UserID != 0 {
			roleAssignments[ra.UserID] = *ra
		}
	}
	return roleAssignments
}

// WaitPropagationLock blocks until a propagation lock can be acquired for the Role. This function
// is only safe to call from a transaction. The lock will be released at the end of the transaction.
//
// See https://www.postgresql.org/docs/current/functions-admin.html#FUNCTIONS-ADVISORY-LOCKS
func (r *Role) WaitPropagationLock(tx *gorm.DB) error {
	if err := tx.Exec("SELECT pg_advisory_xact_lock(?, ?)", advisory_locks.ROLE_PROPAGATION, r.ID).Error; err != nil {
		return fmt.Errorf("failed to lock Role %d for propagation: %w", r.ID, err)
	}
	return nil
}

// TryPropagationLock attempts to acquire a propagation lock for the Role. It returns a boolean for
// whether the lock was obtained; it does not block. This function is only safe to call from a
// transaction. The lock will be released at the end of the transaction.
//
// See https://www.postgresql.org/docs/current/functions-admin.html#FUNCTIONS-ADVISORY-LOCKS
func (r *Role) TryPropagationLock(tx *gorm.DB) (bool, error) {
	var locked bool
	if err := tx.Raw("SELECT pg_try_advisory_xact_lock(?, ?)", advisory_locks.ROLE_PROPAGATION, r.ID).Scan(&locked).Error; err != nil {
		return false, fmt.Errorf("failed to try lock Role %d for propagation: %w", r.ID, err)
	}
	return locked, nil
}

// UpdatePropagatedAt sets the Role's PropagatedAt field to the current time without triggering any
// hooks or other Gorm behavior (like setting the gorm.Model UpdatedAt field) since we're not
// semantically making a change to the Role.
func (r *Role) UpdatePropagatedAt(tx *gorm.DB) error {
	if err := tx.Model(&r).UpdateColumns(&Role{PropagatedAt: sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}}).Error; err != nil {
		return err
	}
	return nil
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
	} else {
		slack.SendPermissionChangeNotification(tx.Statement.Context, user.SlackReference(true), slack.PermissionChangeNotificationInputs{
			Summary: fmt.Sprintf("created Role %s", *r.Name),
			Results: []string{
				"Fields: " + slack.EscapeText(litter.Sdump(r.RoleFields)),
			},
		})
	}
	return nil
}

func (r *Role) BeforeUpdate(tx *gorm.DB) error {
	if user, err := GetCurrentUserForDB(tx); err != nil {
		return err
	} else if err = user.ErrIfNotSuperAdmin(); err != nil {
		return err
	} else if err = copier.CopyWithOption(&r.previousFields, &r.RoleFields, copier.Option{DeepCopy: true}); err != nil {
		return fmt.Errorf("failed to copy RoleFields: %w", err)
	}
	return nil
}

func (r *Role) AfterUpdate(tx *gorm.DB) error {
	if user, err := GetCurrentUserForDB(tx); err != nil {
		return err
	} else if err = user.ErrIfNotSuperAdmin(); err != nil {
		return err
	} else if err = tx.Create(&RoleOperation{
		RoleID:    r.ID,
		AuthorID:  user.ID,
		Operation: "update",
		From:      r.previousFields,
		To:        r.RoleFields,
	}).Error; err != nil {
		return fmt.Errorf("failed to create RoleOperation: %w", err)
	} else {
		slack.SendPermissionChangeNotification(tx.Statement.Context, user.SlackReference(true), slack.PermissionChangeNotificationInputs{
			Summary: fmt.Sprintf("edited Role %s", *r.Name),
			Results: []string{
				"Old fields: " + slack.EscapeText(litter.Sdump(r.previousFields)),
				"New fields: " + slack.EscapeText(litter.Sdump(r.RoleFields)),
			},
		})
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
	} else {
		slack.SendPermissionChangeNotification(tx.Statement.Context, user.SlackReference(true), slack.PermissionChangeNotificationInputs{
			Summary: fmt.Sprintf("deleted Role %s", *r.Name),
			Results: []string{
				"Fields: " + slack.EscapeText(litter.Sdump(r.RoleFields)),
			},
		})
	}
	return nil
}
