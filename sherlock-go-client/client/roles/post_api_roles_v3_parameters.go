// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// NewPostAPIRolesV3Params creates a new PostAPIRolesV3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAPIRolesV3Params() *PostAPIRolesV3Params {
	return &PostAPIRolesV3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIRolesV3ParamsWithTimeout creates a new PostAPIRolesV3Params object
// with the ability to set a timeout on a request.
func NewPostAPIRolesV3ParamsWithTimeout(timeout time.Duration) *PostAPIRolesV3Params {
	return &PostAPIRolesV3Params{
		timeout: timeout,
	}
}

// NewPostAPIRolesV3ParamsWithContext creates a new PostAPIRolesV3Params object
// with the ability to set a context for a request.
func NewPostAPIRolesV3ParamsWithContext(ctx context.Context) *PostAPIRolesV3Params {
	return &PostAPIRolesV3Params{
		Context: ctx,
	}
}

// NewPostAPIRolesV3ParamsWithHTTPClient creates a new PostAPIRolesV3Params object
// with the ability to set a custom HTTPClient for a request.
func NewPostAPIRolesV3ParamsWithHTTPClient(client *http.Client) *PostAPIRolesV3Params {
	return &PostAPIRolesV3Params{
		HTTPClient: client,
	}
}

/* PostAPIRolesV3Params contains all the parameters to send to the API endpoint
   for the post API roles v3 operation.

   Typically these are written to a http.Request.
*/
type PostAPIRolesV3Params struct {

	/* Role.

	   The initial fields the Role should have set
	*/
	Role *models.SherlockRoleV3Edit

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post API roles v3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIRolesV3Params) WithDefaults() *PostAPIRolesV3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post API roles v3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIRolesV3Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post API roles v3 params
func (o *PostAPIRolesV3Params) WithTimeout(timeout time.Duration) *PostAPIRolesV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API roles v3 params
func (o *PostAPIRolesV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API roles v3 params
func (o *PostAPIRolesV3Params) WithContext(ctx context.Context) *PostAPIRolesV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API roles v3 params
func (o *PostAPIRolesV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API roles v3 params
func (o *PostAPIRolesV3Params) WithHTTPClient(client *http.Client) *PostAPIRolesV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API roles v3 params
func (o *PostAPIRolesV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRole adds the role to the post API roles v3 params
func (o *PostAPIRolesV3Params) WithRole(role *models.SherlockRoleV3Edit) *PostAPIRolesV3Params {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the post API roles v3 params
func (o *PostAPIRolesV3Params) SetRole(role *models.SherlockRoleV3Edit) {
	o.Role = role
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIRolesV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Role != nil {
		if err := r.SetBodyParam(o.Role); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}