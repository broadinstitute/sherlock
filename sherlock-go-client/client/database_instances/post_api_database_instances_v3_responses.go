// Code generated by go-swagger; DO NOT EDIT.

package database_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// PostAPIDatabaseInstancesV3Reader is a Reader for the PostAPIDatabaseInstancesV3 structure.
type PostAPIDatabaseInstancesV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIDatabaseInstancesV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostAPIDatabaseInstancesV3Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAPIDatabaseInstancesV3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAPIDatabaseInstancesV3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAPIDatabaseInstancesV3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewPostAPIDatabaseInstancesV3ProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostAPIDatabaseInstancesV3Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAPIDatabaseInstancesV3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAPIDatabaseInstancesV3Created creates a PostAPIDatabaseInstancesV3Created with default headers values
func NewPostAPIDatabaseInstancesV3Created() *PostAPIDatabaseInstancesV3Created {
	return &PostAPIDatabaseInstancesV3Created{}
}

/* PostAPIDatabaseInstancesV3Created describes a response with status code 201, with default header values.

Created
*/
type PostAPIDatabaseInstancesV3Created struct {
	Payload *models.SherlockDatabaseInstanceV3
}

func (o *PostAPIDatabaseInstancesV3Created) Error() string {
	return fmt.Sprintf("[POST /api/database-instances/v3][%d] postApiDatabaseInstancesV3Created  %+v", 201, o.Payload)
}
func (o *PostAPIDatabaseInstancesV3Created) GetPayload() *models.SherlockDatabaseInstanceV3 {
	return o.Payload
}

func (o *PostAPIDatabaseInstancesV3Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SherlockDatabaseInstanceV3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIDatabaseInstancesV3BadRequest creates a PostAPIDatabaseInstancesV3BadRequest with default headers values
func NewPostAPIDatabaseInstancesV3BadRequest() *PostAPIDatabaseInstancesV3BadRequest {
	return &PostAPIDatabaseInstancesV3BadRequest{}
}

/* PostAPIDatabaseInstancesV3BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostAPIDatabaseInstancesV3BadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIDatabaseInstancesV3BadRequest) Error() string {
	return fmt.Sprintf("[POST /api/database-instances/v3][%d] postApiDatabaseInstancesV3BadRequest  %+v", 400, o.Payload)
}
func (o *PostAPIDatabaseInstancesV3BadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIDatabaseInstancesV3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIDatabaseInstancesV3Forbidden creates a PostAPIDatabaseInstancesV3Forbidden with default headers values
func NewPostAPIDatabaseInstancesV3Forbidden() *PostAPIDatabaseInstancesV3Forbidden {
	return &PostAPIDatabaseInstancesV3Forbidden{}
}

/* PostAPIDatabaseInstancesV3Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostAPIDatabaseInstancesV3Forbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIDatabaseInstancesV3Forbidden) Error() string {
	return fmt.Sprintf("[POST /api/database-instances/v3][%d] postApiDatabaseInstancesV3Forbidden  %+v", 403, o.Payload)
}
func (o *PostAPIDatabaseInstancesV3Forbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIDatabaseInstancesV3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIDatabaseInstancesV3NotFound creates a PostAPIDatabaseInstancesV3NotFound with default headers values
func NewPostAPIDatabaseInstancesV3NotFound() *PostAPIDatabaseInstancesV3NotFound {
	return &PostAPIDatabaseInstancesV3NotFound{}
}

/* PostAPIDatabaseInstancesV3NotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostAPIDatabaseInstancesV3NotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIDatabaseInstancesV3NotFound) Error() string {
	return fmt.Sprintf("[POST /api/database-instances/v3][%d] postApiDatabaseInstancesV3NotFound  %+v", 404, o.Payload)
}
func (o *PostAPIDatabaseInstancesV3NotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIDatabaseInstancesV3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIDatabaseInstancesV3ProxyAuthenticationRequired creates a PostAPIDatabaseInstancesV3ProxyAuthenticationRequired with default headers values
func NewPostAPIDatabaseInstancesV3ProxyAuthenticationRequired() *PostAPIDatabaseInstancesV3ProxyAuthenticationRequired {
	return &PostAPIDatabaseInstancesV3ProxyAuthenticationRequired{}
}

/* PostAPIDatabaseInstancesV3ProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type PostAPIDatabaseInstancesV3ProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIDatabaseInstancesV3ProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[POST /api/database-instances/v3][%d] postApiDatabaseInstancesV3ProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *PostAPIDatabaseInstancesV3ProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIDatabaseInstancesV3ProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIDatabaseInstancesV3Conflict creates a PostAPIDatabaseInstancesV3Conflict with default headers values
func NewPostAPIDatabaseInstancesV3Conflict() *PostAPIDatabaseInstancesV3Conflict {
	return &PostAPIDatabaseInstancesV3Conflict{}
}

/* PostAPIDatabaseInstancesV3Conflict describes a response with status code 409, with default header values.

Conflict
*/
type PostAPIDatabaseInstancesV3Conflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIDatabaseInstancesV3Conflict) Error() string {
	return fmt.Sprintf("[POST /api/database-instances/v3][%d] postApiDatabaseInstancesV3Conflict  %+v", 409, o.Payload)
}
func (o *PostAPIDatabaseInstancesV3Conflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIDatabaseInstancesV3Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIDatabaseInstancesV3InternalServerError creates a PostAPIDatabaseInstancesV3InternalServerError with default headers values
func NewPostAPIDatabaseInstancesV3InternalServerError() *PostAPIDatabaseInstancesV3InternalServerError {
	return &PostAPIDatabaseInstancesV3InternalServerError{}
}

/* PostAPIDatabaseInstancesV3InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostAPIDatabaseInstancesV3InternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIDatabaseInstancesV3InternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/database-instances/v3][%d] postApiDatabaseInstancesV3InternalServerError  %+v", 500, o.Payload)
}
func (o *PostAPIDatabaseInstancesV3InternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIDatabaseInstancesV3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
