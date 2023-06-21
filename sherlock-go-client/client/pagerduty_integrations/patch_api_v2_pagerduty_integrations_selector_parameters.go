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

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// NewPatchAPIV2PagerdutyIntegrationsSelectorParams creates a new PatchAPIV2PagerdutyIntegrationsSelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAPIV2PagerdutyIntegrationsSelectorParams() *PatchAPIV2PagerdutyIntegrationsSelectorParams {
	return &PatchAPIV2PagerdutyIntegrationsSelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPIV2PagerdutyIntegrationsSelectorParamsWithTimeout creates a new PatchAPIV2PagerdutyIntegrationsSelectorParams object
// with the ability to set a timeout on a request.
func NewPatchAPIV2PagerdutyIntegrationsSelectorParamsWithTimeout(timeout time.Duration) *PatchAPIV2PagerdutyIntegrationsSelectorParams {
	return &PatchAPIV2PagerdutyIntegrationsSelectorParams{
		timeout: timeout,
	}
}

// NewPatchAPIV2PagerdutyIntegrationsSelectorParamsWithContext creates a new PatchAPIV2PagerdutyIntegrationsSelectorParams object
// with the ability to set a context for a request.
func NewPatchAPIV2PagerdutyIntegrationsSelectorParamsWithContext(ctx context.Context) *PatchAPIV2PagerdutyIntegrationsSelectorParams {
	return &PatchAPIV2PagerdutyIntegrationsSelectorParams{
		Context: ctx,
	}
}

// NewPatchAPIV2PagerdutyIntegrationsSelectorParamsWithHTTPClient creates a new PatchAPIV2PagerdutyIntegrationsSelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAPIV2PagerdutyIntegrationsSelectorParamsWithHTTPClient(client *http.Client) *PatchAPIV2PagerdutyIntegrationsSelectorParams {
	return &PatchAPIV2PagerdutyIntegrationsSelectorParams{
		HTTPClient: client,
	}
}

/* PatchAPIV2PagerdutyIntegrationsSelectorParams contains all the parameters to send to the API endpoint
   for the patch API v2 pagerduty integrations selector operation.

   Typically these are written to a http.Request.
*/
type PatchAPIV2PagerdutyIntegrationsSelectorParams struct {

	/* PagerdutyIntegration.

	   The edits to make to the PagerdutyIntegration
	*/
	PagerdutyIntegration *models.V2controllersEditablePagerdutyIntegration

	/* Selector.

	   The PagerdutyIntegration to edit's selector: chart/version or numeric ID
	*/
	Selector string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch API v2 pagerduty integrations selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIV2PagerdutyIntegrationsSelectorParams) WithDefaults() *PatchAPIV2PagerdutyIntegrationsSelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch API v2 pagerduty integrations selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIV2PagerdutyIntegrationsSelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch API v2 pagerduty integrations selector params
func (o *PatchAPIV2PagerdutyIntegrationsSelectorParams) WithTimeout(timeout time.Duration) *PatchAPIV2PagerdutyIntegrationsSelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API v2 pagerduty integrations selector params
func (o *PatchAPIV2PagerdutyIntegrationsSelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API v2 pagerduty integrations selector params
func (o *PatchAPIV2PagerdutyIntegrationsSelectorParams) WithContext(ctx context.Context) *PatchAPIV2PagerdutyIntegrationsSelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API v2 pagerduty integrations selector params
func (o *PatchAPIV2PagerdutyIntegrationsSelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API v2 pagerduty integrations selector params
func (o *PatchAPIV2PagerdutyIntegrationsSelectorParams) WithHTTPClient(client *http.Client) *PatchAPIV2PagerdutyIntegrationsSelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API v2 pagerduty integrations selector params
func (o *PatchAPIV2PagerdutyIntegrationsSelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPagerdutyIntegration adds the pagerdutyIntegration to the patch API v2 pagerduty integrations selector params
func (o *PatchAPIV2PagerdutyIntegrationsSelectorParams) WithPagerdutyIntegration(pagerdutyIntegration *models.V2controllersEditablePagerdutyIntegration) *PatchAPIV2PagerdutyIntegrationsSelectorParams {
	o.SetPagerdutyIntegration(pagerdutyIntegration)
	return o
}

// SetPagerdutyIntegration adds the pagerdutyIntegration to the patch API v2 pagerduty integrations selector params
func (o *PatchAPIV2PagerdutyIntegrationsSelectorParams) SetPagerdutyIntegration(pagerdutyIntegration *models.V2controllersEditablePagerdutyIntegration) {
	o.PagerdutyIntegration = pagerdutyIntegration
}

// WithSelector adds the selector to the patch API v2 pagerduty integrations selector params
func (o *PatchAPIV2PagerdutyIntegrationsSelectorParams) WithSelector(selector string) *PatchAPIV2PagerdutyIntegrationsSelectorParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the patch API v2 pagerduty integrations selector params
func (o *PatchAPIV2PagerdutyIntegrationsSelectorParams) SetSelector(selector string) {
	o.Selector = selector
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPIV2PagerdutyIntegrationsSelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PagerdutyIntegration != nil {
		if err := r.SetBodyParam(o.PagerdutyIntegration); err != nil {
			return err
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