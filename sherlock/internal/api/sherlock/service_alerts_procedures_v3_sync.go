package sherlock

import (
	"fmt"

	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
)

type ServiceAlertV3SyncRequest struct {
	OnEnvironment *string `json:"onEnvironment,omitempty" form:"onEnvironment"`
}

// usersProceduresV3Deactivate godoc
//
//	@summary		Sync service alerts
//	@description	Method to get all currently active service alerts from Sherlock's DB and ensure that the service alert json files placed in Google Buckets for Terra match.
//	@tags			ServiceAlert
//	@accept			json
//	@produce		json
//	@param			environment					body		ServiceAlertV3SyncRequest	true	"Information on Service Alert environment"
//	@success		200						{array}		ServiceAlertV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/service-alerts/procedures/v3/sync [post]
func syncServiceAlerts(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	var body ServiceAlertV3SyncRequest
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	var activeAlerts []models.ServiceAlert
	// Only return service alerts that haven't been deleted
	if err = db.Model(&models.ServiceAlert{}).Where("DeletedAt = ''").Find(&activeAlerts).Error; err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) error querying for Service Alerts: %w", errors.InternalServerError, err))
		return
	}

}
