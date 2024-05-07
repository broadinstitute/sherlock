package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
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
	errors.AbortRequest(ctx, fmt.Errorf("not implemented"))
}
