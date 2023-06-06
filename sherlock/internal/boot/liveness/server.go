package liveness

import (
	"context"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

var livenessServer = &http.Server{
	Addr:    ":8081",
	Handler: livenessHandler{},
}

type livenessHandler struct{}

func (h livenessHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("OK"))
}

func Start() {
	if err := livenessServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Msgf("livenessServer.ListenAndServe() err: %v", err)
	}
}

func Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := livenessServer.Shutdown(ctx); err != nil {
		log.Fatal().Msgf("livenessServer.Shutdown() err: %v", err)
	}
}
