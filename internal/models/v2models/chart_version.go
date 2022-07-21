package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/errors"
	"gorm.io/gorm"
	"strconv"
)

type ChartVersion struct {
	gorm.Model
	Chart        Chart
	ChartID      uint   `gorm:"not null: default:null"`
	ChartVersion string `gorm:"not null: default:null"`
}

func (c ChartVersion) TableName() string {
	return "v2_chart_versions"
}

func newChartVersionStore(db *gorm.DB) Store[ChartVersion] {
	return Store[ChartVersion]{
		db:                   db,
		selectorToQueryModel: chartVersionSelectorToQuery,
		modelToSelectors:     chartVersionToSelectors,
	}
}

func chartVersionSelectorToQuery(_ *gorm.DB, selector string) (ChartVersion, error) {
	var query ChartVersion
	if isNumeric(selector) { // ID
		id, err := strconv.Atoi(selector)
		if err != nil {
			return ChartVersion{}, fmt.Errorf("(%s) string to int conversion error of '%s': %v", errors.BadRequest, selector, err)
		}
		query.ID = uint(id)
		return query, nil
	}
	return ChartVersion{}, fmt.Errorf("(%s) invalid chart version selector '%s'", errors.BadRequest, selector)
}

func chartVersionToSelectors(chartVersion ChartVersion) []string {
	var selectors []string
	if chartVersion.ID != 0 {
		selectors = append(selectors, fmt.Sprintf("%d", chartVersion.ID))
	}
	return selectors
}
