// Code generated by go-swagger; DO NOT EDIT.

package changesets

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

// NewPostAPIV2ChangesetsParams creates a new PostAPIV2ChangesetsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAPIV2ChangesetsParams() *PostAPIV2ChangesetsParams {
	return &PostAPIV2ChangesetsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV2ChangesetsParamsWithTimeout creates a new PostAPIV2ChangesetsParams object
// with the ability to set a timeout on a request.
func NewPostAPIV2ChangesetsParamsWithTimeout(timeout time.Duration) *PostAPIV2ChangesetsParams {
	return &PostAPIV2ChangesetsParams{
		timeout: timeout,
	}
}

// NewPostAPIV2ChangesetsParamsWithContext creates a new PostAPIV2ChangesetsParams object
// with the ability to set a context for a request.
func NewPostAPIV2ChangesetsParamsWithContext(ctx context.Context) *PostAPIV2ChangesetsParams {
	return &PostAPIV2ChangesetsParams{
		Context: ctx,
	}
}

// NewPostAPIV2ChangesetsParamsWithHTTPClient creates a new PostAPIV2ChangesetsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAPIV2ChangesetsParamsWithHTTPClient(client *http.Client) *PostAPIV2ChangesetsParams {
	return &PostAPIV2ChangesetsParams{
		HTTPClient: client,
	}
}

/* PostAPIV2ChangesetsParams contains all the parameters to send to the API endpoint
   for the post API v2 changesets operation.

   Typically these are written to a http.Request.
*/
type PostAPIV2ChangesetsParams struct {

	/* Changeset.

	   The Changeset to create
	*/
	Changeset *models.V2controllersCreatableChangeset

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post API v2 changesets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIV2ChangesetsParams) WithDefaults() *PostAPIV2ChangesetsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post API v2 changesets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIV2ChangesetsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post API v2 changesets params
func (o *PostAPIV2ChangesetsParams) WithTimeout(timeout time.Duration) *PostAPIV2ChangesetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v2 changesets params
func (o *PostAPIV2ChangesetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v2 changesets params
func (o *PostAPIV2ChangesetsParams) WithContext(ctx context.Context) *PostAPIV2ChangesetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v2 changesets params
func (o *PostAPIV2ChangesetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v2 changesets params
func (o *PostAPIV2ChangesetsParams) WithHTTPClient(client *http.Client) *PostAPIV2ChangesetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v2 changesets params
func (o *PostAPIV2ChangesetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChangeset adds the changeset to the post API v2 changesets params
func (o *PostAPIV2ChangesetsParams) WithChangeset(changeset *models.V2controllersCreatableChangeset) *PostAPIV2ChangesetsParams {
	o.SetChangeset(changeset)
	return o
}

// SetChangeset adds the changeset to the post API v2 changesets params
func (o *PostAPIV2ChangesetsParams) SetChangeset(changeset *models.V2controllersCreatableChangeset) {
	o.Changeset = changeset
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV2ChangesetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Changeset != nil {
		if err := r.SetBodyParam(o.Changeset); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
