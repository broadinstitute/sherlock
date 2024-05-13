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

// GetAPIRoleAssignmentsV3RoleSelectorUserSelectorReader is a Reader for the GetAPIRoleAssignmentsV3RoleSelectorUserSelector structure.
type GetAPIRoleAssignmentsV3RoleSelectorUserSelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIRoleAssignmentsV3RoleSelectorUserSelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIRoleAssignmentsV3RoleSelectorUserSelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIRoleAssignmentsV3RoleSelectorUserSelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIRoleAssignmentsV3RoleSelectorUserSelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIRoleAssignmentsV3RoleSelectorUserSelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPIRoleAssignmentsV3RoleSelectorUserSelectorProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPIRoleAssignmentsV3RoleSelectorUserSelectorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIRoleAssignmentsV3RoleSelectorUserSelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIRoleAssignmentsV3RoleSelectorUserSelectorOK creates a GetAPIRoleAssignmentsV3RoleSelectorUserSelectorOK with default headers values
func NewGetAPIRoleAssignmentsV3RoleSelectorUserSelectorOK() *GetAPIRoleAssignmentsV3RoleSelectorUserSelectorOK {
	return &GetAPIRoleAssignmentsV3RoleSelectorUserSelectorOK{}
}

/* GetAPIRoleAssignmentsV3RoleSelectorUserSelectorOK describes a response with status code 200, with default header values.

OK
*/
type GetAPIRoleAssignmentsV3RoleSelectorUserSelectorOK struct {
	Payload *models.SherlockRoleAssignmentV3
}

func (o *GetAPIRoleAssignmentsV3RoleSelectorUserSelectorOK) Error() string {
	return fmt.Sprintf("[GET /api/role-assignments/v3/{role-selector}/{user-selector}][%d] getApiRoleAssignmentsV3RoleSelectorUserSelectorOK  %+v", 200, o.Payload)
}
func (o *GetAPIRoleAssignmentsV3RoleSelectorUserSelectorOK) GetPayload() *models.SherlockRoleAssignmentV3 {
	return o.Payload
}

