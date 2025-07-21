package sherlock

import (
	"fmt"
	"net/http"

	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type AppVersionV3ChangelogResponse struct {
	Changelog []AppVersionV3 `json:"changelog"`
	Complete  bool           `json:"complete"`
}

// appVersionsProceduresV3Changelog godoc
//
//	@summary		Get a changelog between two AppVersions
//	@description	Get the path through parent references from a child AppVersion (inclusive) to a parent AppVersion (exclusive), if possible. Because parent references point from newer children to older parents, the newer AppVersion should be the child. The result will always exclude the parent.
//	@tags			AppVersions
//	@produce		json
//	@param			child					query		string	true	"The selector of the newer AppVersion for the changelog"
//	@param			parent					query		string	true	"The selector of the older AppVersion for the changelog"
//	@success		200						{object}	AppVersionV3ChangelogResponse
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/app-versions/procedures/v3/changelog [get]
func appVersionsProceduresV3Changelog(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	var childQuery, parentQuery models.AppVersion

	if childQuery, err = appVersionModelFromSelector(db, canonicalizeSelector(ctx.Query("child"))); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error parsing child: %w", err))
		return
	}

	if parentQuery, err = appVersionModelFromSelector(db, canonicalizeSelector(ctx.Query("parent"))); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error parsing parent: %w", err))
		return
	}

	var child, parent models.AppVersion

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

	if path, foundPath, err = models.GetAppVersionPathIDs(db, parent.ID, child.ID); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) error calculating changelog components: %w", errors.InternalServerError, err))
		return
	}

	response := AppVersionV3ChangelogResponse{
		Complete: foundPath,
	}

	if foundPath {
		var pathModels []models.AppVersion
		if len(path) > 0 {
			if err = db.Preload(clause.Associations).Order("created_at desc").Find(&pathModels, path).Error; err != nil {
				errors.AbortRequest(ctx, fmt.Errorf("error querying data of calculated changelog components: %w", err))
				return
			}
		}
		response.Changelog = utils.Map(pathModels, func(m models.AppVersion) AppVersionV3 { return appVersionFromModel(m) })
	} else {
		if err = db.Preload(clause.Associations).First(&child, child.ID).Error; err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("error querying data of child: %w", err))
			return
		} else {
			response.Changelog = []AppVersionV3{appVersionFromModel(child)}
		}
	}

	ctx.JSON(http.StatusOK, response)
}
