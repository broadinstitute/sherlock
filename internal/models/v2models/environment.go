package v2models

import (
	"fmt"
	"strconv"

	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/errors"
	"gorm.io/gorm"
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
	BaseDomain          *string
	NamePrefixesDomain  *bool
	HelmfileRef         *string `gorm:"not null; default:null"`
}

func (e Environment) TableName() string {
	return "v2_environments"
}

var environmentStore *internalModelStore[Environment]

func init() {
	environmentStore = &internalModelStore[Environment]{
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

func environmentToSelectors(environment *Environment) []string {
	var selectors []string
	if environment != nil {
		if environment.Name != "" {
			selectors = append(selectors, environment.Name)
		}
		if environment.ID != 0 {
			selectors = append(selectors, fmt.Sprintf("%d", environment.ID))
		}
	}
	return selectors
}

func environmentRequiresSuitability(_ *gorm.DB, environment *Environment) bool {
	// RequiresSuitability is a required field and shouldn't ever actually be stored as nil, but if it is we fail-safe
	return environment.RequiresSuitability == nil || *environment.RequiresSuitability
}

func validateEnvironment(environment *Environment) error {
	if environment == nil {
		return fmt.Errorf("the model passed was nil")
	}
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

	if environment.HelmfileRef == nil || *environment.HelmfileRef == "" {
		return fmt.Errorf("a %T must have a non-empty terra-helmfile ref", environment)
	}
	return nil
}

func postCreateEnvironment(db *gorm.DB, environment *Environment, user *auth.User) error {
	if environment.Lifecycle == "dynamic" &&
		environment.ChartReleasesFromTemplate != nil &&
		*environment.ChartReleasesFromTemplate &&
		environment.TemplateEnvironmentID != nil {
		// This is a dynamic environment that is getting created right now, let's copy the chart releases from the template too
		chartReleases, err := chartReleaseStore.listAllMatching(db, 0, ChartRelease{EnvironmentID: environment.TemplateEnvironmentID})
		if err != nil {
			return fmt.Errorf("wasn't able to list chart releases of template %s: %v", environment.TemplateEnvironment.Name, err)
		}
		for _, chartRelease := range chartReleases {
			_, _, err := chartReleaseStore.create(db,
				ChartRelease{
					ChartID:             chartRelease.ChartID,
					ClusterID:           environment.DefaultClusterID,
					DestinationType:     "environment",
					EnvironmentID:       &environment.ID,
					Name:                fmt.Sprintf("%s-%s", chartRelease.Chart.Name, environment.Name),
					Namespace:           *environment.DefaultNamespace,
					ChartReleaseVersion: chartRelease.ChartReleaseVersion,
					Subdomain:           chartRelease.Subdomain,
					Protocol:            chartRelease.Protocol,
					Port:                chartRelease.Port,
				}, user)
			if err != nil {
				return fmt.Errorf("wasn't able to copy template's release of the %s chart: %v", chartRelease.Chart.Name, err)
			}
		}
	}
	return nil
}
