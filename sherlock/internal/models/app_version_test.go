package models

import "github.com/broadinstitute/sherlock/go-shared/pkg/utils"

func (s *modelSuite) TestAppVersionChartIdValidationSqlMissing() {
	err := s.DB.Create(&AppVersion{AppVersion: "version"}).Error
	s.ErrorContains(err, "fk_v2_app_versions_chart")
}

func (s *modelSuite) TestAppVersionVersionValidationSqlMissing() {
	chart := Chart{Name: "name", ChartRepo: utils.PointerTo("repo")}
	err := s.DB.Create(&chart).Error
	s.NoError(err)
	err = s.DB.Create(&AppVersion{ChartID: chart.ID}).Error
	s.ErrorContains(err, "app_version")
}

func (s *modelSuite) TestAppVersionVersionValidationSqlEmpty() {
	chart := Chart{Name: "name", ChartRepo: utils.PointerTo("repo")}
	err := s.DB.Create(&chart).Error
	s.NoError(err)
	err = s.DB.Create(&AppVersion{ChartID: chart.ID, AppVersion: ""}).Error
	s.ErrorContains(err, "app_version")
}

func (s *modelSuite) TestAppVersionUniquenessSql() {
	chart := Chart{Name: "name", ChartRepo: utils.PointerTo("repo")}
	err := s.DB.Create(&chart).Error
	s.NoError(err)
	err = s.DB.Create(&AppVersion{ChartID: chart.ID, AppVersion: "version"}).Error
	s.NoError(err)
	err = s.DB.Create(&AppVersion{ChartID: chart.ID, AppVersion: "version"}).Error
	s.ErrorContains(err, "app_versions_selector_unique_constraint")
}

func (s *modelSuite) TestAppVersionValid() {
	chart := Chart{Name: "name", ChartRepo: utils.PointerTo("repo")}
	err := s.DB.Create(&chart).Error
	s.NoError(err)
	err = s.DB.Create(&AppVersion{ChartID: chart.ID, AppVersion: "version"}).Error
	s.NoError(err)
}

func (s *modelSuite) TestAppVersionCiIdentifiers() {
	chart := Chart{Name: "name", ChartRepo: utils.PointerTo("repo")}
	s.NoError(s.DB.Create(&chart).Error)
	appVersion := AppVersion{ChartID: chart.ID, AppVersion: "version"}
	s.NoError(s.DB.Create(&appVersion).Error)
	ciIdentifier := appVersion.GetCiIdentifier()
	s.NoError(s.DB.Create(&ciIdentifier).Error)
	s.NotZero(ciIdentifier.ID)
	s.Equal("app-version", ciIdentifier.ResourceType)
	s.Run("loads association", func() {
		var result AppVersion
		s.NoError(s.DB.Preload("CiIdentifier").First(&result, appVersion.ID).Error)
		s.NotNil(result.CiIdentifier)
		s.NotZero(result.CiIdentifier.ID)
		s.NotZero(result.GetCiIdentifier().ID)
		s.Equal(appVersion.ID, result.CiIdentifier.ResourceID)
		s.Equal("app-version", result.CiIdentifier.ResourceType)
	})
}
