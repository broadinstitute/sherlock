package deploys

import (
	"fmt"
	"time"

	"github.com/broadinstitute/sherlock/internal/builds"
	"github.com/broadinstitute/sherlock/internal/environments"
	"github.com/broadinstitute/sherlock/internal/services"
	"gorm.io/gorm"
)

var (
	// ErrServiceInstanceNotFound is a wrapper around gorms failed lookup error specifically
	// for failure to find a service instance
	ErrServiceInstanceNotFound = gorm.ErrRecordNotFound
)

type dataStore struct {
	*gorm.DB
}

// ServiceInstance is the model type for interacting with the database
// representation of service instances
type ServiceInstance struct {
	ID            int
	ServiceID     int
	Service       services.Service `gorm:"foreignKey:ServiceID;references:ID"`
	EnvironmentID int
	Environment   environments.Environment `gorm:"foreignKey:EnvironmentID;references:ID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type serviceInstanceStore interface {
	listAll() ([]ServiceInstance, error)
	createNew(environmentID int, serviceID int) (ServiceInstance, error)
	getByEnvironmentAndServiceName(environmentName, serviceName string) (ServiceInstance, error)
}

func newServiceInstanceStore(dbConn *gorm.DB) dataStore {
	return dataStore{dbConn}
}

func (db dataStore) createNew(environmentID, serviceID int) (ServiceInstance, error) {
	newServiceInstance := ServiceInstance{
		ServiceID:     serviceID,
		EnvironmentID: environmentID,
	}

	if err := db.Create(&newServiceInstance).Error; err != nil {
		return ServiceInstance{}, fmt.Errorf("error persisting service instance: %v", err)
	}

	// retrieve the same service instance record back from the db but now with all the
	// associations populated.
	err := db.Preload("Service").
		Preload("Environment").
		First(&newServiceInstance, newServiceInstance.ID).Error

	return newServiceInstance, err
}

func (db dataStore) getByEnvironmentAndServiceName(environmentName, serviceName string) (ServiceInstance, error) {
	var serviceInstance ServiceInstance

	// using gorms struct query features to set the WHERE clause
	queryStruct := ServiceInstance{
		Environment: environments.Environment{
			Name: environmentName,
		},
		Service: services.Service{
			Name: serviceName,
		},
	}

	err := db.Preload("Service").Preload("Environment").Where(&queryStruct).First(&serviceInstance).Error
	return serviceInstance, err
}

func (db dataStore) listAll() ([]ServiceInstance, error) {
	serviceInstances := make([]ServiceInstance, 0)

	err := db.Preload("Service").Preload("Environment").Find(&serviceInstances).Error
	if err != nil {
		return []ServiceInstance{}, fmt.Errorf("error listing service instances: %v", err)
	}

	return serviceInstances, nil
}

// Deploy is the type  defining the database model for a deployment. It is an association
// between a service instance and a build
type Deploy struct {
	ID                int
	ServiceInstanceID int
	ServiceInstance   ServiceInstance `gorm:"foreignKey:ServiceInstanceID;references:ID"`
	BuildID           int
	Build             builds.Build `gorm:"foreignKey:BuildID;references:ID"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type deployStore interface {
	createDeploy(buildID, serviceInstanceID int) (Deploy, error)
	getDeploysByServiceInstance(serviceInstanceID int) ([]Deploy, error)
}

func newDeployStore(dbConn *gorm.DB) dataStore {
	return dataStore{dbConn}
}

func (db dataStore) createDeploy(buildID, serviceInstanceID int) (Deploy, error) {
	newDeploy := Deploy{
		ServiceInstanceID: serviceInstanceID,
		BuildID:           buildID,
	}

	if err := db.Create(&newDeploy).Error; err != nil {
		return Deploy{}, err
	}

	// retrieve the same deploy record back from the db but now with all the
	// associations populated.
	err := db.Preload("ServiceInstance").
		Preload("ServiceInstance.Service").
		Preload("ServiceInstance.Environment").
		Preload("Build").
		Preload("Build.Service").
		First(&newDeploy, newDeploy.ID).Error

	return newDeploy, err
}

func (db dataStore) getDeploysByServiceInstance(serviceInstanceID int) ([]Deploy, error) {
	var deploys []Deploy

	err := db.Preload("ServiceInstance").
		Preload("ServiceInstance.Service").
		Preload("ServiceInstance.Environment").
		Preload("Build").
		Preload("Build.Service").
		Where(&Deploy{ServiceInstanceID: serviceInstanceID}).
		Find(&deploys).
		Error

	return deploys, err
}
