package boot

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/db"
	"github.com/stretchr/testify/suite"
	"net/http"
	"sync"
	"testing"
	"time"
)

type applicationSuite struct {
	suite.Suite
}

func TestApplicationSuite(t *testing.T) {
	suite.Run(t, new(applicationSuite))
}

func (s *applicationSuite) SetupSuite() {
	config.LoadTestConfig()
}

func (s *applicationSuite) TestApplication_StartStop() {
	application := &Application{
		runInsideDatabaseTransaction: true,
	}

	s.Run("start and then stop", func() {
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
		s.Truef(livenessSucceeded, ":8081 returned 200")
		s.Truef(readinessSucceeded, ":8080/status returned 200")
	})
}

func (s *applicationSuite) TestApplication_dbMigrationLock() {
	gormDB, cleanup, err := db.Connect()
	s.NoError(err)
	application := &Application{
		gormDB: gormDB,
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
		s.True(completionDesired)
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

	// Have to make the linter stand down
	//nolint:staticcheck
	completedMarker.Unlock()

	sqlDB, err := gormDB.DB()
	s.NoError(err)
	s.NotNil(sqlDB)
	err = sqlDB.Close()
	s.NoError(err)
	err = cleanup()
	s.NoError(err)
}
