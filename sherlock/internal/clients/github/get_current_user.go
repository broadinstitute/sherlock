package github

import (
	"context"
	"strconv"

	"github.com/google/go-github/v58/github"
	"golang.org/x/oauth2"
)

// GetCurrentUser gets the currently authenticated user's info. It is a bit weird because
// this often won't use Sherlock's own auth, it'll use the authTokenOverride so we get the
// info for some other user.
func GetCurrentUser(ctx context.Context, authTokenOverride ...string) (githubID string, username string, name string, err error) {
	if isEnabled() {
		var user *github.User
		// If we're mocked, always use the mock.
		// If we don't have an auth token to use, always use our built-in client regardless.
		if isMocked() || len(authTokenOverride) == 0 {
			user, _, err = client.Users.Get(ctx, "")
		} else if len(authTokenOverride) > 0 {
			user, _, err = github.NewClient(oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: authTokenOverride[0]}))).Users.Get(ctx, "")
		}
		if err == nil && user != nil {
			if user.ID != nil {
				githubID = strconv.FormatInt(*user.ID, 10)
			}
			if user.Login != nil {
				username = *user.Login
			}
			if user.Name != nil {
				name = *user.Name
			}
		}
	}
	return
}
