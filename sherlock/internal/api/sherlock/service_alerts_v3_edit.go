package sherlock

import (
	"fmt"
	"net/http"

	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
)

// serviceAlertV3Edit godoc~
//
//	@summary		Edit a service alert
//	@description	Update a service alert with new information.
//	@tags			ServiceAlert
//	@produce		json
//	@param			selector				path		string							true	"The selector of the ServiceAlert, which is the guid for a given alert"
//	@param			service-alert			body		ServiceAlertV3EditableFields	true	"The edits to make to the ServiceAlert"
//	@success		200						{object}	ServiceAlertV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/service-alerts/v3/{selector} [patch]
func serviceAlertV3Edit(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	query, err := serviceAlertFromSelector(canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	var body ServiceAlertV3
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	edits, err := body.toModel(db)
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	var toEdit models.ServiceAlert
	if err = db.Where(&query).First(&toEdit).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	var user *models.User
	if user, err = models.GetCurrentUserForDB(db); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("unable to get current user for deleting service alert: %w", err))
		return
	}
	edits.UpdatedBy = &user.Email

	if err = db.Model(&toEdit).Updates(&edits).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, ServiceAlertFromModel(toEdit))

}
