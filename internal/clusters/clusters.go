// ClusterController, public interface for Cluster objects

// Package clusters defines data structure representing
// a cluster instance and methods for interacting with them
// it is left to concrete implementations in package db or others to
// implement these interfaces
package clusters

import (
	"errors"

	"gorm.io/gorm"
)

// ClusterController is the management layer for clusters
type ClusterController struct {
	store clusterStore
}

// NewController accepts a gorm DB connection and returns a new instance
// of the cluster controller
func NewController(dbConn *gorm.DB) *ClusterController {
	clusterStore := newClusterStore(dbConn)
	return &ClusterController{
		store: clusterStore,
	}
}

// DoesClusterExist is a helper method to check if a cluster with the given name
// already exists in sherlock's data storage
func (clusterController ClusterController) DoesClusterExist(name string) (id int, ok bool) {
	cluster, err := clusterController.GetByName(name)
	if errors.Is(err, ErrClusterNotFound) {
		return 0, false
	}
	return cluster.ID, true
}

// CreateNew is the public api on the clusterController for persisting a new service entity to
// the data store
func (clusterController *ClusterController) CreateNew(newCluster CreateClusterRequest) (Cluster, error) {
	return clusterController.store.createNew(newCluster)

}

// ListAll is the public api for listing out all clusters tracked by sherlock
func (clusterController *ClusterController) ListAll() ([]Cluster, error) {
	return clusterController.store.listAll()

}

// GetByName is the public API for looking up a cluster from the data store by name
func (clusterController *ClusterController) GetByName(name string) (Cluster, error) {
	return clusterController.store.getByName(name)

}

// AddEnvironmentByID takes a ClusterObject and associates an existing environment to it.
func (clusterController *ClusterController) AddEnvironmentByID(currentCluster Cluster, environmentID int) (Cluster, error) {
	return clusterController.store.addEnvironmentByID(currentCluster, environmentID)
}
