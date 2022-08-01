package v2models

import (
	"github.com/broadinstitute/sherlock/internal/testutils"
	"gorm.io/gorm"
	"testing"
)

func Test_clusterToSelectors(t *testing.T) {
	type args struct {
		cluster Cluster
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "empty",
			args: args{},
			want: nil,
		},
		{
			name: "with name",
			args: args{cluster: Cluster{
				Name: "foobar",
			}},
			want: []string{"foobar"},
		},
		{
			name: "with id",
			args: args{cluster: Cluster{
				Model: gorm.Model{ID: 123},
			}},
			want: []string{"123"},
		},
		{
			name: "with name and id",
			args: args{cluster: Cluster{
				Model: gorm.Model{ID: 123},
				Name:  "foobar",
			}},
			want: []string{"foobar", "123"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := clusterToSelectors(tt.args.cluster)
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

func Test_clusterRequiresSuitability(t *testing.T) {
	tru := true
	fal := false
	type args struct {
		cluster Cluster
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "does require",
			args: args{cluster: Cluster{RequiresSuitability: &tru}},
			want: true,
		},
		{
			name: "does not require",
			args: args{cluster: Cluster{RequiresSuitability: &fal}},
			want: false,
		},
		{
			name: "fail safe",
			args: args{cluster: Cluster{}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := clusterRequiresSuitability(tt.args.cluster)
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

func Test_validateCluster(t *testing.T) {
	type args struct {
		cluster Cluster
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "missing name",
			args: args{cluster: Cluster{
				Provider:            "google",
				GoogleProject:       "broad-dsde-dev",
				Base:                testutils.PointerTo("live"),
				Address:             testutils.PointerTo("1.2.3.4"),
				RequiresSuitability: testutils.PointerTo(false),
			}},
			wantErr: true,
		},
		{
			name: "google provider but no project",
			args: args{cluster: Cluster{
				Name:                "terra-dev",
				Provider:            "google",
				Base:                testutils.PointerTo("live"),
				Address:             testutils.PointerTo("1.2.3.4"),
				RequiresSuitability: testutils.PointerTo(false),
			}},
			wantErr: true,
		},
		{
			name: "azure provider but no subscription",
			args: args{cluster: Cluster{
				Name:                "terra-dev",
				Provider:            "azure",
				Base:                testutils.PointerTo("live"),
				Address:             testutils.PointerTo("1.2.3.4"),
				RequiresSuitability: testutils.PointerTo(false),
			}},
			wantErr: true,
		},
		{
			name: "invalid provider",
			args: args{cluster: Cluster{
				Name:                "terra-dev",
				Provider:            "foo-bar",
				GoogleProject:       "broad-dsde-dev",
				AzureSubscription:   "",
				Base:                testutils.PointerTo("live"),
				Address:             testutils.PointerTo("1.2.3.4"),
				RequiresSuitability: testutils.PointerTo(false),
			}},
			wantErr: true,
		},
		{
			name: "no base",
			args: args{cluster: Cluster{
				Name:                "terra-dev",
				Provider:            "google",
				GoogleProject:       "broad-dsde-dev",
				Address:             testutils.PointerTo("1.2.3.4"),
				RequiresSuitability: testutils.PointerTo(false),
			}},
			wantErr: true,
		},
		{
			name: "no address",
			args: args{cluster: Cluster{
				Name:                "terra-dev",
				Provider:            "google",
				GoogleProject:       "broad-dsde-dev",
				Base:                testutils.PointerTo("live"),
				RequiresSuitability: testutils.PointerTo(false),
			}},
			wantErr: true,
		},
		{
			name: "no requires suitability",
			args: args{cluster: Cluster{
				Name:          "terra-dev",
				Provider:      "google",
				GoogleProject: "broad-dsde-dev",
				Base:          testutils.PointerTo("live"),
				Address:       testutils.PointerTo("1.2.3.4"),
			}},
			wantErr: true,
		},
		{
			name: "valid with google",
			args: args{cluster: Cluster{
				Name:                "terra-dev",
				Provider:            "google",
				GoogleProject:       "broad-dsde-dev",
				Base:                testutils.PointerTo("live"),
				Address:             testutils.PointerTo("1.2.3.4"),
				RequiresSuitability: testutils.PointerTo(false),
			}},
			wantErr: false,
		},
		{
			name: "valid with azure",
			args: args{cluster: Cluster{
				Name:                "terra-dev",
				Provider:            "azure",
				AzureSubscription:   "some uuid probably",
				Base:                testutils.PointerTo("live"),
				Address:             testutils.PointerTo("1.2.3.4"),
				RequiresSuitability: testutils.PointerTo(false),
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateCluster(tt.args.cluster); (err != nil) != tt.wantErr {
				t.Errorf("validateCluster() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
