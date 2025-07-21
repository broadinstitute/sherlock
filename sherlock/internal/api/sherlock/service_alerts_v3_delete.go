package sherlock

import (
	"fmt"
	"net/http"

	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
)

// serviceAlertV3Delete godoc
//
//	@summary		Delete a ServiceAlert
//	@description	Delete an individual ServiceAlert.
//	@tags			ServiceAlert
//	@produce		json
//	@param			selector				path		string	true	"The selector of the ServiceAlert, ServiceAlert, which is the guid for a given alert"
//	@success		200						{object}	ServiceAlertV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/service-alerts/v3/{selector} [delete]
func serviceAlertV3Delete(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	query, err := serviceAlertFromSelector(canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var result models.ServiceAlert
	if err = db.Where(&query).First(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	var user *models.User
	if user, err = models.GetCurrentUserForDB(db); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("unable to get current user for deleting service alert: %w", err))
		return
	}
	result.DeletedById = &user.ID

	if err = db.Delete(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, ServiceAlertFromModel(result))
}
