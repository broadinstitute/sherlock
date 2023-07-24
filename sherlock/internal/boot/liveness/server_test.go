package liveness

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/db"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer_Lifecycle(t *testing.T) {
	config.LoadTestConfig()
	sqlDB, err := db.Connect()
	assert.NoError(t, err)
	server := &Server{}
	go server.Start(sqlDB)
	t.Run("online probe", func(t *testing.T) {
		var livenessSucceeded bool
		attemptsRemaining := 4 * 20
		for ; attemptsRemaining >= 0 && !livenessSucceeded; attemptsRemaining-- {
			resp, err := http.Get("http://localhost:8081")
			if err == nil && resp.StatusCode == http.StatusOK {
				livenessSucceeded = true
			} else {
				time.Sleep(time.Second / 4)
			}
		}
		assert.True(t, livenessSucceeded)
	})
	t.Run("offline probe", func(t *testing.T) {
		server.MakeAlwaysReturnOK()
		resp, err := http.Get("http://localhost:8081")
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
	server.Stop()
}

func TestServer_MakeAlwaysReturnOK(t *testing.T) {
	var cancelCalled bool
	server := &Server{
		cancelPingCtx: func() {
			cancelCalled = true
		},
		handler: &handler{
			returnOK: false,
		},
	}
	server.MakeAlwaysReturnOK()
	assert.True(t, cancelCalled)
	assert.True(t, server.handler.returnOK)
}

func Test_handler_ServeHTTP(t *testing.T) {
	t.Run("returns OK", func(t *testing.T) {
		handlerInstance := &handler{}
		handlerInstance.returnOK = true
		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		handlerInstance.ServeHTTP(w, req)
		resp := w.Result()
		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "OK", string(body))
	})
	t.Run("returns NOT OK", func(t *testing.T) {
		handlerInstance := &handler{}
		handlerInstance.returnOK = false
		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		handlerInstance.ServeHTTP(w, req)
		resp := w.Result()
		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		assert.Equal(t, "NOT OK", string(body))
	})
}
