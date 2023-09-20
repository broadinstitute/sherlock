package models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/authentication_method"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/test_users"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"gorm.io/gorm"
)

func (s *modelSuite) TestUserRejectEditImmutableField() {
	userToEdit := s.SetNonSuitableTestUserForDB()
	s.Run("email", func() {
		err := s.DB.Model(&userToEdit).Updates(&User{Email: "some-other-email@example.com"}).Error
		s.ErrorContains(err, "cannot be changed")
	})
	s.Run("google ID", func() {
		err := s.DB.Model(&userToEdit).Updates(&User{GoogleID: "some google ID?"}).Error
		s.ErrorContains(err, "cannot be changed")
	})
	s.Run("ID", func() {
		err := s.DB.Model(&userToEdit).Updates(&User{Model: gorm.Model{ID: userToEdit.ID + 1}}).Error
		s.ErrorContains(err, "cannot be changed")
	})
}

func (s *modelSuite) TestUserCatchBadEdit() {
	s.SetNonSuitableTestUserForDB()
	err := s.DB.Where(&User{Email: test_users.NonSuitableTestUserEmail}).Updates(&User{Name: utils.PointerTo("new name")}).Error
	s.ErrorContains(err, "user ID in BeforeEdit was nil, possibly a bad database call")
}

func (s *modelSuite) TestUserOnlySelfEdit() {
	nonSuitable := s.SetNonSuitableTestUserForDB()
	s.SetSuitableTestUserForDB()
	err := s.DB.Model(nonSuitable).Updates(&User{Name: utils.PointerTo("new name")}).Error
	s.ErrorContains(err, errors.Forbidden)
	s.SetNonSuitableTestUserForDB()
	err = s.DB.Model(nonSuitable).Updates(&User{Name: utils.PointerTo("new name")}).Error
	s.NoError(err)
	s.Equal("new name", *nonSuitable.Name)
}

func (s *modelSuite) TestUserEditAuthMethodCheck() {
	user := s.SetNonSuitableTestUserForDB()
	user.AuthenticationMethod = authentication_method.UNKNOWN
	var userToEdit User
	err := s.DB.Where(&User{Email: test_users.NonSuitableTestUserEmail}).First(&userToEdit).Error
	s.NoError(err)
	err = s.DB.Model(&userToEdit).Updates(&User{Name: utils.PointerTo("new name")}).Error
	s.ErrorContains(err, "users cannot be edited via this authentication method")
}

func (s *modelSuite) TestUserNoDelete() {
	s.SetNonSuitableTestUserForDB()
	s.SetSuitableTestUserForDB()
	var user User
	err := s.DB.Where(&User{Email: test_users.NonSuitableTestUserEmail}).Delete(&user).Error
	s.ErrorContains(err, errors.Forbidden)
	err = s.DB.Where(&User{Email: test_users.NonSuitableTestUserEmail}).First(&user).Error
	s.NoError(err)
	s.NotZero(user.ID)
}

func (s *modelSuite) TestUserSuitabilityAccess() {
	s.Run("suitable", func() {
		suitable := &User{Email: test_users.SuitableTestUserEmail}
		s.True(suitable.Suitability().Suitable())
		s.NotZero(suitable.cachedSuitability)
		s.True(suitable.cachedSuitability.Suitable())
	})
	s.Run("not suitable", func() {
		notSuitable := &User{Email: test_users.NonSuitableTestUserEmail}
		s.False(notSuitable.Suitability().Suitable())
		s.NotZero(notSuitable.cachedSuitability)
		s.False(notSuitable.cachedSuitability.Suitable())
	})
}

func (s *modelSuite) TestUserUsername() {
	tests := []struct {
		name  string
		email string
		want  string
	}{
		{
			name:  "normal BI username",
			email: "someone@broadinstitute.org",
			want:  "someone",
		},
		{
			name:  "with separators",
			email: "someone.else-blah_blah@somewhere.info",
			want:  "someone-else-blah-blah",
		},
		{
			name:  "strips invalid",
			email: "1a bc?de.23",
			want:  "1abcde-23",
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			testutils.AssertNoDiff(s.T(), tt.want, (&User{Email: tt.email}).AlphaNumericHyphenatedUsername())
		})
	}
}

func (s *modelSuite) TestUserGithubValidationSqlOnlyName() {
	user := s.SetNonSuitableTestUserForDB()
	err := s.DB.Model(user).Updates(&User{GithubUsername: utils.PointerTo("foo")}).Error
	s.ErrorContains(err, "violates check constraint \"github_info_together\"")
}

func (s *modelSuite) TestUserGithubValidationSqlOnlyID() {
	user := s.SetNonSuitableTestUserForDB()
	err := s.DB.Model(user).Updates(&User{GithubID: utils.PointerTo("bar")}).Error
	s.ErrorContains(err, "violates check constraint \"github_info_together\"")
}

func (s *modelSuite) TestUserGithubValidationSqlValid() {
	user := s.SetNonSuitableTestUserForDB()
	err := s.DB.Model(user).Updates(&User{
		GithubUsername: utils.PointerTo("foo"),
		GithubID:       utils.PointerTo("bar"),
	}).Error
	s.NoError(err)
}

func (s *modelSuite) TestUserEmailValidationSqlInvalid() {
	err := s.DB.Create(&User{Email: "invalid", GoogleID: "some value"}).Error
	s.ErrorContains(err, "violates check constraint \"email_format\"")
}

func (s *modelSuite) TestUserEmailValidationSqlValid() {
	err := s.DB.Create(&User{Email: "valid@example.com", GoogleID: "some value"}).Error
	s.NoError(err)
}

func (s *modelSuite) TestUserEmailUniquenessSql() {
	err := s.DB.Create(&User{Email: "valid@example.com", GoogleID: "some value"}).Error
	s.NoError(err)
	err = s.DB.Create(&User{Email: "valid@example.com", GoogleID: "some other value"}).Error
	s.ErrorContains(err, "violates unique constraint")
}

func (s *modelSuite) TestUserGoogleIdUniquenessSql() {
	err := s.DB.Create(&User{Email: "valid@example.com", GoogleID: "some value"}).Error
	s.NoError(err)
	err = s.DB.Create(&User{Email: "valid-2@example.com", GoogleID: "some value"}).Error
	s.ErrorContains(err, "violates unique constraint")
}

func (s *modelSuite) TestUserGithubUsernameUniquenessSql() {
	err := s.DB.Create(&User{
		Email:          "valid@example.com",
		GoogleID:       "some value",
		GithubUsername: utils.PointerTo("valid"),
		GithubID:       utils.PointerTo("some value"),
	}).Error
	s.NoError(err)
	err = s.DB.Create(&User{
		Email:          "valid-2@example.com",
		GoogleID:       "some other value",
		GithubUsername: utils.PointerTo("valid"),
		GithubID:       utils.PointerTo("some other value"),
	}).Error
	s.ErrorContains(err, "violates unique constraint")
}

func (s *modelSuite) TestUserGithubIdUniquenessSql() {
	err := s.DB.Create(&User{
		Email:          "valid@example.com",
		GoogleID:       "some value",
		GithubUsername: utils.PointerTo("valid"),
		GithubID:       utils.PointerTo("some value"),
	}).Error
	s.NoError(err)
	err = s.DB.Create(&User{
		Email:          "valid-2@example.com",
		GoogleID:       "some other value",
		GithubUsername: utils.PointerTo("valid-2"),
		GithubID:       utils.PointerTo("some value"),
	}).Error
	s.ErrorContains(err, "violates unique constraint")
}
