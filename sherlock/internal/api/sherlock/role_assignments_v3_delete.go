package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

// roleAssignmentsV3Delete godoc
//
//	@summary		Delete a RoleAssignment
//	@description	Delete the RoleAssignment between a given Role and User.
//	@description	Non-super-admins may only mutate RoleAssignments for themselves, only for roles they can break-glass into, and only with an expiry no further than the role's default break-glass duration in the future.
//	@description	Propagation will be triggered after this operation.
//	@tags			RoleAssignments
//	@produce		json
//	@param			role-selector			path		string	true	"The selector of the Role, which can be either the numeric ID or the name"
//	@param			user-selector			path		string	true	"The selector of the User, which can be either a numeric ID, the email, 'google-id/{google subject ID}', 'github/{github username}', or 'github-id/{github numeric ID}'."//	@success	200	{object}	RoleAssignmentV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/role-assignments/v3/{role-selector}/{user-selector} [delete]
func roleAssignmentsV3Delete(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	var toDelete models.RoleAssignment
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
	toDelete.RoleID = role.ID

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
	toDelete.UserID = user.ID

	var result models.RoleAssignment
	if err = db.Preload(clause.Associations).Where(&toDelete).First(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	if err = db.Omit(clause.Associations).Delete(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, roleAssignmentFromModel(result))

	go role_propagation.WaitToPropagate(ctx, db, role.ID)
}
