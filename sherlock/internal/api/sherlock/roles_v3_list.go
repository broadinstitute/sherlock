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

// rolesV3List godoc
//
//	@summary		List Roles matching a filter
//	@description	List Roles matching a filter.
//	@tags			Roles
//	@produce		json
//	@param			filter					query		RoleV3	false	"Filter the returned Roles"
//	@param			limit					query		int		false	"Control how many Roles are returned (default 0, no limit)"
//	@param			offset					query		int		false	"Control the offset for the returned Roles (default 0)"
//	@success		200						{array}		RoleV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/roles/v3 [get]
func rolesV3List(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	var filter RoleV3
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
	var results []models.Role
	chain := db.
		Where(&modelFilter)
	if limit > 0 {
		chain = chain.Limit(limit)
	}
	if err = chain.
		Offset(offset).
		Order("name asc").
		Scopes(models.ReadRoleScope).
		Find(&results).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, utils.Map(results, roleFromModel))
}
