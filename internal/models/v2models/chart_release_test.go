package v2models

import (
	"github.com/broadinstitute/sherlock/internal/testutils"
	"gorm.io/gorm"
	"testing"
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
		chartRelease ChartRelease
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "none",
			args: args{chartRelease: ChartRelease{}},
			want: nil,
		},
		{
			name: "name",
			args: args{chartRelease: ChartRelease{Name: "foobar"}},
			want: []string{"foobar"},
		},
		{
			name: "id",
			args: args{chartRelease: ChartRelease{Model: gorm.Model{ID: 123}}},
			want: []string{"123"},
		},
		{
			name: "name and id",
			args: args{chartRelease: ChartRelease{Model: gorm.Model{ID: 123}, Name: "foobar"}},
			want: []string{"foobar", "123"},
		},
		{
			name: "environment-based selectors",
			args: args{chartRelease: ChartRelease{
				Chart:       Chart{Name: "chart-a", Model: gorm.Model{ID: 246}},
				Environment: &Environment{Name: "environment-a", Model: gorm.Model{ID: 135}},
			}},
			want: []string{"environment-a/chart-a", "environment-a/246", "135/chart-a", "135/246"},
		},
		{
			name: "environment-based selectors from environment id only",
			args: args{chartRelease: ChartRelease{
				Chart:         Chart{Name: "chart-a", Model: gorm.Model{ID: 246}},
				EnvironmentID: testutils.PointerTo[uint](135),
			}},
			want: []string{"135/chart-a", "135/246"},
		},
		{
			name: "environment-based selectors from chart and environment id only",
			args: args{chartRelease: ChartRelease{
				ChartID:       246,
				EnvironmentID: testutils.PointerTo[uint](135),
			}},
			want: []string{"135/246"},
		},
		{
			name: "environment-based selectors with ids available",
			args: args{chartRelease: ChartRelease{
				ChartID:       246,
				Chart:         Chart{Name: "chart-a", Model: gorm.Model{ID: 246}},
				EnvironmentID: testutils.PointerTo[uint](135),
				Environment:   &Environment{Name: "environment-a", Model: gorm.Model{ID: 135}},
			}},
			want: []string{"environment-a/chart-a", "environment-a/246", "135/chart-a", "135/246"},
		},
		{
			name: "cluster-based selectors",
			args: args{chartRelease: ChartRelease{
				Chart:     Chart{Name: "chart-a", Model: gorm.Model{ID: 246}},
				Cluster:   &Cluster{Name: "cluster-a", Model: gorm.Model{ID: 789}},
				Namespace: "namespace-a",
			}},
			want: []string{"cluster-a/namespace-a/chart-a", "cluster-a/namespace-a/246", "789/namespace-a/chart-a", "789/namespace-a/246"},
		},
		{
			name: "cluster-based selectors from cluster id only",
			args: args{chartRelease: ChartRelease{
				Chart:     Chart{Name: "chart-a", Model: gorm.Model{ID: 246}},
				ClusterID: testutils.PointerTo[uint](789),
				Namespace: "namespace-a",
			}},
			want: []string{"789/namespace-a/chart-a", "789/namespace-a/246"},
		},
		{
			name: "cluster-based selectors from chart and cluster id only",
			args: args{chartRelease: ChartRelease{
				ChartID:   246,
				ClusterID: testutils.PointerTo[uint](789),
				Namespace: "namespace-a",
			}},
			want: []string{"789/namespace-a/246"},
		},
		{
			name: "cluster-based selectors with ids available",
			args: args{chartRelease: ChartRelease{
				ChartID:   246,
				Chart:     Chart{Name: "chart-a", Model: gorm.Model{ID: 246}},
				ClusterID: testutils.PointerTo[uint](789),
				Cluster:   &Cluster{Name: "cluster-a", Model: gorm.Model{ID: 789}},
				Namespace: "namespace-a",
			}},
			want: []string{"cluster-a/namespace-a/chart-a", "cluster-a/namespace-a/246", "789/namespace-a/chart-a", "789/namespace-a/246"},
		},
		{
			name: "all possible selectors",
			args: args{chartRelease: ChartRelease{
				Model:       gorm.Model{ID: 123},
				Name:        "foobar",
				Chart:       Chart{Name: "chart-a", Model: gorm.Model{ID: 246}},
				Cluster:     &Cluster{Name: "cluster-a", Model: gorm.Model{ID: 789}},
				Namespace:   "namespace-a",
				Environment: &Environment{Name: "environment-a", Model: gorm.Model{ID: 135}},
			}},
			want: []string{"environment-a/chart-a", "environment-a/246", "135/chart-a", "135/246", "cluster-a/namespace-a/chart-a", "cluster-a/namespace-a/246", "789/namespace-a/chart-a", "789/namespace-a/246", "foobar", "123"},
		},
		{
			name: "all possible selectors with IDs available",
			args: args{chartRelease: ChartRelease{
				Model:         gorm.Model{ID: 123},
				Name:          "foobar",
				ChartID:       246,
				Chart:         Chart{Name: "chart-a", Model: gorm.Model{ID: 246}},
				ClusterID:     testutils.PointerTo[uint](789),
				Cluster:       &Cluster{Name: "cluster-a", Model: gorm.Model{ID: 789}},
				Namespace:     "namespace-a",
				EnvironmentID: testutils.PointerTo[uint](135),
				Environment:   &Environment{Name: "environment-a", Model: gorm.Model{ID: 135}},
			}},
			want: []string{"environment-a/chart-a", "environment-a/246", "135/chart-a", "135/246", "cluster-a/namespace-a/chart-a", "cluster-a/namespace-a/246", "789/namespace-a/chart-a", "789/namespace-a/246", "foobar", "123"},
		},
		{
			name: "cluster selectors disappear if namespace is empty",
			args: args{chartRelease: ChartRelease{
				Model:       gorm.Model{ID: 123},
				Name:        "foobar",
				Chart:       Chart{Name: "chart-a", Model: gorm.Model{ID: 246}},
				Cluster:     &Cluster{Name: "cluster-a", Model: gorm.Model{ID: 789}},
				Namespace:   "",
				Environment: &Environment{Name: "environment-a", Model: gorm.Model{ID: 135}},
			}},
			want: []string{"environment-a/chart-a", "environment-a/246", "135/chart-a", "135/246", "foobar", "123"},
		},
		{
			name: "all possible selectors without full selectors on associations",
			args: args{chartRelease: ChartRelease{
				Model:       gorm.Model{ID: 123},
				Name:        "foobar",
				Chart:       Chart{Model: gorm.Model{ID: 246}},
				Cluster:     &Cluster{Name: "cluster-a"},
				Namespace:   "namespace-a",
				Environment: &Environment{Name: "environment-a"},
			}},
			want: []string{"environment-a/246", "cluster-a/namespace-a/246", "foobar", "123"},
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
		chartRelease ChartRelease
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid with all fields set",
			args: args{chartRelease: ChartRelease{
				Name:                     "foobar",
				ChartID:                  246,
				ClusterID:                testutils.PointerTo[uint](789),
				Namespace:                "a-namespace",
				EnvironmentID:            testutils.PointerTo[uint](135),
				DestinationType:          "environment",
				HelmfileRef:              testutils.PointerTo("HEAD"),
				TargetAppVersionUse:      testutils.PointerTo("branch"),
				TargetAppVersionBranch:   testutils.PointerTo("main"),
				TargetAppVersionCommit:   testutils.PointerTo("abcd"),
				TargetAppVersionExact:    testutils.PointerTo("abcd2"),
				TargetChartVersionUse:    testutils.PointerTo("latest"),
				TargetChartVersionExact:  testutils.PointerTo("zxy"),
				ThelmaMode:               testutils.PointerTo("deploy"),
				CurrentChartVersionExact: testutils.PointerTo("zxy2"),
				CurrentAppVersionExact:   testutils.PointerTo("abcd3"),
			}},
			wantErr: false,
		},
		{
			name: "valid with app branch",
			args: args{chartRelease: ChartRelease{
				Name:                   "foobar",
				ChartID:                246,
				ClusterID:              testutils.PointerTo[uint](789),
				Namespace:              "a-namespace",
				EnvironmentID:          testutils.PointerTo[uint](135),
				DestinationType:        "environment",
				HelmfileRef:            testutils.PointerTo("HEAD"),
				TargetAppVersionUse:    testutils.PointerTo("branch"),
				TargetAppVersionBranch: testutils.PointerTo("main"),
				TargetChartVersionUse:  testutils.PointerTo("latest"),
			}},
			wantErr: false,
		},
		{
			name: "valid with app commit",
			args: args{chartRelease: ChartRelease{
				Name:                   "foobar",
				ChartID:                246,
				ClusterID:              testutils.PointerTo[uint](789),
				Namespace:              "a-namespace",
				EnvironmentID:          testutils.PointerTo[uint](135),
				DestinationType:        "environment",
				HelmfileRef:            testutils.PointerTo("HEAD"),
				TargetAppVersionUse:    testutils.PointerTo("commit"),
				TargetAppVersionCommit: testutils.PointerTo("abcd"),
				TargetChartVersionUse:  testutils.PointerTo("latest"),
			}},
			wantErr: false,
		},
		{
			name: "valid with exact versions",
			args: args{chartRelease: ChartRelease{
				Name:                    "foobar",
				ChartID:                 246,
				ClusterID:               testutils.PointerTo[uint](789),
				Namespace:               "a-namespace",
				EnvironmentID:           testutils.PointerTo[uint](135),
				DestinationType:         "environment",
				HelmfileRef:             testutils.PointerTo("HEAD"),
				TargetAppVersionUse:     testutils.PointerTo("exact"),
				TargetAppVersionExact:   testutils.PointerTo("abcd"),
				TargetChartVersionUse:   testutils.PointerTo("exact"),
				TargetChartVersionExact: testutils.PointerTo("zxy"),
			}},
			wantErr: false,
		},
		{
			name: "valid without app",
			args: args{chartRelease: ChartRelease{
				Name:                  "foobar",
				ChartID:               246,
				ClusterID:             testutils.PointerTo[uint](789),
				Namespace:             "a-namespace",
				EnvironmentID:         testutils.PointerTo[uint](135),
				DestinationType:       "environment",
				HelmfileRef:           testutils.PointerTo("HEAD"),
				TargetChartVersionUse: testutils.PointerTo("latest"),
			}},
			wantErr: false,
		},
		{
			name: "valid with only cluster",
			args: args{chartRelease: ChartRelease{
				Name:                  "foobar",
				ChartID:               246,
				ClusterID:             testutils.PointerTo[uint](789),
				Namespace:             "a-namespace",
				DestinationType:       "cluster",
				HelmfileRef:           testutils.PointerTo("HEAD"),
				TargetChartVersionUse: testutils.PointerTo("latest"),
			}},
			wantErr: false,
		},
		{
			name: "valid with only environment",
			args: args{chartRelease: ChartRelease{
				Name:                  "foobar",
				ChartID:               246,
				EnvironmentID:         testutils.PointerTo[uint](135),
				DestinationType:       "environment",
				HelmfileRef:           testutils.PointerTo("HEAD"),
				TargetChartVersionUse: testutils.PointerTo("latest"),
			}},
			wantErr: false,
		},
		{
			name: "no name",
			args: args{chartRelease: ChartRelease{
				ChartID:                246,
				ClusterID:              testutils.PointerTo[uint](789),
				Namespace:              "a-namespace",
				EnvironmentID:          testutils.PointerTo[uint](135),
				DestinationType:        "environment",
				HelmfileRef:            testutils.PointerTo("HEAD"),
				TargetAppVersionUse:    testutils.PointerTo("branch"),
				TargetAppVersionBranch: testutils.PointerTo("main"),
				TargetChartVersionUse:  testutils.PointerTo("latest"),
			}},
			wantErr: true,
		},
		{
			name: "no chart",
			args: args{chartRelease: ChartRelease{
				Name:                   "foobar",
				ClusterID:              testutils.PointerTo[uint](789),
				Namespace:              "a-namespace",
				EnvironmentID:          testutils.PointerTo[uint](135),
				DestinationType:        "environment",
				HelmfileRef:            testutils.PointerTo("HEAD"),
				TargetAppVersionUse:    testutils.PointerTo("branch"),
				TargetAppVersionBranch: testutils.PointerTo("main"),
				TargetChartVersionUse:  testutils.PointerTo("latest"),
			}},
			wantErr: true,
		},
		{
			name: "wrong destination for environment",
			args: args{chartRelease: ChartRelease{
				Name:                   "foobar",
				ChartID:                246,
				ClusterID:              testutils.PointerTo[uint](789),
				Namespace:              "a-namespace",
				EnvironmentID:          testutils.PointerTo[uint](135),
				DestinationType:        "cluster",
				HelmfileRef:            testutils.PointerTo("HEAD"),
				TargetAppVersionUse:    testutils.PointerTo("branch"),
				TargetAppVersionBranch: testutils.PointerTo("main"),
				TargetChartVersionUse:  testutils.PointerTo("latest"),
			}},
			wantErr: true,
		},
		{
			name: "wrong destination for only having cluster",
			args: args{chartRelease: ChartRelease{
				Name:                   "foobar",
				ChartID:                246,
				ClusterID:              testutils.PointerTo[uint](789),
				Namespace:              "a-namespace",
				DestinationType:        "environment",
				HelmfileRef:            testutils.PointerTo("HEAD"),
				TargetAppVersionUse:    testutils.PointerTo("branch"),
				TargetAppVersionBranch: testutils.PointerTo("main"),
				TargetChartVersionUse:  testutils.PointerTo("latest"),
			}},
			wantErr: true,
		},
		{
			name: "no environment or cluster",
			args: args{chartRelease: ChartRelease{
				Name:                   "foobar",
				ChartID:                246,
				DestinationType:        "environment",
				HelmfileRef:            testutils.PointerTo("HEAD"),
				TargetAppVersionUse:    testutils.PointerTo("branch"),
				TargetAppVersionBranch: testutils.PointerTo("main"),
				TargetChartVersionUse:  testutils.PointerTo("latest"),
			}},
			wantErr: true,
		},
		{
			name: "cluster but no namespace",
			args: args{chartRelease: ChartRelease{
				Name:                   "foobar",
				ChartID:                246,
				ClusterID:              testutils.PointerTo[uint](789),
				EnvironmentID:          testutils.PointerTo[uint](135),
				DestinationType:        "environment",
				HelmfileRef:            testutils.PointerTo("HEAD"),
				TargetAppVersionUse:    testutils.PointerTo("branch"),
				TargetAppVersionBranch: testutils.PointerTo("main"),
				TargetChartVersionUse:  testutils.PointerTo("latest"),
			}},
			wantErr: true,
		},
		{
			name: "namespace but no cluster",
			args: args{chartRelease: ChartRelease{
				Name:                   "foobar",
				ChartID:                246,
				Namespace:              "a-namespace",
				EnvironmentID:          testutils.PointerTo[uint](135),
				DestinationType:        "environment",
				HelmfileRef:            testutils.PointerTo("HEAD"),
				TargetAppVersionUse:    testutils.PointerTo("branch"),
				TargetAppVersionBranch: testutils.PointerTo("main"),
				TargetChartVersionUse:  testutils.PointerTo("latest"),
			}},
			wantErr: true,
		},
		{
			name: "no helmfile ref",
			args: args{chartRelease: ChartRelease{
				Name:                   "foobar",
				ChartID:                246,
				ClusterID:              testutils.PointerTo[uint](789),
				Namespace:              "a-namespace",
				EnvironmentID:          testutils.PointerTo[uint](135),
				DestinationType:        "environment",
				TargetAppVersionUse:    testutils.PointerTo("branch"),
				TargetAppVersionBranch: testutils.PointerTo("main"),
				TargetChartVersionUse:  testutils.PointerTo("latest"),
			}},
			wantErr: true,
		},
		{
			name: "app version use branch but no branch provided",
			args: args{chartRelease: ChartRelease{
				Name:                  "foobar",
				ChartID:               246,
				ClusterID:             testutils.PointerTo[uint](789),
				Namespace:             "a-namespace",
				EnvironmentID:         testutils.PointerTo[uint](135),
				DestinationType:       "environment",
				HelmfileRef:           testutils.PointerTo("HEAD"),
				TargetAppVersionUse:   testutils.PointerTo("branch"),
				TargetChartVersionUse: testutils.PointerTo("latest"),
			}},
			wantErr: true,
		},
		{
			name: "app version use commit but no commit provided",
			args: args{chartRelease: ChartRelease{
				Name:                  "foobar",
				ChartID:               246,
				ClusterID:             testutils.PointerTo[uint](789),
				Namespace:             "a-namespace",
				EnvironmentID:         testutils.PointerTo[uint](135),
				DestinationType:       "environment",
				HelmfileRef:           testutils.PointerTo("HEAD"),
				TargetAppVersionUse:   testutils.PointerTo("commit"),
				TargetChartVersionUse: testutils.PointerTo("latest"),
			}},
			wantErr: true,
		},
		{
			name: "app version use exact but no exact provided",
			args: args{chartRelease: ChartRelease{
				Name:                  "foobar",
				ChartID:               246,
				ClusterID:             testutils.PointerTo[uint](789),
				Namespace:             "a-namespace",
				EnvironmentID:         testutils.PointerTo[uint](135),
				DestinationType:       "environment",
				HelmfileRef:           testutils.PointerTo("HEAD"),
				TargetAppVersionUse:   testutils.PointerTo("exact"),
				TargetChartVersionUse: testutils.PointerTo("latest"),
			}},
			wantErr: true,
		},
		{
			name: "invalid app version use",
			args: args{chartRelease: ChartRelease{
				Name:                   "foobar",
				ChartID:                246,
				ClusterID:              testutils.PointerTo[uint](789),
				Namespace:              "a-namespace",
				EnvironmentID:          testutils.PointerTo[uint](135),
				DestinationType:        "environment",
				HelmfileRef:            testutils.PointerTo("HEAD"),
				TargetAppVersionUse:    testutils.PointerTo("something obviously incorrect"),
				TargetAppVersionBranch: testutils.PointerTo("main"),
				TargetChartVersionUse:  testutils.PointerTo("latest"),
			}},
			wantErr: true,
		},
		{
			name: "omitted chart version use",
			args: args{chartRelease: ChartRelease{
				Name:                   "foobar",
				ChartID:                246,
				ClusterID:              testutils.PointerTo[uint](789),
				Namespace:              "a-namespace",
				EnvironmentID:          testutils.PointerTo[uint](135),
				DestinationType:        "environment",
				HelmfileRef:            testutils.PointerTo("HEAD"),
				TargetAppVersionUse:    testutils.PointerTo("branch"),
				TargetAppVersionBranch: testutils.PointerTo("main"),
			}},
			wantErr: true,
		},
		{
			name: "invalid chart version use",
			args: args{chartRelease: ChartRelease{
				Name:                   "foobar",
				ChartID:                246,
				ClusterID:              testutils.PointerTo[uint](789),
				Namespace:              "a-namespace",
				EnvironmentID:          testutils.PointerTo[uint](135),
				DestinationType:        "environment",
				HelmfileRef:            testutils.PointerTo("HEAD"),
				TargetAppVersionUse:    testutils.PointerTo("branch"),
				TargetAppVersionBranch: testutils.PointerTo("main"),
				TargetChartVersionUse:  testutils.PointerTo("something obviously incorrect"),
			}},
			wantErr: true,
		},
		{
			name: "chart version use exact but no exact provided",
			args: args{chartRelease: ChartRelease{
				Name:                   "foobar",
				ChartID:                246,
				ClusterID:              testutils.PointerTo[uint](789),
				Namespace:              "a-namespace",
				EnvironmentID:          testutils.PointerTo[uint](135),
				DestinationType:        "environment",
				HelmfileRef:            testutils.PointerTo("HEAD"),
				TargetAppVersionUse:    testutils.PointerTo("branch"),
				TargetAppVersionBranch: testutils.PointerTo("main"),
				TargetChartVersionUse:  testutils.PointerTo("exact"),
			}},
			wantErr: true,
		},
		{
			name: "specifically empty (thus invalid) thelma mode",
			args: args{chartRelease: ChartRelease{
				Name:                   "foobar",
				ChartID:                246,
				ClusterID:              testutils.PointerTo[uint](789),
				Namespace:              "a-namespace",
				EnvironmentID:          testutils.PointerTo[uint](135),
				DestinationType:        "environment",
				HelmfileRef:            testutils.PointerTo("HEAD"),
				TargetAppVersionUse:    testutils.PointerTo("branch"),
				TargetAppVersionBranch: testutils.PointerTo("main"),
				TargetChartVersionUse:  testutils.PointerTo("latest"),
				ThelmaMode:             testutils.PointerTo(""),
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
