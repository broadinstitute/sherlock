package auth

import "testing"

func GenerateUser(t *testing.T, suitable bool) *User {
	if t == nil {
		t.Errorf("refusing to generate user, testing.T was nil")
		return nil
	}
	var suspensionReason string
	if !suitable {
		suspensionReason = "user was generated to be non-suitable"
	}
	return &User{
		AuthenticatedEmail: "sherlock.holmes@broadinstitute.org",
		MatchedFirecloudAccount: &FirecloudAccount{
			Email:               "sherlock.holmes@firecloud.org",
			AcceptedGoogleTerms: true,
			EnrolledIn2fa:       true,
			Suspended:           !suitable,
			Archived:            false,
			SuspensionReason:    suspensionReason,
			Groups: &FirecloudGroupMembership{
				FcAdmins:               true,
				FirecloudProjectOwners: true,
			},
		},
		offline: true,
	}
}
