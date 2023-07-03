package v2models

import (
	"reflect"
	"testing"
	"time"

	"github.com/broadinstitute/sherlock/sherlock/internal/testutils"
	"gorm.io/gorm"
)

func Test_ChangesetSelectorToQuery(t *testing.T) {
	type args struct {
		db       *gorm.DB
		selector string
	}
	tests := []struct {
		name    string
		args    args
		want    Changeset
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
			want: Changeset{Model: gorm.Model{ID: 123}},
		},
		{
			name:    "invalid id",
			args:    args{selector: testutils.StringNumberTooBigForInt},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := changesetSelectorToQuery(tt.args.db, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChangesetSelectorToQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

func Test_ChangesetToSelectors(t *testing.T) {
	type args struct {
		Changeset *Changeset
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "none",
			args: args{Changeset: &Changeset{}},
			want: nil,
		},
		{
			name: "id",
			args: args{Changeset: &Changeset{Model: gorm.Model{ID: 123}}},
			want: []string{"123"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := changesetToSelectors(tt.args.Changeset)
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

func Test_validateChangeset(t *testing.T) {
	type args struct {
		changeset *Changeset
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "empty",
			args:    args{changeset: nil},
			wantErr: true,
		},
		{
			name: "empty from; there is no validation on from",
			args: args{changeset: &Changeset{
				To: ChartReleaseVersion{
					ResolvedAt: testutils.PointerTo(time.Now()),

					AppVersionResolver: testutils.PointerTo("branch"),
					AppVersionExact:    testutils.PointerTo("v1.2.3"),
					AppVersionCommit:   testutils.PointerTo("a1b2c3d4"),
					AppVersionBranch:   testutils.PointerTo("main"),
					AppVersion: &AppVersion{
						Model:      gorm.Model{ID: 1},
						AppVersion: "v1.2.3",
						GitCommit:  "a1b2c3d4",
						GitBranch:  "main",
					},
					AppVersionID: testutils.PointerTo[uint](1),

					ChartVersionResolver: testutils.PointerTo("latest"),
					ChartVersionExact:    testutils.PointerTo("v0.0.100"),
					ChartVersion: &ChartVersion{
						Model:        gorm.Model{ID: 2},
						ChartVersion: "v0.0.100",
					},
					ChartVersionID: testutils.PointerTo[uint](2),

					HelmfileRef:         testutils.PointerTo("e5f6g7h8"),
					FirecloudDevelopRef: testutils.PointerTo("dev"),
				},
			}},
			wantErr: false,
		},
		{
			name: "invalid from; there is no validation on from",
			args: args{changeset: &Changeset{
				From: ChartReleaseVersion{
					ResolvedAt: testutils.PointerTo(time.Now()),

					AppVersionResolver: testutils.PointerTo("none"),
					AppVersionCommit:   testutils.PointerTo("a1b2c3d4"),

					ChartVersionResolver: testutils.PointerTo("latest"),
					ChartVersionExact:    testutils.PointerTo("v0.0.100"),
					ChartVersion: &ChartVersion{
						Model:        gorm.Model{ID: 2},
						ChartVersion: "v0.0.100",
					},
					ChartVersionID: testutils.PointerTo[uint](2),

					HelmfileRef:         testutils.PointerTo("e5f6g7h8"),
					FirecloudDevelopRef: testutils.PointerTo("dev"),
				},
				To: ChartReleaseVersion{
					ResolvedAt: testutils.PointerTo(time.Now()),

					AppVersionResolver: testutils.PointerTo("branch"),
					AppVersionExact:    testutils.PointerTo("v1.2.3"),
					AppVersionCommit:   testutils.PointerTo("a1b2c3d4"),
					AppVersionBranch:   testutils.PointerTo("main"),
					AppVersion: &AppVersion{
						Model:      gorm.Model{ID: 1},
						AppVersion: "v1.2.3",
						GitCommit:  "a1b2c3d4",
						GitBranch:  "main",
					},
					AppVersionID: testutils.PointerTo[uint](1),

					ChartVersionResolver: testutils.PointerTo("latest"),
					ChartVersionExact:    testutils.PointerTo("v0.0.100"),
					ChartVersion: &ChartVersion{
						Model:        gorm.Model{ID: 2},
						ChartVersion: "v0.0.100",
					},
					ChartVersionID: testutils.PointerTo[uint](2),

					HelmfileRef:         testutils.PointerTo("e5f6g7h8"),
					FirecloudDevelopRef: testutils.PointerTo("dev"),
				},
			}},
			wantErr: false,
		},
		{
			name: "empty to",
			args: args{changeset: &Changeset{
				From: ChartReleaseVersion{
					ResolvedAt: testutils.PointerTo(time.Now()),

					AppVersionResolver: testutils.PointerTo("branch"),
					AppVersionExact:    testutils.PointerTo("v1.2.3"),
					AppVersionCommit:   testutils.PointerTo("a1b2c3d4"),
					AppVersionBranch:   testutils.PointerTo("main"),
					AppVersion: &AppVersion{
						Model:      gorm.Model{ID: 1},
						AppVersion: "v1.2.3",
						GitCommit:  "a1b2c3d4",
						GitBranch:  "main",
					},
					AppVersionID: testutils.PointerTo[uint](1),

					ChartVersionResolver: testutils.PointerTo("latest"),
					ChartVersionExact:    testutils.PointerTo("v0.0.100"),
					ChartVersion: &ChartVersion{
						Model:        gorm.Model{ID: 2},
						ChartVersion: "v0.0.100",
					},
					ChartVersionID: testutils.PointerTo[uint](2),

					HelmfileRef:         testutils.PointerTo("e5f6g7h8"),
					FirecloudDevelopRef: testutils.PointerTo("dev"),
				},
			}},
			wantErr: true,
		},
		{
			name: "invalid to",
			args: args{changeset: &Changeset{
				From: ChartReleaseVersion{
					ResolvedAt: testutils.PointerTo(time.Now()),

					AppVersionResolver: testutils.PointerTo("branch"),
					AppVersionExact:    testutils.PointerTo("v1.2.3"),
					AppVersionCommit:   testutils.PointerTo("a1b2c3d4"),
					AppVersionBranch:   testutils.PointerTo("main"),
					AppVersion: &AppVersion{
						Model:      gorm.Model{ID: 1},
						AppVersion: "v1.2.3",
						GitCommit:  "a1b2c3d4",
						GitBranch:  "main",
					},
					AppVersionID: testutils.PointerTo[uint](1),

					ChartVersionResolver: testutils.PointerTo("latest"),
					ChartVersionExact:    testutils.PointerTo("v0.0.100"),
					ChartVersion: &ChartVersion{
						Model:        gorm.Model{ID: 2},
						ChartVersion: "v0.0.100",
					},
					ChartVersionID: testutils.PointerTo[uint](2),

					HelmfileRef:         testutils.PointerTo("e5f6g7h8"),
					FirecloudDevelopRef: testutils.PointerTo("dev"),
				},
				To: ChartReleaseVersion{
					ResolvedAt: testutils.PointerTo(time.Now()),

					AppVersionResolver: testutils.PointerTo("none"),
					AppVersionCommit:   testutils.PointerTo("a1b2c3d4"),

					ChartVersionResolver: testutils.PointerTo("latest"),
					ChartVersionExact:    testutils.PointerTo("v0.0.100"),
					ChartVersion: &ChartVersion{
						Model:        gorm.Model{ID: 2},
						ChartVersion: "v0.0.100",
					},
					ChartVersionID: testutils.PointerTo[uint](2),

					HelmfileRef:         testutils.PointerTo("e5f6g7h8"),
					FirecloudDevelopRef: testutils.PointerTo("dev"),
				},
			}},
			wantErr: true,
		},
		{
			name: "valid",
			args: args{changeset: &Changeset{
				From: ChartReleaseVersion{
					ResolvedAt: testutils.PointerTo(time.Now()),

					AppVersionResolver: testutils.PointerTo("branch"),
					AppVersionExact:    testutils.PointerTo("v1.2.3"),
					AppVersionCommit:   testutils.PointerTo("a1b2c3d4"),
					AppVersionBranch:   testutils.PointerTo("main"),
					AppVersion: &AppVersion{
						Model:      gorm.Model{ID: 1},
						AppVersion: "v1.2.3",
						GitCommit:  "a1b2c3d4",
						GitBranch:  "main",
					},
					AppVersionID: testutils.PointerTo[uint](1),

					ChartVersionResolver: testutils.PointerTo("latest"),
					ChartVersionExact:    testutils.PointerTo("v0.0.100"),
					ChartVersion: &ChartVersion{
						Model:        gorm.Model{ID: 2},
						ChartVersion: "v0.0.100",
					},
					ChartVersionID: testutils.PointerTo[uint](2),

					HelmfileRef:         testutils.PointerTo("e5f6g7h8"),
					FirecloudDevelopRef: testutils.PointerTo("dev"),
				},
				To: ChartReleaseVersion{
					ResolvedAt: testutils.PointerTo(time.Now()),

					AppVersionResolver: testutils.PointerTo("branch"),
					AppVersionExact:    testutils.PointerTo("v1.2.3"),
					AppVersionCommit:   testutils.PointerTo("a1b2c3d4"),
					AppVersionBranch:   testutils.PointerTo("main"),
					AppVersion: &AppVersion{
						Model:      gorm.Model{ID: 1},
						AppVersion: "v1.2.3",
						GitCommit:  "a1b2c3d4",
						GitBranch:  "main",
					},
					AppVersionID: testutils.PointerTo[uint](1),

					ChartVersionResolver: testutils.PointerTo("latest"),
					ChartVersionExact:    testutils.PointerTo("v0.0.100"),
					ChartVersion: &ChartVersion{
						Model:        gorm.Model{ID: 2},
						ChartVersion: "v0.0.100",
					},
					ChartVersionID: testutils.PointerTo[uint](2),

					HelmfileRef:         testutils.PointerTo("e5f6g7h8"),
					FirecloudDevelopRef: testutils.PointerTo("dev"),
				},
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateChangeset(tt.args.changeset); (err != nil) != tt.wantErr {
				t.Errorf("validateChangeset() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChangeset_GetCiIdentifier(t *testing.T) {
	type fields struct {
		Model            gorm.Model
		CiIdentifier     *CiIdentifier
		ChartRelease     *ChartRelease
		ChartReleaseID   uint
		From             ChartReleaseVersion
		To               ChartReleaseVersion
		AppliedAt        *time.Time
		SupersededAt     *time.Time
		NewAppVersions   []*AppVersion
		NewChartVersions []*ChartVersion
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
				ResourceType: "changeset",
				ResourceID:   123,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Changeset{
				Model:            tt.fields.Model,
				CiIdentifier:     tt.fields.CiIdentifier,
				ChartRelease:     tt.fields.ChartRelease,
				ChartReleaseID:   tt.fields.ChartReleaseID,
				From:             tt.fields.From,
				To:               tt.fields.To,
				AppliedAt:        tt.fields.AppliedAt,
				SupersededAt:     tt.fields.SupersededAt,
				NewAppVersions:   tt.fields.NewAppVersions,
				NewChartVersions: tt.fields.NewChartVersions,
			}
			if got := c.GetCiIdentifier(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCiIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}
