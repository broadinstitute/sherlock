package errors

import (
	"errors"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"net/http"
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
	ProxyAuthenticationRequired = "HTTP Proxy Authentication Required" // 407
	Conflict                    = "HTTP Conflict"                      // 409
	InternalServerError         = "HTTP Internal Server Error"         // 500
)

type ErrorResponse struct {
	ToBlame string `json:"toBlame"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

func ErrorToApiResponse(err error) (int, ErrorResponse) {
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
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return http.StatusConflict, ErrorResponse{
			ToBlame: "client",
			Type:    Conflict,
			Message: errorString,
		}
	}

	return http.StatusInternalServerError, ErrorResponse{
		ToBlame: "server",
		Type:    InternalServerError,
		Message: errorString,
	}
}
