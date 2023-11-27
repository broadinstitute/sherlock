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

type DatabaseInstance struct {
	gorm.Model
	ChartRelease   *ChartRelease
	ChartReleaseID uint

	Platform      *string
	GoogleProject *string
	InstanceName  *string

	DefaultDatabase *string
}

func (d DatabaseInstance) getID() uint {
	return d.ID
}

var InternalDatabaseInstanceStore *internalModelStore[DatabaseInstance]

func init() {
	InternalDatabaseInstanceStore = &internalModelStore[DatabaseInstance]{
		selectorToQueryModel: databaseInstanceSelectorToQuery,
		modelToSelectors:     databaseInstanceToSelectors,
		errorIfForbidden:     databaseInstanceErrorIfForbidden,
		validateModel:        validateDatabaseInstance,
	}
}

func databaseInstanceSelectorToQuery(db *gorm.DB, selector string) (DatabaseInstance, error) {
	if len(selector) == 0 {
		return DatabaseInstance{}, fmt.Errorf("(%s) database instance selector cannot be empty", errors.BadRequest)
	}
	var query DatabaseInstance
	if utils.IsNumeric(selector) { // ID
		id, err := strconv.Atoi(selector)
		if err != nil {
			return DatabaseInstance{}, fmt.Errorf("(%s) string to int conversion error of '%s': %w", errors.BadRequest, selector, err)
		}
		query.ID = uint(id)
		return query, nil
	} else if strings.HasPrefix(selector, "chart-release/") { // "chart-release/" + chart release
		chartReleaseSubSelector := strings.TrimPrefix(selector, "chart-release/")
		chartReleaseID, err := InternalChartReleaseStore.ResolveSelector(db, chartReleaseSubSelector)
		if err != nil {
			return DatabaseInstance{}, fmt.Errorf("error handling database instance subselector %s: %w", chartReleaseSubSelector, err)
		}
		query.ChartReleaseID = chartReleaseID
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

func databaseInstanceErrorIfForbidden(db *gorm.DB, databaseInstance *DatabaseInstance, action model_actions.ActionType, user *models.User) error {
	if chartRelease, err := InternalChartReleaseStore.Get(db, *databaseInstance.ChartRelease); err != nil {
		return err
	} else if err = chartReleaseErrorIfForbidden(db, &chartRelease, action, user); err != nil {
		return err
	}
	return nil
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
			fallthrough
		case "azure":
			if databaseInstance.InstanceName == nil || *databaseInstance.InstanceName == "" {
				return fmt.Errorf("a %T with a 'google' or 'azure' provider must have an instance name", databaseInstance)
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
