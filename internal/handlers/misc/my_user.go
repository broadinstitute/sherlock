package misc

import (
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/auth/auth_models"
	"github.com/broadinstitute/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MyUserResponse struct {
	Email       string            `json:"email"`
	Suitability string            `json:"suitability"`
	RawInfo     *auth_models.User `json:"rawInfo"`
}

// MyUserHandler godoc
//
//	@summary		Get information about the calling user
//	@description	Get Sherlock's understanding of the calling user based on IAP and the Firecloud.org Google Workspace organization.
//	@tags			Misc
//	@produce		json
//	@success		200		{object}	misc.MyUserResponse
//	@failure		407,500	{object}	errors.ErrorResponse
//	@router			/my-user [get]
func MyUserHandler(ctx *gin.Context) {
	user, err := auth.ExtractUserFromContext(ctx)
	if err != nil {
		ctx.JSON(errors.ErrorToApiResponse(err))
		return
	}

	var suitabilityDescription string
	if nonsuitableErr := user.SuitableOrError(); nonsuitableErr != nil {
		suitabilityDescription = nonsuitableErr.Error()
	} else {
		suitabilityDescription = "user is suitable"
	}

	ctx.JSON(http.StatusOK, MyUserResponse{
		Email:       user.Email,
		Suitability: suitabilityDescription,
		RawInfo:     user,
	})
}
