// db-level managemenet for Cluster Struct
// APIs should not interact with this file and should user ClusterController instead
// all gorm related methods should live in this file.

package v1models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// ErrClusterNotFound is the error to represent a failed lookup of a cluster db record
var ErrClusterNotFound = gorm.ErrRecordNotFound

// clusterStore is a wrapper around a gorm postgres client
// which can be used as an implementor of the ClusterStore interface
type clusterStore struct {
	*gorm.DB
}

// Cluster is the data structure that models a persisted to a database via gorm
type Cluster struct {
	ID            int    `gorm:"primaryKey;uniqueIndex"`
	Name          string `gorm:"not null;default:null"`
	GoogleProject string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// clusterStore is the interface defining allowed db actions for Cluster
type ClusterStore interface {
	ListAll() ([]Cluster, error)
	CreateNew(CreateClusterRequest) (Cluster, error)
	GetByID(int) (Cluster, error)
	GetByName(string) (Cluster, error)
}

// creates a db connection via gorm
func NewClusterStore(dbconn *gorm.DB) clusterStore {
	return clusterStore{dbconn}
}

// CreateClusterRequest struct defines the data required to create a new cluster in db
type CreateClusterRequest struct {
	Name string `json:"name" binding:"required"`
}

// creates a cluster entity object to be persisted with the database from a
// request to create a cluster.
func (createClusterRequest CreateClusterRequest) ClusterReq() Cluster {
	return Cluster{
		Name: createClusterRequest.Name,
	}
}

//
// db methods
//

// Returns ALL Clusters in the db
func (db clusterStore) ListAll() ([]Cluster, error) {
	clusters := []Cluster{}

	err := db.Find(&clusters).Error
	if err != nil {
		return []Cluster{}, fmt.Errorf("error retreiving clusters: %v", err)
	}

	return clusters, nil
}

// Saves an Cluster object to the db, returns the object if successful, nil otherwise
func (db clusterStore) CreateNew(newClusterReq CreateClusterRequest) (Cluster, error) {
	newCluster := newClusterReq.ClusterReq()

	if err := db.Create(&newCluster).Error; err != nil {
		return Cluster{}, fmt.Errorf("error saving to database: %v", err)
	}
	return newCluster, nil
}

// Get is used to retrieve a specific cluster entity from a postgres database using
// id (primary key) as the lookup mechanism
func (db clusterStore) GetByID(id int) (Cluster, error) {
	cluster := Cluster{}

	if err := db.First(&cluster, id).Error; err != nil {
		return Cluster{}, err
	}
	return cluster, nil
}

// get an Cluster by name column
func (db clusterStore) GetByName(name string) (Cluster, error) {
	cluster := Cluster{}

	if err := db.Where(&Cluster{Name: name}).First(&cluster).Error; err != nil {
		return Cluster{}, err
	}
	return cluster, nil
}
