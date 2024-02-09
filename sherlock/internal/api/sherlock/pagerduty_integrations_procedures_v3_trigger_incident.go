package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/pagerduty"
	"github.com/gin-gonic/gin"
	"net/http"
)

// pagerdutyIntegrationsV3Get godoc
//
//	@summary		Get an individual PagerdutyIntegration
//	@description	Get an individual PagerdutyIntegration.
//	@tags			PagerdutyIntegrations
//	@produce		json
//	@param			selector				path		string					true	"The selector of the PagerdutyIntegration, which can be either a numeric ID or pd-id/<pagerduty-id>."
//	@param			summary					body		pagerduty.AlertSummary	true	"Summary of the incident"
//	@success		200						{object}	pagerduty.SendAlertResponse
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/pagerduty-integrations/procedures/v3/trigger-incident/{selector} [post]
func pagerdutyIntegrationsProceduresV3TriggerIncident(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	query, err := pagerdutyIntegrationModelFromSelector(canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var request pagerduty.AlertSummary
	if err = ctx.ShouldBindJSON(&request); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) JSON error parsing to %T: %w", errors.BadRequest, request, err))
		return
	}
	var pagerdutyIntegration models.PagerdutyIntegration
	if err = db.Where(&query).First(&pagerdutyIntegration).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	if pagerdutyIntegration.Key == nil {
		// Impossible to hit this case because of the model validation, but we'll check anyway to prevent NPE
		errors.AbortRequest(ctx, fmt.Errorf("(%s) PagerdutyIntegration %d lacks key", errors.BadRequest, pagerdutyIntegration.ID))
		return
	}
	if request.SourceLink == "" {
		// If we don't have a source link, we're still trying to declare an incident, so now isn't the time to quibble about input validation
		request.SourceLink = "https://broad.io/beehive"
	}

	response, err := pagerduty.SendAlert(*pagerdutyIntegration.Key, request)
	if err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error from pagerduty: %w", err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
