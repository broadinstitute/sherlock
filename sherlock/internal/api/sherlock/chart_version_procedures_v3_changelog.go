package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

type ChartVersionV3ChangelogResponse struct {
	Changelog []ChartVersionV3 `json:"changelog"`
	Complete  bool             `json:"complete"`
}

// chartVersionsProceduresV3Changelog godoc
//
//	@summary		Get a changelog between two ChartVersions
//	@description	Get the path through parent references from a child ChartVersion (inclusive) to a parent ChartVersion (exclusive), if possible. Because parent references point from newer children to older parents, the newer ChartVersion should be the child. The result will always exclude the parent.
//	@tags			ChartVersions
//	@produce		json
//	@param			child					query		string	true	"The selector of the newer ChartVersion for the changelog"
//	@param			parent					query		string	true	"The selector of the older ChartVersion for the changelog"
//	@success		200						{object}	ChartVersionV3ChangelogResponse
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/chart-versions/procedures/v3/changelog [get]
func chartVersionsProceduresV3Changelog(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	var childQuery, parentQuery models.ChartVersion

	if childQuery, err = chartVersionModelFromSelector(db, canonicalizeSelector(ctx.Query("child"))); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error parsing child: %w", err))
		return
	}

	if parentQuery, err = chartVersionModelFromSelector(db, canonicalizeSelector(ctx.Query("parent"))); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error parsing parent: %w", err))
		return
	}

	var child, parent models.ChartVersion

	if err = db.Select("id").Where(&childQuery).First(&child).Error; err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error querying child %s: %w", ctx.Query("child"), err))
		return
	}

	if err = db.Select("id").Where(&parentQuery).First(&parent).Error; err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error querying parent %s: %w", ctx.Query("parent"), err))
		return
	}

	var path []uint
	var foundPath bool

	if path, foundPath, err = models.GetChartVersionPathIDs(db, child.ID, parent.ID); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) error calculating changelog components: %w", errors.InternalServerError, err))
		return
	}

	response := ChartVersionV3ChangelogResponse{
		Complete: foundPath,
	}

	if foundPath {
		var pathModels []models.ChartVersion
		if len(path) > 0 {
			if err = db.Preload(clause.Associations).Order("created_at desc").Find(&pathModels, path).Error; err != nil {
				errors.AbortRequest(ctx, fmt.Errorf("error querying data of calculated changelog components: %w", err))
				return
			}
		}
		response.Changelog = utils.Map(pathModels, func(m models.ChartVersion) ChartVersionV3 { return chartVersionFromModel(m) })
	} else {
		if err = db.Preload(clause.Associations).First(&child, child.ID).Error; err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("error querying data of child: %w", err))
			return
		} else {
			response.Changelog = []ChartVersionV3{chartVersionFromModel(child)}
		}
	}

	ctx.JSON(http.StatusOK, response)
}
