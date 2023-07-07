package sherlock

import (
	"gorm.io/gorm"
	"time"
)

type commonFields struct {
	ID        uint      `json:"id" form:"id"`
	CreatedAt time.Time `json:"createdAt" form:"-" format:"date-time"`
	UpdatedAt time.Time `json:"updatedAt" form:"-" format:"date-time"`
}

func (f commonFields) toGormModel() gorm.Model {
	return gorm.Model{
		ID:        f.ID,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
	}
}

func commonFieldsFromGormModel(model gorm.Model) commonFields {
	return commonFields{
		ID:        model.ID,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
