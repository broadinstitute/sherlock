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
	googleoauth "google.golang.org/api/oauth2/v2"
)

var (
	// Email is Sherlock's own email address. This variable shouldn't be
	// accessed until Load has been called.
	Email = "sherlock-uninitialized@broadinstitute.org"
	// GoogleID is Sherlock's own subject ID. This variable shouldn't be
	// accessed until Load has been called.
	GoogleID = "sherlock-uninitialized"
)

func Load(ctx context.Context) error {
	email := config.Config.String("self.overrideEmail")
	subjectID := config.Config.String("self.overrideSubjectID")
	if email != "" && subjectID != "" {
		Email = email
		GoogleID = subjectID
		return nil
	}

	service, err := googleoauth.NewService(ctx)
	if err != nil {
		return fmt.Errorf("failed to create Google OAuth service: %w", err)
	}
	userinfo, err := service.Userinfo.V2.Me.Get().Do()
	if err != nil {
		return fmt.Errorf("failed to get self userinfo: %w", err)
	}
	Email = userinfo.Email
	GoogleID = userinfo.Id
	return nil
}
