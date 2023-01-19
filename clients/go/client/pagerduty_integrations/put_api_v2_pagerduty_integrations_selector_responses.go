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

// PutAPIV2PagerdutyIntegrationsSelectorReader is a Reader for the PutAPIV2PagerdutyIntegrationsSelector structure.
type PutAPIV2PagerdutyIntegrationsSelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAPIV2PagerdutyIntegrationsSelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutAPIV2PagerdutyIntegrationsSelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPutAPIV2PagerdutyIntegrationsSelectorCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutAPIV2PagerdutyIntegrationsSelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutAPIV2PagerdutyIntegrationsSelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutAPIV2PagerdutyIntegrationsSelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewPutAPIV2PagerdutyIntegrationsSelectorProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutAPIV2PagerdutyIntegrationsSelectorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutAPIV2PagerdutyIntegrationsSelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutAPIV2PagerdutyIntegrationsSelectorOK creates a PutAPIV2PagerdutyIntegrationsSelectorOK with default headers values
func NewPutAPIV2PagerdutyIntegrationsSelectorOK() *PutAPIV2PagerdutyIntegrationsSelectorOK {
	return &PutAPIV2PagerdutyIntegrationsSelectorOK{}
}

/* PutAPIV2PagerdutyIntegrationsSelectorOK describes a response with status code 200, with default header values.

OK
*/
type PutAPIV2PagerdutyIntegrationsSelectorOK struct {
	Payload *models.V2controllersPagerdutyIntegration
}

func (o *PutAPIV2PagerdutyIntegrationsSelectorOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/pagerduty-integrations/{selector}][%d] putApiV2PagerdutyIntegrationsSelectorOK  %+v", 200, o.Payload)
}
func (o *PutAPIV2PagerdutyIntegrationsSelectorOK) GetPayload() *models.V2controllersPagerdutyIntegration {
	return o.Payload
}

