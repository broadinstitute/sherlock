// Code generated by go-swagger; DO NOT EDIT.

package pagerduty_integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/clients/go/client/models"
)

// PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorReader is a Reader for the PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelector structure.
type PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewPostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorAccepted creates a PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorAccepted with default headers values
func NewPostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorAccepted() *PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorAccepted {
	return &PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorAccepted{}
}

/* PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorAccepted struct {
	Payload *models.PagerdutySendAlertResponse
}

func (o *PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/procedures/pagerduty-integrations/trigger-incident/{selector}][%d] postApiV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorAccepted  %+v", 202, o.Payload)
}
func (o *PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorAccepted) GetPayload() *models.PagerdutySendAlertResponse {
	return o.Payload
}

func (o *PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PagerdutySendAlertResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorBadRequest creates a PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorBadRequest with default headers values
func NewPostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorBadRequest() *PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorBadRequest {
	return &PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorBadRequest{}
}

/* PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/procedures/pagerduty-integrations/trigger-incident/{selector}][%d] postApiV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorBadRequest  %+v", 400, o.Payload)
}
func (o *PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorForbidden creates a PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorForbidden with default headers values
func NewPostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorForbidden() *PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorForbidden {
	return &PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorForbidden{}
}

/* PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/procedures/pagerduty-integrations/trigger-incident/{selector}][%d] postApiV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorForbidden  %+v", 403, o.Payload)
}
func (o *PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorNotFound creates a PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorNotFound with default headers values
func NewPostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorNotFound() *PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorNotFound {
	return &PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorNotFound{}
}

/* PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/procedures/pagerduty-integrations/trigger-incident/{selector}][%d] postApiV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorNotFound  %+v", 404, o.Payload)
}
func (o *PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorProxyAuthenticationRequired creates a PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorProxyAuthenticationRequired with default headers values
func NewPostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorProxyAuthenticationRequired() *PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorProxyAuthenticationRequired {
	return &PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorProxyAuthenticationRequired{}
}

/* PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[POST /api/v2/procedures/pagerduty-integrations/trigger-incident/{selector}][%d] postApiV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorConflict creates a PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorConflict with default headers values
func NewPostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorConflict() *PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorConflict {
	return &PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorConflict{}
}

/* PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorConflict describes a response with status code 409, with default header values.

Conflict
*/
type PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/procedures/pagerduty-integrations/trigger-incident/{selector}][%d] postApiV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorConflict  %+v", 409, o.Payload)
}
func (o *PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorInternalServerError creates a PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorInternalServerError with default headers values
func NewPostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorInternalServerError() *PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorInternalServerError {
	return &PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorInternalServerError{}
}

/* PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/procedures/pagerduty-integrations/trigger-incident/{selector}][%d] postApiV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorInternalServerError  %+v", 500, o.Payload)
}
func (o *PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2ProceduresPagerdutyIntegrationsTriggerIncidentSelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
