package models

// TestTestHelperItself checks that modelSuite's user helper
// methods work properly (and thus also that the database
// connection is working).
func (s *modelSuite) TestTestHelperItself() {
	s.Run("no user by default", func() {
		user, err := GetCurrentUserForDB(s.DB)
		s.ErrorContains(err, "database user not available")
		s.Nil(user)
	})
	s.Run("suitable test user", func() {
		s.SetSuitableTestUserForDB()
		user, err := GetCurrentUserForDB(s.DB)
		s.NoError(err)
		s.Equal(s.TestData.User_Suitable().Email, user.Email)
		s.NotZero(user.ID)
		if s.NotNil(user.Suitability) {
			if s.NotNil(user.Suitability.Suitable) {
				s.True(*user.Suitability.Suitable)
			}
		}
	})
	s.Run("non-suitable test user", func() {
		s.SetNonSuitableTestUserForDB()
		user, err := GetCurrentUserForDB(s.DB)
		s.NoError(err)
		s.Equal(s.TestData.User_NonSuitable().Email, user.Email)
		s.NotZero(user.ID)
		if s.NotNil(user.Suitability) {
			if s.NotNil(user.Suitability.Suitable) {
				s.False(*user.Suitability.Suitable)
			}
		}
	})
}
