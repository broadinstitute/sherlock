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
)

func (s *handlerSuite) TestChartReleaseV3Get_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/chart-releases/v3/something/with/too/many/slashes", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChartReleaseV3Get_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/chart-releases/v3/123", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestChartReleaseV3Get() {
	chartRelease := s.TestData.ChartRelease_LeonardoDev()

	var got ChartReleaseV3
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/chart-releases/v3/"+chartRelease.Name, nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Equal(chartRelease.Name, got.Name)
}

func (s *handlerSuite) Test_chartReleaseModelFromSelector() {
	type args struct {
		selector string
	}
	tests := []struct {
		name      string
		args      args
		wantQuery models.ChartRelease
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
			wantQuery: models.ChartRelease{
				Model: gorm.Model{ID: 123},
			},
		},
		{
			name:    "invalid id",
			args:    args{selector: testutils.StringNumberTooBigForInt},
			wantErr: assert.Error,
		},
		{
			name:    "environment + chart",
			args:    args{selector: s.TestData.Environment_Prod().Name + "/" + s.TestData.Chart_Leonardo().Name},
			wantErr: assert.NoError,
			wantQuery: models.ChartRelease{
				EnvironmentID: utils.PointerTo(s.TestData.Environment_Prod().ID),
				ChartID:       s.TestData.Chart_Leonardo().ID,
			},
		},
		{
			name:    "environment + chart, empty environment",
			args:    args{selector: "/" + s.TestData.Chart_Leonardo().Name},
			wantErr: assert.Error,
		},
		{
			name:    "environment + chart, invalid environment",
			args:    args{selector: "!!!!!!!/" + s.TestData.Chart_Leonardo().Name},
			wantErr: assert.Error,
		},
		{
			name:    "environment + chart, not found environment",
			args:    args{selector: "not-found/" + s.TestData.Chart_Leonardo().Name},
			wantErr: assert.Error,
		},
		{
			name:    "environment + chart, empty chart",
			args:    args{selector: s.TestData.Environment_Prod().Name + "/"},
			wantErr: assert.Error,
		},
		{
			name:    "environment + chart, invalid chart",
			args:    args{selector: s.TestData.Environment_Prod().Name + "/!!!!!!!"},
			wantErr: assert.Error,
		},
		{
			name:    "environment + chart, not found chart",
			args:    args{selector: s.TestData.Environment_Prod().Name + "/not-found"},
			wantErr: assert.Error,
		},
		{
			name: "cluster + namespace + chart",
			args: args{
				selector: s.TestData.Cluster_TerraProd().Name + "/" + "terra-prod" + "/" + s.TestData.Chart_Leonardo().Name,
			},
			wantErr: assert.NoError,
			wantQuery: models.ChartRelease{
				ClusterID: utils.PointerTo(s.TestData.Cluster_TerraProd().ID),
				Namespace: "terra-prod",
				ChartID:   s.TestData.Chart_Leonardo().ID,
			},
		},
		{
			name: "cluster + namespace + chart, empty cluster",
			args: args{
				selector: "/" + "terra-prod" + "/" + s.TestData.Chart_Leonardo().Name,
			},
			wantErr: assert.Error,
		},
		{
			name: "cluster + namespace + chart, invalid cluster",
			args: args{
				selector: "!!!!!!/" + "terra-prod" + "/" + s.TestData.Chart_Leonardo().Name,
			},
			wantErr: assert.Error,
		},
		{
			name: "cluster + namespace + chart, not found cluster",
			args: args{
				selector: "not-found/" + "terra-prod" + "/" + s.TestData.Chart_Leonardo().Name,
			},
			wantErr: assert.Error,
		},
		{
			name: "cluster + namespace + chart, empty namespace",
			args: args{
				selector: s.TestData.Cluster_TerraProd().Name + "/" + "/" + s.TestData.Chart_Leonardo().Name,
			},
			wantErr: assert.Error,
		},
		{
			name: "cluster + namespace + chart, invalid namespace",
			args: args{
				selector: s.TestData.Cluster_TerraProd().Name + "/" + "!!!!!!" + "/" + s.TestData.Chart_Leonardo().Name,
			},
			wantErr: assert.Error,
		},
		{
			name: "cluster + namespace + chart, empty chart",
			args: args{
				selector: s.TestData.Cluster_TerraProd().Name + "/" + "terra-prod" + "/",
			},
			wantErr: assert.Error,
		},
		{
			name: "cluster + namespace + chart, invalid chart",
			args: args{
				selector: s.TestData.Cluster_TerraProd().Name + "/" + "terra-prod" + "/!!!!!!",
			},
			wantErr: assert.Error,
		},
		{
			name: "cluster + namespace + chart, not found chart",
			args: args{
				selector: s.TestData.Cluster_TerraProd().Name + "/" + "terra-prod" + "/not-found",
			},
			wantErr: assert.Error,
		},
		{
			name:    "name",
			args:    args{selector: "a-name"},
			wantErr: assert.NoError,
			wantQuery: models.ChartRelease{
				Name: "a-name",
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			gotQuery, err := chartReleaseModelFromSelector(s.DB, tt.args.selector)
			if !tt.wantErr(s.T(), err, fmt.Sprintf("chartReleaseModelFromSelector(%v)", tt.args.selector)) {
				return
			}
			s.Equalf(tt.wantQuery, gotQuery, "chartReleaseModelFromSelector(%v)", tt.args.selector)
		})
	}
}
