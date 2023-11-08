package models

import "github.com/broadinstitute/sherlock/sherlock/internal/errors"

func (s *modelSuite) TestDatabaseInstanceForbidden() {
	databaseInstance := s.TestData.DatabaseInstance_LeonardoProd()
	s.SetNonSuitableTestUserForDB()
	s.ErrorContains(databaseInstance.errorIfForbidden(s.DB), errors.Forbidden)
}

func (s *modelSuite) TestDatabaseInstanceAllowed() {
	databaseInstance := s.TestData.DatabaseInstance_LeonardoDev()
	s.SetNonSuitableTestUserForDB()
	s.NoError(databaseInstance.errorIfForbidden(s.DB))
}
