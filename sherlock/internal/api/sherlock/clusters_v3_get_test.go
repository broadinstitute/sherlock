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

func (s *handlerSuite) TestClusterV3Get_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/clusters/v3/something/with/slashes", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestClusterV3Get_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/clusters/v3/my-cluster", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestClusterV3Get() {
	s.SetNonSuitableTestUserForDB()
	s.NoError(s.DB.Create(&models.Cluster{
		Name:                "some-name",
		Provider:            "azure",
		AzureSubscription:   "some-subscription",
		Location:            "some-location",
		Base:                utils.PointerTo("some base"),
		Address:             utils.PointerTo("0.0.0.0"),
		RequiresSuitability: utils.PointerTo(false),
		HelmfileRef:         utils.PointerTo("some-ref"),
	}).Error)

	var got ClusterV3
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/clusters/v3/some-name", nil),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.Base) {
		s.Equal("some base", *got.Base)
	}
}

func Test_clusterModelFromSelector(t *testing.T) {
	type args struct {
		selector string
	}
	tests := []struct {
		name      string
		args      args
		wantQuery models.Cluster
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
			wantQuery: models.Cluster{Model: gorm.Model{ID: 123}},
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
			wantQuery: models.Cluster{Name: "foo-bar-2"},
			wantErr:   assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQuery, err := clusterModelFromSelector(tt.args.selector)
			if !tt.wantErr(t, err, fmt.Sprintf("clusterModelFromSelector(%v)", tt.args.selector)) {
				return
			}
			assert.Equalf(t, tt.wantQuery, gotQuery, "clusterModelFromSelector(%v)", tt.args.selector)
		})
	}
}
