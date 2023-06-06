// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// GetAPIV2SelectorsUsersSelectorReader is a Reader for the GetAPIV2SelectorsUsersSelector structure.
type GetAPIV2SelectorsUsersSelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV2SelectorsUsersSelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV2SelectorsUsersSelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIV2SelectorsUsersSelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIV2SelectorsUsersSelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIV2SelectorsUsersSelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPIV2SelectorsUsersSelectorProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPIV2SelectorsUsersSelectorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIV2SelectorsUsersSelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIV2SelectorsUsersSelectorOK creates a GetAPIV2SelectorsUsersSelectorOK with default headers values
func NewGetAPIV2SelectorsUsersSelectorOK() *GetAPIV2SelectorsUsersSelectorOK {
	return &GetAPIV2SelectorsUsersSelectorOK{}
}

/* GetAPIV2SelectorsUsersSelectorOK describes a response with status code 200, with default header values.

OK
*/
type GetAPIV2SelectorsUsersSelectorOK struct {
	Payload []string
}

func (o *GetAPIV2SelectorsUsersSelectorOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/selectors/users/{selector}][%d] getApiV2SelectorsUsersSelectorOK  %+v", 200, o.Payload)
}
func (o *GetAPIV2SelectorsUsersSelectorOK) GetPayload() []string {
	return o.Payload
}

func (o *GetAPIV2SelectorsUsersSelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2SelectorsUsersSelectorBadRequest creates a GetAPIV2SelectorsUsersSelectorBadRequest with default headers values
func NewGetAPIV2SelectorsUsersSelectorBadRequest() *GetAPIV2SelectorsUsersSelectorBadRequest {
	return &GetAPIV2SelectorsUsersSelectorBadRequest{}
}

/* GetAPIV2SelectorsUsersSelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIV2SelectorsUsersSelectorBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2SelectorsUsersSelectorBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/selectors/users/{selector}][%d] getApiV2SelectorsUsersSelectorBadRequest  %+v", 400, o.Payload)
}
func (o *GetAPIV2SelectorsUsersSelectorBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2SelectorsUsersSelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2SelectorsUsersSelectorForbidden creates a GetAPIV2SelectorsUsersSelectorForbidden with default headers values
func NewGetAPIV2SelectorsUsersSelectorForbidden() *GetAPIV2SelectorsUsersSelectorForbidden {
	return &GetAPIV2SelectorsUsersSelectorForbidden{}
}

/* GetAPIV2SelectorsUsersSelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPIV2SelectorsUsersSelectorForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2SelectorsUsersSelectorForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/selectors/users/{selector}][%d] getApiV2SelectorsUsersSelectorForbidden  %+v", 403, o.Payload)
}
func (o *GetAPIV2SelectorsUsersSelectorForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2SelectorsUsersSelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2SelectorsUsersSelectorNotFound creates a GetAPIV2SelectorsUsersSelectorNotFound with default headers values
func NewGetAPIV2SelectorsUsersSelectorNotFound() *GetAPIV2SelectorsUsersSelectorNotFound {
	return &GetAPIV2SelectorsUsersSelectorNotFound{}
}

/* GetAPIV2SelectorsUsersSelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIV2SelectorsUsersSelectorNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2SelectorsUsersSelectorNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/selectors/users/{selector}][%d] getApiV2SelectorsUsersSelectorNotFound  %+v", 404, o.Payload)
}
func (o *GetAPIV2SelectorsUsersSelectorNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2SelectorsUsersSelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2SelectorsUsersSelectorProxyAuthenticationRequired creates a GetAPIV2SelectorsUsersSelectorProxyAuthenticationRequired with default headers values
func NewGetAPIV2SelectorsUsersSelectorProxyAuthenticationRequired() *GetAPIV2SelectorsUsersSelectorProxyAuthenticationRequired {
	return &GetAPIV2SelectorsUsersSelectorProxyAuthenticationRequired{}
}

/* GetAPIV2SelectorsUsersSelectorProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPIV2SelectorsUsersSelectorProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2SelectorsUsersSelectorProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/v2/selectors/users/{selector}][%d] getApiV2SelectorsUsersSelectorProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPIV2SelectorsUsersSelectorProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2SelectorsUsersSelectorProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2SelectorsUsersSelectorConflict creates a GetAPIV2SelectorsUsersSelectorConflict with default headers values
func NewGetAPIV2SelectorsUsersSelectorConflict() *GetAPIV2SelectorsUsersSelectorConflict {
	return &GetAPIV2SelectorsUsersSelectorConflict{}
}

/* GetAPIV2SelectorsUsersSelectorConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPIV2SelectorsUsersSelectorConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2SelectorsUsersSelectorConflict) Error() string {
	return fmt.Sprintf("[GET /api/v2/selectors/users/{selector}][%d] getApiV2SelectorsUsersSelectorConflict  %+v", 409, o.Payload)
}
func (o *GetAPIV2SelectorsUsersSelectorConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2SelectorsUsersSelectorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2SelectorsUsersSelectorInternalServerError creates a GetAPIV2SelectorsUsersSelectorInternalServerError with default headers values
func NewGetAPIV2SelectorsUsersSelectorInternalServerError() *GetAPIV2SelectorsUsersSelectorInternalServerError {
	return &GetAPIV2SelectorsUsersSelectorInternalServerError{}
}

/* GetAPIV2SelectorsUsersSelectorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPIV2SelectorsUsersSelectorInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2SelectorsUsersSelectorInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/selectors/users/{selector}][%d] getApiV2SelectorsUsersSelectorInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPIV2SelectorsUsersSelectorInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2SelectorsUsersSelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
