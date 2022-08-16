// Code generated by go-swagger; DO NOT EDIT.

package environments

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

	"github.com/broadinstitute/sherlock/client/go/client/models"
)

// NewPostAPIV2EnvironmentsParams creates a new PostAPIV2EnvironmentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAPIV2EnvironmentsParams() *PostAPIV2EnvironmentsParams {
	return &PostAPIV2EnvironmentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV2EnvironmentsParamsWithTimeout creates a new PostAPIV2EnvironmentsParams object
// with the ability to set a timeout on a request.
func NewPostAPIV2EnvironmentsParamsWithTimeout(timeout time.Duration) *PostAPIV2EnvironmentsParams {
	return &PostAPIV2EnvironmentsParams{
		timeout: timeout,
	}
}

// NewPostAPIV2EnvironmentsParamsWithContext creates a new PostAPIV2EnvironmentsParams object
// with the ability to set a context for a request.
func NewPostAPIV2EnvironmentsParamsWithContext(ctx context.Context) *PostAPIV2EnvironmentsParams {
	return &PostAPIV2EnvironmentsParams{
		Context: ctx,
	}
}

// NewPostAPIV2EnvironmentsParamsWithHTTPClient creates a new PostAPIV2EnvironmentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAPIV2EnvironmentsParamsWithHTTPClient(client *http.Client) *PostAPIV2EnvironmentsParams {
	return &PostAPIV2EnvironmentsParams{
		HTTPClient: client,
	}
}

/* PostAPIV2EnvironmentsParams contains all the parameters to send to the API endpoint
   for the post API v2 environments operation.

   Typically these are written to a http.Request.
*/
type PostAPIV2EnvironmentsParams struct {

	/* Environment.

	   The Environment to create
	*/
	Environment *models.V2controllersCreatableEnvironment

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post API v2 environments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIV2EnvironmentsParams) WithDefaults() *PostAPIV2EnvironmentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post API v2 environments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIV2EnvironmentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post API v2 environments params
func (o *PostAPIV2EnvironmentsParams) WithTimeout(timeout time.Duration) *PostAPIV2EnvironmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v2 environments params
func (o *PostAPIV2EnvironmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v2 environments params
func (o *PostAPIV2EnvironmentsParams) WithContext(ctx context.Context) *PostAPIV2EnvironmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v2 environments params
func (o *PostAPIV2EnvironmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v2 environments params
func (o *PostAPIV2EnvironmentsParams) WithHTTPClient(client *http.Client) *PostAPIV2EnvironmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v2 environments params
func (o *PostAPIV2EnvironmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironment adds the environment to the post API v2 environments params
func (o *PostAPIV2EnvironmentsParams) WithEnvironment(environment *models.V2controllersCreatableEnvironment) *PostAPIV2EnvironmentsParams {
	o.SetEnvironment(environment)
	return o
}

// SetEnvironment adds the environment to the post API v2 environments params
func (o *PostAPIV2EnvironmentsParams) SetEnvironment(environment *models.V2controllersCreatableEnvironment) {
	o.Environment = environment
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV2EnvironmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Environment != nil {
		if err := r.SetBodyParam(o.Environment); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
