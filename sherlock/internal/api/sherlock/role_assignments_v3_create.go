package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

// roleAssignmentsV3Create godoc
//
//	@summary		Create a RoleAssignment
//	@description	Create the RoleAssignment between a given Role and User.
//	@description	Non-super-admins may only mutate RoleAssignments for themselves, only for roles they can break-glass into, and only with an expiry no further than the role's default break-glass duration in the future.
//	@tags			RoleAssignments
//	@produce		json
//	@param			role-id					path		uint					true	"The numeric ID of the role"
//	@param			user-selector			path		string					true	"The selector of the User, which can be either a numeric ID, the email, 'google-id/{google subject ID}', 'github/{github username}', or 'github-id/{github numeric ID}'."
//	@param			role-assignment			body		RoleAssignmentV3Edit	true	"The initial fields to set for the new RoleAssignment"
//	@success		201						{object}	RoleAssignmentV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/role-assignments/v3/{role-id}/{user-selector} [post]
func roleAssignmentsV3Create(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	var body RoleAssignmentV3Edit
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	toCreate := body.toModel()
	toCreate.RoleID, err = utils.ParseUint(ctx.Param("role-id"))
	if err != nil {
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
	toCreate.UserID = user.ID

	if err = db.Create(&toCreate).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	if err = db.Preload(clause.Associations).Where(&models.RoleAssignment{RoleID: toCreate.RoleID, UserID: toCreate.UserID}).First(&toCreate).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, roleAssignmentFromModel(toCreate))
}
