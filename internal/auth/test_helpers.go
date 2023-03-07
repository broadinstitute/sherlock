package auth

import (
	"github.com/broadinstitute/sherlock/internal/auth/auth_models"
	"testing"
)

func GenerateUser(t *testing.T, suitable bool) *auth_models.User {
	if t == nil {
		t.Errorf("refusing to generate user, testing.T was nil")
		return nil
	}
	var suspensionReason string
	if !suitable {
		suspensionReason = "user was generated to be non-suitable"
	}
	return &auth_models.User{
		StoredUserFields: auth_models.StoredUserFields{
			Email:    "sherlock.holmes@broadinstitute.org",
			GoogleID: "someidwouldgohere",
		},
		InferredUserFields: auth_models.InferredUserFields{
			MatchedFirecloudAccount: &auth_models.FirecloudAccount{
				Email:               "sherlock.holmes@firecloud.org",
				AcceptedGoogleTerms: true,
				EnrolledIn2fa:       true,
				Suspended:           !suitable,
				Archived:            false,
				SuspensionReason:    suspensionReason,
				Groups: &auth_models.FirecloudGroupMembership{
					FcAdmins:               true,
					FirecloudProjectOwners: true,
				},
			},
		},
	}
}
