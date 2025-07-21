package sherlock

import (
	"time"

	"gorm.io/gorm"
)

type CommonFields struct {
	ID        uint      `json:"id" form:"id"`
	CreatedAt time.Time `json:"createdAt" form:"createdAt" format:"date-time"`
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt" format:"date-time"`
}

func (f CommonFields) toGormModel() gorm.Model {
	return gorm.Model{
		ID:        f.ID,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
	}
}

func commonFieldsFromGormModel(model gorm.Model) CommonFields {
	return CommonFields{
		ID:        model.ID,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
