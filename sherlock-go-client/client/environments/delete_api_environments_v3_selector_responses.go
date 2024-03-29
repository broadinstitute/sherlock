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

// DeleteAPIEnvironmentsV3SelectorReader is a Reader for the DeleteAPIEnvironmentsV3Selector structure.
type DeleteAPIEnvironmentsV3SelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIEnvironmentsV3SelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAPIEnvironmentsV3SelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteAPIEnvironmentsV3SelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAPIEnvironmentsV3SelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAPIEnvironmentsV3SelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewDeleteAPIEnvironmentsV3SelectorProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteAPIEnvironmentsV3SelectorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAPIEnvironmentsV3SelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAPIEnvironmentsV3SelectorOK creates a DeleteAPIEnvironmentsV3SelectorOK with default headers values
func NewDeleteAPIEnvironmentsV3SelectorOK() *DeleteAPIEnvironmentsV3SelectorOK {
	return &DeleteAPIEnvironmentsV3SelectorOK{}
}

/* DeleteAPIEnvironmentsV3SelectorOK describes a response with status code 200, with default header values.

OK
*/
type DeleteAPIEnvironmentsV3SelectorOK struct {
	Payload *models.SherlockEnvironmentV3
}

func (o *DeleteAPIEnvironmentsV3SelectorOK) Error() string {
	return fmt.Sprintf("[DELETE /api/environments/v3/{selector}][%d] deleteApiEnvironmentsV3SelectorOK  %+v", 200, o.Payload)
}
func (o *DeleteAPIEnvironmentsV3SelectorOK) GetPayload() *models.SherlockEnvironmentV3 {
	return o.Payload
}

func (o *DeleteAPIEnvironmentsV3SelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SherlockEnvironmentV3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIEnvironmentsV3SelectorBadRequest creates a DeleteAPIEnvironmentsV3SelectorBadRequest with default headers values
func NewDeleteAPIEnvironmentsV3SelectorBadRequest() *DeleteAPIEnvironmentsV3SelectorBadRequest {
	return &DeleteAPIEnvironmentsV3SelectorBadRequest{}
}

/* DeleteAPIEnvironmentsV3SelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteAPIEnvironmentsV3SelectorBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIEnvironmentsV3SelectorBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/environments/v3/{selector}][%d] deleteApiEnvironmentsV3SelectorBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteAPIEnvironmentsV3SelectorBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIEnvironmentsV3SelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIEnvironmentsV3SelectorForbidden creates a DeleteAPIEnvironmentsV3SelectorForbidden with default headers values
func NewDeleteAPIEnvironmentsV3SelectorForbidden() *DeleteAPIEnvironmentsV3SelectorForbidden {
	return &DeleteAPIEnvironmentsV3SelectorForbidden{}
}

/* DeleteAPIEnvironmentsV3SelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteAPIEnvironmentsV3SelectorForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIEnvironmentsV3SelectorForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/environments/v3/{selector}][%d] deleteApiEnvironmentsV3SelectorForbidden  %+v", 403, o.Payload)
}
func (o *DeleteAPIEnvironmentsV3SelectorForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIEnvironmentsV3SelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIEnvironmentsV3SelectorNotFound creates a DeleteAPIEnvironmentsV3SelectorNotFound with default headers values
func NewDeleteAPIEnvironmentsV3SelectorNotFound() *DeleteAPIEnvironmentsV3SelectorNotFound {
	return &DeleteAPIEnvironmentsV3SelectorNotFound{}
}

/* DeleteAPIEnvironmentsV3SelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteAPIEnvironmentsV3SelectorNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIEnvironmentsV3SelectorNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/environments/v3/{selector}][%d] deleteApiEnvironmentsV3SelectorNotFound  %+v", 404, o.Payload)
}
func (o *DeleteAPIEnvironmentsV3SelectorNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIEnvironmentsV3SelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIEnvironmentsV3SelectorProxyAuthenticationRequired creates a DeleteAPIEnvironmentsV3SelectorProxyAuthenticationRequired with default headers values
func NewDeleteAPIEnvironmentsV3SelectorProxyAuthenticationRequired() *DeleteAPIEnvironmentsV3SelectorProxyAuthenticationRequired {
	return &DeleteAPIEnvironmentsV3SelectorProxyAuthenticationRequired{}
}

/* DeleteAPIEnvironmentsV3SelectorProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type DeleteAPIEnvironmentsV3SelectorProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIEnvironmentsV3SelectorProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[DELETE /api/environments/v3/{selector}][%d] deleteApiEnvironmentsV3SelectorProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *DeleteAPIEnvironmentsV3SelectorProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIEnvironmentsV3SelectorProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIEnvironmentsV3SelectorConflict creates a DeleteAPIEnvironmentsV3SelectorConflict with default headers values
func NewDeleteAPIEnvironmentsV3SelectorConflict() *DeleteAPIEnvironmentsV3SelectorConflict {
	return &DeleteAPIEnvironmentsV3SelectorConflict{}
}

/* DeleteAPIEnvironmentsV3SelectorConflict describes a response with status code 409, with default header values.

Conflict
*/
type DeleteAPIEnvironmentsV3SelectorConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIEnvironmentsV3SelectorConflict) Error() string {
	return fmt.Sprintf("[DELETE /api/environments/v3/{selector}][%d] deleteApiEnvironmentsV3SelectorConflict  %+v", 409, o.Payload)
}
func (o *DeleteAPIEnvironmentsV3SelectorConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIEnvironmentsV3SelectorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIEnvironmentsV3SelectorInternalServerError creates a DeleteAPIEnvironmentsV3SelectorInternalServerError with default headers values
func NewDeleteAPIEnvironmentsV3SelectorInternalServerError() *DeleteAPIEnvironmentsV3SelectorInternalServerError {
	return &DeleteAPIEnvironmentsV3SelectorInternalServerError{}
}

/* DeleteAPIEnvironmentsV3SelectorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteAPIEnvironmentsV3SelectorInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIEnvironmentsV3SelectorInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/environments/v3/{selector}][%d] deleteApiEnvironmentsV3SelectorInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteAPIEnvironmentsV3SelectorInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIEnvironmentsV3SelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
