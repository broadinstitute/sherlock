// Code generated by go-swagger; DO NOT EDIT.

package deploy_hooks

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

// NewPatchAPIDeployHooksSlackV3SelectorParams creates a new PatchAPIDeployHooksSlackV3SelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAPIDeployHooksSlackV3SelectorParams() *PatchAPIDeployHooksSlackV3SelectorParams {
	return &PatchAPIDeployHooksSlackV3SelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPIDeployHooksSlackV3SelectorParamsWithTimeout creates a new PatchAPIDeployHooksSlackV3SelectorParams object
// with the ability to set a timeout on a request.
func NewPatchAPIDeployHooksSlackV3SelectorParamsWithTimeout(timeout time.Duration) *PatchAPIDeployHooksSlackV3SelectorParams {
	return &PatchAPIDeployHooksSlackV3SelectorParams{
		timeout: timeout,
	}
}

// NewPatchAPIDeployHooksSlackV3SelectorParamsWithContext creates a new PatchAPIDeployHooksSlackV3SelectorParams object
// with the ability to set a context for a request.
func NewPatchAPIDeployHooksSlackV3SelectorParamsWithContext(ctx context.Context) *PatchAPIDeployHooksSlackV3SelectorParams {
	return &PatchAPIDeployHooksSlackV3SelectorParams{
		Context: ctx,
	}
}

// NewPatchAPIDeployHooksSlackV3SelectorParamsWithHTTPClient creates a new PatchAPIDeployHooksSlackV3SelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAPIDeployHooksSlackV3SelectorParamsWithHTTPClient(client *http.Client) *PatchAPIDeployHooksSlackV3SelectorParams {
	return &PatchAPIDeployHooksSlackV3SelectorParams{
		HTTPClient: client,
	}
}

/* PatchAPIDeployHooksSlackV3SelectorParams contains all the parameters to send to the API endpoint
   for the patch API deploy hooks slack v3 selector operation.

   Typically these are written to a http.Request.
*/
type PatchAPIDeployHooksSlackV3SelectorParams struct {

	/* Selector.

	   The ID of the SlackDeployHook to edit
	*/
	Selector string

	/* SlackDeployHook.

	   The edits to make to the SlackDeployHook
	*/
	SlackDeployHook *models.SherlockSlackDeployHookV3Edit

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch API deploy hooks slack v3 selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIDeployHooksSlackV3SelectorParams) WithDefaults() *PatchAPIDeployHooksSlackV3SelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch API deploy hooks slack v3 selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIDeployHooksSlackV3SelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch API deploy hooks slack v3 selector params
func (o *PatchAPIDeployHooksSlackV3SelectorParams) WithTimeout(timeout time.Duration) *PatchAPIDeployHooksSlackV3SelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API deploy hooks slack v3 selector params
func (o *PatchAPIDeployHooksSlackV3SelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API deploy hooks slack v3 selector params
func (o *PatchAPIDeployHooksSlackV3SelectorParams) WithContext(ctx context.Context) *PatchAPIDeployHooksSlackV3SelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API deploy hooks slack v3 selector params
func (o *PatchAPIDeployHooksSlackV3SelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API deploy hooks slack v3 selector params
func (o *PatchAPIDeployHooksSlackV3SelectorParams) WithHTTPClient(client *http.Client) *PatchAPIDeployHooksSlackV3SelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API deploy hooks slack v3 selector params
func (o *PatchAPIDeployHooksSlackV3SelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSelector adds the selector to the patch API deploy hooks slack v3 selector params
func (o *PatchAPIDeployHooksSlackV3SelectorParams) WithSelector(selector string) *PatchAPIDeployHooksSlackV3SelectorParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the patch API deploy hooks slack v3 selector params
func (o *PatchAPIDeployHooksSlackV3SelectorParams) SetSelector(selector string) {
	o.Selector = selector
}

// WithSlackDeployHook adds the slackDeployHook to the patch API deploy hooks slack v3 selector params
func (o *PatchAPIDeployHooksSlackV3SelectorParams) WithSlackDeployHook(slackDeployHook *models.SherlockSlackDeployHookV3Edit) *PatchAPIDeployHooksSlackV3SelectorParams {
	o.SetSlackDeployHook(slackDeployHook)
	return o
}

// SetSlackDeployHook adds the slackDeployHook to the patch API deploy hooks slack v3 selector params
func (o *PatchAPIDeployHooksSlackV3SelectorParams) SetSlackDeployHook(slackDeployHook *models.SherlockSlackDeployHookV3Edit) {
	o.SlackDeployHook = slackDeployHook
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPIDeployHooksSlackV3SelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param selector
	if err := r.SetPathParam("selector", o.Selector); err != nil {
		return err
	}
	if o.SlackDeployHook != nil {
		if err := r.SetBodyParam(o.SlackDeployHook); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}