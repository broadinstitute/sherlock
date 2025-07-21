package sherlock

import (
	"fmt"
	"net/http"

	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

// pagerdutyIntegrationsV3Edit godoc
//
//	@summary		Edit an individual PagerdutyIntegration
//	@description	Edit an individual PagerdutyIntegration.
//	@tags			PagerdutyIntegrations
//	@produce		json
//	@param			selector				path		string						true	"The selector of the PagerdutyIntegration, which can be either a numeric ID or pd-id/<pagerduty-id>."
//	@param			pagerdutyIntegration	body		PagerdutyIntegrationV3Edit	true	"The edits to make to the PagerdutyIntegration"
//	@success		200						{object}	PagerdutyIntegrationV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/pagerduty-integrations/v3/{selector} [patch]
func pagerdutyIntegrationsV3Edit(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	query, err := pagerdutyIntegrationModelFromSelector(canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	var body PagerdutyIntegrationV3Edit
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	edits := body.toModel()

	var toEdit models.PagerdutyIntegration
	if err = db.Preload(clause.Associations).Where(&query).First(&toEdit).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	if err = db.Model(&toEdit).Omit(clause.Associations).Updates(&edits).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, pagerdutyIntegrationFromModel(toEdit))
}
