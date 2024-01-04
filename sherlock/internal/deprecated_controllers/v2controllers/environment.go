package v2controllers

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models/environment"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"strconv"
	"strings"
	"time"

	petname "github.com/dustinkirkland/golang-petname"
	"gorm.io/gorm"
)

type Environment struct {
	ReadableBaseType
	CiIdentifier             *CiIdentifier         `json:"ciIdentifier,omitempty" form:"-"`
	TemplateEnvironmentInfo  *Environment          `json:"templateEnvironmentInfo,omitempty" swaggertype:"object" form:"-"` // Single-layer recursive; provides info of the template environment if this environment has one
	DefaultClusterInfo       *Cluster              `json:"defaultClusterInfo,omitempty" form:"-"`
	PagerdutyIntegrationInfo *PagerdutyIntegration `json:"pagerdutyIntegrationInfo,omitempty" form:"-"`
	OwnerInfo                *User                 `json:"ownerInfo,omitempty" form:"-"`
	CreatableEnvironment
}

type CreatableEnvironment struct {
	Base                      string `json:"base" form:"base"`                                                          // Required when creating
	AutoPopulateChartReleases *bool  `json:"autoPopulateChartReleases" form:"autoPopulateChartReleases" default:"true"` // If true when creating, dynamic environments copy from template and template environments get the honeycomb chart
	Lifecycle                 string `json:"lifecycle" form:"lifecycle" default:"dynamic"`
	Name                      string `json:"name" form:"name"`                                 // When creating, will be calculated if dynamic, required otherwise
	TemplateEnvironment       string `json:"templateEnvironment" form:"templateEnvironment"`   // Required for dynamic environments
	UniqueResourcePrefix      string `json:"uniqueResourcePrefix" form:"uniqueResourcePrefix"` // When creating, will be calculated if left empty
	DefaultNamespace          string `json:"defaultNamespace" form:"defaultNamespace"`         // When creating, will be calculated if left empty
	NamePrefix                string `json:"namePrefix" form:"namePrefix"`                     // Used for dynamic environment name generation only, to override using the owner email handle and template name
	ValuesName                string `json:"valuesName" form:"valuesName"`                     // When creating, defaults to template name or environment name
	EditableEnvironment
}

type EditableEnvironment struct {
	DefaultCluster              *string                 `json:"defaultCluster" form:"defaultCluster"`
	DefaultFirecloudDevelopRef  *string                 `json:"defaultFirecloudDevelopRef" form:"defaultFirecloudDevelopRef" default:"dev"` // should be the environment branch for live envs. Is usually dev for template/dynamic but not necessarily
	Owner                       *string                 `json:"owner" form:"owner"`                                                         // When creating, will default to you
	RequiresSuitability         *bool                   `json:"requiresSuitability" form:"requiresSuitability" default:"false"`
	BaseDomain                  *string                 `json:"baseDomain" form:"baseDomain" default:"bee.envs-terra.bio"`
	NamePrefixesDomain          *bool                   `json:"namePrefixesDomain" form:"namePrefixesDomain" default:"true"`
	HelmfileRef                 *string                 `json:"helmfileRef" form:"helmfileRef" default:"HEAD"`
	PreventDeletion             *bool                   `json:"preventDeletion" form:"preventDeletion" default:"false"` // Used to protect specific BEEs from deletion (thelma checks this field)
	AutoDelete                  *environment.AutoDelete `json:"autoDelete" form:"autoDelete"`
	Description                 *string                 `json:"description" form:"description"`
	PactIdentifier              *uuid.UUID              `json:"pactIdentifier" form:"PactIdentifier" default:""`
	PagerdutyIntegration        *string                 `json:"pagerdutyIntegration,omitempty" form:"pagerdutyIntegration"`
	Offline                     *bool                   `json:"offline" form:"offline" default:"false"`                                                 // Applicable for BEEs only, whether Thelma should render the BEE as "offline" zero replicas (this field is a target state, not a status)
	OfflineScheduleBeginEnabled *bool                   `json:"offlineScheduleBeginEnabled,omitempty" form:"offlineScheduleBeginEnabled"`               // When enabled, the BEE will be slated to go offline around the begin time each day
	OfflineScheduleBeginTime    *time.Time              `json:"offlineScheduleBeginTime,omitempty" form:"offlineScheduleBeginTime"  format:"date-time"` // Stored with timezone to determine day of the week
	OfflineScheduleEndEnabled   *bool                   `json:"offlineScheduleEndEnabled,omitempty" form:"offlineScheduleEndEnabled"`                   // When enabled, the BEE will be slated to come online around the end time each weekday (each day if weekends enabled)
	OfflineScheduleEndTime      *time.Time              `json:"offlineScheduleEndTime,omitempty" form:"offlineScheduleEndTime"  format:"date-time"`     // Stored with timezone to determine day of the week
	OfflineScheduleEndWeekends  *bool                   `json:"offlineScheduleEndWeekends,omitempty" form:"offlineScheduleEndWeekends"`
}

