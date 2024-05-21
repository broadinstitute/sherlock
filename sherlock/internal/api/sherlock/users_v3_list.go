package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// usersV3List godoc
//
//	@summary		List Users matching a filter
//	@description	List Users matching a filter. The results will include suitability and other information.
//	@description	Note that the suitability info can't directly be filtered for at this time.
//	@tags			Users
//	@produce		json
//	@param			filter					query		UserV3	false	"Filter the returned Users"
//	@param			limit					query		int		false	"Control how many Users are returned (default 0, no limit)"
//	@param			offset					query		int		false	"Control the offset for the returned Users (default 0)"
//	@success		200						{array}		UserV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/users/v3 [get]
func usersV3List(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	var filter UserV3
	if err = ctx.ShouldBindQuery(&filter); err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	modelFilter := filter.toModel()
	limit, err := utils.ParseInt(ctx.DefaultQuery("limit", "0"))
	if err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) %v", errors.BadRequest, err))
		return
	}
	offset, err := utils.ParseInt(ctx.DefaultQuery("offset", "0"))
	if err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) %v", errors.BadRequest, err))
		return
	}
	var results []models.User
	chain := db.Where(&modelFilter)
	if limit > 0 {
		chain = chain.Limit(limit)
	}
	if err = chain.
		Offset(offset).
		Order("email asc").
		Scopes(models.ReadUserScope).
		Find(&results).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, utils.Map(results, userFromModel))
}
