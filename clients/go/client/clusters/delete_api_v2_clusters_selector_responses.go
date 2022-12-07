// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/clients/go/client/models"
)

// DeleteAPIV2ClustersSelectorReader is a Reader for the DeleteAPIV2ClustersSelector structure.
type DeleteAPIV2ClustersSelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIV2ClustersSelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAPIV2ClustersSelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteAPIV2ClustersSelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAPIV2ClustersSelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAPIV2ClustersSelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewDeleteAPIV2ClustersSelectorProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteAPIV2ClustersSelectorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAPIV2ClustersSelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAPIV2ClustersSelectorOK creates a DeleteAPIV2ClustersSelectorOK with default headers values
func NewDeleteAPIV2ClustersSelectorOK() *DeleteAPIV2ClustersSelectorOK {
	return &DeleteAPIV2ClustersSelectorOK{}
}

/*
DeleteAPIV2ClustersSelectorOK describes a response with status code 200, with default header values.

OK
*/
type DeleteAPIV2ClustersSelectorOK struct {
	Payload *models.V2controllersCluster
}

// IsSuccess returns true when this delete Api v2 clusters selector o k response has a 2xx status code
func (o *DeleteAPIV2ClustersSelectorOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete Api v2 clusters selector o k response has a 3xx status code
func (o *DeleteAPIV2ClustersSelectorOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api v2 clusters selector o k response has a 4xx status code
func (o *DeleteAPIV2ClustersSelectorOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Api v2 clusters selector o k response has a 5xx status code
func (o *DeleteAPIV2ClustersSelectorOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api v2 clusters selector o k response a status code equal to that given
func (o *DeleteAPIV2ClustersSelectorOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteAPIV2ClustersSelectorOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/clusters/{selector}][%d] deleteApiV2ClustersSelectorOK  %+v", 200, o.Payload)
}

func (o *DeleteAPIV2ClustersSelectorOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/clusters/{selector}][%d] deleteApiV2ClustersSelectorOK  %+v", 200, o.Payload)
}

func (o *DeleteAPIV2ClustersSelectorOK) GetPayload() *models.V2controllersCluster {
	return o.Payload
}

func (o *DeleteAPIV2ClustersSelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V2controllersCluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIV2ClustersSelectorBadRequest creates a DeleteAPIV2ClustersSelectorBadRequest with default headers values
func NewDeleteAPIV2ClustersSelectorBadRequest() *DeleteAPIV2ClustersSelectorBadRequest {
	return &DeleteAPIV2ClustersSelectorBadRequest{}
}

/*
DeleteAPIV2ClustersSelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteAPIV2ClustersSelectorBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

// IsSuccess returns true when this delete Api v2 clusters selector bad request response has a 2xx status code
func (o *DeleteAPIV2ClustersSelectorBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Api v2 clusters selector bad request response has a 3xx status code
func (o *DeleteAPIV2ClustersSelectorBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api v2 clusters selector bad request response has a 4xx status code
func (o *DeleteAPIV2ClustersSelectorBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Api v2 clusters selector bad request response has a 5xx status code
func (o *DeleteAPIV2ClustersSelectorBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api v2 clusters selector bad request response a status code equal to that given
func (o *DeleteAPIV2ClustersSelectorBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteAPIV2ClustersSelectorBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/clusters/{selector}][%d] deleteApiV2ClustersSelectorBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteAPIV2ClustersSelectorBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/clusters/{selector}][%d] deleteApiV2ClustersSelectorBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteAPIV2ClustersSelectorBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIV2ClustersSelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIV2ClustersSelectorForbidden creates a DeleteAPIV2ClustersSelectorForbidden with default headers values
func NewDeleteAPIV2ClustersSelectorForbidden() *DeleteAPIV2ClustersSelectorForbidden {
	return &DeleteAPIV2ClustersSelectorForbidden{}
}

/*
DeleteAPIV2ClustersSelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteAPIV2ClustersSelectorForbidden struct {
	Payload *models.ErrorsErrorResponse
}

// IsSuccess returns true when this delete Api v2 clusters selector forbidden response has a 2xx status code
func (o *DeleteAPIV2ClustersSelectorForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Api v2 clusters selector forbidden response has a 3xx status code
func (o *DeleteAPIV2ClustersSelectorForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api v2 clusters selector forbidden response has a 4xx status code
func (o *DeleteAPIV2ClustersSelectorForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Api v2 clusters selector forbidden response has a 5xx status code
func (o *DeleteAPIV2ClustersSelectorForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api v2 clusters selector forbidden response a status code equal to that given
func (o *DeleteAPIV2ClustersSelectorForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteAPIV2ClustersSelectorForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/clusters/{selector}][%d] deleteApiV2ClustersSelectorForbidden  %+v", 403, o.Payload)
}

func (o *DeleteAPIV2ClustersSelectorForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/clusters/{selector}][%d] deleteApiV2ClustersSelectorForbidden  %+v", 403, o.Payload)
}

func (o *DeleteAPIV2ClustersSelectorForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIV2ClustersSelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIV2ClustersSelectorNotFound creates a DeleteAPIV2ClustersSelectorNotFound with default headers values
func NewDeleteAPIV2ClustersSelectorNotFound() *DeleteAPIV2ClustersSelectorNotFound {
	return &DeleteAPIV2ClustersSelectorNotFound{}
}

/*
DeleteAPIV2ClustersSelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteAPIV2ClustersSelectorNotFound struct {
	Payload *models.ErrorsErrorResponse
}

// IsSuccess returns true when this delete Api v2 clusters selector not found response has a 2xx status code
func (o *DeleteAPIV2ClustersSelectorNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Api v2 clusters selector not found response has a 3xx status code
func (o *DeleteAPIV2ClustersSelectorNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api v2 clusters selector not found response has a 4xx status code
func (o *DeleteAPIV2ClustersSelectorNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Api v2 clusters selector not found response has a 5xx status code
func (o *DeleteAPIV2ClustersSelectorNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api v2 clusters selector not found response a status code equal to that given
func (o *DeleteAPIV2ClustersSelectorNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteAPIV2ClustersSelectorNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/clusters/{selector}][%d] deleteApiV2ClustersSelectorNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAPIV2ClustersSelectorNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/clusters/{selector}][%d] deleteApiV2ClustersSelectorNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAPIV2ClustersSelectorNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIV2ClustersSelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIV2ClustersSelectorProxyAuthenticationRequired creates a DeleteAPIV2ClustersSelectorProxyAuthenticationRequired with default headers values
func NewDeleteAPIV2ClustersSelectorProxyAuthenticationRequired() *DeleteAPIV2ClustersSelectorProxyAuthenticationRequired {
	return &DeleteAPIV2ClustersSelectorProxyAuthenticationRequired{}
}

/*
DeleteAPIV2ClustersSelectorProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type DeleteAPIV2ClustersSelectorProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

// IsSuccess returns true when this delete Api v2 clusters selector proxy authentication required response has a 2xx status code
func (o *DeleteAPIV2ClustersSelectorProxyAuthenticationRequired) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Api v2 clusters selector proxy authentication required response has a 3xx status code
func (o *DeleteAPIV2ClustersSelectorProxyAuthenticationRequired) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api v2 clusters selector proxy authentication required response has a 4xx status code
func (o *DeleteAPIV2ClustersSelectorProxyAuthenticationRequired) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Api v2 clusters selector proxy authentication required response has a 5xx status code
func (o *DeleteAPIV2ClustersSelectorProxyAuthenticationRequired) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api v2 clusters selector proxy authentication required response a status code equal to that given
func (o *DeleteAPIV2ClustersSelectorProxyAuthenticationRequired) IsCode(code int) bool {
	return code == 407
}

func (o *DeleteAPIV2ClustersSelectorProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/clusters/{selector}][%d] deleteApiV2ClustersSelectorProxyAuthenticationRequired  %+v", 407, o.Payload)
}

func (o *DeleteAPIV2ClustersSelectorProxyAuthenticationRequired) String() string {
	return fmt.Sprintf("[DELETE /api/v2/clusters/{selector}][%d] deleteApiV2ClustersSelectorProxyAuthenticationRequired  %+v", 407, o.Payload)
}

func (o *DeleteAPIV2ClustersSelectorProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIV2ClustersSelectorProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIV2ClustersSelectorConflict creates a DeleteAPIV2ClustersSelectorConflict with default headers values
func NewDeleteAPIV2ClustersSelectorConflict() *DeleteAPIV2ClustersSelectorConflict {
	return &DeleteAPIV2ClustersSelectorConflict{}
}

/*
DeleteAPIV2ClustersSelectorConflict describes a response with status code 409, with default header values.

Conflict
*/
type DeleteAPIV2ClustersSelectorConflict struct {
	Payload *models.ErrorsErrorResponse
}

// IsSuccess returns true when this delete Api v2 clusters selector conflict response has a 2xx status code
func (o *DeleteAPIV2ClustersSelectorConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Api v2 clusters selector conflict response has a 3xx status code
func (o *DeleteAPIV2ClustersSelectorConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api v2 clusters selector conflict response has a 4xx status code
func (o *DeleteAPIV2ClustersSelectorConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Api v2 clusters selector conflict response has a 5xx status code
func (o *DeleteAPIV2ClustersSelectorConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api v2 clusters selector conflict response a status code equal to that given
func (o *DeleteAPIV2ClustersSelectorConflict) IsCode(code int) bool {
	return code == 409
}

func (o *DeleteAPIV2ClustersSelectorConflict) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/clusters/{selector}][%d] deleteApiV2ClustersSelectorConflict  %+v", 409, o.Payload)
}

func (o *DeleteAPIV2ClustersSelectorConflict) String() string {
	return fmt.Sprintf("[DELETE /api/v2/clusters/{selector}][%d] deleteApiV2ClustersSelectorConflict  %+v", 409, o.Payload)
}

func (o *DeleteAPIV2ClustersSelectorConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIV2ClustersSelectorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIV2ClustersSelectorInternalServerError creates a DeleteAPIV2ClustersSelectorInternalServerError with default headers values
func NewDeleteAPIV2ClustersSelectorInternalServerError() *DeleteAPIV2ClustersSelectorInternalServerError {
	return &DeleteAPIV2ClustersSelectorInternalServerError{}
}

/*
DeleteAPIV2ClustersSelectorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteAPIV2ClustersSelectorInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

// IsSuccess returns true when this delete Api v2 clusters selector internal server error response has a 2xx status code
func (o *DeleteAPIV2ClustersSelectorInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Api v2 clusters selector internal server error response has a 3xx status code
func (o *DeleteAPIV2ClustersSelectorInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api v2 clusters selector internal server error response has a 4xx status code
func (o *DeleteAPIV2ClustersSelectorInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Api v2 clusters selector internal server error response has a 5xx status code
func (o *DeleteAPIV2ClustersSelectorInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete Api v2 clusters selector internal server error response a status code equal to that given
func (o *DeleteAPIV2ClustersSelectorInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteAPIV2ClustersSelectorInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/clusters/{selector}][%d] deleteApiV2ClustersSelectorInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteAPIV2ClustersSelectorInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/clusters/{selector}][%d] deleteApiV2ClustersSelectorInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteAPIV2ClustersSelectorInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *DeleteAPIV2ClustersSelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
