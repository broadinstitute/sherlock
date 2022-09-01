package v2controllers

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"gorm.io/gorm"
)

type ChartRelease struct {
	ReadableBaseType
	ChartInfo       Chart        `json:"chartInfo" form:"-"`
	ClusterInfo     *Cluster     `json:"clusterInfo,omitempty" form:"-"`
	EnvironmentInfo *Environment `json:"environmentInfo,omitempty" form:"-"`
	DestinationType string       `json:"destinationType" form:"destinationType" enum:"environment,cluster"` // Calculated field
	CreatableChartRelease
}

type CreatableChartRelease struct {
	Chart       string `json:"chart" form:"chart"`             // Required when creating
	Cluster     string `json:"cluster" form:"cluster"`         // When creating, will default the environment's default cluster, if provided. Either this or environment must be provided.
	Environment string `json:"environment" form:"environment"` // Either this or cluster must be provided.
	Name        string `json:"name" form:"name"`               // When creating, will be calculated if left empty
	Namespace   string `json:"namespace" form:"namespace"`     // When creating, will default to the environment's default namespace, if provided
	EditableChartRelease
}

// EditableChartRelease
// There's indeed some grouped fields here. Trying to nest them in an object or something will quickly make Swaggo
// upset. There's four things to worry about: Gin's json parsing, Swaggo's json tag parsing, Gin's form parsing for
// query params, and Swaggo's apparent parsing of json tags (instead of form ones) for query params. The query param
// part only shows up on the list method; all the times Jack tried to add nesting, something broke.
// https://broadinstitute.slack.com/archives/CQ6SL4N5T/p1660059822037769
type EditableChartRelease struct {
	CurrentAppVersionExact   *string `json:"currentAppVersionExact" form:"currentAppVersionExact"`
	CurrentChartVersionExact *string `json:"currentChartVersionExact" form:"currentChartVersionExact"`
	HelmfileRef              *string `json:"helmfileRef" form:"helmfileRef" default:"HEAD"`
	TargetAppVersionBranch   *string `json:"targetAppVersionBranch" form:"targetAppVersionBranch"` // When creating, will default to the app's main branch if it has one recorded
	TargetAppVersionCommit   *string `json:"targetAppVersionCommit" form:"targetAppVersionCommit"`
	TargetAppVersionExact    *string `json:"targetAppVersionExact" form:"targetAppVersionExact"`
	TargetAppVersionUse      *string `json:"targetAppVersionUse" form:"targetAppVersionUse" enums:"branch,commit,exact"` // When creating, will default to referencing any provided target app version field (exact, then commit, then branch)
	TargetChartVersionExact  *string `json:"targetChartVersionExact" form:"targetChartVersionExact"`
	TargetChartVersionUse    *string `json:"targetChartVersionUse" form:"targetChartVersionUse" enums:"latest,exact"` // When creating, will default to latest unless an exact target chart version is provided
	ThelmaMode               *string `json:"thelmaMode,omitempty" form:"thelmaMode"`
}

//nolint:unused
func (c CreatableChartRelease) toReadable() ChartRelease {
	return ChartRelease{CreatableChartRelease: c}
}

//nolint:unused
func (e EditableChartRelease) toCreatable() CreatableChartRelease {
	return CreatableChartRelease{EditableChartRelease: e}
}

type ChartReleaseController = ModelController[v2models.ChartRelease, ChartRelease, CreatableChartRelease, EditableChartRelease]

func newChartReleaseController(stores *v2models.StoreSet) *ChartReleaseController {
	return &ChartReleaseController{
		primaryStore:       stores.ChartReleaseStore,
		allStores:          stores,
		modelToReadable:    modelChartReleaseToChartRelease,
		readableToModel:    chartReleaseToModelChartRelease,
		setDynamicDefaults: setChartReleaseDynamicDefaults,
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
		DestinationType: model.DestinationType,
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

func chartReleaseToModelChartRelease(chartRelease ChartRelease, stores *v2models.StoreSet) (v2models.ChartRelease, error) {
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
		DestinationType:          chartRelease.DestinationType,
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

func setChartReleaseDynamicDefaults(chartRelease *ChartRelease, stores *v2models.StoreSet, _ *auth.User) error {
	chart, err := stores.ChartStore.Get(chartRelease.Chart)
	if err != nil {
		return err
	}
	if chart.AppImageGitMainBranch != nil && *chart.AppImageGitMainBranch != "" && chartRelease.TargetAppVersionBranch == nil {
		chartRelease.TargetAppVersionBranch = chart.AppImageGitMainBranch
	}

	if chartRelease.TargetAppVersionUse == nil {
		var temp string
		if chartRelease.TargetAppVersionExact != nil {
			temp = "exact"
		} else if chartRelease.TargetAppVersionCommit != nil {
			temp = "commit"
		} else if chartRelease.TargetAppVersionBranch != nil {
			temp = "branch"
		}
		if temp != "" {
			chartRelease.TargetAppVersionUse = &temp
		}
	}

	if chartRelease.TargetChartVersionUse == nil {
		temp := "latest"
		if chartRelease.TargetChartVersionExact != nil {
			temp = "exact"
		}
		chartRelease.TargetChartVersionUse = &temp
	}

	if chartRelease.Environment != "" {
		environment, err := stores.EnvironmentStore.Get(chartRelease.Environment)
		if err != nil {
			return err
		}
		if chartRelease.Name == "" {
			chartRelease.Name = fmt.Sprintf("%s-%s", chart.Name, environment.Name)
		}
		if chartRelease.Cluster == "" && environment.DefaultCluster != nil {
			chartRelease.Cluster = environment.DefaultCluster.Name
		}
		if chartRelease.Namespace == "" && environment.DefaultNamespace != nil {
			chartRelease.Namespace = *environment.DefaultNamespace
		}
		if chartRelease.DestinationType == "" {
			chartRelease.DestinationType = "environment"
		}
	}

	if chartRelease.Cluster != "" {
		cluster, err := stores.ClusterStore.Get(chartRelease.Cluster)
		if err != nil {
			return err
		}
		if chartRelease.Name == "" {
			if chartRelease.Namespace == "" || chartRelease.Namespace == cluster.Name {
				chartRelease.Name = fmt.Sprintf("%s-%s", chart.Name, cluster.Name)
			} else {
				chartRelease.Name = fmt.Sprintf("%s-%s-%s", chart.Name, chartRelease.Namespace, cluster.Name)
			}
		}
		if chartRelease.DestinationType == "" {
			chartRelease.DestinationType = "cluster"
		}
	}
	return nil
}
