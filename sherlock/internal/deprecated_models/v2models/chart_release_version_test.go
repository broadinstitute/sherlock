package v2models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"testing"
	"time"

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
		},
		{
			name:    "chartReleaseVersionValidCommitLatest",
			wantErr: false,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("commit"),
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
		},
		{
			name:    "chartReleaseVersionValidExactLatest",
			wantErr: false,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("commit"),
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
		},
		{
			name:    "chartReleaseVersionValidNoneLatest",
			wantErr: false,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("none"),

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
		},
		{
			name:    "chartReleaseVersionValidBranchExact",
			wantErr: false,
			obj: ChartReleaseVersion{
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

				ChartVersionResolver: utils.PointerTo("exact"),
				ChartVersionExact:    utils.PointerTo("v0.0.100"),
				ChartVersion: &ChartVersion{
					Model:        gorm.Model{ID: 2},
					ChartVersion: "v0.0.100",
				},
				ChartVersionID: utils.PointerTo[uint](2),

				HelmfileRef:         utils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: utils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionValidCommitExact",
			wantErr: false,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("commit"),
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

				ChartVersionResolver: utils.PointerTo("exact"),
				ChartVersionExact:    utils.PointerTo("v0.0.100"),
				ChartVersion: &ChartVersion{
					Model:        gorm.Model{ID: 2},
					ChartVersion: "v0.0.100",
				},
				ChartVersionID: utils.PointerTo[uint](2),

				HelmfileRef:         utils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: utils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionValidExactExact",
			wantErr: false,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("commit"),
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

				ChartVersionResolver: utils.PointerTo("exact"),
				ChartVersionExact:    utils.PointerTo("v0.0.100"),
				ChartVersion: &ChartVersion{
					Model:        gorm.Model{ID: 2},
					ChartVersion: "v0.0.100",
				},
				ChartVersionID: utils.PointerTo[uint](2),

				HelmfileRef:         utils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: utils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionValidNoneExact",
			wantErr: false,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("none"),

				ChartVersionResolver: utils.PointerTo("exact"),
				ChartVersionExact:    utils.PointerTo("v0.0.100"),
				ChartVersion: &ChartVersion{
					Model:        gorm.Model{ID: 2},
					ChartVersion: "v0.0.100",
				},
				ChartVersionID: utils.PointerTo[uint](2),

				HelmfileRef:         utils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: utils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionValidBranchLatestMin",
			wantErr: false,
			obj: ChartReleaseVersion{
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
		},
		{
			name:    "chartReleaseVersionValidCommitLatestMin",
			wantErr: false,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("commit"),
				AppVersionExact:    utils.PointerTo("v1.2.3"),
				AppVersionCommit:   utils.PointerTo("a1b2c3d4"),
				AppVersion: &AppVersion{
					Model:      gorm.Model{ID: 1},
					AppVersion: "v1.2.3",
					GitCommit:  "a1b2c3d4",
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
		},
		{
			name:    "chartReleaseVersionValidExactLatestMin",
			wantErr: false,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("exact"),
				AppVersionExact:    utils.PointerTo("v1.2.3"),

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
		},
		{
			name:    "chartReleaseVersionValidNoneLatestMin",
			wantErr: false,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("none"),

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
		},
		{
			name:    "chartReleaseVersionValidBranchExactMin",
			wantErr: false,
			obj: ChartReleaseVersion{
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

				ChartVersionResolver: utils.PointerTo("exact"),
				ChartVersionExact:    utils.PointerTo("v0.0.100"),

				HelmfileRef:         utils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: utils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionValidCommitExactMin",
			wantErr: false,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("commit"),
				AppVersionExact:    utils.PointerTo("v1.2.3"),
				AppVersionCommit:   utils.PointerTo("a1b2c3d4"),
				AppVersion: &AppVersion{
					Model:      gorm.Model{ID: 1},
					AppVersion: "v1.2.3",
					GitCommit:  "a1b2c3d4",
				},
				AppVersionID: utils.PointerTo[uint](1),

				ChartVersionResolver: utils.PointerTo("exact"),
				ChartVersionExact:    utils.PointerTo("v0.0.100"),

				HelmfileRef:         utils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: utils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionValidExactExactMin",
			wantErr: false,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("exact"),
				AppVersionExact:    utils.PointerTo("v1.2.3"),

				ChartVersionResolver: utils.PointerTo("exact"),
				ChartVersionExact:    utils.PointerTo("v0.0.100"),

				HelmfileRef:         utils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: utils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionValidNoneExactMin",
			wantErr: false,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("none"),

				ChartVersionResolver: utils.PointerTo("exact"),
				ChartVersionExact:    utils.PointerTo("v0.0.100"),

				HelmfileRef:         utils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: utils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionValidNoFirecloudDevelopRef",
			wantErr: false,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("none"),

				ChartVersionResolver: utils.PointerTo("exact"),
				ChartVersionExact:    utils.PointerTo("v0.0.100"),

				HelmfileRef: utils.PointerTo("e5f6g7h8"),
			},
		},
		{
			name:    "chartReleaseVersionInvalidUnresolved",
			wantErr: true,
			obj: ChartReleaseVersion{
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
		},
		{
			name:    "chartReleaseVersionInvalidNoAppVersionResolver",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionExact:  utils.PointerTo("v1.2.3"),
				AppVersionCommit: utils.PointerTo("a1b2c3d4"),
				AppVersionBranch: utils.PointerTo("main"),
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
		},
		{
			name:    "chartReleaseVersionInvalidNoBranch",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("branch"),
				AppVersionExact:    utils.PointerTo("v1.2.3"),
				AppVersionCommit:   utils.PointerTo("a1b2c3d4"),
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
		},
		{
			name:    "chartReleaseVersionInvalidNoBranchMatch",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("branch"),
				AppVersionExact:    utils.PointerTo("v1.2.3"),
				AppVersionCommit:   utils.PointerTo("a1b2c3d4"),
				AppVersionBranch:   utils.PointerTo("main"),
				AppVersion: &AppVersion{
					Model:      gorm.Model{ID: 1},
					AppVersion: "v1.2.3",
					GitCommit:  "a1b2c3d4",
					GitBranch:  "branchy-branch",
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
		},
		{
			name:    "chartReleaseVersionInvalidNoBranchAppVersionMatch",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("branch"),
				AppVersionExact:    utils.PointerTo("v1.2.3"),
				AppVersionCommit:   utils.PointerTo("a1b2c3d4"),
				AppVersionBranch:   utils.PointerTo("main"),
				AppVersionID:       utils.PointerTo[uint](1),

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
		},
		{
			name:    "chartReleaseVersionInvalidNoBranchCommit",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("branch"),
				AppVersionExact:    utils.PointerTo("v1.2.3"),
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
		},
		{
			name:    "chartReleaseVersionInvalidNoBranchCommitMatch",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("branch"),
				AppVersionExact:    utils.PointerTo("v1.2.3"),
				AppVersionCommit:   utils.PointerTo("a1b2c3d4"),
				AppVersionBranch:   utils.PointerTo("main"),
				AppVersion: &AppVersion{
					Model:      gorm.Model{ID: 1},
					AppVersion: "v1.2.3",
					GitCommit:  "commitycommit",
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
		},
		{
			name:    "chartReleaseVersionInvalidNoBranchExact",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("branch"),
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
		},
		{
			name:    "chartReleaseVersionInvalidNoBranchExactMatch",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("branch"),
				AppVersionExact:    utils.PointerTo("v1.2.3"),
				AppVersionCommit:   utils.PointerTo("a1b2c3d4"),
				AppVersionBranch:   utils.PointerTo("main"),
				AppVersion: &AppVersion{
					Model:      gorm.Model{ID: 1},
					AppVersion: "v1.2.3-abc",
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
		},
		{
			name:    "chartReleaseVersionInvalidNoCommit",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("commit"),
				AppVersionExact:    utils.PointerTo("v1.2.3"),
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
		},
		{
			name:    "chartReleaseVersionInvalidNoCommitMatch",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("commit"),
				AppVersionExact:    utils.PointerTo("v1.2.3"),
				AppVersionCommit:   utils.PointerTo("a1b2c3d4"),
				AppVersionBranch:   utils.PointerTo("main"),
				AppVersion: &AppVersion{
					Model:      gorm.Model{ID: 1},
					AppVersion: "v1.2.3",
					GitCommit:  "commitycommit",
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
		},
		{
			name:    "chartReleaseVersionInvalidNoCommitAppVersionMatch",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("commit"),
				AppVersionExact:    utils.PointerTo("v1.2.3"),
				AppVersionCommit:   utils.PointerTo("a1b2c3d4"),
				AppVersionBranch:   utils.PointerTo("main"),
				AppVersionID:       utils.PointerTo[uint](1),

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
		},
		{
			name:    "chartReleaseVersionInvalidNoCommitExact",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("commit"),
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
		},
		{
			name:    "chartReleaseVersionInvalidNoCommitExactMatch",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("commit"),
				AppVersionExact:    utils.PointerTo("v1.2.3"),
				AppVersionCommit:   utils.PointerTo("a1b2c3d4"),
				AppVersionBranch:   utils.PointerTo("main"),
				AppVersion: &AppVersion{
					Model:      gorm.Model{ID: 1},
					AppVersion: "v1.2.3-abc",
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
		},
		{
			name:    "chartReleaseVersionInvalidNoExact",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("exact"),
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
		},
		{
			name:    "chartReleaseVersionInvalidNoExactMatch",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("exact"),
				AppVersionExact:    utils.PointerTo("v1.2.3"),
				AppVersionCommit:   utils.PointerTo("a1b2c3d4"),
				AppVersionBranch:   utils.PointerTo("main"),
				AppVersion: &AppVersion{
					Model:      gorm.Model{ID: 1},
					AppVersion: "v1.2.3-abc",
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
		},
		{
			name:    "chartReleaseVersionInvalidNoAppVersionFollowID",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("follow"),
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
		},
		{
			name:    "chartReleaseVersionInvalidNoneWithBranch",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("none"),
				AppVersionBranch:   utils.PointerTo("main"),

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
		},
		{
			name:    "chartReleaseVersionInvalidNoneWithCommit",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("none"),
				AppVersionCommit:   utils.PointerTo("a1b2c3d4"),

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
		},
		{
			name:    "chartReleaseVersionInvalidNoneWithExact",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("none"),
				AppVersionExact:    utils.PointerTo("v1.2.3"),

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
		},
		{
			name:    "chartReleaseVersionInvalidNoneWithMatch",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("none"),
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
		},
		{
			name:    "chartReleaseVersionInvalidNoneWithFollowID",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver:             utils.PointerTo("none"),
				AppVersionFollowChartReleaseID: utils.PointerTo[uint](1),

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
		},
		{
			name:    "chartReleaseVersionInvalidUnknownAppVersionResolver",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("some obviously incorrect value"),
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
		},
		{
			name:    "chartReleaseVersionInvalidAppVersionConflict",
			wantErr: true,
			obj: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("branch"),
				AppVersionExact:    utils.PointerTo("v1.2.3"),
				AppVersionCommit:   utils.PointerTo("a1b2c3d4"),
				AppVersionBranch:   utils.PointerTo("main"),
				AppVersion: &AppVersion{
					Model:      gorm.Model{ID: 111},
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
		},
		{
			name:    "chartReleaseVersionInvalidNoChartVersionResolver",
			wantErr: true,
			obj: ChartReleaseVersion{
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

				ChartVersionExact: utils.PointerTo("v0.0.100"),
				ChartVersion: &ChartVersion{
					Model:        gorm.Model{ID: 2},
					ChartVersion: "v0.0.100",
				},
				ChartVersionID: utils.PointerTo[uint](2),

				HelmfileRef:         utils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: utils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionInvalidNoLatestChartVersion",
			wantErr: true,
			obj: ChartReleaseVersion{
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

				HelmfileRef:         utils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: utils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionInvalidNoChartVersionFollowID",
			wantErr: true,
			obj: ChartReleaseVersion{
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

				ChartVersionResolver: utils.PointerTo("follow"),

				HelmfileRef:         utils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: utils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionInvalidExactConflict",
			wantErr: true,
			obj: ChartReleaseVersion{
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

				ChartVersionResolver: utils.PointerTo("exact"),
				ChartVersionExact:    utils.PointerTo("v0.0.111"),
				ChartVersion: &ChartVersion{
					Model:        gorm.Model{ID: 2},
					ChartVersion: "v0.0.100",
				},
				ChartVersionID: utils.PointerTo[uint](2),

				HelmfileRef:         utils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: utils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionInvalidUnknownChartVersionResolver",
			wantErr: true,
			obj: ChartReleaseVersion{
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

				ChartVersionResolver: utils.PointerTo("some obviously incorrect value"),
				ChartVersionExact:    utils.PointerTo("v0.0.100"),
				ChartVersion: &ChartVersion{
					Model:        gorm.Model{ID: 2},
					ChartVersion: "v0.0.100",
				},
				ChartVersionID: utils.PointerTo[uint](2),

				HelmfileRef:         utils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: utils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionInvalidNoExactChartVersion",
			wantErr: true,
			obj: ChartReleaseVersion{
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
				ChartVersion: &ChartVersion{
					Model:        gorm.Model{ID: 2},
					ChartVersion: "v0.0.100",
				},
				ChartVersionID: utils.PointerTo[uint](2),

				HelmfileRef:         utils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: utils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionInvalidChartVersionConflict",
			wantErr: true,
			obj: ChartReleaseVersion{
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
					Model:        gorm.Model{ID: 222},
					ChartVersion: "v0.0.100",
				},
				ChartVersionID: utils.PointerTo[uint](2),

				HelmfileRef:         utils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: utils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionValidNoHelmfileRef",
			wantErr: false,
			obj: ChartReleaseVersion{
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
				ChartVersionID:      utils.PointerTo[uint](2),
				FirecloudDevelopRef: utils.PointerTo("dev"),
			},
		},
		{
			name:    "chartReleaseVersionInvalidHelmfileRefEnabledNoHelmfileRef",
			wantErr: true,
			obj: ChartReleaseVersion{
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
				ChartVersionID:      utils.PointerTo[uint](2),
				FirecloudDevelopRef: utils.PointerTo("dev"),
				HelmfileRefEnabled:  utils.PointerTo(true),
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
		ResolvedAt                       *time.Time
		AppVersionResolver               *string
		AppVersionExact                  *string
		AppVersionBranch                 *string
		AppVersionCommit                 *string
		AppVersionFollowChartRelease     *ChartRelease
		AppVersionFollowChartReleaseID   *uint
		AppVersion                       *AppVersion
		AppVersionID                     *uint
		ChartVersionResolver             *string
		ChartVersionExact                *string
		ChartVersionFollowChartRelease   *ChartRelease
		ChartVersionFollowChartReleaseID *uint
		ChartVersion                     *ChartVersion
		ChartVersionID                   *uint
		HelmfileRef                      *string
		FirecloudDevelopRef              *string
		ThelmaMode                       *string
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
				ResolvedAt: utils.PointerTo(time.Now()),
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
			args: args{other: ChartReleaseVersion{
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
			}},
			want: true,
		},
		{
			name: "false for not equal",
			fields: fields{
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
			args: args{other: ChartReleaseVersion{
				ResolvedAt: utils.PointerTo(time.Now()),

				AppVersionResolver: utils.PointerTo("commit"),
				AppVersionExact:    utils.PointerTo("v1.2.3"),
				AppVersionCommit:   utils.PointerTo("a1b2c3d4"),
				AppVersion: &AppVersion{
					Model:      gorm.Model{ID: 1},
					AppVersion: "v1.2.3",
					GitCommit:  "a1b2c3d4",
				},
				AppVersionID: utils.PointerTo[uint](1),

				ChartVersionResolver: utils.PointerTo("exact"),
				ChartVersionExact:    utils.PointerTo("v0.0.100"),

				HelmfileRef:         utils.PointerTo("e5f6g7h8"),
				FirecloudDevelopRef: utils.PointerTo("dev"),
			}},
			want: false,
		},
		{
			name: "checks firecloud develop ref",
			fields: fields{
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
			args: args{other: ChartReleaseVersion{
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
				FirecloudDevelopRef: utils.PointerTo("prod"),
			}},
			want: false,
		},
		{
			name: "checks app version follow id",
			fields: fields{
				AppVersionFollowChartReleaseID: utils.PointerTo[uint](1),
			},
			args: args{other: ChartReleaseVersion{
				AppVersionFollowChartReleaseID: utils.PointerTo[uint](2),
			}},
			want: false,
		},
		{
			name: "checks chart version follow id",
			fields: fields{
				ChartVersionFollowChartReleaseID: utils.PointerTo[uint](1),
			},
			args: args{other: ChartReleaseVersion{
				ChartVersionFollowChartReleaseID: utils.PointerTo[uint](2),
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chartReleaseVersion := &ChartReleaseVersion{
				ResolvedAt:                       tt.fields.ResolvedAt,
				AppVersionResolver:               tt.fields.AppVersionResolver,
				AppVersionExact:                  tt.fields.AppVersionExact,
				AppVersionBranch:                 tt.fields.AppVersionBranch,
				AppVersionCommit:                 tt.fields.AppVersionCommit,
				AppVersionFollowChartRelease:     tt.fields.AppVersionFollowChartRelease,
				AppVersionFollowChartReleaseID:   tt.fields.AppVersionFollowChartReleaseID,
				AppVersion:                       tt.fields.AppVersion,
				AppVersionID:                     tt.fields.AppVersionID,
				ChartVersionResolver:             tt.fields.ChartVersionResolver,
				ChartVersionExact:                tt.fields.ChartVersionExact,
				ChartVersionFollowChartRelease:   tt.fields.ChartVersionFollowChartRelease,
				ChartVersionFollowChartReleaseID: tt.fields.ChartVersionFollowChartReleaseID,
				ChartVersion:                     tt.fields.ChartVersion,
				ChartVersionID:                   tt.fields.ChartVersionID,
				HelmfileRef:                      tt.fields.HelmfileRef,
				FirecloudDevelopRef:              tt.fields.FirecloudDevelopRef,
			}
			if got := chartReleaseVersion.equalTo(tt.args.other); got != tt.want {
				t.Errorf("equalTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
