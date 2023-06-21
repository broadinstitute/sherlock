// Code generated by go-swagger; DO NOT EDIT.

package ci_runs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// PostAPIV2CiRunsReader is a Reader for the PostAPIV2CiRuns structure.
type PostAPIV2CiRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIV2CiRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAPIV2CiRunsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostAPIV2CiRunsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAPIV2CiRunsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAPIV2CiRunsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAPIV2CiRunsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewPostAPIV2CiRunsProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostAPIV2CiRunsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAPIV2CiRunsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAPIV2CiRunsOK creates a PostAPIV2CiRunsOK with default headers values
func NewPostAPIV2CiRunsOK() *PostAPIV2CiRunsOK {
	return &PostAPIV2CiRunsOK{}
}

/* PostAPIV2CiRunsOK describes a response with status code 200, with default header values.

OK
*/
type PostAPIV2CiRunsOK struct {
	Payload *models.V2controllersCiRun
}

func (o *PostAPIV2CiRunsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/ci-runs][%d] postApiV2CiRunsOK  %+v", 200, o.Payload)
}
func (o *PostAPIV2CiRunsOK) GetPayload() *models.V2controllersCiRun {
	return o.Payload
}

func (o *PostAPIV2CiRunsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V2controllersCiRun)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2CiRunsCreated creates a PostAPIV2CiRunsCreated with default headers values
func NewPostAPIV2CiRunsCreated() *PostAPIV2CiRunsCreated {
	return &PostAPIV2CiRunsCreated{}
}

/* PostAPIV2CiRunsCreated describes a response with status code 201, with default header values.

Created
*/
type PostAPIV2CiRunsCreated struct {
	Payload *models.V2controllersCiRun
}

func (o *PostAPIV2CiRunsCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/ci-runs][%d] postApiV2CiRunsCreated  %+v", 201, o.Payload)
}
func (o *PostAPIV2CiRunsCreated) GetPayload() *models.V2controllersCiRun {
	return o.Payload
}

func (o *PostAPIV2CiRunsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V2controllersCiRun)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2CiRunsBadRequest creates a PostAPIV2CiRunsBadRequest with default headers values
func NewPostAPIV2CiRunsBadRequest() *PostAPIV2CiRunsBadRequest {
	return &PostAPIV2CiRunsBadRequest{}
}

/* PostAPIV2CiRunsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostAPIV2CiRunsBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIV2CiRunsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/ci-runs][%d] postApiV2CiRunsBadRequest  %+v", 400, o.Payload)
}
func (o *PostAPIV2CiRunsBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2CiRunsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2CiRunsForbidden creates a PostAPIV2CiRunsForbidden with default headers values
func NewPostAPIV2CiRunsForbidden() *PostAPIV2CiRunsForbidden {
	return &PostAPIV2CiRunsForbidden{}
}

/* PostAPIV2CiRunsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostAPIV2CiRunsForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIV2CiRunsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/ci-runs][%d] postApiV2CiRunsForbidden  %+v", 403, o.Payload)
}
func (o *PostAPIV2CiRunsForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2CiRunsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2CiRunsNotFound creates a PostAPIV2CiRunsNotFound with default headers values
func NewPostAPIV2CiRunsNotFound() *PostAPIV2CiRunsNotFound {
	return &PostAPIV2CiRunsNotFound{}
}

/* PostAPIV2CiRunsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostAPIV2CiRunsNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIV2CiRunsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/ci-runs][%d] postApiV2CiRunsNotFound  %+v", 404, o.Payload)
}
func (o *PostAPIV2CiRunsNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2CiRunsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2CiRunsProxyAuthenticationRequired creates a PostAPIV2CiRunsProxyAuthenticationRequired with default headers values
func NewPostAPIV2CiRunsProxyAuthenticationRequired() *PostAPIV2CiRunsProxyAuthenticationRequired {
	return &PostAPIV2CiRunsProxyAuthenticationRequired{}
}

/* PostAPIV2CiRunsProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type PostAPIV2CiRunsProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIV2CiRunsProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[POST /api/v2/ci-runs][%d] postApiV2CiRunsProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *PostAPIV2CiRunsProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2CiRunsProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2CiRunsConflict creates a PostAPIV2CiRunsConflict with default headers values
func NewPostAPIV2CiRunsConflict() *PostAPIV2CiRunsConflict {
	return &PostAPIV2CiRunsConflict{}
}

/* PostAPIV2CiRunsConflict describes a response with status code 409, with default header values.

Conflict
*/
type PostAPIV2CiRunsConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIV2CiRunsConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/ci-runs][%d] postApiV2CiRunsConflict  %+v", 409, o.Payload)
}
func (o *PostAPIV2CiRunsConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2CiRunsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2CiRunsInternalServerError creates a PostAPIV2CiRunsInternalServerError with default headers values
func NewPostAPIV2CiRunsInternalServerError() *PostAPIV2CiRunsInternalServerError {
	return &PostAPIV2CiRunsInternalServerError{}
}

/* PostAPIV2CiRunsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostAPIV2CiRunsInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIV2CiRunsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/ci-runs][%d] postApiV2CiRunsInternalServerError  %+v", 500, o.Payload)
}
func (o *PostAPIV2CiRunsInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2CiRunsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}