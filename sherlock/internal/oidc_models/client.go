package oidc_models

import (
	"fmt"
	"time"

	"github.com/zitadel/oidc/v3/pkg/oidc"
	"github.com/zitadel/oidc/v3/pkg/op"
)

// "Client implements the op.Client interface"
var _ op.Client = &Client{}

type Client struct {
	ID                           string                   `gorm:"primaryKey"`
	ClientSecretHash             []byte                   // PBKDF2 derived key, HMAC-SHA-512; should be empty for PKCE; Sherlock will derive the same number of bytes as the length of this field automatically
	ClientSecretSalt             []byte                   // Salt for ClientSecretHash; should be empty for PKCE
	ClientSecretIterations       int                      // Number of iterations for ClientSecretHash; should be empty for PKCE. https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html#pbkdf2
	ClientRedirectURIs           oidc.SpaceDelimitedArray // In dev mode, may include globs or use http
	ClientPostLogoutRedirectURIs oidc.SpaceDelimitedArray // In dev mode, may include globs or use http
	ClientApplicationType        op.ApplicationType       // "web", "user_agent", or "native"
	ClientAuthMethod             oidc.AuthMethod          // "client_secret_basic", "client_secret_post", "none", or "private_key_jwt"
	ClientIDTokenLifetime        int64                    // Nanoseconds
	ClientDevMode                bool                     // Allow insecure/nonstandard redirect URIs, like globs and http
	ClientClockSkew              int64                    // Nanoseconds
}

func (c *Client) GetID() string {
	return c.ID
}

func (c *Client) RedirectURIs() []string {
	return c.ClientRedirectURIs
}

func (c *Client) PostLogoutRedirectURIs() []string {
	return c.ClientPostLogoutRedirectURIs
}

func (c *Client) ApplicationType() op.ApplicationType {
	return c.ClientApplicationType
}

func (c *Client) AuthMethod() oidc.AuthMethod {
	return c.ClientAuthMethod
}

func (c *Client) ResponseTypes() []oidc.ResponseType {
	return []oidc.ResponseType{oidc.ResponseTypeCode, oidc.ResponseTypeIDTokenOnly, oidc.ResponseTypeIDToken}
}

func (c *Client) GrantTypes() []oidc.GrantType {
	return []oidc.GrantType{oidc.GrantTypeCode, oidc.GrantTypeRefreshToken}
}

func (c *Client) LoginURL(authRequestID string) string {
	// To sherlock/sherlock/internal/api/login/login.go
	return fmt.Sprintf("/login?id=%s", authRequestID)
}

func (c *Client) AccessTokenType() op.AccessTokenType {
	return op.AccessTokenTypeJWT
}

func (c *Client) IDTokenLifetime() time.Duration {
	return time.Duration(c.ClientIDTokenLifetime)
}

func (c *Client) DevMode() bool {
	return c.ClientDevMode
}

func (c *Client) IDTokenUserinfoClaimsAssertion() bool {
	// It technically violates spec to return userinfo claims in ID tokens when an access token is provided.
	// The caller is expected to call the userinfo endpoint with the access token to get the userinfo data.
	// But... Dex supports and actually expects these claims in the initial ID token. It can be configured
	// to call the userinfo endpoint, but why make the extra roundtrip?
	//
	// In the future we could always have this be a bool field on Client itself to customize the behavior
	// per client.
	return true
}

func (c *Client) ClockSkew() time.Duration {
	return time.Duration(c.ClientClockSkew)
}

func (c *Client) RestrictAdditionalIdTokenScopes() func(scopes []string) []string {
	return func(scopes []string) []string { return scopes }
}

func (c *Client) RestrictAdditionalAccessTokenScopes() func(scopes []string) []string {
	// Hardcode no client-level extra restrictions on access token scopes
	return func(scopes []string) []string { return scopes }
}

func (c *Client) IsScopeAllowed(scope string) bool {
	// Hardcode that groups are allowed, because currently that's the whole point of Sherlock doing OIDC
	return scope == groupsClaim
}

// "devModeGlobClient implements the op.HasRedirectGlobs interface"
var _ op.HasRedirectGlobs = &devModeGlobClient{}

type devModeGlobClient struct {
	Client
}

func (c *devModeGlobClient) RedirectURIGlobs() []string {
	return c.RedirectURIs()
}

func (c *devModeGlobClient) PostLogoutRedirectURIGlobs() []string {
	return c.PostLogoutRedirectURIs()
}

// wrapPossibleDevModeClient takes a Client from the database returns an op.Client, possibly wrapped with
// devModeGlobClient based on Client.ClientDevMode.
func wrapPossibleDevModeClient(client Client) op.Client {
	if client.DevMode() {
		return &devModeGlobClient{client}
	} else {
		return &client
	}
}
