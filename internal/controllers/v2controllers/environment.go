package v2controllers

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"github.com/creasty/defaults"
	"github.com/dustinkirkland/golang-petname"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Environment struct {
	ReadableBaseType
	TemplateEnvironmentInfo *Environment `json:"templateEnvironmentInfo,omitempty" swaggertype:"object" form:"templateEnvironmentInfo"` // Single-layer recursive; provides info of the template environment if this environment has one
	DefaultClusterInfo      *Cluster     `json:"defaultClusterInfo,omitempty" form:"defaultClusterInfo"`
	ValuesName              string       `json:"valuesName" form:"valuesName"`
	CreatableEnvironment
}

type CreatableEnvironment struct {
	Base                string `json:"base" form:"base"` // Required when creating
	Lifecycle           string `json:"lifecycle" form:"lifecycle" default:"dynamic"`
	Name                string `json:"name" form:"name"`                               // When creating, will be calculated if dynamic, required otherwise
	TemplateEnvironment string `json:"templateEnvironment" form:"templateEnvironment"` // Required for dynamic environments
	EditableEnvironment
}

type EditableEnvironment struct {
	DefaultCluster      *string `json:"defaultCluster" form:"defaultCluster"`
	DefaultNamespace    *string `json:"defaultNamespace" form:"defaultNamespace"`
	Owner               *string `json:"owner" form:"owner"` // When creating, will be set to your email
	RequiresSuitability *bool   `json:"requiresSuitability" default:"false" form:"requiresSuitability"`
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

func modelEnvironmentToEnvironment(model v2models.Environment) *Environment {
	var templateEnvironment *Environment
	var templateEnvironmentName string
	if model.TemplateEnvironment != nil {
		templateEnvironment = modelEnvironmentToEnvironment(*model.TemplateEnvironment)
		templateEnvironmentName = templateEnvironment.Name
	}
	var defaultCluster *Cluster
	var defaultClusterName string
	if model.DefaultCluster != nil {
		defaultCluster = modelClusterToCluster(*model.DefaultCluster)
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
			Base:                model.Base,
			Lifecycle:           model.Lifecycle,
			Name:                model.Name,
			TemplateEnvironment: templateEnvironmentName,
			EditableEnvironment: EditableEnvironment{
				DefaultCluster:      &defaultClusterName,
				DefaultNamespace:    model.DefaultNamespace,
				Owner:               model.Owner,
				RequiresSuitability: model.RequiresSuitability,
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
		Base:                  environment.Base,
		Lifecycle:             environment.Lifecycle,
		Name:                  environment.Name,
		TemplateEnvironmentID: templateEnvironmentID,
		ValuesName:            environment.ValuesName,
		DefaultClusterID:      defaultClusterID,
		DefaultNamespace:      environment.DefaultNamespace,
		Owner:                 environment.Owner,
		RequiresSuitability:   environment.RequiresSuitability,
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
			rand.Seed(time.Now().UnixNano())
			for suffixLength := 3; suffixLength >= 1; suffixLength-- {
				environment.Name = fmt.Sprintf("%s-%s-%s", user.AlphaNumericHyphenatedUsername(), templateEnvironment.Name, petname.Generate(suffixLength, "-"))
				if len(environment.Name) <= 32 {
					break
				}
			}
			if len(environment.Name) > 32 {
				environment.Name = strings.TrimSuffix(environment.Name[0:31], "-")
			}
		}
	} else {
		// If there's no template, the valuesName to use is the name of this environment
		environment.ValuesName = environment.Name
	}
	if environment.DefaultNamespace == nil {
		environment.DefaultNamespace = &environment.Name
	}
	if defaults.CanUpdate(environment.Owner) {
		environment.Owner = &user.AuthenticatedEmail
	}
	return nil
}
