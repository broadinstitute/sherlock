package sherlock

import (
	"fmt"

	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ServiceAlertV3 struct {
	CommonFields
	Uuid      *string `json:"uuid,omitempty" form:"uuid"`
	CreatedBy *string `json:"createdBy,omitempty" form:"createdBy"`
	UpdatedBy *string `json:"updatedBy,omitempty" form:"updatedBy"`
	DeletedBy *string `json:"deletedBy,omitempty" form:"deltedBy"`
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
	Severity     *string `json:"severity" form:"severity" enums:"blocker, critical, minor"`
}

func doUserQuery(id string, db *gorm.DB) (uint, error) {
	userQuery, err := userModelFromSelector(id)
	if err != nil {
		return 0, err
	}
	var userResult models.User
	if err = db.Where(&userQuery).First(&userResult).Error; err != nil {
		return 0, err
	} else {
		return userResult.ID, nil
	}
}

// TO And FROM, do conversion from / to string from/to UUID
func (i ServiceAlertV3) toModel(db *gorm.DB) (models.ServiceAlert, error) {
	ret := models.ServiceAlert{
		Model:        i.toGormModel(),
		Title:        i.Title,
		AlertMessage: i.AlertMessage,
		Link:         i.Link,
		Severity:     i.Severity,
		CreatedBy:    i.CreatedBy,
		UpdatedBy:    i.UpdatedBy,
		DeletedBy:    i.DeletedBy,
	}

	if i.Uuid != nil {
		svc_alert_uuid, err := uuid.Parse(*i.Uuid)
		if err != nil {
			return models.ServiceAlert{}, fmt.Errorf("error parsing service alert UUID '%s': %w", *i.Uuid, err)
		}
		ret.Uuid = &svc_alert_uuid
	}

	if i.OnEnvironment != nil {
		environmentQuery, err := environmentModelFromSelector(*i.OnEnvironment)
		if err != nil {
			return models.ServiceAlert{}, fmt.Errorf("error parsing environment selector '%s': %w", *i.OnEnvironment, err)
		}
		var result models.Environment
		if err = db.Where(&environmentQuery).Select("id").First(&result).Error; err != nil {
			return models.ServiceAlert{}, fmt.Errorf("error fetching environment '%s': %w", *i.OnEnvironment, err)
		}
		ret.OnEnvironmentID = &result.ID
	}

	return ret, nil
}

func ServiceAlertFromModel(model models.ServiceAlert) ServiceAlertV3 {

	var onEnvironment *string
	if model.OnEnvironment != nil && model.OnEnvironment.Name != "" {
		onEnvironment = &model.OnEnvironment.Name
	} else if model.OnEnvironmentID != nil {
		s := utils.UintToString(*model.OnEnvironmentID)
		onEnvironment = &s
	}
	var alertUuidString *string
	if model.Uuid != nil {
		s := uuid.UUID.String(*model.Uuid)
		alertUuidString = &s
	}

	return ServiceAlertV3{
		CommonFields: commonFieldsFromGormModel(model.Model),
		Uuid:         alertUuidString,
		CreatedBy:    model.CreatedBy,
		UpdatedBy:    model.UpdatedBy,
		DeletedBy:    model.DeletedBy,
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
