// Code generated by go-swagger; DO NOT EDIT.

package roles

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

// NewGetAPIRolesV3SelectorParams creates a new GetAPIRolesV3SelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIRolesV3SelectorParams() *GetAPIRolesV3SelectorParams {
	return &GetAPIRolesV3SelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIRolesV3SelectorParamsWithTimeout creates a new GetAPIRolesV3SelectorParams object
// with the ability to set a timeout on a request.
func NewGetAPIRolesV3SelectorParamsWithTimeout(timeout time.Duration) *GetAPIRolesV3SelectorParams {
	return &GetAPIRolesV3SelectorParams{
		timeout: timeout,
	}
}

// NewGetAPIRolesV3SelectorParamsWithContext creates a new GetAPIRolesV3SelectorParams object
// with the ability to set a context for a request.
func NewGetAPIRolesV3SelectorParamsWithContext(ctx context.Context) *GetAPIRolesV3SelectorParams {
	return &GetAPIRolesV3SelectorParams{
		Context: ctx,
	}
}

// NewGetAPIRolesV3SelectorParamsWithHTTPClient creates a new GetAPIRolesV3SelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIRolesV3SelectorParamsWithHTTPClient(client *http.Client) *GetAPIRolesV3SelectorParams {
	return &GetAPIRolesV3SelectorParams{
		HTTPClient: client,
	}
}

/* GetAPIRolesV3SelectorParams contains all the parameters to send to the API endpoint
   for the get API roles v3 selector operation.

   Typically these are written to a http.Request.
*/
type GetAPIRolesV3SelectorParams struct {

	/* Selector.

	   The selector of the Role, which can be either the numeric ID or the name
	*/
	Selector string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API roles v3 selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIRolesV3SelectorParams) WithDefaults() *GetAPIRolesV3SelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API roles v3 selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIRolesV3SelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API roles v3 selector params
func (o *GetAPIRolesV3SelectorParams) WithTimeout(timeout time.Duration) *GetAPIRolesV3SelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API roles v3 selector params
func (o *GetAPIRolesV3SelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API roles v3 selector params
func (o *GetAPIRolesV3SelectorParams) WithContext(ctx context.Context) *GetAPIRolesV3SelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API roles v3 selector params
func (o *GetAPIRolesV3SelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API roles v3 selector params
func (o *GetAPIRolesV3SelectorParams) WithHTTPClient(client *http.Client) *GetAPIRolesV3SelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API roles v3 selector params
func (o *GetAPIRolesV3SelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSelector adds the selector to the get API roles v3 selector params
func (o *GetAPIRolesV3SelectorParams) WithSelector(selector string) *GetAPIRolesV3SelectorParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the get API roles v3 selector params
func (o *GetAPIRolesV3SelectorParams) SetSelector(selector string) {
	o.Selector = selector
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIRolesV3SelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
