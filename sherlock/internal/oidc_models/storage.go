package oidc_models

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha512"
	"errors"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/go-jose/go-jose/v4"
	"github.com/google/uuid"
	"github.com/zitadel/oidc/v3/pkg/oidc"
	"github.com/zitadel/oidc/v3/pkg/op"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/text/language"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
	"time"
)

const groupsClaim = "groups"

var _ op.Storage = &storageImpl{}
var _ op.CanSetUserinfoFromRequest = &storageImpl{}

type storageImpl struct {
	db *gorm.DB
}

func (s *storageImpl) CreateAuthRequest(_ context.Context, authRequest *oidc.AuthRequest, hintedUserID string) (op.AuthRequest, error) {

	requestUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to create UUID: %w", err)
	}

	request := AuthRequest{
		ID:                  requestUUID,
		ClientID:            authRequest.ClientID,
		Nonce:               authRequest.Nonce,
		RedirectURI:         authRequest.RedirectURI,
		ResponseType:        authRequest.ResponseType,
		ResponseMode:        authRequest.ResponseMode,
		Scopes:              authRequest.Scopes,
		State:               authRequest.State,
		CodeChallenge:       authRequest.CodeChallenge,
		CodeChallengeMethod: authRequest.CodeChallengeMethod,
	}

	if hintedUserID != "" {
		parsedHintedUserID, err := utils.ParseUint(hintedUserID)
		if err != nil {
			return nil, fmt.Errorf("failed to parse user ID: %w", err)
		}
		request.UserID = &parsedHintedUserID
	}

	err = s.db.Omit(clause.Associations).Create(&request).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create auth request: %w", err)
	}

	return &request, nil
}

func (s *storageImpl) AuthRequestByID(_ context.Context, requestID string) (op.AuthRequest, error) {
	validRequestUUID, err := uuid.Parse(requestID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse request UUID: %w", err)
	}
	var request AuthRequest
	err = s.db.Omit(clause.Associations).Where(&AuthRequest{ID: validRequestUUID}).Take(&request).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get auth request: %w", err)
	}
	return &request, nil
}

func (s *storageImpl) AuthRequestByCode(_ context.Context, code string) (op.AuthRequest, error) {
	var requestCode AuthRequestCode
	err := s.db.Preload("AuthRequest").Where(&AuthRequestCode{Code: code}).Take(&requestCode).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get auth request code: %w", err)
	} else if requestCode.AuthRequest == nil {
		return nil, fmt.Errorf("auth request not found for code: %s", code)
	}
	return requestCode.AuthRequest, nil
}

func (s *storageImpl) SaveAuthCode(_ context.Context, requestID string, code string) error {
	validRequestUUID, err := uuid.Parse(requestID)
	if err != nil {
		return fmt.Errorf("failed to parse request UUID: %w", err)
	}
	requestCode := AuthRequestCode{
		Code:          code,
		AuthRequestID: validRequestUUID,
	}
	err = s.db.Omit(clause.Associations).Create(&requestCode).Error
	if err != nil {
		return fmt.Errorf("failed to save auth code: %w", err)
	}
	return nil
}

func (s *storageImpl) DeleteAuthRequest(_ context.Context, requestID string) error {
	validRequestUUID, err := uuid.Parse(requestID)
	if err != nil {
		return fmt.Errorf("failed to parse request UUID: %w", err)
	}
	// Auth codes deleted via foreign key cascade
	err = s.db.Omit(clause.Associations).Where(&AuthRequest{ID: validRequestUUID}).Delete(&AuthRequest{}).Error
	if err != nil {
		return fmt.Errorf("failed to delete auth request: %w", err)
	}
	return nil
}

func (s *storageImpl) CreateAccessToken(_ context.Context, request op.TokenRequest) (accessTokenID string, expiration time.Time, err error) {
	var clientID string
	var userID uint
	switch req := request.(type) {
	case *AuthRequest:
		if !req.Done() || req.UserID == nil {
			return "", time.Time{}, oidc.ErrLoginRequired()
		}
		clientID = req.ClientID
		userID = *req.UserID
		// It is possible for requests to be of another type -- like op.TokenExchangeRequest -- but we don't support
		// that currently
	default:
		return "", time.Time{}, fmt.Errorf("unsupported request type: %T", request)
	}

	token, err := s.createAccessToken(clientID, nil, request.GetScopes(), userID)
	if err != nil {
		return "", time.Time{}, err
	}
	return token.ID.String(), token.Expiry, nil
}

