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

// PutAPIGithubActionsJobsV3Reader is a Reader for the PutAPIGithubActionsJobsV3 structure.
type PutAPIGithubActionsJobsV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAPIGithubActionsJobsV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutAPIGithubActionsJobsV3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutAPIGithubActionsJobsV3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutAPIGithubActionsJobsV3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutAPIGithubActionsJobsV3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewPutAPIGithubActionsJobsV3ProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutAPIGithubActionsJobsV3Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutAPIGithubActionsJobsV3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutAPIGithubActionsJobsV3OK creates a PutAPIGithubActionsJobsV3OK with default headers values
func NewPutAPIGithubActionsJobsV3OK() *PutAPIGithubActionsJobsV3OK {
	return &PutAPIGithubActionsJobsV3OK{}
}

/* PutAPIGithubActionsJobsV3OK describes a response with status code 200, with default header values.

OK
*/
type PutAPIGithubActionsJobsV3OK struct {
	Payload *models.SherlockGithubActionsJobV3
}

func (o *PutAPIGithubActionsJobsV3OK) Error() string {
	return fmt.Sprintf("[PUT /api/github-actions-jobs/v3][%d] putApiGithubActionsJobsV3OK  %+v", 200, o.Payload)
}
func (o *PutAPIGithubActionsJobsV3OK) GetPayload() *models.SherlockGithubActionsJobV3 {
	return o.Payload
}

func (o *PutAPIGithubActionsJobsV3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SherlockGithubActionsJobV3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIGithubActionsJobsV3BadRequest creates a PutAPIGithubActionsJobsV3BadRequest with default headers values
func NewPutAPIGithubActionsJobsV3BadRequest() *PutAPIGithubActionsJobsV3BadRequest {
	return &PutAPIGithubActionsJobsV3BadRequest{}
}

/* PutAPIGithubActionsJobsV3BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutAPIGithubActionsJobsV3BadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PutAPIGithubActionsJobsV3BadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/github-actions-jobs/v3][%d] putApiGithubActionsJobsV3BadRequest  %+v", 400, o.Payload)
}
func (o *PutAPIGithubActionsJobsV3BadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PutAPIGithubActionsJobsV3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIGithubActionsJobsV3Forbidden creates a PutAPIGithubActionsJobsV3Forbidden with default headers values
func NewPutAPIGithubActionsJobsV3Forbidden() *PutAPIGithubActionsJobsV3Forbidden {
	return &PutAPIGithubActionsJobsV3Forbidden{}
}

/* PutAPIGithubActionsJobsV3Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutAPIGithubActionsJobsV3Forbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PutAPIGithubActionsJobsV3Forbidden) Error() string {
	return fmt.Sprintf("[PUT /api/github-actions-jobs/v3][%d] putApiGithubActionsJobsV3Forbidden  %+v", 403, o.Payload)
}
func (o *PutAPIGithubActionsJobsV3Forbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PutAPIGithubActionsJobsV3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIGithubActionsJobsV3NotFound creates a PutAPIGithubActionsJobsV3NotFound with default headers values
func NewPutAPIGithubActionsJobsV3NotFound() *PutAPIGithubActionsJobsV3NotFound {
	return &PutAPIGithubActionsJobsV3NotFound{}
}

/* PutAPIGithubActionsJobsV3NotFound describes a response with status code 404, with default header values.

Not Found
*/
type PutAPIGithubActionsJobsV3NotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PutAPIGithubActionsJobsV3NotFound) Error() string {
	return fmt.Sprintf("[PUT /api/github-actions-jobs/v3][%d] putApiGithubActionsJobsV3NotFound  %+v", 404, o.Payload)
}
func (o *PutAPIGithubActionsJobsV3NotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PutAPIGithubActionsJobsV3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIGithubActionsJobsV3ProxyAuthenticationRequired creates a PutAPIGithubActionsJobsV3ProxyAuthenticationRequired with default headers values
func NewPutAPIGithubActionsJobsV3ProxyAuthenticationRequired() *PutAPIGithubActionsJobsV3ProxyAuthenticationRequired {
	return &PutAPIGithubActionsJobsV3ProxyAuthenticationRequired{}
}

/* PutAPIGithubActionsJobsV3ProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type PutAPIGithubActionsJobsV3ProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PutAPIGithubActionsJobsV3ProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[PUT /api/github-actions-jobs/v3][%d] putApiGithubActionsJobsV3ProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *PutAPIGithubActionsJobsV3ProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PutAPIGithubActionsJobsV3ProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIGithubActionsJobsV3Conflict creates a PutAPIGithubActionsJobsV3Conflict with default headers values
func NewPutAPIGithubActionsJobsV3Conflict() *PutAPIGithubActionsJobsV3Conflict {
	return &PutAPIGithubActionsJobsV3Conflict{}
}

/* PutAPIGithubActionsJobsV3Conflict describes a response with status code 409, with default header values.

Conflict
*/
type PutAPIGithubActionsJobsV3Conflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PutAPIGithubActionsJobsV3Conflict) Error() string {
	return fmt.Sprintf("[PUT /api/github-actions-jobs/v3][%d] putApiGithubActionsJobsV3Conflict  %+v", 409, o.Payload)
}
func (o *PutAPIGithubActionsJobsV3Conflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PutAPIGithubActionsJobsV3Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIGithubActionsJobsV3InternalServerError creates a PutAPIGithubActionsJobsV3InternalServerError with default headers values
func NewPutAPIGithubActionsJobsV3InternalServerError() *PutAPIGithubActionsJobsV3InternalServerError {
	return &PutAPIGithubActionsJobsV3InternalServerError{}
}

/* PutAPIGithubActionsJobsV3InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PutAPIGithubActionsJobsV3InternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PutAPIGithubActionsJobsV3InternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/github-actions-jobs/v3][%d] putApiGithubActionsJobsV3InternalServerError  %+v", 500, o.Payload)
}
func (o *PutAPIGithubActionsJobsV3InternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PutAPIGithubActionsJobsV3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
