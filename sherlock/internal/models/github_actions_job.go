package models

import (
	"gorm.io/gorm"
	"time"
)

type GithubActionsJob struct {
	gorm.Model
	GithubActionsOwner         string
	GithubActionsRepo          string
	GithubActionsRunID         uint
	GithubActionsAttemptNumber uint
	GithubActionsJobID         uint

	// Mutable
	JobCreatedAt  *time.Time
	JobStartedAt  *time.Time
	JobTerminalAt *time.Time
	Status        *string
}
