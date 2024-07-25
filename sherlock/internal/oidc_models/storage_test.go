package oidc_models

import (
	"context"
	"database/sql"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/go-jose/go-jose/v4"
	"github.com/google/uuid"
	"github.com/zitadel/oidc/v3/pkg/oidc"
	"golang.org/x/text/language"
	"gorm.io/gorm/clause"
	"strings"
	"time"
)

func (s *oidcModelsSuite) TestStorageImpl_CreateAuthRequest() {
	clientID, _, err := s.GenerateClient(s.DB)
	s.NoError(err)

	authRequest, err := s.storage.CreateAuthRequest(context.Background(), &oidc.AuthRequest{
		Scopes:              []string{oidc.ScopeOpenID, oidc.ScopeProfile, oidc.ScopeEmail, groupsClaim},
		ResponseType:        oidc.ResponseTypeIDTokenOnly,
		ClientID:            clientID,
		RedirectURI:         s.GeneratedClientRedirectURI(),
		State:               "some-state",
		Nonce:               "some-nonce",
		ResponseMode:        oidc.ResponseModeQuery,
		CodeChallenge:       "code-challenge",
		CodeChallengeMethod: oidc.CodeChallengeMethodS256,
	}, "")
	s.NoError(err)

	var authRequests []AuthRequest
	s.NoError(s.DB.Find(&authRequests).Error)
	s.Len(authRequests, 1)

	s.Equal(authRequest.GetID(), authRequests[0].ID.String())
	s.Equal(clientID, authRequests[0].ClientID)
	s.Equal(oidc.SpaceDelimitedArray{oidc.ScopeOpenID, oidc.ScopeProfile, oidc.ScopeEmail, groupsClaim}, authRequests[0].Scopes)
	s.Equal(oidc.ResponseTypeIDTokenOnly, authRequests[0].ResponseType)
	s.Equal(s.GeneratedClientRedirectURI(), authRequests[0].RedirectURI)
	s.Equal("some-state", authRequests[0].State)
	s.Equal("some-nonce", authRequests[0].Nonce)
	s.Equal(oidc.ResponseModeQuery, authRequests[0].ResponseMode)
	s.Equal("code-challenge", authRequests[0].CodeChallenge)
	s.Equal(oidc.CodeChallengeMethodS256, authRequests[0].CodeChallengeMethod)
}

func (s *oidcModelsSuite) TestStorageImpl_AuthRequestByID() {
	clientID, _, err := s.GenerateClient(s.DB)
	s.NoError(err)

	authRequest := &AuthRequest{
		ID:          uuid.New(),
		DoneAt:      sql.NullTime{Time: time.Now(), Valid: true},
		ClientID:    clientID,
		Nonce:       "some-nonce",
		RedirectURI: s.GeneratedClientRedirectURI(),
		Scopes:      []string{oidc.ScopeOpenID, oidc.ScopeProfile, oidc.ScopeEmail, groupsClaim},
		State:       "some-state",
		UserID:      utils.PointerTo(s.TestData.User_Suitable().ID),
	}
	s.NoError(s.DB.Omit(clause.Associations).Create(authRequest).Error)

	authRequestByID, err := s.storage.AuthRequestByID(context.Background(), authRequest.ID.String())
	s.NoError(err)
	s.Equal(authRequest.ID.String(), authRequestByID.GetID())
}

func (s *oidcModelsSuite) TestStorageImpl_AuthRequestByCode() {
	clientID, _, err := s.GenerateClient(s.DB)
	s.NoError(err)

	authRequest := &AuthRequest{
		ID:          uuid.New(),
		DoneAt:      sql.NullTime{Time: time.Now(), Valid: true},
		ClientID:    clientID,
		Nonce:       "some-nonce",
		RedirectURI: s.GeneratedClientRedirectURI(),
		Scopes:      []string{oidc.ScopeOpenID, oidc.ScopeProfile, oidc.ScopeEmail, groupsClaim},
		State:       "some-state",
		UserID:      utils.PointerTo(s.TestData.User_Suitable().ID),
	}
	s.NoError(s.DB.Omit(clause.Associations).Create(authRequest).Error)

	code := "some-code"
	s.NoError(s.storage.SaveAuthCode(context.Background(), authRequest.ID.String(), code))

	authRequestByCode, err := s.storage.AuthRequestByCode(context.Background(), code)
	s.NoError(err)
	s.Equal(authRequest.ID.String(), authRequestByCode.GetID())
}

