package v2models

import "gorm.io/gorm"

type StoreSet struct {
	ClusterStore           Store[Cluster]
	EnvironmentStore       Store[Environment]
	ChartStore             Store[Chart]
	ChartVersionStore      Store[ChartVersion]
	AppVersionStore        Store[AppVersion]
	ChartReleaseStore      Store[ChartRelease]
	ChartDeployRecordStore Store[ChartDeployRecord]
}

func NewStoreSet(db *gorm.DB) StoreSet {
	// TODO this is Jack's hack to not write SQL while iterating on the database
	err := db.AutoMigrate(&Cluster{}, &Environment{}, &Chart{}, &ChartVersion{}, &AppVersion{}, &ChartRelease{}, &ChartDeployRecord{})
	if err != nil {
		panic(err)
	}
	return StoreSet{
		ClusterStore:           newClusterStore(db),
		EnvironmentStore:       newEnvironmentStore(db),
		ChartStore:             newChartStore(db),
		ChartVersionStore:      newChartVersionStore(db),
		AppVersionStore:        newAppVersionStore(db),
		ChartReleaseStore:      newChartReleaseStore(db),
		ChartDeployRecordStore: newChartDeployRecordStore(db),
	}
}
