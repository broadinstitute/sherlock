package suitability_synchronization

import (
	"github.com/stretchr/testify/assert"
	admin "google.golang.org/api/admin/directory/v1"
	"testing"
)

func Test_parseFirecloudUser(t *testing.T) {
	type args struct {
		workspaceUser *admin.User
	}
	tests := []struct {
		name                     string
		args                     args
		wantSuitable             bool
		wantDescriptionSubstring string
	}{
		{
			name: "email missing",
			args: args{
				workspaceUser: &admin.User{},
			},
			wantSuitable:             false,
			wantDescriptionSubstring: "doesn't appear to have a primary email",
		},
		{
			name: "not agreed to terms",
			args: args{
				workspaceUser: &admin.User{
					PrimaryEmail: "foo@example.com",
				},
			},
			wantSuitable:             false,
			wantDescriptionSubstring: "hasn't accepted Google Workspace terms (suggesting they've never logged in",
		},
		{
			name: "no second factor",
			args: args{
				workspaceUser: &admin.User{
					PrimaryEmail:  "foo@example.com",
					AgreedToTerms: true,
				},
			},
			wantSuitable:             false,
			wantDescriptionSubstring: "hasn't enrolled in two-factor authentication",
		},
		{
			name: "suspended",
			args: args{
				workspaceUser: &admin.User{
					PrimaryEmail:    "foo@example.com",
					AgreedToTerms:   true,
					IsEnrolledIn2Sv: true,
					Suspended:       true,
				},
			},
			wantSuitable:             false,
			wantDescriptionSubstring: "is suspended, probably due to inactivity",
		},
		{
			name: "archived",
			args: args{
				workspaceUser: &admin.User{
					PrimaryEmail:    "foo@example.com",
					AgreedToTerms:   true,
					IsEnrolledIn2Sv: true,
					Suspended:       false,
					Archived:        true,
				},
			},
			wantSuitable:             false,
			wantDescriptionSubstring: "is archived",
		},
		{
			name: "suitable",
			args: args{
				workspaceUser: &admin.User{
					PrimaryEmail:    "foo@example.com",
					AgreedToTerms:   true,
					IsEnrolledIn2Sv: true,
					Suspended:       false,
					Archived:        false,
				},
			},
			wantSuitable:             true,
			wantDescriptionSubstring: "is suitable",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSuitable, gotDescription := parseFirecloudUser(tt.args.workspaceUser)
			assert.Equalf(t, tt.wantSuitable, gotSuitable, "parseFirecloudUser(%v)", tt.args.workspaceUser)
			assert.Containsf(t, gotDescription, tt.wantDescriptionSubstring, "parseFirecloudUser(%v)", tt.args.workspaceUser)
		})
	}
}
