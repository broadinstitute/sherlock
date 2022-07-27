package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/errors"
	"gorm.io/gorm"
	"strconv"
)

type ChartDeployRecord struct {
	gorm.Model
	ChartRelease      ChartRelease
	ChartReleaseID    uint   `gorm:"not null; default:null"`
	ExactChartVersion string `gorm:"not null; default:null"`
	ExactAppVersion   string `gorm:"not null; default:null"`
	HelmfileRef       string
}

func (c ChartDeployRecord) TableName() string {
	return "v2_chart_deploy_records"
}

func newChartDeployRecordStore(db *gorm.DB) Store[ChartDeployRecord] {
	return Store[ChartDeployRecord]{
		db:                       db,
		selectorToQueryModel:     chartDeployRecordSelectorToQuery,
		modelToSelectors:         chartDeployRecordToSelectors,
		modelRequiresSuitability: chartDeployRecordRequiresSuitability,
		validateModel:            validateChartDeployRecord,
	}
}

func chartDeployRecordSelectorToQuery(_ *gorm.DB, selector string) (ChartDeployRecord, error) {
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

func chartDeployRecordRequiresSuitability(chartDeployRecord ChartDeployRecord) bool {
	return chartReleaseRequiresSuitability(chartDeployRecord.ChartRelease)
}

func validateChartDeployRecord(chartDeployRecord ChartDeployRecord) error {
	if chartDeployRecord.ChartReleaseID == 0 {
		return fmt.Errorf("a %T must have an associated chart release", chartDeployRecord)
	}
	if chartDeployRecord.ExactChartVersion == "" {
		return fmt.Errorf("a %T must have a non-empty chart version", chartDeployRecord)
	}
	if chartDeployRecord.ExactAppVersion == "" {
		return fmt.Errorf("a %T must have a non-empty app version", chartDeployRecord)
	}
	return nil
}
