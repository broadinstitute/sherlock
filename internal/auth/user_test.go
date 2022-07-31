package auth

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/internal/testutils"
	"testing"
)

func TestUser_Username(t *testing.T) {
	type fields struct {
		AuthenticatedEmail      string
		MatchedFirecloudAccount *FirecloudAccount
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "reasonable email",
			fields: fields{AuthenticatedEmail: "basic@gmail.com"},
			want:   "basic",
		},
		{
			name:   "hi there RFC5321",
			fields: fields{AuthenticatedEmail: "\"foo % bar\"@I'm breaking relay syntax but only barely@[IPv6:::1]"},
			want:   "\"foo % bar\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				AuthenticatedEmail:      tt.fields.AuthenticatedEmail,
				MatchedFirecloudAccount: tt.fields.MatchedFirecloudAccount,
			}
			got := u.Username()
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

func TestUser_checkSuitability(t *testing.T) {
	type fields struct {
		AuthenticatedEmail      string
		MatchedFirecloudAccount *FirecloudAccount
	}
	tests := []struct {
		name            string
		fields          fields
		wantDescription string
		wantBool        bool
	}{
		{
			name: "no FC account",
			fields: fields{
				AuthenticatedEmail: "name@broadinstitute.org",
			},
			wantDescription: "name is not known suitable as a matching Firecloud account wasn't found",
			wantBool:        false,
		},
		{
			name: "no FC workspace ToS accepted",
			fields: fields{
				AuthenticatedEmail: "name@broadinstitute.org",
				MatchedFirecloudAccount: &FirecloudAccount{
					AcceptedGoogleTerms: false,
					EnrolledIn2fa:       true,
					Suspended:           false,
					Archived:            false,
					Groups: &FirecloudGroupMembership{
						FcAdmins:               true,
						FirecloudProjectOwners: true,
					},
				},
			},
			wantDescription: "name may be suitable but the matching Firecloud account has issues: user hasn't accepted Google Workspace terms (suggesting they've never logged in)",
			wantBool:        false,
		},
		{
			name: "no FC workspace 2FA",
			fields: fields{
				AuthenticatedEmail: "name@broadinstitute.org",
				MatchedFirecloudAccount: &FirecloudAccount{
					AcceptedGoogleTerms: true,
					EnrolledIn2fa:       false,
					Suspended:           false,
					Archived:            false,
					Groups: &FirecloudGroupMembership{
						FcAdmins:               true,
						FirecloudProjectOwners: true,
					},
				},
			},
			wantDescription: "name may be suitable but the matching Firecloud account has issues: user hasn't enrolled in two-factor authentication",
			wantBool:        false,
		},
		{
			name: "FC account suspended, no reason",
			fields: fields{
				AuthenticatedEmail: "name@broadinstitute.org",
				MatchedFirecloudAccount: &FirecloudAccount{
					AcceptedGoogleTerms: true,
					EnrolledIn2fa:       true,
					Suspended:           true,
					Archived:            false,
					Groups: &FirecloudGroupMembership{
						FcAdmins:               true,
						FirecloudProjectOwners: true,
					},
				},
			},
			wantDescription: "name may be suitable but the matching Firecloud account has issues: user is currently suspended (no reason given)",
			wantBool:        false,
		},
		{
			name: "FC account suspended, reason given",
			fields: fields{
				AuthenticatedEmail: "name@broadinstitute.org",
				MatchedFirecloudAccount: &FirecloudAccount{
					AcceptedGoogleTerms: true,
					EnrolledIn2fa:       true,
					Suspended:           true,
					Archived:            false,
					SuspensionReason:    "user is inactive",
					Groups: &FirecloudGroupMembership{
						FcAdmins:               true,
						FirecloudProjectOwners: true,
					},
				},
			},
			wantDescription: "name may be suitable but the matching Firecloud account has issues: user is currently suspended (user is inactive)",
			wantBool:        false,
		},
		{
			name: "FC account archived",
			fields: fields{
				AuthenticatedEmail: "name@broadinstitute.org",
				MatchedFirecloudAccount: &FirecloudAccount{
					AcceptedGoogleTerms: true,
					EnrolledIn2fa:       true,
					Suspended:           false,
					Archived:            true,
					Groups: &FirecloudGroupMembership{
						FcAdmins:               true,
						FirecloudProjectOwners: true,
					},
				},
			},
			wantDescription: "name may be suitable but the matching Firecloud account has issues: user is currently archived",
			wantBool:        false,
		},
		{
			name: "FC account not in fc-admins",
			fields: fields{
				AuthenticatedEmail: "name@broadinstitute.org",
				MatchedFirecloudAccount: &FirecloudAccount{
					AcceptedGoogleTerms: true,
					EnrolledIn2fa:       true,
					Suspended:           false,
					Archived:            false,
					Groups: &FirecloudGroupMembership{
						FcAdmins:               false,
						FirecloudProjectOwners: true,
					},
				},
			},
			wantDescription: fmt.Sprintf("name may be suitable but the matching Firecloud account has issues: user is not in the %s group", config.Config.MustString("auth.firecloud.groups.fcAdmins")),
			wantBool:        false,
		},
		{
			name: "FC account not in firecloud-project-owners",
			fields: fields{
				AuthenticatedEmail: "name@broadinstitute.org",
				MatchedFirecloudAccount: &FirecloudAccount{
					AcceptedGoogleTerms: true,
					EnrolledIn2fa:       true,
					Suspended:           false,
					Archived:            false,
					Groups: &FirecloudGroupMembership{
						FcAdmins:               true,
						FirecloudProjectOwners: false,
					},
				},
			},
			wantDescription: fmt.Sprintf("name may be suitable but the matching Firecloud account has issues: user is not in the %s group", config.Config.MustString("auth.firecloud.groups.firecloudProjectOwners")),
			wantBool:        false,
		},
		{
			name: "multiple problems",
			fields: fields{
				AuthenticatedEmail: "name@broadinstitute.org",
				MatchedFirecloudAccount: &FirecloudAccount{
					AcceptedGoogleTerms: false,
					EnrolledIn2fa:       true,
					Suspended:           true,
					Archived:            false,
					SuspensionReason:    "user is inactive",
					Groups: &FirecloudGroupMembership{
						FcAdmins:               true,
						FirecloudProjectOwners: false,
					},
				},
			},
			wantDescription: "name may be suitable but the matching Firecloud account has issues: user hasn't accepted Google Workspace terms (suggesting they've never logged in), user is currently suspended (user is inactive), user is not in the firecloud-project-owners@firecloud.org group",
			wantBool:        false,
		},
		{
			name: "suitable",
			fields: fields{
				AuthenticatedEmail: "name@broadinstitute.org",
				MatchedFirecloudAccount: &FirecloudAccount{
					AcceptedGoogleTerms: true,
					EnrolledIn2fa:       true,
					Suspended:           false,
					Archived:            false,
					Groups: &FirecloudGroupMembership{
						FcAdmins:               true,
						FirecloudProjectOwners: true,
					},
				},
			},
			wantDescription: "name is known suitable",
			wantBool:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				AuthenticatedEmail:      tt.fields.AuthenticatedEmail,
				MatchedFirecloudAccount: tt.fields.MatchedFirecloudAccount,
			}
			gotDescription := u.describeSuitability()
			testutils.AssertNoDiff(t, tt.wantDescription, gotDescription)
			gotBool := u.isKnownSuitable()
			testutils.AssertNoDiff(t, tt.wantBool, gotBool)
		})
	}
}
