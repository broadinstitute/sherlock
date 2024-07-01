// Package self is responsible for telling Sherlock who it is.
//
// This is necessary because Sherlock is its own source-of-truth for
// access control. When Sherlock takes actions of its own accord, it
// needs to know how to attribute those actions to itself in logs
// or journaled database entries.
//
// For simplicity's sake, this package's goal is for Sherlock to be
// able to upsert its own models.User record into the database. This
// package doesn't actually *do* that (import cycles...), but it tries
// to expose the necessary information in an encapsulated way.
package self

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	oauth2google "golang.org/x/oauth2/google"
	googleoauth "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

var (
	// Email is Sherlock's own email address. This variable shouldn't be
	// accessed until Load has been called.
	Email = "sherlock-uninitialized@broadinstitute.org"
	// GoogleID is Sherlock's own subject ID. This variable shouldn't be
	// accessed until Load has been called.
	GoogleID = "accounts.google.com:sherlock-uninitialized"
)

func Load(ctx context.Context) error {
	email := config.Config.String("self.overrideEmail")
	subjectID := config.Config.String("self.overrideSubjectID")
	if email != "" && subjectID != "" {
		Email = email
		GoogleID = subjectID
		return nil
	}

	// You might think we could rely on googleoauth.NewService to pick up ADC, but you'd actually be wrong
	// for local usage. For whatever reason, doing that with local gcloud ADC fails, complaining about your
	// permissions for whatever your quota project is configured as. This happens even if you try to pass
	// an option to googleoauth to set the project to an empty string.
	//
	// The solution is to make a tokenSource based on ADC and pass it to googleoauth.NewService. This works
	// because it loses the information about the project. It's a hack but it does work, so whatever.
	tokenSource, err := oauth2google.DefaultTokenSource(ctx)
	if err != nil {
		return fmt.Errorf("failed to get default token source: %w", err)
	}
	service, err := googleoauth.NewService(ctx, option.WithTokenSource(tokenSource))
	if err != nil {
		return fmt.Errorf("failed to create Google OAuth service: %w", err)
	}
	userinfo, err := service.Userinfo.V2.Me.Get().Do()
	if err != nil {
		return fmt.Errorf("failed to get self userinfo: %w", err)
	}
	Email = userinfo.Email
	// Sherlock stores Google Subject IDs the way it observes them from IAP, which means with the leading
	// "accounts.google.com:" prefix. There's no way to get that full prefixed ID from the userinfo endpoint,
	// so we add it ourselves for consistency. Maybe a decade down the road Google will add other possible
	// prefixes, but presumably they won't change the meaning of this field of this endpoint.
	GoogleID = "accounts.google.com:" + userinfo.Id
	return nil
}
