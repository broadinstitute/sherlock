package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

// databaseInstancesV3Delete godoc
//
//	@summary		Delete an individual DatabaseInstance
//	@description	Delete an individual DatabaseInstance by its selector.
//	@tags			DatabaseInstances
//	@produce		json
//	@param			selector				path		string	true	"The selector of the DatabaseInstance, which can be either a numeric ID or 'chart-release/' followed by a chart release selector."
//	@success		200						{object}	DatabaseInstanceV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/database-instances/v3/{selector} [delete]
func databaseInstancesV3Delete(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	query, err := databaseInstanceModelFromSelector(db, canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var result models.DatabaseInstance
	if err = db.Preload(clause.Associations).Where(&query).First(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	if err = db.Omit(clause.Associations).Delete(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, databaseInstanceFromModel(result))
}
