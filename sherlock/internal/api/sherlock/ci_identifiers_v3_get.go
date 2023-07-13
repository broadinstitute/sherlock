package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

func ciIdentifiersV3Get(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	query, err := ciIdentifierModelFromSelector(db, canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
		return
	}
	var result models.CiIdentifier
	if err = db.Where(&query).First(&result).Error; err != nil {
		ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, ciIdentifierFromModel(result))
}

func ciIdentifierModelFromSelector(db *gorm.DB, selector string) (query models.CiIdentifier, err error) {
	if len(selector) == 0 {
		return models.CiIdentifier{}, fmt.Errorf("(%s) selector cannot be empty", errors.BadRequest)
	} else if utils.IsNumeric(selector) {
		query.ID, err = utils.ParseUint(selector)
		return query, err
	} else if strings.Count(selector, "/") > 0 {
		// resource type + type's selector ...
		parts := strings.Split(selector, "/")
		query.ResourceType = parts[0]
		resourceSelector := strings.Join(parts[1:], "/")
		switch query.ResourceType {
		// We're calling out to the "old code" here to handle resolving the selectors.
		// Not horribly inelegant but we'll want to call refactored mechanisms here when they're available.
		case "chart":
			query.ResourceID, err = v2models.InternalChartStore.ResolveSelector(db, resourceSelector)
		case "chart-version":
			query.ResourceID, err = v2models.InternalChartVersionStore.ResolveSelector(db, resourceSelector)
		case "app-version":
			query.ResourceID, err = v2models.InternalAppVersionStore.ResolveSelector(db, resourceSelector)
		case "cluster":
			query.ResourceID, err = v2models.InternalClusterStore.ResolveSelector(db, resourceSelector)
		case "environment":
			query.ResourceID, err = v2models.InternalEnvironmentStore.ResolveSelector(db, resourceSelector)
		case "chart-release":
			query.ResourceID, err = v2models.InternalChartReleaseStore.ResolveSelector(db, resourceSelector)
		case "changeset":
			query.ResourceID, err = v2models.InternalChangesetStore.ResolveSelector(db, resourceSelector)
		default:
			err = fmt.Errorf("(%s) invalid CI identifier selector '%s', resource type sub-selector '%s' wasn't recognized", errors.BadRequest, selector, query.ResourceType)
		}
		return query, err
	} else {
		return models.CiIdentifier{}, fmt.Errorf("(%s) invalid CI identifier selector '%s'", errors.BadRequest, selector)
	}
}