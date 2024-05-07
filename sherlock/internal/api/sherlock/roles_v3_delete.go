package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
)

// rolesV3Delete godoc
//
//	@summary		Delete a Role
//	@description	Delete an individual Role.
//	@tags			Roles
//	@produce		json
//	@param			id						path		uint	true	"The numeric ID of the role"
//	@success		200						{object}	RoleV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/roles/v3/{id} [delete]
func rolesV3Delete(ctx *gin.Context) {
	errors.AbortRequest(ctx, fmt.Errorf("not implemented"))
}
