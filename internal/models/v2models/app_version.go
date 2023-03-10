package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/internal/utils"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type AppVersion struct {
	gorm.Model
	Chart              *Chart
	ChartID            uint   `gorm:"not null: default:null"`
	AppVersion         string `gorm:"not null: default:null"`
	GitCommit          string
	GitBranch          string
	Description        string
	ParentAppVersion   *AppVersion
	ParentAppVersionID *uint
}

func (a AppVersion) TableName() string {
	return "v2_app_versions"
}

func (a AppVersion) getID() uint {
	return a.ID
}

var appVersionStore *internalModelStore[AppVersion]

func init() {
	appVersionStore = &internalModelStore[AppVersion]{
		selectorToQueryModel:    appVersionSelectorToQuery,
		modelToSelectors:        appVersionToSelectors,
		validateModel:           validateAppVersion,
		handleIncomingDuplicate: rejectDuplicateAppVersion,
	}
}

func appVersionSelectorToQuery(db *gorm.DB, selector string) (AppVersion, error) {
	if len(selector) == 0 {
		return AppVersion{}, fmt.Errorf("(%s) app version selector cannot be empty", errors.BadRequest)
	}
	var query AppVersion
	if utils.IsNumeric(selector) { // ID
		id, err := strconv.Atoi(selector)
		if err != nil {
			return AppVersion{}, fmt.Errorf("(%s) string to int conversion error of '%s': %v", errors.BadRequest, selector, err)
		}
		query.ID = uint(id)
		return query, nil
	} else if strings.Count(selector, "/") == 1 { // chart + version
		parts := strings.Split(selector, "/")

		// chart
		chartQuery, err := chartSelectorToQuery(db, parts[0])
		if err != nil {
			return AppVersion{}, fmt.Errorf("invalid chart release selector %s, chart sub-selector error: %v", selector, err)
		}
		chart, err := chartStore.get(db, chartQuery)
		if err != nil {
			return AppVersion{}, fmt.Errorf("error handling chart sub-selector %s: %v", parts[0], err)
		}
		query.ChartID = chart.ID

		// version
		version := parts[1]
		if len(version) == 0 {
			return AppVersion{}, fmt.Errorf("(%s) invalid app version selector %s, version sub-selector was empty", errors.BadRequest, selector)
		}
		query.AppVersion = version

		return query, nil
	}
	return AppVersion{}, fmt.Errorf("(%s) invalid app version selector '%s'", errors.BadRequest, selector)
}

func appVersionToSelectors(appVersion *AppVersion) []string {
	var selectors []string
	if appVersion != nil {
		if appVersion.ID != 0 {
			selectors = append(selectors, fmt.Sprintf("%d", appVersion.ID))
		}
		if appVersion.AppVersion != "" {
			chartSelectors := chartToSelectors(appVersion.Chart)
			if len(chartSelectors) == 0 && appVersion.ChartID != 0 {
				// Chart not filled so chartToSelectors gives nothing, but we have the chart ID and it is a selector anyway
				chartSelectors = []string{fmt.Sprintf("%d", appVersion.ChartID)}
			}
			for _, chartSelector := range chartSelectors {
				selectors = append(selectors, fmt.Sprintf("%s/%s", chartSelector, appVersion.AppVersion))
			}
		}
	}
	return selectors
}

func validateAppVersion(appVersion *AppVersion) error {
	if appVersion == nil {
		return fmt.Errorf("the model passed was nil")
	}
	if appVersion.ChartID == 0 {
		return fmt.Errorf("an %T must have an associated chart", appVersion)
	}
	if appVersion.AppVersion == "" {
		return fmt.Errorf("an %T must have a non-empty app version", appVersion)
	}
	return nil
}

func rejectDuplicateAppVersion(existing *AppVersion, new *AppVersion) error {
	if existing.AppVersion != new.AppVersion {
		return fmt.Errorf("new %T has chart version '%s', which is mismatched with the existing value of %s", new, new.AppVersion, existing.AppVersion)
	}
	if existing.ChartID != new.ChartID {
		return fmt.Errorf("new %T has chart ID '%d', which is mismatched with the existing value of '%d'", new, new.ChartID, existing.ChartID)
	}
	if existing.GitBranch != new.GitBranch {
		return fmt.Errorf("new %T has git branch '%s', which is mismatched with the existing value of '%s'", new, new.GitBranch, existing.GitBranch)
	}
	if existing.GitCommit != new.GitCommit {
		return fmt.Errorf("new %T has git commit '%s', which is mismatched with the existing value of '%s'", new, new.GitCommit, existing.GitCommit)
	}
	if (existing.ParentAppVersionID != nil) && (new.ParentAppVersionID == nil) {
		return fmt.Errorf("new %T has no parent ID, which is mismatched with the existing having one", new)
	} else if (existing.ParentAppVersionID == nil) && (new.ParentAppVersionID != nil) {
		return fmt.Errorf("new %T has a parent ID, which is mismatched with the existing not having one", new)
	} else if existing.ParentAppVersionID != nil && new.ParentAppVersionID != nil && *existing.ParentAppVersionID != *new.ParentAppVersionID {
		return fmt.Errorf("new %T has parent ID '%d', which is mismatched with the existing value of '%d'", new, *new.ParentAppVersionID, *existing.ParentAppVersionID)
	}
	return nil
}
