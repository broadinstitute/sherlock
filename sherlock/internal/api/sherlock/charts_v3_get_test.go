package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"testing"
)

func (s *handlerSuite) TestChartV3Get_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/charts/v3/something/with/slashes", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChartV3Get_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/charts/v3/my-chart", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestChartV3Get() {
	s.NoError(s.DB.Create(&models.Chart{Name: "my-chart", ChartRepo: utils.PointerTo("some-repo")}).Error)

	var got ChartV3
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/charts/v3/my-chart", nil),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.ChartRepo) {
		s.Equal("some-repo", *got.ChartRepo)
	}
}

func Test_chartModelFromSelector(t *testing.T) {
	type args struct {
		selector string
	}
	tests := []struct {
		name      string
		args      args
		wantQuery models.Chart
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
			name:      "valid id",
			args:      args{selector: "123"},
			wantQuery: models.Chart{Model: gorm.Model{ID: 123}},
			wantErr:   assert.NoError,
		},
		{
			name:    "invalid id",
			args:    args{selector: testutils.StringNumberTooBigForInt},
			wantErr: assert.Error,
		},
		{
			name:      "valid name",
			args:      args{selector: "foo-bar-2"},
			wantQuery: models.Chart{Name: "foo-bar-2"},
			wantErr:   assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQuery, err := chartModelFromSelector(tt.args.selector)
			if !tt.wantErr(t, err, fmt.Sprintf("chartModelFromSelector(%v)", tt.args.selector)) {
				return
			}
			assert.Equalf(t, tt.wantQuery, gotQuery, "chartModelFromSelector(%v)", tt.args.selector)
		})
	}
}
