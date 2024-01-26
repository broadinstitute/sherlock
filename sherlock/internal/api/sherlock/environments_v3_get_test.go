package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"testing"
)

func (s *handlerSuite) TestEnvironmentV3Get_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/environments/v3/something/with/too/many/slashes", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestEnvironmentV3Get_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/environments/v3/123", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestEnvironmentV3Get() {
	environment := s.TestData.Environment_Prod()

	var got EnvironmentV3
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/environments/v3/"+environment.Name, nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Equal(environment.Name, got.Name)
}

func Test_environmentModelFromSelector(t *testing.T) {
	type args struct {
		selector string
	}
	tests := []struct {
		name      string
		args      args
		wantQuery models.Environment
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name:    "empty",
			args:    args{selector: ""},
			wantErr: assert.Error,
		},
		{
			name:    "invalid",
			args:    args{selector: "something obviously invalid!"},
			wantErr: assert.Error,
		},
		{
			name: "valid id",
			args: args{selector: "123"},
			wantQuery: models.Environment{
				Model: gorm.Model{ID: 123},
			},
			wantErr: assert.NoError,
		},
		{
			name:    "invalid id",
			args:    args{selector: testutils.StringNumberTooBigForInt},
			wantErr: assert.Error,
		},
		{
			name: "name",
			args: args{selector: "prod"},
			wantQuery: models.Environment{
				Name: "prod",
			},
			wantErr: assert.NoError,
		},
		{
			name: "unique resource prefix",
			args: args{selector: "resource-prefix/abcd"},
			wantQuery: models.Environment{
				UniqueResourcePrefix: "abcd",
			},
			wantErr: assert.NoError,
		},
		{
			name:    "invalid unique resource prefix",
			args:    args{selector: "resource-prefix/abcde"},
			wantErr: assert.Error,
		},
		{
			name:    "invalid differentiator on selector",
			args:    args{selector: "something-else/abcd"},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQuery, err := environmentModelFromSelector(tt.args.selector)
			if !tt.wantErr(t, err, fmt.Sprintf("environmentModelFromSelector(%v)", tt.args.selector)) {
				return
			}
			assert.Equalf(t, tt.wantQuery, gotQuery, "environmentModelFromSelector(%v)", tt.args.selector)
		})
	}
}
