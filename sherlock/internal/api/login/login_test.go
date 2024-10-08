package login

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/oidc_models"
	"github.com/google/uuid"
	"github.com/zitadel/oidc/v3/pkg/oidc"
	"gorm.io/gorm/clause"
	"net/http"
	"net/http/httptest"
	"time"
)

func (s *handlerSuite) TestLoginGet_noAuthRequestID() {
	request, err := http.NewRequest("GET", "/login", nil)
	s.NoError(err)
	recorder := httptest.NewRecorder()
	s.internalRouter.ServeHTTP(recorder, request)

	s.Equal(http.StatusBadRequest, recorder.Code)
}

func (s *handlerSuite) TestLoginGet_invalidAuthRequestID() {
	request, err := http.NewRequest("GET", "/login?id=invalid", nil)
	s.NoError(err)
	recorder := httptest.NewRecorder()
	s.internalRouter.ServeHTTP(recorder, request)

	s.Equal(http.StatusBadRequest, recorder.Code)
}

func (s *handlerSuite) TestLoginGet() {
	clientID, _, err := s.GenerateClient(s.DB)
	s.NoError(err)

	authRequest := oidc_models.AuthRequest{
		ID:          uuid.New(),
		ClientID:    clientID,
		Nonce:       "some-nonce",
		RedirectURI: s.GeneratedClientRedirectURI(),
		Scopes:      []string{oidc.ScopeOpenID, oidc.ScopeProfile, oidc.ScopeEmail, "groups"},
		State:       "some-state",
	}

	s.NoError(s.DB.Omit(clause.Associations).Create(&authRequest).Error)

	request, err := http.NewRequest("GET", "/login?id="+authRequest.GetID(), nil)
	s.NoError(err)
	s.UseSuitableUserFor(request)
	recorder := httptest.NewRecorder()
	s.internalRouter.ServeHTTP(recorder, request)

	s.Equal(http.StatusFound, recorder.Code)

	// Check that the auth request was marked as done
	var reloadedAuthRequest oidc_models.AuthRequest
	s.NoError(s.DB.Where("id = ?", authRequest.ID.String()).First(&reloadedAuthRequest).Error)
	s.True(reloadedAuthRequest.DoneAt.Valid)
	s.Equal(s.TestData.User_Suitable().ID, *reloadedAuthRequest.UserID)
}

func (s *handlerSuite) TestLoginGet_DeactivatedUser() {
	clientID, _, err := s.GenerateClient(s.DB)
	s.NoError(err)

	authRequest := oidc_models.AuthRequest{
		ID:          uuid.New(),
		ClientID:    clientID,
		Nonce:       "some-nonce",
		RedirectURI: s.GeneratedClientRedirectURI(),
		Scopes:      []string{oidc.ScopeOpenID, oidc.ScopeProfile, oidc.ScopeEmail, "groups"},
		State:       "some-state",
	}

	s.NoError(s.DB.Omit(clause.Associations).Create(&authRequest).Error)

	request, err := http.NewRequest("GET", "/login?id="+authRequest.GetID(), nil)
	s.NoError(err)
	s.UseSuitableUserFor(request)

	// Deactivate the user right before making the request
	s.SetUserForDB(utils.PointerTo(s.TestData.User_SuperAdmin()))
	s.NoError(s.DB.Model(utils.PointerTo(s.TestData.User_Suitable())).Omit(clause.Associations).Update("deactivated_at", utils.PointerTo(time.Now())).Error)

	recorder := httptest.NewRecorder()
	s.internalRouter.ServeHTTP(recorder, request)

	s.Equal(http.StatusForbidden, recorder.Code)

	// Check that the auth request was not marked as done
	var reloadedAuthRequest oidc_models.AuthRequest
	s.NoError(s.DB.Where("id = ?", authRequest.ID.String()).First(&reloadedAuthRequest).Error)
	s.False(reloadedAuthRequest.DoneAt.Valid)
}
