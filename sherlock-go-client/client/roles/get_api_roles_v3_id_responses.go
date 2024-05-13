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

// GetAPIRolesV3IDReader is a Reader for the GetAPIRolesV3ID structure.
type GetAPIRolesV3IDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIRolesV3IDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIRolesV3IDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIRolesV3IDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIRolesV3IDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIRolesV3IDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPIRolesV3IDProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPIRolesV3IDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIRolesV3IDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIRolesV3IDOK creates a GetAPIRolesV3IDOK with default headers values
func NewGetAPIRolesV3IDOK() *GetAPIRolesV3IDOK {
	return &GetAPIRolesV3IDOK{}
}

/* GetAPIRolesV3IDOK describes a response with status code 200, with default header values.

OK
*/
type GetAPIRolesV3IDOK struct {
	Payload *models.SherlockRoleV3
}

func (o *GetAPIRolesV3IDOK) Error() string {
	return fmt.Sprintf("[GET /api/roles/v3/{id}][%d] getApiRolesV3IdOK  %+v", 200, o.Payload)
}
func (o *GetAPIRolesV3IDOK) GetPayload() *models.SherlockRoleV3 {
	return o.Payload
}

func (o *GetAPIRolesV3IDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SherlockRoleV3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIRolesV3IDBadRequest creates a GetAPIRolesV3IDBadRequest with default headers values
func NewGetAPIRolesV3IDBadRequest() *GetAPIRolesV3IDBadRequest {
	return &GetAPIRolesV3IDBadRequest{}
}

/* GetAPIRolesV3IDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIRolesV3IDBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIRolesV3IDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/roles/v3/{id}][%d] getApiRolesV3IdBadRequest  %+v", 400, o.Payload)
}
func (o *GetAPIRolesV3IDBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIRolesV3IDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIRolesV3IDForbidden creates a GetAPIRolesV3IDForbidden with default headers values
func NewGetAPIRolesV3IDForbidden() *GetAPIRolesV3IDForbidden {
	return &GetAPIRolesV3IDForbidden{}
}

/* GetAPIRolesV3IDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPIRolesV3IDForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIRolesV3IDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/roles/v3/{id}][%d] getApiRolesV3IdForbidden  %+v", 403, o.Payload)
}
func (o *GetAPIRolesV3IDForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIRolesV3IDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIRolesV3IDNotFound creates a GetAPIRolesV3IDNotFound with default headers values
func NewGetAPIRolesV3IDNotFound() *GetAPIRolesV3IDNotFound {
	return &GetAPIRolesV3IDNotFound{}
}

/* GetAPIRolesV3IDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIRolesV3IDNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIRolesV3IDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/roles/v3/{id}][%d] getApiRolesV3IdNotFound  %+v", 404, o.Payload)
}
func (o *GetAPIRolesV3IDNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIRolesV3IDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIRolesV3IDProxyAuthenticationRequired creates a GetAPIRolesV3IDProxyAuthenticationRequired with default headers values
func NewGetAPIRolesV3IDProxyAuthenticationRequired() *GetAPIRolesV3IDProxyAuthenticationRequired {
	return &GetAPIRolesV3IDProxyAuthenticationRequired{}
}

/* GetAPIRolesV3IDProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPIRolesV3IDProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIRolesV3IDProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/roles/v3/{id}][%d] getApiRolesV3IdProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPIRolesV3IDProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIRolesV3IDProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIRolesV3IDConflict creates a GetAPIRolesV3IDConflict with default headers values
func NewGetAPIRolesV3IDConflict() *GetAPIRolesV3IDConflict {
	return &GetAPIRolesV3IDConflict{}
}

/* GetAPIRolesV3IDConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPIRolesV3IDConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIRolesV3IDConflict) Error() string {
	return fmt.Sprintf("[GET /api/roles/v3/{id}][%d] getApiRolesV3IdConflict  %+v", 409, o.Payload)
}
func (o *GetAPIRolesV3IDConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIRolesV3IDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIRolesV3IDInternalServerError creates a GetAPIRolesV3IDInternalServerError with default headers values
func NewGetAPIRolesV3IDInternalServerError() *GetAPIRolesV3IDInternalServerError {
	return &GetAPIRolesV3IDInternalServerError{}
}

/* GetAPIRolesV3IDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPIRolesV3IDInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIRolesV3IDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/roles/v3/{id}][%d] getApiRolesV3IdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPIRolesV3IDInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIRolesV3IDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}