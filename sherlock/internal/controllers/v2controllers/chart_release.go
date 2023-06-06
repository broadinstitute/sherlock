package v2controllers

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth/auth_models"
	"github.com/broadinstitute/sherlock/internal/config"
	"strconv"

	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"gorm.io/gorm"
)

type ChartRelease struct {
	ReadableBaseType
	CiIdentifier             *CiIdentifier         `json:"ciIdentifier,omitempty" form:"-"`
	ChartInfo                *Chart                `json:"chartInfo,omitempty" form:"-"`
	ClusterInfo              *Cluster              `json:"clusterInfo,omitempty" form:"-"`
	EnvironmentInfo          *Environment          `json:"environmentInfo,omitempty" form:"-"`
	AppVersionReference      string                `json:"appVersionReference,omitempty" form:"appVersionReference"`
	AppVersionInfo           *AppVersion           `json:"appVersionInfo,omitempty" form:"-"`
	ChartVersionReference    string                `json:"chartVersionReference,omitempty" form:"chartVersionReference"`
	ChartVersionInfo         *ChartVersion         `json:"chartVersionInfo,omitempty" form:"-"`
	PagerdutyIntegrationInfo *PagerdutyIntegration `json:"pagerdutyIntegrationInfo,omitempty" form:"-"`
	DestinationType          string                `json:"destinationType" form:"destinationType" enum:"environment,cluster"` // Calculated field
	CreatableChartRelease
}

type CreatableChartRelease struct {
	Chart       string `json:"chart" form:"chart"`             // Required when creating
	Cluster     string `json:"cluster" form:"cluster"`         // When creating, will default the environment's default cluster, if provided. Either this or environment must be provided.
	Environment string `json:"environment" form:"environment"` // Either this or cluster must be provided.
	Name        string `json:"name" form:"name"`               // When creating, will be calculated if left empty
	Namespace   string `json:"namespace" form:"namespace"`     // When creating, will default to the environment's default namespace, if provided

	AppVersionResolver             *string `json:"appVersionResolver" form:"appVersionResolver" enums:"branch,commit,exact,follow,none"` // // When creating, will default to automatically reference any provided app version fields
	AppVersionExact                *string `json:"appVersionExact" form:"appVersionExact"`
	AppVersionBranch               *string `json:"appVersionBranch" form:"appVersionBranch"` // When creating, will default to the app's mainline branch if no other app version info is present
	AppVersionCommit               *string `json:"appVersionCommit" form:"appVersionCommit"`
	AppVersionFollowChartRelease   string  `json:"appVersionFollowChartRelease" form:"appVersionFollowChartRelease"`
	ChartVersionResolver           *string `json:"chartVersionResolver" form:"chartVersionResolver" enums:"latest,exact,follow"` // When creating, will default to automatically reference any provided chart version
	ChartVersionExact              *string `json:"chartVersionExact" form:"chartVersionExact"`
	ChartVersionFollowChartRelease string  `json:"chartVersionFollowChartRelease" form:"chartVersionFollowChartRelease"`
	HelmfileRef                    *string `json:"helmfileRef" form:"helmfileRef" default:"HEAD"`
	FirecloudDevelopRef            *string `json:"firecloudDevelopRef" form:"firecloudDevelopRef"`
	EditableChartRelease
}

type EditableChartRelease struct {
	Subdomain               *string `json:"subdomain,omitempty" form:"subdomain"` // When creating, will use the chart's default if left empty
	Protocol                *string `json:"protocol,omitempty" form:"protocol"`   // When creating, will use the chart's default if left empty
	Port                    *uint   `json:"port,omitempty" form:"port"`           // When creating, will use the chart's default if left empty
	PagerdutyIntegration    *string `json:"pagerdutyIntegration,omitempty" form:"pagerdutyIntegration"`
	IncludeInBulkChangesets *bool   `json:"includedInBulkChangesets" form:"includedInBulkChangesets" default:"true"`
}

