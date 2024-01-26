package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"time"
)

type IncidentV3 struct {
	CommonFields
	IncidentV3Create
}

type IncidentV3Create struct {
	IncidentV3Edit
}

type IncidentV3Edit struct {
	Ticket            *string    `json:"ticket" form:"ticket"`
	Description       *string    `json:"description" form:"description"`
	StartedAt         *time.Time `json:"startedAt" form:"startedAt"`
	RemediatedAt      *time.Time `json:"remediatedAt" form:"remediatedAt"`
	ReviewCompletedAt *time.Time `json:"reviewCompletedAt" form:"reviewCompletedAt"`
}

func (i IncidentV3) toModel() models.Incident {
	return models.Incident{
		Model:             i.toGormModel(),
		Ticket:            i.Ticket,
		Description:       i.Description,
		StartedAt:         i.StartedAt,
		RemediatedAt:      i.RemediatedAt,
		ReviewCompletedAt: i.ReviewCompletedAt,
	}
}

func (i IncidentV3Create) toModel() models.Incident {
	return IncidentV3{IncidentV3Create: i}.toModel()
}

func (i IncidentV3Edit) toModel() models.Incident {
	return IncidentV3Create{IncidentV3Edit: i}.toModel()
}

func incidentFromModel(model models.Incident) IncidentV3 {
	return IncidentV3{
		CommonFields: commonFieldsFromGormModel(model.Model),
		IncidentV3Create: IncidentV3Create{
			IncidentV3Edit: IncidentV3Edit{
				Ticket:            model.Ticket,
				Description:       model.Description,
				StartedAt:         model.StartedAt,
				RemediatedAt:      model.RemediatedAt,
				ReviewCompletedAt: model.ReviewCompletedAt,
			},
		},
	}
}
