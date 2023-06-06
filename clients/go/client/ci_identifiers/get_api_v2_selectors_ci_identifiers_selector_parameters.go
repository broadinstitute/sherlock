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
)

// NewGetAPIV2SelectorsCiIdentifiersSelectorParams creates a new GetAPIV2SelectorsCiIdentifiersSelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIV2SelectorsCiIdentifiersSelectorParams() *GetAPIV2SelectorsCiIdentifiersSelectorParams {
	return &GetAPIV2SelectorsCiIdentifiersSelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV2SelectorsCiIdentifiersSelectorParamsWithTimeout creates a new GetAPIV2SelectorsCiIdentifiersSelectorParams object
// with the ability to set a timeout on a request.
func NewGetAPIV2SelectorsCiIdentifiersSelectorParamsWithTimeout(timeout time.Duration) *GetAPIV2SelectorsCiIdentifiersSelectorParams {
	return &GetAPIV2SelectorsCiIdentifiersSelectorParams{
		timeout: timeout,
	}
}

// NewGetAPIV2SelectorsCiIdentifiersSelectorParamsWithContext creates a new GetAPIV2SelectorsCiIdentifiersSelectorParams object
// with the ability to set a context for a request.
func NewGetAPIV2SelectorsCiIdentifiersSelectorParamsWithContext(ctx context.Context) *GetAPIV2SelectorsCiIdentifiersSelectorParams {
	return &GetAPIV2SelectorsCiIdentifiersSelectorParams{
		Context: ctx,
	}
}

// NewGetAPIV2SelectorsCiIdentifiersSelectorParamsWithHTTPClient creates a new GetAPIV2SelectorsCiIdentifiersSelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIV2SelectorsCiIdentifiersSelectorParamsWithHTTPClient(client *http.Client) *GetAPIV2SelectorsCiIdentifiersSelectorParams {
	return &GetAPIV2SelectorsCiIdentifiersSelectorParams{
		HTTPClient: client,
	}
}

/* GetAPIV2SelectorsCiIdentifiersSelectorParams contains all the parameters to send to the API endpoint
   for the get API v2 selectors ci identifiers selector operation.

   Typically these are written to a http.Request.
*/
type GetAPIV2SelectorsCiIdentifiersSelectorParams struct {

	/* Selector.

	   The selector of the CiIdentifier to list other selectors for
	*/
	Selector string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API v2 selectors ci identifiers selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2SelectorsCiIdentifiersSelectorParams) WithDefaults() *GetAPIV2SelectorsCiIdentifiersSelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API v2 selectors ci identifiers selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2SelectorsCiIdentifiersSelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API v2 selectors ci identifiers selector params
func (o *GetAPIV2SelectorsCiIdentifiersSelectorParams) WithTimeout(timeout time.Duration) *GetAPIV2SelectorsCiIdentifiersSelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v2 selectors ci identifiers selector params
func (o *GetAPIV2SelectorsCiIdentifiersSelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v2 selectors ci identifiers selector params
func (o *GetAPIV2SelectorsCiIdentifiersSelectorParams) WithContext(ctx context.Context) *GetAPIV2SelectorsCiIdentifiersSelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v2 selectors ci identifiers selector params
func (o *GetAPIV2SelectorsCiIdentifiersSelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v2 selectors ci identifiers selector params
func (o *GetAPIV2SelectorsCiIdentifiersSelectorParams) WithHTTPClient(client *http.Client) *GetAPIV2SelectorsCiIdentifiersSelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v2 selectors ci identifiers selector params
func (o *GetAPIV2SelectorsCiIdentifiersSelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSelector adds the selector to the get API v2 selectors ci identifiers selector params
func (o *GetAPIV2SelectorsCiIdentifiersSelectorParams) WithSelector(selector string) *GetAPIV2SelectorsCiIdentifiersSelectorParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the get API v2 selectors ci identifiers selector params
func (o *GetAPIV2SelectorsCiIdentifiersSelectorParams) SetSelector(selector string) {
	o.Selector = selector
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV2SelectorsCiIdentifiersSelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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