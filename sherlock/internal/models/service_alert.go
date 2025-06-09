package models

import (
	"gorm.io/gorm"
)

type ServiceAlert struct {
	gorm.Model
	Title           *string
	AlertMessage    *string
	Link            *string
	Severity        *string
	OnEnvironmentID *uint
	OnEnvironment   *Environment
}
