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

// GetAPIChartReleasesV3SelectorReader is a Reader for the GetAPIChartReleasesV3Selector structure.
type GetAPIChartReleasesV3SelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIChartReleasesV3SelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIChartReleasesV3SelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIChartReleasesV3SelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIChartReleasesV3SelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIChartReleasesV3SelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPIChartReleasesV3SelectorProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPIChartReleasesV3SelectorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIChartReleasesV3SelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIChartReleasesV3SelectorOK creates a GetAPIChartReleasesV3SelectorOK with default headers values
func NewGetAPIChartReleasesV3SelectorOK() *GetAPIChartReleasesV3SelectorOK {
	return &GetAPIChartReleasesV3SelectorOK{}
}

/* GetAPIChartReleasesV3SelectorOK describes a response with status code 200, with default header values.

OK
*/
type GetAPIChartReleasesV3SelectorOK struct {
	Payload *models.SherlockChartReleaseV3
}

func (o *GetAPIChartReleasesV3SelectorOK) Error() string {
	return fmt.Sprintf("[GET /api/chart-releases/v3/{selector}][%d] getApiChartReleasesV3SelectorOK  %+v", 200, o.Payload)
}
func (o *GetAPIChartReleasesV3SelectorOK) GetPayload() *models.SherlockChartReleaseV3 {
	return o.Payload
}

func (o *GetAPIChartReleasesV3SelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SherlockChartReleaseV3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChartReleasesV3SelectorBadRequest creates a GetAPIChartReleasesV3SelectorBadRequest with default headers values
func NewGetAPIChartReleasesV3SelectorBadRequest() *GetAPIChartReleasesV3SelectorBadRequest {
	return &GetAPIChartReleasesV3SelectorBadRequest{}
}

/* GetAPIChartReleasesV3SelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIChartReleasesV3SelectorBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChartReleasesV3SelectorBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/chart-releases/v3/{selector}][%d] getApiChartReleasesV3SelectorBadRequest  %+v", 400, o.Payload)
}
func (o *GetAPIChartReleasesV3SelectorBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChartReleasesV3SelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChartReleasesV3SelectorForbidden creates a GetAPIChartReleasesV3SelectorForbidden with default headers values
func NewGetAPIChartReleasesV3SelectorForbidden() *GetAPIChartReleasesV3SelectorForbidden {
	return &GetAPIChartReleasesV3SelectorForbidden{}
}

/* GetAPIChartReleasesV3SelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPIChartReleasesV3SelectorForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChartReleasesV3SelectorForbidden) Error() string {
	return fmt.Sprintf("[GET /api/chart-releases/v3/{selector}][%d] getApiChartReleasesV3SelectorForbidden  %+v", 403, o.Payload)
}
func (o *GetAPIChartReleasesV3SelectorForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChartReleasesV3SelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChartReleasesV3SelectorNotFound creates a GetAPIChartReleasesV3SelectorNotFound with default headers values
func NewGetAPIChartReleasesV3SelectorNotFound() *GetAPIChartReleasesV3SelectorNotFound {
	return &GetAPIChartReleasesV3SelectorNotFound{}
}

/* GetAPIChartReleasesV3SelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIChartReleasesV3SelectorNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChartReleasesV3SelectorNotFound) Error() string {
	return fmt.Sprintf("[GET /api/chart-releases/v3/{selector}][%d] getApiChartReleasesV3SelectorNotFound  %+v", 404, o.Payload)
}
func (o *GetAPIChartReleasesV3SelectorNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChartReleasesV3SelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChartReleasesV3SelectorProxyAuthenticationRequired creates a GetAPIChartReleasesV3SelectorProxyAuthenticationRequired with default headers values
func NewGetAPIChartReleasesV3SelectorProxyAuthenticationRequired() *GetAPIChartReleasesV3SelectorProxyAuthenticationRequired {
	return &GetAPIChartReleasesV3SelectorProxyAuthenticationRequired{}
}

/* GetAPIChartReleasesV3SelectorProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPIChartReleasesV3SelectorProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChartReleasesV3SelectorProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/chart-releases/v3/{selector}][%d] getApiChartReleasesV3SelectorProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPIChartReleasesV3SelectorProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChartReleasesV3SelectorProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChartReleasesV3SelectorConflict creates a GetAPIChartReleasesV3SelectorConflict with default headers values
func NewGetAPIChartReleasesV3SelectorConflict() *GetAPIChartReleasesV3SelectorConflict {
	return &GetAPIChartReleasesV3SelectorConflict{}
}

/* GetAPIChartReleasesV3SelectorConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPIChartReleasesV3SelectorConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChartReleasesV3SelectorConflict) Error() string {
	return fmt.Sprintf("[GET /api/chart-releases/v3/{selector}][%d] getApiChartReleasesV3SelectorConflict  %+v", 409, o.Payload)
}
func (o *GetAPIChartReleasesV3SelectorConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChartReleasesV3SelectorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChartReleasesV3SelectorInternalServerError creates a GetAPIChartReleasesV3SelectorInternalServerError with default headers values
func NewGetAPIChartReleasesV3SelectorInternalServerError() *GetAPIChartReleasesV3SelectorInternalServerError {
	return &GetAPIChartReleasesV3SelectorInternalServerError{}
}

/* GetAPIChartReleasesV3SelectorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPIChartReleasesV3SelectorInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChartReleasesV3SelectorInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/chart-releases/v3/{selector}][%d] getApiChartReleasesV3SelectorInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPIChartReleasesV3SelectorInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChartReleasesV3SelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
