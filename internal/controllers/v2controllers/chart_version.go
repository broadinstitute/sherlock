package v2controllers

import (
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"gorm.io/gorm"
)

type ChartVersion struct {
	ReadableBaseType
	ChartInfo Chart `json:"chartInfo" form:"chartInfo"`
	CreatableChartVersion
}

type CreatableChartVersion struct {
	Chart        string `json:"chart" form:"chart"`               // Required when creating
	ChartVersion string `json:"chartVersion" form:"chartVersion"` // Required when creating
	EditableChartVersion
}

type EditableChartVersion struct{}

//nolint:unused
func (c CreatableChartVersion) toReadable() ChartVersion {
	return ChartVersion{CreatableChartVersion: c}
}

//nolint:unused
func (e EditableChartVersion) toCreatable() CreatableChartVersion {
	return CreatableChartVersion{EditableChartVersion: e}
}

type ChartVersionController = ModelController[v2models.ChartVersion, ChartVersion, CreatableChartVersion, EditableChartVersion]

func newChartVersionController(stores *v2models.StoreSet) *ChartVersionController {
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
		ChartInfo: *chart,
		CreatableChartVersion: CreatableChartVersion{
			Chart:        chart.Name,
			ChartVersion: model.ChartVersion,
		},
	}
}

func chartVersionToModelChartVersion(chartVersion ChartVersion, stores *v2models.StoreSet) (v2models.ChartVersion, error) {
	var chartID uint
	if chartVersion.Chart != "" {
		chart, err := stores.ChartStore.Get(chartVersion.Chart)
		if err != nil {
			return v2models.ChartVersion{}, err
		}
		chartID = chart.ID
	}
	return v2models.ChartVersion{
		Model: gorm.Model{
			ID:        chartVersion.ID,
			CreatedAt: chartVersion.CreatedAt,
			UpdatedAt: chartVersion.UpdatedAt,
		},
		ChartID:      chartID,
		ChartVersion: chartVersion.ChartVersion,
	}, nil
}
