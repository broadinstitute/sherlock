// Package services defines data structure representing
// a service instance and methods for interacting with them
// it is left to concrete implementations in package db or others to
// implement these interfaces
package services

import (
	"time"
)

// Service is the data structure representing an indvidual applicaiton
type Service struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name" binding:"required"`
	RepoURL   string    `json:"repo_url" binding:"required"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// ServiceModel is the  interface used to model operations relating to services in the backend datastore
type ServiceModel interface {
	ListAll() ([]Service, error)
	Create(*Service) (*Service, error)
	Get(string) (*Service, error)
}
