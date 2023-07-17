// Code generated by go-swagger; DO NOT EDIT.

package ci_runs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// GetAPICiRunsV3Reader is a Reader for the GetAPICiRunsV3 structure.
type GetAPICiRunsV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPICiRunsV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPICiRunsV3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPICiRunsV3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPICiRunsV3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPICiRunsV3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPICiRunsV3ProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPICiRunsV3Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPICiRunsV3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPICiRunsV3OK creates a GetAPICiRunsV3OK with default headers values
func NewGetAPICiRunsV3OK() *GetAPICiRunsV3OK {
	return &GetAPICiRunsV3OK{}
}

/* GetAPICiRunsV3OK describes a response with status code 200, with default header values.

OK
*/
type GetAPICiRunsV3OK struct {
	Payload []*models.SherlockCiRunV3
}

func (o *GetAPICiRunsV3OK) Error() string {
	return fmt.Sprintf("[GET /api/ci-runs/v3][%d] getApiCiRunsV3OK  %+v", 200, o.Payload)
}
func (o *GetAPICiRunsV3OK) GetPayload() []*models.SherlockCiRunV3 {
	return o.Payload
}

func (o *GetAPICiRunsV3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPICiRunsV3BadRequest creates a GetAPICiRunsV3BadRequest with default headers values
func NewGetAPICiRunsV3BadRequest() *GetAPICiRunsV3BadRequest {
	return &GetAPICiRunsV3BadRequest{}
}

/* GetAPICiRunsV3BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPICiRunsV3BadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPICiRunsV3BadRequest) Error() string {
	return fmt.Sprintf("[GET /api/ci-runs/v3][%d] getApiCiRunsV3BadRequest  %+v", 400, o.Payload)
}
func (o *GetAPICiRunsV3BadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPICiRunsV3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPICiRunsV3Forbidden creates a GetAPICiRunsV3Forbidden with default headers values
func NewGetAPICiRunsV3Forbidden() *GetAPICiRunsV3Forbidden {
	return &GetAPICiRunsV3Forbidden{}
}

/* GetAPICiRunsV3Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPICiRunsV3Forbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPICiRunsV3Forbidden) Error() string {
	return fmt.Sprintf("[GET /api/ci-runs/v3][%d] getApiCiRunsV3Forbidden  %+v", 403, o.Payload)
}
func (o *GetAPICiRunsV3Forbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPICiRunsV3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPICiRunsV3NotFound creates a GetAPICiRunsV3NotFound with default headers values
func NewGetAPICiRunsV3NotFound() *GetAPICiRunsV3NotFound {
	return &GetAPICiRunsV3NotFound{}
}

/* GetAPICiRunsV3NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPICiRunsV3NotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPICiRunsV3NotFound) Error() string {
	return fmt.Sprintf("[GET /api/ci-runs/v3][%d] getApiCiRunsV3NotFound  %+v", 404, o.Payload)
}
func (o *GetAPICiRunsV3NotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPICiRunsV3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPICiRunsV3ProxyAuthenticationRequired creates a GetAPICiRunsV3ProxyAuthenticationRequired with default headers values
func NewGetAPICiRunsV3ProxyAuthenticationRequired() *GetAPICiRunsV3ProxyAuthenticationRequired {
	return &GetAPICiRunsV3ProxyAuthenticationRequired{}
}

/* GetAPICiRunsV3ProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPICiRunsV3ProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPICiRunsV3ProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/ci-runs/v3][%d] getApiCiRunsV3ProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPICiRunsV3ProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPICiRunsV3ProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPICiRunsV3Conflict creates a GetAPICiRunsV3Conflict with default headers values
func NewGetAPICiRunsV3Conflict() *GetAPICiRunsV3Conflict {
	return &GetAPICiRunsV3Conflict{}
}

/* GetAPICiRunsV3Conflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPICiRunsV3Conflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPICiRunsV3Conflict) Error() string {
	return fmt.Sprintf("[GET /api/ci-runs/v3][%d] getApiCiRunsV3Conflict  %+v", 409, o.Payload)
}
func (o *GetAPICiRunsV3Conflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPICiRunsV3Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPICiRunsV3InternalServerError creates a GetAPICiRunsV3InternalServerError with default headers values
func NewGetAPICiRunsV3InternalServerError() *GetAPICiRunsV3InternalServerError {
	return &GetAPICiRunsV3InternalServerError{}
}

/* GetAPICiRunsV3InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPICiRunsV3InternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPICiRunsV3InternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/ci-runs/v3][%d] getApiCiRunsV3InternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPICiRunsV3InternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPICiRunsV3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
