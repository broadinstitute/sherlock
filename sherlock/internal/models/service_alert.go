package models

import (
	"gorm.io/gorm"
)

type ServiceAlert struct {
	gorm.Model
	Title           *string
	Message         *string
	Link            *string
	Severity        *string
	UUID            *uint
	Environment     *Environment
	OnEnvironmentID *uint
}
