package db

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// SetupMockRepository is a test utility function that can be used to
// supply a stub postgres backend that can in turn be used for simulating queries
func SetupMockRepository(t *testing.T, useRegexQueryMatcher bool) (*Repository, sqlmock.Sqlmock) {
	t.Helper()
	var (
		stubDB *sql.DB
		mock   sqlmock.Sqlmock
		err    error
	)
	if useRegexQueryMatcher {
		stubDB, mock, err = sqlmock.New()
	} else {
		stubDB, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	}
	if err != nil {
		t.Fatalf("error opening stub db: %v\n", err)
	}

	// creates a gorm instance with a stub db connection
	gdb, err := gorm.Open(
		postgres.New(postgres.Config{
			Conn: stubDB,
		}),
		&gorm.Config{},
	)
	if err != nil {
		t.Fatalf("err using gorm with stub db: %v", err)
	}

	stubRepository := NewRepository(gdb)
	return stubRepository, mock
}