func (s *oidcModelsSuite) TestStorageImpl_SaveAuthCode() {
	clientID, _, err := s.GenerateClient(s.DB)
	s.NoError(err)

	authRequest := &AuthRequest{
		ID:          uuid.New(),
		DoneAt:      sql.NullTime{Time: time.Now(), Valid: true},
		ClientID:    clientID,
		Nonce:       "some-nonce",
		RedirectURI: s.GeneratedClientRedirectURI(),
		Scopes:      []string{oidc.ScopeOpenID, oidc.ScopeProfile, oidc.ScopeEmail, groupsClaim},
		State:       "some-state",
		UserID:      utils.PointerTo(s.TestData.User_Suitable().ID),
	}
	s.NoError(s.DB.Omit(clause.Associations).Create(authRequest).Error)

	code := "some-code"
	s.NoError(s.storage.SaveAuthCode(context.Background(), authRequest.ID.String(), code))

	var authRequestCodes []AuthRequestCode
	s.NoError(s.DB.Find(&authRequestCodes).Error)
	s.Len(authRequestCodes, 1)
	s.Equal(code, authRequestCodes[0].Code)
	s.Equal(authRequest.ID.String(), authRequestCodes[0].AuthRequestID.String())
}

func (s *oidcModelsSuite) TestStorageImpl_DeleteAuthRequest() {
	clientID, _, err := s.GenerateClient(s.DB)
	s.NoError(err)

	authRequest := &AuthRequest{
		ID:          uuid.New(),
		DoneAt:      sql.NullTime{Time: time.Now(), Valid: true},
		ClientID:    clientID,
		Nonce:       "some-nonce",
		RedirectURI: s.GeneratedClientRedirectURI(),
		Scopes:      []string{oidc.ScopeOpenID, oidc.ScopeProfile, oidc.ScopeEmail, groupsClaim},
		State:       "some-state",
		UserID:      utils.PointerTo(s.TestData.User_Suitable().ID),
	}
	s.NoError(s.DB.Omit(clause.Associations).Create(authRequest).Error)

	s.NoError(s.storage.DeleteAuthRequest(context.Background(), authRequest.ID.String()))

	var authRequests []AuthRequest
	s.NoError(s.DB.Find(&authRequests).Error)
	s.Len(authRequests, 0)
}

func (s *oidcModelsSuite) TestStorageImpl_CreateAccessToken() {
	clientID, _, err := s.GenerateClient(s.DB)
	s.NoError(err)
	accessTokenID, _, err := s.storage.CreateAccessToken(context.Background(), &AuthRequest{
		ID:          uuid.New(),
		DoneAt:      sql.NullTime{Time: time.Now(), Valid: true},
		ClientID:    clientID,
		Nonce:       "some-nonce",
		RedirectURI: s.GeneratedClientRedirectURI(),
		Scopes:      []string{oidc.ScopeOpenID, oidc.ScopeProfile, oidc.ScopeEmail, groupsClaim},
		State:       "some-state",
		UserID:      utils.PointerTo(s.TestData.User_Suitable().ID),
	})
	s.NoError(err)

	var tokens []Token
	s.NoError(s.DB.Find(&tokens).Error)
	s.Len(tokens, 1)
	s.Equal(accessTokenID, tokens[0].ID.String())
	s.Equal(clientID, tokens[0].ClientID)
	s.Equal(oidc.SpaceDelimitedArray{oidc.ScopeOpenID, oidc.ScopeProfile, oidc.ScopeEmail, groupsClaim}, tokens[0].Scopes)

}

func (s *oidcModelsSuite) TestStorageImpl_CreateAccessAndRefreshTokens() {
	clientID, _, err := s.GenerateClient(s.DB)
	s.NoError(err)
	authRequest := &AuthRequest{
		ID:          uuid.New(),
		DoneAt:      sql.NullTime{Time: time.Now(), Valid: true},
		ClientID:    clientID,
		Nonce:       "some-nonce",
		RedirectURI: s.GeneratedClientRedirectURI(),
		Scopes:      []string{oidc.ScopeOpenID, oidc.ScopeProfile, oidc.ScopeEmail, groupsClaim},
		State:       "some-state",
		UserID:      utils.PointerTo(s.TestData.User_Suitable().ID),
	}
	accessTokenID, _, _, err := s.storage.CreateAccessAndRefreshTokens(context.Background(), authRequest, "")
	s.NoError(err)

	var tokens []Token
	s.NoError(s.DB.Find(&tokens).Error)
	s.Len(tokens, 1)
	s.Equal(accessTokenID, tokens[0].ID.String())
	s.Equal(clientID, tokens[0].ClientID)
	s.Equal(oidc.SpaceDelimitedArray{oidc.ScopeOpenID, oidc.ScopeProfile, oidc.ScopeEmail, groupsClaim}, tokens[0].Scopes)

	var refreshTokens []RefreshToken
	s.NoError(s.DB.Find(&refreshTokens).Error)
	s.Len(refreshTokens, 1)
}

