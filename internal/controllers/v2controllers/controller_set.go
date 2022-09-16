package v2controllers

import "github.com/broadinstitute/sherlock/internal/models/v2models"

type ControllerSet struct {
	ClusterController      *ClusterController
	EnvironmentController  *EnvironmentController
	ChartController        *ChartController
	ChartVersionController *ChartVersionController
	AppVersionController   *AppVersionController
	ChartReleaseController *ChartReleaseController
	ChangesetController    *ChangesetController
}

func NewControllerSet(stores *v2models.StoreSet) *ControllerSet {
	return &ControllerSet{
		ClusterController:      newClusterController(stores),
		EnvironmentController:  newEnvironmentController(stores),
		ChartController:        newChartController(stores),
		ChartVersionController: newChartVersionController(stores),
		AppVersionController:   newAppVersionController(stores),
		ChartReleaseController: newChartReleaseController(stores),
		ChangesetController:    newChangesetController(stores),
	}
}
