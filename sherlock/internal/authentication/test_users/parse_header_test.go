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
		email, googleID := ParseHeader(&gin.Context{Request: req})
		assert.Equal(t, SuitableTestUserEmail, email)
		assert.Equal(t, SuitableTestUserGoogleID, googleID)
	})
	t.Run("explicitly suitable", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		assert.NoError(t, err)
		req.Header.Set(SuitabilityControlHeader, strconv.FormatBool(true))
		email, googleID := ParseHeader(&gin.Context{Request: req})
		assert.Equal(t, SuitableTestUserEmail, email)
		assert.Equal(t, SuitableTestUserGoogleID, googleID)
	})
	t.Run("explicitly non-suitable", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		assert.NoError(t, err)
		req.Header.Set(SuitabilityControlHeader, strconv.FormatBool(false))
		email, googleID := ParseHeader(&gin.Context{Request: req})
		assert.Equal(t, NonSuitableTestUserEmail, email)
		assert.Equal(t, NonSuitableTestUserGoogleID, googleID)
	})
	t.Run("panics if can't parse header", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		assert.NoError(t, err)
		req.Header.Set(SuitabilityControlHeader, "something that isn't boolean")
		assert.Panics(t, func() {
			ParseHeader(&gin.Context{Request: req})
		})
	})
}
