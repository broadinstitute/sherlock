// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/client/models"
)

// PatchAPIV2ClustersSelectorReader is a Reader for the PatchAPIV2ClustersSelector structure.
type PatchAPIV2ClustersSelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPIV2ClustersSelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAPIV2ClustersSelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchAPIV2ClustersSelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchAPIV2ClustersSelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchAPIV2ClustersSelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewPatchAPIV2ClustersSelectorProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchAPIV2ClustersSelectorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchAPIV2ClustersSelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchAPIV2ClustersSelectorOK creates a PatchAPIV2ClustersSelectorOK with default headers values
func NewPatchAPIV2ClustersSelectorOK() *PatchAPIV2ClustersSelectorOK {
	return &PatchAPIV2ClustersSelectorOK{}
}

/* PatchAPIV2ClustersSelectorOK describes a response with status code 200, with default header values.

OK
*/
type PatchAPIV2ClustersSelectorOK struct {
	Payload *models.V2controllersCluster
}

func (o *PatchAPIV2ClustersSelectorOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/clusters/{selector}][%d] patchApiV2ClustersSelectorOK  %+v", 200, o.Payload)
}
func (o *PatchAPIV2ClustersSelectorOK) GetPayload() *models.V2controllersCluster {
	return o.Payload
}

func (o *PatchAPIV2ClustersSelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V2controllersCluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIV2ClustersSelectorBadRequest creates a PatchAPIV2ClustersSelectorBadRequest with default headers values
func NewPatchAPIV2ClustersSelectorBadRequest() *PatchAPIV2ClustersSelectorBadRequest {
	return &PatchAPIV2ClustersSelectorBadRequest{}
}

/* PatchAPIV2ClustersSelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PatchAPIV2ClustersSelectorBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIV2ClustersSelectorBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/clusters/{selector}][%d] patchApiV2ClustersSelectorBadRequest  %+v", 400, o.Payload)
}
func (o *PatchAPIV2ClustersSelectorBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIV2ClustersSelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIV2ClustersSelectorForbidden creates a PatchAPIV2ClustersSelectorForbidden with default headers values
func NewPatchAPIV2ClustersSelectorForbidden() *PatchAPIV2ClustersSelectorForbidden {
	return &PatchAPIV2ClustersSelectorForbidden{}
}

/* PatchAPIV2ClustersSelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PatchAPIV2ClustersSelectorForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIV2ClustersSelectorForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/clusters/{selector}][%d] patchApiV2ClustersSelectorForbidden  %+v", 403, o.Payload)
}
func (o *PatchAPIV2ClustersSelectorForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIV2ClustersSelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIV2ClustersSelectorNotFound creates a PatchAPIV2ClustersSelectorNotFound with default headers values
func NewPatchAPIV2ClustersSelectorNotFound() *PatchAPIV2ClustersSelectorNotFound {
	return &PatchAPIV2ClustersSelectorNotFound{}
}

/* PatchAPIV2ClustersSelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PatchAPIV2ClustersSelectorNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIV2ClustersSelectorNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/clusters/{selector}][%d] patchApiV2ClustersSelectorNotFound  %+v", 404, o.Payload)
}
func (o *PatchAPIV2ClustersSelectorNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIV2ClustersSelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIV2ClustersSelectorProxyAuthenticationRequired creates a PatchAPIV2ClustersSelectorProxyAuthenticationRequired with default headers values
func NewPatchAPIV2ClustersSelectorProxyAuthenticationRequired() *PatchAPIV2ClustersSelectorProxyAuthenticationRequired {
	return &PatchAPIV2ClustersSelectorProxyAuthenticationRequired{}
}

/* PatchAPIV2ClustersSelectorProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type PatchAPIV2ClustersSelectorProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIV2ClustersSelectorProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/clusters/{selector}][%d] patchApiV2ClustersSelectorProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *PatchAPIV2ClustersSelectorProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIV2ClustersSelectorProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIV2ClustersSelectorConflict creates a PatchAPIV2ClustersSelectorConflict with default headers values
func NewPatchAPIV2ClustersSelectorConflict() *PatchAPIV2ClustersSelectorConflict {
	return &PatchAPIV2ClustersSelectorConflict{}
}

/* PatchAPIV2ClustersSelectorConflict describes a response with status code 409, with default header values.

Conflict
*/
type PatchAPIV2ClustersSelectorConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIV2ClustersSelectorConflict) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/clusters/{selector}][%d] patchApiV2ClustersSelectorConflict  %+v", 409, o.Payload)
}
func (o *PatchAPIV2ClustersSelectorConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIV2ClustersSelectorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIV2ClustersSelectorInternalServerError creates a PatchAPIV2ClustersSelectorInternalServerError with default headers values
func NewPatchAPIV2ClustersSelectorInternalServerError() *PatchAPIV2ClustersSelectorInternalServerError {
	return &PatchAPIV2ClustersSelectorInternalServerError{}
}

/* PatchAPIV2ClustersSelectorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PatchAPIV2ClustersSelectorInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PatchAPIV2ClustersSelectorInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/clusters/{selector}][%d] patchApiV2ClustersSelectorInternalServerError  %+v", 500, o.Payload)
}
func (o *PatchAPIV2ClustersSelectorInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PatchAPIV2ClustersSelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
