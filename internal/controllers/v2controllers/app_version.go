package v2controllers

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type AppVersion struct {
	ReadableBaseType
	ChartInfo            *Chart      `json:"chartInfo,omitempty"  form:"-"`
	ParentAppVersionInfo *AppVersion `json:"parentAppVersionInfo,omitempty" swaggertype:"object" form:"-"`
	CreatableAppVersion
}

type CreatableAppVersion struct {
	Chart            string `json:"chart" form:"chart"`           // Required when creating
	AppVersion       string `json:"appVersion" form:"appVersion"` // Required when creating
	GitCommit        string `json:"gitCommit" form:"gitCommit"`
	GitBranch        string `json:"gitBranch" form:"gitBranch"`
	ParentAppVersion string `json:"parentAppVersion" form:"parentAppVersion"`
	EditableAppVersion
}

type EditableAppVersion struct{}

//nolint:unused
func (c CreatableAppVersion) toReadable() AppVersion {
	return AppVersion{CreatableAppVersion: c}
}

//nolint:unused
func (e EditableAppVersion) toCreatable() CreatableAppVersion {
	return CreatableAppVersion{EditableAppVersion: e}
}

type AppVersionController = ModelController[v2models.AppVersion, AppVersion, CreatableAppVersion, EditableAppVersion]

func newAppVersionController(stores *v2models.StoreSet) *AppVersionController {
	return &AppVersionController{
		primaryStore:    stores.AppVersionStore,
		allStores:       stores,
		modelToReadable: modelAppVersionToAppVersion,
		readableToModel: appVersionToModelAppVersion,
	}
}

func modelAppVersionToAppVersion(model *v2models.AppVersion) *AppVersion {
	if model == nil {
		return nil
	}

	var chartName string
	chart := modelChartToChart(model.Chart)
	if chart != nil {
		chartName = chart.Name
	}

	var parentAppVersionSelector string
	parentAppVersion := modelAppVersionToAppVersion(model.ParentAppVersion)
	if parentAppVersion != nil {
		// The parent's associations might not be loaded, so we can't safely get the chart name of the parent, but
		// we know that the parent's chart name is the same as ours.
		parentAppVersionSelector = fmt.Sprintf("%s/%s", chartName, parentAppVersion.AppVersion)
	}

	return &AppVersion{
		ReadableBaseType: ReadableBaseType{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
		ChartInfo:            chart,
		ParentAppVersionInfo: parentAppVersion,
		CreatableAppVersion: CreatableAppVersion{
			Chart:            chartName,
			AppVersion:       model.AppVersion,
			GitCommit:        model.GitCommit,
			GitBranch:        model.GitBranch,
			ParentAppVersion: parentAppVersionSelector,
		},
	}
}

func appVersionToModelAppVersion(appVersion AppVersion, stores *v2models.StoreSet) (v2models.AppVersion, error) {
	var chartID uint
	if appVersion.Chart != "" {
		chart, err := stores.ChartStore.Get(appVersion.Chart)
		if err != nil {
			return v2models.AppVersion{}, err
		}
		chartID = chart.ID
	}
	var parentAppVersionID *uint
	if appVersion.ParentAppVersion != "" {
		parentAppVersion, err := stores.AppVersionStore.Get(appVersion.ParentAppVersion)
		if err != nil {
			log.Debug().Msgf("while handling %T was given parent %s that didn't have a match, ignoring: %v", appVersion, appVersion.ParentAppVersion, err)
		} else {
			if chartID != 0 && parentAppVersion.ChartID != chartID {
				return v2models.AppVersion{}, fmt.Errorf("(%s) given parent matches a different chart (%s, ID %d) than this one does (%s, ID %d)", errors.BadRequest, parentAppVersion.Chart.Name, parentAppVersion.ChartID, appVersion.Chart, chartID)
			}
			parentAppVersionID = &parentAppVersion.ID
		}
	}
	return v2models.AppVersion{
		Model: gorm.Model{
			ID:        appVersion.ID,
			CreatedAt: appVersion.CreatedAt,
			UpdatedAt: appVersion.UpdatedAt,
		},
		ChartID:            chartID,
		AppVersion:         appVersion.AppVersion,
		GitCommit:          appVersion.GitCommit,
		GitBranch:          appVersion.GitBranch,
		ParentAppVersionID: parentAppVersionID,
	}, nil
}
