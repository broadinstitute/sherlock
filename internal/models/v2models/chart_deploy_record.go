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
