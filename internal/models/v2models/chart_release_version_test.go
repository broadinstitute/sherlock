package v2models

import (
	"testing"
	"time"

	"github.com/broadinstitute/sherlock/internal/testutils"
	"gorm.io/gorm"
)

func TestChartReleaseVersion_validate(t *testing.T) {
	tests := []struct {
		name    string
		obj     ChartReleaseVersion
		wantErr bool
	}{
		{
			name:    "chartReleaseVersionValidBranchLatest",
			wantErr: false,
			obj: ChartReleaseVersion{
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
		},
		{
			name:    "chartReleaseVersionValidCommitLatest",
			wantErr: false,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("commit"),
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
		},
		{
			name:    "chartReleaseVersionValidExactLatest",
			wantErr: false,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("commit"),
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
		},
		{
			name:    "chartReleaseVersionValidNoneLatest",
			wantErr: false,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("none"),

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
		},
		{
			name:    "chartReleaseVersionValidBranchExact",
			wantErr: false,
			obj: ChartReleaseVersion{
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

				ChartVersionResolver: testutils.PointerTo("exact"),
				ChartVersionExact:    testutils.PointerTo("v0.0.100"),
				ChartVersion: &ChartVersion{
					Model:        gorm.Model{ID: 2},
					ChartVersion: "v0.0.100",
				},
				ChartVersionID: testutils.PointerTo[uint](2),

				HelmfileRef:         testutils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: testutils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionValidCommitExact",
			wantErr: false,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("commit"),
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

				ChartVersionResolver: testutils.PointerTo("exact"),
				ChartVersionExact:    testutils.PointerTo("v0.0.100"),
				ChartVersion: &ChartVersion{
					Model:        gorm.Model{ID: 2},
					ChartVersion: "v0.0.100",
				},
				ChartVersionID: testutils.PointerTo[uint](2),

				HelmfileRef:         testutils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: testutils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionValidExactExact",
			wantErr: false,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("commit"),
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

				ChartVersionResolver: testutils.PointerTo("exact"),
				ChartVersionExact:    testutils.PointerTo("v0.0.100"),
				ChartVersion: &ChartVersion{
					Model:        gorm.Model{ID: 2},
					ChartVersion: "v0.0.100",
				},
				ChartVersionID: testutils.PointerTo[uint](2),

				HelmfileRef:         testutils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: testutils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionValidNoneExact",
			wantErr: false,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("none"),

				ChartVersionResolver: testutils.PointerTo("exact"),
				ChartVersionExact:    testutils.PointerTo("v0.0.100"),
				ChartVersion: &ChartVersion{
					Model:        gorm.Model{ID: 2},
					ChartVersion: "v0.0.100",
				},
				ChartVersionID: testutils.PointerTo[uint](2),

				HelmfileRef:         testutils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: testutils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionValidBranchLatestMin",
			wantErr: false,
			obj: ChartReleaseVersion{
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
		},
		{
			name:    "chartReleaseVersionValidCommitLatestMin",
			wantErr: false,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("commit"),
				AppVersionExact:    testutils.PointerTo("v1.2.3"),
				AppVersionCommit:   testutils.PointerTo("a1b2c3d4"),
				AppVersion: &AppVersion{
					Model:      gorm.Model{ID: 1},
					AppVersion: "v1.2.3",
					GitCommit:  "a1b2c3d4",
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
		},
		{
			name:    "chartReleaseVersionValidExactLatestMin",
			wantErr: false,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("exact"),
				AppVersionExact:    testutils.PointerTo("v1.2.3"),

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
		},
		{
			name:    "chartReleaseVersionValidNoneLatestMin",
			wantErr: false,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("none"),

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
		},
		{
			name:    "chartReleaseVersionValidBranchExactMin",
			wantErr: false,
			obj: ChartReleaseVersion{
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

				ChartVersionResolver: testutils.PointerTo("exact"),
				ChartVersionExact:    testutils.PointerTo("v0.0.100"),

				HelmfileRef:         testutils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: testutils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionValidCommitExactMin",
			wantErr: false,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("commit"),
				AppVersionExact:    testutils.PointerTo("v1.2.3"),
				AppVersionCommit:   testutils.PointerTo("a1b2c3d4"),
				AppVersion: &AppVersion{
					Model:      gorm.Model{ID: 1},
					AppVersion: "v1.2.3",
					GitCommit:  "a1b2c3d4",
				},
				AppVersionID: testutils.PointerTo[uint](1),

				ChartVersionResolver: testutils.PointerTo("exact"),
				ChartVersionExact:    testutils.PointerTo("v0.0.100"),

				HelmfileRef:         testutils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: testutils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionValidExactExactMin",
			wantErr: false,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("exact"),
				AppVersionExact:    testutils.PointerTo("v1.2.3"),

				ChartVersionResolver: testutils.PointerTo("exact"),
				ChartVersionExact:    testutils.PointerTo("v0.0.100"),

				HelmfileRef:         testutils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: testutils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionValidNoneExactMin",
			wantErr: false,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("none"),

				ChartVersionResolver: testutils.PointerTo("exact"),
				ChartVersionExact:    testutils.PointerTo("v0.0.100"),

				HelmfileRef:         testutils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: testutils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionInvalidUnresolved",
			wantErr: true,
			obj: ChartReleaseVersion{
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
		},
		{
			name:    "chartReleaseVersionInvalidNoAppVersionResolver",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionExact:  testutils.PointerTo("v1.2.3"),
				AppVersionCommit: testutils.PointerTo("a1b2c3d4"),
				AppVersionBranch: testutils.PointerTo("main"),
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
		},
		{
			name:    "chartReleaseVersionInvalidNoBranch",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("branch"),
				AppVersionExact:    testutils.PointerTo("v1.2.3"),
				AppVersionCommit:   testutils.PointerTo("a1b2c3d4"),
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
		},
		{
			name:    "chartReleaseVersionInvalidNoBranchMatch",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("branch"),
				AppVersionExact:    testutils.PointerTo("v1.2.3"),
				AppVersionCommit:   testutils.PointerTo("a1b2c3d4"),
				AppVersionBranch:   testutils.PointerTo("main"),
				AppVersion: &AppVersion{
					Model:      gorm.Model{ID: 1},
					AppVersion: "v1.2.3",
					GitCommit:  "a1b2c3d4",
					GitBranch:  "branchy-branch",
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
		},
		{
			name:    "chartReleaseVersionInvalidNoBranchAppVersionMatch",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("branch"),
				AppVersionExact:    testutils.PointerTo("v1.2.3"),
				AppVersionCommit:   testutils.PointerTo("a1b2c3d4"),
				AppVersionBranch:   testutils.PointerTo("main"),
				AppVersionID:       testutils.PointerTo[uint](1),

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
		},
		{
			name:    "chartReleaseVersionInvalidNoBranchCommit",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("branch"),
				AppVersionExact:    testutils.PointerTo("v1.2.3"),
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
		},
		{
			name:    "chartReleaseVersionInvalidNoBranchCommitMatch",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("branch"),
				AppVersionExact:    testutils.PointerTo("v1.2.3"),
				AppVersionCommit:   testutils.PointerTo("a1b2c3d4"),
				AppVersionBranch:   testutils.PointerTo("main"),
				AppVersion: &AppVersion{
					Model:      gorm.Model{ID: 1},
					AppVersion: "v1.2.3",
					GitCommit:  "commitycommit",
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
		},
		{
			name:    "chartReleaseVersionInvalidNoBranchExact",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("branch"),
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
		},
		{
			name:    "chartReleaseVersionInvalidNoBranchExactMatch",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("branch"),
				AppVersionExact:    testutils.PointerTo("v1.2.3"),
				AppVersionCommit:   testutils.PointerTo("a1b2c3d4"),
				AppVersionBranch:   testutils.PointerTo("main"),
				AppVersion: &AppVersion{
					Model:      gorm.Model{ID: 1},
					AppVersion: "v1.2.3-abc",
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
		},
		{
			name:    "chartReleaseVersionInvalidNoCommit",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("commit"),
				AppVersionExact:    testutils.PointerTo("v1.2.3"),
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
		},
		{
			name:    "chartReleaseVersionInvalidNoCommitMatch",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("commit"),
				AppVersionExact:    testutils.PointerTo("v1.2.3"),
				AppVersionCommit:   testutils.PointerTo("a1b2c3d4"),
				AppVersionBranch:   testutils.PointerTo("main"),
				AppVersion: &AppVersion{
					Model:      gorm.Model{ID: 1},
					AppVersion: "v1.2.3",
					GitCommit:  "commitycommit",
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
		},
		{
			name:    "chartReleaseVersionInvalidNoCommitAppVersionMatch",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("commit"),
				AppVersionExact:    testutils.PointerTo("v1.2.3"),
				AppVersionCommit:   testutils.PointerTo("a1b2c3d4"),
				AppVersionBranch:   testutils.PointerTo("main"),
				AppVersionID:       testutils.PointerTo[uint](1),

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
		},
		{
			name:    "chartReleaseVersionInvalidNoCommitExact",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("commit"),
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
		},
		{
			name:    "chartReleaseVersionInvalidNoCommitExactMatch",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("commit"),
				AppVersionExact:    testutils.PointerTo("v1.2.3"),
				AppVersionCommit:   testutils.PointerTo("a1b2c3d4"),
				AppVersionBranch:   testutils.PointerTo("main"),
				AppVersion: &AppVersion{
					Model:      gorm.Model{ID: 1},
					AppVersion: "v1.2.3-abc",
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
		},
		{
			name:    "chartReleaseVersionInvalidNoExact",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("exact"),
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
		},
		{
			name:    "chartReleaseVersionInvalidNoExactMatch",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("exact"),
				AppVersionExact:    testutils.PointerTo("v1.2.3"),
				AppVersionCommit:   testutils.PointerTo("a1b2c3d4"),
				AppVersionBranch:   testutils.PointerTo("main"),
				AppVersion: &AppVersion{
					Model:      gorm.Model{ID: 1},
					AppVersion: "v1.2.3-abc",
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
		},
		{
			name:    "chartReleaseVersionInvalidNoneWithBranch",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("none"),
				AppVersionBranch:   testutils.PointerTo("main"),

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
		},
		{
			name:    "chartReleaseVersionInvalidNoneWithCommit",
			wantErr: true,
			obj: ChartReleaseVersion{
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
		},
		{
			name:    "chartReleaseVersionInvalidNoneWithExact",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("none"),
				AppVersionExact:    testutils.PointerTo("v1.2.3"),

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
		},
		{
			name:    "chartReleaseVersionInvalidNoneWithMatch",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("none"),
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
		},
		{
			name:    "chartReleaseVersionInvalidUnknownAppVersionResolver",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("some obviously incorrect value"),
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
		},
		{
			name:    "chartReleaseVersionInvalidAppVersionConflict",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("branch"),
				AppVersionExact:    testutils.PointerTo("v1.2.3"),
				AppVersionCommit:   testutils.PointerTo("a1b2c3d4"),
				AppVersionBranch:   testutils.PointerTo("main"),
				AppVersion: &AppVersion{
					Model:      gorm.Model{ID: 111},
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
		},
		{
			name:    "chartReleaseVersionInvalidNoChartVersionResolver",
			wantErr: true,
			obj: ChartReleaseVersion{
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

				ChartVersionExact: testutils.PointerTo("v0.0.100"),
				ChartVersion: &ChartVersion{
					Model:        gorm.Model{ID: 2},
					ChartVersion: "v0.0.100",
				},
				ChartVersionID: testutils.PointerTo[uint](2),

				HelmfileRef:         testutils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: testutils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionInvalidNoLatestChartVersion",
			wantErr: true,
			obj: ChartReleaseVersion{
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

				HelmfileRef:         testutils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: testutils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionInvalidExactConflict",
			wantErr: true,
			obj: ChartReleaseVersion{
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

				ChartVersionResolver: testutils.PointerTo("exact"),
				ChartVersionExact:    testutils.PointerTo("v0.0.111"),
				ChartVersion: &ChartVersion{
					Model:        gorm.Model{ID: 2},
					ChartVersion: "v0.0.100",
				},
				ChartVersionID: testutils.PointerTo[uint](2),

				HelmfileRef:         testutils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: testutils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionInvalidUnknownChartVersionResolver",
			wantErr: true,
			obj: ChartReleaseVersion{
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

				ChartVersionResolver: testutils.PointerTo("some obviously incorrect value"),
				ChartVersionExact:    testutils.PointerTo("v0.0.100"),
				ChartVersion: &ChartVersion{
					Model:        gorm.Model{ID: 2},
					ChartVersion: "v0.0.100",
				},
				ChartVersionID: testutils.PointerTo[uint](2),

				HelmfileRef:         testutils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: testutils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionInvalidNoExactChartVersion",
			wantErr: true,
			obj: ChartReleaseVersion{
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
				ChartVersion: &ChartVersion{
					Model:        gorm.Model{ID: 2},
					ChartVersion: "v0.0.100",
				},
				ChartVersionID: testutils.PointerTo[uint](2),

				HelmfileRef:         testutils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: testutils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionInvalidChartVersionConflict",
			wantErr: true,
			obj: ChartReleaseVersion{
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
					Model:        gorm.Model{ID: 222},
					ChartVersion: "v0.0.100",
				},
				ChartVersionID: testutils.PointerTo[uint](2),

				HelmfileRef:         testutils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: testutils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionInvalidNoHelmfileRef",
			wantErr: true,
			obj: ChartReleaseVersion{
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.obj.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChartReleaseVersion_equalTo(t *testing.T) {
	type fields struct {
		ResolvedAt           *time.Time
		AppVersionResolver   *string
		AppVersionExact      *string
		AppVersionBranch     *string
		AppVersionCommit     *string
		AppVersion           *AppVersion
		AppVersionID         *uint
		ChartVersionResolver *string
		ChartVersionExact    *string
		ChartVersion         *ChartVersion
		ChartVersionID       *uint
		HelmfileRef          *string
		FirecloudDevelopRef  *string
		ThelmaMode           *string
	}
	type args struct {
		other ChartReleaseVersion
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "ignores ResolvedAt field",
			fields: fields{
				ResolvedAt: nil,
			},
			args: args{other: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),
			}},
			want: true,
		},
		{
			name: "empty equal to empty",
			args: args{other: ChartReleaseVersion{}},
			want: true,
		},
		{
			name: "true for equal",
			fields: fields{
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
			args: args{other: ChartReleaseVersion{
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
			}},
			want: true,
		},
		{
			name: "false for not equal",
			fields: fields{
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
			args: args{other: ChartReleaseVersion{
				ResolvedAt: testutils.PointerTo(time.Now()),

				AppVersionResolver: testutils.PointerTo("commit"),
				AppVersionExact:    testutils.PointerTo("v1.2.3"),
				AppVersionCommit:   testutils.PointerTo("a1b2c3d4"),
				AppVersion: &AppVersion{
					Model:      gorm.Model{ID: 1},
					AppVersion: "v1.2.3",
					GitCommit:  "a1b2c3d4",
				},
				AppVersionID: testutils.PointerTo[uint](1),

				ChartVersionResolver: testutils.PointerTo("exact"),
				ChartVersionExact:    testutils.PointerTo("v0.0.100"),

				HelmfileRef:         testutils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: testutils.PointerTo("dev"),
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chartReleaseVersion := &ChartReleaseVersion{
				ResolvedAt:           tt.fields.ResolvedAt,
				AppVersionResolver:   tt.fields.AppVersionResolver,
				AppVersionExact:      tt.fields.AppVersionExact,
				AppVersionBranch:     tt.fields.AppVersionBranch,
				AppVersionCommit:     tt.fields.AppVersionCommit,
				AppVersion:           tt.fields.AppVersion,
				AppVersionID:         tt.fields.AppVersionID,
				ChartVersionResolver: tt.fields.ChartVersionResolver,
				ChartVersionExact:    tt.fields.ChartVersionExact,
				ChartVersion:         tt.fields.ChartVersion,
				ChartVersionID:       tt.fields.ChartVersionID,
				HelmfileRef:          tt.fields.HelmfileRef,
				FirecloudDevelopRef:  tt.fields.FirecloudDevelopRef,
			}
			if got := chartReleaseVersion.equalTo(tt.args.other); got != tt.want {
				t.Errorf("equalTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
