package v2controllers

import (
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"gorm.io/gorm"
)

type Chart struct {
	ReadableBaseType
	CreatableChart
}

type CreatableChart struct {
	Name string `json:"name" form:"name"`
	EditableChart
}

type EditableChart struct {
	ChartRepo             *string `json:"chartRepo" form:"chartRepo"`
	AppImageGitRepo       *string `json:"appImageGitRepo" form:"appImageGitRepo"`
	AppImageGitMainBranch *string `json:"appImageGitMainBranch" form:"appImageGitMainBranch"`
}

func (c CreatableChart) toReadable() Chart {
	return Chart{CreatableChart: c}
}

func (e EditableChart) toCreatable() CreatableChart {
	return CreatableChart{EditableChart: e}
}

type ChartController = ModelController[v2models.Chart, Chart, CreatableChart, EditableChart]

func NewChartController(stores v2models.StoreSet) *ChartController {
	return &ChartController{
		primaryStore:    stores.ChartStore,
		allStores:       stores,
		modelToReadable: modelChartToChart,
		readableToModel: chartToModelChart,
	}
}

func modelChartToChart(model v2models.Chart) *Chart {
	return &Chart{
		ReadableBaseType: ReadableBaseType{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
		CreatableChart: CreatableChart{
			Name: model.Name,
			EditableChart: EditableChart{
				ChartRepo:             model.ChartRepo,
				AppImageGitRepo:       model.AppImageGitRepo,
				AppImageGitMainBranch: model.AppImageGitMainBranch,
			},
		},
	}
}

func chartToModelChart(chart Chart, _ v2models.StoreSet) (v2models.Chart, error) {
	return v2models.Chart{
		Model: gorm.Model{
			ID:        chart.ID,
			CreatedAt: chart.CreatedAt,
			UpdatedAt: chart.UpdatedAt,
		},
		Name:                  chart.Name,
		ChartRepo:             chart.ChartRepo,
		AppImageGitRepo:       chart.AppImageGitRepo,
		AppImageGitMainBranch: chart.AppImageGitMainBranch,
	}, nil
}
