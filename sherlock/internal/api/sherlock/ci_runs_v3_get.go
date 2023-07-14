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
	"strings"
)

func ciRunsV3Get(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	query, err := ciRunModelFromSelector(canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
		return
	}
	var result models.CiRun
	if err = db.Preload(clause.Associations).Where(&query).First(&result).Error; err != nil {
		ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, ciRunFromModel(result))
}

func ciRunModelFromSelector(selector string) (query models.CiRun, err error) {
	if len(selector) == 0 {
		return models.CiRun{}, fmt.Errorf("(%s) selector cannot be empty", errors.BadRequest)
	} else if utils.IsNumeric(selector) {
		query.ID, err = utils.ParseUint(selector)
		return query, err
	} else if strings.HasPrefix(selector, "github-actions/") && strings.Count(selector, "/") == 4 {
		// "github-actions" + owner + repo + run ID + attempt number
		parts := strings.Split(selector, "/")
		query.Platform = "github-actions"
		query.GithubActionsOwner = parts[1]
		if query.GithubActionsOwner == "" {
			return models.CiRun{}, fmt.Errorf("(%s) invalid CI run selector %s, owner sub-selector was empty", errors.BadRequest, selector)
		}
		query.GithubActionsRepo = parts[2]
		if query.GithubActionsRepo == "" {
			return models.CiRun{}, fmt.Errorf("(%s) invalid CI run selector %s, repo sub-selector was empty", errors.BadRequest, selector)
		}
		query.GithubActionsRunID, err = utils.ParseUint(parts[3])
		if err != nil {
			return models.CiRun{}, fmt.Errorf("(%s) invalid CI run selector %s, run ID sub-selector '%s' had string to int conversion error: %v", errors.BadRequest, selector, parts[3], err)
		}
		query.GithubActionsAttemptNumber, err = utils.ParseUint(parts[4])
		if err != nil {
			return models.CiRun{}, fmt.Errorf("(%s) invalid CI run selector %s, attempt number sub-selector '%s' had string to int conversion error: %v", errors.BadRequest, selector, parts[3], err)
		}
		return query, nil
	} else if strings.HasPrefix(selector, "argo-workflows/") && strings.Count(selector, "/") == 2 {
		// "argo-workflows" + namespace + name
		parts := strings.Split(selector, "/")
		query.Platform = "argo-workflows"
		query.ArgoWorkflowsNamespace = parts[1]
		if query.ArgoWorkflowsNamespace == "" {
			return models.CiRun{}, fmt.Errorf("(%s) invalid CI run selector %s, namespace sub-selector was empty", errors.BadRequest, selector)
		}
		query.ArgoWorkflowsName = parts[2]
		if query.ArgoWorkflowsName == "" {
			return models.CiRun{}, fmt.Errorf("(%s) invalid CI run selector %s, name sub-selector was empty", errors.BadRequest, selector)
		}
		return query, nil
	} else {
		return models.CiRun{}, fmt.Errorf("(%s) invalid CI run selector '%s'", errors.BadRequest, selector)
	}
}
