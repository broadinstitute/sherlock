package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

// roleAssignmentsV3Edit godoc
//
//	@summary		Edit a RoleAssignment
//	@description	Edit the RoleAssignment between a given Role and User.
//	@description	Non-super-admins may only mutate RoleAssignments for themselves, only for roles they can break-glass into, and only with an expiry no further than the role's default break-glass duration in the future.
//	@tags			RoleAssignments
//	@produce		json
//	@param			role-selector			path		string					true	"The selector of the Role, which can be either the numeric ID or the name"
//	@param			user-selector			path		string					true	"The selector of the User, which can be either a numeric ID, the email, 'google-id/{google subject ID}', 'github/{github username}', or 'github-id/{github numeric ID}'."//	@param	role-assignment	body	RoleAssignmentV3Edit	true	"The edits to make to the RoleAssignment"
//	@param			role-assignment			body		RoleAssignmentV3Edit	true	"The edits to make to the RoleAssignment"
//	@success		200						{object}	RoleAssignmentV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/role-assignments/v3/{role-selector}/{user-selector} [patch]
func roleAssignmentsV3Edit(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	var body RoleAssignmentV3Edit
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}
	edits := body.toModel()

	roleQuery, err := roleModelFromSelector(canonicalizeSelector(ctx.Param("role-selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var role models.Role
	if err = db.Where(&roleQuery).Select("id").First(&role).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	userQuery, err := userModelFromSelector(canonicalizeSelector(ctx.Param("user-selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var user models.User
	if err = db.Where(&userQuery).Select("id").First(&user).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	var toEdit models.RoleAssignment
	if err = db.Preload(clause.Associations).Where(&models.RoleAssignment{RoleID: role.ID, UserID: user.ID}).First(&toEdit).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	if err = db.Model(&toEdit).Omit(clause.Associations).Updates(&edits).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, roleAssignmentFromModel(toEdit))
}
