// Code generated by go-swagger; DO NOT EDIT.

package charts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// DeleteAPIV2ChartsSelectorReader is a Reader for the DeleteAPIV2ChartsSelector structure.
type DeleteAPIV2ChartsSelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIV2ChartsSelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAPIV2ChartsSelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteAPIV2ChartsSelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAPIV2ChartsSelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAPIV2ChartsSelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewDeleteAPIV2ChartsSelectorProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteAPIV2ChartsSelectorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAPIV2ChartsSelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAPIV2ChartsSelectorOK creates a DeleteAPIV2ChartsSelectorOK with default headers values
func NewDeleteAPIV2ChartsSelectorOK() *DeleteAPIV2ChartsSelectorOK {
	return &DeleteAPIV2ChartsSelectorOK{}
}

/* DeleteAPIV2ChartsSelectorOK describes a response with status code 200, with default header values.

OK
*/
type DeleteAPIV2ChartsSelectorOK struct {
	Payload *models.V2controllersChart
}

func (o *DeleteAPIV2ChartsSelectorOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/charts/{selector}][%d] deleteApiV2ChartsSelectorOK  %+v", 200, o.Payload)
}
func (o *DeleteAPIV2ChartsSelectorOK) GetPayload() *models.V2controllersChart {
	return o.Payload
}

func (o *DeleteAPIV2ChartsSelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V2controllersChart)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIV2ChartsSelectorBadRequest creates a DeleteAPIV2ChartsSelectorBadRequest with default headers values
func NewDeleteAPIV2ChartsSelectorBadRequest() *DeleteAPIV2ChartsSelectorBadRequest {
	return &DeleteAPIV2ChartsSelectorBadRequest{}
}

/* DeleteAPIV2ChartsSelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteAPIV2ChartsSelectorBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIV2ChartsSelectorBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/charts/{selector}][%d] deleteApiV2ChartsSelectorBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteAPIV2ChartsSelectorBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIV2ChartsSelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIV2ChartsSelectorForbidden creates a DeleteAPIV2ChartsSelectorForbidden with default headers values
func NewDeleteAPIV2ChartsSelectorForbidden() *DeleteAPIV2ChartsSelectorForbidden {
	return &DeleteAPIV2ChartsSelectorForbidden{}
}

/* DeleteAPIV2ChartsSelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteAPIV2ChartsSelectorForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIV2ChartsSelectorForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/charts/{selector}][%d] deleteApiV2ChartsSelectorForbidden  %+v", 403, o.Payload)
}
func (o *DeleteAPIV2ChartsSelectorForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIV2ChartsSelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIV2ChartsSelectorNotFound creates a DeleteAPIV2ChartsSelectorNotFound with default headers values
func NewDeleteAPIV2ChartsSelectorNotFound() *DeleteAPIV2ChartsSelectorNotFound {
	return &DeleteAPIV2ChartsSelectorNotFound{}
}

/* DeleteAPIV2ChartsSelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteAPIV2ChartsSelectorNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIV2ChartsSelectorNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/charts/{selector}][%d] deleteApiV2ChartsSelectorNotFound  %+v", 404, o.Payload)
}
func (o *DeleteAPIV2ChartsSelectorNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIV2ChartsSelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIV2ChartsSelectorProxyAuthenticationRequired creates a DeleteAPIV2ChartsSelectorProxyAuthenticationRequired with default headers values
func NewDeleteAPIV2ChartsSelectorProxyAuthenticationRequired() *DeleteAPIV2ChartsSelectorProxyAuthenticationRequired {
	return &DeleteAPIV2ChartsSelectorProxyAuthenticationRequired{}
}

/* DeleteAPIV2ChartsSelectorProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type DeleteAPIV2ChartsSelectorProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIV2ChartsSelectorProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/charts/{selector}][%d] deleteApiV2ChartsSelectorProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *DeleteAPIV2ChartsSelectorProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIV2ChartsSelectorProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIV2ChartsSelectorConflict creates a DeleteAPIV2ChartsSelectorConflict with default headers values
func NewDeleteAPIV2ChartsSelectorConflict() *DeleteAPIV2ChartsSelectorConflict {
	return &DeleteAPIV2ChartsSelectorConflict{}
}

/* DeleteAPIV2ChartsSelectorConflict describes a response with status code 409, with default header values.

Conflict
*/
type DeleteAPIV2ChartsSelectorConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIV2ChartsSelectorConflict) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/charts/{selector}][%d] deleteApiV2ChartsSelectorConflict  %+v", 409, o.Payload)
}
func (o *DeleteAPIV2ChartsSelectorConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIV2ChartsSelectorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIV2ChartsSelectorInternalServerError creates a DeleteAPIV2ChartsSelectorInternalServerError with default headers values
func NewDeleteAPIV2ChartsSelectorInternalServerError() *DeleteAPIV2ChartsSelectorInternalServerError {
	return &DeleteAPIV2ChartsSelectorInternalServerError{}
}

/* DeleteAPIV2ChartsSelectorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteAPIV2ChartsSelectorInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIV2ChartsSelectorInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/charts/{selector}][%d] deleteApiV2ChartsSelectorInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteAPIV2ChartsSelectorInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIV2ChartsSelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}