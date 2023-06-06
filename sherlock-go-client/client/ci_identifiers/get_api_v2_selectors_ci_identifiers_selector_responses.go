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

// GetAPIV2SelectorsCiIdentifiersSelectorReader is a Reader for the GetAPIV2SelectorsCiIdentifiersSelector structure.
type GetAPIV2SelectorsCiIdentifiersSelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV2SelectorsCiIdentifiersSelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV2SelectorsCiIdentifiersSelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIV2SelectorsCiIdentifiersSelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIV2SelectorsCiIdentifiersSelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIV2SelectorsCiIdentifiersSelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPIV2SelectorsCiIdentifiersSelectorProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPIV2SelectorsCiIdentifiersSelectorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIV2SelectorsCiIdentifiersSelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIV2SelectorsCiIdentifiersSelectorOK creates a GetAPIV2SelectorsCiIdentifiersSelectorOK with default headers values
func NewGetAPIV2SelectorsCiIdentifiersSelectorOK() *GetAPIV2SelectorsCiIdentifiersSelectorOK {
	return &GetAPIV2SelectorsCiIdentifiersSelectorOK{}
}

/* GetAPIV2SelectorsCiIdentifiersSelectorOK describes a response with status code 200, with default header values.

OK
*/
type GetAPIV2SelectorsCiIdentifiersSelectorOK struct {
	Payload []string
}

func (o *GetAPIV2SelectorsCiIdentifiersSelectorOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/selectors/ci-identifiers/{selector}][%d] getApiV2SelectorsCiIdentifiersSelectorOK  %+v", 200, o.Payload)
}
func (o *GetAPIV2SelectorsCiIdentifiersSelectorOK) GetPayload() []string {
	return o.Payload
}

func (o *GetAPIV2SelectorsCiIdentifiersSelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2SelectorsCiIdentifiersSelectorBadRequest creates a GetAPIV2SelectorsCiIdentifiersSelectorBadRequest with default headers values
func NewGetAPIV2SelectorsCiIdentifiersSelectorBadRequest() *GetAPIV2SelectorsCiIdentifiersSelectorBadRequest {
	return &GetAPIV2SelectorsCiIdentifiersSelectorBadRequest{}
}

/* GetAPIV2SelectorsCiIdentifiersSelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIV2SelectorsCiIdentifiersSelectorBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2SelectorsCiIdentifiersSelectorBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/selectors/ci-identifiers/{selector}][%d] getApiV2SelectorsCiIdentifiersSelectorBadRequest  %+v", 400, o.Payload)
}
func (o *GetAPIV2SelectorsCiIdentifiersSelectorBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2SelectorsCiIdentifiersSelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2SelectorsCiIdentifiersSelectorForbidden creates a GetAPIV2SelectorsCiIdentifiersSelectorForbidden with default headers values
func NewGetAPIV2SelectorsCiIdentifiersSelectorForbidden() *GetAPIV2SelectorsCiIdentifiersSelectorForbidden {
	return &GetAPIV2SelectorsCiIdentifiersSelectorForbidden{}
}

/* GetAPIV2SelectorsCiIdentifiersSelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPIV2SelectorsCiIdentifiersSelectorForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2SelectorsCiIdentifiersSelectorForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/selectors/ci-identifiers/{selector}][%d] getApiV2SelectorsCiIdentifiersSelectorForbidden  %+v", 403, o.Payload)
}
func (o *GetAPIV2SelectorsCiIdentifiersSelectorForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2SelectorsCiIdentifiersSelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2SelectorsCiIdentifiersSelectorNotFound creates a GetAPIV2SelectorsCiIdentifiersSelectorNotFound with default headers values
func NewGetAPIV2SelectorsCiIdentifiersSelectorNotFound() *GetAPIV2SelectorsCiIdentifiersSelectorNotFound {
	return &GetAPIV2SelectorsCiIdentifiersSelectorNotFound{}
}

/* GetAPIV2SelectorsCiIdentifiersSelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIV2SelectorsCiIdentifiersSelectorNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2SelectorsCiIdentifiersSelectorNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/selectors/ci-identifiers/{selector}][%d] getApiV2SelectorsCiIdentifiersSelectorNotFound  %+v", 404, o.Payload)
}
func (o *GetAPIV2SelectorsCiIdentifiersSelectorNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2SelectorsCiIdentifiersSelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2SelectorsCiIdentifiersSelectorProxyAuthenticationRequired creates a GetAPIV2SelectorsCiIdentifiersSelectorProxyAuthenticationRequired with default headers values
func NewGetAPIV2SelectorsCiIdentifiersSelectorProxyAuthenticationRequired() *GetAPIV2SelectorsCiIdentifiersSelectorProxyAuthenticationRequired {
	return &GetAPIV2SelectorsCiIdentifiersSelectorProxyAuthenticationRequired{}
}

/* GetAPIV2SelectorsCiIdentifiersSelectorProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPIV2SelectorsCiIdentifiersSelectorProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2SelectorsCiIdentifiersSelectorProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/v2/selectors/ci-identifiers/{selector}][%d] getApiV2SelectorsCiIdentifiersSelectorProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPIV2SelectorsCiIdentifiersSelectorProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2SelectorsCiIdentifiersSelectorProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2SelectorsCiIdentifiersSelectorConflict creates a GetAPIV2SelectorsCiIdentifiersSelectorConflict with default headers values
func NewGetAPIV2SelectorsCiIdentifiersSelectorConflict() *GetAPIV2SelectorsCiIdentifiersSelectorConflict {
	return &GetAPIV2SelectorsCiIdentifiersSelectorConflict{}
}

/* GetAPIV2SelectorsCiIdentifiersSelectorConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPIV2SelectorsCiIdentifiersSelectorConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2SelectorsCiIdentifiersSelectorConflict) Error() string {
	return fmt.Sprintf("[GET /api/v2/selectors/ci-identifiers/{selector}][%d] getApiV2SelectorsCiIdentifiersSelectorConflict  %+v", 409, o.Payload)
}
func (o *GetAPIV2SelectorsCiIdentifiersSelectorConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2SelectorsCiIdentifiersSelectorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2SelectorsCiIdentifiersSelectorInternalServerError creates a GetAPIV2SelectorsCiIdentifiersSelectorInternalServerError with default headers values
func NewGetAPIV2SelectorsCiIdentifiersSelectorInternalServerError() *GetAPIV2SelectorsCiIdentifiersSelectorInternalServerError {
	return &GetAPIV2SelectorsCiIdentifiersSelectorInternalServerError{}
}

/* GetAPIV2SelectorsCiIdentifiersSelectorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPIV2SelectorsCiIdentifiersSelectorInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2SelectorsCiIdentifiersSelectorInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/selectors/ci-identifiers/{selector}][%d] getApiV2SelectorsCiIdentifiersSelectorInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPIV2SelectorsCiIdentifiersSelectorInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2SelectorsCiIdentifiersSelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
