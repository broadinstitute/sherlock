// Code generated by go-swagger; DO NOT EDIT.

package app_versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/clients/go/client/models"
)

// PostAPIV2AppVersionsReader is a Reader for the PostAPIV2AppVersions structure.
type PostAPIV2AppVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIV2AppVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAPIV2AppVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostAPIV2AppVersionsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAPIV2AppVersionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAPIV2AppVersionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAPIV2AppVersionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewPostAPIV2AppVersionsProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostAPIV2AppVersionsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAPIV2AppVersionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAPIV2AppVersionsOK creates a PostAPIV2AppVersionsOK with default headers values
func NewPostAPIV2AppVersionsOK() *PostAPIV2AppVersionsOK {
	return &PostAPIV2AppVersionsOK{}
}

/*
PostAPIV2AppVersionsOK describes a response with status code 200, with default header values.

OK
*/
type PostAPIV2AppVersionsOK struct {
	Payload *models.V2controllersAppVersion
}

// IsSuccess returns true when this post Api v2 app versions o k response has a 2xx status code
func (o *PostAPIV2AppVersionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post Api v2 app versions o k response has a 3xx status code
func (o *PostAPIV2AppVersionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Api v2 app versions o k response has a 4xx status code
func (o *PostAPIV2AppVersionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post Api v2 app versions o k response has a 5xx status code
func (o *PostAPIV2AppVersionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post Api v2 app versions o k response a status code equal to that given
func (o *PostAPIV2AppVersionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostAPIV2AppVersionsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/app-versions][%d] postApiV2AppVersionsOK  %+v", 200, o.Payload)
}

func (o *PostAPIV2AppVersionsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/app-versions][%d] postApiV2AppVersionsOK  %+v", 200, o.Payload)
}

func (o *PostAPIV2AppVersionsOK) GetPayload() *models.V2controllersAppVersion {
	return o.Payload
}

func (o *PostAPIV2AppVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V2controllersAppVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2AppVersionsCreated creates a PostAPIV2AppVersionsCreated with default headers values
func NewPostAPIV2AppVersionsCreated() *PostAPIV2AppVersionsCreated {
	return &PostAPIV2AppVersionsCreated{}
}

/*
PostAPIV2AppVersionsCreated describes a response with status code 201, with default header values.

Created
*/
type PostAPIV2AppVersionsCreated struct {
	Payload *models.V2controllersAppVersion
}

// IsSuccess returns true when this post Api v2 app versions created response has a 2xx status code
func (o *PostAPIV2AppVersionsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post Api v2 app versions created response has a 3xx status code
func (o *PostAPIV2AppVersionsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Api v2 app versions created response has a 4xx status code
func (o *PostAPIV2AppVersionsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post Api v2 app versions created response has a 5xx status code
func (o *PostAPIV2AppVersionsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post Api v2 app versions created response a status code equal to that given
func (o *PostAPIV2AppVersionsCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostAPIV2AppVersionsCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/app-versions][%d] postApiV2AppVersionsCreated  %+v", 201, o.Payload)
}

func (o *PostAPIV2AppVersionsCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/app-versions][%d] postApiV2AppVersionsCreated  %+v", 201, o.Payload)
}

func (o *PostAPIV2AppVersionsCreated) GetPayload() *models.V2controllersAppVersion {
	return o.Payload
}

func (o *PostAPIV2AppVersionsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V2controllersAppVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2AppVersionsBadRequest creates a PostAPIV2AppVersionsBadRequest with default headers values
func NewPostAPIV2AppVersionsBadRequest() *PostAPIV2AppVersionsBadRequest {
	return &PostAPIV2AppVersionsBadRequest{}
}

/*
PostAPIV2AppVersionsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostAPIV2AppVersionsBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

// IsSuccess returns true when this post Api v2 app versions bad request response has a 2xx status code
func (o *PostAPIV2AppVersionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Api v2 app versions bad request response has a 3xx status code
func (o *PostAPIV2AppVersionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Api v2 app versions bad request response has a 4xx status code
func (o *PostAPIV2AppVersionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post Api v2 app versions bad request response has a 5xx status code
func (o *PostAPIV2AppVersionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post Api v2 app versions bad request response a status code equal to that given
func (o *PostAPIV2AppVersionsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostAPIV2AppVersionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/app-versions][%d] postApiV2AppVersionsBadRequest  %+v", 400, o.Payload)
}

func (o *PostAPIV2AppVersionsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/app-versions][%d] postApiV2AppVersionsBadRequest  %+v", 400, o.Payload)
}

func (o *PostAPIV2AppVersionsBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2AppVersionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2AppVersionsForbidden creates a PostAPIV2AppVersionsForbidden with default headers values
func NewPostAPIV2AppVersionsForbidden() *PostAPIV2AppVersionsForbidden {
	return &PostAPIV2AppVersionsForbidden{}
}

/*
PostAPIV2AppVersionsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostAPIV2AppVersionsForbidden struct {
	Payload *models.ErrorsErrorResponse
}

// IsSuccess returns true when this post Api v2 app versions forbidden response has a 2xx status code
func (o *PostAPIV2AppVersionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Api v2 app versions forbidden response has a 3xx status code
func (o *PostAPIV2AppVersionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Api v2 app versions forbidden response has a 4xx status code
func (o *PostAPIV2AppVersionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post Api v2 app versions forbidden response has a 5xx status code
func (o *PostAPIV2AppVersionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post Api v2 app versions forbidden response a status code equal to that given
func (o *PostAPIV2AppVersionsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostAPIV2AppVersionsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/app-versions][%d] postApiV2AppVersionsForbidden  %+v", 403, o.Payload)
}

func (o *PostAPIV2AppVersionsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/app-versions][%d] postApiV2AppVersionsForbidden  %+v", 403, o.Payload)
}

func (o *PostAPIV2AppVersionsForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2AppVersionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2AppVersionsNotFound creates a PostAPIV2AppVersionsNotFound with default headers values
func NewPostAPIV2AppVersionsNotFound() *PostAPIV2AppVersionsNotFound {
	return &PostAPIV2AppVersionsNotFound{}
}

/*
PostAPIV2AppVersionsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostAPIV2AppVersionsNotFound struct {
	Payload *models.ErrorsErrorResponse
}

// IsSuccess returns true when this post Api v2 app versions not found response has a 2xx status code
func (o *PostAPIV2AppVersionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Api v2 app versions not found response has a 3xx status code
func (o *PostAPIV2AppVersionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Api v2 app versions not found response has a 4xx status code
func (o *PostAPIV2AppVersionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post Api v2 app versions not found response has a 5xx status code
func (o *PostAPIV2AppVersionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post Api v2 app versions not found response a status code equal to that given
func (o *PostAPIV2AppVersionsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostAPIV2AppVersionsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/app-versions][%d] postApiV2AppVersionsNotFound  %+v", 404, o.Payload)
}

func (o *PostAPIV2AppVersionsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/app-versions][%d] postApiV2AppVersionsNotFound  %+v", 404, o.Payload)
}

func (o *PostAPIV2AppVersionsNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2AppVersionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2AppVersionsProxyAuthenticationRequired creates a PostAPIV2AppVersionsProxyAuthenticationRequired with default headers values
func NewPostAPIV2AppVersionsProxyAuthenticationRequired() *PostAPIV2AppVersionsProxyAuthenticationRequired {
	return &PostAPIV2AppVersionsProxyAuthenticationRequired{}
}

/*
PostAPIV2AppVersionsProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type PostAPIV2AppVersionsProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

// IsSuccess returns true when this post Api v2 app versions proxy authentication required response has a 2xx status code
func (o *PostAPIV2AppVersionsProxyAuthenticationRequired) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Api v2 app versions proxy authentication required response has a 3xx status code
func (o *PostAPIV2AppVersionsProxyAuthenticationRequired) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Api v2 app versions proxy authentication required response has a 4xx status code
func (o *PostAPIV2AppVersionsProxyAuthenticationRequired) IsClientError() bool {
	return true
}

// IsServerError returns true when this post Api v2 app versions proxy authentication required response has a 5xx status code
func (o *PostAPIV2AppVersionsProxyAuthenticationRequired) IsServerError() bool {
	return false
}

// IsCode returns true when this post Api v2 app versions proxy authentication required response a status code equal to that given
func (o *PostAPIV2AppVersionsProxyAuthenticationRequired) IsCode(code int) bool {
	return code == 407
}

func (o *PostAPIV2AppVersionsProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[POST /api/v2/app-versions][%d] postApiV2AppVersionsProxyAuthenticationRequired  %+v", 407, o.Payload)
}

func (o *PostAPIV2AppVersionsProxyAuthenticationRequired) String() string {
	return fmt.Sprintf("[POST /api/v2/app-versions][%d] postApiV2AppVersionsProxyAuthenticationRequired  %+v", 407, o.Payload)
}

func (o *PostAPIV2AppVersionsProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2AppVersionsProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2AppVersionsConflict creates a PostAPIV2AppVersionsConflict with default headers values
func NewPostAPIV2AppVersionsConflict() *PostAPIV2AppVersionsConflict {
	return &PostAPIV2AppVersionsConflict{}
}

/*
PostAPIV2AppVersionsConflict describes a response with status code 409, with default header values.

Conflict
*/
type PostAPIV2AppVersionsConflict struct {
	Payload *models.ErrorsErrorResponse
}

// IsSuccess returns true when this post Api v2 app versions conflict response has a 2xx status code
func (o *PostAPIV2AppVersionsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Api v2 app versions conflict response has a 3xx status code
func (o *PostAPIV2AppVersionsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Api v2 app versions conflict response has a 4xx status code
func (o *PostAPIV2AppVersionsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post Api v2 app versions conflict response has a 5xx status code
func (o *PostAPIV2AppVersionsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post Api v2 app versions conflict response a status code equal to that given
func (o *PostAPIV2AppVersionsConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PostAPIV2AppVersionsConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/app-versions][%d] postApiV2AppVersionsConflict  %+v", 409, o.Payload)
}

func (o *PostAPIV2AppVersionsConflict) String() string {
	return fmt.Sprintf("[POST /api/v2/app-versions][%d] postApiV2AppVersionsConflict  %+v", 409, o.Payload)
}

func (o *PostAPIV2AppVersionsConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2AppVersionsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2AppVersionsInternalServerError creates a PostAPIV2AppVersionsInternalServerError with default headers values
func NewPostAPIV2AppVersionsInternalServerError() *PostAPIV2AppVersionsInternalServerError {
	return &PostAPIV2AppVersionsInternalServerError{}
}

/*
PostAPIV2AppVersionsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostAPIV2AppVersionsInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

// IsSuccess returns true when this post Api v2 app versions internal server error response has a 2xx status code
func (o *PostAPIV2AppVersionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Api v2 app versions internal server error response has a 3xx status code
func (o *PostAPIV2AppVersionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Api v2 app versions internal server error response has a 4xx status code
func (o *PostAPIV2AppVersionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post Api v2 app versions internal server error response has a 5xx status code
func (o *PostAPIV2AppVersionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post Api v2 app versions internal server error response a status code equal to that given
func (o *PostAPIV2AppVersionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostAPIV2AppVersionsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/app-versions][%d] postApiV2AppVersionsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAPIV2AppVersionsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/app-versions][%d] postApiV2AppVersionsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAPIV2AppVersionsInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2AppVersionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