func (s *storageImpl) CreateAccessAndRefreshTokens(_ context.Context, request op.TokenRequest, currentRefreshToken string) (accessTokenID string, newRefreshToken string, expiration time.Time, err error) {
	var clientID string
	var userID uint
	var authTime time.Time
	switch req := request.(type) {
	case *AuthRequest:
		if !req.Done() || req.UserID == nil {
			return "", "", time.Time{}, oidc.ErrLoginRequired()
		}
		clientID = req.ClientID
		userID = *req.UserID
		authTime = req.GetAuthTime()
	case *RefreshToken:
		clientID = req.ClientID
		userID = req.UserID
		authTime = req.GetAuthTime()
	default:
		return "", "", time.Time{}, fmt.Errorf("unsupported request type: %T", request)
	}

	var refreshTokenModel *RefreshToken
	if currentRefreshToken == "" {
		refreshTokenModel, newRefreshToken, err = s.createRefreshToken(clientID, request.GetScopes(), userID, authTime)
	} else {
		refreshTokenModel, newRefreshToken, err = s.renewRefreshToken(currentRefreshToken)
	}
	if err != nil {
		return "", "", time.Time{}, fmt.Errorf("failed to create refresh token: %w", err)
	}

	var tokenModel *Token
	tokenModel, err = s.createAccessToken(clientID, &refreshTokenModel.ID, request.GetScopes(), userID)
	if err != nil {
		return "", "", time.Time{}, fmt.Errorf("failed to create access token: %w", err)
	}

	return tokenModel.ID.String(), newRefreshToken, tokenModel.Expiry, nil
}

func (s *storageImpl) TokenRequestByRefreshToken(_ context.Context, refreshTokenID string) (op.RefreshTokenRequest, error) {
	parsedRefreshTokenID, err := uuid.Parse(refreshTokenID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse refresh token ID: %w", err)
	}
	var refreshToken RefreshToken
	err = s.db.Omit(clause.Associations).Where(&RefreshToken{ID: parsedRefreshTokenID}).Take(&refreshToken).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get refresh token: %w", err)
	}
	return &refreshToken, nil
}

func (s *storageImpl) TerminateSession(_ context.Context, userID string, clientID string) error {
	parsedUserID, err := utils.ParseUint(userID)
	if err != nil {
		return fmt.Errorf("failed to parse user ID: %w", err)
	}
	err1 := s.db.Omit(clause.Associations).Where(&RefreshToken{ClientID: clientID, UserID: parsedUserID}).Delete(&RefreshToken{}).Error
	// Most tokens should've been caught by the foreign key cascade, but just in case we'll make sure we wipe all matching tokens too
	err2 := s.db.Omit(clause.Associations).Where(&Token{ClientID: clientID, UserID: parsedUserID}).Delete(&Token{}).Error
	if err1 != nil && !errors.Is(err1, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to delete refresh tokens: %w", err1)
	}
	if err2 != nil && !errors.Is(err2, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to delete access tokens: %w", err2)
	}
	return nil
}

func (s *storageImpl) RevokeToken(_ context.Context, tokenOrTokenID string, userID string, clientID string) *oidc.Error {
	var tokenUUIDToRevoke, refreshTokenUUIDToRevoke *uuid.UUID

	// UserID can be empty! If it's passed we'll filter on it.
	var queryUserID uint
	if userID != "" {
		var err error
		queryUserID, err = utils.ParseUint(userID)
		if err != nil {
			return oidc.ErrServerError().WithParent(err).WithDescription("failed to parse user ID")
		}
	}

	tokenUUID, err := uuid.Parse(tokenOrTokenID)
	if err != nil {
		// Token's not a UUID, so all we can do is revoke a matching refresh token
		hash := sha512.Sum512([]byte(tokenOrTokenID))
		var refreshToken RefreshToken
		err = s.db.Omit(clause.Associations).Where(&RefreshToken{TokenHash: hash[:], ClientID: clientID, UserID: queryUserID}).Take(&refreshToken).Error
		if err == nil {
			refreshTokenUUIDToRevoke = &refreshToken.ID
		}
	} else {
		// Look up as token UUID
		var token Token
		err = s.db.Omit(clause.Associations).Where(&Token{ID: tokenUUID, ClientID: clientID, UserID: queryUserID}).Take(&token).Error
		if err == nil {
			tokenUUIDToRevoke = &token.ID
			refreshTokenUUIDToRevoke = token.RefreshTokenID
		} else {
			// Look up as refresh token UUID
			var refreshToken RefreshToken
			err = s.db.Omit(clause.Associations).Where(&RefreshToken{ID: tokenUUID, ClientID: clientID, UserID: queryUserID}).Take(&refreshToken).Error
			if err == nil {
				refreshTokenUUIDToRevoke = &refreshToken.ID
			}
		}
	}

	if tokenUUIDToRevoke != nil {
		err = s.db.Omit(clause.Associations).Where(&Token{ID: *tokenUUIDToRevoke}).Delete(&Token{}).Error
	}
	if refreshTokenUUIDToRevoke != nil {
		err = s.revokeRefreshToken(*refreshTokenUUIDToRevoke)
	}
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return oidc.ErrInvalidRequest().WithDescription("token not found")
		} else {
			return oidc.ErrServerError().WithParent(err).WithDescription("failed to revoke token")
		}
	}
	return nil
}

