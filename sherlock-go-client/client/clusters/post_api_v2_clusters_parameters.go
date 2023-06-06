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

// NewPostAPIV2ClustersParams creates a new PostAPIV2ClustersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAPIV2ClustersParams() *PostAPIV2ClustersParams {
	return &PostAPIV2ClustersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV2ClustersParamsWithTimeout creates a new PostAPIV2ClustersParams object
// with the ability to set a timeout on a request.
func NewPostAPIV2ClustersParamsWithTimeout(timeout time.Duration) *PostAPIV2ClustersParams {
	return &PostAPIV2ClustersParams{
		timeout: timeout,
	}
}

// NewPostAPIV2ClustersParamsWithContext creates a new PostAPIV2ClustersParams object
// with the ability to set a context for a request.
func NewPostAPIV2ClustersParamsWithContext(ctx context.Context) *PostAPIV2ClustersParams {
	return &PostAPIV2ClustersParams{
		Context: ctx,
	}
}

// NewPostAPIV2ClustersParamsWithHTTPClient creates a new PostAPIV2ClustersParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAPIV2ClustersParamsWithHTTPClient(client *http.Client) *PostAPIV2ClustersParams {
	return &PostAPIV2ClustersParams{
		HTTPClient: client,
	}
}

/* PostAPIV2ClustersParams contains all the parameters to send to the API endpoint
   for the post API v2 clusters operation.

   Typically these are written to a http.Request.
*/
type PostAPIV2ClustersParams struct {

	/* Cluster.

	   The Cluster to create
	*/
	Cluster *models.V2controllersCreatableCluster

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post API v2 clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIV2ClustersParams) WithDefaults() *PostAPIV2ClustersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post API v2 clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIV2ClustersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post API v2 clusters params
func (o *PostAPIV2ClustersParams) WithTimeout(timeout time.Duration) *PostAPIV2ClustersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v2 clusters params
func (o *PostAPIV2ClustersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v2 clusters params
func (o *PostAPIV2ClustersParams) WithContext(ctx context.Context) *PostAPIV2ClustersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v2 clusters params
func (o *PostAPIV2ClustersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v2 clusters params
func (o *PostAPIV2ClustersParams) WithHTTPClient(client *http.Client) *PostAPIV2ClustersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v2 clusters params
func (o *PostAPIV2ClustersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCluster adds the cluster to the post API v2 clusters params
func (o *PostAPIV2ClustersParams) WithCluster(cluster *models.V2controllersCreatableCluster) *PostAPIV2ClustersParams {
	o.SetCluster(cluster)
	return o
}

// SetCluster adds the cluster to the post API v2 clusters params
func (o *PostAPIV2ClustersParams) SetCluster(cluster *models.V2controllersCreatableCluster) {
	o.Cluster = cluster
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV2ClustersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Cluster != nil {
		if err := r.SetBodyParam(o.Cluster); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
