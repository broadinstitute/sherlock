package models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/authentication_method"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/db"
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
	// outside tests. You'll run into circular
	// dependency issues if you try to import something
	// from this package into the db package. You might
	// be able to run this auto-migrate process from
	// the boot package, but at that point you should
	// consider just going ahead and writing proper SQL
	// to represent your model.
	//
	//err = h.internalDB.AutoMigrate(&SomeNewModel{})
	//if err != nil {
	//	panic(err)
	//}
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
func (h *TestSuiteHelper) SetUserForDB(user *User) *User {
	user.AuthenticationMethod = authentication_method.TEST
	h.DB = SetCurrentUserForDB(h.DB, user)
	return user
}

// SetSuitableTestUserForDB is a helper function, calling SetUserForDB with
// TestData.User_Suitable
func (h *TestSuiteHelper) SetSuitableTestUserForDB() *User {
	return h.SetUserForDB(utils.PointerTo(h.TestData.User_Suitable()))
}

// SetNonSuitableTestUserForDB is a helper function, calling SetUserForDB with
// TestData.User_NonSuitable
func (h *TestSuiteHelper) SetNonSuitableTestUserForDB() *User {
	return h.SetUserForDB(utils.PointerTo(h.TestData.User_NonSuitable()))
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
