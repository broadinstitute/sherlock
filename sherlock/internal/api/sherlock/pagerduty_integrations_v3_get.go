package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
	"strings"
)

// pagerdutyIntegrationsV3Get godoc
//
//	@summary		Get an individual PagerdutyIntegration
//	@description	Get an individual PagerdutyIntegration.
//	@tags			PagerdutyIntegrations
//	@produce		json
//	@param			selector				path		string	true	"The selector of the PagerdutyIntegration, which can be either a numeric ID or pd-id/<pagerduty-id>."
//	@success		200						{object}	PagerdutyIntegrationV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/pagerduty-integrations/v3/{selector} [get]
func pagerdutyIntegrationsV3Get(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	query, err := pagerdutyIntegrationModelFromSelector(canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var result models.PagerdutyIntegration
	if err = db.Preload(clause.Associations).Where(&query).First(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, pagerdutyIntegrationFromModel(result))
}

func pagerdutyIntegrationModelFromSelector(selector string) (query models.PagerdutyIntegration, err error) {
	if len(selector) == 0 {
		return models.PagerdutyIntegration{}, fmt.Errorf("(%s) pagerduty integration selector cannot be empty", errors.BadRequest)
	}
	if utils.IsNumeric(selector) { // ID
		query.ID, err = utils.ParseUint(selector)
		return query, err
	} else if strings.Count(selector, "/") == 1 { // "pd-id" + pagerduty id
		parts := strings.Split(selector, "/")

		// "pd-id"
		// The reason we have this is to make sure we can tell the difference between our actual IDs and what
		// Pagerduty uses as IDs. They don't say anything about the contents of their string IDs so a prefix of
		// "pd-id/" seems like a good enough bet; better than just not numeric. See Environment's resource prefix
		// selector for another example of this same pattern.
		selectorLabel := parts[0]
		if selectorLabel != "pd-id" {
			return models.PagerdutyIntegration{}, fmt.Errorf("(%s) invalid pagerduty integration selector %s, pagerduty id selector needed to start with 'pd-id/' but was '%s/'", errors.BadRequest, selector, selectorLabel)
		}

		// pagerduty id
		pagerdutyID := parts[1]
		if len(pagerdutyID) == 0 {
			return models.PagerdutyIntegration{}, fmt.Errorf("(%s) invalid pagerduty integration selector %s, sub-selector was invalid (nothing after the '/')", errors.BadRequest, selector)
		}
		query.PagerdutyID = pagerdutyID
		return query, nil
	}
	return models.PagerdutyIntegration{}, fmt.Errorf("(%s) invalid pagerduty integration selector '%s'", errors.BadRequest, selector)
}
