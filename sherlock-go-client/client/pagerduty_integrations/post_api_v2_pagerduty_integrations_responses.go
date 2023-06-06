// Code generated by go-swagger; DO NOT EDIT.

package pagerduty_integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// PostAPIV2PagerdutyIntegrationsReader is a Reader for the PostAPIV2PagerdutyIntegrations structure.
type PostAPIV2PagerdutyIntegrationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIV2PagerdutyIntegrationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAPIV2PagerdutyIntegrationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostAPIV2PagerdutyIntegrationsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAPIV2PagerdutyIntegrationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAPIV2PagerdutyIntegrationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAPIV2PagerdutyIntegrationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewPostAPIV2PagerdutyIntegrationsProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostAPIV2PagerdutyIntegrationsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAPIV2PagerdutyIntegrationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAPIV2PagerdutyIntegrationsOK creates a PostAPIV2PagerdutyIntegrationsOK with default headers values
func NewPostAPIV2PagerdutyIntegrationsOK() *PostAPIV2PagerdutyIntegrationsOK {
	return &PostAPIV2PagerdutyIntegrationsOK{}
}

/* PostAPIV2PagerdutyIntegrationsOK describes a response with status code 200, with default header values.

OK
*/
type PostAPIV2PagerdutyIntegrationsOK struct {
	Payload *models.V2controllersPagerdutyIntegration
}

func (o *PostAPIV2PagerdutyIntegrationsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/pagerduty-integrations][%d] postApiV2PagerdutyIntegrationsOK  %+v", 200, o.Payload)
}
func (o *PostAPIV2PagerdutyIntegrationsOK) GetPayload() *models.V2controllersPagerdutyIntegration {
	return o.Payload
}

func (o *PostAPIV2PagerdutyIntegrationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V2controllersPagerdutyIntegration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2PagerdutyIntegrationsCreated creates a PostAPIV2PagerdutyIntegrationsCreated with default headers values
func NewPostAPIV2PagerdutyIntegrationsCreated() *PostAPIV2PagerdutyIntegrationsCreated {
	return &PostAPIV2PagerdutyIntegrationsCreated{}
}

/* PostAPIV2PagerdutyIntegrationsCreated describes a response with status code 201, with default header values.

Created
*/
type PostAPIV2PagerdutyIntegrationsCreated struct {
	Payload *models.V2controllersPagerdutyIntegration
}

func (o *PostAPIV2PagerdutyIntegrationsCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/pagerduty-integrations][%d] postApiV2PagerdutyIntegrationsCreated  %+v", 201, o.Payload)
}
func (o *PostAPIV2PagerdutyIntegrationsCreated) GetPayload() *models.V2controllersPagerdutyIntegration {
	return o.Payload
}

func (o *PostAPIV2PagerdutyIntegrationsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V2controllersPagerdutyIntegration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2PagerdutyIntegrationsBadRequest creates a PostAPIV2PagerdutyIntegrationsBadRequest with default headers values
func NewPostAPIV2PagerdutyIntegrationsBadRequest() *PostAPIV2PagerdutyIntegrationsBadRequest {
	return &PostAPIV2PagerdutyIntegrationsBadRequest{}
}

/* PostAPIV2PagerdutyIntegrationsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostAPIV2PagerdutyIntegrationsBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIV2PagerdutyIntegrationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/pagerduty-integrations][%d] postApiV2PagerdutyIntegrationsBadRequest  %+v", 400, o.Payload)
}
func (o *PostAPIV2PagerdutyIntegrationsBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2PagerdutyIntegrationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2PagerdutyIntegrationsForbidden creates a PostAPIV2PagerdutyIntegrationsForbidden with default headers values
func NewPostAPIV2PagerdutyIntegrationsForbidden() *PostAPIV2PagerdutyIntegrationsForbidden {
	return &PostAPIV2PagerdutyIntegrationsForbidden{}
}

/* PostAPIV2PagerdutyIntegrationsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostAPIV2PagerdutyIntegrationsForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIV2PagerdutyIntegrationsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/pagerduty-integrations][%d] postApiV2PagerdutyIntegrationsForbidden  %+v", 403, o.Payload)
}
func (o *PostAPIV2PagerdutyIntegrationsForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2PagerdutyIntegrationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2PagerdutyIntegrationsNotFound creates a PostAPIV2PagerdutyIntegrationsNotFound with default headers values
func NewPostAPIV2PagerdutyIntegrationsNotFound() *PostAPIV2PagerdutyIntegrationsNotFound {
	return &PostAPIV2PagerdutyIntegrationsNotFound{}
}

/* PostAPIV2PagerdutyIntegrationsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostAPIV2PagerdutyIntegrationsNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIV2PagerdutyIntegrationsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/pagerduty-integrations][%d] postApiV2PagerdutyIntegrationsNotFound  %+v", 404, o.Payload)
}
func (o *PostAPIV2PagerdutyIntegrationsNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2PagerdutyIntegrationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2PagerdutyIntegrationsProxyAuthenticationRequired creates a PostAPIV2PagerdutyIntegrationsProxyAuthenticationRequired with default headers values
func NewPostAPIV2PagerdutyIntegrationsProxyAuthenticationRequired() *PostAPIV2PagerdutyIntegrationsProxyAuthenticationRequired {
	return &PostAPIV2PagerdutyIntegrationsProxyAuthenticationRequired{}
}

/* PostAPIV2PagerdutyIntegrationsProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type PostAPIV2PagerdutyIntegrationsProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIV2PagerdutyIntegrationsProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[POST /api/v2/pagerduty-integrations][%d] postApiV2PagerdutyIntegrationsProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *PostAPIV2PagerdutyIntegrationsProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2PagerdutyIntegrationsProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2PagerdutyIntegrationsConflict creates a PostAPIV2PagerdutyIntegrationsConflict with default headers values
func NewPostAPIV2PagerdutyIntegrationsConflict() *PostAPIV2PagerdutyIntegrationsConflict {
	return &PostAPIV2PagerdutyIntegrationsConflict{}
}

/* PostAPIV2PagerdutyIntegrationsConflict describes a response with status code 409, with default header values.

Conflict
*/
type PostAPIV2PagerdutyIntegrationsConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIV2PagerdutyIntegrationsConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/pagerduty-integrations][%d] postApiV2PagerdutyIntegrationsConflict  %+v", 409, o.Payload)
}
func (o *PostAPIV2PagerdutyIntegrationsConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2PagerdutyIntegrationsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2PagerdutyIntegrationsInternalServerError creates a PostAPIV2PagerdutyIntegrationsInternalServerError with default headers values
func NewPostAPIV2PagerdutyIntegrationsInternalServerError() *PostAPIV2PagerdutyIntegrationsInternalServerError {
	return &PostAPIV2PagerdutyIntegrationsInternalServerError{}
}

/* PostAPIV2PagerdutyIntegrationsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostAPIV2PagerdutyIntegrationsInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIV2PagerdutyIntegrationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/pagerduty-integrations][%d] postApiV2PagerdutyIntegrationsInternalServerError  %+v", 500, o.Payload)
}
func (o *PostAPIV2PagerdutyIntegrationsInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2PagerdutyIntegrationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
