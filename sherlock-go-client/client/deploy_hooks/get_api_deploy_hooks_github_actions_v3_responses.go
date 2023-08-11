// Code generated by go-swagger; DO NOT EDIT.

package deploy_hooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// GetAPIDeployHooksGithubActionsV3Reader is a Reader for the GetAPIDeployHooksGithubActionsV3 structure.
type GetAPIDeployHooksGithubActionsV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIDeployHooksGithubActionsV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIDeployHooksGithubActionsV3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIDeployHooksGithubActionsV3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIDeployHooksGithubActionsV3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIDeployHooksGithubActionsV3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPIDeployHooksGithubActionsV3ProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPIDeployHooksGithubActionsV3Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIDeployHooksGithubActionsV3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIDeployHooksGithubActionsV3OK creates a GetAPIDeployHooksGithubActionsV3OK with default headers values
func NewGetAPIDeployHooksGithubActionsV3OK() *GetAPIDeployHooksGithubActionsV3OK {
	return &GetAPIDeployHooksGithubActionsV3OK{}
}

/* GetAPIDeployHooksGithubActionsV3OK describes a response with status code 200, with default header values.

OK
*/
type GetAPIDeployHooksGithubActionsV3OK struct {
	Payload []*models.SherlockGithubActionsDeployHookV3
}

func (o *GetAPIDeployHooksGithubActionsV3OK) Error() string {
	return fmt.Sprintf("[GET /api/deploy-hooks/github-actions/v3][%d] getApiDeployHooksGithubActionsV3OK  %+v", 200, o.Payload)
}
func (o *GetAPIDeployHooksGithubActionsV3OK) GetPayload() []*models.SherlockGithubActionsDeployHookV3 {
	return o.Payload
}

func (o *GetAPIDeployHooksGithubActionsV3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIDeployHooksGithubActionsV3BadRequest creates a GetAPIDeployHooksGithubActionsV3BadRequest with default headers values
func NewGetAPIDeployHooksGithubActionsV3BadRequest() *GetAPIDeployHooksGithubActionsV3BadRequest {
	return &GetAPIDeployHooksGithubActionsV3BadRequest{}
}

/* GetAPIDeployHooksGithubActionsV3BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIDeployHooksGithubActionsV3BadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIDeployHooksGithubActionsV3BadRequest) Error() string {
	return fmt.Sprintf("[GET /api/deploy-hooks/github-actions/v3][%d] getApiDeployHooksGithubActionsV3BadRequest  %+v", 400, o.Payload)
}
func (o *GetAPIDeployHooksGithubActionsV3BadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIDeployHooksGithubActionsV3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIDeployHooksGithubActionsV3Forbidden creates a GetAPIDeployHooksGithubActionsV3Forbidden with default headers values
func NewGetAPIDeployHooksGithubActionsV3Forbidden() *GetAPIDeployHooksGithubActionsV3Forbidden {
	return &GetAPIDeployHooksGithubActionsV3Forbidden{}
}

/* GetAPIDeployHooksGithubActionsV3Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPIDeployHooksGithubActionsV3Forbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIDeployHooksGithubActionsV3Forbidden) Error() string {
	return fmt.Sprintf("[GET /api/deploy-hooks/github-actions/v3][%d] getApiDeployHooksGithubActionsV3Forbidden  %+v", 403, o.Payload)
}
func (o *GetAPIDeployHooksGithubActionsV3Forbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIDeployHooksGithubActionsV3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIDeployHooksGithubActionsV3NotFound creates a GetAPIDeployHooksGithubActionsV3NotFound with default headers values
func NewGetAPIDeployHooksGithubActionsV3NotFound() *GetAPIDeployHooksGithubActionsV3NotFound {
	return &GetAPIDeployHooksGithubActionsV3NotFound{}
}

/* GetAPIDeployHooksGithubActionsV3NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIDeployHooksGithubActionsV3NotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIDeployHooksGithubActionsV3NotFound) Error() string {
	return fmt.Sprintf("[GET /api/deploy-hooks/github-actions/v3][%d] getApiDeployHooksGithubActionsV3NotFound  %+v", 404, o.Payload)
}
func (o *GetAPIDeployHooksGithubActionsV3NotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIDeployHooksGithubActionsV3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIDeployHooksGithubActionsV3ProxyAuthenticationRequired creates a GetAPIDeployHooksGithubActionsV3ProxyAuthenticationRequired with default headers values
func NewGetAPIDeployHooksGithubActionsV3ProxyAuthenticationRequired() *GetAPIDeployHooksGithubActionsV3ProxyAuthenticationRequired {
	return &GetAPIDeployHooksGithubActionsV3ProxyAuthenticationRequired{}
}

/* GetAPIDeployHooksGithubActionsV3ProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPIDeployHooksGithubActionsV3ProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIDeployHooksGithubActionsV3ProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/deploy-hooks/github-actions/v3][%d] getApiDeployHooksGithubActionsV3ProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPIDeployHooksGithubActionsV3ProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIDeployHooksGithubActionsV3ProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIDeployHooksGithubActionsV3Conflict creates a GetAPIDeployHooksGithubActionsV3Conflict with default headers values
func NewGetAPIDeployHooksGithubActionsV3Conflict() *GetAPIDeployHooksGithubActionsV3Conflict {
	return &GetAPIDeployHooksGithubActionsV3Conflict{}
}

/* GetAPIDeployHooksGithubActionsV3Conflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPIDeployHooksGithubActionsV3Conflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIDeployHooksGithubActionsV3Conflict) Error() string {
	return fmt.Sprintf("[GET /api/deploy-hooks/github-actions/v3][%d] getApiDeployHooksGithubActionsV3Conflict  %+v", 409, o.Payload)
}
func (o *GetAPIDeployHooksGithubActionsV3Conflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIDeployHooksGithubActionsV3Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIDeployHooksGithubActionsV3InternalServerError creates a GetAPIDeployHooksGithubActionsV3InternalServerError with default headers values
func NewGetAPIDeployHooksGithubActionsV3InternalServerError() *GetAPIDeployHooksGithubActionsV3InternalServerError {
	return &GetAPIDeployHooksGithubActionsV3InternalServerError{}
}

/* GetAPIDeployHooksGithubActionsV3InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPIDeployHooksGithubActionsV3InternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIDeployHooksGithubActionsV3InternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/deploy-hooks/github-actions/v3][%d] getApiDeployHooksGithubActionsV3InternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPIDeployHooksGithubActionsV3InternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIDeployHooksGithubActionsV3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}