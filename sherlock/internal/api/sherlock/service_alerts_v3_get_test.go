package sherlock

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func (s *handlerSuite) TestServiceAlertV3Get_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/service-alerts/v3/something/with/slashes", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestServiceAlertV3Get_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/service-alerts/v3/123", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestServiceAlertV3Get() {
	s.SetNonSuitableTestUserForDB()
	serviceAlert1 := s.TestData.ServiceAlert_1()

	var got ServiceAlertV3
	code := s.HandleRequest(
		s.NewRequest("GET", fmt.Sprintf("/api/service-alerts/v3/%d", serviceAlert1.ID), nil),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.AlertMessage) {
		s.Equal(*serviceAlert1.AlertMessage, *got.AlertMessage)
	}
}

func Test_serviceAlertModelFromSelector(t *testing.T) {
	type args struct {
		selector string
	}
	tests := []struct {
		name      string
		args      args
		wantQuery models.ServiceAlert
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
			wantQuery: models.ServiceAlert{Model: gorm.Model{ID: 123}},
			wantErr:   assert.NoError,
		},
		{
			name:    "invalid id",
			args:    args{selector: testutils.StringNumberTooBigForInt},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQuery, err := serviceAlertFromSelector(tt.args.selector)
			if !tt.wantErr(t, err, fmt.Sprintf("serviceAlertFromSelector(%v)", tt.args.selector)) {
				return
			}
			assert.Equalf(t, tt.wantQuery, gotQuery, "serviceAlertFromSelector(%v)", tt.args.selector)
		})
	}
}
