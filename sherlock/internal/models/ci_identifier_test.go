package models

func (s *modelSuite) TestCiIdentifierResourceValidationSqlInvalid() {
	err := s.DB.Create(&CiIdentifier{}).Error
	s.ErrorContains(err, "violates check constraint \"resource_present\"")
}

func (s *modelSuite) TestCiIdentifierResourceValidationSqlValid() {
	err := s.DB.Create(&CiIdentifier{
		ResourceType: "environment",
		ResourceID:   1,
	}).Error
	s.NoError(err)
}

func (s *modelSuite) TestCiIdentifierUniquenessSqlInvalid() {
	err := s.DB.Create(&CiIdentifier{
		ResourceType: "environment",
		ResourceID:   1,
	}).Error
	s.NoError(err)
	err = s.DB.Create(&CiIdentifier{
		ResourceType: "environment",
		ResourceID:   1,
	}).Error
	s.ErrorContains(err, "violates unique constraint")
}

func (s *modelSuite) TestCiIdentifierUniquenessSqlValid() {
	id1 := CiIdentifier{
		ResourceType: "environment",
		ResourceID:   1,
	}
	err := s.DB.Create(&id1).Error
	s.NoError(err)
	s.NotZero(id1.ID)
	err = s.DB.Delete(&id1).Error
	s.NoError(err)
	id2 := CiIdentifier{
		ResourceType: "environment",
		ResourceID:   1,
	}
	err = s.DB.Create(&id2).Error
	s.NoError(err)
	s.NotZero(id2.ID)
	s.NotEqual(id1.ID, id2.ID)
}
