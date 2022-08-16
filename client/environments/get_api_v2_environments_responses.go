// Code generated by go-swagger; DO NOT EDIT.

package environments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/client/models"
)

// GetAPIV2EnvironmentsReader is a Reader for the GetAPIV2Environments structure.
type GetAPIV2EnvironmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV2EnvironmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV2EnvironmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIV2EnvironmentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIV2EnvironmentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIV2EnvironmentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPIV2EnvironmentsProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPIV2EnvironmentsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIV2EnvironmentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIV2EnvironmentsOK creates a GetAPIV2EnvironmentsOK with default headers values
func NewGetAPIV2EnvironmentsOK() *GetAPIV2EnvironmentsOK {
	return &GetAPIV2EnvironmentsOK{}
}

/* GetAPIV2EnvironmentsOK describes a response with status code 200, with default header values.

OK
*/
type GetAPIV2EnvironmentsOK struct {
	Payload []*models.V2controllersEnvironment
}

func (o *GetAPIV2EnvironmentsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/environments][%d] getApiV2EnvironmentsOK  %+v", 200, o.Payload)
}
func (o *GetAPIV2EnvironmentsOK) GetPayload() []*models.V2controllersEnvironment {
	return o.Payload
}

func (o *GetAPIV2EnvironmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2EnvironmentsBadRequest creates a GetAPIV2EnvironmentsBadRequest with default headers values
func NewGetAPIV2EnvironmentsBadRequest() *GetAPIV2EnvironmentsBadRequest {
	return &GetAPIV2EnvironmentsBadRequest{}
}

/* GetAPIV2EnvironmentsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIV2EnvironmentsBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2EnvironmentsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/environments][%d] getApiV2EnvironmentsBadRequest  %+v", 400, o.Payload)
}
func (o *GetAPIV2EnvironmentsBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2EnvironmentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2EnvironmentsForbidden creates a GetAPIV2EnvironmentsForbidden with default headers values
func NewGetAPIV2EnvironmentsForbidden() *GetAPIV2EnvironmentsForbidden {
	return &GetAPIV2EnvironmentsForbidden{}
}

/* GetAPIV2EnvironmentsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPIV2EnvironmentsForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2EnvironmentsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/environments][%d] getApiV2EnvironmentsForbidden  %+v", 403, o.Payload)
}
func (o *GetAPIV2EnvironmentsForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2EnvironmentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2EnvironmentsNotFound creates a GetAPIV2EnvironmentsNotFound with default headers values
func NewGetAPIV2EnvironmentsNotFound() *GetAPIV2EnvironmentsNotFound {
	return &GetAPIV2EnvironmentsNotFound{}
}

/* GetAPIV2EnvironmentsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIV2EnvironmentsNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2EnvironmentsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/environments][%d] getApiV2EnvironmentsNotFound  %+v", 404, o.Payload)
}
func (o *GetAPIV2EnvironmentsNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2EnvironmentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2EnvironmentsProxyAuthenticationRequired creates a GetAPIV2EnvironmentsProxyAuthenticationRequired with default headers values
func NewGetAPIV2EnvironmentsProxyAuthenticationRequired() *GetAPIV2EnvironmentsProxyAuthenticationRequired {
	return &GetAPIV2EnvironmentsProxyAuthenticationRequired{}
}

/* GetAPIV2EnvironmentsProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPIV2EnvironmentsProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2EnvironmentsProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/v2/environments][%d] getApiV2EnvironmentsProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPIV2EnvironmentsProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2EnvironmentsProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2EnvironmentsConflict creates a GetAPIV2EnvironmentsConflict with default headers values
func NewGetAPIV2EnvironmentsConflict() *GetAPIV2EnvironmentsConflict {
	return &GetAPIV2EnvironmentsConflict{}
}

/* GetAPIV2EnvironmentsConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPIV2EnvironmentsConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2EnvironmentsConflict) Error() string {
	return fmt.Sprintf("[GET /api/v2/environments][%d] getApiV2EnvironmentsConflict  %+v", 409, o.Payload)
}
func (o *GetAPIV2EnvironmentsConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2EnvironmentsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2EnvironmentsInternalServerError creates a GetAPIV2EnvironmentsInternalServerError with default headers values
func NewGetAPIV2EnvironmentsInternalServerError() *GetAPIV2EnvironmentsInternalServerError {
	return &GetAPIV2EnvironmentsInternalServerError{}
}

/* GetAPIV2EnvironmentsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPIV2EnvironmentsInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2EnvironmentsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/environments][%d] getApiV2EnvironmentsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPIV2EnvironmentsInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2EnvironmentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
