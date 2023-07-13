package sherlock

import (
	"bytes"
	"encoding/json"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/test_users"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type handlerSuite struct {
	suite.Suite
	test_users.TestUserHelper
	models.TestSuiteHelper

	internalRouter *gin.Engine
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(handlerSuite))
}

func (s *handlerSuite) SetupSuite() {
	s.TestSuiteHelper.SetupSuite()
	// Reduces console output
	gin.SetMode(gin.ReleaseMode)
}

func (s *handlerSuite) SetupTest() {
	s.TestSuiteHelper.SetupTest()
	s.internalRouter = gin.New()
	apiRouter := s.internalRouter.Group("api", authentication.UserMiddleware(s.DB), authentication.DbMiddleware(s.DB))
	ConfigureRoutes(apiRouter)
}

func (s *handlerSuite) NewRequest(method string, url string, toJsonBody any) *http.Request {
	body, err := json.Marshal(toJsonBody)
	s.NoErrorf(err, "json.Marshal(%v) error", toJsonBody)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	s.NoErrorf(err, "http.NewRequest(%s, %s, ...) error", method, url)
	return req
}

func (s *handlerSuite) HandleRequest(req *http.Request, fromJsonBodyPointer any) int {
	recorder := httptest.NewRecorder()
	s.internalRouter.ServeHTTP(recorder, req)
	s.Equalf("application/json; charset=utf-8", recorder.Header().Get("Content-Type"), "unexpected content type")
	decoder := json.NewDecoder(recorder.Result().Body)
	decoder.DisallowUnknownFields()
	s.NoErrorf(decoder.Decode(fromJsonBodyPointer), "JSON decode error")
	return recorder.Code
}
