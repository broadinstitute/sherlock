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

// GetAPIChartReleasesV3Reader is a Reader for the GetAPIChartReleasesV3 structure.
type GetAPIChartReleasesV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIChartReleasesV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIChartReleasesV3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIChartReleasesV3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIChartReleasesV3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIChartReleasesV3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPIChartReleasesV3ProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPIChartReleasesV3Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIChartReleasesV3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIChartReleasesV3OK creates a GetAPIChartReleasesV3OK with default headers values
func NewGetAPIChartReleasesV3OK() *GetAPIChartReleasesV3OK {
	return &GetAPIChartReleasesV3OK{}
}

/* GetAPIChartReleasesV3OK describes a response with status code 200, with default header values.

OK
*/
type GetAPIChartReleasesV3OK struct {
	Payload []*models.SherlockChartReleaseV3
}

func (o *GetAPIChartReleasesV3OK) Error() string {
	return fmt.Sprintf("[GET /api/chart-releases/v3][%d] getApiChartReleasesV3OK  %+v", 200, o.Payload)
}
func (o *GetAPIChartReleasesV3OK) GetPayload() []*models.SherlockChartReleaseV3 {
	return o.Payload
}

func (o *GetAPIChartReleasesV3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChartReleasesV3BadRequest creates a GetAPIChartReleasesV3BadRequest with default headers values
func NewGetAPIChartReleasesV3BadRequest() *GetAPIChartReleasesV3BadRequest {
	return &GetAPIChartReleasesV3BadRequest{}
}

/* GetAPIChartReleasesV3BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIChartReleasesV3BadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChartReleasesV3BadRequest) Error() string {
	return fmt.Sprintf("[GET /api/chart-releases/v3][%d] getApiChartReleasesV3BadRequest  %+v", 400, o.Payload)
}
func (o *GetAPIChartReleasesV3BadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChartReleasesV3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChartReleasesV3Forbidden creates a GetAPIChartReleasesV3Forbidden with default headers values
func NewGetAPIChartReleasesV3Forbidden() *GetAPIChartReleasesV3Forbidden {
	return &GetAPIChartReleasesV3Forbidden{}
}

/* GetAPIChartReleasesV3Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPIChartReleasesV3Forbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChartReleasesV3Forbidden) Error() string {
	return fmt.Sprintf("[GET /api/chart-releases/v3][%d] getApiChartReleasesV3Forbidden  %+v", 403, o.Payload)
}
func (o *GetAPIChartReleasesV3Forbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChartReleasesV3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChartReleasesV3NotFound creates a GetAPIChartReleasesV3NotFound with default headers values
func NewGetAPIChartReleasesV3NotFound() *GetAPIChartReleasesV3NotFound {
	return &GetAPIChartReleasesV3NotFound{}
}

/* GetAPIChartReleasesV3NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIChartReleasesV3NotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChartReleasesV3NotFound) Error() string {
	return fmt.Sprintf("[GET /api/chart-releases/v3][%d] getApiChartReleasesV3NotFound  %+v", 404, o.Payload)
}
func (o *GetAPIChartReleasesV3NotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChartReleasesV3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChartReleasesV3ProxyAuthenticationRequired creates a GetAPIChartReleasesV3ProxyAuthenticationRequired with default headers values
func NewGetAPIChartReleasesV3ProxyAuthenticationRequired() *GetAPIChartReleasesV3ProxyAuthenticationRequired {
	return &GetAPIChartReleasesV3ProxyAuthenticationRequired{}
}

/* GetAPIChartReleasesV3ProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPIChartReleasesV3ProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChartReleasesV3ProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/chart-releases/v3][%d] getApiChartReleasesV3ProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPIChartReleasesV3ProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChartReleasesV3ProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChartReleasesV3Conflict creates a GetAPIChartReleasesV3Conflict with default headers values
func NewGetAPIChartReleasesV3Conflict() *GetAPIChartReleasesV3Conflict {
	return &GetAPIChartReleasesV3Conflict{}
}

/* GetAPIChartReleasesV3Conflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPIChartReleasesV3Conflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChartReleasesV3Conflict) Error() string {
	return fmt.Sprintf("[GET /api/chart-releases/v3][%d] getApiChartReleasesV3Conflict  %+v", 409, o.Payload)
}
func (o *GetAPIChartReleasesV3Conflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChartReleasesV3Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChartReleasesV3InternalServerError creates a GetAPIChartReleasesV3InternalServerError with default headers values
func NewGetAPIChartReleasesV3InternalServerError() *GetAPIChartReleasesV3InternalServerError {
	return &GetAPIChartReleasesV3InternalServerError{}
}

/* GetAPIChartReleasesV3InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPIChartReleasesV3InternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChartReleasesV3InternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/chart-releases/v3][%d] getApiChartReleasesV3InternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPIChartReleasesV3InternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChartReleasesV3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
