package v2models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"testing"

	"gorm.io/gorm"
)

func Test_clusterSelectorToQuery(t *testing.T) {
	type args struct {
		db       *gorm.DB
		selector string
	}
	tests := []struct {
		name    string
		args    args
		want    Cluster
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
			want: Cluster{Model: gorm.Model{ID: 123}},
		},
		{
			name:    "invalid id",
			args:    args{selector: testutils.StringNumberTooBigForInt},
			wantErr: true,
		},
		{
			name: "valid name",
			args: args{selector: "foo-bar-2"},
			want: Cluster{Name: "foo-bar-2"},
		},
		{
			name:    "invalid name",
			args:    args{selector: "foooooooooooooooooooooooooooooooo"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := clusterSelectorToQuery(tt.args.db, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("clusterSelectorToQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

func Test_clusterToSelectors(t *testing.T) {
	type args struct {
		cluster *Cluster
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
			args: args{cluster: &Cluster{
				Name: "foobar",
			}},
			want: []string{"foobar"},
		},
		{
			name: "with id",
			args: args{cluster: &Cluster{
				Model: gorm.Model{ID: 123},
			}},
			want: []string{"123"},
		},
		{
			name: "with name and id",
			args: args{cluster: &Cluster{
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

func Test_validateCluster(t *testing.T) {
	type args struct {
		cluster *Cluster
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "missing name",
			args: args{cluster: &Cluster{
				Provider:            "google",
				GoogleProject:       "broad-dsde-dev",
				Base:                testutils.PointerTo("live"),
				Address:             testutils.PointerTo("1.2.3.4"),
				RequiresSuitability: testutils.PointerTo(false),
				Location:            "us-central1-a",
				HelmfileRef:         testutils.PointerTo("HEAD"),
			}},
			wantErr: true,
		},
		{
			name: "google provider but no project",
			args: args{cluster: &Cluster{
				Name:                "terra-dev",
				Provider:            "google",
				Base:                testutils.PointerTo("live"),
				Address:             testutils.PointerTo("1.2.3.4"),
				RequiresSuitability: testutils.PointerTo(false),
				Location:            "us-central1-a",
				HelmfileRef:         testutils.PointerTo("HEAD"),
			}},
			wantErr: true,
		},
		{
			name: "azure provider but no subscription",
			args: args{cluster: &Cluster{
				Name:                "terra-dev",
				Provider:            "azure",
				Base:                testutils.PointerTo("live"),
				Address:             testutils.PointerTo("1.2.3.4"),
				RequiresSuitability: testutils.PointerTo(false),
				Location:            "US-EAST",
				HelmfileRef:         testutils.PointerTo("HEAD"),
			}},
			wantErr: true,
		},
		{
			name: "invalid provider",
			args: args{cluster: &Cluster{
				Name:                "terra-dev",
				Provider:            "foo-bar",
				GoogleProject:       "broad-dsde-dev",
				AzureSubscription:   "",
				Base:                testutils.PointerTo("live"),
				Address:             testutils.PointerTo("1.2.3.4"),
				RequiresSuitability: testutils.PointerTo(false),
				Location:            "us-central1-a",
				HelmfileRef:         testutils.PointerTo("HEAD"),
			}},
			wantErr: true,
		},
		{
			name: "no base",
			args: args{cluster: &Cluster{
				Name:                "terra-dev",
				Provider:            "google",
				GoogleProject:       "broad-dsde-dev",
				Address:             testutils.PointerTo("1.2.3.4"),
				RequiresSuitability: testutils.PointerTo(false),
				Location:            "us-central1-a",
				HelmfileRef:         testutils.PointerTo("HEAD"),
			}},
			wantErr: true,
		},
		{
			name: "no address",
			args: args{cluster: &Cluster{
				Name:                "terra-dev",
				Provider:            "google",
				GoogleProject:       "broad-dsde-dev",
				Base:                testutils.PointerTo("live"),
				RequiresSuitability: testutils.PointerTo(false),
				Location:            "us-central1-a",
				HelmfileRef:         testutils.PointerTo("HEAD"),
			}},
			wantErr: true,
		},
		{
			name: "no requires suitability",
			args: args{cluster: &Cluster{
				Name:          "terra-dev",
				Provider:      "google",
				GoogleProject: "broad-dsde-dev",
				Base:          testutils.PointerTo("live"),
				Address:       testutils.PointerTo("1.2.3.4"),
				Location:      "us-central1-a",
			}},
			wantErr: true,
		},
		{
			name: "valid with google",
			args: args{cluster: &Cluster{
				Name:                "terra-dev",
				Provider:            "google",
				GoogleProject:       "broad-dsde-dev",
				Base:                testutils.PointerTo("live"),
				Address:             testutils.PointerTo("1.2.3.4"),
				RequiresSuitability: testutils.PointerTo(false),
				Location:            "us-central1-a",
				HelmfileRef:         testutils.PointerTo("HEAD"),
			}},
			wantErr: false,
		},
		{
			name: "valid with azure",
			args: args{cluster: &Cluster{
				Name:                "terra-dev",
				Provider:            "azure",
				AzureSubscription:   "some uuid probably",
				Base:                testutils.PointerTo("live"),
				Address:             testutils.PointerTo("1.2.3.4"),
				RequiresSuitability: testutils.PointerTo(false),
				Location:            "us-central1-a",
				HelmfileRef:         testutils.PointerTo("HEAD"),
			}},
			wantErr: false,
		},
		{
			name: "missing location",
			args: args{cluster: &Cluster{
				Name:                "terra-dev",
				Provider:            "azure",
				AzureSubscription:   "some uuid probably",
				Base:                testutils.PointerTo("live"),
				Address:             testutils.PointerTo("1.2.3.4"),
				RequiresSuitability: testutils.PointerTo(false),
				HelmfileRef:         testutils.PointerTo("HEAD"),
			}},
			wantErr: true,
		},
		{
			name: "valid with location",
			args: args{cluster: &Cluster{
				Name:                "terra-dev",
				Provider:            "google",
				GoogleProject:       "broad-dsde-dev",
				Location:            "us-central1-a",
				Base:                testutils.PointerTo("live"),
				Address:             testutils.PointerTo("1.2.3.4"),
				RequiresSuitability: testutils.PointerTo(false),
				HelmfileRef:         testutils.PointerTo("HEAD"),
			}},
			wantErr: false,
		},
		{
			name: "missing helfileRef",
			args: args{cluster: &Cluster{
				Name:                "terra-dev",
				Provider:            "azure",
				AzureSubscription:   "some uuid probably",
				Base:                testutils.PointerTo("live"),
				Address:             testutils.PointerTo("1.2.3.4"),
				RequiresSuitability: testutils.PointerTo(false),
			}},
			wantErr: true,
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
