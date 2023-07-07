package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"gorm.io/gorm"
)

type Model interface {
	TableName() string
	getID() uint
}

type ModelStore[M Model] struct {
	db       *gorm.DB
	internal *internalModelStore[M]
}

func (s ModelStore[M]) Create(model M, user *models.User) (M, bool, error) {
	return s.internal.Create(s.db, model, user)
}

func (s ModelStore[M]) ListAllMatchingByUpdated(filter M, limit int) ([]M, error) {
	return s.internal.ListAllMatchingByUpdated(s.db, limit, &filter)
}

func (s ModelStore[M]) ListAllMatchingByCreated(filter M, limit int) ([]M, error) {
	return s.internal.ListAllMatchingByCreated(s.db, limit, &filter)
}

func (s ModelStore[M]) Get(selector string) (M, error) {
	return s.internal.GetBySelector(s.db, selector)
}

func (s ModelStore[M]) Edit(selector string, editsToMake M, user *models.User) (M, error) {
	return s.internal.EditBySelector(s.db, selector, editsToMake, user)
}

func (s ModelStore[M]) Delete(selector string, user *models.User) (M, error) {
	return s.internal.DeleteBySelector(s.db, selector, user)
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
	return s.internal.modelToSelectors(&query), nil
}