//nolint:unused
func (e Environment) toModel(storeSet *v2models.StoreSet) (v2models.Environment, error) {
	var templateEnvironmentID *uint
	if e.TemplateEnvironment != "" {
		templateEnvironment, err := storeSet.EnvironmentStore.Get(e.TemplateEnvironment)
		if err != nil {
			return v2models.Environment{}, err
		}
		templateEnvironmentID = &templateEnvironment.ID
	}
	var defaultClusterID *uint
	if e.DefaultCluster != nil && *e.DefaultCluster != "" {
		defaultCluster, err := storeSet.ClusterStore.Get(*e.DefaultCluster)
		if err != nil {
			return v2models.Environment{}, err
		}
		defaultClusterID = &defaultCluster.ID
	}
	var pagerdutyIntegrationID *uint
	if e.PagerdutyIntegration != nil && *e.PagerdutyIntegration != "" {
		pagerdutyIntegration, err := storeSet.PagerdutyIntegration.Get(*e.PagerdutyIntegration)
		if err != nil {
			return v2models.Environment{}, err
		}
		pagerdutyIntegrationID = &pagerdutyIntegration.ID
	}

	// The model has LegacyOwner from when we just stored a string. We'll read that out into Owner if there's no
	// OwnerID in the model, but we won't read back in to LegacyOwner. In other words, that field of the model
	// is becoming read-only, and it won't ever be visible if there's an owner reference.
	var ownerID *uint
	if e.Owner != nil {
		user, err := storeSet.UserStore.Get(*e.Owner)
		if err != nil {
			return v2models.Environment{}, err
		}
		ownerID = &user.ID
	}
	return v2models.Environment{
		Model: gorm.Model{
			ID:        e.ID,
			CreatedAt: e.CreatedAt,
			UpdatedAt: e.UpdatedAt,
		},
		Base:                        e.Base,
		AutoPopulateChartReleases:   e.AutoPopulateChartReleases,
		Lifecycle:                   e.Lifecycle,
		Name:                        e.Name,
		TemplateEnvironmentID:       templateEnvironmentID,
		ValuesName:                  e.ValuesName,
		UniqueResourcePrefix:        e.UniqueResourcePrefix,
		DefaultClusterID:            defaultClusterID,
		DefaultNamespace:            e.DefaultNamespace,
		NamePrefix:                  e.NamePrefix,
		DefaultFirecloudDevelopRef:  e.DefaultFirecloudDevelopRef,
		OwnerID:                     ownerID,
		RequiresSuitability:         e.RequiresSuitability,
		BaseDomain:                  e.BaseDomain,
		NamePrefixesDomain:          e.NamePrefixesDomain,
		HelmfileRef:                 e.HelmfileRef,
		PreventDeletion:             e.PreventDeletion,
		AutoDelete:                  e.AutoDelete,
		Description:                 e.Description,
		PactIdentifier:              e.PactIdentifier,
		PagerdutyIntegrationID:      pagerdutyIntegrationID,
		Offline:                     e.Offline,
		OfflineScheduleBeginEnabled: e.OfflineScheduleBeginEnabled,
		OfflineScheduleBeginTime:    utils.TimePtrToISO8601(e.OfflineScheduleBeginTime),
		OfflineScheduleEndEnabled:   e.OfflineScheduleEndEnabled,
		OfflineScheduleEndTime:      utils.TimePtrToISO8601(e.OfflineScheduleEndTime),
		OfflineScheduleEndWeekends:  e.OfflineScheduleEndWeekends,
	}, nil
}

