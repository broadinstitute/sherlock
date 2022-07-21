package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/errors"
	"gorm.io/gorm"
	"strconv"
)

type Chart struct {
	gorm.Model
	Name string `gorm:"not null; default:null; unique"`
	// Mutable
	ChartRepo             *string `gorm:"not null; default:null"`
	AppImageGitRepo       *string
	AppImageGitMainBranch *string
}

func (c Chart) TableName() string {
	return "v2_charts"
}

func newChartStore(db *gorm.DB) Store[Chart] {
	return Store[Chart]{
		db:                   db,
		selectorToQueryModel: chartSelectorToQuery,
		modelToSelectors:     chartToSelectors,
	}
}

func chartSelectorToQuery(_ *gorm.DB, selector string) (Chart, error) {
	var query Chart
	if isNumeric(selector) { // ID
		id, err := strconv.Atoi(selector)
		if err != nil {
			return Chart{}, fmt.Errorf("(%s) string to int conversion error of '%s': %v", errors.BadRequest, selector, err)
		}
		query.ID = uint(id)
		return query, nil
	} else if isAlphaNumericWithHyphens(selector) &&
		len(selector) > 0 && len(selector) <= 24 &&
		isStartingWithLetter(selector) &&
		isEndingWithAlphaNumeric(selector) { // Name
		query.Name = selector
		return query, nil
	}
	return Chart{}, fmt.Errorf("(%s) invalid chart selector '%s'", errors.BadRequest, selector)
}

func chartToSelectors(chart Chart) []string {
	var selectors []string
	if chart.Name != "" {
		selectors = append(selectors, chart.Name)
	}
	if chart.ID != 0 {
		selectors = append(selectors, fmt.Sprintf("%d", chart.ID))
	}
	return selectors
}
