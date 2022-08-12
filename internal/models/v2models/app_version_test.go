package v2models

import (
	"github.com/broadinstitute/sherlock/internal/testutils"
	"gorm.io/gorm"
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
		appVersion AppVersion
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "none",
			args: args{appVersion: AppVersion{}},
			want: nil,
		},
		{
			name: "id",
			args: args{appVersion: AppVersion{Model: gorm.Model{ID: 123}}},
			want: []string{"123"},
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
		appVersion AppVersion
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "no chart",
			args: args{appVersion: AppVersion{
				AppVersion: "1.2.3",
				GitCommit:  "abcd",
				GitBranch:  "main",
			}},
			wantErr: true,
		},
		{
			name: "no version",
			args: args{appVersion: AppVersion{
				ChartID:   123,
				GitCommit: "abcd",
				GitBranch: "main",
			}},
			wantErr: true,
		},
		{
			name: "valid",
			args: args{appVersion: AppVersion{
				ChartID:    123,
				AppVersion: "1.2.3",
				GitCommit:  "abcd",
				GitBranch:  "main",
			}},
			wantErr: false,
		},
		{
			name: "valid without commit",
			args: args{appVersion: AppVersion{
				ChartID:    123,
				AppVersion: "1.2.3",
				GitBranch:  "main",
			}},
			wantErr: false,
		},
		{
			name: "valid without branch",
			args: args{appVersion: AppVersion{
				ChartID:    123,
				AppVersion: "1.2.3",
				GitCommit:  "abcd",
			}},
			wantErr: false,
		},
		{
			name: "valid without git at all",
			args: args{appVersion: AppVersion{
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
