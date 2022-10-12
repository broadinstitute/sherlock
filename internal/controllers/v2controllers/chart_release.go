package v2controllers

import (
	"fmt"
	"strconv"

	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"gorm.io/gorm"
)

type ChartRelease struct {
	ReadableBaseType
	ChartInfo             *Chart        `json:"chartInfo,omitempty" form:"-"`
	ClusterInfo           *Cluster      `json:"clusterInfo,omitempty" form:"-"`
	EnvironmentInfo       *Environment  `json:"environmentInfo,omitempty" form:"-"`
	AppVersionReference   string        `json:"appVersionReference,omitempty" form:"appVersionReference"`
	AppVersionInfo        *AppVersion   `json:"appVersionInfo,omitempty" form:"-"`
	ChartVersionReference string        `json:"chartVersionReference,omitempty" form:"chartVersionReference"`
	ChartVersionInfo      *ChartVersion `json:"chartVersionInfo,omitempty" form:"-"`
	DestinationType       string        `json:"destinationType" form:"destinationType" enum:"environment,cluster"` // Calculated field
	CreatableChartRelease
}

type CreatableChartRelease struct {
	Chart       string `json:"chart" form:"chart"`             // Required when creating
	Cluster     string `json:"cluster" form:"cluster"`         // When creating, will default the environment's default cluster, if provided. Either this or environment must be provided.
	Environment string `json:"environment" form:"environment"` // Either this or cluster must be provided.
	Name        string `json:"name" form:"name"`               // When creating, will be calculated if left empty
	Namespace   string `json:"namespace" form:"namespace"`     // When creating, will default to the environment's default namespace, if provided

	AppVersionResolver   *string `json:"appVersionResolver" form:"appVersionResolver" enums:"branch,commit,exact,none"` // // When creating, will default to automatically reference any provided app version fields
	AppVersionExact      *string `json:"appVersionExact" form:"appVersionExact"`
	AppVersionBranch     *string `json:"appVersionBranch" form:"appVersionBranch"` // When creating, will default to the app's mainline branch if no other app version info is present
	AppVersionCommit     *string `json:"appVersionCommit" form:"appVersionCommit"`
	ChartVersionResolver *string `json:"chartVersionResolver" form:"chartVersionResolver" enums:"latest,exact"` // When creating, will default to automatically reference any provided chart version
	ChartVersionExact    *string `json:"chartVersionExact" form:"chartVersionExact"`
	HelmfileRef          *string `json:"helmfileRef" form:"helmfileRef" default:"HEAD"`
	FirecloudDevelopRef  *string `json:"firecloudDevelopRef" form:"firecloudDevelopRef" default:"dev"`
	EditableChartRelease
}

