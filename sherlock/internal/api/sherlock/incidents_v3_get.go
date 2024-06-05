package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

// incidentsV3Get godoc
//
//	@summary		Get an individual Incident
//	@description	Get an individual Incident.
//	@tags			Incidents
//	@produce		json
//	@param			selector				path		string	true	"The ID of the Incident"
//	@success		200						{object}	IncidentV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/incidents/v3/{selector} [get]
func incidentsV3Get(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	query, err := incidentModelFromSelector(canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var result models.Incident
	if err = db.Preload(clause.Associations).Where(&query).First(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, incidentFromModel(result))
}

func incidentModelFromSelector(selector string) (query models.Incident, err error) {
	if len(selector) == 0 {
		return models.Incident{}, fmt.Errorf("(%s) incident selector cannot be empty", errors.BadRequest)
	}
	if utils.IsNumeric(selector) { // ID
		query.ID, err = utils.ParseUint(selector)
		return query, err
	}
	return models.Incident{}, fmt.Errorf("(%s) invalid incident selector '%s'", errors.BadRequest, selector)
}
