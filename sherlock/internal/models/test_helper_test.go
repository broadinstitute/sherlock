package models

import "github.com/broadinstitute/sherlock/sherlock/internal/authentication/test_users"

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
		s.Equal(test_users.SuitableTestUserEmail, user.Email)
		s.NotZero(user.ID)
		s.True(user.DeprecatedSuitability().Suitable())
	})
	s.Run("non-suitable test user", func() {
		s.SetNonSuitableTestUserForDB()
		user, err := GetCurrentUserForDB(s.DB)
		s.NoError(err)
		s.Equal(test_users.NonSuitableTestUserEmail, user.Email)
		s.NotZero(user.ID)
		s.False(user.DeprecatedSuitability().Suitable())
	})
}
