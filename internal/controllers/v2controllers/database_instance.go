package v2controllers

import (
	"github.com/broadinstitute/sherlock/internal/auth/auth_models"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"gorm.io/gorm"
)

type DatabaseInstance struct {
	ReadableBaseType
	ChartReleaseInfo *ChartRelease `json:"chartReleaseInfo,omitempty" form:"-"`
	CreatableDatabaseInstance
}

type CreatableDatabaseInstance struct {
	ChartRelease string `json:"chartRelease" form:"chartRelease"` // Required when creating
	EditableDatabaseInstance
}

type EditableDatabaseInstance struct {
	Platform                  *string `json:"platform" form:"platform" default:"kubernetes"`              // 'google', 'azure', or default 'kubernetes'
	GoogleProject             *string `json:"googleProject" form:"googleProject"`                         // Required if platform is 'google'
	GoogleLocation            *string `json:"googleLocation" form:"googleLocation"`                       // Required if platform is 'google'
	AzureSubscription         *string `json:"azureSubscription" form:"azureSubscription"`                 // Required if platform is 'azure'
	AzureManagedResourceGroup *string `json:"azureManagedResourceGroup" form:"azureManagedResourceGroup"` // Required if platform is 'azure'
	InstanceName              *string `json:"instanceName" form:"instanceName"`                           // Required if platform is 'google' or 'azure'

	DefaultDatabase *string `json:"defaultDatabase" form:"defaultDatabase" ` // When creating, defaults to the chart name
}

//nolint:unused
func (d DatabaseInstance) toModel(storeSet *v2models.StoreSet) (v2models.DatabaseInstance, error) {
	var chartReleaseID uint
	if d.ChartRelease != "" {
		chartRelease, err := storeSet.ChartReleaseStore.Get(d.ChartRelease)
		if err != nil {
			return v2models.DatabaseInstance{}, err
		}
		chartReleaseID = chartRelease.ID
	}
	return v2models.DatabaseInstance{
		Model: gorm.Model{
			ID:        d.ID,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		},
		ChartReleaseID:            chartReleaseID,
		Platform:                  d.Platform,
		GoogleProject:             d.GoogleProject,
		GoogleLocation:            d.GoogleLocation,
		AzureSubscription:         d.AzureSubscription,
		AzureManagedResourceGroup: d.AzureManagedResourceGroup,
		InstanceName:              d.InstanceName,
		DefaultDatabase:           d.DefaultDatabase,
	}, nil
}

//nolint:unused
func (d CreatableDatabaseInstance) toModel(storeSet *v2models.StoreSet) (v2models.DatabaseInstance, error) {
	return DatabaseInstance{CreatableDatabaseInstance: d}.toModel(storeSet)
}

//nolint:unused
func (d EditableDatabaseInstance) toModel(storeSet *v2models.StoreSet) (v2models.DatabaseInstance, error) {
	return CreatableDatabaseInstance{EditableDatabaseInstance: d}.toModel(storeSet)
}

type DatabaseInstanceController = ModelController[v2models.DatabaseInstance, DatabaseInstance, CreatableDatabaseInstance, EditableDatabaseInstance]

func newDatabaseInstanceController(stores *v2models.StoreSet) *DatabaseInstanceController {
	return &DatabaseInstanceController{
		primaryStore:       stores.DatabaseInstanceStore,
		allStores:          stores,
		modelToReadable:    modelDatabaseInstanceToDatabaseInstance,
		setDynamicDefaults: setDatabaseInstanceDynamicDefaults,
	}
}

func modelDatabaseInstanceToDatabaseInstance(model *v2models.DatabaseInstance) *DatabaseInstance {
	if model == nil {
		return nil
	}

	var chartReleaseName string
	chartRelease := modelChartReleaseToChartRelease(model.ChartRelease)
	if chartRelease != nil {
		chartReleaseName = chartRelease.Name
	}

	return &DatabaseInstance{
		ReadableBaseType: ReadableBaseType{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
		ChartReleaseInfo: chartRelease,
		CreatableDatabaseInstance: CreatableDatabaseInstance{
			ChartRelease: chartReleaseName,
			EditableDatabaseInstance: EditableDatabaseInstance{
				Platform:                  model.Platform,
				GoogleProject:             model.GoogleProject,
				GoogleLocation:            model.GoogleLocation,
				AzureSubscription:         model.AzureSubscription,
				AzureManagedResourceGroup: model.AzureManagedResourceGroup,
				InstanceName:              model.InstanceName,
				DefaultDatabase:           model.DefaultDatabase,
			},
		},
	}
}

func setDatabaseInstanceDynamicDefaults(databaseInstance *CreatableDatabaseInstance, stores *v2models.StoreSet, _ *auth_models.User) error {
	if (databaseInstance.DefaultDatabase == nil || *databaseInstance.DefaultDatabase == "") && databaseInstance.ChartRelease != "" {
		chartRelease, err := stores.ChartReleaseStore.Get(databaseInstance.ChartRelease)
		if err != nil {
			return err
		}
		if chartRelease.Chart != nil {
			databaseInstance.DefaultDatabase = &chartRelease.Chart.Name
		}
	}
	return nil
}