func (s *storageImpl) GetRefreshTokenInfo(_ context.Context, clientID string, token string) (userID string, tokenID string, err error) {
	hash := sha512.Sum512([]byte(token))
	var refreshToken RefreshToken
	err = s.db.Omit(clause.Associations).Where(&RefreshToken{TokenHash: hash[:], ClientID: clientID}).Take(&refreshToken).Error
	if err != nil {
		return "", "", fmt.Errorf("failed to get refresh token: %w", err)
	}
	return utils.UintToString(refreshToken.UserID), refreshToken.ID.String(), nil
}

func (s *storageImpl) SigningKey(_ context.Context) (op.SigningKey, error) {
	var key SigningKey
	err := s.db.Omit(clause.Associations).Order("created_at desc").First(&key).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get signing key: %w", err)
	}
	return decryptPrivateSigningKey(&key)
}

func (s *storageImpl) KeySet(_ context.Context) ([]op.Key, error) {
	var keys []SigningKey
	err := s.db.Omit(clause.Associations).Order("created_at desc").Find(&keys).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get signing keys: %w", err)
	}
	return utils.Map(keys, func(rawKey SigningKey) op.Key {
		return &publicSigningKey{SigningKey: rawKey}
	}), nil
}

func (s *storageImpl) GetClientByClientID(_ context.Context, clientID string) (op.Client, error) {
	var client Client
	err := s.db.Omit(clause.Associations).Where(&Client{ID: clientID}).Take(&client).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get client: %w", err)
	}
	return wrapPossibleDevModeClient(client), nil
}

func (s *storageImpl) AuthorizeClientIDSecret(_ context.Context, clientID, clientSecret string) error {
	var client Client
	err := s.db.Omit(clause.Associations).Where(&Client{ID: clientID}).Take(&client).Error
	if err != nil {
		return fmt.Errorf("failed to get client: %w", err)
	}
	if len(client.ClientSecretSalt) == 0 || len(client.ClientSecretHash) == 0 || client.ClientSecretIterations == 0 {
		return fmt.Errorf("client secret not configured; perhaps client should be using PKCE")
	}
	derivedKey := pbkdf2.Key([]byte(clientSecret), client.ClientSecretSalt, client.ClientSecretIterations, len(client.ClientSecretHash), sha512.New)
	if !bytes.Equal(derivedKey, client.ClientSecretHash) {
		return fmt.Errorf("client secret does not match")
	}
	return nil
}

// SetUserinfoFromScopes is an empty implementation because according to Zitadel's example, SetUserinfoFromRequest
// should be implemented instead.
func (s *storageImpl) SetUserinfoFromScopes(_ context.Context, _ *oidc.UserInfo, _, _ string, _ []string) error {
	return nil
}

func (s *storageImpl) SetUserinfoFromRequest(_ context.Context, userinfo *oidc.UserInfo, request op.IDTokenRequest, scopes []string) error {
	userID, err := utils.ParseUint(request.GetSubject())
	if err != nil {
		return fmt.Errorf("failed to parse user ID: %w", err)
	}
	return s.setUserinfo(userinfo, userID, scopes)
}

func (s *storageImpl) SetUserinfoFromToken(_ context.Context, userinfo *oidc.UserInfo, tokenID, subject, _ string) error {
	parsedTokenID, err := uuid.Parse(tokenID)
	if err != nil {
		return fmt.Errorf("failed to parse token ID: %w", err)
	}
	var token Token
	err = s.db.Omit(clause.Associations).Where(&Token{ID: parsedTokenID}).Take(&token).Error
	if err != nil {
		return fmt.Errorf("failed to get token: %w", err)
	}
	if utils.UintToString(token.UserID) != subject {
		// This check is theoretically unnecessary because the library should be covering this,
		// but we have the info so why not.
		return fmt.Errorf("token mismatched with subject")
	}

	// We ignore the origin argument to this function because CORS is handled by Gin middleware, no reason for us
	// to implement that here.

	return s.setUserinfo(userinfo, token.UserID, token.Scopes)
}

