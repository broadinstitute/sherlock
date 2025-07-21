package oidc_models

import (
	"time"

	"github.com/google/uuid"
)

type AuthRequestCode struct {
	Code          string `gorm:"primaryKey"`
	CreatedAt     time.Time
	AuthRequestID uuid.UUID
	AuthRequest   *AuthRequest `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