//nolint:unused
func (c ChartRelease) toModel(storeSet *v2models.StoreSet) (v2models.ChartRelease, error) {
	var chartID uint
	if c.Chart != "" {
		chart, err := storeSet.ChartStore.Get(c.Chart)
		if err != nil {
			return v2models.ChartRelease{}, err
		}
		chartID = chart.ID
	}
	var environmentID *uint
	if c.Environment != "" {
		environment, err := storeSet.EnvironmentStore.Get(c.Environment)
		if err != nil {
			return v2models.ChartRelease{}, err
		}
		environmentID = &environment.ID
	}
	var clusterID *uint
	if c.Cluster != "" {
		cluster, err := storeSet.ClusterStore.Get(c.Cluster)
		if err != nil {
			return v2models.ChartRelease{}, err
		}
		clusterID = &cluster.ID
	}
	var appVersionID *uint
	if c.AppVersionReference != "" {
		appVersion, err := storeSet.AppVersionStore.Get(c.AppVersionReference)
		if err != nil {
			return v2models.ChartRelease{}, err
		}
		appVersionID = &appVersion.ID
	}
	var chartVersionID *uint
	if c.ChartVersionReference != "" {
		chartVersion, err := storeSet.ChartVersionStore.Get(c.ChartVersionReference)
		if err != nil {
			return v2models.ChartRelease{}, err
		}
		chartVersionID = &chartVersion.ID
	}
	var appVersionFollowChartReleaseID *uint
	if c.AppVersionFollowChartRelease != "" {
		followChartRelease, err := storeSet.ChartReleaseStore.Get(c.AppVersionFollowChartRelease)
		if err != nil {
			return v2models.ChartRelease{}, err
		}
		appVersionFollowChartReleaseID = &followChartRelease.ID
	}
	var chartVersionFollowChartReleaseID *uint
	if c.ChartVersionFollowChartRelease != "" {
		followChartRelease, err := storeSet.ChartReleaseStore.Get(c.ChartVersionFollowChartRelease)
		if err != nil {
			return v2models.ChartRelease{}, err
		}
		chartVersionFollowChartReleaseID = &followChartRelease.ID
	}
	var pagerdutyIntegrationID *uint
	if c.PagerdutyIntegration != nil && *c.PagerdutyIntegration != "" {
		pagerdutyIntegration, err := storeSet.PagerdutyIntegration.Get(*c.PagerdutyIntegration)
		if err != nil {
			return v2models.ChartRelease{}, err
		}
		pagerdutyIntegrationID = &pagerdutyIntegration.ID
	}
	return v2models.ChartRelease{
		Model: gorm.Model{
			ID:        c.ID,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		},
		ChartID:         chartID,
		ClusterID:       clusterID,
		EnvironmentID:   environmentID,
		DestinationType: c.DestinationType,
		Name:            c.Name,
		Namespace:       c.Namespace,
		ChartReleaseVersion: v2models.ChartReleaseVersion{
			AppVersionResolver:               c.AppVersionResolver,
			AppVersionExact:                  c.AppVersionExact,
			AppVersionBranch:                 c.AppVersionBranch,
			AppVersionCommit:                 c.AppVersionCommit,
			AppVersionFollowChartReleaseID:   appVersionFollowChartReleaseID,
			AppVersionID:                     appVersionID,
			ChartVersionResolver:             c.ChartVersionResolver,
			ChartVersionExact:                c.ChartVersionExact,
			ChartVersionFollowChartReleaseID: chartVersionFollowChartReleaseID,
			ChartVersionID:                   chartVersionID,
			HelmfileRef:                      c.HelmfileRef,
			FirecloudDevelopRef:              c.FirecloudDevelopRef,
		},
		Subdomain:               c.Subdomain,
		Protocol:                c.Protocol,
		Port:                    c.Port,
		PagerdutyIntegrationID:  pagerdutyIntegrationID,
		IncludeInBulkChangesets: c.IncludeInBulkChangesets,
	}, nil
}

//nolint:unused
func (c CreatableChartRelease) toModel(storeSet *v2models.StoreSet) (v2models.ChartRelease, error) {
	return ChartRelease{CreatableChartRelease: c}.toModel(storeSet)
}

//nolint:unused
func (c EditableChartRelease) toModel(storeSet *v2models.StoreSet) (v2models.ChartRelease, error) {
	return CreatableChartRelease{EditableChartRelease: c}.toModel(storeSet)
}

type ChartReleaseController = ModelController[v2models.ChartRelease, ChartRelease, CreatableChartRelease, EditableChartRelease]

