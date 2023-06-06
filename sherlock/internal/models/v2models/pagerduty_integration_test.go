package v2models

import (
	"github.com/broadinstitute/sherlock/internal/testutils"
	"gorm.io/gorm"
	"testing"
)

func Test_pagerdutyIntegrationSelectorToQuery(t *testing.T) {
	type args struct {
		db       *gorm.DB
		selector string
	}
	tests := []struct {
		name    string
		args    args
		want    PagerdutyIntegration
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
			want: PagerdutyIntegration{Model: gorm.Model{ID: 123}},
		},
		{
			name:    "invalid id",
			args:    args{selector: testutils.StringNumberTooBigForInt},
			wantErr: true,
		},
		{
			name: "valid pd id",
			args: args{selector: "pd-id/a1B2c3d"},
			want: PagerdutyIntegration{PagerdutyID: "a1B2c3d"},
		},
		{
			name:    "invalid pd id prefix",
			args:    args{selector: "abc/def"},
			wantErr: true,
		},
		{
			name:    "invalid pd id suffix",
			args:    args{selector: "pd-id/"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := pagerdutyIntegrationSelectorToQuery(tt.args.db, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("pagerdutyIntegrationSelectorToQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

func Test_pagerdutyIntegrationToSelectors(t *testing.T) {
	type args struct {
		pagerdutyIntegration *PagerdutyIntegration
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "none",
			args: args{pagerdutyIntegration: &PagerdutyIntegration{}},
			want: nil,
		},
		{
			name: "id",
			args: args{pagerdutyIntegration: &PagerdutyIntegration{Model: gorm.Model{ID: 123}}},
			want: []string{"123"},
		},
		{
			name: "pd id",
			args: args{pagerdutyIntegration: &PagerdutyIntegration{PagerdutyID: "a1B2c3d"}},
			want: []string{"pd-id/a1B2c3d"},
		},
		{
			name: "id and pd id",
			args: args{pagerdutyIntegration: &PagerdutyIntegration{Model: gorm.Model{ID: 123}, PagerdutyID: "a1B2c3d"}},
			want: []string{"123", "pd-id/a1B2c3d"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pagerdutyIntegrationToSelectors(tt.args.pagerdutyIntegration)
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

func Test_validatePagerdutyIntegration(t *testing.T) {
	type args struct {
		pagerdutyIntegration *PagerdutyIntegration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid with type",
			args: args{pagerdutyIntegration: &PagerdutyIntegration{
				PagerdutyID: "a1B2c3d",
				Name:        testutils.PointerTo("foo"),
				Key:         testutils.PointerTo("foo"),
				Type:        testutils.PointerTo("service"),
			}},
			wantErr: false,
		},
		{
			name: "no type",
			args: args{pagerdutyIntegration: &PagerdutyIntegration{
				PagerdutyID: "a1B2c3d",
				Name:        testutils.PointerTo("foo"),
				Key:         testutils.PointerTo("foo"),
			}},
			wantErr: true,
		},
		{
			name: "no pd id",
			args: args{pagerdutyIntegration: &PagerdutyIntegration{
				Name: testutils.PointerTo("foo"),
				Key:  testutils.PointerTo("foo"),
			}},
			wantErr: true,
		},
		{
			name: "no name",
			args: args{pagerdutyIntegration: &PagerdutyIntegration{
				PagerdutyID: "a1B2c3d",
				Key:         testutils.PointerTo("foo"),
			}},
			wantErr: true,
		},
		{
			name: "no key",
			args: args{pagerdutyIntegration: &PagerdutyIntegration{
				PagerdutyID: "a1B2c3d",
				Name:        testutils.PointerTo("foo"),
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validatePagerdutyIntegration(tt.args.pagerdutyIntegration); (err != nil) != tt.wantErr {
				t.Errorf("validatePagerdutyIntegration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
