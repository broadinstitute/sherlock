package misc

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MyUserResponse struct {
	Email       string     `json:"email"`
	Suitability string     `json:"suitability"`
	RawInfo     *auth.User `json:"rawInfo"`
}

// MyUserHandler godoc
// @summary      Get information about the calling user
// @description  Get Sherlock's understanding of the calling user based on IAP and the Firecloud.org Google Workspace organization.
// @tags         Misc
// @produce      json
// @success      200      {object}  misc.MyUserResponse
// @failure      407,500  {object}  errors.ErrorResponse
// @router       /my-user [get]
func MyUserHandler(ctx *gin.Context) {
	userValue, exists := ctx.Get(auth.ContextUserKey)
	if !exists {
		ctx.JSON(errors.ErrorToApiResponse(fmt.Errorf("(%s) authentication middleware not present", errors.InternalServerError)))
		return
	}
	user, ok := userValue.(*auth.User)
	if !ok {
		ctx.JSON(errors.ErrorToApiResponse(fmt.Errorf("(%s) authentication middleware misconfigured: suitability represented as %T", errors.InternalServerError, userValue)))
	}

	var suitabilityDescription string
	if nonsuitableErr := user.SuitableOrError(); nonsuitableErr != nil {
		suitabilityDescription = nonsuitableErr.Error()
	} else {
		suitabilityDescription = "user is suitable"
	}

	ctx.JSON(http.StatusOK, MyUserResponse{
		Email:       user.AuthenticatedEmail,
		Suitability: suitabilityDescription,
		RawInfo:     user,
	})
}
