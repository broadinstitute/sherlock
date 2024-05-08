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

// GetAPIRoleAssignmentsV3Reader is a Reader for the GetAPIRoleAssignmentsV3 structure.
type GetAPIRoleAssignmentsV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIRoleAssignmentsV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIRoleAssignmentsV3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIRoleAssignmentsV3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIRoleAssignmentsV3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIRoleAssignmentsV3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPIRoleAssignmentsV3ProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPIRoleAssignmentsV3Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIRoleAssignmentsV3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIRoleAssignmentsV3OK creates a GetAPIRoleAssignmentsV3OK with default headers values
func NewGetAPIRoleAssignmentsV3OK() *GetAPIRoleAssignmentsV3OK {
	return &GetAPIRoleAssignmentsV3OK{}
}

/* GetAPIRoleAssignmentsV3OK describes a response with status code 200, with default header values.

OK
*/
type GetAPIRoleAssignmentsV3OK struct {
	Payload []*models.SherlockRoleAssignmentV3
}

func (o *GetAPIRoleAssignmentsV3OK) Error() string {
	return fmt.Sprintf("[GET /api/role-assignments/v3][%d] getApiRoleAssignmentsV3OK  %+v", 200, o.Payload)
}
func (o *GetAPIRoleAssignmentsV3OK) GetPayload() []*models.SherlockRoleAssignmentV3 {
	return o.Payload
}

func (o *GetAPIRoleAssignmentsV3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIRoleAssignmentsV3BadRequest creates a GetAPIRoleAssignmentsV3BadRequest with default headers values
func NewGetAPIRoleAssignmentsV3BadRequest() *GetAPIRoleAssignmentsV3BadRequest {
	return &GetAPIRoleAssignmentsV3BadRequest{}
}

/* GetAPIRoleAssignmentsV3BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIRoleAssignmentsV3BadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIRoleAssignmentsV3BadRequest) Error() string {
	return fmt.Sprintf("[GET /api/role-assignments/v3][%d] getApiRoleAssignmentsV3BadRequest  %+v", 400, o.Payload)
}
func (o *GetAPIRoleAssignmentsV3BadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIRoleAssignmentsV3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIRoleAssignmentsV3Forbidden creates a GetAPIRoleAssignmentsV3Forbidden with default headers values
func NewGetAPIRoleAssignmentsV3Forbidden() *GetAPIRoleAssignmentsV3Forbidden {
	return &GetAPIRoleAssignmentsV3Forbidden{}
}

/* GetAPIRoleAssignmentsV3Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPIRoleAssignmentsV3Forbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIRoleAssignmentsV3Forbidden) Error() string {
	return fmt.Sprintf("[GET /api/role-assignments/v3][%d] getApiRoleAssignmentsV3Forbidden  %+v", 403, o.Payload)
}
func (o *GetAPIRoleAssignmentsV3Forbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIRoleAssignmentsV3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIRoleAssignmentsV3NotFound creates a GetAPIRoleAssignmentsV3NotFound with default headers values
func NewGetAPIRoleAssignmentsV3NotFound() *GetAPIRoleAssignmentsV3NotFound {
	return &GetAPIRoleAssignmentsV3NotFound{}
}

/* GetAPIRoleAssignmentsV3NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIRoleAssignmentsV3NotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIRoleAssignmentsV3NotFound) Error() string {
	return fmt.Sprintf("[GET /api/role-assignments/v3][%d] getApiRoleAssignmentsV3NotFound  %+v", 404, o.Payload)
}
func (o *GetAPIRoleAssignmentsV3NotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIRoleAssignmentsV3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIRoleAssignmentsV3ProxyAuthenticationRequired creates a GetAPIRoleAssignmentsV3ProxyAuthenticationRequired with default headers values
func NewGetAPIRoleAssignmentsV3ProxyAuthenticationRequired() *GetAPIRoleAssignmentsV3ProxyAuthenticationRequired {
	return &GetAPIRoleAssignmentsV3ProxyAuthenticationRequired{}
}

/* GetAPIRoleAssignmentsV3ProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPIRoleAssignmentsV3ProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIRoleAssignmentsV3ProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/role-assignments/v3][%d] getApiRoleAssignmentsV3ProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPIRoleAssignmentsV3ProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIRoleAssignmentsV3ProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIRoleAssignmentsV3Conflict creates a GetAPIRoleAssignmentsV3Conflict with default headers values
func NewGetAPIRoleAssignmentsV3Conflict() *GetAPIRoleAssignmentsV3Conflict {
	return &GetAPIRoleAssignmentsV3Conflict{}
}

/* GetAPIRoleAssignmentsV3Conflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPIRoleAssignmentsV3Conflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIRoleAssignmentsV3Conflict) Error() string {
	return fmt.Sprintf("[GET /api/role-assignments/v3][%d] getApiRoleAssignmentsV3Conflict  %+v", 409, o.Payload)
}
func (o *GetAPIRoleAssignmentsV3Conflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIRoleAssignmentsV3Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIRoleAssignmentsV3InternalServerError creates a GetAPIRoleAssignmentsV3InternalServerError with default headers values
func NewGetAPIRoleAssignmentsV3InternalServerError() *GetAPIRoleAssignmentsV3InternalServerError {
	return &GetAPIRoleAssignmentsV3InternalServerError{}
}

/* GetAPIRoleAssignmentsV3InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPIRoleAssignmentsV3InternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIRoleAssignmentsV3InternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/role-assignments/v3][%d] getApiRoleAssignmentsV3InternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPIRoleAssignmentsV3InternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIRoleAssignmentsV3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
