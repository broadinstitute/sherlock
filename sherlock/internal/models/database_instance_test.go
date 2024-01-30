package models

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"gorm.io/gorm"
)

func (s *modelSuite) TestDatabaseInstance_Forbidden() {
	databaseInstance := s.TestData.DatabaseInstance_LeonardoProd()
	s.SetNonSuitableTestUserForDB()
	s.ErrorContains(databaseInstance.errorIfForbidden(s.DB), errors.Forbidden)
}

func (s *modelSuite) TestDatabaseInstance_Allowed() {
	databaseInstance := s.TestData.DatabaseInstance_LeonardoDev()
	s.SetNonSuitableTestUserForDB()
	s.NoError(databaseInstance.errorIfForbidden(s.DB))
}

func (s *modelSuite) TestDatabaseInstance_PermissionsUsesChartRelease() {
	databaseInstance := s.TestData.DatabaseInstance_LeonardoDev()
	err := s.DB.Delete(&DatabaseInstance{Model: gorm.Model{ID: databaseInstance.ID}}).Error
	s.ErrorContains(err, "wasn't properly loaded")
}

func (s *modelSuite) TestDatabaseInstance_RequiresChartRelease() {
	databaseInstance := s.TestData.DatabaseInstance_LeonardoDev()
	err := s.DB.Model(&databaseInstance).Update("chart_release_id", nil).Error
	s.ErrorContains(err, "chart_release_id_present")
}

func (s *modelSuite) TestDatabaseInstance_GoogleRequiresProject() {
	databaseInstance := s.TestData.DatabaseInstance_LeonardoDev()
	err := s.DB.Model(&databaseInstance).Update("google_project", "").Error
	s.ErrorContains(err, "google_info_present")
}

func (s *modelSuite) TestDatabaseInstance_GoogleRequiresInstanceName() {
	databaseInstance := s.TestData.DatabaseInstance_LeonardoDev()
	err := s.DB.Model(&databaseInstance).Update("instance_name", "").Error
	s.ErrorContains(err, "google_info_present")
}

func (s *modelSuite) TestDatabaseInstance_AzureRequiresInstanceName() {
	databaseInstance := s.TestData.DatabaseInstance_LeonardoDev()
	err := s.DB.Model(&databaseInstance).Updates(map[string]any{
		"platform":       "azure",
		"instance_name":  "",
		"google_project": "",
	}).Error
	s.ErrorContains(err, "azure_info_present")
}

func (s *modelSuite) TestDatabaseInstance_BlocksDuplicates() {
	databaseInstance := s.TestData.DatabaseInstance_LeonardoDev()
	databaseInstance.ID = 0
	err := s.DB.Create(&databaseInstance).Error
	s.ErrorContains(err, "duplicate")
}

func (s *modelSuite) TestDatabaseInstance_RequiresValidPlatform() {
	databaseInstance := s.TestData.DatabaseInstance_LeonardoDev()
	err := s.DB.Model(&databaseInstance).Update("platform", "invalid").Error
	s.ErrorContains(err, "platform")
}
