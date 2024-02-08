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

// GetAPIChangesetsV3Reader is a Reader for the GetAPIChangesetsV3 structure.
type GetAPIChangesetsV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIChangesetsV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIChangesetsV3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIChangesetsV3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIChangesetsV3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIChangesetsV3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPIChangesetsV3ProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPIChangesetsV3Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIChangesetsV3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIChangesetsV3OK creates a GetAPIChangesetsV3OK with default headers values
func NewGetAPIChangesetsV3OK() *GetAPIChangesetsV3OK {
	return &GetAPIChangesetsV3OK{}
}

/* GetAPIChangesetsV3OK describes a response with status code 200, with default header values.

OK
*/
type GetAPIChangesetsV3OK struct {
	Payload []*models.SherlockChangesetV3
}

func (o *GetAPIChangesetsV3OK) Error() string {
	return fmt.Sprintf("[GET /api/changesets/v3][%d] getApiChangesetsV3OK  %+v", 200, o.Payload)
}
func (o *GetAPIChangesetsV3OK) GetPayload() []*models.SherlockChangesetV3 {
	return o.Payload
}

func (o *GetAPIChangesetsV3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChangesetsV3BadRequest creates a GetAPIChangesetsV3BadRequest with default headers values
func NewGetAPIChangesetsV3BadRequest() *GetAPIChangesetsV3BadRequest {
	return &GetAPIChangesetsV3BadRequest{}
}

/* GetAPIChangesetsV3BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIChangesetsV3BadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChangesetsV3BadRequest) Error() string {
	return fmt.Sprintf("[GET /api/changesets/v3][%d] getApiChangesetsV3BadRequest  %+v", 400, o.Payload)
}
func (o *GetAPIChangesetsV3BadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChangesetsV3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChangesetsV3Forbidden creates a GetAPIChangesetsV3Forbidden with default headers values
func NewGetAPIChangesetsV3Forbidden() *GetAPIChangesetsV3Forbidden {
	return &GetAPIChangesetsV3Forbidden{}
}

/* GetAPIChangesetsV3Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPIChangesetsV3Forbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChangesetsV3Forbidden) Error() string {
	return fmt.Sprintf("[GET /api/changesets/v3][%d] getApiChangesetsV3Forbidden  %+v", 403, o.Payload)
}
func (o *GetAPIChangesetsV3Forbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChangesetsV3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChangesetsV3NotFound creates a GetAPIChangesetsV3NotFound with default headers values
func NewGetAPIChangesetsV3NotFound() *GetAPIChangesetsV3NotFound {
	return &GetAPIChangesetsV3NotFound{}
}

/* GetAPIChangesetsV3NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIChangesetsV3NotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChangesetsV3NotFound) Error() string {
	return fmt.Sprintf("[GET /api/changesets/v3][%d] getApiChangesetsV3NotFound  %+v", 404, o.Payload)
}
func (o *GetAPIChangesetsV3NotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChangesetsV3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChangesetsV3ProxyAuthenticationRequired creates a GetAPIChangesetsV3ProxyAuthenticationRequired with default headers values
func NewGetAPIChangesetsV3ProxyAuthenticationRequired() *GetAPIChangesetsV3ProxyAuthenticationRequired {
	return &GetAPIChangesetsV3ProxyAuthenticationRequired{}
}

/* GetAPIChangesetsV3ProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPIChangesetsV3ProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChangesetsV3ProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/changesets/v3][%d] getApiChangesetsV3ProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPIChangesetsV3ProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChangesetsV3ProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChangesetsV3Conflict creates a GetAPIChangesetsV3Conflict with default headers values
func NewGetAPIChangesetsV3Conflict() *GetAPIChangesetsV3Conflict {
	return &GetAPIChangesetsV3Conflict{}
}

/* GetAPIChangesetsV3Conflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPIChangesetsV3Conflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChangesetsV3Conflict) Error() string {
	return fmt.Sprintf("[GET /api/changesets/v3][%d] getApiChangesetsV3Conflict  %+v", 409, o.Payload)
}
func (o *GetAPIChangesetsV3Conflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChangesetsV3Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChangesetsV3InternalServerError creates a GetAPIChangesetsV3InternalServerError with default headers values
func NewGetAPIChangesetsV3InternalServerError() *GetAPIChangesetsV3InternalServerError {
	return &GetAPIChangesetsV3InternalServerError{}
}

/* GetAPIChangesetsV3InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPIChangesetsV3InternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChangesetsV3InternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/changesets/v3][%d] getApiChangesetsV3InternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPIChangesetsV3InternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChangesetsV3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