func (o *PutAPIV2PagerdutyIntegrationsSelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V2controllersPagerdutyIntegration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIV2PagerdutyIntegrationsSelectorCreated creates a PutAPIV2PagerdutyIntegrationsSelectorCreated with default headers values
func NewPutAPIV2PagerdutyIntegrationsSelectorCreated() *PutAPIV2PagerdutyIntegrationsSelectorCreated {
	return &PutAPIV2PagerdutyIntegrationsSelectorCreated{}
}

/* PutAPIV2PagerdutyIntegrationsSelectorCreated describes a response with status code 201, with default header values.

Created
*/
type PutAPIV2PagerdutyIntegrationsSelectorCreated struct {
	Payload *models.V2controllersPagerdutyIntegration
}

func (o *PutAPIV2PagerdutyIntegrationsSelectorCreated) Error() string {
	return fmt.Sprintf("[PUT /api/v2/pagerduty-integrations/{selector}][%d] putApiV2PagerdutyIntegrationsSelectorCreated  %+v", 201, o.Payload)
}
func (o *PutAPIV2PagerdutyIntegrationsSelectorCreated) GetPayload() *models.V2controllersPagerdutyIntegration {
	return o.Payload
}

func (o *PutAPIV2PagerdutyIntegrationsSelectorCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V2controllersPagerdutyIntegration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIV2PagerdutyIntegrationsSelectorBadRequest creates a PutAPIV2PagerdutyIntegrationsSelectorBadRequest with default headers values
func NewPutAPIV2PagerdutyIntegrationsSelectorBadRequest() *PutAPIV2PagerdutyIntegrationsSelectorBadRequest {
	return &PutAPIV2PagerdutyIntegrationsSelectorBadRequest{}
}

/* PutAPIV2PagerdutyIntegrationsSelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutAPIV2PagerdutyIntegrationsSelectorBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PutAPIV2PagerdutyIntegrationsSelectorBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/pagerduty-integrations/{selector}][%d] putApiV2PagerdutyIntegrationsSelectorBadRequest  %+v", 400, o.Payload)
}
func (o *PutAPIV2PagerdutyIntegrationsSelectorBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PutAPIV2PagerdutyIntegrationsSelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIV2PagerdutyIntegrationsSelectorForbidden creates a PutAPIV2PagerdutyIntegrationsSelectorForbidden with default headers values
func NewPutAPIV2PagerdutyIntegrationsSelectorForbidden() *PutAPIV2PagerdutyIntegrationsSelectorForbidden {
	return &PutAPIV2PagerdutyIntegrationsSelectorForbidden{}
}

/* PutAPIV2PagerdutyIntegrationsSelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutAPIV2PagerdutyIntegrationsSelectorForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PutAPIV2PagerdutyIntegrationsSelectorForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/pagerduty-integrations/{selector}][%d] putApiV2PagerdutyIntegrationsSelectorForbidden  %+v", 403, o.Payload)
}
func (o *PutAPIV2PagerdutyIntegrationsSelectorForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PutAPIV2PagerdutyIntegrationsSelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIV2PagerdutyIntegrationsSelectorNotFound creates a PutAPIV2PagerdutyIntegrationsSelectorNotFound with default headers values
func NewPutAPIV2PagerdutyIntegrationsSelectorNotFound() *PutAPIV2PagerdutyIntegrationsSelectorNotFound {
	return &PutAPIV2PagerdutyIntegrationsSelectorNotFound{}
}

/* PutAPIV2PagerdutyIntegrationsSelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PutAPIV2PagerdutyIntegrationsSelectorNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PutAPIV2PagerdutyIntegrationsSelectorNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/pagerduty-integrations/{selector}][%d] putApiV2PagerdutyIntegrationsSelectorNotFound  %+v", 404, o.Payload)
}
func (o *PutAPIV2PagerdutyIntegrationsSelectorNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PutAPIV2PagerdutyIntegrationsSelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIV2PagerdutyIntegrationsSelectorProxyAuthenticationRequired creates a PutAPIV2PagerdutyIntegrationsSelectorProxyAuthenticationRequired with default headers values
func NewPutAPIV2PagerdutyIntegrationsSelectorProxyAuthenticationRequired() *PutAPIV2PagerdutyIntegrationsSelectorProxyAuthenticationRequired {
	return &PutAPIV2PagerdutyIntegrationsSelectorProxyAuthenticationRequired{}
}

/* PutAPIV2PagerdutyIntegrationsSelectorProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type PutAPIV2PagerdutyIntegrationsSelectorProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PutAPIV2PagerdutyIntegrationsSelectorProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[PUT /api/v2/pagerduty-integrations/{selector}][%d] putApiV2PagerdutyIntegrationsSelectorProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *PutAPIV2PagerdutyIntegrationsSelectorProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PutAPIV2PagerdutyIntegrationsSelectorProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIV2PagerdutyIntegrationsSelectorConflict creates a PutAPIV2PagerdutyIntegrationsSelectorConflict with default headers values
func NewPutAPIV2PagerdutyIntegrationsSelectorConflict() *PutAPIV2PagerdutyIntegrationsSelectorConflict {
	return &PutAPIV2PagerdutyIntegrationsSelectorConflict{}
}

/* PutAPIV2PagerdutyIntegrationsSelectorConflict describes a response with status code 409, with default header values.

Conflict
*/
type PutAPIV2PagerdutyIntegrationsSelectorConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PutAPIV2PagerdutyIntegrationsSelectorConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/pagerduty-integrations/{selector}][%d] putApiV2PagerdutyIntegrationsSelectorConflict  %+v", 409, o.Payload)
}
func (o *PutAPIV2PagerdutyIntegrationsSelectorConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PutAPIV2PagerdutyIntegrationsSelectorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIV2PagerdutyIntegrationsSelectorInternalServerError creates a PutAPIV2PagerdutyIntegrationsSelectorInternalServerError with default headers values
func NewPutAPIV2PagerdutyIntegrationsSelectorInternalServerError() *PutAPIV2PagerdutyIntegrationsSelectorInternalServerError {
	return &PutAPIV2PagerdutyIntegrationsSelectorInternalServerError{}
}

/* PutAPIV2PagerdutyIntegrationsSelectorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PutAPIV2PagerdutyIntegrationsSelectorInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PutAPIV2PagerdutyIntegrationsSelectorInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/pagerduty-integrations/{selector}][%d] putApiV2PagerdutyIntegrationsSelectorInternalServerError  %+v", 500, o.Payload)
}
func (o *PutAPIV2PagerdutyIntegrationsSelectorInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PutAPIV2PagerdutyIntegrationsSelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
