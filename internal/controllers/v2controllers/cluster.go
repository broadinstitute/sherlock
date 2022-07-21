package v2controllers

import (
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"gorm.io/gorm"
)

// Cluster
// @description  The full set of Cluster fields that can be read or used for filtering queries
type Cluster struct {
	ReadableBaseType
	CreatableCluster
}

// CreatableCluster
// @description  The subset of Cluster fields that can be set upon creation
type CreatableCluster struct {
	Name              string `json:"name" validate:"required"`
	Provider          string `json:"provider" enums:"google,azure" validate:"required"`
	GoogleProject     string `json:"googleProject"`
	AzureSubscription string `json:"azureSubscription"`
	EditableCluster
}

// EditableCluster
// @description  The subset of Cluster fields that can be edited after creation
type EditableCluster struct {
	Base                *string `json:"base" validate:"required"`
	Address             *string `json:"address" validate:"required"`
	RequiresSuitability *bool   `json:"requiresSuitability" default:"false"`
}

func (c CreatableCluster) toReadable() Cluster {
	return Cluster{CreatableCluster: c}
}

func (e EditableCluster) toCreatable() CreatableCluster {
	return CreatableCluster{EditableCluster: e}
}

type ClusterController = MutableModelController[v2models.Cluster, Cluster, CreatableCluster, EditableCluster]

func NewClusterController(stores v2models.StoreSet) *ClusterController {
	return &ClusterController{
		ImmutableModelController[v2models.Cluster, Cluster, CreatableCluster]{
			primaryStore:    stores.ClusterStore,
			allStores:       stores,
			modelToReadable: modelClusterToCluster,
			readableToModel: clusterToModelCluster,
		},
	}
}

func modelClusterToCluster(model v2models.Cluster) *Cluster {
	return &Cluster{
		ReadableBaseType: ReadableBaseType{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
		CreatableCluster: CreatableCluster{
			Name:              model.Name,
			Provider:          model.Provider,
			GoogleProject:     model.GoogleProject,
			AzureSubscription: model.AzureSubscription,
			EditableCluster: EditableCluster{
				Base:                model.Base,
				Address:             model.Address,
				RequiresSuitability: model.RequiresSuitability,
			},
		},
	}
}

func clusterToModelCluster(cluster Cluster, _ v2models.StoreSet) (v2models.Cluster, error) {
	return v2models.Cluster{
		Model: gorm.Model{
			ID:        cluster.ID,
			CreatedAt: cluster.CreatedAt,
			UpdatedAt: cluster.UpdatedAt,
		},
		Name:                cluster.Name,
		Provider:            cluster.Provider,
		GoogleProject:       cluster.GoogleProject,
		AzureSubscription:   cluster.AzureSubscription,
		Base:                cluster.Base,
		Address:             cluster.Address,
		RequiresSuitability: cluster.RequiresSuitability,
	}, nil
}
