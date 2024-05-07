package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
)

// rolesV3Create godoc
//
//	@summary		Create a Role
//	@description	Create an individual Role with no one assigned to it.
//	@tags			Roles
//	@produce		json
//	@param			role					body		RoleV3Edit	true	"The initial fields the Role should have set"
//	@success		201						{object}	RoleV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/roles/v3 [post]
func rolesV3Create(ctx *gin.Context) {
	errors.AbortRequest(ctx, fmt.Errorf("not implemented"))
}
