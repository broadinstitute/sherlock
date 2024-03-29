// Code generated by go-swagger; DO NOT EDIT.

package environments

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

// NewPatchAPIEnvironmentsV3SelectorParams creates a new PatchAPIEnvironmentsV3SelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAPIEnvironmentsV3SelectorParams() *PatchAPIEnvironmentsV3SelectorParams {
	return &PatchAPIEnvironmentsV3SelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPIEnvironmentsV3SelectorParamsWithTimeout creates a new PatchAPIEnvironmentsV3SelectorParams object
// with the ability to set a timeout on a request.
func NewPatchAPIEnvironmentsV3SelectorParamsWithTimeout(timeout time.Duration) *PatchAPIEnvironmentsV3SelectorParams {
	return &PatchAPIEnvironmentsV3SelectorParams{
		timeout: timeout,
	}
}

// NewPatchAPIEnvironmentsV3SelectorParamsWithContext creates a new PatchAPIEnvironmentsV3SelectorParams object
// with the ability to set a context for a request.
func NewPatchAPIEnvironmentsV3SelectorParamsWithContext(ctx context.Context) *PatchAPIEnvironmentsV3SelectorParams {
	return &PatchAPIEnvironmentsV3SelectorParams{
		Context: ctx,
	}
}

// NewPatchAPIEnvironmentsV3SelectorParamsWithHTTPClient creates a new PatchAPIEnvironmentsV3SelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAPIEnvironmentsV3SelectorParamsWithHTTPClient(client *http.Client) *PatchAPIEnvironmentsV3SelectorParams {
	return &PatchAPIEnvironmentsV3SelectorParams{
		HTTPClient: client,
	}
}

/* PatchAPIEnvironmentsV3SelectorParams contains all the parameters to send to the API endpoint
   for the patch API environments v3 selector operation.

   Typically these are written to a http.Request.
*/
type PatchAPIEnvironmentsV3SelectorParams struct {

	/* Environment.

	   The edits to make to the Environment
	*/
	Environment *models.SherlockEnvironmentV3Edit

	/* Selector.

	   The selector of the Environment, which can be either a numeric ID, the name, or 'resource-prefix' + / + the unique resource prefix.
	*/
	Selector string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch API environments v3 selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIEnvironmentsV3SelectorParams) WithDefaults() *PatchAPIEnvironmentsV3SelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch API environments v3 selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIEnvironmentsV3SelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch API environments v3 selector params
func (o *PatchAPIEnvironmentsV3SelectorParams) WithTimeout(timeout time.Duration) *PatchAPIEnvironmentsV3SelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API environments v3 selector params
func (o *PatchAPIEnvironmentsV3SelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API environments v3 selector params
func (o *PatchAPIEnvironmentsV3SelectorParams) WithContext(ctx context.Context) *PatchAPIEnvironmentsV3SelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API environments v3 selector params
func (o *PatchAPIEnvironmentsV3SelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API environments v3 selector params
func (o *PatchAPIEnvironmentsV3SelectorParams) WithHTTPClient(client *http.Client) *PatchAPIEnvironmentsV3SelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API environments v3 selector params
func (o *PatchAPIEnvironmentsV3SelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironment adds the environment to the patch API environments v3 selector params
func (o *PatchAPIEnvironmentsV3SelectorParams) WithEnvironment(environment *models.SherlockEnvironmentV3Edit) *PatchAPIEnvironmentsV3SelectorParams {
	o.SetEnvironment(environment)
	return o
}

// SetEnvironment adds the environment to the patch API environments v3 selector params
func (o *PatchAPIEnvironmentsV3SelectorParams) SetEnvironment(environment *models.SherlockEnvironmentV3Edit) {
	o.Environment = environment
}

// WithSelector adds the selector to the patch API environments v3 selector params
func (o *PatchAPIEnvironmentsV3SelectorParams) WithSelector(selector string) *PatchAPIEnvironmentsV3SelectorParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the patch API environments v3 selector params
func (o *PatchAPIEnvironmentsV3SelectorParams) SetSelector(selector string) {
	o.Selector = selector
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPIEnvironmentsV3SelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Environment != nil {
		if err := r.SetBodyParam(o.Environment); err != nil {
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
