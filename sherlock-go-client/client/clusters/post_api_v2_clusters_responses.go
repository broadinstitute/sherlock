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

// PostAPIV2ClustersReader is a Reader for the PostAPIV2Clusters structure.
type PostAPIV2ClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIV2ClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAPIV2ClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostAPIV2ClustersCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAPIV2ClustersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAPIV2ClustersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAPIV2ClustersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewPostAPIV2ClustersProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostAPIV2ClustersConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAPIV2ClustersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAPIV2ClustersOK creates a PostAPIV2ClustersOK with default headers values
func NewPostAPIV2ClustersOK() *PostAPIV2ClustersOK {
	return &PostAPIV2ClustersOK{}
}

/* PostAPIV2ClustersOK describes a response with status code 200, with default header values.

OK
*/
type PostAPIV2ClustersOK struct {
	Payload *models.V2controllersCluster
}

func (o *PostAPIV2ClustersOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/clusters][%d] postApiV2ClustersOK  %+v", 200, o.Payload)
}
func (o *PostAPIV2ClustersOK) GetPayload() *models.V2controllersCluster {
	return o.Payload
}

func (o *PostAPIV2ClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V2controllersCluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2ClustersCreated creates a PostAPIV2ClustersCreated with default headers values
func NewPostAPIV2ClustersCreated() *PostAPIV2ClustersCreated {
	return &PostAPIV2ClustersCreated{}
}

/* PostAPIV2ClustersCreated describes a response with status code 201, with default header values.

Created
*/
type PostAPIV2ClustersCreated struct {
	Payload *models.V2controllersCluster
}

func (o *PostAPIV2ClustersCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/clusters][%d] postApiV2ClustersCreated  %+v", 201, o.Payload)
}
func (o *PostAPIV2ClustersCreated) GetPayload() *models.V2controllersCluster {
	return o.Payload
}

func (o *PostAPIV2ClustersCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V2controllersCluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2ClustersBadRequest creates a PostAPIV2ClustersBadRequest with default headers values
func NewPostAPIV2ClustersBadRequest() *PostAPIV2ClustersBadRequest {
	return &PostAPIV2ClustersBadRequest{}
}

/* PostAPIV2ClustersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostAPIV2ClustersBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIV2ClustersBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/clusters][%d] postApiV2ClustersBadRequest  %+v", 400, o.Payload)
}
func (o *PostAPIV2ClustersBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2ClustersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2ClustersForbidden creates a PostAPIV2ClustersForbidden with default headers values
func NewPostAPIV2ClustersForbidden() *PostAPIV2ClustersForbidden {
	return &PostAPIV2ClustersForbidden{}
}

/* PostAPIV2ClustersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostAPIV2ClustersForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIV2ClustersForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/clusters][%d] postApiV2ClustersForbidden  %+v", 403, o.Payload)
}
func (o *PostAPIV2ClustersForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2ClustersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2ClustersNotFound creates a PostAPIV2ClustersNotFound with default headers values
func NewPostAPIV2ClustersNotFound() *PostAPIV2ClustersNotFound {
	return &PostAPIV2ClustersNotFound{}
}

/* PostAPIV2ClustersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostAPIV2ClustersNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIV2ClustersNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/clusters][%d] postApiV2ClustersNotFound  %+v", 404, o.Payload)
}
func (o *PostAPIV2ClustersNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2ClustersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2ClustersProxyAuthenticationRequired creates a PostAPIV2ClustersProxyAuthenticationRequired with default headers values
func NewPostAPIV2ClustersProxyAuthenticationRequired() *PostAPIV2ClustersProxyAuthenticationRequired {
	return &PostAPIV2ClustersProxyAuthenticationRequired{}
}

/* PostAPIV2ClustersProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type PostAPIV2ClustersProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIV2ClustersProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[POST /api/v2/clusters][%d] postApiV2ClustersProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *PostAPIV2ClustersProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2ClustersProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2ClustersConflict creates a PostAPIV2ClustersConflict with default headers values
func NewPostAPIV2ClustersConflict() *PostAPIV2ClustersConflict {
	return &PostAPIV2ClustersConflict{}
}

/* PostAPIV2ClustersConflict describes a response with status code 409, with default header values.

Conflict
*/
type PostAPIV2ClustersConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIV2ClustersConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/clusters][%d] postApiV2ClustersConflict  %+v", 409, o.Payload)
}
func (o *PostAPIV2ClustersConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2ClustersConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV2ClustersInternalServerError creates a PostAPIV2ClustersInternalServerError with default headers values
func NewPostAPIV2ClustersInternalServerError() *PostAPIV2ClustersInternalServerError {
	return &PostAPIV2ClustersInternalServerError{}
}

/* PostAPIV2ClustersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostAPIV2ClustersInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIV2ClustersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/clusters][%d] postApiV2ClustersInternalServerError  %+v", 500, o.Payload)
}
func (o *PostAPIV2ClustersInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIV2ClustersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}