func (s *oidcModelsSuite) TestStorageImpl_TokenRequestByRefreshToken() {
	clientID, _, err := s.GenerateClient(s.DB)
	s.NoError(err)
	refreshTokenModel, _, err := s.storage.createRefreshToken(clientID, []string{oidc.ScopeOpenID, oidc.ScopeProfile, oidc.ScopeEmail, groupsClaim}, s.TestData.User_Suitable().ID, time.Now())
	s.NoError(err)

	tokenRequest, err := s.storage.TokenRequestByRefreshToken(context.Background(), refreshTokenModel.ID.String())
	s.NoError(err)
	s.Equal(refreshTokenModel.OriginalAuthAt.Second(), tokenRequest.GetAuthTime().Second())
}

func (s *oidcModelsSuite) TestStorageImpl_TerminateSession() {
	clientID, _, err := s.GenerateClient(s.DB)
	s.NoError(err)
	_, _, _, err = s.storage.CreateAccessAndRefreshTokens(context.Background(), &AuthRequest{
		ID:          uuid.New(),
		DoneAt:      sql.NullTime{Time: time.Now(), Valid: true},
		ClientID:    clientID,
		Nonce:       "some-nonce",
		RedirectURI: s.GeneratedClientRedirectURI(),
		Scopes:      []string{oidc.ScopeOpenID, oidc.ScopeProfile, oidc.ScopeEmail, groupsClaim},
		State:       "some-state",
		UserID:      utils.PointerTo(s.TestData.User_Suitable().ID),
	}, "")
	s.NoError(err)
	s.NoError(s.storage.TerminateSession(context.Background(), utils.UintToString(s.TestData.User_Suitable().ID), clientID))
	var refreshTokens []RefreshToken
	s.NoError(s.DB.Find(&refreshTokens).Error)
	s.Len(refreshTokens, 0)
	var tokens []Token
	s.NoError(s.DB.Find(&tokens).Error)
	s.Len(tokens, 0)
}

func (s *oidcModelsSuite) TestStorageImpl_RevokeToken() {
	clientID, _, err := s.GenerateClient(s.DB)
	s.NoError(err)
	s.Run("refresh token", func() {
		_, refreshToken, _, err := s.storage.CreateAccessAndRefreshTokens(context.Background(), &AuthRequest{
			ID:          uuid.New(),
			DoneAt:      sql.NullTime{Time: time.Now(), Valid: true},
			ClientID:    clientID,
			Nonce:       "some-nonce",
			RedirectURI: s.GeneratedClientRedirectURI(),
			Scopes:      []string{oidc.ScopeOpenID, oidc.ScopeProfile, oidc.ScopeEmail, groupsClaim},
			State:       "some-state",
			UserID:      utils.PointerTo(s.TestData.User_Suitable().ID),
		}, "")
		s.NoError(err)

		s.Nil(s.storage.RevokeToken(context.Background(), refreshToken, utils.UintToString(s.TestData.User_Suitable().ID), clientID))

		var refreshTokens []RefreshToken
		s.NoError(s.DB.Find(&refreshTokens).Error)
		s.Len(refreshTokens, 0)
		var tokens []Token
		s.NoError(s.DB.Find(&tokens).Error)
		s.Len(tokens, 0)
	})
	s.Run("token", func() {
		accessTokenID, _, err := s.storage.CreateAccessToken(context.Background(), &AuthRequest{
			ID:          uuid.New(),
			DoneAt:      sql.NullTime{Time: time.Now(), Valid: true},
			ClientID:    clientID,
			Nonce:       "some-nonce",
			RedirectURI: s.GeneratedClientRedirectURI(),
			Scopes:      []string{oidc.ScopeOpenID, oidc.ScopeProfile, oidc.ScopeEmail, groupsClaim},
			State:       "some-state",
			UserID:      utils.PointerTo(s.TestData.User_Suitable().ID),
		})
		s.NoError(err)

		s.Nil(s.storage.RevokeToken(context.Background(), accessTokenID, utils.UintToString(s.TestData.User_Suitable().ID), clientID))

		var refreshTokens []RefreshToken
		s.NoError(s.DB.Find(&refreshTokens).Error)
		s.Len(refreshTokens, 0)
		var tokens []Token
		s.NoError(s.DB.Find(&tokens).Error)
		s.Len(tokens, 0)
	})
}

