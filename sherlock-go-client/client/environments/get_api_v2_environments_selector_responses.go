// Code generated by go-swagger; DO NOT EDIT.

package environments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// GetAPIV2EnvironmentsSelectorReader is a Reader for the GetAPIV2EnvironmentsSelector structure.
type GetAPIV2EnvironmentsSelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV2EnvironmentsSelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV2EnvironmentsSelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIV2EnvironmentsSelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIV2EnvironmentsSelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIV2EnvironmentsSelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPIV2EnvironmentsSelectorProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPIV2EnvironmentsSelectorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIV2EnvironmentsSelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIV2EnvironmentsSelectorOK creates a GetAPIV2EnvironmentsSelectorOK with default headers values
func NewGetAPIV2EnvironmentsSelectorOK() *GetAPIV2EnvironmentsSelectorOK {
	return &GetAPIV2EnvironmentsSelectorOK{}
}

/* GetAPIV2EnvironmentsSelectorOK describes a response with status code 200, with default header values.

OK
*/
type GetAPIV2EnvironmentsSelectorOK struct {
	Payload *models.V2controllersEnvironment
}

func (o *GetAPIV2EnvironmentsSelectorOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/environments/{selector}][%d] getApiV2EnvironmentsSelectorOK  %+v", 200, o.Payload)
}
func (o *GetAPIV2EnvironmentsSelectorOK) GetPayload() *models.V2controllersEnvironment {
	return o.Payload
}

func (o *GetAPIV2EnvironmentsSelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V2controllersEnvironment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2EnvironmentsSelectorBadRequest creates a GetAPIV2EnvironmentsSelectorBadRequest with default headers values
func NewGetAPIV2EnvironmentsSelectorBadRequest() *GetAPIV2EnvironmentsSelectorBadRequest {
	return &GetAPIV2EnvironmentsSelectorBadRequest{}
}

/* GetAPIV2EnvironmentsSelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIV2EnvironmentsSelectorBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2EnvironmentsSelectorBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/environments/{selector}][%d] getApiV2EnvironmentsSelectorBadRequest  %+v", 400, o.Payload)
}
func (o *GetAPIV2EnvironmentsSelectorBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2EnvironmentsSelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2EnvironmentsSelectorForbidden creates a GetAPIV2EnvironmentsSelectorForbidden with default headers values
func NewGetAPIV2EnvironmentsSelectorForbidden() *GetAPIV2EnvironmentsSelectorForbidden {
	return &GetAPIV2EnvironmentsSelectorForbidden{}
}

/* GetAPIV2EnvironmentsSelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPIV2EnvironmentsSelectorForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2EnvironmentsSelectorForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/environments/{selector}][%d] getApiV2EnvironmentsSelectorForbidden  %+v", 403, o.Payload)
}
func (o *GetAPIV2EnvironmentsSelectorForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2EnvironmentsSelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2EnvironmentsSelectorNotFound creates a GetAPIV2EnvironmentsSelectorNotFound with default headers values
func NewGetAPIV2EnvironmentsSelectorNotFound() *GetAPIV2EnvironmentsSelectorNotFound {
	return &GetAPIV2EnvironmentsSelectorNotFound{}
}

/* GetAPIV2EnvironmentsSelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIV2EnvironmentsSelectorNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2EnvironmentsSelectorNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/environments/{selector}][%d] getApiV2EnvironmentsSelectorNotFound  %+v", 404, o.Payload)
}
func (o *GetAPIV2EnvironmentsSelectorNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2EnvironmentsSelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2EnvironmentsSelectorProxyAuthenticationRequired creates a GetAPIV2EnvironmentsSelectorProxyAuthenticationRequired with default headers values
func NewGetAPIV2EnvironmentsSelectorProxyAuthenticationRequired() *GetAPIV2EnvironmentsSelectorProxyAuthenticationRequired {
	return &GetAPIV2EnvironmentsSelectorProxyAuthenticationRequired{}
}

/* GetAPIV2EnvironmentsSelectorProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPIV2EnvironmentsSelectorProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2EnvironmentsSelectorProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/v2/environments/{selector}][%d] getApiV2EnvironmentsSelectorProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPIV2EnvironmentsSelectorProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2EnvironmentsSelectorProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2EnvironmentsSelectorConflict creates a GetAPIV2EnvironmentsSelectorConflict with default headers values
func NewGetAPIV2EnvironmentsSelectorConflict() *GetAPIV2EnvironmentsSelectorConflict {
	return &GetAPIV2EnvironmentsSelectorConflict{}
}

/* GetAPIV2EnvironmentsSelectorConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPIV2EnvironmentsSelectorConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2EnvironmentsSelectorConflict) Error() string {
	return fmt.Sprintf("[GET /api/v2/environments/{selector}][%d] getApiV2EnvironmentsSelectorConflict  %+v", 409, o.Payload)
}
func (o *GetAPIV2EnvironmentsSelectorConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2EnvironmentsSelectorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2EnvironmentsSelectorInternalServerError creates a GetAPIV2EnvironmentsSelectorInternalServerError with default headers values
func NewGetAPIV2EnvironmentsSelectorInternalServerError() *GetAPIV2EnvironmentsSelectorInternalServerError {
	return &GetAPIV2EnvironmentsSelectorInternalServerError{}
}

/* GetAPIV2EnvironmentsSelectorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPIV2EnvironmentsSelectorInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2EnvironmentsSelectorInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/environments/{selector}][%d] getApiV2EnvironmentsSelectorInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPIV2EnvironmentsSelectorInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2EnvironmentsSelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
