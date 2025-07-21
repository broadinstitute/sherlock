package sherlock

import (
	"fmt"

	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"gorm.io/gorm"
)

type ClusterV3 struct {
	CommonFields
	CiIdentifier     *CiIdentifierV3 `json:"ciIdentifier,omitempty" form:"-"`
	RequiredRoleInfo *RoleV3         `json:"requiredRoleInfo,omitempty" form:"-"`
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
	RequiresSuitability *bool   `json:"requiresSuitability" form:"requiresSuitability"`
	RequiredRole        *string `json:"requiredRole" form:"requiredRole"` // If present, requires membership in the given role for mutations. Set to an empty string to clear.
	HelmfileRef         *string `json:"helmfileRef" form:"helmfileRef" default:"HEAD"`
}

func (c ClusterV3) toModel(db *gorm.DB) (models.Cluster, error) {
	ret := models.Cluster{
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
	if c.RequiredRole != nil && *c.RequiredRole != "" {
		requiredRoleModel, err := roleModelFromSelector(*c.RequiredRole)
		if err != nil {
			return models.Cluster{}, err
		}
		var requiredRole models.Role
		if err = db.Where(requiredRoleModel).Select("id").First(&requiredRole).Error; err != nil {
			return models.Cluster{}, fmt.Errorf("required role '%s' not found: %w", *c.RequiredRole, err)
		} else {
			ret.RequiredRoleID = &requiredRole.ID
		}
	}
	return ret, nil
}

func (c ClusterV3Create) toModel(db *gorm.DB) (models.Cluster, error) {
	return ClusterV3{ClusterV3Create: c}.toModel(db)
}

func (c ClusterV3Edit) toModel(db *gorm.DB) (models.Cluster, error) {
	return ClusterV3Create{ClusterV3Edit: c}.toModel(db)
}

func clusterFromModel(model models.Cluster) ClusterV3 {
	ret := ClusterV3{
		CommonFields:     commonFieldsFromGormModel(model.Model),
		CiIdentifier:     utils.NilOrCall(ciIdentifierFromModel, model.CiIdentifier),
		RequiredRoleInfo: utils.NilOrCall(roleFromModel, model.RequiredRole),
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
	if model.RequiredRole != nil && model.RequiredRole.Name != nil {
		ret.RequiredRole = model.RequiredRole.Name
	} else if model.RequiredRoleID != nil {
		ret.RequiredRole = utils.PointerTo(utils.UintToString(*model.RequiredRoleID))
	} else if substituteEmptyRole := config.Config.String("model.roles.substituteEmptyRequiredRoleWithValue"); substituteEmptyRole != "" {
		ret.RequiredRole = &substituteEmptyRole
	}
	return ret
}
