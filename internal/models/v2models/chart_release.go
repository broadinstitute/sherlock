package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/errors"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type ChartRelease struct {
	gorm.Model
	Chart         Chart
	ChartID       uint
	Cluster       *Cluster
	ClusterID     *uint
	Environment   *Environment
	EnvironmentID *uint
	Name          string `gorm:"not null; default:null; unique"`
	Namespace     string
	// Mutable
	CurrentAppVersionExact   *string
	CurrentChartVersionExact *string
	HelmfileRef              *string
	TargetAppVersionBranch   *string
	TargetAppVersionCommit   *string
	TargetAppVersionExact    *string
	TargetAppVersionUse      *string
	TargetChartVersionExact  *string
	TargetChartVersionUse    *string `gorm:"not null; default:null"`
	ThelmaMode               *string
}

func (c ChartRelease) TableName() string {
	return "v2_chart_releases"
}

func newChartReleaseStore(db *gorm.DB) Store[ChartRelease] {
	return Store[ChartRelease]{
		db:                       db,
		selectorToQueryModel:     chartReleaseSelectorToQuery,
		modelToSelectors:         chartReleaseToSelectors,
		modelRequiresSuitability: chartReleaseRequiresSuitability,
	}
}

func chartReleaseSelectorToQuery(db *gorm.DB, selector string) (ChartRelease, error) {
	var ret ChartRelease
	if isNumeric(selector) { // ID
		id, err := strconv.Atoi(selector)
		if err != nil {
			return ChartRelease{}, fmt.Errorf("(%s) string to int conversion error of '%s': %v", errors.BadRequest, selector, err)
		}
		ret.ID = uint(id)
		return ret, nil
	} else if strings.Count(selector, "/") == 1 { // environment + chart
		parts := strings.Split(selector, "/")

		// environment
		environmentQuery, err := environmentSelectorToQuery(db, parts[0])
		if err != nil {
			return ChartRelease{}, fmt.Errorf("invalid chart release selector %s, environment sub-selector error: %v", selector, err)
		}
		environment, err := getFromQuery(db, environmentQuery)
		if err != nil {
			return ChartRelease{}, fmt.Errorf("error handling environment sub-selector %s: %v", parts[0], err)
		}
		ret.EnvironmentID = &environment.ID

		// chart
		chartQuery, err := chartSelectorToQuery(db, parts[1])
		if err != nil {
			return ChartRelease{}, fmt.Errorf("invalid chart release selector %s, chart sub-selector error: %v", selector, err)
		}
		chart, err := getFromQuery(db, chartQuery)
		if err != nil {
			return ChartRelease{}, fmt.Errorf("error handling chart sub-selector %s: %v", parts[1], err)
		}
		ret.ChartID = chart.ID

		return ret, nil
	} else if strings.Count(selector, "/") == 2 { // cluster + namespace + chart
		parts := strings.Split(selector, "/")

		// cluster
		clusterQuery, err := clusterSelectorToQuery(db, parts[0])
		if err != nil {
			return ChartRelease{}, fmt.Errorf("invalid chart release selector %s, cluster sub-selector error: %v", selector, err)
		}
		cluster, err := getFromQuery(db, clusterQuery)
		if err != nil {
			return ChartRelease{}, fmt.Errorf("error handling cluster sub-selector %s: %v", parts[0], err)
		}
		ret.ClusterID = &cluster.ID

		// namespace
		namespace := parts[1]
		if !(isAlphaNumericWithHyphens(namespace) &&
			len(namespace) > 0 &&
			isStartingWithLetter(namespace) &&
			isEndingWithAlphaNumeric(namespace)) {
			return ChartRelease{}, fmt.Errorf("(%s) invalid chart release selector %s, namespace sub-selector %s was invalid", errors.BadRequest, selector, namespace)
		}
		ret.Namespace = namespace

		// chart
		chartQuery, err := chartSelectorToQuery(db, parts[2])
		if err != nil {
			return ChartRelease{}, fmt.Errorf("invalid chart release selector %s, chart sub-selector error: %v", selector, err)
		}
		chart, err := getFromQuery(db, chartQuery)
		if err != nil {
			return ChartRelease{}, fmt.Errorf("error handling chart sub-selector %s: %v", parts[1], err)
		}
		ret.ChartID = chart.ID

		return ret, nil
	} else if isAlphaNumericWithHyphens(selector) &&
		len(selector) > 0 &&
		isStartingWithLetter(selector) &&
		isEndingWithAlphaNumeric(selector) { // name
		ret.Name = selector
		return ret, nil
	}
	return ChartRelease{}, fmt.Errorf("(%s) invalid chart release selector '%s'", errors.BadRequest, selector)
}

