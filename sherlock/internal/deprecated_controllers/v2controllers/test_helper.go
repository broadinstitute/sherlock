package v2controllers

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/authentication_method"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/test_users"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"gorm.io/gorm"
	"testing"
)

func generateUser(t *testing.T, db *gorm.DB, suitable bool) *models.User {
	var email, googleID string
	if suitable {
		email = test_users.SuitableTestUserEmail
		googleID = test_users.SuitableTestUserGoogleID
	} else {
		email = test_users.NonSuitableTestUserEmail
		googleID = test_users.NonSuitableTestUserGoogleID
	}
	var result models.User
	if err := db.Where(&models.User{Email: email, GoogleID: googleID}).FirstOrCreate(&result).Error; err != nil {
		t.Error(err)
		return nil
	}
	result.AuthenticationMethod = authentication_method.TEST
	return &result
}
