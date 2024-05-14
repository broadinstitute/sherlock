package sherlock

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/github"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm/clause"
	"net/http"
	"strings"
	"sync"
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
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	callingUser, err := authentication.MustUseUser(ctx)
	if err != nil {
		return
	}

	// Copying isn't *strictly* necessary, but it's defensive programming, since the
	// callingUser is also used for authorization and we wouldn't want to accidentally
	// mutate it in a way that impacts that behavior for this request.
	copiedUser := &models.User{}
	if err = copier.CopyWithOption(copiedUser, callingUser, copier.Option{DeepCopy: true}); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) error copying user struct: %w", errors.InternalServerError, err))
		return
	}

	copiedUser, hasUpdates := processUserEdits(copiedUser, body.userDirectlyEditableFields, body.GithubAccessToken)
	var statusCode int
	if hasUpdates {
		if err = db.Omit(clause.Associations).Save(copiedUser).Error; err != nil {
			errors.AbortRequest(ctx, err)
			return
		}
		statusCode = http.StatusCreated
	} else {
		statusCode = http.StatusOK
	}

	var result models.User
	if err = db.Scopes(models.ReadUserScope).Take(&result, callingUser.ID).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(statusCode, userFromModel(result))
}

func processUserEdits(callingUser *models.User, directEdits userDirectlyEditableFields, userGithubToken *string) (resultingUser *models.User, hasUpdates bool) {
	githubString := "github"
	sherlockString := "sherlock"
	slackString := "slack"
	var githubID, githubUsername, githubName, slackID, slackUsername, slackName string
	var wg sync.WaitGroup
	timeoutContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if config.Config.Bool("github.behaviors.collectUserInfo.enable") && userGithubToken != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var err error
			githubID, githubUsername, githubName, err = github.GetCurrentUser(timeoutContext, *userGithubToken)
			if err != nil {
				log.Warn().Err(err).Msgf("error using %s's github auth to collect user info", callingUser.Email)
			}
		}()
	}

	if config.Config.Bool("slack.behaviors.collectUserInfo.enable") {
		emailDomains := config.Config.Strings("slack.behaviors.collectUserInfo.restrictToEmailDomains")
		callingUserEmailParts := strings.Split(callingUser.Email, "@")
		if len(emailDomains) == 0 || utils.Contains(emailDomains, callingUserEmailParts[len(callingUserEmailParts)-1]) {
			wg.Add(1)
			go func() {
				defer wg.Done()
				var err error
				slackID, slackUsername, slackName, err = slack.GetUser(timeoutContext, callingUser.Email)
				if err != nil {
					log.Warn().Err(err).Msgf("error getting %s's slack info", callingUser.Email)
				}
			}()
		}
	}

	wg.Wait()

	// If nameFrom wasn't set but nameInferredFromGithub was, be compatible and convert it to nameFrom
	if directEdits.NameFrom == nil && directEdits.NameInferredFromGithub != nil {
		if *directEdits.NameInferredFromGithub {
			directEdits.NameFrom = &githubString
		} else {
			directEdits.NameFrom = &sherlockString
		}
	}

	// Set nameFrom if it was provided and is different from what we already have
	if directEdits.NameFrom != nil && (callingUser.NameFrom == nil || *directEdits.NameFrom != *callingUser.NameFrom) {
		callingUser.NameFrom = directEdits.NameFrom
		hasUpdates = true
	}

	// Set name directly if we should, it was provided, and is different from what we already have
	if (callingUser.NameFrom == nil || *callingUser.NameFrom == "sherlock") && directEdits.Name != nil && (callingUser.Name == nil || *directEdits.Name != *callingUser.Name) {
		callingUser.Name = directEdits.Name
		callingUser.NameFrom = &sherlockString
		hasUpdates = true
	}

	// Set Slack ID if we got it and it is different from what we already have
	if slackID != "" && (callingUser.SlackID == nil || slackID != *callingUser.SlackID) {
		callingUser.SlackID = &slackID
		hasUpdates = true
	}

	// Set Slack username if we got it and it is different from what we already have
	if slackUsername != "" && (callingUser.SlackUsername == nil || slackUsername != *callingUser.SlackUsername) {
		callingUser.SlackUsername = &slackUsername
		hasUpdates = true
	}

	// Set name from Slack if we should, it was provided, and is different from what we already have
	if (callingUser.NameFrom == nil || *callingUser.NameFrom == "slack") && slackName != "" && (callingUser.Name == nil || slackName != *callingUser.Name) {
		callingUser.Name = &slackName
		callingUser.NameFrom = &slackString
		hasUpdates = true
	}

	// Set Github ID if we got it and it is different from what we already have
	if githubID != "" && (callingUser.GithubID == nil || githubID != *callingUser.GithubID) {
		callingUser.GithubID = &githubID
		hasUpdates = true
	}

	// Set Github username if we got it and it is different from what we already have
	if githubUsername != "" && (callingUser.GithubUsername == nil || githubUsername != *callingUser.GithubUsername) {
		callingUser.GithubUsername = &githubUsername
		hasUpdates = true
	}

	// Set name from Github if we should, it was provided, and is different from what we already have
	if (callingUser.NameFrom == nil || *callingUser.NameFrom == "github") && githubName != "" && (callingUser.Name == nil || githubName != *callingUser.Name) {
		callingUser.Name = &githubName
		callingUser.NameFrom = &githubString
		hasUpdates = true
	}
	return callingUser, hasUpdates
}
