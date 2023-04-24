package v2controllers

import (
	"github.com/broadinstitute/sherlock/internal/auth/auth_models"
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
	LegacyConfigsEnabled  *bool   `json:"legacyConfigsEnabled" form:"legacyConfigsEnbled" default:"false"`  // Indicates whether a chart requires config rendering from firecloud-develop
	DefaultSubdomain      *string `json:"defaultSubdomain" form:"defaultSubdomain"`                         // When creating, will default to the name of the chart
	DefaultProtocol       *string `json:"defaultProtocol" form:"defaultProtocol" default:"https"`
	DefaultPort           *uint   `json:"defaultPort" form:"defaultPort" default:"443"`
	Description           *string `json:"description" form:"description"`
	PlaybookURL           *string `json:"playbookURL" form:"playbookURL"`
}

//nolint:unused
func (c Chart) toModel(_ *v2models.StoreSet) (v2models.Chart, error) {
	return v2models.Chart{
		Model: gorm.Model{
			ID:        c.ID,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		},
		Name:                  c.Name,
		ChartRepo:             c.ChartRepo,
		AppImageGitRepo:       c.AppImageGitRepo,
		AppImageGitMainBranch: c.AppImageGitMainBranch,
		ChartExposesEndpoint:  c.ChartExposesEndpoint,
		LegacyConfigsEnabled:  c.LegacyConfigsEnabled,
		DefaultSubdomain:      c.DefaultSubdomain,
		DefaultProtocol:       c.DefaultProtocol,
		DefaultPort:           c.DefaultPort,
		Description:           c.Description,
		PlaybookURL:           c.PlaybookURL,
	}, nil
}

//nolint:unused
func (c CreatableChart) toModel(storeSet *v2models.StoreSet) (v2models.Chart, error) {
	return Chart{CreatableChart: c}.toModel(storeSet)
}

//nolint:unused
func (c EditableChart) toModel(storeSet *v2models.StoreSet) (v2models.Chart, error) {
	return CreatableChart{EditableChart: c}.toModel(storeSet)
}

type ChartController = ModelController[v2models.Chart, Chart, CreatableChart, EditableChart]

func newChartController(stores *v2models.StoreSet) *ChartController {
	return &ChartController{
		primaryStore:       stores.ChartStore,
		allStores:          stores,
		modelToReadable:    modelChartToChart,
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
				LegacyConfigsEnabled:  model.LegacyConfigsEnabled,
				DefaultSubdomain:      model.DefaultSubdomain,
				DefaultProtocol:       model.DefaultProtocol,
				DefaultPort:           model.DefaultPort,
				Description:           model.Description,
				PlaybookURL:           model.PlaybookURL,
			},
		},
	}
}

func setChartDynamicDefaults(chart *CreatableChart, _ *v2models.StoreSet, _ *auth_models.User) error {
	if chart.DefaultSubdomain == nil {
		chart.DefaultSubdomain = &chart.Name
	}
	return nil
}
