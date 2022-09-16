package v2controllers

import (
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"gorm.io/gorm"
)

type Chart struct {
	ReadableBaseType
	CreatableChart
}

type CreatableChart struct {
	Name string `json:"name" form:"name"` // Required when creating
	EditableChart
}

type EditableChart struct {
	ChartRepo             *string `json:"chartRepo" form:"chartRepo" default:"terra-helm"`
	AppImageGitRepo       *string `json:"appImageGitRepo" form:"appImageGitRepo"`
	AppImageGitMainBranch *string `json:"appImageGitMainBranch" form:"appImageGitMainBranch"`
	ChartExposesEndpoint  *bool   `json:"chartExposesEndpoint" form:"chartExposesEndpoint" default:"false"` // Indicates if the default subdomain, protocol, and port fields are relevant for this chart
	DefaultSubdomain      *string `json:"defaultSubdomain" form:"defaultSubdomain"`                         // When creating, will default to the name of the chart
	DefaultProtocol       *string `json:"defaultProtocol" form:"defaultProtocol" default:"https"`
	DefaultPort           *uint   `json:"defaultPort" form:"defaultPort" default:"443"`
}

//nolint:unused
func (c CreatableChart) toReadable() Chart {
	return Chart{CreatableChart: c}
}

//nolint:unused
func (e EditableChart) toCreatable() CreatableChart {
	return CreatableChart{EditableChart: e}
}

type ChartController = ModelController[v2models.Chart, Chart, CreatableChart, EditableChart]

func newChartController(stores *v2models.StoreSet) *ChartController {
	return &ChartController{
		primaryStore:       stores.ChartStore,
		allStores:          stores,
		modelToReadable:    modelChartToChart,
		readableToModel:    chartToModelChart,
		setDynamicDefaults: setChartDynamicDefaults,
	}
}

func modelChartToChart(model *v2models.Chart) *Chart {
	if model == nil {
		return nil
	}

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
				ChartExposesEndpoint:  model.ChartExposesEndpoint,
				DefaultSubdomain:      model.DefaultSubdomain,
				DefaultProtocol:       model.DefaultProtocol,
				DefaultPort:           model.DefaultPort,
			},
		},
	}
}

func chartToModelChart(chart Chart, _ *v2models.StoreSet) (v2models.Chart, error) {
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
		ChartExposesEndpoint:  chart.ChartExposesEndpoint,
		DefaultSubdomain:      chart.DefaultSubdomain,
		DefaultProtocol:       chart.DefaultProtocol,
		DefaultPort:           chart.DefaultPort,
	}, nil
}

func setChartDynamicDefaults(chart *Chart, _ *v2models.StoreSet, _ *auth.User) error {
	if chart.DefaultSubdomain == nil {
		chart.DefaultSubdomain = &chart.Name
	}
	return nil
}
