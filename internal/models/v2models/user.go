package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth/auth_models"
	"github.com/broadinstitute/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/internal/models/model_actions"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type User struct {
	gorm.Model
	auth_models.StoredUserFields
}

func (u User) TableName() string {
	return "v2_users"
}

var userStore *internalModelStore[User]

func init() {
	userStore = &internalModelStore[User]{
		selectorToQueryModel: userSelectorToQuery,
		modelToSelectors:     userToSelectors,
		errorIfForbidden:     userErrorIfForbidden,
		validateModel:        validateUser,
	}
}

func userSelectorToQuery(_ *gorm.DB, selector string) (User, error) {
	if len(selector) == 0 {
		return User{}, fmt.Errorf("(%s) user selector cannot be empty", errors.BadRequest)
	}
	var query User
	if isNumeric(selector) { // ID
		id, err := strconv.Atoi(selector)
		if err != nil {
			return User{}, fmt.Errorf("(%s) string to int conversion error of '%s': %v", errors.BadRequest, selector, err)
		}
		query.ID = uint(id)
		return query, nil
	} else if strings.Contains(selector, "@") { // email
		query.Email = selector
		return query, nil
	} else if strings.HasPrefix(selector, "google-id/") { // "google-id/" + Google Subject ID
		query.GoogleID = strings.TrimPrefix(selector, "google-id/")
		return query, nil
	}
	return User{}, fmt.Errorf("(%s) invalid user selector '%s'", errors.BadRequest, selector)
}

func userToSelectors(user *User) []string {
	var selectors []string
	if user != nil {
		if user.Email != "" {
			selectors = append(selectors, user.Email)
		}
		if user.GoogleID != "" {
			selectors = append(selectors, fmt.Sprintf("google-id/%s", user.GoogleID))
		}
		if user.ID != 0 {
			selectors = append(selectors, fmt.Sprintf("%d", user.ID))
		}
	}
	return selectors
}

func userErrorIfForbidden(_ *gorm.DB, modelUser *User, action model_actions.ActionType, user *auth_models.User) error {
	switch action {
	case model_actions.CREATE:
		if user != nil {
			// The controller always sets the user, so it being nil means we're still at the auth middleware.
			return fmt.Errorf("users can only be created during their first request")
		}
	case model_actions.EDIT:
		if modelUser.Email != user.Email {
			return fmt.Errorf("users can only edit themselves")
		}
	case model_actions.DELETE:
		return fmt.Errorf("users cannot be deleted")
	}
	return nil
}

func validateUser(user *User) error {
	if user == nil {
		return fmt.Errorf("the model passed was nil")
	}
	if user.Email == "" {
		return fmt.Errorf("a %T must have an email", user)
	} else if !strings.Contains(user.Email, "@") {
		return fmt.Errorf("a %T must have an email: '%s' did not contain an '@'", user, user.Email)
	}
	if user.GoogleID == "" {
		return fmt.Errorf("a %T must have a Google subject ID", user)
	}
	return nil
}
