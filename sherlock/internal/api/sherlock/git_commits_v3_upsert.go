package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type GitCommitV3Upsert struct {
	GitRepo      string    `json:"gitRepo"`
	GitCommit    string    `json:"gitCommit"`
	GitBranch    string    `json:"gitBranch"`
	IsMainBranch bool      `json:"isMainBranch"`
	CommittedAt  time.Time `json:"committedAt"`
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
		errors.AbortRequest(ctx, fmt.Errorf("failed to query for previous GitCommit: %w", err))
		return
	}

	var timeSince *uint

	if len(previous) > 0 && !previous[0].CommittedAt.IsZero() && !body.CommittedAt.IsZero() && previous[0].CommittedAt.Before(body.CommittedAt) {
		timeSince = utils.PointerTo(uint(body.CommittedAt.Sub(previous[0].CommittedAt).Seconds()))
	}

	var result models.GitCommit
	where := models.GitCommit{GitRepo: body.GitRepo, GitBranch: body.GitBranch, GitCommit: body.GitCommit}
	attrs := models.GitCommit{IsMainBranch: body.IsMainBranch, SecSincePrev: timeSince, CommittedAt: body.CommittedAt}
	if err = db.Where(&where).Attrs(&attrs).FirstOrCreate(&result).Error; err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("failed to upsert GitCommit (WHERE %+v, ATTRS %+v): %w", where, attrs, err))
		return
	}

	ctx.JSON(http.StatusCreated, gitCommitFromModel(result))
}
