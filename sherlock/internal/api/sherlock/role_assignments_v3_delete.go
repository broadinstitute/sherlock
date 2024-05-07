package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
)

// roleAssignmentsV3Delete godoc
//
//	@summary		Delete a RoleAssignment
//	@description	Delete the RoleAssignment between a given Role and User.
//	@tags			RoleAssignments
//	@produce		json
//	@param			role-id					path		uint	true	"The numeric ID of the role"
//	@param			user-selector			path		string	true	"The selector of the User, which can be either a numeric ID, the email, 'google-id/{google subject ID}', 'github/{github username}', or 'github-id/{github numeric ID}'."
//	@success		200						{object}	RoleAssignmentV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/role-assignments/v3/{role-id}/{user-selector} [delete]
func roleAssignmentsV3Delete(ctx *gin.Context) {
	errors.AbortRequest(ctx, fmt.Errorf("not implemented"))
}
