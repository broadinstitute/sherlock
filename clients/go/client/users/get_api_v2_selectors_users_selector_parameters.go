// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewGetAPIV2SelectorsUsersSelectorParams creates a new GetAPIV2SelectorsUsersSelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIV2SelectorsUsersSelectorParams() *GetAPIV2SelectorsUsersSelectorParams {
	return &GetAPIV2SelectorsUsersSelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV2SelectorsUsersSelectorParamsWithTimeout creates a new GetAPIV2SelectorsUsersSelectorParams object
// with the ability to set a timeout on a request.
func NewGetAPIV2SelectorsUsersSelectorParamsWithTimeout(timeout time.Duration) *GetAPIV2SelectorsUsersSelectorParams {
	return &GetAPIV2SelectorsUsersSelectorParams{
		timeout: timeout,
	}
}

// NewGetAPIV2SelectorsUsersSelectorParamsWithContext creates a new GetAPIV2SelectorsUsersSelectorParams object
// with the ability to set a context for a request.
func NewGetAPIV2SelectorsUsersSelectorParamsWithContext(ctx context.Context) *GetAPIV2SelectorsUsersSelectorParams {
	return &GetAPIV2SelectorsUsersSelectorParams{
		Context: ctx,
	}
}

// NewGetAPIV2SelectorsUsersSelectorParamsWithHTTPClient creates a new GetAPIV2SelectorsUsersSelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIV2SelectorsUsersSelectorParamsWithHTTPClient(client *http.Client) *GetAPIV2SelectorsUsersSelectorParams {
	return &GetAPIV2SelectorsUsersSelectorParams{
		HTTPClient: client,
	}
}

/* GetAPIV2SelectorsUsersSelectorParams contains all the parameters to send to the API endpoint
   for the get API v2 selectors users selector operation.

   Typically these are written to a http.Request.
*/
type GetAPIV2SelectorsUsersSelectorParams struct {

	/* Selector.

	   The selector of the User to list other selectors for
	*/
	Selector string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API v2 selectors users selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2SelectorsUsersSelectorParams) WithDefaults() *GetAPIV2SelectorsUsersSelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API v2 selectors users selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2SelectorsUsersSelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API v2 selectors users selector params
func (o *GetAPIV2SelectorsUsersSelectorParams) WithTimeout(timeout time.Duration) *GetAPIV2SelectorsUsersSelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v2 selectors users selector params
func (o *GetAPIV2SelectorsUsersSelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v2 selectors users selector params
func (o *GetAPIV2SelectorsUsersSelectorParams) WithContext(ctx context.Context) *GetAPIV2SelectorsUsersSelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v2 selectors users selector params
func (o *GetAPIV2SelectorsUsersSelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v2 selectors users selector params
func (o *GetAPIV2SelectorsUsersSelectorParams) WithHTTPClient(client *http.Client) *GetAPIV2SelectorsUsersSelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v2 selectors users selector params
func (o *GetAPIV2SelectorsUsersSelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSelector adds the selector to the get API v2 selectors users selector params
func (o *GetAPIV2SelectorsUsersSelectorParams) WithSelector(selector string) *GetAPIV2SelectorsUsersSelectorParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the get API v2 selectors users selector params
func (o *GetAPIV2SelectorsUsersSelectorParams) SetSelector(selector string) {
	o.Selector = selector
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV2SelectorsUsersSelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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