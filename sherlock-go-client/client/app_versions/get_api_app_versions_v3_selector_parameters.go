// Code generated by go-swagger; DO NOT EDIT.

package app_versions

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

// NewGetAPIAppVersionsV3SelectorParams creates a new GetAPIAppVersionsV3SelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIAppVersionsV3SelectorParams() *GetAPIAppVersionsV3SelectorParams {
	return &GetAPIAppVersionsV3SelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIAppVersionsV3SelectorParamsWithTimeout creates a new GetAPIAppVersionsV3SelectorParams object
// with the ability to set a timeout on a request.
func NewGetAPIAppVersionsV3SelectorParamsWithTimeout(timeout time.Duration) *GetAPIAppVersionsV3SelectorParams {
	return &GetAPIAppVersionsV3SelectorParams{
		timeout: timeout,
	}
}

// NewGetAPIAppVersionsV3SelectorParamsWithContext creates a new GetAPIAppVersionsV3SelectorParams object
// with the ability to set a context for a request.
func NewGetAPIAppVersionsV3SelectorParamsWithContext(ctx context.Context) *GetAPIAppVersionsV3SelectorParams {
	return &GetAPIAppVersionsV3SelectorParams{
		Context: ctx,
	}
}

// NewGetAPIAppVersionsV3SelectorParamsWithHTTPClient creates a new GetAPIAppVersionsV3SelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIAppVersionsV3SelectorParamsWithHTTPClient(client *http.Client) *GetAPIAppVersionsV3SelectorParams {
	return &GetAPIAppVersionsV3SelectorParams{
		HTTPClient: client,
	}
}

/* GetAPIAppVersionsV3SelectorParams contains all the parameters to send to the API endpoint
   for the get API app versions v3 selector operation.

   Typically these are written to a http.Request.
*/
type GetAPIAppVersionsV3SelectorParams struct {

	/* Selector.

	   The selector of the AppVersion, which can be either a numeric ID or chart/version.
	*/
	Selector string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API app versions v3 selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIAppVersionsV3SelectorParams) WithDefaults() *GetAPIAppVersionsV3SelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API app versions v3 selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIAppVersionsV3SelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API app versions v3 selector params
func (o *GetAPIAppVersionsV3SelectorParams) WithTimeout(timeout time.Duration) *GetAPIAppVersionsV3SelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API app versions v3 selector params
func (o *GetAPIAppVersionsV3SelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API app versions v3 selector params
func (o *GetAPIAppVersionsV3SelectorParams) WithContext(ctx context.Context) *GetAPIAppVersionsV3SelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API app versions v3 selector params
func (o *GetAPIAppVersionsV3SelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API app versions v3 selector params
func (o *GetAPIAppVersionsV3SelectorParams) WithHTTPClient(client *http.Client) *GetAPIAppVersionsV3SelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API app versions v3 selector params
func (o *GetAPIAppVersionsV3SelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSelector adds the selector to the get API app versions v3 selector params
func (o *GetAPIAppVersionsV3SelectorParams) WithSelector(selector string) *GetAPIAppVersionsV3SelectorParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the get API app versions v3 selector params
func (o *GetAPIAppVersionsV3SelectorParams) SetSelector(selector string) {
	o.Selector = selector
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIAppVersionsV3SelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
