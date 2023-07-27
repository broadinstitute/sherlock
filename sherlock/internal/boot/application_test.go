package boot

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/db"
	"github.com/stretchr/testify/assert"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestApplication_StartStop(t *testing.T) {
	config.LoadTestConfig()
	application := &Application{
		runInsideDatabaseTransaction: true,
	}

	t.Run("start and then stop", func(t *testing.T) {
		go application.Start()
		var livenessSucceeded, readinessSucceeded bool
		attemptsRemaining := 4 * 20
		for ; attemptsRemaining >= 0 && !livenessSucceeded; attemptsRemaining-- {
			resp, err := http.Get("http://localhost:8081")
			if err == nil && resp.StatusCode == http.StatusOK {
				livenessSucceeded = true
			} else {
				time.Sleep(time.Second / 4)
			}
		}
		for ; attemptsRemaining >= 0 && !readinessSucceeded; attemptsRemaining-- {
			resp, err := http.Get("http://localhost:8080/status")
			if err == nil && resp.StatusCode == http.StatusOK {
				readinessSucceeded = true
			} else {
				time.Sleep(time.Second / 4)
			}
		}
		application.Stop()
		assert.Truef(t, livenessSucceeded, ":8081 returned 200")
		assert.Truef(t, readinessSucceeded, ":8080/status returned 200")
	})
}

func TestApplication_dbMigrationLock(t *testing.T) {
	config.LoadTestConfig()
	sqlDB, err := db.Connect()
	assert.NoError(t, err)
	application := &Application{
		sqlDB: sqlDB,
	}
	// Pretend that a migration is ongoing
	application.dbMigrationLock.Lock()

	// We want to block until the Stop() goroutine finishes, so we'll have it unlock this mutex
	completedMarker := sync.Mutex{}
	completedMarker.Lock()

	// We want to know that the Stop() call blocks until it shouldn't, so we have a boolean
	// with a mutex around it so that we can change whether the assertion passes or fails from
	// outside.
	completionDesired := false
	completionDesiredMutex := sync.Mutex{}
	go func() {
		application.Stop()
		completionDesiredMutex.Lock()
		assert.True(t, completionDesired)
		completionDesiredMutex.Unlock()

		// Unlock this to indicate that the test can exit
		completedMarker.Unlock()
	}()

	// Sleep so that we make sure the goroutine should be blocking on Stop()
	time.Sleep(time.Second)

	// Pretend that the migration just completed
	completionDesiredMutex.Lock()
	application.dbMigrationLock.Unlock()
	completionDesired = true
	completionDesiredMutex.Unlock()

	// Block for completion of the goroutine
	completedMarker.Lock()
	completedMarker.Unlock()
}
