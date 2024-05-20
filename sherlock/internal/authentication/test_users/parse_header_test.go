package test_users

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestParseHeader(t *testing.T) {
	superAdminTestUserEmail := "super admin test user email"
	superAdminTestUserGoogleID := "super admin test user google id"
	suitableTestUserEmail := "suitable test user email"
	suitableTestUserGoogleID := "suitable test user google id"
	nonSuitableTestUserEmail := "non-suitable test user email"
	nonSuitableTestUserGoogleID := "non-suitable test user google id"
	parseHeader := MakeHeaderParser(
		models.User{
			Email:    superAdminTestUserEmail,
			GoogleID: superAdminTestUserGoogleID,
		},
		models.User{
			Email:    suitableTestUserEmail,
			GoogleID: suitableTestUserGoogleID,
		},
		models.User{
			Email:    nonSuitableTestUserEmail,
			GoogleID: nonSuitableTestUserGoogleID,
		},
	)
	t.Run("default", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		assert.NoError(t, err)
		email, googleID, err := parseHeader(&gin.Context{Request: req})
		assert.NoError(t, err)
		assert.Equal(t, suitableTestUserEmail, email)
		assert.Equal(t, suitableTestUserGoogleID, googleID)
	})
	t.Run("explicitly super admin", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		assert.NoError(t, err)
		req.Header.Set(superAdminControlHeader, strconv.FormatBool(true))
		email, googleID, err := parseHeader(&gin.Context{Request: req})
		assert.NoError(t, err)
		assert.Equal(t, superAdminTestUserEmail, email)
		assert.Equal(t, superAdminTestUserGoogleID, googleID)
	})
	t.Run("explicitly non-super admin", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		assert.NoError(t, err)
		req.Header.Set(superAdminControlHeader, strconv.FormatBool(false))
		email, googleID, err := parseHeader(&gin.Context{Request: req})
		assert.NoError(t, err)
		assert.Equal(t, suitableTestUserEmail, email)
		assert.Equal(t, suitableTestUserGoogleID, googleID)
	})
	t.Run("explicitly suitable", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		assert.NoError(t, err)
		req.Header.Set(suitableControlHeader, strconv.FormatBool(true))
		email, googleID, err := parseHeader(&gin.Context{Request: req})
		assert.NoError(t, err)
		assert.Equal(t, suitableTestUserEmail, email)
		assert.Equal(t, suitableTestUserGoogleID, googleID)
	})
	t.Run("explicitly non-suitable", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		assert.NoError(t, err)
		req.Header.Set(suitableControlHeader, strconv.FormatBool(false))
		email, googleID, err := parseHeader(&gin.Context{Request: req})
		assert.NoError(t, err)
		assert.Equal(t, nonSuitableTestUserEmail, email)
		assert.Equal(t, nonSuitableTestUserGoogleID, googleID)
	})
	t.Run("errors if can't parse superAdminControlHeader", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		assert.NoError(t, err)
		req.Header.Set(superAdminControlHeader, "something that isn't boolean")
		_, _, err = parseHeader(&gin.Context{Request: req})
		assert.ErrorContains(t, err, "failed to parse boolean")
	})
	t.Run("errors if can't parse suitableControlHeader", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		assert.NoError(t, err)
		req.Header.Set(suitableControlHeader, "something that isn't boolean")
		_, _, err = parseHeader(&gin.Context{Request: req})
		assert.ErrorContains(t, err, "failed to parse boolean")
	})
	t.Run("errors if super admin is asked to be non-suitable", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		assert.NoError(t, err)
		req.Header.Set(superAdminControlHeader, strconv.FormatBool(true))
		req.Header.Set(suitableControlHeader, strconv.FormatBool(false))
		_, _, err = parseHeader(&gin.Context{Request: req})
		assert.ErrorContains(t, err, "super admin cannot be non-suitable")
	})
}
