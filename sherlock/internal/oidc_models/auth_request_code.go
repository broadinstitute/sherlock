package oidc_models

import (
	"github.com/google/uuid"
	"time"
)

type AuthRequestCode struct {
	Code          string `gorm:"primaryKey"`
	CreatedAt     time.Time
	AuthRequestID uuid.UUID
	AuthRequest   *AuthRequest `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
