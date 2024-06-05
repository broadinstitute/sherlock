package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

// environmentsV3Delete godoc
//
//	@summary		Delete an individual Environment
//	@description	Delete an individual Environment by its ID.
//	@tags			Environments
//	@produce		json
//	@param			selector				path		string	true	"The selector of the Environment, which can be either a numeric ID, the name, or 'resource-prefix' + / + the unique resource prefix."
//	@success		200						{object}	EnvironmentV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/environments/v3/{selector} [delete]
func environmentsV3Delete(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	query, err := environmentModelFromSelector(canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var result models.Environment
	if err = db.Preload(clause.Associations).Where(&query).First(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	if err = db.Omit(clause.Associations).Delete(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, environmentFromModel(result))
}
