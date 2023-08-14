package slack

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
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
		err error
	}
	tests := []struct {
		name       string
		args       args
		mockConfig func(client *mockMockableClient)
	}{
		{
			name: "normal case",
			args: args{err: fmt.Errorf("some error")},
			mockConfig: func(client *mockMockableClient) {
				client.On("SendMessageContext", ctx, "channel 1",
					mock.AnythingOfType("slack.MsgOption"),
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
				client.On("SendMessageContext", ctx, "channel 2",
					mock.AnythingOfType("slack.MsgOption"),
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
			},
		},
		{
			name: "1 errors",
			args: args{err: fmt.Errorf("some error")},
			mockConfig: func(client *mockMockableClient) {
				client.On("SendMessageContext", ctx, "channel 1",
					mock.AnythingOfType("slack.MsgOption"),
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", fmt.Errorf("some send error"))
				client.On("SendMessageContext", ctx, "channel 2",
					mock.AnythingOfType("slack.MsgOption"),
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
			},
		},
		{
			name: "2 errors",
			args: args{err: fmt.Errorf("some error")},
			mockConfig: func(client *mockMockableClient) {
				client.On("SendMessageContext", ctx, "channel 1",
					mock.AnythingOfType("slack.MsgOption"),
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
				client.On("SendMessageContext", ctx, "channel 2",
					mock.AnythingOfType("slack.MsgOption"),
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", fmt.Errorf("some send error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newMockMockableClient(t)
			tt.mockConfig(c)
			client = c
			ReportError(ctx, tt.args.err)
			c.AssertExpectations(t)
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
		c := newMockMockableClient(t)
		client = c

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/400", nil)
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		c.AssertExpectations(t)
	})
	t.Run("407 sends", func(t *testing.T) {
		c := newMockMockableClient(t)
		client = c

		c.On("SendMessageContext", ctx, "channel 1",
			mock.AnythingOfType("slack.MsgOption"),
			mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
		c.On("SendMessageContext", ctx, "channel 2",
			mock.AnythingOfType("slack.MsgOption"),
			mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/407", nil)
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusProxyAuthRequired, recorder.Code)
		c.AssertExpectations(t)
	})
	t.Run("500 sends", func(t *testing.T) {
		c := newMockMockableClient(t)
		client = c

		c.On("SendMessageContext", ctx, "channel 1",
			mock.AnythingOfType("slack.MsgOption"),
			mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
		c.On("SendMessageContext", ctx, "channel 2",
			mock.AnythingOfType("slack.MsgOption"),
			mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/500", nil)
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
		c.AssertExpectations(t)
	})
}
