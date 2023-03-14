// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/clients/go/client/models"
)

// PatchAPIV2UsersSelectorReader is a Reader for the PatchAPIV2UsersSelector structure.
type PatchAPIV2UsersSelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPIV2UsersSelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAPIV2UsersSelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchAPIV2UsersSelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchAPIV2UsersSelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchAPIV2UsersSelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewPatchAPIV2UsersSelectorProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchAPIV2UsersSelectorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchAPIV2UsersSelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchAPIV2UsersSelectorOK creates a PatchAPIV2UsersSelectorOK with default headers values
func NewPatchAPIV2UsersSelectorOK() *PatchAPIV2UsersSelectorOK {
	return &PatchAPIV2UsersSelectorOK{}
}

/* PatchAPIV2UsersSelectorOK describes a response with status code 200, with default header values.

OK
*/
type PatchAPIV2UsersSelectorOK struct {
	Payload *models.V2controllersUser
}

func (o *PatchAPIV2UsersSelectorOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{selector}][%d] patchApiV2UsersSelectorOK  %+v", 200, o.Payload)
}
func (o *PatchAPIV2UsersSelectorOK) GetPayload() *models.V2controllersUser {
	return o.Payload
}

func (o *PatchAPIV2UsersSelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V2controllersUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIV2UsersSelectorBadRequest creates a PatchAPIV2UsersSelectorBadRequest with default headers values
func NewPatchAPIV2UsersSelectorBadRequest() *PatchAPIV2UsersSelectorBadRequest {
	return &PatchAPIV2UsersSelectorBadRequest{}
}

/* PatchAPIV2UsersSelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PatchAPIV2UsersSelectorBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIV2UsersSelectorBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{selector}][%d] patchApiV2UsersSelectorBadRequest  %+v", 400, o.Payload)
}
func (o *PatchAPIV2UsersSelectorBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIV2UsersSelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIV2UsersSelectorForbidden creates a PatchAPIV2UsersSelectorForbidden with default headers values
func NewPatchAPIV2UsersSelectorForbidden() *PatchAPIV2UsersSelectorForbidden {
	return &PatchAPIV2UsersSelectorForbidden{}
}

/* PatchAPIV2UsersSelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PatchAPIV2UsersSelectorForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIV2UsersSelectorForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{selector}][%d] patchApiV2UsersSelectorForbidden  %+v", 403, o.Payload)
}
func (o *PatchAPIV2UsersSelectorForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIV2UsersSelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIV2UsersSelectorNotFound creates a PatchAPIV2UsersSelectorNotFound with default headers values
func NewPatchAPIV2UsersSelectorNotFound() *PatchAPIV2UsersSelectorNotFound {
	return &PatchAPIV2UsersSelectorNotFound{}
}

/* PatchAPIV2UsersSelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PatchAPIV2UsersSelectorNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIV2UsersSelectorNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{selector}][%d] patchApiV2UsersSelectorNotFound  %+v", 404, o.Payload)
}
func (o *PatchAPIV2UsersSelectorNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIV2UsersSelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIV2UsersSelectorProxyAuthenticationRequired creates a PatchAPIV2UsersSelectorProxyAuthenticationRequired with default headers values
func NewPatchAPIV2UsersSelectorProxyAuthenticationRequired() *PatchAPIV2UsersSelectorProxyAuthenticationRequired {
	return &PatchAPIV2UsersSelectorProxyAuthenticationRequired{}
}

/* PatchAPIV2UsersSelectorProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type PatchAPIV2UsersSelectorProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIV2UsersSelectorProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{selector}][%d] patchApiV2UsersSelectorProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *PatchAPIV2UsersSelectorProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIV2UsersSelectorProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIV2UsersSelectorConflict creates a PatchAPIV2UsersSelectorConflict with default headers values
func NewPatchAPIV2UsersSelectorConflict() *PatchAPIV2UsersSelectorConflict {
	return &PatchAPIV2UsersSelectorConflict{}
}

/* PatchAPIV2UsersSelectorConflict describes a response with status code 409, with default header values.

Conflict
*/
type PatchAPIV2UsersSelectorConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIV2UsersSelectorConflict) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{selector}][%d] patchApiV2UsersSelectorConflict  %+v", 409, o.Payload)
}
func (o *PatchAPIV2UsersSelectorConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIV2UsersSelectorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIV2UsersSelectorInternalServerError creates a PatchAPIV2UsersSelectorInternalServerError with default headers values
func NewPatchAPIV2UsersSelectorInternalServerError() *PatchAPIV2UsersSelectorInternalServerError {
	return &PatchAPIV2UsersSelectorInternalServerError{}
}

/* PatchAPIV2UsersSelectorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PatchAPIV2UsersSelectorInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIV2UsersSelectorInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{selector}][%d] patchApiV2UsersSelectorInternalServerError  %+v", 500, o.Payload)
}
func (o *PatchAPIV2UsersSelectorInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIV2UsersSelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
