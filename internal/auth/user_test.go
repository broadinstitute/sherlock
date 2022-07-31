package auth

import (
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

func TestUser_describeSuitability(t *testing.T) {
	type fields struct {
		AuthenticatedEmail      string
		MatchedFirecloudAccount *FirecloudAccount
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				AuthenticatedEmail:      tt.fields.AuthenticatedEmail,
				MatchedFirecloudAccount: tt.fields.MatchedFirecloudAccount,
			}
			if got := u.describeSuitability(); got != tt.want {
				t.Errorf("describeSuitability() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_isKnownSuitable(t *testing.T) {
	type fields struct {
		AuthenticatedEmail      string
		MatchedFirecloudAccount *FirecloudAccount
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				AuthenticatedEmail:      tt.fields.AuthenticatedEmail,
				MatchedFirecloudAccount: tt.fields.MatchedFirecloudAccount,
			}
			if got := u.isKnownSuitable(); got != tt.want {
				t.Errorf("isKnownSuitable() = %v, want %v", got, tt.want)
			}
		})
	}
}
