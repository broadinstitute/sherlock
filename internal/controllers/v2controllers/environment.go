package v2controllers

import (
	"fmt"
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
	TemplateEnvironmentInfo *Environment `json:"templateEnvironmentInfo,omitempty" swaggertype:"object" form:"-"` // Single-layer recursive; provides info of the template environment if this environment has one
	DefaultClusterInfo      *Cluster     `json:"defaultClusterInfo,omitempty" form:"-"`
	ValuesName              string       `json:"valuesName" form:"valuesName"`
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
	EditableEnvironment
}

type EditableEnvironment struct {
	DefaultCluster             *string `json:"defaultCluster" form:"defaultCluster"`
	DefaultFirecloudDevelopRef *string `json:"defaultFirecloudDevelopRef" form:"defaultFirecloudDevelopRef" default:"dev"` // should be the environment branch for live envs. Is usually dev for template/dynamic but not necessarily
	Owner                      *string `json:"owner" form:"owner"`                                                         // When creating, will be set to your email
	RequiresSuitability        *bool   `json:"requiresSuitability" form:"requiresSuitability" default:"false"`
	BaseDomain                 *string `json:"baseDomain" form:"baseDomain" default:"bee.envs-terra.bio"`
	NamePrefixesDomain         *bool   `json:"namePrefixesDomain" form:"namePrefixesDomain" default:"true"`
	HelmfileRef                *string `json:"helmfileRef" form:"helmfileRef" default:"HEAD"`
}

//nolint:unused
func (c CreatableEnvironment) toReadable() Environment {
	return Environment{CreatableEnvironment: c}
}

//nolint:unused
func (e EditableEnvironment) toCreatable() CreatableEnvironment {
	return CreatableEnvironment{EditableEnvironment: e}
}

type EnvironmentController = ModelController[v2models.Environment, Environment, CreatableEnvironment, EditableEnvironment]

func newEnvironmentController(stores *v2models.StoreSet) *EnvironmentController {
	return &EnvironmentController{
		primaryStore:       stores.EnvironmentStore,
		allStores:          stores,
		modelToReadable:    modelEnvironmentToEnvironment,
		readableToModel:    environmentToModelEnvironment,
		setDynamicDefaults: setEnvironmentDynamicDefaults,
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

	return &Environment{
		ReadableBaseType: ReadableBaseType{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
		TemplateEnvironmentInfo: templateEnvironment,
		DefaultClusterInfo:      defaultCluster,
		ValuesName:              model.ValuesName,
		CreatableEnvironment: CreatableEnvironment{
			Base:                      model.Base,
			ChartReleasesFromTemplate: model.ChartReleasesFromTemplate,
			Lifecycle:                 model.Lifecycle,
			Name:                      model.Name,
			TemplateEnvironment:       templateEnvironmentName,
			UniqueResourcePrefix:      model.UniqueResourcePrefix,
			DefaultNamespace:          model.DefaultNamespace,
			NamePrefix:                model.NamePrefix,
			EditableEnvironment: EditableEnvironment{
				DefaultCluster:             &defaultClusterName,
				DefaultFirecloudDevelopRef: model.DefaultFirecloudDevelopRef,
				Owner:                      model.Owner,
				RequiresSuitability:        model.RequiresSuitability,
				BaseDomain:                 model.BaseDomain,
				NamePrefixesDomain:         model.NamePrefixesDomain,
				HelmfileRef:                model.HelmfileRef,
			},
		},
	}
}

func environmentToModelEnvironment(environment Environment, stores *v2models.StoreSet) (v2models.Environment, error) {
	var templateEnvironmentID *uint
	if environment.TemplateEnvironment != "" {
		templateEnvironment, err := stores.EnvironmentStore.Get(environment.TemplateEnvironment)
		if err != nil {
			return v2models.Environment{}, err
		}
		templateEnvironmentID = &templateEnvironment.ID
	}
	var defaultClusterID *uint
	if environment.DefaultCluster != nil && *environment.DefaultCluster != "" {
		defaultCluster, err := stores.ClusterStore.Get(*environment.DefaultCluster)
		if err != nil {
			return v2models.Environment{}, err
		}
		defaultClusterID = &defaultCluster.ID
	}
	return v2models.Environment{
		Model: gorm.Model{
			ID:        environment.ID,
			CreatedAt: environment.CreatedAt,
			UpdatedAt: environment.UpdatedAt,
		},
		Base:                       environment.Base,
		ChartReleasesFromTemplate:  environment.ChartReleasesFromTemplate,
		Lifecycle:                  environment.Lifecycle,
		Name:                       environment.Name,
		TemplateEnvironmentID:      templateEnvironmentID,
		ValuesName:                 environment.ValuesName,
		UniqueResourcePrefix:       environment.UniqueResourcePrefix,
		DefaultClusterID:           defaultClusterID,
		DefaultNamespace:           environment.DefaultNamespace,
		NamePrefix:                 environment.NamePrefix,
		DefaultFirecloudDevelopRef: environment.DefaultFirecloudDevelopRef,
		Owner:                      environment.Owner,
		RequiresSuitability:        environment.RequiresSuitability,
		BaseDomain:                 environment.BaseDomain,
		NamePrefixesDomain:         environment.NamePrefixesDomain,
		HelmfileRef:                environment.HelmfileRef,
	}, nil
}

// setEnvironmentDynamicDefaults doesn't need to worry about validation, nor does it need to worry about any
// static defaults defined in struct tags. The model handles validation, and the caller will handle struct tags
// after this function runs.
func setEnvironmentDynamicDefaults(environment *Environment, stores *v2models.StoreSet, user *auth.User) error {
	if environment.TemplateEnvironment != "" {
		templateEnvironment, err := stores.EnvironmentStore.Get(environment.TemplateEnvironment)
		if err != nil {
			return err
		}
		// If there's a template, the valuesName to use is the name of the template
		environment.ValuesName = templateEnvironment.Name
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
