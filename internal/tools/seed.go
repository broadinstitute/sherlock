package tools

import (
	"fmt"

	"github.com/broadinstitute/sherlock/internal/services"
	"github.com/jmoiron/sqlx"
)

// SeedServices is a test utility that will populate a database with a predetermined list of "services"
// to be used for running integration tests against a real database
func SeedServices(db *sqlx.DB) ([]services.Service, error) {
	services := []services.Service{
		{
			Name:    "cromwell",
			RepoURL: "https://github.com/broadinstitute/cromwell",
		},
		{
			Name:    "leonardo",
			RepoURL: "https://github.com/DataBiosphere/leonardo",
		},
		{
			Name:    "workspacemanager",
			RepoURL: "https://github.com/DataBiosphere/terra-workspace-manager",
		},
	}

	statement, err := db.Prepare("INSERT INTO services (name, repo_url) VALUES ($1, $2) RETURNING id, created_at;")
	if err != nil {
		return nil, fmt.Errorf("error preparing statement %v", err)
	}
	defer statement.Close()

	for i, service := range services {
		row := statement.QueryRow(service.Name, service.RepoURL)

		if err = row.Scan(&services[i].ID, &services[i].CreatedAt); err != nil {
			return nil, fmt.Errorf("error seeding data %v", err)
		}

	}
	return services, nil
}

// Truncate cleans up tables after integration tests
func Truncate(db *sqlx.DB) error {
	statement := "TRUNCATE TABLE services, builds, environments, service_instances, deploys"

	if _, err := db.Exec(statement); err != nil {
		return fmt.Errorf("error truncating test db %v", err)
	}

	return nil
}
