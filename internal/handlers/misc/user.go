package misc

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserHandler godoc
// @summary      Get information about the calling user
// @description  Get Sherlock's understanding of the calling user based on IAP and Google Groups.
// @tags         Misc
// @produce      json
// @success      200      {object}  auth.User
// @failure      407,500  {object}  errors.ErrorResponse
// @router       /user [get]
func UserHandler(ctx *gin.Context) {
	userValue, exists := ctx.Get(auth.ContextUserKey)
	if !exists {
		ctx.JSON(errors.ErrorToApiResponse(fmt.Errorf("(%s) authentication middleware not present", errors.InternalServerError)))
		return
	}
	user, ok := userValue.(auth.User)
	if !ok {
		ctx.JSON(errors.ErrorToApiResponse(fmt.Errorf("(%s) authentication middleware misconfigured: suitability represented as %T", errors.InternalServerError, userValue)))
	}
	ctx.JSON(http.StatusOK, user)
}
