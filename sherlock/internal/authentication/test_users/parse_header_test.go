package test_users

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestParseHeader(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		assert.NoError(t, err)
		email, googleID, err := ParseHeader(&gin.Context{Request: req})
		assert.NoError(t, err)
		assert.Equal(t, SuitableTestUserEmail, email)
		assert.Equal(t, SuitableTestUserGoogleID, googleID)
	})
	t.Run("explicitly suitable", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		assert.NoError(t, err)
		req.Header.Set(suitableControlHeader, strconv.FormatBool(true))
		email, googleID, err := ParseHeader(&gin.Context{Request: req})
		assert.NoError(t, err)
		assert.Equal(t, SuitableTestUserEmail, email)
		assert.Equal(t, SuitableTestUserGoogleID, googleID)
	})
	t.Run("explicitly non-suitable", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		assert.NoError(t, err)
		req.Header.Set(suitableControlHeader, strconv.FormatBool(false))
		email, googleID, err := ParseHeader(&gin.Context{Request: req})
		assert.NoError(t, err)
		assert.Equal(t, NonSuitableTestUserEmail, email)
		assert.Equal(t, NonSuitableTestUserGoogleID, googleID)
	})
	t.Run("errors if can't parse suitableControlHeader", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		assert.NoError(t, err)
		req.Header.Set(suitableControlHeader, "something that isn't boolean")
		_, _, err = ParseHeader(&gin.Context{Request: req})
		assert.ErrorContains(t, err, "failed to parse boolean")
	})
}
