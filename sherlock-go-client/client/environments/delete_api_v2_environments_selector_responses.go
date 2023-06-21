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

// DeleteAPIV2EnvironmentsSelectorReader is a Reader for the DeleteAPIV2EnvironmentsSelector structure.
type DeleteAPIV2EnvironmentsSelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIV2EnvironmentsSelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAPIV2EnvironmentsSelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteAPIV2EnvironmentsSelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAPIV2EnvironmentsSelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAPIV2EnvironmentsSelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewDeleteAPIV2EnvironmentsSelectorProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteAPIV2EnvironmentsSelectorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAPIV2EnvironmentsSelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAPIV2EnvironmentsSelectorOK creates a DeleteAPIV2EnvironmentsSelectorOK with default headers values
func NewDeleteAPIV2EnvironmentsSelectorOK() *DeleteAPIV2EnvironmentsSelectorOK {
	return &DeleteAPIV2EnvironmentsSelectorOK{}
}

/* DeleteAPIV2EnvironmentsSelectorOK describes a response with status code 200, with default header values.

OK
*/
type DeleteAPIV2EnvironmentsSelectorOK struct {
	Payload *models.V2controllersEnvironment
}

func (o *DeleteAPIV2EnvironmentsSelectorOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/environments/{selector}][%d] deleteApiV2EnvironmentsSelectorOK  %+v", 200, o.Payload)
}
func (o *DeleteAPIV2EnvironmentsSelectorOK) GetPayload() *models.V2controllersEnvironment {
	return o.Payload
}

func (o *DeleteAPIV2EnvironmentsSelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V2controllersEnvironment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIV2EnvironmentsSelectorBadRequest creates a DeleteAPIV2EnvironmentsSelectorBadRequest with default headers values
func NewDeleteAPIV2EnvironmentsSelectorBadRequest() *DeleteAPIV2EnvironmentsSelectorBadRequest {
	return &DeleteAPIV2EnvironmentsSelectorBadRequest{}
}

/* DeleteAPIV2EnvironmentsSelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteAPIV2EnvironmentsSelectorBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIV2EnvironmentsSelectorBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/environments/{selector}][%d] deleteApiV2EnvironmentsSelectorBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteAPIV2EnvironmentsSelectorBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIV2EnvironmentsSelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIV2EnvironmentsSelectorForbidden creates a DeleteAPIV2EnvironmentsSelectorForbidden with default headers values
func NewDeleteAPIV2EnvironmentsSelectorForbidden() *DeleteAPIV2EnvironmentsSelectorForbidden {
	return &DeleteAPIV2EnvironmentsSelectorForbidden{}
}

/* DeleteAPIV2EnvironmentsSelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteAPIV2EnvironmentsSelectorForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIV2EnvironmentsSelectorForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/environments/{selector}][%d] deleteApiV2EnvironmentsSelectorForbidden  %+v", 403, o.Payload)
}
func (o *DeleteAPIV2EnvironmentsSelectorForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIV2EnvironmentsSelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIV2EnvironmentsSelectorNotFound creates a DeleteAPIV2EnvironmentsSelectorNotFound with default headers values
func NewDeleteAPIV2EnvironmentsSelectorNotFound() *DeleteAPIV2EnvironmentsSelectorNotFound {
	return &DeleteAPIV2EnvironmentsSelectorNotFound{}
}

/* DeleteAPIV2EnvironmentsSelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteAPIV2EnvironmentsSelectorNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIV2EnvironmentsSelectorNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/environments/{selector}][%d] deleteApiV2EnvironmentsSelectorNotFound  %+v", 404, o.Payload)
}
func (o *DeleteAPIV2EnvironmentsSelectorNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIV2EnvironmentsSelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIV2EnvironmentsSelectorProxyAuthenticationRequired creates a DeleteAPIV2EnvironmentsSelectorProxyAuthenticationRequired with default headers values
func NewDeleteAPIV2EnvironmentsSelectorProxyAuthenticationRequired() *DeleteAPIV2EnvironmentsSelectorProxyAuthenticationRequired {
	return &DeleteAPIV2EnvironmentsSelectorProxyAuthenticationRequired{}
}

/* DeleteAPIV2EnvironmentsSelectorProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type DeleteAPIV2EnvironmentsSelectorProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIV2EnvironmentsSelectorProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/environments/{selector}][%d] deleteApiV2EnvironmentsSelectorProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *DeleteAPIV2EnvironmentsSelectorProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIV2EnvironmentsSelectorProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIV2EnvironmentsSelectorConflict creates a DeleteAPIV2EnvironmentsSelectorConflict with default headers values
func NewDeleteAPIV2EnvironmentsSelectorConflict() *DeleteAPIV2EnvironmentsSelectorConflict {
	return &DeleteAPIV2EnvironmentsSelectorConflict{}
}

/* DeleteAPIV2EnvironmentsSelectorConflict describes a response with status code 409, with default header values.

Conflict
*/
type DeleteAPIV2EnvironmentsSelectorConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIV2EnvironmentsSelectorConflict) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/environments/{selector}][%d] deleteApiV2EnvironmentsSelectorConflict  %+v", 409, o.Payload)
}
func (o *DeleteAPIV2EnvironmentsSelectorConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIV2EnvironmentsSelectorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIV2EnvironmentsSelectorInternalServerError creates a DeleteAPIV2EnvironmentsSelectorInternalServerError with default headers values
func NewDeleteAPIV2EnvironmentsSelectorInternalServerError() *DeleteAPIV2EnvironmentsSelectorInternalServerError {
	return &DeleteAPIV2EnvironmentsSelectorInternalServerError{}
}

/* DeleteAPIV2EnvironmentsSelectorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteAPIV2EnvironmentsSelectorInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIV2EnvironmentsSelectorInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/environments/{selector}][%d] deleteApiV2EnvironmentsSelectorInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteAPIV2EnvironmentsSelectorInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIV2EnvironmentsSelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}