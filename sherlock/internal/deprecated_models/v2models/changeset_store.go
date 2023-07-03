package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/auth/auth_models"
	"gorm.io/gorm"
)

type ChangesetStore struct {
	*ModelStore[Changeset]
}

func (s *ChangesetStore) PlanAndApply(changesets []Changeset, user *auth_models.User) ([]Changeset, error) {
	var ret []Changeset
	err := s.db.Transaction(func(tx *gorm.DB) error {
		planned, err := changesetStore.plan(tx, changesets, user)
		if err != nil {
			return err
		}
		ret, err = changesetStore.apply(tx, planned, user)
		return err
	})
	return ret, err
}

func (s *ChangesetStore) Plan(changesets []Changeset, user *auth_models.User) ([]Changeset, error) {
	return changesetStore.plan(s.db, changesets, user)
}

func (s *ChangesetStore) Apply(selectors []string, user *auth_models.User) ([]Changeset, error) {
	var queries []Changeset
	for index, selector := range selectors {
		query, err := changesetStore.selectorToQueryModel(s.db, selector)
		if err != nil {
			return []Changeset{}, fmt.Errorf("pre-apply error parsing selector %d '%s': %v", index+1, selector, err)
		}
		queries = append(queries, query)
	}
	return changesetStore.apply(s.db, queries, user)
}

func (s *ChangesetStore) QueryApplied(chartReleaseSelector string, offset int, limit int) ([]Changeset, error) {
	chartReleaseQuery, err := chartReleaseStore.selectorToQueryModel(s.db, chartReleaseSelector)
	if err != nil {
		return nil, err
	}
	chartRelease, err := chartReleaseStore.get(s.db, chartReleaseQuery)
	if err != nil {
		return nil, err
	}
	return changesetStore.queryApplied(s.db, chartRelease.ID, offset, limit)
}
