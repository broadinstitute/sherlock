package oidc_models

import (
	"time"

	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/google/uuid"
	"github.com/zitadel/oidc/v3/pkg/oidc"
	"github.com/zitadel/oidc/v3/pkg/op"
)

var _ op.TokenRequest = &Token{}

type Token struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	CreatedAt time.Time

	RefreshToken   *RefreshToken `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	RefreshTokenID *uuid.UUID

	Client   *Client `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ClientID string  // AKA Audience, Application ID
	Scopes   oidc.SpaceDelimitedArray
	Expiry   time.Time

	User   *models.User
	UserID uint
}

func (t *Token) GetSubject() string {
	return utils.UintToString(t.UserID)
}

func (t *Token) GetAudience() []string {
	return []string{t.ClientID}
}

func (t *Token) GetScopes() []string {
	return t.Scopes
}
