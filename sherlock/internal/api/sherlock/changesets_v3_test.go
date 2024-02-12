package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func (s *handlerSuite) TestChangesetV3_toModel() {
	now := time.Now()
	type fields struct {
		CommonFields                       CommonFields
		CiIdentifier                       *CiIdentifierV3
		ChartReleaseInfo                   *ChartReleaseV3
		AppliedAt                          *time.Time
		SupersededAt                       *time.Time
		NewAppVersions                     []AppVersionV3
		NewChartVersions                   []ChartVersionV3
		FromResolvedAt                     *time.Time
		FromAppVersionResolver             *string
		FromAppVersionExact                *string
		FromAppVersionBranch               *string
		FromAppVersionCommit               *string
		FromAppVersionFollowChartRelease   string
		FromAppVersionReference            string
		FromChartVersionResolver           *string
		FromChartVersionExact              *string
		FromChartVersionFollowChartRelease string
		FromChartVersionReference          string
		FromHelmfileRef                    *string
		FromHelmfileRefEnabled             *bool
		ToResolvedAt                       *time.Time
		ToAppVersionReference              string
		ToChartVersionReference            string
		PlannedBy                          *string
		PlannedByInfo                      *UserV3
		AppliedBy                          *string
		AppliedByInfo                      *UserV3
		ChangesetV3Create                  ChangesetV3Create
	}
	tests := []struct {
		name    string
		fields  fields
		want    models.Changeset
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "empty",
			fields:  fields{},
			wantErr: assert.NoError,
			want:    models.Changeset{},
		},
		{
			name: "invalid chart release selector",
			fields: fields{
				ChangesetV3Create: ChangesetV3Create{
					ChartRelease: "!!!!",
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "not found chart release",
			fields: fields{
				ChangesetV3Create: ChangesetV3Create{
					ChartRelease: "not-found",
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "invalid to app version follow chart release selector",
			fields: fields{
				ChangesetV3Create: ChangesetV3Create{
					ToAppVersionFollowChartRelease: "!!!!",
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "not found to app version follow chart release",
			fields: fields{
				ChangesetV3Create: ChangesetV3Create{
					ToAppVersionFollowChartRelease: "not-found",
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "invalid to app version reference selector",
			fields: fields{
				ToAppVersionReference: "!!!!",
			},
			wantErr: assert.Error,
		},
		{
			name: "not found to app version reference selector",
			fields: fields{
				ToAppVersionReference: s.TestData.Chart_Leonardo().Name + "/not-found",
			},
			wantErr: assert.Error,
		},
		{
			name: "invalid to chart version follow chart release selector",
			fields: fields{
				ChangesetV3Create: ChangesetV3Create{
					ToChartVersionFollowChartRelease: "!!!!",
				},
			},
			wantErr: assert.Error,
		}, {
			name: "not found to chart version follow chart release",
			fields: fields{
				ChangesetV3Create: ChangesetV3Create{
					ToChartVersionFollowChartRelease: "not-found",
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "invalid to chart version reference selector",
			fields: fields{
				ToChartVersionReference: "!!!!",
			},
			wantErr: assert.Error,
		},
		{
			name: "not found to chart version reference selector",
			fields: fields{
				ToChartVersionReference: s.TestData.Chart_Leonardo().Name + "/not-found",
			},
			wantErr: assert.Error,
		},
		{
			name: "invalid from app version follow chart release selector",
			fields: fields{
				FromAppVersionFollowChartRelease: "!!!!",
			},
			wantErr: assert.Error,
		},
		{
			name: "not found from app version follow chart release",
			fields: fields{
				FromAppVersionFollowChartRelease: "not-found",
			},
			wantErr: assert.Error,
		},
		{
			name: "invalid from app version reference selector",
			fields: fields{
				FromAppVersionReference: "!!!!",
			},
			wantErr: assert.Error,
		},
		{
			name: "not found from app version reference selector",
			fields: fields{
				FromAppVersionReference: s.TestData.Chart_Leonardo().Name + "/not-found",
			},
			wantErr: assert.Error,
		},
		{
			name: "invalid from chart version follow chart release selector",
			fields: fields{
				FromChartVersionFollowChartRelease: "!!!!",
			},
			wantErr: assert.Error,
		}, {
			name: "not found from chart version follow chart release",
			fields: fields{
				FromChartVersionFollowChartRelease: "not-found",
			},
			wantErr: assert.Error,
		},
		{
			name: "invalid from chart version reference selector",
			fields: fields{
				FromChartVersionReference: "!!!!",
			},
			wantErr: assert.Error,
		},
		{
			name: "not found from chart version reference selector",
			fields: fields{
				FromChartVersionReference: s.TestData.Chart_Leonardo().Name + "/not-found",
			},
			wantErr: assert.Error,
		},
		{
			name: "invalid planned by selector",
			fields: fields{
				PlannedBy: utils.PointerTo("!!!!"),
			},
			wantErr: assert.Error,
		},
		{
			name: "not found planned by selector",
			fields: fields{
				PlannedBy: utils.PointerTo("not-found@example.com"),
			},
			wantErr: assert.Error,
		},
		{
			name: "invalid applied by selector",
			fields: fields{
				AppliedBy: utils.PointerTo("!!!!"),
			},
			wantErr: assert.Error,
		},
		{
			name: "not found applied by selector",
			fields: fields{
				AppliedBy: utils.PointerTo("not-found@example.com"),
			},
			wantErr: assert.Error,
		},
		{
			name: "full",
			fields: fields{
				CommonFields: CommonFields{
					ID:        1,
					CreatedAt: now.Add(-time.Hour),
					UpdatedAt: now.Add(-time.Minute),
				},
				AppliedAt:                          utils.PointerTo(now.Add(time.Hour)),
				SupersededAt:                       utils.PointerTo(now.Add(time.Minute)),
				FromResolvedAt:                     utils.PointerTo(now.Add(-time.Hour)),
				FromAppVersionResolver:             utils.PointerTo("resolver"),
				FromAppVersionExact:                utils.PointerTo("exact"),
				FromAppVersionBranch:               utils.PointerTo("branch"),
				FromAppVersionCommit:               utils.PointerTo("commit"),
				FromAppVersionFollowChartRelease:   s.TestData.ChartRelease_LeonardoStaging().Name,
				FromAppVersionReference:            s.TestData.Chart_Leonardo().Name + "/" + s.TestData.AppVersion_Leonardo_V3().AppVersion,
				FromChartVersionResolver:           utils.PointerTo("resolver"),
				FromChartVersionExact:              utils.PointerTo("exact"),
				FromChartVersionFollowChartRelease: s.TestData.ChartRelease_LeonardoStaging().Name,
				FromChartVersionReference:          s.TestData.Chart_Leonardo().Name + "/" + s.TestData.ChartVersion_Leonardo_V3().ChartVersion,
				FromHelmfileRef:                    utils.PointerTo("helmfile-ref"),
				FromHelmfileRefEnabled:             utils.PointerTo(true),
				ToResolvedAt:                       utils.PointerTo(now.Add(time.Hour)),
				ToAppVersionReference:              s.TestData.Chart_Leonardo().Name + "/" + s.TestData.AppVersion_Leonardo_V2().AppVersion,
				ToChartVersionReference:            s.TestData.Chart_Leonardo().Name + "/" + s.TestData.ChartVersion_Leonardo_V2().ChartVersion,
				PlannedBy:                          utils.PointerTo(s.TestData.User_NonSuitable().Email),
				AppliedBy:                          utils.PointerTo(s.TestData.User_Suitable().Email),
				ChangesetV3Create: ChangesetV3Create{
					ToAppVersionResolver:             utils.PointerTo("resolver"),
					ToAppVersionExact:                utils.PointerTo("exact"),
					ToAppVersionBranch:               utils.PointerTo("branch"),
					ToAppVersionCommit:               utils.PointerTo("commit"),
					ToAppVersionFollowChartRelease:   s.TestData.ChartRelease_LeonardoProd().Name,
					ToChartVersionResolver:           utils.PointerTo("resolver"),
					ToChartVersionExact:              utils.PointerTo("exact"),
					ToChartVersionFollowChartRelease: s.TestData.ChartRelease_LeonardoProd().Name,
					ToHelmfileRef:                    utils.PointerTo("helmfile-ref"),
					ToHelmfileRefEnabled:             utils.PointerTo(true),
					ChartRelease:                     s.TestData.ChartRelease_LeonardoDev().Name,
				},
			},
			wantErr: assert.NoError,
			want: models.Changeset{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: now.Add(-time.Hour),
					UpdatedAt: now.Add(-time.Minute),
				},
				ChartReleaseID: s.TestData.ChartRelease_LeonardoDev().ID,
				From: models.ChartReleaseVersion{
					ResolvedAt:                       utils.PointerTo(now.Add(-time.Hour)),
					AppVersionResolver:               utils.PointerTo("resolver"),
					AppVersionExact:                  utils.PointerTo("exact"),
					AppVersionBranch:                 utils.PointerTo("branch"),
					AppVersionCommit:                 utils.PointerTo("commit"),
					AppVersionFollowChartReleaseID:   utils.PointerTo(s.TestData.ChartRelease_LeonardoStaging().ID),
					AppVersionID:                     utils.PointerTo(s.TestData.AppVersion_Leonardo_V3().ID),
					ChartVersionResolver:             utils.PointerTo("resolver"),
					ChartVersionExact:                utils.PointerTo("exact"),
					ChartVersionFollowChartReleaseID: utils.PointerTo(s.TestData.ChartRelease_LeonardoStaging().ID),
					ChartVersionID:                   utils.PointerTo(s.TestData.ChartVersion_Leonardo_V3().ID),
					HelmfileRef:                      utils.PointerTo("helmfile-ref"),
					HelmfileRefEnabled:               utils.PointerTo(true),
				},
				To: models.ChartReleaseVersion{
					ResolvedAt:                       utils.PointerTo(now.Add(time.Hour)),
					AppVersionResolver:               utils.PointerTo("resolver"),
					AppVersionExact:                  utils.PointerTo("exact"),
					AppVersionBranch:                 utils.PointerTo("branch"),
					AppVersionCommit:                 utils.PointerTo("commit"),
					AppVersionFollowChartReleaseID:   utils.PointerTo(s.TestData.ChartRelease_LeonardoProd().ID),
					AppVersionID:                     utils.PointerTo(s.TestData.AppVersion_Leonardo_V2().ID),
					ChartVersionResolver:             utils.PointerTo("resolver"),
					ChartVersionExact:                utils.PointerTo("exact"),
					ChartVersionFollowChartReleaseID: utils.PointerTo(s.TestData.ChartRelease_LeonardoProd().ID),
					ChartVersionID:                   utils.PointerTo(s.TestData.ChartVersion_Leonardo_V2().ID),
					HelmfileRef:                      utils.PointerTo("helmfile-ref"),
					HelmfileRefEnabled:               utils.PointerTo(true),
				},
				AppliedAt:    utils.PointerTo(now.Add(time.Hour)),
				SupersededAt: utils.PointerTo(now.Add(time.Minute)),
				PlannedByID:  utils.PointerTo(s.TestData.User_NonSuitable().ID),
				AppliedByID:  utils.PointerTo(s.TestData.User_Suitable().ID),
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			c := ChangesetV3{
				CommonFields: tt.fields.CommonFields,
				ChangesetV3Query: ChangesetV3Query{
					CiIdentifier:                       tt.fields.CiIdentifier,
					ChartReleaseInfo:                   tt.fields.ChartReleaseInfo,
					AppliedAt:                          tt.fields.AppliedAt,
					SupersededAt:                       tt.fields.SupersededAt,
					NewAppVersions:                     tt.fields.NewAppVersions,
					NewChartVersions:                   tt.fields.NewChartVersions,
					FromResolvedAt:                     tt.fields.FromResolvedAt,
					FromAppVersionResolver:             tt.fields.FromAppVersionResolver,
					FromAppVersionExact:                tt.fields.FromAppVersionExact,
					FromAppVersionBranch:               tt.fields.FromAppVersionBranch,
					FromAppVersionCommit:               tt.fields.FromAppVersionCommit,
					FromAppVersionFollowChartRelease:   tt.fields.FromAppVersionFollowChartRelease,
					FromAppVersionReference:            tt.fields.FromAppVersionReference,
					FromChartVersionResolver:           tt.fields.FromChartVersionResolver,
					FromChartVersionExact:              tt.fields.FromChartVersionExact,
					FromChartVersionFollowChartRelease: tt.fields.FromChartVersionFollowChartRelease,
					FromChartVersionReference:          tt.fields.FromChartVersionReference,
					FromHelmfileRef:                    tt.fields.FromHelmfileRef,
					FromHelmfileRefEnabled:             tt.fields.FromHelmfileRefEnabled,
					ToResolvedAt:                       tt.fields.ToResolvedAt,
					ToAppVersionReference:              tt.fields.ToAppVersionReference,
					ToChartVersionReference:            tt.fields.ToChartVersionReference,
					PlannedBy:                          tt.fields.PlannedBy,
					PlannedByInfo:                      tt.fields.PlannedByInfo,
					AppliedBy:                          tt.fields.AppliedBy,
					AppliedByInfo:                      tt.fields.AppliedByInfo,
					ChangesetV3Create:                  tt.fields.ChangesetV3Create,
				},
			}
			got, err := c.toModel(s.DB)
			if !tt.wantErr(s.T(), err, "toModel()") {
				return
			}
			s.Equalf(tt.want, got, "toModel()")
		})
	}
}

func Test_changesetFromModel(t *testing.T) {
	now := time.Now()
	type args struct {
		model models.Changeset
	}
	tests := []struct {
		name string
		args args
		want ChangesetV3
	}{
		{
			name: "empty",
			args: args{},
			want: ChangesetV3{},
		},
		{
			name: "app version list",
			args: args{
				model: models.Changeset{
					NewAppVersions: []*models.AppVersion{
						nil,
						{
							Model: gorm.Model{
								ID: 1,
							},
						},
						nil,
					},
				},
			},
			want: ChangesetV3{
				ChangesetV3Query: ChangesetV3Query{
					NewAppVersions: []AppVersionV3{
						{CommonFields: CommonFields{ID: 1}},
					},
				},
			},
		},
		{
			name: "chart version list",
			args: args{
				model: models.Changeset{
					NewChartVersions: []*models.ChartVersion{
						nil,
						{
							Model: gorm.Model{
								ID: 1,
							},
						},
						nil,
					},
				},
			},
			want: ChangesetV3{
				ChangesetV3Query: ChangesetV3Query{
					NewChartVersions: []ChartVersionV3{
						{CommonFields: CommonFields{ID: 1}},
					},
				},
			},
		},
		{
			name: "from app version follow chart release name",
			args: args{
				model: models.Changeset{
					From: models.ChartReleaseVersion{
						AppVersionFollowChartRelease: &models.ChartRelease{
							Name: "name",
						},
					},
				},
			},
			want: ChangesetV3{
				ChangesetV3Query: ChangesetV3Query{
					FromAppVersionFollowChartRelease: "name",
				},
			},
		},
		{
			name: "from app version follow chart release id",
			args: args{
				model: models.Changeset{
					From: models.ChartReleaseVersion{
						AppVersionFollowChartReleaseID: utils.PointerTo[uint](1),
					},
				},
			},
			want: ChangesetV3{
				ChangesetV3Query: ChangesetV3Query{
					FromAppVersionFollowChartRelease: "1",
				},
			},
		},
		{
			name: "from app version chart reference",
			args: args{
				model: models.Changeset{
					From: models.ChartReleaseVersion{
						AppVersion: &models.AppVersion{
							Chart: &models.Chart{
								Name: "name",
							},
							AppVersion: "version",
						},
					},
				},
			},
			want: ChangesetV3{
				ChangesetV3Query: ChangesetV3Query{
					FromAppVersionReference: "name/version",
				},
			},
		},
		{
			name: "from app version chart id reference",
			args: args{
				model: models.Changeset{
					From: models.ChartReleaseVersion{
						AppVersion: &models.AppVersion{
							ChartID:    1,
							AppVersion: "version",
						},
					},
				},
			},
			want: ChangesetV3{
				ChangesetV3Query: ChangesetV3Query{
					FromAppVersionReference: "1/version",
				},
			},
		},
		{
			name: "from app version id reference",
			args: args{
				model: models.Changeset{
					From: models.ChartReleaseVersion{
						AppVersionID: utils.PointerTo[uint](1),
					},
				},
			},
			want: ChangesetV3{
				ChangesetV3Query: ChangesetV3Query{
					FromAppVersionReference: "1",
				},
			},
		},
		{
			name: "from chart version follow chart release name",
			args: args{
				model: models.Changeset{
					From: models.ChartReleaseVersion{
						ChartVersionFollowChartRelease: &models.ChartRelease{
							Name: "name",
						},
					},
				},
			},
			want: ChangesetV3{
				ChangesetV3Query: ChangesetV3Query{
					FromChartVersionFollowChartRelease: "name",
				},
			},
		},
		{
			name: "from chart version follow chart release id",
			args: args{
				model: models.Changeset{
					From: models.ChartReleaseVersion{
						ChartVersionFollowChartReleaseID: utils.PointerTo[uint](1),
					},
				},
			},
			want: ChangesetV3{
				ChangesetV3Query: ChangesetV3Query{
					FromChartVersionFollowChartRelease: "1",
				},
			},
		},
		{
			name: "from chart version chart reference",
			args: args{
				model: models.Changeset{
					From: models.ChartReleaseVersion{
						ChartVersion: &models.ChartVersion{
							Chart: &models.Chart{
								Name: "name",
							},
							ChartVersion: "version",
						},
					},
				},
			},
			want: ChangesetV3{
				ChangesetV3Query: ChangesetV3Query{
					FromChartVersionReference: "name/version",
				},
			},
		},
		{
			name: "from chart version chart id reference",
			args: args{
				model: models.Changeset{
					From: models.ChartReleaseVersion{
						ChartVersion: &models.ChartVersion{
							ChartID:      1,
							ChartVersion: "version",
						},
					},
				},
			},
			want: ChangesetV3{
				ChangesetV3Query: ChangesetV3Query{
					FromChartVersionReference: "1/version",
				},
			},
		},
		{
			name: "from chart version id reference",
			args: args{
				model: models.Changeset{
					From: models.ChartReleaseVersion{
						ChartVersionID: utils.PointerTo[uint](1),
					},
				},
			},
			want: ChangesetV3{
				ChangesetV3Query: ChangesetV3Query{
					FromChartVersionReference: "1",
				},
			},
		},

		{
			name: "to app version follow chart release name",
			args: args{
				model: models.Changeset{
					To: models.ChartReleaseVersion{
						AppVersionFollowChartRelease: &models.ChartRelease{
							Name: "name",
						},
					},
				},
			},
			want: ChangesetV3{
				ChangesetV3Query: ChangesetV3Query{
					ChangesetV3Create: ChangesetV3Create{
						ToAppVersionFollowChartRelease: "name",
					},
				},
			},
		},
		{
			name: "to app version follow chart release id",
			args: args{
				model: models.Changeset{
					To: models.ChartReleaseVersion{
						AppVersionFollowChartReleaseID: utils.PointerTo[uint](1),
					},
				},
			},
			want: ChangesetV3{
				ChangesetV3Query: ChangesetV3Query{
					ChangesetV3Create: ChangesetV3Create{
						ToAppVersionFollowChartRelease: "1",
					},
				},
			},
		},
		{
			name: "to app version chart reference",
			args: args{
				model: models.Changeset{
					To: models.ChartReleaseVersion{
						AppVersion: &models.AppVersion{
							Chart: &models.Chart{
								Name: "name",
							},
							AppVersion: "version",
						},
					},
				},
			},
			want: ChangesetV3{
				ChangesetV3Query: ChangesetV3Query{
					ToAppVersionReference: "name/version",
				},
			},
		},
		{
			name: "to app version chart id reference",
			args: args{
				model: models.Changeset{
					To: models.ChartReleaseVersion{
						AppVersion: &models.AppVersion{
							ChartID:    1,
							AppVersion: "version",
						},
					},
				},
			},
			want: ChangesetV3{
				ChangesetV3Query: ChangesetV3Query{
					ToAppVersionReference: "1/version",
				},
			},
		},
		{
			name: "to app version id reference",
			args: args{
				model: models.Changeset{
					To: models.ChartReleaseVersion{
						AppVersionID: utils.PointerTo[uint](1),
					},
				},
			},
			want: ChangesetV3{
				ChangesetV3Query: ChangesetV3Query{
					ToAppVersionReference: "1",
				},
			},
		},
		{
			name: "to chart version follow chart release name",
			args: args{
				model: models.Changeset{
					To: models.ChartReleaseVersion{
						ChartVersionFollowChartRelease: &models.ChartRelease{
							Name: "name",
						},
					},
				},
			},
			want: ChangesetV3{
				ChangesetV3Query: ChangesetV3Query{
					ChangesetV3Create: ChangesetV3Create{
						ToChartVersionFollowChartRelease: "name",
					},
				},
			},
		},
		{
			name: "to chart version follow chart release id",
			args: args{
				model: models.Changeset{
					To: models.ChartReleaseVersion{
						ChartVersionFollowChartReleaseID: utils.PointerTo[uint](1),
					},
				},
			},
			want: ChangesetV3{
				ChangesetV3Query: ChangesetV3Query{
					ChangesetV3Create: ChangesetV3Create{
						ToChartVersionFollowChartRelease: "1",
					},
				},
			},
		},
		{
			name: "to chart version chart reference",
			args: args{
				model: models.Changeset{
					To: models.ChartReleaseVersion{
						ChartVersion: &models.ChartVersion{
							Chart: &models.Chart{
								Name: "name",
							},
							ChartVersion: "version",
						},
					},
				},
			},
			want: ChangesetV3{
				ChangesetV3Query: ChangesetV3Query{
					ToChartVersionReference: "name/version",
				},
			},
		},
		{
			name: "to chart version chart id reference",
			args: args{
				model: models.Changeset{
					To: models.ChartReleaseVersion{
						ChartVersion: &models.ChartVersion{
							ChartID:      1,
							ChartVersion: "version",
						},
					},
				},
			},
			want: ChangesetV3{
				ChangesetV3Query: ChangesetV3Query{
					ToChartVersionReference: "1/version",
				},
			},
		},
		{
			name: "to chart version id reference",
			args: args{
				model: models.Changeset{
					To: models.ChartReleaseVersion{
						ChartVersionID: utils.PointerTo[uint](1),
					},
				},
			},
			want: ChangesetV3{
				ChangesetV3Query: ChangesetV3Query{
					ToChartVersionReference: "1",
				},
			},
		},
		{
			name: "planned by email",
			args: args{
				model: models.Changeset{
					PlannedBy: &models.User{
						Email: "email",
					},
				},
			},
			want: ChangesetV3{
				ChangesetV3Query: ChangesetV3Query{
					PlannedBy: utils.PointerTo("email"),
					PlannedByInfo: &UserV3{
						Email:                  "email",
						SuitabilityDescription: utils.PointerTo("user email lacks production suitability"),
						Suitable:               utils.PointerTo(false),
					},
				},
			},
		},
		{
			name: "planned by id",
			args: args{
				model: models.Changeset{
					PlannedByID: utils.PointerTo[uint](1),
				},
			},
			want: ChangesetV3{
				ChangesetV3Query: ChangesetV3Query{
					PlannedBy: utils.PointerTo("1"),
				},
			},
		},
		{
			name: "applied by email",
			args: args{
				model: models.Changeset{
					AppliedBy: &models.User{
						Email: "email",
					},
				},
			},
			want: ChangesetV3{
				ChangesetV3Query: ChangesetV3Query{
					AppliedBy: utils.PointerTo("email"),
					AppliedByInfo: &UserV3{
						Email:                  "email",
						SuitabilityDescription: utils.PointerTo("user email lacks production suitability"),
						Suitable:               utils.PointerTo(false),
					},
				},
			},
		},
		{
			name: "applied by id",
			args: args{
				model: models.Changeset{
					AppliedByID: utils.PointerTo[uint](1),
				},
			},
			want: ChangesetV3{
				ChangesetV3Query: ChangesetV3Query{
					AppliedBy: utils.PointerTo("1"),
				},
			},
		},
		{
			name: "chart release by name",
			args: args{
				model: models.Changeset{
					ChartRelease: &models.ChartRelease{
						Name: "name",
					},
				},
			},
			want: ChangesetV3{
				ChangesetV3Query: ChangesetV3Query{
					ChangesetV3Create: ChangesetV3Create{
						ChartRelease: "name",
					},
					ChartReleaseInfo: &ChartReleaseV3{
						ChartReleaseV3Create: ChartReleaseV3Create{
							Name: "name",
						},
					},
				},
			},
		},
		{
			name: "chart release by id",
			args: args{
				model: models.Changeset{
					ChartReleaseID: 1,
				},
			},
			want: ChangesetV3{
				ChangesetV3Query: ChangesetV3Query{
					ChangesetV3Create: ChangesetV3Create{
						ChartRelease: "1",
					},
				},
			},
		},
		{
			name: "normal fields",
			args: args{
				model: models.Changeset{
					Model: gorm.Model{
						ID: 1,
					},
					CiIdentifier: &models.CiIdentifier{
						Model: gorm.Model{ID: 2},
					},
					From: models.ChartReleaseVersion{
						ResolvedAt:           utils.PointerTo(now.Add(-time.Hour)),
						AppVersionResolver:   utils.PointerTo("resolver"),
						AppVersionExact:      utils.PointerTo("exact"),
						AppVersionBranch:     utils.PointerTo("branch"),
						AppVersionCommit:     utils.PointerTo("commit"),
						ChartVersionResolver: utils.PointerTo("resolver"),
						ChartVersionExact:    utils.PointerTo("exact"),
						HelmfileRef:          utils.PointerTo("helmfile-ref"),
						HelmfileRefEnabled:   utils.PointerTo(true),
					},
					To: models.ChartReleaseVersion{
						ResolvedAt:           utils.PointerTo(now.Add(time.Hour)),
						AppVersionResolver:   utils.PointerTo("resolver"),
						AppVersionExact:      utils.PointerTo("exact"),
						AppVersionBranch:     utils.PointerTo("branch"),
						AppVersionCommit:     utils.PointerTo("commit"),
						ChartVersionResolver: utils.PointerTo("resolver"),
						ChartVersionExact:    utils.PointerTo("exact"),
						HelmfileRef:          utils.PointerTo("helmfile-ref"),
						HelmfileRefEnabled:   utils.PointerTo(true),
					},
					AppliedAt:    utils.PointerTo(now.Add(time.Hour)),
					SupersededAt: utils.PointerTo(now.Add(time.Minute)),
				},
			},
			want: ChangesetV3{
				CommonFields: CommonFields{ID: 1},
				ChangesetV3Query: ChangesetV3Query{
					CiIdentifier:             &CiIdentifierV3{CommonFields: CommonFields{ID: 2}},
					AppliedAt:                utils.PointerTo(now.Add(time.Hour)),
					SupersededAt:             utils.PointerTo(now.Add(time.Minute)),
					FromResolvedAt:           utils.PointerTo(now.Add(-time.Hour)),
					FromAppVersionResolver:   utils.PointerTo("resolver"),
					FromAppVersionExact:      utils.PointerTo("exact"),
					FromAppVersionBranch:     utils.PointerTo("branch"),
					FromAppVersionCommit:     utils.PointerTo("commit"),
					FromChartVersionResolver: utils.PointerTo("resolver"),
					FromChartVersionExact:    utils.PointerTo("exact"),
					FromHelmfileRef:          utils.PointerTo("helmfile-ref"),
					FromHelmfileRefEnabled:   utils.PointerTo(true),
					ToResolvedAt:             utils.PointerTo(now.Add(time.Hour)),
					ChangesetV3Create: ChangesetV3Create{
						ToAppVersionResolver:   utils.PointerTo("resolver"),
						ToAppVersionExact:      utils.PointerTo("exact"),
						ToAppVersionBranch:     utils.PointerTo("branch"),
						ToAppVersionCommit:     utils.PointerTo("commit"),
						ToChartVersionResolver: utils.PointerTo("resolver"),
						ToChartVersionExact:    utils.PointerTo("exact"),
						ToHelmfileRef:          utils.PointerTo("helmfile-ref"),
						ToHelmfileRefEnabled:   utils.PointerTo(true),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, changesetFromModel(tt.args.model), "changesetFromModel(%v)", tt.args.model)
		})
	}
}
