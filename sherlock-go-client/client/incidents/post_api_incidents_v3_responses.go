// Code generated by go-swagger; DO NOT EDIT.

package incidents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// PostAPIIncidentsV3Reader is a Reader for the PostAPIIncidentsV3 structure.
type PostAPIIncidentsV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIIncidentsV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostAPIIncidentsV3Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAPIIncidentsV3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAPIIncidentsV3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAPIIncidentsV3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewPostAPIIncidentsV3ProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostAPIIncidentsV3Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAPIIncidentsV3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAPIIncidentsV3Created creates a PostAPIIncidentsV3Created with default headers values
func NewPostAPIIncidentsV3Created() *PostAPIIncidentsV3Created {
	return &PostAPIIncidentsV3Created{}
}

/* PostAPIIncidentsV3Created describes a response with status code 201, with default header values.

Created
*/
type PostAPIIncidentsV3Created struct {
	Payload *models.SherlockIncidentV3
}

func (o *PostAPIIncidentsV3Created) Error() string {
	return fmt.Sprintf("[POST /api/incidents/v3][%d] postApiIncidentsV3Created  %+v", 201, o.Payload)
}
func (o *PostAPIIncidentsV3Created) GetPayload() *models.SherlockIncidentV3 {
	return o.Payload
}

func (o *PostAPIIncidentsV3Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SherlockIncidentV3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIIncidentsV3BadRequest creates a PostAPIIncidentsV3BadRequest with default headers values
func NewPostAPIIncidentsV3BadRequest() *PostAPIIncidentsV3BadRequest {
	return &PostAPIIncidentsV3BadRequest{}
}

/* PostAPIIncidentsV3BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostAPIIncidentsV3BadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIIncidentsV3BadRequest) Error() string {
	return fmt.Sprintf("[POST /api/incidents/v3][%d] postApiIncidentsV3BadRequest  %+v", 400, o.Payload)
}
func (o *PostAPIIncidentsV3BadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIIncidentsV3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIIncidentsV3Forbidden creates a PostAPIIncidentsV3Forbidden with default headers values
func NewPostAPIIncidentsV3Forbidden() *PostAPIIncidentsV3Forbidden {
	return &PostAPIIncidentsV3Forbidden{}
}

/* PostAPIIncidentsV3Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostAPIIncidentsV3Forbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIIncidentsV3Forbidden) Error() string {
	return fmt.Sprintf("[POST /api/incidents/v3][%d] postApiIncidentsV3Forbidden  %+v", 403, o.Payload)
}
func (o *PostAPIIncidentsV3Forbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIIncidentsV3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIIncidentsV3NotFound creates a PostAPIIncidentsV3NotFound with default headers values
func NewPostAPIIncidentsV3NotFound() *PostAPIIncidentsV3NotFound {
	return &PostAPIIncidentsV3NotFound{}
}

/* PostAPIIncidentsV3NotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostAPIIncidentsV3NotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIIncidentsV3NotFound) Error() string {
	return fmt.Sprintf("[POST /api/incidents/v3][%d] postApiIncidentsV3NotFound  %+v", 404, o.Payload)
}
func (o *PostAPIIncidentsV3NotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIIncidentsV3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIIncidentsV3ProxyAuthenticationRequired creates a PostAPIIncidentsV3ProxyAuthenticationRequired with default headers values
func NewPostAPIIncidentsV3ProxyAuthenticationRequired() *PostAPIIncidentsV3ProxyAuthenticationRequired {
	return &PostAPIIncidentsV3ProxyAuthenticationRequired{}
}

/* PostAPIIncidentsV3ProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type PostAPIIncidentsV3ProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIIncidentsV3ProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[POST /api/incidents/v3][%d] postApiIncidentsV3ProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *PostAPIIncidentsV3ProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIIncidentsV3ProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIIncidentsV3Conflict creates a PostAPIIncidentsV3Conflict with default headers values
func NewPostAPIIncidentsV3Conflict() *PostAPIIncidentsV3Conflict {
	return &PostAPIIncidentsV3Conflict{}
}

/* PostAPIIncidentsV3Conflict describes a response with status code 409, with default header values.

Conflict
*/
type PostAPIIncidentsV3Conflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIIncidentsV3Conflict) Error() string {
	return fmt.Sprintf("[POST /api/incidents/v3][%d] postApiIncidentsV3Conflict  %+v", 409, o.Payload)
}
func (o *PostAPIIncidentsV3Conflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIIncidentsV3Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIIncidentsV3InternalServerError creates a PostAPIIncidentsV3InternalServerError with default headers values
func NewPostAPIIncidentsV3InternalServerError() *PostAPIIncidentsV3InternalServerError {
	return &PostAPIIncidentsV3InternalServerError{}
}

/* PostAPIIncidentsV3InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostAPIIncidentsV3InternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIIncidentsV3InternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/incidents/v3][%d] postApiIncidentsV3InternalServerError  %+v", 500, o.Payload)
}
func (o *PostAPIIncidentsV3InternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIIncidentsV3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}