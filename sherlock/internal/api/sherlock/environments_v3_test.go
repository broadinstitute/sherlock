package sherlock

import (
	"database/sql"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func (s *handlerSuite) TestEnvironmentV3_toModel() {
	nowString := time.Now().Truncate(time.Second).Format(time.RFC3339)
	now, err := time.Parse(time.RFC3339, nowString)
	s.NoError(err)
	templateEnvironment := s.TestData.Environment_Swatomation()
	defaultCluster := s.TestData.Cluster_TerraQaBees()
	pagerdutyIntegration := s.TestData.PagerdutyIntegration_ManuallyTriggeredTerraIncident()
	owner := s.TestData.User_Suitable()
	pactUuid := uuid.New()
	type fields struct {
		CommonFields        CommonFields
		EnvironmentV3Create EnvironmentV3Create
	}
	tests := []struct {
		name    string
		fields  fields
		want    models.Environment
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "empty",
			fields:  fields{},
			wantErr: assert.NoError,
			want:    models.Environment{},
		},
		{
			name: "delete after",
			fields: fields{
				EnvironmentV3Create: EnvironmentV3Create{
					EnvironmentV3Edit: EnvironmentV3Edit{
						DeleteAfter: utils.PointerTo(now.Format(time.RFC3339)),
					},
				},
			},
			wantErr: assert.NoError,
			want: models.Environment{
				DeleteAfter: sql.NullTime{
					Time:  now,
					Valid: true,
				},
			},
		},
		{
			name: "invalid template env selector",
			fields: fields{
				EnvironmentV3Create: EnvironmentV3Create{
					TemplateEnvironment: "!!!!!",
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "not found template env",
			fields: fields{
				EnvironmentV3Create: EnvironmentV3Create{
					TemplateEnvironment: "not-found",
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "invalid cluster selector",
			fields: fields{
				EnvironmentV3Create: EnvironmentV3Create{
					EnvironmentV3Edit: EnvironmentV3Edit{
						DefaultCluster: utils.PointerTo("!!!!!"),
					},
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "not found cluster",
			fields: fields{
				EnvironmentV3Create: EnvironmentV3Create{
					EnvironmentV3Edit: EnvironmentV3Edit{
						DefaultCluster: utils.PointerTo("not-found"),
					},
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "invalid pagerduty integration selector",
			fields: fields{
				EnvironmentV3Create: EnvironmentV3Create{
					EnvironmentV3Edit: EnvironmentV3Edit{
						PagerdutyIntegration: utils.PointerTo("!!!!!"),
					},
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "not found pagerduty integration",
			fields: fields{
				EnvironmentV3Create: EnvironmentV3Create{
					EnvironmentV3Edit: EnvironmentV3Edit{
						PagerdutyIntegration: utils.PointerTo("pd-id/not-found"),
					},
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "invalid owner selector",
			fields: fields{
				EnvironmentV3Create: EnvironmentV3Create{
					EnvironmentV3Edit: EnvironmentV3Edit{
						Owner: utils.PointerTo("!!!!!"),
					},
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "not found owner",
			fields: fields{
				EnvironmentV3Create: EnvironmentV3Create{
					EnvironmentV3Edit: EnvironmentV3Edit{
						Owner: utils.PointerTo("not-found@example.com"),
					},
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "valid",
			fields: fields{
				CommonFields: CommonFields{
					ID:        1,
					CreatedAt: now.Add(-time.Hour),
					UpdatedAt: now.Add(-time.Minute),
				},
				EnvironmentV3Create: EnvironmentV3Create{
					Base:                      "base",
					AutoPopulateChartReleases: utils.PointerTo(true),
					Lifecycle:                 "lifecycle",
					Name:                      "name",
					TemplateEnvironment:       templateEnvironment.Name,
					UniqueResourcePrefix:      "unique-resource-prefix",
					DefaultNamespace:          "default-namespace",
					ValuesName:                "values-name",
					EnvironmentV3Edit: EnvironmentV3Edit{
						DefaultCluster:              utils.PointerTo(defaultCluster.Name),
						Owner:                       utils.PointerTo(owner.Email),
						RequiresSuitability:         utils.PointerTo(true),
						RequiredRole:                s.TestData.Role_TerraSuitableEngineer().Name,
						BaseDomain:                  utils.PointerTo("base-domain"),
						NamePrefixesDomain:          utils.PointerTo(true),
						HelmfileRef:                 utils.PointerTo("HEAD"),
						PreventDeletion:             utils.PointerTo(true),
						DeleteAfter:                 utils.PointerTo(nowString),
						Description:                 utils.PointerTo("description"),
						PactIdentifier:              utils.PointerTo(pactUuid),
						PagerdutyIntegration:        utils.PointerTo(utils.UintToString(pagerdutyIntegration.ID)),
						Offline:                     utils.PointerTo(true),
						OfflineScheduleBeginEnabled: utils.PointerTo(true),
						OfflineScheduleBeginTime:    utils.PointerTo(now),
						OfflineScheduleEndEnabled:   utils.PointerTo(true),
						OfflineScheduleEndTime:      utils.PointerTo(now),
						OfflineScheduleEndWeekends:  utils.PointerTo(true),
						EnableJanitor:               utils.PointerTo(true),
						ServiceBannerBucket: 		utils.PointerTo("firecloud-alerts-qa-bees"),
					},
				},
			},
			wantErr: assert.NoError,
			want: models.Environment{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: now.Add(-time.Hour),
					UpdatedAt: now.Add(-time.Minute),
				},
				Base:                      "base",
				Lifecycle:                 "lifecycle",
				Name:                      "name",
				TemplateEnvironmentID:     &templateEnvironment.ID,
				ValuesName:                "values-name",
				AutoPopulateChartReleases: utils.PointerTo(true),
				UniqueResourcePrefix:      "unique-resource-prefix",
				DefaultNamespace:          "default-namespace",
				DefaultClusterID:          &defaultCluster.ID,
				OwnerID:                   &owner.ID,
				RequiresSuitability:       utils.PointerTo(true),
				RequiredRoleID:            utils.PointerTo(s.TestData.Role_TerraSuitableEngineer().ID),
				BaseDomain:                utils.PointerTo("base-domain"),
				NamePrefixesDomain:        utils.PointerTo(true),
				HelmfileRef:               utils.PointerTo("HEAD"),
				PreventDeletion:           utils.PointerTo(true),
				DeleteAfter: sql.NullTime{
					Time:  now,
					Valid: true,
				},
				Description:                 utils.PointerTo("description"),
				PagerdutyIntegrationID:      &pagerdutyIntegration.ID,
				Offline:                     utils.PointerTo(true),
				OfflineScheduleBeginEnabled: utils.PointerTo(true),
				OfflineScheduleBeginTime:    utils.TimePtrToISO8601(&now),
				OfflineScheduleEndEnabled:   utils.PointerTo(true),
				OfflineScheduleEndTime:      utils.TimePtrToISO8601(&now),
				OfflineScheduleEndWeekends:  utils.PointerTo(true),
				PactIdentifier:              &pactUuid,
				EnableJanitor:               utils.PointerTo(true),
				ServiceBannerBucket:         utils.PointerTo("firecloud-alerts-qa-bees"),
			},
		},
		{
			name: "full with whitespace",
			fields: fields{
				CommonFields: CommonFields{
					ID:        1,
					CreatedAt: now.Add(-time.Hour),
					UpdatedAt: now.Add(-time.Minute),
				},
				EnvironmentV3Create: EnvironmentV3Create{
					Base:                      "base",
					AutoPopulateChartReleases: utils.PointerTo(true),
					Lifecycle:                 "lifecycle",
					Name:                      "name",
					TemplateEnvironment:       templateEnvironment.Name,
					UniqueResourcePrefix:      "unique-resource-prefix",
					DefaultNamespace:          "default-namespace",
					ValuesName:                "values-name",
					EnvironmentV3Edit: EnvironmentV3Edit{
						DefaultCluster:              utils.PointerTo(defaultCluster.Name),
						Owner:                       utils.PointerTo(owner.Email),
						RequiresSuitability:         utils.PointerTo(true),
						BaseDomain:                  utils.PointerTo("base-domain"),
						NamePrefixesDomain:          utils.PointerTo(true),
						HelmfileRef:                 utils.PointerTo(" HEAD "),
						PreventDeletion:             utils.PointerTo(true),
						DeleteAfter:                 utils.PointerTo(nowString),
						Description:                 utils.PointerTo("description"),
						PactIdentifier:              utils.PointerTo(pactUuid),
						PagerdutyIntegration:        utils.PointerTo(utils.UintToString(pagerdutyIntegration.ID)),
						Offline:                     utils.PointerTo(true),
						OfflineScheduleBeginEnabled: utils.PointerTo(true),
						OfflineScheduleBeginTime:    utils.PointerTo(now),
						OfflineScheduleEndEnabled:   utils.PointerTo(true),
						OfflineScheduleEndTime:      utils.PointerTo(now),
						OfflineScheduleEndWeekends:  utils.PointerTo(true),
						EnableJanitor:               utils.PointerTo(true),
						ServiceBannerBucket:         utils.PointerTo("firecloud-alerts-qa-bees"),
					},
				},
			},
			wantErr: assert.NoError,
			want: models.Environment{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: now.Add(-time.Hour),
					UpdatedAt: now.Add(-time.Minute),
				},
				Base:                      "base",
				Lifecycle:                 "lifecycle",
				Name:                      "name",
				TemplateEnvironmentID:     &templateEnvironment.ID,
				ValuesName:                "values-name",
				AutoPopulateChartReleases: utils.PointerTo(true),
				UniqueResourcePrefix:      "unique-resource-prefix",
				DefaultNamespace:          "default-namespace",
				DefaultClusterID:          &defaultCluster.ID,
				OwnerID:                   &owner.ID,
				RequiresSuitability:       utils.PointerTo(true),
				BaseDomain:                utils.PointerTo("base-domain"),
				NamePrefixesDomain:        utils.PointerTo(true),
				HelmfileRef:               utils.PointerTo("HEAD"),
				PreventDeletion:           utils.PointerTo(true),
				DeleteAfter: sql.NullTime{
					Time:  now,
					Valid: true,
				},
				Description:                 utils.PointerTo("description"),
				PagerdutyIntegrationID:      &pagerdutyIntegration.ID,
				Offline:                     utils.PointerTo(true),
				OfflineScheduleBeginEnabled: utils.PointerTo(true),
				OfflineScheduleBeginTime:    utils.TimePtrToISO8601(&now),
				OfflineScheduleEndEnabled:   utils.PointerTo(true),
				OfflineScheduleEndTime:      utils.TimePtrToISO8601(&now),
				OfflineScheduleEndWeekends:  utils.PointerTo(true),
				PactIdentifier:              &pactUuid,
				EnableJanitor:               utils.PointerTo(true),
				ServiceBannerBucket:         utils.PointerTo("firecloud-alerts-qa-bees"),
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			e := EnvironmentV3{
				CommonFields:        tt.fields.CommonFields,
				EnvironmentV3Create: tt.fields.EnvironmentV3Create,
			}
			got, err := e.toModel(s.DB)
			if !tt.wantErr(s.T(), err, "toModel()") {
				return
			}
			s.Equalf(tt.want, got, "toModel()")
		})
	}
}

func Test_environmentFromModel(t *testing.T) {
	config.LoadTestConfig()
	now := time.Now()
	nowString := utils.TimePtrToISO8601(&now)
	nowTimeParsedAgain, err := utils.ISO8601PtrToTime(nowString)
	assert.NoError(t, err)
	pactUuid := uuid.New()
	type args struct {
		model models.Environment
	}
	tests := []struct {
		name string
		args args
		want EnvironmentV3
	}{
		{
			name: "empty",
			args: args{},
			want: EnvironmentV3{
				EnvironmentV3Create: EnvironmentV3Create{
					EnvironmentV3Edit: EnvironmentV3Edit{
						RequiredRole: utils.PointerTo(config.Config.String("model.roles.substituteEmptyRequiredRoleWithValue")),
					},
				},
			},
		},
		{
			name: "full",
			args: args{model: models.Environment{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: now.Add(-time.Hour),
					UpdatedAt: now.Add(-time.Minute),
				},
				CiIdentifier:              &models.CiIdentifier{Model: gorm.Model{ID: 2}},
				Base:                      "base",
				Lifecycle:                 "lifecycle",
				Name:                      "name",
				TemplateEnvironment:       &models.Environment{Model: gorm.Model{ID: 3}, Name: "name-3"},
				TemplateEnvironmentID:     utils.PointerTo[uint](3),
				ValuesName:                "values-name",
				AutoPopulateChartReleases: utils.PointerTo(true),
				UniqueResourcePrefix:      "unique-resource-prefix",
				DefaultNamespace:          "default-namespace",
				DefaultCluster:            &models.Cluster{Model: gorm.Model{ID: 4}, Name: "name-4"},
				DefaultClusterID:          utils.PointerTo[uint](4),
				Owner:                     &models.User{Model: gorm.Model{ID: 5}, Email: "example@example.com"},
				OwnerID:                   utils.PointerTo[uint](5),
				RequiresSuitability:       utils.PointerTo(true),
				RequiredRole:              &models.Role{Model: gorm.Model{ID: 999}, RoleFields: models.RoleFields{Name: utils.PointerTo("role-name")}},
				RequiredRoleID:            utils.PointerTo[uint](999),
				BaseDomain:                utils.PointerTo("base-domain"),
				NamePrefixesDomain:        utils.PointerTo(true),
				HelmfileRef:               utils.PointerTo("HEAD"),
				PreventDeletion:           utils.PointerTo(true),
				DeleteAfter: sql.NullTime{
					Time:  now,
					Valid: true,
				},
				Description:                 utils.PointerTo("description"),
				PagerdutyIntegration:        &models.PagerdutyIntegration{Model: gorm.Model{ID: 6}, PagerdutyID: "blah"},
				PagerdutyIntegrationID:      utils.PointerTo[uint](6),
				Offline:                     utils.PointerTo(true),
				OfflineScheduleBeginEnabled: utils.PointerTo(true),
				OfflineScheduleBeginTime:    utils.TimePtrToISO8601(&now),
				OfflineScheduleEndEnabled:   utils.PointerTo(true),
				OfflineScheduleEndTime:      utils.TimePtrToISO8601(&now),
				OfflineScheduleEndWeekends:  utils.PointerTo(true),
				PactIdentifier:              &pactUuid,
				EnableJanitor:               utils.PointerTo(true),
				ServiceBannerBucket:         utils.PointerTo("firecloud-alerts-qa-bees"),
			}},
			want: EnvironmentV3{
				CommonFields: CommonFields{
					ID:        1,
					CreatedAt: now.Add(-time.Hour),
					UpdatedAt: now.Add(-time.Minute),
				},
				CiIdentifier:             &CiIdentifierV3{CommonFields: CommonFields{ID: 2}},
				TemplateEnvironmentInfo:  &EnvironmentV3{CommonFields: CommonFields{ID: 3}, EnvironmentV3Create: EnvironmentV3Create{Name: "name-3", EnvironmentV3Edit: EnvironmentV3Edit{RequiredRole: utils.PointerTo(config.Config.String("model.roles.substituteEmptyRequiredRoleWithValue"))}}},
				DefaultClusterInfo:       &ClusterV3{CommonFields: CommonFields{ID: 4}, ClusterV3Create: ClusterV3Create{Name: "name-4", ClusterV3Edit: ClusterV3Edit{RequiredRole: utils.PointerTo(config.Config.String("model.roles.substituteEmptyRequiredRoleWithValue"))}}},
				PagerdutyIntegrationInfo: &PagerdutyIntegrationV3{CommonFields: CommonFields{ID: 6}, PagerdutyID: "blah"},
				OwnerInfo: &UserV3{CommonFields: CommonFields{ID: 5}, Email: "example@example.com", Suitable: utils.PointerTo(false),
					SuitabilityDescription: utils.PointerTo("no matching suitability record found or loaded; assuming unsuitable")},
				RequiredRoleInfo: &RoleV3{CommonFields: CommonFields{ID: 999}, RoleV3Edit: RoleV3Edit{Name: utils.PointerTo("role-name")}},
				EnvironmentV3Create: EnvironmentV3Create{
					Base:                      "base",
					AutoPopulateChartReleases: utils.PointerTo(true),
					Lifecycle:                 "lifecycle",
					Name:                      "name",
					TemplateEnvironment:       "name-3",
					UniqueResourcePrefix:      "unique-resource-prefix",
					DefaultNamespace:          "default-namespace",
					ValuesName:                "values-name",
					EnvironmentV3Edit: EnvironmentV3Edit{
						DefaultCluster:              utils.PointerTo("name-4"),
						Owner:                       utils.PointerTo("example@example.com"),
						RequiresSuitability:         utils.PointerTo(true),
						RequiredRole:                utils.PointerTo("role-name"),
						BaseDomain:                  utils.PointerTo("base-domain"),
						NamePrefixesDomain:          utils.PointerTo(true),
						HelmfileRef:                 utils.PointerTo("HEAD"),
						PreventDeletion:             utils.PointerTo(true),
						DeleteAfter:                 utils.PointerTo(now.Format(time.RFC3339)),
						Description:                 utils.PointerTo("description"),
						PactIdentifier:              utils.PointerTo(pactUuid),
						PagerdutyIntegration:        utils.PointerTo("pd-id/blah"),
						Offline:                     utils.PointerTo(true),
						OfflineScheduleBeginEnabled: utils.PointerTo(true),
						OfflineScheduleBeginTime:    nowTimeParsedAgain,
						OfflineScheduleEndEnabled:   utils.PointerTo(true),
						OfflineScheduleEndTime:      nowTimeParsedAgain,
						OfflineScheduleEndWeekends:  utils.PointerTo(true),
						EnableJanitor:               utils.PointerTo(true),
						ServiceBannerBucket:         utils.PointerTo("firecloud-alerts-qa-bees"),
					},
				},
			},
		},
		{
			name: "pagerduty ID case",
			args: args{model: models.Environment{
				PagerdutyIntegrationID: utils.PointerTo[uint](6),
			}},
			want: EnvironmentV3{
				EnvironmentV3Create: EnvironmentV3Create{
					EnvironmentV3Edit: EnvironmentV3Edit{
						PagerdutyIntegration: utils.PointerTo("6"),
						RequiredRole:         utils.PointerTo(config.Config.String("model.roles.substituteEmptyRequiredRoleWithValue")),
					},
				},
			},
		},
		{
			name: "role ID case",
			args: args{model: models.Environment{
				RequiredRoleID: utils.PointerTo[uint](999),
			}},
			want: EnvironmentV3{
				EnvironmentV3Create: EnvironmentV3Create{
					EnvironmentV3Edit: EnvironmentV3Edit{
						RequiredRole: utils.PointerTo("999"),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, environmentFromModel(tt.args.model), "environmentFromModel(%v)", tt.args.model)
		})
	}
}
