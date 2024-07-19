package oidc_models

import (
	"errors"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/google/uuid"
	"github.com/zitadel/oidc/v3/pkg/oidc"
	"github.com/zitadel/oidc/v3/pkg/op"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

var _ op.RefreshTokenRequest = &RefreshToken{}

type RefreshToken struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	CreatedAt time.Time

	TokenHash []byte // SHA-512 hash

	ClientID       string // AKA Audience, Application ID
	Scopes         oidc.SpaceDelimitedArray
	OriginalAuthAt time.Time

	User   *models.User
	UserID uint
}

func expireRefreshTokens(db *gorm.DB) error {
	err := db.
		Omit(clause.Associations).
		Where("created_at < ?", time.Now().Add(-config.Config.MustDuration("oidc.refreshTokenDuration"))).
		Delete(&RefreshToken{}).
		Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	} else {
		return nil
	}
}

func (r *RefreshToken) GetAMR() []string {
	return []string{}
}

func (r *RefreshToken) GetAudience() []string {
	return []string{r.ClientID}
}

func (r *RefreshToken) GetAuthTime() time.Time {
	return r.OriginalAuthAt
}

func (r *RefreshToken) GetClientID() string {
	return r.ClientID
}

func (r *RefreshToken) GetScopes() []string {
	return r.Scopes
}

func (r *RefreshToken) GetSubject() string {
	return utils.UintToString(r.UserID)
}

func (r *RefreshToken) SetCurrentScopes(scopes []string) {
	r.Scopes = scopes
}
