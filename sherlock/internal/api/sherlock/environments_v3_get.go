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
	"strings"
)

// environmentsV3Get godoc
//
//	@summary		Get an individual Environment
//	@description	Get an individual Environment.
//	@tags			Environments
//	@produce		json
//	@param			selector				path		string	true	"The selector of the Environment, which can be either a numeric ID, the name, or 'resource-prefix' + / + the unique resource prefix."
//	@success		200						{object}	EnvironmentV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/environments/v3/{selector} [get]
func environmentsV3Get(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	query, err := environmentModelFromSelector(canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var result models.Environment
	if err = db.Preload(clause.Associations).Where(&query).First(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, environmentFromModel(result))
}

func environmentModelFromSelector(selector string) (query models.Environment, err error) {
	if len(selector) == 0 {
		return models.Environment{}, fmt.Errorf("(%s) selector cannot be empty", errors.BadRequest)
	} else if utils.IsNumeric(selector) { // ID
		query.ID, err = utils.ParseUint(selector)
		return query, err
	} else if utils.IsAlphaNumericWithHyphens(selector) &&
		utils.IsStartingWithLetter(selector) &&
		utils.IsEndingWithAlphaNumeric(selector) { // Name
		query.Name = selector
		return query, nil
	} else if strings.Count(selector, "/") == 1 { // "resource-prefix" + unique resource prefix
		parts := strings.Split(selector, "/")

		// "resource-prefix"
		// The reason we have this at all is so that we can differentiate resource prefix selectors from name selectors.
		// In other words, a name can't have the slash that "resource-prefix/<blah>" has, so that's our hack to tell
		// incoming selectors apart
		selectorLabel := parts[0]
		if selectorLabel != "resource-prefix" {
			return models.Environment{}, fmt.Errorf("(%s) invalid environment selector %s, unique resource prefix selector needed to start with 'resource-prefix/' but was '%s/'", errors.BadRequest, selector, selectorLabel)
		}

		// unique resource prefix
		uniqueResourcePrefix := parts[1]
		if !(utils.IsLowerAlphaNumeric(uniqueResourcePrefix) &&
			utils.IsStartingWithLetter(uniqueResourcePrefix) &&
			utils.IsEndingWithAlphaNumeric(uniqueResourcePrefix) &&
			len(uniqueResourcePrefix) == 4) {
			return models.Environment{}, fmt.Errorf("(%s) invalid environment selector %s, unique resource prefix sub-selector %s was invalid", errors.BadRequest, selector, uniqueResourcePrefix)
		}
		query.UniqueResourcePrefix = uniqueResourcePrefix
		return query, nil
	} else {
		return models.Environment{}, fmt.Errorf("(%s) invalid environment selector '%s'", errors.BadRequest, selector)
	}
}
