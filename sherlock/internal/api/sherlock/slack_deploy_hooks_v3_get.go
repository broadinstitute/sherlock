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

// slackDeployHooksV3Get godoc
//
//	@summary		Get an individual SlackDeployHook
//	@description	Get an individual SlackDeployHook by its ID.
//	@tags			DeployHooks
//	@produce		json
//	@param			selector				path		string	true	"The ID of the SlackDeployHook"
//	@success		200						{object}	SlackDeployHookV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/deploy-hooks/slack/v3/{selector} [get]
func slackDeployHooksV3Get(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	query, err := slackDeployHookModelFromSelector(canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var result models.SlackDeployHook
	if err = db.Preload(clause.Associations).Where(&query).First(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, slackDeployHookFromModel(result))
}

func slackDeployHookModelFromSelector(selector string) (query models.SlackDeployHook, err error) {
	if len(selector) == 0 {
		return models.SlackDeployHook{}, fmt.Errorf("(%s) selector cannot be empty", errors.BadRequest)
	} else if utils.IsNumeric(selector) {
		query.ID, err = utils.ParseUint(selector)
		return query, err
	} else {
		return models.SlackDeployHook{}, fmt.Errorf("(%s) invalid slack deploy hook selector '%s'", errors.BadRequest, selector)
	}
}
