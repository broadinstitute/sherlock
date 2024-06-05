package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

// incidentsV3Delete godoc
//
//	@summary		Delete an individual Incident
//	@description	Delete an individual Incident by its ID.
//	@tags			Incidents
//	@produce		json
//	@param			selector				path		string	true	"The ID of the Incident"
//	@success		200						{object}	IncidentV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/incidents/v3/{selector} [delete]
func incidentsV3Delete(ctx *gin.Context) {
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
	if err = db.Omit(clause.Associations).Delete(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, incidentFromModel(result))
}
