// Code generated by go-swagger; DO NOT EDIT.

package misc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// GetStatusReader is a Reader for the GetStatus structure.
type GetStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetStatusOK creates a GetStatusOK with default headers values
func NewGetStatusOK() *GetStatusOK {
	return &GetStatusOK{}
}

/* GetStatusOK describes a response with status code 200, with default header values.

OK
*/
type GetStatusOK struct {
	Payload *models.MiscStatusResponse
}

func (o *GetStatusOK) Error() string {
	return fmt.Sprintf("[GET /status][%d] getStatusOK  %+v", 200, o.Payload)
}
func (o *GetStatusOK) GetPayload() *models.MiscStatusResponse {
	return o.Payload
}

func (o *GetStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MiscStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}