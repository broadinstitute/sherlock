// Package services defines data structure representing
// a service instance and methods for interacting with them
// it is left to concrete implementations in package db or others to
// implement these interfaces
package services

import "gorm.io/gorm"

// // Service is the data structure representing an indvidual applicaiton
// type Service struct {
// 	ID        int       `json:"id,omitempty"`
// 	Name      string    `json:"name" binding:"required"`
// 	RepoURL   string    `json:"repo_url" binding:"required"`
// 	CreatedAt time.Time `json:"created_at,omitempty"`
// 	UpdatedAt time.Time `json:"updated_at,omitempty"`
// }

// ServiceController is the  interface used to model operations relating to services in the backend datastore
type ServiceController struct {
	store serviceStore
}

// NewController accepts a gorm DB connection and returns a new instance
// of the service controller
func NewController(dbConn *gorm.DB) *ServiceController {
	serviceStore := newServiceStore(dbConn)
	return &ServiceController{
		store: serviceStore,
	}
}
