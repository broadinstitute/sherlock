package v2controllers

import (
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"gorm.io/gorm"
)

// Cluster
//
//	@description	The full set of Cluster fields that can be read or used for filtering queries
type Cluster struct {
	ReadableBaseType
	CreatableCluster
}

// CreatableCluster
//
//	@description	The subset of Cluster fields that can be set upon creation
type CreatableCluster struct {
	Name              string `json:"name" form:"name"` // Required when creating
	Provider          string `json:"provider" form:"provider" enums:"google,azure" default:"google"`
	GoogleProject     string `json:"googleProject" form:"googleProject"`         // Required when creating if provider is 'google'
	AzureSubscription string `json:"azureSubscription" form:"azureSubscription"` // Required when creating if providers is 'azure'
	Location          string `json:"location" form:"location" default:"us-central1-a"`
	EditableCluster
}

// EditableCluster
//
//	@description	The subset of Cluster fields that can be edited after creation
type EditableCluster struct {
	Base                *string `json:"base"  form:"base"`      // Required when creating
	Address             *string `json:"address" form:"address"` // Required when creating
	RequiresSuitability *bool   `json:"requiresSuitability" form:"requiresSuitability" default:"false"`
	HelmfileRef         *string `json:"helmfileRef" form:"helmfileRef" default:"HEAD"`
}

//nolint:unused
func (c Cluster) toModel(_ *v2models.StoreSet) (v2models.Cluster, error) {
	return v2models.Cluster{
		Model: gorm.Model{
			ID:        c.ID,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		},
		Name:                c.Name,
		Provider:            c.Provider,
		GoogleProject:       c.GoogleProject,
		AzureSubscription:   c.AzureSubscription,
		Location:            c.Location,
		Base:                c.Base,
		Address:             c.Address,
		RequiresSuitability: c.RequiresSuitability,
		HelmfileRef:         c.HelmfileRef,
	}, nil
}

//nolint:unused
func (c CreatableCluster) toModel(storeSet *v2models.StoreSet) (v2models.Cluster, error) {
	return Cluster{CreatableCluster: c}.toModel(storeSet)
}

//nolint:unused
func (c EditableCluster) toModel(storeSet *v2models.StoreSet) (v2models.Cluster, error) {
	return CreatableCluster{EditableCluster: c}.toModel(storeSet)
}

type ClusterController = ModelController[v2models.Cluster, Cluster, CreatableCluster, EditableCluster]

func newClusterController(stores *v2models.StoreSet) *ClusterController {
	return &ClusterController{
		primaryStore:    stores.ClusterStore,
		allStores:       stores,
		modelToReadable: modelClusterToCluster,
	}
}

func modelClusterToCluster(model *v2models.Cluster) *Cluster {
	if model == nil {
		return nil
	}

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
			Location:          model.Location,
			EditableCluster: EditableCluster{
				Base:                model.Base,
				Address:             model.Address,
				RequiresSuitability: model.RequiresSuitability,
				HelmfileRef:         model.HelmfileRef,
			},
		},
	}
}
