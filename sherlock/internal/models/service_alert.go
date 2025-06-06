package models

import (
	"gorm.io/gorm"
)

type ServiceAlert struct {
	gorm.Model
	Title       *string
	Message     *string
	Link        *string
	Severity    *string
	UUID        *string
	Environment *Environment
}
