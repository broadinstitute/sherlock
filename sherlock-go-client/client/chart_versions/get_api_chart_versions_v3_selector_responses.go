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

// GetAPIChartVersionsV3SelectorReader is a Reader for the GetAPIChartVersionsV3Selector structure.
type GetAPIChartVersionsV3SelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIChartVersionsV3SelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIChartVersionsV3SelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIChartVersionsV3SelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIChartVersionsV3SelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIChartVersionsV3SelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPIChartVersionsV3SelectorProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPIChartVersionsV3SelectorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIChartVersionsV3SelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIChartVersionsV3SelectorOK creates a GetAPIChartVersionsV3SelectorOK with default headers values
func NewGetAPIChartVersionsV3SelectorOK() *GetAPIChartVersionsV3SelectorOK {
	return &GetAPIChartVersionsV3SelectorOK{}
}

/* GetAPIChartVersionsV3SelectorOK describes a response with status code 200, with default header values.

OK
*/
type GetAPIChartVersionsV3SelectorOK struct {
	Payload *models.SherlockChartVersionV3
}

func (o *GetAPIChartVersionsV3SelectorOK) Error() string {
	return fmt.Sprintf("[GET /api/chart-versions/v3/{selector}][%d] getApiChartVersionsV3SelectorOK  %+v", 200, o.Payload)
}
func (o *GetAPIChartVersionsV3SelectorOK) GetPayload() *models.SherlockChartVersionV3 {
	return o.Payload
}

func (o *GetAPIChartVersionsV3SelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SherlockChartVersionV3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChartVersionsV3SelectorBadRequest creates a GetAPIChartVersionsV3SelectorBadRequest with default headers values
func NewGetAPIChartVersionsV3SelectorBadRequest() *GetAPIChartVersionsV3SelectorBadRequest {
	return &GetAPIChartVersionsV3SelectorBadRequest{}
}

/* GetAPIChartVersionsV3SelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIChartVersionsV3SelectorBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChartVersionsV3SelectorBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/chart-versions/v3/{selector}][%d] getApiChartVersionsV3SelectorBadRequest  %+v", 400, o.Payload)
}
func (o *GetAPIChartVersionsV3SelectorBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChartVersionsV3SelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChartVersionsV3SelectorForbidden creates a GetAPIChartVersionsV3SelectorForbidden with default headers values
func NewGetAPIChartVersionsV3SelectorForbidden() *GetAPIChartVersionsV3SelectorForbidden {
	return &GetAPIChartVersionsV3SelectorForbidden{}
}

/* GetAPIChartVersionsV3SelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPIChartVersionsV3SelectorForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChartVersionsV3SelectorForbidden) Error() string {
	return fmt.Sprintf("[GET /api/chart-versions/v3/{selector}][%d] getApiChartVersionsV3SelectorForbidden  %+v", 403, o.Payload)
}
func (o *GetAPIChartVersionsV3SelectorForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChartVersionsV3SelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChartVersionsV3SelectorNotFound creates a GetAPIChartVersionsV3SelectorNotFound with default headers values
func NewGetAPIChartVersionsV3SelectorNotFound() *GetAPIChartVersionsV3SelectorNotFound {
	return &GetAPIChartVersionsV3SelectorNotFound{}
}

/* GetAPIChartVersionsV3SelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIChartVersionsV3SelectorNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChartVersionsV3SelectorNotFound) Error() string {
	return fmt.Sprintf("[GET /api/chart-versions/v3/{selector}][%d] getApiChartVersionsV3SelectorNotFound  %+v", 404, o.Payload)
}
func (o *GetAPIChartVersionsV3SelectorNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChartVersionsV3SelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChartVersionsV3SelectorProxyAuthenticationRequired creates a GetAPIChartVersionsV3SelectorProxyAuthenticationRequired with default headers values
func NewGetAPIChartVersionsV3SelectorProxyAuthenticationRequired() *GetAPIChartVersionsV3SelectorProxyAuthenticationRequired {
	return &GetAPIChartVersionsV3SelectorProxyAuthenticationRequired{}
}

/* GetAPIChartVersionsV3SelectorProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPIChartVersionsV3SelectorProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChartVersionsV3SelectorProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/chart-versions/v3/{selector}][%d] getApiChartVersionsV3SelectorProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPIChartVersionsV3SelectorProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChartVersionsV3SelectorProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChartVersionsV3SelectorConflict creates a GetAPIChartVersionsV3SelectorConflict with default headers values
func NewGetAPIChartVersionsV3SelectorConflict() *GetAPIChartVersionsV3SelectorConflict {
	return &GetAPIChartVersionsV3SelectorConflict{}
}

/* GetAPIChartVersionsV3SelectorConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPIChartVersionsV3SelectorConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChartVersionsV3SelectorConflict) Error() string {
	return fmt.Sprintf("[GET /api/chart-versions/v3/{selector}][%d] getApiChartVersionsV3SelectorConflict  %+v", 409, o.Payload)
}
func (o *GetAPIChartVersionsV3SelectorConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChartVersionsV3SelectorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChartVersionsV3SelectorInternalServerError creates a GetAPIChartVersionsV3SelectorInternalServerError with default headers values
func NewGetAPIChartVersionsV3SelectorInternalServerError() *GetAPIChartVersionsV3SelectorInternalServerError {
	return &GetAPIChartVersionsV3SelectorInternalServerError{}
}

/* GetAPIChartVersionsV3SelectorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPIChartVersionsV3SelectorInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChartVersionsV3SelectorInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/chart-versions/v3/{selector}][%d] getApiChartVersionsV3SelectorInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPIChartVersionsV3SelectorInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChartVersionsV3SelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
