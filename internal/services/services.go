package services

import (
	"fmt"
	"time"

	"gorm.io/gorm"
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
func ListAll(store *gorm.DB) ([]Service, error) {
	services := []Service{}

	store.Find(&services)
	if store.Error != nil {
		return nil, fmt.Errorf("Error retriving services: %v", store.Error)
	}

	return services, nil
}
