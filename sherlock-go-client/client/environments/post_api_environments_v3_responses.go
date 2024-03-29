// Code generated by go-swagger; DO NOT EDIT.

package environments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// PostAPIEnvironmentsV3Reader is a Reader for the PostAPIEnvironmentsV3 structure.
type PostAPIEnvironmentsV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIEnvironmentsV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostAPIEnvironmentsV3Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAPIEnvironmentsV3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAPIEnvironmentsV3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAPIEnvironmentsV3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewPostAPIEnvironmentsV3ProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostAPIEnvironmentsV3Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAPIEnvironmentsV3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAPIEnvironmentsV3Created creates a PostAPIEnvironmentsV3Created with default headers values
func NewPostAPIEnvironmentsV3Created() *PostAPIEnvironmentsV3Created {
	return &PostAPIEnvironmentsV3Created{}
}

/* PostAPIEnvironmentsV3Created describes a response with status code 201, with default header values.

Created
*/
type PostAPIEnvironmentsV3Created struct {
	Payload *models.SherlockEnvironmentV3
}

func (o *PostAPIEnvironmentsV3Created) Error() string {
	return fmt.Sprintf("[POST /api/environments/v3][%d] postApiEnvironmentsV3Created  %+v", 201, o.Payload)
}
func (o *PostAPIEnvironmentsV3Created) GetPayload() *models.SherlockEnvironmentV3 {
	return o.Payload
}

func (o *PostAPIEnvironmentsV3Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SherlockEnvironmentV3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIEnvironmentsV3BadRequest creates a PostAPIEnvironmentsV3BadRequest with default headers values
func NewPostAPIEnvironmentsV3BadRequest() *PostAPIEnvironmentsV3BadRequest {
	return &PostAPIEnvironmentsV3BadRequest{}
}

/* PostAPIEnvironmentsV3BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostAPIEnvironmentsV3BadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIEnvironmentsV3BadRequest) Error() string {
	return fmt.Sprintf("[POST /api/environments/v3][%d] postApiEnvironmentsV3BadRequest  %+v", 400, o.Payload)
}
func (o *PostAPIEnvironmentsV3BadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIEnvironmentsV3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIEnvironmentsV3Forbidden creates a PostAPIEnvironmentsV3Forbidden with default headers values
func NewPostAPIEnvironmentsV3Forbidden() *PostAPIEnvironmentsV3Forbidden {
	return &PostAPIEnvironmentsV3Forbidden{}
}

/* PostAPIEnvironmentsV3Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostAPIEnvironmentsV3Forbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIEnvironmentsV3Forbidden) Error() string {
	return fmt.Sprintf("[POST /api/environments/v3][%d] postApiEnvironmentsV3Forbidden  %+v", 403, o.Payload)
}
func (o *PostAPIEnvironmentsV3Forbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIEnvironmentsV3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIEnvironmentsV3NotFound creates a PostAPIEnvironmentsV3NotFound with default headers values
func NewPostAPIEnvironmentsV3NotFound() *PostAPIEnvironmentsV3NotFound {
	return &PostAPIEnvironmentsV3NotFound{}
}

/* PostAPIEnvironmentsV3NotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostAPIEnvironmentsV3NotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIEnvironmentsV3NotFound) Error() string {
	return fmt.Sprintf("[POST /api/environments/v3][%d] postApiEnvironmentsV3NotFound  %+v", 404, o.Payload)
}
func (o *PostAPIEnvironmentsV3NotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIEnvironmentsV3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIEnvironmentsV3ProxyAuthenticationRequired creates a PostAPIEnvironmentsV3ProxyAuthenticationRequired with default headers values
func NewPostAPIEnvironmentsV3ProxyAuthenticationRequired() *PostAPIEnvironmentsV3ProxyAuthenticationRequired {
	return &PostAPIEnvironmentsV3ProxyAuthenticationRequired{}
}

/* PostAPIEnvironmentsV3ProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type PostAPIEnvironmentsV3ProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIEnvironmentsV3ProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[POST /api/environments/v3][%d] postApiEnvironmentsV3ProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *PostAPIEnvironmentsV3ProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIEnvironmentsV3ProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIEnvironmentsV3Conflict creates a PostAPIEnvironmentsV3Conflict with default headers values
func NewPostAPIEnvironmentsV3Conflict() *PostAPIEnvironmentsV3Conflict {
	return &PostAPIEnvironmentsV3Conflict{}
}

/* PostAPIEnvironmentsV3Conflict describes a response with status code 409, with default header values.

Conflict
*/
type PostAPIEnvironmentsV3Conflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIEnvironmentsV3Conflict) Error() string {
	return fmt.Sprintf("[POST /api/environments/v3][%d] postApiEnvironmentsV3Conflict  %+v", 409, o.Payload)
}
func (o *PostAPIEnvironmentsV3Conflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIEnvironmentsV3Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIEnvironmentsV3InternalServerError creates a PostAPIEnvironmentsV3InternalServerError with default headers values
func NewPostAPIEnvironmentsV3InternalServerError() *PostAPIEnvironmentsV3InternalServerError {
	return &PostAPIEnvironmentsV3InternalServerError{}
}

/* PostAPIEnvironmentsV3InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostAPIEnvironmentsV3InternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIEnvironmentsV3InternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/environments/v3][%d] postApiEnvironmentsV3InternalServerError  %+v", 500, o.Payload)
}
func (o *PostAPIEnvironmentsV3InternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIEnvironmentsV3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
