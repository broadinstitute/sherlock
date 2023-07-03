package v2handlers

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/auth"
	v2controllers2 "github.com/broadinstitute/sherlock/sherlock/internal/deprecated_controllers/v2controllers"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/pagerduty"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

func handleCreate[M v2models.Model, R v2controllers2.Readable[M], C v2controllers2.Creatable[M], E v2controllers2.Editable[M]](controller *v2controllers2.ModelController[M, R, C, E]) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user, err := auth.ShouldUseUser(ctx)
		if err != nil {
			ctx.JSON(errors.ErrorToApiResponse(err))
			return
		}
		var creatable C
		if err := ctx.ShouldBindJSON(&creatable); err != nil {
			ctx.JSON(errors.ErrorToApiResponse(fmt.Errorf("(%s) JSON error parsing to %T: %v", errors.BadRequest, creatable, err)))
			return
		}
		result, created, err := controller.Create(creatable, user)
		if err != nil {
			ctx.JSON(errors.ErrorToApiResponse(err))
			return
		}
		if created {
			ctx.JSON(http.StatusCreated, result)
		} else {
			ctx.JSON(http.StatusOK, result)
		}
	}
}

func handleList[M v2models.Model, R v2controllers2.Readable[M], C v2controllers2.Creatable[M], E v2controllers2.Editable[M]](controller *v2controllers2.ModelController[M, R, C, E]) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var filter R
		if err := ctx.ShouldBindQuery(&filter); err != nil {
			ctx.JSON(errors.ErrorToApiResponse(fmt.Errorf("(%s) error parsing a filtering %T from the query parameters: %v", errors.BadRequest, filter, err)))
			return
		} else {
			log.Trace().Msgf("parsing query params to %T: '%s' => %+v", filter, ctx.Request.URL.RawQuery, filter)
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

func handleGet[M v2models.Model, R v2controllers2.Readable[M], C v2controllers2.Creatable[M], E v2controllers2.Editable[M]](controller *v2controllers2.ModelController[M, R, C, E]) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		result, err := controller.Get(formatSelector(ctx.Param("selector")))
		if err != nil {
			ctx.JSON(errors.ErrorToApiResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, result)
	}
}

func handleEdit[M v2models.Model, R v2controllers2.Readable[M], C v2controllers2.Creatable[M], E v2controllers2.Editable[M]](controller *v2controllers2.ModelController[M, R, C, E]) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user, err := auth.ShouldUseUser(ctx)
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

func handleUpsert[M v2models.Model, R v2controllers2.Readable[M], C v2controllers2.Creatable[M], E v2controllers2.Editable[M]](controller *v2controllers2.ModelController[M, R, C, E]) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user, err := auth.ShouldUseUser(ctx)
		if err != nil {
			ctx.JSON(errors.ErrorToApiResponse(err))
			return
		}
		var creatable C
		if err := ctx.ShouldBindBodyWith(&creatable, binding.JSON); err != nil {
			ctx.JSON(errors.ErrorToApiResponse(fmt.Errorf("(%s) JSON error parsing to %T: %v", errors.BadRequest, creatable, err)))
			return
		}
		var editable E
		if err := ctx.ShouldBindBodyWith(&editable, binding.JSON); err != nil {
			ctx.JSON(errors.ErrorToApiResponse(fmt.Errorf("(%s) JSON error parsing to %T: %v", errors.BadRequest, editable, err)))
			return
		}
		result, created, err := controller.Upsert(formatSelector(ctx.Param("selector")), creatable, editable, user)
		if err != nil {
			ctx.JSON(errors.ErrorToApiResponse(err))
			return
		}
		if created {
			ctx.JSON(http.StatusCreated, result)
		} else {
			ctx.JSON(http.StatusOK, result)
		}
	}
}

func handleDelete[M v2models.Model, R v2controllers2.Readable[M], C v2controllers2.Creatable[M], E v2controllers2.Editable[M]](controller *v2controllers2.ModelController[M, R, C, E]) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user, err := auth.ShouldUseUser(ctx)
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

func handleSelectorList[M v2models.Model, R v2controllers2.Readable[M], C v2controllers2.Creatable[M], E v2controllers2.Editable[M]](controller *v2controllers2.ModelController[M, R, C, E]) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		result, err := controller.GetOtherValidSelectors(formatSelector(ctx.Param("selector")))
		if err != nil {
			ctx.JSON(errors.ErrorToApiResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, result)
	}
}

func handleTriggerPagerdutyIncident[M v2models.Model, R v2controllers2.Readable[M], C v2controllers2.Creatable[M], E v2controllers2.Editable[M]](controller *v2controllers2.ModelController[M, R, C, E]) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		_, err := auth.ShouldUseUser(ctx)
		if err != nil {
			ctx.JSON(errors.ErrorToApiResponse(err))
			return
		}
		var request pagerduty.AlertSummary
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(errors.ErrorToApiResponse(fmt.Errorf("(%s) JSON error parsing to %T: %v", errors.BadRequest, request, err)))
			return
		}
		result, err := controller.TriggerPagerdutyIncident(formatSelector(ctx.Param("selector")), request)
		if err != nil {
			ctx.JSON(errors.ErrorToApiResponse(err))
		}
		ctx.JSON(http.StatusAccepted, result)
	}
}

func handleGetChildrenPathToParent[M v2models.TreeModel, R v2controllers2.Readable[M], C v2controllers2.Creatable[M], E v2controllers2.Editable[M]](controller *v2controllers2.TreeModelController[M, R, C, E]) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		results, connected, err := controller.GetChildrenPathToParent(ctx.Query("child"), ctx.Query("parent"))
		if err != nil {
			ctx.JSON(errors.ErrorToApiResponse(err))
			return
		}
		if connected {
			ctx.JSON(http.StatusOK, results)
		} else {
			ctx.JSON(http.StatusNoContent, results)
		}
	}
}
