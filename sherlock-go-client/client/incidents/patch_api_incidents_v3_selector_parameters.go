// Code generated by go-swagger; DO NOT EDIT.

package incidents

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

// NewPatchAPIIncidentsV3SelectorParams creates a new PatchAPIIncidentsV3SelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAPIIncidentsV3SelectorParams() *PatchAPIIncidentsV3SelectorParams {
	return &PatchAPIIncidentsV3SelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPIIncidentsV3SelectorParamsWithTimeout creates a new PatchAPIIncidentsV3SelectorParams object
// with the ability to set a timeout on a request.
func NewPatchAPIIncidentsV3SelectorParamsWithTimeout(timeout time.Duration) *PatchAPIIncidentsV3SelectorParams {
	return &PatchAPIIncidentsV3SelectorParams{
		timeout: timeout,
	}
}

// NewPatchAPIIncidentsV3SelectorParamsWithContext creates a new PatchAPIIncidentsV3SelectorParams object
// with the ability to set a context for a request.
func NewPatchAPIIncidentsV3SelectorParamsWithContext(ctx context.Context) *PatchAPIIncidentsV3SelectorParams {
	return &PatchAPIIncidentsV3SelectorParams{
		Context: ctx,
	}
}

// NewPatchAPIIncidentsV3SelectorParamsWithHTTPClient creates a new PatchAPIIncidentsV3SelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAPIIncidentsV3SelectorParamsWithHTTPClient(client *http.Client) *PatchAPIIncidentsV3SelectorParams {
	return &PatchAPIIncidentsV3SelectorParams{
		HTTPClient: client,
	}
}

/* PatchAPIIncidentsV3SelectorParams contains all the parameters to send to the API endpoint
   for the patch API incidents v3 selector operation.

   Typically these are written to a http.Request.
*/
type PatchAPIIncidentsV3SelectorParams struct {

	/* Incident.

	   The edits to make to the Incident
	*/
	Incident *models.SherlockIncidentV3Edit

	/* Selector.

	   The ID of the Incident
	*/
	Selector string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch API incidents v3 selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIIncidentsV3SelectorParams) WithDefaults() *PatchAPIIncidentsV3SelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch API incidents v3 selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIIncidentsV3SelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch API incidents v3 selector params
func (o *PatchAPIIncidentsV3SelectorParams) WithTimeout(timeout time.Duration) *PatchAPIIncidentsV3SelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API incidents v3 selector params
func (o *PatchAPIIncidentsV3SelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API incidents v3 selector params
func (o *PatchAPIIncidentsV3SelectorParams) WithContext(ctx context.Context) *PatchAPIIncidentsV3SelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API incidents v3 selector params
func (o *PatchAPIIncidentsV3SelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API incidents v3 selector params
func (o *PatchAPIIncidentsV3SelectorParams) WithHTTPClient(client *http.Client) *PatchAPIIncidentsV3SelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API incidents v3 selector params
func (o *PatchAPIIncidentsV3SelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncident adds the incident to the patch API incidents v3 selector params
func (o *PatchAPIIncidentsV3SelectorParams) WithIncident(incident *models.SherlockIncidentV3Edit) *PatchAPIIncidentsV3SelectorParams {
	o.SetIncident(incident)
	return o
}

// SetIncident adds the incident to the patch API incidents v3 selector params
func (o *PatchAPIIncidentsV3SelectorParams) SetIncident(incident *models.SherlockIncidentV3Edit) {
	o.Incident = incident
}

// WithSelector adds the selector to the patch API incidents v3 selector params
func (o *PatchAPIIncidentsV3SelectorParams) WithSelector(selector string) *PatchAPIIncidentsV3SelectorParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the patch API incidents v3 selector params
func (o *PatchAPIIncidentsV3SelectorParams) SetSelector(selector string) {
	o.Selector = selector
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPIIncidentsV3SelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Incident != nil {
		if err := r.SetBodyParam(o.Incident); err != nil {
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
