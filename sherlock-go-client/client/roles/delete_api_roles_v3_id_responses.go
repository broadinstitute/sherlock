// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// DeleteAPIRolesV3IDReader is a Reader for the DeleteAPIRolesV3ID structure.
type DeleteAPIRolesV3IDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIRolesV3IDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAPIRolesV3IDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteAPIRolesV3IDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAPIRolesV3IDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAPIRolesV3IDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewDeleteAPIRolesV3IDProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteAPIRolesV3IDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAPIRolesV3IDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAPIRolesV3IDOK creates a DeleteAPIRolesV3IDOK with default headers values
func NewDeleteAPIRolesV3IDOK() *DeleteAPIRolesV3IDOK {
	return &DeleteAPIRolesV3IDOK{}
}

/* DeleteAPIRolesV3IDOK describes a response with status code 200, with default header values.

OK
*/
type DeleteAPIRolesV3IDOK struct {
	Payload *models.SherlockRoleV3
}

func (o *DeleteAPIRolesV3IDOK) Error() string {
	return fmt.Sprintf("[DELETE /api/roles/v3/{id}][%d] deleteApiRolesV3IdOK  %+v", 200, o.Payload)
}
func (o *DeleteAPIRolesV3IDOK) GetPayload() *models.SherlockRoleV3 {
	return o.Payload
}

func (o *DeleteAPIRolesV3IDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SherlockRoleV3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIRolesV3IDBadRequest creates a DeleteAPIRolesV3IDBadRequest with default headers values
func NewDeleteAPIRolesV3IDBadRequest() *DeleteAPIRolesV3IDBadRequest {
	return &DeleteAPIRolesV3IDBadRequest{}
}

/* DeleteAPIRolesV3IDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteAPIRolesV3IDBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIRolesV3IDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/roles/v3/{id}][%d] deleteApiRolesV3IdBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteAPIRolesV3IDBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIRolesV3IDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIRolesV3IDForbidden creates a DeleteAPIRolesV3IDForbidden with default headers values
func NewDeleteAPIRolesV3IDForbidden() *DeleteAPIRolesV3IDForbidden {
	return &DeleteAPIRolesV3IDForbidden{}
}

/* DeleteAPIRolesV3IDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteAPIRolesV3IDForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIRolesV3IDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/roles/v3/{id}][%d] deleteApiRolesV3IdForbidden  %+v", 403, o.Payload)
}
func (o *DeleteAPIRolesV3IDForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIRolesV3IDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIRolesV3IDNotFound creates a DeleteAPIRolesV3IDNotFound with default headers values
func NewDeleteAPIRolesV3IDNotFound() *DeleteAPIRolesV3IDNotFound {
	return &DeleteAPIRolesV3IDNotFound{}
}

/* DeleteAPIRolesV3IDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteAPIRolesV3IDNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIRolesV3IDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/roles/v3/{id}][%d] deleteApiRolesV3IdNotFound  %+v", 404, o.Payload)
}
func (o *DeleteAPIRolesV3IDNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIRolesV3IDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIRolesV3IDProxyAuthenticationRequired creates a DeleteAPIRolesV3IDProxyAuthenticationRequired with default headers values
func NewDeleteAPIRolesV3IDProxyAuthenticationRequired() *DeleteAPIRolesV3IDProxyAuthenticationRequired {
	return &DeleteAPIRolesV3IDProxyAuthenticationRequired{}
}

/* DeleteAPIRolesV3IDProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type DeleteAPIRolesV3IDProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIRolesV3IDProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[DELETE /api/roles/v3/{id}][%d] deleteApiRolesV3IdProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *DeleteAPIRolesV3IDProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIRolesV3IDProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIRolesV3IDConflict creates a DeleteAPIRolesV3IDConflict with default headers values
func NewDeleteAPIRolesV3IDConflict() *DeleteAPIRolesV3IDConflict {
	return &DeleteAPIRolesV3IDConflict{}
}

/* DeleteAPIRolesV3IDConflict describes a response with status code 409, with default header values.

Conflict
*/
type DeleteAPIRolesV3IDConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIRolesV3IDConflict) Error() string {
	return fmt.Sprintf("[DELETE /api/roles/v3/{id}][%d] deleteApiRolesV3IdConflict  %+v", 409, o.Payload)
}
func (o *DeleteAPIRolesV3IDConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIRolesV3IDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIRolesV3IDInternalServerError creates a DeleteAPIRolesV3IDInternalServerError with default headers values
func NewDeleteAPIRolesV3IDInternalServerError() *DeleteAPIRolesV3IDInternalServerError {
	return &DeleteAPIRolesV3IDInternalServerError{}
}

/* DeleteAPIRolesV3IDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteAPIRolesV3IDInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIRolesV3IDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/roles/v3/{id}][%d] deleteApiRolesV3IdInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteAPIRolesV3IDInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIRolesV3IDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
