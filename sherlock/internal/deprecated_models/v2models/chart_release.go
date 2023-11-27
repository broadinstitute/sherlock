package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/model_actions"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type ChartRelease struct {
	gorm.Model
	CiIdentifier    *CiIdentifier `gorm:"polymorphic:Resource; polymorphicValue:chart-release"`
	Chart           *Chart
	ChartID         uint
	Cluster         *Cluster
	ClusterID       *uint
	DestinationType string
	Environment     *Environment
	EnvironmentID   *uint
	Name            string `gorm:"not null; default:null"`
	Namespace       string
	ChartReleaseVersion
	Subdomain               *string
	Protocol                *string
	Port                    *uint
	PagerdutyIntegration    *PagerdutyIntegration
	PagerdutyIntegrationID  *uint
	IncludeInBulkChangesets *bool
}

func (c ChartRelease) getID() uint {
	return c.ID
}

func (c ChartRelease) GetCiIdentifier() *CiIdentifier {
	if c.CiIdentifier != nil {
		return c.CiIdentifier
	} else {
		return &CiIdentifier{ResourceType: "chart-release", ResourceID: c.ID}
	}
}

var InternalChartReleaseStore *internalModelStore[ChartRelease]

func init() {
	InternalChartReleaseStore = &internalModelStore[ChartRelease]{
		selectorToQueryModel:  chartReleaseSelectorToQuery,
		modelToSelectors:      chartReleaseToSelectors,
		errorIfForbidden:      chartReleaseErrorIfForbidden,
		validateModel:         validateChartRelease,
		preCreate:             preCreateChartRelease,
		preDeletePostValidate: preDeletePostValidateChartRelease,
	}
}