func (o *GetAPIRoleAssignmentsV3RoleSelectorUserSelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SherlockRoleAssignmentV3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIRoleAssignmentsV3RoleSelectorUserSelectorBadRequest creates a GetAPIRoleAssignmentsV3RoleSelectorUserSelectorBadRequest with default headers values
func NewGetAPIRoleAssignmentsV3RoleSelectorUserSelectorBadRequest() *GetAPIRoleAssignmentsV3RoleSelectorUserSelectorBadRequest {
	return &GetAPIRoleAssignmentsV3RoleSelectorUserSelectorBadRequest{}
}

/* GetAPIRoleAssignmentsV3RoleSelectorUserSelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIRoleAssignmentsV3RoleSelectorUserSelectorBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIRoleAssignmentsV3RoleSelectorUserSelectorBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/role-assignments/v3/{role-selector}/{user-selector}][%d] getApiRoleAssignmentsV3RoleSelectorUserSelectorBadRequest  %+v", 400, o.Payload)
}
func (o *GetAPIRoleAssignmentsV3RoleSelectorUserSelectorBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIRoleAssignmentsV3RoleSelectorUserSelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIRoleAssignmentsV3RoleSelectorUserSelectorForbidden creates a GetAPIRoleAssignmentsV3RoleSelectorUserSelectorForbidden with default headers values
func NewGetAPIRoleAssignmentsV3RoleSelectorUserSelectorForbidden() *GetAPIRoleAssignmentsV3RoleSelectorUserSelectorForbidden {
	return &GetAPIRoleAssignmentsV3RoleSelectorUserSelectorForbidden{}
}

/* GetAPIRoleAssignmentsV3RoleSelectorUserSelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPIRoleAssignmentsV3RoleSelectorUserSelectorForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIRoleAssignmentsV3RoleSelectorUserSelectorForbidden) Error() string {
	return fmt.Sprintf("[GET /api/role-assignments/v3/{role-selector}/{user-selector}][%d] getApiRoleAssignmentsV3RoleSelectorUserSelectorForbidden  %+v", 403, o.Payload)
}
func (o *GetAPIRoleAssignmentsV3RoleSelectorUserSelectorForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIRoleAssignmentsV3RoleSelectorUserSelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIRoleAssignmentsV3RoleSelectorUserSelectorNotFound creates a GetAPIRoleAssignmentsV3RoleSelectorUserSelectorNotFound with default headers values
func NewGetAPIRoleAssignmentsV3RoleSelectorUserSelectorNotFound() *GetAPIRoleAssignmentsV3RoleSelectorUserSelectorNotFound {
	return &GetAPIRoleAssignmentsV3RoleSelectorUserSelectorNotFound{}
}

/* GetAPIRoleAssignmentsV3RoleSelectorUserSelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIRoleAssignmentsV3RoleSelectorUserSelectorNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIRoleAssignmentsV3RoleSelectorUserSelectorNotFound) Error() string {
	return fmt.Sprintf("[GET /api/role-assignments/v3/{role-selector}/{user-selector}][%d] getApiRoleAssignmentsV3RoleSelectorUserSelectorNotFound  %+v", 404, o.Payload)
}
func (o *GetAPIRoleAssignmentsV3RoleSelectorUserSelectorNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIRoleAssignmentsV3RoleSelectorUserSelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIRoleAssignmentsV3RoleSelectorUserSelectorProxyAuthenticationRequired creates a GetAPIRoleAssignmentsV3RoleSelectorUserSelectorProxyAuthenticationRequired with default headers values
func NewGetAPIRoleAssignmentsV3RoleSelectorUserSelectorProxyAuthenticationRequired() *GetAPIRoleAssignmentsV3RoleSelectorUserSelectorProxyAuthenticationRequired {
	return &GetAPIRoleAssignmentsV3RoleSelectorUserSelectorProxyAuthenticationRequired{}
}

/* GetAPIRoleAssignmentsV3RoleSelectorUserSelectorProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPIRoleAssignmentsV3RoleSelectorUserSelectorProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIRoleAssignmentsV3RoleSelectorUserSelectorProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/role-assignments/v3/{role-selector}/{user-selector}][%d] getApiRoleAssignmentsV3RoleSelectorUserSelectorProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPIRoleAssignmentsV3RoleSelectorUserSelectorProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIRoleAssignmentsV3RoleSelectorUserSelectorProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIRoleAssignmentsV3RoleSelectorUserSelectorConflict creates a GetAPIRoleAssignmentsV3RoleSelectorUserSelectorConflict with default headers values
func NewGetAPIRoleAssignmentsV3RoleSelectorUserSelectorConflict() *GetAPIRoleAssignmentsV3RoleSelectorUserSelectorConflict {
	return &GetAPIRoleAssignmentsV3RoleSelectorUserSelectorConflict{}
}

/* GetAPIRoleAssignmentsV3RoleSelectorUserSelectorConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPIRoleAssignmentsV3RoleSelectorUserSelectorConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIRoleAssignmentsV3RoleSelectorUserSelectorConflict) Error() string {
	return fmt.Sprintf("[GET /api/role-assignments/v3/{role-selector}/{user-selector}][%d] getApiRoleAssignmentsV3RoleSelectorUserSelectorConflict  %+v", 409, o.Payload)
}
func (o *GetAPIRoleAssignmentsV3RoleSelectorUserSelectorConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIRoleAssignmentsV3RoleSelectorUserSelectorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIRoleAssignmentsV3RoleSelectorUserSelectorInternalServerError creates a GetAPIRoleAssignmentsV3RoleSelectorUserSelectorInternalServerError with default headers values
func NewGetAPIRoleAssignmentsV3RoleSelectorUserSelectorInternalServerError() *GetAPIRoleAssignmentsV3RoleSelectorUserSelectorInternalServerError {
	return &GetAPIRoleAssignmentsV3RoleSelectorUserSelectorInternalServerError{}
}

/* GetAPIRoleAssignmentsV3RoleSelectorUserSelectorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPIRoleAssignmentsV3RoleSelectorUserSelectorInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIRoleAssignmentsV3RoleSelectorUserSelectorInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/role-assignments/v3/{role-selector}/{user-selector}][%d] getApiRoleAssignmentsV3RoleSelectorUserSelectorInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPIRoleAssignmentsV3RoleSelectorUserSelectorInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIRoleAssignmentsV3RoleSelectorUserSelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}