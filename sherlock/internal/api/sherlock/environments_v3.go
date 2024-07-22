package sherlock

import (
	"database/sql"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"strings"
	"time"
)

type EnvironmentV3 struct {
	CommonFields
	CiIdentifier             *CiIdentifierV3         `json:"ciIdentifier,omitempty" form:"-"`
	TemplateEnvironmentInfo  *EnvironmentV3          `json:"templateEnvironmentInfo,omitempty" swaggertype:"object" form:"-"`
	DefaultClusterInfo       *ClusterV3              `json:"defaultClusterInfo,omitempty" form:"-"`
	PagerdutyIntegrationInfo *PagerdutyIntegrationV3 `json:"pagerdutyIntegrationInfo,omitempty" form:"-"`
	OwnerInfo                *UserV3                 `json:"ownerInfo,omitempty" form:"-"`
	RequiredRoleInfo         *RoleV3                 `json:"requiredRoleInfo,omitempty" form:"-"`
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
	ValuesName                string `json:"valuesName" form:"valuesName"`                     // When creating, defaults to template name or environment name
	EnvironmentV3Edit
}

type EnvironmentV3Edit struct {
	DefaultCluster              *string    `json:"defaultCluster" form:"defaultCluster"`
	Owner                       *string    `json:"owner" form:"owner"` // When creating, will default to you
	RequiresSuitability         *bool      `json:"requiresSuitability" form:"requiresSuitability"`
	RequiredRole                *string    `json:"requiredRole" form:"requiredRole"` // If present, requires membership in the given role for mutations. Set to an empty string to clear.
	BaseDomain                  *string    `json:"baseDomain" form:"baseDomain" default:"bee.envs-terra.bio"`
	NamePrefixesDomain          *bool      `json:"namePrefixesDomain" form:"namePrefixesDomain" default:"true"`
	HelmfileRef                 *string    `json:"helmfileRef" form:"helmfileRef" default:"HEAD"`
	PreventDeletion             *bool      `json:"preventDeletion" form:"preventDeletion" default:"false"`      // Used to protect specific BEEs from deletion (thelma checks this field)
	DeleteAfter                 *string    `json:"deleteAfter,omitempty" form:"deleteAfter" format:"date-time"` // If set, the BEE will be automatically deleted after this time. Can be set to "" or Go's zero time value to clear the field.
	Description                 *string    `json:"description" form:"description"`
	PactIdentifier              *uuid.UUID `json:"pactIdentifier" form:"PactIdentifier"`
	PagerdutyIntegration        *string    `json:"pagerdutyIntegration,omitempty" form:"pagerdutyIntegration"`
	Offline                     *bool      `json:"offline" form:"offline" default:"false"`                                                 // Applicable for BEEs only, whether Thelma should render the BEE as "offline" zero replicas (this field is a target state, not a status)
	OfflineScheduleBeginEnabled *bool      `json:"offlineScheduleBeginEnabled,omitempty" form:"offlineScheduleBeginEnabled"`               // When enabled, the BEE will be slated to go offline around the begin time each day
	OfflineScheduleBeginTime    *time.Time `json:"offlineScheduleBeginTime,omitempty" form:"offlineScheduleBeginTime"  format:"date-time"` // Stored with timezone to determine day of the week
	OfflineScheduleEndEnabled   *bool      `json:"offlineScheduleEndEnabled,omitempty" form:"offlineScheduleEndEnabled"`                   // When enabled, the BEE will be slated to come online around the end time each weekday (each day if weekends enabled)
	OfflineScheduleEndTime      *time.Time `json:"offlineScheduleEndTime,omitempty" form:"offlineScheduleEndTime"  format:"date-time"`     // Stored with timezone to determine day of the week
	OfflineScheduleEndWeekends  *bool      `json:"offlineScheduleEndWeekends,omitempty" form:"offlineScheduleEndWeekends"`
	EnableJanitor               *bool      `json:"enableJanitor,omitempty" form:"enableJanitor"` // If true, janitor resource cleanup will be enabled for this environment. BEEs default to template's value, templates default to true, and static/live environments default to false.
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
		RequiresSuitability:         e.RequiresSuitability,
		BaseDomain:                  e.BaseDomain,
		NamePrefixesDomain:          e.NamePrefixesDomain,
		HelmfileRef:                 utils.NilOrCall(strings.TrimSpace, e.HelmfileRef),
		PreventDeletion:             e.PreventDeletion,
		Description:                 e.Description,
		Offline:                     e.Offline,
		OfflineScheduleBeginEnabled: e.OfflineScheduleBeginEnabled,
		OfflineScheduleBeginTime:    utils.TimePtrToISO8601(e.OfflineScheduleBeginTime),
		OfflineScheduleEndEnabled:   e.OfflineScheduleEndEnabled,
		OfflineScheduleEndTime:      utils.TimePtrToISO8601(e.OfflineScheduleEndTime),
		OfflineScheduleEndWeekends:  e.OfflineScheduleEndWeekends,
		PactIdentifier:              e.PactIdentifier,
		EnableJanitor:               e.EnableJanitor,
	}
	if e.DeleteAfter != nil {
		if *e.DeleteAfter == "" {
			// DeleteAfter explicitly set to an empty string; this means the field should be cleared.
			// We do that by explicitly storing the zero timestamp in the database
			ret.DeleteAfter = sql.NullTime{Time: time.Time{}, Valid: true}
		} else {
			deleteAfter, err := time.Parse(time.RFC3339, *e.DeleteAfter)
			if err != nil {
				return models.Environment{}, fmt.Errorf("(%s) failed to parse deleteAfter '%s': %w", errors.BadRequest, *e.DeleteAfter, err)
			}
			ret.DeleteAfter = sql.NullTime{Time: deleteAfter, Valid: true}
		}
	}
	if e.TemplateEnvironment != "" {
		templateEnvironmentModel, err := environmentModelFromSelector(e.TemplateEnvironment)
		if err != nil {
			return models.Environment{}, err
		}
		var templateEnvironment models.Environment
		if err = db.Where(&templateEnvironmentModel).Select("id").First(&templateEnvironment).Error; err != nil {
			return models.Environment{}, fmt.Errorf("template environment '%s' not found: %w", e.TemplateEnvironment, err)
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
			return models.Environment{}, fmt.Errorf("default cluster '%s' not found: %w", *e.DefaultCluster, err)
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
			return models.Environment{}, fmt.Errorf("pagerduty integration '%s' not found: %w", *e.PagerdutyIntegration, err)
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
			return models.Environment{}, fmt.Errorf("owner '%s' not found: %w", *e.Owner, err)
		} else {
			ret.OwnerID = &owner.ID
		}
	}
	if e.RequiredRole != nil && *e.RequiredRole != "" {
		requiredRoleModel, err := roleModelFromSelector(*e.RequiredRole)
		if err != nil {
			return models.Environment{}, err
		}
		var requiredRole models.Role
		if err = db.Where(&requiredRoleModel).Select("id").First(&requiredRole).Error; err != nil {
			return models.Environment{}, fmt.Errorf("required role '%s' not found: %w", *e.RequiredRole, err)
		} else {
			ret.RequiredRoleID = &requiredRole.ID
		}
	}
	return ret, nil
}

