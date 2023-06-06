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

// GetAPIV2ChartReleasesReader is a Reader for the GetAPIV2ChartReleases structure.
type GetAPIV2ChartReleasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV2ChartReleasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV2ChartReleasesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIV2ChartReleasesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIV2ChartReleasesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIV2ChartReleasesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPIV2ChartReleasesProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPIV2ChartReleasesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIV2ChartReleasesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIV2ChartReleasesOK creates a GetAPIV2ChartReleasesOK with default headers values
func NewGetAPIV2ChartReleasesOK() *GetAPIV2ChartReleasesOK {
	return &GetAPIV2ChartReleasesOK{}
}

/* GetAPIV2ChartReleasesOK describes a response with status code 200, with default header values.

OK
*/
type GetAPIV2ChartReleasesOK struct {
	Payload []*models.V2controllersChartRelease
}

func (o *GetAPIV2ChartReleasesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/chart-releases][%d] getApiV2ChartReleasesOK  %+v", 200, o.Payload)
}
func (o *GetAPIV2ChartReleasesOK) GetPayload() []*models.V2controllersChartRelease {
	return o.Payload
}

func (o *GetAPIV2ChartReleasesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ChartReleasesBadRequest creates a GetAPIV2ChartReleasesBadRequest with default headers values
func NewGetAPIV2ChartReleasesBadRequest() *GetAPIV2ChartReleasesBadRequest {
	return &GetAPIV2ChartReleasesBadRequest{}
}

/* GetAPIV2ChartReleasesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIV2ChartReleasesBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ChartReleasesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/chart-releases][%d] getApiV2ChartReleasesBadRequest  %+v", 400, o.Payload)
}
func (o *GetAPIV2ChartReleasesBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ChartReleasesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ChartReleasesForbidden creates a GetAPIV2ChartReleasesForbidden with default headers values
func NewGetAPIV2ChartReleasesForbidden() *GetAPIV2ChartReleasesForbidden {
	return &GetAPIV2ChartReleasesForbidden{}
}

/* GetAPIV2ChartReleasesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPIV2ChartReleasesForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ChartReleasesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/chart-releases][%d] getApiV2ChartReleasesForbidden  %+v", 403, o.Payload)
}
func (o *GetAPIV2ChartReleasesForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ChartReleasesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ChartReleasesNotFound creates a GetAPIV2ChartReleasesNotFound with default headers values
func NewGetAPIV2ChartReleasesNotFound() *GetAPIV2ChartReleasesNotFound {
	return &GetAPIV2ChartReleasesNotFound{}
}

/* GetAPIV2ChartReleasesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIV2ChartReleasesNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ChartReleasesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/chart-releases][%d] getApiV2ChartReleasesNotFound  %+v", 404, o.Payload)
}
func (o *GetAPIV2ChartReleasesNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ChartReleasesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ChartReleasesProxyAuthenticationRequired creates a GetAPIV2ChartReleasesProxyAuthenticationRequired with default headers values
func NewGetAPIV2ChartReleasesProxyAuthenticationRequired() *GetAPIV2ChartReleasesProxyAuthenticationRequired {
	return &GetAPIV2ChartReleasesProxyAuthenticationRequired{}
}

/* GetAPIV2ChartReleasesProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPIV2ChartReleasesProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ChartReleasesProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/v2/chart-releases][%d] getApiV2ChartReleasesProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPIV2ChartReleasesProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ChartReleasesProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ChartReleasesConflict creates a GetAPIV2ChartReleasesConflict with default headers values
func NewGetAPIV2ChartReleasesConflict() *GetAPIV2ChartReleasesConflict {
	return &GetAPIV2ChartReleasesConflict{}
}

/* GetAPIV2ChartReleasesConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPIV2ChartReleasesConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ChartReleasesConflict) Error() string {
	return fmt.Sprintf("[GET /api/v2/chart-releases][%d] getApiV2ChartReleasesConflict  %+v", 409, o.Payload)
}
func (o *GetAPIV2ChartReleasesConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ChartReleasesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ChartReleasesInternalServerError creates a GetAPIV2ChartReleasesInternalServerError with default headers values
func NewGetAPIV2ChartReleasesInternalServerError() *GetAPIV2ChartReleasesInternalServerError {
	return &GetAPIV2ChartReleasesInternalServerError{}
}

/* GetAPIV2ChartReleasesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPIV2ChartReleasesInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ChartReleasesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/chart-releases][%d] getApiV2ChartReleasesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPIV2ChartReleasesInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ChartReleasesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
