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

// NewPostAPIV2ChartReleasesParams creates a new PostAPIV2ChartReleasesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAPIV2ChartReleasesParams() *PostAPIV2ChartReleasesParams {
	return &PostAPIV2ChartReleasesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV2ChartReleasesParamsWithTimeout creates a new PostAPIV2ChartReleasesParams object
// with the ability to set a timeout on a request.
func NewPostAPIV2ChartReleasesParamsWithTimeout(timeout time.Duration) *PostAPIV2ChartReleasesParams {
	return &PostAPIV2ChartReleasesParams{
		timeout: timeout,
	}
}

// NewPostAPIV2ChartReleasesParamsWithContext creates a new PostAPIV2ChartReleasesParams object
// with the ability to set a context for a request.
func NewPostAPIV2ChartReleasesParamsWithContext(ctx context.Context) *PostAPIV2ChartReleasesParams {
	return &PostAPIV2ChartReleasesParams{
		Context: ctx,
	}
}

// NewPostAPIV2ChartReleasesParamsWithHTTPClient creates a new PostAPIV2ChartReleasesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAPIV2ChartReleasesParamsWithHTTPClient(client *http.Client) *PostAPIV2ChartReleasesParams {
	return &PostAPIV2ChartReleasesParams{
		HTTPClient: client,
	}
}

/* PostAPIV2ChartReleasesParams contains all the parameters to send to the API endpoint
   for the post API v2 chart releases operation.

   Typically these are written to a http.Request.
*/
type PostAPIV2ChartReleasesParams struct {

	/* ChartRelease.

	   The ChartRelease to create
	*/
	ChartRelease *models.V2controllersCreatableChartRelease

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post API v2 chart releases params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIV2ChartReleasesParams) WithDefaults() *PostAPIV2ChartReleasesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post API v2 chart releases params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIV2ChartReleasesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post API v2 chart releases params
func (o *PostAPIV2ChartReleasesParams) WithTimeout(timeout time.Duration) *PostAPIV2ChartReleasesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v2 chart releases params
func (o *PostAPIV2ChartReleasesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v2 chart releases params
func (o *PostAPIV2ChartReleasesParams) WithContext(ctx context.Context) *PostAPIV2ChartReleasesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v2 chart releases params
func (o *PostAPIV2ChartReleasesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v2 chart releases params
func (o *PostAPIV2ChartReleasesParams) WithHTTPClient(client *http.Client) *PostAPIV2ChartReleasesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v2 chart releases params
func (o *PostAPIV2ChartReleasesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChartRelease adds the chartRelease to the post API v2 chart releases params
func (o *PostAPIV2ChartReleasesParams) WithChartRelease(chartRelease *models.V2controllersCreatableChartRelease) *PostAPIV2ChartReleasesParams {
	o.SetChartRelease(chartRelease)
	return o
}

// SetChartRelease adds the chartRelease to the post API v2 chart releases params
func (o *PostAPIV2ChartReleasesParams) SetChartRelease(chartRelease *models.V2controllersCreatableChartRelease) {
	o.ChartRelease = chartRelease
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV2ChartReleasesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ChartRelease != nil {
		if err := r.SetBodyParam(o.ChartRelease); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
