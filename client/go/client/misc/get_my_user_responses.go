// Code generated by go-swagger; DO NOT EDIT.

package misc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/client/go/client/models"
)

// GetMyUserReader is a Reader for the GetMyUser structure.
type GetMyUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMyUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMyUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 407:
		result := NewGetMyUserProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetMyUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMyUserOK creates a GetMyUserOK with default headers values
func NewGetMyUserOK() *GetMyUserOK {
	return &GetMyUserOK{}
}

/* GetMyUserOK describes a response with status code 200, with default header values.

OK
*/
type GetMyUserOK struct {
	Payload *models.MiscMyUserResponse
}

func (o *GetMyUserOK) Error() string {
	return fmt.Sprintf("[GET /my-user][%d] getMyUserOK  %+v", 200, o.Payload)
}
func (o *GetMyUserOK) GetPayload() *models.MiscMyUserResponse {
	return o.Payload
}

func (o *GetMyUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MiscMyUserResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMyUserProxyAuthenticationRequired creates a GetMyUserProxyAuthenticationRequired with default headers values
func NewGetMyUserProxyAuthenticationRequired() *GetMyUserProxyAuthenticationRequired {
	return &GetMyUserProxyAuthenticationRequired{}
}

/* GetMyUserProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetMyUserProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetMyUserProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /my-user][%d] getMyUserProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetMyUserProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetMyUserProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMyUserInternalServerError creates a GetMyUserInternalServerError with default headers values
func NewGetMyUserInternalServerError() *GetMyUserInternalServerError {
	return &GetMyUserInternalServerError{}
}

/* GetMyUserInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetMyUserInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetMyUserInternalServerError) Error() string {
	return fmt.Sprintf("[GET /my-user][%d] getMyUserInternalServerError  %+v", 500, o.Payload)
}
func (o *GetMyUserInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetMyUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
