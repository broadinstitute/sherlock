package sherlock

import "github.com/broadinstitute/sherlock/sherlock/internal/models"

type PagerdutyIntegrationV3 struct {
	CommonFields
	PagerdutyID string  `json:"pagerdutyID" form:"pagerdutyID"`
	Name        *string `json:"name" form:"name"`
	Type        *string `json:"type" form:"type"`
}

type PagerdutyIntegrationV3Create struct {
	PagerdutyID string `json:"pagerdutyID" form:"pagerdutyID"`
	PagerdutyIntegrationV3Edit
}

type PagerdutyIntegrationV3Edit struct {
	Name *string `json:"name" form:"name"`
	Key  *string `json:"key" form:"key"`
	Type *string `json:"type" form:"type"`
}

func (p PagerdutyIntegrationV3) toModel() models.PagerdutyIntegration {
	return models.PagerdutyIntegration{
		Model:       p.toGormModel(),
		PagerdutyID: p.PagerdutyID,
		Name:        p.Name,
		Type:        p.Type,
	}
}

func (p PagerdutyIntegrationV3Create) toModel() models.PagerdutyIntegration {
	return models.PagerdutyIntegration{
		PagerdutyID: p.PagerdutyID,
		Name:        p.Name,
		Key:         p.Key,
		Type:        p.Type,
	}
}

func (p PagerdutyIntegrationV3Edit) toModel() models.PagerdutyIntegration {
	return PagerdutyIntegrationV3Create{PagerdutyIntegrationV3Edit: p}.toModel()
}

func pagerdutyIntegrationFromModel(model models.PagerdutyIntegration) PagerdutyIntegrationV3 {
	return PagerdutyIntegrationV3{
		CommonFields: commonFieldsFromGormModel(model.Model),
		PagerdutyID:  model.PagerdutyID,
		Name:         model.Name,
		Type:         model.Type,
	}
}
