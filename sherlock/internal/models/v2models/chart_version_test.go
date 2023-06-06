package v2models

import (
	"github.com/broadinstitute/sherlock/internal/testutils"
	"gorm.io/gorm"
	"reflect"
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
		chartVersion *ChartVersion
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "none",
			args: args{chartVersion: &ChartVersion{}},
			want: nil,
		},
		{
			name: "id",
			args: args{chartVersion: &ChartVersion{Model: gorm.Model{ID: 123}}},
			want: []string{"123"},
		},
		{
			name: "id and chart + version",
			args: args{chartVersion: &ChartVersion{
				Model: gorm.Model{ID: 123},
				Chart: &Chart{
					Model: gorm.Model{ID: 456},
					Name:  "datarepo",
				},
				ChartVersion: "1.2.3",
			}},
			want: []string{"123", "datarepo/1.2.3", "456/1.2.3"},
		},
		{
			name: "id and chart id + version",
			args: args{chartVersion: &ChartVersion{
				Model:        gorm.Model{ID: 123},
				ChartID:      456,
				ChartVersion: "1.2.3",
			}},
			want: []string{"123", "456/1.2.3"},
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
		chartVersion *ChartVersion
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "no chart",
			args: args{chartVersion: &ChartVersion{
				ChartVersion: "1.2.3",
			}},
			wantErr: true,
		},
		{
			name: "no version",
			args: args{chartVersion: &ChartVersion{
				ChartID: 123,
			}},
			wantErr: true,
		},
		{
			name: "valid",
			args: args{chartVersion: &ChartVersion{
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

func Test_rejectDuplicateChartVersion(t *testing.T) {
	type args struct {
		existing *ChartVersion
		new      *ChartVersion
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "mismatch chart version",
			args: args{
				existing: &ChartVersion{ChartVersion: "123"},
				new:      &ChartVersion{ChartVersion: "456"},
			},
			wantErr: true,
		},
		{
			name: "mismatch chart",
			args: args{
				existing: &ChartVersion{ChartID: 123},
				new:      &ChartVersion{ChartID: 456},
			},
			wantErr: true,
		},
		{
			name: "only existing has parent",
			args: args{
				existing: &ChartVersion{ParentChartVersionID: testutils.PointerTo[uint](123)},
				new:      &ChartVersion{},
			},
			wantErr: true,
		},
		{
			name: "only new has parent",
			args: args{
				existing: &ChartVersion{},
				new:      &ChartVersion{ParentChartVersionID: testutils.PointerTo[uint](456)},
			},
			wantErr: true,
		},
		{
			name: "mismatch parent",
			args: args{
				existing: &ChartVersion{ParentChartVersionID: testutils.PointerTo[uint](123)},
				new:      &ChartVersion{ParentChartVersionID: testutils.PointerTo[uint](456)},
			},
			wantErr: true,
		},
		{
			name: "no mismatch",
			args: args{
				existing: &ChartVersion{},
				new:      &ChartVersion{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := rejectDuplicateChartVersion(tt.args.existing, tt.args.new); (err != nil) != tt.wantErr {
				t.Errorf("rejectDuplicateChartVersion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChartVersion_GetCiIdentifier(t *testing.T) {
	type fields struct {
		Model                gorm.Model
		CiIdentifier         *CiIdentifier
		Chart                *Chart
		ChartID              uint
		ChartVersion         string
		Description          string
		ParentChartVersion   *ChartVersion
		ParentChartVersionID *uint
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
				ResourceType: "chart-version",
				ResourceID:   123,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ChartVersion{
				Model:                tt.fields.Model,
				CiIdentifier:         tt.fields.CiIdentifier,
				Chart:                tt.fields.Chart,
				ChartID:              tt.fields.ChartID,
				ChartVersion:         tt.fields.ChartVersion,
				Description:          tt.fields.Description,
				ParentChartVersion:   tt.fields.ParentChartVersion,
				ParentChartVersionID: tt.fields.ParentChartVersionID,
			}
			if got := c.GetCiIdentifier(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCiIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}
