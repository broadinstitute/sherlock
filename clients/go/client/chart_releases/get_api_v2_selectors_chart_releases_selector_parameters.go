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
)

// NewGetAPIV2SelectorsChartReleasesSelectorParams creates a new GetAPIV2SelectorsChartReleasesSelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIV2SelectorsChartReleasesSelectorParams() *GetAPIV2SelectorsChartReleasesSelectorParams {
	return &GetAPIV2SelectorsChartReleasesSelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV2SelectorsChartReleasesSelectorParamsWithTimeout creates a new GetAPIV2SelectorsChartReleasesSelectorParams object
// with the ability to set a timeout on a request.
func NewGetAPIV2SelectorsChartReleasesSelectorParamsWithTimeout(timeout time.Duration) *GetAPIV2SelectorsChartReleasesSelectorParams {
	return &GetAPIV2SelectorsChartReleasesSelectorParams{
		timeout: timeout,
	}
}

// NewGetAPIV2SelectorsChartReleasesSelectorParamsWithContext creates a new GetAPIV2SelectorsChartReleasesSelectorParams object
// with the ability to set a context for a request.
func NewGetAPIV2SelectorsChartReleasesSelectorParamsWithContext(ctx context.Context) *GetAPIV2SelectorsChartReleasesSelectorParams {
	return &GetAPIV2SelectorsChartReleasesSelectorParams{
		Context: ctx,
	}
}

// NewGetAPIV2SelectorsChartReleasesSelectorParamsWithHTTPClient creates a new GetAPIV2SelectorsChartReleasesSelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIV2SelectorsChartReleasesSelectorParamsWithHTTPClient(client *http.Client) *GetAPIV2SelectorsChartReleasesSelectorParams {
	return &GetAPIV2SelectorsChartReleasesSelectorParams{
		HTTPClient: client,
	}
}

/*
GetAPIV2SelectorsChartReleasesSelectorParams contains all the parameters to send to the API endpoint

	for the get API v2 selectors chart releases selector operation.

	Typically these are written to a http.Request.
*/
type GetAPIV2SelectorsChartReleasesSelectorParams struct {

	/* Selector.

	   The selector of the ChartRelease to list other selectors for
	*/
	Selector string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API v2 selectors chart releases selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2SelectorsChartReleasesSelectorParams) WithDefaults() *GetAPIV2SelectorsChartReleasesSelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API v2 selectors chart releases selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2SelectorsChartReleasesSelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API v2 selectors chart releases selector params
func (o *GetAPIV2SelectorsChartReleasesSelectorParams) WithTimeout(timeout time.Duration) *GetAPIV2SelectorsChartReleasesSelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v2 selectors chart releases selector params
func (o *GetAPIV2SelectorsChartReleasesSelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v2 selectors chart releases selector params
func (o *GetAPIV2SelectorsChartReleasesSelectorParams) WithContext(ctx context.Context) *GetAPIV2SelectorsChartReleasesSelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v2 selectors chart releases selector params
func (o *GetAPIV2SelectorsChartReleasesSelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v2 selectors chart releases selector params
func (o *GetAPIV2SelectorsChartReleasesSelectorParams) WithHTTPClient(client *http.Client) *GetAPIV2SelectorsChartReleasesSelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v2 selectors chart releases selector params
func (o *GetAPIV2SelectorsChartReleasesSelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSelector adds the selector to the get API v2 selectors chart releases selector params
func (o *GetAPIV2SelectorsChartReleasesSelectorParams) WithSelector(selector string) *GetAPIV2SelectorsChartReleasesSelectorParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the get API v2 selectors chart releases selector params
func (o *GetAPIV2SelectorsChartReleasesSelectorParams) SetSelector(selector string) {
	o.Selector = selector
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV2SelectorsChartReleasesSelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
