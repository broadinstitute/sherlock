package sherlock

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// serviceAlertV3Create godoc~
//
//	@summary		Create a service alert
//	@description	Create a service alert to be displayed within terra.
//	@tags			ServiceAlert
//	@produce		json
//	@param			serviceAlert			body		ServiceAlertV3Create	true	"The initial fields the ServiceAlert should have set"
//	@success		200						{object}	ServiceAlertV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/service-alerts/v3 [post]
func serviceAlertV3Create(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)

	if err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) Issue connecting to db", err))
		return
	}
	var body ServiceAlertV3
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	toCreate, err := body.toModel(db)
	if err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) Issue creating body", err))
		return
	}

	var severity_types = []string{"blocker", "critical", "minor"}
	if !slices.Contains(severity_types, *toCreate.Severity) {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) invalid severity", errors.BadRequest))
		return
	}
	toCreate.Uuid = utils.PointerTo(uuid.New())
	if toCreate.OnEnvironmentID == nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) environment is required", errors.BadRequest))
		return
	}

	var user *models.User
	if user, err = models.GetCurrentUserForDB(db); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("unable to get current user for creating service alert: %w", err))
		return
	}
	toCreate.CreatedBy = &user.Email

	if err = db.Create(&toCreate).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	if err = db.First(&toCreate, toCreate.ID).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, ServiceAlertFromModel(toCreate))

}
