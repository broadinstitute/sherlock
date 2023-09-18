package v2models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func Test_chartSelectorToQuery(t *testing.T) {
	type args struct {
		db       *gorm.DB
		selector string
	}
	tests := []struct {
		name    string
		args    args
		want    Chart
		wantErr bool
	}{
		{
			name:    "empty",
			args:    args{selector: ""},
			wantErr: true,
		},
		{
			name:    "invalid",
			args:    args{selector: "something obviously invalid!"},
			wantErr: true,
		},
		{
			name: "valid id",
			args: args{selector: "123"},
			want: Chart{Model: gorm.Model{ID: 123}},
		},
		{
			name:    "invalid id",
			args:    args{selector: testutils.StringNumberTooBigForInt},
			wantErr: true,
		},
		{
			name: "valid name",
			args: args{selector: "foo-bar-2"},
			want: Chart{Name: "foo-bar-2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := chartSelectorToQuery(tt.args.db, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("chartSelectorToQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

func Test_chartToSelectors(t *testing.T) {
	type args struct {
		chart *Chart
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "none",
			args: args{chart: &Chart{}},
			want: nil,
		},
		{
			name: "name",
			args: args{chart: &Chart{Name: "foobar"}},
			want: []string{"foobar"},
		},
		{
			name: "id",
			args: args{chart: &Chart{Model: gorm.Model{ID: 123}}},
			want: []string{"123"},
		},
		{
			name: "name and id",
			args: args{chart: &Chart{Name: "foobar", Model: gorm.Model{ID: 123}}},
			want: []string{"foobar", "123"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := chartToSelectors(tt.args.chart)
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

func Test_validateChart(t *testing.T) {
	type args struct {
		chart *Chart
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "no name",
			args: args{chart: &Chart{
				ChartRepo:             utils.PointerTo("terra-helm"),
				AppImageGitRepo:       utils.PointerTo("broadinstitute/leonardo"),
				AppImageGitMainBranch: utils.PointerTo("main"),
			}},
			wantErr: true,
		},
		{
			name: "no chart repo",
			args: args{chart: &Chart{
				Name:                  "leonardo",
				AppImageGitRepo:       utils.PointerTo("broadinstitute/leonardo"),
				AppImageGitMainBranch: utils.PointerTo("main"),
			}},
			wantErr: true,
		},
		{
			name: "valid with git info",
			args: args{chart: &Chart{
				Name:                  "leonardo",
				ChartRepo:             utils.PointerTo("terra-helm"),
				AppImageGitRepo:       utils.PointerTo("broadinstitute/leonardo"),
				AppImageGitMainBranch: utils.PointerTo("main"),
			}},
			wantErr: false,
		},
		{
			name: "valid without git info",
			args: args{chart: &Chart{
				Name:      "leonardo",
				ChartRepo: utils.PointerTo("terra-helm"),
			}},
			wantErr: false,
		},
		{
			name: "valid with legacy configs enabled",
			args: args{chart: &Chart{
				Name:                 "leonardo",
				ChartRepo:            utils.PointerTo("terra-helm"),
				LegacyConfigsEnabled: utils.PointerTo(true),
			}},
			wantErr: false,
		},
		{
			name: "valid without legacy configs set",
			args: args{chart: &Chart{
				Name:      "leonardo",
				ChartRepo: utils.PointerTo("terra-helm"),
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateChart(tt.args.chart); (err != nil) != tt.wantErr {
				t.Errorf("validateChart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChart_GetCiIdentifier(t *testing.T) {
	type fields struct {
		Model                 gorm.Model
		CiIdentifier          *CiIdentifier
		Name                  string
		ChartRepo             *string
		AppImageGitRepo       *string
		AppImageGitMainBranch *string
		ChartExposesEndpoint  *bool
		DefaultSubdomain      *string
		DefaultProtocol       *string
		DefaultPort           *uint
		LegacyConfigsEnabled  *bool
		Description           *string
		PlaybookURL           *string
	}
	tests := []struct {
		name   string
		fields fields
		want   *CiIdentifier
	}{
		{
			name: "returns existing",
			fields: fields{
				CiIdentifier: &CiIdentifier{
					Model: gorm.Model{
						ID: 123,
					},
				},
			},
			want: &CiIdentifier{
				Model: gorm.Model{
					ID: 123,
				},
			},
		},
		{
			name: "returns generated if no existing",
			fields: fields{
				Model: gorm.Model{
					ID: 123,
				},
			},
			want: &CiIdentifier{
				ResourceType: "chart",
				ResourceID:   123,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Chart{
				Model:                 tt.fields.Model,
				CiIdentifier:          tt.fields.CiIdentifier,
				Name:                  tt.fields.Name,
				ChartRepo:             tt.fields.ChartRepo,
				AppImageGitRepo:       tt.fields.AppImageGitRepo,
				AppImageGitMainBranch: tt.fields.AppImageGitMainBranch,
				ChartExposesEndpoint:  tt.fields.ChartExposesEndpoint,
				DefaultSubdomain:      tt.fields.DefaultSubdomain,
				DefaultProtocol:       tt.fields.DefaultProtocol,
				DefaultPort:           tt.fields.DefaultPort,
				LegacyConfigsEnabled:  tt.fields.LegacyConfigsEnabled,
				Description:           tt.fields.Description,
				PlaybookURL:           tt.fields.PlaybookURL,
			}
			if got := c.GetCiIdentifier(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCiIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}
