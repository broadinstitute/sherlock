package boot

import (
	"context"
	"errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/boot/liveness"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/bits_data_warehouse"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/github"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/sherlock/internal/metrics"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication/gha_oidc"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/oidc_models"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation"
	"github.com/broadinstitute/sherlock/sherlock/internal/suitability_synchronization"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"net/http"
	"sync"
	"time"
)

type Application struct {
	dbDriverCleanup func() error
	gormDB          *gorm.DB
	livenessServer  *liveness.Server
	// dbMigrationLock lets us manually protect database migrations by trying to block shutdown until it
	// completes. If we drain the database connection pool while a migration is running, the migration
	// could fail or it could be unable to make a new query to release its lock even if it succeeded.
	dbMigrationLock sync.Mutex
	cancelCtx       context.CancelFunc
	server          *http.Server

	// runInsideDatabaseTransaction begins a transaction on the gorm.DB after migration and rolls it back
	// before closing the connection. This makes Start + Stop safe to run from tests, because they won't
	// leave the database in a dirty state.
	runInsideDatabaseTransaction bool
}

func (a *Application) Start() {
	log.Info().Msgf("BOOT | registering database driver...")
	if dbDriverCleanup, err := db.RegisterDriver(); err != nil {
		log.Fatal().Err(err).Msgf("db.RegisterDriver() error")
	} else {
		a.dbDriverCleanup = dbDriverCleanup
	}

	log.Info().Msgf("BOOT | connecting to database...")
	if gormDB, err := db.Connect(); err != nil {
		log.Fatal().Err(err).Msgf("db.Connect() error")
	} else {
		a.gormDB = gormDB
	}

	log.Info().Msgf("BOOT | starting liveness endpoint...")
	a.livenessServer = &liveness.Server{}
	go a.livenessServer.Start(a.gormDB)

	log.Info().Msgf("BOOT | migrating database and configuring Gorm...")
	a.dbMigrationLock.Lock()
	err := db.Migrate(a.gormDB)
	a.dbMigrationLock.Unlock()
	if err != nil {
		log.Fatal().Err(err).Msgf("db.Configure() error")
	} else if err = models.Init(a.gormDB); err != nil {
		log.Fatal().Err(err).Msgf("models.Init() error")
	}

	if a.runInsideDatabaseTransaction {
		a.gormDB = a.gormDB.Begin()
	}

	log.Info().Msgf("BOOT | creating global context...")
	ctx, cancelCtx := context.WithCancel(context.Background())
	a.cancelCtx = cancelCtx

	log.Info().Msgf("BOOT | initializing GitHub Actions OIDC token verification...")
	if err = gha_oidc.InitVerifier(ctx); err != nil {
		log.Fatal().Err(err).Msgf("gha_oidc_auth.InitVerifier() error")
	}

	if config.Config.Bool("metrics.v2.enable") {
		log.Info().Msgf("BOOT | registering metric views...")
		if err = metrics.RegisterViews(); err != nil {
			log.Fatal().Err(err).Msgf("metrics.RegisterViews() error")
		}

		log.Info().Msgf("BOOT | calculating metric values...")
		if err = models.UpdateMetrics(ctx, a.gormDB); err != nil {
			log.Fatal().Err(err).Msgf("models.UpdateMetrics() error")
		}

		go models.KeepMetricsUpdated(ctx, a.gormDB)
	}

	if config.Config.Bool("slack.enable") {
		log.Info().Msgf("BOOT | initializing Slack socket...")
		if err = slack.Init(ctx); err != nil {
			log.Fatal().Err(err).Msgf("slack.Init() error")
		}
		go slack.Start(ctx)
	}

	if config.Config.Bool("github.enable") {
		log.Info().Msgf("BOOT | initializing GitHub client...")
		if err = github.Init(ctx); err != nil {
			log.Fatal().Err(err).Msgf("github.Init() error")
		}
	}

	if config.Config.Bool("bitsDataWarehouse.enable") {
		log.Info().Msgf("BOOT | initializing BITS Data Warehouse connection...")
		if err = bits_data_warehouse.Init(ctx); err != nil {
			// Don't fail the boot if we can't connect
			log.Warn().Err(err).Msgf("bits_data_warehouse.Init() error")
		}
	}

	if config.Config.Bool("rolePropagation.enable") {
		log.Info().Msgf("BOOT | initializing role propagators...")
		if err = role_propagation.Init(ctx); err != nil {
			log.Fatal().Err(err).Msgf("role_propagation.Init() error")
		}
		go role_propagation.KeepPropagatingStale(ctx, a.gormDB)
	}

	if config.Config.Bool("suitabilitySynchronization.enable") {
		log.Info().Msgf("BOOT | initializing suitability synchronization...")
		if config.Config.MustString("mode") != "debug" && config.Config.Bool("suitabilitySynchronization.behaviors.loadIntoDB.enable") {
			log.Info().Msgf("BOOT | loading suitability from external stores...")
			if err = suitability_synchronization.LoadIntoDB(ctx, a.gormDB); err != nil {
				log.Fatal().Err(err).Msgf("suitability_synchronization.LoadIntoDB() error")
			}
			go suitability_synchronization.KeepLoadingIntoDB(ctx, a.gormDB)
		}
		if config.Config.Bool("suitabilitySynchronization.behaviors.suspendRoleAssignments.enable") {
			log.Info().Msgf("BOOT | beginning automatic role assignment suspension...")
			go suitability_synchronization.KeepSuspendingRoleAssignments(ctx, a.gormDB)
		}
	}

	if config.Config.Bool("oidc.enable") {
		log.Info().Msgf("BOOT | initializing OIDC provider...")
		if err = oidc_models.Init(ctx, a.gormDB); err != nil {
			log.Fatal().Err(err).Msgf("oidc_models.Init() error")
		}
		go oidc_models.KeepSigningKeysRotated(ctx, a.gormDB)
		go oidc_models.KeepExpiringRefreshTokens(ctx, a.gormDB)
	}

	go models.KeepAutoAssigningRoles(ctx, a.gormDB)

	log.Info().Msgf("BOOT | building Gin router...")
	gin.SetMode(gin.ReleaseMode) // gin.DebugMode can help resolve routing issues
	a.server = &http.Server{
		Addr:    ":8080",
		Handler: BuildRouter(ctx, a.gormDB),
	}

	log.Info().Msgf("BOOT | boot complete; now serving...")
	if err = a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal().Err(err).Msgf("server.ListenAndServe() error")
	}
}

