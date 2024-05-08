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

// GetAPIRoleAssignmentsV3RoleIDUserSelectorReader is a Reader for the GetAPIRoleAssignmentsV3RoleIDUserSelector structure.
type GetAPIRoleAssignmentsV3RoleIDUserSelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIRoleAssignmentsV3RoleIDUserSelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIRoleAssignmentsV3RoleIDUserSelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIRoleAssignmentsV3RoleIDUserSelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIRoleAssignmentsV3RoleIDUserSelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIRoleAssignmentsV3RoleIDUserSelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPIRoleAssignmentsV3RoleIDUserSelectorProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPIRoleAssignmentsV3RoleIDUserSelectorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIRoleAssignmentsV3RoleIDUserSelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIRoleAssignmentsV3RoleIDUserSelectorOK creates a GetAPIRoleAssignmentsV3RoleIDUserSelectorOK with default headers values
func NewGetAPIRoleAssignmentsV3RoleIDUserSelectorOK() *GetAPIRoleAssignmentsV3RoleIDUserSelectorOK {
	return &GetAPIRoleAssignmentsV3RoleIDUserSelectorOK{}
}

/* GetAPIRoleAssignmentsV3RoleIDUserSelectorOK describes a response with status code 200, with default header values.

OK
*/
type GetAPIRoleAssignmentsV3RoleIDUserSelectorOK struct {
	Payload *models.SherlockRoleAssignmentV3
}

func (o *GetAPIRoleAssignmentsV3RoleIDUserSelectorOK) Error() string {
	return fmt.Sprintf("[GET /api/role-assignments/v3/{role-id}/{user-selector}][%d] getApiRoleAssignmentsV3RoleIdUserSelectorOK  %+v", 200, o.Payload)
}
func (o *GetAPIRoleAssignmentsV3RoleIDUserSelectorOK) GetPayload() *models.SherlockRoleAssignmentV3 {
	return o.Payload
}

