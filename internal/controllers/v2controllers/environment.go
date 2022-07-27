package v2controllers

import (
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"github.com/creasty/defaults"
	"gorm.io/gorm"
	"strconv"
)

type Environment struct {
	ReadableBaseType
	TemplateEnvironmentInfo *Environment `json:"templateEnvironmentInfo,omitempty" swaggertype:"object" form:"templateEnvironmentInfo"` // Single-layer recursive; provides info of the template environment if this environment has one
	DefaultClusterInfo      *Cluster     `json:"defaultClusterInfo,omitempty" form:"defaultClusterInfo"`
	ValuesName              string       `json:"valuesName" form:"valuesName"`
	CreatableEnvironment
}

type CreatableEnvironment struct {
	Base                string `json:"base" form:"base"`
	Lifecycle           string `json:"lifecycle" default:"dynamic" form:"lifecycle"`
	Name                string `json:"name" form:"name"`
	TemplateEnvironment string `json:"templateEnvironment" form:"templateEnvironment"`
	EditableEnvironment
}

type EditableEnvironment struct {
	DefaultCluster      *string `json:"defaultCluster" form:"defaultCluster"`
	DefaultNamespace    *string `json:"defaultNamespace" form:"defaultNamespace"`
	Owner               *string `json:"owner" form:"owner"`
	RequiresSuitability *bool   `json:"requiresSuitability" default:"false" form:"requiresSuitability"`
}

func (c CreatableEnvironment) toReadable() Environment {
	return Environment{CreatableEnvironment: c}
}

func (e EditableEnvironment) toCreatable() CreatableEnvironment {
	return CreatableEnvironment{EditableEnvironment: e}
}

type EnvironmentController = ModelController[v2models.Environment, Environment, CreatableEnvironment, EditableEnvironment]

func NewEnvironmentController(stores v2models.StoreSet) *EnvironmentController {
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

func environmentToModelEnvironment(environment Environment, stores v2models.StoreSet) (v2models.Environment, error) {
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
func setEnvironmentDynamicDefaults(environment *Environment, stores v2models.StoreSet, user *auth.User) error {
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
		if environment.DefaultNamespace == nil {
			environment.DefaultNamespace = templateEnvironment.DefaultNamespace
		}
		if environment.RequiresSuitability == nil {
			environment.RequiresSuitability = templateEnvironment.RequiresSuitability
		}
	} else {
		// If there's no template, the valuesName to use is the name of this environment
		environment.ValuesName = environment.Name
	}
	if defaults.CanUpdate(environment.Owner) {
		environment.Owner = &user.AuthenticatedEmail
	}
	return nil
}
