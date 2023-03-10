package auth_models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/internal/testutils"
	"testing"
)

func TestUser_Username(t *testing.T) {
	type fields struct {
		Email string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "reasonable email",
			fields: fields{Email: "basic@gmail.com"},
			want:   "basic",
		},
		{
			name:   "hi there RFC5321",
			fields: fields{Email: "\"foo % bar\"@I'm breaking relay syntax but only barely@[IPv6:::1]"},
			want:   "\"foo % bar\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				StoredControlledUserFields: StoredControlledUserFields{
					Email: tt.fields.Email,
				},
			}
			got := u.Username()
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

func TestUser_checkSuitability(t *testing.T) {
	type fields struct {
		Email                   string
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
				Email: "name@broadinstitute.org",
			},
			wantDescription: "name is not known suitable as a matching Firecloud account wasn't found",
			wantBool:        false,
		},
		{
			name: "no FC workspace ToS accepted",
			fields: fields{
				Email: "name@broadinstitute.org",
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
				Email: "name@broadinstitute.org",
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
				Email: "name@broadinstitute.org",
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
				Email: "name@broadinstitute.org",
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
				Email: "name@broadinstitute.org",
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
				Email: "name@broadinstitute.org",
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
				Email: "name@broadinstitute.org",
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
				Email: "name@broadinstitute.org",
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
				Email: "name@broadinstitute.org",
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
				StoredControlledUserFields: StoredControlledUserFields{
					Email: tt.fields.Email,
				},
				InferredUserFields: InferredUserFields{
					MatchedFirecloudAccount: tt.fields.MatchedFirecloudAccount,
				},
			}
			gotDescription := u.describeSuitability()
			testutils.AssertNoDiff(t, tt.wantDescription, gotDescription)
			gotBool := u.isKnownSuitable()
			testutils.AssertNoDiff(t, tt.wantBool, gotBool)
		})
	}
}

func TestUser_AlphaNumericHyphenatedUsername(t *testing.T) {
	type fields struct {
		Email                   string
		MatchedFirecloudAccount *FirecloudAccount
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "normal BI username",
			fields: fields{Email: "someone@broadinstitute.org"},
			want:   "someone",
		},
		{
			name:   "with separators",
			fields: fields{Email: "someone.else-blah_blah@somewhere.info"},
			want:   "someone-else-blah-blah",
		},
		{
			name:   "strips invalid",
			fields: fields{Email: "1a bc?de.23"},
			want:   "1abcde-23",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				StoredControlledUserFields: StoredControlledUserFields{
					Email: tt.fields.Email,
				},
				InferredUserFields: InferredUserFields{
					MatchedFirecloudAccount: tt.fields.MatchedFirecloudAccount,
				},
			}
			got := u.AlphaNumericHyphenatedUsername()
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}
