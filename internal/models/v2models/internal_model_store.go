package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/errors"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

type internalModelStore[M Model] struct {
	// Required:

	// selectorToQueryModel is meant to turn an arbitrary selector (possibly user-provided) into a database query.
	// The function is given a reference to the current stores so it can resolve indirect selectors, like those of
	// ChartRelease.
	// This function should not resolve its own output using the stores reference, the caller will resolve it against
	// the database directly.
	// This function doesn't need to particularly worry about sanitizing user input--struct fields are safe as-is for
	// database queries with Gorm.
	selectorToQueryModel func(db *gorm.DB, selector string) (M, error)
	// modelToSelectors is a "debugging" and validation function. It should generate as many selectors as possible from
	// its input. This is exposed to users so they can get aliases for a given selector for an existing model entry,
	// but it is also used by Create to do a uniqueness check across all selectors for a given input.
	modelToSelectors func(model *M) []string

	// Optional:

	// modelRequiresSuitability lets a particular type flag that the given model (which will always come from the
	// database) requires suitability for any mutations. If no function is provided, the model type is assumed
	// to have no suitability restrictions. To support hopping from association to association and to fail-safe,
	// the db reference should be used to load associations if they're used.
	modelRequiresSuitability func(db *gorm.DB, model *M) bool
	// validateModel lets a type enforce restrictions upon data entry. Associated data will not be present but foreign
	// keys themselves can be checked. There's no need to validate the grammar of selectors, that can be checked
	// automatically.
	// Note that this function has no mechanism to query the database to check the value of a foreign key--this is by
	// design, as setting foreign keys is done by the controller and a non-zero value will be a valid one. This function
	// should only worry about the presence of a foreign key, if an association is required.
	validateModel func(model *M) error
	// postCreate lets a type run perform additional actions once the model has been created but before the database
	// transaction finishes. Errors returned by this function will roll back the entire transaction.
	postCreate func(db *gorm.DB, created *M, user *auth.User) error
	// preCreate is similar to postCreate but it runs before even validation does--before the model has entered the
	// database at all.
	preCreate func(db *gorm.DB, toCreate *M, user *auth.User) error
	// rejectDuplicateCreate lets a type provide custom handling for when a new entry has selectors that match an
	// entry that's already in the database. Typically, this is always considered an error. If this function is
	// provided and does not error, the database will not be changed and the already-stored entry will be returned.
	// This means that duplicate create calls would all return successfully, while still maintaining selector-uniqueness
	// inside the database.
	rejectDuplicate func(existing *M, new *M) error
}

func (s internalModelStore[M]) create(db *gorm.DB, model M, user *auth.User) (M, bool, error) {
	if s.preCreate != nil {
		if err := s.preCreate(db, &model, user); err != nil {
			return model, false, fmt.Errorf("pre-create error: %v", err)
		}
	}
	if s.validateModel != nil {
		if err := s.validateModel(&model); err != nil {
			return model, false, fmt.Errorf("creation validation error: (%s) new %T: %v", errors.BadRequest, model, err)
		}
	}
	selectorsThatShouldNotCurrentlyExist := s.modelToSelectors(&model)
	log.Debug().Msgf("about to add new %T, checking that %d selectors don't already exist: %+v", model, len(selectorsThatShouldNotCurrentlyExist), selectorsThatShouldNotCurrentlyExist)
	for _, selector := range selectorsThatShouldNotCurrentlyExist {
		queryThatShouldNotMatch, err := s.selectorToQueryModel(db, selector)
		if err != nil {
			return model, false, fmt.Errorf("creation validation error: new %T would have an invalid selector %s: %v", model, selector, err)
		}
		var shouldStayEmpty []M
		result := db.Where(&queryThatShouldNotMatch).Find(&shouldStayEmpty)
		if result.Error != nil {
			return model, false, fmt.Errorf("(%s) unexpected creation error: new %T's selector %s couldn't be uniqueness-checked against the database due to an error: %v", errors.InternalServerError, model, selector, result.Error)
		} else if result.RowsAffected > 0 {
			// There's entries in the database already; if there's custom handling run that instead of just erroring
			if s.rejectDuplicate != nil {
				for _, existingMatch := range shouldStayEmpty {
					if err = s.rejectDuplicate(&existingMatch, &model); err != nil {
						return existingMatch, false, fmt.Errorf("creation validation error: (%s) new %T's selector %s matches an entry already in the database and there was an error resolving the duplicates: %v", errors.Conflict, model, selector, err)
					}
				}
				log.Debug().Msgf("won't add new %T, selector %s already exists but rejectDuplicateCreate didn't error so accepting and returning the first accepting match", model, selector)
				return shouldStayEmpty[0], false, nil
			} else {
				return shouldStayEmpty[0], false, fmt.Errorf("creation validation error: (%s) new %T's selector %s already matches an entry in the database", errors.Conflict, model, selector)
			}
		}
	}
	var ret M
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&model).Error; err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				// post-MVP TODO: We could pretty easily add APIs to expose soft-deletion controls to users; right now they'd have to have us go into the DB to un-delete stuff.
				return fmt.Errorf("creation error: (%s) new %T violated a database uniqueness constraint (are you recreating something with the same name? Contact DevOps) original error: %v", errors.BadRequest, model, err)
			}
			return fmt.Errorf("creation error: new %T couldn't be created in the database due to an error: %v", model, err)
		}
		result, err := s.get(tx, model)
		if err != nil {
			return fmt.Errorf("(%s) unexpected creation error: mid-transaction validation on %T failed: %v", errors.InternalServerError, model, err)
		}
		// Use db instead of tx here because tx is dirty and user-modified. Determination of suitability should
		// never be recursive, so this is safer.
		// (Why do we check suitability here rather than before adding at all? Adding and then querying lets us load
		// associations, and while suitability isn't recursive, it can be associative. Ex: a chart release's suitability
		// is defined in terms of the environment and cluster's suitability)
		if s.modelRequiresSuitability != nil && s.modelRequiresSuitability(db, &result) {
			if err = user.SuitableOrError(); err != nil {
				return fmt.Errorf("creation error: (%s) suitability is required to create this %T: %v", errors.Forbidden, model, err)
			}
		}
		if s.postCreate != nil {
			if err = s.postCreate(tx, &result, user); err != nil {
				return fmt.Errorf("post-create error: the %T itself was valid but an error occured running post-creation actions so creation was rolled back: %v", model, err)
			}
		}
		ret = result
		return nil
	})
	return ret, err == nil, err
}

