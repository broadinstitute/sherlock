package models

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication/authentication_method"
	"github.com/broadinstitute/sherlock/sherlock/internal/models/self"
	"gorm.io/gorm"
	"strings"
	"unicode"
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
	userMakingRequest, err := GetCurrentUserForDB(tx)
	if err != nil {
		return err
	}
	if userMakingRequest.ID != u.ID {
		if u.ID == 0 {
			// We can smartly catch a coding error here.
			// This is probably a result of doing something like db.Where().Updates() when you need to do
			// db.Where().First() followed by db.Model().Updates().
			// https://gorm.io/docs/update.html
			return fmt.Errorf("(%s) user ID in BeforeEdit was nil, possibly a bad database call", errors.InternalServerError)
		}
		return fmt.Errorf("(%s) users can only edit themselves", errors.Forbidden)
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

func (u *User) NameOrEmailHandle() string {
	if u.Name != nil {
		return *u.Name
	} else {
		return strings.Split(u.Email, "@")[0]
	}
}

func (u *User) SlackReference(mention bool) string {
	if u.SlackID != nil && mention {
		return fmt.Sprintf("<@%s>", *u.SlackID)
	} else {
		return fmt.Sprintf("<https://broad.io/beehive/r/user/%s|%s>", u.Email, u.NameOrEmailHandle())
	}
}

// ErrIfNotSuitable may be removed in the future: we'll want to be doing Sherlock's internal RBAC based on
// Role and RoleAssignment instead of Suitability, and the only usage of Suitability *should* be for
// suspending RoleAssignment entries. An error isn't really helpful for us there, so we may have no need
// for this method. For the moment, though, it offers a very similar logical interface to the now-removed
// `User.Suitable().SuitableOrError()` method, so it makes sense for now.
func (u *User) ErrIfNotSuitable() error {
	if u.Email == self.Email && u.GoogleID == self.GoogleID && u.AuthenticationMethod == authentication_method.SHERLOCK_INTERNAL {
		// Short-circuit to respect Sherlock's own user; see SelfUser.
		// We only respect this with an internal authentication method as defense-in-depth (it should be impossible to
		// actually make a request as Sherlock, but we don't want to find out).
		return nil
	}
	if u.Suitability != nil && u.Suitability.Suitable != nil && u.Suitability.Description != nil {
		if *u.Suitability.Suitable {
			return nil
		} else {
			return fmt.Errorf("(%s) user is unsuitable: %s", errors.Forbidden, *u.Suitability.Description)
		}
	} else {
		return fmt.Errorf("(%s) no matching suitability record found or loaded; assuming unsuitable", errors.Forbidden)
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
			*assignment.Role.GrantsSherlockSuperAdmin &&
			assignment.IsActive() {
			return nil
		}
	}
	return fmt.Errorf("(%s) caller is not a super-admin", errors.Forbidden)
}
