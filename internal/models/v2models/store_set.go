package v2models

import (
	"gorm.io/gorm"
)

type StoreSet struct {
	db *gorm.DB

	ClusterStore          *ModelStore[Cluster]
	EnvironmentStore      *ModelStore[Environment]
	ChartStore            *ModelStore[Chart]
	ChartVersionStore     *TreeModelStore[ChartVersion]
	AppVersionStore       *TreeModelStore[AppVersion]
	ChartReleaseStore     *ModelStore[ChartRelease]
	PagerdutyIntegration  *ModelStore[PagerdutyIntegration]
	DatabaseInstanceStore *ModelStore[DatabaseInstance]
	UserStore             *ModelStore[User]
	ChangesetStore        *ChangesetStore
}

func NewStoreSet(db *gorm.DB) *StoreSet {
	return &StoreSet{
		db: db,

		ClusterStore:     &ModelStore[Cluster]{db: db, internalModelStore: clusterStore},
		EnvironmentStore: &ModelStore[Environment]{db: db, internalModelStore: environmentStore},
		ChartStore:       &ModelStore[Chart]{db: db, internalModelStore: chartStore},
		ChartVersionStore: &TreeModelStore[ChartVersion]{
			ModelStore:             &ModelStore[ChartVersion]{db: db, internalModelStore: chartVersionStore.internalModelStore},
			internalTreeModelStore: chartVersionStore,
		},
		AppVersionStore: &TreeModelStore[AppVersion]{
			ModelStore:             &ModelStore[AppVersion]{db: db, internalModelStore: appVersionStore.internalModelStore},
			internalTreeModelStore: appVersionStore,
		},
		ChartReleaseStore:     &ModelStore[ChartRelease]{db: db, internalModelStore: chartReleaseStore},
		PagerdutyIntegration:  &ModelStore[PagerdutyIntegration]{db: db, internalModelStore: pagerdutyIntegrationStore},
		DatabaseInstanceStore: &ModelStore[DatabaseInstance]{db: db, internalModelStore: databaseInstanceStore},
		UserStore:             &ModelStore[User]{db: db, internalModelStore: userStore},
		ChangesetStore: &ChangesetStore{
			ModelStore: &ModelStore[Changeset]{db: db, internalModelStore: changesetStore.internalModelStore},
		},
	}
}
