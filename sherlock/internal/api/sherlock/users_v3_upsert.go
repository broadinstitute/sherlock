package sherlock

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v50/github"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
	"net/http"
	"strconv"
	"time"
)

type UserV3Upsert struct {
	userDirectlyEditableFields
	// An access token for the GitHub account to associate with the calling user. The access token isn't stored.
	// The design here ensures that an association is only built when someone controls both accounts (Google via
	// IAP and GitHub via this access token).
	GithubAccessToken *string `json:"githubAccessToken"`
}

// usersV3Upsert godoc
//
//	@summary		Update the calling User's information
//	@description	Update the calling User's information. As with all authenticated Sherlock endpoints,
//	@description	newly-observed callers will have a User record added, meaning that this endpoint
//	@description	behaves like an upsert.
//	@tags			Users
//	@accept			json
//	@produce		json
//	@param			user					body		UserV3Upsert	false	"The User data to update"
//	@success		200,201					{object}	UserV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/users/v3 [put]
func usersV3Upsert(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	var body UserV3Upsert
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %v", errors.BadRequest, err))
		return
	}

	var githubUser *github.User
	if body.GithubAccessToken != nil {
		tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: *body.GithubAccessToken})
		timeoutContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		githubClient := github.NewClient(oauth2.NewClient(timeoutContext, tokenSource))
		githubUser, _, err = githubClient.Users.Get(timeoutContext, "")
		cancel()
		if err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("(%s) GitHub API responded to getting access token's identity with an error: %v", errors.BadRequest, err))
			return
		} else if githubUser == nil {
			errors.AbortRequest(ctx, fmt.Errorf("(%s) GitHub API responded to getting access token's identity with an empty response but no error", errors.BadRequest))
			return
		} else if githubUser.ID == nil || githubUser.Login == nil {
			errors.AbortRequest(ctx, fmt.Errorf("(%s) GitHub API responded to getting access token's identity without providing an ID and/or username", errors.InternalServerError))
			return
		}
	}

	callingUser, err := authentication.MustUseUser(ctx)
	if err != nil {
		return
	}

	callingUser, hasUpdates := processUserEdits(callingUser, githubUser, body.userDirectlyEditableFields)
	if hasUpdates {
		if err = db.Save(callingUser).Error; err != nil {
			errors.AbortRequest(ctx, err)
			return
		}
		ctx.JSON(http.StatusCreated, userFromModel(*callingUser))
	} else {
		ctx.JSON(http.StatusOK, userFromModel(*callingUser))
	}
}

func processUserEdits(callingUser *models.User, githubUser *github.User, directEdits userDirectlyEditableFields) (*models.User, bool) {
	hasUpdates := false

	// If direct edits set NameInferredFromGithub and callingUser lacks it or has a different value, set it
	if directEdits.NameInferredFromGithub != nil &&
		(callingUser.NameInferredFromGithub == nil || *directEdits.NameInferredFromGithub != *callingUser.NameInferredFromGithub) {
		callingUser.NameInferredFromGithub = directEdits.NameInferredFromGithub
		hasUpdates = true
	}

	// If callingUser lacks NameInferredFromGithub or has it false, and direct edits set Name, and callingUser lacks it or has a different value, set it
	if (callingUser.NameInferredFromGithub == nil || !*callingUser.NameInferredFromGithub) &&
		directEdits.Name != nil &&
		(callingUser.Name == nil || *directEdits.Name != *callingUser.Name) {
		callingUser.Name = directEdits.Name
		hasUpdates = true
	}

	// If we have a githubUser:
	if githubUser != nil {

		githubUserIdString := strconv.FormatInt(*githubUser.ID, 10)

		// If callingUser lacks GitHub info, set it and log
		if callingUser.GithubID == nil || callingUser.GithubUsername == nil {
			callingUser.GithubID = &githubUserIdString
			callingUser.GithubUsername = githubUser.Login
			hasUpdates = true
			log.Info().Msgf("GH   | first-time github account linking from %s to github account %s (ID: %s)", callingUser.Email, *githubUser.Login, githubUserIdString)
		}

		// If callingUser has different IDed GitHub info stored, update it and log
		if *callingUser.GithubID != githubUserIdString {
			callingUser.GithubID = &githubUserIdString
			callingUser.GithubUsername = githubUser.Login
			hasUpdates = true
			log.Info().Msgf("GH   | github account linking from %s changing from github account %s (ID: %s) to %s (ID: %s)", callingUser.Email, *callingUser.GithubUsername, *callingUser.GithubID, *githubUser.Login, githubUserIdString)
		}

		// If callingUser has same IDed GitHub info stored but new name, update it and log
		if *callingUser.GithubID == githubUserIdString && *callingUser.GithubUsername != *githubUser.Login {
			callingUser.GithubUsername = githubUser.Login
			hasUpdates = true
			log.Info().Msgf("GH   | github account linking from %s had new username for same github account (ID: %s), %s to %s", callingUser.Email, githubUserIdString, *callingUser.GithubUsername, *githubUser.Login)
		}

		// If callingUser lacks NameInferredFromGithub, default it to whether or not the callingUser already has a name
		if callingUser.NameInferredFromGithub == nil {
			shouldInferFromGithub := callingUser.Name == nil
			callingUser.NameInferredFromGithub = &shouldInferFromGithub
			hasUpdates = true
		}

		// If githubUser has Name, and callingUser has true NameInferredFromGithub, and callingUser lacks Name or has a different value, set it
		if githubUser.Name != nil &&
			callingUser.NameInferredFromGithub != nil && *callingUser.NameInferredFromGithub &&
			(callingUser.Name == nil || *githubUser.Name != *callingUser.Name) {
			callingUser.Name = githubUser.Name
			hasUpdates = true
		}
	}
	return callingUser, hasUpdates
}