// The signature here is really loose so we don't need to go down to the raw db's Where method. The signature exposed
// outside this package by ModelStore's ListAllMatching is more restrictive.
func (s internalModelStore[M]) listAllMatching(db *gorm.DB, limit int, query interface{}, args ...interface{}) ([]M, error) {
	var modelRef M
	var matching []M
	tx := db.Model(&modelRef).Where(query, args...).Preload(clause.Associations).Order("updated_at desc")
	if limit > 0 {
		tx = tx.Limit(limit)
	}
	if err := tx.Find(&matching).Error; err != nil {
		return matching, fmt.Errorf("(%s) unexpected list-all-matching error: %v", errors.InternalServerError, err)
	}
	return matching, nil
}

func (s internalModelStore[M]) get(db *gorm.DB, query M) (M, error) {
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

func (s internalModelStore[M]) edit(db *gorm.DB, query M, editsToMake M, user *auth.User, updateAllFields bool) (M, error) {
	toEdit, err := s.get(db, query)
	if err != nil {
		return toEdit, err
	}
	if s.modelRequiresSuitability != nil && s.modelRequiresSuitability(db, &toEdit) {
		if err = user.SuitableOrError(); err != nil {
			return toEdit, fmt.Errorf("edit error: (%s) suitability is required to edit %T: %v", errors.Forbidden, toEdit, err)
		}
	}
	var ret M
	err = db.Transaction(func(tx *gorm.DB) error {
		var chain = tx.Model(&toEdit)
		if updateAllFields {
			chain.Select("*")
		}
		if err = chain.Updates(&editsToMake).Error; err != nil {
			return fmt.Errorf("edit error editing %T: %v", toEdit, err)
		}
		result, err := s.get(tx, toEdit)
		if err != nil {
			return fmt.Errorf("(%s) unexpected edit error: mid-transaction validation on %T failed: %v", errors.InternalServerError, toEdit, err)
		}
		if s.validateModel != nil {
			if err := s.validateModel(&result); err != nil {
				return fmt.Errorf("edit validation error: (%s) resulting %T: %v", errors.BadRequest, result, err)
			}
		}
		// We check suitability *again* to prevent a user from editing an entry in a way that makes it require
		// suitability in the future, if they aren't themselves suitable.
		if s.modelRequiresSuitability != nil && s.modelRequiresSuitability(db, &result) {
			if err = user.SuitableOrError(); err != nil {
				return fmt.Errorf("edit error: (%s) suitability is required to edit %T in this way: %v", errors.Forbidden, toEdit, err)
			}
		}
		ret = result
		return nil
	})
	return ret, err
}

func (s internalModelStore[M]) delete(db *gorm.DB, query M, user *auth.User) (M, error) {
	toDelete, err := s.get(db, query)
	if err != nil {
		return toDelete, err
	}
	if s.modelRequiresSuitability != nil && s.modelRequiresSuitability(db, &toDelete) {
		if err = user.SuitableOrError(); err != nil {
			return toDelete, fmt.Errorf("delete error: (%s) suitability is required to delete %T: %v", errors.Forbidden, toDelete, err)
		}
	}
	if err = db.Delete(&toDelete).Error; err != nil {
		return toDelete, fmt.Errorf("delete error deleting %T: %v", toDelete, err)
	}
	return toDelete, nil
}
