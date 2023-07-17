// Code generated by go-swagger; DO NOT EDIT.

package ci_identifiers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// GetAPICiIdentifiersV3SelectorReader is a Reader for the GetAPICiIdentifiersV3Selector structure.
type GetAPICiIdentifiersV3SelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPICiIdentifiersV3SelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPICiIdentifiersV3SelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPICiIdentifiersV3SelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPICiIdentifiersV3SelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPICiIdentifiersV3SelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPICiIdentifiersV3SelectorProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPICiIdentifiersV3SelectorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPICiIdentifiersV3SelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPICiIdentifiersV3SelectorOK creates a GetAPICiIdentifiersV3SelectorOK with default headers values
func NewGetAPICiIdentifiersV3SelectorOK() *GetAPICiIdentifiersV3SelectorOK {
	return &GetAPICiIdentifiersV3SelectorOK{}
}

/* GetAPICiIdentifiersV3SelectorOK describes a response with status code 200, with default header values.

OK
*/
type GetAPICiIdentifiersV3SelectorOK struct {
	Payload *models.SherlockCiIdentifierV3
}

func (o *GetAPICiIdentifiersV3SelectorOK) Error() string {
	return fmt.Sprintf("[GET /api/ci-identifiers/v3/{selector}][%d] getApiCiIdentifiersV3SelectorOK  %+v", 200, o.Payload)
}
func (o *GetAPICiIdentifiersV3SelectorOK) GetPayload() *models.SherlockCiIdentifierV3 {
	return o.Payload
}

func (o *GetAPICiIdentifiersV3SelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SherlockCiIdentifierV3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPICiIdentifiersV3SelectorBadRequest creates a GetAPICiIdentifiersV3SelectorBadRequest with default headers values
func NewGetAPICiIdentifiersV3SelectorBadRequest() *GetAPICiIdentifiersV3SelectorBadRequest {
	return &GetAPICiIdentifiersV3SelectorBadRequest{}
}

/* GetAPICiIdentifiersV3SelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPICiIdentifiersV3SelectorBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPICiIdentifiersV3SelectorBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/ci-identifiers/v3/{selector}][%d] getApiCiIdentifiersV3SelectorBadRequest  %+v", 400, o.Payload)
}
func (o *GetAPICiIdentifiersV3SelectorBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPICiIdentifiersV3SelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPICiIdentifiersV3SelectorForbidden creates a GetAPICiIdentifiersV3SelectorForbidden with default headers values
func NewGetAPICiIdentifiersV3SelectorForbidden() *GetAPICiIdentifiersV3SelectorForbidden {
	return &GetAPICiIdentifiersV3SelectorForbidden{}
}

/* GetAPICiIdentifiersV3SelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPICiIdentifiersV3SelectorForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPICiIdentifiersV3SelectorForbidden) Error() string {
	return fmt.Sprintf("[GET /api/ci-identifiers/v3/{selector}][%d] getApiCiIdentifiersV3SelectorForbidden  %+v", 403, o.Payload)
}
func (o *GetAPICiIdentifiersV3SelectorForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPICiIdentifiersV3SelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPICiIdentifiersV3SelectorNotFound creates a GetAPICiIdentifiersV3SelectorNotFound with default headers values
func NewGetAPICiIdentifiersV3SelectorNotFound() *GetAPICiIdentifiersV3SelectorNotFound {
	return &GetAPICiIdentifiersV3SelectorNotFound{}
}

/* GetAPICiIdentifiersV3SelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPICiIdentifiersV3SelectorNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPICiIdentifiersV3SelectorNotFound) Error() string {
	return fmt.Sprintf("[GET /api/ci-identifiers/v3/{selector}][%d] getApiCiIdentifiersV3SelectorNotFound  %+v", 404, o.Payload)
}
func (o *GetAPICiIdentifiersV3SelectorNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPICiIdentifiersV3SelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPICiIdentifiersV3SelectorProxyAuthenticationRequired creates a GetAPICiIdentifiersV3SelectorProxyAuthenticationRequired with default headers values
func NewGetAPICiIdentifiersV3SelectorProxyAuthenticationRequired() *GetAPICiIdentifiersV3SelectorProxyAuthenticationRequired {
	return &GetAPICiIdentifiersV3SelectorProxyAuthenticationRequired{}
}

/* GetAPICiIdentifiersV3SelectorProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPICiIdentifiersV3SelectorProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPICiIdentifiersV3SelectorProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/ci-identifiers/v3/{selector}][%d] getApiCiIdentifiersV3SelectorProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPICiIdentifiersV3SelectorProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPICiIdentifiersV3SelectorProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPICiIdentifiersV3SelectorConflict creates a GetAPICiIdentifiersV3SelectorConflict with default headers values
func NewGetAPICiIdentifiersV3SelectorConflict() *GetAPICiIdentifiersV3SelectorConflict {
	return &GetAPICiIdentifiersV3SelectorConflict{}
}

/* GetAPICiIdentifiersV3SelectorConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPICiIdentifiersV3SelectorConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPICiIdentifiersV3SelectorConflict) Error() string {
	return fmt.Sprintf("[GET /api/ci-identifiers/v3/{selector}][%d] getApiCiIdentifiersV3SelectorConflict  %+v", 409, o.Payload)
}
func (o *GetAPICiIdentifiersV3SelectorConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPICiIdentifiersV3SelectorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPICiIdentifiersV3SelectorInternalServerError creates a GetAPICiIdentifiersV3SelectorInternalServerError with default headers values
func NewGetAPICiIdentifiersV3SelectorInternalServerError() *GetAPICiIdentifiersV3SelectorInternalServerError {
	return &GetAPICiIdentifiersV3SelectorInternalServerError{}
}

/* GetAPICiIdentifiersV3SelectorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPICiIdentifiersV3SelectorInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPICiIdentifiersV3SelectorInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/ci-identifiers/v3/{selector}][%d] getApiCiIdentifiersV3SelectorInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPICiIdentifiersV3SelectorInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPICiIdentifiersV3SelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}