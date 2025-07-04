// Code generated by go-swagger; DO NOT EDIT.

package service_alert

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// PostAPIServiceAlertsV3Reader is a Reader for the PostAPIServiceAlertsV3 structure.
type PostAPIServiceAlertsV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIServiceAlertsV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAPIServiceAlertsV3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAPIServiceAlertsV3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAPIServiceAlertsV3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAPIServiceAlertsV3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewPostAPIServiceAlertsV3ProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostAPIServiceAlertsV3Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAPIServiceAlertsV3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAPIServiceAlertsV3OK creates a PostAPIServiceAlertsV3OK with default headers values
func NewPostAPIServiceAlertsV3OK() *PostAPIServiceAlertsV3OK {
	return &PostAPIServiceAlertsV3OK{}
}

/* PostAPIServiceAlertsV3OK describes a response with status code 200, with default header values.

OK
*/
type PostAPIServiceAlertsV3OK struct {
	Payload *models.SherlockServiceAlertV3
}

func (o *PostAPIServiceAlertsV3OK) Error() string {
	return fmt.Sprintf("[POST /api/service-alerts/v3][%d] postApiServiceAlertsV3OK  %+v", 200, o.Payload)
}
func (o *PostAPIServiceAlertsV3OK) GetPayload() *models.SherlockServiceAlertV3 {
	return o.Payload
}

func (o *PostAPIServiceAlertsV3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SherlockServiceAlertV3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIServiceAlertsV3BadRequest creates a PostAPIServiceAlertsV3BadRequest with default headers values
func NewPostAPIServiceAlertsV3BadRequest() *PostAPIServiceAlertsV3BadRequest {
	return &PostAPIServiceAlertsV3BadRequest{}
}

/* PostAPIServiceAlertsV3BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostAPIServiceAlertsV3BadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIServiceAlertsV3BadRequest) Error() string {
	return fmt.Sprintf("[POST /api/service-alerts/v3][%d] postApiServiceAlertsV3BadRequest  %+v", 400, o.Payload)
}
func (o *PostAPIServiceAlertsV3BadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIServiceAlertsV3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIServiceAlertsV3Forbidden creates a PostAPIServiceAlertsV3Forbidden with default headers values
func NewPostAPIServiceAlertsV3Forbidden() *PostAPIServiceAlertsV3Forbidden {
	return &PostAPIServiceAlertsV3Forbidden{}
}

/* PostAPIServiceAlertsV3Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostAPIServiceAlertsV3Forbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIServiceAlertsV3Forbidden) Error() string {
	return fmt.Sprintf("[POST /api/service-alerts/v3][%d] postApiServiceAlertsV3Forbidden  %+v", 403, o.Payload)
}
func (o *PostAPIServiceAlertsV3Forbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIServiceAlertsV3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIServiceAlertsV3NotFound creates a PostAPIServiceAlertsV3NotFound with default headers values
func NewPostAPIServiceAlertsV3NotFound() *PostAPIServiceAlertsV3NotFound {
	return &PostAPIServiceAlertsV3NotFound{}
}

/* PostAPIServiceAlertsV3NotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostAPIServiceAlertsV3NotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIServiceAlertsV3NotFound) Error() string {
	return fmt.Sprintf("[POST /api/service-alerts/v3][%d] postApiServiceAlertsV3NotFound  %+v", 404, o.Payload)
}
func (o *PostAPIServiceAlertsV3NotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIServiceAlertsV3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIServiceAlertsV3ProxyAuthenticationRequired creates a PostAPIServiceAlertsV3ProxyAuthenticationRequired with default headers values
func NewPostAPIServiceAlertsV3ProxyAuthenticationRequired() *PostAPIServiceAlertsV3ProxyAuthenticationRequired {
	return &PostAPIServiceAlertsV3ProxyAuthenticationRequired{}
}

/* PostAPIServiceAlertsV3ProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type PostAPIServiceAlertsV3ProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIServiceAlertsV3ProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[POST /api/service-alerts/v3][%d] postApiServiceAlertsV3ProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *PostAPIServiceAlertsV3ProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIServiceAlertsV3ProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIServiceAlertsV3Conflict creates a PostAPIServiceAlertsV3Conflict with default headers values
func NewPostAPIServiceAlertsV3Conflict() *PostAPIServiceAlertsV3Conflict {
	return &PostAPIServiceAlertsV3Conflict{}
}

/* PostAPIServiceAlertsV3Conflict describes a response with status code 409, with default header values.

Conflict
*/
type PostAPIServiceAlertsV3Conflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIServiceAlertsV3Conflict) Error() string {
	return fmt.Sprintf("[POST /api/service-alerts/v3][%d] postApiServiceAlertsV3Conflict  %+v", 409, o.Payload)
}
func (o *PostAPIServiceAlertsV3Conflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIServiceAlertsV3Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIServiceAlertsV3InternalServerError creates a PostAPIServiceAlertsV3InternalServerError with default headers values
func NewPostAPIServiceAlertsV3InternalServerError() *PostAPIServiceAlertsV3InternalServerError {
	return &PostAPIServiceAlertsV3InternalServerError{}
}

/* PostAPIServiceAlertsV3InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostAPIServiceAlertsV3InternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIServiceAlertsV3InternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/service-alerts/v3][%d] postApiServiceAlertsV3InternalServerError  %+v", 500, o.Payload)
}
func (o *PostAPIServiceAlertsV3InternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIServiceAlertsV3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
