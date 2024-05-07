package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
)

// rolesV3Edit godoc
//
//	@summary		Edit a Role
//	@description	Edit an individual Role.
//	@tags			Roles
//	@produce		json
//	@param			id						path		uint		true	"The numeric ID of the role"
//	@param			role					body		RoleV3Edit	true	"The edits to make to the Role"
//	@success		200						{object}	RoleV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/roles/v3/{id} [patch]
func rolesV3Edit(ctx *gin.Context) {
	errors.AbortRequest(ctx, fmt.Errorf("not implemented"))
}
