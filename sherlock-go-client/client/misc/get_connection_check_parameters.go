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

// NewGetConnectionCheckParams creates a new GetConnectionCheckParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetConnectionCheckParams() *GetConnectionCheckParams {
	return &GetConnectionCheckParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetConnectionCheckParamsWithTimeout creates a new GetConnectionCheckParams object
// with the ability to set a timeout on a request.
func NewGetConnectionCheckParamsWithTimeout(timeout time.Duration) *GetConnectionCheckParams {
	return &GetConnectionCheckParams{
		timeout: timeout,
	}
}

// NewGetConnectionCheckParamsWithContext creates a new GetConnectionCheckParams object
// with the ability to set a context for a request.
func NewGetConnectionCheckParamsWithContext(ctx context.Context) *GetConnectionCheckParams {
	return &GetConnectionCheckParams{
		Context: ctx,
	}
}

// NewGetConnectionCheckParamsWithHTTPClient creates a new GetConnectionCheckParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetConnectionCheckParamsWithHTTPClient(client *http.Client) *GetConnectionCheckParams {
	return &GetConnectionCheckParams{
		HTTPClient: client,
	}
}

/* GetConnectionCheckParams contains all the parameters to send to the API endpoint
   for the get connection check operation.

   Typically these are written to a http.Request.
*/
type GetConnectionCheckParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get connection check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConnectionCheckParams) WithDefaults() *GetConnectionCheckParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get connection check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConnectionCheckParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get connection check params
func (o *GetConnectionCheckParams) WithTimeout(timeout time.Duration) *GetConnectionCheckParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get connection check params
func (o *GetConnectionCheckParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get connection check params
func (o *GetConnectionCheckParams) WithContext(ctx context.Context) *GetConnectionCheckParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get connection check params
func (o *GetConnectionCheckParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get connection check params
func (o *GetConnectionCheckParams) WithHTTPClient(client *http.Client) *GetConnectionCheckParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get connection check params
func (o *GetConnectionCheckParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetConnectionCheckParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
