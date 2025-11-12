package sherlock

import (
	"fmt"

	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"gorm.io/gorm"
)

type DatabaseInstanceV3 struct {
	CommonFields
	ChartReleaseInfo *ChartReleaseV3 `json:"chartReleaseInfo,omitempty" form:"-"`
	DatabaseInstanceV3Create
}

type DatabaseInstanceV3Create struct {
	DatabaseInstanceV3Edit
	ChartRelease string `json:"chartRelease" form:"chartRelease"` // Required when creating
}

type DatabaseInstanceV3Edit struct {
	Platform        *string `json:"platform" form:"platform" default:"kubernetes"` // 'google', 'azure', or default 'kubernetes'
	GoogleProject   *string `json:"googleProject" form:"googleProject"`            // Required if platform is 'google'
	InstanceName    *string `json:"instanceName" form:"instanceName"`              // Required if platform is 'google' or 'azure'
	DefaultDatabase *string `json:"defaultDatabase" form:"defaultDatabase"`        // When creating, defaults to the chart name
}

func (d DatabaseInstanceV3) toModel(db *gorm.DB) (models.DatabaseInstance, error) {
	ret := models.DatabaseInstance{
		Model:           d.toGormModel(),
		Platform:        d.Platform,
		GoogleProject:   d.GoogleProject,
		InstanceName:    d.InstanceName,
		DefaultDatabase: d.DefaultDatabase,
	}
	if d.ChartRelease != "" {
		chartReleaseModel, err := chartReleaseModelFromSelector(db, d.ChartRelease)
		if err != nil {
			return models.DatabaseInstance{}, err
		}
		var chartRelease models.ChartRelease
		if err = db.Where(&chartReleaseModel).Select("id").First(&chartRelease).Error; err != nil {
			return models.DatabaseInstance{}, fmt.Errorf("unable to get chart release '%s': %w", d.ChartRelease, err)
		} else {
			ret.ChartReleaseID = chartRelease.ID
		}
	}
	return ret, nil
}

func (d DatabaseInstanceV3Create) toModel(db *gorm.DB) (models.DatabaseInstance, error) {
	return DatabaseInstanceV3{DatabaseInstanceV3Create: d}.toModel(db)
}

func (d DatabaseInstanceV3Edit) toModel(db *gorm.DB) (models.DatabaseInstance, error) {
	return DatabaseInstanceV3Create{DatabaseInstanceV3Edit: d}.toModel(db)
}

func databaseInstanceFromModel(model models.DatabaseInstance) DatabaseInstanceV3 {
	ret := DatabaseInstanceV3{
		CommonFields:     commonFieldsFromGormModel(model.Model),
		ChartReleaseInfo: utils.NilOrCall(chartReleaseFromModel, model.ChartRelease),
		DatabaseInstanceV3Create: DatabaseInstanceV3Create{
			DatabaseInstanceV3Edit: DatabaseInstanceV3Edit{
				Platform:        model.Platform,
				GoogleProject:   model.GoogleProject,
				InstanceName:    model.InstanceName,
				DefaultDatabase: model.DefaultDatabase,
			},
		},
	}
	if model.ChartRelease != nil && model.ChartRelease.Name != "" {
		ret.ChartRelease = model.ChartRelease.Name
	} else if model.ChartReleaseID != 0 {
		ret.ChartRelease = utils.UintToString(model.ChartReleaseID)
	}
	return ret
}