//nolint:unused
func (e EnvironmentV3Create) toModel(db *gorm.DB) (models.Environment, error) {
	return EnvironmentV3{EnvironmentV3Create: e}.toModel(db)
}

//nolint:unused
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
		RequiredRoleInfo:         utils.NilOrCall(roleFromModel, model.RequiredRole),
		EnvironmentV3Create: EnvironmentV3Create{
			Base:                      model.Base,
			AutoPopulateChartReleases: model.AutoPopulateChartReleases,
			Lifecycle:                 model.Lifecycle,
			Name:                      model.Name,
			UniqueResourcePrefix:      model.UniqueResourcePrefix,
			DefaultNamespace:          model.DefaultNamespace,
			ValuesName:                model.ValuesName,
			EnvironmentV3Edit: EnvironmentV3Edit{
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
				EnableJanitor:               model.EnableJanitor,
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
	if !model.DeleteAfter.Time.IsZero() {
		ret.DeleteAfter = utils.PointerTo(model.DeleteAfter.Time.Format(time.RFC3339))
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
	if model.RequiredRole != nil && model.RequiredRole.Name != nil {
		ret.RequiredRole = model.RequiredRole.Name
	} else if model.RequiredRoleID != nil {
		ret.RequiredRole = utils.PointerTo(utils.UintToString(*model.RequiredRoleID))
	}
	return ret
}
