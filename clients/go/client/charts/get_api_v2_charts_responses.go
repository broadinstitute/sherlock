// Code generated by go-swagger; DO NOT EDIT.

package charts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/clients/go/client/models"
)

// GetAPIV2ChartsReader is a Reader for the GetAPIV2Charts structure.
type GetAPIV2ChartsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV2ChartsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV2ChartsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIV2ChartsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIV2ChartsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIV2ChartsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPIV2ChartsProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPIV2ChartsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIV2ChartsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIV2ChartsOK creates a GetAPIV2ChartsOK with default headers values
func NewGetAPIV2ChartsOK() *GetAPIV2ChartsOK {
	return &GetAPIV2ChartsOK{}
}

/*
GetAPIV2ChartsOK describes a response with status code 200, with default header values.

OK
*/
type GetAPIV2ChartsOK struct {
	Payload []*models.V2controllersChart
}

// IsSuccess returns true when this get Api v2 charts o k response has a 2xx status code
func (o *GetAPIV2ChartsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api v2 charts o k response has a 3xx status code
func (o *GetAPIV2ChartsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api v2 charts o k response has a 4xx status code
func (o *GetAPIV2ChartsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api v2 charts o k response has a 5xx status code
func (o *GetAPIV2ChartsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api v2 charts o k response a status code equal to that given
func (o *GetAPIV2ChartsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAPIV2ChartsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/charts][%d] getApiV2ChartsOK  %+v", 200, o.Payload)
}

func (o *GetAPIV2ChartsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/charts][%d] getApiV2ChartsOK  %+v", 200, o.Payload)
}

func (o *GetAPIV2ChartsOK) GetPayload() []*models.V2controllersChart {
	return o.Payload
}

func (o *GetAPIV2ChartsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ChartsBadRequest creates a GetAPIV2ChartsBadRequest with default headers values
func NewGetAPIV2ChartsBadRequest() *GetAPIV2ChartsBadRequest {
	return &GetAPIV2ChartsBadRequest{}
}

/*
GetAPIV2ChartsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIV2ChartsBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

// IsSuccess returns true when this get Api v2 charts bad request response has a 2xx status code
func (o *GetAPIV2ChartsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Api v2 charts bad request response has a 3xx status code
func (o *GetAPIV2ChartsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api v2 charts bad request response has a 4xx status code
func (o *GetAPIV2ChartsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Api v2 charts bad request response has a 5xx status code
func (o *GetAPIV2ChartsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api v2 charts bad request response a status code equal to that given
func (o *GetAPIV2ChartsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetAPIV2ChartsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/charts][%d] getApiV2ChartsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAPIV2ChartsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/charts][%d] getApiV2ChartsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAPIV2ChartsBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ChartsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ChartsForbidden creates a GetAPIV2ChartsForbidden with default headers values
func NewGetAPIV2ChartsForbidden() *GetAPIV2ChartsForbidden {
	return &GetAPIV2ChartsForbidden{}
}

/*
GetAPIV2ChartsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPIV2ChartsForbidden struct {
	Payload *models.ErrorsErrorResponse
}

// IsSuccess returns true when this get Api v2 charts forbidden response has a 2xx status code
func (o *GetAPIV2ChartsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Api v2 charts forbidden response has a 3xx status code
func (o *GetAPIV2ChartsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api v2 charts forbidden response has a 4xx status code
func (o *GetAPIV2ChartsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Api v2 charts forbidden response has a 5xx status code
func (o *GetAPIV2ChartsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api v2 charts forbidden response a status code equal to that given
func (o *GetAPIV2ChartsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAPIV2ChartsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/charts][%d] getApiV2ChartsForbidden  %+v", 403, o.Payload)
}

func (o *GetAPIV2ChartsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/charts][%d] getApiV2ChartsForbidden  %+v", 403, o.Payload)
}

func (o *GetAPIV2ChartsForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ChartsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ChartsNotFound creates a GetAPIV2ChartsNotFound with default headers values
func NewGetAPIV2ChartsNotFound() *GetAPIV2ChartsNotFound {
	return &GetAPIV2ChartsNotFound{}
}

/*
GetAPIV2ChartsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIV2ChartsNotFound struct {
	Payload *models.ErrorsErrorResponse
}

// IsSuccess returns true when this get Api v2 charts not found response has a 2xx status code
func (o *GetAPIV2ChartsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Api v2 charts not found response has a 3xx status code
func (o *GetAPIV2ChartsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api v2 charts not found response has a 4xx status code
func (o *GetAPIV2ChartsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Api v2 charts not found response has a 5xx status code
func (o *GetAPIV2ChartsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api v2 charts not found response a status code equal to that given
func (o *GetAPIV2ChartsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetAPIV2ChartsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/charts][%d] getApiV2ChartsNotFound  %+v", 404, o.Payload)
}

func (o *GetAPIV2ChartsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/charts][%d] getApiV2ChartsNotFound  %+v", 404, o.Payload)
}

func (o *GetAPIV2ChartsNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ChartsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ChartsProxyAuthenticationRequired creates a GetAPIV2ChartsProxyAuthenticationRequired with default headers values
func NewGetAPIV2ChartsProxyAuthenticationRequired() *GetAPIV2ChartsProxyAuthenticationRequired {
	return &GetAPIV2ChartsProxyAuthenticationRequired{}
}

/*
GetAPIV2ChartsProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPIV2ChartsProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

// IsSuccess returns true when this get Api v2 charts proxy authentication required response has a 2xx status code
func (o *GetAPIV2ChartsProxyAuthenticationRequired) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Api v2 charts proxy authentication required response has a 3xx status code
func (o *GetAPIV2ChartsProxyAuthenticationRequired) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api v2 charts proxy authentication required response has a 4xx status code
func (o *GetAPIV2ChartsProxyAuthenticationRequired) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Api v2 charts proxy authentication required response has a 5xx status code
func (o *GetAPIV2ChartsProxyAuthenticationRequired) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api v2 charts proxy authentication required response a status code equal to that given
func (o *GetAPIV2ChartsProxyAuthenticationRequired) IsCode(code int) bool {
	return code == 407
}

func (o *GetAPIV2ChartsProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/v2/charts][%d] getApiV2ChartsProxyAuthenticationRequired  %+v", 407, o.Payload)
}

func (o *GetAPIV2ChartsProxyAuthenticationRequired) String() string {
	return fmt.Sprintf("[GET /api/v2/charts][%d] getApiV2ChartsProxyAuthenticationRequired  %+v", 407, o.Payload)
}

func (o *GetAPIV2ChartsProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ChartsProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ChartsConflict creates a GetAPIV2ChartsConflict with default headers values
func NewGetAPIV2ChartsConflict() *GetAPIV2ChartsConflict {
	return &GetAPIV2ChartsConflict{}
}

/*
GetAPIV2ChartsConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPIV2ChartsConflict struct {
	Payload *models.ErrorsErrorResponse
}

// IsSuccess returns true when this get Api v2 charts conflict response has a 2xx status code
func (o *GetAPIV2ChartsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Api v2 charts conflict response has a 3xx status code
func (o *GetAPIV2ChartsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api v2 charts conflict response has a 4xx status code
func (o *GetAPIV2ChartsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Api v2 charts conflict response has a 5xx status code
func (o *GetAPIV2ChartsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api v2 charts conflict response a status code equal to that given
func (o *GetAPIV2ChartsConflict) IsCode(code int) bool {
	return code == 409
}

func (o *GetAPIV2ChartsConflict) Error() string {
	return fmt.Sprintf("[GET /api/v2/charts][%d] getApiV2ChartsConflict  %+v", 409, o.Payload)
}

func (o *GetAPIV2ChartsConflict) String() string {
	return fmt.Sprintf("[GET /api/v2/charts][%d] getApiV2ChartsConflict  %+v", 409, o.Payload)
}

func (o *GetAPIV2ChartsConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ChartsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ChartsInternalServerError creates a GetAPIV2ChartsInternalServerError with default headers values
func NewGetAPIV2ChartsInternalServerError() *GetAPIV2ChartsInternalServerError {
	return &GetAPIV2ChartsInternalServerError{}
}

/*
GetAPIV2ChartsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPIV2ChartsInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

// IsSuccess returns true when this get Api v2 charts internal server error response has a 2xx status code
func (o *GetAPIV2ChartsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Api v2 charts internal server error response has a 3xx status code
func (o *GetAPIV2ChartsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api v2 charts internal server error response has a 4xx status code
func (o *GetAPIV2ChartsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api v2 charts internal server error response has a 5xx status code
func (o *GetAPIV2ChartsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get Api v2 charts internal server error response a status code equal to that given
func (o *GetAPIV2ChartsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetAPIV2ChartsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/charts][%d] getApiV2ChartsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAPIV2ChartsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/charts][%d] getApiV2ChartsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAPIV2ChartsInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ChartsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
