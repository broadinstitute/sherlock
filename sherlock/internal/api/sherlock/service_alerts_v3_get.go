package sherlock

import (
	"fmt"
	"net/http"

	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

// rolesV3Get godoc~
//
//	@summary		Get a Service Alert
//	@description	Get an individual Service Alert and it's metadata.
//	@tags			ServiceAlert
//	@produce		json
//	@param			selector				path		string	true	"The selector of the ServiceAlert, which is the guid for a given alert"
//	@success		200						{object}	ServiceAlertV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/service-alerts/v3/{selector} [get]
func serviceAlertV3Get(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	query, err := serviceAlertFromSelector(canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var result models.ServiceAlert
	if err = db.Preload(clause.Associations).Where(&query).First(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, ServiceAlertFromModel(result))
}

func serviceAlertFromSelector(selector string) (query models.ServiceAlert, err error) {
	if len(selector) == 0 {
		return models.ServiceAlert{}, fmt.Errorf("(%s) selector cannot be empty", errors.BadRequest)
	} else if utils.IsNumeric(selector) {
		query.ID, err = utils.ParseUint(selector)
		return query, err
	} else if utils.IsAlphaNumericWithHyphens(selector) {
		return models.ServiceAlert{}, fmt.Errorf("(%s) selector must be of type uint", errors.BadRequest)
	} else {
		return models.ServiceAlert{}, fmt.Errorf("(%s) role selector must be a numeric ID or a name; '%s' invalid", errors.BadRequest, selector)
	}
}
