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

// PostAPIChangesetsProceduresV3PlanAndApplyReader is a Reader for the PostAPIChangesetsProceduresV3PlanAndApply structure.
type PostAPIChangesetsProceduresV3PlanAndApplyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIChangesetsProceduresV3PlanAndApplyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAPIChangesetsProceduresV3PlanAndApplyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostAPIChangesetsProceduresV3PlanAndApplyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAPIChangesetsProceduresV3PlanAndApplyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAPIChangesetsProceduresV3PlanAndApplyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAPIChangesetsProceduresV3PlanAndApplyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 407:
		result := NewPostAPIChangesetsProceduresV3PlanAndApplyProxyAuthenticationRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostAPIChangesetsProceduresV3PlanAndApplyConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAPIChangesetsProceduresV3PlanAndApplyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAPIChangesetsProceduresV3PlanAndApplyOK creates a PostAPIChangesetsProceduresV3PlanAndApplyOK with default headers values
func NewPostAPIChangesetsProceduresV3PlanAndApplyOK() *PostAPIChangesetsProceduresV3PlanAndApplyOK {
	return &PostAPIChangesetsProceduresV3PlanAndApplyOK{}
}

/* PostAPIChangesetsProceduresV3PlanAndApplyOK describes a response with status code 200, with default header values.

OK
*/
type PostAPIChangesetsProceduresV3PlanAndApplyOK struct {
	Payload []*models.SherlockChangesetV3
}

func (o *PostAPIChangesetsProceduresV3PlanAndApplyOK) Error() string {
	return fmt.Sprintf("[POST /api/changesets/procedures/v3/plan-and-apply][%d] postApiChangesetsProceduresV3PlanAndApplyOK  %+v", 200, o.Payload)
}
func (o *PostAPIChangesetsProceduresV3PlanAndApplyOK) GetPayload() []*models.SherlockChangesetV3 {
	return o.Payload
}

