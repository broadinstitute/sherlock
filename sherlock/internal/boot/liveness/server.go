package liveness

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// Server runs a liveness endpoint designed to integrate with Sherlock's lifecycle.
//
// Server.Start should be run once the sql.DB is connected. It will begin monitoring the health of the connection
// pool courtesy of sql.DB's Ping method and reporting that health on port 8081.
// Server.MakeAlwaysReturnOK should be run during shutdown before the sql.DB connection pool is drained. It stops
// monitoring of the sql.DB and sets the endpoint to always report healthy, so Kubernetes won't think Sherlock is
// unhealthy as it shuts down.
// Server.Stop should be called right before Sherlock exits. It shuts down the monitoring endpoint.
//
// Note that we're potentially being overly correct here. If Kubernetes were to detect Sherlock as unhealthy
// during shutdown, it could potentially send another SIGTERM, but that's not really an issue. We'd already
// be reacting to a SIGTERM, and the new one would have a longer deadline than the one we'd already be on.
type Server struct {
	db            *gorm.DB
	handler       *handler
	server        *http.Server
	pingCtx       context.Context
	cancelPingCtx context.CancelFunc
}

func evaluate(db *gorm.DB, ctx context.Context) bool {
	if sqlDB, err := db.DB(); err != nil {
		log.Error().Err(err).Msgf("LIVE | liveness.evaluate()...db.DB() error")
		return false
	} else if err = sqlDB.PingContext(ctx); err != nil {
		log.Error().Err(err).Msgf("LIVE | liveness.evaluate()...sqlDB.PingContext() error")
		return false
	} else {
		return true
	}
}

func (s *Server) Start(db *gorm.DB) {
	s.db = db
	s.pingCtx, s.cancelPingCtx = context.WithCancel(context.Background())
	s.handler = &handler{
		returnOK: evaluate(s.db, s.pingCtx),
	}
	s.server = &http.Server{
		Addr:    ":8081",
		Handler: s.handler,
	}
	go s.repeatedlyPingDatabase()
	if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal().Msgf("LIVE | liveness.Server.server.ListenAndServe() err: %v", err)
	}
}

func (s *Server) repeatedlyPingDatabase() {
	interval := config.Config.MustDuration("db.livenessPingInterval")
	for {
		s.handler.mutex.Lock()
		s.handler.returnOK = evaluate(s.db, s.pingCtx)
		s.handler.mutex.Unlock()
		select {
		case <-time.After(interval):
		case <-s.pingCtx.Done():
			return
		}
	}
}

// MakeAlwaysReturnOK exists to make the liveness Server suddenly not care about the database connection anymore.
// This is to allow the database connection to be shut down without the liveness endpoint reporting that Sherlock
// is unhealthy.
// In other words, there's actually one time when we want Sherlock to report as "alive" even if its database
// connection is offline: during shutdown. This function exists to facilitate that.
func (s *Server) MakeAlwaysReturnOK() {
	if s.cancelPingCtx != nil {
		s.cancelPingCtx()
	}
	if s.handler != nil {
		s.handler.mutex.Lock()
		s.handler.returnOK = true
		s.handler.mutex.Unlock()
	}
}

func (s *Server) Stop() {
	if s.cancelPingCtx != nil {
		s.cancelPingCtx()
	}
	if s.server != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		if err := s.server.Shutdown(ctx); err != nil {
			log.Fatal().Msgf("LIVE | liveness.Server.server.Shutdown() err: %v", err)
		}
	}
}

type handler struct {
	// mutex is here because technically both ServeHTTP and Server.repeatedlyPingDatabase goroutines hit returnOK.
	// It's a sync.RWMutex because it's read from ServeHTTP and only written to by Server.repeatedlyPingDatabase on
	// an interval.
	mutex    sync.RWMutex
	returnOK bool
}

func (h *handler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	h.mutex.RLock()
	ok := h.returnOK
	h.mutex.RUnlock()
	if ok {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("NOT OK"))
	}
}
