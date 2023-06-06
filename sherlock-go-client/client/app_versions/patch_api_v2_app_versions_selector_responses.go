// Code generated by go-swagger; DO NOT EDIT.

package app_versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// PatchAPIV2AppVersionsSelectorReader is a Reader for the PatchAPIV2AppVersionsSelector structure.
type PatchAPIV2AppVersionsSelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPIV2AppVersionsSelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAPIV2AppVersionsSelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchAPIV2AppVersionsSelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchAPIV2AppVersionsSelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchAPIV2AppVersionsSelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewPatchAPIV2AppVersionsSelectorProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchAPIV2AppVersionsSelectorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchAPIV2AppVersionsSelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchAPIV2AppVersionsSelectorOK creates a PatchAPIV2AppVersionsSelectorOK with default headers values
func NewPatchAPIV2AppVersionsSelectorOK() *PatchAPIV2AppVersionsSelectorOK {
	return &PatchAPIV2AppVersionsSelectorOK{}
}

/* PatchAPIV2AppVersionsSelectorOK describes a response with status code 200, with default header values.

OK
*/
type PatchAPIV2AppVersionsSelectorOK struct {
	Payload *models.V2controllersAppVersion
}

func (o *PatchAPIV2AppVersionsSelectorOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/app-versions/{selector}][%d] patchApiV2AppVersionsSelectorOK  %+v", 200, o.Payload)
}
func (o *PatchAPIV2AppVersionsSelectorOK) GetPayload() *models.V2controllersAppVersion {
	return o.Payload
}

func (o *PatchAPIV2AppVersionsSelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V2controllersAppVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIV2AppVersionsSelectorBadRequest creates a PatchAPIV2AppVersionsSelectorBadRequest with default headers values
func NewPatchAPIV2AppVersionsSelectorBadRequest() *PatchAPIV2AppVersionsSelectorBadRequest {
	return &PatchAPIV2AppVersionsSelectorBadRequest{}
}

/* PatchAPIV2AppVersionsSelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PatchAPIV2AppVersionsSelectorBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIV2AppVersionsSelectorBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/app-versions/{selector}][%d] patchApiV2AppVersionsSelectorBadRequest  %+v", 400, o.Payload)
}
func (o *PatchAPIV2AppVersionsSelectorBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIV2AppVersionsSelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIV2AppVersionsSelectorForbidden creates a PatchAPIV2AppVersionsSelectorForbidden with default headers values
func NewPatchAPIV2AppVersionsSelectorForbidden() *PatchAPIV2AppVersionsSelectorForbidden {
	return &PatchAPIV2AppVersionsSelectorForbidden{}
}

/* PatchAPIV2AppVersionsSelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PatchAPIV2AppVersionsSelectorForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIV2AppVersionsSelectorForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/app-versions/{selector}][%d] patchApiV2AppVersionsSelectorForbidden  %+v", 403, o.Payload)
}
func (o *PatchAPIV2AppVersionsSelectorForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIV2AppVersionsSelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIV2AppVersionsSelectorNotFound creates a PatchAPIV2AppVersionsSelectorNotFound with default headers values
func NewPatchAPIV2AppVersionsSelectorNotFound() *PatchAPIV2AppVersionsSelectorNotFound {
	return &PatchAPIV2AppVersionsSelectorNotFound{}
}

/* PatchAPIV2AppVersionsSelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PatchAPIV2AppVersionsSelectorNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIV2AppVersionsSelectorNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/app-versions/{selector}][%d] patchApiV2AppVersionsSelectorNotFound  %+v", 404, o.Payload)
}
func (o *PatchAPIV2AppVersionsSelectorNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIV2AppVersionsSelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIV2AppVersionsSelectorProxyAuthenticationRequired creates a PatchAPIV2AppVersionsSelectorProxyAuthenticationRequired with default headers values
func NewPatchAPIV2AppVersionsSelectorProxyAuthenticationRequired() *PatchAPIV2AppVersionsSelectorProxyAuthenticationRequired {
	return &PatchAPIV2AppVersionsSelectorProxyAuthenticationRequired{}
}

/* PatchAPIV2AppVersionsSelectorProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type PatchAPIV2AppVersionsSelectorProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIV2AppVersionsSelectorProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/app-versions/{selector}][%d] patchApiV2AppVersionsSelectorProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *PatchAPIV2AppVersionsSelectorProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIV2AppVersionsSelectorProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIV2AppVersionsSelectorConflict creates a PatchAPIV2AppVersionsSelectorConflict with default headers values
func NewPatchAPIV2AppVersionsSelectorConflict() *PatchAPIV2AppVersionsSelectorConflict {
	return &PatchAPIV2AppVersionsSelectorConflict{}
}

/* PatchAPIV2AppVersionsSelectorConflict describes a response with status code 409, with default header values.

Conflict
*/
type PatchAPIV2AppVersionsSelectorConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIV2AppVersionsSelectorConflict) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/app-versions/{selector}][%d] patchApiV2AppVersionsSelectorConflict  %+v", 409, o.Payload)
}
func (o *PatchAPIV2AppVersionsSelectorConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIV2AppVersionsSelectorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIV2AppVersionsSelectorInternalServerError creates a PatchAPIV2AppVersionsSelectorInternalServerError with default headers values
func NewPatchAPIV2AppVersionsSelectorInternalServerError() *PatchAPIV2AppVersionsSelectorInternalServerError {
	return &PatchAPIV2AppVersionsSelectorInternalServerError{}
}

/* PatchAPIV2AppVersionsSelectorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PatchAPIV2AppVersionsSelectorInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIV2AppVersionsSelectorInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/app-versions/{selector}][%d] patchApiV2AppVersionsSelectorInternalServerError  %+v", 500, o.Payload)
}
func (o *PatchAPIV2AppVersionsSelectorInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIV2AppVersionsSelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
