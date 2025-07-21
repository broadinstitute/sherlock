package sherlock

import (
	"net/http"

	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

// pagerdutyIntegrationsV3Delete godoc
//
//	@summary		Delete an individual PagerdutyIntegration
//	@description	Delete an individual PagerdutyIntegration by its ID.
//	@tags			PagerdutyIntegrations
//	@produce		json
//	@param			selector				path		string	true	"The selector of the PagerdutyIntegration, which can be either a numeric ID or pd-id/<pagerduty-id>."
//	@success		200						{object}	PagerdutyIntegrationV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/pagerduty-integrations/v3/{selector} [delete]
func pagerdutyIntegrationsV3Delete(ctx *gin.Context) {
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
	if err = db.Omit(clause.Associations).Delete(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, pagerdutyIntegrationFromModel(result))
}
