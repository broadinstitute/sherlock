package sherlock

import (
	"database/sql"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"time"
)

type EnvironmentV3 struct {
	CommonFields
	CiIdentifier             *CiIdentifierV3         `json:"ciIdentifier,omitempty" form:"-"`
	TemplateEnvironmentInfo  *EnvironmentV3          `json:"templateEnvironmentInfo,omitempty" swaggertype:"object" form:"-"`
	DefaultClusterInfo       *ClusterV3              `json:"defaultClusterInfo,omitempty" form:"-"`
	PagerdutyIntegrationInfo *PagerdutyIntegrationV3 `json:"pagerdutyIntegrationInfo,omitempty" form:"-"`
	OwnerInfo                *UserV3                 `json:"ownerInfo,omitempty" form:"-"`
	EnvironmentV3Create
}

type EnvironmentV3Create struct {
	Base                      string `json:"base" form:"base"`                                                          // Required when creating
	AutoPopulateChartReleases *bool  `json:"autoPopulateChartReleases" form:"autoPopulateChartReleases" default:"true"` // If true when creating, dynamic environments copy from template and template environments get the honeycomb chart
	Lifecycle                 string `json:"lifecycle" form:"lifecycle" default:"dynamic"`
	Name                      string `json:"name" form:"name"`                                 // When creating, will be calculated if dynamic, required otherwise
	TemplateEnvironment       string `json:"templateEnvironment" form:"templateEnvironment"`   // Required for dynamic environments
	UniqueResourcePrefix      string `json:"uniqueResourcePrefix" form:"uniqueResourcePrefix"` // When creating, will be calculated if left empty
	DefaultNamespace          string `json:"defaultNamespace" form:"defaultNamespace"`         // When creating, will be calculated if left empty
	NamePrefix                string `json:"namePrefix" form:"namePrefix"`                     // Used for dynamic environment name generation only, to override using the owner email handle and template name
	ValuesName                string `json:"valuesName" form:"valuesName"`                     // When creating, defaults to template name or environment name
	EnvironmentV3Edit
}

type EnvironmentV3Edit struct {
	DefaultCluster              *string    `json:"defaultCluster" form:"defaultCluster"`
	DefaultFirecloudDevelopRef  *string    `json:"defaultFirecloudDevelopRef" form:"defaultFirecloudDevelopRef" default:"dev"` // should be the environment branch for live envs. Is usually dev for template/dynamic but not necessarily
	Owner                       *string    `json:"owner" form:"owner"`                                                         // When creating, will default to you
	RequiresSuitability         *bool      `json:"requiresSuitability" form:"requiresSuitability" default:"false"`
	BaseDomain                  *string    `json:"baseDomain" form:"baseDomain" default:"bee.envs-terra.bio"`
	NamePrefixesDomain          *bool      `json:"namePrefixesDomain" form:"namePrefixesDomain" default:"true"`
	HelmfileRef                 *string    `json:"helmfileRef" form:"helmfileRef" default:"HEAD"`
	PreventDeletion             *bool      `json:"preventDeletion" form:"preventDeletion" default:"false"` // Used to protect specific BEEs from deletion (thelma checks this field)
	DeleteAfter                 *time.Time `json:"deleteAfter,omitempty" form:"deleteAfter"`               // If set, the BEE will be automatically deleted after this time (thelma checks this field)
	Description                 *string    `json:"description" form:"description"`
	PactIdentifier              *uuid.UUID `json:"pactIdentifier" form:"PactIdentifier"`
	PagerdutyIntegration        *string    `json:"pagerdutyIntegration,omitempty" form:"pagerdutyIntegration"`
	Offline                     *bool      `json:"offline" form:"offline" default:"false"`                                                 // Applicable for BEEs only, whether Thelma should render the BEE as "offline" zero replicas (this field is a target state, not a status)
	OfflineScheduleBeginEnabled *bool      `json:"offlineScheduleBeginEnabled,omitempty" form:"offlineScheduleBeginEnabled"`               // When enabled, the BEE will be slated to go offline around the begin time each day
	OfflineScheduleBeginTime    *time.Time `json:"offlineScheduleBeginTime,omitempty" form:"offlineScheduleBeginTime"  format:"date-time"` // Stored with timezone to determine day of the week
	OfflineScheduleEndEnabled   *bool      `json:"offlineScheduleEndEnabled,omitempty" form:"offlineScheduleEndEnabled"`                   // When enabled, the BEE will be slated to come online around the end time each weekday (each day if weekends enabled)
	OfflineScheduleEndTime      *time.Time `json:"offlineScheduleEndTime,omitempty" form:"offlineScheduleEndTime"  format:"date-time"`     // Stored with timezone to determine day of the week
	OfflineScheduleEndWeekends  *bool      `json:"offlineScheduleEndWeekends,omitempty" form:"offlineScheduleEndWeekends"`
}

