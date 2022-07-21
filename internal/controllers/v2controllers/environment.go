package v2controllers

import (
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"gorm.io/gorm"
)

type Environment struct {
	ReadableBaseType
	TemplateEnvironmentInfo *Environment `swaggertype:"object"` // Single-layer recursive; provides info of the template environment if this environment has one
	DefaultClusterInfo      *Cluster
	CreatableEnvironment
}

type CreatableEnvironment struct {
	Base                string
	Lifecycle           string
	Name                string
	TemplateEnvironment string
	ValuesName          string
	EditableEnvironment
}

type EditableEnvironment struct {
	DefaultCluster      *string
	DefaultNamespace    *string
	Owner               *string
	RequiresSuitability *bool
}

func (c CreatableEnvironment) toReadable() Environment {
	return Environment{CreatableEnvironment: c}
}

func (e EditableEnvironment) toCreatable() CreatableEnvironment {
	return CreatableEnvironment{EditableEnvironment: e}
}

type EnvironmentController = MutableModelController[v2models.Environment, Environment, CreatableEnvironment, EditableEnvironment]

func NewEnvironmentController(stores v2models.StoreSet) *EnvironmentController {
	return &EnvironmentController{
		ImmutableModelController[v2models.Environment, Environment, CreatableEnvironment]{
			primaryStore:    stores.EnvironmentStore,
			allStores:       stores,
			modelToReadable: modelEnvironmentToEnvironment,
			readableToModel: environmentToModelEnvironment,
		},
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
		CreatableEnvironment: CreatableEnvironment{
			Base:                model.Base,
			Lifecycle:           model.Lifecycle,
			Name:                model.Name,
			TemplateEnvironment: templateEnvironmentName,
			ValuesName:          model.ValuesName,
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
