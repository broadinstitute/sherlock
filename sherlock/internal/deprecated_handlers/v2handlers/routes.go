package v2handlers

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_controllers/v2controllers"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ConfigureRoutes(v2apiRouter *gin.RouterGroup, db *gorm.DB) {
	v2ControllerSet := v2controllers.NewControllerSet(v2models.NewStoreSet(db))
	registerClusterHandlers(v2apiRouter, v2ControllerSet.ClusterController)
	registerEnvironmentHandlers(v2apiRouter, v2ControllerSet.EnvironmentController)
	registerChartHandlers(v2apiRouter, v2ControllerSet.ChartController)
	registerChartVersionHandlers(v2apiRouter, v2ControllerSet.ChartVersionController)
	registerAppVersionHandlers(v2apiRouter, v2ControllerSet.AppVersionController)
	registerChartReleaseHandlers(v2apiRouter, v2ControllerSet.ChartReleaseController)
	registerChangesetHandlers(v2apiRouter, v2ControllerSet.ChangesetController)
	registerPagerdutyIntegrationHandlers(v2apiRouter, v2ControllerSet.PagerdutyIntegrationController)
	registerDatabaseInstanceHandlers(v2apiRouter, v2ControllerSet.DatabaseInstanceController)
	registerUserHandlers(v2apiRouter, v2ControllerSet.UserController)
	registerCiIdentifierHandlers(v2apiRouter, v2ControllerSet.CiIdentifierController)
	registerCiRunHandlers(v2apiRouter, v2ControllerSet.CiRunController)
}
