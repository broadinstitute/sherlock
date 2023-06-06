// Code generated by go-swagger; DO NOT EDIT.

package clusters

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

// NewPatchAPIV2ClustersSelectorParams creates a new PatchAPIV2ClustersSelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAPIV2ClustersSelectorParams() *PatchAPIV2ClustersSelectorParams {
	return &PatchAPIV2ClustersSelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPIV2ClustersSelectorParamsWithTimeout creates a new PatchAPIV2ClustersSelectorParams object
// with the ability to set a timeout on a request.
func NewPatchAPIV2ClustersSelectorParamsWithTimeout(timeout time.Duration) *PatchAPIV2ClustersSelectorParams {
	return &PatchAPIV2ClustersSelectorParams{
		timeout: timeout,
	}
}

// NewPatchAPIV2ClustersSelectorParamsWithContext creates a new PatchAPIV2ClustersSelectorParams object
// with the ability to set a context for a request.
func NewPatchAPIV2ClustersSelectorParamsWithContext(ctx context.Context) *PatchAPIV2ClustersSelectorParams {
	return &PatchAPIV2ClustersSelectorParams{
		Context: ctx,
	}
}

// NewPatchAPIV2ClustersSelectorParamsWithHTTPClient creates a new PatchAPIV2ClustersSelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAPIV2ClustersSelectorParamsWithHTTPClient(client *http.Client) *PatchAPIV2ClustersSelectorParams {
	return &PatchAPIV2ClustersSelectorParams{
		HTTPClient: client,
	}
}

/* PatchAPIV2ClustersSelectorParams contains all the parameters to send to the API endpoint
   for the patch API v2 clusters selector operation.

   Typically these are written to a http.Request.
*/
type PatchAPIV2ClustersSelectorParams struct {

	/* Cluster.

	   The edits to make to the Cluster
	*/
	Cluster *models.V2controllersEditableCluster

	/* Selector.

	   The Cluster to edit's selector: name or numeric ID
	*/
	Selector string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch API v2 clusters selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIV2ClustersSelectorParams) WithDefaults() *PatchAPIV2ClustersSelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch API v2 clusters selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIV2ClustersSelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch API v2 clusters selector params
func (o *PatchAPIV2ClustersSelectorParams) WithTimeout(timeout time.Duration) *PatchAPIV2ClustersSelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API v2 clusters selector params
func (o *PatchAPIV2ClustersSelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API v2 clusters selector params
func (o *PatchAPIV2ClustersSelectorParams) WithContext(ctx context.Context) *PatchAPIV2ClustersSelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API v2 clusters selector params
func (o *PatchAPIV2ClustersSelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API v2 clusters selector params
func (o *PatchAPIV2ClustersSelectorParams) WithHTTPClient(client *http.Client) *PatchAPIV2ClustersSelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API v2 clusters selector params
func (o *PatchAPIV2ClustersSelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCluster adds the cluster to the patch API v2 clusters selector params
func (o *PatchAPIV2ClustersSelectorParams) WithCluster(cluster *models.V2controllersEditableCluster) *PatchAPIV2ClustersSelectorParams {
	o.SetCluster(cluster)
	return o
}

// SetCluster adds the cluster to the patch API v2 clusters selector params
func (o *PatchAPIV2ClustersSelectorParams) SetCluster(cluster *models.V2controllersEditableCluster) {
	o.Cluster = cluster
}

// WithSelector adds the selector to the patch API v2 clusters selector params
func (o *PatchAPIV2ClustersSelectorParams) WithSelector(selector string) *PatchAPIV2ClustersSelectorParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the patch API v2 clusters selector params
func (o *PatchAPIV2ClustersSelectorParams) SetSelector(selector string) {
	o.Selector = selector
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPIV2ClustersSelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Cluster != nil {
		if err := r.SetBodyParam(o.Cluster); err != nil {
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
