// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// PostAPIClustersV3Reader is a Reader for the PostAPIClustersV3 structure.
type PostAPIClustersV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIClustersV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostAPIClustersV3Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAPIClustersV3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAPIClustersV3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAPIClustersV3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewPostAPIClustersV3ProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostAPIClustersV3Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAPIClustersV3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAPIClustersV3Created creates a PostAPIClustersV3Created with default headers values
func NewPostAPIClustersV3Created() *PostAPIClustersV3Created {
	return &PostAPIClustersV3Created{}
}

/* PostAPIClustersV3Created describes a response with status code 201, with default header values.

Created
*/
type PostAPIClustersV3Created struct {
	Payload *models.SherlockClusterV3
}

func (o *PostAPIClustersV3Created) Error() string {
	return fmt.Sprintf("[POST /api/clusters/v3][%d] postApiClustersV3Created  %+v", 201, o.Payload)
}
func (o *PostAPIClustersV3Created) GetPayload() *models.SherlockClusterV3 {
	return o.Payload
}

func (o *PostAPIClustersV3Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SherlockClusterV3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIClustersV3BadRequest creates a PostAPIClustersV3BadRequest with default headers values
func NewPostAPIClustersV3BadRequest() *PostAPIClustersV3BadRequest {
	return &PostAPIClustersV3BadRequest{}
}

/* PostAPIClustersV3BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostAPIClustersV3BadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIClustersV3BadRequest) Error() string {
	return fmt.Sprintf("[POST /api/clusters/v3][%d] postApiClustersV3BadRequest  %+v", 400, o.Payload)
}
func (o *PostAPIClustersV3BadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIClustersV3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIClustersV3Forbidden creates a PostAPIClustersV3Forbidden with default headers values
func NewPostAPIClustersV3Forbidden() *PostAPIClustersV3Forbidden {
	return &PostAPIClustersV3Forbidden{}
}

/* PostAPIClustersV3Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostAPIClustersV3Forbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIClustersV3Forbidden) Error() string {
	return fmt.Sprintf("[POST /api/clusters/v3][%d] postApiClustersV3Forbidden  %+v", 403, o.Payload)
}
func (o *PostAPIClustersV3Forbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIClustersV3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIClustersV3NotFound creates a PostAPIClustersV3NotFound with default headers values
func NewPostAPIClustersV3NotFound() *PostAPIClustersV3NotFound {
	return &PostAPIClustersV3NotFound{}
}

/* PostAPIClustersV3NotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostAPIClustersV3NotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIClustersV3NotFound) Error() string {
	return fmt.Sprintf("[POST /api/clusters/v3][%d] postApiClustersV3NotFound  %+v", 404, o.Payload)
}
func (o *PostAPIClustersV3NotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIClustersV3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIClustersV3ProxyAuthenticationRequired creates a PostAPIClustersV3ProxyAuthenticationRequired with default headers values
func NewPostAPIClustersV3ProxyAuthenticationRequired() *PostAPIClustersV3ProxyAuthenticationRequired {
	return &PostAPIClustersV3ProxyAuthenticationRequired{}
}

/* PostAPIClustersV3ProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type PostAPIClustersV3ProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIClustersV3ProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[POST /api/clusters/v3][%d] postApiClustersV3ProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *PostAPIClustersV3ProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIClustersV3ProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIClustersV3Conflict creates a PostAPIClustersV3Conflict with default headers values
func NewPostAPIClustersV3Conflict() *PostAPIClustersV3Conflict {
	return &PostAPIClustersV3Conflict{}
}

/* PostAPIClustersV3Conflict describes a response with status code 409, with default header values.

Conflict
*/
type PostAPIClustersV3Conflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIClustersV3Conflict) Error() string {
	return fmt.Sprintf("[POST /api/clusters/v3][%d] postApiClustersV3Conflict  %+v", 409, o.Payload)
}
func (o *PostAPIClustersV3Conflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIClustersV3Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIClustersV3InternalServerError creates a PostAPIClustersV3InternalServerError with default headers values
func NewPostAPIClustersV3InternalServerError() *PostAPIClustersV3InternalServerError {
	return &PostAPIClustersV3InternalServerError{}
}

/* PostAPIClustersV3InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostAPIClustersV3InternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIClustersV3InternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/clusters/v3][%d] postApiClustersV3InternalServerError  %+v", 500, o.Payload)
}
func (o *PostAPIClustersV3InternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIClustersV3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
