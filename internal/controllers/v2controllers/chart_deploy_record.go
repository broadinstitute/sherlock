package v2controllers

import (
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"gorm.io/gorm"
)

type ChartDeployRecord struct {
	ReadableBaseType
	ChartReleaseInfo ChartRelease `json:"chartReleaseInfo" form:"chartReleaseInfo"`
	CreatableChartDeployRecord
}

type CreatableChartDeployRecord struct {
	ChartRelease      string `json:"chartRelease" form:"chartRelease"`           // Required when creating
	ExactChartVersion string `json:"exactChartVersion" form:"exactChartVersion"` // When creating, will default to the value currently held by the chart release
	ExactAppVersion   string `json:"exactAppVersion" form:"exactAppVersion"`     // When creating, will default to the value currently held by the chart release
	HelmfileRef       string `json:"helmfileRef" form:"helmfileRef"`             // When creating, will default to the value currently held by the chart release
	EditableChartDeployRecord
}

type EditableChartDeployRecord struct{}

//nolint:unused
func (c CreatableChartDeployRecord) toReadable() ChartDeployRecord {
	return ChartDeployRecord{CreatableChartDeployRecord: c}
}

//nolint:unused
func (e EditableChartDeployRecord) toCreatable() CreatableChartDeployRecord {
	return CreatableChartDeployRecord{EditableChartDeployRecord: e}
}

type ChartDeployRecordController = ModelController[v2models.ChartDeployRecord, ChartDeployRecord, CreatableChartDeployRecord, EditableChartDeployRecord]

func newChartDeployRecordController(stores *v2models.StoreSet) *ChartDeployRecordController {
	return &ChartDeployRecordController{
		primaryStore:       stores.ChartDeployRecordStore,
		allStores:          stores,
		modelToReadable:    modelChartDeployRecordToChartDeployRecord,
		readableToModel:    chartDeployRecordToModelChartDeployRecord,
		setDynamicDefaults: setChartDeployRecordDynamicDefaults,
	}
}

func modelChartDeployRecordToChartDeployRecord(model v2models.ChartDeployRecord) *ChartDeployRecord {
	chartRelease := modelChartReleaseToChartRelease(model.ChartRelease)
	return &ChartDeployRecord{
		ReadableBaseType: ReadableBaseType{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
		ChartReleaseInfo: *chartRelease,
		CreatableChartDeployRecord: CreatableChartDeployRecord{
			ChartRelease:      chartRelease.Name,
			ExactChartVersion: model.ExactChartVersion,
			ExactAppVersion:   model.ExactAppVersion,
			HelmfileRef:       model.HelmfileRef,
		},
	}
}

func chartDeployRecordToModelChartDeployRecord(chartDeployRecord ChartDeployRecord, stores *v2models.StoreSet) (v2models.ChartDeployRecord, error) {
	var chartReleaseID uint
	if chartDeployRecord.ChartRelease != "" {
		chartRelease, err := stores.ChartReleaseStore.Get(chartDeployRecord.ChartRelease)
		if err != nil {
			return v2models.ChartDeployRecord{}, err
		}
		chartReleaseID = chartRelease.ID
	}
	return v2models.ChartDeployRecord{
		Model: gorm.Model{
			ID:        chartDeployRecord.ID,
			CreatedAt: chartDeployRecord.CreatedAt,
			UpdatedAt: chartDeployRecord.UpdatedAt,
			DeletedAt: gorm.DeletedAt{},
		},
		ChartReleaseID:    chartReleaseID,
		ExactChartVersion: chartDeployRecord.ExactChartVersion,
		ExactAppVersion:   chartDeployRecord.ExactAppVersion,
		HelmfileRef:       chartDeployRecord.HelmfileRef,
	}, nil
}

func setChartDeployRecordDynamicDefaults(chartDeployRecord *ChartDeployRecord, stores *v2models.StoreSet, _ *auth.User) error {
	chartRelease, err := stores.ChartReleaseStore.Get(chartDeployRecord.ChartRelease)
	if err != nil {
		return err
	}
	if chartDeployRecord.ExactChartVersion == "" && chartRelease.CurrentChartVersionExact != nil {
		chartDeployRecord.ExactChartVersion = *chartRelease.CurrentChartVersionExact
	}
	if chartDeployRecord.ExactAppVersion == "" && chartRelease.CurrentAppVersionExact != nil {
		chartDeployRecord.ExactAppVersion = *chartRelease.CurrentAppVersionExact
	}
	if chartDeployRecord.HelmfileRef == "" && chartRelease.HelmfileRef != nil {
		chartDeployRecord.HelmfileRef = *chartRelease.HelmfileRef
	}
	return nil
}
