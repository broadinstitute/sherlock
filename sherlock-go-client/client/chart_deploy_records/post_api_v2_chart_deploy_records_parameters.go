// Code generated by go-swagger; DO NOT EDIT.

package chart_deploy_records

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

// NewPostAPIV2ChartDeployRecordsParams creates a new PostAPIV2ChartDeployRecordsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAPIV2ChartDeployRecordsParams() *PostAPIV2ChartDeployRecordsParams {
	return &PostAPIV2ChartDeployRecordsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV2ChartDeployRecordsParamsWithTimeout creates a new PostAPIV2ChartDeployRecordsParams object
// with the ability to set a timeout on a request.
func NewPostAPIV2ChartDeployRecordsParamsWithTimeout(timeout time.Duration) *PostAPIV2ChartDeployRecordsParams {
	return &PostAPIV2ChartDeployRecordsParams{
		timeout: timeout,
	}
}

// NewPostAPIV2ChartDeployRecordsParamsWithContext creates a new PostAPIV2ChartDeployRecordsParams object
// with the ability to set a context for a request.
func NewPostAPIV2ChartDeployRecordsParamsWithContext(ctx context.Context) *PostAPIV2ChartDeployRecordsParams {
	return &PostAPIV2ChartDeployRecordsParams{
		Context: ctx,
	}
}

// NewPostAPIV2ChartDeployRecordsParamsWithHTTPClient creates a new PostAPIV2ChartDeployRecordsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAPIV2ChartDeployRecordsParamsWithHTTPClient(client *http.Client) *PostAPIV2ChartDeployRecordsParams {
	return &PostAPIV2ChartDeployRecordsParams{
		HTTPClient: client,
	}
}

/* PostAPIV2ChartDeployRecordsParams contains all the parameters to send to the API endpoint
   for the post API v2 chart deploy records operation.

   Typically these are written to a http.Request.
*/
type PostAPIV2ChartDeployRecordsParams struct {

	/* ChartDeployRecord.

	   The ChartDeployRecord to create
	*/
	ChartDeployRecord *models.V2controllersCreatableChartDeployRecord

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post API v2 chart deploy records params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIV2ChartDeployRecordsParams) WithDefaults() *PostAPIV2ChartDeployRecordsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post API v2 chart deploy records params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIV2ChartDeployRecordsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post API v2 chart deploy records params
func (o *PostAPIV2ChartDeployRecordsParams) WithTimeout(timeout time.Duration) *PostAPIV2ChartDeployRecordsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v2 chart deploy records params
func (o *PostAPIV2ChartDeployRecordsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v2 chart deploy records params
func (o *PostAPIV2ChartDeployRecordsParams) WithContext(ctx context.Context) *PostAPIV2ChartDeployRecordsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v2 chart deploy records params
func (o *PostAPIV2ChartDeployRecordsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v2 chart deploy records params
func (o *PostAPIV2ChartDeployRecordsParams) WithHTTPClient(client *http.Client) *PostAPIV2ChartDeployRecordsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v2 chart deploy records params
func (o *PostAPIV2ChartDeployRecordsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChartDeployRecord adds the chartDeployRecord to the post API v2 chart deploy records params
func (o *PostAPIV2ChartDeployRecordsParams) WithChartDeployRecord(chartDeployRecord *models.V2controllersCreatableChartDeployRecord) *PostAPIV2ChartDeployRecordsParams {
	o.SetChartDeployRecord(chartDeployRecord)
	return o
}

// SetChartDeployRecord adds the chartDeployRecord to the post API v2 chart deploy records params
func (o *PostAPIV2ChartDeployRecordsParams) SetChartDeployRecord(chartDeployRecord *models.V2controllersCreatableChartDeployRecord) {
	o.ChartDeployRecord = chartDeployRecord
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV2ChartDeployRecordsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ChartDeployRecord != nil {
		if err := r.SetBodyParam(o.ChartDeployRecord); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
