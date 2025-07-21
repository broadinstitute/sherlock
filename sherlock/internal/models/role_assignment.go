package models

import (
	"fmt"
	"time"

	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RoleAssignmentFields struct {
	// Suspended represents whether the RoleAssignment is active or not.
	//
	// This field is closely related to RoleFields.SuspendNonSuitableUsers. When that field is true,
	// this field effectively becomes a computed field, managed entirely by Sherlock based on the
	// Suitability of each User. This accomplishes effectively revoking a User's permissions when
	// they become non-suitable.
	//
	// Why have a concept of suspension vs. just deleting the RoleAssignment? There's two reasons:
	// 1. Some role grants, like a Firecloud.org account, actually have a concept of suspension.
	//    When this is available, it's extremely useful: e.g. we can create and delete Firecloud.org
	//    accounts based on presence of a RoleAssignment, and as suitability fluctuates we just
	//    suspend it. If we deleted it, the Google subject ID would change, and people would need
	//    to set their password and two-factor.
	// 2. It allows the super-admins to manage role membership independent of suitability. We don't
	//    need to wait for a new employee to actually become suitable before adding them to a bunch
	//    of roles: the roles will just automatically not grant permissions until the user becomes
	//    suitable.
	Suspended *bool
	ExpiresAt *time.Time
}

func (a *RoleAssignmentFields) String() string {
	suspended := false
	if a.Suspended != nil {
		suspended = *a.Suspended
	}
	expiresAt := "(never)"
	if a.ExpiresAt != nil {
		expiresAt = a.ExpiresAt.Format(time.RFC3339)
	}
	return fmt.Sprintf("{Suspended: %t, ExpiresAt: %s}", suspended, expiresAt)
}

type RoleAssignment struct {
	Role   *Role
	RoleID uint `gorm:"primaryKey"`
	User   *User
	UserID uint `gorm:"primaryKey"`

	RoleAssignmentFields

	// previousFields is an unexported field ignored by Gorm. It exists so that the BeforeUpdate hook can
	// copy and store the current state so the AfterUpdate hook can correctly journal the change into a
	// RoleAssignmentOperation record.
	//
	// (The struct tag to have Gorm ignore it is theoretically unnecessary because it's unexported, but
	// it's included for clarity.)
	previousFields RoleAssignmentFields `gorm:"-:all"`
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

func (ra *RoleAssignment) IsActive() bool {
	return ra.Suspended != nil && !*ra.Suspended && (ra.ExpiresAt == nil || ra.ExpiresAt.After(time.Now()))
}

func (ra *RoleAssignment) ErrIfNotActive() error {
	if ra.Suspended != nil && *ra.Suspended {
		return fmt.Errorf("(%s) role assignment is suspended", errors.Forbidden)
	} else if ra.ExpiresAt != nil && ra.ExpiresAt.Before(time.Now()) {
		return fmt.Errorf("(%s) role assignment has expired", errors.Forbidden)
	} else {
		return nil
	}
}

func (ra *RoleAssignment) Description(db *gorm.DB) string {
	var role Role
	var user User
	var roleName, userName string
	if ra.Role == nil {
		db.Take(&role, ra.RoleID)
	} else {
		role = *ra.Role
	}
	if role.Name != nil {
		roleName = *role.Name
	} else {
		roleName = fmt.Sprintf("Role %s", utils.UintToString(ra.RoleID))
	}
	if ra.User == nil {
		db.Take(&user, ra.UserID)
	} else {
		user = *ra.User
	}
	if user.Email != "" {
		userName = user.SlackReference(true)
	} else {
		userName = fmt.Sprintf("User %s", utils.UintToString(ra.UserID))
	}
	return fmt.Sprintf("%s in %s", userName, roleName)
}

func (ra *RoleAssignment) errorIfForbidden(tx *gorm.DB) error {
	user, err := GetCurrentUserForDB(tx)
	if err != nil {
		return err
	}

	if errIfNotSuperAdmin := user.ErrIfNotSuperAdmin(); errIfNotSuperAdmin == nil {
		// If the super-admin check doesn't return an error, short-circuit and return nil.
		return nil
	}

	// Otherwise, we go through the steps to check if the user is allowed to break-glass
	// We get the current RoleAssignment from the database instead of using the receiver in case it was mutated
	var current RoleAssignment
	if err = tx.Where(&RoleAssignment{RoleID: ra.RoleID, UserID: ra.UserID}).First(&current).Error; err != nil {
		return fmt.Errorf("failed to find current RoleAssignment: %w", err)
	}
	var targetRole Role
	if err = tx.Preload("CanBeGlassBrokenByRole").First(&targetRole, current.RoleID).Error; err != nil {
		return fmt.Errorf("failed to find target role: %w", err)
	}
	if targetRole.CanBeGlassBrokenByRoleID == nil {
		return fmt.Errorf("(%s) role %s (%d) cannot be glass-broken and the caller is not a super-admin who can make non-glass-break assignments", errors.Forbidden, *targetRole.Name, current.RoleID)
	} else if ra.UserID != user.ID {
		return fmt.Errorf("(%s) a caller may only make break-glass assignments for themselves", errors.Forbidden)
	} else if targetRole.DefaultGlassBreakDuration == nil {
		return fmt.Errorf("role %s (%d) is misconfigured: it declares that it can be glass-broken but defines no default duration", *targetRole.Name, current.RoleID)
	} else {
		for _, userAssignment := range user.Assignments {
			// If the user has the role that can break-glass the target role:
			if userAssignment.RoleID == *targetRole.CanBeGlassBrokenByRoleID {
				if userAssignment.Suspended == nil || *userAssignment.Suspended {
					// If suspended, bail
					return fmt.Errorf("(%s) role %s (%d) can only be glass-broken by role %s (%d) and the caller has that role but their assignment is suspended", errors.Forbidden, *targetRole.Name, current.RoleID, *targetRole.CanBeGlassBrokenByRole.Name, *targetRole.CanBeGlassBrokenByRoleID)
				} else if current.ExpiresAt == nil {
					// If not working with an expiry, bail, break-glass requires one
					return fmt.Errorf("(%s) role %s (%d) can be glass-broken by the caller but glass-break assignments require an expiry", errors.Forbidden, *targetRole.Name, current.RoleID)
				} else if current.ExpiresAt.Before(time.Now()) {
					// If the expiry is in the past, bail
					return fmt.Errorf("(%s) role %s (%d) can be glass-broken by the caller but the expiry on the break-glass assignment is in the past", errors.Forbidden, *targetRole.Name, current.RoleID)
				} else if current.ExpiresAt.After(time.Now().Add(time.Duration(*targetRole.DefaultGlassBreakDuration) + time.Second)) {
					// If the expiry is too far in the future, bail
					// (We fudge the timing by a second to account for the fact that the expiry is inclusive)
					return fmt.Errorf("(%s) role %s (%d) can be glass-broken by the caller but the expiry on the break-glass assignment is too far in the future (must be within %s of now)", errors.Forbidden, *targetRole.Name, current.RoleID, time.Duration(*targetRole.DefaultGlassBreakDuration).String())
				} else if current.Suspended != nil && *current.Suspended {
					// If the assigment is suspended, bail, break-glass and suspensions don't mix
					return fmt.Errorf("(%s) role %s (%d) can be glass-broken by the caller but the break-glass assignment is suspended (break-glass and suspensions don't mix)", errors.Forbidden, *targetRole.Name, current.RoleID)
				} else {
					// If we get all the way here, the operation is valid
					return nil
				}
			}
		}
		return fmt.Errorf("(%s) role %s (%d) can only be glass-broken by role %s (%d) and the caller has neither that role nor super-admin", errors.Forbidden, *targetRole.Name, current.RoleID, *targetRole.CanBeGlassBrokenByRole.Name, *targetRole.CanBeGlassBrokenByRoleID)
	}
}

// errorIfUserDeactivated does what it says on the tin, but we don't want to call it often. We
// only want to forbid creation of role assignments with deactivated users, because we want to
// rely on
func (ra *RoleAssignment) errorIfUserDeactivated(tx *gorm.DB) error {
	var user User
	if err := tx.First(&user, ra.UserID).Error; err != nil {
		return fmt.Errorf("failed to find user to determine if it is deactivated: %w", err)
	} else if user.DeactivatedAt != nil {
		return fmt.Errorf("(%s) cannot create role assignment targeting deactivated user", errors.Forbidden)
	}
	return nil
}

func (ra *RoleAssignment) AfterCreate(tx *gorm.DB) error {
	if user, err := GetCurrentUserForDB(tx); err != nil {
		return err
	} else if err = ra.errorIfForbidden(tx); err != nil {
		return err
	} else if err = ra.errorIfUserDeactivated(tx); err != nil {
		return err
	} else if err = tx.Omit(clause.Associations).Create(&RoleAssignmentOperation{
		RoleID:    ra.RoleID,
		UserID:    ra.UserID,
		AuthorID:  user.ID,
		Operation: "create",
		To:        ra.RoleAssignmentFields,
	}).Error; err != nil {
		return fmt.Errorf("failed to create RoleAssignmentOperation: %w", err)
	} else {
		slack.SendPermissionChangeNotification(tx.Statement.Context, user.SlackReference(true), slack.PermissionChangeNotificationInputs{
			Summary: fmt.Sprintf("created RoleAssignment for %s", ra.Description(tx)),
			Results: []string{
				"Fields: " + slack.EscapeText(ra.RoleAssignmentFields.String()),
			},
		})
	}
	return nil
}

func (ra *RoleAssignment) BeforeUpdate(tx *gorm.DB) error {
	if err := ra.errorIfForbidden(tx); err != nil {
		return err
	} else if err = copier.CopyWithOption(&ra.previousFields, &ra.RoleAssignmentFields, copier.Option{DeepCopy: true}); err != nil {
		return fmt.Errorf("failed to copy RoleAssignmentFields: %w", err)
	}
	return nil
}

func (ra *RoleAssignment) AfterUpdate(tx *gorm.DB) error {
	if user, err := GetCurrentUserForDB(tx); err != nil {
		return err
	} else if err = ra.errorIfForbidden(tx); err != nil {
		return err
	} else if err = tx.Omit(clause.Associations).Create(&RoleAssignmentOperation{
		RoleID:    ra.RoleID,
		UserID:    ra.UserID,
		AuthorID:  user.ID,
		Operation: "update",
		From:      ra.previousFields,
		To:        ra.RoleAssignmentFields,
	}).Error; err != nil {
		return fmt.Errorf("failed to create RoleAssignmentOperation: %w", err)
	} else {
		slack.SendPermissionChangeNotification(tx.Statement.Context, user.SlackReference(true), slack.PermissionChangeNotificationInputs{
			Summary: fmt.Sprintf("edited RoleAssignment for %s", ra.Description(tx)),
			Results: []string{
				"Old fields: " + slack.EscapeText(ra.previousFields.String()),
				"New fields: " + slack.EscapeText(ra.RoleAssignmentFields.String()),
			},
		})
	}
	return nil
}

func (ra *RoleAssignment) BeforeDelete(tx *gorm.DB) error {
	var current RoleAssignment
	if user, err := GetCurrentUserForDB(tx); err != nil {
		return err
	} else if err = ra.errorIfForbidden(tx); err != nil {
		return err
	} else if err = tx.Where(&RoleAssignment{RoleID: ra.RoleID, UserID: ra.UserID}).First(&current).Error; err != nil {
		return fmt.Errorf("failed to find current RoleAssignment: %w", err)
	} else if err = tx.Omit(clause.Associations).Create(&RoleAssignmentOperation{
		RoleID:    ra.RoleID,
		UserID:    ra.UserID,
		AuthorID:  user.ID,
		Operation: "delete",
		From:      current.RoleAssignmentFields,
	}).Error; err != nil {
		return fmt.Errorf("failed to create RoleAssignmentOperation: %w", err)
	} else {
		slack.SendPermissionChangeNotification(tx.Statement.Context, user.SlackReference(true), slack.PermissionChangeNotificationInputs{
			Summary: fmt.Sprintf("deleted RoleAssignment for %s", ra.Description(tx)),
			Results: []string{
				"Fields: " + slack.EscapeText(ra.RoleAssignmentFields.String()),
			},
		})
	}
	return nil
}