// SetIntrospectionFromToken has some arguments we ignore. We ignore the subject because we can get the user ID in
// a more verified way from the token.
func (s *storageImpl) SetIntrospectionFromToken(_ context.Context, introspection *oidc.IntrospectionResponse, tokenID, subject, clientID string) error {
	parsedUserID, err := utils.ParseUint(subject)
	if err != nil {
		return fmt.Errorf("failed to parse user ID: %w", err)
	}

	parsedTokenID, err := uuid.Parse(tokenID)
	if err != nil {
		return fmt.Errorf("failed to parse token ID: %w", err)
	}
	var token Token
	err = s.db.Omit(clause.Associations).Where(&Token{ID: parsedTokenID, ClientID: clientID, UserID: parsedUserID}).Take(&token).Error
	if err != nil {
		return fmt.Errorf("failed to get token: %w", err)
	}

	// In our case, we basically equate audience, client ID, and application.
	// This kind of audience == client ID check is "the right way" to do this,
	// though, so we do it to avoid a gotcha down the line.
	for _, aud := range token.GetAudience() {
		if aud == clientID {
			userInfo := new(oidc.UserInfo)
			err = s.setUserinfo(userInfo, token.UserID, token.Scopes)
			if err != nil {
				return fmt.Errorf("failed to set userinfo: %w", err)
			}
			introspection.SetUserInfo(userInfo)
			introspection.Scope = token.Scopes
			introspection.ClientID = token.ClientID
			introspection.Active = true
			return nil
		}
	}
	return fmt.Errorf("token not valid for client")
}

func (s *storageImpl) GetPrivateClaimsFromScopes(_ context.Context, userID, _ string, scopes []string) (map[string]any, error) {
	claims := make(map[string]any)
	for _, scope := range scopes {
		switch scope {
		case groupsClaim:
			// Only get the user if we need to
			parsedUserID, err := utils.ParseUint(userID)
			if err != nil {
				return nil, fmt.Errorf("failed to parse user ID: %w", err)
			}
			var user models.User
			err = s.db.Where(&models.User{Model: gorm.Model{ID: parsedUserID}}).Scopes(models.ReadUserScope).Take(&user).Error
			if err != nil {
				return nil, fmt.Errorf("failed to get user: %w", err)
			}
			groups := make([]string, 0, len(user.Assignments))
			for _, assignment := range user.Assignments {
				if assignment != nil && assignment.IsActive() && assignment.Role != nil && assignment.Role.Name != nil {
					groups = append(groups, *assignment.Role.Name)
				}
			}
			claims[groupsClaim] = groups
		}
	}
	return claims, nil
}

func (s *storageImpl) SignatureAlgorithms(_ context.Context) ([]jose.SignatureAlgorithm, error) {
	return []jose.SignatureAlgorithm{jose.RS256}, nil
}

func (s *storageImpl) GetKeyByIDAndClientID(_ context.Context, _, _ string) (*jose.JSONWebKey, error) {
	// What we're meant to do here is define a list of (client ID, key ID, public key) tuples in configuration or something.
	// The idea is that a client application would have the private key, and it could then sign JWTs *to send to us*, per
	// RFC 7523 (urn:ietf:params:oauth:grant-type:jwt-bearer). zitadel-oidc would call this function to get the public key,
	// validate the signature, and then return the requested access token. This allows a client to get an access token for
	// a user without user interaction.
	//
	// We don't support that grant type for Sherlock so we leave this unimplemented. If somehow this does get called,
	// the error will bubble up as if the assertion was invalid.
	//
	// https://datatracker.ietf.org/doc/html/rfc7523
	return nil, fmt.Errorf("JWT Profile Authorization Grants (RFC 7523) aren't currently implemented")
}

func (s *storageImpl) ValidateJWTProfileScopes(_ context.Context, _ string, scopes []string) ([]string, error) {
	allowedScopes := make([]string, 0)
	for _, scope := range scopes {
		if scope == oidc.ScopeOpenID || scope == oidc.ScopeProfile || scope == oidc.ScopeEmail {
			allowedScopes = append(allowedScopes, scope)
		}
	}
	return allowedScopes, nil
}

