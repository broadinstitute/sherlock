package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/errors"
	"gorm.io/gorm"
	"strconv"
)

type Environment struct {
	gorm.Model
	Base                  string
	Lifecycle             string
	Name                  string `gorm:"not null; default:null; unique"`
	TemplateEnvironment   *Environment
	TemplateEnvironmentID *uint
	ValuesName            string
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

func newEnvironmentStore(db *gorm.DB) Store[Environment] {
	return Store[Environment]{
		db:                   db,
		selectorToQueryModel: environmentSelectorToQuery,
		modelToSelectors:     environmentToSelectors,
	}
}

func environmentSelectorToQuery(_ *gorm.DB, selector string) (Environment, error) {
	var query Environment
	if isNumeric(selector) { // ID
		id, err := strconv.Atoi(selector)
		if err != nil {
			return Environment{}, fmt.Errorf("(%s) string to int conversion error of '%s': %v", errors.BadRequest, selector, err)
		}
		query.ID = uint(id)
		return query, nil
	} else if isAlphaNumericWithHyphens(selector) &&
		len(selector) > 0 && len(selector) <= 32 &&
		isStartingWithLetter(selector) &&
		isEndingWithAlphaNumeric(selector) { // Name
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

func environmentRequiresSuitability(environment Environment) bool {
	return environment.RequiresSuitability == nil || *environment.RequiresSuitability
}
