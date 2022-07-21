package v2controllers

import (
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"gorm.io/gorm"
)

type ChartVersion struct {
	ReadableBaseType
	ChartInfo *Chart
	CreatableChartVersion
}

type CreatableChartVersion struct {
	Chart        string
	ChartVersion string
}

func (c CreatableChartVersion) toReadable() ChartVersion {
	return ChartVersion{CreatableChartVersion: c}
}

type ChartVersionController = ImmutableModelController[v2models.ChartVersion, ChartVersion, CreatableChartVersion]

func NewChartVersionController(stores v2models.StoreSet) *ChartVersionController {
	return &ChartVersionController{
		primaryStore:    stores.ChartVersionStore,
		allStores:       stores,
		modelToReadable: modelChartVersionToChartVersion,
		readableToModel: chartVersionToModelChartVersion,
	}
}

func modelChartVersionToChartVersion(model v2models.ChartVersion) *ChartVersion {
	chart := modelChartToChart(model.Chart)
	return &ChartVersion{
		ReadableBaseType: ReadableBaseType{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
		ChartInfo: chart,
		CreatableChartVersion: CreatableChartVersion{
			Chart:        chart.Name,
			ChartVersion: model.ChartVersion,
		},
	}
}

func chartVersionToModelChartVersion(chartVersion ChartVersion, stores v2models.StoreSet) (v2models.ChartVersion, error) {
	chart, err := stores.ChartStore.Get(chartVersion.Chart)
	if err != nil {
		return v2models.ChartVersion{}, err
	}
	return v2models.ChartVersion{
		Model: gorm.Model{
			ID:        chartVersion.ID,
			CreatedAt: chartVersion.CreatedAt,
			UpdatedAt: chartVersion.UpdatedAt,
		},
		ChartID:      chart.ID,
		ChartVersion: chartVersion.ChartVersion,
	}, nil
}
