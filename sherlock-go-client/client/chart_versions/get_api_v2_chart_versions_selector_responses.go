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

// GetAPIV2ChartVersionsSelectorReader is a Reader for the GetAPIV2ChartVersionsSelector structure.
type GetAPIV2ChartVersionsSelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV2ChartVersionsSelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV2ChartVersionsSelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIV2ChartVersionsSelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIV2ChartVersionsSelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIV2ChartVersionsSelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPIV2ChartVersionsSelectorProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPIV2ChartVersionsSelectorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIV2ChartVersionsSelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIV2ChartVersionsSelectorOK creates a GetAPIV2ChartVersionsSelectorOK with default headers values
func NewGetAPIV2ChartVersionsSelectorOK() *GetAPIV2ChartVersionsSelectorOK {
	return &GetAPIV2ChartVersionsSelectorOK{}
}

/* GetAPIV2ChartVersionsSelectorOK describes a response with status code 200, with default header values.

OK
*/
type GetAPIV2ChartVersionsSelectorOK struct {
	Payload *models.V2controllersChartVersion
}

func (o *GetAPIV2ChartVersionsSelectorOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/chart-versions/{selector}][%d] getApiV2ChartVersionsSelectorOK  %+v", 200, o.Payload)
}
func (o *GetAPIV2ChartVersionsSelectorOK) GetPayload() *models.V2controllersChartVersion {
	return o.Payload
}

func (o *GetAPIV2ChartVersionsSelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V2controllersChartVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ChartVersionsSelectorBadRequest creates a GetAPIV2ChartVersionsSelectorBadRequest with default headers values
func NewGetAPIV2ChartVersionsSelectorBadRequest() *GetAPIV2ChartVersionsSelectorBadRequest {
	return &GetAPIV2ChartVersionsSelectorBadRequest{}
}

/* GetAPIV2ChartVersionsSelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIV2ChartVersionsSelectorBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ChartVersionsSelectorBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/chart-versions/{selector}][%d] getApiV2ChartVersionsSelectorBadRequest  %+v", 400, o.Payload)
}
func (o *GetAPIV2ChartVersionsSelectorBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ChartVersionsSelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ChartVersionsSelectorForbidden creates a GetAPIV2ChartVersionsSelectorForbidden with default headers values
func NewGetAPIV2ChartVersionsSelectorForbidden() *GetAPIV2ChartVersionsSelectorForbidden {
	return &GetAPIV2ChartVersionsSelectorForbidden{}
}

/* GetAPIV2ChartVersionsSelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPIV2ChartVersionsSelectorForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ChartVersionsSelectorForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/chart-versions/{selector}][%d] getApiV2ChartVersionsSelectorForbidden  %+v", 403, o.Payload)
}
func (o *GetAPIV2ChartVersionsSelectorForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ChartVersionsSelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ChartVersionsSelectorNotFound creates a GetAPIV2ChartVersionsSelectorNotFound with default headers values
func NewGetAPIV2ChartVersionsSelectorNotFound() *GetAPIV2ChartVersionsSelectorNotFound {
	return &GetAPIV2ChartVersionsSelectorNotFound{}
}

/* GetAPIV2ChartVersionsSelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIV2ChartVersionsSelectorNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ChartVersionsSelectorNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/chart-versions/{selector}][%d] getApiV2ChartVersionsSelectorNotFound  %+v", 404, o.Payload)
}
func (o *GetAPIV2ChartVersionsSelectorNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ChartVersionsSelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ChartVersionsSelectorProxyAuthenticationRequired creates a GetAPIV2ChartVersionsSelectorProxyAuthenticationRequired with default headers values
func NewGetAPIV2ChartVersionsSelectorProxyAuthenticationRequired() *GetAPIV2ChartVersionsSelectorProxyAuthenticationRequired {
	return &GetAPIV2ChartVersionsSelectorProxyAuthenticationRequired{}
}

/* GetAPIV2ChartVersionsSelectorProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPIV2ChartVersionsSelectorProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ChartVersionsSelectorProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/v2/chart-versions/{selector}][%d] getApiV2ChartVersionsSelectorProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPIV2ChartVersionsSelectorProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ChartVersionsSelectorProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ChartVersionsSelectorConflict creates a GetAPIV2ChartVersionsSelectorConflict with default headers values
func NewGetAPIV2ChartVersionsSelectorConflict() *GetAPIV2ChartVersionsSelectorConflict {
	return &GetAPIV2ChartVersionsSelectorConflict{}
}

/* GetAPIV2ChartVersionsSelectorConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPIV2ChartVersionsSelectorConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ChartVersionsSelectorConflict) Error() string {
	return fmt.Sprintf("[GET /api/v2/chart-versions/{selector}][%d] getApiV2ChartVersionsSelectorConflict  %+v", 409, o.Payload)
}
func (o *GetAPIV2ChartVersionsSelectorConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ChartVersionsSelectorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ChartVersionsSelectorInternalServerError creates a GetAPIV2ChartVersionsSelectorInternalServerError with default headers values
func NewGetAPIV2ChartVersionsSelectorInternalServerError() *GetAPIV2ChartVersionsSelectorInternalServerError {
	return &GetAPIV2ChartVersionsSelectorInternalServerError{}
}

/* GetAPIV2ChartVersionsSelectorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPIV2ChartVersionsSelectorInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ChartVersionsSelectorInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/chart-versions/{selector}][%d] getApiV2ChartVersionsSelectorInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPIV2ChartVersionsSelectorInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ChartVersionsSelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}