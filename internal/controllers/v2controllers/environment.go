package v2controllers

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/internal/models/v2models/environment"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	petname "github.com/dustinkirkland/golang-petname"
	"gorm.io/gorm"
)

type Environment struct {
	ReadableBaseType
	TemplateEnvironmentInfo  *Environment          `json:"templateEnvironmentInfo,omitempty" swaggertype:"object" form:"-"` // Single-layer recursive; provides info of the template environment if this environment has one
	DefaultClusterInfo       *Cluster              `json:"defaultClusterInfo,omitempty" form:"-"`
	PagerdutyIntegrationInfo *PagerdutyIntegration `json:"pagerdutyIntegrationInfo,omitempty" form:"-"`
	CreatableEnvironment
}

type CreatableEnvironment struct {
	Base                      string `json:"base" form:"base"`                                                          // Required when creating
	ChartReleasesFromTemplate *bool  `json:"chartReleasesFromTemplate" form:"chartReleasesFromTemplate" default:"true"` // Upon creation of a dynamic environment, if this is true the template's chart releases will be copied to the new environment
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
	DefaultCluster             *string                 `json:"defaultCluster" form:"defaultCluster"`
	DefaultFirecloudDevelopRef *string                 `json:"defaultFirecloudDevelopRef" form:"defaultFirecloudDevelopRef" default:"dev"` // should be the environment branch for live envs. Is usually dev for template/dynamic but not necessarily
	Owner                      *string                 `json:"owner" form:"owner"`                                                         // When creating, will be set to your email
	RequiresSuitability        *bool                   `json:"requiresSuitability" form:"requiresSuitability" default:"false"`
	BaseDomain                 *string                 `json:"baseDomain" form:"baseDomain" default:"bee.envs-terra.bio"`
	NamePrefixesDomain         *bool                   `json:"namePrefixesDomain" form:"namePrefixesDomain" default:"true"`
	HelmfileRef                *string                 `json:"helmfileRef" form:"helmfileRef" default:"HEAD"`
	PreventDeletion            *bool                   `json:"preventDeletion" form:"preventDeletion" default:"false"` // Used to protect specific BEEs from deletion (thelma checks this field)
	AutoDelete                 *environment.AutoDelete `json:"autoDelete" form:"autoDelete"`
	Description                *string                 `json:"description" form:"description"`
	PagerdutyIntegration       *string                 `json:"pagerdutyIntegration,omitempty" form:"pagerdutyIntegration"`
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
	return v2models.Environment{
		Model: gorm.Model{
			ID:        e.ID,
			CreatedAt: e.CreatedAt,
			UpdatedAt: e.UpdatedAt,
		},
		Base:                       e.Base,
		ChartReleasesFromTemplate:  e.ChartReleasesFromTemplate,
		Lifecycle:                  e.Lifecycle,
		Name:                       e.Name,
		TemplateEnvironmentID:      templateEnvironmentID,
		ValuesName:                 e.ValuesName,
		UniqueResourcePrefix:       e.UniqueResourcePrefix,
		DefaultClusterID:           defaultClusterID,
		DefaultNamespace:           e.DefaultNamespace,
		NamePrefix:                 e.NamePrefix,
		DefaultFirecloudDevelopRef: e.DefaultFirecloudDevelopRef,
		Owner:                      e.Owner,
		RequiresSuitability:        e.RequiresSuitability,
		BaseDomain:                 e.BaseDomain,
		NamePrefixesDomain:         e.NamePrefixesDomain,
		HelmfileRef:                e.HelmfileRef,
		PreventDeletion:            e.PreventDeletion,
		AutoDelete:                 e.AutoDelete,
		Description:                e.Description,
		PagerdutyIntegrationID:     pagerdutyIntegrationID,
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

	return &Environment{
		ReadableBaseType: ReadableBaseType{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
		TemplateEnvironmentInfo:  templateEnvironment,
		DefaultClusterInfo:       defaultCluster,
		PagerdutyIntegrationInfo: pagerdutyIntegration,
		CreatableEnvironment: CreatableEnvironment{
			Base:                      model.Base,
			ChartReleasesFromTemplate: model.ChartReleasesFromTemplate,
			Lifecycle:                 model.Lifecycle,
			Name:                      model.Name,
			TemplateEnvironment:       templateEnvironmentName,
			UniqueResourcePrefix:      model.UniqueResourcePrefix,
			DefaultNamespace:          model.DefaultNamespace,
			NamePrefix:                model.NamePrefix,
			ValuesName:                model.ValuesName,
			EditableEnvironment: EditableEnvironment{
				DefaultCluster:             &defaultClusterName,
				DefaultFirecloudDevelopRef: model.DefaultFirecloudDevelopRef,
				Owner:                      model.Owner,
				RequiresSuitability:        model.RequiresSuitability,
				BaseDomain:                 model.BaseDomain,
				NamePrefixesDomain:         model.NamePrefixesDomain,
				HelmfileRef:                model.HelmfileRef,
				PreventDeletion:            model.PreventDeletion,
				AutoDelete:                 model.AutoDelete,
				Description:                model.Description,
				PagerdutyIntegration:       &pagerdutyIntegrationID,
			},
		},
	}
}

// setEnvironmentDynamicDefaults doesn't need to worry about validation, nor does it need to worry about any
// static defaults defined in struct tags. The model handles validation, and the caller will handle struct tags
// after this function runs.
func setEnvironmentDynamicDefaults(environment *CreatableEnvironment, stores *v2models.StoreSet, user *auth.User) error {
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
			rand.Seed(time.Now().UnixNano())
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
		environment.Owner = &user.AuthenticatedEmail
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
