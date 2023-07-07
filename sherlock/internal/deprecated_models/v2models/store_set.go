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
	CiIdentifierStore     *ModelStore[CiIdentifier]
	CiRunStore            *ModelStore[CiRun]
}

func NewStoreSet(db *gorm.DB) *StoreSet {
	return &StoreSet{
		db: db,

		ClusterStore:     &ModelStore[Cluster]{db: db, internal: InternalClusterStore},
		EnvironmentStore: &ModelStore[Environment]{db: db, internal: InternalEnvironmentStore},
		ChartStore:       &ModelStore[Chart]{db: db, internal: InternalChartStore},
		ChartVersionStore: &TreeModelStore[ChartVersion]{
			ModelStore:             &ModelStore[ChartVersion]{db: db, internal: InternalChartVersionStore.internalModelStore},
			internalTreeModelStore: InternalChartVersionStore,
		},
		AppVersionStore: &TreeModelStore[AppVersion]{
			ModelStore:             &ModelStore[AppVersion]{db: db, internal: InternalAppVersionStore.internalModelStore},
			internalTreeModelStore: InternalAppVersionStore,
		},
		ChartReleaseStore:     &ModelStore[ChartRelease]{db: db, internal: InternalChartReleaseStore},
		PagerdutyIntegration:  &ModelStore[PagerdutyIntegration]{db: db, internal: InternalPagerdutyIntegrationStore},
		DatabaseInstanceStore: &ModelStore[DatabaseInstance]{db: db, internal: InternalDatabaseInstanceStore},
		UserStore:             &ModelStore[User]{db: db, internal: InternalUserStore},
		ChangesetStore: &ChangesetStore{
			ModelStore: &ModelStore[Changeset]{db: db, internal: InternalChangesetStore.internalModelStore},
		},
		CiIdentifierStore: &ModelStore[CiIdentifier]{db: db, internal: InternalCiIdentifierStore},
		CiRunStore:        &ModelStore[CiRun]{db: db, internal: InternalCiRunStore},
	}
}
