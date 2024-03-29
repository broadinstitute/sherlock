// Code generated by go-swagger; DO NOT EDIT.

package chart_versions

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

// NewPatchAPIChartVersionsV3SelectorParams creates a new PatchAPIChartVersionsV3SelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAPIChartVersionsV3SelectorParams() *PatchAPIChartVersionsV3SelectorParams {
	return &PatchAPIChartVersionsV3SelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPIChartVersionsV3SelectorParamsWithTimeout creates a new PatchAPIChartVersionsV3SelectorParams object
// with the ability to set a timeout on a request.
func NewPatchAPIChartVersionsV3SelectorParamsWithTimeout(timeout time.Duration) *PatchAPIChartVersionsV3SelectorParams {
	return &PatchAPIChartVersionsV3SelectorParams{
		timeout: timeout,
	}
}

// NewPatchAPIChartVersionsV3SelectorParamsWithContext creates a new PatchAPIChartVersionsV3SelectorParams object
// with the ability to set a context for a request.
func NewPatchAPIChartVersionsV3SelectorParamsWithContext(ctx context.Context) *PatchAPIChartVersionsV3SelectorParams {
	return &PatchAPIChartVersionsV3SelectorParams{
		Context: ctx,
	}
}

// NewPatchAPIChartVersionsV3SelectorParamsWithHTTPClient creates a new PatchAPIChartVersionsV3SelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAPIChartVersionsV3SelectorParamsWithHTTPClient(client *http.Client) *PatchAPIChartVersionsV3SelectorParams {
	return &PatchAPIChartVersionsV3SelectorParams{
		HTTPClient: client,
	}
}

/* PatchAPIChartVersionsV3SelectorParams contains all the parameters to send to the API endpoint
   for the patch API chart versions v3 selector operation.

   Typically these are written to a http.Request.
*/
type PatchAPIChartVersionsV3SelectorParams struct {

	/* ChartVersion.

	   The edits to make to the ChartVersion
	*/
	ChartVersion *models.SherlockChartVersionV3Edit

	/* Selector.

	   The selector of the ChartVersion, which can be either a numeric ID or chart/version.
	*/
	Selector string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch API chart versions v3 selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIChartVersionsV3SelectorParams) WithDefaults() *PatchAPIChartVersionsV3SelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch API chart versions v3 selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIChartVersionsV3SelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch API chart versions v3 selector params
func (o *PatchAPIChartVersionsV3SelectorParams) WithTimeout(timeout time.Duration) *PatchAPIChartVersionsV3SelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API chart versions v3 selector params
func (o *PatchAPIChartVersionsV3SelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API chart versions v3 selector params
func (o *PatchAPIChartVersionsV3SelectorParams) WithContext(ctx context.Context) *PatchAPIChartVersionsV3SelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API chart versions v3 selector params
func (o *PatchAPIChartVersionsV3SelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API chart versions v3 selector params
func (o *PatchAPIChartVersionsV3SelectorParams) WithHTTPClient(client *http.Client) *PatchAPIChartVersionsV3SelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API chart versions v3 selector params
func (o *PatchAPIChartVersionsV3SelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChartVersion adds the chartVersion to the patch API chart versions v3 selector params
func (o *PatchAPIChartVersionsV3SelectorParams) WithChartVersion(chartVersion *models.SherlockChartVersionV3Edit) *PatchAPIChartVersionsV3SelectorParams {
	o.SetChartVersion(chartVersion)
	return o
}

// SetChartVersion adds the chartVersion to the patch API chart versions v3 selector params
func (o *PatchAPIChartVersionsV3SelectorParams) SetChartVersion(chartVersion *models.SherlockChartVersionV3Edit) {
	o.ChartVersion = chartVersion
}

// WithSelector adds the selector to the patch API chart versions v3 selector params
func (o *PatchAPIChartVersionsV3SelectorParams) WithSelector(selector string) *PatchAPIChartVersionsV3SelectorParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the patch API chart versions v3 selector params
func (o *PatchAPIChartVersionsV3SelectorParams) SetSelector(selector string) {
	o.Selector = selector
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPIChartVersionsV3SelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ChartVersion != nil {
		if err := r.SetBodyParam(o.ChartVersion); err != nil {
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
