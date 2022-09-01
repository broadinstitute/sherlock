package v2controllers

import (
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"gorm.io/gorm"
)

type AppVersion struct {
	ReadableBaseType
	ChartInfo Chart `json:"chartInfo" form:"-"`
	CreatableAppVersion
}

type CreatableAppVersion struct {
	Chart      string `json:"chart" form:"chart"`           // Required when creating
	AppVersion string `json:"appVersion" form:"appVersion"` // Required when creating
	GitCommit  string `json:"gitCommit" form:"gitCommit"`
	GitBranch  string `json:"gitBranch" form:"gitBranch"`
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

func modelAppVersionToAppVersion(model v2models.AppVersion) *AppVersion {
	chart := modelChartToChart(model.Chart)
	return &AppVersion{
		ReadableBaseType: ReadableBaseType{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
		ChartInfo: *chart,
		CreatableAppVersion: CreatableAppVersion{
			Chart:      chart.Name,
			AppVersion: model.AppVersion,
			GitCommit:  model.GitCommit,
			GitBranch:  model.GitBranch,
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
	return v2models.AppVersion{
		Model: gorm.Model{
			ID:        appVersion.ID,
			CreatedAt: appVersion.CreatedAt,
			UpdatedAt: appVersion.UpdatedAt,
		},
		ChartID:    chartID,
		AppVersion: appVersion.AppVersion,
		GitCommit:  appVersion.GitCommit,
		GitBranch:  appVersion.GitBranch,
	}, nil
}
