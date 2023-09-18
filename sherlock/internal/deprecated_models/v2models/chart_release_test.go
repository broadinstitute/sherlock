package v2models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"
)

// Test_chartReleaseSelectorToQuery can't fully test chartReleaseSelectorToQuery because that function actually
// makes use of its database reference to resolve associative selectors. We do unit test the other SelectorToQuery
// functions that would be called, and the untested execution-paths can be checked with larger functional tests.
func Test_chartReleaseSelectorToQuery(t *testing.T) {
	type args struct {
		db       *gorm.DB
		selector string
	}
	tests := []struct {
		name    string
		args    args
		want    ChartRelease
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
			want: ChartRelease{Model: gorm.Model{ID: 123}},
		},
		{
			name:    "invalid id",
			args:    args{selector: testutils.StringNumberTooBigForInt},
			wantErr: true,
		},
		{
			name: "valid name",
			args: args{selector: "terra-prod"},
			want: ChartRelease{Name: "terra-prod"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := chartReleaseSelectorToQuery(tt.args.db, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("chartReleaseSelectorToQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

// Test_chartReleaseToSelectors is thorough because chartReleaseToSelectors is what helps uniqueness check new
// chart releases as they're added. If this function misses possible selectors, those selectors will be unusable at
// the very least--most likely, a somehow-invalid chart release is getting added and will cause infrastructure-level
// errors.
func Test_chartReleaseToSelectors(t *testing.T) {
	type args struct {
		chartRelease *ChartRelease
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "none",
			args: args{chartRelease: &ChartRelease{}},
			want: nil,
		},
		{
			name: "name",
			args: args{chartRelease: &ChartRelease{Name: "foobar"}},
			want: []string{"foobar"},
		},
		{
			name: "id",
			args: args{chartRelease: &ChartRelease{Model: gorm.Model{ID: 123}}},
			want: []string{"123"},
		},
		{
			name: "name and id",
			args: args{chartRelease: &ChartRelease{Model: gorm.Model{ID: 123}, Name: "foobar"}},
			want: []string{"foobar", "123"},
		},
		{
			name: "environment-based selectors",
			args: args{chartRelease: &ChartRelease{
				Chart:       &Chart{Name: "chart-a", Model: gorm.Model{ID: 246}},
				Environment: &Environment{Name: "environment-a", Model: gorm.Model{ID: 135}},
			}},
			want: []string{"environment-a/chart-a", "environment-a/246", "135/chart-a", "135/246"},
		},
		{
			name: "environment-based selectors from environment id only",
			args: args{chartRelease: &ChartRelease{
				Chart:         &Chart{Name: "chart-a", Model: gorm.Model{ID: 246}},
				EnvironmentID: utils.PointerTo[uint](135),
			}},
			want: []string{"135/chart-a", "135/246"},
		},
		{
			name: "environment-based selectors from chart and environment id only",
			args: args{chartRelease: &ChartRelease{
				ChartID:       246,
				EnvironmentID: utils.PointerTo[uint](135),
			}},
			want: []string{"135/246"},
		},
		{
			name: "environment-based selectors with ids available",
			args: args{chartRelease: &ChartRelease{
				ChartID:       246,
				Chart:         &Chart{Name: "chart-a", Model: gorm.Model{ID: 246}},
				EnvironmentID: utils.PointerTo[uint](135),
				Environment:   &Environment{Name: "environment-a", Model: gorm.Model{ID: 135}},
			}},
			want: []string{"environment-a/chart-a", "environment-a/246", "135/chart-a", "135/246"},
		},
		{
			name: "cluster-based selectors",
			args: args{chartRelease: &ChartRelease{
				Chart:     &Chart{Name: "chart-a", Model: gorm.Model{ID: 246}},
				Cluster:   &Cluster{Name: "cluster-a", Model: gorm.Model{ID: 789}},
				Namespace: "namespace-a",
			}},
			want: []string{"cluster-a/namespace-a/chart-a", "cluster-a/namespace-a/246", "789/namespace-a/chart-a", "789/namespace-a/246"},
		},
		{
			name: "cluster-based selectors from cluster id only",
			args: args{chartRelease: &ChartRelease{
				Chart:     &Chart{Name: "chart-a", Model: gorm.Model{ID: 246}},
				ClusterID: utils.PointerTo[uint](789),
				Namespace: "namespace-a",
			}},
			want: []string{"789/namespace-a/chart-a", "789/namespace-a/246"},
		},
		{
			name: "cluster-based selectors from chart and cluster id only",
			args: args{chartRelease: &ChartRelease{
				ChartID:   246,
				ClusterID: utils.PointerTo[uint](789),
				Namespace: "namespace-a",
			}},
			want: []string{"789/namespace-a/246"},
		},
		{
			name: "cluster-based selectors with ids available",
			args: args{chartRelease: &ChartRelease{
				ChartID:   246,
				Chart:     &Chart{Name: "chart-a", Model: gorm.Model{ID: 246}},
				ClusterID: utils.PointerTo[uint](789),
				Cluster:   &Cluster{Name: "cluster-a", Model: gorm.Model{ID: 789}},
				Namespace: "namespace-a",
			}},
			want: []string{"cluster-a/namespace-a/chart-a", "cluster-a/namespace-a/246", "789/namespace-a/chart-a", "789/namespace-a/246"},
		},
		{
			name: "all possible selectors",
			args: args{chartRelease: &ChartRelease{
				Model:       gorm.Model{ID: 123},
				Name:        "foobar",
				Chart:       &Chart{Name: "chart-a", Model: gorm.Model{ID: 246}},
				Cluster:     &Cluster{Name: "cluster-a", Model: gorm.Model{ID: 789}},
				Namespace:   "namespace-a",
				Environment: &Environment{Name: "environment-a", Model: gorm.Model{ID: 135}},
			}},
			want: []string{"foobar", "environment-a/chart-a", "environment-a/246", "135/chart-a", "135/246", "cluster-a/namespace-a/chart-a", "cluster-a/namespace-a/246", "789/namespace-a/chart-a", "789/namespace-a/246", "123"},
		},
		{
			name: "all possible selectors with IDs available",
			args: args{chartRelease: &ChartRelease{
				Model:         gorm.Model{ID: 123},
				Name:          "foobar",
				ChartID:       246,
				Chart:         &Chart{Name: "chart-a", Model: gorm.Model{ID: 246}},
				ClusterID:     utils.PointerTo[uint](789),
				Cluster:       &Cluster{Name: "cluster-a", Model: gorm.Model{ID: 789}},
				Namespace:     "namespace-a",
				EnvironmentID: utils.PointerTo[uint](135),
				Environment:   &Environment{Name: "environment-a", Model: gorm.Model{ID: 135}},
			}},
			want: []string{"foobar", "environment-a/chart-a", "environment-a/246", "135/chart-a", "135/246", "cluster-a/namespace-a/chart-a", "cluster-a/namespace-a/246", "789/namespace-a/chart-a", "789/namespace-a/246", "123"},
		},
		{
			name: "cluster selectors disappear if namespace is empty",
			args: args{chartRelease: &ChartRelease{
				Model:       gorm.Model{ID: 123},
				Name:        "foobar",
				Chart:       &Chart{Name: "chart-a", Model: gorm.Model{ID: 246}},
				Cluster:     &Cluster{Name: "cluster-a", Model: gorm.Model{ID: 789}},
				Namespace:   "",
				Environment: &Environment{Name: "environment-a", Model: gorm.Model{ID: 135}},
			}},
			want: []string{"foobar", "environment-a/chart-a", "environment-a/246", "135/chart-a", "135/246", "123"},
		},
		{
			name: "all possible selectors without full selectors on associations",
			args: args{chartRelease: &ChartRelease{
				Model:       gorm.Model{ID: 123},
				Name:        "foobar",
				Chart:       &Chart{Model: gorm.Model{ID: 246}},
				Cluster:     &Cluster{Name: "cluster-a"},
				Namespace:   "namespace-a",
				Environment: &Environment{Name: "environment-a"},
			}},
			want: []string{"foobar", "environment-a/246", "cluster-a/namespace-a/246", "123"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := chartReleaseToSelectors(tt.args.chartRelease)
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

func Test_validateChartRelease(t *testing.T) {
	type args struct {
		chartRelease *ChartRelease
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid with all fields set",
			args: args{chartRelease: &ChartRelease{
				Name:            "foobar",
				ChartID:         246,
				ClusterID:       utils.PointerTo[uint](789),
				Namespace:       "a-namespace",
				EnvironmentID:   utils.PointerTo[uint](135),
				DestinationType: "environment",
				ChartReleaseVersion: ChartReleaseVersion{
					ResolvedAt: utils.PointerTo(time.Now()),

					AppVersionResolver: utils.PointerTo("branch"),
					AppVersionExact:    utils.PointerTo("v1.2.3"),
					AppVersionCommit:   utils.PointerTo("a1b2c3d4"),
					AppVersionBranch:   utils.PointerTo("main"),
					AppVersion: &AppVersion{
						Model:      gorm.Model{ID: 1},
						AppVersion: "v1.2.3",
						GitCommit:  "a1b2c3d4",
						GitBranch:  "main",
					},
					AppVersionID: utils.PointerTo[uint](1),

					ChartVersionResolver: utils.PointerTo("latest"),
					ChartVersionExact:    utils.PointerTo("v0.0.100"),
					ChartVersion: &ChartVersion{
						Model:        gorm.Model{ID: 2},
						ChartVersion: "v0.0.100",
					},
					ChartVersionID: utils.PointerTo[uint](2),

					HelmfileRef:         utils.PointerTo("e5f6g7h8"),
					FirecloudDevelopRef: utils.PointerTo("dev"),
				},
			}},
			wantErr: false,
		},
		{
			name: "valid with only cluster",
			args: args{chartRelease: &ChartRelease{
				Name:            "foobar",
				ChartID:         246,
				ClusterID:       utils.PointerTo[uint](789),
				Namespace:       "a-namespace",
				DestinationType: "cluster",
				ChartReleaseVersion: ChartReleaseVersion{
					ResolvedAt: utils.PointerTo(time.Now()),

					AppVersionResolver: utils.PointerTo("branch"),
					AppVersionExact:    utils.PointerTo("v1.2.3"),
					AppVersionCommit:   utils.PointerTo("a1b2c3d4"),
					AppVersionBranch:   utils.PointerTo("main"),
					AppVersion: &AppVersion{
						Model:      gorm.Model{ID: 1},
						AppVersion: "v1.2.3",
						GitCommit:  "a1b2c3d4",
						GitBranch:  "main",
					},
					AppVersionID: utils.PointerTo[uint](1),

					ChartVersionResolver: utils.PointerTo("latest"),
					ChartVersionExact:    utils.PointerTo("v0.0.100"),
					ChartVersion: &ChartVersion{
						Model:        gorm.Model{ID: 2},
						ChartVersion: "v0.0.100",
					},
					ChartVersionID: utils.PointerTo[uint](2),

					HelmfileRef:         utils.PointerTo("e5f6g7h8"),
					FirecloudDevelopRef: utils.PointerTo("dev"),
				},
			}},
			wantErr: false,
		},
		{
			name: "valid with only environment",
			args: args{chartRelease: &ChartRelease{
				Name:            "foobar",
				ChartID:         246,
				EnvironmentID:   utils.PointerTo[uint](135),
				DestinationType: "environment",
				ChartReleaseVersion: ChartReleaseVersion{
					ResolvedAt: utils.PointerTo(time.Now()),

					AppVersionResolver: utils.PointerTo("branch"),
					AppVersionExact:    utils.PointerTo("v1.2.3"),
					AppVersionCommit:   utils.PointerTo("a1b2c3d4"),
					AppVersionBranch:   utils.PointerTo("main"),
					AppVersion: &AppVersion{
						Model:      gorm.Model{ID: 1},
						AppVersion: "v1.2.3",
						GitCommit:  "a1b2c3d4",
						GitBranch:  "main",
					},
					AppVersionID: utils.PointerTo[uint](1),

					ChartVersionResolver: utils.PointerTo("latest"),
					ChartVersionExact:    utils.PointerTo("v0.0.100"),
					ChartVersion: &ChartVersion{
						Model:        gorm.Model{ID: 2},
						ChartVersion: "v0.0.100",
					},
					ChartVersionID: utils.PointerTo[uint](2),

					HelmfileRef:         utils.PointerTo("e5f6g7h8"),
					FirecloudDevelopRef: utils.PointerTo("dev"),
				},
			}},
			wantErr: false,
		},
		{
			name: "no name",
			args: args{chartRelease: &ChartRelease{
				ChartID:         246,
				ClusterID:       utils.PointerTo[uint](789),
				Namespace:       "a-namespace",
				EnvironmentID:   utils.PointerTo[uint](135),
				DestinationType: "environment",
				ChartReleaseVersion: ChartReleaseVersion{
					ResolvedAt: utils.PointerTo(time.Now()),

					AppVersionResolver: utils.PointerTo("branch"),
					AppVersionExact:    utils.PointerTo("v1.2.3"),
					AppVersionCommit:   utils.PointerTo("a1b2c3d4"),
					AppVersionBranch:   utils.PointerTo("main"),
					AppVersion: &AppVersion{
						Model:      gorm.Model{ID: 1},
						AppVersion: "v1.2.3",
						GitCommit:  "a1b2c3d4",
						GitBranch:  "main",
					},
					AppVersionID: utils.PointerTo[uint](1),

					ChartVersionResolver: utils.PointerTo("latest"),
					ChartVersionExact:    utils.PointerTo("v0.0.100"),
					ChartVersion: &ChartVersion{
						Model:        gorm.Model{ID: 2},
						ChartVersion: "v0.0.100",
					},
					ChartVersionID: utils.PointerTo[uint](2),

					HelmfileRef:         utils.PointerTo("e5f6g7h8"),
					FirecloudDevelopRef: utils.PointerTo("dev"),
				},
			}},
			wantErr: true,
		},
		{
			name: "no chart",
			args: args{chartRelease: &ChartRelease{
				Name:            "foobar",
				ClusterID:       utils.PointerTo[uint](789),
				Namespace:       "a-namespace",
				EnvironmentID:   utils.PointerTo[uint](135),
				DestinationType: "environment",
				ChartReleaseVersion: ChartReleaseVersion{
					ResolvedAt: utils.PointerTo(time.Now()),

					AppVersionResolver: utils.PointerTo("branch"),
					AppVersionExact:    utils.PointerTo("v1.2.3"),
					AppVersionCommit:   utils.PointerTo("a1b2c3d4"),
					AppVersionBranch:   utils.PointerTo("main"),
					AppVersion: &AppVersion{
						Model:      gorm.Model{ID: 1},
						AppVersion: "v1.2.3",
						GitCommit:  "a1b2c3d4",
						GitBranch:  "main",
					},
					AppVersionID: utils.PointerTo[uint](1),

					ChartVersionResolver: utils.PointerTo("latest"),
					ChartVersionExact:    utils.PointerTo("v0.0.100"),
					ChartVersion: &ChartVersion{
						Model:        gorm.Model{ID: 2},
						ChartVersion: "v0.0.100",
					},
					ChartVersionID: utils.PointerTo[uint](2),

					HelmfileRef:         utils.PointerTo("e5f6g7h8"),
					FirecloudDevelopRef: utils.PointerTo("dev"),
				},
			}},
			wantErr: true,
		},
		{
			name: "wrong destination for environment",
			args: args{chartRelease: &ChartRelease{
				Name:            "foobar",
				ChartID:         246,
				ClusterID:       utils.PointerTo[uint](789),
				Namespace:       "a-namespace",
				EnvironmentID:   utils.PointerTo[uint](135),
				DestinationType: "cluster",
				ChartReleaseVersion: ChartReleaseVersion{
					ResolvedAt: utils.PointerTo(time.Now()),

					AppVersionResolver: utils.PointerTo("branch"),
					AppVersionExact:    utils.PointerTo("v1.2.3"),
					AppVersionCommit:   utils.PointerTo("a1b2c3d4"),
					AppVersionBranch:   utils.PointerTo("main"),
					AppVersion: &AppVersion{
						Model:      gorm.Model{ID: 1},
						AppVersion: "v1.2.3",
						GitCommit:  "a1b2c3d4",
						GitBranch:  "main",
					},
					AppVersionID: utils.PointerTo[uint](1),

					ChartVersionResolver: utils.PointerTo("latest"),
					ChartVersionExact:    utils.PointerTo("v0.0.100"),
					ChartVersion: &ChartVersion{
						Model:        gorm.Model{ID: 2},
						ChartVersion: "v0.0.100",
					},
					ChartVersionID: utils.PointerTo[uint](2),

					HelmfileRef:         utils.PointerTo("e5f6g7h8"),
					FirecloudDevelopRef: utils.PointerTo("dev"),
				},
			}},
			wantErr: true,
		},
		{
			name: "wrong destination for only having cluster",
			args: args{chartRelease: &ChartRelease{
				Name:            "foobar",
				ChartID:         246,
				ClusterID:       utils.PointerTo[uint](789),
				Namespace:       "a-namespace",
				DestinationType: "environment",
				ChartReleaseVersion: ChartReleaseVersion{
					ResolvedAt: utils.PointerTo(time.Now()),

					AppVersionResolver: utils.PointerTo("branch"),
					AppVersionExact:    utils.PointerTo("v1.2.3"),
					AppVersionCommit:   utils.PointerTo("a1b2c3d4"),
					AppVersionBranch:   utils.PointerTo("main"),
					AppVersion: &AppVersion{
						Model:      gorm.Model{ID: 1},
						AppVersion: "v1.2.3",
						GitCommit:  "a1b2c3d4",
						GitBranch:  "main",
					},
					AppVersionID: utils.PointerTo[uint](1),

					ChartVersionResolver: utils.PointerTo("latest"),
					ChartVersionExact:    utils.PointerTo("v0.0.100"),
					ChartVersion: &ChartVersion{
						Model:        gorm.Model{ID: 2},
						ChartVersion: "v0.0.100",
					},
					ChartVersionID: utils.PointerTo[uint](2),

					HelmfileRef:         utils.PointerTo("e5f6g7h8"),
					FirecloudDevelopRef: utils.PointerTo("dev"),
				},
			}},
			wantErr: true,
		},
		{
			name: "no environment or cluster",
			args: args{chartRelease: &ChartRelease{
				Name:            "foobar",
				ChartID:         246,
				DestinationType: "environment",
				ChartReleaseVersion: ChartReleaseVersion{
					ResolvedAt: utils.PointerTo(time.Now()),

					AppVersionResolver: utils.PointerTo("branch"),
					AppVersionExact:    utils.PointerTo("v1.2.3"),
					AppVersionCommit:   utils.PointerTo("a1b2c3d4"),
					AppVersionBranch:   utils.PointerTo("main"),
					AppVersion: &AppVersion{
						Model:      gorm.Model{ID: 1},
						AppVersion: "v1.2.3",
						GitCommit:  "a1b2c3d4",
						GitBranch:  "main",
					},
					AppVersionID: utils.PointerTo[uint](1),

					ChartVersionResolver: utils.PointerTo("latest"),
					ChartVersionExact:    utils.PointerTo("v0.0.100"),
					ChartVersion: &ChartVersion{
						Model:        gorm.Model{ID: 2},
						ChartVersion: "v0.0.100",
					},
					ChartVersionID: utils.PointerTo[uint](2),

					HelmfileRef:         utils.PointerTo("e5f6g7h8"),
					FirecloudDevelopRef: utils.PointerTo("dev"),
				},
			}},
			wantErr: true,
		},
		{
			name: "cluster but no namespace",
			args: args{chartRelease: &ChartRelease{
				Name:            "foobar",
				ChartID:         246,
				ClusterID:       utils.PointerTo[uint](789),
				EnvironmentID:   utils.PointerTo[uint](135),
				DestinationType: "environment",
				ChartReleaseVersion: ChartReleaseVersion{
					ResolvedAt: utils.PointerTo(time.Now()),

					AppVersionResolver: utils.PointerTo("branch"),
					AppVersionExact:    utils.PointerTo("v1.2.3"),
					AppVersionCommit:   utils.PointerTo("a1b2c3d4"),
					AppVersionBranch:   utils.PointerTo("main"),
					AppVersion: &AppVersion{
						Model:      gorm.Model{ID: 1},
						AppVersion: "v1.2.3",
						GitCommit:  "a1b2c3d4",
						GitBranch:  "main",
					},
					AppVersionID: utils.PointerTo[uint](1),

					ChartVersionResolver: utils.PointerTo("latest"),
					ChartVersionExact:    utils.PointerTo("v0.0.100"),
					ChartVersion: &ChartVersion{
						Model:        gorm.Model{ID: 2},
						ChartVersion: "v0.0.100",
					},
					ChartVersionID: utils.PointerTo[uint](2),

					HelmfileRef:         utils.PointerTo("e5f6g7h8"),
					FirecloudDevelopRef: utils.PointerTo("dev"),
				},
			}},
			wantErr: true,
		},
		{
			name: "namespace but no cluster",
			args: args{chartRelease: &ChartRelease{
				Name:            "foobar",
				ChartID:         246,
				Namespace:       "a-namespace",
				EnvironmentID:   utils.PointerTo[uint](135),
				DestinationType: "environment",
				ChartReleaseVersion: ChartReleaseVersion{
					ResolvedAt: utils.PointerTo(time.Now()),

					AppVersionResolver: utils.PointerTo("branch"),
					AppVersionExact:    utils.PointerTo("v1.2.3"),
					AppVersionCommit:   utils.PointerTo("a1b2c3d4"),
					AppVersionBranch:   utils.PointerTo("main"),
					AppVersion: &AppVersion{
						Model:      gorm.Model{ID: 1},
						AppVersion: "v1.2.3",
						GitCommit:  "a1b2c3d4",
						GitBranch:  "main",
					},
					AppVersionID: utils.PointerTo[uint](1),

					ChartVersionResolver: utils.PointerTo("latest"),
					ChartVersionExact:    utils.PointerTo("v0.0.100"),
					ChartVersion: &ChartVersion{
						Model:        gorm.Model{ID: 2},
						ChartVersion: "v0.0.100",
					},
					ChartVersionID: utils.PointerTo[uint](2),

					HelmfileRef:         utils.PointerTo("e5f6g7h8"),
					FirecloudDevelopRef: utils.PointerTo("dev"),
				},
			}},
			wantErr: true,
		},
		{
			name: "invalid inside the ChartReleaseVersion",
			args: args{chartRelease: &ChartRelease{
				Name:            "foobar",
				ChartID:         246,
				ClusterID:       utils.PointerTo[uint](789),
				Namespace:       "a-namespace",
				EnvironmentID:   utils.PointerTo[uint](135),
				DestinationType: "environment",
				ChartReleaseVersion: ChartReleaseVersion{
					// This not being nil is something only checked inside ChartReleaseVersion.validate--we won't
					// re-enumerate the thousand lines of tests, just check that we're calling it
					ResolvedAt: nil,

					AppVersionResolver: utils.PointerTo("branch"),
					AppVersionExact:    utils.PointerTo("v1.2.3"),
					AppVersionCommit:   utils.PointerTo("a1b2c3d4"),
					AppVersionBranch:   utils.PointerTo("main"),
					AppVersion: &AppVersion{
						Model:      gorm.Model{ID: 1},
						AppVersion: "v1.2.3",
						GitCommit:  "a1b2c3d4",
						GitBranch:  "main",
					},
					AppVersionID: utils.PointerTo[uint](1),

					ChartVersionResolver: utils.PointerTo("latest"),
					ChartVersionExact:    utils.PointerTo("v0.0.100"),
					ChartVersion: &ChartVersion{
						Model:        gorm.Model{ID: 2},
						ChartVersion: "v0.0.100",
					},
					ChartVersionID: utils.PointerTo[uint](2),

					HelmfileRef:         utils.PointerTo("e5f6g7h8"),
					FirecloudDevelopRef: utils.PointerTo("dev"),
				},
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateChartRelease(tt.args.chartRelease); (err != nil) != tt.wantErr {
				t.Errorf("validateChartRelease() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChartRelease_GetCiIdentifier(t *testing.T) {
	type fields struct {
		Model                   gorm.Model
		CiIdentifier            *CiIdentifier
		Chart                   *Chart
		ChartID                 uint
		Cluster                 *Cluster
		ClusterID               *uint
		DestinationType         string
		Environment             *Environment
		EnvironmentID           *uint
		Name                    string
		Namespace               string
		ChartReleaseVersion     ChartReleaseVersion
		Subdomain               *string
		Protocol                *string
		Port                    *uint
		PagerdutyIntegration    *PagerdutyIntegration
		PagerdutyIntegrationID  *uint
		IncludeInBulkChangesets *bool
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
				ResourceType: "chart-release",
				ResourceID:   123,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ChartRelease{
				Model:                   tt.fields.Model,
				CiIdentifier:            tt.fields.CiIdentifier,
				Chart:                   tt.fields.Chart,
				ChartID:                 tt.fields.ChartID,
				Cluster:                 tt.fields.Cluster,
				ClusterID:               tt.fields.ClusterID,
				DestinationType:         tt.fields.DestinationType,
				Environment:             tt.fields.Environment,
				EnvironmentID:           tt.fields.EnvironmentID,
				Name:                    tt.fields.Name,
				Namespace:               tt.fields.Namespace,
				ChartReleaseVersion:     tt.fields.ChartReleaseVersion,
				Subdomain:               tt.fields.Subdomain,
				Protocol:                tt.fields.Protocol,
				Port:                    tt.fields.Port,
				PagerdutyIntegration:    tt.fields.PagerdutyIntegration,
				PagerdutyIntegrationID:  tt.fields.PagerdutyIntegrationID,
				IncludeInBulkChangesets: tt.fields.IncludeInBulkChangesets,
			}
			if got := c.GetCiIdentifier(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCiIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}
