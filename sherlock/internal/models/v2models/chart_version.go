package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/utils"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type ChartVersion struct {
	gorm.Model
	CiIdentifier         *CiIdentifier `gorm:"polymorphic:Resource; polymorphicValue:chart-version"`
	Chart                *Chart
	ChartID              uint   `gorm:"not null: default:null"`
	ChartVersion         string `gorm:"not null: default:null"`
	Description          string
	ParentChartVersion   *ChartVersion
	ParentChartVersionID *uint
}

func (c ChartVersion) TableName() string {
	return "v2_chart_versions"
}

func (c ChartVersion) getID() uint {
	return c.ID
}

func (c ChartVersion) getParentID() *uint {
	return c.ParentChartVersionID
}

func (c ChartVersion) GetCiIdentifier() *CiIdentifier {
	if c.CiIdentifier != nil {
		return c.CiIdentifier
	} else {
		return &CiIdentifier{ResourceType: "chart-version", ResourceID: c.ID}
	}
}

var chartVersionStore *internalTreeModelStore[ChartVersion]

func init() {
	chartVersionStore = &internalTreeModelStore[ChartVersion]{
		internalModelStore: &internalModelStore[ChartVersion]{
			selectorToQueryModel:    chartVersionSelectorToQuery,
			modelToSelectors:        chartVersionToSelectors,
			validateModel:           validateChartVersion,
			handleIncomingDuplicate: rejectDuplicateChartVersion,
		},
	}
}

func chartVersionSelectorToQuery(db *gorm.DB, selector string) (ChartVersion, error) {
	if len(selector) == 0 {
		return ChartVersion{}, fmt.Errorf("(%s) chart version selector cannot be empty", errors.BadRequest)
	}
	var query ChartVersion
	if utils.IsNumeric(selector) { // ID
		id, err := strconv.Atoi(selector)
		if err != nil {
			return ChartVersion{}, fmt.Errorf("(%s) string to int conversion error of '%s': %v", errors.BadRequest, selector, err)
		}
		query.ID = uint(id)
		return query, nil
	} else if strings.Count(selector, "/") == 1 { // chart + version
		parts := strings.Split(selector, "/")

		// chart
		chartID, err := chartStore.resolveSelector(db, parts[0])
		if err != nil {
			return ChartVersion{}, fmt.Errorf("error handling chart sub-selector %s: %v", parts[0], err)
		}
		query.ChartID = chartID

		// version
		version := parts[1]
		if len(version) == 0 {
			return ChartVersion{}, fmt.Errorf("(%s) invalid chart version selector %s, version sub-selector was empty", errors.BadRequest, selector)
		}
		query.ChartVersion = version

		return query, nil
	}
	return ChartVersion{}, fmt.Errorf("(%s) invalid chart version selector '%s'", errors.BadRequest, selector)
}

func chartVersionToSelectors(chartVersion *ChartVersion) []string {
	var selectors []string
	if chartVersion != nil {
		if chartVersion.ID != 0 {
			selectors = append(selectors, fmt.Sprintf("%d", chartVersion.ID))
		}
		if chartVersion.ChartVersion != "" {
			chartSelectors := chartToSelectors(chartVersion.Chart)
			if len(chartSelectors) == 0 && chartVersion.ChartID != 0 {
				// Chart not filled so chartToSelectors gives nothing, but we have the chart ID and it is a selector anyway
				chartSelectors = []string{fmt.Sprintf("%d", chartVersion.ChartID)}
			}
			for _, chartSelector := range chartSelectors {
				selectors = append(selectors, fmt.Sprintf("%s/%s", chartSelector, chartVersion.ChartVersion))
			}
		}
	}
	return selectors
}

func validateChartVersion(chartVersion *ChartVersion) error {
	if chartVersion == nil {
		return fmt.Errorf("the model passed was nil")
	}
	if chartVersion.ChartID == 0 {
		return fmt.Errorf("an %T must have an associated chart", chartVersion)
	}
	if chartVersion.ChartVersion == "" {
		return fmt.Errorf("an %T must have a non-empty chart version", chartVersion)
	}
	return nil
}

func rejectDuplicateChartVersion(existing *ChartVersion, new *ChartVersion) error {
	if existing.ChartVersion != new.ChartVersion {
		return fmt.Errorf("new %T has chart version %s, which is mismatched with the existing value of %s", new, new.ChartVersion, existing.ChartVersion)
	}
	if existing.ChartID != new.ChartID {
		return fmt.Errorf("new %T has chart ID %d, which is mismatched with the existing value of %d", new, new.ChartID, existing.ChartID)
	}
	if (existing.ParentChartVersionID != nil) && (new.ParentChartVersionID == nil) {
		return fmt.Errorf("new %T has no parent ID, which is mismatched with the existing having one", new)
	} else if (existing.ParentChartVersionID == nil) && (new.ParentChartVersionID != nil) {
		return fmt.Errorf("new %T has a parent ID, which is mismatched with the existing not having one", new)
	} else if existing.ParentChartVersionID != nil && new.ParentChartVersionID != nil && *existing.ParentChartVersionID != *new.ParentChartVersionID {
		return fmt.Errorf("new %T has parent ID %d, which is mismatched with the existing value of %d", new, *new.ParentChartVersionID, *existing.ParentChartVersionID)
	}
	return nil
}
