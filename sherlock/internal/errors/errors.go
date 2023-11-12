package errors

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

// This package is designed to help influence the status code sent when an error is returned via an API.
// Just include one of the strings below in your error message text and the appropriate status code will
// get set if that error is ever serialized into an API response.
// InternalServerError takes precedence, and if an error doesn't include any of the strings at all then
// it is treated like an InternalServerError.
const (
	BadRequest                  = "HTTP Bad Request"                   // 400
	Forbidden                   = "HTTP Forbidden"                     // 403
	NotFound                    = "HTTP Not Found"                     // 404
	MethodNotAllowed            = "HTTP Method Not Allowed"            // 405
	ProxyAuthenticationRequired = "HTTP Proxy Authentication Required" // 407
	Conflict                    = "HTTP Conflict"                      // 409
	InternalServerError         = "HTTP Internal Server Error"         // 500
)

type ErrorResponse struct {
	ToBlame string `json:"toBlame"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

func errorToApiResponse(err error) (int, ErrorResponse) {
	code, response := convert(err)
	if response.ToBlame == "server" {
		log.Error().Msgf("BODY | %3d | %s", code, response.Message)
	}
	return code, response
}

func convert(err error) (int, ErrorResponse) {
	errorString := err.Error()

	// Even though InternalServerError is the catch-all, we want it to take precedence
	// over all others if it was actually explicitly set.
	if strings.Contains(errorString, InternalServerError) {
		return http.StatusInternalServerError, ErrorResponse{
			ToBlame: "server",
			Type:    InternalServerError,
			Message: errorString,
		}
	}

	if strings.Contains(errorString, BadRequest) {
		return http.StatusBadRequest, ErrorResponse{
			ToBlame: "client",
			Type:    BadRequest,
			Message: errorString,
		}
	}
	if strings.Contains(errorString, Forbidden) {
		return http.StatusForbidden, ErrorResponse{
			ToBlame: "client",
			Type:    Forbidden,
			Message: errorString,
		}
	}
	if strings.Contains(errorString, NotFound) {
		return http.StatusNotFound, ErrorResponse{
			ToBlame: "client",
			Type:    NotFound,
			Message: errorString,
		}
	}
	if strings.Contains(errorString, MethodNotAllowed) {
		return http.StatusMethodNotAllowed, ErrorResponse{
			ToBlame: "client",
			Type:    MethodNotAllowed,
			Message: errorString,
		}
	}
	if strings.Contains(errorString, ProxyAuthenticationRequired) {
		return http.StatusProxyAuthRequired, ErrorResponse{
			ToBlame: "client",
			Type:    ProxyAuthenticationRequired,
			Message: errorString,
		}
	}
	if strings.Contains(errorString, Conflict) {
		return http.StatusConflict, ErrorResponse{
			ToBlame: "client",
			Type:    Conflict,
			Message: errorString,
		}
	}

	// If we're about to return a 500 to the client, quickly check if
	// we can infer a better response code from a Gorm error.
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusNotFound, ErrorResponse{
			ToBlame: "client",
			Type:    NotFound,
			Message: errorString,
		}
	}
	if errors.Is(err, gorm.ErrDuplicatedKey) || strings.Contains(err.Error(), "violates unique constraint") {
		return http.StatusConflict, ErrorResponse{
			ToBlame: "client",
			Type:    Conflict,
			Message: errorString,
		}
	}
	if strings.Contains(err.Error(), "violates check constraint") || strings.Contains(err.Error(), "violates not-null constraint") {
		return http.StatusBadRequest, ErrorResponse{
			ToBlame: "client",
			Type:    BadRequest,
			Message: errorString,
		}
	}

	// If we're still about to return a 500 to the client, check if
	// there was a Go-internal conversion error and assume that would
	// be on the client.
	if errors.Is(err, strconv.ErrRange) || errors.Is(err, strconv.ErrSyntax) {
		return http.StatusBadRequest, ErrorResponse{
			ToBlame: "client",
			Type:    BadRequest,
			Message: fmt.Sprintf("(%s) string conversion error: %v", BadRequest, err),
		}
	}

	return http.StatusInternalServerError, ErrorResponse{
		ToBlame: "server",
		Type:    InternalServerError,
		Message: errorString,
	}
}
