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

// rolesV3Edit godoc
//
//	@summary		Edit a Role
//	@description	Edit an individual Role.
//	@description	Only super-admins may mutate Roles.
//	@tags			Roles
//	@produce		json
//	@param			id						path		uint		true	"The numeric ID of the role"
//	@param			role					body		RoleV3Edit	true	"The edits to make to the Role"
//	@success		200						{object}	RoleV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/roles/v3/{id} [patch]
func rolesV3Edit(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	id, err := utils.ParseUint(ctx.Param("id"))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	var body RoleV3Edit
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	edits := body.toModel()

	var toEdit models.Role
	if err = db.Scopes(models.ReadRoleScope).First(&toEdit, id).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	if err = db.Model(&toEdit).Omit(clause.Associations).Updates(&edits).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, roleFromModel(toEdit))
}