// Health isn't something we particularly need to rely on, because this OIDC provider exists in the context of a larger
// application that already has liveness and readiness probes configured. If the OIDC API is available, the OIDC provider
// will be healthy.
func (s *storageImpl) Health(_ context.Context) error {
	if s.db == nil {
		return fmt.Errorf("no database connection")
	}
	return nil
}

func (s *storageImpl) createRefreshToken(clientID string, scopes []string, userID uint, authAt time.Time) (*RefreshToken, string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, "", fmt.Errorf("failed to create UUID: %w", err)
	}
	refreshToken := make([]byte, 256)
	_, err = rand.Read(refreshToken)
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate refresh token: %w", err)
	}
	hash := sha512.Sum512(refreshToken)
	refreshTokenModel := &RefreshToken{
		ID:             id,
		TokenHash:      hash[:],
		ClientID:       clientID,
		Scopes:         scopes,
		OriginalAuthAt: authAt,
		UserID:         userID,
	}
	err = s.db.Omit(clause.Associations).Create(refreshTokenModel).Error
	if err != nil {
		return nil, "", fmt.Errorf("failed to create refresh token: %w", err)
	}
	return refreshTokenModel, string(refreshToken), nil
}

func (s *storageImpl) revokeRefreshToken(refreshTokenID uuid.UUID) error {
	// Normal tokens are deleted via foreign key cascade
	err := s.db.Omit(clause.Associations).Where(&RefreshToken{ID: refreshTokenID}).Delete(&RefreshToken{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to delete refresh token: %w", err)
	}
	return nil
}

func (s *storageImpl) renewRefreshToken(refreshToken string) (*RefreshToken, string, error) {
	hash := sha512.Sum512([]byte(refreshToken))
	var refreshTokenModel RefreshToken
	err := s.db.Omit(clause.Associations).Where(&RefreshToken{TokenHash: hash[:]}).Take(&refreshTokenModel).Error
	if err != nil {
		return nil, "", fmt.Errorf("failed to get refresh token: %w", err)
	}

	err = s.revokeRefreshToken(refreshTokenModel.ID)
	if err != nil {
		return nil, "", fmt.Errorf("failed to revoke existing refresh token: %w", err)
	}

	return s.createRefreshToken(refreshTokenModel.ClientID, refreshTokenModel.Scopes, refreshTokenModel.UserID, refreshTokenModel.OriginalAuthAt)
}

func (s *storageImpl) createAccessToken(clientID string, refreshTokenID *uuid.UUID, scopes []string, userID uint) (*Token, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to create UUID: %w", err)
	}
	tokenModel := &Token{
		ID:             id,
		RefreshTokenID: refreshTokenID,
		ClientID:       clientID,
		Scopes:         scopes,
		Expiry:         time.Now().Add(config.Config.MustDuration("oidc.tokenDuration")),
		UserID:         userID,
	}

	err = s.db.Omit(clause.Associations).Create(tokenModel).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create access token: %w", err)
	}

	return tokenModel, nil
}

func (s *storageImpl) setUserinfo(userInfo *oidc.UserInfo, userID uint, scopes []string) error {
	var user models.User
	err := s.db.Where(&models.User{Model: gorm.Model{ID: userID}}).Scopes(models.ReadUserScope).Take(&user).Error
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}
	for _, scope := range scopes {
		switch scope {
		case oidc.ScopeOpenID:
			userInfo.Subject = utils.UintToString(user.ID)
		case oidc.ScopeEmail:
			userInfo.Email = user.Email
			userInfo.EmailVerified = true
		case oidc.ScopeProfile:
			userInfo.PreferredUsername = user.AlphaNumericHyphenatedUsername()
			userInfo.Nickname = user.NameOrEmailHandle()
			userInfo.Locale = oidc.NewLocale(language.AmericanEnglish)
			userInfo.UpdatedAt = oidc.FromTime(user.UpdatedAt)
			if user.Name != nil {
				userInfo.Name = *user.Name
				nameParts := strings.Split(*user.Name, " ")
				if len(nameParts) > 0 {
					userInfo.GivenName = nameParts[0]
				}
				if len(nameParts) > 1 {
					userInfo.FamilyName = nameParts[len(nameParts)-1]
				}
			}
			userInfo.Website = "https://broad.io/beehive/r/user/" + user.Email // :shrug:
		case groupsClaim:
			groups := make([]string, 0, len(user.Assignments))
			for _, assignment := range user.Assignments {
				if assignment != nil && assignment.IsActive() && assignment.Role != nil && assignment.Role.Name != nil {
					groups = append(groups, *assignment.Role.Name)
				}
			}
			userInfo.AppendClaims(groupsClaim, groups)
		}
	}
	return nil
}
