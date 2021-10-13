package deploys

import (
	"fmt"
	"time"

	"github.com/broadinstitute/sherlock/internal/builds"
	"github.com/broadinstitute/sherlock/internal/environments"
	"github.com/broadinstitute/sherlock/internal/models"
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
	ClusterID     int                      `gorm:"default:null"`
	Cluster       models.Cluster           `gorm:"foreignKey:ClusterID;references:ID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type serviceInstanceStore interface {
	listAll() ([]ServiceInstance, error)
	createNew(clusterID, environmentID, serviceID int) (ServiceInstance, error)
	getByEnvironmentAndServiceName(environmentName, serviceName string) (ServiceInstance, error)
	getByEnvironmentAndServiceID(environmentID, serviceID int) (ServiceInstance, error)
}

func newServiceInstanceStore(dbConn *gorm.DB) dataStore {
	return dataStore{dbConn}
}

func (db dataStore) createNew(clusterID, environmentID, serviceID int) (ServiceInstance, error) {
	newServiceInstance := ServiceInstance{
		ServiceID:     serviceID,
		EnvironmentID: environmentID,
		ClusterID:     clusterID,
	}

	if err := db.Create(&newServiceInstance).Error; err != nil {
		return ServiceInstance{}, fmt.Errorf("error persisting service instance: %v", err)
	}

	return newServiceInstance, nil
}

func (db dataStore) getByEnvironmentAndServiceID(environmentID, serviceID int) (ServiceInstance, error) {
	var serviceInstance ServiceInstance

	err := db.Preload("Service").Preload("Environment").First(&serviceInstance, &ServiceInstance{ServiceID: serviceID, EnvironmentID: environmentID}).Error
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

func (db dataStore) getDeploysByServiceInstance(serviceInstanceID int) ([]Deploy, error) {
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