type EditableChartRelease struct {
	Subdomain *string `json:"subdomain,omitempty" form:"subdomain"` // When creating, will use the chart's default if left empty
	Protocol  *string `json:"protocol,omitempty" form:"protocol"`   // When creating, will use the chart's default if left empty
	Port      *uint   `json:"port,omitempty" form:"port"`           // When creating, will use the chart's default if left empty
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

func modelChartReleaseToChartRelease(model *v2models.ChartRelease) *ChartRelease {
	if model == nil {
		return nil
	}

	var chartName string
	chart := modelChartToChart(model.Chart)
	if chart != nil {
		chartName = chart.Name
	}

	var environmentName string
	environment := modelEnvironmentToEnvironment(model.Environment)
	if environment != nil {
		environmentName = environment.Name
	}

	var clusterName string
	cluster := modelClusterToCluster(model.Cluster)
	if cluster != nil {
		clusterName = cluster.Name
	}

	var appVersionReference string
	appVersion := modelAppVersionToAppVersion(model.AppVersion)
	if appVersion != nil {
		appVersionReference = fmt.Sprintf("%s/%s", chartName, appVersion.AppVersion)
	} else if model.AppVersionID != nil {
		appVersionReference = strconv.FormatUint(uint64(*model.AppVersionID), 10)
	}

	var chartVersionReference string
	chartVersion := modelChartVersionToChartVersion(model.ChartVersion)
	if chartVersion != nil {
		chartVersionReference = fmt.Sprintf("%s/%s", chartName, chartVersion.ChartVersion)
	} else if model.ChartVersionID != nil {
		chartVersionReference = strconv.FormatUint(uint64(*model.ChartVersionID), 10)
	}

	return &ChartRelease{
		ReadableBaseType: ReadableBaseType{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
		ChartInfo:             chart,
		ClusterInfo:           cluster,
		EnvironmentInfo:       environment,
		AppVersionReference:   appVersionReference,
		AppVersionInfo:        appVersion,
		ChartVersionReference: chartVersionReference,
		ChartVersionInfo:      chartVersion,
		DestinationType:       model.DestinationType,
		CreatableChartRelease: CreatableChartRelease{
			Chart:                chartName,
			Cluster:              clusterName,
			Environment:          environmentName,
			Name:                 model.Name,
			Namespace:            model.Namespace,
			AppVersionResolver:   model.AppVersionResolver,
			AppVersionExact:      model.AppVersionExact,
			AppVersionBranch:     model.AppVersionBranch,
			AppVersionCommit:     model.AppVersionCommit,
			ChartVersionResolver: model.ChartVersionResolver,
			ChartVersionExact:    model.ChartVersionExact,
			HelmfileRef:          model.HelmfileRef,
			FirecloudDevelopRef:  model.FirecloudDevelopRef,
			EditableChartRelease: EditableChartRelease{
				Subdomain: model.Subdomain,
				Protocol:  model.Protocol,
				Port:      model.Port,
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
	var appVersionID *uint
	if chartRelease.AppVersionReference != "" {
		appVersion, err := stores.AppVersionStore.Get(chartRelease.AppVersionReference)
		if err != nil {
			return v2models.ChartRelease{}, err
		}
		appVersionID = &appVersion.ID
	}
	var chartVersionID *uint
	if chartRelease.ChartVersionReference != "" {
		chartVersion, err := stores.ChartVersionStore.Get(chartRelease.ChartVersionReference)
		if err != nil {
			return v2models.ChartRelease{}, err
		}
		chartVersionID = &chartVersion.ID
	}
	return v2models.ChartRelease{
		Model: gorm.Model{
			ID:        chartRelease.ID,
			CreatedAt: chartRelease.CreatedAt,
			UpdatedAt: chartRelease.UpdatedAt,
		},
		ChartID:         chartID,
		ClusterID:       clusterID,
		EnvironmentID:   environmentID,
		DestinationType: chartRelease.DestinationType,
		Name:            chartRelease.Name,
		Namespace:       chartRelease.Namespace,
		ChartReleaseVersion: v2models.ChartReleaseVersion{
			AppVersionResolver:   chartRelease.AppVersionResolver,
			AppVersionExact:      chartRelease.AppVersionExact,
			AppVersionBranch:     chartRelease.AppVersionBranch,
			AppVersionCommit:     chartRelease.AppVersionCommit,
			AppVersionID:         appVersionID,
			ChartVersionResolver: chartRelease.ChartVersionResolver,
			ChartVersionExact:    chartRelease.ChartVersionExact,
			ChartVersionID:       chartVersionID,
			HelmfileRef:          chartRelease.HelmfileRef,
			FirecloudDevelopRef:  chartRelease.FirecloudDevelopRef,
		},
		Subdomain: chartRelease.Subdomain,
		Protocol:  chartRelease.Protocol,
		Port:      chartRelease.Port,
	}, nil
}

func setChartReleaseDynamicDefaults(chartRelease *ChartRelease, stores *v2models.StoreSet, _ *auth.User) error {
	chart, err := stores.ChartStore.Get(chartRelease.Chart)
	if err != nil {
		return err
	}
	if chart.AppImageGitMainBranch != nil && *chart.AppImageGitMainBranch != "" && chartRelease.AppVersionBranch == nil {
		chartRelease.AppVersionBranch = chart.AppImageGitMainBranch
	}
	if chart.ChartExposesEndpoint != nil && *chart.ChartExposesEndpoint {
		if chartRelease.Subdomain == nil {
			chartRelease.Subdomain = chart.DefaultSubdomain
		}
		if chartRelease.Protocol == nil {
			chartRelease.Protocol = chart.DefaultProtocol
		}
		if chartRelease.Port == nil {
			chartRelease.Port = chart.DefaultPort
		}
	}

	if chartRelease.AppVersionResolver == nil {
		resolver := "none"
		if chartRelease.AppVersionExact != nil {
			resolver = "exact"
		} else if chartRelease.AppVersionCommit != nil {
			resolver = "commit"
		} else if chartRelease.AppVersionBranch != nil {
			resolver = "branch"
		}
		if resolver != "" {
			chartRelease.AppVersionResolver = &resolver
		}
	}

	if chartRelease.ChartVersionResolver == nil {
		resolver := "latest"
		if chartRelease.ChartVersionExact != nil {
			resolver = "exact"
		}
		chartRelease.ChartVersionResolver = &resolver
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
