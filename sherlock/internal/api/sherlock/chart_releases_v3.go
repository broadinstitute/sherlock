package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"gorm.io/gorm"
	"time"
)

type ChartReleaseV3 struct {
	CommonFields
	CiIdentifier             *CiIdentifierV3         `json:"ciIdentifier,omitempty" form:"-"`
	ChartInfo                *ChartV3                `json:"chartInfo,omitempty" form:"-"`
	ClusterInfo              *ClusterV3              `json:"clusterInfo,omitempty" form:"-"`
	EnvironmentInfo          *EnvironmentV3          `json:"environmentInfo,omitempty" form:"-"`
	AppVersionReference      string                  `json:"appVersionReference,omitempty" form:"appVersionReference"`
	AppVersionInfo           *AppVersionV3           `json:"appVersionInfo,omitempty" form:"-"`
	ChartVersionReference    string                  `json:"chartVersionReference,omitempty" form:"chartVersionReference"`
	ChartVersionInfo         *ChartVersionV3         `json:"chartVersionInfo,omitempty" form:"-"`
	PagerdutyIntegrationInfo *PagerdutyIntegrationV3 `json:"pagerdutyIntegrationInfo,omitempty" form:"-"`
	DestinationType          string                  `json:"destinationType" form:"destinationType" enum:"environment,cluster"` // Calculated field
	ResolvedAt               *time.Time              `json:"resolvedAt,omitempty" form:"resolvedAt" format:"date-time"`
	ChartReleaseV3Create
}

type ChartReleaseV3Create struct {
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
	HelmfileRefEnabled             *bool   `json:"helmfileRefEnabled" form:"helmfileRefEnabled" default:"false"`
	FirecloudDevelopRef            *string `json:"firecloudDevelopRef" form:"firecloudDevelopRef"`
	ChartReleaseV3Edit
}

type ChartReleaseV3Edit struct {
	Subdomain               *string `json:"subdomain,omitempty" form:"subdomain"` // When creating, will use the chart's default if left empty
	Protocol                *string `json:"protocol,omitempty" form:"protocol"`   // When creating, will use the chart's default if left empty
	Port                    *uint   `json:"port,omitempty" form:"port"`           // When creating, will use the chart's default if left empty
	PagerdutyIntegration    *string `json:"pagerdutyIntegration,omitempty" form:"pagerdutyIntegration"`
	IncludeInBulkChangesets *bool   `json:"includedInBulkChangesets" form:"includedInBulkChangesets" default:"true"`
}