// chartReleaseToSelectors is subtly more complex than some of the other modelToSelectors functions. A ChartRelease
// is special in that its selectors vary based on optionally provided associations.
// The contract for this function is that it generate as many selectors as possible based on the input, and *usually* it
// is just sufficient to call the modelToSelectors functions on any associations and compose the output. Here, though,
// it is possible for the Environment or Cluster to be nil *but the EnvironmentID or ClusterID to be present!* That
// would be a sign that the associations just weren't actually loaded in while assembling the ChartRelease (maybe we're
// validating something not in the database yet?). In that case, we should use the EnvironmentID or ClusterID directly
// as the numeric selectors they are.
//
// (This "ID present but the association wasn't loaded" case is actually just a general thing across most types here,
// but ChartRelease is the only type where those associations actually influence the selectors, so the modelToSelectors
// functions for other types don't need to care)
func chartReleaseToSelectors(chartRelease ChartRelease) []string {
	var selectors []string
	if (chartRelease.Environment != nil || chartRelease.EnvironmentID != nil) || ((chartRelease.Cluster != nil || chartRelease.ClusterID != nil) && chartRelease.Namespace != "") {
		chartSelectors := chartToSelectors(chartRelease.Chart)
		if chartRelease.Environment != nil {
			for _, environmentSelector := range environmentToSelectors(*chartRelease.Environment) {
				for _, chartSelector := range chartSelectors {
					selectors = append(selectors, fmt.Sprintf("%s/%s", environmentSelector, chartSelector))
				}
			}
		} else if chartRelease.EnvironmentID != nil {
			// Environment not present but ID is, we can't call environmentToSelectors but we know the ID is a selector anyway
			for _, chartSelector := range chartSelectors {
				selectors = append(selectors, fmt.Sprintf("%d/%s", *chartRelease.EnvironmentID, chartSelector))
			}
		}
		if chartRelease.Cluster != nil && chartRelease.Namespace != "" {
			for _, clusterSelector := range clusterToSelectors(*chartRelease.Cluster) {
				for _, chartSelector := range chartSelectors {
					selectors = append(selectors, fmt.Sprintf("%s/%s/%s", clusterSelector, chartRelease.Namespace, chartSelector))
				}
			}
		} else if chartRelease.ClusterID != nil && chartRelease.Namespace != "" {
			// Cluster not present but ID is, we can't call clusterToSelectors but we know the ID is a selector anyway
			for _, chartSelector := range chartSelectors {
				selectors = append(selectors, fmt.Sprintf("%d/%s/%s", *chartRelease.ClusterID, chartRelease.Namespace, chartSelector))
			}
		}
	}
	if chartRelease.Name != "" {
		selectors = append(selectors, chartRelease.Name)
	}
	if chartRelease.ID != 0 {
		selectors = append(selectors, fmt.Sprintf("%d", chartRelease.ID))
	}
	return selectors
}

func chartReleaseRequiresSuitability(chartRelease ChartRelease) bool {
	return chartRelease.Cluster == nil || clusterRequiresSuitability(*chartRelease.Cluster) || (chartRelease.Environment == nil || environmentRequiresSuitability(*chartRelease.Environment))
}
