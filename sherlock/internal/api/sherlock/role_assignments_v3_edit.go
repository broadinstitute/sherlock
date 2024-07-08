package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

// roleAssignmentsV3Edit godoc
//
//	@summary		Edit a RoleAssignment
//	@description	Edit the RoleAssignment between a given Role and User.
//	@description	Non-super-admins may only mutate RoleAssignments for themselves, only for roles they can break-glass into, and only with an expiry no further than the role's default break-glass duration in the future.
//	@description	Propagation will be triggered after this operation.
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
	edits, err := body.toModel()
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

	var toEdit models.RoleAssignment
	if err = db.Preload(clause.Associations).Where(&models.RoleAssignment{RoleID: role.ID, UserID: user.ID}).First(&toEdit).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	if edits.Suspended != nil && role.SuspendNonSuitableUsers != nil && *role.SuspendNonSuitableUsers {
		// If Role.SuspendNonSuitableUsers is true and RoleAssignment.Suspended is being set, we make sure it matches the computation
		// We do this just so we don't unnecessarily error if someone accidentally sends the field as a no-op in the request
		shouldBeSuspended := user.Suitability == nil || !*user.Suitability.Suitable
		if *edits.Suspended != shouldBeSuspended {
			errors.AbortRequest(ctx, fmt.Errorf("(%s) request manually set suspended to %v, but for this role it's a computed field and is expected to be %v (please omit setting it or set it to %v)",
				errors.BadRequest, *edits.Suspended, shouldBeSuspended, shouldBeSuspended))
			return
		}
		// We don't bother actually setting RoleAssignment.Suspended here since we don't want to attribute an edit to the user that they didn't make
		// Our cron that updates suspensions will make the update very soon and correctly attribute it to Sherlock itself
	}

	if err = db.Model(&toEdit).Omit(clause.Associations).Updates(&edits).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, roleAssignmentFromModel(toEdit))

	role_propagation.DoOnDemandPropagation(ctx, db, role.ID)
}
