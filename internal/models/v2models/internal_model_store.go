package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth/auth_models"
	"github.com/broadinstitute/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/internal/models/model_actions"
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

	// errorIfForbidden controls whether a user may perform a certain action on the model instance in question.
	// If not provided, it is assumed any Sherlock user may perform any action on the model isntance.
	// The db reference should be used to fully load any associations that are used.
	errorIfForbidden func(db *gorm.DB, model *M, action model_actions.ActionType, user *auth_models.User) error
	// validateModel lets a type enforce restrictions upon data entry. Associated data will not be present but foreign
	// keys themselves can be checked. There's no need to validate the grammar of selectors, that can be checked
	// automatically.
	// Note that this function has no mechanism to query the database to check the value of a foreign key--this is by
	// design, as setting foreign keys is done by the controller and a non-zero value will be a valid one. This function
	// should only worry about the presence of a foreign key, if an association is required.
	validateModel func(model *M) error
	// postCreate lets a type run perform additional actions once the model has been created but before the database
	// transaction finishes. Errors returned by this function will roll back the entire transaction.
	postCreate func(db *gorm.DB, created *M, user *auth_models.User) error
	// preCreate is similar to postCreate but it runs before even validation does--before the model has entered the
	// database at all.
	preCreate func(db *gorm.DB, toCreate *M, user *auth_models.User) error
	// preDeletePostValidate runs after validation right before deletion. Since it runs after validation, it is important
	// that this function not change toDelete in a way that would require re-validation.
	preDeletePostValidate func(db *gorm.DB, toDelete *M, user *auth_models.User) error
	// handleIncomingDuplicate lets a type determine if an error should actually be thrown when a duplicate is detected
	// at creation-time. Normally, this is always an error, but if this function is defined and doesn't error, the
	// database will be unaltered and the existing entry will be returned to the user as the result.
	handleIncomingDuplicate func(existing *M, new *M) error
	// editsMayChangeSelectors lets a type declare that its selectors are impacted by mutable fields, so edits should
	// have the selector-uniqueness property enforced.
	editsMayChangeSelectors bool
}

func (s internalModelStore[M]) wrappedErrorIfForbidden(db *gorm.DB, model *M, action model_actions.ActionType, user *auth_models.User) error {
	if s.errorIfForbidden == nil {
		return nil
	} else if err := s.errorIfForbidden(db, model, action, user); err != nil {
		return fmt.Errorf("%s permissions error for %T (%s): %v", model_actions.ActionTypeToString(action), *model, errors.Forbidden, err)
	} else {
		return nil
	}
}

func (s internalModelStore[M]) requireSameModel(existing *M, new *M) error {
	if existing != nil && new != nil && (*existing).getID() == (*new).getID() {
		return nil
	} else {
		return fmt.Errorf("mismatch")
	}
}

func (s internalModelStore[M]) enforceSelectorUniqueness(db *gorm.DB, model *M, handleConflicts func(existing *M, new *M) error) (acceptedDuplicate *M, err error) {
	selectors := s.modelToSelectors(model)
	log.Debug().Msgf("enforcing %T selector uniqueness across %d selectors (%+v); handleConflicts provided=%t", model, len(selectors), selectors, handleConflicts != nil)
	for _, selector := range s.modelToSelectors(model) {
		query, err := s.selectorToQueryModel(db, selector)
		if err != nil {
			return nil, fmt.Errorf("selector validation error: resulting model has invalid selector '%s': %v", selector, err)
		}
		var results []M
		if err := db.Where(&query).Find(&results).Error; err != nil {
			return nil, fmt.Errorf("(%s) unexpected selector validation error: failed to query possible selector conflicts: %v", errors.InternalServerError, err)
		} else {
			for _, result := range results {
				if handleConflicts == nil { // if we can't handle conflicts
					return nil, fmt.Errorf("(%s) selector conflict: new %T selector '%s' already matches an entry in the database (ID %d)", errors.Conflict, result, selector, result.getID())
				} else if err := handleConflicts(&result, model); err != nil { // if handling a conflict still errors
					return nil, fmt.Errorf("(%s) selector conflict: new %T selector '%s' already matches an entry in the database (ID %d): conflict handler reported %v", errors.Conflict, result, selector, result.getID(), err)
				} else if acceptedDuplicate == nil { // if we don't have a duplicate recorded, fine
					acceptedDuplicate = &result
				} else if (*acceptedDuplicate).getID() != (result).getID() { // if we do have a duplicate it's different, still error
					// I'm not sure it's possible to hit this case, but maybe if handleConflicts was changed in-flight then duplicates could "appear" in the database
					return nil, fmt.Errorf("(%s) selector conflict: new %T matched multiple duplicates in the database (at least IDs %d and %d)", errors.Conflict, result, (*acceptedDuplicate).getID(), (result).getID())
				}
			}
		}
	}
	return acceptedDuplicate, nil
}