func (c ChartReleaseV3) toModel(db *gorm.DB) (models.ChartRelease, error) {
	ret := models.ChartRelease{
		Model:           c.CommonFields.toGormModel(),
		DestinationType: c.DestinationType,
		Name:            c.Name,
		Namespace:       c.Namespace,
		ChartReleaseVersion: models.ChartReleaseVersion{
			ResolvedAt:           c.ResolvedAt,
			AppVersionResolver:   c.AppVersionResolver,
			AppVersionExact:      c.AppVersionExact,
			AppVersionBranch:     c.AppVersionBranch,
			AppVersionCommit:     c.AppVersionCommit,
			ChartVersionResolver: c.ChartVersionResolver,
			ChartVersionExact:    c.ChartVersionExact,
			HelmfileRef:          c.HelmfileRef,
			HelmfileRefEnabled:   c.HelmfileRefEnabled,
			FirecloudDevelopRef:  c.FirecloudDevelopRef,
		},
		Subdomain:               c.Subdomain,
		Protocol:                c.Protocol,
		Port:                    c.Port,
		PagerdutyIntegrationID:  nil, //
		IncludeInBulkChangesets: c.IncludeInBulkChangesets,
	}
	if c.Chart != "" {
		chartModel, err := chartModelFromSelector(c.Chart)
		if err != nil {
			return models.ChartRelease{}, err
		}
		var chart models.Chart
		if err = db.Where(&chartModel).Select("id").First(&chart).Error; err != nil {
			return models.ChartRelease{}, err
		} else {
			ret.ChartID = chart.ID
		}
	}
	if c.Cluster != "" {
		clusterModel, err := clusterModelFromSelector(c.Cluster)
		if err != nil {
			return models.ChartRelease{}, err
		}
		var cluster models.Cluster
		if err = db.Where(&clusterModel).Select("id").First(&cluster).Error; err != nil {
			return models.ChartRelease{}, err
		} else {
			ret.ClusterID = &cluster.ID
		}
	}
	if c.Environment != "" {
		environmentModel, err := environmentModelFromSelector(c.Environment)
		if err != nil {
			return models.ChartRelease{}, err
		}
		var environment models.Environment
		if err = db.Where(&environmentModel).Select("id").First(&environment).Error; err != nil {
			return models.ChartRelease{}, err
		} else {
			ret.EnvironmentID = &environment.ID
		}
	}
	if c.AppVersionFollowChartRelease != "" {
		chartReleaseModel, err := chartReleaseModelFromSelector(db, c.AppVersionFollowChartRelease)
		if err != nil {
			return models.ChartRelease{}, err
		}
		var chartRelease models.ChartRelease
		if err = db.Where(&chartReleaseModel).Select("id").First(&chartRelease).Error; err != nil {
			return models.ChartRelease{}, err
		} else {
			ret.AppVersionFollowChartReleaseID = &chartRelease.ID
		}
	}
	if c.AppVersionReference != "" {
		appVersionModel, err := appVersionModelFromSelector(db, c.AppVersionReference)
		if err != nil {
			return models.ChartRelease{}, err
		}
		var appVersion models.AppVersion
		if err = db.Where(&appVersionModel).Select("id").First(&appVersion).Error; err != nil {
			return models.ChartRelease{}, err
		} else {
			ret.AppVersionID = &appVersion.ID
		}
	}
	if c.ChartVersionFollowChartRelease != "" {
		chartReleaseModel, err := chartReleaseModelFromSelector(db, c.ChartVersionFollowChartRelease)
		if err != nil {
			return models.ChartRelease{}, err
		}
		var chartRelease models.ChartRelease
		if err = db.Where(&chartReleaseModel).Select("id").First(&chartRelease).Error; err != nil {
			return models.ChartRelease{}, err
		} else {
			ret.ChartVersionFollowChartReleaseID = &chartRelease.ID
		}
	}
	if c.ChartVersionReference != "" {
		chartVersionModel, err := chartVersionModelFromSelector(db, c.ChartVersionReference)
		if err != nil {
			return models.ChartRelease{}, err
		}
		var chartVersion models.ChartVersion
		if err = db.Where(&chartVersionModel).Select("id").First(&chartVersion).Error; err != nil {
			return models.ChartRelease{}, err
		} else {
			ret.ChartVersionID = &chartVersion.ID
		}
	}
	if c.PagerdutyIntegration != nil {
		pagerdutyIntegrationModel, err := pagerdutyIntegrationModelFromSelector(*c.PagerdutyIntegration)
		if err != nil {
			return models.ChartRelease{}, err
		}
		var pagerdutyIntegration models.PagerdutyIntegration
		if err = db.Where(&pagerdutyIntegrationModel).Select("id").First(&pagerdutyIntegration).Error; err != nil {
			return models.ChartRelease{}, err
		} else {
			ret.PagerdutyIntegrationID = &pagerdutyIntegration.ID
		}
	}
	return ret, nil
}

func (c ChartReleaseV3Create) toModel(db *gorm.DB) (models.ChartRelease, error) {
	return ChartReleaseV3{ChartReleaseV3Create: c}.toModel(db)
}

func (c ChartReleaseV3Edit) toModel(db *gorm.DB) (models.ChartRelease, error) {
	return ChartReleaseV3Create{ChartReleaseV3Edit: c}.toModel(db)
}

