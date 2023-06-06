package v2controllers

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/models/v2models"
	"gorm.io/gorm"
)

type CiIdentifier struct {
	ReadableBaseType
	CiRuns       []CiRun `json:"ciRuns,omitempty" form:"-"`
	ResourceType string  `json:"resourceType" form:"resourceType"`
	ResourceID   uint    `json:"resourceID" form:"resourceID"`
}

type CreatableCiIdentifier struct {
	ResourceType string `json:"resourceType" form:"resourceType"`
	ResourceID   uint   `json:"resourceID" form:"resourceID"`
	EditableCiIdentifier
}

type EditableCiIdentifier struct {
	CiRuns []string `json:"ciRuns" form:"-"` // Always appends; will eliminate duplicates
}

//nolint:unused
func (c CiIdentifier) toModel(_ *v2models.StoreSet) (v2models.CiIdentifier, error) {
	return v2models.CiIdentifier{
		Model: gorm.Model{
			ID:        c.ID,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		},
		ResourceType: c.ResourceType,
		ResourceID:   c.ResourceID,
	}, nil
}

//nolint:unused
func (c CreatableCiIdentifier) toModel(storeSet *v2models.StoreSet) (v2models.CiIdentifier, error) {
	ciRuns := make([]*v2models.CiRun, len(c.EditableCiIdentifier.CiRuns))
append:
	for _, ciRunSelector := range c.CiRuns {
		ciRun, err := storeSet.CiRunStore.Get(ciRunSelector)
		if err != nil {
			return v2models.CiIdentifier{}, err
		}
		// If we already have this one in the list, don't add it again
		for _, existingCiRun := range ciRuns {
			if existingCiRun.ID == ciRun.ID {
				continue append
			}
		}
		ciRuns = append(ciRuns, &ciRun)
	}
	return v2models.CiIdentifier{
		CiRuns:       ciRuns,
		ResourceType: c.ResourceType,
		ResourceID:   c.ResourceID,
	}, nil
}

//nolint:unused
func (c EditableCiIdentifier) toModel(storeSet *v2models.StoreSet) (v2models.CiIdentifier, error) {
	// We don't need to do anything special to handle the append behavior of the associations, the model will do that
	// for us
	return CreatableCiIdentifier{EditableCiIdentifier: c}.toModel(storeSet)
}

type CiIdentifierController = ModelController[v2models.CiIdentifier, CiIdentifier, CreatableCiIdentifier, EditableCiIdentifier]

func newCiIdentifierController(stores *v2models.StoreSet) *CiIdentifierController {
	return &CiIdentifierController{
		primaryStore:    stores.CiIdentifierStore,
		allStores:       stores,
		modelToReadable: modelCiIdentifierToCiIdentifier,
	}
}

func modelCiIdentifierToCiIdentifier(model *v2models.CiIdentifier) *CiIdentifier {
	if model == nil {
		return nil
	}

	var ciRuns []CiRun
	for _, modelCiRun := range model.CiRuns {
		ciRun := modelCiRunToCiRun(modelCiRun)
		if ciRun != nil {
			ciRuns = append(ciRuns, *ciRun)
		}
	}

	return &CiIdentifier{
		ReadableBaseType: ReadableBaseType{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
		CiRuns:       ciRuns,
		ResourceType: model.ResourceType,
		ResourceID:   model.ResourceID,
	}
}
