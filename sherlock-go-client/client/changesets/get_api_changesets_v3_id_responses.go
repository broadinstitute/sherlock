// Code generated by go-swagger; DO NOT EDIT.

package changesets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// GetAPIChangesetsV3IDReader is a Reader for the GetAPIChangesetsV3ID structure.
type GetAPIChangesetsV3IDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIChangesetsV3IDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIChangesetsV3IDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIChangesetsV3IDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIChangesetsV3IDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIChangesetsV3IDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPIChangesetsV3IDProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPIChangesetsV3IDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIChangesetsV3IDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIChangesetsV3IDOK creates a GetAPIChangesetsV3IDOK with default headers values
func NewGetAPIChangesetsV3IDOK() *GetAPIChangesetsV3IDOK {
	return &GetAPIChangesetsV3IDOK{}
}

/* GetAPIChangesetsV3IDOK describes a response with status code 200, with default header values.

OK
*/
type GetAPIChangesetsV3IDOK struct {
	Payload *models.SherlockChangesetV3
}

func (o *GetAPIChangesetsV3IDOK) Error() string {
	return fmt.Sprintf("[GET /api/changesets/v3/{id}][%d] getApiChangesetsV3IdOK  %+v", 200, o.Payload)
}
func (o *GetAPIChangesetsV3IDOK) GetPayload() *models.SherlockChangesetV3 {
	return o.Payload
}

func (o *GetAPIChangesetsV3IDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SherlockChangesetV3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChangesetsV3IDBadRequest creates a GetAPIChangesetsV3IDBadRequest with default headers values
func NewGetAPIChangesetsV3IDBadRequest() *GetAPIChangesetsV3IDBadRequest {
	return &GetAPIChangesetsV3IDBadRequest{}
}

/* GetAPIChangesetsV3IDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIChangesetsV3IDBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChangesetsV3IDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/changesets/v3/{id}][%d] getApiChangesetsV3IdBadRequest  %+v", 400, o.Payload)
}
func (o *GetAPIChangesetsV3IDBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChangesetsV3IDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChangesetsV3IDForbidden creates a GetAPIChangesetsV3IDForbidden with default headers values
func NewGetAPIChangesetsV3IDForbidden() *GetAPIChangesetsV3IDForbidden {
	return &GetAPIChangesetsV3IDForbidden{}
}

/* GetAPIChangesetsV3IDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPIChangesetsV3IDForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChangesetsV3IDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/changesets/v3/{id}][%d] getApiChangesetsV3IdForbidden  %+v", 403, o.Payload)
}
func (o *GetAPIChangesetsV3IDForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChangesetsV3IDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChangesetsV3IDNotFound creates a GetAPIChangesetsV3IDNotFound with default headers values
func NewGetAPIChangesetsV3IDNotFound() *GetAPIChangesetsV3IDNotFound {
	return &GetAPIChangesetsV3IDNotFound{}
}

/* GetAPIChangesetsV3IDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIChangesetsV3IDNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChangesetsV3IDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/changesets/v3/{id}][%d] getApiChangesetsV3IdNotFound  %+v", 404, o.Payload)
}
func (o *GetAPIChangesetsV3IDNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChangesetsV3IDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChangesetsV3IDProxyAuthenticationRequired creates a GetAPIChangesetsV3IDProxyAuthenticationRequired with default headers values
func NewGetAPIChangesetsV3IDProxyAuthenticationRequired() *GetAPIChangesetsV3IDProxyAuthenticationRequired {
	return &GetAPIChangesetsV3IDProxyAuthenticationRequired{}
}

/* GetAPIChangesetsV3IDProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPIChangesetsV3IDProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChangesetsV3IDProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/changesets/v3/{id}][%d] getApiChangesetsV3IdProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPIChangesetsV3IDProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChangesetsV3IDProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChangesetsV3IDConflict creates a GetAPIChangesetsV3IDConflict with default headers values
func NewGetAPIChangesetsV3IDConflict() *GetAPIChangesetsV3IDConflict {
	return &GetAPIChangesetsV3IDConflict{}
}

/* GetAPIChangesetsV3IDConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPIChangesetsV3IDConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChangesetsV3IDConflict) Error() string {
	return fmt.Sprintf("[GET /api/changesets/v3/{id}][%d] getApiChangesetsV3IdConflict  %+v", 409, o.Payload)
}
func (o *GetAPIChangesetsV3IDConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChangesetsV3IDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChangesetsV3IDInternalServerError creates a GetAPIChangesetsV3IDInternalServerError with default headers values
func NewGetAPIChangesetsV3IDInternalServerError() *GetAPIChangesetsV3IDInternalServerError {
	return &GetAPIChangesetsV3IDInternalServerError{}
}

/* GetAPIChangesetsV3IDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPIChangesetsV3IDInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChangesetsV3IDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/changesets/v3/{id}][%d] getApiChangesetsV3IdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPIChangesetsV3IDInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChangesetsV3IDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
