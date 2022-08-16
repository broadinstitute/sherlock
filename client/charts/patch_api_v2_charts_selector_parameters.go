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

	"github.com/broadinstitute/sherlock/client/models"
)

// NewPatchAPIV2ChartsSelectorParams creates a new PatchAPIV2ChartsSelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAPIV2ChartsSelectorParams() *PatchAPIV2ChartsSelectorParams {
	return &PatchAPIV2ChartsSelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPIV2ChartsSelectorParamsWithTimeout creates a new PatchAPIV2ChartsSelectorParams object
// with the ability to set a timeout on a request.
func NewPatchAPIV2ChartsSelectorParamsWithTimeout(timeout time.Duration) *PatchAPIV2ChartsSelectorParams {
	return &PatchAPIV2ChartsSelectorParams{
		timeout: timeout,
	}
}

// NewPatchAPIV2ChartsSelectorParamsWithContext creates a new PatchAPIV2ChartsSelectorParams object
// with the ability to set a context for a request.
func NewPatchAPIV2ChartsSelectorParamsWithContext(ctx context.Context) *PatchAPIV2ChartsSelectorParams {
	return &PatchAPIV2ChartsSelectorParams{
		Context: ctx,
	}
}

// NewPatchAPIV2ChartsSelectorParamsWithHTTPClient creates a new PatchAPIV2ChartsSelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAPIV2ChartsSelectorParamsWithHTTPClient(client *http.Client) *PatchAPIV2ChartsSelectorParams {
	return &PatchAPIV2ChartsSelectorParams{
		HTTPClient: client,
	}
}

/* PatchAPIV2ChartsSelectorParams contains all the parameters to send to the API endpoint
   for the patch API v2 charts selector operation.

   Typically these are written to a http.Request.
*/
type PatchAPIV2ChartsSelectorParams struct {

	/* Chart.

	   The edits to make to the Chart
	*/
	Chart *models.V2controllersEditableChart

	/* Selector.

	   The Chart to edit's selector: name or numeric ID
	*/
	Selector string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch API v2 charts selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIV2ChartsSelectorParams) WithDefaults() *PatchAPIV2ChartsSelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch API v2 charts selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIV2ChartsSelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch API v2 charts selector params
func (o *PatchAPIV2ChartsSelectorParams) WithTimeout(timeout time.Duration) *PatchAPIV2ChartsSelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API v2 charts selector params
func (o *PatchAPIV2ChartsSelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API v2 charts selector params
func (o *PatchAPIV2ChartsSelectorParams) WithContext(ctx context.Context) *PatchAPIV2ChartsSelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API v2 charts selector params
func (o *PatchAPIV2ChartsSelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API v2 charts selector params
func (o *PatchAPIV2ChartsSelectorParams) WithHTTPClient(client *http.Client) *PatchAPIV2ChartsSelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API v2 charts selector params
func (o *PatchAPIV2ChartsSelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChart adds the chart to the patch API v2 charts selector params
func (o *PatchAPIV2ChartsSelectorParams) WithChart(chart *models.V2controllersEditableChart) *PatchAPIV2ChartsSelectorParams {
	o.SetChart(chart)
	return o
}

// SetChart adds the chart to the patch API v2 charts selector params
func (o *PatchAPIV2ChartsSelectorParams) SetChart(chart *models.V2controllersEditableChart) {
	o.Chart = chart
}

// WithSelector adds the selector to the patch API v2 charts selector params
func (o *PatchAPIV2ChartsSelectorParams) WithSelector(selector string) *PatchAPIV2ChartsSelectorParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the patch API v2 charts selector params
func (o *PatchAPIV2ChartsSelectorParams) SetSelector(selector string) {
	o.Selector = selector
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPIV2ChartsSelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
