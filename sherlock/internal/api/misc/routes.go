package misc

import "github.com/gin-gonic/gin"

func ConfigureRoutes(rootRouter *gin.RouterGroup) {
	rootRouter.GET("/status", statusGet)
	rootRouter.GET("/version", versionGet)
	rootRouter.GET("/connection-check", connectionCheckGet)
}
