package models

import (
	"cmp"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func (s *modelSuite) TestChangesetToResolvedAtPresent() {
	s.SetSuitableTestUserForDB()
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	err := s.DB.Model(&changeset).Select("to_resolved_at").Updates(&Changeset{To: ChartReleaseVersion{ResolvedAt: nil}}).Error
	s.ErrorContains(err, "violates check constraint \"to_resolved_at_present\"")
}

func (s *modelSuite) TestChangesetToAppVersionResolverEmpty() {
	s.SetSuitableTestUserForDB()
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	err := s.DB.Model(&changeset).Select("to_app_version_resolver").Updates(&Changeset{To: ChartReleaseVersion{AppVersionResolver: utils.PointerTo("")}}).Error
	s.ErrorContains(err, "violates check constraint \"to_app_version_resolver_valid\"")
}

func (s *modelSuite) TestChangesetToAppVersionResolverBranchBranchNull() {
	s.SetSuitableTestUserForDB()
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	err := s.DB.Model(&changeset).Select("to_app_version_resolver").Updates(&Changeset{To: ChartReleaseVersion{AppVersionResolver: utils.PointerTo("branch")}}).Error
	s.NoError(err)
	err = s.DB.Model(&changeset).Select("to_app_version_branch").Updates(&Changeset{To: ChartReleaseVersion{AppVersionBranch: nil}}).Error
	s.ErrorContains(err, "violates check constraint \"to_app_version_resolver_valid\"")
}

func (s *modelSuite) TestChangesetToAppVersionResolverBranchBranchEmpty() {
	s.SetSuitableTestUserForDB()
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	err := s.DB.Model(&changeset).Select("to_app_version_resolver").Updates(&Changeset{To: ChartReleaseVersion{AppVersionResolver: utils.PointerTo("branch")}}).Error
	s.NoError(err)
	err = s.DB.Model(&changeset).Select("to_app_version_branch").Updates(&Changeset{To: ChartReleaseVersion{AppVersionBranch: utils.PointerTo("")}}).Error
	s.ErrorContains(err, "violates check constraint \"to_app_version_resolver_valid\"")
}

func (s *modelSuite) TestChangesetToAppVersionResolverBranchAppVersionIDNull() {
	s.SetSuitableTestUserForDB()
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	err := s.DB.Model(&changeset).Select("to_app_version_resolver").Updates(&Changeset{To: ChartReleaseVersion{AppVersionResolver: utils.PointerTo("branch")}}).Error
	s.NoError(err)
	err = s.DB.Model(&changeset).Select("to_app_version_id").Updates(&Changeset{To: ChartReleaseVersion{AppVersionID: nil}}).Error
	s.ErrorContains(err, "violates check constraint \"to_app_version_resolver_valid\"")
}

func (s *modelSuite) TestChangesetToAppVersionResolverBranchCommitNull() {
	s.SetSuitableTestUserForDB()
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	err := s.DB.Model(&changeset).Select("to_app_version_resolver").Updates(&Changeset{To: ChartReleaseVersion{AppVersionResolver: utils.PointerTo("branch")}}).Error
	s.NoError(err)
	err = s.DB.Model(&changeset).Select("to_app_version_commit").Updates(&Changeset{To: ChartReleaseVersion{AppVersionCommit: nil}}).Error
	s.ErrorContains(err, "violates check constraint \"to_app_version_resolver_valid\"")
}

func (s *modelSuite) TestChangesetToAppVersionResolverBranchCommitEmpty() {
	s.SetSuitableTestUserForDB()
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	err := s.DB.Model(&changeset).Select("to_app_version_resolver").Updates(&Changeset{To: ChartReleaseVersion{AppVersionResolver: utils.PointerTo("branch")}}).Error
	s.NoError(err)
	err = s.DB.Model(&changeset).Select("to_app_version_commit").Updates(&Changeset{To: ChartReleaseVersion{AppVersionCommit: utils.PointerTo("")}}).Error
	s.ErrorContains(err, "violates check constraint \"to_app_version_resolver_valid\"")
}

func (s *modelSuite) TestChangesetToAppVersionResolverBranchExactNull() {
	s.SetSuitableTestUserForDB()
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	err := s.DB.Model(&changeset).Select("to_app_version_resolver").Updates(&Changeset{To: ChartReleaseVersion{AppVersionResolver: utils.PointerTo("branch")}}).Error
	s.NoError(err)
	err = s.DB.Model(&changeset).Select("to_app_version_exact").Updates(&Changeset{To: ChartReleaseVersion{AppVersionExact: nil}}).Error
	s.ErrorContains(err, "violates check constraint \"to_app_version_resolver_valid\"")
}

func (s *modelSuite) TestChangesetToAppVersionResolverBranchExactEmpty() {
	s.SetSuitableTestUserForDB()
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	err := s.DB.Model(&changeset).Select("to_app_version_resolver").Updates(&Changeset{To: ChartReleaseVersion{AppVersionResolver: utils.PointerTo("branch")}}).Error
	s.NoError(err)
	err = s.DB.Model(&changeset).Select("to_app_version_exact").Updates(&Changeset{To: ChartReleaseVersion{AppVersionExact: utils.PointerTo("")}}).Error
	s.ErrorContains(err, "violates check constraint \"to_app_version_resolver_valid\"")
}

func (s *modelSuite) TestChangesetToAppVersionResolverCommitCommitNull() {
	s.SetSuitableTestUserForDB()
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	err := s.DB.Model(&changeset).Select("to_app_version_resolver").Updates(&Changeset{To: ChartReleaseVersion{AppVersionResolver: utils.PointerTo("commit")}}).Error
	s.NoError(err)
	err = s.DB.Model(&changeset).Select("to_app_version_commit").Updates(&Changeset{To: ChartReleaseVersion{AppVersionCommit: nil}}).Error
	s.ErrorContains(err, "violates check constraint \"to_app_version_resolver_valid\"")
}

func (s *modelSuite) TestChangesetToAppVersionResolverCommitCommitEmpty() {
	s.SetSuitableTestUserForDB()
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	err := s.DB.Model(&changeset).Select("to_app_version_resolver").Updates(&Changeset{To: ChartReleaseVersion{AppVersionResolver: utils.PointerTo("commit")}}).Error
	s.NoError(err)
	err = s.DB.Model(&changeset).Select("to_app_version_commit").Updates(&Changeset{To: ChartReleaseVersion{AppVersionCommit: utils.PointerTo("")}}).Error
	s.ErrorContains(err, "violates check constraint \"to_app_version_resolver_valid\"")
}

func (s *modelSuite) TestChangesetToAppVersionResolverCommitExactNull() {
	s.SetSuitableTestUserForDB()
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	err := s.DB.Model(&changeset).Select("to_app_version_resolver").Updates(&Changeset{To: ChartReleaseVersion{AppVersionResolver: utils.PointerTo("commit")}}).Error
	s.NoError(err)
	err = s.DB.Model(&changeset).Select("to_app_version_exact").Updates(&Changeset{To: ChartReleaseVersion{AppVersionExact: nil}}).Error
	s.ErrorContains(err, "violates check constraint \"to_app_version_resolver_valid\"")
}

func (s *modelSuite) TestChangesetToAppVersionResolverCommitExactEmpty() {
	s.SetSuitableTestUserForDB()
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	err := s.DB.Model(&changeset).Select("to_app_version_resolver").Updates(&Changeset{To: ChartReleaseVersion{AppVersionResolver: utils.PointerTo("commit")}}).Error
	s.NoError(err)
	err = s.DB.Model(&changeset).Select("to_app_version_exact").Updates(&Changeset{To: ChartReleaseVersion{AppVersionExact: utils.PointerTo("")}}).Error
	s.ErrorContains(err, "violates check constraint \"to_app_version_resolver_valid\"")
}

func (s *modelSuite) TestChangesetToAppVersionResolverExactNull() {
	s.SetSuitableTestUserForDB()
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	err := s.DB.Model(&changeset).Select("to_app_version_exact").Updates(&Changeset{To: ChartReleaseVersion{AppVersionExact: nil}}).Error
	s.ErrorContains(err, "violates check constraint \"to_app_version_resolver_valid\"")
}

func (s *modelSuite) TestChangesetToAppVersionResolverExactEmpty() {
	s.SetSuitableTestUserForDB()
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	err := s.DB.Model(&changeset).Select("to_app_version_exact").Updates(&Changeset{To: ChartReleaseVersion{AppVersionExact: utils.PointerTo("")}}).Error
	s.ErrorContains(err, "violates check constraint \"to_app_version_resolver_valid\"")
}

func (s *modelSuite) TestChangesetToAppVersionResolverFollowNull() {
	s.SetSuitableTestUserForDB()
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	changeset.To.AppVersionFollowChartReleaseID = utils.PointerTo(s.TestData.ChartRelease_LeonardoStaging().ID)
	changeset.To.AppVersionResolver = utils.PointerTo("follow")
	s.NoError(changeset.To.resolve(s.DB, s.TestData.Chart_Leonardo().ID))
	s.NoError(s.DB.Save(&changeset).Error)
	err := s.DB.Model(&changeset).Select("to_app_version_follow_chart_release_id").Updates(&Changeset{To: ChartReleaseVersion{AppVersionFollowChartReleaseID: nil}}).Error
	s.ErrorContains(err, "violates check constraint \"to_app_version_resolver_valid\"")
}

func (s *modelSuite) TestChangesetToAppVersionResolverNoneBranchNotNullEmpty() {
	s.SetSuitableTestUserForDB()
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	changeset.To.AppVersionResolver = utils.PointerTo("none")
	s.NoError(changeset.To.resolve(s.DB, s.TestData.Chart_Leonardo().ID))
	s.NoError(s.DB.Save(&changeset).Error)
	err := s.DB.Model(&changeset).Select("to_app_version_branch").Updates(&Changeset{To: ChartReleaseVersion{AppVersionBranch: utils.PointerTo("dev")}}).Error
	s.ErrorContains(err, "violates check constraint \"to_app_version_resolver_valid\"")
}

func (s *modelSuite) TestChangesetToAppVersionResolverNoneCommitNotNullEmpty() {
	s.SetSuitableTestUserForDB()
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	changeset.To.AppVersionResolver = utils.PointerTo("none")
	s.NoError(changeset.To.resolve(s.DB, s.TestData.Chart_Leonardo().ID))
	s.NoError(s.DB.Save(&changeset).Error)
	err := s.DB.Model(&changeset).Select("to_app_version_commit").Updates(&Changeset{To: ChartReleaseVersion{AppVersionCommit: utils.PointerTo("commit")}}).Error
	s.ErrorContains(err, "violates check constraint \"to_app_version_resolver_valid\"")
}

func (s *modelSuite) TestChangesetToAppVersionResolverNoneExactNotNullEmpty() {
	s.SetSuitableTestUserForDB()
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	changeset.To.AppVersionResolver = utils.PointerTo("none")
	s.NoError(changeset.To.resolve(s.DB, s.TestData.Chart_Leonardo().ID))
	s.NoError(s.DB.Save(&changeset).Error)
	err := s.DB.Model(&changeset).Select("to_app_version_exact").Updates(&Changeset{To: ChartReleaseVersion{AppVersionExact: utils.PointerTo("exact")}}).Error
	s.ErrorContains(err, "violates check constraint \"to_app_version_resolver_valid\"")
}

func (s *modelSuite) TestChangesetToAppVersionResolverNoneIDNotNull() {
	s.SetSuitableTestUserForDB()
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	changeset.To.AppVersionResolver = utils.PointerTo("none")
	s.NoError(changeset.To.resolve(s.DB, s.TestData.Chart_Leonardo().ID))
	s.NoError(s.DB.Save(&changeset).Error)
	err := s.DB.Model(&changeset).Select("to_app_version_id").Updates(&Changeset{To: ChartReleaseVersion{AppVersionID: utils.PointerTo(uint(1))}}).Error
	s.ErrorContains(err, "violates check constraint \"to_app_version_resolver_valid\"")
}

func (s *modelSuite) TestChangesetToAppVersionResolverNoneFollowIDNotNull() {
	s.SetSuitableTestUserForDB()
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	changeset.To.AppVersionResolver = utils.PointerTo("none")
	s.NoError(changeset.To.resolve(s.DB, s.TestData.Chart_Leonardo().ID))
	s.NoError(s.DB.Save(&changeset).Error)
	err := s.DB.Model(&changeset).Select("to_app_version_follow_chart_release_id").Updates(&Changeset{To: ChartReleaseVersion{AppVersionFollowChartReleaseID: utils.PointerTo(uint(1))}}).Error
	s.ErrorContains(err, "violates check constraint \"to_app_version_resolver_valid\"")
}

func (s *modelSuite) TestChangesetToChartVersionResolverNotNull() {
	s.SetSuitableTestUserForDB()
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	err := s.DB.Model(&changeset).Select("to_chart_version_resolver").Updates(&Changeset{To: ChartReleaseVersion{ChartVersionResolver: nil}}).Error
	s.ErrorContains(err, "violates check constraint \"to_chart_version_resolver_valid\"")
}

func (s *modelSuite) TestChangesetToChartVersionResolverNotEmpty() {
	s.SetSuitableTestUserForDB()
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	err := s.DB.Model(&changeset).Select("to_chart_version_resolver").Updates(&Changeset{To: ChartReleaseVersion{ChartVersionResolver: utils.PointerTo("")}}).Error
	s.ErrorContains(err, "violates check constraint \"to_chart_version_resolver_valid\"")
}

func (s *modelSuite) TestChangesetToChartVersionExactNotNull() {
	s.SetSuitableTestUserForDB()
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	err := s.DB.Model(&changeset).Select("to_chart_version_exact").Updates(&Changeset{To: ChartReleaseVersion{ChartVersionExact: nil}}).Error
	s.ErrorContains(err, "violates check constraint \"to_chart_version_resolver_valid\"")
}

func (s *modelSuite) TestChangesetToChartVersionExactNotEmpty() {
	s.SetSuitableTestUserForDB()
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	err := s.DB.Model(&changeset).Select("to_chart_version_exact").Updates(&Changeset{To: ChartReleaseVersion{ChartVersionExact: utils.PointerTo("")}}).Error
	s.ErrorContains(err, "violates check constraint \"to_chart_version_resolver_valid\"")
}

func (s *modelSuite) TestChangesetToHelmfileRefValidRefNull() {
	s.SetSuitableTestUserForDB()
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	err := s.DB.Model(&changeset).Select("to_helmfile_ref_enabled").Updates(&Changeset{To: ChartReleaseVersion{HelmfileRefEnabled: utils.PointerTo(true)}}).Error
	s.NoError(err)
	err = s.DB.Model(&changeset).Select("to_helmfile_ref").Updates(&Changeset{To: ChartReleaseVersion{HelmfileRef: nil}}).Error
	s.ErrorContains(err, "violates check constraint \"to_helmfile_ref_valid\"")
}

func (s *modelSuite) TestChangesetToHelmfileRefValidRefFalse() {
	s.SetSuitableTestUserForDB()
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	err := s.DB.Model(&changeset).Select("to_helmfile_ref_enabled").Updates(&Changeset{To: ChartReleaseVersion{HelmfileRefEnabled: utils.PointerTo(true)}}).Error
	s.NoError(err)
	err = s.DB.Model(&changeset).Select("to_helmfile_ref").Updates(&Changeset{To: ChartReleaseVersion{HelmfileRef: utils.PointerTo("")}}).Error
	s.ErrorContains(err, "violates check constraint \"to_helmfile_ref_valid\"")
}

func TestChangeset_Summarize(t *testing.T) {
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
		PlannedBy        *User
		PlannedByID      *uint
		AppliedBy        *User
		AppliedByID      *uint
	}
	type args struct {
		includeFrom bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "empty",
			fields: fields{},
			args:   args{includeFrom: false},
			want:   "configuration change",
		},
		{
			name: "app change, no include from",
			fields: fields{
				To: ChartReleaseVersion{
					AppVersionExact: utils.PointerTo("1.2.3"),
				},
				From: ChartReleaseVersion{
					AppVersionExact: utils.PointerTo("1.2.2"),
				},
			},
			args: args{includeFrom: false},
			want: "app⭢1.2.3",
		},
		{
			name: "app change, include from",
			fields: fields{
				To: ChartReleaseVersion{
					AppVersionExact: utils.PointerTo("1.2.3"),
				},
				From: ChartReleaseVersion{
					AppVersionExact: utils.PointerTo("1.2.2"),
				},
			},
			args: args{includeFrom: true},
			want: "app 1.2.2⭢1.2.3",
		},
		{
			name: "app change, include from but no data",
			fields: fields{
				To: ChartReleaseVersion{
					AppVersionExact: utils.PointerTo("1.2.3"),
				},
			},
			args: args{includeFrom: true},
			want: "app⭢1.2.3",
		},
		{
			name: "chart change, no include from",
			fields: fields{
				To: ChartReleaseVersion{
					ChartVersionExact: utils.PointerTo("1.2.3"),
				},
				From: ChartReleaseVersion{
					ChartVersionExact: utils.PointerTo("1.2.2"),
				},
			},
			args: args{includeFrom: false},
			want: "chart⭢1.2.3",
		},
		{
			name: "chart change, include from",
			fields: fields{
				To: ChartReleaseVersion{
					ChartVersionExact: utils.PointerTo("1.2.3"),
				},
				From: ChartReleaseVersion{
					ChartVersionExact: utils.PointerTo("1.2.2"),
				},
			},
			args: args{includeFrom: true},
			want: "chart 1.2.2⭢1.2.3",
		},
		{
			name: "chart change, include from but no data",
			fields: fields{
				To: ChartReleaseVersion{
					ChartVersionExact: utils.PointerTo("1.2.3"),
				},
			},
			args: args{includeFrom: true},
			want: "chart⭢1.2.3",
		},
		{
			name: "app and chart change, no include from",
			fields: fields{
				To: ChartReleaseVersion{
					AppVersionExact:   utils.PointerTo("1.2.3"),
					ChartVersionExact: utils.PointerTo("1.2.3"),
				},
				From: ChartReleaseVersion{
					AppVersionExact:   utils.PointerTo("1.2.2"),
					ChartVersionExact: utils.PointerTo("1.2.2"),
				},
			},
			args: args{includeFrom: false},
			want: "app⭢1.2.3, chart⭢1.2.3",
		},
		{
			name: "app and chart change, include from",
			fields: fields{
				To: ChartReleaseVersion{
					AppVersionExact:   utils.PointerTo("1.2.3"),
					ChartVersionExact: utils.PointerTo("1.2.3"),
				},
				From: ChartReleaseVersion{
					AppVersionExact:   utils.PointerTo("1.2.2"),
					ChartVersionExact: utils.PointerTo("1.2.2"),
				},
			},
			args: args{includeFrom: true},
			want: "app 1.2.2⭢1.2.3, chart 1.2.2⭢1.2.3",
		},
		{
			name: "app and chart change, include from but no data",
			fields: fields{
				To: ChartReleaseVersion{
					AppVersionExact:   utils.PointerTo("1.2.3"),
					ChartVersionExact: utils.PointerTo("1.2.3"),
				},
			},
			args: args{includeFrom: true},
			want: "app⭢1.2.3, chart⭢1.2.3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Changeset{
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
				PlannedBy:        tt.fields.PlannedBy,
				PlannedByID:      tt.fields.PlannedByID,
				AppliedBy:        tt.fields.AppliedBy,
				AppliedByID:      tt.fields.AppliedByID,
			}
			assert.Equalf(t, tt.want, c.Summarize(tt.args.includeFrom), "Summarize(%v)", tt.args.includeFrom)
		})
	}
}

