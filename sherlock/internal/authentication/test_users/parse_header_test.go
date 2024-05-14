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
	suitableTestUserEmail := "suitable test user email"
	suitableTestUserGoogleID := "suitable test user google id"
	nonSuitableTestUserEmail := "non-suitable test user email"
	nonSuitableTestUserGoogleID := "non-suitable test user google id"
	parseHeader := MakeHeaderParser(
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
	t.Run("errors if can't parse suitableControlHeader", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		assert.NoError(t, err)
		req.Header.Set(suitableControlHeader, "something that isn't boolean")
		_, _, err = parseHeader(&gin.Context{Request: req})
		assert.ErrorContains(t, err, "failed to parse boolean")
	})
}
