package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"gorm.io/gorm"
)

type ChangesetStore struct {
	*ModelStore[Changeset]
}

func (s *ChangesetStore) PlanAndApply(changesets []Changeset, user *models.User) ([]Changeset, error) {
	var ret []Changeset
	err := s.db.Transaction(func(tx *gorm.DB) error {
		planned, err := InternalChangesetStore.plan(tx, changesets, user)
		if err != nil {
			return err
		}
		ret, err = InternalChangesetStore.apply(tx, planned, user)
		return err
	})
	return ret, err
}

func (s *ChangesetStore) Plan(changesets []Changeset, user *models.User) ([]Changeset, error) {
	return InternalChangesetStore.plan(s.db, changesets, user)
}

func (s *ChangesetStore) Apply(selectors []string, user *models.User) ([]Changeset, error) {
	var queries []Changeset
	for index, selector := range selectors {
		query, err := InternalChangesetStore.selectorToQueryModel(s.db, selector)
		if err != nil {
			return []Changeset{}, fmt.Errorf("pre-apply error parsing selector %d '%s': %v", index+1, selector, err)
		}
		queries = append(queries, query)
	}
	return InternalChangesetStore.apply(s.db, queries, user)
}

func (s *ChangesetStore) QueryApplied(chartReleaseSelector string, offset int, limit int) ([]Changeset, error) {
	chartReleaseQuery, err := InternalChartReleaseStore.selectorToQueryModel(s.db, chartReleaseSelector)
	if err != nil {
		return nil, err
	}
	chartRelease, err := InternalChartReleaseStore.Get(s.db, chartReleaseQuery)
	if err != nil {
		return nil, err
	}
	return InternalChangesetStore.queryApplied(s.db, chartRelease.ID, offset, limit)
}
