// Code generated by go-swagger; DO NOT EDIT.

package role_assignments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorReader is a Reader for the PatchAPIRoleAssignmentsV3RoleSelectorUserSelector structure.
type PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAPIRoleAssignmentsV3RoleSelectorUserSelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchAPIRoleAssignmentsV3RoleSelectorUserSelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchAPIRoleAssignmentsV3RoleSelectorUserSelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchAPIRoleAssignmentsV3RoleSelectorUserSelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewPatchAPIRoleAssignmentsV3RoleSelectorUserSelectorProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchAPIRoleAssignmentsV3RoleSelectorUserSelectorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchAPIRoleAssignmentsV3RoleSelectorUserSelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchAPIRoleAssignmentsV3RoleSelectorUserSelectorOK creates a PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorOK with default headers values
func NewPatchAPIRoleAssignmentsV3RoleSelectorUserSelectorOK() *PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorOK {
	return &PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorOK{}
}

/* PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorOK describes a response with status code 200, with default header values.

OK
*/
type PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorOK struct {
	Payload *models.SherlockRoleAssignmentV3
}

func (o *PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorOK) Error() string {
	return fmt.Sprintf("[PATCH /api/role-assignments/v3/{role-selector}/{user-selector}][%d] patchApiRoleAssignmentsV3RoleSelectorUserSelectorOK  %+v", 200, o.Payload)
}
func (o *PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorOK) GetPayload() *models.SherlockRoleAssignmentV3 {
	return o.Payload
}

func (o *PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SherlockRoleAssignmentV3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIRoleAssignmentsV3RoleSelectorUserSelectorBadRequest creates a PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorBadRequest with default headers values
func NewPatchAPIRoleAssignmentsV3RoleSelectorUserSelectorBadRequest() *PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorBadRequest {
	return &PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorBadRequest{}
}

/* PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/role-assignments/v3/{role-selector}/{user-selector}][%d] patchApiRoleAssignmentsV3RoleSelectorUserSelectorBadRequest  %+v", 400, o.Payload)
}
func (o *PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIRoleAssignmentsV3RoleSelectorUserSelectorForbidden creates a PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorForbidden with default headers values
func NewPatchAPIRoleAssignmentsV3RoleSelectorUserSelectorForbidden() *PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorForbidden {
	return &PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorForbidden{}
}

/* PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/role-assignments/v3/{role-selector}/{user-selector}][%d] patchApiRoleAssignmentsV3RoleSelectorUserSelectorForbidden  %+v", 403, o.Payload)
}
func (o *PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIRoleAssignmentsV3RoleSelectorUserSelectorNotFound creates a PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorNotFound with default headers values
func NewPatchAPIRoleAssignmentsV3RoleSelectorUserSelectorNotFound() *PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorNotFound {
	return &PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorNotFound{}
}

/* PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/role-assignments/v3/{role-selector}/{user-selector}][%d] patchApiRoleAssignmentsV3RoleSelectorUserSelectorNotFound  %+v", 404, o.Payload)
}
func (o *PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIRoleAssignmentsV3RoleSelectorUserSelectorProxyAuthenticationRequired creates a PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorProxyAuthenticationRequired with default headers values
func NewPatchAPIRoleAssignmentsV3RoleSelectorUserSelectorProxyAuthenticationRequired() *PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorProxyAuthenticationRequired {
	return &PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorProxyAuthenticationRequired{}
}

/* PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[PATCH /api/role-assignments/v3/{role-selector}/{user-selector}][%d] patchApiRoleAssignmentsV3RoleSelectorUserSelectorProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIRoleAssignmentsV3RoleSelectorUserSelectorConflict creates a PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorConflict with default headers values
func NewPatchAPIRoleAssignmentsV3RoleSelectorUserSelectorConflict() *PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorConflict {
	return &PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorConflict{}
}

/* PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorConflict describes a response with status code 409, with default header values.

Conflict
*/
type PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorConflict) Error() string {
	return fmt.Sprintf("[PATCH /api/role-assignments/v3/{role-selector}/{user-selector}][%d] patchApiRoleAssignmentsV3RoleSelectorUserSelectorConflict  %+v", 409, o.Payload)
}
func (o *PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIRoleAssignmentsV3RoleSelectorUserSelectorInternalServerError creates a PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorInternalServerError with default headers values
func NewPatchAPIRoleAssignmentsV3RoleSelectorUserSelectorInternalServerError() *PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorInternalServerError {
	return &PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorInternalServerError{}
}

/* PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/role-assignments/v3/{role-selector}/{user-selector}][%d] patchApiRoleAssignmentsV3RoleSelectorUserSelectorInternalServerError  %+v", 500, o.Payload)
}
func (o *PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIRoleAssignmentsV3RoleSelectorUserSelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
