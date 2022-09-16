package v2models

import (
	"gorm.io/gorm"
)

type StoreSet struct {
	db *gorm.DB

	ClusterStore      *ModelStore[Cluster]
	EnvironmentStore  *ModelStore[Environment]
	ChartStore        *ModelStore[Chart]
	ChartVersionStore *ModelStore[ChartVersion]
	AppVersionStore   *ModelStore[AppVersion]
	ChartReleaseStore *ModelStore[ChartRelease]

	ChangesetEventStore *ChangesetEventStore
}

func NewStoreSet(db *gorm.DB) *StoreSet {
	return &StoreSet{
		db: db,

		ClusterStore:      &ModelStore[Cluster]{db: db, internalModelStore: clusterStore},
		EnvironmentStore:  &ModelStore[Environment]{db: db, internalModelStore: environmentStore},
		ChartStore:        &ModelStore[Chart]{db: db, internalModelStore: chartStore},
		ChartVersionStore: &ModelStore[ChartVersion]{db: db, internalModelStore: chartVersionStore},
		AppVersionStore:   &ModelStore[AppVersion]{db: db, internalModelStore: appVersionStore},
		ChartReleaseStore: &ModelStore[ChartRelease]{db: db, internalModelStore: chartReleaseStore},

		ChangesetEventStore: &ChangesetEventStore{
			ModelStore:                  &ModelStore[Changeset]{db: db, internalModelStore: changesetStore.internalModelStore},
			internalChangesetEventStore: changesetStore,
		},
	}
}

// WithRollbackStoreSet is a last-resort for achieving reasonable error behavior outside the model. Functions inside
// the model will reasonably rollback when they encounter an error. This function is intended to be used when a
// controller composes model functions together and needs to rollback the entire composition when an error occurs in
// any of them.
func WithRollbackStoreSet[T any](storeSet *StoreSet, f func(*StoreSet) (T, error)) (T, error) {
	var ret T
	err := storeSet.db.Transaction(func(tx *gorm.DB) error {
		var err error
		ret, err = f(NewStoreSet(tx))
		return err
	})
	return ret, err
}
