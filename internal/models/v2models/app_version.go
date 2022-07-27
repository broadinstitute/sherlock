package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/errors"
	"gorm.io/gorm"
	"strconv"
)

type AppVersion struct {
	gorm.Model
	Chart      Chart
	ChartID    uint   `gorm:"not null: default:null"`
	AppVersion string `gorm:"not null: default:null"`
	GitCommit  string
	GitBranch  string
}

func (c AppVersion) TableName() string {
	return "v2_app_versions"
}

func newAppVersionStore(db *gorm.DB) Store[AppVersion] {
	return Store[AppVersion]{
		db:                   db,
		selectorToQueryModel: appVersionSelectorToQuery,
		modelToSelectors:     appVersionToSelectors,
		validateModel:        validateAppVersion,
	}
}

func appVersionSelectorToQuery(_ *gorm.DB, selector string) (AppVersion, error) {
	var query AppVersion
	if isNumeric(selector) { // ID
		id, err := strconv.Atoi(selector)
		if err != nil {
			return AppVersion{}, fmt.Errorf("(%s) string to int conversion error of '%s': %v", errors.BadRequest, selector, err)
		}
		query.ID = uint(id)
		return query, nil
	}
	return AppVersion{}, fmt.Errorf("(%s) invalid app version selector '%s'", errors.BadRequest, selector)
}

func appVersionToSelectors(appVersion AppVersion) []string {
	var selectors []string
	if appVersion.ID != 0 {
		selectors = append(selectors, fmt.Sprintf("%d", appVersion.ID))
	}
	return selectors
}

func validateAppVersion(appVersion AppVersion) error {
	if appVersion.ChartID == 0 {
		return fmt.Errorf("an %T must have an associated chart", appVersion)
	}
	if appVersion.AppVersion == "" {
		return fmt.Errorf("an %T must have a non-empty app version", appVersion)
	}
	return nil
}
