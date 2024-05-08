package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

// rolesV3Delete godoc
//
//	@summary		Delete a Role
//	@description	Delete an individual Role.
//	@description	Only super-admins may mutate Roles.
//	@tags			Roles
//	@produce		json
//	@param			id						path		uint	true	"The numeric ID of the role"
//	@success		200						{object}	RoleV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/roles/v3/{id} [delete]
func rolesV3Delete(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	id, err := utils.ParseUint(ctx.Param("id"))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var result models.Role
	if err = db.Scopes(models.ReadRoleScope).First(&result, id).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	if err = db.Omit(clause.Associations).Delete(&models.Role{}, id).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, roleFromModel(result))
}
