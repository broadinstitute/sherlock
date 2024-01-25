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

func (s *handlerSuite) TestIncidentV3Get_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/incidents/v3/something/with/slashes", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestIncidentV3Get_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/incidents/v3/123", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestIncidentV3Get() {
	s.SetNonSuitableTestUserForDB()
	incident1 := s.TestData.Incident_1()

	var got IncidentV3
	code := s.HandleRequest(
		s.NewRequest("GET", fmt.Sprintf("/api/incidents/v3/%d", incident1.ID), nil),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.Description) {
		s.Equal(*incident1.Description, *got.Description)
	}
}

func Test_incidentModelFromSelector(t *testing.T) {
	type args struct {
		selector string
	}
	tests := []struct {
		name      string
		args      args
		wantQuery models.Incident
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
			wantQuery: models.Incident{Model: gorm.Model{ID: 123}},
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
			gotQuery, err := incidentModelFromSelector(tt.args.selector)
			if !tt.wantErr(t, err, fmt.Sprintf("incidentModelFromSelector(%v)", tt.args.selector)) {
				return
			}
			assert.Equalf(t, tt.wantQuery, gotQuery, "incidentModelFromSelector(%v)", tt.args.selector)
		})
	}
}
