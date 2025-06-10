package sherlock

import (
	"fmt"

	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"gorm.io/gorm"
)

type ServiceAlertV3 struct {
	CommonFields
	UUID *string `json:"uuid,omitempty" form:"uuid"`
	ServiceAlertV3Create
}

type ServiceAlertV3Create struct {
	OnEnvironment *string `json:"onEnvironment,omitempty" form:"onEnvironment"`
	ServiceAlertV3EditableFields
}

type ServiceAlertV3EditableFields struct {
	Title        *string `json:"title" form:"title"`
	AlertMessage *string `json:"message" form:"message"`
	Link         *string `json:"link" form:"link"`
	Severity     *string `json:"severity" form:"severity"`
}

func (i ServiceAlertV3) toModel(db *gorm.DB) models.ServiceAlert {
	ret := models.ServiceAlert{
		Model:        i.toGormModel(),
		Title:        i.Title,
		AlertMessage: i.AlertMessage,
		Link:         i.Link,
		Severity:     i.Severity,
	}
	if i.OnEnvironment != nil {
		environmentQuery, err := environmentModelFromSelector(*i.OnEnvironment)
		if err != nil {
			fmt.Errorf("error parsing environment selector '%s': %w", *i.OnEnvironment, err)
		}
		var result models.Environment
		if err = db.Where(&environmentQuery).Select("id").First(&result).Error; err != nil {
			fmt.Errorf("error fetching environment '%s': %w", *i.OnEnvironment, err)
		}
		ret.OnEnvironmentID = &result.ID
	}

	return ret
}

func ServiceAlertFromModel(model models.ServiceAlert) ServiceAlertV3 {
	var onEnvironment *string
	if model.OnEnvironment != nil && model.OnEnvironment.Name != "" {
		onEnvironment = &model.OnEnvironment.Name
	} else if model.OnEnvironmentID != nil {
		s := utils.UintToString(*model.OnEnvironmentID)
		onEnvironment = &s
	}
	return ServiceAlertV3{
		CommonFields: commonFieldsFromGormModel(model.Model),
		// UUID:          model.UUID, TODO
		ServiceAlertV3Create: ServiceAlertV3Create{
			OnEnvironment: onEnvironment,
			ServiceAlertV3EditableFields: ServiceAlertV3EditableFields{
				Title:        model.Title,
				AlertMessage: model.AlertMessage,
				Link:         model.Link,
				Severity:     model.Severity,
			},
		},
	}
}
