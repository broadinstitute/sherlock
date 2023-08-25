package slack

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack/slack_mocks"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReportError(t *testing.T) {
	config.LoadTestConfig()
	zerolog.SetGlobalLevel(zerolog.Disabled)
	ctx := context.Background()
	type args struct {
		errs []error
	}
	tests := []struct {
		name       string
		args       args
		mockConfig func(c *slack_mocks.MockMockableClient)
	}{
		{
			name: "normal case",
			args: args{errs: []error{fmt.Errorf("some error")}},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.On("SendMessageContext", ctx, "channel 1",
					mock.AnythingOfType("slack.MsgOption"),
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
				c.On("SendMessageContext", ctx, "channel 2",
					mock.AnythingOfType("slack.MsgOption"),
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
			},
		},
		{
			name:       "sends no errors",
			args:       args{errs: []error{}},
			mockConfig: func(c *slack_mocks.MockMockableClient) {},
		},
		{
			name: "sends multiple errors",
			args: args{errs: []error{fmt.Errorf("some error"), fmt.Errorf("some second error")}},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.On("SendMessageContext", ctx, "channel 1",
					mock.AnythingOfType("slack.MsgOption"),
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
				c.On("SendMessageContext", ctx, "channel 2",
					mock.AnythingOfType("slack.MsgOption"),
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
				c.On("SendMessageContext", ctx, "channel 1",
					mock.AnythingOfType("slack.MsgOption"),
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
				c.On("SendMessageContext", ctx, "channel 2",
					mock.AnythingOfType("slack.MsgOption"),
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
			},
		},
		{
			name: "sending on channel 1 errors",
			args: args{errs: []error{fmt.Errorf("some error")}},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.On("SendMessageContext", ctx, "channel 1",
					mock.AnythingOfType("slack.MsgOption"),
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", fmt.Errorf("some send error"))
				c.On("SendMessageContext", ctx, "channel 2",
					mock.AnythingOfType("slack.MsgOption"),
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
			},
		},
		{
			name: "sending on channel 2 errors",
			args: args{errs: []error{fmt.Errorf("some error")}},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.On("SendMessageContext", ctx, "channel 1",
					mock.AnythingOfType("slack.MsgOption"),
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
				c.On("SendMessageContext", ctx, "channel 2",
					mock.AnythingOfType("slack.MsgOption"),
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", fmt.Errorf("some send error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UseMockedClient(t, tt.mockConfig, func() {
				ReportError(ctx, tt.args.errs...)
			})
		})
	}
}

func TestErrorReportingMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	config.LoadTestConfig()
	zerolog.SetGlobalLevel(zerolog.Disabled)
	ctx := context.Background()
	router := gin.New()
	router.Use(ErrorReportingMiddleware(ctx))
	router.GET("/:code", func(c *gin.Context) {
		code, err := utils.ParseInt(c.Param("code"))
		assert.NoError(t, err)
		if code > 399 {
			_ = c.Error(fmt.Errorf("code %d error", code))
		}
		c.JSON(code, gin.H{})
	})

	t.Run("400 doesn't send anything", func(t *testing.T) {
		UseMockedClient(t, func(_ *slack_mocks.MockMockableClient) {}, func() {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/400", nil)
			router.ServeHTTP(recorder, request)

			assert.Equal(t, http.StatusBadRequest, recorder.Code)
		})
	})
	t.Run("407 sends", func(t *testing.T) {
		UseMockedClient(t, func(c *slack_mocks.MockMockableClient) {
			c.On("SendMessageContext", ctx, "channel 1",
				mock.AnythingOfType("slack.MsgOption"),
				mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
			c.On("SendMessageContext", ctx, "channel 2",
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
		UseMockedClient(t, func(c *slack_mocks.MockMockableClient) {
			c.On("SendMessageContext", ctx, "channel 1",
				mock.AnythingOfType("slack.MsgOption"),
				mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
			c.On("SendMessageContext", ctx, "channel 2",
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
