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

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// NewPutAPIV2CiIdentifiersSelectorParams creates a new PutAPIV2CiIdentifiersSelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutAPIV2CiIdentifiersSelectorParams() *PutAPIV2CiIdentifiersSelectorParams {
	return &PutAPIV2CiIdentifiersSelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutAPIV2CiIdentifiersSelectorParamsWithTimeout creates a new PutAPIV2CiIdentifiersSelectorParams object
// with the ability to set a timeout on a request.
func NewPutAPIV2CiIdentifiersSelectorParamsWithTimeout(timeout time.Duration) *PutAPIV2CiIdentifiersSelectorParams {
	return &PutAPIV2CiIdentifiersSelectorParams{
		timeout: timeout,
	}
}

// NewPutAPIV2CiIdentifiersSelectorParamsWithContext creates a new PutAPIV2CiIdentifiersSelectorParams object
// with the ability to set a context for a request.
func NewPutAPIV2CiIdentifiersSelectorParamsWithContext(ctx context.Context) *PutAPIV2CiIdentifiersSelectorParams {
	return &PutAPIV2CiIdentifiersSelectorParams{
		Context: ctx,
	}
}

// NewPutAPIV2CiIdentifiersSelectorParamsWithHTTPClient creates a new PutAPIV2CiIdentifiersSelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutAPIV2CiIdentifiersSelectorParamsWithHTTPClient(client *http.Client) *PutAPIV2CiIdentifiersSelectorParams {
	return &PutAPIV2CiIdentifiersSelectorParams{
		HTTPClient: client,
	}
}

/* PutAPIV2CiIdentifiersSelectorParams contains all the parameters to send to the API endpoint
   for the put API v2 ci identifiers selector operation.

   Typically these are written to a http.Request.
*/
type PutAPIV2CiIdentifiersSelectorParams struct {

	/* CiIdentifier.

	   The CiIdentifier to upsert
	*/
	CiIdentifier *models.V2controllersCreatableCiIdentifier

	/* Selector.

	   The CiIdentifier to upsert's selector: ID or type + '/' + selector of the referenced type
	*/
	Selector string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put API v2 ci identifiers selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutAPIV2CiIdentifiersSelectorParams) WithDefaults() *PutAPIV2CiIdentifiersSelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put API v2 ci identifiers selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutAPIV2CiIdentifiersSelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put API v2 ci identifiers selector params
func (o *PutAPIV2CiIdentifiersSelectorParams) WithTimeout(timeout time.Duration) *PutAPIV2CiIdentifiersSelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put API v2 ci identifiers selector params
func (o *PutAPIV2CiIdentifiersSelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put API v2 ci identifiers selector params
func (o *PutAPIV2CiIdentifiersSelectorParams) WithContext(ctx context.Context) *PutAPIV2CiIdentifiersSelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put API v2 ci identifiers selector params
func (o *PutAPIV2CiIdentifiersSelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put API v2 ci identifiers selector params
func (o *PutAPIV2CiIdentifiersSelectorParams) WithHTTPClient(client *http.Client) *PutAPIV2CiIdentifiersSelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put API v2 ci identifiers selector params
func (o *PutAPIV2CiIdentifiersSelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCiIdentifier adds the ciIdentifier to the put API v2 ci identifiers selector params
func (o *PutAPIV2CiIdentifiersSelectorParams) WithCiIdentifier(ciIdentifier *models.V2controllersCreatableCiIdentifier) *PutAPIV2CiIdentifiersSelectorParams {
	o.SetCiIdentifier(ciIdentifier)
	return o
}

// SetCiIdentifier adds the ciIdentifier to the put API v2 ci identifiers selector params
func (o *PutAPIV2CiIdentifiersSelectorParams) SetCiIdentifier(ciIdentifier *models.V2controllersCreatableCiIdentifier) {
	o.CiIdentifier = ciIdentifier
}

// WithSelector adds the selector to the put API v2 ci identifiers selector params
func (o *PutAPIV2CiIdentifiersSelectorParams) WithSelector(selector string) *PutAPIV2CiIdentifiersSelectorParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the put API v2 ci identifiers selector params
func (o *PutAPIV2CiIdentifiersSelectorParams) SetSelector(selector string) {
	o.Selector = selector
}

// WriteToRequest writes these params to a swagger request
func (o *PutAPIV2CiIdentifiersSelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.CiIdentifier != nil {
		if err := r.SetBodyParam(o.CiIdentifier); err != nil {
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