func (o *PostAPIChangesetsProceduresV3PlanAndApplyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIChangesetsProceduresV3PlanAndApplyCreated creates a PostAPIChangesetsProceduresV3PlanAndApplyCreated with default headers values
func NewPostAPIChangesetsProceduresV3PlanAndApplyCreated() *PostAPIChangesetsProceduresV3PlanAndApplyCreated {
	return &PostAPIChangesetsProceduresV3PlanAndApplyCreated{}
}

/* PostAPIChangesetsProceduresV3PlanAndApplyCreated describes a response with status code 201, with default header values.

Created
*/
type PostAPIChangesetsProceduresV3PlanAndApplyCreated struct {
	Payload []*models.SherlockChangesetV3
}

func (o *PostAPIChangesetsProceduresV3PlanAndApplyCreated) Error() string {
	return fmt.Sprintf("[POST /api/changesets/procedures/v3/plan-and-apply][%d] postApiChangesetsProceduresV3PlanAndApplyCreated  %+v", 201, o.Payload)
}
func (o *PostAPIChangesetsProceduresV3PlanAndApplyCreated) GetPayload() []*models.SherlockChangesetV3 {
	return o.Payload
}

func (o *PostAPIChangesetsProceduresV3PlanAndApplyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIChangesetsProceduresV3PlanAndApplyBadRequest creates a PostAPIChangesetsProceduresV3PlanAndApplyBadRequest with default headers values
func NewPostAPIChangesetsProceduresV3PlanAndApplyBadRequest() *PostAPIChangesetsProceduresV3PlanAndApplyBadRequest {
	return &PostAPIChangesetsProceduresV3PlanAndApplyBadRequest{}
}

/* PostAPIChangesetsProceduresV3PlanAndApplyBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostAPIChangesetsProceduresV3PlanAndApplyBadRequest struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIChangesetsProceduresV3PlanAndApplyBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/changesets/procedures/v3/plan-and-apply][%d] postApiChangesetsProceduresV3PlanAndApplyBadRequest  %+v", 400, o.Payload)
}
func (o *PostAPIChangesetsProceduresV3PlanAndApplyBadRequest) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIChangesetsProceduresV3PlanAndApplyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIChangesetsProceduresV3PlanAndApplyForbidden creates a PostAPIChangesetsProceduresV3PlanAndApplyForbidden with default headers values
func NewPostAPIChangesetsProceduresV3PlanAndApplyForbidden() *PostAPIChangesetsProceduresV3PlanAndApplyForbidden {
	return &PostAPIChangesetsProceduresV3PlanAndApplyForbidden{}
}

/* PostAPIChangesetsProceduresV3PlanAndApplyForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostAPIChangesetsProceduresV3PlanAndApplyForbidden struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIChangesetsProceduresV3PlanAndApplyForbidden) Error() string {
	return fmt.Sprintf("[POST /api/changesets/procedures/v3/plan-and-apply][%d] postApiChangesetsProceduresV3PlanAndApplyForbidden  %+v", 403, o.Payload)
}
func (o *PostAPIChangesetsProceduresV3PlanAndApplyForbidden) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIChangesetsProceduresV3PlanAndApplyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIChangesetsProceduresV3PlanAndApplyNotFound creates a PostAPIChangesetsProceduresV3PlanAndApplyNotFound with default headers values
func NewPostAPIChangesetsProceduresV3PlanAndApplyNotFound() *PostAPIChangesetsProceduresV3PlanAndApplyNotFound {
	return &PostAPIChangesetsProceduresV3PlanAndApplyNotFound{}
}

/* PostAPIChangesetsProceduresV3PlanAndApplyNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostAPIChangesetsProceduresV3PlanAndApplyNotFound struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIChangesetsProceduresV3PlanAndApplyNotFound) Error() string {
	return fmt.Sprintf("[POST /api/changesets/procedures/v3/plan-and-apply][%d] postApiChangesetsProceduresV3PlanAndApplyNotFound  %+v", 404, o.Payload)
}
func (o *PostAPIChangesetsProceduresV3PlanAndApplyNotFound) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIChangesetsProceduresV3PlanAndApplyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIChangesetsProceduresV3PlanAndApplyProxyAuthenticationRequired creates a PostAPIChangesetsProceduresV3PlanAndApplyProxyAuthenticationRequired with default headers values
func NewPostAPIChangesetsProceduresV3PlanAndApplyProxyAuthenticationRequired() *PostAPIChangesetsProceduresV3PlanAndApplyProxyAuthenticationRequired {
	return &PostAPIChangesetsProceduresV3PlanAndApplyProxyAuthenticationRequired{}
}

/* PostAPIChangesetsProceduresV3PlanAndApplyProxyAuthenticationRequired describes a response with status code 407, with default header values.

Proxy Authentication Required
*/
type PostAPIChangesetsProceduresV3PlanAndApplyProxyAuthenticationRequired struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIChangesetsProceduresV3PlanAndApplyProxyAuthenticationRequired) Error() string {
	return fmt.Sprintf("[POST /api/changesets/procedures/v3/plan-and-apply][%d] postApiChangesetsProceduresV3PlanAndApplyProxyAuthenticationRequired  %+v", 407, o.Payload)
}
func (o *PostAPIChangesetsProceduresV3PlanAndApplyProxyAuthenticationRequired) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIChangesetsProceduresV3PlanAndApplyProxyAuthenticationRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIChangesetsProceduresV3PlanAndApplyConflict creates a PostAPIChangesetsProceduresV3PlanAndApplyConflict with default headers values
func NewPostAPIChangesetsProceduresV3PlanAndApplyConflict() *PostAPIChangesetsProceduresV3PlanAndApplyConflict {
	return &PostAPIChangesetsProceduresV3PlanAndApplyConflict{}
}

/* PostAPIChangesetsProceduresV3PlanAndApplyConflict describes a response with status code 409, with default header values.

Conflict
*/
type PostAPIChangesetsProceduresV3PlanAndApplyConflict struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIChangesetsProceduresV3PlanAndApplyConflict) Error() string {
	return fmt.Sprintf("[POST /api/changesets/procedures/v3/plan-and-apply][%d] postApiChangesetsProceduresV3PlanAndApplyConflict  %+v", 409, o.Payload)
}
func (o *PostAPIChangesetsProceduresV3PlanAndApplyConflict) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIChangesetsProceduresV3PlanAndApplyConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIChangesetsProceduresV3PlanAndApplyInternalServerError creates a PostAPIChangesetsProceduresV3PlanAndApplyInternalServerError with default headers values
func NewPostAPIChangesetsProceduresV3PlanAndApplyInternalServerError() *PostAPIChangesetsProceduresV3PlanAndApplyInternalServerError {
	return &PostAPIChangesetsProceduresV3PlanAndApplyInternalServerError{}
}

/* PostAPIChangesetsProceduresV3PlanAndApplyInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostAPIChangesetsProceduresV3PlanAndApplyInternalServerError struct {
	Payload *models.ErrorsErrorResponse
}

func (o *PostAPIChangesetsProceduresV3PlanAndApplyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/changesets/procedures/v3/plan-and-apply][%d] postApiChangesetsProceduresV3PlanAndApplyInternalServerError  %+v", 500, o.Payload)
}
func (o *PostAPIChangesetsProceduresV3PlanAndApplyInternalServerError) GetPayload() *models.ErrorsErrorResponse {
	return o.Payload
}

func (o *PostAPIChangesetsProceduresV3PlanAndApplyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}