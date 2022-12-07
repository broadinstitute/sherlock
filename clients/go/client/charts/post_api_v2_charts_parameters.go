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

	"github.com/broadinstitute/sherlock/clients/go/client/models"
)

// NewPostAPIV2ChartsParams creates a new PostAPIV2ChartsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAPIV2ChartsParams() *PostAPIV2ChartsParams {
	return &PostAPIV2ChartsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV2ChartsParamsWithTimeout creates a new PostAPIV2ChartsParams object
// with the ability to set a timeout on a request.
func NewPostAPIV2ChartsParamsWithTimeout(timeout time.Duration) *PostAPIV2ChartsParams {
	return &PostAPIV2ChartsParams{
		timeout: timeout,
	}
}

// NewPostAPIV2ChartsParamsWithContext creates a new PostAPIV2ChartsParams object
// with the ability to set a context for a request.
func NewPostAPIV2ChartsParamsWithContext(ctx context.Context) *PostAPIV2ChartsParams {
	return &PostAPIV2ChartsParams{
		Context: ctx,
	}
}

// NewPostAPIV2ChartsParamsWithHTTPClient creates a new PostAPIV2ChartsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAPIV2ChartsParamsWithHTTPClient(client *http.Client) *PostAPIV2ChartsParams {
	return &PostAPIV2ChartsParams{
		HTTPClient: client,
	}
}

/*
PostAPIV2ChartsParams contains all the parameters to send to the API endpoint

	for the post API v2 charts operation.

	Typically these are written to a http.Request.
*/
type PostAPIV2ChartsParams struct {

	/* Chart.

	   The Chart to create
	*/
	Chart *models.V2controllersCreatableChart

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post API v2 charts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIV2ChartsParams) WithDefaults() *PostAPIV2ChartsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post API v2 charts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIV2ChartsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post API v2 charts params
func (o *PostAPIV2ChartsParams) WithTimeout(timeout time.Duration) *PostAPIV2ChartsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v2 charts params
func (o *PostAPIV2ChartsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v2 charts params
func (o *PostAPIV2ChartsParams) WithContext(ctx context.Context) *PostAPIV2ChartsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v2 charts params
func (o *PostAPIV2ChartsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v2 charts params
func (o *PostAPIV2ChartsParams) WithHTTPClient(client *http.Client) *PostAPIV2ChartsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v2 charts params
func (o *PostAPIV2ChartsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChart adds the chart to the post API v2 charts params
func (o *PostAPIV2ChartsParams) WithChart(chart *models.V2controllersCreatableChart) *PostAPIV2ChartsParams {
	o.SetChart(chart)
	return o
}

// SetChart adds the chart to the post API v2 charts params
func (o *PostAPIV2ChartsParams) SetChart(chart *models.V2controllersCreatableChart) {
	o.Chart = chart
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV2ChartsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Chart != nil {
		if err := r.SetBodyParam(o.Chart); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
