package auth

import (
	"github.com/broadinstitute/sherlock/internal/auth/auth_models"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"gorm.io/gorm"
	"testing"
)

func GenerateUser(t *testing.T, db *gorm.DB, suitable bool) *auth_models.User {
	if t == nil {
		t.Errorf("refusing to generate user, testing.T was nil")
		return nil
	}

	middleware := v2models.NewMiddlewareUserStore(db)
	user, err := middleware.GetOrCreateUser("sherlock.holmes@broadinstitute.org", "someidwouldgohere")
	if err != nil {
		t.Error(err)
		return nil
	}

	var suspensionReason string
	if !suitable {
		suspensionReason = "user was generated to be non-suitable"
	}
	return &auth_models.User{
		StoredControlledUserFields: user.StoredControlledUserFields,
		StoredMutableUserFields:    user.StoredMutableUserFields,
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
