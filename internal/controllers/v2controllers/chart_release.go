package v2controllers

import (
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"gorm.io/gorm"
)

type ChartRelease struct {
	ReadableBaseType
	ChartInfo       Chart        `json:"chartInfo" form:"chartInfo"`
	ClusterInfo     *Cluster     `json:"clusterInfo,omitempty" form:"clusterInfo"`
	EnvironmentInfo *Environment `json:"environmentInfo,omitempty" form:"environmentInfo"`
	CreatableChartRelease
}

type CreatableChartRelease struct {
	Chart       string `json:"chart" form:"chart"`
	Cluster     string `json:"cluster" form:"cluster"`
	Environment string `json:"environment" form:"environment"`
	Name        string `json:"name" form:"name"`
	Namespace   string `json:"namespace" form:"namespace"`
	EditableChartRelease
}

type EditableChartRelease struct {
	CurrentAppVersionExact   *string `json:"currentAppVersionExact" form:"currentAppVersionExact"`
	CurrentChartVersionExact *string `json:"currentChartVersionExact" form:"currentChartVersionExact"`
	HelmfileRef              *string `json:"helmfileRef" form:"helmfileRef"`
	TargetAppVersionBranch   *string `json:"targetAppVersionBranch" form:"targetAppVersionBranch"`
	TargetAppVersionCommit   *string `json:"targetAppVersionCommit" form:"targetAppVersionCommit"`
	TargetAppVersionExact    *string `json:"targetAppVersionExact" form:"targetAppVersionExact"`
	TargetAppVersionUse      *string `json:"targetAppVersionUse" form:"targetAppVersionUse"`
	TargetChartVersionExact  *string `json:"targetChartVersionExact" form:"targetChartVersionExact"`
	TargetChartVersionUse    *string `json:"targetChartVersionUse" form:"targetChartVersionUse"`
	ThelmaMode               *string `json:"thelmaMode,omitempty" form:"thelmaMode"`
}

func (c CreatableChartRelease) toReadable() ChartRelease {
	return ChartRelease{CreatableChartRelease: c}
}

func (e EditableChartRelease) toCreatable() CreatableChartRelease {
	return CreatableChartRelease{EditableChartRelease: e}
}

type ChartReleaseController = ModelController[v2models.ChartRelease, ChartRelease, CreatableChartRelease, EditableChartRelease]

func NewChartReleaseController(stores v2models.StoreSet) *ChartReleaseController {
	return &ChartReleaseController{
		primaryStore:    stores.ChartReleaseStore,
		allStores:       stores,
		modelToReadable: modelChartReleaseToChartRelease,
		readableToModel: chartReleaseToModelChartRelease,
	}
}

func modelChartReleaseToChartRelease(model v2models.ChartRelease) *ChartRelease {
	chart := modelChartToChart(model.Chart)
	var environment *Environment
	var environmentName string
	if model.Environment != nil {
		environment = modelEnvironmentToEnvironment(*model.Environment)
		environmentName = environment.Name
	}
	var cluster *Cluster
	var clusterName string
	if model.Cluster != nil {
		cluster = modelClusterToCluster(*model.Cluster)
		clusterName = cluster.Name
	}
	return &ChartRelease{
		ReadableBaseType: ReadableBaseType{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
		ChartInfo:       *chart,
		ClusterInfo:     cluster,
		EnvironmentInfo: environment,
		CreatableChartRelease: CreatableChartRelease{
			Chart:       chart.Name,
			Cluster:     clusterName,
			Environment: environmentName,
			Name:        model.Name,
			Namespace:   model.Namespace,
			EditableChartRelease: EditableChartRelease{
				CurrentAppVersionExact:   model.CurrentAppVersionExact,
				CurrentChartVersionExact: model.CurrentChartVersionExact,
				HelmfileRef:              model.HelmfileRef,
				TargetAppVersionBranch:   model.TargetAppVersionBranch,
				TargetAppVersionCommit:   model.TargetAppVersionCommit,
				TargetAppVersionExact:    model.TargetAppVersionExact,
				TargetAppVersionUse:      model.TargetAppVersionUse,
				TargetChartVersionExact:  model.TargetChartVersionExact,
				TargetChartVersionUse:    model.TargetChartVersionUse,
				ThelmaMode:               model.ThelmaMode,
			},
		},
	}
}

func chartReleaseToModelChartRelease(chartRelease ChartRelease, stores v2models.StoreSet) (v2models.ChartRelease, error) {
	var chartID uint
	if chartRelease.Chart != "" {
		chart, err := stores.ChartStore.Get(chartRelease.Chart)
		if err != nil {
			return v2models.ChartRelease{}, err
		}
		chartID = chart.ID
	}
	var environmentID *uint
	if chartRelease.Environment != "" {
		environment, err := stores.EnvironmentStore.Get(chartRelease.Environment)
		if err != nil {
			return v2models.ChartRelease{}, err
		}
		environmentID = &environment.ID
	}
	var clusterID *uint
	if chartRelease.Cluster != "" {
		cluster, err := stores.ClusterStore.Get(chartRelease.Cluster)
		if err != nil {
			return v2models.ChartRelease{}, err
		}
		clusterID = &cluster.ID
	}
	return v2models.ChartRelease{
		Model: gorm.Model{
			ID:        chartRelease.ID,
			CreatedAt: chartRelease.CreatedAt,
			UpdatedAt: chartRelease.UpdatedAt,
		},
		ChartID:                  chartID,
		ClusterID:                clusterID,
		EnvironmentID:            environmentID,
		Name:                     chartRelease.Name,
		Namespace:                chartRelease.Namespace,
		CurrentAppVersionExact:   chartRelease.CurrentAppVersionExact,
		CurrentChartVersionExact: chartRelease.CurrentChartVersionExact,
		HelmfileRef:              chartRelease.HelmfileRef,
		TargetAppVersionBranch:   chartRelease.TargetAppVersionBranch,
		TargetAppVersionCommit:   chartRelease.TargetAppVersionCommit,
		TargetAppVersionExact:    chartRelease.TargetAppVersionExact,
		TargetAppVersionUse:      chartRelease.TargetAppVersionUse,
		TargetChartVersionExact:  chartRelease.TargetChartVersionExact,
		TargetChartVersionUse:    chartRelease.TargetChartVersionUse,
		ThelmaMode:               chartRelease.ThelmaMode,
	}, nil
}