func (s internalModelStore[M]) create(db *gorm.DB, model M, user *auth_models.User) (M, bool, error) {
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
	if allowedDuplicate, err := s.enforceSelectorUniqueness(db, &model, s.handleIncomingDuplicate); err != nil {
		return model, false, fmt.Errorf("create validation error: %v", err)
	} else if allowedDuplicate != nil {
		return *allowedDuplicate, false, nil
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
		// Use db instead of tx here because tx is dirty and user-modified. Determination of permissions should
		// never be recursive, so this is safer.
		// (Why do we check permissions here rather than before adding at all? Adding and then querying lets us load
		// associations, and while permissions aren't recursive, they can be associative. Ex: the ability to affect
		// a chart release is dependent on the chart release's environment and cluster)
		if err = s.wrappedErrorIfForbidden(db, &result, model_actions.CREATE, user); err != nil {
			return err
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
// outside this package by ModelStore's ListAllMatchingByUpdated is more restrictive.
func (s internalModelStore[M]) listAllMatchingByUpdated(db *gorm.DB, limit int, query interface{}, args ...interface{}) ([]M, error) {
	return s.listAllMatchingOrdered(db, limit, "updated_at desc", query, args...)
}

func (s internalModelStore[M]) listAllMatchingByCreated(db *gorm.DB, limit int, query interface{}, args ...interface{}) ([]M, error) {
	return s.listAllMatchingOrdered(db, limit, "created_at desc", query, args...)
}

func (s internalModelStore[M]) listAllMatchingOrdered(db *gorm.DB, limit int, order string, query interface{}, args ...interface{}) ([]M, error) {
	var modelRef M
	var matching []M
	tx := db.Model(&modelRef).Where(query, args...).Preload(clause.Associations).Order(order)
	if limit > 0 {
		tx = tx.Limit(limit)
	}
	if err := tx.Find(&matching).Error; err != nil {
		return matching, fmt.Errorf("(%s) unexpected list-all-matching error: %v", errors.InternalServerError, err)
	}
	return matching, nil
}

func (s internalModelStore[M]) getIfExists(db *gorm.DB, query M) (*M, error) {
	var matching []M
	tx := db.Where(&query).Preload(clause.Associations)
	if err := tx.Find(&matching).Error; err != nil {
		return nil, fmt.Errorf("(%s) unexpected query error: failed to run query %T %+v against the database: %v", errors.InternalServerError, query, query, err)
	} else if len(matching) > 1 {
		return nil, fmt.Errorf("query result error: (%s) more than one entry (%d total) matched non-zero values of %T %+v", errors.BadRequest, len(matching), query, query)
	} else if len(matching) == 1 {
		return &matching[0], nil
	} else {
		return nil, nil
	}
}

func (s internalModelStore[M]) get(db *gorm.DB, query M) (M, error) {
	var zeroValue M
	if result, err := s.getIfExists(db, query); err != nil {
		return zeroValue, err
	} else if result == nil {
		return zeroValue, fmt.Errorf("no result for query (%s)", errors.NotFound)
	} else {
		return *result, nil
	}
}

func (s internalModelStore[M]) edit(db *gorm.DB, query M, editsToMake M, user *auth_models.User, updateAllFields bool) (M, error) {
	toEdit, err := s.get(db, query)
	if err != nil {
		return toEdit, err
	}
	if err = s.wrappedErrorIfForbidden(db, &toEdit, model_actions.EDIT, user); err != nil {
		return toEdit, err
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
		// We check permissions *again* to prevent a user from editing an entry in a way that makes it require
		// permissions above theirs in the future.
		if err = s.wrappedErrorIfForbidden(db, &result, model_actions.EDIT, user); err != nil {
			return err
		}
		if s.editsMayChangeSelectors {
			if _, err := s.enforceSelectorUniqueness(tx, &result, s.requireSameModel); err != nil {
				return fmt.Errorf("edit validation error: %v", err)
			}
		}
		ret = result
		return nil
	})
	return ret, err
}

func (s internalModelStore[M]) deleteIfExists(db *gorm.DB, query M, user *auth_models.User) (*M, error) {
	if toDelete, err := s.getIfExists(db, query); err != nil || toDelete == nil {
		return toDelete, err
	} else {
		if err = s.wrappedErrorIfForbidden(db, toDelete, model_actions.DELETE, user); err != nil {
			return toDelete, err
		}
		err = db.Transaction(func(tx *gorm.DB) error {
			if s.preDeletePostValidate != nil {
				if err := s.preDeletePostValidate(tx, toDelete, user); err != nil {
					return fmt.Errorf("pre-delete post-validate error: %v", err)
				}
			}
			return db.Delete(&toDelete).Error
		})
		if err != nil {
			return toDelete, fmt.Errorf("delete error deleting %T: %v", toDelete, err)
		}
		return toDelete, nil
	}
}

func (s internalModelStore[M]) delete(db *gorm.DB, query M, user *auth_models.User) (M, error) {
	var zeroValue M
	if result, err := s.deleteIfExists(db, query, user); err != nil {
		return zeroValue, err
	} else if result == nil {
		return zeroValue, fmt.Errorf("delete error: no result for query (%s)", errors.NotFound)
	} else {
		return *result, nil
	}
}
