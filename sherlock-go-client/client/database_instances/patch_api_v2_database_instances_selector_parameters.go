// Code generated by go-swagger; DO NOT EDIT.

package database_instances

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

// NewPatchAPIV2DatabaseInstancesSelectorParams creates a new PatchAPIV2DatabaseInstancesSelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAPIV2DatabaseInstancesSelectorParams() *PatchAPIV2DatabaseInstancesSelectorParams {
	return &PatchAPIV2DatabaseInstancesSelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPIV2DatabaseInstancesSelectorParamsWithTimeout creates a new PatchAPIV2DatabaseInstancesSelectorParams object
// with the ability to set a timeout on a request.
func NewPatchAPIV2DatabaseInstancesSelectorParamsWithTimeout(timeout time.Duration) *PatchAPIV2DatabaseInstancesSelectorParams {
	return &PatchAPIV2DatabaseInstancesSelectorParams{
		timeout: timeout,
	}
}

// NewPatchAPIV2DatabaseInstancesSelectorParamsWithContext creates a new PatchAPIV2DatabaseInstancesSelectorParams object
// with the ability to set a context for a request.
func NewPatchAPIV2DatabaseInstancesSelectorParamsWithContext(ctx context.Context) *PatchAPIV2DatabaseInstancesSelectorParams {
	return &PatchAPIV2DatabaseInstancesSelectorParams{
		Context: ctx,
	}
}

// NewPatchAPIV2DatabaseInstancesSelectorParamsWithHTTPClient creates a new PatchAPIV2DatabaseInstancesSelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAPIV2DatabaseInstancesSelectorParamsWithHTTPClient(client *http.Client) *PatchAPIV2DatabaseInstancesSelectorParams {
	return &PatchAPIV2DatabaseInstancesSelectorParams{
		HTTPClient: client,
	}
}

/* PatchAPIV2DatabaseInstancesSelectorParams contains all the parameters to send to the API endpoint
   for the patch API v2 database instances selector operation.

   Typically these are written to a http.Request.
*/
type PatchAPIV2DatabaseInstancesSelectorParams struct {

	/* Chart.

	   The edits to make to the DatabaseInstance
	*/
	Chart *models.V2controllersEditableDatabaseInstance

	/* Selector.

	   The DatabaseInstance to edit's selector: numeric ID or 'chart-release/' followed by a chart release selector
	*/
	Selector string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch API v2 database instances selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIV2DatabaseInstancesSelectorParams) WithDefaults() *PatchAPIV2DatabaseInstancesSelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch API v2 database instances selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIV2DatabaseInstancesSelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch API v2 database instances selector params
func (o *PatchAPIV2DatabaseInstancesSelectorParams) WithTimeout(timeout time.Duration) *PatchAPIV2DatabaseInstancesSelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API v2 database instances selector params
func (o *PatchAPIV2DatabaseInstancesSelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API v2 database instances selector params
func (o *PatchAPIV2DatabaseInstancesSelectorParams) WithContext(ctx context.Context) *PatchAPIV2DatabaseInstancesSelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API v2 database instances selector params
func (o *PatchAPIV2DatabaseInstancesSelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API v2 database instances selector params
func (o *PatchAPIV2DatabaseInstancesSelectorParams) WithHTTPClient(client *http.Client) *PatchAPIV2DatabaseInstancesSelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API v2 database instances selector params
func (o *PatchAPIV2DatabaseInstancesSelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChart adds the chart to the patch API v2 database instances selector params
func (o *PatchAPIV2DatabaseInstancesSelectorParams) WithChart(chart *models.V2controllersEditableDatabaseInstance) *PatchAPIV2DatabaseInstancesSelectorParams {
	o.SetChart(chart)
	return o
}

// SetChart adds the chart to the patch API v2 database instances selector params
func (o *PatchAPIV2DatabaseInstancesSelectorParams) SetChart(chart *models.V2controllersEditableDatabaseInstance) {
	o.Chart = chart
}

// WithSelector adds the selector to the patch API v2 database instances selector params
func (o *PatchAPIV2DatabaseInstancesSelectorParams) WithSelector(selector string) *PatchAPIV2DatabaseInstancesSelectorParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the patch API v2 database instances selector params
func (o *PatchAPIV2DatabaseInstancesSelectorParams) SetSelector(selector string) {
	o.Selector = selector
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPIV2DatabaseInstancesSelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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