package models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication/authentication_method"
	"gorm.io/gorm"
)

// TestSuiteHelper can be embedded in a test suite struct to
// help manage the database connection, test transaction,
// and the Gorm instance's current user.
type TestSuiteHelper struct {
	DB         *gorm.DB
	internalDB *gorm.DB
	TestData   TestData
}

// SetupSuite runs once before all tests. It connects to the
// database and runs migrations.
func (h *TestSuiteHelper) SetupSuite() {
	config.LoadTestConfig()
	sqlDB, err := db.Connect()
	if err != nil {
		panic(err)
	}
	h.internalDB, err = db.Configure(sqlDB)
	if err != nil {
		panic(err)
	}

	if err = Init(h.internalDB); err != nil {
		panic(err)
	}
}

// SetupTest begins a transaction and sets the main database
// reference on modelSuite accordingly. This means that
// TearDownTest will be able to roll back the entire test's
// changes.
func (h *TestSuiteHelper) SetupTest() {
	h.DB = h.internalDB.Begin()
	h.TestData = &testDataImpl{h: h}
}

// SetUserForDB is a low-level helper function, setting with given user as
// the current principal for the database. You'll usually want to call
// SetSuitableTestUserForDB or SetNonSuitableTestUserForDB instead.
func (h *TestSuiteHelper) SetUserForDB(user *User, reload ...bool) *User {
	if len(reload) > 0 && reload[0] {
		err := h.DB.Scopes(ReadUserScope).First(&user, user.ID).Error
		if err != nil {
			panic(fmt.Errorf("failed to reload user: %w", err))
		}
	}
	if user != nil && user.AuthenticationMethod != authentication_method.SHERLOCK_INTERNAL {
		user.AuthenticationMethod = authentication_method.TEST
	}
	h.DB = SetCurrentUserForDB(h.DB, user)
	return user
}

// SetSuitableTestUserForDB is a helper function, calling SetUserForDB with
// TestData.User_Suitable
func (h *TestSuiteHelper) SetSuitableTestUserForDB(reload ...bool) *User {
	return h.SetUserForDB(utils.PointerTo(h.TestData.User_Suitable()), reload...)
}

// SetNonSuitableTestUserForDB is a helper function, calling SetUserForDB with
// TestData.User_NonSuitable
func (h *TestSuiteHelper) SetNonSuitableTestUserForDB(reload ...bool) *User {
	return h.SetUserForDB(utils.PointerTo(h.TestData.User_NonSuitable()), reload...)
}

// SetSelfSuperAdminForDB is a helper function, calling SetUserForDB with
// SelfUser. This is different from other similar helpers in that the user
// being set isn't coming from TestData; this is a system-level user that
// is necessary for Sherlock's actual runtime. In tests, it's a convenient
// way to get super-user privileges when the "who" isn't important.
func (h *TestSuiteHelper) SetSelfSuperAdminForDB() *User {
	return h.SetUserForDB(SelfUser)
}

// TearDownTest takes advantage of SetupTest having begun a
// transaction to roll back the entire test's changes. It
// sets the modelSuite's main database reference to nil to
// help surface any concurrency issues.
func (h *TestSuiteHelper) TearDownTest() {
	h.TestData = nil
	h.DB.Rollback()
	h.DB = nil
}

// TearDownSuite closes the entire database connection once
// all tests have completed.
func (h *TestSuiteHelper) TearDownSuite() {
	sqlDB, err := h.internalDB.DB()
	if err != nil {
		panic(err)
	}
	err = sqlDB.Close()
	if err != nil {
		panic(err)
	}
}
