// Code generated by go-swagger; DO NOT EDIT.

package misc

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

// NewGetStatusParams creates a new GetStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetStatusParams() *GetStatusParams {
	return &GetStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetStatusParamsWithTimeout creates a new GetStatusParams object
// with the ability to set a timeout on a request.
func NewGetStatusParamsWithTimeout(timeout time.Duration) *GetStatusParams {
	return &GetStatusParams{
		timeout: timeout,
	}
}

// NewGetStatusParamsWithContext creates a new GetStatusParams object
// with the ability to set a context for a request.
func NewGetStatusParamsWithContext(ctx context.Context) *GetStatusParams {
	return &GetStatusParams{
		Context: ctx,
	}
}

// NewGetStatusParamsWithHTTPClient creates a new GetStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetStatusParamsWithHTTPClient(client *http.Client) *GetStatusParams {
	return &GetStatusParams{
		HTTPClient: client,
	}
}

/* GetStatusParams contains all the parameters to send to the API endpoint
   for the get status operation.

   Typically these are written to a http.Request.
*/
type GetStatusParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStatusParams) WithDefaults() *GetStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get status params
func (o *GetStatusParams) WithTimeout(timeout time.Duration) *GetStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get status params
func (o *GetStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get status params
func (o *GetStatusParams) WithContext(ctx context.Context) *GetStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get status params
func (o *GetStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get status params
func (o *GetStatusParams) WithHTTPClient(client *http.Client) *GetStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get status params
func (o *GetStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
