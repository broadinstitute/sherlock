package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation"
	"github.com/creasty/defaults"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

// roleAssignmentsV3Create godoc
//
//	@summary		Create a RoleAssignment
//	@description	Create the RoleAssignment between a given Role and User.
//	@description	Non-super-admins may only mutate RoleAssignments for themselves, only for roles they can break-glass into, and only with an expiry no further than the role's default break-glass duration in the future.
//	@description	Propagation will be triggered after this operation.
//	@tags			RoleAssignments
//	@produce		json
//	@param			role-selector			path		string					true	"The selector of the Role, which can be either the numeric ID or the name"
//	@param			user-selector			path		string					true	"The selector of the User, which can be either a numeric ID, the email, 'google-id/{google subject ID}', 'github/{github username}', or 'github-id/{github numeric ID}'."
//	@param			role-assignment			body		RoleAssignmentV3Edit	true	"The initial fields to set for the new RoleAssignment"
//	@success		201						{object}	RoleAssignmentV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/role-assignments/v3/{role-selector}/{user-selector} [post]
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
	if err = defaults.Set(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error setting defaults: %w", err))
		return
	}

	toCreate, err := body.toModel()
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	roleQuery, err := roleModelFromSelector(canonicalizeSelector(ctx.Param("role-selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var role models.Role
	if err = db.Where(&roleQuery).First(&role).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	toCreate.RoleID = role.ID

	userQuery, err := userModelFromSelector(canonicalizeSelector(ctx.Param("user-selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var user models.User
	if err = db.Where(&userQuery).Preload("Suitability").First(&user).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	toCreate.UserID = user.ID

	if role.SuspendNonSuitableUsers != nil && *role.SuspendNonSuitableUsers {
		// If Role.SuspendNonSuitableUsers is true, we compute RoleAssignments.Suspended
		shouldBeSuspended := user.Suitability == nil || !*user.Suitability.Suitable
		if toCreate.Suspended != nil && *toCreate.Suspended != shouldBeSuspended {
			errors.AbortRequest(ctx, fmt.Errorf("(%s) request manually set suspended to %v, but for this role it's a computed field and is expected to be %v (please omit setting it or set it to %v)",
				errors.BadRequest, *toCreate.Suspended, shouldBeSuspended, shouldBeSuspended))
			return
		} else {
			toCreate.Suspended = &shouldBeSuspended
		}
	} else if toCreate.Suspended == nil {
		// If it's not a computed field but is empty, we default to false
		toCreate.Suspended = utils.PointerTo(false)
	}

	if err = db.Create(&toCreate).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	if err = db.Preload(clause.Associations).Where(&models.RoleAssignment{RoleID: toCreate.RoleID, UserID: toCreate.UserID}).First(&toCreate).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, roleAssignmentFromModel(toCreate))

	role_propagation.DoOnDemandPropagation(ctx, db, role.ID)
}
