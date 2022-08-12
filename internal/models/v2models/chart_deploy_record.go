package v2models

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/internal/metrics"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"strconv"
)

type ChartDeployRecord struct {
	gorm.Model
	ChartRelease      ChartRelease
	ChartReleaseID    uint   `gorm:"not null; default:null"`
	ExactChartVersion string `gorm:"not null; default:null"`
	ExactAppVersion   string
	HelmfileRef       string `gorm:"not null; default:null"`
}

func (c ChartDeployRecord) TableName() string {
	return "v2_chart_deploy_records"
}

func newChartDeployRecordStore(db *gorm.DB) *Store[ChartDeployRecord] {
	return &Store[ChartDeployRecord]{
		db:                       db,
		selectorToQueryModel:     chartDeployRecordSelectorToQuery,
		modelToSelectors:         chartDeployRecordToSelectors,
		modelRequiresSuitability: chartDeployRecordRequiresSuitability,
		validateModel:            validateChartDeployRecord,
		postCreate:               postCreateChartDeployRecord,
	}
}

func chartDeployRecordSelectorToQuery(_ *gorm.DB, selector string) (ChartDeployRecord, error) {
	if len(selector) == 0 {
		return ChartDeployRecord{}, fmt.Errorf("(%s) chart deploy record selector cannot be empty", errors.BadRequest)
	}
	var query ChartDeployRecord
	if isNumeric(selector) { // ID
		id, err := strconv.Atoi(selector)
		if err != nil {
			return ChartDeployRecord{}, fmt.Errorf("(%s) string to int conversion error of '%s': %v", errors.BadRequest, selector, err)
		}
		query.ID = uint(id)
		return query, nil
	}
	return ChartDeployRecord{}, fmt.Errorf("(%s) invalid chart deploy record selector '%s'", errors.BadRequest, selector)
}

func chartDeployRecordToSelectors(chartDeployRecord ChartDeployRecord) []string {
	var selectors []string
	if chartDeployRecord.ID != 0 {
		selectors = append(selectors, fmt.Sprintf("%d", chartDeployRecord.ID))
	}
	return selectors
}

func chartDeployRecordRequiresSuitability(db *gorm.DB, chartDeployRecord ChartDeployRecord) bool {
	chartRelease, err := getFromQuery(db, chartDeployRecord.ChartRelease)
	if err != nil {
		return true
	}
	return chartReleaseRequiresSuitability(db, chartRelease)
}

func validateChartDeployRecord(chartDeployRecord ChartDeployRecord) error {
	if chartDeployRecord.ChartReleaseID == 0 {
		return fmt.Errorf("a %T must have an associated chart release", chartDeployRecord)
	}
	if chartDeployRecord.ExactChartVersion == "" {
		return fmt.Errorf("a %T must have a non-empty chart version", chartDeployRecord)
	}
	if chartDeployRecord.HelmfileRef == "" {
		return fmt.Errorf("a %T must have a non-empty terra-helmfile ref", chartDeployRecord)
	}
	return nil
}

func postCreateChartDeployRecord(db *gorm.DB, chartDeployRecord ChartDeployRecord, _ *auth.User) error {
	if config.Config.String("metrics.accelerate.fromAPI") == "v2" && chartDeployRecord.ExactAppVersion != "" {
		ctx := context.Background()
		storeSet := NewStoreSet(db)
		// reload to pull in the associations of the chart release
		chartRelease, err := getFromQuery(db, chartDeployRecord.ChartRelease)
		if err != nil {
			return fmt.Errorf("error recording metrics: couldn't fully load %T %s: %v", chartDeployRecord.ChartRelease, chartDeployRecord.ChartRelease.Name, err)
		} else if chartRelease.Environment == nil {
			log.Debug().Msgf("skipping recording accelerate metrics for %s's new %T for %s %s: it wasn't associated to an environment",
				chartRelease.Name, chartDeployRecord, chartRelease.Chart.Name, chartDeployRecord.ExactAppVersion)
			return nil
		} else if chartRelease.Environment.Lifecycle != "static" {
			log.Debug().Msgf("skipping recording accelerate metrics for %s's new %T for %s %s: the environment was '%s' instead of 'static'",
				chartRelease.Name, chartDeployRecord, chartRelease.Chart.Name, chartDeployRecord.ExactAppVersion, chartRelease.Environment.Lifecycle)
			return nil
		}

		var leadTime float64
		deploysOfThisAppVersion, err := storeSet.ChartDeployRecordStore.ListAllMatching(
			ChartDeployRecord{ExactAppVersion: chartDeployRecord.ExactAppVersion, ChartReleaseID: chartRelease.ID}, 2)
		if err != nil {
			return fmt.Errorf("error recording metrics: couldn't look for other deployments of the same version %s: %v", chartDeployRecord.ExactAppVersion, err)
		} else if len(deploysOfThisAppVersion) > 1 {
			log.Debug().Msgf("skipping recording lead-time metric for %s's new %T for %s %s: it was previous deployed here at %s",
				chartRelease.Name, chartDeployRecord, chartRelease.Chart.Name, chartDeployRecord.ExactAppVersion, deploysOfThisAppVersion[1].CreatedAt.String())
		} else {
			matchingAppVersions, err := storeSet.AppVersionStore.ListAllMatching(
				AppVersion{AppVersion: chartDeployRecord.ExactAppVersion, ChartID: chartRelease.ChartID}, 1)
			if err != nil {
				return fmt.Errorf("error recording metrics: couldn't look for app versions matching deployed %s: %v", chartDeployRecord.ExactAppVersion, err)
			} else if len(matchingAppVersions) == 0 {
				log.Debug().Msgf("skipping recording lead-time metric for %s's new %T for %s %s: no app versions found, so we have no build time",
					chartRelease.Name, chartDeployRecord, chartRelease.Chart.Name, chartDeployRecord.ExactAppVersion)
			} else {
				leadTime = chartDeployRecord.CreatedAt.Sub(matchingAppVersions[0].CreatedAt).Hours()
			}
		}

		metrics.RecordDeployFrequency(ctx,
			chartRelease.Environment.Name,
			chartRelease.Chart.Name)
		if leadTime != 0 {
			metrics.RecordLeadTime(ctx,
				leadTime,
				chartRelease.Environment.Name,
				chartRelease.Chart.Name)
		}
	}
	return nil
}
