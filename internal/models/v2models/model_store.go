package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/errors"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

type Model interface {
	TableName() string
}

type Store[M Model] struct {
	db *gorm.DB
	// selectorToQueryModel is meant to turn an arbitrary selector (possibly user-provided) into a database query.
	// The function is given a reference to the current stores so it can resolve indirect selectors, like those of
	// ChartRelease.
	// This function should not resolve its own output using the stores reference, the caller will resolve it against
	// the database directly.
	// This function doesn't need to particularly worry about sanitizing user input--struct fields are safe as-is for
	// database queries with Gorm.
	selectorToQueryModel func(db *gorm.DB, selector string) (M, error)
	// modelToSelectors is a "debugging" and validation function. It should generate as many selectors possible from
	// its input. This is exposed to users so they can get aliases for a given selector for an existing model entry,
	// but it is also used by Create to do a uniqueness check across all selectors for a given input.
	modelToSelectors func(model M) []string
	// modelRequiresSuitability lets a particular type flag that the given model (which will always come from the
	// database) requires suitability for any mutations. If no function is provided, the model type is assumed
	// to have no suitability restrictions.
	modelRequiresSuitability func(model M) bool
	// validateModel lets a type enforce restrictions upon data entry. This function should not expect associated data
	// to be present (it will sometimes be called ahead of when such data can be automatically supplied by the
	// database). If no function is provided, the model type is assumed to have no validation beyond selectors,
	// suitability, and any association-setting done by a controller.
	validateModel func(model M) error
}

func (s Store[M]) Create(model M, userSuitable bool) (M, error) {
	if s.validateModel != nil {
		if err := s.validateModel(model); err != nil {
			return model, fmt.Errorf("creation validation error: (%s) new %T: %v", errors.BadRequest, model, err)
		}
	}
	selectorsThatShouldNotCurrentlyExist := s.modelToSelectors(model)
	log.Debug().Msgf("about to add new %T, checking that %d selectors don't already exist: %+v", model, len(selectorsThatShouldNotCurrentlyExist), selectorsThatShouldNotCurrentlyExist)
	for _, selector := range selectorsThatShouldNotCurrentlyExist {
		queryThatShouldNotMatch, err := s.selectorToQueryModel(s.db, selector)
		if err != nil {
			return model, fmt.Errorf("creation validation error: new %T would have an invalid selector %s: %v", model, selector, err)
		}
		var shouldStayEmpty M
		result := s.db.Where(&queryThatShouldNotMatch).Limit(1).Find(&shouldStayEmpty)
		if result.Error != nil {
			return model, fmt.Errorf("(%s) unexpected creation error: new %T's selector %s couldn't be uniqueness-checked against the database due to an error: %v", errors.InternalServerError, model, selector, result.Error)
		} else if result.RowsAffected > 0 {
			log.Debug().Msgf("can't add new %T, selector %s already exists", model, selector)
			return shouldStayEmpty, fmt.Errorf("creation validation error: (%s) new %T's selector %s already matches an entry in the database", errors.Conflict, model, selector)
		}
	}
	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&model).Error; err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				// post-MVP TODO: We could pretty easily add APIs to expose soft-deletion controls to users; right now they'd have to have us go into the DB to un-delete stuff.
				return fmt.Errorf("creation error: (%s) new %T violated a database uniqueness constraint (are you recreating something with the same name? Contact DevOps) original error: %v", errors.BadRequest, model, err)
			}
			return fmt.Errorf("creation error: new %T couldn't be created in the database due to an error: %v", model, err)
		}
		if s.modelRequiresSuitability != nil {
			result, err := getFromQuery(tx, model)
			if err != nil {
				return fmt.Errorf("(%s) unexpected creation error: mid-transaction validation on %T failed: %v", errors.InternalServerError, model, err)
			}
			if s.modelRequiresSuitability(result) && !userSuitable {
				return fmt.Errorf("create error: (%s) user is not suitable but suitability is required to create this %T", errors.Forbidden, model)
			}
		}
		return nil
	})
	return model, err
}

func (s Store[M]) ListAllMatching(filter M, limit int) ([]M, error) {
	var matching []M
	tx := s.db.Where(&filter).Preload(clause.Associations).Order("updated_at desc")
	if limit > 0 {
		tx = tx.Limit(limit)
	}
	if err := tx.Find(&matching).Error; err != nil {
		return matching, fmt.Errorf("(%s) unexpected list-all-matching error: %v", errors.InternalServerError, err)
	}
	return matching, nil
}

