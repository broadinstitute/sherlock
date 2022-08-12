package v2handlers

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/controllers/v2controllers"
	"github.com/broadinstitute/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
	"strings"
)

// formatSelector standardizes the input string so the rest of Sherlock doesn't need to worry about, for instance,
// Gin including a leading slash when parsing a selector out of the URL.
func formatSelector(selector string) string {
	return strings.Trim(selector, "/")
}

func handleCreate[M v2models.Model, R v2controllers.Readable, C v2controllers.Creatable[R], E v2controllers.Editable[R, C]](controller *v2controllers.ModelController[M, R, C, E]) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user, err := auth.ExtractUserFromContext(ctx)
		if err != nil {
			ctx.JSON(errors.ErrorToApiResponse(err))
			return
		}
		var creatable C
		if err := ctx.ShouldBindJSON(&creatable); err != nil {
			ctx.JSON(errors.ErrorToApiResponse(fmt.Errorf("(%s) JSON error parsing to %T: %v", errors.BadRequest, creatable, err)))
			return
		}
		result, err := controller.Create(creatable, user)
		if err != nil {
			ctx.JSON(errors.ErrorToApiResponse(err))
			return
		}
		ctx.JSON(http.StatusCreated, result)
	}
}

func handleList[M v2models.Model, R v2controllers.Readable, C v2controllers.Creatable[R], E v2controllers.Editable[R, C]](controller *v2controllers.ModelController[M, R, C, E]) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var filter R
		if err := ctx.ShouldBindQuery(&filter); err != nil {
			ctx.JSON(errors.ErrorToApiResponse(fmt.Errorf("(%s) error parsing a filtering %T from the query parameters: %v", errors.BadRequest, filter, err)))
			return
		} else {
			log.Debug().Msgf("parsing query params to %T: '%s' => %+v", filter, ctx.Request.URL.RawQuery, filter)
		}
		limitString := ctx.DefaultQuery("limit", "0")
		limit, err := strconv.Atoi(limitString)
		if err != nil {
			ctx.JSON(errors.ErrorToApiResponse(fmt.Errorf("(%s) error parsing limit parameter: %v", errors.BadRequest, err)))
			return
		}
		result, err := controller.ListAllMatching(filter, limit)
		if err != nil {
			ctx.JSON(errors.ErrorToApiResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, result)
	}
}

func handleGet[M v2models.Model, R v2controllers.Readable, C v2controllers.Creatable[R], E v2controllers.Editable[R, C]](controller *v2controllers.ModelController[M, R, C, E]) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		result, err := controller.Get(formatSelector(ctx.Param("selector")))
		if err != nil {
			ctx.JSON(errors.ErrorToApiResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, result)
	}
}

func handleEdit[M v2models.Model, R v2controllers.Readable, C v2controllers.Creatable[R], E v2controllers.Editable[R, C]](controller *v2controllers.ModelController[M, R, C, E]) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user, err := auth.ExtractUserFromContext(ctx)
		if err != nil {
			ctx.JSON(errors.ErrorToApiResponse(err))
			return
		}
		var editable E
		if err := ctx.ShouldBindJSON(&editable); err != nil {
			ctx.JSON(errors.ErrorToApiResponse(fmt.Errorf("(%s) JSON error parsing to %T: %v", errors.BadRequest, editable, err)))
			return
		}
		result, err := controller.Edit(formatSelector(ctx.Param("selector")), editable, user)
		if err != nil {
			ctx.JSON(errors.ErrorToApiResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, result)
	}
}

func handleDelete[M v2models.Model, R v2controllers.Readable, C v2controllers.Creatable[R], E v2controllers.Editable[R, C]](controller *v2controllers.ModelController[M, R, C, E]) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user, err := auth.ExtractUserFromContext(ctx)
		if err != nil {
			ctx.JSON(errors.ErrorToApiResponse(err))
			return
		}
		result, err := controller.Delete(formatSelector(ctx.Param("selector")), user)
		if err != nil {
			ctx.JSON(errors.ErrorToApiResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, result)
	}
}

func handleSelectorList[M v2models.Model, R v2controllers.Readable, C v2controllers.Creatable[R], E v2controllers.Editable[R, C]](controller *v2controllers.ModelController[M, R, C, E]) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		result, err := controller.GetOtherValidSelectors(formatSelector(ctx.Param("selector")))
		if err != nil {
			ctx.JSON(errors.ErrorToApiResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, result)
	}
}