func (s *oidcModelsSuite) TestStorageImpl_GetRefreshTokenInfo() {
	clientID, _, err := s.GenerateClient(s.DB)
	s.NoError(err)
	refreshTokenModel, rawRefreshToken, err := s.storage.createRefreshToken(clientID, []string{oidc.ScopeOpenID, oidc.ScopeProfile, oidc.ScopeEmail, groupsClaim}, s.TestData.User_Suitable().ID, time.Now())
	s.NoError(err)

	userID, tokenID, err := s.storage.GetRefreshTokenInfo(context.Background(), clientID, rawRefreshToken)
	s.NoError(err)
	s.Equal(utils.UintToString(s.TestData.User_Suitable().ID), userID)
	s.Equal(refreshTokenModel.ID.String(), tokenID)
}

func (s *oidcModelsSuite) TestStorageImpl_SigningKey() {
	key1, err := saveNewSigningKey(context.Background(), s.DB)
	s.NoError(err)
	key2, err := saveNewSigningKey(context.Background(), s.DB)
	s.NoError(err)
	s.NoError(s.DB.Model(&key1).UpdateColumn("created_at", time.Now().Add(-time.Hour)).Error)
	signingKey, err := s.storage.SigningKey(context.Background())
	s.NoError(err)
	s.Equal(key2.ID.String(), signingKey.ID())
}

func (s *oidcModelsSuite) TestStorageImpl_KeySet() {
	key1, err := saveNewSigningKey(context.Background(), s.DB)
	s.NoError(err)
	key2, err := saveNewSigningKey(context.Background(), s.DB)
	s.NoError(err)
	keyset, err := s.storage.KeySet(context.Background())
	s.NoError(err)
	s.Len(keyset, 2)
	var key1Found, key2Found bool
	for _, key := range keyset {
		if key.ID() == key1.ID.String() {
			key1Found = true
		}
		if key.ID() == key2.ID.String() {
			key2Found = true
		}
	}
	s.True(key1Found)
	s.True(key2Found)
}

func (s *oidcModelsSuite) TestStorageImpl_GetClientByClientID() {
	clientID, _, err := s.GenerateClient(s.DB)
	s.NoError(err)
	client, err := s.storage.GetClientByClientID(context.Background(), clientID)
	s.NoError(err)
	s.Equal(clientID, client.GetID())
}

func (s *oidcModelsSuite) TestStorageImpl_AuthorizeClientIDSecret() {
	clientID, clientSecret, err := s.GenerateClient(s.DB)
	s.NoError(err)
	s.Run("valid", func() {
		s.NoError(s.storage.AuthorizeClientIDSecret(context.Background(), clientID, clientSecret))
	})
	s.Run("invalid", func() {
		s.Error(s.storage.AuthorizeClientIDSecret(context.Background(), clientID, "invalid"))
	})
}

func (s *oidcModelsSuite) TestStorageImpl_SetUserinfoFromScopes() {
	clientID, _, err := s.GenerateClient(s.DB)
	s.NoError(err)
	userinfo := &oidc.UserInfo{}
	s.NoError(s.storage.SetUserinfoFromScopes(context.Background(), userinfo, utils.UintToString(s.TestData.User_Suitable().ID), clientID, []string{oidc.ScopeOpenID, oidc.ScopeProfile, oidc.ScopeEmail, groupsClaim}))
	s.Empty(userinfo.Subject) // This method does nothing intentionally; the library says to implement the request handling instead
}

func (s *oidcModelsSuite) TestStorageImpl_SetUserinfoFromRequest() {
	userinfo := &oidc.UserInfo{}
	s.NoError(s.storage.SetUserinfoFromRequest(context.Background(), userinfo, &AuthRequest{
		UserID: utils.PointerTo(s.TestData.User_Suitable().ID),
	}, []string{oidc.ScopeOpenID, oidc.ScopeProfile, oidc.ScopeEmail, groupsClaim}))
	s.Equal(utils.UintToString(s.TestData.User_Suitable().ID), userinfo.Subject)
}