func (e EnvironmentV3) toModel(db *gorm.DB) (models.Environment, error) {
	ret := models.Environment{
		Model:                       e.CommonFields.toGormModel(),
		Base:                        e.Base,
		Lifecycle:                   e.Lifecycle,
		Name:                        e.Name,
		ValuesName:                  e.ValuesName,
		AutoPopulateChartReleases:   e.AutoPopulateChartReleases,
		UniqueResourcePrefix:        e.UniqueResourcePrefix,
		DefaultNamespace:            e.DefaultNamespace,
		DefaultFirecloudDevelopRef:  e.DefaultFirecloudDevelopRef,
		RequiresSuitability:         e.RequiresSuitability,
		BaseDomain:                  e.BaseDomain,
		NamePrefixesDomain:          e.NamePrefixesDomain,
		HelmfileRef:                 e.HelmfileRef,
		PreventDeletion:             e.PreventDeletion,
		Description:                 e.Description,
		Offline:                     e.Offline,
		OfflineScheduleBeginEnabled: e.OfflineScheduleBeginEnabled,
		OfflineScheduleBeginTime:    utils.TimePtrToISO8601(e.OfflineScheduleBeginTime),
		OfflineScheduleEndEnabled:   e.OfflineScheduleEndEnabled,
		OfflineScheduleEndTime:      utils.TimePtrToISO8601(e.OfflineScheduleEndTime),
		OfflineScheduleEndWeekends:  e.OfflineScheduleEndWeekends,
		PactIdentifier:              e.PactIdentifier,
	}
	if e.DeleteAfter != nil {
		ret.DeleteAfter = sql.NullTime{Time: *e.DeleteAfter, Valid: true}
	}
	if e.TemplateEnvironment != "" {
		templateEnvironmentModel, err := environmentModelFromSelector(e.TemplateEnvironment)
		if err != nil {
			return models.Environment{}, err
		}
		var templateEnvironment models.Environment
		if err = db.Where(&templateEnvironmentModel).Select("id").First(&templateEnvironment).Error; err != nil {
			return models.Environment{}, err
		} else {
			ret.TemplateEnvironmentID = &templateEnvironment.ID
		}
	}
	if e.DefaultCluster != nil && *e.DefaultCluster != "" {
		defaultClusterModel, err := clusterModelFromSelector(*e.DefaultCluster)
		if err != nil {
			return models.Environment{}, err
		}
		var defaultCluster models.Cluster
		if err = db.Where(&defaultClusterModel).Select("id").First(&defaultCluster).Error; err != nil {
			return models.Environment{}, err
		} else {
			ret.DefaultClusterID = &defaultCluster.ID
		}
	}
	if e.PagerdutyIntegration != nil && *e.PagerdutyIntegration != "" {
		pagerdutyIntegrationModel, err := pagerdutyIntegrationModelFromSelector(*e.PagerdutyIntegration)
		if err != nil {
			return models.Environment{}, err
		}
		var pagerdutyIntegration models.PagerdutyIntegration
		if err = db.Where(&pagerdutyIntegrationModel).Select("id").First(&pagerdutyIntegration).Error; err != nil {
			return models.Environment{}, err
		} else {
			ret.PagerdutyIntegrationID = &pagerdutyIntegration.ID
		}
	}
	if e.Owner != nil && *e.Owner != "" {
		ownerModel, err := userModelFromSelector(*e.Owner)
		if err != nil {
			return models.Environment{}, err
		}
		var owner models.User
		if err = db.Where(&ownerModel).Select("id").First(&owner).Error; err != nil {
			return models.Environment{}, err
		} else {
			ret.OwnerID = &owner.ID
		}
	}
	return ret, nil
}

