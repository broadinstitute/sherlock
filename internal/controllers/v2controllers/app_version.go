package v2controllers

import (
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"gorm.io/gorm"
)

type AppVersion struct {
	ReadableBaseType
	ChartInfo *Chart
	CreatableAppVersion
}

type CreatableAppVersion struct {
	Chart      string
	AppVersion string
	GitCommit  string
	GitBranch  string
}

func (c CreatableAppVersion) toReadable() AppVersion {
	return AppVersion{CreatableAppVersion: c}
}

type AppVersionController = ImmutableModelController[v2models.AppVersion, AppVersion, CreatableAppVersion]

func NewAppVersionController(stores v2models.StoreSet) *AppVersionController {
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
		ChartInfo: chart,
		CreatableAppVersion: CreatableAppVersion{
			Chart:      chart.Name,
			AppVersion: model.AppVersion,
			GitCommit:  model.GitCommit,
			GitBranch:  model.GitBranch,
		},
	}
}

func appVersionToModelAppVersion(appVersion AppVersion, stores v2models.StoreSet) (v2models.AppVersion, error) {
	chart, err := stores.ChartStore.Get(appVersion.Chart)
	if err != nil {
		return v2models.AppVersion{}, err
	}
	return v2models.AppVersion{
		Model: gorm.Model{
			ID:        appVersion.ID,
			CreatedAt: appVersion.CreatedAt,
			UpdatedAt: appVersion.UpdatedAt,
		},
		ChartID:    chart.ID,
		AppVersion: appVersion.AppVersion,
		GitCommit:  appVersion.GitCommit,
		GitBranch:  appVersion.GitBranch,
	}, nil
}
