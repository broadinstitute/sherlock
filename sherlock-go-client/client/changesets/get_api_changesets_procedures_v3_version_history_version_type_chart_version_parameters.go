// Code generated by go-swagger; DO NOT EDIT.

package changesets

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

// NewGetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams creates a new GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams() *GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams {
	return &GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParamsWithTimeout creates a new GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams object
// with the ability to set a timeout on a request.
func NewGetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParamsWithTimeout(timeout time.Duration) *GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams {
	return &GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams{
		timeout: timeout,
	}
}

// NewGetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParamsWithContext creates a new GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams object
// with the ability to set a context for a request.
func NewGetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParamsWithContext(ctx context.Context) *GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams {
	return &GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams{
		Context: ctx,
	}
}

// NewGetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParamsWithHTTPClient creates a new GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParamsWithHTTPClient(client *http.Client) *GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams {
	return &GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams{
		HTTPClient: client,
	}
}

/* GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams contains all the parameters to send to the API endpoint
   for the get API changesets procedures v3 version history version type chart version operation.

   Typically these are written to a http.Request.
*/
type GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams struct {

	/* Chart.

	   The chart the version belongs to
	*/
	Chart string

	/* Version.

	   The version to look for
	*/
	Version string

	/* VersionType.

	   The type of the version, either 'app' or 'chart'
	*/
	VersionType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API changesets procedures v3 version history version type chart version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams) WithDefaults() *GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API changesets procedures v3 version history version type chart version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API changesets procedures v3 version history version type chart version params
func (o *GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams) WithTimeout(timeout time.Duration) *GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API changesets procedures v3 version history version type chart version params
func (o *GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API changesets procedures v3 version history version type chart version params
func (o *GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams) WithContext(ctx context.Context) *GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API changesets procedures v3 version history version type chart version params
func (o *GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API changesets procedures v3 version history version type chart version params
func (o *GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams) WithHTTPClient(client *http.Client) *GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API changesets procedures v3 version history version type chart version params
func (o *GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChart adds the chart to the get API changesets procedures v3 version history version type chart version params
func (o *GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams) WithChart(chart string) *GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams {
	o.SetChart(chart)
	return o
}

// SetChart adds the chart to the get API changesets procedures v3 version history version type chart version params
func (o *GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams) SetChart(chart string) {
	o.Chart = chart
}

// WithVersion adds the version to the get API changesets procedures v3 version history version type chart version params
func (o *GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams) WithVersion(version string) *GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get API changesets procedures v3 version history version type chart version params
func (o *GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams) SetVersion(version string) {
	o.Version = version
}

// WithVersionType adds the versionType to the get API changesets procedures v3 version history version type chart version params
func (o *GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams) WithVersionType(versionType string) *GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams {
	o.SetVersionType(versionType)
	return o
}

// SetVersionType adds the versionType to the get API changesets procedures v3 version history version type chart version params
func (o *GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams) SetVersionType(versionType string) {
	o.VersionType = versionType
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIChangesetsProceduresV3VersionHistoryVersionTypeChartVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param chart
	if err := r.SetPathParam("chart", o.Chart); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	// path param version-type
	if err := r.SetPathParam("version-type", o.VersionType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
