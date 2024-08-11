package db

import (
	"cloud.google.com/go/cloudsqlconn"
	"cloud.google.com/go/cloudsqlconn/postgres/pgxv5"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
)

// RegisterDriver handles setup for the cloudsql-postgres driver, if it's what'll be used.
// That driver can't be used for tests (as a Cloud SQL database can't be used for tests)
// so it is safe for tests to skip calling this and handling the resulting cleanup function.
func RegisterDriver() (cleanup func() error, err error) {
	if config.Config.MustString("db.driver") == "cloudsql-postgres" {
		opts := make([]cloudsqlconn.Option, 0)
		if config.Config.Bool("db.cloudSql.automaticIamAuthEnabled") {
			opts = append(opts, cloudsqlconn.WithIAMAuthN())
		}
		return pgxv5.RegisterDriver("cloudsql-postgres", opts...)
	} else {
		return nil, nil
	}
}
