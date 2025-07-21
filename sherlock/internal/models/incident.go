package models

import (
	"time"

	"gorm.io/gorm"
)

type Incident struct {
	gorm.Model
	Ticket            *string
	Description       *string
	StartedAt         *time.Time
	RemediatedAt      *time.Time
	ReviewCompletedAt *time.Time
}