func (s *oidcModelsSuite) TestStorageImpl_SetUserinfoFromToken() {
	clientID, _, err := s.GenerateClient(s.DB)
	s.NoError(err)
	token := Token{
		ID:       uuid.New(),
		ClientID: clientID,
		Scopes:   []string{oidc.ScopeOpenID, oidc.ScopeProfile, oidc.ScopeEmail, groupsClaim},
		Expiry:   time.Now().Add(time.Hour),
		UserID:   s.TestData.User_Suitable().ID,
	}
	s.NoError(s.DB.Omit(clause.Associations).Create(&token).Error)
	userinfo := &oidc.UserInfo{}
	s.NoError(s.storage.SetUserinfoFromToken(context.Background(), userinfo, token.ID.String(), utils.UintToString(s.TestData.User_Suitable().ID), ""))
	s.Equal(utils.UintToString(s.TestData.User_Suitable().ID), userinfo.Subject)
}

func (s *oidcModelsSuite) TestStorageImpl_SetIntrospectionFromToken() {
	clientID, _, err := s.GenerateClient(s.DB)
	s.NoError(err)
	token := Token{
		ID:       uuid.New(),
		ClientID: clientID,
		Scopes:   []string{oidc.ScopeOpenID, oidc.ScopeProfile, oidc.ScopeEmail, groupsClaim},
		Expiry:   time.Now().Add(time.Hour),
		UserID:   s.TestData.User_Suitable().ID,
	}
	s.NoError(s.DB.Omit(clause.Associations).Create(&token).Error)
	introspection := &oidc.IntrospectionResponse{}
	s.NoError(s.storage.SetIntrospectionFromToken(context.Background(), introspection, token.ID.String(), utils.UintToString(s.TestData.User_Suitable().ID), clientID))
	s.Equal(utils.UintToString(s.TestData.User_Suitable().ID), introspection.Subject)
	s.Equal(token.Scopes, introspection.Scope)
	s.Equal(clientID, introspection.ClientID)
	s.True(introspection.Active)
}

func (s *oidcModelsSuite) TestStorageImpl_GetPrivateClaimsFromScopes() {
	s.Run("empty", func() {
		claims, err := s.storage.GetPrivateClaimsFromScopes(context.Background(), utils.UintToString(s.TestData.User_Suitable().ID), "", []string{})
		s.NoError(err)
		s.Empty(claims)
	})
	s.Run("groups", func() {
		claims, err := s.storage.GetPrivateClaimsFromScopes(context.Background(), utils.UintToString(s.TestData.User_Suitable().ID), "", []string{groupsClaim})
		s.NoError(err)
		if s.NotEmpty(claims) && s.Contains(claims, groupsClaim) && s.IsType([]string{}, claims[groupsClaim]) {
			groups := claims[groupsClaim].([]string)
			for _, ra := range s.TestData.User_Suitable().Assignments {
				if ra != nil && ra.IsActive() {
					s.Contains(groups, *ra.Role.Name)
				}
			}
		}
	})
	s.Run("extra", func() {
		claims, err := s.storage.GetPrivateClaimsFromScopes(context.Background(), utils.UintToString(s.TestData.User_Suitable().ID), "", []string{groupsClaim, "extra"})
		s.NoError(err)
		if s.NotEmpty(claims) && s.Contains(claims, groupsClaim) && s.IsType([]string{}, claims[groupsClaim]) {
			groups := claims[groupsClaim].([]string)
			for _, ra := range s.TestData.User_Suitable().Assignments {
				if ra != nil && ra.IsActive() {
					s.Contains(groups, *ra.Role.Name)
				}
			}
		}
	})
}

func (s *oidcModelsSuite) TestStorageImpl_SignatureAlgorithms() {
	algorithms, err := s.storage.SignatureAlgorithms(context.Background())
	s.NoError(err)
	s.Equal([]jose.SignatureAlgorithm{jose.ES512}, algorithms)
}

func (s *oidcModelsSuite) TestStorageImpl_GetKeyByIDAndClientID() {
	_, err := s.storage.GetKeyByIDAndClientID(context.Background(), "", "")
	s.ErrorContains(err, "aren't currently implemented")
}

func (s *oidcModelsSuite) TestStorageImpl_ValidateJWTProfileScopes() {
	s.Run("empty", func() {
		scopes, err := s.storage.ValidateJWTProfileScopes(context.Background(), "", []string{})
		s.NoError(err)
		s.Empty(scopes)
	})
	s.Run("all allowed", func() {
		scopes, err := s.storage.ValidateJWTProfileScopes(context.Background(), "", []string{oidc.ScopeOpenID, oidc.ScopeEmail, oidc.ScopeProfile})
		s.NoError(err)
		s.Equal([]string{oidc.ScopeOpenID, oidc.ScopeEmail, oidc.ScopeProfile}, scopes)
	})
	s.Run("extra stuff", func() {
		scopes, err := s.storage.ValidateJWTProfileScopes(context.Background(), "", []string{oidc.ScopeOpenID, oidc.ScopeEmail, oidc.ScopeProfile, "extra"})
		s.NoError(err)
		s.Equal([]string{oidc.ScopeOpenID, oidc.ScopeEmail, oidc.ScopeProfile}, scopes)
	})
}

