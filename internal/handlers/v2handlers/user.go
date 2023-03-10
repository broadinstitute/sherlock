package v2handlers

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/controllers/v2controllers"
	"github.com/broadinstitute/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterUserHandlers(routerGroup *gin.RouterGroup, controller *v2controllers.UserController) {
	routerGroup.POST("/procedures/users/link-github", updateUserGithubAssociation(controller))
}

// updateUserGithubAssociation godoc
//
//	@summary		Update the User's GitHub account link
//	@description	Update the authenticated User's associated personal GitHub account
//	@tags			Users
//	@accept			json
//	@produce		json
//	@param			github-access-payload-request	body		v2controllers.GithubAccessPayload	true	"Access to the GitHub account to link"
//	@success		200,202							{object}	v2controllers.User
//
//	@failure		400,403,404,407,409,500			{object}	errors.ErrorResponse
//
//	@router			/api/v2/procedures/users/link-github [post]
func updateUserGithubAssociation(controller *v2controllers.UserController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user, err := auth.ExtractUserFromContext(ctx)
		if err != nil {
			ctx.JSON(errors.ErrorToApiResponse(err))
			return
		}
		var request v2controllers.GithubAccessPayload
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(errors.ErrorToApiResponse(fmt.Errorf("(%s) JSON error parsing to %T: %v", errors.BadRequest, request, err)))
			return
		}
		result, updated, err := controller.UpdateUserGithubAssociation(request, user)
		if err != nil {
			ctx.JSON(errors.ErrorToApiResponse(err))
			return
		}
		if updated {
			ctx.JSON(http.StatusAccepted, result)
		} else {
			ctx.JSON(http.StatusOK, result)
		}
	}
}
