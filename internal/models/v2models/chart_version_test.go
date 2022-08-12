package v2models

import (
	"github.com/broadinstitute/sherlock/internal/testutils"
	"gorm.io/gorm"
	"testing"
)

func Test_chartVersionSelectorToQuery(t *testing.T) {
	type args struct {
		db       *gorm.DB
		selector string
	}
	tests := []struct {
		name    string
		args    args
		want    ChartVersion
		wantErr bool
	}{
		{
			name:    "empty",
			args:    args{selector: ""},
			wantErr: true,
		},
		{
			name:    "invalid",
			args:    args{selector: "noAlphanumericSelectorsAllowed"},
			wantErr: true,
		},
		{
			name: "valid id",
			args: args{selector: "123"},
			want: ChartVersion{Model: gorm.Model{ID: 123}},
		},
		{
			name:    "invalid id",
			args:    args{selector: testutils.StringNumberTooBigForInt},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := chartVersionSelectorToQuery(tt.args.db, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("chartVersionSelectorToQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

func Test_chartVersionToSelectors(t *testing.T) {
	type args struct {
		chartVersion ChartVersion
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "none",
			args: args{chartVersion: ChartVersion{}},
			want: nil,
		},
		{
			name: "id",
			args: args{chartVersion: ChartVersion{Model: gorm.Model{ID: 123}}},
			want: []string{"123"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := chartVersionToSelectors(tt.args.chartVersion)
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

func Test_validateChartVersion(t *testing.T) {
	type args struct {
		chartVersion ChartVersion
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "no chart",
			args: args{chartVersion: ChartVersion{
				ChartVersion: "1.2.3",
			}},
			wantErr: true,
		},
		{
			name: "no version",
			args: args{chartVersion: ChartVersion{
				ChartID: 123,
			}},
			wantErr: true,
		},
		{
			name: "valid",
			args: args{chartVersion: ChartVersion{
				ChartID:      123,
				ChartVersion: "1.2.3",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateChartVersion(tt.args.chartVersion); (err != nil) != tt.wantErr {
				t.Errorf("validateChartVersion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