func (s *oidcModelsSuite) TestStorageImpl_Health() {
	s.Run("error when nil db", func() {
		storage := &storageImpl{}
		s.Error(storage.Health(context.Background()))
	})
	s.Run("no error", func() {
		s.NoError(s.storage.Health(context.Background()))
	})
}

func (s *oidcModelsSuite) TestStorageImpl_createRefreshToken() {
	clientID, _, err := s.GenerateClient(s.DB)
	s.NoError(err)
	refreshToken, _, err := s.storage.createRefreshToken(clientID, []string{oidc.ScopeOpenID, oidc.ScopeProfile, oidc.ScopeEmail, groupsClaim}, s.TestData.User_Suitable().ID, time.Now())
	s.NoError(err)
	s.NotEmpty(refreshToken.ID)
	s.Equal(clientID, refreshToken.ClientID)
	s.Equal(oidc.SpaceDelimitedArray{oidc.ScopeOpenID, oidc.ScopeProfile, oidc.ScopeEmail, groupsClaim}, refreshToken.Scopes)
	s.Equal(s.TestData.User_Suitable().ID, refreshToken.UserID)

	var refreshTokens []RefreshToken
	s.NoError(s.DB.Find(&refreshTokens).Error)
	s.Len(refreshTokens, 1)
}

func (s *oidcModelsSuite) TestStorageImpl_revokeRefreshToken() {
	clientID, _, err := s.GenerateClient(s.DB)
	s.NoError(err)
	refreshToken, _, err := s.storage.createRefreshToken(clientID, []string{oidc.ScopeOpenID, oidc.ScopeProfile, oidc.ScopeEmail, groupsClaim}, s.TestData.User_Suitable().ID, time.Now())
	s.NoError(err)

	s.NoError(s.storage.revokeRefreshToken(refreshToken.ID))

	var refreshTokens []RefreshToken
	s.NoError(s.DB.Find(&refreshTokens).Error)
	s.Len(refreshTokens, 0)
}

func (s *oidcModelsSuite) TestStorageImpl_renewRefreshToken() {
	clientID, _, err := s.GenerateClient(s.DB)
	s.NoError(err)
	refreshToken, rawRefreshToken, err := s.storage.createRefreshToken(clientID, []string{oidc.ScopeOpenID, oidc.ScopeProfile, oidc.ScopeEmail, groupsClaim}, s.TestData.User_Suitable().ID, time.Now())
	s.NoError(err)

	newRefreshToken, _, err := s.storage.renewRefreshToken(rawRefreshToken)
	s.NoError(err)

	s.NotEqual(refreshToken.ID, newRefreshToken.ID)
	s.Equal(refreshToken.ClientID, newRefreshToken.ClientID)
	s.Equal(refreshToken.Scopes, newRefreshToken.Scopes)
	s.Equal(refreshToken.UserID, newRefreshToken.UserID)

	var refreshTokens []RefreshToken
	s.NoError(s.DB.Find(&refreshTokens).Error)
	s.Len(refreshTokens, 1)
}

func (s *oidcModelsSuite) TestStorageImpl_createAccessToken() {
	clientID, _, err := s.GenerateClient(s.DB)
	s.NoError(err)
	accessToken, err := s.storage.createAccessToken(clientID, nil, []string{oidc.ScopeOpenID, oidc.ScopeProfile, oidc.ScopeEmail, groupsClaim}, s.TestData.User_Suitable().ID)
	s.NoError(err)

	var tokens []Token
	s.NoError(s.DB.Find(&tokens).Error)
	s.Len(tokens, 1)
	s.Equal(accessToken.ID.String(), tokens[0].ID.String())
	s.Equal(clientID, tokens[0].ClientID)
	s.Equal(oidc.SpaceDelimitedArray{oidc.ScopeOpenID, oidc.ScopeProfile, oidc.ScopeEmail, groupsClaim}, tokens[0].Scopes)
	s.Equal(s.TestData.User_Suitable().ID, tokens[0].UserID)
}

