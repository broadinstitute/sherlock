package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
)

type ClusterV3 struct {
	CommonFields
	CiIdentifier *CiIdentifierV3 `json:"ciIdentifier,omitempty" form:"-"`
	ClusterV3Create
}

type ClusterV3Create struct {
	Name              string `json:"name" form:"name"` // Required when creating
	Provider          string `json:"provider" form:"provider" enums:"google,azure" default:"google"`
	GoogleProject     string `json:"googleProject" form:"googleProject"`         // Required when creating if provider is 'google'
	AzureSubscription string `json:"azureSubscription" form:"azureSubscription"` // Required when creating if provider is 'azure'
	Location          string `json:"location" form:"location" default:"us-central1-a"`
	ClusterV3Edit
}

type ClusterV3Edit struct {
	Base                *string `json:"base"  form:"base"`      // Required when creating
	Address             *string `json:"address" form:"address"` // Required when creating
	RequiresSuitability *bool   `json:"requiresSuitability" form:"requiresSuitability" default:"false"`
	HelmfileRef         *string `json:"helmfileRef" form:"helmfileRef" default:"HEAD"`
}

func (c ClusterV3) toModel() models.Cluster {
	return models.Cluster{
		Model:               c.toGormModel(),
		Name:                c.Name,
		Provider:            c.Provider,
		GoogleProject:       c.GoogleProject,
		AzureSubscription:   c.AzureSubscription,
		Location:            c.Location,
		Base:                c.Base,
		Address:             c.Address,
		RequiresSuitability: c.RequiresSuitability,
		HelmfileRef:         c.HelmfileRef,
	}
}

func (c ClusterV3Create) toModel() models.Cluster {
	return ClusterV3{ClusterV3Create: c}.toModel()
}

func (c ClusterV3Edit) toModel() models.Cluster {
	return ClusterV3Create{ClusterV3Edit: c}.toModel()
}

func clusterFromModel(model models.Cluster) ClusterV3 {
	var ciIdentifier *CiIdentifierV3
	if model.CiIdentifier != nil {
		ciIdentifier = utils.PointerTo(ciIdentifierFromModel(*model.CiIdentifier))
	}
	return ClusterV3{
		CommonFields: commonFieldsFromGormModel(model.Model),
		CiIdentifier: ciIdentifier,
		ClusterV3Create: ClusterV3Create{
			Name:              model.Name,
			Provider:          model.Provider,
			GoogleProject:     model.GoogleProject,
			AzureSubscription: model.AzureSubscription,
			Location:          model.Location,
			ClusterV3Edit: ClusterV3Edit{
				Base:                model.Base,
				Address:             model.Address,
				RequiresSuitability: model.RequiresSuitability,
				HelmfileRef:         model.HelmfileRef,
			},
		},
	}
}
