package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/errors"
	"gorm.io/gorm"
	"strconv"
)

type Environment struct {
	gorm.Model
	Base                      string
	Lifecycle                 string `gorm:"not null; default:null"`
	Name                      string `gorm:"not null; default:null; unique"`
	TemplateEnvironment       *Environment
	TemplateEnvironmentID     *uint
	ValuesName                string
	ChartReleasesFromTemplate *bool
	// Mutable
	DefaultCluster      *Cluster
	DefaultClusterID    *uint
	DefaultNamespace    *string
	Owner               *string `gorm:"not null;default:null"`
	RequiresSuitability *bool
}

func (e Environment) TableName() string {
	return "v2_environments"
}

func newEnvironmentStore(db *gorm.DB) *Store[Environment] {
	return &Store[Environment]{
		db:                       db,
		selectorToQueryModel:     environmentSelectorToQuery,
		modelToSelectors:         environmentToSelectors,
		modelRequiresSuitability: environmentRequiresSuitability,
		validateModel:            validateEnvironment,
		postCreate:               postCreateEnvironment,
	}
}

func environmentSelectorToQuery(_ *gorm.DB, selector string) (Environment, error) {
	if len(selector) == 0 {
		return Environment{}, fmt.Errorf("(%s) environment selector cannot be empty", errors.BadRequest)
	}
	var query Environment
	if isNumeric(selector) { // ID
		id, err := strconv.Atoi(selector)
		if err != nil {
			return Environment{}, fmt.Errorf("(%s) string to int conversion error of '%s': %v", errors.BadRequest, selector, err)
		}
		query.ID = uint(id)
		return query, nil
	} else if isAlphaNumericWithHyphens(selector) &&
		isStartingWithLetter(selector) &&
		isEndingWithAlphaNumeric(selector) { // Name
		if len(selector) > 32 {
			return Environment{}, fmt.Errorf("(%s) %T name is too long, was %d characters and the maximum is 32", errors.BadRequest, Environment{}, len(selector))
		}
		query.Name = selector
		return query, nil
	}
	return Environment{}, fmt.Errorf("(%s) invalid environment selector '%s'", errors.BadRequest, selector)
}

func environmentToSelectors(environment Environment) []string {
	var selectors []string
	if environment.Name != "" {
		selectors = append(selectors, environment.Name)
	}
	if environment.ID != 0 {
		selectors = append(selectors, fmt.Sprintf("%d", environment.ID))
	}
	return selectors
}

func environmentRequiresSuitability(_ *gorm.DB, environment Environment) bool {
	// RequiresSuitability is a required field and shouldn't ever actually be stored as nil, but if it is we fail-safe
	return environment.RequiresSuitability == nil || *environment.RequiresSuitability
}

func validateEnvironment(environment Environment) error {
	if environment.Name == "" {
		return fmt.Errorf("a %T must have a non-empty name", environment)
	}
	switch environment.Lifecycle {
	case "template":
		if environment.TemplateEnvironmentID != nil {
			return fmt.Errorf("a template %T cannot itself have a template", environment)
		}
	case "dynamic":
		if environment.TemplateEnvironmentID == nil {
			return fmt.Errorf("a dynamic %T must have a template", environment)
		}
		fallthrough
	case "static":
		if environment.Base == "" {
			return fmt.Errorf("a non-template %T must have a base", environment)
		}
		if environment.DefaultClusterID == nil {
			return fmt.Errorf("a non-template %T must have a default cluster", environment)
		}
		if environment.DefaultNamespace == nil || *environment.DefaultNamespace == "" {
			return fmt.Errorf("a non-template %T must have a default namespace", environment)
		}
		if environment.Owner == nil || *environment.Owner == "" {
			return fmt.Errorf("a non-template %T must have an owner", environment)
		}
		if environment.RequiresSuitability == nil {
			return fmt.Errorf("a non-template %T must set whether it requires suitability or not", environment)
		}
	default:
		return fmt.Errorf("a %T must have a lifecycle of either 'template', 'static', or 'dynamic'", environment)
	}
	return nil
}

func postCreateEnvironment(db *gorm.DB, environment Environment, user *auth.User) error {
	if environment.Lifecycle == "dynamic" &&
		environment.ChartReleasesFromTemplate != nil &&
		*environment.ChartReleasesFromTemplate &&
		environment.TemplateEnvironmentID != nil {
		storeSet := NewStoreSet(db)
		// This is a dynamic environment that is getting created right now, let's copy the chart releases from the template too
		chartReleases, err := storeSet.ChartReleaseStore.ListAllMatching(
			ChartRelease{EnvironmentID: environment.TemplateEnvironmentID}, 0)
		if err != nil {
			return fmt.Errorf("wasn't able to list chart releases of template %s: %v", environment.TemplateEnvironment.Name, err)
		}
		for _, chartRelease := range chartReleases {
			_, err := storeSet.ChartReleaseStore.Create(
				ChartRelease{
					ChartID:                  chartRelease.ChartID,
					ClusterID:                environment.DefaultClusterID,
					DestinationType:          "environment",
					EnvironmentID:            &environment.ID,
					Name:                     fmt.Sprintf("%s-%s", chartRelease.Chart.Name, environment.Name),
					Namespace:                *environment.DefaultNamespace,
					CurrentAppVersionExact:   chartRelease.CurrentAppVersionExact,
					CurrentChartVersionExact: chartRelease.CurrentChartVersionExact,
					HelmfileRef:              chartRelease.HelmfileRef,
					TargetAppVersionBranch:   chartRelease.TargetAppVersionBranch,
					TargetAppVersionCommit:   chartRelease.TargetAppVersionCommit,
					TargetAppVersionExact:    chartRelease.TargetAppVersionExact,
					TargetAppVersionUse:      chartRelease.TargetAppVersionUse,
					TargetChartVersionExact:  chartRelease.TargetChartVersionExact,
					TargetChartVersionUse:    chartRelease.TargetChartVersionUse,
					ThelmaMode:               chartRelease.ThelmaMode,
				}, user)
			if err != nil {
				return fmt.Errorf("wasn't able to copy template's release of the %s chart: %v", chartRelease.Chart.Name, err)
			}
		}
	}
	return nil
}
