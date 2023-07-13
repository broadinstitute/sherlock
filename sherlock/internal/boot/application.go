package boot

import (
	"context"
	"database/sql"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/gha_oidc"
	"github.com/broadinstitute/sherlock/sherlock/internal/authorization"
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

	// runInsideDatabaseTransaction begins a transaction on the gorm.DB after migration and rolls it back
	// before closing the connection. This makes Start + Stop safe to run from tests, because they won't
	// leave the database in a dirty state.
	runInsideDatabaseTransaction bool
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

	log.Info().Msgf("BOOT | migrating database and configuring Gorm...")
	if gormDB, err := db.Configure(a.sqlDB); err != nil {
		log.Fatal().Msgf("db.Configure() err: %v", err)
	} else {
		a.gormDB = gormDB
	}

	if a.runInsideDatabaseTransaction {
		a.gormDB = a.gormDB.Begin()
	}

	log.Info().Msgf("BOOT | creating global context...")
	ctx, cancelCtx := context.WithCancel(context.Background())
	a.cancelCtx = cancelCtx

	if config.Config.MustString("mode") != "debug" {
		log.Info().Msgf("BOOT | caching Firecloud accounts...")
		if err := authorization.CacheFirecloudSuitability(ctx); err != nil {
			log.Fatal().Msgf("authorization.CacheFirecloudSuitability() err: %v", err)
		}
		go authorization.KeepFirecloudCacheUpdated(ctx)
	}

	log.Info().Msgf("BOOT | reading extra permissions defined in configuration...")
	if err := authorization.CacheConfigSuitability(); err != nil {
		log.Fatal().Msgf("authorization.CacheConfigSuitability() err: %v", err)
	}

	log.Info().Msgf("BOOT | initializing GitHub Actions OIDC token verification...")
	if err := gha_oidc.InitVerifier(ctx); err != nil {
		log.Fatal().Msgf("gha_oidc_auth.InitVerifier() err: %v", err)
	}

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

	log.Info().Msgf("BOOT | boot complete; now serving...")
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
			log.Warn().Msgf("BOOT | server shutdown error: %v", err)
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

	if a.runInsideDatabaseTransaction && a.gormDB != nil {
		a.gormDB.Rollback()
	}

	if a.sqlDB != nil {
		log.Info().Msgf("BOOT | closing database connections...")
		if err := a.sqlDB.Close(); err != nil {
			log.Warn().Msgf("BOOT | database connection close error: %v", err)
		}
	} else {
		log.Info().Msgf("BOOT | no SQL database reference, skipping closing database connections")
	}

	log.Info().Msgf("BOOT | stopping liveness endpoint...")
	liveness.Stop()
	log.Info().Msgf("BOOT | exiting...")
}
