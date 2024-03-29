// Code generated by go-swagger; DO NOT EDIT.

package changesets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseReader is a Reader for the GetAPIChangesetsProceduresV3ChartReleaseHistoryChartRelease structure.
type GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseOK creates a GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseOK with default headers values
func NewGetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseOK() *GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseOK {
	return &GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseOK{}
}

/* GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseOK describes a response with status code 200, with default header values.

OK
*/
type GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseOK struct {
	Payload []*models.SherlockChangesetV3
}

func (o *GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseOK) Error() string {
	return fmt.Sprintf("[GET /api/changesets/procedures/v3/chart-release-history/{chart-release}][%d] getApiChangesetsProceduresV3ChartReleaseHistoryChartReleaseOK  %+v", 200, o.Payload)
}
func (o *GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseOK) GetPayload() []*models.SherlockChangesetV3 {
	return o.Payload
}

func (o *GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseBadRequest creates a GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseBadRequest with default headers values
func NewGetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseBadRequest() *GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseBadRequest {
	return &GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseBadRequest{}
}

/* GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/changesets/procedures/v3/chart-release-history/{chart-release}][%d] getApiChangesetsProceduresV3ChartReleaseHistoryChartReleaseBadRequest  %+v", 400, o.Payload)
}
func (o *GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseForbidden creates a GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseForbidden with default headers values
func NewGetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseForbidden() *GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseForbidden {
	return &GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseForbidden{}
}

/* GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseForbidden) Error() string {
	return fmt.Sprintf("[GET /api/changesets/procedures/v3/chart-release-history/{chart-release}][%d] getApiChangesetsProceduresV3ChartReleaseHistoryChartReleaseForbidden  %+v", 403, o.Payload)
}
func (o *GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseNotFound creates a GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseNotFound with default headers values
func NewGetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseNotFound() *GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseNotFound {
	return &GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseNotFound{}
}

/* GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseNotFound) Error() string {
	return fmt.Sprintf("[GET /api/changesets/procedures/v3/chart-release-history/{chart-release}][%d] getApiChangesetsProceduresV3ChartReleaseHistoryChartReleaseNotFound  %+v", 404, o.Payload)
}
func (o *GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseProxyAuthenticationRequired creates a GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseProxyAuthenticationRequired with default headers values
func NewGetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseProxyAuthenticationRequired() *GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseProxyAuthenticationRequired {
	return &GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseProxyAuthenticationRequired{}
}

/* GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/changesets/procedures/v3/chart-release-history/{chart-release}][%d] getApiChangesetsProceduresV3ChartReleaseHistoryChartReleaseProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseConflict creates a GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseConflict with default headers values
func NewGetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseConflict() *GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseConflict {
	return &GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseConflict{}
}

/* GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseConflict) Error() string {
	return fmt.Sprintf("[GET /api/changesets/procedures/v3/chart-release-history/{chart-release}][%d] getApiChangesetsProceduresV3ChartReleaseHistoryChartReleaseConflict  %+v", 409, o.Payload)
}
func (o *GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseInternalServerError creates a GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseInternalServerError with default headers values
func NewGetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseInternalServerError() *GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseInternalServerError {
	return &GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseInternalServerError{}
}

/* GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/changesets/procedures/v3/chart-release-history/{chart-release}][%d] getApiChangesetsProceduresV3ChartReleaseHistoryChartReleaseInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIChangesetsProceduresV3ChartReleaseHistoryChartReleaseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
