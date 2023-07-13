package boot

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestApplication(t *testing.T) {
	config.LoadTestConfig()
	application := &Application{
		runInsideDatabaseTransaction: true,
	}
	go application.Start()
	var livenessSucceeded, readinessSucceeded bool
	attemptsRemaining := 4 * 20
	for ; attemptsRemaining >= 0 && !livenessSucceeded; attemptsRemaining-- {
		resp, err := http.Get("http://localhost:8081")
		if err == nil && resp.StatusCode == 200 {
			livenessSucceeded = true
		} else {
			time.Sleep(time.Second / 4)
		}
	}
	for ; attemptsRemaining >= 0 && !readinessSucceeded; attemptsRemaining-- {
		resp, err := http.Get("http://localhost:8080/status")
		if err == nil && resp.StatusCode == 200 {
			readinessSucceeded = true
		} else {
			time.Sleep(time.Second / 4)
		}
	}
	application.Stop()
	assert.Truef(t, livenessSucceeded, ":8081 returned 200")
	assert.Truef(t, readinessSucceeded, ":8080/status returned 200")
}
