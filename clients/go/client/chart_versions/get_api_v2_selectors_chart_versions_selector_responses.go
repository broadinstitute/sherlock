// Code generated by go-swagger; DO NOT EDIT.

package chart_versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/clients/go/client/models"
)

// GetAPIV2SelectorsChartVersionsSelectorReader is a Reader for the GetAPIV2SelectorsChartVersionsSelector structure.
type GetAPIV2SelectorsChartVersionsSelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV2SelectorsChartVersionsSelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV2SelectorsChartVersionsSelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIV2SelectorsChartVersionsSelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIV2SelectorsChartVersionsSelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIV2SelectorsChartVersionsSelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPIV2SelectorsChartVersionsSelectorProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPIV2SelectorsChartVersionsSelectorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIV2SelectorsChartVersionsSelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIV2SelectorsChartVersionsSelectorOK creates a GetAPIV2SelectorsChartVersionsSelectorOK with default headers values
func NewGetAPIV2SelectorsChartVersionsSelectorOK() *GetAPIV2SelectorsChartVersionsSelectorOK {
	return &GetAPIV2SelectorsChartVersionsSelectorOK{}
}

/* GetAPIV2SelectorsChartVersionsSelectorOK describes a response with status code 200, with default header values.

OK
*/
type GetAPIV2SelectorsChartVersionsSelectorOK struct {
	Payload []string
}

func (o *GetAPIV2SelectorsChartVersionsSelectorOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/selectors/chart-versions/{selector}][%d] getApiV2SelectorsChartVersionsSelectorOK  %+v", 200, o.Payload)
}
func (o *GetAPIV2SelectorsChartVersionsSelectorOK) GetPayload() []string {
	return o.Payload
}

func (o *GetAPIV2SelectorsChartVersionsSelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2SelectorsChartVersionsSelectorBadRequest creates a GetAPIV2SelectorsChartVersionsSelectorBadRequest with default headers values
func NewGetAPIV2SelectorsChartVersionsSelectorBadRequest() *GetAPIV2SelectorsChartVersionsSelectorBadRequest {
	return &GetAPIV2SelectorsChartVersionsSelectorBadRequest{}
}

/* GetAPIV2SelectorsChartVersionsSelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIV2SelectorsChartVersionsSelectorBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2SelectorsChartVersionsSelectorBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/selectors/chart-versions/{selector}][%d] getApiV2SelectorsChartVersionsSelectorBadRequest  %+v", 400, o.Payload)
}
func (o *GetAPIV2SelectorsChartVersionsSelectorBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2SelectorsChartVersionsSelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2SelectorsChartVersionsSelectorForbidden creates a GetAPIV2SelectorsChartVersionsSelectorForbidden with default headers values
func NewGetAPIV2SelectorsChartVersionsSelectorForbidden() *GetAPIV2SelectorsChartVersionsSelectorForbidden {
	return &GetAPIV2SelectorsChartVersionsSelectorForbidden{}
}

/* GetAPIV2SelectorsChartVersionsSelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPIV2SelectorsChartVersionsSelectorForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2SelectorsChartVersionsSelectorForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/selectors/chart-versions/{selector}][%d] getApiV2SelectorsChartVersionsSelectorForbidden  %+v", 403, o.Payload)
}
func (o *GetAPIV2SelectorsChartVersionsSelectorForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2SelectorsChartVersionsSelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2SelectorsChartVersionsSelectorNotFound creates a GetAPIV2SelectorsChartVersionsSelectorNotFound with default headers values
func NewGetAPIV2SelectorsChartVersionsSelectorNotFound() *GetAPIV2SelectorsChartVersionsSelectorNotFound {
	return &GetAPIV2SelectorsChartVersionsSelectorNotFound{}
}

/* GetAPIV2SelectorsChartVersionsSelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIV2SelectorsChartVersionsSelectorNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2SelectorsChartVersionsSelectorNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/selectors/chart-versions/{selector}][%d] getApiV2SelectorsChartVersionsSelectorNotFound  %+v", 404, o.Payload)
}
func (o *GetAPIV2SelectorsChartVersionsSelectorNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2SelectorsChartVersionsSelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2SelectorsChartVersionsSelectorProxyAuthenticationRequired creates a GetAPIV2SelectorsChartVersionsSelectorProxyAuthenticationRequired with default headers values
func NewGetAPIV2SelectorsChartVersionsSelectorProxyAuthenticationRequired() *GetAPIV2SelectorsChartVersionsSelectorProxyAuthenticationRequired {
	return &GetAPIV2SelectorsChartVersionsSelectorProxyAuthenticationRequired{}
}

/* GetAPIV2SelectorsChartVersionsSelectorProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPIV2SelectorsChartVersionsSelectorProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2SelectorsChartVersionsSelectorProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/v2/selectors/chart-versions/{selector}][%d] getApiV2SelectorsChartVersionsSelectorProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPIV2SelectorsChartVersionsSelectorProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2SelectorsChartVersionsSelectorProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2SelectorsChartVersionsSelectorConflict creates a GetAPIV2SelectorsChartVersionsSelectorConflict with default headers values
func NewGetAPIV2SelectorsChartVersionsSelectorConflict() *GetAPIV2SelectorsChartVersionsSelectorConflict {
	return &GetAPIV2SelectorsChartVersionsSelectorConflict{}
}

/* GetAPIV2SelectorsChartVersionsSelectorConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPIV2SelectorsChartVersionsSelectorConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2SelectorsChartVersionsSelectorConflict) Error() string {
	return fmt.Sprintf("[GET /api/v2/selectors/chart-versions/{selector}][%d] getApiV2SelectorsChartVersionsSelectorConflict  %+v", 409, o.Payload)
}
func (o *GetAPIV2SelectorsChartVersionsSelectorConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2SelectorsChartVersionsSelectorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2SelectorsChartVersionsSelectorInternalServerError creates a GetAPIV2SelectorsChartVersionsSelectorInternalServerError with default headers values
func NewGetAPIV2SelectorsChartVersionsSelectorInternalServerError() *GetAPIV2SelectorsChartVersionsSelectorInternalServerError {
	return &GetAPIV2SelectorsChartVersionsSelectorInternalServerError{}
}

/* GetAPIV2SelectorsChartVersionsSelectorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPIV2SelectorsChartVersionsSelectorInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2SelectorsChartVersionsSelectorInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/selectors/chart-versions/{selector}][%d] getApiV2SelectorsChartVersionsSelectorInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPIV2SelectorsChartVersionsSelectorInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2SelectorsChartVersionsSelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
