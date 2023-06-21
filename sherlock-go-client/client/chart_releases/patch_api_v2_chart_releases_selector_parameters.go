// Code generated by go-swagger; DO NOT EDIT.

package chart_releases

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

// NewPatchAPIV2ChartReleasesSelectorParams creates a new PatchAPIV2ChartReleasesSelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAPIV2ChartReleasesSelectorParams() *PatchAPIV2ChartReleasesSelectorParams {
	return &PatchAPIV2ChartReleasesSelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPIV2ChartReleasesSelectorParamsWithTimeout creates a new PatchAPIV2ChartReleasesSelectorParams object
// with the ability to set a timeout on a request.
func NewPatchAPIV2ChartReleasesSelectorParamsWithTimeout(timeout time.Duration) *PatchAPIV2ChartReleasesSelectorParams {
	return &PatchAPIV2ChartReleasesSelectorParams{
		timeout: timeout,
	}
}

// NewPatchAPIV2ChartReleasesSelectorParamsWithContext creates a new PatchAPIV2ChartReleasesSelectorParams object
// with the ability to set a context for a request.
func NewPatchAPIV2ChartReleasesSelectorParamsWithContext(ctx context.Context) *PatchAPIV2ChartReleasesSelectorParams {
	return &PatchAPIV2ChartReleasesSelectorParams{
		Context: ctx,
	}
}

// NewPatchAPIV2ChartReleasesSelectorParamsWithHTTPClient creates a new PatchAPIV2ChartReleasesSelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAPIV2ChartReleasesSelectorParamsWithHTTPClient(client *http.Client) *PatchAPIV2ChartReleasesSelectorParams {
	return &PatchAPIV2ChartReleasesSelectorParams{
		HTTPClient: client,
	}
}

/* PatchAPIV2ChartReleasesSelectorParams contains all the parameters to send to the API endpoint
   for the patch API v2 chart releases selector operation.

   Typically these are written to a http.Request.
*/
type PatchAPIV2ChartReleasesSelectorParams struct {

	/* ChartRelease.

	   The edits to make to the ChartRelease
	*/
	ChartRelease *models.V2controllersEditableChartRelease

	/* Selector.

	   The ChartRelease to edit's selector
	*/
	Selector string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch API v2 chart releases selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIV2ChartReleasesSelectorParams) WithDefaults() *PatchAPIV2ChartReleasesSelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch API v2 chart releases selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIV2ChartReleasesSelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch API v2 chart releases selector params
func (o *PatchAPIV2ChartReleasesSelectorParams) WithTimeout(timeout time.Duration) *PatchAPIV2ChartReleasesSelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API v2 chart releases selector params
func (o *PatchAPIV2ChartReleasesSelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API v2 chart releases selector params
func (o *PatchAPIV2ChartReleasesSelectorParams) WithContext(ctx context.Context) *PatchAPIV2ChartReleasesSelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API v2 chart releases selector params
func (o *PatchAPIV2ChartReleasesSelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API v2 chart releases selector params
func (o *PatchAPIV2ChartReleasesSelectorParams) WithHTTPClient(client *http.Client) *PatchAPIV2ChartReleasesSelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API v2 chart releases selector params
func (o *PatchAPIV2ChartReleasesSelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChartRelease adds the chartRelease to the patch API v2 chart releases selector params
func (o *PatchAPIV2ChartReleasesSelectorParams) WithChartRelease(chartRelease *models.V2controllersEditableChartRelease) *PatchAPIV2ChartReleasesSelectorParams {
	o.SetChartRelease(chartRelease)
	return o
}

// SetChartRelease adds the chartRelease to the patch API v2 chart releases selector params
func (o *PatchAPIV2ChartReleasesSelectorParams) SetChartRelease(chartRelease *models.V2controllersEditableChartRelease) {
	o.ChartRelease = chartRelease
}

// WithSelector adds the selector to the patch API v2 chart releases selector params
func (o *PatchAPIV2ChartReleasesSelectorParams) WithSelector(selector string) *PatchAPIV2ChartReleasesSelectorParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the patch API v2 chart releases selector params
func (o *PatchAPIV2ChartReleasesSelectorParams) SetSelector(selector string) {
	o.Selector = selector
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPIV2ChartReleasesSelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ChartRelease != nil {
		if err := r.SetBodyParam(o.ChartRelease); err != nil {
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