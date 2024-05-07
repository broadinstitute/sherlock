package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
)

// roleAssignmentsV3List godoc
//
//	@summary		List RoleAssignments matching a filter
//	@description	List RoleAssignments matching a filter. The correct way to list RoleAssignments for a particular Role or User is to get that Role or User specifically, not to use this endpoint.
//	@tags			RoleAssignments
//	@produce		json
//	@param			filter					query		RoleAssignmentV3	false	"Filter the returned RoleAssignments"
//	@param			limit					query		int					false	"Control how many RoleAssignments are returned (default 0, no limit)"
//	@param			offset					query		int					false	"Control the offset for the returned RoleAssignments (default 0)"
//	@success		200						{array}		RoleAssignmentV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/role-assignments/v3 [get]
func roleAssignmentsV3List(ctx *gin.Context) {
	errors.AbortRequest(ctx, fmt.Errorf("not implemented"))
}