func newChartReleaseController(stores *v2models.StoreSet) *ChartReleaseController {
	return &ChartReleaseController{
		primaryStore:                   stores.ChartReleaseStore,
		allStores:                      stores,
		modelToReadable:                modelChartReleaseToChartRelease,
		setDynamicDefaults:             setChartReleaseDynamicDefaults,
		extractPagerdutyIntegrationKey: extractPagerdutyIntegrationKeyFromChartRelease,
		beehiveUrlFormatString:         config.Config.MustString("beehive.chartReleaseUrlFormatString"),
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

	var appVersionFollowChartRelease string
	if model.AppVersionFollowChartRelease != nil && model.AppVersionFollowChartRelease.Name != "" {
		appVersionFollowChartRelease = model.AppVersionFollowChartRelease.Name
	} else if model.AppVersionFollowChartReleaseID != nil {
		appVersionFollowChartRelease = strconv.FormatUint(uint64(*model.AppVersionFollowChartReleaseID), 10)
	}

	var chartVersionFollowChartRelease string
	if model.ChartVersionFollowChartRelease != nil && model.ChartVersionFollowChartRelease.Name != "" {
		chartVersionFollowChartRelease = model.ChartVersionFollowChartRelease.Name
	} else if model.ChartVersionFollowChartReleaseID != nil {
		chartVersionFollowChartRelease = strconv.FormatUint(uint64(*model.ChartVersionFollowChartReleaseID), 10)
	}

	var pagerdutyIntegrationID string
	pagerdutyIntegration := modelPagerdutyIntegrationToPagerdutyIntegration(model.PagerdutyIntegration)
	if pagerdutyIntegration != nil {
		pagerdutyIntegrationID = strconv.FormatUint(uint64(pagerdutyIntegration.ID), 10)
	} else if model.PagerdutyIntegrationID != nil {
		pagerdutyIntegrationID = strconv.FormatUint(uint64(*model.PagerdutyIntegrationID), 10)
	}

	return &ChartRelease{
		ReadableBaseType: ReadableBaseType{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
		CiIdentifier:             modelCiIdentifierToCiIdentifier(model.CiIdentifier),
		ChartInfo:                chart,
		ClusterInfo:              cluster,
		EnvironmentInfo:          environment,
		AppVersionReference:      appVersionReference,
		AppVersionInfo:           appVersion,
		ChartVersionReference:    chartVersionReference,
		ChartVersionInfo:         chartVersion,
		PagerdutyIntegrationInfo: pagerdutyIntegration,
		DestinationType:          model.DestinationType,
		CreatableChartRelease: CreatableChartRelease{
			Chart:                          chartName,
			Cluster:                        clusterName,
			Environment:                    environmentName,
			Name:                           model.Name,
			Namespace:                      model.Namespace,
			AppVersionResolver:             model.AppVersionResolver,
			AppVersionExact:                model.AppVersionExact,
			AppVersionBranch:               model.AppVersionBranch,
			AppVersionCommit:               model.AppVersionCommit,
			AppVersionFollowChartRelease:   appVersionFollowChartRelease,
			ChartVersionResolver:           model.ChartVersionResolver,
			ChartVersionExact:              model.ChartVersionExact,
			ChartVersionFollowChartRelease: chartVersionFollowChartRelease,
			HelmfileRef:                    model.HelmfileRef,
			FirecloudDevelopRef:            model.FirecloudDevelopRef,
			EditableChartRelease: EditableChartRelease{
				Subdomain:               model.Subdomain,
				Protocol:                model.Protocol,
				Port:                    model.Port,
				PagerdutyIntegration:    &pagerdutyIntegrationID,
				IncludeInBulkChangesets: model.IncludeInBulkChangesets,
			},
		},
	}
}

func setChartReleaseDynamicDefaults(chartRelease *CreatableChartRelease, stores *v2models.StoreSet, _ *auth_models.User) error {
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
		} else if chartRelease.AppVersionFollowChartRelease != "" {
			resolver = "follow"
		}
		if resolver != "" {
			chartRelease.AppVersionResolver = &resolver
		}
	}

	if chartRelease.ChartVersionResolver == nil {
		resolver := "latest"
		if chartRelease.ChartVersionExact != nil {
			resolver = "exact"
		} else if chartRelease.ChartVersionFollowChartRelease != "" {
			resolver = "follow"
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
		if chartRelease.Namespace == "" && environment.DefaultNamespace != "" {
			chartRelease.Namespace = environment.DefaultNamespace
		}
		if chartRelease.FirecloudDevelopRef == nil {
			if chart.LegacyConfigsEnabled != nil && *chart.LegacyConfigsEnabled {
				chartRelease.FirecloudDevelopRef = environment.DefaultFirecloudDevelopRef
			}
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
	}
	return nil
}

func extractPagerdutyIntegrationKeyFromChartRelease(model *v2models.ChartRelease) *string {
	if model != nil && model.PagerdutyIntegration != nil {
		return model.PagerdutyIntegration.Key
	}
	return nil
}
