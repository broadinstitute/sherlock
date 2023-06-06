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
	// DDO-2709 Note from Jack:
	// I wrote out some handlers that currently return a 403 in all cases. There's tests for that happening at the
	// controller level but I wanted to be sure. I've just commented them out here instead of actually deleting them
	// because maybe we'll want them in the future -- no reason to wire them up twice.

	//routerGroup.POST("/users", createUser(controller))
	routerGroup.GET("/users", listUser(controller))
	routerGroup.GET("/users/*selector", getUser(controller))
	routerGroup.PATCH("/users/*selector", editUser(controller))
	//routerGroup.PUT("/users/*selector", upsertUser(controller))
	//routerGroup.DELETE("/users/*selector", deleteUser(controller))
	routerGroup.GET("/selectors/users/*selector", listUserSelectors(controller))
	routerGroup.POST("/procedures/users/link-github", updateUserGithubAssociation(controller))
	routerGroup.GET("/procedures/users/me", getOwnUser(controller))
}

//// createUser godoc
////
////	@summary		Create a new User entry
////	@description	Create a new User entry. Note that some fields are immutable after creation; /edit lists mutable fields.
////	@tags			Users
////	@accept			json
////	@produce		json
////	@param			user					body		v2controllers.CreatableUser	true	"The User to create"
////	@success		200,201					{object}	v2controllers.User
////	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
////	@router			/api/v2/users [post]
//func createUser(controller *v2controllers.UserController) func(ctx *gin.Context) {
//	return handleCreate(&controller.ModelController)
//}

// listUser godoc
//
//	@summary		List User entries
//	@description	List existing User entries, ordered by most recently updated.
//	@tags			Users
//	@produce		json
//	@param			filter					query		v2controllers.User	false	"Optional filters to apply to the returned entries"
//	@param			limit					query		int					false	"An optional limit to the number of entries returned"
//	@success		200						{array}		v2controllers.User
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/users [get]
func listUser(controller *v2controllers.UserController) func(ctx *gin.Context) {
	return handleList(&controller.ModelController)
}

// getUser godoc
//
//	@summary		Get a User entry
//	@description	Get an existing User entry via one of its "selectors": email, numeric id, 'github/' + GitHub username, 'github-id/' + GitHub numeric id, or 'google-id/' + Google numeric id.
//	@tags			Users
//	@produce		json
//	@param			selector				path		string	true	"The User to get's selector: email, numeric id, 'github/' + GitHub username, 'github-id/' + GitHub numeric id, or 'google-id/' + Google numeric id"
//	@success		200						{object}	v2controllers.User
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/users/{selector} [get]
func getUser(controller *v2controllers.UserController) func(ctx *gin.Context) {
	return handleGet(&controller.ModelController)
}

// editUser godoc
//
//	@summary		Edit a User entry
//	@description	Edit an existing User entry via one of its "selectors": email, numeric id, 'github/' + GitHub username, 'github-id/' + GitHub numeric id, or 'google-id/' + Google numeric id. Note that only mutable fields are available here, immutable fields can only be set using /create.
//	@tags			Users
//	@accept			json
//	@produce		json
//	@param			selector				path		string						true	"The User to edit's selector: email, numeric id, 'github/' + GitHub username, 'github-id/' + GitHub numeric id, or 'google-id/' + Google numeric id"
//	@param			user					body		v2controllers.EditableUser	true	"The edits to make to the User"
//	@success		200						{object}	v2controllers.User
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/users/{selector} [patch]
func editUser(controller *v2controllers.UserController) func(ctx *gin.Context) {
	return handleEdit(&controller.ModelController)
}

//// upsertUser godoc
////
////	@summary		Create or edit a User entry
////	@description	Create or edit a User entry. Attempts to edit and will attempt to create upon an error.
////	@description	If an edit was made or the creation process de-duplicates, this method will return normally with a 200.
////	@tags			Users
////	@accept			json
////	@produce		json
////	@param			selector				path		string							true	"The User to upsert's selector: email, numeric id, 'github/' + GitHub username, 'github-id/' + GitHub numeric id, or 'google-id/' + Google numeric id"
////	@param			user					body		v2controllers.CreatableUser	true	"The User to upsert"
////	@success		200,201					{object}	v2controllers.User
////	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
////	@router			/api/v2/users/{selector} [put]
//func upsertUser(controller *v2controllers.UserController) func(ctx *gin.Context) {
//	return handleUpsert(&controller.ModelController)
//}

//// deleteUser godoc
////
////	@summary		Delete a User entry
////	@description	Delete an existing User entry via one of its "selectors": email, numeric id, 'github/' + GitHub username, 'github-id/' + GitHub numeric id, or 'google-id/' + Google numeric id.
////	@tags			Users
////	@produce		json
////	@param			selector				path		string	true	"The User to delete's selector: email, numeric id, 'github/' + GitHub username, 'github-id/' + GitHub numeric id, or 'google-id/' + Google numeric id"
////	@success		200						{object}	v2controllers.User
////	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
////	@router			/api/v2/users/{selector} [delete]
//func deleteUser(controller *v2controllers.UserController) func(ctx *gin.Context) {
//	return handleDelete(&controller.ModelController)
//}

// listUserSelectors godoc
//
//	@summary		List User selectors
//	@description	Validate a given User selector and provide any other selectors that would match the same User.
//	@tags			Users
//	@produce		json
//	@param			selector				path		string	true	"The selector of the User to list other selectors for"
//	@success		200						{array}		string
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/selectors/users/{selector} [get]
func listUserSelectors(controller *v2controllers.UserController) func(ctx *gin.Context) {
	return handleSelectorList(&controller.ModelController)
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

// getOwnUser godoc
//
//	@summary		Get your own User entry
//	@description	Get your own User entry
//	@tags			Users
//	@produce		json
//	@success		200						{object}	v2controllers.User
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/procedures/users/me [get]
func getOwnUser(controller *v2controllers.UserController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user, err := auth.ExtractUserFromContext(ctx)
		if err != nil {
			ctx.JSON(errors.ErrorToApiResponse(err))
			return
		}
		result, err := controller.Get(user.Email)
		if err != nil {
			ctx.JSON(errors.ErrorToApiResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, result)
	}
}