func (a *Application) Stop() {
	if a.server != nil {
		log.Info().Msgf("BOOT | shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := a.server.Shutdown(ctx); err != nil {
			log.Warn().Err(err).Msgf("BOOT | server shutdown error")
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

	if a.gormDB != nil {
		log.Info().Msgf("BOOT | closing database connections...")
		if wasUnlocked := a.dbMigrationLock.TryLock(); !wasUnlocked {
			log.Info().Msgf("BOOT | detected database migration underway, attempting to wait until it completes...")
			a.dbMigrationLock.Lock()
			log.Info().Msgf("BOOT | database migration complete, proceeding with connection close...")
		}
		if sqlDB, err := a.gormDB.DB(); err != nil {
			log.Warn().Err(err).Msgf("BOOT | database error obtaining *sql.DB to shut down")
		} else if err = sqlDB.Close(); err != nil {
			log.Warn().Err(err).Msgf("BOOT | database connection close error")
		}
		a.dbMigrationLock.Unlock()
	} else {
		log.Info().Msgf("BOOT | no SQL database reference, skipping closing database connections")
	}

	if a.dbDriverCleanup != nil {
		log.Info().Msgf("BOOT | cleaning up database driver...")
		if err := a.dbDriverCleanup(); err != nil {
			log.Warn().Err(err).Msgf("BOOT | database driver clean up error")
		}
	} else {
		log.Info().Msgf("BOOT | no database driver cleanup function, skipping cleanup")
	}

	if a.livenessServer != nil {
		log.Info().Msgf("BOOT | stopping liveness endpoint...")
		a.livenessServer.Stop()
	} else {
		log.Info().Msgf("BOOT | no liveness server reference, skipping stopping liveness endpoint")
	}

	log.Info().Msgf("BOOT | exiting...")
}
