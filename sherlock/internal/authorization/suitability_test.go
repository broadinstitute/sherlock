package authorization

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/local_user"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/test_users"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

// some of the tests here aren't particularly enlightened but if we mess up getters we'd have a bad time,
// so may as well check

func TestGetSuitabilityFor(t *testing.T) {
	config.LoadTestConfig()
	cachedFirecloudSuitability = map[string]*Suitability{
		"suitable@firecloud.org":          {suitable: true, source: FIRECLOUD, description: "yes"},
		"suitable@broadinstitute.org":     {suitable: true, source: FIRECLOUD, description: "yes"},
		"not-suitable@firecloud.org":      {suitable: false, source: FIRECLOUD, description: "no"},
		"not-suitable@broadinstitute.org": {suitable: false, source: FIRECLOUD, description: "no"},
		"no-bi-email@firecloud.org":       {suitable: true, source: FIRECLOUD, description: "yes"},
		"in-both@firecloud.org":           {suitable: false, source: FIRECLOUD, description: "no"},
	}
	cachedConfigSuitability = map[string]*Suitability{
		"has-extra-permissions-suitable@example.com":     {suitable: true, source: CONFIG, description: "yes"},
		"has-extra-permissions-non-suitable@example.com": {suitable: false, source: CONFIG, description: "no"},
		"in-both@firecloud.org":                          {suitable: true, source: CONFIG, description: "yes"},
	}

	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want *Suitability
	}{
		{
			name: "suitable FC",
			args: args{email: "suitable@firecloud.org"},
			want: &Suitability{suitable: true, source: FIRECLOUD, description: "yes"},
		},
		{
			name: "suitable BI",
			args: args{email: "suitable@broadinstitute.org"},
			want: &Suitability{suitable: true, source: FIRECLOUD, description: "yes"},
		},
		{
			name: "not suitable FC",
			args: args{email: "not-suitable@firecloud.org"},
			want: &Suitability{suitable: false, source: FIRECLOUD, description: "no"},
		},
		{
			name: "not suitable BI",
			args: args{email: "not-suitable@broadinstitute.org"},
			want: &Suitability{suitable: false, source: FIRECLOUD, description: "no"},
		},
		{
			name: "misconfigured FC recovery email",
			args: args{email: "no-bi-email@broadinstitute.org"},
			want: &Suitability{suitable: false, source: NONE, description: fmt.Sprintf("SUPER IMPORTANT: %s isn't recognized, but the same-named %s is. That almost definitely means the %s Firecloud user account is misconfigured and should have %s set as the associated recovery email. DevOps needs to fix this. Once they do, you'll need to wait %d minutes for Sherlock to pick up the update.",
				"no-bi-email@broadinstitute.org", "no-bi-email@firecloud.org", "no-bi-email@firecloud.org", "no-bi-email@broadinstitute.org", config.Config.MustInt("auth.updateIntervalMinutes"))},
		},
		{
			name: "suitable config",
			args: args{email: "has-extra-permissions-suitable@example.com"},
			want: &Suitability{suitable: true, source: CONFIG, description: "yes"},
		},
		{
			name: "not suitable config",
			args: args{email: "has-extra-permissions-non-suitable@example.com"},
			want: &Suitability{suitable: false, source: CONFIG, description: "no"},
		},
		{
			name: "FC overrides config if conflict",
			args: args{email: "in-both@firecloud.org"},
			want: &Suitability{suitable: false, source: FIRECLOUD, description: "no"},
		},
		{
			name: "unknown",
			args: args{email: "random@example.com"},
			want: &Suitability{suitable: false, source: NONE, description: "user random@example.com lacks production suitability"},
		},
		{
			name: "suitable test user",
			args: args{email: test_users.SuitableTestUserEmail},
			want: &Suitability{suitable: true, source: CONFIG, description: fmt.Sprintf("test user; email %s equal to suitable %s: true",
				test_users.SuitableTestUserEmail, test_users.SuitableTestUserEmail)},
		},
		{
			name: "non-suitable test user",
			args: args{email: test_users.NonSuitableTestUserEmail},
			want: &Suitability{suitable: false, source: CONFIG, description: fmt.Sprintf("test user; email %s equal to suitable %s: false",
				test_users.NonSuitableTestUserEmail, test_users.SuitableTestUserEmail)},
		},
		{
			name: "local user",
			args: args{email: local_user.LocalUserEmail},
			want: &Suitability{suitable: local_user.LocalUserSuitable, source: CONFIG, description: fmt.Sprintf("local user; suitable: %v",
				local_user.LocalUserSuitable)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GetSuitabilityFor(tt.args.email), "GetSuitabilityFor(%v)", tt.args.email)
		})
	}
}

