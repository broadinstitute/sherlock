package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// rolesV3Create godoc
//
//	@summary		Create a Role
//	@description	Create an individual Role with no one assigned to it.
//	@description	Only super-admins may mutate Roles.
//	@tags			Roles
//	@produce		json
//	@param			role					body		RoleV3Edit	true	"The initial fields the Role should have set"
//	@success		201						{object}	RoleV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/roles/v3 [post]
func rolesV3Create(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	var body RoleV3Edit
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	toCreate := body.toModel()
	if err = db.Create(&toCreate).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	if err = db.Scopes(models.ReadRoleScope).First(&toCreate, toCreate.ID).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, roleFromModel(toCreate))
}