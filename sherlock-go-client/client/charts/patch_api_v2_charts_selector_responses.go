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

// PatchAPIV2ChartsSelectorReader is a Reader for the PatchAPIV2ChartsSelector structure.
type PatchAPIV2ChartsSelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPIV2ChartsSelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAPIV2ChartsSelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchAPIV2ChartsSelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchAPIV2ChartsSelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchAPIV2ChartsSelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewPatchAPIV2ChartsSelectorProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchAPIV2ChartsSelectorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchAPIV2ChartsSelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchAPIV2ChartsSelectorOK creates a PatchAPIV2ChartsSelectorOK with default headers values
func NewPatchAPIV2ChartsSelectorOK() *PatchAPIV2ChartsSelectorOK {
	return &PatchAPIV2ChartsSelectorOK{}
}

/* PatchAPIV2ChartsSelectorOK describes a response with status code 200, with default header values.

OK
*/
type PatchAPIV2ChartsSelectorOK struct {
	Payload *models.V2controllersChart
}

func (o *PatchAPIV2ChartsSelectorOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/charts/{selector}][%d] patchApiV2ChartsSelectorOK  %+v", 200, o.Payload)
}
func (o *PatchAPIV2ChartsSelectorOK) GetPayload() *models.V2controllersChart {
	return o.Payload
}

func (o *PatchAPIV2ChartsSelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V2controllersChart)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIV2ChartsSelectorBadRequest creates a PatchAPIV2ChartsSelectorBadRequest with default headers values
func NewPatchAPIV2ChartsSelectorBadRequest() *PatchAPIV2ChartsSelectorBadRequest {
	return &PatchAPIV2ChartsSelectorBadRequest{}
}

/* PatchAPIV2ChartsSelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PatchAPIV2ChartsSelectorBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIV2ChartsSelectorBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/charts/{selector}][%d] patchApiV2ChartsSelectorBadRequest  %+v", 400, o.Payload)
}
func (o *PatchAPIV2ChartsSelectorBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIV2ChartsSelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIV2ChartsSelectorForbidden creates a PatchAPIV2ChartsSelectorForbidden with default headers values
func NewPatchAPIV2ChartsSelectorForbidden() *PatchAPIV2ChartsSelectorForbidden {
	return &PatchAPIV2ChartsSelectorForbidden{}
}

/* PatchAPIV2ChartsSelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PatchAPIV2ChartsSelectorForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIV2ChartsSelectorForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/charts/{selector}][%d] patchApiV2ChartsSelectorForbidden  %+v", 403, o.Payload)
}
func (o *PatchAPIV2ChartsSelectorForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIV2ChartsSelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIV2ChartsSelectorNotFound creates a PatchAPIV2ChartsSelectorNotFound with default headers values
func NewPatchAPIV2ChartsSelectorNotFound() *PatchAPIV2ChartsSelectorNotFound {
	return &PatchAPIV2ChartsSelectorNotFound{}
}

/* PatchAPIV2ChartsSelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PatchAPIV2ChartsSelectorNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIV2ChartsSelectorNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/charts/{selector}][%d] patchApiV2ChartsSelectorNotFound  %+v", 404, o.Payload)
}
func (o *PatchAPIV2ChartsSelectorNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIV2ChartsSelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIV2ChartsSelectorProxyAuthenticationRequired creates a PatchAPIV2ChartsSelectorProxyAuthenticationRequired with default headers values
func NewPatchAPIV2ChartsSelectorProxyAuthenticationRequired() *PatchAPIV2ChartsSelectorProxyAuthenticationRequired {
	return &PatchAPIV2ChartsSelectorProxyAuthenticationRequired{}
}

/* PatchAPIV2ChartsSelectorProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type PatchAPIV2ChartsSelectorProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIV2ChartsSelectorProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/charts/{selector}][%d] patchApiV2ChartsSelectorProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *PatchAPIV2ChartsSelectorProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIV2ChartsSelectorProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIV2ChartsSelectorConflict creates a PatchAPIV2ChartsSelectorConflict with default headers values
func NewPatchAPIV2ChartsSelectorConflict() *PatchAPIV2ChartsSelectorConflict {
	return &PatchAPIV2ChartsSelectorConflict{}
}

/* PatchAPIV2ChartsSelectorConflict describes a response with status code 409, with default header values.

Conflict
*/
type PatchAPIV2ChartsSelectorConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIV2ChartsSelectorConflict) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/charts/{selector}][%d] patchApiV2ChartsSelectorConflict  %+v", 409, o.Payload)
}
func (o *PatchAPIV2ChartsSelectorConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIV2ChartsSelectorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIV2ChartsSelectorInternalServerError creates a PatchAPIV2ChartsSelectorInternalServerError with default headers values
func NewPatchAPIV2ChartsSelectorInternalServerError() *PatchAPIV2ChartsSelectorInternalServerError {
	return &PatchAPIV2ChartsSelectorInternalServerError{}
}

/* PatchAPIV2ChartsSelectorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PatchAPIV2ChartsSelectorInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIV2ChartsSelectorInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/charts/{selector}][%d] patchApiV2ChartsSelectorInternalServerError  %+v", 500, o.Payload)
}
func (o *PatchAPIV2ChartsSelectorInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIV2ChartsSelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}