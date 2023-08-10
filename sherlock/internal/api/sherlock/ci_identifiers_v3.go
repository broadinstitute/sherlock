package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
)

type CiIdentifierV3 struct {
	CommonFields
	CiRuns       []CiRunV3 `json:"ciRuns,omitempty" form:"-"`
	ResourceType string    `json:"resourceType" form:"resourceType"`
	ResourceID   uint      `json:"resourceID" form:"resourceID"`
}

func (c CiIdentifierV3) toModel() models.CiIdentifier {
	return models.CiIdentifier{
		Model:        c.toGormModel(),
		ResourceType: c.ResourceType,
		ResourceID:   c.ResourceID,
	}
}

func ciIdentifierFromModel(model models.CiIdentifier) CiIdentifierV3 {
	var ciRuns []CiRunV3
	if len(model.CiRuns) > 0 {
		ciRuns = make([]CiRunV3, len(model.CiRuns))
		for index, modelCiRun := range model.CiRuns {
			ciRuns[index] = ciRunFromModel(modelCiRun)
		}
	}
	return CiIdentifierV3{
		CommonFields: commonFieldsFromGormModel(model.Model),
		CiRuns:       ciRuns,
		ResourceType: model.ResourceType,
		ResourceID:   model.ResourceID,
	}
}

// This function is an absolute hack. There's some other places where we want to read the CiIdentifiers
// from un-refactored models (v2models.CiIdentifiable) as the refactored models.CiIdentifier type.
// To help avoid circular package dependencies, we put that function here. Once this package doesn't
// rely on un-refactored code we can summarily throw this out.
func ciIdentifierModelFromOldModel(source v2models.CiIdentifiable) models.CiIdentifier {
	oldIdentifier := source.GetCiIdentifier()
	return models.CiIdentifier{
		Model:        oldIdentifier.Model,
		ResourceType: oldIdentifier.ResourceType,
		ResourceID:   oldIdentifier.ResourceID,
	}
}