func chartReleaseFromModel(model models.ChartRelease) ChartReleaseV3 {
	ret := ChartReleaseV3{
		CommonFields:             commonFieldsFromGormModel(model.Model),
		CiIdentifier:             utils.NilOrCall(ciIdentifierFromModel, model.CiIdentifier),
		ChartInfo:                utils.NilOrCall(chartFromModel, model.Chart),
		ClusterInfo:              utils.NilOrCall(clusterFromModel, model.Cluster),
		EnvironmentInfo:          utils.NilOrCall(environmentFromModel, model.Environment),
		AppVersionInfo:           utils.NilOrCall(appVersionFromModel, model.AppVersion),
		ChartVersionInfo:         utils.NilOrCall(chartVersionFromModel, model.ChartVersion),
		PagerdutyIntegrationInfo: utils.NilOrCall(pagerdutyIntegrationFromModel, model.PagerdutyIntegration),
		DestinationType:          model.DestinationType,
		ResolvedAt:               model.ResolvedAt,
		ChartReleaseV3Create: ChartReleaseV3Create{
			Name:                 model.Name,
			Namespace:            model.Namespace,
			AppVersionResolver:   model.AppVersionResolver,
			AppVersionExact:      model.AppVersionExact,
			AppVersionBranch:     model.AppVersionBranch,
			AppVersionCommit:     model.AppVersionCommit,
			ChartVersionResolver: model.ChartVersionResolver,
			ChartVersionExact:    model.ChartVersionExact,
			HelmfileRef:          model.HelmfileRef,
			HelmfileRefEnabled:   model.HelmfileRefEnabled,
			FirecloudDevelopRef:  model.FirecloudDevelopRef,
			ChartReleaseV3Edit: ChartReleaseV3Edit{
				Subdomain:               model.Subdomain,
				Protocol:                model.Protocol,
				Port:                    model.Port,
				IncludeInBulkChangesets: model.IncludeInBulkChangesets,
			},
		},
	}
	if model.AppVersion != nil && model.AppVersion.AppVersion != "" && model.Chart != nil && model.Chart.Name != "" {
		ret.AppVersionReference = fmt.Sprintf("%s/%s", model.Chart.Name, *model.AppVersionExact)
	} else if model.AppVersionID != nil {
		ret.AppVersionReference = utils.UintToString(*model.AppVersionID)
	}
	if model.ChartVersion != nil && model.ChartVersion.ChartVersion != "" && model.Chart != nil && model.Chart.Name != "" {
		ret.ChartVersionReference = fmt.Sprintf("%s/%s", model.Chart.Name, *model.ChartVersionExact)
	} else if model.ChartVersionID != nil {
		ret.ChartVersionReference = utils.UintToString(*model.ChartVersionID)
	}
	if model.Chart != nil && model.Chart.Name != "" {
		ret.Chart = model.Chart.Name
	} else if model.ChartID != 0 {
		ret.Chart = utils.UintToString(model.ChartID)
	}
	if model.Cluster != nil && model.Cluster.Name != "" {
		ret.Cluster = model.Cluster.Name
	} else if model.ClusterID != nil {
		ret.Cluster = utils.UintToString(*model.ClusterID)
	}
	if model.Environment != nil && model.Environment.Name != "" {
		ret.Environment = model.Environment.Name
	} else if model.EnvironmentID != nil {
		ret.Environment = utils.UintToString(*model.EnvironmentID)
	}
	if model.AppVersionFollowChartRelease != nil && model.AppVersionFollowChartRelease.Name != "" {
		ret.AppVersionFollowChartRelease = model.AppVersionFollowChartRelease.Name
	} else if model.AppVersionFollowChartReleaseID != nil {
		ret.AppVersionFollowChartRelease = utils.UintToString(*model.AppVersionFollowChartReleaseID)
	}
	if model.ChartVersionFollowChartRelease != nil && model.ChartVersionFollowChartRelease.Name != "" {
		ret.ChartVersionFollowChartRelease = model.ChartVersionFollowChartRelease.Name
	} else if model.ChartVersionFollowChartReleaseID != nil {
		ret.ChartVersionFollowChartRelease = utils.UintToString(*model.ChartVersionFollowChartReleaseID)
	}
	if model.PagerdutyIntegration != nil && model.PagerdutyIntegration.PagerdutyID != "" {
		ret.PagerdutyIntegration = utils.PointerTo(fmt.Sprintf("pd-id/%s", model.PagerdutyIntegration.PagerdutyID))
	} else if model.PagerdutyIntegrationID != nil {
		ret.PagerdutyIntegration = utils.PointerTo(utils.UintToString(*model.PagerdutyIntegrationID))
	}
	return ret
}
