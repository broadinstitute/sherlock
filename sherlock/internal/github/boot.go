package github

import (
	"context"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/google/go-github/v50/github"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
)

func Init(ctx context.Context) error {
	if config.Config.Bool("github.enable") {
		tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: config.Config.MustString("github.token")})
		rawClient = github.NewClient(oauth2.NewClient(ctx, tokenSource))
		githubUser, _, err := rawClient.Users.Get(ctx, "")
		if err != nil {
			return err
		} else {
			log.Info().Msgf("GH   | successfully authenticated to GitHub as \"%s\"", *githubUser.Login)
		}
		setClientFromRawClient()
	}
	return nil
}
