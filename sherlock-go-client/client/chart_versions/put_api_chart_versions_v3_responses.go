// Code generated by go-swagger; DO NOT EDIT.

package chart_versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// PutAPIChartVersionsV3Reader is a Reader for the PutAPIChartVersionsV3 structure.
type PutAPIChartVersionsV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAPIChartVersionsV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPutAPIChartVersionsV3Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutAPIChartVersionsV3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutAPIChartVersionsV3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutAPIChartVersionsV3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewPutAPIChartVersionsV3ProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutAPIChartVersionsV3Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutAPIChartVersionsV3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutAPIChartVersionsV3Created creates a PutAPIChartVersionsV3Created with default headers values
func NewPutAPIChartVersionsV3Created() *PutAPIChartVersionsV3Created {
	return &PutAPIChartVersionsV3Created{}
}

/* PutAPIChartVersionsV3Created describes a response with status code 201, with default header values.

Created
*/
type PutAPIChartVersionsV3Created struct {
	Payload *models.SherlockChartVersionV3
}

func (o *PutAPIChartVersionsV3Created) Error() string {
	return fmt.Sprintf("[PUT /api/chart-versions/v3][%d] putApiChartVersionsV3Created  %+v", 201, o.Payload)
}
func (o *PutAPIChartVersionsV3Created) GetPayload() *models.SherlockChartVersionV3 {
	return o.Payload
}

func (o *PutAPIChartVersionsV3Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SherlockChartVersionV3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIChartVersionsV3BadRequest creates a PutAPIChartVersionsV3BadRequest with default headers values
func NewPutAPIChartVersionsV3BadRequest() *PutAPIChartVersionsV3BadRequest {
	return &PutAPIChartVersionsV3BadRequest{}
}

/* PutAPIChartVersionsV3BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutAPIChartVersionsV3BadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PutAPIChartVersionsV3BadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/chart-versions/v3][%d] putApiChartVersionsV3BadRequest  %+v", 400, o.Payload)
}
func (o *PutAPIChartVersionsV3BadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PutAPIChartVersionsV3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIChartVersionsV3Forbidden creates a PutAPIChartVersionsV3Forbidden with default headers values
func NewPutAPIChartVersionsV3Forbidden() *PutAPIChartVersionsV3Forbidden {
	return &PutAPIChartVersionsV3Forbidden{}
}

/* PutAPIChartVersionsV3Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutAPIChartVersionsV3Forbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PutAPIChartVersionsV3Forbidden) Error() string {
	return fmt.Sprintf("[PUT /api/chart-versions/v3][%d] putApiChartVersionsV3Forbidden  %+v", 403, o.Payload)
}
func (o *PutAPIChartVersionsV3Forbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PutAPIChartVersionsV3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIChartVersionsV3NotFound creates a PutAPIChartVersionsV3NotFound with default headers values
func NewPutAPIChartVersionsV3NotFound() *PutAPIChartVersionsV3NotFound {
	return &PutAPIChartVersionsV3NotFound{}
}

/* PutAPIChartVersionsV3NotFound describes a response with status code 404, with default header values.

Not Found
*/
type PutAPIChartVersionsV3NotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PutAPIChartVersionsV3NotFound) Error() string {
	return fmt.Sprintf("[PUT /api/chart-versions/v3][%d] putApiChartVersionsV3NotFound  %+v", 404, o.Payload)
}
func (o *PutAPIChartVersionsV3NotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PutAPIChartVersionsV3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIChartVersionsV3ProxyAuthenticationRequired creates a PutAPIChartVersionsV3ProxyAuthenticationRequired with default headers values
func NewPutAPIChartVersionsV3ProxyAuthenticationRequired() *PutAPIChartVersionsV3ProxyAuthenticationRequired {
	return &PutAPIChartVersionsV3ProxyAuthenticationRequired{}
}

/* PutAPIChartVersionsV3ProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type PutAPIChartVersionsV3ProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PutAPIChartVersionsV3ProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[PUT /api/chart-versions/v3][%d] putApiChartVersionsV3ProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *PutAPIChartVersionsV3ProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PutAPIChartVersionsV3ProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIChartVersionsV3Conflict creates a PutAPIChartVersionsV3Conflict with default headers values
func NewPutAPIChartVersionsV3Conflict() *PutAPIChartVersionsV3Conflict {
	return &PutAPIChartVersionsV3Conflict{}
}

/* PutAPIChartVersionsV3Conflict describes a response with status code 409, with default header values.

Conflict
*/
type PutAPIChartVersionsV3Conflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PutAPIChartVersionsV3Conflict) Error() string {
	return fmt.Sprintf("[PUT /api/chart-versions/v3][%d] putApiChartVersionsV3Conflict  %+v", 409, o.Payload)
}
func (o *PutAPIChartVersionsV3Conflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PutAPIChartVersionsV3Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIChartVersionsV3InternalServerError creates a PutAPIChartVersionsV3InternalServerError with default headers values
func NewPutAPIChartVersionsV3InternalServerError() *PutAPIChartVersionsV3InternalServerError {
	return &PutAPIChartVersionsV3InternalServerError{}
}

/* PutAPIChartVersionsV3InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PutAPIChartVersionsV3InternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PutAPIChartVersionsV3InternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/chart-versions/v3][%d] putApiChartVersionsV3InternalServerError  %+v", 500, o.Payload)
}
func (o *PutAPIChartVersionsV3InternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PutAPIChartVersionsV3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}