package models

import (
	"context"
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication/authentication_method"
	"github.com/broadinstitute/sherlock/sherlock/internal/models/self"
	"gorm.io/gorm"
)

// SelfUser is the User that Sherlock itself is conceptually running as, as derived from
// self.Email and self.GoogleID.
//
// This reference won't be updated during Sherlock's runtime, but that's okay because
// the User.ErrIfNotSuperAdmin method short-circuits when the User.Email and
// User.GoogleID match self.Email and self.GoogleID respectively, automatically
// considering the User a super-admin.
var SelfUser *User

// sherlockDbContextField exists as a type to avoid collisions with other keys;
// linting errors if you don't do this.
type sherlockDbContextField string

const dbUserField sherlockDbContextField = "SherlockUser"

func SetCurrentUserForDB(db *gorm.DB, user *User) *gorm.DB {
	return db.WithContext(context.WithValue(db.Statement.Context, dbUserField, user))
}

func GetCurrentUserForDB(db *gorm.DB) (*User, error) {
	user, ok := db.Statement.Context.Value(dbUserField).(*User)
	if !ok {
		return nil, fmt.Errorf("(%s) database user not available (was %T)", errors.InternalServerError, db.Statement.Context.Value(dbUserField))
	}
	if user == nil {
		return nil, fmt.Errorf("(%s) database user was nil", errors.InternalServerError)
	}
	return user, nil
}

// User is a more complex type than the other types in this package simply because we don't
// just use it as a CRUD model -- it's also the struct used as the touch-point for
// authentication and authorization.
type User struct {
	gorm.Model
	Email          string `gorm:"index"`
	GoogleID       string
	GithubUsername *string
	GithubID       *string
	SlackUsername  *string
	SlackID        *string

	Name *string
	// NameFrom must be either "sherlock", "github", or "slack"
	NameFrom *string

	// DeactivatedAt is a nullable timestamp that, when not null, indicates that the User is deactivated.
	// This is like a soft-delete but it has semantic meaning within Sherlock (the User can't authenticate
	// and can't have any RoleAssignments but could have in the past and does still exist), unlike
	// gorm.Model's DeletedAt which already has a meaning enforced by Gorm.
	//
	// Nothing uses this time -- it's essentially just for reference.
	DeactivatedAt *time.Time

	// Suitability is a potential reference to a matching Suitability record, which in turn
	// potentially indicates that the User is "suitable" for production access. If this is
	// nil the User should be assumed to be unsuitable.
	Suitability *Suitability `gorm:"foreignKey:Email;references:Email"`

	// Assignments lists Role records that this User is assigned to. A RoleAssignment can potentially be suspended,
	// which indicates that the User should not presently have any access commensurate with the corresponding Role.
	// More information on this behavior is available on the Role type.
	Assignments []*RoleAssignment

	// Via is ignored by Gorm and isn't stored in the database -- it should be set at the application layer
	// when the User is changed for a given request to represent the previous User.
	// It won't be defined when the User is queried out of the database, only as they authenticate.
	// An example is when substituting a User from GHA OIDC over a User from IAP: the GHA OIDC User would be
	// "via" the initially-derived IAP User.
	// See the authentication package for more information.
	Via *User `gorm:"-:all"`

	// AuthenticationMethod is ignored by Gorm and isn't stored in the database -- it should be set at the
	// application layer to describe how this User was authenticated.
	// It won't be defined when the User is queried out of the database, only as they authenticate.
	// See the authentication package for more information.
	AuthenticationMethod authentication_method.Method `gorm:"-:all"`
}

// ReadUserScope should be used in place of `db.Preload(clause.Associations)` for reading User records, as it
// properly loads the Role records opposite the many-to-many RoleAssignment relationship.
func ReadUserScope(db *gorm.DB) *gorm.DB {
	return db.
		Preload("Suitability").
		Preload("Assignments").
		Preload("Assignments.Role")
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	if tx.Statement.Changed("ID", "Email", "GoogleID") {
		// (Jack) I don't think it's possible to hit this case unless some internal code tried it,
		// since these fields aren't editable in the API, but might as well just chuck an error here.
		// If this ever fires I suppose we'd be glad it did.
		// I don't think it's worth being this exhaustive elsewhere, but with this table I think it makes
		// sense because of how intertwined it is with authentication.
		return fmt.Errorf("(%s) email and google ID cannot be changed", errors.BadRequest)
	}
	if u.ID == 0 {
		// We can smartly catch a coding error here.
		// This is probably a result of doing something like db.Where().Updates() when you need to do
		// db.Where().First() followed by db.Model().Updates().
		// https://gorm.io/docs/update.html
		return fmt.Errorf("(%s) user ID in BeforeEdit was nil, possibly a bad database call", errors.InternalServerError)
	}

	userMakingRequest, err := GetCurrentUserForDB(tx)
	if err != nil {
		return err
	}

	if tx.Statement.Changed("DeactivatedAt") && userMakingRequest.ID == u.ID {
		// We shouldn't be able to hit this because anything that updates the deactivated_at field should
		// also be checking this... but better safe than sorry.
		return fmt.Errorf("(%s) users cannot deactivate themselves", errors.BadRequest)
	}

	if errWhenNotSuperAdmin := userMakingRequest.ErrIfNotSuperAdmin(); errWhenNotSuperAdmin != nil && userMakingRequest.ID != u.ID {
		return fmt.Errorf("only super admins can edit other users: %w", errWhenNotSuperAdmin)
	}

	if (config.Config.MustString("mode") != "debug" && userMakingRequest.AuthenticationMethod != authentication_method.IAP) ||
		(config.Config.MustString("mode") == "debug" && userMakingRequest.AuthenticationMethod != authentication_method.TEST && userMakingRequest.AuthenticationMethod != authentication_method.LOCAL) {
		// For non-debug, require IAP to edit users. For debug, require TEST or LOCAL so we can still hit this error case if we really try.
		return fmt.Errorf("(%s) users cannot be edited via this authentication method while sherlock is in %s mode", errors.Forbidden, config.Config.MustString("mode"))
	}
	return nil
}

