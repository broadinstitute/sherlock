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

// DeleteAPIV2PagerdutyIntegrationsSelectorReader is a Reader for the DeleteAPIV2PagerdutyIntegrationsSelector structure.
type DeleteAPIV2PagerdutyIntegrationsSelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIV2PagerdutyIntegrationsSelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAPIV2PagerdutyIntegrationsSelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteAPIV2PagerdutyIntegrationsSelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAPIV2PagerdutyIntegrationsSelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAPIV2PagerdutyIntegrationsSelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewDeleteAPIV2PagerdutyIntegrationsSelectorProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteAPIV2PagerdutyIntegrationsSelectorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAPIV2PagerdutyIntegrationsSelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAPIV2PagerdutyIntegrationsSelectorOK creates a DeleteAPIV2PagerdutyIntegrationsSelectorOK with default headers values
func NewDeleteAPIV2PagerdutyIntegrationsSelectorOK() *DeleteAPIV2PagerdutyIntegrationsSelectorOK {
	return &DeleteAPIV2PagerdutyIntegrationsSelectorOK{}
}

/* DeleteAPIV2PagerdutyIntegrationsSelectorOK describes a response with status code 200, with default header values.

OK
*/
type DeleteAPIV2PagerdutyIntegrationsSelectorOK struct {
	Payload *models.V2controllersPagerdutyIntegration
}

func (o *DeleteAPIV2PagerdutyIntegrationsSelectorOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/pagerduty-integrations/{selector}][%d] deleteApiV2PagerdutyIntegrationsSelectorOK  %+v", 200, o.Payload)
}
func (o *DeleteAPIV2PagerdutyIntegrationsSelectorOK) GetPayload() *models.V2controllersPagerdutyIntegration {
	return o.Payload
}

func (o *DeleteAPIV2PagerdutyIntegrationsSelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V2controllersPagerdutyIntegration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIV2PagerdutyIntegrationsSelectorBadRequest creates a DeleteAPIV2PagerdutyIntegrationsSelectorBadRequest with default headers values
func NewDeleteAPIV2PagerdutyIntegrationsSelectorBadRequest() *DeleteAPIV2PagerdutyIntegrationsSelectorBadRequest {
	return &DeleteAPIV2PagerdutyIntegrationsSelectorBadRequest{}
}

/* DeleteAPIV2PagerdutyIntegrationsSelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteAPIV2PagerdutyIntegrationsSelectorBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIV2PagerdutyIntegrationsSelectorBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/pagerduty-integrations/{selector}][%d] deleteApiV2PagerdutyIntegrationsSelectorBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteAPIV2PagerdutyIntegrationsSelectorBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIV2PagerdutyIntegrationsSelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIV2PagerdutyIntegrationsSelectorForbidden creates a DeleteAPIV2PagerdutyIntegrationsSelectorForbidden with default headers values
func NewDeleteAPIV2PagerdutyIntegrationsSelectorForbidden() *DeleteAPIV2PagerdutyIntegrationsSelectorForbidden {
	return &DeleteAPIV2PagerdutyIntegrationsSelectorForbidden{}
}

/* DeleteAPIV2PagerdutyIntegrationsSelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteAPIV2PagerdutyIntegrationsSelectorForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIV2PagerdutyIntegrationsSelectorForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/pagerduty-integrations/{selector}][%d] deleteApiV2PagerdutyIntegrationsSelectorForbidden  %+v", 403, o.Payload)
}
func (o *DeleteAPIV2PagerdutyIntegrationsSelectorForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIV2PagerdutyIntegrationsSelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIV2PagerdutyIntegrationsSelectorNotFound creates a DeleteAPIV2PagerdutyIntegrationsSelectorNotFound with default headers values
func NewDeleteAPIV2PagerdutyIntegrationsSelectorNotFound() *DeleteAPIV2PagerdutyIntegrationsSelectorNotFound {
	return &DeleteAPIV2PagerdutyIntegrationsSelectorNotFound{}
}

/* DeleteAPIV2PagerdutyIntegrationsSelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteAPIV2PagerdutyIntegrationsSelectorNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIV2PagerdutyIntegrationsSelectorNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/pagerduty-integrations/{selector}][%d] deleteApiV2PagerdutyIntegrationsSelectorNotFound  %+v", 404, o.Payload)
}
func (o *DeleteAPIV2PagerdutyIntegrationsSelectorNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIV2PagerdutyIntegrationsSelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIV2PagerdutyIntegrationsSelectorProxyAuthenticationRequired creates a DeleteAPIV2PagerdutyIntegrationsSelectorProxyAuthenticationRequired with default headers values
func NewDeleteAPIV2PagerdutyIntegrationsSelectorProxyAuthenticationRequired() *DeleteAPIV2PagerdutyIntegrationsSelectorProxyAuthenticationRequired {
	return &DeleteAPIV2PagerdutyIntegrationsSelectorProxyAuthenticationRequired{}
}

/* DeleteAPIV2PagerdutyIntegrationsSelectorProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type DeleteAPIV2PagerdutyIntegrationsSelectorProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIV2PagerdutyIntegrationsSelectorProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/pagerduty-integrations/{selector}][%d] deleteApiV2PagerdutyIntegrationsSelectorProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *DeleteAPIV2PagerdutyIntegrationsSelectorProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIV2PagerdutyIntegrationsSelectorProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIV2PagerdutyIntegrationsSelectorConflict creates a DeleteAPIV2PagerdutyIntegrationsSelectorConflict with default headers values
func NewDeleteAPIV2PagerdutyIntegrationsSelectorConflict() *DeleteAPIV2PagerdutyIntegrationsSelectorConflict {
	return &DeleteAPIV2PagerdutyIntegrationsSelectorConflict{}
}

/* DeleteAPIV2PagerdutyIntegrationsSelectorConflict describes a response with status code 409, with default header values.

Conflict
*/
type DeleteAPIV2PagerdutyIntegrationsSelectorConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIV2PagerdutyIntegrationsSelectorConflict) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/pagerduty-integrations/{selector}][%d] deleteApiV2PagerdutyIntegrationsSelectorConflict  %+v", 409, o.Payload)
}
func (o *DeleteAPIV2PagerdutyIntegrationsSelectorConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIV2PagerdutyIntegrationsSelectorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIV2PagerdutyIntegrationsSelectorInternalServerError creates a DeleteAPIV2PagerdutyIntegrationsSelectorInternalServerError with default headers values
func NewDeleteAPIV2PagerdutyIntegrationsSelectorInternalServerError() *DeleteAPIV2PagerdutyIntegrationsSelectorInternalServerError {
	return &DeleteAPIV2PagerdutyIntegrationsSelectorInternalServerError{}
}

/* DeleteAPIV2PagerdutyIntegrationsSelectorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteAPIV2PagerdutyIntegrationsSelectorInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *DeleteAPIV2PagerdutyIntegrationsSelectorInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/pagerduty-integrations/{selector}][%d] deleteApiV2PagerdutyIntegrationsSelectorInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteAPIV2PagerdutyIntegrationsSelectorInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIV2PagerdutyIntegrationsSelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
