// Code generated by go-swagger; DO NOT EDIT.

package ci_identifiers

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

// NewGetAPICiIdentifiersV3SelectorParams creates a new GetAPICiIdentifiersV3SelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPICiIdentifiersV3SelectorParams() *GetAPICiIdentifiersV3SelectorParams {
	return &GetAPICiIdentifiersV3SelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPICiIdentifiersV3SelectorParamsWithTimeout creates a new GetAPICiIdentifiersV3SelectorParams object
// with the ability to set a timeout on a request.
func NewGetAPICiIdentifiersV3SelectorParamsWithTimeout(timeout time.Duration) *GetAPICiIdentifiersV3SelectorParams {
	return &GetAPICiIdentifiersV3SelectorParams{
		timeout: timeout,
	}
}

// NewGetAPICiIdentifiersV3SelectorParamsWithContext creates a new GetAPICiIdentifiersV3SelectorParams object
// with the ability to set a context for a request.
func NewGetAPICiIdentifiersV3SelectorParamsWithContext(ctx context.Context) *GetAPICiIdentifiersV3SelectorParams {
	return &GetAPICiIdentifiersV3SelectorParams{
		Context: ctx,
	}
}

// NewGetAPICiIdentifiersV3SelectorParamsWithHTTPClient creates a new GetAPICiIdentifiersV3SelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPICiIdentifiersV3SelectorParamsWithHTTPClient(client *http.Client) *GetAPICiIdentifiersV3SelectorParams {
	return &GetAPICiIdentifiersV3SelectorParams{
		HTTPClient: client,
	}
}

/* GetAPICiIdentifiersV3SelectorParams contains all the parameters to send to the API endpoint
   for the get API ci identifiers v3 selector operation.

   Typically these are written to a http.Request.
*/
type GetAPICiIdentifiersV3SelectorParams struct {

	/* LimitCiRuns.

	   Control how many CiRuns are returned (default 10)
	*/
	LimitCiRuns *int64

	/* OffsetCiRuns.

	   Control the offset for the returned CiRuns (default 0)
	*/
	OffsetCiRuns *int64

	/* Selector.

	   The selector of CiIdentifier, which can be referenced either by numeric ID or indirectly by '{type}/{selector...}'
	*/
	Selector string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API ci identifiers v3 selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPICiIdentifiersV3SelectorParams) WithDefaults() *GetAPICiIdentifiersV3SelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API ci identifiers v3 selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPICiIdentifiersV3SelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API ci identifiers v3 selector params
func (o *GetAPICiIdentifiersV3SelectorParams) WithTimeout(timeout time.Duration) *GetAPICiIdentifiersV3SelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API ci identifiers v3 selector params
func (o *GetAPICiIdentifiersV3SelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API ci identifiers v3 selector params
func (o *GetAPICiIdentifiersV3SelectorParams) WithContext(ctx context.Context) *GetAPICiIdentifiersV3SelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API ci identifiers v3 selector params
func (o *GetAPICiIdentifiersV3SelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API ci identifiers v3 selector params
func (o *GetAPICiIdentifiersV3SelectorParams) WithHTTPClient(client *http.Client) *GetAPICiIdentifiersV3SelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API ci identifiers v3 selector params
func (o *GetAPICiIdentifiersV3SelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimitCiRuns adds the limitCiRuns to the get API ci identifiers v3 selector params
func (o *GetAPICiIdentifiersV3SelectorParams) WithLimitCiRuns(limitCiRuns *int64) *GetAPICiIdentifiersV3SelectorParams {
	o.SetLimitCiRuns(limitCiRuns)
	return o
}

// SetLimitCiRuns adds the limitCiRuns to the get API ci identifiers v3 selector params
func (o *GetAPICiIdentifiersV3SelectorParams) SetLimitCiRuns(limitCiRuns *int64) {
	o.LimitCiRuns = limitCiRuns
}

// WithOffsetCiRuns adds the offsetCiRuns to the get API ci identifiers v3 selector params
func (o *GetAPICiIdentifiersV3SelectorParams) WithOffsetCiRuns(offsetCiRuns *int64) *GetAPICiIdentifiersV3SelectorParams {
	o.SetOffsetCiRuns(offsetCiRuns)
	return o
}

// SetOffsetCiRuns adds the offsetCiRuns to the get API ci identifiers v3 selector params
func (o *GetAPICiIdentifiersV3SelectorParams) SetOffsetCiRuns(offsetCiRuns *int64) {
	o.OffsetCiRuns = offsetCiRuns
}

// WithSelector adds the selector to the get API ci identifiers v3 selector params
func (o *GetAPICiIdentifiersV3SelectorParams) WithSelector(selector string) *GetAPICiIdentifiersV3SelectorParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the get API ci identifiers v3 selector params
func (o *GetAPICiIdentifiersV3SelectorParams) SetSelector(selector string) {
	o.Selector = selector
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPICiIdentifiersV3SelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.LimitCiRuns != nil {

		// query param limitCiRuns
		var qrLimitCiRuns int64

		if o.LimitCiRuns != nil {
			qrLimitCiRuns = *o.LimitCiRuns
		}
		qLimitCiRuns := swag.FormatInt64(qrLimitCiRuns)
		if qLimitCiRuns != "" {

			if err := r.SetQueryParam("limitCiRuns", qLimitCiRuns); err != nil {
				return err
			}
		}
	}

	if o.OffsetCiRuns != nil {

		// query param offsetCiRuns
		var qrOffsetCiRuns int64

		if o.OffsetCiRuns != nil {
			qrOffsetCiRuns = *o.OffsetCiRuns
		}
		qOffsetCiRuns := swag.FormatInt64(qrOffsetCiRuns)
		if qOffsetCiRuns != "" {

			if err := r.SetQueryParam("offsetCiRuns", qOffsetCiRuns); err != nil {
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
