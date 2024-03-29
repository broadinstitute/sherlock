// Code generated by go-swagger; DO NOT EDIT.

package chart_versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// GetAPIChartVersionsProceduresV3ChangelogReader is a Reader for the GetAPIChartVersionsProceduresV3Changelog structure.
type GetAPIChartVersionsProceduresV3ChangelogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIChartVersionsProceduresV3ChangelogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIChartVersionsProceduresV3ChangelogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIChartVersionsProceduresV3ChangelogBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIChartVersionsProceduresV3ChangelogForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIChartVersionsProceduresV3ChangelogNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPIChartVersionsProceduresV3ChangelogProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPIChartVersionsProceduresV3ChangelogConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIChartVersionsProceduresV3ChangelogInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIChartVersionsProceduresV3ChangelogOK creates a GetAPIChartVersionsProceduresV3ChangelogOK with default headers values
func NewGetAPIChartVersionsProceduresV3ChangelogOK() *GetAPIChartVersionsProceduresV3ChangelogOK {
	return &GetAPIChartVersionsProceduresV3ChangelogOK{}
}

/* GetAPIChartVersionsProceduresV3ChangelogOK describes a response with status code 200, with default header values.

OK
*/
type GetAPIChartVersionsProceduresV3ChangelogOK struct {
	Payload *models.SherlockChartVersionV3ChangelogResponse
}

func (o *GetAPIChartVersionsProceduresV3ChangelogOK) Error() string {
	return fmt.Sprintf("[GET /api/chart-versions/procedures/v3/changelog][%d] getApiChartVersionsProceduresV3ChangelogOK  %+v", 200, o.Payload)
}
func (o *GetAPIChartVersionsProceduresV3ChangelogOK) GetPayload() *models.SherlockChartVersionV3ChangelogResponse {
	return o.Payload
}

func (o *GetAPIChartVersionsProceduresV3ChangelogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SherlockChartVersionV3ChangelogResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChartVersionsProceduresV3ChangelogBadRequest creates a GetAPIChartVersionsProceduresV3ChangelogBadRequest with default headers values
func NewGetAPIChartVersionsProceduresV3ChangelogBadRequest() *GetAPIChartVersionsProceduresV3ChangelogBadRequest {
	return &GetAPIChartVersionsProceduresV3ChangelogBadRequest{}
}

/* GetAPIChartVersionsProceduresV3ChangelogBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIChartVersionsProceduresV3ChangelogBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChartVersionsProceduresV3ChangelogBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/chart-versions/procedures/v3/changelog][%d] getApiChartVersionsProceduresV3ChangelogBadRequest  %+v", 400, o.Payload)
}
func (o *GetAPIChartVersionsProceduresV3ChangelogBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChartVersionsProceduresV3ChangelogBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChartVersionsProceduresV3ChangelogForbidden creates a GetAPIChartVersionsProceduresV3ChangelogForbidden with default headers values
func NewGetAPIChartVersionsProceduresV3ChangelogForbidden() *GetAPIChartVersionsProceduresV3ChangelogForbidden {
	return &GetAPIChartVersionsProceduresV3ChangelogForbidden{}
}

/* GetAPIChartVersionsProceduresV3ChangelogForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPIChartVersionsProceduresV3ChangelogForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChartVersionsProceduresV3ChangelogForbidden) Error() string {
	return fmt.Sprintf("[GET /api/chart-versions/procedures/v3/changelog][%d] getApiChartVersionsProceduresV3ChangelogForbidden  %+v", 403, o.Payload)
}
func (o *GetAPIChartVersionsProceduresV3ChangelogForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChartVersionsProceduresV3ChangelogForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChartVersionsProceduresV3ChangelogNotFound creates a GetAPIChartVersionsProceduresV3ChangelogNotFound with default headers values
func NewGetAPIChartVersionsProceduresV3ChangelogNotFound() *GetAPIChartVersionsProceduresV3ChangelogNotFound {
	return &GetAPIChartVersionsProceduresV3ChangelogNotFound{}
}

/* GetAPIChartVersionsProceduresV3ChangelogNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIChartVersionsProceduresV3ChangelogNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChartVersionsProceduresV3ChangelogNotFound) Error() string {
	return fmt.Sprintf("[GET /api/chart-versions/procedures/v3/changelog][%d] getApiChartVersionsProceduresV3ChangelogNotFound  %+v", 404, o.Payload)
}
func (o *GetAPIChartVersionsProceduresV3ChangelogNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChartVersionsProceduresV3ChangelogNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChartVersionsProceduresV3ChangelogProxyAuthenticationRequired creates a GetAPIChartVersionsProceduresV3ChangelogProxyAuthenticationRequired with default headers values
func NewGetAPIChartVersionsProceduresV3ChangelogProxyAuthenticationRequired() *GetAPIChartVersionsProceduresV3ChangelogProxyAuthenticationRequired {
	return &GetAPIChartVersionsProceduresV3ChangelogProxyAuthenticationRequired{}
}

/* GetAPIChartVersionsProceduresV3ChangelogProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPIChartVersionsProceduresV3ChangelogProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChartVersionsProceduresV3ChangelogProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/chart-versions/procedures/v3/changelog][%d] getApiChartVersionsProceduresV3ChangelogProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPIChartVersionsProceduresV3ChangelogProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChartVersionsProceduresV3ChangelogProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChartVersionsProceduresV3ChangelogConflict creates a GetAPIChartVersionsProceduresV3ChangelogConflict with default headers values
func NewGetAPIChartVersionsProceduresV3ChangelogConflict() *GetAPIChartVersionsProceduresV3ChangelogConflict {
	return &GetAPIChartVersionsProceduresV3ChangelogConflict{}
}

/* GetAPIChartVersionsProceduresV3ChangelogConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPIChartVersionsProceduresV3ChangelogConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChartVersionsProceduresV3ChangelogConflict) Error() string {
	return fmt.Sprintf("[GET /api/chart-versions/procedures/v3/changelog][%d] getApiChartVersionsProceduresV3ChangelogConflict  %+v", 409, o.Payload)
}
func (o *GetAPIChartVersionsProceduresV3ChangelogConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChartVersionsProceduresV3ChangelogConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChartVersionsProceduresV3ChangelogInternalServerError creates a GetAPIChartVersionsProceduresV3ChangelogInternalServerError with default headers values
func NewGetAPIChartVersionsProceduresV3ChangelogInternalServerError() *GetAPIChartVersionsProceduresV3ChangelogInternalServerError {
	return &GetAPIChartVersionsProceduresV3ChangelogInternalServerError{}
}

/* GetAPIChartVersionsProceduresV3ChangelogInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPIChartVersionsProceduresV3ChangelogInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChartVersionsProceduresV3ChangelogInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/chart-versions/procedures/v3/changelog][%d] getApiChartVersionsProceduresV3ChangelogInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPIChartVersionsProceduresV3ChangelogInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChartVersionsProceduresV3ChangelogInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
