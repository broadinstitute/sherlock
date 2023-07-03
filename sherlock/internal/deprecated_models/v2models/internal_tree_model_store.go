package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"gorm.io/gorm"
	"strconv"
)

type internalTreeModelStore[M TreeModel] struct {
	*internalModelStore[M]
}

func (s internalTreeModelStore[M]) getChildrenPathToParent(db *gorm.DB, originChild *M, destinationParentID *uint) (path []*M, connected bool, err error) {
	if originChild == nil || destinationParentID == nil {
		return nil, false, nil
	} else if (*originChild).getID() == *destinationParentID {
		return []*M{}, true, nil
	}
	var childrenPath []*M
	currentChild := originChild
search:
	for currentChild != nil && *destinationParentID != (*currentChild).getID() {
		for _, traversedChild := range childrenPath {
			if (*traversedChild).getID() == (*currentChild).getID() {
				break search
			}
		}
		childrenPath = append(childrenPath, currentChild)
		if (*currentChild).getParentID() != nil {
			potentialNewHeadQuery, err := s.selectorToQueryModel(db, strconv.FormatUint(uint64(*(*currentChild).getParentID()), 10))
			if err != nil {
				return nil, false, fmt.Errorf("(%s) failed to build query for next %T tree node: %v", errors.InternalServerError, originChild, err)
			}
			potentialNewHead, err := s.getIfExists(db, potentialNewHeadQuery)
			if err != nil {
				return nil, false, fmt.Errorf("(%s) failed to query next %T tree node: %v", errors.InternalServerError, originChild, err)
			} else if potentialNewHead != nil {
				currentChild = potentialNewHead
			} else {
				break search
			}
		} else {
			break search
		}
	}
	if len(childrenPath) > 0 &&
		childrenPath[len(childrenPath)-1] != nil &&
		(*childrenPath[len(childrenPath)-1]).getParentID() != nil &&
		*(*childrenPath[len(childrenPath)-1]).getParentID() == *destinationParentID {
		return childrenPath, true, nil
	} else {
		return []*M{originChild}, false, nil
	}
}
