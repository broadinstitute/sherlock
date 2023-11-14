package models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
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

func (s *modelSuite) TestEnvironmentCreationForbidden() {
	s.SetNonSuitableTestUserForDB()
	environment := Environment{
		Base:                       "live",
		Lifecycle:                  "static",
		Name:                       "prod",
		ValuesName:                 "prod",
		AutoPopulateChartReleases:  utils.PointerTo(false),
		DefaultNamespace:           "terra-prod",
		DefaultFirecloudDevelopRef: utils.PointerTo("prod"),
		RequiresSuitability:        utils.PointerTo(true),
		BaseDomain:                 utils.PointerTo("dsde-prod.broadinstitute.org"),
		NamePrefixesDomain:         utils.PointerTo(false),
		HelmfileRef:                utils.PointerTo("HEAD"),
		PreventDeletion:            utils.PointerTo(true),
		Description:                utils.PointerTo("Terra's production environment"),
		Offline:                    utils.PointerTo(false),
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

func (s *modelSuite) TestEnvironmentValidationSqlName() {
	//TODO add more
	s.SetNonSuitableTestUserForDB()
	environment := Environment{
		Base:                       "live",
		Lifecycle:                  "template",
		Name:                       "prod_env",
		ValuesName:                 "prod",
		AutoPopulateChartReleases:  utils.PointerTo(false),
		DefaultNamespace:           "terra-prod",
		DefaultFirecloudDevelopRef: utils.PointerTo("prod"),
		RequiresSuitability:        utils.PointerTo(false),
		BaseDomain:                 utils.PointerTo("dsde-prod.broadinstitute.org"),
		NamePrefixesDomain:         utils.PointerTo(false),
		HelmfileRef:                utils.PointerTo("HEAD"),
		PreventDeletion:            utils.PointerTo(true),
		Description:                utils.PointerTo("Terra's production environment"),
		Offline:                    utils.PointerTo(false),
	}
	s.ErrorContains(s.DB.Create(&environment).Error, "violates check constraint \"name_valid\"")
}

func (s *modelSuite) TestEnvironmentValidationSqlOwnerID() {
	s.SetSuitableTestUserForDB()
	environment := s.TestData.Environment_Prod()
	err := s.DB.Model(&environment).Updates(&Environment{OwnerID: nil}).Error
	println(err)
}

func (s *modelSuite) TestEnvironmentValidationSqlLifecycle() {
	s.SetSuitableTestUserForDB()
	environment := s.TestData.Environment_Staging()
	err := s.DB.Model(&environment).Updates(&Environment{RequiresSuitability: nil}).Error
	//	s.ErrorContains(err, "violates check constraint \"lifecycle_valid\"")
	err = s.DB.Model(&environment).Updates(&Environment{TemplateEnvironmentID: nil, Lifecycle: "dynamic"}).Error
	s.ErrorContains(err, "violates check constraint \"lifecycle_valid\"")
	templateID := uint(2)
	err = s.DB.Model(&environment).Updates(&Environment{TemplateEnvironmentID: &templateID, Lifecycle: "template"}).Error
	//s.ErrorContains(err, "violates check constraint \"lifecycle_valid\"")

}
