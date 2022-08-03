package v2models

import (
	"github.com/broadinstitute/sherlock/internal/testutils"
	"gorm.io/gorm"
	"testing"
)

func Test_chartDeployRecordSelectorToQuery(t *testing.T) {
	type args struct {
		db       *gorm.DB
		selector string
	}
	tests := []struct {
		name    string
		args    args
		want    ChartDeployRecord
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
			want: ChartDeployRecord{Model: gorm.Model{ID: 123}},
		},
		{
			name:    "invalid id",
			args:    args{selector: testutils.StringNumberTooBigForInt},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := chartDeployRecordSelectorToQuery(tt.args.db, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("chartDeployRecordSelectorToQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

func Test_chartDeployRecordToSelectors(t *testing.T) {
	type args struct {
		chartDeployRecord ChartDeployRecord
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "none",
			args: args{chartDeployRecord: ChartDeployRecord{}},
			want: nil,
		},
		{
			name: "id",
			args: args{chartDeployRecord: ChartDeployRecord{Model: gorm.Model{ID: 123}}},
			want: []string{"123"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := chartDeployRecordToSelectors(tt.args.chartDeployRecord)
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

func Test_validateChartDeployRecord(t *testing.T) {
	type args struct {
		chartDeployRecord ChartDeployRecord
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "no chart release",
			args: args{chartDeployRecord: ChartDeployRecord{
				ExactChartVersion: "1.2.3",
				ExactAppVersion:   "4.5.6",
				HelmfileRef:       "a1b2c3d4",
			}},
			wantErr: true,
		},
		{
			name: "no chart version",
			args: args{chartDeployRecord: ChartDeployRecord{
				ChartReleaseID:  123,
				ExactAppVersion: "4.5.6",
				HelmfileRef:     "a1b2c3d4",
			}},
			wantErr: true,
		},
		{
			name: "no app version",
			args: args{chartDeployRecord: ChartDeployRecord{
				ChartReleaseID:    123,
				ExactChartVersion: "1.2.3",
				HelmfileRef:       "a1b2c3d4",
			}},
			wantErr: true,
		},
		{
			name: "no helmfile ref",
			args: args{chartDeployRecord: ChartDeployRecord{
				ChartReleaseID:    123,
				ExactChartVersion: "1.2.3",
				ExactAppVersion:   "4.5.6",
			}},
			wantErr: true,
		},
		{
			name: "valid",
			args: args{chartDeployRecord: ChartDeployRecord{
				ChartReleaseID:    123,
				ExactChartVersion: "1.2.3",
				ExactAppVersion:   "4.5.6",
				HelmfileRef:       "a1b2c3d4",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateChartDeployRecord(tt.args.chartDeployRecord); (err != nil) != tt.wantErr {
				t.Errorf("validateChartDeployRecord() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
