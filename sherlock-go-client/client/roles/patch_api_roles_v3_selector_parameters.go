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

// NewPatchAPIRolesV3SelectorParams creates a new PatchAPIRolesV3SelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAPIRolesV3SelectorParams() *PatchAPIRolesV3SelectorParams {
	return &PatchAPIRolesV3SelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPIRolesV3SelectorParamsWithTimeout creates a new PatchAPIRolesV3SelectorParams object
// with the ability to set a timeout on a request.
func NewPatchAPIRolesV3SelectorParamsWithTimeout(timeout time.Duration) *PatchAPIRolesV3SelectorParams {
	return &PatchAPIRolesV3SelectorParams{
		timeout: timeout,
	}
}

// NewPatchAPIRolesV3SelectorParamsWithContext creates a new PatchAPIRolesV3SelectorParams object
// with the ability to set a context for a request.
func NewPatchAPIRolesV3SelectorParamsWithContext(ctx context.Context) *PatchAPIRolesV3SelectorParams {
	return &PatchAPIRolesV3SelectorParams{
		Context: ctx,
	}
}

// NewPatchAPIRolesV3SelectorParamsWithHTTPClient creates a new PatchAPIRolesV3SelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAPIRolesV3SelectorParamsWithHTTPClient(client *http.Client) *PatchAPIRolesV3SelectorParams {
	return &PatchAPIRolesV3SelectorParams{
		HTTPClient: client,
	}
}

/* PatchAPIRolesV3SelectorParams contains all the parameters to send to the API endpoint
   for the patch API roles v3 selector operation.

   Typically these are written to a http.Request.
*/
type PatchAPIRolesV3SelectorParams struct {

	/* Role.

	   The edits to make to the Role
	*/
	Role *models.SherlockRoleV3Edit

	/* Selector.

	   The selector of the Role, which can be either the numeric ID or the name
	*/
	Selector string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch API roles v3 selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIRolesV3SelectorParams) WithDefaults() *PatchAPIRolesV3SelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch API roles v3 selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIRolesV3SelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch API roles v3 selector params
func (o *PatchAPIRolesV3SelectorParams) WithTimeout(timeout time.Duration) *PatchAPIRolesV3SelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API roles v3 selector params
func (o *PatchAPIRolesV3SelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API roles v3 selector params
func (o *PatchAPIRolesV3SelectorParams) WithContext(ctx context.Context) *PatchAPIRolesV3SelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API roles v3 selector params
func (o *PatchAPIRolesV3SelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API roles v3 selector params
func (o *PatchAPIRolesV3SelectorParams) WithHTTPClient(client *http.Client) *PatchAPIRolesV3SelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API roles v3 selector params
func (o *PatchAPIRolesV3SelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRole adds the role to the patch API roles v3 selector params
func (o *PatchAPIRolesV3SelectorParams) WithRole(role *models.SherlockRoleV3Edit) *PatchAPIRolesV3SelectorParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the patch API roles v3 selector params
func (o *PatchAPIRolesV3SelectorParams) SetRole(role *models.SherlockRoleV3Edit) {
	o.Role = role
}

// WithSelector adds the selector to the patch API roles v3 selector params
func (o *PatchAPIRolesV3SelectorParams) WithSelector(selector string) *PatchAPIRolesV3SelectorParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the patch API roles v3 selector params
func (o *PatchAPIRolesV3SelectorParams) SetSelector(selector string) {
	o.Selector = selector
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPIRolesV3SelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Role != nil {
		if err := r.SetBodyParam(o.Role); err != nil {
			return err
		}
	}

	// path param selector
	if err := r.SetPathParam("selector", o.Selector); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}