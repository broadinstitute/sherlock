// Code generated by go-swagger; DO NOT EDIT.

package charts

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

// NewDeleteAPIV2ChartsSelectorParams creates a new DeleteAPIV2ChartsSelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAPIV2ChartsSelectorParams() *DeleteAPIV2ChartsSelectorParams {
	return &DeleteAPIV2ChartsSelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPIV2ChartsSelectorParamsWithTimeout creates a new DeleteAPIV2ChartsSelectorParams object
// with the ability to set a timeout on a request.
func NewDeleteAPIV2ChartsSelectorParamsWithTimeout(timeout time.Duration) *DeleteAPIV2ChartsSelectorParams {
	return &DeleteAPIV2ChartsSelectorParams{
		timeout: timeout,
	}
}

// NewDeleteAPIV2ChartsSelectorParamsWithContext creates a new DeleteAPIV2ChartsSelectorParams object
// with the ability to set a context for a request.
func NewDeleteAPIV2ChartsSelectorParamsWithContext(ctx context.Context) *DeleteAPIV2ChartsSelectorParams {
	return &DeleteAPIV2ChartsSelectorParams{
		Context: ctx,
	}
}

// NewDeleteAPIV2ChartsSelectorParamsWithHTTPClient creates a new DeleteAPIV2ChartsSelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAPIV2ChartsSelectorParamsWithHTTPClient(client *http.Client) *DeleteAPIV2ChartsSelectorParams {
	return &DeleteAPIV2ChartsSelectorParams{
		HTTPClient: client,
	}
}

/* DeleteAPIV2ChartsSelectorParams contains all the parameters to send to the API endpoint
   for the delete API v2 charts selector operation.

   Typically these are written to a http.Request.
*/
type DeleteAPIV2ChartsSelectorParams struct {

	/* Selector.

	   The Chart to delete's selector: name or numeric ID
	*/
	Selector string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete API v2 charts selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPIV2ChartsSelectorParams) WithDefaults() *DeleteAPIV2ChartsSelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete API v2 charts selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPIV2ChartsSelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete API v2 charts selector params
func (o *DeleteAPIV2ChartsSelectorParams) WithTimeout(timeout time.Duration) *DeleteAPIV2ChartsSelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API v2 charts selector params
func (o *DeleteAPIV2ChartsSelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API v2 charts selector params
func (o *DeleteAPIV2ChartsSelectorParams) WithContext(ctx context.Context) *DeleteAPIV2ChartsSelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API v2 charts selector params
func (o *DeleteAPIV2ChartsSelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API v2 charts selector params
func (o *DeleteAPIV2ChartsSelectorParams) WithHTTPClient(client *http.Client) *DeleteAPIV2ChartsSelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API v2 charts selector params
func (o *DeleteAPIV2ChartsSelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSelector adds the selector to the delete API v2 charts selector params
func (o *DeleteAPIV2ChartsSelectorParams) WithSelector(selector string) *DeleteAPIV2ChartsSelectorParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the delete API v2 charts selector params
func (o *DeleteAPIV2ChartsSelectorParams) SetSelector(selector string) {
	o.Selector = selector
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPIV2ChartsSelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
