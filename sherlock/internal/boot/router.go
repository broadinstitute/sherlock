package boot

import (
	"context"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"strings"

	"github.com/broadinstitute/sherlock/go-shared/pkg/version"
	"github.com/broadinstitute/sherlock/sherlock/docs"
	"github.com/broadinstitute/sherlock/sherlock/html"
	"github.com/broadinstitute/sherlock/sherlock/internal/api/login"
	"github.com/broadinstitute/sherlock/sherlock/internal/api/misc"
	"github.com/broadinstitute/sherlock/sherlock/internal/api/sherlock"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/metrics"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/cors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/csrf_protection"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/headers"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/logger"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/security"
	"github.com/broadinstitute/sherlock/sherlock/internal/oidc_models"
	"github.com/gin-gonic/gin"
	swaggo_files "github.com/swaggo/files/v2"
	swaggo_gin "github.com/swaggo/gin-swagger"
	"golang.org/x/net/webdav"
	"gorm.io/gorm"
)

//	@title			Sherlock
//	@description	The Data Science Platform's source-of-truth service.
//	@description	Note: this API will try to load and return associations in responses, so clients won't need to make as many requests. This behavior isn't recursive, though, so associations of associations are *not* fully loaded (even if it might seem that way from looking at the data types).
//	@version		development
//	@schemes		https
//	@accept			json
//	@produce		json
//	@host			sherlock.dsp-devops-prod.broadinstitute.org
//	@BasePath		/

//	@contact.name	DSP DevOps
//	@contact.email	dsp-devops@broadinstitute.org

//	@license.name	BSD-3-Clause
//	@license.url	https://github.com/broadinstitute/sherlock/blob/main/LICENSE.txt

func BuildRouter(ctx context.Context, db *gorm.DB) *gin.Engine {
	// primaryHost may be unset or empty, which means this could be an empty string. That's potentially
	// okay, because the Swagger page itself will still work. In production, though, having this set
	// is important to provide a fully valid config that other tools like security scanners can use.
	// See default_config.yaml's primaryHost for more information.
	docs.SwaggerInfo.Host = config.Config.String("primaryHost")
	docs.SwaggerInfo.Version = version.BuildVersion
	if config.Config.String("mode") == "debug" {
		// When running locally, make the Swagger page have a scheme dropdown with http as the default
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
	}

	router := gin.New()

	router.Use(
		gin.Recovery(),
		logger.Logger(),
		cors.Cors(),
		headers.Headers(),
		security.Security())

	resourceMiddleware := make(gin.HandlersChain, 0)
	resourceMiddleware = append(resourceMiddleware, csrf_protection.CsrfProtection())
	resourceMiddleware = append(resourceMiddleware, authentication.Middleware(db)...)

	// Replace Gin's standard fallback responses with our standard error format for friendlier client behavior
	router.NoRoute(func(ctx *gin.Context) {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) no handler for %s found", errors.NotFound, ctx.Request.URL.Path))
	})
	router.NoMethod(func(ctx *gin.Context) {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) method %s not allowed for %s", errors.MethodNotAllowed, ctx.Request.Method, ctx.Request.URL.Path))
	})

	// /status, /version
	misc.ConfigureRoutes(&router.RouterGroup)

	router.GET("/metrics", metrics.PrometheusHandler())

	router.StaticFS("/static", http.FS(html.StaticHtmlFiles))

	router.GET("/swagger/*any", swaggo_gin.WrapHandler(
		&webdav.Handler{
			FileSystem: &webdavFilesystem{http.FS(swaggo_files.FS)},
			LockSystem: webdav.NewMemLS(),
		}, func(c *swaggo_gin.Config) {
			c.Title = "Sherlock Swagger UI"
			c.URL = "/swagger/doc.json"
			c.DefaultModelsExpandDepth = 2
			c.DocExpansion = "none"
		}))
	router.GET("", func(ctx *gin.Context) { ctx.Redirect(http.StatusMovedPermanently, "/swagger/index.html") })

	if config.Config.Bool("oidc.enable") {
		// delegate /oidc/* to OIDC library, trimming the path prefix because of how the library expects to receive requests
		// https://broadinstitute.slack.com/archives/CQ6SL4N5T/p1721406732128199
		router.Any("/oidc/*any", func(ctx *gin.Context) {
			req := ctx.Request.Clone(ctx)
			req.RequestURI = strings.TrimPrefix(req.RequestURI, "/oidc")
			req.URL.Path = strings.TrimPrefix(req.URL.Path, "/oidc")
			oidc_models.Provider.ServeHTTP(ctx.Writer, req)
		})
		// authenticate /login handler to complete OIDC auth requests
		router.GET("/login", append(resourceMiddleware, login.LoginGet)...)
	}

	// routes under /api require authentication and may use the database
	apiRouter := router.Group("/api", resourceMiddleware...)

	// refactored sherlock API, under /api/{type}/v3
	sherlock.ConfigureRoutes(apiRouter)

	// special error for the removed "v2" API, under /api/v2/{type}
	apiRouter.Any("/v2/*path", func(ctx *gin.Context) {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) sherlock's v2 API has been removed; reach out to #dsp-devops-champions for help updating your client", errors.NotFound))
	})

	return router
}

// webdavFilesystem adapts a http.FileSystem to a webdav.FileSystem. This is unfortunately necessary because
// swaggo_gin.WrapHandler expects a webdav.Handler that needs a webdav.FileSystem (even though we don't need
// most of what webdav.FileSystem provides). There's actually a package online that does this but it's not
// popular and this is so small I judge it makes sense to just write what we actually need ourselves.
//
// For context -- why does swaggo_gin.WrapHandler expect a webdav.Handler if swaggo_files.FS is a http.FileSystem?
// Because we're using an updated version of swaggo_files to pull in a more recent non-vulnerable version of
// swagger-ui. swaggo_gin never got an update to match, so we're stuck with this weirdness.
type webdavFilesystem struct {
	http.FileSystem
}

func (w *webdavFilesystem) Mkdir(_ context.Context, _ string, _ os.FileMode) error {
	return os.ErrPermission
}

func (w *webdavFilesystem) RemoveAll(_ context.Context, _ string) error {
	return os.ErrPermission
}

func (w *webdavFilesystem) Rename(_ context.Context, _, _ string) error {
	return os.ErrPermission
}

func (w *webdavFilesystem) Stat(_ context.Context, name string) (os.FileInfo, error) {
	if f, err := w.FileSystem.Open(name); err != nil {
		return nil, err
	} else {
		defer func(f fs.File) {
			_ = f.Close()
		}(f)
		return f.Stat()
	}
}

func (w *webdavFilesystem) OpenFile(_ context.Context, name string, flag int, _ os.FileMode) (webdav.File, error) {
	if flag != os.O_RDONLY {
		return nil, os.ErrPermission
	} else if f, err := w.FileSystem.Open(name); err != nil {
		return nil, err
	} else {
		return &webdavFile{f}, nil
	}
}

type webdavFile struct {
	http.File
}

func (w *webdavFile) Write(_ []byte) (n int, err error) {
	return 0, &os.PathError{Op: "write", Err: os.ErrPermission}
}
