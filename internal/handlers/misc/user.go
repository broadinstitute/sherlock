package misc

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserResponse struct {
	Email    string `json:"email"`
	Suitable string `json:"suitable"`
}

// UserHandler godoc
// @summary      Get information about the calling user
// @description  Get Sherlock's understanding of the calling user based on IAP and the Firecloud.org Google Workspace.
// @tags         Misc
// @produce      json
// @success      200      {object}  misc.UserResponse
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
	ctx.JSON(http.StatusOK, UserResponse{
		Email:    user.AuthenticatedEmail,
		Suitable: user.DescribeSuitability(),
	})
}
