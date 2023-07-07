package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/model_actions"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/utils"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

// CiIdentifiable is an interface to help align functionality across other types that can have a CiIdentifier.
type CiIdentifiable interface {
	// GetCiIdentifier should return either a type's loaded CiIdentifier or a generated one based on resource type and
	// resource ID. It's important for the caller to understand that this function can return things not in the
	// database yet.
	//
	// In reality, this function is meant to go hand-in-hand with createCiRunIdentifiersJustInTime. This function will
	// return CiIdentifier instances that may or may not exist, and that function will create them all before the
	// CiRun actually enters the database. If there's a breakdown in communication, it's actually fine for superfluous
	// creates to be attempted -- rejectDuplicateCiIdentifier will make it so that the creates will behave just like
	// gets.
	GetCiIdentifier() *CiIdentifier
}

type CiIdentifier struct {
	gorm.Model
	ResourceType string `gorm:"index:idx_v2_ci_identifiers_polymorphic_index,priority:1"`
	ResourceID   uint   `gorm:"index:idx_v2_ci_identifiers_polymorphic_index,priority:2"`
	// Mutable
	CiRuns []*CiRun `gorm:"many2many:v2_ci_runs_for_identifiers; constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
}

func (c CiIdentifier) TableName() string {
	return "v2_ci_identifiers"
}

func (c CiIdentifier) getID() uint {
	return c.ID
}

var InternalCiIdentifierStore *internalModelStore[CiIdentifier]

func init() {
	InternalCiIdentifierStore = &internalModelStore[CiIdentifier]{
		selectorToQueryModel:    ciIdentifierSelectorToQuery,
		modelToSelectors:        ciIdentifierToSelectors,
		errorIfForbidden:        ciIdentifierErrorIfForbidden,
		validateModel:           validateCiIdentifier,
		handleIncomingDuplicate: rejectDuplicateCiIdentifier,
		editsAppendManyToMany: map[string]func(edits *CiIdentifier) any{
			"CiRuns": func(edits *CiIdentifier) any { return edits.CiRuns },
		},
	}
}

func ciIdentifierSelectorToQuery(db *gorm.DB, selector string) (CiIdentifier, error) {
	if len(selector) == 0 {
		return CiIdentifier{}, fmt.Errorf("(%s) CI identifier selector cannot be empty", errors.BadRequest)
	}
	var query CiIdentifier
	if utils.IsNumeric(selector) {
		// ID
		id, err := strconv.Atoi(selector)
		if err != nil {
			return CiIdentifier{}, fmt.Errorf("(%s) string to int conversion error of '%s': %v", errors.BadRequest, selector, err)
		}
		query.ID = uint(id)
		return query, nil
	} else if strings.Count(selector, "/") > 0 {
		// resource type + type's selector ...
		parts := strings.Split(selector, "/")
		query.ResourceType = parts[0]
		resourceSelector := strings.Join(parts[1:], "/")
		var resolver SelectorResolver
		switch query.ResourceType {
		case "chart":
			resolver = InternalChartStore
		case "chart-version":
			resolver = InternalChartVersionStore
		case "app-version":
			resolver = InternalAppVersionStore
		case "cluster":
			resolver = InternalClusterStore
		case "environment":
			resolver = InternalEnvironmentStore
		case "chart-release":
			resolver = InternalChartReleaseStore
		case "changeset":
			resolver = InternalChangesetStore
		default:
			return CiIdentifier{}, fmt.Errorf("(%s) invalid CI identifier selector '%s', resource type sub-selector '%s' wasn't recognized", errors.BadRequest, selector, query.ResourceType)
		}
		id, err := resolver.ResolveSelector(db, resourceSelector)
		if err != nil {
			return CiIdentifier{}, fmt.Errorf("invalid CI identifier selector '%s', resource sub-selector '%s' had an error: %v", selector, resourceSelector, err)
		}
		query.ResourceID = id
		return query, nil
	}
	return CiIdentifier{}, fmt.Errorf("(%s) invalid CI identifier selector '%s'", errors.BadRequest, selector)
}

func ciIdentifierToSelectors(ciIdentifier *CiIdentifier) []string {
	var selectors []string
	if ciIdentifier != nil {
		if ciIdentifier.ID != 0 {
			selectors = append(selectors, strconv.FormatUint(uint64(ciIdentifier.ID), 10))
		}
		if ciIdentifier.ResourceType != "" && ciIdentifier.ResourceID != 0 {
			selectors = append(selectors, fmt.Sprintf("%s/%d", ciIdentifier.ResourceType, ciIdentifier.ResourceID))
		}
	}
	return selectors
}

func ciIdentifierErrorIfForbidden(_ *gorm.DB, ciIdentifier *CiIdentifier, actionType model_actions.ActionType, _ *models.User) error {
	if actionType == model_actions.DELETE {
		return fmt.Errorf("(%s) deleting a %T is not allowed", errors.Forbidden, ciIdentifier)
	} else {
		return nil
	}
}

func validateCiIdentifier(ciIdentifier *CiIdentifier) error {
	if ciIdentifier == nil {
		return fmt.Errorf("the model passed was nil")
	}
	if ciIdentifier.ResourceType == "" {
		return fmt.Errorf("a %T must have a resource type", ciIdentifier)
	}
	if ciIdentifier.ResourceID == 0 {
		return fmt.Errorf("a %T must have a resource ID", ciIdentifier)
	}
	return nil
}

func rejectDuplicateCiIdentifier(existing *CiIdentifier, new *CiIdentifier) error {
	if existing.ResourceType != new.ResourceType {
		return fmt.Errorf("%T resource type mismatch during upsert attempt, new was %s but existing was %s", new, new.ResourceType, existing.ResourceType)
	}
	if existing.ResourceID != new.ResourceID {
		return fmt.Errorf("%T resource ID mismatch during upsert attempt, new was %d but existing was %d", new, new.ResourceID, existing.ResourceID)
	}
	if new.CiRuns != nil && len(new.CiRuns) > 0 {
		// We care about this because if this function *doesn't* error, the new CiIdentifier actually gets thrown out
		// and the existing one is returned as what was created. If the new CiIdentifier has CiRuns that the existing
		// one doesn't, we'd be returning a 200 while having forgotten about those CiRuns. We could potentially build
		// fancy handling here to determine if the new CiRuns list is a subset of the existing one, but I can't think
		// of how a client would ever even make realistic use of that behavior. For now, we just say that the new
		// CiIdentifier specifying any CiRuns at all is potentially unsafe and we bail out.
		return fmt.Errorf("%d new CiRuns would be lost during upsert attempt on %T", len(new.CiRuns), new)
	}
	return nil
}