func chartReleaseSelectorToQuery(db *gorm.DB, selector string) (ChartRelease, error) {
	if len(selector) == 0 {
		return ChartRelease{}, fmt.Errorf("(%s) chart release selector cannot be empty", errors.BadRequest)
	}
	var query ChartRelease
	if utils.IsNumeric(selector) { // ID
		id, err := strconv.Atoi(selector)
		if err != nil {
			return ChartRelease{}, fmt.Errorf("(%s) string to int conversion error of '%s': %w", errors.BadRequest, selector, err)
		}
		query.ID = uint(id)
		return query, nil
	} else if strings.Count(selector, "/") == 1 { // environment + chart
		parts := strings.Split(selector, "/")

		// environment
		environmentID, err := InternalEnvironmentStore.ResolveSelector(db, parts[0])
		if err != nil {
			return ChartRelease{}, fmt.Errorf("error handling environment sub-selector %s: %w", parts[0], err)
		}
		query.EnvironmentID = &environmentID

		// chart
		chartID, err := InternalChartStore.ResolveSelector(db, parts[1])
		if err != nil {
			return ChartRelease{}, fmt.Errorf("error handling chart sub-selector %s: %w", parts[1], err)
		}
		query.ChartID = chartID

		return query, nil
	} else if strings.Count(selector, "/") == 2 { // cluster + namespace + chart
		parts := strings.Split(selector, "/")

		// cluster
		clusterID, err := InternalClusterStore.ResolveSelector(db, parts[0])
		if err != nil {
			return ChartRelease{}, fmt.Errorf("error handling cluster sub-selector %s: %w", parts[0], err)
		}
		query.ClusterID = &clusterID

		// namespace
		namespace := parts[1]
		if !(utils.IsAlphaNumericWithHyphens(namespace) &&
			len(namespace) > 0 &&
			utils.IsStartingWithLetter(namespace) &&
			utils.IsEndingWithAlphaNumeric(namespace)) {
			return ChartRelease{}, fmt.Errorf("(%s) invalid chart release selector %s, namespace sub-selector %s was invalid", errors.BadRequest, selector, namespace)
		}
		query.Namespace = namespace

		// chart
		chartID, err := InternalChartStore.ResolveSelector(db, parts[2])
		if err != nil {
			return ChartRelease{}, fmt.Errorf("error handling chart sub-selector %s: %w", parts[1], err)
		}
		query.ChartID = chartID

		return query, nil
	} else if utils.IsAlphaNumericWithHyphens(selector) &&
		utils.IsStartingWithLetter(selector) &&
		utils.IsEndingWithAlphaNumeric(selector) { // name
		query.Name = selector
		return query, nil
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
func chartReleaseToSelectors(chartRelease *ChartRelease) []string {
	var selectors []string
	if chartRelease != nil {
		if chartRelease.Name != "" {
			selectors = append(selectors, chartRelease.Name)
		}
		if (chartRelease.Environment != nil || chartRelease.EnvironmentID != nil) || ((chartRelease.Cluster != nil || chartRelease.ClusterID != nil) && chartRelease.Namespace != "") {
			chartSelectors := chartToSelectors(chartRelease.Chart)
			if len(chartSelectors) == 0 && chartRelease.ChartID != 0 {
				// Chart not filled so chartToSelectors gives nothing, but we have the chart ID and it is a selector anyway
				chartSelectors = []string{fmt.Sprintf("%d", chartRelease.ChartID)}
			}
			if chartRelease.Environment != nil {
				for _, environmentSelector := range environmentToSelectors(chartRelease.Environment) {
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
				for _, clusterSelector := range clusterToSelectors(chartRelease.Cluster) {
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
		if chartRelease.ID != 0 {
			selectors = append(selectors, fmt.Sprintf("%d", chartRelease.ID))
		}
	}
	return selectors
}

func chartReleaseErrorIfForbidden(db *gorm.DB, chartRelease *ChartRelease, action model_actions.ActionType, user *models.User) error {
	if chartRelease.Cluster != nil {
		if cluster, err := InternalClusterStore.Get(db, *chartRelease.Cluster); err != nil {
			return err
		} else if err = clusterErrorIfForbidden(db, &cluster, action, user); err != nil {
			return err
		}
	}
	if chartRelease.Environment != nil {
		if environment, err := InternalEnvironmentStore.Get(db, *chartRelease.Environment); err != nil {
			return err
		} else if err = environmentErrorIfForbidden(db, &environment, action, user); err != nil {
			return err
		}
	}
	return nil
}

func validateChartRelease(chartRelease *ChartRelease) error {
	if chartRelease == nil {
		return fmt.Errorf("the model passed was nil")
	}
	if chartRelease.Name == "" {
		return fmt.Errorf("a %T must have a non-empty Name", chartRelease)
	}
	if chartRelease.ChartID == 0 {
		return fmt.Errorf("a %T must have an associated chart", chartRelease)
	}
	if chartRelease.EnvironmentID != nil {
		if chartRelease.DestinationType != "environment" {
			return fmt.Errorf("(%s) calculated field for %T destination should be 'environment' if an environment is present", errors.InternalServerError, chartRelease)
		}
	} else if chartRelease.ClusterID != nil {
		if chartRelease.DestinationType != "cluster" {
			return fmt.Errorf("(%s) calculated field for %T destination should be 'cluster' if a cluster and no environment is present", errors.InternalServerError, chartRelease)
		}
	} else {
		return fmt.Errorf("a %T must have either an associated environment or an associated cluster", chartRelease)
	}

	if chartRelease.ClusterID != nil && chartRelease.Namespace == "" {
		return fmt.Errorf("a %T that has a cluster must have a namespace", chartRelease)
	} else if chartRelease.ClusterID == nil && chartRelease.Namespace != "" {
		return fmt.Errorf("a %T that doesn't have a cluster must not have a namespace", chartRelease)
	}

	return chartRelease.ChartReleaseVersion.validate()
}

func preCreateChartRelease(db *gorm.DB, toCreate *ChartRelease, _ *models.User) error {
	if toCreate != nil {
		if toCreate.EnvironmentID != nil {
			toCreate.DestinationType = "environment"
		} else if toCreate.ClusterID != nil {
			toCreate.DestinationType = "cluster"
		}
		return toCreate.ChartReleaseVersion.resolve(db, Chart{Model: gorm.Model{ID: toCreate.ChartID}})
	}
	return nil
}

func preDeletePostValidateChartRelease(db *gorm.DB, chartRelease *ChartRelease, user *models.User) error {
	_, err := InternalDatabaseInstanceStore.DeleteIfExists(db, DatabaseInstance{ChartReleaseID: chartRelease.ID}, user)
	return err
}
