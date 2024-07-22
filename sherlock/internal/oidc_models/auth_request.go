package oidc_models

import (
	"database/sql"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/google/uuid"
	"github.com/zitadel/oidc/v3/pkg/oidc"
	"github.com/zitadel/oidc/v3/pkg/op"
	"time"
)

// "AuthRequest implements the op.AuthRequest interface"
var _ op.AuthRequest = &AuthRequest{}

type AuthRequest struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	CreatedAt time.Time
	DoneAt    sql.NullTime

	Client       *Client `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ClientID     string  // AKA Audience, Application ID
	Nonce        string
	RedirectURI  string            // AKA CallbackURI
	ResponseType oidc.ResponseType `gorm:"string"`
	ResponseMode oidc.ResponseMode `gorm:"string"`
	Scopes       oidc.SpaceDelimitedArray
	State        string

	CodeChallenge       string
	CodeChallengeMethod oidc.CodeChallengeMethod `gorm:"string"`

	// The user won't be filled until the request has been logged-in to
	User   *models.User
	UserID *uint
}

func (r *AuthRequest) GetID() string {
	return r.ID.String()
}

func (r *AuthRequest) GetACR() string {
	return ""
}

func (r *AuthRequest) GetAMR() []string {
	// Return an empty array because we don't know for sure what AMRs IAP enforced on the caller.
	// https://openid.net/specs/openid-connect-core-1_0.html#IDToken
	// https://www.rfc-editor.org/info/rfc8176
	return []string{}
}

func (r *AuthRequest) GetAudience() []string {
	return []string{r.ClientID}
}

func (r *AuthRequest) GetAuthTime() time.Time {
	return r.CreatedAt
}

func (r *AuthRequest) GetClientID() string {
	return r.ClientID
}

func (r *AuthRequest) GetCodeChallenge() *oidc.CodeChallenge {
	return &oidc.CodeChallenge{
		Challenge: r.CodeChallenge,
		Method:    r.CodeChallengeMethod,
	}
}

func (r *AuthRequest) GetNonce() string {
	return r.Nonce
}

func (r *AuthRequest) GetRedirectURI() string {
	return r.RedirectURI
}

func (r *AuthRequest) GetResponseType() oidc.ResponseType {
	return r.ResponseType
}

func (r *AuthRequest) GetResponseMode() oidc.ResponseMode {
	return r.ResponseMode
}

func (r *AuthRequest) GetScopes() []string {
	return r.Scopes
}

func (r *AuthRequest) GetState() string {
	return r.State
}

func (r *AuthRequest) GetSubject() string {
	if r.UserID == nil {
		return ""
	}
	return utils.UintToString(*r.UserID)
}

func (r *AuthRequest) Done() bool {
	return r.DoneAt.Valid
}
