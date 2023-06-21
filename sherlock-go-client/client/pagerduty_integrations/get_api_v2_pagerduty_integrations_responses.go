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

// GetAPIV2PagerdutyIntegrationsReader is a Reader for the GetAPIV2PagerdutyIntegrations structure.
type GetAPIV2PagerdutyIntegrationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV2PagerdutyIntegrationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV2PagerdutyIntegrationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIV2PagerdutyIntegrationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIV2PagerdutyIntegrationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIV2PagerdutyIntegrationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPIV2PagerdutyIntegrationsProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPIV2PagerdutyIntegrationsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIV2PagerdutyIntegrationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIV2PagerdutyIntegrationsOK creates a GetAPIV2PagerdutyIntegrationsOK with default headers values
func NewGetAPIV2PagerdutyIntegrationsOK() *GetAPIV2PagerdutyIntegrationsOK {
	return &GetAPIV2PagerdutyIntegrationsOK{}
}

/* GetAPIV2PagerdutyIntegrationsOK describes a response with status code 200, with default header values.

OK
*/
type GetAPIV2PagerdutyIntegrationsOK struct {
	Payload []*models.V2controllersPagerdutyIntegration
}

func (o *GetAPIV2PagerdutyIntegrationsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/pagerduty-integrations][%d] getApiV2PagerdutyIntegrationsOK  %+v", 200, o.Payload)
}
func (o *GetAPIV2PagerdutyIntegrationsOK) GetPayload() []*models.V2controllersPagerdutyIntegration {
	return o.Payload
}

func (o *GetAPIV2PagerdutyIntegrationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2PagerdutyIntegrationsBadRequest creates a GetAPIV2PagerdutyIntegrationsBadRequest with default headers values
func NewGetAPIV2PagerdutyIntegrationsBadRequest() *GetAPIV2PagerdutyIntegrationsBadRequest {
	return &GetAPIV2PagerdutyIntegrationsBadRequest{}
}

/* GetAPIV2PagerdutyIntegrationsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIV2PagerdutyIntegrationsBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2PagerdutyIntegrationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/pagerduty-integrations][%d] getApiV2PagerdutyIntegrationsBadRequest  %+v", 400, o.Payload)
}
func (o *GetAPIV2PagerdutyIntegrationsBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2PagerdutyIntegrationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2PagerdutyIntegrationsForbidden creates a GetAPIV2PagerdutyIntegrationsForbidden with default headers values
func NewGetAPIV2PagerdutyIntegrationsForbidden() *GetAPIV2PagerdutyIntegrationsForbidden {
	return &GetAPIV2PagerdutyIntegrationsForbidden{}
}

/* GetAPIV2PagerdutyIntegrationsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPIV2PagerdutyIntegrationsForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2PagerdutyIntegrationsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/pagerduty-integrations][%d] getApiV2PagerdutyIntegrationsForbidden  %+v", 403, o.Payload)
}
func (o *GetAPIV2PagerdutyIntegrationsForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2PagerdutyIntegrationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2PagerdutyIntegrationsNotFound creates a GetAPIV2PagerdutyIntegrationsNotFound with default headers values
func NewGetAPIV2PagerdutyIntegrationsNotFound() *GetAPIV2PagerdutyIntegrationsNotFound {
	return &GetAPIV2PagerdutyIntegrationsNotFound{}
}

/* GetAPIV2PagerdutyIntegrationsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIV2PagerdutyIntegrationsNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2PagerdutyIntegrationsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/pagerduty-integrations][%d] getApiV2PagerdutyIntegrationsNotFound  %+v", 404, o.Payload)
}
func (o *GetAPIV2PagerdutyIntegrationsNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2PagerdutyIntegrationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2PagerdutyIntegrationsProxyAuthenticationRequired creates a GetAPIV2PagerdutyIntegrationsProxyAuthenticationRequired with default headers values
func NewGetAPIV2PagerdutyIntegrationsProxyAuthenticationRequired() *GetAPIV2PagerdutyIntegrationsProxyAuthenticationRequired {
	return &GetAPIV2PagerdutyIntegrationsProxyAuthenticationRequired{}
}

/* GetAPIV2PagerdutyIntegrationsProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPIV2PagerdutyIntegrationsProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2PagerdutyIntegrationsProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/v2/pagerduty-integrations][%d] getApiV2PagerdutyIntegrationsProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPIV2PagerdutyIntegrationsProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2PagerdutyIntegrationsProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2PagerdutyIntegrationsConflict creates a GetAPIV2PagerdutyIntegrationsConflict with default headers values
func NewGetAPIV2PagerdutyIntegrationsConflict() *GetAPIV2PagerdutyIntegrationsConflict {
	return &GetAPIV2PagerdutyIntegrationsConflict{}
}

/* GetAPIV2PagerdutyIntegrationsConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPIV2PagerdutyIntegrationsConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2PagerdutyIntegrationsConflict) Error() string {
	return fmt.Sprintf("[GET /api/v2/pagerduty-integrations][%d] getApiV2PagerdutyIntegrationsConflict  %+v", 409, o.Payload)
}
func (o *GetAPIV2PagerdutyIntegrationsConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2PagerdutyIntegrationsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2PagerdutyIntegrationsInternalServerError creates a GetAPIV2PagerdutyIntegrationsInternalServerError with default headers values
func NewGetAPIV2PagerdutyIntegrationsInternalServerError() *GetAPIV2PagerdutyIntegrationsInternalServerError {
	return &GetAPIV2PagerdutyIntegrationsInternalServerError{}
}

/* GetAPIV2PagerdutyIntegrationsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPIV2PagerdutyIntegrationsInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2PagerdutyIntegrationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/pagerduty-integrations][%d] getApiV2PagerdutyIntegrationsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPIV2PagerdutyIntegrationsInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2PagerdutyIntegrationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}