func TestSuitability_Description(t *testing.T) {
	type fields struct {
		suitable    bool
		description string
		source      SuitabilitySource
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "A",
			fields: fields{description: "A"},
			want:   "A",
		},
		{
			name:   "B",
			fields: fields{description: "B"},
			want:   "B",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Suitability{
				suitable:    tt.fields.suitable,
				description: tt.fields.description,
				source:      tt.fields.source,
			}
			assert.Equalf(t, tt.want, s.Description(), "Description()")
		})
	}
}

func TestSuitability_Source(t *testing.T) {
	type fields struct {
		suitable    bool
		description string
		source      SuitabilitySource
	}
	tests := []struct {
		name   string
		fields fields
		want   SuitabilitySource
	}{
		{
			name:   "nil = none",
			fields: fields{},
			want:   NONE,
		},
		{
			name:   "none",
			fields: fields{source: NONE},
			want:   NONE,
		},
		{
			name:   "firecloud",
			fields: fields{source: FIRECLOUD},
			want:   FIRECLOUD,
		},
		{
			name:   "config",
			fields: fields{source: CONFIG},
			want:   CONFIG,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Suitability{
				suitable:    tt.fields.suitable,
				description: tt.fields.description,
				source:      tt.fields.source,
			}
			assert.Equalf(t, tt.want, s.Source(), "Source()")
		})
	}
}

func TestSuitability_Suitable(t *testing.T) {
	type fields struct {
		suitable    bool
		description string
		source      SuitabilitySource
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "nil = false",
			fields: fields{},
			want:   false,
		},
		{
			name:   "false",
			fields: fields{suitable: false},
			want:   false,
		},
		{
			name:   "true",
			fields: fields{suitable: true},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Suitability{
				suitable:    tt.fields.suitable,
				description: tt.fields.description,
				source:      tt.fields.source,
			}
			assert.Equalf(t, tt.want, s.Suitable(), "Suitable()")
		})
	}
}

func TestSuitability_SuitableOrError(t *testing.T) {
	type fields struct {
		suitable    bool
		description string
		source      SuitabilitySource
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "nil",
			fields:  fields{},
			wantErr: assert.Error,
		},
		{
			name: "not suitable",
			fields: fields{
				suitable: false,
				source:   FIRECLOUD,
			},
			wantErr: assert.Error,
		},
		{
			name: "no source (?)",
			fields: fields{
				suitable: true,
				source:   NONE,
			},
			wantErr: assert.Error,
		},
		{
			name: "description used as error",
			fields: fields{
				description: "some error",
			},
			wantErr: func(t assert.TestingT, err error, _ ...interface{}) bool {
				return assert.EqualError(t, err, "some error")
			},
		},
		{
			name: "suitable",
			fields: fields{
				suitable: true,
				source:   FIRECLOUD,
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Suitability{
				suitable:    tt.fields.suitable,
				description: tt.fields.description,
				source:      tt.fields.source,
			}
			tt.wantErr(t, s.SuitableOrError(), "SuitableOrError()")
		})
	}
}
