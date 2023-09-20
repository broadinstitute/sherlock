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

// NewGetAPIChartsV3SelectorParams creates a new GetAPIChartsV3SelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIChartsV3SelectorParams() *GetAPIChartsV3SelectorParams {
	return &GetAPIChartsV3SelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIChartsV3SelectorParamsWithTimeout creates a new GetAPIChartsV3SelectorParams object
// with the ability to set a timeout on a request.
func NewGetAPIChartsV3SelectorParamsWithTimeout(timeout time.Duration) *GetAPIChartsV3SelectorParams {
	return &GetAPIChartsV3SelectorParams{
		timeout: timeout,
	}
}

// NewGetAPIChartsV3SelectorParamsWithContext creates a new GetAPIChartsV3SelectorParams object
// with the ability to set a context for a request.
func NewGetAPIChartsV3SelectorParamsWithContext(ctx context.Context) *GetAPIChartsV3SelectorParams {
	return &GetAPIChartsV3SelectorParams{
		Context: ctx,
	}
}

// NewGetAPIChartsV3SelectorParamsWithHTTPClient creates a new GetAPIChartsV3SelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIChartsV3SelectorParamsWithHTTPClient(client *http.Client) *GetAPIChartsV3SelectorParams {
	return &GetAPIChartsV3SelectorParams{
		HTTPClient: client,
	}
}

/* GetAPIChartsV3SelectorParams contains all the parameters to send to the API endpoint
   for the get API charts v3 selector operation.

   Typically these are written to a http.Request.
*/
type GetAPIChartsV3SelectorParams struct {

	/* Selector.

	   The selector of the Chart, which can be either a numeric ID or the name.
	*/
	Selector string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API charts v3 selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIChartsV3SelectorParams) WithDefaults() *GetAPIChartsV3SelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API charts v3 selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIChartsV3SelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API charts v3 selector params
func (o *GetAPIChartsV3SelectorParams) WithTimeout(timeout time.Duration) *GetAPIChartsV3SelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API charts v3 selector params
func (o *GetAPIChartsV3SelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API charts v3 selector params
func (o *GetAPIChartsV3SelectorParams) WithContext(ctx context.Context) *GetAPIChartsV3SelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API charts v3 selector params
func (o *GetAPIChartsV3SelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API charts v3 selector params
func (o *GetAPIChartsV3SelectorParams) WithHTTPClient(client *http.Client) *GetAPIChartsV3SelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API charts v3 selector params
func (o *GetAPIChartsV3SelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSelector adds the selector to the get API charts v3 selector params
func (o *GetAPIChartsV3SelectorParams) WithSelector(selector string) *GetAPIChartsV3SelectorParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the get API charts v3 selector params
func (o *GetAPIChartsV3SelectorParams) SetSelector(selector string) {
	o.Selector = selector
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIChartsV3SelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
