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

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// NewPutAPIV2ChartsSelectorParams creates a new PutAPIV2ChartsSelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutAPIV2ChartsSelectorParams() *PutAPIV2ChartsSelectorParams {
	return &PutAPIV2ChartsSelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutAPIV2ChartsSelectorParamsWithTimeout creates a new PutAPIV2ChartsSelectorParams object
// with the ability to set a timeout on a request.
func NewPutAPIV2ChartsSelectorParamsWithTimeout(timeout time.Duration) *PutAPIV2ChartsSelectorParams {
	return &PutAPIV2ChartsSelectorParams{
		timeout: timeout,
	}
}

// NewPutAPIV2ChartsSelectorParamsWithContext creates a new PutAPIV2ChartsSelectorParams object
// with the ability to set a context for a request.
func NewPutAPIV2ChartsSelectorParamsWithContext(ctx context.Context) *PutAPIV2ChartsSelectorParams {
	return &PutAPIV2ChartsSelectorParams{
		Context: ctx,
	}
}

// NewPutAPIV2ChartsSelectorParamsWithHTTPClient creates a new PutAPIV2ChartsSelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutAPIV2ChartsSelectorParamsWithHTTPClient(client *http.Client) *PutAPIV2ChartsSelectorParams {
	return &PutAPIV2ChartsSelectorParams{
		HTTPClient: client,
	}
}

/* PutAPIV2ChartsSelectorParams contains all the parameters to send to the API endpoint
   for the put API v2 charts selector operation.

   Typically these are written to a http.Request.
*/
type PutAPIV2ChartsSelectorParams struct {

	/* Chart.

	   The Chart to upsert
	*/
	Chart *models.V2controllersCreatableChart

	/* Selector.

	   The Chart to upsert's selector: name or numeric ID
	*/
	Selector string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put API v2 charts selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutAPIV2ChartsSelectorParams) WithDefaults() *PutAPIV2ChartsSelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put API v2 charts selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutAPIV2ChartsSelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put API v2 charts selector params
func (o *PutAPIV2ChartsSelectorParams) WithTimeout(timeout time.Duration) *PutAPIV2ChartsSelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put API v2 charts selector params
func (o *PutAPIV2ChartsSelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put API v2 charts selector params
func (o *PutAPIV2ChartsSelectorParams) WithContext(ctx context.Context) *PutAPIV2ChartsSelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put API v2 charts selector params
func (o *PutAPIV2ChartsSelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put API v2 charts selector params
func (o *PutAPIV2ChartsSelectorParams) WithHTTPClient(client *http.Client) *PutAPIV2ChartsSelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put API v2 charts selector params
func (o *PutAPIV2ChartsSelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChart adds the chart to the put API v2 charts selector params
func (o *PutAPIV2ChartsSelectorParams) WithChart(chart *models.V2controllersCreatableChart) *PutAPIV2ChartsSelectorParams {
	o.SetChart(chart)
	return o
}

// SetChart adds the chart to the put API v2 charts selector params
func (o *PutAPIV2ChartsSelectorParams) SetChart(chart *models.V2controllersCreatableChart) {
	o.Chart = chart
}

// WithSelector adds the selector to the put API v2 charts selector params
func (o *PutAPIV2ChartsSelectorParams) WithSelector(selector string) *PutAPIV2ChartsSelectorParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the put API v2 charts selector params
func (o *PutAPIV2ChartsSelectorParams) SetSelector(selector string) {
	o.Selector = selector
}

// WriteToRequest writes these params to a swagger request
func (o *PutAPIV2ChartsSelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Chart != nil {
		if err := r.SetBodyParam(o.Chart); err != nil {
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