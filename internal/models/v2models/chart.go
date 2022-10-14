package v2models

import (
	"fmt"
	"strconv"

	"github.com/broadinstitute/sherlock/internal/errors"
	"gorm.io/gorm"
)

type Chart struct {
	gorm.Model
	Name string `gorm:"not null; default:null; unique"`
	// Mutable
	ChartRepo             *string `gorm:"not null; default:null"`
	AppImageGitRepo       *string
	AppImageGitMainBranch *string
	ChartExposesEndpoint  *bool
	DefaultSubdomain      *string
	DefaultProtocol       *string
	DefaultPort           *uint
	LegacyConfigsEnabled  *bool
}

func (c Chart) TableName() string {
	return "v2_charts"
}

var chartStore *internalModelStore[Chart]

func init() {
	chartStore = &internalModelStore[Chart]{
		selectorToQueryModel: chartSelectorToQuery,
		modelToSelectors:     chartToSelectors,
		validateModel:        validateChart,
	}
}

func chartSelectorToQuery(_ *gorm.DB, selector string) (Chart, error) {
	if len(selector) == 0 {
		return Chart{}, fmt.Errorf("(%s) chart selector cannot be empty", errors.BadRequest)
	}
	var query Chart
	if isNumeric(selector) { // ID
		id, err := strconv.Atoi(selector)
		if err != nil {
			return Chart{}, fmt.Errorf("(%s) string to int conversion error of '%s': %v", errors.BadRequest, selector, err)
		}
		query.ID = uint(id)
		return query, nil
	} else if isAlphaNumericWithHyphens(selector) &&
		isStartingWithLetter(selector) &&
		isEndingWithAlphaNumeric(selector) { // Name
		query.Name = selector
		return query, nil
	}
	return Chart{}, fmt.Errorf("(%s) invalid chart selector '%s'", errors.BadRequest, selector)
}

func chartToSelectors(chart *Chart) []string {
	var selectors []string
	if chart != nil {
		if chart.Name != "" {
			selectors = append(selectors, chart.Name)
		}
		if chart.ID != 0 {
			selectors = append(selectors, fmt.Sprintf("%d", chart.ID))
		}
	}
	return selectors
}

func validateChart(chart *Chart) error {
	if chart == nil {
		return fmt.Errorf("the model passed was nil")
	}
	if chart.Name == "" {
		return fmt.Errorf("a %T must have a non-empty name", chart)
	}
	if chart.ChartRepo == nil || *chart.ChartRepo == "" {
		return fmt.Errorf("a %T must have a non-empty chart repo", chart)
	}
	return nil
}