func TestCompareChangesetsByName(t *testing.T) {
	type args struct {
		a Changeset
		b Changeset
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "both with nil chart releases",
			args: args{
				a: Changeset{},
				b: Changeset{},
			},
			want: 0,
		},
		{
			name: "a with nil chart release",
			args: args{
				a: Changeset{},
				b: Changeset{ChartRelease: &ChartRelease{Name: "b"}},
			},
			want: -1,
		},
		{
			name: "b with nil chart release",
			args: args{
				a: Changeset{ChartRelease: &ChartRelease{Name: "a"}},
				b: Changeset{},
			},
			want: 1,
		},
		{
			name: "compare chart release names",
			args: args{
				a: Changeset{ChartRelease: &ChartRelease{Name: "a"}},
				b: Changeset{ChartRelease: &ChartRelease{Name: "b"}},
			},
			want: cmp.Compare("a", "b"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CompareChangesetsByName(tt.args.a, tt.args.b), "CompareChangesetsByName(%v, %v)", tt.args.a, tt.args.b)
		})
	}
}

func TestUsersFromChangesets(t *testing.T) {
	type args struct {
		changesets []Changeset
	}
	tests := []struct {
		name string
		args args
		want []User
	}{
		{
			name: "empty",
			args: args{changesets: []Changeset{}},
			want: nil,
		},
		{
			name: "sample case",
			args: args{changesets: []Changeset{
				{AppliedBy: &User{Model: gorm.Model{ID: 1}}},
				{AppliedBy: &User{Model: gorm.Model{ID: 2}}, PlannedBy: &User{Model: gorm.Model{ID: 1}}},
				{AppliedBy: &User{Model: gorm.Model{ID: 1}}, PlannedBy: &User{Model: gorm.Model{ID: 1}}},
				{AppliedBy: &User{Model: gorm.Model{ID: 3}}, PlannedBy: &User{Model: gorm.Model{ID: 4}}},
				{PlannedBy: &User{Model: gorm.Model{ID: 4}}},
			}},
			want: []User{
				{Model: gorm.Model{ID: 1}},
				{Model: gorm.Model{ID: 2}},
				{Model: gorm.Model{ID: 3}},
				{Model: gorm.Model{ID: 4}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, UsersFromChangesets(tt.args.changesets), "UsersFromChangesets(%v)", tt.args.changesets)
		})
	}
}
