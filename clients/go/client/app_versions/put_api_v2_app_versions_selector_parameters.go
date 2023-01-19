// Code generated by go-swagger; DO NOT EDIT.

package app_versions

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

	"github.com/broadinstitute/sherlock/clients/go/client/models"
)

// NewPutAPIV2AppVersionsSelectorParams creates a new PutAPIV2AppVersionsSelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutAPIV2AppVersionsSelectorParams() *PutAPIV2AppVersionsSelectorParams {
	return &PutAPIV2AppVersionsSelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutAPIV2AppVersionsSelectorParamsWithTimeout creates a new PutAPIV2AppVersionsSelectorParams object
// with the ability to set a timeout on a request.
func NewPutAPIV2AppVersionsSelectorParamsWithTimeout(timeout time.Duration) *PutAPIV2AppVersionsSelectorParams {
	return &PutAPIV2AppVersionsSelectorParams{
		timeout: timeout,
	}
}

// NewPutAPIV2AppVersionsSelectorParamsWithContext creates a new PutAPIV2AppVersionsSelectorParams object
// with the ability to set a context for a request.
func NewPutAPIV2AppVersionsSelectorParamsWithContext(ctx context.Context) *PutAPIV2AppVersionsSelectorParams {
	return &PutAPIV2AppVersionsSelectorParams{
		Context: ctx,
	}
}

// NewPutAPIV2AppVersionsSelectorParamsWithHTTPClient creates a new PutAPIV2AppVersionsSelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutAPIV2AppVersionsSelectorParamsWithHTTPClient(client *http.Client) *PutAPIV2AppVersionsSelectorParams {
	return &PutAPIV2AppVersionsSelectorParams{
		HTTPClient: client,
	}
}

/* PutAPIV2AppVersionsSelectorParams contains all the parameters to send to the API endpoint
   for the put API v2 app versions selector operation.

   Typically these are written to a http.Request.
*/
type PutAPIV2AppVersionsSelectorParams struct {

	/* AppVersion.

	   The AppVersion to upsert
	*/
	AppVersion *models.V2controllersCreatableAppVersion

	/* Selector.

	   The AppVersion to upsert's selector: chart/version or numeric ID
	*/
	Selector string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put API v2 app versions selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutAPIV2AppVersionsSelectorParams) WithDefaults() *PutAPIV2AppVersionsSelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put API v2 app versions selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutAPIV2AppVersionsSelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put API v2 app versions selector params
func (o *PutAPIV2AppVersionsSelectorParams) WithTimeout(timeout time.Duration) *PutAPIV2AppVersionsSelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put API v2 app versions selector params
func (o *PutAPIV2AppVersionsSelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put API v2 app versions selector params
func (o *PutAPIV2AppVersionsSelectorParams) WithContext(ctx context.Context) *PutAPIV2AppVersionsSelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put API v2 app versions selector params
func (o *PutAPIV2AppVersionsSelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put API v2 app versions selector params
func (o *PutAPIV2AppVersionsSelectorParams) WithHTTPClient(client *http.Client) *PutAPIV2AppVersionsSelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put API v2 app versions selector params
func (o *PutAPIV2AppVersionsSelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppVersion adds the appVersion to the put API v2 app versions selector params
func (o *PutAPIV2AppVersionsSelectorParams) WithAppVersion(appVersion *models.V2controllersCreatableAppVersion) *PutAPIV2AppVersionsSelectorParams {
	o.SetAppVersion(appVersion)
	return o
}

// SetAppVersion adds the appVersion to the put API v2 app versions selector params
func (o *PutAPIV2AppVersionsSelectorParams) SetAppVersion(appVersion *models.V2controllersCreatableAppVersion) {
	o.AppVersion = appVersion
}

// WithSelector adds the selector to the put API v2 app versions selector params
func (o *PutAPIV2AppVersionsSelectorParams) WithSelector(selector string) *PutAPIV2AppVersionsSelectorParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the put API v2 app versions selector params
func (o *PutAPIV2AppVersionsSelectorParams) SetSelector(selector string) {
	o.Selector = selector
}

// WriteToRequest writes these params to a swagger request
func (o *PutAPIV2AppVersionsSelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.AppVersion != nil {
		if err := r.SetBodyParam(o.AppVersion); err != nil {
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
