package models

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/authentication_method"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/test_users"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/db"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
)

// modelSuite has two responsibilities:
//   - database connection and test cleanup
//   - helping fulfil the "user in database context" contract that
//     models rely on
type modelSuite struct {
	suite.Suite
	db *gorm.DB

	// internalDB shouldn't be called by tests; it's the database
	// reference that exists outside the test transaction set up
	// by SetupTest.
	internalDB *gorm.DB
}

// TestModelSuite makes `go test` aware of the modelSuite tests.
func TestModelSuite(t *testing.T) {
	suite.Run(t, new(modelSuite))
}

// SetupSuite runs once before all tests. It connects to the
// database and runs migrations.
func (s *modelSuite) SetupSuite() {
	config.LoadTestConfig()
	sqlDB, err := db.Connect()
	if err != nil {
		panic(err)
	}
	s.internalDB, err = db.Configure(sqlDB)
	if err != nil {
		panic(err)
	}

	// If you're working on a new model and want to have
	// Gorm basically fudge the database schema from your
	// struct, here's a good place to do that. This is
	// useful in two ways:
	// 1. You can write tests and play with your model
	//    before going and writing actual SQL to do the
	//    migration.
	// 2. This isn't cleaned up, so you can inspect the
	//    database structure after tests to get a head
	//    start on writing the migration SQL. Tools
	//    like DataGrip and GoLand can diff two database
	//    structures and generate migration SQL, for
	//    example.
	// Right now, you will still have to write migration
	// SQL for Sherlock to use your new model when run
	// from outside tests. You'll run into circular
	// dependency issues if you try to import something
	// from this package into the db package. You might
	// be able to run this auto-migrate process from
	// the boot package, but at that point you should
	// consider just going ahead and writing proper SQL
	// to represent your model.
	//
	//err = s.internalDB.AutoMigrate(SomeNewModel)
	//if err != nil {
	//	panic(err)
	//}
}

// SetupTest begins a transaction and sets the main database
// reference on modelSuite accordingly. This means that
// TearDownTest will be able to roll back the entire test's
// changes.
func (s *modelSuite) SetupTest() {
	s.db = s.internalDB.Begin()
}

// UseUser is intended to be called from within a test function.
// It will upsert a User with the given email and googleID and
// set it as the modelSuite db's current user.
func (s *modelSuite) UseUser(email, googleID string) *User {
	var user User
	if err := s.db.
		Where(&User{Email: email, GoogleID: googleID}).
		FirstOrCreate(&user).Error; err != nil {
		panic(err)
	} else {
		user.AuthenticationMethod = authentication_method.TEST
		s.db = SetCurrentUserForDB(s.db, &user)
		return &user
	}
}

// UseSuitableTestUser is intended to be called from within a
// test function. It calls UseUser with the suitable test user
// info from the test_users package, which will be recognized
// as suitable by the authorization package.
func (s *modelSuite) UseSuitableTestUser() *User {
	return s.UseUser(test_users.SuitableTestUserEmail, test_users.SuitableTestUserGoogleID)
}

// UseNonSuitableTestUser is intended to be called from within
// a test function. It calls UseUser with the non-suitable
// test user from the test_users package, which will be
// recognized but considered non-suitable by the authorization
// package.
func (s *modelSuite) UseNonSuitableTestUser() *User {
	return s.UseUser(test_users.NonSuitableTestUserEmail, test_users.NonSuitableTestUserGoogleID)
}

// TearDownTest takes advantage of SetupTest having begun a
// transaction to roll back the entire test's changes. It
// sets the modelSuite's main database reference to nil to
// help surface any concurrency issues.
func (s *modelSuite) TearDownTest() {
	s.db.Rollback()
	s.db = nil
}

// TearDownSuite closes the entire database connection once
// all tests have completed.
func (s *modelSuite) TearDownSuite() {
	sqlDB, err := s.internalDB.DB()
	if err != nil {
		panic(err)
	}
	err = sqlDB.Close()
	if err != nil {
		panic(err)
	}
}

// TestModelSuiteItself checks that modelSuite's user helper
// methods work properly (and thus also that the database
// connection is working).
func (s *modelSuite) TestModelSuiteItself() {
	s.Run("no user by default", func() {
		user, err := GetCurrentUserForDB(s.db)
		s.ErrorContains(err, "database user not available")
		s.Nil(user)
	})
	s.Run("suitable test user", func() {
		s.UseSuitableTestUser()
		user, err := GetCurrentUserForDB(s.db)
		s.NoError(err)
		s.Equal(test_users.SuitableTestUserEmail, user.Email)
		s.NotZero(user.ID)
		s.True(user.Suitability().Suitable())
	})
	s.Run("non-suitable test user", func() {
		s.UseNonSuitableTestUser()
		user, err := GetCurrentUserForDB(s.db)
		s.NoError(err)
		s.Equal(test_users.NonSuitableTestUserEmail, user.Email)
		s.NotZero(user.ID)
		s.False(user.Suitability().Suitable())
	})
}
