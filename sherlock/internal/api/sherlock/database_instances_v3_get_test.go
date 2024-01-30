package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
)

func (s *handlerSuite) TestDatabaseInstanceV3Get_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/database-instances/v3/something/with/too/many/slashes", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestDatabaseInstanceV3Get_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/database-instances/v3/chart-release/not-found", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestDatabaseInstanceV3Get() {
	databaseInstance := s.TestData.DatabaseInstance_LeonardoDev()

	var got DatabaseInstanceV3
	code := s.HandleRequest(
		s.NewRequest("GET", fmt.Sprintf("/api/database-instances/v3/%d", databaseInstance.ID), nil),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.Platform) {
		s.Equal(*databaseInstance.Platform, *got.Platform)
	}
}

func (s *handlerSuite) Test_databaseInstanceModelFromSelector() {
	type args struct {
		selector string
	}
	tests := []struct {
		name      string
		args      args
		wantQuery models.DatabaseInstance
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
			name:    "valid id",
			args:    args{selector: "123"},
			wantErr: assert.NoError,
			wantQuery: models.DatabaseInstance{
				Model: gorm.Model{ID: 123},
			},
		},
		{
			name:    "invalid id",
			args:    args{selector: testutils.StringNumberTooBigForInt},
			wantErr: assert.Error,
		},
		{
			name:    "chart release",
			args:    args{selector: "chart-release/" + s.TestData.ChartRelease_LeonardoProd().Name},
			wantErr: assert.NoError,
			wantQuery: models.DatabaseInstance{
				ChartReleaseID: s.TestData.ChartRelease_LeonardoProd().ID,
			},
		},
		{
			name:    "invalid chart release",
			args:    args{selector: "chart-release/!!!!!"},
			wantErr: assert.Error,
		},
		{
			name:    "not found chart release",
			args:    args{selector: "chart-release/not-found"},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			gotQuery, err := databaseInstanceModelFromSelector(s.DB, tt.args.selector)
			if !tt.wantErr(s.T(), err, fmt.Sprintf("databaseInstanceModelFromSelector(%v)", tt.args.selector)) {
				return
			}
			s.Equalf(tt.wantQuery, gotQuery, "databaseInstanceModelFromSelector(%v)", tt.args.selector)
		})
	}
}
