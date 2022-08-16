// Code generated by go-swagger; DO NOT EDIT.

package app_versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/client/go/client/models"
)

// GetAPIV2AppVersionsReader is a Reader for the GetAPIV2AppVersions structure.
type GetAPIV2AppVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV2AppVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV2AppVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIV2AppVersionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIV2AppVersionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIV2AppVersionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPIV2AppVersionsProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPIV2AppVersionsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIV2AppVersionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIV2AppVersionsOK creates a GetAPIV2AppVersionsOK with default headers values
func NewGetAPIV2AppVersionsOK() *GetAPIV2AppVersionsOK {
	return &GetAPIV2AppVersionsOK{}
}

/* GetAPIV2AppVersionsOK describes a response with status code 200, with default header values.

OK
*/
type GetAPIV2AppVersionsOK struct {
	Payload []*models.V2controllersAppVersion
}

func (o *GetAPIV2AppVersionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/app-versions][%d] getApiV2AppVersionsOK  %+v", 200, o.Payload)
}
func (o *GetAPIV2AppVersionsOK) GetPayload() []*models.V2controllersAppVersion {
	return o.Payload
}

func (o *GetAPIV2AppVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2AppVersionsBadRequest creates a GetAPIV2AppVersionsBadRequest with default headers values
func NewGetAPIV2AppVersionsBadRequest() *GetAPIV2AppVersionsBadRequest {
	return &GetAPIV2AppVersionsBadRequest{}
}

/* GetAPIV2AppVersionsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIV2AppVersionsBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2AppVersionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/app-versions][%d] getApiV2AppVersionsBadRequest  %+v", 400, o.Payload)
}
func (o *GetAPIV2AppVersionsBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2AppVersionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2AppVersionsForbidden creates a GetAPIV2AppVersionsForbidden with default headers values
func NewGetAPIV2AppVersionsForbidden() *GetAPIV2AppVersionsForbidden {
	return &GetAPIV2AppVersionsForbidden{}
}

/* GetAPIV2AppVersionsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPIV2AppVersionsForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2AppVersionsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/app-versions][%d] getApiV2AppVersionsForbidden  %+v", 403, o.Payload)
}
func (o *GetAPIV2AppVersionsForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2AppVersionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2AppVersionsNotFound creates a GetAPIV2AppVersionsNotFound with default headers values
func NewGetAPIV2AppVersionsNotFound() *GetAPIV2AppVersionsNotFound {
	return &GetAPIV2AppVersionsNotFound{}
}

/* GetAPIV2AppVersionsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIV2AppVersionsNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2AppVersionsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/app-versions][%d] getApiV2AppVersionsNotFound  %+v", 404, o.Payload)
}
func (o *GetAPIV2AppVersionsNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2AppVersionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2AppVersionsProxyAuthenticationRequired creates a GetAPIV2AppVersionsProxyAuthenticationRequired with default headers values
func NewGetAPIV2AppVersionsProxyAuthenticationRequired() *GetAPIV2AppVersionsProxyAuthenticationRequired {
	return &GetAPIV2AppVersionsProxyAuthenticationRequired{}
}

/* GetAPIV2AppVersionsProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPIV2AppVersionsProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2AppVersionsProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/v2/app-versions][%d] getApiV2AppVersionsProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPIV2AppVersionsProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2AppVersionsProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2AppVersionsConflict creates a GetAPIV2AppVersionsConflict with default headers values
func NewGetAPIV2AppVersionsConflict() *GetAPIV2AppVersionsConflict {
	return &GetAPIV2AppVersionsConflict{}
}

/* GetAPIV2AppVersionsConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPIV2AppVersionsConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2AppVersionsConflict) Error() string {
	return fmt.Sprintf("[GET /api/v2/app-versions][%d] getApiV2AppVersionsConflict  %+v", 409, o.Payload)
}
func (o *GetAPIV2AppVersionsConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2AppVersionsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2AppVersionsInternalServerError creates a GetAPIV2AppVersionsInternalServerError with default headers values
func NewGetAPIV2AppVersionsInternalServerError() *GetAPIV2AppVersionsInternalServerError {
	return &GetAPIV2AppVersionsInternalServerError{}
}

/* GetAPIV2AppVersionsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPIV2AppVersionsInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2AppVersionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/app-versions][%d] getApiV2AppVersionsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPIV2AppVersionsInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2AppVersionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
