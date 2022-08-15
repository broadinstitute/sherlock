package v2controllers

import (
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"gorm.io/gorm"
)

// Cluster
// @description The full set of Cluster fields that can be read or used for filtering queries
type Cluster struct {
	ReadableBaseType
	CreatableCluster
}

// CreatableCluster
// @description The subset of Cluster fields that can be set upon creation
type CreatableCluster struct {
	Name              string `json:"name" validate:"required" form:"name"` // Required when creating
	Provider          string `json:"provider" form:"provider" enums:"google,azure" default:"google"`
	GoogleProject     string `json:"googleProject" form:"googleProject"`         // Required when creating if provider is 'google'
	AzureSubscription string `json:"azureSubscription" form:"azureSubscription"` // Required when creating if providers is 'azure'
	EditableCluster
}

// EditableCluster
// @description The subset of Cluster fields that can be edited after creation
type EditableCluster struct {
	Base                *string `json:"base"  form:"base"`      // Required when creating
	Address             *string `json:"address" form:"address"` // Required when creating
	RequiresSuitability *bool   `json:"requiresSuitability" form:"requiresSuitability" default:"false"`
}

//nolint:unused
func (c CreatableCluster) toReadable() Cluster {
	return Cluster{CreatableCluster: c}
}

//nolint:unused
func (e EditableCluster) toCreatable() CreatableCluster {
	return CreatableCluster{EditableCluster: e}
}

type ClusterController = ModelController[v2models.Cluster, Cluster, CreatableCluster, EditableCluster]

func newClusterController(stores *v2models.StoreSet) *ClusterController {
	return &ClusterController{
		primaryStore:    stores.ClusterStore,
		allStores:       stores,
		modelToReadable: modelClusterToCluster,
		readableToModel: clusterToModelCluster,
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

func clusterToModelCluster(cluster Cluster, _ *v2models.StoreSet) (v2models.Cluster, error) {
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
