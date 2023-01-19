package v2models

import (
	"gorm.io/gorm"
)

type StoreSet struct {
	db *gorm.DB

	ClusterStore         *ModelStore[Cluster]
	EnvironmentStore     *ModelStore[Environment]
	ChartStore           *ModelStore[Chart]
	ChartVersionStore    *ModelStore[ChartVersion]
	AppVersionStore      *ModelStore[AppVersion]
	ChartReleaseStore    *ModelStore[ChartRelease]
	PagerdutyIntegration *ModelStore[PagerdutyIntegration]

	ChangesetEventStore *ChangesetEventStore
}

func NewStoreSet(db *gorm.DB) *StoreSet {
	return &StoreSet{
		db: db,

		ClusterStore:         &ModelStore[Cluster]{db: db, internalModelStore: clusterStore},
		EnvironmentStore:     &ModelStore[Environment]{db: db, internalModelStore: environmentStore},
		ChartStore:           &ModelStore[Chart]{db: db, internalModelStore: chartStore},
		ChartVersionStore:    &ModelStore[ChartVersion]{db: db, internalModelStore: chartVersionStore},
		AppVersionStore:      &ModelStore[AppVersion]{db: db, internalModelStore: appVersionStore},
		ChartReleaseStore:    &ModelStore[ChartRelease]{db: db, internalModelStore: chartReleaseStore},
		PagerdutyIntegration: &ModelStore[PagerdutyIntegration]{db: db, internalModelStore: pagerdutyIntegrationStore},

		ChangesetEventStore: &ChangesetEventStore{
			ModelStore:                  &ModelStore[Changeset]{db: db, internalModelStore: changesetStore.internalModelStore},
			internalChangesetEventStore: changesetStore,
		},
	}
}
