package v2controllers

import (
	"github.com/broadinstitute/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"gorm.io/gorm"
)

type PagerdutyIntegration struct {
	ReadableBaseType
	PagerdutyID string  `json:"pagerdutyID" form:"pagerdutyID"`
	Name        *string `json:"name" form:"name"`
	Type        *string `json:"type" form:"type"`
}

type CreatablePagerdutyIntegration struct {
	PagerdutyID string `json:"pagerdutyID" form:"pagerdutyID"`
	EditablePagerdutyIntegration
}

type EditablePagerdutyIntegration struct {
	Name *string `json:"name" form:"name"`
	Key  *string `json:"key" form:"key"`
	Type *string `json:"type" form:"type"`
}

func (p PagerdutyIntegration) toModel(_ *v2models.StoreSet) (v2models.PagerdutyIntegration, error) {
	return v2models.PagerdutyIntegration{
		Model: gorm.Model{
			ID:        p.ID,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		},
		Name:        p.Name,
		PagerdutyID: p.PagerdutyID,
		Type:        p.Type,
	}, nil
}

func (p CreatablePagerdutyIntegration) toModel(_ *v2models.StoreSet) (v2models.PagerdutyIntegration, error) {
	return v2models.PagerdutyIntegration{
		Name:        p.Name,
		PagerdutyID: p.PagerdutyID,
		Key:         p.Key,
		Type:        p.Type,
	}, nil
}

func (p EditablePagerdutyIntegration) toModel(storeSet *v2models.StoreSet) (v2models.PagerdutyIntegration, error) {
	return CreatablePagerdutyIntegration{EditablePagerdutyIntegration: p}.toModel(storeSet)
}

type PagerdutyIntegrationController = ModelController[v2models.PagerdutyIntegration, PagerdutyIntegration, CreatablePagerdutyIntegration, EditablePagerdutyIntegration]

func newPagerdutyIntegrationController(stores *v2models.StoreSet) *PagerdutyIntegrationController {
	return &PagerdutyIntegrationController{
		primaryStore:                   stores.PagerdutyIntegration,
		allStores:                      stores,
		modelToReadable:                modelPagerdutyIntegrationToPagerdutyIntegration,
		extractPagerdutyIntegrationKey: extractPagerdutyIntegrationKeyFromPagerdutyIntegration,
		beehiveUrlFormatString:         config.Config.MustString("beehive.pagerdutyIntegrationUrlFormatString"),
	}
}

func modelPagerdutyIntegrationToPagerdutyIntegration(model *v2models.PagerdutyIntegration) *PagerdutyIntegration {
	if model == nil {
		return nil
	}

	return &PagerdutyIntegration{
		ReadableBaseType: ReadableBaseType{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
		Name:        model.Name,
		PagerdutyID: model.PagerdutyID,
		Type:        model.Type,
	}
}

func extractPagerdutyIntegrationKeyFromPagerdutyIntegration(model *v2models.PagerdutyIntegration) *string {
	if model != nil {
		return model.Key
	}
	return nil
}
