package sherlock

import "github.com/gin-gonic/gin"

func ConfigureRoutes(apiRouter *gin.RouterGroup) {
	ciIdentifiersV3 := apiRouter.Group("ci-identifiers/v3")
	{
		ciIdentifiersV3.GET("", ciIdentifiersV3List)
		ciIdentifiersV3.GET("*selector", ciIdentifiersV3Get)
	}
	ciRunsV3 := apiRouter.Group("ci-runs/v3")
	{
		ciRunsV3.GET("", ciRunsV3List)
		ciRunsV3.GET("*selector", ciRunsV3Get)
		ciRunsV3.PUT("", ciRunsV3Upsert)
	}
	usersV3 := apiRouter.Group("users/v3")
	{
		usersV3.GET("", usersV3List)
		usersV3.GET("*selector", usersV3Get)
		usersV3.PUT("", usersV3Upsert)
	}
}
