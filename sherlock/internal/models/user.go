package models

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/authentication_method"
	"github.com/broadinstitute/sherlock/sherlock/internal/authorization"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"gorm.io/gorm"
	"strings"
	"unicode"
)

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
	Email          string
	GoogleID       string
	GithubUsername *string
	GithubID       *string
	SlackUsername  *string
	SlackID        *string

	Name *string
	// NameFrom must be either "sherlock", "github", or "slack"
	NameFrom *string

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

	// cachedSuitability is ignored by Gorm and isn't stored in the database -- it is used to cache calls to Suitability,
	// which looks up the user's authorization.Suitability.
	// See the authorization package for more information.
	// In the future, Sherlock will become its own source of truth for suitability and other authorization, in
	// which case this behavior will become database-persistent and may be entirely represented in the database.
	cachedSuitability *authorization.Suitability `gorm:"-:all"`
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

func (u *User) Suitability() *authorization.Suitability {
	if u.Email != "" && u.cachedSuitability == nil {
		u.cachedSuitability = authorization.GetSuitabilityFor(u.Email)
	}
	return u.cachedSuitability
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

func (u *User) SlackReference() string {
	if u.SlackID != nil {
		return fmt.Sprintf("<@%s>", *u.SlackID)
	} else {
		return fmt.Sprintf("<https://broad.io/beehive/r/user/%s|%s>", u.Email, u.NameOrEmailHandle())
	}
}
