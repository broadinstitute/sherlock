package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func ciIdentifiersV3List(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	var filter CiIdentifierV3
	if err = ctx.MustBindWith(&filter, binding.Query); err != nil {
		return
	}
	modelFilter := filter.toModel()
	var results []models.CiIdentifier
	if err = db.Where(&modelFilter).Find(&results).Error; err != nil {
		ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, utils.Map(results, ciIdentifierFromModel))
}