func (u *User) BeforeDelete(_ *gorm.DB) error {
	return fmt.Errorf("(%s) users cannot be deleted", errors.Forbidden)
}

func (u *User) AfterCreate(db *gorm.DB) error {
	var actor string
	if dbUser, err := GetCurrentUserForDB(db); err == nil {
		actor = dbUser.SlackReference(true)
	} else {
		actor = u.SlackReference(true)
	}
	slack.SendPermissionChangeNotification(db.Statement.Context, actor, slack.PermissionChangeNotificationInputs{
		Summary: fmt.Sprintf("created User %s", u.SlackReference(true)),
	})
	return nil
}

func (u *User) AlphaNumericHyphenatedUsername() string {
	var ret []rune
	for _, r := range strings.Split(u.Email, "@")[0] {
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			ret = append(ret, r)
		} else if r == '.' || r == '-' || r == '_' {
			ret = append(ret, '-')
		}
	}
	return string(ret)
}

func (u *User) NameOrUsername() string {
	if u.Name != nil {
		return *u.Name
	} else {
		return u.AlphaNumericHyphenatedUsername()
	}
}

func (u *User) SlackReference(mention bool) string {
	if u.SlackID != nil && mention {
		return fmt.Sprintf("<@%s>", *u.SlackID)
	} else {
		return fmt.Sprintf("<https://broad.io/beehive/r/user/%s|%s>", u.Email, u.NameOrUsername())
	}
}

func (u *User) ErrIfNotActiveInRole(db *gorm.DB, roleID *uint) error {
	// If a role ID isn't actually provided, no role is required, so we can short-circuit.
	if roleID == nil {
		return nil
	}

	// Search for matching assignment
	for _, assignment := range u.Assignments {
		if assignment.RoleID == *roleID {
			if assignment.Role == nil {
				return fmt.Errorf("(%s) issue loading required role %d for caller", errors.InternalServerError, *roleID)
			} else if assignment.Role.Name == nil {
				return fmt.Errorf("(%s) issue loading required role %d for caller (name missing)", errors.InternalServerError, *roleID)
			} else if err := assignment.ErrIfNotActive(); err != nil {
				return fmt.Errorf("(%s) caller is in required role '%s' but they're not active: %w", errors.Forbidden, *assignment.Role.Name, err)
			} else {
				return nil
			}
		}
	}

	// If no matching assignment, check if the caller is a super-admin
	if err := u.ErrIfNotSuperAdmin(); err == nil {
		return nil
	}

	// If no matching assignment, get the name for the error message
	var role Role
	if err := db.Select("name").Take(&role, roleID).Error; err != nil {
		return fmt.Errorf("(%s) role %d required not found", errors.InternalServerError, *roleID)
	} else if role.Name == nil {
		return fmt.Errorf("(%s) caller is not in required role %d (name missing)", errors.InternalServerError, *roleID)
	} else {
		return fmt.Errorf("(%s) caller is not in required role '%s'", errors.Forbidden, *role.Name)
	}
}

func (u *User) ErrIfNotSuperAdmin() error {
	if u.Email == self.Email && u.GoogleID == self.GoogleID && u.AuthenticationMethod == authentication_method.SHERLOCK_INTERNAL {
		// Short-circuit to respect Sherlock's own user; see SelfUser.
		// We only respect this with an internal authentication method as defense-in-depth (it should be impossible to
		// actually make a request as Sherlock, but we don't want to find out).
		return nil
	}
	for _, assignment := range u.Assignments {
		if assignment.Role.GrantsSherlockSuperAdmin != nil &&
			*assignment.Role.GrantsSherlockSuperAdmin {
			// Only one Sherlock role can grant super admin, so we just check if the match is active
			if err := assignment.ErrIfNotActive(); err != nil {
				return fmt.Errorf("(%s) caller is in a role that grants super-admin but they're not active: %w", errors.Forbidden, err)
			} else {
				return nil
			}
		}
	}
	return fmt.Errorf("(%s) caller is not a super-admin", errors.Forbidden)
}
