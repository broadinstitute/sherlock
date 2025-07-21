package sherlock

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
)

// usersV3Get godoc
//
//	@summary		Get an individual User
//	@description	Get an individual User. As a special case, "me" or "self" can be passed as the selector to get the current user.
//	@tags			Users
//	@produce		json
//	@param			selector				path		string	true	"The selector of the User, which can be either a numeric ID, the email, 'google-id/{google subject ID}', 'github/{github username}', or 'github-id/{github numeric ID}'. As a special case, 'me' or 'self' can be passed to get the calling user."
//	@success		200						{object}	UserV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/users/v3/{selector} [get]
func usersV3Get(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	selector := canonicalizeSelector(ctx.Param("selector"))

	if selector == "me" || selector == "self" {
		user, err := authentication.MustUseUser(ctx)
		if err != nil {
			return
		}
		ctx.JSON(http.StatusOK, userFromModel(*user))
	} else {
		query, err := userModelFromSelector(selector)
		if err != nil {
			errors.AbortRequest(ctx, err)
			return
		}
		var result models.User
		if err = db.Where(&query).Scopes(models.ReadUserScope).First(&result).Error; err != nil {
			errors.AbortRequest(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, userFromModel(result))
	}
}

func userModelFromSelector(selector string) (query models.User, err error) {
	if len(selector) == 0 {
		return models.User{}, fmt.Errorf("(%s) selector cannot be empty", errors.BadRequest)
	} else if utils.IsNumeric(selector) {
		query.ID, err = utils.ParseUint(selector)
		return query, err
	} else if strings.Contains(selector, "@") {
		query.Email = selector
		return query, nil
	} else if strings.HasPrefix(selector, "google-id/") {
		query.GoogleID = strings.TrimPrefix(selector, "google-id/")
		if len(query.GoogleID) == 0 {
			return models.User{}, fmt.Errorf("(%s) google ID selector can't be empty", errors.BadRequest)
		}
		return query, nil
	} else if strings.HasPrefix(selector, "github/") {
		githubUsername := strings.TrimPrefix(selector, "github/")
		if len(githubUsername) == 0 {
			return models.User{}, fmt.Errorf("(%s) github username selector can't be empty", errors.BadRequest)
		}
		query.GithubUsername = &githubUsername
		return query, nil
	} else if strings.HasPrefix(selector, "github-id/") {
		githubID := strings.TrimPrefix(selector, "github-id/")
		if len(githubID) == 0 {
			return models.User{}, fmt.Errorf("(%s) github ID selector can't be empty", errors.BadRequest)
		}
		query.GithubID = &githubID
		return query, nil
	} else {
		return models.User{}, fmt.Errorf("(%s) invalid user selector '%s'", errors.BadRequest, selector)
	}
}
