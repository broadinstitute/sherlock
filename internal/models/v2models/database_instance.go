package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/errors"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type DatabaseInstance struct {
	gorm.Model
	ChartRelease   *ChartRelease
	ChartReleaseID uint

	Platform                  *string
	GoogleProject             *string
	GoogleLocation            *string
	AzureSubscription         *string
	AzureManagedResourceGroup *string
	InstanceName              *string

	DefaultDatabase *string
}

func (d DatabaseInstance) TableName() string {
	return "v2_database_instances"
}

var databaseInstanceStore *internalModelStore[DatabaseInstance]

func init() {
	databaseInstanceStore = &internalModelStore[DatabaseInstance]{
		selectorToQueryModel:     databaseInstanceSelectorToQuery,
		modelToSelectors:         databaseInstanceToSelectors,
		modelRequiresSuitability: databaseInstanceRequiresSuitability,
		validateModel:            validateDatabaseInstance,
	}
}

func databaseInstanceSelectorToQuery(db *gorm.DB, selector string) (DatabaseInstance, error) {
	if len(selector) == 0 {
		return DatabaseInstance{}, fmt.Errorf("(%s) database instance selector cannot be empty", errors.BadRequest)
	}
	var query DatabaseInstance
	if isNumeric(selector) { // ID
		id, err := strconv.Atoi(selector)
		if err != nil {
			return DatabaseInstance{}, fmt.Errorf("(%s) string to int conversion error of '%s': %v", errors.BadRequest, selector, err)
		}
		query.ID = uint(id)
		return query, nil
	} else if strings.HasPrefix(selector, "chart-release/") { // "chart-release/" + chart release
		chartReleaseSubSelector := strings.TrimPrefix(selector, "chart-release/")
		chartReleaseQuery, err := chartReleaseSelectorToQuery(db, chartReleaseSubSelector)
		if err != nil {
			return DatabaseInstance{}, fmt.Errorf("invalid database instance selector %s, chart release sub-selector error: %v", selector, err)
		}
		chartRelease, err := chartReleaseStore.get(db, chartReleaseQuery)
		if err != nil {
			return DatabaseInstance{}, fmt.Errorf("error handling database instance subselector %s: %v", chartReleaseSubSelector, err)
		}
		query.ChartReleaseID = chartRelease.ID
		return query, nil
	}
	return DatabaseInstance{}, fmt.Errorf("(%s) invalid database instance selector '%s'", errors.BadRequest, selector)
}

func databaseInstanceToSelectors(databaseInstance *DatabaseInstance) []string {
	var selectors []string
	if databaseInstance != nil {
		if chartReleaseSelectors := chartReleaseToSelectors(databaseInstance.ChartRelease); len(chartReleaseSelectors) > 0 {
			for _, chartReleaseSelector := range chartReleaseSelectors {
				selectors = append(selectors, fmt.Sprintf("chart-release/%s", chartReleaseSelector))
			}
		} else if databaseInstance.ChartReleaseID != 0 {
			selectors = append(selectors, fmt.Sprintf("chart-release/%d", databaseInstance.ChartReleaseID))
		}
		if databaseInstance.ID != 0 {
			selectors = append(selectors, fmt.Sprintf("%d", databaseInstance.ID))
		}
	}
	return selectors
}

func databaseInstanceRequiresSuitability(db *gorm.DB, databaseInstance *DatabaseInstance) bool {
	if chartRelease, err := chartReleaseStore.get(db, *databaseInstance.ChartRelease); err != nil {
		return true
	} else {
		return chartReleaseRequiresSuitability(db, &chartRelease)
	}
}

func validateDatabaseInstance(databaseInstance *DatabaseInstance) error {
	if databaseInstance == nil {
		return fmt.Errorf("the model passed was nil")
	}
	if databaseInstance.ChartReleaseID == 0 {
		return fmt.Errorf("a %T must have an associated chart release", databaseInstance)
	}
	if databaseInstance.Platform == nil {
		return fmt.Errorf("a %T must have a platform", databaseInstance)
	} else {
		switch *databaseInstance.Platform {
		case "google":
			if databaseInstance.GoogleProject == nil || *databaseInstance.GoogleProject == "" {
				return fmt.Errorf("a %T with a 'google' provider must have a google project", databaseInstance)
			}
			if databaseInstance.GoogleLocation == nil || *databaseInstance.GoogleProject == "" {
				return fmt.Errorf("a %T with a 'google' provider must have a location", databaseInstance)
			}
			if databaseInstance.InstanceName == nil || *databaseInstance.InstanceName == "" {
				return fmt.Errorf("a %T with an 'google' provider must have an instance name", databaseInstance)
			}
		case "azure":
			if databaseInstance.AzureSubscription == nil || *databaseInstance.AzureSubscription == "" {
				return fmt.Errorf("a %T with an 'azure' provider must have a subscription", databaseInstance)
			}
			if databaseInstance.AzureManagedResourceGroup == nil || *databaseInstance.AzureManagedResourceGroup == "" {
				return fmt.Errorf("a %T with an 'azure' provider must have an MRG", databaseInstance)
			}
			if databaseInstance.InstanceName == nil || *databaseInstance.InstanceName == "" {
				return fmt.Errorf("a %T with an 'azure' provider must have an instance name", databaseInstance)
			}
		case "kubernetes":
		default:
			return fmt.Errorf("a %T must have a platform of 'google', 'azure', or 'kubernetes'", databaseInstance)
		}
	}
	if databaseInstance.DefaultDatabase == nil || *databaseInstance.DefaultDatabase == "" {
		return fmt.Errorf("a %T must have a default database to connect to inside the instance", databaseInstance)
	}
	return nil
}
