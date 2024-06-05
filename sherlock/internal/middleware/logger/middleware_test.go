package logger

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack/slack_mocks"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogger(t *testing.T) {
	gin.SetMode(gin.TestMode)
	config.LoadTestConfig()
	zerolog.SetGlobalLevel(zerolog.Disabled)
	router := gin.New()
	router.Use(Logger())
	router.GET("/:code", func(c *gin.Context) {
		code, err := utils.ParseInt(c.Param("code"))
		assert.NoError(t, err)
		if code > 399 {
			_ = c.Error(fmt.Errorf("code %d error", code))
		}
		c.JSON(code, gin.H{})
	})

	t.Run("400 doesn't send anything", func(t *testing.T) {
		slack.UseMockedClient(t, func(_ *slack_mocks.MockMockableClient) {}, func() {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/400", nil)
			router.ServeHTTP(recorder, request)

			assert.Equal(t, http.StatusBadRequest, recorder.Code)
		})
	})
	t.Run("407 sends", func(t *testing.T) {
		slack.UseMockedClient(t, func(c *slack_mocks.MockMockableClient) {
			c.On("SendMessageContext", mock.Anything, "channel 1",
				mock.AnythingOfType("slack.MsgOption"),
				mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
			c.On("SendMessageContext", mock.Anything, "channel 2",
				mock.AnythingOfType("slack.MsgOption"),
				mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
		}, func() {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/407", nil)
			router.ServeHTTP(recorder, request)

			assert.Equal(t, http.StatusProxyAuthRequired, recorder.Code)
		})
	})
	t.Run("500 sends", func(t *testing.T) {
		slack.UseMockedClient(t, func(c *slack_mocks.MockMockableClient) {
			c.On("SendMessageContext", mock.Anything, "channel 1",
				mock.AnythingOfType("slack.MsgOption"),
				mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
			c.On("SendMessageContext", mock.Anything, "channel 2",
				mock.AnythingOfType("slack.MsgOption"),
				mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
		}, func() {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/500", nil)
			router.ServeHTTP(recorder, request)

			assert.Equal(t, http.StatusInternalServerError, recorder.Code)
		})
	})
}