//nolint:unused
func (e CreatableEnvironment) toModel(storeSet *v2models.StoreSet) (v2models.Environment, error) {
	return Environment{CreatableEnvironment: e}.toModel(storeSet)
}

//nolint:unused
func (e EditableEnvironment) toModel(storeSet *v2models.StoreSet) (v2models.Environment, error) {
	return CreatableEnvironment{EditableEnvironment: e}.toModel(storeSet)
}

type EnvironmentController = ModelController[v2models.Environment, Environment, CreatableEnvironment, EditableEnvironment]

func newEnvironmentController(stores *v2models.StoreSet) *EnvironmentController {
	return &EnvironmentController{
		primaryStore:                   stores.EnvironmentStore,
		allStores:                      stores,
		modelToReadable:                modelEnvironmentToEnvironment,
		setDynamicDefaults:             setEnvironmentDynamicDefaults,
		beehiveUrlFormatString:         config.Config.MustString("beehive.environmentUrlFormatString"),
		extractPagerdutyIntegrationKey: extractPagerdutyIntegrationKeyFromEnvironment,
	}
}

func modelEnvironmentToEnvironment(model *v2models.Environment) *Environment {
	if model == nil {
		return nil
	}

	var templateEnvironmentName string
	templateEnvironment := modelEnvironmentToEnvironment(model.TemplateEnvironment)
	if templateEnvironment != nil {
		templateEnvironmentName = templateEnvironment.Name
	}

	var defaultClusterName string
	defaultCluster := modelClusterToCluster(model.DefaultCluster)
	if defaultCluster != nil {
		defaultClusterName = defaultCluster.Name
	}
	var pagerdutyIntegrationID string
	pagerdutyIntegration := modelPagerdutyIntegrationToPagerdutyIntegration(model.PagerdutyIntegration)
	if pagerdutyIntegration != nil {
		pagerdutyIntegrationID = strconv.FormatUint(uint64(pagerdutyIntegration.ID), 10)
	} else if model.PagerdutyIntegrationID != nil {
		pagerdutyIntegrationID = strconv.FormatUint(uint64(*model.PagerdutyIntegrationID), 10)
	}
	offlineScheduleBeginTime, err := utils.ISO8601PtrToTime(model.OfflineScheduleBeginTime)
	if err != nil {
		log.Error().Err(err).Msgf("couldn't parse %s's offlineScheduleBeginTime coming from the database, swallowing: %v", model.Name, err)
	}
	offlineScheduleEndTime, err := utils.ISO8601PtrToTime(model.OfflineScheduleEndTime)
	if err != nil {
		log.Error().Err(err).Msgf("couldn't parse %s's offlineScheduleEndTime coming from the database, swallowing: %v", model.Name, err)
	}

	var ownerEmail string
	owner := modelUserToUser(model.Owner)
	if owner != nil {
		ownerEmail = owner.Email
	} else if model.LegacyOwner != nil {
		ownerEmail = *model.LegacyOwner
	}

	return &Environment{
		ReadableBaseType: ReadableBaseType{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
		CiIdentifier:             modelCiIdentifierToCiIdentifier(model.CiIdentifier),
		TemplateEnvironmentInfo:  templateEnvironment,
		DefaultClusterInfo:       defaultCluster,
		PagerdutyIntegrationInfo: pagerdutyIntegration,
		OwnerInfo:                owner,
		CreatableEnvironment: CreatableEnvironment{
			Base:                      model.Base,
			AutoPopulateChartReleases: model.AutoPopulateChartReleases,
			Lifecycle:                 model.Lifecycle,
			Name:                      model.Name,
			TemplateEnvironment:       templateEnvironmentName,
			UniqueResourcePrefix:      model.UniqueResourcePrefix,
			DefaultNamespace:          model.DefaultNamespace,
			NamePrefix:                model.NamePrefix,
			ValuesName:                model.ValuesName,
			EditableEnvironment: EditableEnvironment{
				DefaultCluster:              &defaultClusterName,
				DefaultFirecloudDevelopRef:  model.DefaultFirecloudDevelopRef,
				Owner:                       &ownerEmail,
				RequiresSuitability:         model.RequiresSuitability,
				BaseDomain:                  model.BaseDomain,
				NamePrefixesDomain:          model.NamePrefixesDomain,
				HelmfileRef:                 model.HelmfileRef,
				PreventDeletion:             model.PreventDeletion,
				AutoDelete:                  model.AutoDelete,
				Description:                 model.Description,
				PactIdentifier:              model.PactIdentifier,
				PagerdutyIntegration:        &pagerdutyIntegrationID,
				Offline:                     model.Offline,
				OfflineScheduleBeginEnabled: model.OfflineScheduleBeginEnabled,
				OfflineScheduleBeginTime:    offlineScheduleBeginTime,
				OfflineScheduleEndEnabled:   model.OfflineScheduleEndEnabled,
				OfflineScheduleEndTime:      offlineScheduleEndTime,
				OfflineScheduleEndWeekends:  model.OfflineScheduleEndWeekends,
			},
		},
	}
}

// setEnvironmentDynamicDefaults doesn't need to worry about validation, nor does it need to worry about any
// static defaults defined in struct tags. The model handles validation, and the caller will handle struct tags
// after this function runs.
func setEnvironmentDynamicDefaults(environment *CreatableEnvironment, stores *v2models.StoreSet, user *models.User) error {
	if environment.TemplateEnvironment != "" {
		templateEnvironment, err := stores.EnvironmentStore.Get(environment.TemplateEnvironment)
		if err != nil {
			return err
		}
		if environment.ValuesName == "" {
			// If there's a template, the valuesName to use is the name of the template
			environment.ValuesName = templateEnvironment.Name
		}
		if environment.Base == "" {
			environment.Base = templateEnvironment.Base
		}
		if environment.DefaultCluster == nil && templateEnvironment.DefaultClusterID != nil {
			id := strconv.FormatUint(uint64(*templateEnvironment.DefaultClusterID), 10)
			environment.DefaultCluster = &id
		}
		if environment.RequiresSuitability == nil {
			environment.RequiresSuitability = templateEnvironment.RequiresSuitability
		}
		if environment.Name == "" {
			namePrefix := environment.NamePrefix
			if namePrefix == "" {
				namePrefix = fmt.Sprintf("%s-%s", user.AlphaNumericHyphenatedUsername(), templateEnvironment.Name)
			}
			for suffixLength := 3; suffixLength >= 1; suffixLength-- {
				environment.Name = fmt.Sprintf("%s-%s", namePrefix, petname.Generate(suffixLength, "-"))
				if len(environment.Name) <= 32 {
					break
				}
			}
			if len(environment.Name) > 32 {
				environment.Name = strings.TrimSuffix(environment.Name[0:31], "-")
			}
		}
		if environment.BaseDomain == nil {
			environment.BaseDomain = templateEnvironment.BaseDomain
		}
		if environment.NamePrefixesDomain == nil {
			environment.NamePrefixesDomain = templateEnvironment.NamePrefixesDomain
		}
		// if a default firecloud develop ref is not specified check the template
		if environment.DefaultFirecloudDevelopRef == nil && templateEnvironment.DefaultFirecloudDevelopRef != nil {
			environment.DefaultFirecloudDevelopRef = templateEnvironment.DefaultFirecloudDevelopRef
		}

	} else {
		// If there's no template, the valuesName to use is the name of this environment
		environment.ValuesName = environment.Name
	}
	if environment.DefaultNamespace == "" {
		environment.DefaultNamespace = fmt.Sprintf("terra-%s", environment.Name)
	}
	if environment.Owner == nil {
		environment.Owner = &user.Email
	}

	// set default firecloud-develop ref for live terra envs
	if environment.DefaultFirecloudDevelopRef == nil {
		// if we are in a live terra env, the fc-develop ref should be the env name
		if environment.Lifecycle == "static" && environment.Base == "live" {
			environment.DefaultFirecloudDevelopRef = &environment.Name
		}
	}

	return nil
}

func extractPagerdutyIntegrationKeyFromEnvironment(model *v2models.Environment) *string {
	if model != nil && model.PagerdutyIntegration != nil {
		return model.PagerdutyIntegration.Key
	}
	return nil
}
