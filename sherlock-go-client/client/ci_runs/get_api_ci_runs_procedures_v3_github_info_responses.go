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

// GetAPICiRunsProceduresV3GithubInfoReader is a Reader for the GetAPICiRunsProceduresV3GithubInfo structure.
type GetAPICiRunsProceduresV3GithubInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPICiRunsProceduresV3GithubInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPICiRunsProceduresV3GithubInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPICiRunsProceduresV3GithubInfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPICiRunsProceduresV3GithubInfoForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPICiRunsProceduresV3GithubInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPICiRunsProceduresV3GithubInfoProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPICiRunsProceduresV3GithubInfoConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPICiRunsProceduresV3GithubInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPICiRunsProceduresV3GithubInfoOK creates a GetAPICiRunsProceduresV3GithubInfoOK with default headers values
func NewGetAPICiRunsProceduresV3GithubInfoOK() *GetAPICiRunsProceduresV3GithubInfoOK {
	return &GetAPICiRunsProceduresV3GithubInfoOK{}
}

/* GetAPICiRunsProceduresV3GithubInfoOK describes a response with status code 200, with default header values.

OK
*/
type GetAPICiRunsProceduresV3GithubInfoOK struct {
	Payload map[string]map[string][]string
}

func (o *GetAPICiRunsProceduresV3GithubInfoOK) Error() string {
	return fmt.Sprintf("[GET /api/ci-runs/procedures/v3/github-info][%d] getApiCiRunsProceduresV3GithubInfoOK  %+v", 200, o.Payload)
}
func (o *GetAPICiRunsProceduresV3GithubInfoOK) GetPayload() map[string]map[string][]string {
	return o.Payload
}

func (o *GetAPICiRunsProceduresV3GithubInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPICiRunsProceduresV3GithubInfoBadRequest creates a GetAPICiRunsProceduresV3GithubInfoBadRequest with default headers values
func NewGetAPICiRunsProceduresV3GithubInfoBadRequest() *GetAPICiRunsProceduresV3GithubInfoBadRequest {
	return &GetAPICiRunsProceduresV3GithubInfoBadRequest{}
}

/* GetAPICiRunsProceduresV3GithubInfoBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPICiRunsProceduresV3GithubInfoBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPICiRunsProceduresV3GithubInfoBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/ci-runs/procedures/v3/github-info][%d] getApiCiRunsProceduresV3GithubInfoBadRequest  %+v", 400, o.Payload)
}
func (o *GetAPICiRunsProceduresV3GithubInfoBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPICiRunsProceduresV3GithubInfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPICiRunsProceduresV3GithubInfoForbidden creates a GetAPICiRunsProceduresV3GithubInfoForbidden with default headers values
func NewGetAPICiRunsProceduresV3GithubInfoForbidden() *GetAPICiRunsProceduresV3GithubInfoForbidden {
	return &GetAPICiRunsProceduresV3GithubInfoForbidden{}
}

/* GetAPICiRunsProceduresV3GithubInfoForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPICiRunsProceduresV3GithubInfoForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPICiRunsProceduresV3GithubInfoForbidden) Error() string {
	return fmt.Sprintf("[GET /api/ci-runs/procedures/v3/github-info][%d] getApiCiRunsProceduresV3GithubInfoForbidden  %+v", 403, o.Payload)
}
func (o *GetAPICiRunsProceduresV3GithubInfoForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPICiRunsProceduresV3GithubInfoForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPICiRunsProceduresV3GithubInfoNotFound creates a GetAPICiRunsProceduresV3GithubInfoNotFound with default headers values
func NewGetAPICiRunsProceduresV3GithubInfoNotFound() *GetAPICiRunsProceduresV3GithubInfoNotFound {
	return &GetAPICiRunsProceduresV3GithubInfoNotFound{}
}

/* GetAPICiRunsProceduresV3GithubInfoNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPICiRunsProceduresV3GithubInfoNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPICiRunsProceduresV3GithubInfoNotFound) Error() string {
	return fmt.Sprintf("[GET /api/ci-runs/procedures/v3/github-info][%d] getApiCiRunsProceduresV3GithubInfoNotFound  %+v", 404, o.Payload)
}
func (o *GetAPICiRunsProceduresV3GithubInfoNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPICiRunsProceduresV3GithubInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPICiRunsProceduresV3GithubInfoProxyAuthenticationRequired creates a GetAPICiRunsProceduresV3GithubInfoProxyAuthenticationRequired with default headers values
func NewGetAPICiRunsProceduresV3GithubInfoProxyAuthenticationRequired() *GetAPICiRunsProceduresV3GithubInfoProxyAuthenticationRequired {
	return &GetAPICiRunsProceduresV3GithubInfoProxyAuthenticationRequired{}
}

/* GetAPICiRunsProceduresV3GithubInfoProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPICiRunsProceduresV3GithubInfoProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPICiRunsProceduresV3GithubInfoProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/ci-runs/procedures/v3/github-info][%d] getApiCiRunsProceduresV3GithubInfoProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPICiRunsProceduresV3GithubInfoProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPICiRunsProceduresV3GithubInfoProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPICiRunsProceduresV3GithubInfoConflict creates a GetAPICiRunsProceduresV3GithubInfoConflict with default headers values
func NewGetAPICiRunsProceduresV3GithubInfoConflict() *GetAPICiRunsProceduresV3GithubInfoConflict {
	return &GetAPICiRunsProceduresV3GithubInfoConflict{}
}

/* GetAPICiRunsProceduresV3GithubInfoConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPICiRunsProceduresV3GithubInfoConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPICiRunsProceduresV3GithubInfoConflict) Error() string {
	return fmt.Sprintf("[GET /api/ci-runs/procedures/v3/github-info][%d] getApiCiRunsProceduresV3GithubInfoConflict  %+v", 409, o.Payload)
}
func (o *GetAPICiRunsProceduresV3GithubInfoConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPICiRunsProceduresV3GithubInfoConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPICiRunsProceduresV3GithubInfoInternalServerError creates a GetAPICiRunsProceduresV3GithubInfoInternalServerError with default headers values
func NewGetAPICiRunsProceduresV3GithubInfoInternalServerError() *GetAPICiRunsProceduresV3GithubInfoInternalServerError {
	return &GetAPICiRunsProceduresV3GithubInfoInternalServerError{}
}

/* GetAPICiRunsProceduresV3GithubInfoInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPICiRunsProceduresV3GithubInfoInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPICiRunsProceduresV3GithubInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/ci-runs/procedures/v3/github-info][%d] getApiCiRunsProceduresV3GithubInfoInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPICiRunsProceduresV3GithubInfoInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPICiRunsProceduresV3GithubInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