func (e EnvironmentV3Create) toModel(db *gorm.DB) (models.Environment, error) {
	return EnvironmentV3{EnvironmentV3Create: e}.toModel(db)
}

func (e EnvironmentV3Edit) toModel(db *gorm.DB) (models.Environment, error) {
	return EnvironmentV3Create{EnvironmentV3Edit: e}.toModel(db)
}

func environmentFromModel(model models.Environment) EnvironmentV3 {
	ret := EnvironmentV3{
		CommonFields:             commonFieldsFromGormModel(model.Model),
		CiIdentifier:             utils.NilOrCall(ciIdentifierFromModel, model.CiIdentifier),
		TemplateEnvironmentInfo:  utils.NilOrCall(environmentFromModel, model.TemplateEnvironment),
		DefaultClusterInfo:       utils.NilOrCall(clusterFromModel, model.DefaultCluster),
		PagerdutyIntegrationInfo: utils.NilOrCall(pagerdutyIntegrationFromModel, model.PagerdutyIntegration),
		OwnerInfo:                utils.NilOrCall(userFromModel, model.Owner),
		EnvironmentV3Create: EnvironmentV3Create{
			Base:                      model.Base,
			AutoPopulateChartReleases: model.AutoPopulateChartReleases,
			Lifecycle:                 model.Lifecycle,
			Name:                      model.Name,
			UniqueResourcePrefix:      model.UniqueResourcePrefix,
			DefaultNamespace:          model.DefaultNamespace,
			ValuesName:                model.ValuesName,
			EnvironmentV3Edit: EnvironmentV3Edit{
				DefaultFirecloudDevelopRef:  model.DefaultFirecloudDevelopRef,
				RequiresSuitability:         model.RequiresSuitability,
				BaseDomain:                  model.BaseDomain,
				NamePrefixesDomain:          model.NamePrefixesDomain,
				HelmfileRef:                 model.HelmfileRef,
				PreventDeletion:             model.PreventDeletion,
				Description:                 model.Description,
				PactIdentifier:              model.PactIdentifier,
				Offline:                     model.Offline,
				OfflineScheduleBeginEnabled: model.OfflineScheduleBeginEnabled,
				OfflineScheduleEndEnabled:   model.OfflineScheduleEndEnabled,
				OfflineScheduleEndWeekends:  model.OfflineScheduleEndWeekends,
			},
		},
	}
	if model.TemplateEnvironment != nil {
		ret.TemplateEnvironment = model.TemplateEnvironment.Name
	}
	if model.DefaultCluster != nil {
		ret.DefaultCluster = &model.DefaultCluster.Name
	}
	if model.Owner != nil {
		ret.Owner = &model.Owner.Email
	}
	if model.DeleteAfter.Valid {
		ret.DeleteAfter = &model.DeleteAfter.Time
	}
	if model.PagerdutyIntegration != nil && model.PagerdutyIntegration.PagerdutyID != "" {
		ret.PagerdutyIntegration = utils.PointerTo(fmt.Sprintf("pd-id/%s", model.PagerdutyIntegration.PagerdutyID))
	} else if model.PagerdutyIntegrationID != nil {
		ret.PagerdutyIntegration = utils.PointerTo(utils.UintToString(*model.PagerdutyIntegrationID))
	}
	var err error
	if ret.OfflineScheduleBeginTime, err = utils.ISO8601PtrToTime(model.OfflineScheduleBeginTime); err != nil {
		log.Error().Uint("environment", model.ID).Err(err).Msg("failed to parse offline schedule begin time")
		ret.OfflineScheduleBeginTime = nil
	}
	if ret.OfflineScheduleEndTime, err = utils.ISO8601PtrToTime(model.OfflineScheduleEndTime); err != nil {
		log.Error().Uint("environment", model.ID).Err(err).Msg("failed to parse offline schedule end time")
		ret.OfflineScheduleEndTime = nil
	}
	return ret
}