func (o *GetAPIRoleAssignmentsV3RoleIDUserSelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SherlockRoleAssignmentV3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIRoleAssignmentsV3RoleIDUserSelectorBadRequest creates a GetAPIRoleAssignmentsV3RoleIDUserSelectorBadRequest with default headers values
func NewGetAPIRoleAssignmentsV3RoleIDUserSelectorBadRequest() *GetAPIRoleAssignmentsV3RoleIDUserSelectorBadRequest {
	return &GetAPIRoleAssignmentsV3RoleIDUserSelectorBadRequest{}
}

/* GetAPIRoleAssignmentsV3RoleIDUserSelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIRoleAssignmentsV3RoleIDUserSelectorBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIRoleAssignmentsV3RoleIDUserSelectorBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/role-assignments/v3/{role-id}/{user-selector}][%d] getApiRoleAssignmentsV3RoleIdUserSelectorBadRequest  %+v", 400, o.Payload)
}
func (o *GetAPIRoleAssignmentsV3RoleIDUserSelectorBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIRoleAssignmentsV3RoleIDUserSelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIRoleAssignmentsV3RoleIDUserSelectorForbidden creates a GetAPIRoleAssignmentsV3RoleIDUserSelectorForbidden with default headers values
func NewGetAPIRoleAssignmentsV3RoleIDUserSelectorForbidden() *GetAPIRoleAssignmentsV3RoleIDUserSelectorForbidden {
	return &GetAPIRoleAssignmentsV3RoleIDUserSelectorForbidden{}
}

/* GetAPIRoleAssignmentsV3RoleIDUserSelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPIRoleAssignmentsV3RoleIDUserSelectorForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIRoleAssignmentsV3RoleIDUserSelectorForbidden) Error() string {
	return fmt.Sprintf("[GET /api/role-assignments/v3/{role-id}/{user-selector}][%d] getApiRoleAssignmentsV3RoleIdUserSelectorForbidden  %+v", 403, o.Payload)
}
func (o *GetAPIRoleAssignmentsV3RoleIDUserSelectorForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIRoleAssignmentsV3RoleIDUserSelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIRoleAssignmentsV3RoleIDUserSelectorNotFound creates a GetAPIRoleAssignmentsV3RoleIDUserSelectorNotFound with default headers values
func NewGetAPIRoleAssignmentsV3RoleIDUserSelectorNotFound() *GetAPIRoleAssignmentsV3RoleIDUserSelectorNotFound {
	return &GetAPIRoleAssignmentsV3RoleIDUserSelectorNotFound{}
}

/* GetAPIRoleAssignmentsV3RoleIDUserSelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIRoleAssignmentsV3RoleIDUserSelectorNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIRoleAssignmentsV3RoleIDUserSelectorNotFound) Error() string {
	return fmt.Sprintf("[GET /api/role-assignments/v3/{role-id}/{user-selector}][%d] getApiRoleAssignmentsV3RoleIdUserSelectorNotFound  %+v", 404, o.Payload)
}
func (o *GetAPIRoleAssignmentsV3RoleIDUserSelectorNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIRoleAssignmentsV3RoleIDUserSelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIRoleAssignmentsV3RoleIDUserSelectorProxyAuthenticationRequired creates a GetAPIRoleAssignmentsV3RoleIDUserSelectorProxyAuthenticationRequired with default headers values
func NewGetAPIRoleAssignmentsV3RoleIDUserSelectorProxyAuthenticationRequired() *GetAPIRoleAssignmentsV3RoleIDUserSelectorProxyAuthenticationRequired {
	return &GetAPIRoleAssignmentsV3RoleIDUserSelectorProxyAuthenticationRequired{}
}

/* GetAPIRoleAssignmentsV3RoleIDUserSelectorProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPIRoleAssignmentsV3RoleIDUserSelectorProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIRoleAssignmentsV3RoleIDUserSelectorProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/role-assignments/v3/{role-id}/{user-selector}][%d] getApiRoleAssignmentsV3RoleIdUserSelectorProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPIRoleAssignmentsV3RoleIDUserSelectorProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIRoleAssignmentsV3RoleIDUserSelectorProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIRoleAssignmentsV3RoleIDUserSelectorConflict creates a GetAPIRoleAssignmentsV3RoleIDUserSelectorConflict with default headers values
func NewGetAPIRoleAssignmentsV3RoleIDUserSelectorConflict() *GetAPIRoleAssignmentsV3RoleIDUserSelectorConflict {
	return &GetAPIRoleAssignmentsV3RoleIDUserSelectorConflict{}
}

/* GetAPIRoleAssignmentsV3RoleIDUserSelectorConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPIRoleAssignmentsV3RoleIDUserSelectorConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIRoleAssignmentsV3RoleIDUserSelectorConflict) Error() string {
	return fmt.Sprintf("[GET /api/role-assignments/v3/{role-id}/{user-selector}][%d] getApiRoleAssignmentsV3RoleIdUserSelectorConflict  %+v", 409, o.Payload)
}
func (o *GetAPIRoleAssignmentsV3RoleIDUserSelectorConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIRoleAssignmentsV3RoleIDUserSelectorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIRoleAssignmentsV3RoleIDUserSelectorInternalServerError creates a GetAPIRoleAssignmentsV3RoleIDUserSelectorInternalServerError with default headers values
func NewGetAPIRoleAssignmentsV3RoleIDUserSelectorInternalServerError() *GetAPIRoleAssignmentsV3RoleIDUserSelectorInternalServerError {
	return &GetAPIRoleAssignmentsV3RoleIDUserSelectorInternalServerError{}
}

/* GetAPIRoleAssignmentsV3RoleIDUserSelectorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPIRoleAssignmentsV3RoleIDUserSelectorInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIRoleAssignmentsV3RoleIDUserSelectorInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/role-assignments/v3/{role-id}/{user-selector}][%d] getApiRoleAssignmentsV3RoleIdUserSelectorInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPIRoleAssignmentsV3RoleIDUserSelectorInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIRoleAssignmentsV3RoleIDUserSelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
