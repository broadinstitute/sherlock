// ClusterController, public interface for Cluster objects

package controllers

import (
	"errors"
	"fmt"

	"github.com/broadinstitute/sherlock/internal/models"
	"gorm.io/gorm"
)

// ClusterController is the management layer for clusters
type ClusterController struct {
	store models.ClusterStore
	// other derived data
}

// NewController accepts a gorm DB connection and returns a new instance
// of the cluster controller
func NewClusterController(dbConn *gorm.DB) *ClusterController {
	clusterStore := models.NewClusterStore(dbConn)
	return &ClusterController{
		store: clusterStore,
	}
}

// DoesClusterExist is a helper method to check if a cluster with the given name
// already exists in sherlock's data storage
func (clusterController ClusterController) DoesClusterExist(name string) (id int, ok bool) {
	cluster, err := clusterController.GetByName(name)
	if errors.Is(err, models.ErrClusterNotFound) {
		return 0, false
	}
	return cluster.ID, true
}

// CreateNew is the public api on the clusterController for persisting a new service entity to
// the data store
func (clusterController *ClusterController) CreateNew(newCluster models.CreateClusterRequest) (models.Cluster, error) {
	return clusterController.store.CreateNew(newCluster)
}

// ListAll is the public api for listing out all clusters tracked by sherlock
func (clusterController *ClusterController) ListAll() ([]models.Cluster, error) {
	return clusterController.store.ListAll()
}

// GetByName is the public API for looking up a cluster from the data store by name
func (clusterController *ClusterController) GetByName(name string) (models.Cluster, error) {
	return clusterController.store.GetByName(name)
}

// GetByID is the public API for looking up a cluster from the data store by name
func (clusterController *ClusterController) GetByID(id int) (models.Cluster, error) {
	return clusterController.store.GetByID(id)
}

// FindOrCreate will attempt to look an cluster by name and return its ID if successful
// if unsuccessful it will create a new cluster from the provider name and return that id
func (clusterController *ClusterController) FindOrCreate(name string) (int, error) {
	clusterID, exists := clusterController.DoesClusterExist(name)

	if !exists {
		newCluster := models.CreateClusterRequest{Name: name}
		createdCluster, err := clusterController.CreateNew(newCluster)
		if err != nil {
			return 0, fmt.Errorf("error creating cluster %s: %v", name, err)
		}
		clusterID = createdCluster.ID
	}
	return clusterID, nil
}
