package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth/auth_models"
	"github.com/broadinstitute/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/internal/models/model_actions"
	"github.com/broadinstitute/sherlock/internal/utils"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type PagerdutyIntegration struct {
	gorm.Model
	PagerdutyID string
	Name        *string
	Key         *string
	Type        *string
}

func (p PagerdutyIntegration) TableName() string {
	return "v2_pagerduty_integrations"
}

func (p PagerdutyIntegration) getID() uint {
	return p.ID
}

var pagerdutyIntegrationStore *internalModelStore[PagerdutyIntegration]

func init() {
	pagerdutyIntegrationStore = &internalModelStore[PagerdutyIntegration]{
		selectorToQueryModel:  pagerdutyIntegrationSelectorToQuery,
		modelToSelectors:      pagerdutyIntegrationToSelectors,
		errorIfForbidden:      pagerdutyIntegrationErrorIfForbidden,
		validateModel:         validatePagerdutyIntegration,
		preDeletePostValidate: preDeletePostValidatePagerdutyIntegration,
	}
}

func pagerdutyIntegrationSelectorToQuery(_ *gorm.DB, selector string) (PagerdutyIntegration, error) {
	if len(selector) == 0 {
		return PagerdutyIntegration{}, fmt.Errorf("(%s) pagerduty integration selector cannot be empty", errors.BadRequest)
	}
	var query PagerdutyIntegration
	if utils.IsNumeric(selector) {
		id, err := strconv.Atoi(selector)
		if err != nil {
			return PagerdutyIntegration{}, fmt.Errorf("(%s) string to int conversion error of '%s': %v", errors.BadRequest, selector, err)
		}
		query.ID = uint(id)
		return query, nil
	} else if strings.Count(selector, "/") == 1 { // "pd-id" + pagerduty id
		parts := strings.Split(selector, "/")

		// "pd-id"
		// The reason we have this is to make sure we can tell the difference between our actual IDs and what
		// Pagerduty uses as IDs. They don't say anything about the contents of their string IDs so a prefix of
		// "pd-id/" seems like a good enough bet; better than just not numeric. See Environment's resource prefix
		// selector for another example of this same pattern.
		selectorLabel := parts[0]
		if selectorLabel != "pd-id" {
			return PagerdutyIntegration{}, fmt.Errorf("(%s) invalid pagerduty integration selector %s, pagerduty id selector needed to start with 'pd-id/' but was '%s/'", errors.BadRequest, selector, selectorLabel)
		}

		// pagerduty id
		pagerdutyID := parts[1]
		if len(pagerdutyID) == 0 {
			return PagerdutyIntegration{}, fmt.Errorf("(%s) invalid pagerduty integration selector %s, sub-selector was invalid (nothing after the '/')", errors.BadRequest, selector)
		}
		query.PagerdutyID = pagerdutyID
		return query, nil
	}
	return PagerdutyIntegration{}, fmt.Errorf("(%s) invalid pagerduty integration selector '%s'", errors.BadRequest, selector)
}

func pagerdutyIntegrationToSelectors(pagerdutyIntegration *PagerdutyIntegration) []string {
	var selectors []string
	if pagerdutyIntegration != nil {
		if pagerdutyIntegration.ID != 0 {
			selectors = append(selectors, fmt.Sprintf("%d", pagerdutyIntegration.ID))
		}
		if pagerdutyIntegration.PagerdutyID != "" {
			selectors = append(selectors, fmt.Sprintf("pd-id/%s", pagerdutyIntegration.PagerdutyID))
		}
	}
	return selectors
}

func pagerdutyIntegrationErrorIfForbidden(_ *gorm.DB, _ *PagerdutyIntegration, _ model_actions.ActionType, user *auth_models.User) error {
	return user.SuitableOrError()
}

func validatePagerdutyIntegration(pagerdutyIntegration *PagerdutyIntegration) error {
	if pagerdutyIntegration == nil {
		return fmt.Errorf("the model passed was nil")
	}
	if pagerdutyIntegration.PagerdutyID == "" {
		return fmt.Errorf("a %T must have a non-empty pagerduty id", pagerdutyIntegration)
	}
	if pagerdutyIntegration.Name == nil || *pagerdutyIntegration.Name == "" {
		return fmt.Errorf("a %T must have a non-empty name", pagerdutyIntegration)
	}
	if pagerdutyIntegration.Key == nil || *pagerdutyIntegration.Key == "" {
		return fmt.Errorf("a %T must have a non-empty key", pagerdutyIntegration)
	}
	if pagerdutyIntegration.Type == nil || *pagerdutyIntegration.Type == "" {
		return fmt.Errorf("a %T must have a non-empty type", pagerdutyIntegration)
	}
	return nil
}

func preDeletePostValidatePagerdutyIntegration(db *gorm.DB, pagerdutyIntegration *PagerdutyIntegration, _ *auth_models.User) error {
	chartReleases, err := chartReleaseStore.listAllMatchingByUpdated(db, 0, ChartRelease{PagerdutyIntegrationID: &pagerdutyIntegration.ID})
	if err != nil {
		return fmt.Errorf("wasn't able to check for chart releases that use this integration: %v", err)
	} else if len(chartReleases) > 0 {
		return fmt.Errorf("the following chart release uses this integration: %s", chartReleases[0].Name)
	}
	environments, err := environmentStore.listAllMatchingByUpdated(db, 0, Environment{PagerdutyIntegrationID: &pagerdutyIntegration.ID})
	if err != nil {
		return fmt.Errorf("wasn't able to check for environments that use this integration: %v", err)
	} else if len(environments) > 0 {
		return fmt.Errorf("the following environment uses this integration: %s", environments[0].Name)
	}
	return nil
}
