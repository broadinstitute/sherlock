package sherlock

import (
	"bytes"
	"encoding/json"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication/test_users"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"io"
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
	gin.SetMode(gin.TestMode)
}

func (s *handlerSuite) SetupTest() {
	s.TestSuiteHelper.SetupTest()
	s.internalRouter = gin.New()
	apiRouter := s.internalRouter.Group("api", authentication.TestMiddleware(s.DB, s.TestData)...)
	ConfigureRoutes(apiRouter)
}

func (s *handlerSuite) NewRequest(method string, url string, toJsonBody any) *http.Request {
	body, err := json.Marshal(toJsonBody)
	s.NoErrorf(err, "json.Marshal(%v) error", toJsonBody)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	s.NoErrorf(err, "http.NewRequest(%s, %s, ...) error", method, url)
	return req
}

func (s *handlerSuite) NewSuperAdminRequest(method string, url string, toJsonBody any) *http.Request {
	req := s.NewRequest(method, url, toJsonBody)
	s.UseSuperAdminUserFor(req)
	return req
}

func (s *handlerSuite) NewSuitableRequest(method string, url string, toJsonBody any) *http.Request {
	req := s.NewRequest(method, url, toJsonBody)
	s.UseSuitableUserFor(req)
	return req
}

func (s *handlerSuite) NewNonSuitableRequest(method string, url string, toJsonBody any) *http.Request {
	req := s.NewRequest(method, url, toJsonBody)
	s.UseNonSuitableUserFor(req)
	return req
}

func (s *handlerSuite) HandleRequest(req *http.Request, fromJsonBodyPointer any) int {
	recorder := httptest.NewRecorder()
	s.internalRouter.ServeHTTP(recorder, req)
	s.Equalf("application/json; charset=utf-8", recorder.Header().Get("Content-Type"), "unexpected content type")

	// Read the body out specifically so we can use it repeatedly
	body, err := io.ReadAll(recorder.Result().Body)
	if s.NoErrorf(err, "body read error") {

		// First try to decode to the expected type. We'll error if there were any unexpected fields.
		intendedDecoder := json.NewDecoder(bytes.NewBuffer(body))
		intendedDecoder.DisallowUnknownFields()
		err = intendedDecoder.Decode(fromJsonBodyPointer)

		// Assert that there shouldn't be an error, but enter this conditional if there was
		if !s.NoErrorf(err, "failed to decode body to %T", fromJsonBodyPointer) {

			// Now try to decode to the universal error type. Again, we'll error if there were any unexpected fields.
			var errorResponse errors.ErrorResponse
			errorDecoder := json.NewDecoder(bytes.NewBuffer(body))
			errorDecoder.DisallowUnknownFields()
			err = errorDecoder.Decode(&errorResponse)

			// Assert that there shouldn't be a parse error, and that the parsed error response isn't empty.
			if s.NoErrorf(err, "failed to decode unknown body to %T", errorResponse) && s.NotZerof(errorResponse, "unknown body didn't supply %T fields", errorResponse) {
				// If those assertions held, that means the test failed and just returned an error. We'll fail the test with an explicit error to help debug.
				s.FailNowf("unexpected error response", "%s blamed on %s: %s", errorResponse.Type, errorResponse.ToBlame, errorResponse.Message)
			} else {
				// If those assertions didn't hold, that means there was a totally unexpected type in the response. We'll fail the test with the body to help debug.
				s.FailNowf("fully unexpected response type", "body: %s", string(body))
			}
		}
	}

	return recorder.Code
}
