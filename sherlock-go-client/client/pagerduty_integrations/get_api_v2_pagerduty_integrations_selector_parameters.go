// Code generated by go-swagger; DO NOT EDIT.

package pagerduty_integrations

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

// NewGetAPIV2PagerdutyIntegrationsSelectorParams creates a new GetAPIV2PagerdutyIntegrationsSelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIV2PagerdutyIntegrationsSelectorParams() *GetAPIV2PagerdutyIntegrationsSelectorParams {
	return &GetAPIV2PagerdutyIntegrationsSelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV2PagerdutyIntegrationsSelectorParamsWithTimeout creates a new GetAPIV2PagerdutyIntegrationsSelectorParams object
// with the ability to set a timeout on a request.
func NewGetAPIV2PagerdutyIntegrationsSelectorParamsWithTimeout(timeout time.Duration) *GetAPIV2PagerdutyIntegrationsSelectorParams {
	return &GetAPIV2PagerdutyIntegrationsSelectorParams{
		timeout: timeout,
	}
}

// NewGetAPIV2PagerdutyIntegrationsSelectorParamsWithContext creates a new GetAPIV2PagerdutyIntegrationsSelectorParams object
// with the ability to set a context for a request.
func NewGetAPIV2PagerdutyIntegrationsSelectorParamsWithContext(ctx context.Context) *GetAPIV2PagerdutyIntegrationsSelectorParams {
	return &GetAPIV2PagerdutyIntegrationsSelectorParams{
		Context: ctx,
	}
}

// NewGetAPIV2PagerdutyIntegrationsSelectorParamsWithHTTPClient creates a new GetAPIV2PagerdutyIntegrationsSelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIV2PagerdutyIntegrationsSelectorParamsWithHTTPClient(client *http.Client) *GetAPIV2PagerdutyIntegrationsSelectorParams {
	return &GetAPIV2PagerdutyIntegrationsSelectorParams{
		HTTPClient: client,
	}
}

/* GetAPIV2PagerdutyIntegrationsSelectorParams contains all the parameters to send to the API endpoint
   for the get API v2 pagerduty integrations selector operation.

   Typically these are written to a http.Request.
*/
type GetAPIV2PagerdutyIntegrationsSelectorParams struct {

	/* Selector.

	   The PagerdutyIntegration to get's selector: chart/version or numeric ID
	*/
	Selector string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API v2 pagerduty integrations selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2PagerdutyIntegrationsSelectorParams) WithDefaults() *GetAPIV2PagerdutyIntegrationsSelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API v2 pagerduty integrations selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2PagerdutyIntegrationsSelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API v2 pagerduty integrations selector params
func (o *GetAPIV2PagerdutyIntegrationsSelectorParams) WithTimeout(timeout time.Duration) *GetAPIV2PagerdutyIntegrationsSelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v2 pagerduty integrations selector params
func (o *GetAPIV2PagerdutyIntegrationsSelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v2 pagerduty integrations selector params
func (o *GetAPIV2PagerdutyIntegrationsSelectorParams) WithContext(ctx context.Context) *GetAPIV2PagerdutyIntegrationsSelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v2 pagerduty integrations selector params
func (o *GetAPIV2PagerdutyIntegrationsSelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v2 pagerduty integrations selector params
func (o *GetAPIV2PagerdutyIntegrationsSelectorParams) WithHTTPClient(client *http.Client) *GetAPIV2PagerdutyIntegrationsSelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v2 pagerduty integrations selector params
func (o *GetAPIV2PagerdutyIntegrationsSelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSelector adds the selector to the get API v2 pagerduty integrations selector params
func (o *GetAPIV2PagerdutyIntegrationsSelectorParams) WithSelector(selector string) *GetAPIV2PagerdutyIntegrationsSelectorParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the get API v2 pagerduty integrations selector params
func (o *GetAPIV2PagerdutyIntegrationsSelectorParams) SetSelector(selector string) {
	o.Selector = selector
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV2PagerdutyIntegrationsSelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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