package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func Test_userModelFromSelector(t *testing.T) {
	type args struct {
		selector string
	}
	tests := []struct {
		name      string
		args      args
		wantQuery models.User
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name:    "empty",
			args:    args{selector: ""},
			wantErr: assert.Error,
		},
		{
			name:      "id",
			args:      args{selector: "123"},
			wantQuery: models.User{Model: gorm.Model{ID: 123}},
			wantErr:   assert.NoError,
		},
		{
			name:    "invalid id",
			args:    args{selector: testutils.StringNumberTooBigForInt},
			wantErr: assert.Error,
		},
		{
			name:      "email",
			args:      args{selector: "foo@example.com"},
			wantQuery: models.User{Email: "foo@example.com"},
			wantErr:   assert.NoError,
		},
		{
			name:      "google id",
			args:      args{selector: "google-id/foo"},
			wantQuery: models.User{GoogleID: "foo"},
			wantErr:   assert.NoError,
		},
		{
			name:    "empty google id",
			args:    args{selector: "google-id/"},
			wantErr: assert.Error,
		},
		{
			name:      "github username",
			args:      args{selector: "github/foo"},
			wantQuery: models.User{GithubUsername: testutils.PointerTo("foo")},
			wantErr:   assert.NoError,
		},
		{
			name:    "empty github username",
			args:    args{selector: "github/"},
			wantErr: assert.Error,
		},
		{
			name:      "github id",
			args:      args{selector: "github-id/foo"},
			wantQuery: models.User{GithubID: testutils.PointerTo("foo")},
			wantErr:   assert.NoError,
		},
		{
			name:    "empty github id",
			args:    args{selector: "github-id/"},
			wantErr: assert.Error,
		},
		{
			name:    "invalid",
			args:    args{selector: "foo"},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQuery, err := userModelFromSelector(tt.args.selector)
			if !tt.wantErr(t, err, fmt.Sprintf("userModelFromSelector(%v)", tt.args.selector)) {
				return
			}
			assert.Equalf(t, tt.wantQuery, gotQuery, "userModelFromSelector(%v)", tt.args.selector)
		})
	}
}
