package v2controllers

import (
	"context"
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/internal/metrics"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"time"
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
		postCreate:         postCreateChartDeployRecord,
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

func postCreateChartDeployRecord(chartDeployRecord ChartDeployRecord, stores *v2models.StoreSet, _ *auth.User) {
	if config.Config.String("metrics.accelerate.fromAPI") == "v2" && chartDeployRecord.ExactAppVersion != "" {
		ctx := context.Background()
		metrics.RecordDeployFrequency(ctx,
			chartDeployRecord.ChartReleaseInfo.Environment,
			chartDeployRecord.ChartReleaseInfo.Chart)

		chartDeployRecordQuery := v2models.ChartDeployRecord{ExactAppVersion: chartDeployRecord.ExactAppVersion, ChartReleaseID: chartDeployRecord.ChartReleaseInfo.ID}
		deploysOfThisAppVersion, err := stores.ChartDeployRecordStore.ListAllMatching(
			chartDeployRecordQuery, 2)
		if err != nil {
			log.Error().Msgf("error recording lead-time metric: couldn't determine if this is a re-deploy: %v", err)
		} else if len(deploysOfThisAppVersion) == 0 {
			log.Warn().Msgf("skipping recording lead-time metric: this ChartDeployRecord %v was not returned in query %v", chartDeployRecord, chartDeployRecordQuery)
		} else if len(deploysOfThisAppVersion) > 1 {
			log.Debug().Msgf("skipping recording lead-time metric: this ChartDeployRecord is a re-deploy, app version previously deployed at %s", deploysOfThisAppVersion[1].CreatedAt.Format(time.RFC1123))
		} else {
			chartRelease, err := stores.ChartReleaseStore.Get(chartDeployRecord.ChartRelease)
			if err != nil {
				log.Error().Msgf("error recording lead-time metric: couldn't get ChartRelease of ChartDeployRecord: %v", err)
			}
			appVersionQuery := v2models.AppVersion{AppVersion: chartDeployRecord.ExactAppVersion, ChartID: chartRelease.ChartID}
			matchingAppVersions, err := stores.AppVersionStore.ListAllMatching(
				appVersionQuery, 1)
			if err != nil {
				log.Error().Msgf("error recording lead-time metric: couldn't query matching app versions: %v", err)
			} else if len(matchingAppVersions) == 0 {
				log.Debug().Msgf("skipping recording lead-time metric: no matching app versions returned by query %v", appVersionQuery)
			} else {
				metrics.RecordLeadTime(ctx,
					chartDeployRecord.CreatedAt.Sub(matchingAppVersions[0].CreatedAt).Hours(),
					chartDeployRecord.ChartReleaseInfo.Environment,
					chartDeployRecord.ChartReleaseInfo.Chart)
			}
		}
	}
}
