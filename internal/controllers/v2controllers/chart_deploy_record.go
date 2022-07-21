package v2controllers

import (
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"gorm.io/gorm"
)

type ChartDeployRecord struct {
	ReadableBaseType
	ChartReleaseInfo ChartRelease
	CreatableChartDeployRecord
}

type CreatableChartDeployRecord struct {
	ChartRelease      string
	ExactChartVersion string
	ExactAppVersion   string
	HelmfileRef       string
}

func (c CreatableChartDeployRecord) toReadable() ChartDeployRecord {
	return ChartDeployRecord{CreatableChartDeployRecord: c}
}

type ChartDeployRecordController = ImmutableModelController[v2models.ChartDeployRecord, ChartDeployRecord, CreatableChartDeployRecord]

func NewChartDeployRecordController(stores v2models.StoreSet) *ChartDeployRecordController {
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

func chartDeployRecordToModelChartDeployRecord(chartDeployRecord ChartDeployRecord, stores v2models.StoreSet) (v2models.ChartDeployRecord, error) {
	chartRelease, err := stores.ChartReleaseStore.Get(chartDeployRecord.ChartRelease)
	if err != nil {
		return v2models.ChartDeployRecord{}, err
	}
	return v2models.ChartDeployRecord{
		Model: gorm.Model{
			ID:        chartDeployRecord.ID,
			CreatedAt: chartDeployRecord.CreatedAt,
			UpdatedAt: chartDeployRecord.UpdatedAt,
			DeletedAt: gorm.DeletedAt{},
		},
		ChartReleaseID:    chartRelease.ID,
		ExactChartVersion: chartDeployRecord.ExactChartVersion,
		ExactAppVersion:   chartDeployRecord.ExactAppVersion,
		HelmfileRef:       chartDeployRecord.HelmfileRef,
	}, nil
}
