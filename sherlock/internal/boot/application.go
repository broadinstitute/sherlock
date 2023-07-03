package boot

import (
	"context"
	"database/sql"
	"github.com/broadinstitute/sherlock/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/sherlock/internal/auth/gha_oidc_auth"
	"github.com/broadinstitute/sherlock/sherlock/internal/boot/liveness"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"github.com/broadinstitute/sherlock/sherlock/internal/metrics"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type Application struct {
	sqlDB     *sql.DB
	gormDB    *gorm.DB
	cancelCtx context.CancelFunc
	server    *http.Server
}

func (a *Application) Start() {
	log.Info().Msgf("BOOT | connecting to database...")
	if sqlDB, err := db.Connect(); err != nil {
		log.Fatal().Msgf("db.Connect() err: %v", err)
	} else {
		a.sqlDB = sqlDB
	}

	log.Info().Msgf("BOOT | starting liveness endpoint...")
	go liveness.Start()

	log.Info().Msgf("BOOT | migrating database and configuring ORM...")
	if gormDB, err := db.Configure(a.sqlDB); err != nil {
		log.Fatal().Msgf("db.Configure() err: %v", err)
	} else {
		a.gormDB = gormDB
	}

	log.Info().Msgf("BOOT | creating global context...")
	ctx, cancelCtx := context.WithCancel(context.Background())
	a.cancelCtx = cancelCtx

	if config.Config.MustString("mode") != "debug" {
		log.Info().Msgf("BOOT | caching Firecloud accounts...")
		if err := auth.CacheFirecloudAccounts(ctx); err != nil {
			log.Fatal().Msgf("auth.CacheFirecloudAccounts() err: %v", err)
		}
		go auth.KeepFirecloudCacheUpdated(ctx)
	}

	log.Info().Msgf("BOOT | initializing GitHub Actions OIDC token verification...")
	if err := gha_oidc_auth.InitVerifier(ctx); err != nil {
		log.Fatal().Msgf("gha_oidc_auth.InitVerifier() err: %v", err)
	}

	log.Info().Msgf("BOOT | reading extra permissions defined in configuration...")
	auth.CacheExtraPermissions()

	if config.Config.Bool("metrics.v2.enable") {
		log.Info().Msgf("BOOT | registering metric views...")
		if err := metrics.RegisterViews(); err != nil {
			log.Fatal().Msgf("v2metrics.RegisterViews() err: %v", err)
		}

		log.Info().Msgf("BOOT | calculating metric values...")
		if err := v2models.UpdateMetrics(ctx, a.gormDB); err != nil {
			log.Fatal().Msgf("v2models.UpdateMetrics() err: %v", err)
		}

		go v2models.KeepMetricsUpdated(ctx, a.gormDB)
	}

	log.Info().Msgf("BOOT | building Gin router...")
	if router, err := buildRouter(a.gormDB); err != nil {
		log.Fatal().Msgf("buildRouter() err: %v", err)
	} else {
		a.server = &http.Server{
			Addr:    ":8080",
			Handler: router,
		}
	}

	log.Info().Msgf("BOOT | starting server...")
	if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Msgf("server.ListenAndServe() err: %v", err)
	}
}

func (a *Application) Stop() {
	if a.server != nil {
		log.Info().Msgf("BOOT | shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := a.server.Shutdown(ctx); err != nil {
			log.Warn().Msgf("server.Shutdown() err: %v", err)
		}
	} else {
		log.Info().Msgf("BOOT | no server; skipping shutting down server")
	}

	if a.cancelCtx != nil {
		log.Info().Msgf("BOOT | canceling global context...")
		a.cancelCtx()
	} else {
		log.Info().Msgf("BOOT | no context cancellation function, skipping canceling global context")
	}

	if a.sqlDB != nil {
		log.Info().Msgf("BOOT | closing database connections...")
		if err := a.sqlDB.Close(); err != nil {
			log.Warn().Msgf("sqlDB.Close() err: %v", err)
		}
	} else {
		log.Info().Msgf("BOOT | no SQL database reference, skipping closing database connections")
	}

	log.Info().Msgf("BOOT | stopping liveness endpoint...")
	liveness.Stop()
}
