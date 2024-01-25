package models

import (
	"gorm.io/gorm"
	"time"
)

type Incident struct {
	gorm.Model
	Ticket            *string
	Description       *string
	StartedAt         *time.Time
	RemediatedAt      *time.Time
	ReviewCompletedAt *time.Time
}
