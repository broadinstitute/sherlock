// Code generated by go-swagger; DO NOT EDIT.

package chart_releases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// PostAPIChartReleasesV3Reader is a Reader for the PostAPIChartReleasesV3 structure.
type PostAPIChartReleasesV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIChartReleasesV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostAPIChartReleasesV3Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAPIChartReleasesV3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAPIChartReleasesV3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAPIChartReleasesV3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewPostAPIChartReleasesV3ProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostAPIChartReleasesV3Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAPIChartReleasesV3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAPIChartReleasesV3Created creates a PostAPIChartReleasesV3Created with default headers values
func NewPostAPIChartReleasesV3Created() *PostAPIChartReleasesV3Created {
	return &PostAPIChartReleasesV3Created{}
}

/* PostAPIChartReleasesV3Created describes a response with status code 201, with default header values.

Created
*/
type PostAPIChartReleasesV3Created struct {
	Payload *models.SherlockChartReleaseV3
}

func (o *PostAPIChartReleasesV3Created) Error() string {
	return fmt.Sprintf("[POST /api/chart-releases/v3][%d] postApiChartReleasesV3Created  %+v", 201, o.Payload)
}
func (o *PostAPIChartReleasesV3Created) GetPayload() *models.SherlockChartReleaseV3 {
	return o.Payload
}

func (o *PostAPIChartReleasesV3Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SherlockChartReleaseV3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIChartReleasesV3BadRequest creates a PostAPIChartReleasesV3BadRequest with default headers values
func NewPostAPIChartReleasesV3BadRequest() *PostAPIChartReleasesV3BadRequest {
	return &PostAPIChartReleasesV3BadRequest{}
}

/* PostAPIChartReleasesV3BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostAPIChartReleasesV3BadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIChartReleasesV3BadRequest) Error() string {
	return fmt.Sprintf("[POST /api/chart-releases/v3][%d] postApiChartReleasesV3BadRequest  %+v", 400, o.Payload)
}
func (o *PostAPIChartReleasesV3BadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIChartReleasesV3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIChartReleasesV3Forbidden creates a PostAPIChartReleasesV3Forbidden with default headers values
func NewPostAPIChartReleasesV3Forbidden() *PostAPIChartReleasesV3Forbidden {
	return &PostAPIChartReleasesV3Forbidden{}
}

/* PostAPIChartReleasesV3Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostAPIChartReleasesV3Forbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIChartReleasesV3Forbidden) Error() string {
	return fmt.Sprintf("[POST /api/chart-releases/v3][%d] postApiChartReleasesV3Forbidden  %+v", 403, o.Payload)
}
func (o *PostAPIChartReleasesV3Forbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIChartReleasesV3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIChartReleasesV3NotFound creates a PostAPIChartReleasesV3NotFound with default headers values
func NewPostAPIChartReleasesV3NotFound() *PostAPIChartReleasesV3NotFound {
	return &PostAPIChartReleasesV3NotFound{}
}

/* PostAPIChartReleasesV3NotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostAPIChartReleasesV3NotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIChartReleasesV3NotFound) Error() string {
	return fmt.Sprintf("[POST /api/chart-releases/v3][%d] postApiChartReleasesV3NotFound  %+v", 404, o.Payload)
}
func (o *PostAPIChartReleasesV3NotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIChartReleasesV3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIChartReleasesV3ProxyAuthenticationRequired creates a PostAPIChartReleasesV3ProxyAuthenticationRequired with default headers values
func NewPostAPIChartReleasesV3ProxyAuthenticationRequired() *PostAPIChartReleasesV3ProxyAuthenticationRequired {
	return &PostAPIChartReleasesV3ProxyAuthenticationRequired{}
}

/* PostAPIChartReleasesV3ProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type PostAPIChartReleasesV3ProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIChartReleasesV3ProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[POST /api/chart-releases/v3][%d] postApiChartReleasesV3ProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *PostAPIChartReleasesV3ProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIChartReleasesV3ProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIChartReleasesV3Conflict creates a PostAPIChartReleasesV3Conflict with default headers values
func NewPostAPIChartReleasesV3Conflict() *PostAPIChartReleasesV3Conflict {
	return &PostAPIChartReleasesV3Conflict{}
}

/* PostAPIChartReleasesV3Conflict describes a response with status code 409, with default header values.

Conflict
*/
type PostAPIChartReleasesV3Conflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIChartReleasesV3Conflict) Error() string {
	return fmt.Sprintf("[POST /api/chart-releases/v3][%d] postApiChartReleasesV3Conflict  %+v", 409, o.Payload)
}
func (o *PostAPIChartReleasesV3Conflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIChartReleasesV3Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIChartReleasesV3InternalServerError creates a PostAPIChartReleasesV3InternalServerError with default headers values
func NewPostAPIChartReleasesV3InternalServerError() *PostAPIChartReleasesV3InternalServerError {
	return &PostAPIChartReleasesV3InternalServerError{}
}

/* PostAPIChartReleasesV3InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostAPIChartReleasesV3InternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIChartReleasesV3InternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/chart-releases/v3][%d] postApiChartReleasesV3InternalServerError  %+v", 500, o.Payload)
}
func (o *PostAPIChartReleasesV3InternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIChartReleasesV3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}