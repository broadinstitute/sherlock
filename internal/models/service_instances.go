package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

var (
	// ErrServiceInstanceNotFound is a wrapper around gorms failed lookup error specifically
	// for failure to find a service instance
	ErrServiceInstanceNotFound = gorm.ErrRecordNotFound
)

type serviceInstanceStore struct {
	*gorm.DB
}

// ServiceInstance is the model type for interacting with the database
// representation of service instances
type ServiceInstance struct {
	ID            int
	ServiceID     int
	Service       Service `gorm:"foreignKey:ServiceID;references:ID"`
	EnvironmentID int
	Environment   Environment `gorm:"foreignKey:EnvironmentID;references:ID"`
	ClusterID     int         `gorm:"default:null"`
	Cluster       Cluster     `gorm:"foreignKey:ClusterID;references:ID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type ServiceInstanceStore interface {
	ListAll() ([]ServiceInstance, error)
	CreateNew(clusterID, environmentID, serviceID int) (ServiceInstance, error)
	GetByEnvironmentAndServiceID(environmentID, serviceID int) (ServiceInstance, error)
	Reload(serviceInstance ServiceInstance, reloadCluster bool, reloadEnvironment bool, reloadService bool) (ServiceInstance, error)
}

func NewServiceInstanceStore(dbConn *gorm.DB) serviceInstanceStore {
	return serviceInstanceStore{dbConn}
}

func (db serviceInstanceStore) CreateNew(clusterID, environmentID, serviceID int) (ServiceInstance, error) {
	newServiceInstance := ServiceInstance{
		ServiceID:     serviceID,
		EnvironmentID: environmentID,
		ClusterID:     clusterID,
	}

	if err := db.Create(&newServiceInstance).Error; err != nil {
		return ServiceInstance{}, fmt.Errorf("error persisting service instance: %v", err)
	}

	db.Preload("Service").Preload("Environment").Preload("Cluster").First(&newServiceInstance)

	return newServiceInstance, nil
}

// reload a ServiceInstance from database, optionally grab linked data objects.
func (db serviceInstanceStore) Reload(serviceInstance ServiceInstance, reloadCluster bool, reloadEnvironment bool, reloadService bool) (ServiceInstance, error) {
	if reloadCluster {
		db.Preload("Cluster").Find(&serviceInstance)
	}

	if reloadEnvironment {
		db.Preload("Environment").Find(&serviceInstance)
	}

	if reloadService {
		db.Preload("Service").Find(&serviceInstance)
	}

	return serviceInstance, nil
}

func (db serviceInstanceStore) GetByEnvironmentAndServiceID(environmentID, serviceID int) (ServiceInstance, error) {
	var serviceInstance ServiceInstance

	err := db.
		Preload("Service").
		Preload("Environment").
		Preload("Cluster").
		Where("service_id = ? AND environment_id = ?", serviceID, environmentID).
		First(&serviceInstance).
		Error
	return serviceInstance, err
}

func (db serviceInstanceStore) ListAll() ([]ServiceInstance, error) {
	serviceInstances := make([]ServiceInstance, 0)

	err := db.Preload("Service").Preload("Environment").Find(&serviceInstances).Error
	if err != nil {
		return []ServiceInstance{}, fmt.Errorf("error listing service instances: %v", err)
	}

	return serviceInstances, nil
}
