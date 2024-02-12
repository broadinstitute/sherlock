package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
)

type CiIdentifierV3 struct {
	CommonFields
	CiRuns       []CiRunV3 `json:"ciRuns,omitempty" form:"-"`
	ResourceType string    `json:"resourceType" form:"resourceType"`
	ResourceID   uint      `json:"resourceID" form:"resourceID"`

	// Available only when querying a CiIdentifier via a CiRun, indicates the status of the run for that resource
	ResourceStatus *string `json:"resourceStatus,omitempty" form:"resourceStatus"`
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
		CommonFields:   commonFieldsFromGormModel(model.Model),
		CiRuns:         ciRuns,
		ResourceType:   model.ResourceType,
		ResourceID:     model.ResourceID,
		ResourceStatus: model.ResourceStatus,
	}
}
