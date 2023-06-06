// Code generated by go-swagger; DO NOT EDIT.

package ci_runs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/clients/go/client/models"
)

// GetAPIV2CiRunsSelectorReader is a Reader for the GetAPIV2CiRunsSelector structure.
type GetAPIV2CiRunsSelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV2CiRunsSelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV2CiRunsSelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIV2CiRunsSelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIV2CiRunsSelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIV2CiRunsSelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPIV2CiRunsSelectorProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPIV2CiRunsSelectorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIV2CiRunsSelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIV2CiRunsSelectorOK creates a GetAPIV2CiRunsSelectorOK with default headers values
func NewGetAPIV2CiRunsSelectorOK() *GetAPIV2CiRunsSelectorOK {
	return &GetAPIV2CiRunsSelectorOK{}
}

/* GetAPIV2CiRunsSelectorOK describes a response with status code 200, with default header values.

OK
*/
type GetAPIV2CiRunsSelectorOK struct {
	Payload *models.V2controllersCiRun
}

func (o *GetAPIV2CiRunsSelectorOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/ci-runs/{selector}][%d] getApiV2CiRunsSelectorOK  %+v", 200, o.Payload)
}
func (o *GetAPIV2CiRunsSelectorOK) GetPayload() *models.V2controllersCiRun {
	return o.Payload
}

func (o *GetAPIV2CiRunsSelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V2controllersCiRun)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2CiRunsSelectorBadRequest creates a GetAPIV2CiRunsSelectorBadRequest with default headers values
func NewGetAPIV2CiRunsSelectorBadRequest() *GetAPIV2CiRunsSelectorBadRequest {
	return &GetAPIV2CiRunsSelectorBadRequest{}
}

/* GetAPIV2CiRunsSelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIV2CiRunsSelectorBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2CiRunsSelectorBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/ci-runs/{selector}][%d] getApiV2CiRunsSelectorBadRequest  %+v", 400, o.Payload)
}
func (o *GetAPIV2CiRunsSelectorBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2CiRunsSelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2CiRunsSelectorForbidden creates a GetAPIV2CiRunsSelectorForbidden with default headers values
func NewGetAPIV2CiRunsSelectorForbidden() *GetAPIV2CiRunsSelectorForbidden {
	return &GetAPIV2CiRunsSelectorForbidden{}
}

/* GetAPIV2CiRunsSelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPIV2CiRunsSelectorForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2CiRunsSelectorForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/ci-runs/{selector}][%d] getApiV2CiRunsSelectorForbidden  %+v", 403, o.Payload)
}
func (o *GetAPIV2CiRunsSelectorForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2CiRunsSelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2CiRunsSelectorNotFound creates a GetAPIV2CiRunsSelectorNotFound with default headers values
func NewGetAPIV2CiRunsSelectorNotFound() *GetAPIV2CiRunsSelectorNotFound {
	return &GetAPIV2CiRunsSelectorNotFound{}
}

/* GetAPIV2CiRunsSelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIV2CiRunsSelectorNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2CiRunsSelectorNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/ci-runs/{selector}][%d] getApiV2CiRunsSelectorNotFound  %+v", 404, o.Payload)
}
func (o *GetAPIV2CiRunsSelectorNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2CiRunsSelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2CiRunsSelectorProxyAuthenticationRequired creates a GetAPIV2CiRunsSelectorProxyAuthenticationRequired with default headers values
func NewGetAPIV2CiRunsSelectorProxyAuthenticationRequired() *GetAPIV2CiRunsSelectorProxyAuthenticationRequired {
	return &GetAPIV2CiRunsSelectorProxyAuthenticationRequired{}
}

/* GetAPIV2CiRunsSelectorProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPIV2CiRunsSelectorProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2CiRunsSelectorProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/v2/ci-runs/{selector}][%d] getApiV2CiRunsSelectorProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPIV2CiRunsSelectorProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2CiRunsSelectorProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2CiRunsSelectorConflict creates a GetAPIV2CiRunsSelectorConflict with default headers values
func NewGetAPIV2CiRunsSelectorConflict() *GetAPIV2CiRunsSelectorConflict {
	return &GetAPIV2CiRunsSelectorConflict{}
}

/* GetAPIV2CiRunsSelectorConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPIV2CiRunsSelectorConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2CiRunsSelectorConflict) Error() string {
	return fmt.Sprintf("[GET /api/v2/ci-runs/{selector}][%d] getApiV2CiRunsSelectorConflict  %+v", 409, o.Payload)
}
func (o *GetAPIV2CiRunsSelectorConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2CiRunsSelectorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2CiRunsSelectorInternalServerError creates a GetAPIV2CiRunsSelectorInternalServerError with default headers values
func NewGetAPIV2CiRunsSelectorInternalServerError() *GetAPIV2CiRunsSelectorInternalServerError {
	return &GetAPIV2CiRunsSelectorInternalServerError{}
}

/* GetAPIV2CiRunsSelectorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPIV2CiRunsSelectorInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2CiRunsSelectorInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/ci-runs/{selector}][%d] getApiV2CiRunsSelectorInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPIV2CiRunsSelectorInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2CiRunsSelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}