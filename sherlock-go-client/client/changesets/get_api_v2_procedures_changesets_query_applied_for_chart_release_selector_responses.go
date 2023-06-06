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

// GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorReader is a Reader for the GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelector structure.
type GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorOK creates a GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorOK with default headers values
func NewGetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorOK() *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorOK {
	return &GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorOK{}
}

/* GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorOK describes a response with status code 200, with default header values.

OK
*/
type GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorOK struct {
	Payload []*models.V2controllersChangeset
}

func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/procedures/changesets/query-applied-for-chart-release/{selector}][%d] getApiV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorOK  %+v", 200, o.Payload)
}
func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorOK) GetPayload() []*models.V2controllersChangeset {
	return o.Payload
}

func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorBadRequest creates a GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorBadRequest with default headers values
func NewGetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorBadRequest() *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorBadRequest {
	return &GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorBadRequest{}
}

/* GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/procedures/changesets/query-applied-for-chart-release/{selector}][%d] getApiV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorBadRequest  %+v", 400, o.Payload)
}
func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorForbidden creates a GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorForbidden with default headers values
func NewGetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorForbidden() *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorForbidden {
	return &GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorForbidden{}
}

/* GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/procedures/changesets/query-applied-for-chart-release/{selector}][%d] getApiV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorForbidden  %+v", 403, o.Payload)
}
func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorNotFound creates a GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorNotFound with default headers values
func NewGetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorNotFound() *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorNotFound {
	return &GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorNotFound{}
}

/* GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/procedures/changesets/query-applied-for-chart-release/{selector}][%d] getApiV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorNotFound  %+v", 404, o.Payload)
}
func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorProxyAuthenticationRequired creates a GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorProxyAuthenticationRequired with default headers values
func NewGetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorProxyAuthenticationRequired() *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorProxyAuthenticationRequired {
	return &GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorProxyAuthenticationRequired{}
}

/* GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/v2/procedures/changesets/query-applied-for-chart-release/{selector}][%d] getApiV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorConflict creates a GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorConflict with default headers values
func NewGetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorConflict() *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorConflict {
	return &GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorConflict{}
}

/* GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorConflict) Error() string {
	return fmt.Sprintf("[GET /api/v2/procedures/changesets/query-applied-for-chart-release/{selector}][%d] getApiV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorConflict  %+v", 409, o.Payload)
}
func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorInternalServerError creates a GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorInternalServerError with default headers values
func NewGetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorInternalServerError() *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorInternalServerError {
	return &GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorInternalServerError{}
}

/* GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/procedures/changesets/query-applied-for-chart-release/{selector}][%d] getApiV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
