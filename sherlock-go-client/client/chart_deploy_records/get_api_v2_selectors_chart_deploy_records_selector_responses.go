// Code generated by go-swagger; DO NOT EDIT.

package chart_deploy_records

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// GetAPIV2SelectorsChartDeployRecordsSelectorReader is a Reader for the GetAPIV2SelectorsChartDeployRecordsSelector structure.
type GetAPIV2SelectorsChartDeployRecordsSelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV2SelectorsChartDeployRecordsSelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV2SelectorsChartDeployRecordsSelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIV2SelectorsChartDeployRecordsSelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIV2SelectorsChartDeployRecordsSelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIV2SelectorsChartDeployRecordsSelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPIV2SelectorsChartDeployRecordsSelectorProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPIV2SelectorsChartDeployRecordsSelectorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIV2SelectorsChartDeployRecordsSelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIV2SelectorsChartDeployRecordsSelectorOK creates a GetAPIV2SelectorsChartDeployRecordsSelectorOK with default headers values
func NewGetAPIV2SelectorsChartDeployRecordsSelectorOK() *GetAPIV2SelectorsChartDeployRecordsSelectorOK {
	return &GetAPIV2SelectorsChartDeployRecordsSelectorOK{}
}

/* GetAPIV2SelectorsChartDeployRecordsSelectorOK describes a response with status code 200, with default header values.

OK
*/
type GetAPIV2SelectorsChartDeployRecordsSelectorOK struct {
	Payload []string
}

func (o *GetAPIV2SelectorsChartDeployRecordsSelectorOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/selectors/chart-deploy-records/{selector}][%d] getApiV2SelectorsChartDeployRecordsSelectorOK  %+v", 200, o.Payload)
}
func (o *GetAPIV2SelectorsChartDeployRecordsSelectorOK) GetPayload() []string {
	return o.Payload
}

func (o *GetAPIV2SelectorsChartDeployRecordsSelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2SelectorsChartDeployRecordsSelectorBadRequest creates a GetAPIV2SelectorsChartDeployRecordsSelectorBadRequest with default headers values
func NewGetAPIV2SelectorsChartDeployRecordsSelectorBadRequest() *GetAPIV2SelectorsChartDeployRecordsSelectorBadRequest {
	return &GetAPIV2SelectorsChartDeployRecordsSelectorBadRequest{}
}

/* GetAPIV2SelectorsChartDeployRecordsSelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIV2SelectorsChartDeployRecordsSelectorBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2SelectorsChartDeployRecordsSelectorBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/selectors/chart-deploy-records/{selector}][%d] getApiV2SelectorsChartDeployRecordsSelectorBadRequest  %+v", 400, o.Payload)
}
func (o *GetAPIV2SelectorsChartDeployRecordsSelectorBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2SelectorsChartDeployRecordsSelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2SelectorsChartDeployRecordsSelectorForbidden creates a GetAPIV2SelectorsChartDeployRecordsSelectorForbidden with default headers values
func NewGetAPIV2SelectorsChartDeployRecordsSelectorForbidden() *GetAPIV2SelectorsChartDeployRecordsSelectorForbidden {
	return &GetAPIV2SelectorsChartDeployRecordsSelectorForbidden{}
}

/* GetAPIV2SelectorsChartDeployRecordsSelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPIV2SelectorsChartDeployRecordsSelectorForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2SelectorsChartDeployRecordsSelectorForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/selectors/chart-deploy-records/{selector}][%d] getApiV2SelectorsChartDeployRecordsSelectorForbidden  %+v", 403, o.Payload)
}
func (o *GetAPIV2SelectorsChartDeployRecordsSelectorForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2SelectorsChartDeployRecordsSelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2SelectorsChartDeployRecordsSelectorNotFound creates a GetAPIV2SelectorsChartDeployRecordsSelectorNotFound with default headers values
func NewGetAPIV2SelectorsChartDeployRecordsSelectorNotFound() *GetAPIV2SelectorsChartDeployRecordsSelectorNotFound {
	return &GetAPIV2SelectorsChartDeployRecordsSelectorNotFound{}
}

/* GetAPIV2SelectorsChartDeployRecordsSelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIV2SelectorsChartDeployRecordsSelectorNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2SelectorsChartDeployRecordsSelectorNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/selectors/chart-deploy-records/{selector}][%d] getApiV2SelectorsChartDeployRecordsSelectorNotFound  %+v", 404, o.Payload)
}
func (o *GetAPIV2SelectorsChartDeployRecordsSelectorNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2SelectorsChartDeployRecordsSelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2SelectorsChartDeployRecordsSelectorProxyAuthenticationRequired creates a GetAPIV2SelectorsChartDeployRecordsSelectorProxyAuthenticationRequired with default headers values
func NewGetAPIV2SelectorsChartDeployRecordsSelectorProxyAuthenticationRequired() *GetAPIV2SelectorsChartDeployRecordsSelectorProxyAuthenticationRequired {
	return &GetAPIV2SelectorsChartDeployRecordsSelectorProxyAuthenticationRequired{}
}

/* GetAPIV2SelectorsChartDeployRecordsSelectorProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPIV2SelectorsChartDeployRecordsSelectorProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2SelectorsChartDeployRecordsSelectorProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/v2/selectors/chart-deploy-records/{selector}][%d] getApiV2SelectorsChartDeployRecordsSelectorProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPIV2SelectorsChartDeployRecordsSelectorProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2SelectorsChartDeployRecordsSelectorProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2SelectorsChartDeployRecordsSelectorConflict creates a GetAPIV2SelectorsChartDeployRecordsSelectorConflict with default headers values
func NewGetAPIV2SelectorsChartDeployRecordsSelectorConflict() *GetAPIV2SelectorsChartDeployRecordsSelectorConflict {
	return &GetAPIV2SelectorsChartDeployRecordsSelectorConflict{}
}

/* GetAPIV2SelectorsChartDeployRecordsSelectorConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPIV2SelectorsChartDeployRecordsSelectorConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2SelectorsChartDeployRecordsSelectorConflict) Error() string {
	return fmt.Sprintf("[GET /api/v2/selectors/chart-deploy-records/{selector}][%d] getApiV2SelectorsChartDeployRecordsSelectorConflict  %+v", 409, o.Payload)
}
func (o *GetAPIV2SelectorsChartDeployRecordsSelectorConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2SelectorsChartDeployRecordsSelectorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2SelectorsChartDeployRecordsSelectorInternalServerError creates a GetAPIV2SelectorsChartDeployRecordsSelectorInternalServerError with default headers values
func NewGetAPIV2SelectorsChartDeployRecordsSelectorInternalServerError() *GetAPIV2SelectorsChartDeployRecordsSelectorInternalServerError {
	return &GetAPIV2SelectorsChartDeployRecordsSelectorInternalServerError{}
}

/* GetAPIV2SelectorsChartDeployRecordsSelectorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPIV2SelectorsChartDeployRecordsSelectorInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2SelectorsChartDeployRecordsSelectorInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/selectors/chart-deploy-records/{selector}][%d] getApiV2SelectorsChartDeployRecordsSelectorInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPIV2SelectorsChartDeployRecordsSelectorInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2SelectorsChartDeployRecordsSelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}