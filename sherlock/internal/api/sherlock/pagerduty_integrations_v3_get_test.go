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

func (s *handlerSuite) TestPagerdutyIntegrationV3Get_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/pagerduty-integrations/v3/something/with/slashes", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestPagerdutyIntegrationV3Get_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/pagerduty-integrations/v3/pd-id/blahblahblah", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestPagerdutyIntegrationV3Get() {
	pdi := s.TestData.PagerdutyIntegration_ManuallyTriggeredTerraIncident()
	var got PagerdutyIntegrationV3
	code := s.HandleRequest(
		s.NewSuitableRequest("GET", fmt.Sprintf("/api/pagerduty-integrations/v3/%d", pdi.ID), nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Equal(pdi.ID, got.ID)
}

func Test_pagerdutyIntegrationModelFromSelector(t *testing.T) {
	type args struct {
		selector string
	}
	tests := []struct {
		name      string
		args      args
		wantQuery models.PagerdutyIntegration
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name:    "empty",
			args:    args{selector: ""},
			wantErr: assert.Error,
		},
		{
			name:    "invalid",
			args:    args{selector: "something obviously invalid"},
			wantErr: assert.Error,
		},
		{
			name: "valid id",
			args: args{selector: "123"},
			wantQuery: models.PagerdutyIntegration{
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
			name: "pd-id",
			args: args{selector: "pd-id/blahblahblah"},
			wantQuery: models.PagerdutyIntegration{
				PagerdutyID: "blahblahblah",
			},
			wantErr: assert.NoError,
		},
		{
			name:    "invalid prefix",
			args:    args{selector: "something/blahblahblah"},
			wantErr: assert.Error,
		},
		{
			name:    "empty sub-selector",
			args:    args{selector: "pd-id/"},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQuery, err := pagerdutyIntegrationModelFromSelector(tt.args.selector)
			if !tt.wantErr(t, err, fmt.Sprintf("pagerdutyIntegrationModelFromSelector(%v)", tt.args.selector)) {
				return
			}
			assert.Equalf(t, tt.wantQuery, gotQuery, "pagerdutyIntegrationModelFromSelector(%v)", tt.args.selector)
		})
	}
}
