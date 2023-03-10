package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth/auth_models"
	"github.com/broadinstitute/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/internal/models/model_actions"
	"github.com/broadinstitute/sherlock/internal/utils"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type User struct {
	gorm.Model
	auth_models.StoredControlledUserFields
	auth_models.StoredMutableUserFields
}

func (u User) TableName() string {
	return "v2_users"
}

func (u User) getID() uint {
	return u.ID
}

var userStore *internalModelStore[User]

func init() {
	userStore = &internalModelStore[User]{
		selectorToQueryModel:    userSelectorToQuery,
		modelToSelectors:        userToSelectors,
		errorIfForbidden:        userErrorIfForbidden,
		validateModel:           validateUser,
		editsMayChangeSelectors: true,
	}
}

func userSelectorToQuery(_ *gorm.DB, selector string) (User, error) {
	if len(selector) == 0 {
		return User{}, fmt.Errorf("(%s) user selector cannot be empty", errors.BadRequest)
	}
	var query User
	if utils.IsNumeric(selector) { // ID
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
	} else if strings.HasPrefix(selector, "github/") { // "github/" + GitHub Username
		githubUsername := strings.TrimPrefix(selector, "github/")
		query.GithubUsername = &githubUsername
		return query, nil
	} else if strings.HasPrefix(selector, "github-id/") { // "github-id/" + GitHub ID
		githubID := strings.TrimPrefix(selector, "github-id/")
		query.GithubID = &githubID
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
		if user.GithubUsername != nil {
			selectors = append(selectors, fmt.Sprintf("github/%s", *user.GithubUsername))
		}
		if user.GithubID != nil {
			selectors = append(selectors, fmt.Sprintf("github-id/%s", *user.GithubID))
		}
	}
	return selectors
}

func userErrorIfForbidden(_ *gorm.DB, modelUser *User, action model_actions.ActionType, user *auth_models.User) error {
	switch action {
	case model_actions.CREATE:
		if user != nil {
			// The handler/controller always pass the user, so it being nil means we're still at the auth middleware.
			return fmt.Errorf("users can only be created via middleware")
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
	if (user.GithubUsername == nil) != (user.GithubID == nil) {
		return fmt.Errorf("a %T must either have both a GitHub username and ID set or neither", user)
	}
	return nil
}
