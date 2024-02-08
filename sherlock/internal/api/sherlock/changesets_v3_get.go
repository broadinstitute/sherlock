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

// changesetsV3Get godoc
//
//	@summary		Get an individual Changeset
//	@description	Get an individual Changeset.
//	@tags			Changesets
//	@produce		json
//	@param			id						path		int	true	"The numeric ID of the changeset"
//	@success		200						{object}	ChangesetV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/changesets/v3/{id} [get]
func changesetsV3Get(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	id, err := utils.ParseUint(canonicalizeSelector(ctx.Param("id")))
	if err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) failed to parse '%s' to an ID: %v", errors.BadRequest, canonicalizeSelector(ctx.Param("id")), err))
		return
	}
	var result models.Changeset
	if err = db.Scopes(models.ReadChangesetScope).First(&result, id).Error; err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) failed to get changeset with ID %d: %v", errors.NotFound, id, err))
		return
	}
	ctx.JSON(http.StatusOK, changesetFromModel(result))
}
