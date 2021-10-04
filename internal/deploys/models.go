package deploys

import (
	"fmt"
	"time"

	"github.com/broadinstitute/sherlock/internal/environments"
	"github.com/broadinstitute/sherlock/internal/services"
	"gorm.io/gorm"
)

var (
	// ErrServiceNotFound is a wrapper around gorms failed lookup error specifically
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
