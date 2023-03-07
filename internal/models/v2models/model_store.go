package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth/auth_models"
	"gorm.io/gorm"
)

type Model interface {
	TableName() string
}

type ModelStore[M Model] struct {
	db *gorm.DB
	*internalModelStore[M]
}

func (s ModelStore[M]) Create(model M, user *auth_models.User) (M, bool, error) {
	return s.create(s.db, model, user)
}

func (s ModelStore[M]) ListAllMatchingByUpdated(filter M, limit int) ([]M, error) {
	return s.listAllMatchingByUpdated(s.db, limit, &filter)
}

func (s ModelStore[M]) ListAllMatchingByCreated(filter M, limit int) ([]M, error) {
	return s.listAllMatchingByCreated(s.db, limit, &filter)
}

func (s ModelStore[M]) Get(selector string) (M, error) {
	query, err := s.selectorToQueryModel(s.db, selector)
	if err != nil {
		return query, fmt.Errorf("query error parsing %T selector '%s': %v", query, selector, err)
	}
	ret, err := s.get(s.db, query)
	if err != nil {
		return query, fmt.Errorf("query error using %T selector '%s': %v", query, selector, err)
	}
	return ret, nil
}

func (s ModelStore[M]) Edit(selector string, editsToMake M, user *auth_models.User) (M, error) {
	query, err := s.selectorToQueryModel(s.db, selector)
	if err != nil {
		return query, fmt.Errorf("query error parsing %T selector '%s': %v", query, selector, err)
	}
	ret, err := s.edit(s.db, query, editsToMake, user, false)
	if err != nil {
		return query, fmt.Errorf("edit error using %T selector '%s': %v", query, selector, err)
	}
	return ret, nil
}

func (s ModelStore[M]) Delete(selector string, user *auth_models.User) (M, error) {
	query, err := s.selectorToQueryModel(s.db, selector)
	if err != nil {
		return query, fmt.Errorf("query error parsing %T selector '%s': %v", query, selector, err)
	}
	ret, err := s.delete(s.db, query, user)
	if err != nil {
		return query, fmt.Errorf("delete error using %T selector '%s': %v", query, selector, err)
	}
	return ret, nil
}

// GetOtherValidSelectors is basically just a human debug method. Different model types have different selectors to try
// to make it easier to refer to them than by having to directly query them and use their numeric ID primary key. Under
// the hood, models are already required to be able to generate selectors from an entry for uniqueness-validation
// purposes, so this is a simple method that uses that existing code to translate one selector for an existing entry
// into all possible selectors that would match.
func (s ModelStore[M]) GetOtherValidSelectors(selector string) ([]string, error) {
	query, err := s.Get(selector)
	if err != nil {
		return []string{}, fmt.Errorf("query error parsing %T selector '%s': %v", query, selector, err)
	}
	return s.modelToSelectors(&query), nil
}
