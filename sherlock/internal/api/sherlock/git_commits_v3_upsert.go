package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GitCommitV3Upsert struct {
	CommonFields
	GitRepo      string `json:"gitRepo"`
	GitCommit    string `json:"gitCommit"`
	GitBranch    string `json:"gitBranch"`
	IsMainBranch bool   `json:"isMainBranch"`
}

// gitCommitsV3Upsert godoc
//
//	@summary		Upsert a GitCommit
//	@description	Upsert a GitCommit.
//	@tags			GitCommits
//	@accept			json
//	@produce		json
//	@param			gitCommit				body		GitCommitV3Upsert	true	"The GitCommit to upsert"
//	@success		201						{object}	GitCommitV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/git-commits/v3 [put]
func gitCommitsV3Upsert(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	var body GitCommitV3Upsert
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	var previous []models.GitCommit
	if err = db.
		Where(&models.GitCommit{GitRepo: body.GitRepo, GitBranch: body.GitBranch}).
		Limit(1).
		Order("created_at desc").
		Find(&previous).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	var timeSince *uint

	if len(previous) > 0 {
		timeSince = utils.PointerTo(uint(body.CreatedAt.Sub(previous[0].CreatedAt).Seconds()))
	}

	var result models.GitCommit
	if err = db.Where(&models.GitCommit{GitRepo: body.GitRepo, GitBranch: body.GitBranch, GitCommit: body.GitCommit}).
		Attrs(&models.GitCommit{
			IsMainBranch: body.IsMainBranch,
			SecSincePrev: timeSince,
		}).FirstOrCreate(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gitCommitFromModel(result))
}
