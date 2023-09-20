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

// NewPatchAPIChartsV3SelectorParams creates a new PatchAPIChartsV3SelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAPIChartsV3SelectorParams() *PatchAPIChartsV3SelectorParams {
	return &PatchAPIChartsV3SelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPIChartsV3SelectorParamsWithTimeout creates a new PatchAPIChartsV3SelectorParams object
// with the ability to set a timeout on a request.
func NewPatchAPIChartsV3SelectorParamsWithTimeout(timeout time.Duration) *PatchAPIChartsV3SelectorParams {
	return &PatchAPIChartsV3SelectorParams{
		timeout: timeout,
	}
}

// NewPatchAPIChartsV3SelectorParamsWithContext creates a new PatchAPIChartsV3SelectorParams object
// with the ability to set a context for a request.
func NewPatchAPIChartsV3SelectorParamsWithContext(ctx context.Context) *PatchAPIChartsV3SelectorParams {
	return &PatchAPIChartsV3SelectorParams{
		Context: ctx,
	}
}

// NewPatchAPIChartsV3SelectorParamsWithHTTPClient creates a new PatchAPIChartsV3SelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAPIChartsV3SelectorParamsWithHTTPClient(client *http.Client) *PatchAPIChartsV3SelectorParams {
	return &PatchAPIChartsV3SelectorParams{
		HTTPClient: client,
	}
}

/* PatchAPIChartsV3SelectorParams contains all the parameters to send to the API endpoint
   for the patch API charts v3 selector operation.

   Typically these are written to a http.Request.
*/
type PatchAPIChartsV3SelectorParams struct {

	/* Chart.

	   The edits to make to the Chart
	*/
	Chart *models.SherlockChartV3Edit

	/* Selector.

	   The selector of the Chart, which can be either a numeric ID or the name.
	*/
	Selector string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch API charts v3 selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIChartsV3SelectorParams) WithDefaults() *PatchAPIChartsV3SelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch API charts v3 selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIChartsV3SelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch API charts v3 selector params
func (o *PatchAPIChartsV3SelectorParams) WithTimeout(timeout time.Duration) *PatchAPIChartsV3SelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API charts v3 selector params
func (o *PatchAPIChartsV3SelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API charts v3 selector params
func (o *PatchAPIChartsV3SelectorParams) WithContext(ctx context.Context) *PatchAPIChartsV3SelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API charts v3 selector params
func (o *PatchAPIChartsV3SelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API charts v3 selector params
func (o *PatchAPIChartsV3SelectorParams) WithHTTPClient(client *http.Client) *PatchAPIChartsV3SelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API charts v3 selector params
func (o *PatchAPIChartsV3SelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChart adds the chart to the patch API charts v3 selector params
func (o *PatchAPIChartsV3SelectorParams) WithChart(chart *models.SherlockChartV3Edit) *PatchAPIChartsV3SelectorParams {
	o.SetChart(chart)
	return o
}

// SetChart adds the chart to the patch API charts v3 selector params
func (o *PatchAPIChartsV3SelectorParams) SetChart(chart *models.SherlockChartV3Edit) {
	o.Chart = chart
}

// WithSelector adds the selector to the patch API charts v3 selector params
func (o *PatchAPIChartsV3SelectorParams) WithSelector(selector string) *PatchAPIChartsV3SelectorParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the patch API charts v3 selector params
func (o *PatchAPIChartsV3SelectorParams) SetSelector(selector string) {
	o.Selector = selector
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPIChartsV3SelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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