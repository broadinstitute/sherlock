// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// GetAPIV2ClustersReader is a Reader for the GetAPIV2Clusters structure.
type GetAPIV2ClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV2ClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV2ClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIV2ClustersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIV2ClustersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIV2ClustersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewGetAPIV2ClustersProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPIV2ClustersConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIV2ClustersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIV2ClustersOK creates a GetAPIV2ClustersOK with default headers values
func NewGetAPIV2ClustersOK() *GetAPIV2ClustersOK {
	return &GetAPIV2ClustersOK{}
}

/* GetAPIV2ClustersOK describes a response with status code 200, with default header values.

OK
*/
type GetAPIV2ClustersOK struct {
	Payload []*models.V2controllersCluster
}

func (o *GetAPIV2ClustersOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/clusters][%d] getApiV2ClustersOK  %+v", 200, o.Payload)
}
func (o *GetAPIV2ClustersOK) GetPayload() []*models.V2controllersCluster {
	return o.Payload
}

func (o *GetAPIV2ClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ClustersBadRequest creates a GetAPIV2ClustersBadRequest with default headers values
func NewGetAPIV2ClustersBadRequest() *GetAPIV2ClustersBadRequest {
	return &GetAPIV2ClustersBadRequest{}
}

/* GetAPIV2ClustersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIV2ClustersBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ClustersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/clusters][%d] getApiV2ClustersBadRequest  %+v", 400, o.Payload)
}
func (o *GetAPIV2ClustersBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ClustersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ClustersForbidden creates a GetAPIV2ClustersForbidden with default headers values
func NewGetAPIV2ClustersForbidden() *GetAPIV2ClustersForbidden {
	return &GetAPIV2ClustersForbidden{}
}

/* GetAPIV2ClustersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPIV2ClustersForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ClustersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/clusters][%d] getApiV2ClustersForbidden  %+v", 403, o.Payload)
}
func (o *GetAPIV2ClustersForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ClustersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ClustersNotFound creates a GetAPIV2ClustersNotFound with default headers values
func NewGetAPIV2ClustersNotFound() *GetAPIV2ClustersNotFound {
	return &GetAPIV2ClustersNotFound{}
}

/* GetAPIV2ClustersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIV2ClustersNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ClustersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/clusters][%d] getApiV2ClustersNotFound  %+v", 404, o.Payload)
}
func (o *GetAPIV2ClustersNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ClustersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ClustersProxyAuthenticationRequired creates a GetAPIV2ClustersProxyAuthenticationRequired with default headers values
func NewGetAPIV2ClustersProxyAuthenticationRequired() *GetAPIV2ClustersProxyAuthenticationRequired {
	return &GetAPIV2ClustersProxyAuthenticationRequired{}
}

/* GetAPIV2ClustersProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type GetAPIV2ClustersProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ClustersProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[GET /api/v2/clusters][%d] getApiV2ClustersProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *GetAPIV2ClustersProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ClustersProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ClustersConflict creates a GetAPIV2ClustersConflict with default headers values
func NewGetAPIV2ClustersConflict() *GetAPIV2ClustersConflict {
	return &GetAPIV2ClustersConflict{}
}

/* GetAPIV2ClustersConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPIV2ClustersConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ClustersConflict) Error() string {
	return fmt.Sprintf("[GET /api/v2/clusters][%d] getApiV2ClustersConflict  %+v", 409, o.Payload)
}
func (o *GetAPIV2ClustersConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ClustersConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV2ClustersInternalServerError creates a GetAPIV2ClustersInternalServerError with default headers values
func NewGetAPIV2ClustersInternalServerError() *GetAPIV2ClustersInternalServerError {
	return &GetAPIV2ClustersInternalServerError{}
}

/* GetAPIV2ClustersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPIV2ClustersInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *GetAPIV2ClustersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/clusters][%d] getApiV2ClustersInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPIV2ClustersInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *GetAPIV2ClustersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}