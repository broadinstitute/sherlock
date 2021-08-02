package services

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Service is the data structure representing an indvidual applicaiton
type Service struct {
	ID        int    `json:"id,omitempty" db:"id"`
	Name      string `json:"name" db:"name"`
	RepoURL   string `json:"repo_url" db:"repo_url"`
	CreatedAt string `json:"created_at,omitempty" db:"created_at"`
}

// ListAll uses the underlying datastore to retrieve all services
func ListAll(store *sqlx.DB) ([]Service, error) {
	services := make([]Service, 0)

	if err := store.Select(&services, selectAll); err != nil {
		return nil, fmt.Errorf("error selecting all services: %v", err)
	}

	return services, nil
}
