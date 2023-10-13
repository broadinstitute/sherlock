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

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// NewPutAPIAppVersionsV3Params creates a new PutAPIAppVersionsV3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutAPIAppVersionsV3Params() *PutAPIAppVersionsV3Params {
	return &PutAPIAppVersionsV3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutAPIAppVersionsV3ParamsWithTimeout creates a new PutAPIAppVersionsV3Params object
// with the ability to set a timeout on a request.
func NewPutAPIAppVersionsV3ParamsWithTimeout(timeout time.Duration) *PutAPIAppVersionsV3Params {
	return &PutAPIAppVersionsV3Params{
		timeout: timeout,
	}
}

// NewPutAPIAppVersionsV3ParamsWithContext creates a new PutAPIAppVersionsV3Params object
// with the ability to set a context for a request.
func NewPutAPIAppVersionsV3ParamsWithContext(ctx context.Context) *PutAPIAppVersionsV3Params {
	return &PutAPIAppVersionsV3Params{
		Context: ctx,
	}
}

// NewPutAPIAppVersionsV3ParamsWithHTTPClient creates a new PutAPIAppVersionsV3Params object
// with the ability to set a custom HTTPClient for a request.
func NewPutAPIAppVersionsV3ParamsWithHTTPClient(client *http.Client) *PutAPIAppVersionsV3Params {
	return &PutAPIAppVersionsV3Params{
		HTTPClient: client,
	}
}

/* PutAPIAppVersionsV3Params contains all the parameters to send to the API endpoint
   for the put API app versions v3 operation.

   Typically these are written to a http.Request.
*/
type PutAPIAppVersionsV3Params struct {

	/* AppVersion.

	   The AppVersion to upsert
	*/
	AppVersion *models.SherlockAppVersionV3Create

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put API app versions v3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutAPIAppVersionsV3Params) WithDefaults() *PutAPIAppVersionsV3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put API app versions v3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutAPIAppVersionsV3Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put API app versions v3 params
func (o *PutAPIAppVersionsV3Params) WithTimeout(timeout time.Duration) *PutAPIAppVersionsV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put API app versions v3 params
func (o *PutAPIAppVersionsV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put API app versions v3 params
func (o *PutAPIAppVersionsV3Params) WithContext(ctx context.Context) *PutAPIAppVersionsV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put API app versions v3 params
func (o *PutAPIAppVersionsV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put API app versions v3 params
func (o *PutAPIAppVersionsV3Params) WithHTTPClient(client *http.Client) *PutAPIAppVersionsV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put API app versions v3 params
func (o *PutAPIAppVersionsV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppVersion adds the appVersion to the put API app versions v3 params
func (o *PutAPIAppVersionsV3Params) WithAppVersion(appVersion *models.SherlockAppVersionV3Create) *PutAPIAppVersionsV3Params {
	o.SetAppVersion(appVersion)
	return o
}

// SetAppVersion adds the appVersion to the put API app versions v3 params
func (o *PutAPIAppVersionsV3Params) SetAppVersion(appVersion *models.SherlockAppVersionV3Create) {
	o.AppVersion = appVersion
}

// WriteToRequest writes these params to a swagger request
func (o *PutAPIAppVersionsV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.AppVersion != nil {
		if err := r.SetBodyParam(o.AppVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
