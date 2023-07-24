package liveness

import (
	"context"
	"database/sql"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

type Server struct {
	sqlDB         *sql.DB
	handler       *handler
	server        *http.Server
	pingCtx       context.Context
	cancelPingCtx context.CancelFunc
}

func (s *Server) Start(sqlDB *sql.DB) {
	s.sqlDB = sqlDB
	s.pingCtx, s.cancelPingCtx = context.WithCancel(context.Background())
	s.handler = &handler{
		returnOK: s.sqlDB.PingContext(s.pingCtx) == nil,
	}
	s.server = &http.Server{
		Addr:    ":8081",
		Handler: s.handler,
	}
	go s.repeatedlyPingDatabase()
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Msgf("LIVE | liveness.Server.server.ListenAndServe() err: %v", err)
	}
}

func (s *Server) repeatedlyPingDatabase() {
	interval := config.Config.MustDuration("db.livenessPingInterval")
	for {
		if err := s.sqlDB.PingContext(s.pingCtx); err == nil {
			s.handler.returnOK = true
		} else {
			s.handler.returnOK = false
			log.Error().Msgf("LIVE | liveness.Server.sqlDB.PingContext(liveness.Server.pingCtx) err: %v", err)
		}
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
		s.handler.returnOK = true
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
	returnOK bool
}

func (h *handler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	if h.returnOK {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("NOT OK"))
	}
}
