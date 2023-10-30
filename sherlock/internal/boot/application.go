package boot

import (
	"context"
	"database/sql"
	"errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/gha_oidc"
	"github.com/broadinstitute/sherlock/sherlock/internal/authorization"
	"github.com/broadinstitute/sherlock/sherlock/internal/boot/liveness"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/sherlock/internal/deployhooks"
	"github.com/broadinstitute/sherlock/sherlock/internal/github"
	"github.com/broadinstitute/sherlock/sherlock/internal/metrics"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"net/http"
	"sync"
	"time"
)

type Application struct {
	sqlDB          *sql.DB
	livenessServer *liveness.Server
	// dbMigrationLock lets us manually protect database migrations by trying to block shutdown until it
	// completes. If we drain the database connection pool while a migration is running, the migration
	// could fail or it could be unable to make a new query to release its lock even if it succeeded.
	dbMigrationLock sync.Mutex
	gormDB          *gorm.DB
	cancelCtx       context.CancelFunc
	server          *http.Server

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
	a.livenessServer = &liveness.Server{}
	go a.livenessServer.Start(a.sqlDB)

	log.Info().Msgf("BOOT | migrating database and configuring Gorm...")
	a.dbMigrationLock.Lock()
	gormDB, err := db.Configure(a.sqlDB)
	a.dbMigrationLock.Unlock()
	if err != nil {
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

	log.Info().Msgf("BOOT | initializing deploy detection...")
	if err = deployhooks.Init(); err != nil {
		log.Fatal().Msgf("deployhooks.Init() err: %v", err)
	}

	if config.Config.MustString("mode") != "debug" {
		log.Info().Msgf("BOOT | caching Firecloud accounts...")
		if err = authorization.CacheFirecloudSuitability(ctx); err != nil {
			log.Fatal().Msgf("authorization.CacheFirecloudSuitability() err: %v", err)
		}
		go authorization.KeepFirecloudCacheUpdated(ctx)
	}

	log.Info().Msgf("BOOT | reading extra permissions defined in configuration...")
	if err = authorization.CacheConfigSuitability(); err != nil {
		log.Fatal().Msgf("authorization.CacheConfigSuitability() err: %v", err)
	}

	log.Info().Msgf("BOOT | initializing GitHub Actions OIDC token verification...")
	if err = gha_oidc.InitVerifier(ctx); err != nil {
		log.Fatal().Msgf("gha_oidc_auth.InitVerifier() err: %v", err)
	}

	if config.Config.Bool("metrics.v2.enable") {
		log.Info().Msgf("BOOT | registering metric views...")
		if err = metrics.RegisterViews(); err != nil {
			log.Fatal().Msgf("metrics.RegisterViews() err: %v", err)
		}

		log.Info().Msgf("BOOT | calculating metric values...")
		if err = models.UpdateMetrics(ctx, a.gormDB); err != nil {
			log.Fatal().Msgf("models.UpdateMetrics() err: %v", err)
		}

		go models.KeepMetricsUpdated(ctx, a.gormDB)
	}

	if config.Config.Bool("slack.enable") {
		log.Info().Msgf("BOOT | initializing Slack socket...")
		if err = slack.Init(ctx); err != nil {
			log.Fatal().Msgf("slack.Init() err: %v", err)
		}
		go slack.Start(ctx)
	}

	if config.Config.Bool("github.enable") {
		log.Info().Msgf("BOOT | initializing GitHub client...")
		if err = github.Init(ctx); err != nil {
			log.Fatal().Msgf("github.Init() err: %v", err)
		}
	}

	log.Info().Msgf("BOOT | building Gin router...")
	a.server = &http.Server{
		Addr:    ":8080",
		Handler: buildRouter(ctx, a.gormDB),
	}

	log.Info().Msgf("BOOT | boot complete; now serving...")
	if err = a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
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

	if a.livenessServer != nil {
		log.Info().Msgf("BOOT | making liveness endpoint not check database anymore...")
		a.livenessServer.MakeAlwaysReturnOK()
	} else {
		log.Info().Msgf("BOOT | no liveness server reference, skipping making liveness endpoint not check database")
	}

	if a.sqlDB != nil {
		log.Info().Msgf("BOOT | closing database connections...")
		if wasUnlocked := a.dbMigrationLock.TryLock(); !wasUnlocked {
			log.Info().Msgf("BOOT | detected database migration underway, attempting to wait until it completes...")
			a.dbMigrationLock.Lock()
			log.Info().Msgf("BOOT | database migration complete, proceeding with connection close...")
		}
		if err := a.sqlDB.Close(); err != nil {
			log.Warn().Msgf("BOOT | database connection close error: %v", err)
		}
		a.dbMigrationLock.Unlock()
	} else {
		log.Info().Msgf("BOOT | no SQL database reference, skipping closing database connections")
	}

	if a.livenessServer != nil {
		log.Info().Msgf("BOOT | stopping liveness endpoint...")
		a.livenessServer.Stop()
	} else {
		log.Info().Msgf("BOOT | no liveness server reference, skipping stopping liveness endpoint")
	}

	log.Info().Msgf("BOOT | exiting...")
}
