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

// GetAPICiRunsV3SelectorReader is a Reader for the GetAPICiRunsV3Selector structure.
type GetAPICiRunsV3SelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPICiRunsV3SelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPICiRunsV3SelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPICiRunsV3SelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPICiRunsV3SelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPICiRunsV3SelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPICiRunsV3SelectorProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPICiRunsV3SelectorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPICiRunsV3SelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPICiRunsV3SelectorOK creates a GetAPICiRunsV3SelectorOK with default headers values
func NewGetAPICiRunsV3SelectorOK() *GetAPICiRunsV3SelectorOK {
	return &GetAPICiRunsV3SelectorOK{}
}

/* GetAPICiRunsV3SelectorOK describes a response with status code 200, with default header values.

OK
*/
type GetAPICiRunsV3SelectorOK struct {
	Payload *models.SherlockCiRunV3
}

func (o *GetAPICiRunsV3SelectorOK) Error() string {
	return fmt.Sprintf("[GET /api/ci-runs/v3/{selector}][%d] getApiCiRunsV3SelectorOK  %+v", 200, o.Payload)
}
func (o *GetAPICiRunsV3SelectorOK) GetPayload() *models.SherlockCiRunV3 {
	return o.Payload
}

func (o *GetAPICiRunsV3SelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SherlockCiRunV3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPICiRunsV3SelectorBadRequest creates a GetAPICiRunsV3SelectorBadRequest with default headers values
func NewGetAPICiRunsV3SelectorBadRequest() *GetAPICiRunsV3SelectorBadRequest {
	return &GetAPICiRunsV3SelectorBadRequest{}
}

/* GetAPICiRunsV3SelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPICiRunsV3SelectorBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPICiRunsV3SelectorBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/ci-runs/v3/{selector}][%d] getApiCiRunsV3SelectorBadRequest  %+v", 400, o.Payload)
}
func (o *GetAPICiRunsV3SelectorBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPICiRunsV3SelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPICiRunsV3SelectorForbidden creates a GetAPICiRunsV3SelectorForbidden with default headers values
func NewGetAPICiRunsV3SelectorForbidden() *GetAPICiRunsV3SelectorForbidden {
	return &GetAPICiRunsV3SelectorForbidden{}
}

/* GetAPICiRunsV3SelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPICiRunsV3SelectorForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPICiRunsV3SelectorForbidden) Error() string {
	return fmt.Sprintf("[GET /api/ci-runs/v3/{selector}][%d] getApiCiRunsV3SelectorForbidden  %+v", 403, o.Payload)
}
func (o *GetAPICiRunsV3SelectorForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPICiRunsV3SelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPICiRunsV3SelectorNotFound creates a GetAPICiRunsV3SelectorNotFound with default headers values
func NewGetAPICiRunsV3SelectorNotFound() *GetAPICiRunsV3SelectorNotFound {
	return &GetAPICiRunsV3SelectorNotFound{}
}

/* GetAPICiRunsV3SelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPICiRunsV3SelectorNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPICiRunsV3SelectorNotFound) Error() string {
	return fmt.Sprintf("[GET /api/ci-runs/v3/{selector}][%d] getApiCiRunsV3SelectorNotFound  %+v", 404, o.Payload)
}
func (o *GetAPICiRunsV3SelectorNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPICiRunsV3SelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPICiRunsV3SelectorProxyAuthenticationRequired creates a GetAPICiRunsV3SelectorProxyAuthenticationRequired with default headers values
func NewGetAPICiRunsV3SelectorProxyAuthenticationRequired() *GetAPICiRunsV3SelectorProxyAuthenticationRequired {
	return &GetAPICiRunsV3SelectorProxyAuthenticationRequired{}
}

/* GetAPICiRunsV3SelectorProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPICiRunsV3SelectorProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPICiRunsV3SelectorProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/ci-runs/v3/{selector}][%d] getApiCiRunsV3SelectorProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPICiRunsV3SelectorProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPICiRunsV3SelectorProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPICiRunsV3SelectorConflict creates a GetAPICiRunsV3SelectorConflict with default headers values
func NewGetAPICiRunsV3SelectorConflict() *GetAPICiRunsV3SelectorConflict {
	return &GetAPICiRunsV3SelectorConflict{}
}

/* GetAPICiRunsV3SelectorConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPICiRunsV3SelectorConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPICiRunsV3SelectorConflict) Error() string {
	return fmt.Sprintf("[GET /api/ci-runs/v3/{selector}][%d] getApiCiRunsV3SelectorConflict  %+v", 409, o.Payload)
}
func (o *GetAPICiRunsV3SelectorConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPICiRunsV3SelectorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPICiRunsV3SelectorInternalServerError creates a GetAPICiRunsV3SelectorInternalServerError with default headers values
func NewGetAPICiRunsV3SelectorInternalServerError() *GetAPICiRunsV3SelectorInternalServerError {
	return &GetAPICiRunsV3SelectorInternalServerError{}
}

/* GetAPICiRunsV3SelectorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPICiRunsV3SelectorInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPICiRunsV3SelectorInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/ci-runs/v3/{selector}][%d] getApiCiRunsV3SelectorInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPICiRunsV3SelectorInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPICiRunsV3SelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
