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

// NewGetAPIV2ProceduresUsersMeParams creates a new GetAPIV2ProceduresUsersMeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIV2ProceduresUsersMeParams() *GetAPIV2ProceduresUsersMeParams {
	return &GetAPIV2ProceduresUsersMeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV2ProceduresUsersMeParamsWithTimeout creates a new GetAPIV2ProceduresUsersMeParams object
// with the ability to set a timeout on a request.
func NewGetAPIV2ProceduresUsersMeParamsWithTimeout(timeout time.Duration) *GetAPIV2ProceduresUsersMeParams {
	return &GetAPIV2ProceduresUsersMeParams{
		timeout: timeout,
	}
}

// NewGetAPIV2ProceduresUsersMeParamsWithContext creates a new GetAPIV2ProceduresUsersMeParams object
// with the ability to set a context for a request.
func NewGetAPIV2ProceduresUsersMeParamsWithContext(ctx context.Context) *GetAPIV2ProceduresUsersMeParams {
	return &GetAPIV2ProceduresUsersMeParams{
		Context: ctx,
	}
}

// NewGetAPIV2ProceduresUsersMeParamsWithHTTPClient creates a new GetAPIV2ProceduresUsersMeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIV2ProceduresUsersMeParamsWithHTTPClient(client *http.Client) *GetAPIV2ProceduresUsersMeParams {
	return &GetAPIV2ProceduresUsersMeParams{
		HTTPClient: client,
	}
}

/* GetAPIV2ProceduresUsersMeParams contains all the parameters to send to the API endpoint
   for the get API v2 procedures users me operation.

   Typically these are written to a http.Request.
*/
type GetAPIV2ProceduresUsersMeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API v2 procedures users me params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2ProceduresUsersMeParams) WithDefaults() *GetAPIV2ProceduresUsersMeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API v2 procedures users me params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2ProceduresUsersMeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API v2 procedures users me params
func (o *GetAPIV2ProceduresUsersMeParams) WithTimeout(timeout time.Duration) *GetAPIV2ProceduresUsersMeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v2 procedures users me params
func (o *GetAPIV2ProceduresUsersMeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v2 procedures users me params
func (o *GetAPIV2ProceduresUsersMeParams) WithContext(ctx context.Context) *GetAPIV2ProceduresUsersMeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v2 procedures users me params
func (o *GetAPIV2ProceduresUsersMeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v2 procedures users me params
func (o *GetAPIV2ProceduresUsersMeParams) WithHTTPClient(client *http.Client) *GetAPIV2ProceduresUsersMeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v2 procedures users me params
func (o *GetAPIV2ProceduresUsersMeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV2ProceduresUsersMeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}