package sherlock

import (
	"fmt"

	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"gorm.io/gorm"
)

type ServiceAlertV3 struct {
	CommonFields
	UUID *uint
	ServiceAlertV3EditableFields
}

type ServiceAlertV3EditableFields struct {
	Title         *string `json:"title" form:"title"`
	Message       *string `json:"message" form:"message"`
	Link          *string `json:"link" form:"link"`
	Severity      *string `json:"severtiy" form:"severity"`
	OnEnvironment *string `json:"onEnvironment,omitempty" form:"onEnvironment"`
}

func (i ServiceAlertV3) toModel(db *gorm.DB) models.ServiceAlert {
	if i.OnEnvironment != nil {
		environmentQuery, err := environmentModelFromSelector(*i.OnEnvironment)
		if err != nil {
			fmt.Errorf("error parsing environment selector '%s': %w", *i.OnEnvironment, err)
		}
		var result models.Environment
		if err = db.Where(&environmentQuery).Select("id").First(&result).Error; err != nil {
			fmt.Errorf("error fetching environment '%s': %w", *i.OnEnvironment, err)
		}
		i.UUID = &result.ID
	}

	return models.ServiceAlert{
		Model:    i.toGormModel(),
		Title:    i.Title,
		Message:  i.Message,
		Link:     i.Link,
		Severity: i.Severity,
		UUID:     i.UUID,
	}
}

func ServiceAlertFromModel(model models.ServiceAlert) ServiceAlertV3 {
	return ServiceAlertV3{
		CommonFields: commonFieldsFromGormModel(model.Model),
		UUID:         model.UUID,
		ServiceAlertV3EditableFields: ServiceAlertV3EditableFields{
			Title:    model.Title,
			Message:  model.Message,
			Link:     model.Link,
			Severity: model.Severity,
		},
	}
}
