package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
)

// rolesV3Get godoc
//
//	@summary		Get a Role
//	@description	Get an individual Role and the Users assigned to it.
//	@tags			Roles
//	@produce		json
//	@param			id						path		uint	true	"The numeric ID of the role"
//	@success		200						{object}	RoleV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/roles/v3/{id} [get]
func rolesV3Get(ctx *gin.Context) {
	errors.AbortRequest(ctx, fmt.Errorf("not implemented"))
}
