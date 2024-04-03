// Code generated by go-swagger; DO NOT EDIT.

package github_actions_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// GetAPIGithubActionsJobsV3SelectorReader is a Reader for the GetAPIGithubActionsJobsV3Selector structure.
type GetAPIGithubActionsJobsV3SelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIGithubActionsJobsV3SelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIGithubActionsJobsV3SelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIGithubActionsJobsV3SelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIGithubActionsJobsV3SelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIGithubActionsJobsV3SelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPIGithubActionsJobsV3SelectorProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPIGithubActionsJobsV3SelectorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIGithubActionsJobsV3SelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIGithubActionsJobsV3SelectorOK creates a GetAPIGithubActionsJobsV3SelectorOK with default headers values
func NewGetAPIGithubActionsJobsV3SelectorOK() *GetAPIGithubActionsJobsV3SelectorOK {
	return &GetAPIGithubActionsJobsV3SelectorOK{}
}

/* GetAPIGithubActionsJobsV3SelectorOK describes a response with status code 200, with default header values.

OK
*/
type GetAPIGithubActionsJobsV3SelectorOK struct {
	Payload *models.SherlockGithubActionsJobV3
}

func (o *GetAPIGithubActionsJobsV3SelectorOK) Error() string {
	return fmt.Sprintf("[GET /api/github-actions-jobs/v3/{selector}][%d] getApiGithubActionsJobsV3SelectorOK  %+v", 200, o.Payload)
}
func (o *GetAPIGithubActionsJobsV3SelectorOK) GetPayload() *models.SherlockGithubActionsJobV3 {
	return o.Payload
}

func (o *GetAPIGithubActionsJobsV3SelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SherlockGithubActionsJobV3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIGithubActionsJobsV3SelectorBadRequest creates a GetAPIGithubActionsJobsV3SelectorBadRequest with default headers values
func NewGetAPIGithubActionsJobsV3SelectorBadRequest() *GetAPIGithubActionsJobsV3SelectorBadRequest {
	return &GetAPIGithubActionsJobsV3SelectorBadRequest{}
}

/* GetAPIGithubActionsJobsV3SelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIGithubActionsJobsV3SelectorBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIGithubActionsJobsV3SelectorBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/github-actions-jobs/v3/{selector}][%d] getApiGithubActionsJobsV3SelectorBadRequest  %+v", 400, o.Payload)
}
func (o *GetAPIGithubActionsJobsV3SelectorBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIGithubActionsJobsV3SelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIGithubActionsJobsV3SelectorForbidden creates a GetAPIGithubActionsJobsV3SelectorForbidden with default headers values
func NewGetAPIGithubActionsJobsV3SelectorForbidden() *GetAPIGithubActionsJobsV3SelectorForbidden {
	return &GetAPIGithubActionsJobsV3SelectorForbidden{}
}

/* GetAPIGithubActionsJobsV3SelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPIGithubActionsJobsV3SelectorForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIGithubActionsJobsV3SelectorForbidden) Error() string {
	return fmt.Sprintf("[GET /api/github-actions-jobs/v3/{selector}][%d] getApiGithubActionsJobsV3SelectorForbidden  %+v", 403, o.Payload)
}
func (o *GetAPIGithubActionsJobsV3SelectorForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIGithubActionsJobsV3SelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIGithubActionsJobsV3SelectorNotFound creates a GetAPIGithubActionsJobsV3SelectorNotFound with default headers values
func NewGetAPIGithubActionsJobsV3SelectorNotFound() *GetAPIGithubActionsJobsV3SelectorNotFound {
	return &GetAPIGithubActionsJobsV3SelectorNotFound{}
}

/* GetAPIGithubActionsJobsV3SelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIGithubActionsJobsV3SelectorNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIGithubActionsJobsV3SelectorNotFound) Error() string {
	return fmt.Sprintf("[GET /api/github-actions-jobs/v3/{selector}][%d] getApiGithubActionsJobsV3SelectorNotFound  %+v", 404, o.Payload)
}
func (o *GetAPIGithubActionsJobsV3SelectorNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIGithubActionsJobsV3SelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIGithubActionsJobsV3SelectorProxyAuthenticationRequired creates a GetAPIGithubActionsJobsV3SelectorProxyAuthenticationRequired with default headers values
func NewGetAPIGithubActionsJobsV3SelectorProxyAuthenticationRequired() *GetAPIGithubActionsJobsV3SelectorProxyAuthenticationRequired {
	return &GetAPIGithubActionsJobsV3SelectorProxyAuthenticationRequired{}
}

/* GetAPIGithubActionsJobsV3SelectorProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPIGithubActionsJobsV3SelectorProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIGithubActionsJobsV3SelectorProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/github-actions-jobs/v3/{selector}][%d] getApiGithubActionsJobsV3SelectorProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPIGithubActionsJobsV3SelectorProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIGithubActionsJobsV3SelectorProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIGithubActionsJobsV3SelectorConflict creates a GetAPIGithubActionsJobsV3SelectorConflict with default headers values
func NewGetAPIGithubActionsJobsV3SelectorConflict() *GetAPIGithubActionsJobsV3SelectorConflict {
	return &GetAPIGithubActionsJobsV3SelectorConflict{}
}

/* GetAPIGithubActionsJobsV3SelectorConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPIGithubActionsJobsV3SelectorConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIGithubActionsJobsV3SelectorConflict) Error() string {
	return fmt.Sprintf("[GET /api/github-actions-jobs/v3/{selector}][%d] getApiGithubActionsJobsV3SelectorConflict  %+v", 409, o.Payload)
}
func (o *GetAPIGithubActionsJobsV3SelectorConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIGithubActionsJobsV3SelectorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIGithubActionsJobsV3SelectorInternalServerError creates a GetAPIGithubActionsJobsV3SelectorInternalServerError with default headers values
func NewGetAPIGithubActionsJobsV3SelectorInternalServerError() *GetAPIGithubActionsJobsV3SelectorInternalServerError {
	return &GetAPIGithubActionsJobsV3SelectorInternalServerError{}
}

/* GetAPIGithubActionsJobsV3SelectorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPIGithubActionsJobsV3SelectorInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIGithubActionsJobsV3SelectorInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/github-actions-jobs/v3/{selector}][%d] getApiGithubActionsJobsV3SelectorInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPIGithubActionsJobsV3SelectorInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIGithubActionsJobsV3SelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}