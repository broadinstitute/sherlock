package authorization

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/stretchr/testify/assert"
	admin "google.golang.org/api/admin/directory/v1"
	"testing"
)

func Test_parseSuitability(t *testing.T) {
	type args struct {
		workspaceUser                     *admin.User
		fcAdminsGroupEmails               []string
		firecloudProjectOwnersGroupEmails []string
	}
	tests := []struct {
		name string
		args args
		want *Suitability
	}{
		{
			name: "no email",
			args: args{workspaceUser: &admin.User{}},
			want: &Suitability{
				suitable:    false,
				description: "firecloud user doesn't appear to have a primary email? something's amiss, marking as not suitable",
				source:      FIRECLOUD,
			},
		},
		{
			name: "not logged in",
			args: args{workspaceUser: &admin.User{
				PrimaryEmail: "example@example.com",
			}},
			want: &Suitability{
				suitable: false,
				description: fmt.Sprintf("firecloud user hasn't accepted Google Workspace terms (suggesting they've never logged in; they'll need to wait %d minutes after first login for Sherlock to pick it up)",
					config.Config.MustInt("auth.updateIntervalMinutes")),
				source: FIRECLOUD,
			},
		},
		{
			name: "no 2FA",
			args: args{workspaceUser: &admin.User{
				PrimaryEmail:  "example@example.com",
				AgreedToTerms: true,
			}},
			want: &Suitability{
				suitable:    false,
				description: "firecloud user hasn't enrolled in two-factor authentication",
				source:      FIRECLOUD,
			},
		},
		{
			name: "suspended",
			args: args{workspaceUser: &admin.User{
				PrimaryEmail:    "example@example.com",
				AgreedToTerms:   true,
				IsEnrolledIn2Sv: true,
				Suspended:       true,
			}},
			want: &Suitability{
				suitable: false,
				description: fmt.Sprintf("firecloud user is suspended, probably due to inactivity (reach out to #dsp-devops-champions for help; they'll need to wait %d minutes after reactivation for Sherlock to pick it up)",
					config.Config.MustInt("auth.updateIntervalMinutes")),
				source: FIRECLOUD,
			},
		},
		{
			name: "archived",
			args: args{workspaceUser: &admin.User{
				PrimaryEmail:    "example@example.com",
				AgreedToTerms:   true,
				IsEnrolledIn2Sv: true,
				Archived:        true,
			}},
			want: &Suitability{
				suitable:    false,
				description: "firecloud user is archived",
				source:      FIRECLOUD,
			},
		},
		{
			name: "no fc admins group",
			args: args{workspaceUser: &admin.User{
				PrimaryEmail:    "example@example.com",
				AgreedToTerms:   true,
				IsEnrolledIn2Sv: true,
			}},
			want: &Suitability{
				suitable: false,
				description: fmt.Sprintf("firecloud user isn't in fc-admins group (reach out to #dsp-devops-champions for help; they'll need to wait %d minutes after being added for Sherlock to pick it up)",
					config.Config.MustInt("auth.updateIntervalMinutes")),
				source: FIRECLOUD,
			},
		},
		{
			name: "no firecloud project owners group",
			args: args{
				workspaceUser: &admin.User{
					PrimaryEmail:    "example@example.com",
					AgreedToTerms:   true,
					IsEnrolledIn2Sv: true,
				},
				fcAdminsGroupEmails: []string{"example@example.com"},
			},
			want: &Suitability{
				suitable: false,
				description: fmt.Sprintf("firecloud user isn't in firecloud-project-owners group (reach out to #dsp-devops-champions for help; they'll need to wait %d minutes after being added for Sherlock to pick it up)",
					config.Config.MustInt("auth.updateIntervalMinutes")),
				source: FIRECLOUD,
			},
		},
		{
			name: "suitable",
			args: args{
				workspaceUser: &admin.User{
					PrimaryEmail:    "example@example.com",
					AgreedToTerms:   true,
					IsEnrolledIn2Sv: true,
				},
				fcAdminsGroupEmails:               []string{"example@example.com"},
				firecloudProjectOwnersGroupEmails: []string{"example@example.com"},
			},
			want: &Suitability{
				suitable:    true,
				description: "firecloud user is suitable",
				source:      FIRECLOUD,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, parseSuitability(tt.args.workspaceUser, tt.args.fcAdminsGroupEmails, tt.args.firecloudProjectOwnersGroupEmails), "parseSuitability(%v, %v, %v)", tt.args.workspaceUser, tt.args.fcAdminsGroupEmails, tt.args.firecloudProjectOwnersGroupEmails)
		})
	}
}