func (s *oidcModelsSuite) TestStorageImpl_setUserInfo() {
	s.Run("no scopes", func() {
		userInfo := &oidc.UserInfo{}
		s.NoError(s.storage.setUserinfo(userInfo, s.TestData.User_Suitable().ID, []string{}))
		s.Empty(userInfo.Subject)
		s.Empty(userInfo.Email)
		s.Empty(userInfo.EmailVerified)
		s.Empty(userInfo.Name)
		s.Empty(userInfo.Nickname)
		s.Empty(userInfo.PreferredUsername)
		s.Empty(userInfo.Locale)
		s.Empty(userInfo.UpdatedAt)
		s.Empty(userInfo.GivenName)
		s.Empty(userInfo.FamilyName)
		s.Empty(userInfo.Website)
		s.Empty(userInfo.Claims)
	})
	s.Run("openid", func() {
		userInfo := &oidc.UserInfo{}
		s.NoError(s.storage.setUserinfo(userInfo, s.TestData.User_Suitable().ID, []string{oidc.ScopeOpenID}))
		s.Equal(utils.UintToString(s.TestData.User_Suitable().ID), userInfo.Subject)
		s.Empty(userInfo.Email)
		s.Empty(userInfo.EmailVerified)
		s.Empty(userInfo.Name)
		s.Empty(userInfo.Nickname)
		s.Empty(userInfo.PreferredUsername)
		s.Empty(userInfo.Locale)
		s.Empty(userInfo.UpdatedAt)
		s.Empty(userInfo.GivenName)
		s.Empty(userInfo.FamilyName)
		s.Empty(userInfo.Website)
		s.Empty(userInfo.Claims)
	})
	s.Run("email", func() {
		userInfo := &oidc.UserInfo{}
		s.NoError(s.storage.setUserinfo(userInfo, s.TestData.User_Suitable().ID, []string{oidc.ScopeEmail}))
		s.Empty(userInfo.Subject)
		s.Equal(s.TestData.User_Suitable().Email, userInfo.Email)
		s.True(bool(userInfo.EmailVerified))
		s.Empty(userInfo.Name)
		s.Empty(userInfo.Nickname)
		s.Empty(userInfo.PreferredUsername)
		s.Empty(userInfo.Locale)
		s.Empty(userInfo.UpdatedAt)
		s.Empty(userInfo.GivenName)
		s.Empty(userInfo.FamilyName)
		s.Empty(userInfo.Website)
		s.Empty(userInfo.Claims)
	})
	s.Run("profile", func() {
		userInfo := &oidc.UserInfo{}
		s.NoError(s.storage.setUserinfo(userInfo, s.TestData.User_Suitable().ID, []string{oidc.ScopeProfile}))
		s.Empty(userInfo.Subject)
		s.Empty(userInfo.Email)
		s.Empty(userInfo.EmailVerified)
		s.Equal(utils.PointerTo(s.TestData.User_Suitable()).NameOrUsername(), userInfo.Name)
		s.Equal(utils.PointerTo(s.TestData.User_Suitable()).NameOrUsername(), userInfo.Nickname)
		s.Equal(utils.PointerTo(s.TestData.User_Suitable()).AlphaNumericHyphenatedUsername(), userInfo.PreferredUsername)
		s.Equal(oidc.NewLocale(language.AmericanEnglish), userInfo.Locale)
		s.Equal(oidc.FromTime(s.TestData.User_Suitable().UpdatedAt), userInfo.UpdatedAt)
		if s.TestData.User_Suitable().Name != nil {
			nameParts := strings.Split(*s.TestData.User_Suitable().Name, " ")
			s.Equal(nameParts[0], userInfo.GivenName)
			s.Equal(nameParts[len(nameParts)-1], userInfo.FamilyName)
		} else {
			s.Empty(userInfo.GivenName)
			s.Empty(userInfo.FamilyName)
		}
		s.Equal("https://broad.io/beehive/r/user/"+s.TestData.User_Suitable().Email, userInfo.Website)
		s.Empty(userInfo.Claims)
	})
	s.Run("groups", func() {
		userInfo := &oidc.UserInfo{}
		s.NoError(s.storage.setUserinfo(userInfo, s.TestData.User_Suitable().ID, []string{groupsClaim}))
		s.Empty(userInfo.Subject)
		s.Empty(userInfo.Email)
		s.Empty(userInfo.EmailVerified)
		s.Empty(userInfo.Name)
		s.Empty(userInfo.Nickname)
		s.Empty(userInfo.PreferredUsername)
		s.Empty(userInfo.Locale)
		s.Empty(userInfo.UpdatedAt)
		s.Empty(userInfo.GivenName)
		s.Empty(userInfo.FamilyName)
		s.Empty(userInfo.Website)
		if s.NotEmpty(userInfo.Claims) && s.Contains(userInfo.Claims, groupsClaim) && s.IsType([]string{}, userInfo.Claims[groupsClaim]) {
			groups := userInfo.Claims[groupsClaim].([]string)
			for _, ra := range s.TestData.User_Suitable().Assignments {
				if ra != nil && ra.IsActive() {
					s.Contains(groups, *ra.Role.Name)
				}
			}
		}
	})
	s.Run("openid email profile groups", func() {
		userInfo := &oidc.UserInfo{}
		s.NoError(s.storage.setUserinfo(userInfo, s.TestData.User_Suitable().ID, []string{oidc.ScopeOpenID, oidc.ScopeEmail, oidc.ScopeProfile, groupsClaim}))
		s.Equal(utils.UintToString(s.TestData.User_Suitable().ID), userInfo.Subject)
		s.Equal(s.TestData.User_Suitable().Email, userInfo.Email)
		s.True(bool(userInfo.EmailVerified))
		s.Equal(utils.PointerTo(s.TestData.User_Suitable()).NameOrUsername(), userInfo.Name)
		s.Equal(utils.PointerTo(s.TestData.User_Suitable()).NameOrUsername(), userInfo.Nickname)
		s.Equal(utils.PointerTo(s.TestData.User_Suitable()).AlphaNumericHyphenatedUsername(), userInfo.PreferredUsername)
		s.Equal(oidc.NewLocale(language.AmericanEnglish), userInfo.Locale)
		s.Equal(oidc.FromTime(s.TestData.User_Suitable().UpdatedAt), userInfo.UpdatedAt)
		if s.TestData.User_Suitable().Name != nil {
			nameParts := strings.Split(*s.TestData.User_Suitable().Name, " ")
			s.Equal(nameParts[0], userInfo.GivenName)
			s.Equal(nameParts[len(nameParts)-1], userInfo.FamilyName)
		} else {
			s.Empty(userInfo.GivenName)
			s.Empty(userInfo.FamilyName)
		}
		s.Equal("https://broad.io/beehive/r/user/"+s.TestData.User_Suitable().Email, userInfo.Website)
		if s.NotEmpty(userInfo.Claims) && s.Contains(userInfo.Claims, groupsClaim) && s.IsType([]string{}, userInfo.Claims[groupsClaim]) {
			groups := userInfo.Claims[groupsClaim].([]string)
			for _, ra := range s.TestData.User_Suitable().Assignments {
				if ra != nil && ra.IsActive() {
					s.Contains(groups, *ra.Role.Name)
				}
			}
		}
	})
	s.Run("groups only active role assignments", func() {
		suspendedRole := models.Role{
			RoleFields: models.RoleFields{
				Name: utils.PointerTo("test-role"),
			},
		}
		s.SetSelfSuperAdminForDB()
		s.NoError(s.DB.Create(&suspendedRole).Error)
		suspendedRoleAssignment := models.RoleAssignment{
			UserID: s.TestData.User_Suitable().ID,
			RoleID: suspendedRole.ID,
			RoleAssignmentFields: models.RoleAssignmentFields{
				Suspended: utils.PointerTo(true),
			},
		}
		s.NoError(s.DB.Create(&suspendedRoleAssignment).Error)

		// reload user to get updated assignments
		var user models.User
		s.NoError(s.DB.Scopes(models.ReadUserScope).First(&user, s.TestData.User_Suitable().ID).Error)

		// make this test extra explicit -- set these bools to true at key points
		var suspendedRoleAssignmentOnUser, suspendedRoleAssignmentOmittedFromUserinfo bool

		userInfo := &oidc.UserInfo{}
		s.NoError(s.storage.setUserinfo(userInfo, user.ID, []string{groupsClaim}))
		if s.NotEmpty(userInfo.Claims) && s.Contains(userInfo.Claims, groupsClaim) && s.IsType([]string{}, userInfo.Claims[groupsClaim]) {
			groups := userInfo.Claims[groupsClaim].([]string)
			for _, ra := range user.Assignments {
				if ra != nil {
					if *ra.Role.Name == *suspendedRole.Name {
						suspendedRoleAssignmentOnUser = true
					}

					if ra.IsActive() {
						s.Contains(groups, *ra.Role.Name)
					} else if s.NotContains(groups, *ra.Role.Name) && *ra.Role.Name == *suspendedRole.Name {
						suspendedRoleAssignmentOmittedFromUserinfo = true
					}
				}
			}
		}

		s.True(suspendedRoleAssignmentOnUser)
		s.True(suspendedRoleAssignmentOmittedFromUserinfo)
	})
}
