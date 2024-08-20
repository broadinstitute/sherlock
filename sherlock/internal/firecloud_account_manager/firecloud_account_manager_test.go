package firecloud_account_manager

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func Test_firecloudAccountManager_validate(t *testing.T) {
	type fields struct {
		indexPlusOneForLocking int
		dbForLocking           *gorm.DB
		Domain                 string
		Enable                 bool
		DryRun                 bool
		OnlyAffectEmails       []string
		NeverAffectEmails      []string
		NewAccountGracePeriod  time.Duration
		InactivityThreshold    time.Duration
		workspaceClient        mockableWorkspaceClient
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "valid",
			fields: fields{
				indexPlusOneForLocking: 1,
				dbForLocking:           &gorm.DB{},
				Domain:                 "domain",
				NewAccountGracePeriod:  time.Hour,
				InactivityThreshold:    24 * time.Hour,
				workspaceClient:        &realWorkspaceClient{},
			},
			wantErr: assert.NoError,
		},
		{
			name: "missing indexPlusOneForLocking",
			fields: fields{
				dbForLocking:          &gorm.DB{},
				Domain:                "domain",
				NewAccountGracePeriod: time.Hour,
				InactivityThreshold:   24 * time.Hour,
				workspaceClient:       &realWorkspaceClient{},
			},
			wantErr: assert.Error,
		},
		{
			name: "missing dbForLocking",
			fields: fields{
				indexPlusOneForLocking: 1,
				Domain:                 "domain",
				NewAccountGracePeriod:  time.Hour,
				InactivityThreshold:    24 * time.Hour,
				workspaceClient:        &realWorkspaceClient{},
			},
			wantErr: assert.Error,
		},
		{
			name: "missing Domain",
			fields: fields{
				indexPlusOneForLocking: 1,
				dbForLocking:           &gorm.DB{},
				NewAccountGracePeriod:  time.Hour,
				InactivityThreshold:    24 * time.Hour,
				workspaceClient:        &realWorkspaceClient{},
			},
			wantErr: assert.Error,
		},
		{
			name: "missing NewAccountGracePeriod",
			fields: fields{
				indexPlusOneForLocking: 1,
				dbForLocking:           &gorm.DB{},
				Domain:                 "domain",
				InactivityThreshold:    24 * time.Hour,
				workspaceClient:        &realWorkspaceClient{},
			},
			wantErr: assert.Error,
		},
		{
			name: "missing InactivityThreshold",
			fields: fields{
				indexPlusOneForLocking: 1,
				dbForLocking:           &gorm.DB{},
				Domain:                 "domain",
				NewAccountGracePeriod:  time.Hour,
				workspaceClient:        &realWorkspaceClient{},
			},
			wantErr: assert.Error,
		},
		{
			name: "missing workspaceClient",
			fields: fields{
				indexPlusOneForLocking: 1,
				dbForLocking:           &gorm.DB{},
				Domain:                 "domain",
				NewAccountGracePeriod:  time.Hour,
				InactivityThreshold:    24 * time.Hour,
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &firecloudAccountManager{
				indexPlusOneForLocking: tt.fields.indexPlusOneForLocking,
				dbForLocking:           tt.fields.dbForLocking,
				Domain:                 tt.fields.Domain,
				Enable:                 tt.fields.Enable,
				DryRun:                 tt.fields.DryRun,
				OnlyAffectEmails:       tt.fields.OnlyAffectEmails,
				NeverAffectEmails:      tt.fields.NeverAffectEmails,
				NewAccountGracePeriod:  tt.fields.NewAccountGracePeriod,
				InactivityThreshold:    tt.fields.InactivityThreshold,
				workspaceClient:        tt.fields.workspaceClient,
			}
			tt.wantErr(t, m.validate(), "validate()")
		})
	}
}
