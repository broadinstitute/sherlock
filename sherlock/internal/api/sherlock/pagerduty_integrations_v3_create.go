package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/creasty/defaults"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

// pagerdutyIntegrationsV3Create godoc
//
//	@summary		Create a PagerdutyIntegration
//	@description	Create a PagerdutyIntegration. Duplicate Pagerduty IDs will be gracefully handled by editing the existing entry. This is partially opaque because some fields are writable but not readable.
//	@tags			PagerdutyIntegrations
//	@accept			json
//	@produce		json
//	@param			pagerdutyIntegration	body		PagerdutyIntegrationV3Create	true	"The PagerdutyIntegration to create"
//	@success		201						{object}	PagerdutyIntegrationV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/pagerduty-integrations/v3 [post]
func pagerdutyIntegrationsV3Create(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	var body PagerdutyIntegrationV3Create
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	if err = defaults.Set(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error setting defaults: %w", err))
		return
	}

	toCreate := body.toModel()

	var result models.PagerdutyIntegration
	if err = db.Where(&models.PagerdutyIntegration{
		PagerdutyID: toCreate.PagerdutyID,
	}).Assign(&models.PagerdutyIntegration{
		Name: toCreate.Name,
		Key:  toCreate.Key,
		Type: toCreate.Type,
	}).Preload(clause.Associations).FirstOrCreate(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, pagerdutyIntegrationFromModel(result))
}
