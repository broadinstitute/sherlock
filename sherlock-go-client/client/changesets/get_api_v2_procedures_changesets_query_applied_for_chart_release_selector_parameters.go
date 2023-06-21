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
	"github.com/go-openapi/swag"
)

// NewGetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams creates a new GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams() *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams {
	return &GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParamsWithTimeout creates a new GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams object
// with the ability to set a timeout on a request.
func NewGetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParamsWithTimeout(timeout time.Duration) *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams {
	return &GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams{
		timeout: timeout,
	}
}

// NewGetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParamsWithContext creates a new GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams object
// with the ability to set a context for a request.
func NewGetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParamsWithContext(ctx context.Context) *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams {
	return &GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams{
		Context: ctx,
	}
}

// NewGetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParamsWithHTTPClient creates a new GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParamsWithHTTPClient(client *http.Client) *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams {
	return &GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams{
		HTTPClient: client,
	}
}

/* GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams contains all the parameters to send to the API endpoint
   for the get API v2 procedures changesets query applied for chart release selector operation.

   Typically these are written to a http.Request.
*/
type GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams struct {

	/* Limit.

	   An optional limit to the number of entries returned
	*/
	Limit *int64

	/* Offset.

	   An optional offset to skip a number of latest Changesets
	*/
	Offset *int64

	/* Selector.

	   Selector the Chart Release to find applied Changesets for
	*/
	Selector string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API v2 procedures changesets query applied for chart release selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams) WithDefaults() *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API v2 procedures changesets query applied for chart release selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API v2 procedures changesets query applied for chart release selector params
func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams) WithTimeout(timeout time.Duration) *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v2 procedures changesets query applied for chart release selector params
func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v2 procedures changesets query applied for chart release selector params
func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams) WithContext(ctx context.Context) *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v2 procedures changesets query applied for chart release selector params
func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v2 procedures changesets query applied for chart release selector params
func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams) WithHTTPClient(client *http.Client) *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v2 procedures changesets query applied for chart release selector params
func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get API v2 procedures changesets query applied for chart release selector params
func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams) WithLimit(limit *int64) *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get API v2 procedures changesets query applied for chart release selector params
func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get API v2 procedures changesets query applied for chart release selector params
func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams) WithOffset(offset *int64) *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get API v2 procedures changesets query applied for chart release selector params
func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSelector adds the selector to the get API v2 procedures changesets query applied for chart release selector params
func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams) WithSelector(selector string) *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the get API v2 procedures changesets query applied for chart release selector params
func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams) SetSelector(selector string) {
	o.Selector = selector
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
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