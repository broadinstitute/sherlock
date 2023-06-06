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

// GetAPIV2ChartReleasesSelectorReader is a Reader for the GetAPIV2ChartReleasesSelector structure.
type GetAPIV2ChartReleasesSelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV2ChartReleasesSelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV2ChartReleasesSelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIV2ChartReleasesSelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIV2ChartReleasesSelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIV2ChartReleasesSelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPIV2ChartReleasesSelectorProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPIV2ChartReleasesSelectorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIV2ChartReleasesSelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIV2ChartReleasesSelectorOK creates a GetAPIV2ChartReleasesSelectorOK with default headers values
func NewGetAPIV2ChartReleasesSelectorOK() *GetAPIV2ChartReleasesSelectorOK {
	return &GetAPIV2ChartReleasesSelectorOK{}
}

/* GetAPIV2ChartReleasesSelectorOK describes a response with status code 200, with default header values.

OK
*/
type GetAPIV2ChartReleasesSelectorOK struct {
	Payload *models.V2controllersChartRelease
}

func (o *GetAPIV2ChartReleasesSelectorOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/chart-releases/{selector}][%d] getApiV2ChartReleasesSelectorOK  %+v", 200, o.Payload)
}
func (o *GetAPIV2ChartReleasesSelectorOK) GetPayload() *models.V2controllersChartRelease {
	return o.Payload
}

func (o *GetAPIV2ChartReleasesSelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V2controllersChartRelease)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ChartReleasesSelectorBadRequest creates a GetAPIV2ChartReleasesSelectorBadRequest with default headers values
func NewGetAPIV2ChartReleasesSelectorBadRequest() *GetAPIV2ChartReleasesSelectorBadRequest {
	return &GetAPIV2ChartReleasesSelectorBadRequest{}
}

/* GetAPIV2ChartReleasesSelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIV2ChartReleasesSelectorBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ChartReleasesSelectorBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/chart-releases/{selector}][%d] getApiV2ChartReleasesSelectorBadRequest  %+v", 400, o.Payload)
}
func (o *GetAPIV2ChartReleasesSelectorBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ChartReleasesSelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ChartReleasesSelectorForbidden creates a GetAPIV2ChartReleasesSelectorForbidden with default headers values
func NewGetAPIV2ChartReleasesSelectorForbidden() *GetAPIV2ChartReleasesSelectorForbidden {
	return &GetAPIV2ChartReleasesSelectorForbidden{}
}

/* GetAPIV2ChartReleasesSelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPIV2ChartReleasesSelectorForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ChartReleasesSelectorForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/chart-releases/{selector}][%d] getApiV2ChartReleasesSelectorForbidden  %+v", 403, o.Payload)
}
func (o *GetAPIV2ChartReleasesSelectorForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ChartReleasesSelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ChartReleasesSelectorNotFound creates a GetAPIV2ChartReleasesSelectorNotFound with default headers values
func NewGetAPIV2ChartReleasesSelectorNotFound() *GetAPIV2ChartReleasesSelectorNotFound {
	return &GetAPIV2ChartReleasesSelectorNotFound{}
}

/* GetAPIV2ChartReleasesSelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIV2ChartReleasesSelectorNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ChartReleasesSelectorNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/chart-releases/{selector}][%d] getApiV2ChartReleasesSelectorNotFound  %+v", 404, o.Payload)
}
func (o *GetAPIV2ChartReleasesSelectorNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ChartReleasesSelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ChartReleasesSelectorProxyAuthenticationRequired creates a GetAPIV2ChartReleasesSelectorProxyAuthenticationRequired with default headers values
func NewGetAPIV2ChartReleasesSelectorProxyAuthenticationRequired() *GetAPIV2ChartReleasesSelectorProxyAuthenticationRequired {
	return &GetAPIV2ChartReleasesSelectorProxyAuthenticationRequired{}
}

/* GetAPIV2ChartReleasesSelectorProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPIV2ChartReleasesSelectorProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ChartReleasesSelectorProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/v2/chart-releases/{selector}][%d] getApiV2ChartReleasesSelectorProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPIV2ChartReleasesSelectorProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ChartReleasesSelectorProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ChartReleasesSelectorConflict creates a GetAPIV2ChartReleasesSelectorConflict with default headers values
func NewGetAPIV2ChartReleasesSelectorConflict() *GetAPIV2ChartReleasesSelectorConflict {
	return &GetAPIV2ChartReleasesSelectorConflict{}
}

/* GetAPIV2ChartReleasesSelectorConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPIV2ChartReleasesSelectorConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ChartReleasesSelectorConflict) Error() string {
	return fmt.Sprintf("[GET /api/v2/chart-releases/{selector}][%d] getApiV2ChartReleasesSelectorConflict  %+v", 409, o.Payload)
}
func (o *GetAPIV2ChartReleasesSelectorConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ChartReleasesSelectorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ChartReleasesSelectorInternalServerError creates a GetAPIV2ChartReleasesSelectorInternalServerError with default headers values
func NewGetAPIV2ChartReleasesSelectorInternalServerError() *GetAPIV2ChartReleasesSelectorInternalServerError {
	return &GetAPIV2ChartReleasesSelectorInternalServerError{}
}

/* GetAPIV2ChartReleasesSelectorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPIV2ChartReleasesSelectorInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ChartReleasesSelectorInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/chart-releases/{selector}][%d] getApiV2ChartReleasesSelectorInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPIV2ChartReleasesSelectorInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ChartReleasesSelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
