package v2controllers

import (
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
	ExactChartVersion string `json:"exactChartVersion" form:"exactChartVersion"` // Required when creating
	ExactAppVersion   string `json:"exactAppVersion" form:"exactAppVersion"`     // Required when creating
	HelmfileRef       string `json:"helmfileRef" form:"helmfileRef"`             // Required when creating
	EditableChartDeployRecord
}

type EditableChartDeployRecord struct{}

func (c CreatableChartDeployRecord) toReadable() ChartDeployRecord {
	return ChartDeployRecord{CreatableChartDeployRecord: c}
}

func (e EditableChartDeployRecord) toCreatable() CreatableChartDeployRecord {
	return CreatableChartDeployRecord{EditableChartDeployRecord: e}
}

type ChartDeployRecordController = ModelController[v2models.ChartDeployRecord, ChartDeployRecord, CreatableChartDeployRecord, EditableChartDeployRecord]

func newChartDeployRecordController(stores *v2models.StoreSet) *ChartDeployRecordController {
	return &ChartDeployRecordController{
		primaryStore:    stores.ChartDeployRecordStore,
		allStores:       stores,
		modelToReadable: modelChartDeployRecordToChartDeployRecord,
		readableToModel: chartDeployRecordToModelChartDeployRecord,
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
		chartReleaseID = chartRelease.ChartID
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
