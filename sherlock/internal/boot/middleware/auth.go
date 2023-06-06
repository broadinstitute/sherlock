package middleware

import (
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"github.com/gin-gonic/gin"
)

func Auth(userStore *v2models.MiddlewareUserStore) gin.HandlerFunc {
	if config.Config.String("mode") == "debug" {
		return auth.FakeUserMiddleware(userStore)
	} else {
		return auth.IapUserMiddleware(userStore)
	}
}
