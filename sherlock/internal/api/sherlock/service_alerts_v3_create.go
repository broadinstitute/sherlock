package sherlock

import (
	"fmt"
	"net/http"

	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/gin-gonic/gin"
)

func serviceAlertV3Create(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	var body ServiceAlertV3
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	toCreate := body.toModel(db)
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
