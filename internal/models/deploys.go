package models

import (
	"time"

	"gorm.io/gorm"
)

var (
	// ErrDeployNotFound is retunrned when unable to find a deployment db record that matches a get query
	ErrDeployNotFound = gorm.ErrRecordNotFound
)

type deployStore struct {
	*gorm.DB
}

// Deploy is the type  defining the database model for a deployment. It is an association
// between a service instance and a build
type Deploy struct {
	ID                int
	ServiceInstanceID int
	ServiceInstance   ServiceInstance `gorm:"foreignKey:ServiceInstanceID;references:ID"`
	BuildID           int
	Build             Build `gorm:"foreignKey:BuildID;references:ID"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type DeployStore interface {
	CreateDeploy(buildID, serviceInstanceID int) (Deploy, error)
	GetDeploysByServiceInstance(serviceInstanceID int) ([]Deploy, error)
	GetMostRecentDeployByServiceInstance(serviceInstanceID int) (Deploy, error)
}

func NewDeployStore(dbConn *gorm.DB) deployStore {
	return deployStore{dbConn}
}

func (db deployStore) CreateDeploy(buildID, serviceInstanceID int) (Deploy, error) {
	newDeploy := Deploy{
		ServiceInstanceID: serviceInstanceID,
		BuildID:           buildID,
	}

	if err := db.Create(&newDeploy).Error; err != nil {
		return Deploy{}, err
	}

	// retrieve the same Deploy record back from the db with all the
	// associations populated
	err := db.Preload("ServiceInstance").
		Preload("ServiceInstance.Service").
		Preload("ServiceInstance.Environment").
		Preload("Build").
		First(&newDeploy).
		Error

	return newDeploy, err
}

func (db deployStore) GetDeploysByServiceInstance(serviceInstanceID int) ([]Deploy, error) {
	var deploys []Deploy

	// TODO: If we ever hit DB bottlenecks this is a likely suspect
	err := db.Preload("ServiceInstance").
		Preload("ServiceInstance.Service").
		Preload("ServiceInstance.Environment").
		Preload("Build").
		Preload("Build.Service").
		Find(&deploys, &Deploy{ServiceInstanceID: serviceInstanceID}).
		Error

	return deploys, err
}

func (db deployStore) GetMostRecentDeployByServiceInstance(serviceInstanceID int) (Deploy, error) {
	var mostRecentDeploy Deploy

	err := db.Preload("Build").
		Order("created_at DESC").
		First(&mostRecentDeploy, &Deploy{ServiceInstanceID: serviceInstanceID}).
		Error

	return mostRecentDeploy, err
}

func (deploy *Deploy) CalculateLeadTimeHours() float64 {
	return deploy.CreatedAt.Sub(deploy.Build.BuiltAt).Hours()
}
