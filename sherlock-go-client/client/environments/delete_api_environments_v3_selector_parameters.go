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
)

// NewDeleteAPIEnvironmentsV3SelectorParams creates a new DeleteAPIEnvironmentsV3SelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAPIEnvironmentsV3SelectorParams() *DeleteAPIEnvironmentsV3SelectorParams {
	return &DeleteAPIEnvironmentsV3SelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPIEnvironmentsV3SelectorParamsWithTimeout creates a new DeleteAPIEnvironmentsV3SelectorParams object
// with the ability to set a timeout on a request.
func NewDeleteAPIEnvironmentsV3SelectorParamsWithTimeout(timeout time.Duration) *DeleteAPIEnvironmentsV3SelectorParams {
	return &DeleteAPIEnvironmentsV3SelectorParams{
		timeout: timeout,
	}
}

// NewDeleteAPIEnvironmentsV3SelectorParamsWithContext creates a new DeleteAPIEnvironmentsV3SelectorParams object
// with the ability to set a context for a request.
func NewDeleteAPIEnvironmentsV3SelectorParamsWithContext(ctx context.Context) *DeleteAPIEnvironmentsV3SelectorParams {
	return &DeleteAPIEnvironmentsV3SelectorParams{
		Context: ctx,
	}
}

// NewDeleteAPIEnvironmentsV3SelectorParamsWithHTTPClient creates a new DeleteAPIEnvironmentsV3SelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAPIEnvironmentsV3SelectorParamsWithHTTPClient(client *http.Client) *DeleteAPIEnvironmentsV3SelectorParams {
	return &DeleteAPIEnvironmentsV3SelectorParams{
		HTTPClient: client,
	}
}

/* DeleteAPIEnvironmentsV3SelectorParams contains all the parameters to send to the API endpoint
   for the delete API environments v3 selector operation.

   Typically these are written to a http.Request.
*/
type DeleteAPIEnvironmentsV3SelectorParams struct {

	/* Selector.

	   The selector of the Environment, which can be either a numeric ID, the name, or 'resource-prefix' + / + the unique resource prefix.
	*/
	Selector string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete API environments v3 selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPIEnvironmentsV3SelectorParams) WithDefaults() *DeleteAPIEnvironmentsV3SelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete API environments v3 selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPIEnvironmentsV3SelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete API environments v3 selector params
func (o *DeleteAPIEnvironmentsV3SelectorParams) WithTimeout(timeout time.Duration) *DeleteAPIEnvironmentsV3SelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API environments v3 selector params
func (o *DeleteAPIEnvironmentsV3SelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API environments v3 selector params
func (o *DeleteAPIEnvironmentsV3SelectorParams) WithContext(ctx context.Context) *DeleteAPIEnvironmentsV3SelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API environments v3 selector params
func (o *DeleteAPIEnvironmentsV3SelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API environments v3 selector params
func (o *DeleteAPIEnvironmentsV3SelectorParams) WithHTTPClient(client *http.Client) *DeleteAPIEnvironmentsV3SelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API environments v3 selector params
func (o *DeleteAPIEnvironmentsV3SelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSelector adds the selector to the delete API environments v3 selector params
func (o *DeleteAPIEnvironmentsV3SelectorParams) WithSelector(selector string) *DeleteAPIEnvironmentsV3SelectorParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the delete API environments v3 selector params
func (o *DeleteAPIEnvironmentsV3SelectorParams) SetSelector(selector string) {
	o.Selector = selector
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPIEnvironmentsV3SelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param selector
	if err := r.SetPathParam("selector", o.Selector); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}