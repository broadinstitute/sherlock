package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type AppVersion struct {
	gorm.Model
	CiIdentifier       *CiIdentifier `gorm:"polymorphic:Resource; polymorphicValue:app-version"`
	Chart              *Chart
	ChartID            uint   `gorm:"not null: default:null"`
	AppVersion         string `gorm:"not null: default:null"`
	GitCommit          string
	GitBranch          string
	Description        string
	ParentAppVersion   *AppVersion
	ParentAppVersionID *uint
}

func (a AppVersion) getID() uint {
	return a.ID
}

func (a AppVersion) getParentID() *uint {
	return a.ParentAppVersionID
}

func (a AppVersion) GetCiIdentifier() *CiIdentifier {
	if a.CiIdentifier != nil {
		return a.CiIdentifier
	} else {
		return &CiIdentifier{ResourceType: "app-version", ResourceID: a.ID}
	}
}

var InternalAppVersionStore *internalTreeModelStore[AppVersion]

func init() {
	InternalAppVersionStore = &internalTreeModelStore[AppVersion]{
		internalModelStore: &internalModelStore[AppVersion]{
			selectorToQueryModel:    appVersionSelectorToQuery,
			modelToSelectors:        appVersionToSelectors,
			validateModel:           validateAppVersion,
			handleIncomingDuplicate: rejectDuplicateAppVersion,
		},
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
			return AppVersion{}, fmt.Errorf("(%s) string to int conversion error of '%s': %w", errors.BadRequest, selector, err)
		}
		query.ID = uint(id)
		return query, nil
	} else if strings.Count(selector, "/") == 1 { // chart + version
		parts := strings.Split(selector, "/")

		// chart
		chartID, err := InternalChartStore.ResolveSelector(db, parts[0])
		if err != nil {
			return AppVersion{}, fmt.Errorf("error handling chart sub-selector %s: %w", parts[0], err)
		}
		query.ChartID = chartID

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