func getFromQuery[M Model](db *gorm.DB, query M) (M, error) {
	var matching []M
	tx := db.Where(&query).Preload(clause.Associations)
	if err := tx.Find(&matching).Error; err != nil {
		return query, fmt.Errorf("(%s) unexpected query error: failed to run query %T %+v against the database: %v", errors.InternalServerError, query, query, err)
	} else if len(matching) == 0 {
		return query, fmt.Errorf("query result error: (%s) no entry matched non-zero values of %T %+v", errors.NotFound, query, query)
	} else if len(matching) > 1 {
		return query, fmt.Errorf("query result error: (%s) more than one entry (%d total) matched non-zero values of %T %+v", errors.BadRequest, len(matching), query, query)
	}
	return matching[0], nil
}

func (s Store[M]) Get(selector string) (M, error) {
	query, err := s.selectorToQueryModel(s.db, selector)
	if err != nil {
		return query, fmt.Errorf("query error parsing %T selector %s: %v", query, selector, err)
	}
	ret, err := getFromQuery(s.db, query)
	if err != nil {
		return query, fmt.Errorf("query error using %T selector %s: %v", query, selector, err)
	}
	return ret, nil
}

func (s Store[M]) Edit(selector string, editsToMake M, userSuitable bool) (M, error) {
	toEdit, err := s.Get(selector)
	if err != nil {
		return toEdit, fmt.Errorf("edit error handling %T selector %s: %v", toEdit, selector, err)
	}
	if s.modelRequiresSuitability != nil && s.modelRequiresSuitability(toEdit) && !userSuitable {
		return toEdit, fmt.Errorf("edit error: (%s) user is not suitable but suitability is required to edit %T %s", errors.Forbidden, toEdit, selector)
	}
	err = s.db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Model(&toEdit).Updates(&editsToMake).Error; err != nil {
			return fmt.Errorf("edit error editing %T matched by selector %s: %v", toEdit, selector, err)
		}
		if s.modelRequiresSuitability != nil || s.validateModel != nil {
			result, err := getFromQuery(tx, toEdit)
			if err != nil {
				return fmt.Errorf("(%s) unexpected edit error: mid-transaction validation on %T failed: %v", errors.InternalServerError, toEdit, err)
			}
			if s.validateModel != nil {
				if err := s.validateModel(result); err != nil {
					return fmt.Errorf("edit validation error: (%s) resulting %T: %v", errors.BadRequest, result, err)
				}
			}
			// We check suitability *again* to prevent a user from editing an entry in a way that makes it require
			// suitability in the future, if they aren't themselves suitable.
			if s.modelRequiresSuitability != nil && s.modelRequiresSuitability(result) && !userSuitable {
				return fmt.Errorf("edit error: (%s) user is not suitable but suitability is required to edit the %T in this way", errors.Forbidden, toEdit)
			}
		}
		return nil
	})
	return toEdit, err
}

func (s Store[M]) Delete(selector string, userSuitable bool) (M, error) {
	toDelete, err := s.Get(selector)
	if err != nil {
		return toDelete, fmt.Errorf("delete error handling %T selector %s: %v", toDelete, selector, err)
	}
	if s.modelRequiresSuitability != nil && s.modelRequiresSuitability(toDelete) && !userSuitable {
		return toDelete, fmt.Errorf("delete error: (%s) user is not suitable but suitability is required to delete %T %s", errors.Forbidden, toDelete, selector)
	}
	if err = s.db.Delete(&toDelete).Error; err != nil {
		return toDelete, fmt.Errorf("delete error deleting %T matched by selector %s: %v", toDelete, selector, err)
	}
	return toDelete, nil
}

// GetOtherValidSelectors is basically just a human debug method. Different model types have different selectors to try
// to make it easier to refer to them than by having to directly query them and use their numeric ID primary key. Under
// the hood, models are already required to be able to generate selectors from an entry for uniqueness-validation
// purposes, so this is a simple method that uses that existing code to translate one selector for an existing entry
// into all possible selectors that would match.
func (s Store[M]) GetOtherValidSelectors(selector string) ([]string, error) {
	model, err := s.Get(selector)
	if err != nil {
		return []string{}, fmt.Errorf("error listing other selectors for %T %s: %v", model, selector, err)
	}
	return s.modelToSelectors(model), nil
}
