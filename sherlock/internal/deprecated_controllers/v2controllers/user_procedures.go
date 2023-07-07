package v2controllers

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/google/go-github/v50/github"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
	"time"
)

type GithubAccessPayload struct {
	GithubAccessToken string `json:"githubAccessToken"`
}

// UpdateUserGithubAssociation is a controlled update of the GitHub fields recorded for each user. Rather than
// accepting modifications to the fields itself, it accepts a GitHub access token and uses that to get the username/id,
// which makes sure that users can only register an association with a GitHub account they control.
func (c UserController) UpdateUserGithubAssociation(githubAccess GithubAccessPayload, user *models.User) (User, bool, error) {
	if githubAccess.GithubAccessToken == "" {
		return User{}, false, fmt.Errorf("(%s) no github access token provided", errors.BadRequest)
	} else if user == nil {
		// We shouldn't be able to hit this case, but it is a pointer and this is a bespoke controller method so better
		// to have an error message than hit the nil pointer.
		return User{}, false, fmt.Errorf("(%s) user not passed from middleware", errors.InternalServerError)
	}
	// Arbitrary long-ish timeout; Beehive doesn't call this endpoint synchronously
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubAccess.GithubAccessToken},
	)
	httpClient := oauth2.NewClient(ctx, tokenSource)
	githubClient := github.NewClient(httpClient)

	githubUser, _, err := githubClient.Users.Get(ctx, "")
	if err != nil {
		return User{}, false, fmt.Errorf("(%s) github api error with user's github access token: %v", errors.BadRequest, err)
	} else if githubUser == nil {
		return User{}, false, fmt.Errorf("(%s) github api call with user's github access token didn't error but the response object was nil", errors.BadRequest)
	} else if githubUser.ID == nil || githubUser.Login == nil {
		return User{}, false, fmt.Errorf("(%s) github api call with user's github access token didn't error but the response didn't contain the id and login fields", errors.InternalServerError)
	} else {
		return c.recordGithubInformation(githubUser, user)
	}
}

func (c UserController) recordGithubInformation(githubInformation *github.User, user *models.User) (User, bool, error) {
	var editsToMake v2models.User
	githubUserIdString := fmt.Sprintf("%d", *githubInformation.ID)
	if user.GithubID == nil || user.GithubUsername == nil { // If we don't store a github user
		log.Info().Msgf("GH   | user %s first-time linking github account %s (ID %s)", user.Email, *githubInformation.Login, githubUserIdString)
		editsToMake.GithubUsername = githubInformation.Login
		editsToMake.GithubID = &githubUserIdString
	} else if *user.GithubID != githubUserIdString { // If the stored github user is a different user (different ID)
		log.Info().Msgf("GH   | user %s changing linked github account from %s (ID %s) to %s (ID %s)", user.Email, *user.GithubUsername, *user.GithubID, *githubInformation.Login, githubUserIdString)
		editsToMake.GithubUsername = githubInformation.Login
		editsToMake.GithubID = &githubUserIdString
	} else if *user.GithubUsername != *githubInformation.Login { // If the stored github user is the same, just new username
		log.Info().Msgf("GH   | user %s linked github account (ID %s) has new username, from %s to %s", user.Email, *user.GithubID, *user.GithubUsername, *githubInformation.Login)
		editsToMake.GithubUsername = githubInformation.Login
	}

	// If we have a github name, and we either don't store a name or should infer a different name than what we store
	if githubInformation.Name != nil && (user.Name == nil || ((user.NameInferredFromGithub == nil || *user.NameInferredFromGithub) && *githubInformation.Name != *user.Name)) {
		log.Info().Msgf("GH   | user %s github account data contained name %s, recording", user.Email, *githubInformation.Name)
		editsToMake.Name = githubInformation.Name
		trueValue := true
		editsToMake.NameInferredFromGithub = &trueValue
	}

	if editsToMake.GithubID != nil || editsToMake.GithubUsername != nil || editsToMake.Name != nil {
		modelUser, err := c.primaryStore.Edit(user.Email, editsToMake, user)
		return *c.modelToReadable(&modelUser), true, err
	} else {
		returnableUser, err := c.Get(user.Email)
		return returnableUser, false, err
	}
}
