package v2models

import (
	"github.com/broadinstitute/sherlock/internal/testutils"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func Test_appVersionSelectorToQuery(t *testing.T) {
	type args struct {
		db       *gorm.DB
		selector string
	}
	tests := []struct {
		name    string
		args    args
		want    AppVersion
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
			want: AppVersion{Model: gorm.Model{ID: 123}},
		},
		{
			name:    "invalid id",
			args:    args{selector: testutils.StringNumberTooBigForInt},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := appVersionSelectorToQuery(tt.args.db, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("appVersionSelectorToQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

func Test_appVersionToSelectors(t *testing.T) {
	type args struct {
		appVersion *AppVersion
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "none",
			args: args{appVersion: &AppVersion{}},
			want: nil,
		},
		{
			name: "id",
			args: args{appVersion: &AppVersion{Model: gorm.Model{ID: 123}}},
			want: []string{"123"},
		},
		{
			name: "id and chart + version",
			args: args{appVersion: &AppVersion{
				Model: gorm.Model{ID: 123},
				Chart: &Chart{
					Model: gorm.Model{ID: 456},
					Name:  "datarepo",
				},
				AppVersion: "1.2.3",
			}},
			want: []string{"123", "datarepo/1.2.3", "456/1.2.3"},
		},
		{
			name: "id and chart id + version",
			args: args{appVersion: &AppVersion{
				Model:      gorm.Model{ID: 123},
				ChartID:    456,
				AppVersion: "1.2.3",
			}},
			want: []string{"123", "456/1.2.3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := appVersionToSelectors(tt.args.appVersion)
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

func Test_validateAppVersion(t *testing.T) {
	type args struct {
		appVersion *AppVersion
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "no chart",
			args: args{appVersion: &AppVersion{
				AppVersion: "1.2.3",
				GitCommit:  "abcd",
				GitBranch:  "main",
			}},
			wantErr: true,
		},
		{
			name: "no version",
			args: args{appVersion: &AppVersion{
				ChartID:   123,
				GitCommit: "abcd",
				GitBranch: "main",
			}},
			wantErr: true,
		},
		{
			name: "valid",
			args: args{appVersion: &AppVersion{
				ChartID:    123,
				AppVersion: "1.2.3",
				GitCommit:  "abcd",
				GitBranch:  "main",
			}},
			wantErr: false,
		},
		{
			name: "valid without commit",
			args: args{appVersion: &AppVersion{
				ChartID:    123,
				AppVersion: "1.2.3",
				GitBranch:  "main",
			}},
			wantErr: false,
		},
		{
			name: "valid without branch",
			args: args{appVersion: &AppVersion{
				ChartID:    123,
				AppVersion: "1.2.3",
				GitCommit:  "abcd",
			}},
			wantErr: false,
		},
		{
			name: "valid without git at all",
			args: args{appVersion: &AppVersion{
				ChartID:    123,
				AppVersion: "1.2.3",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateAppVersion(tt.args.appVersion); (err != nil) != tt.wantErr {
				t.Errorf("validateAppVersion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_rejectDuplicateAppVersion(t *testing.T) {
	type args struct {
		existing *AppVersion
		new      *AppVersion
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "mismatch chart version",
			args: args{
				existing: &AppVersion{AppVersion: "123"},
				new:      &AppVersion{AppVersion: "456"},
			},
			wantErr: true,
		},
		{
			name: "mismatch chart",
			args: args{
				existing: &AppVersion{ChartID: 123},
				new:      &AppVersion{ChartID: 456},
			},
			wantErr: true,
		},
		{
			name: "mismatch git branch",
			args: args{
				existing: &AppVersion{GitBranch: "123"},
				new:      &AppVersion{GitBranch: "456"},
			},
			wantErr: true,
		},
		{
			name: "mismatch git commit",
			args: args{
				existing: &AppVersion{GitCommit: "123"},
				new:      &AppVersion{GitCommit: "456"},
			},
			wantErr: true,
		},
		{
			name: "only existing has parent",
			args: args{
				existing: &AppVersion{ParentAppVersionID: testutils.PointerTo[uint](123)},
				new:      &AppVersion{},
			},
			wantErr: true,
		},
		{
			name: "only new has parent",
			args: args{
				existing: &AppVersion{},
				new:      &AppVersion{ParentAppVersionID: testutils.PointerTo[uint](456)},
			},
			wantErr: true,
		},
		{
			name: "mismatch parent",
			args: args{
				existing: &AppVersion{ParentAppVersionID: testutils.PointerTo[uint](123)},
				new:      &AppVersion{ParentAppVersionID: testutils.PointerTo[uint](456)},
			},
			wantErr: true,
		},
		{
			name: "no mismatch",
			args: args{
				existing: &AppVersion{},
				new:      &AppVersion{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := rejectDuplicateAppVersion(tt.args.existing, tt.args.new); (err != nil) != tt.wantErr {
				t.Errorf("rejectDuplicateAppVersion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAppVersion_GetCiIdentifier(t *testing.T) {
	type fields struct {
		Model              gorm.Model
		CiIdentifier       *CiIdentifier
		Chart              *Chart
		ChartID            uint
		AppVersion         string
		GitCommit          string
		GitBranch          string
		Description        string
		ParentAppVersion   *AppVersion
		ParentAppVersionID *uint
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
				ResourceType: "app-version",
				ResourceID:   123,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AppVersion{
				Model:              tt.fields.Model,
				CiIdentifier:       tt.fields.CiIdentifier,
				Chart:              tt.fields.Chart,
				ChartID:            tt.fields.ChartID,
				AppVersion:         tt.fields.AppVersion,
				GitCommit:          tt.fields.GitCommit,
				GitBranch:          tt.fields.GitBranch,
				Description:        tt.fields.Description,
				ParentAppVersion:   tt.fields.ParentAppVersion,
				ParentAppVersionID: tt.fields.ParentAppVersionID,
			}
			if got := a.GetCiIdentifier(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCiIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}
