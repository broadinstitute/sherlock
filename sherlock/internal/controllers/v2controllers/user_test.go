package v2controllers

import (
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/auth/auth_models"
	"github.com/broadinstitute/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"github.com/broadinstitute/sherlock/internal/testutils"
	"github.com/google/go-github/v50/github"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"strconv"
	"testing"
)

func TestUserControllerSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping functional test")
	}
	suite.Run(t, new(userControllerSuite))
}

type userControllerSuite struct {
	suite.Suite
	*ControllerSet
	middleware *v2models.MiddlewareUserStore
	db         *gorm.DB
}

func (suite *userControllerSuite) SetupTest() {
	config.LoadTestConfig(suite.T())
	suite.db = db.ConnectAndConfigureFromTest(suite.T())
	suite.db.Begin()
	suite.ControllerSet = NewControllerSet(v2models.NewStoreSet(suite.db))
	suite.middleware = v2models.NewMiddlewareUserStore(suite.db)
}

func (suite *userControllerSuite) TearDownTest() {
	suite.db.Rollback()
}

func (suite *userControllerSuite) TestUserFlow() {
	db.Truncate(suite.T(), suite.db)

	// It's impossible to create a user via a controller. The creation type doesn't support some required fields.
	// The model also won't allow a normally-authenticated user to make a user entry (it only allows creation when the
	// auth user is nil) but we can't reach that error case from here (see v2models/user_test.go for unit tests)
	_, created, err := suite.UserController.Create(CreatableUser{}, auth.GenerateUser(suite.T(), suite.db, true))
	assert.ErrorContains(suite.T(), err, errors.BadRequest)
	assert.False(suite.T(), created)
	// The controller does check that we can't pass a nil user, so the model's creation method isn't accessible via
	// controller.
	_, created, err = suite.UserController.Create(CreatableUser{}, nil)
	assert.ErrorContains(suite.T(), err, errors.InternalServerError)
	assert.False(suite.T(), created)

	// Instead, the middleware may use its MiddlewareUserStore to get or create a user.
	generatedUser := auth.GenerateUser(suite.T(), suite.db, false)
	modelUser, err := suite.middleware.GetOrCreateUser(generatedUser.Email, generatedUser.GoogleID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), generatedUser.Email, modelUser.Email)
	assert.Equal(suite.T(), generatedUser.GoogleID, modelUser.GoogleID)
	assert.NotEmpty(suite.T(), modelUser.UpdatedAt)

	// The same middleware call can happen repeatedly, and it won't create or edit the user entry.
	suite.Run("subsequent calls just read", func() {
		readUserOne, err := suite.middleware.GetOrCreateUser(generatedUser.Email, generatedUser.GoogleID)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), generatedUser.Email, modelUser.Email)
		assert.Equal(suite.T(), generatedUser.GoogleID, modelUser.GoogleID)
		assert.Equal(suite.T(), modelUser.UpdatedAt, readUserOne.UpdatedAt)
		readUserTwo, err := suite.middleware.GetOrCreateUser(generatedUser.Email, generatedUser.GoogleID)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), generatedUser.Email, modelUser.Email)
		assert.Equal(suite.T(), generatedUser.GoogleID, modelUser.GoogleID)
		assert.Equal(suite.T(), modelUser.UpdatedAt, readUserTwo.UpdatedAt)
	})

	// Let's assume that we're firmly in a subsequent request now--each one will load the user out of the database,
	// so we'll patch up the object we're using like that had happened.
	generatedUser.ID = modelUser.ID
	generatedUser.StoredControlledUserFields = modelUser.StoredControlledUserFields

	// The user can be read back out via the controller once it's created.
	controllerUser, err := suite.UserController.Get(modelUser.Email)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), modelUser.ID, controllerUser.ID)
	assert.Equal(suite.T(), modelUser.GoogleID, controllerUser.GoogleID)

	// Suppose we create some more users:
	thingOneUser, err := suite.middleware.GetOrCreateUser("thing-one@example.com", "1")
	assert.NoError(suite.T(), err)
	_, err = suite.middleware.GetOrCreateUser("thing-two@example.com", "2")
	assert.NoError(suite.T(), err)

	// We can list all users normally.
	suite.Run("listing works", func() {
		results, err := suite.UserController.ListAllMatching(User{}, 0)
		assert.NoError(suite.T(), err)
		assert.Len(suite.T(), results, 3)
		assert.Contains(suite.T(), results, controllerUser)
	})

	// The middleware also has a shortcut to a narrow "get matching github user if present" method.
	// It returns nil when there's no match, rather than an error.
	generatedUserGithubID := int64(12341234)
	nilUser, err := suite.middleware.GetGithubUserIfExists(strconv.FormatInt(generatedUserGithubID, 10))
	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), nilUser)

	// The fields that store github information can't be directly modified. The controller exports a method
	// that accepts a github access token to get information in a verified way before calling an unexported
	// method to write the information. This prevents people from modifying anyone else's github information or
	// changing theirs to that of an account they don't already control.
	// We can't call the github api from this functional test, so we'll call the unexported method directly.
	generatedUserGithubUsername := "example-github"
	githubPayload := &github.User{
		ID:    &generatedUserGithubID,
		Login: &generatedUserGithubUsername,
	}
	controllerUserWithGithub, updated, err := suite.UserController.recordGithubInformation(githubPayload, generatedUser)
	assert.NoError(suite.T(), err)
	assert.True(suite.T(), updated)
	assert.Equal(suite.T(), controllerUser.ID, controllerUserWithGithub.ID)

	// Each request will load the user out of the database, but suppose there's a race and we process another payload
	// first with the old user. The database will accept the write and we won't error either:
	controllerUserWithGithub, updated, err = suite.UserController.recordGithubInformation(githubPayload, generatedUser)
	assert.NoError(suite.T(), err)
	assert.True(suite.T(), updated)
	assert.Equal(suite.T(), controllerUser.ID, controllerUserWithGithub.ID)

	// Just for test purposes, we'll patch up the auth user object again
	generatedUser.StoredControlledUserFields = controllerUserWithGithub.StoredControlledUserFields

	// Once the user has github info, then the middleware's shortcut to read it will return non nil:
	suite.Run("shortcut works", func() {
		shortcutUser, err := suite.middleware.GetGithubUserIfExists(strconv.FormatInt(generatedUserGithubID, 10))
		assert.NoError(suite.T(), err)
		assert.NotNil(suite.T(), shortcutUser)
		assert.Equal(suite.T(), controllerUserWithGithub.ID, shortcutUser.ID)
	})

	// The record github method will then just read instead of writing.
	suite.Run("subsequent github record calls just read", func() {
		readUserOne, updated, err := suite.UserController.recordGithubInformation(githubPayload, generatedUser)
		assert.NoError(suite.T(), err)
		assert.False(suite.T(), updated)
		assert.Equal(suite.T(), controllerUserWithGithub.UpdatedAt, readUserOne.UpdatedAt)
		readUserTwo, updated, err := suite.UserController.recordGithubInformation(githubPayload, generatedUser)
		assert.NoError(suite.T(), err)
		assert.False(suite.T(), updated)
		assert.Equal(suite.T(), controllerUserWithGithub.UpdatedAt, readUserTwo.UpdatedAt)
	})

	// It will write if there's a change, though. Suppose the human adds a name to their github profile:
	generatedUserGithubName := "mallory"
	githubPayload.Name = &generatedUserGithubName
	controllerUserWithGithub, updated, err = suite.UserController.recordGithubInformation(githubPayload, generatedUser)
	assert.NoError(suite.T(), err)
	assert.True(suite.T(), updated)
	assert.Equal(suite.T(), controllerUser.ID, controllerUserWithGithub.ID)
	assert.Equal(suite.T(), generatedUserGithubName, *controllerUserWithGithub.Name)

	// This would also get pulled in again upon a new request.
	generatedUser.StoredMutableUserFields = controllerUserWithGithub.StoredMutableUserFields

	// Consecutive calls still read only (there was a bug here once, hence the weirdly specific test).
	suite.Run("subsequent github record calls with name just read", func() {
		readUserOne, updated, err := suite.UserController.recordGithubInformation(githubPayload, generatedUser)
		assert.NoError(suite.T(), err)
		assert.False(suite.T(), updated)
		assert.Equal(suite.T(), controllerUserWithGithub.UpdatedAt, readUserOne.UpdatedAt)
		readUserTwo, updated, err := suite.UserController.recordGithubInformation(githubPayload, generatedUser)
		assert.NoError(suite.T(), err)
		assert.False(suite.T(), updated)
		assert.Equal(suite.T(), controllerUserWithGithub.UpdatedAt, readUserTwo.UpdatedAt)
	})

	// The name can be edited in the API...
	otherName := "MALLORY"
	controllerUserWithGithub, err = suite.UserController.Edit(generatedUser.Email, EditableUser{
		StoredMutableUserFields: auth_models.StoredMutableUserFields{
			Name: &otherName,
		},
	}, generatedUser)
	assert.NoError(suite.T(), err)
	assert.True(suite.T(), updated)
	assert.Equal(suite.T(), controllerUser.ID, controllerUserWithGithub.ID)
	assert.Equal(suite.T(), otherName, *controllerUserWithGithub.Name)
	generatedUser.StoredMutableUserFields = controllerUserWithGithub.StoredMutableUserFields

	// But it'll get updated by github still...
	controllerUserWithGithub, updated, err = suite.UserController.recordGithubInformation(githubPayload, generatedUser)
	assert.NoError(suite.T(), err)
	assert.True(suite.T(), updated)
	assert.Equal(suite.T(), controllerUser.ID, controllerUserWithGithub.ID)
	assert.Equal(suite.T(), generatedUserGithubName, *controllerUserWithGithub.Name)
	generatedUser.StoredMutableUserFields = controllerUserWithGithub.StoredMutableUserFields

	// So there's an API flag to choose whether github will update it:
	controllerUserWithGithub, err = suite.UserController.Edit(generatedUser.Email, EditableUser{
		StoredMutableUserFields: auth_models.StoredMutableUserFields{
			Name:                   &otherName,
			NameInferredFromGithub: testutils.PointerTo(false),
		},
	}, generatedUser)
	assert.NoError(suite.T(), err)
	assert.True(suite.T(), updated)
	assert.Equal(suite.T(), controllerUser.ID, controllerUserWithGithub.ID)
	assert.Equal(suite.T(), otherName, *controllerUserWithGithub.Name)
	generatedUser.StoredMutableUserFields = controllerUserWithGithub.StoredMutableUserFields
	controllerUserWithGithub, updated, err = suite.UserController.recordGithubInformation(githubPayload, generatedUser)
	assert.NoError(suite.T(), err)
	assert.False(suite.T(), updated)
	assert.Equal(suite.T(), controllerUser.ID, controllerUserWithGithub.ID)
	assert.Equal(suite.T(), otherName, *controllerUserWithGithub.Name)
	generatedUser.StoredMutableUserFields = controllerUserWithGithub.StoredMutableUserFields

	// Editing is guarded at a permissions level, stopping people from modifying each other's accounts.
	_, err = suite.UserController.Edit(thingOneUser.Email, EditableUser{
		StoredMutableUserFields: auth_models.StoredMutableUserFields{
			Name: &otherName,
		},
	}, generatedUser)
	assert.ErrorContains(suite.T(), err, errors.Forbidden)

	// Deleting other people's accounts is also prevented.
	_, err = suite.UserController.Delete(thingOneUser.Email, generatedUser)
	assert.ErrorContains(suite.T(), err, errors.Forbidden)

	// As is deleting your accounts -- no deletions ever.
	_, err = suite.UserController.Delete(generatedUser.Email, generatedUser)
	assert.ErrorContains(suite.T(), err, errors.Forbidden)
}
