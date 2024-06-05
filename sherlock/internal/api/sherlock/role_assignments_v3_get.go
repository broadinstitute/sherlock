package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

// roleAssignmentsV3Get godoc
//
//	@summary		Get a RoleAssignment
//	@description	Get the RoleAssignment between a given Role and User.
//	@tags			RoleAssignments
//	@produce		json
//	@param			role-selector			path		string	true	"The selector of the Role, which can be either the numeric ID or the name"
//	@param			user-selector			path		string	true	"The selector of the User, which can be either a numeric ID, the email, 'google-id/{google subject ID}', 'github/{github username}', or 'github-id/{github numeric ID}'."
//	@success		200						{object}	RoleAssignmentV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/role-assignments/v3/{role-selector}/{user-selector} [get]
func roleAssignmentsV3Get(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

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

	var result models.RoleAssignment
	if err = db.Preload(clause.Associations).Where(&models.RoleAssignment{RoleID: role.ID, UserID: user.ID}).First(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, roleAssignmentFromModel(result))
}
