package services

import (
	"fmt"
	"time"

	"github.com/broadinstitute/sherlock/internal/db"
)

// Service is the data structure representing an indvidual applicaiton
type Service struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name"`
	RepoURL   string    `json:"repo_url"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// ListAll uses the underlying datastore to retrieve all services
func ListAll(repository *db.Repository) ([]Service, error) {
	services := []Service{}

	err := repository.DB.Find(&services).Error
	if err != nil {
		return nil, fmt.Errorf("Error retriving services: %v", err)
	}

	return services, nil
}
