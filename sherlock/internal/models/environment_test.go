package models

import (
	"database/sql"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func (s *modelSuite) TestEnvironmentUniqueResourcePrefixAssigning() {
	s.NotEmpty(s.TestData.Environment_Dev().UniqueResourcePrefix)
	s.NotEmpty(s.TestData.Environment_Staging().UniqueResourcePrefix)
	s.NotEqual(s.TestData.Environment_Dev().UniqueResourcePrefix, s.TestData.Environment_Staging().UniqueResourcePrefix)
	s.NotEmpty(s.TestData.Environment_Prod().UniqueResourcePrefix)
	s.NotEqual(s.TestData.Environment_Staging().UniqueResourcePrefix, s.TestData.Environment_Prod().UniqueResourcePrefix)
	s.NotEqual(s.TestData.Environment_Dev().UniqueResourcePrefix, s.TestData.Environment_Prod().UniqueResourcePrefix)
}

func (s *modelSuite) TestEnvironmentOwnerAssigning() {
	if s.NotNil(s.TestData.Environment_Prod().OwnerID) {
		s.Equal(s.TestData.User_Suitable().ID, *s.TestData.Environment_Prod().OwnerID)
	}
}

func (s *modelSuite) TestEnvironmentTemplateAutoPopulation() {
	environment := s.TestData.Environment_Swatomation()
	var chartReleases []ChartRelease
	err := s.DB.Where(&ChartRelease{EnvironmentID: &environment.ID}).Find(&chartReleases).Error
	s.NoError(err)
	if s.Len(chartReleases, 1) {
		s.Equal(s.TestData.Chart_Honeycomb().ID, chartReleases[0].ChartID)
	}
}

func (s *modelSuite) TestEnvironmentDynamicAutoPopulation() {
	// Being extra explicit about the state we're setting up
	s.TestData.Environment_Swatomation()
	s.TestData.ChartRelease_LeonardoSwatomation()
	s.TestData.DatabaseInstance_LeonardoSwatomation()

	bee := s.TestData.Environment_Swatomation_TestBee()

	s.Run("leonardo exists", func() {
		var chartRelease ChartRelease
		err := s.DB.Where(&ChartRelease{EnvironmentID: &bee.ID, ChartID: s.TestData.Chart_Leonardo().ID}).First(&chartRelease).Error
		s.NoError(err)
		s.Equal("leonardo-swatomation-test-bee", chartRelease.Name)
		s.Run("leonardo's database instance exists", func() {
			var databaseInstance DatabaseInstance
			err = s.DB.Where(&DatabaseInstance{ChartReleaseID: chartRelease.ID}).First(&databaseInstance).Error
			s.NoError(err)
		})
	})
	s.Run("honeycomb exists", func() {
		var chartRelease ChartRelease
		err := s.DB.Where(&ChartRelease{EnvironmentID: &bee.ID, ChartID: s.TestData.Chart_Honeycomb().ID}).First(&chartRelease).Error
		s.NoError(err)
		s.Equal("honeycomb-swatomation-test-bee", chartRelease.Name)
	})
}

func (s *modelSuite) TestEnvironmentPreventDeletion() {
	environment := s.TestData.Environment_Dev()
	if s.NotNil(environment.PreventDeletion) {
		s.True(*environment.PreventDeletion)
	}
	err := s.DB.Model(&environment).Delete(&environment).Error
	s.ErrorContains(err, "deletion protection enabled")
}

func (s *modelSuite) TestEnvironmentAllowDeletion() {
	environment := s.TestData.Environment_Dev()
	if s.NotNil(environment.PreventDeletion) {
		s.True(*environment.PreventDeletion)
	}
	err := s.DB.Model(&environment).Updates(&Environment{PreventDeletion: utils.PointerTo(false)}).Error
	s.NoError(err)
	err = s.DB.Model(&environment).Delete(&environment).Error
	s.NoError(err)
}

func (s *modelSuite) TestEnvironmentBlockDeletionOnStaticPropagation() {
	environment := s.TestData.Environment_Dev()
	s.TestData.ChartRelease_LeonardoDev()
	if s.NotNil(environment.PreventDeletion) {
		s.True(*environment.PreventDeletion)
	}
	err := s.DB.Model(&environment).Updates(&Environment{PreventDeletion: utils.PointerTo(false)}).Error
	s.NoError(err)
	err = s.DB.Model(&environment).Delete(&environment).Error
	s.ErrorContains(err, "chart releases are still inside this static environment")
}

func (s *modelSuite) TestEnvironmentAllowDeletionOnStaticPropagation() {
	environment := s.TestData.Environment_Dev()
	chartRelease := s.TestData.ChartRelease_LeonardoDev()
	if s.NotNil(environment.PreventDeletion) {
		s.True(*environment.PreventDeletion)
	}
	err := s.DB.Model(&environment).Updates(&Environment{PreventDeletion: utils.PointerTo(false)}).Error
	s.NoError(err)
	err = s.DB.Model(&chartRelease).Delete(&chartRelease).Error
	s.NoError(err)
	err = s.DB.Model(&environment).Delete(&environment).Error
	s.NoError(err)
}

func (s *modelSuite) TestEnvironmentBlockTemplateDeletionIfBeesExist() {
	environment := s.TestData.Environment_Swatomation()
	s.TestData.Environment_Swatomation_DevBee()
	if s.NotNil(environment.PreventDeletion) {
		s.True(*environment.PreventDeletion)
	}
	err := s.DB.Model(&environment).Updates(&Environment{PreventDeletion: utils.PointerTo(false)}).Error
	s.NoError(err)
	err = s.DB.Model(&environment).Delete(&environment).Error
	s.ErrorContains(err, "environments are still based on it")
}

func (s *modelSuite) TestEnvironmentPropagateTemplateDeletion() {
	environment := s.TestData.Environment_Swatomation()
	if s.NotNil(environment.PreventDeletion) {
		s.True(*environment.PreventDeletion)
	}
	var chartReleases []ChartRelease
	err := s.DB.Model(&ChartRelease{}).Where(&ChartRelease{EnvironmentID: &environment.ID}).Find(&chartReleases).Error
	s.NoError(err)
	s.NotEmpty(chartReleases)
	err = s.DB.Model(&environment).Updates(&Environment{PreventDeletion: utils.PointerTo(false)}).Error
	s.NoError(err)
	err = s.DB.Model(&environment).Delete(&environment).Error
	s.NoError(err)
	err = s.DB.Model(&ChartRelease{}).Where(&ChartRelease{EnvironmentID: &environment.ID}).Find(&chartReleases).Error
	s.NoError(err)
	s.Empty(chartReleases)
}

func (s *modelSuite) TestEnvironmentPropagateBeeDeletion() {
	environment := s.TestData.Environment_Swatomation_DevBee()
	var chartReleases []ChartRelease
	err := s.DB.Model(&ChartRelease{}).Where(&ChartRelease{EnvironmentID: &environment.ID}).Find(&chartReleases).Error
	s.NoError(err)
	s.NotEmpty(chartReleases)
	err = s.DB.Model(&environment).Delete(&environment).Error
	s.NoError(err)
	err = s.DB.Model(&ChartRelease{}).Where(&ChartRelease{EnvironmentID: &environment.ID}).Find(&chartReleases).Error
	s.NoError(err)
	s.Empty(chartReleases)
}

func (s *modelSuite) TestEnvironmentDeletionPropagateToDatabaseInstances() {
	s.TestData.DatabaseInstance_LeonardoSwatomation()
	environment := s.TestData.Environment_Swatomation_DevBee()
	var chartReleases []ChartRelease
	err := s.DB.Model(&ChartRelease{}).Where(&ChartRelease{EnvironmentID: &environment.ID}).Find(&chartReleases).Error
	s.NoError(err)
	s.NotEmpty(chartReleases)
	var databaseInstances []DatabaseInstance
	err = s.DB.Model(&DatabaseInstance{}).Where("chart_release_id IN ?", utils.Map(chartReleases, func(cr ChartRelease) uint { return cr.ID })).Find(&databaseInstances).Error
	s.NoError(err)
	s.NotEmpty(databaseInstances)
	err = s.DB.Model(&environment).Delete(&environment).Error
	s.NoError(err)
	err = s.DB.Model(&ChartRelease{}).Where(&ChartRelease{EnvironmentID: &environment.ID}).Find(&chartReleases).Error
	s.NoError(err)
	s.Empty(chartReleases)
	err = s.DB.Model(&DatabaseInstance{}).Where("id IN ?", utils.Map(databaseInstances, func(di DatabaseInstance) uint { return di.ID })).Find(&databaseInstances).Error
	s.NoError(err)
	s.Empty(databaseInstances)
}

func (s *modelSuite) TestEnvironment_errorIfForbidden_nullRequiresSuitability() {
	s.SetNonSuitableTestUserForDB()
	environment := Environment{
		Base:                      "live",
		Lifecycle:                 "static",
		Name:                      "prod",
		ValuesName:                "prod",
		AutoPopulateChartReleases: utils.PointerTo(false),
		DefaultNamespace:          "terra-prod",
		BaseDomain:                utils.PointerTo("dsde-prod.broadinstitute.org"),
		NamePrefixesDomain:        utils.PointerTo(false),
		HelmfileRef:               utils.PointerTo("HEAD"),
		PreventDeletion:           utils.PointerTo(true),
		Description:               utils.PointerTo("Terra's production environment"),
		Offline:                   utils.PointerTo(false),
	}
	s.NoError(environment.errorIfForbidden(s.DB))
}

func (s *modelSuite) TestEnvironmentCreationForbidden() {
	s.SetNonSuitableTestUserForDB()
	environment := Environment{
		Base:                      "live",
		Lifecycle:                 "static",
		Name:                      "prod",
		ValuesName:                "prod",
		AutoPopulateChartReleases: utils.PointerTo(false),
		DefaultNamespace:          "terra-prod",
		RequiresSuitability:       utils.PointerTo(true),
		BaseDomain:                utils.PointerTo("dsde-prod.broadinstitute.org"),
		NamePrefixesDomain:        utils.PointerTo(false),
		HelmfileRef:               utils.PointerTo("HEAD"),
		PreventDeletion:           utils.PointerTo(true),
		Description:               utils.PointerTo("Terra's production environment"),
		Offline:                   utils.PointerTo(false),
	}
	s.ErrorContains(s.DB.Create(&environment).Error, errors.Forbidden)
}

func (s *modelSuite) TestEnvironmentEditEscalateForbidden() {
	environment := s.TestData.Environment_Dev()
	s.SetNonSuitableTestUserForDB()
	s.ErrorContains(s.DB.
		Model(&environment).
		Updates(&Environment{RequiresSuitability: utils.PointerTo(true)}).
		Error, errors.Forbidden)
}

func (s *modelSuite) TestEnvironmentEditDeescalateForbidden() {
	environment := s.TestData.Environment_Prod()
	s.SetNonSuitableTestUserForDB()
	s.ErrorContains(s.DB.
		Model(&environment).
		Updates(&Environment{RequiresSuitability: utils.PointerTo(false)}).
		Error, errors.Forbidden)
}

func (s *modelSuite) TestEnvironmentDeleteForbidden() {
	environment := s.TestData.Environment_Prod()
	if s.NotNil(environment.PreventDeletion) {
		s.True(*environment.PreventDeletion)
	}
	err := s.DB.Model(&environment).Updates(&Environment{PreventDeletion: utils.PointerTo(false)}).Error
	s.NoError(err)
	s.SetNonSuitableTestUserForDB()
	s.ErrorContains(s.DB.
		Delete(&environment).
		Error, errors.Forbidden)
}

func (s *modelSuite) TestEnvironmentValidationSqlNameInvalid() {
	s.SetNonSuitableTestUserForDB()
	environment := s.TestData.Environment_Swatomation_TestBee()
	err := s.DB.Model(&environment).Select("Name").Updates(&Environment{Name: "prod_env"}).Error
	s.ErrorContains(err, "violates check constraint \"name_valid\"")
}

func (s *modelSuite) TestEnvironmentValidationSqlNameEmpty() {
	s.SetNonSuitableTestUserForDB()
	environment := s.TestData.Environment_Swatomation_TestBee()
	err := s.DB.Model(&environment).Select("Name").Updates(&Environment{Name: ""}).Error
	s.ErrorContains(err, "violates check constraint \"name_valid\"")
}

func (s *modelSuite) TestEnvironmentValidationSqlOwnerIDEmpty() {
	s.SetSuitableTestUserForDB()
	environment := s.TestData.Environment_Prod()
	err := s.DB.Model(&environment).Select("OwnerID", "LegacyOwner").Updates(&Environment{OwnerID: nil, LegacyOwner: utils.PointerTo("")}).Error
	s.ErrorContains(err, "violates check constraint \"owner_id_present\"")
}

func (s *modelSuite) TestEnvironmentValidationSqlOwnerIDNull() {
	s.SetSuitableTestUserForDB()
	environment := s.TestData.Environment_Prod()
	err := s.DB.Model(&environment).Select("OwnerID", "LegacyOwner").Updates(&Environment{OwnerID: nil, LegacyOwner: nil}).Error
	s.ErrorContains(err, "violates check constraint \"owner_id_present\"")
}

func (s *modelSuite) TestEnvironmentValidationSqlTemplateLifecycle() {
	s.SetSuitableTestUserForDB()
	templateEnv := s.TestData.Environment_Swatomation()
	err := s.DB.Model(&templateEnv).Updates(&Environment{TemplateEnvironmentID: utils.PointerTo(uint(1))}).Error
	s.ErrorContains(err, "violates check constraint \"lifecycle_valid\"")
}

func (s *modelSuite) TestEnvironmentValidationSqlDynamicLifecycle() {
	s.SetSuitableTestUserForDB()
	dynamicEnv := s.TestData.Environment_Swatomation_TestBee()
	err := s.DB.Model(&dynamicEnv).Select("TemplateEnvironmentID").Updates(&Environment{TemplateEnvironmentID: nil}).Error
	s.ErrorContains(err, "violates check constraint \"lifecycle_valid\"")
}

func (s *modelSuite) TestEnvironmentValidationSqlDynamicLifecycleBase() {
	s.SetSuitableTestUserForDB()
	staticEnv := s.TestData.Environment_Swatomation_LongBee()
	err := s.DB.Model(&staticEnv).Select("Base").Updates(&Environment{Base: ""}).Error
	s.ErrorContains(err, "violates check constraint \"lifecycle_valid\"")
}

func (s *modelSuite) TestEnvironmentValidationSqlDynamicLifecycleDefaultClusterID() {
	s.SetSuitableTestUserForDB()
	staticEnv := s.TestData.Environment_Swatomation_DevBee()
	err := s.DB.Model(&staticEnv).Select("DefaultClusterID").Updates(&Environment{DefaultClusterID: nil}).Error
	s.ErrorContains(err, "violates check constraint \"lifecycle_valid\"")
}

func (s *modelSuite) TestEnvironmentValidationSqlStaticLifecycleBase() {
	s.SetSuitableTestUserForDB()
	staticEnv := s.TestData.Environment_Prod()
	err := s.DB.Model(&staticEnv).Select("Base").Updates(&Environment{Base: ""}).Error
	s.ErrorContains(err, "violates check constraint \"lifecycle_valid\"")
}

func (s *modelSuite) TestEnvironmentValidationSqlStaticLifecycleDefaultClusterID() {
	s.SetSuitableTestUserForDB()
	staticEnv := s.TestData.Environment_Prod()
	err := s.DB.Model(&staticEnv).Select("DefaultClusterID").Updates(&Environment{DefaultClusterID: nil}).Error
	s.ErrorContains(err, "violates check constraint \"lifecycle_valid\"")
}

func (s *modelSuite) TestEnvironmentValidationSqlDefaultNamespace() {
	s.SetSuitableTestUserForDB()
	environment := s.TestData.Environment_DdpAzureDev()
	err := s.DB.Model(&environment).Select("DefaultNamespace").Updates(&Environment{DefaultNamespace: ""}).Error
	s.ErrorContains(err, "violates check constraint \"default_namespace_present\"")
}

func (s *modelSuite) TestEnvironmentValidationSqlUniqueResourcePrefix() {
	s.SetSuitableTestUserForDB()
	environment := s.TestData.Environment_DdpAzureProd()
	err := s.DB.Model(&environment).Select("UniqueResourcePrefix").Updates(&Environment{UniqueResourcePrefix: ""}).Error
	s.ErrorContains(err, "violates check constraint \"unique_resource_prefix_present\"")
}

func (s *modelSuite) TestEnvironmentValidationSqlDeleteAfterNotDynamic() {
	s.SetSuitableTestUserForDB()
	environment := s.TestData.Environment_Staging()
	err := s.DB.Model(&environment).Select("DeleteAfter").Updates(&Environment{DeleteAfter: sql.NullTime{Time: time.Now().Add(6 * time.Hour), Valid: true}}).Error
	s.ErrorContains(err, "violates check constraint \"delete_after_valid\"")
}

func (s *modelSuite) TestEnvironmentValidationSqlDeleteAfterPreventDeletion() {
	s.SetSuitableTestUserForDB()
	environment := s.TestData.Environment_Swatomation_TestBee()
	err := s.DB.Model(&environment).Select("PreventDeletion").Updates(&Environment{PreventDeletion: utils.PointerTo(true)}).Error
	s.ErrorContains(err, "violates check constraint \"delete_after_valid\"")
}

func (s *modelSuite) TestEnvironmentValidationSqlOfflineNotDynamic() {
	s.SetSuitableTestUserForDB()
	environment := s.TestData.Environment_Prod()
	err := s.DB.Model(&environment).Select("Offline").Updates(&Environment{Offline: utils.PointerTo(true)}).Error
	s.ErrorContains(err, "violates check constraint \"offline_valid\"")
}

func (s *modelSuite) TestEnvironmentValidationSqlOfflineBegin() {
	s.SetSuitableTestUserForDB()
	environment := s.TestData.Environment_Dev()
	err := s.DB.Model(&environment).Select("OfflineScheduleBeginEnabled", "OfflineScheduleBeginTime").Updates(&Environment{OfflineScheduleBeginEnabled: utils.PointerTo(true), OfflineScheduleBeginTime: utils.PointerTo("begin time")}).Error
	s.ErrorContains(err, "violates check constraint \"offline_valid\"")
}

func (s *modelSuite) TestEnvironmentValidationSqlOfflineEnd() {
	s.SetSuitableTestUserForDB()
	environment := s.TestData.Environment_Staging()
	err := s.DB.Model(&environment).Select("OfflineScheduleEndEnabled", "OfflineScheduleEndTime").Updates(&Environment{OfflineScheduleEndEnabled: utils.PointerTo(true), OfflineScheduleEndTime: utils.PointerTo("end time")}).Error
	s.ErrorContains(err, "violates check constraint \"offline_valid\"")
}

func (s *modelSuite) TestEnvironmentValidationSqlOfflineBeginPresent() {
	s.SetSuitableTestUserForDB()
	environment := s.TestData.Environment_Swatomation_DevBee()
	err := s.DB.Model(&environment).Select("OfflineScheduleBeginTime").Updates(&Environment{OfflineScheduleBeginTime: nil}).Error
	s.ErrorContains(err, "violates check constraint \"offline_schedule_begin_time_present\"")
}

func (s *modelSuite) TestEnvironmentValidationSqlOfflineEndPresent() {
	s.SetSuitableTestUserForDB()
	environment := s.TestData.Environment_Swatomation_DevBee()
	err := s.DB.Model(&environment).Select("OfflineScheduleEndTime").Updates(&Environment{OfflineScheduleEndTime: nil}).Error
	s.ErrorContains(err, "violates check constraint \"offline_schedule_end_time_present\"")
}

func (s *modelSuite) TestEnvironmentValidationSqlNameUnique() {
	s.SetSuitableTestUserForDB()
	a := s.TestData.Environment_Dev()
	b := s.TestData.Environment_Staging()
	err := s.DB.Model(&b).Updates(&Environment{Name: a.Name}).Error
	s.ErrorContains(err, "violates unique constraint \"environments_name_unique\"")
}

func (s *modelSuite) TestEnvironmentValidationSqlUrpUnique() {
	s.SetSuitableTestUserForDB()
	a := s.TestData.Environment_Dev()
	b := s.TestData.Environment_Staging()
	err := s.DB.Model(&b).Updates(&Environment{UniqueResourcePrefix: a.UniqueResourcePrefix}).Error
	s.ErrorContains(err, "violates unique constraint \"environments_urp_unique\"")
}

func TestEnvironment_SlackBeehiveLink(t *testing.T) {
	type fields struct {
		Name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "no name",
			fields: fields{},
			want:   "(unknown environment)",
		},
		{
			name:   "with name",
			fields: fields{Name: "example"},
			want:   "<" + fmt.Sprintf(config.Config.String("beehive.environmentUrlFormatString"), "example") + "|example>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Environment{
				Name: tt.fields.Name,
			}
			assert.Equalf(t, tt.want, e.SlackBeehiveLink(), "SlackBeehiveLink()")
		})
	}
}

func TestEnvironment_ArgoCdUrl(t *testing.T) {
	type fields struct {
		Name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		wantOk bool
	}{
		{
			name:   "no name",
			fields: fields{},
			want:   "",
			wantOk: false,
		},
		{
			name:   "with name",
			fields: fields{Name: "example"},
			want:   fmt.Sprintf(config.Config.String("argoCd.environmentUrlFormatString"), "example"),
			wantOk: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Environment{
				Name: tt.fields.Name,
			}
			got, gotOk := e.ArgoCdUrl()
			assert.Equalf(t, tt.want, got, "ArgoCdUrl()")
			assert.Equalf(t, tt.wantOk, gotOk, "ArgoCdUrl()")
		})
	}
}
