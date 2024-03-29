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

// NewPatchAPIClustersV3SelectorParams creates a new PatchAPIClustersV3SelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAPIClustersV3SelectorParams() *PatchAPIClustersV3SelectorParams {
	return &PatchAPIClustersV3SelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPIClustersV3SelectorParamsWithTimeout creates a new PatchAPIClustersV3SelectorParams object
// with the ability to set a timeout on a request.
func NewPatchAPIClustersV3SelectorParamsWithTimeout(timeout time.Duration) *PatchAPIClustersV3SelectorParams {
	return &PatchAPIClustersV3SelectorParams{
		timeout: timeout,
	}
}

// NewPatchAPIClustersV3SelectorParamsWithContext creates a new PatchAPIClustersV3SelectorParams object
// with the ability to set a context for a request.
func NewPatchAPIClustersV3SelectorParamsWithContext(ctx context.Context) *PatchAPIClustersV3SelectorParams {
	return &PatchAPIClustersV3SelectorParams{
		Context: ctx,
	}
}

// NewPatchAPIClustersV3SelectorParamsWithHTTPClient creates a new PatchAPIClustersV3SelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAPIClustersV3SelectorParamsWithHTTPClient(client *http.Client) *PatchAPIClustersV3SelectorParams {
	return &PatchAPIClustersV3SelectorParams{
		HTTPClient: client,
	}
}

/* PatchAPIClustersV3SelectorParams contains all the parameters to send to the API endpoint
   for the patch API clusters v3 selector operation.

   Typically these are written to a http.Request.
*/
type PatchAPIClustersV3SelectorParams struct {

	/* Cluster.

	   The edits to make to the Cluster
	*/
	Cluster *models.SherlockClusterV3Edit

	/* Selector.

	   The selector of the Cluster, which can be either a numeric ID or the name.
	*/
	Selector string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch API clusters v3 selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIClustersV3SelectorParams) WithDefaults() *PatchAPIClustersV3SelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch API clusters v3 selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIClustersV3SelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch API clusters v3 selector params
func (o *PatchAPIClustersV3SelectorParams) WithTimeout(timeout time.Duration) *PatchAPIClustersV3SelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API clusters v3 selector params
func (o *PatchAPIClustersV3SelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API clusters v3 selector params
func (o *PatchAPIClustersV3SelectorParams) WithContext(ctx context.Context) *PatchAPIClustersV3SelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API clusters v3 selector params
func (o *PatchAPIClustersV3SelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API clusters v3 selector params
func (o *PatchAPIClustersV3SelectorParams) WithHTTPClient(client *http.Client) *PatchAPIClustersV3SelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API clusters v3 selector params
func (o *PatchAPIClustersV3SelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCluster adds the cluster to the patch API clusters v3 selector params
func (o *PatchAPIClustersV3SelectorParams) WithCluster(cluster *models.SherlockClusterV3Edit) *PatchAPIClustersV3SelectorParams {
	o.SetCluster(cluster)
	return o
}

// SetCluster adds the cluster to the patch API clusters v3 selector params
func (o *PatchAPIClustersV3SelectorParams) SetCluster(cluster *models.SherlockClusterV3Edit) {
	o.Cluster = cluster
}

// WithSelector adds the selector to the patch API clusters v3 selector params
func (o *PatchAPIClustersV3SelectorParams) WithSelector(selector string) *PatchAPIClustersV3SelectorParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the patch API clusters v3 selector params
func (o *PatchAPIClustersV3SelectorParams) SetSelector(selector string) {
	o.Selector = selector
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPIClustersV3SelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
