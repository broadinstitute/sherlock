package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// rolesV3Get godoc~
//
//	@summary		Get a Role
//	@description	Get an individual Role and the Users assigned to it.
//	@tags			Roles
//	@produce		json
//	@param			selector				path		string	true	"The selector of the Role, which can be either the numeric ID or the name"
//	@success		200						{object}	RoleV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/roles/v3/{selector} [get]
func rolesV3Get(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	query, err := roleModelFromSelector(canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var result models.Role
	if err = db.Scopes(models.ReadRoleScope).Where(&query).First(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, roleFromModel(result))
}

func roleModelFromSelector(selector string) (query models.Role, err error) {
	if len(selector) == 0 {
		return models.Role{}, fmt.Errorf("(%s) selector cannot be empty", errors.BadRequest)
	} else if utils.IsNumeric(selector) {
		query.ID, err = utils.ParseUint(selector)
		return query, err
	} else if utils.IsAlphaNumericWithHyphens(selector) {
		query.Name = &selector
		return query, nil
	} else {
		return models.Role{}, fmt.Errorf("(%s) role selector must be a numeric ID or a name; '%s' invalid", errors.BadRequest, selector)
	}
}
