// Code generated by go-swagger; DO NOT EDIT.

package chart_releases

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
	"github.com/go-openapi/swag"
)

// NewGetAPIV2ChartReleasesParams creates a new GetAPIV2ChartReleasesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIV2ChartReleasesParams() *GetAPIV2ChartReleasesParams {
	return &GetAPIV2ChartReleasesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV2ChartReleasesParamsWithTimeout creates a new GetAPIV2ChartReleasesParams object
// with the ability to set a timeout on a request.
func NewGetAPIV2ChartReleasesParamsWithTimeout(timeout time.Duration) *GetAPIV2ChartReleasesParams {
	return &GetAPIV2ChartReleasesParams{
		timeout: timeout,
	}
}

// NewGetAPIV2ChartReleasesParamsWithContext creates a new GetAPIV2ChartReleasesParams object
// with the ability to set a context for a request.
func NewGetAPIV2ChartReleasesParamsWithContext(ctx context.Context) *GetAPIV2ChartReleasesParams {
	return &GetAPIV2ChartReleasesParams{
		Context: ctx,
	}
}

// NewGetAPIV2ChartReleasesParamsWithHTTPClient creates a new GetAPIV2ChartReleasesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIV2ChartReleasesParamsWithHTTPClient(client *http.Client) *GetAPIV2ChartReleasesParams {
	return &GetAPIV2ChartReleasesParams{
		HTTPClient: client,
	}
}

/* GetAPIV2ChartReleasesParams contains all the parameters to send to the API endpoint
   for the get API v2 chart releases operation.

   Typically these are written to a http.Request.
*/
type GetAPIV2ChartReleasesParams struct {

	/* Chart.

	   Required when creating
	*/
	Chart *string

	/* Cluster.

	   When creating, will default the environment's default cluster, if provided. Either this or environment must be provided.
	*/
	Cluster *string

	// CreatedAt.
	CreatedAt *string

	// CurrentAppVersionExact.
	CurrentAppVersionExact *string

	// CurrentChartVersionExact.
	CurrentChartVersionExact *string

	/* DestinationType.

	   Calculated field
	*/
	DestinationType *string

	/* Environment.

	   Either this or cluster must be provided.
	*/
	Environment *string

	// HelmfileRef.
	//
	// Default: "HEAD"
	HelmfileRef *string

	// ID.
	ID *int64

	/* Limit.

	   An optional limit to the number of entries returned
	*/
	Limit *int64

	/* Name.

	   When creating, will be calculated if left empty
	*/
	Name *string

	/* Namespace.

	   When creating, will default to the environment's default namespace, if provided
	*/
	Namespace *string

	/* TargetAppVersionBranch.

	   When creating, will default to the app's main branch if it has one recorded
	*/
	TargetAppVersionBranch *string

	// TargetAppVersionCommit.
	TargetAppVersionCommit *string

	// TargetAppVersionExact.
	TargetAppVersionExact *string

	/* TargetAppVersionUse.

	   When creating, will default to referencing any provided target app version field (exact, then commit, then branch)
	*/
	TargetAppVersionUse *string

	// TargetChartVersionExact.
	TargetChartVersionExact *string

	/* TargetChartVersionUse.

	   When creating, will default to latest unless an exact target chart version is provided
	*/
	TargetChartVersionUse *string

	// ThelmaMode.
	ThelmaMode *string

	// UpdatedAt.
	UpdatedAt *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API v2 chart releases params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2ChartReleasesParams) WithDefaults() *GetAPIV2ChartReleasesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API v2 chart releases params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2ChartReleasesParams) SetDefaults() {
	var (
		helmfileRefDefault = string("HEAD")
	)

	val := GetAPIV2ChartReleasesParams{
		HelmfileRef: &helmfileRefDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) WithTimeout(timeout time.Duration) *GetAPIV2ChartReleasesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) WithContext(ctx context.Context) *GetAPIV2ChartReleasesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) WithHTTPClient(client *http.Client) *GetAPIV2ChartReleasesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChart adds the chart to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) WithChart(chart *string) *GetAPIV2ChartReleasesParams {
	o.SetChart(chart)
	return o
}

// SetChart adds the chart to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) SetChart(chart *string) {
	o.Chart = chart
}

// WithCluster adds the cluster to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) WithCluster(cluster *string) *GetAPIV2ChartReleasesParams {
	o.SetCluster(cluster)
	return o
}

// SetCluster adds the cluster to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) SetCluster(cluster *string) {
	o.Cluster = cluster
}

// WithCreatedAt adds the createdAt to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) WithCreatedAt(createdAt *string) *GetAPIV2ChartReleasesParams {
	o.SetCreatedAt(createdAt)
	return o
}

// SetCreatedAt adds the createdAt to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) SetCreatedAt(createdAt *string) {
	o.CreatedAt = createdAt
}

// WithCurrentAppVersionExact adds the currentAppVersionExact to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) WithCurrentAppVersionExact(currentAppVersionExact *string) *GetAPIV2ChartReleasesParams {
	o.SetCurrentAppVersionExact(currentAppVersionExact)
	return o
}

// SetCurrentAppVersionExact adds the currentAppVersionExact to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) SetCurrentAppVersionExact(currentAppVersionExact *string) {
	o.CurrentAppVersionExact = currentAppVersionExact
}

// WithCurrentChartVersionExact adds the currentChartVersionExact to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) WithCurrentChartVersionExact(currentChartVersionExact *string) *GetAPIV2ChartReleasesParams {
	o.SetCurrentChartVersionExact(currentChartVersionExact)
	return o
}

// SetCurrentChartVersionExact adds the currentChartVersionExact to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) SetCurrentChartVersionExact(currentChartVersionExact *string) {
	o.CurrentChartVersionExact = currentChartVersionExact
}

// WithDestinationType adds the destinationType to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) WithDestinationType(destinationType *string) *GetAPIV2ChartReleasesParams {
	o.SetDestinationType(destinationType)
	return o
}

// SetDestinationType adds the destinationType to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) SetDestinationType(destinationType *string) {
	o.DestinationType = destinationType
}

// WithEnvironment adds the environment to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) WithEnvironment(environment *string) *GetAPIV2ChartReleasesParams {
	o.SetEnvironment(environment)
	return o
}

// SetEnvironment adds the environment to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) SetEnvironment(environment *string) {
	o.Environment = environment
}

// WithHelmfileRef adds the helmfileRef to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) WithHelmfileRef(helmfileRef *string) *GetAPIV2ChartReleasesParams {
	o.SetHelmfileRef(helmfileRef)
	return o
}

// SetHelmfileRef adds the helmfileRef to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) SetHelmfileRef(helmfileRef *string) {
	o.HelmfileRef = helmfileRef
}

// WithID adds the id to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) WithID(id *int64) *GetAPIV2ChartReleasesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) SetID(id *int64) {
	o.ID = id
}

// WithLimit adds the limit to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) WithLimit(limit *int64) *GetAPIV2ChartReleasesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) WithName(name *string) *GetAPIV2ChartReleasesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) SetName(name *string) {
	o.Name = name
}

// WithNamespace adds the namespace to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) WithNamespace(namespace *string) *GetAPIV2ChartReleasesParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) SetNamespace(namespace *string) {
	o.Namespace = namespace
}

// WithTargetAppVersionBranch adds the targetAppVersionBranch to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) WithTargetAppVersionBranch(targetAppVersionBranch *string) *GetAPIV2ChartReleasesParams {
	o.SetTargetAppVersionBranch(targetAppVersionBranch)
	return o
}

// SetTargetAppVersionBranch adds the targetAppVersionBranch to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) SetTargetAppVersionBranch(targetAppVersionBranch *string) {
	o.TargetAppVersionBranch = targetAppVersionBranch
}

// WithTargetAppVersionCommit adds the targetAppVersionCommit to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) WithTargetAppVersionCommit(targetAppVersionCommit *string) *GetAPIV2ChartReleasesParams {
	o.SetTargetAppVersionCommit(targetAppVersionCommit)
	return o
}

// SetTargetAppVersionCommit adds the targetAppVersionCommit to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) SetTargetAppVersionCommit(targetAppVersionCommit *string) {
	o.TargetAppVersionCommit = targetAppVersionCommit
}

// WithTargetAppVersionExact adds the targetAppVersionExact to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) WithTargetAppVersionExact(targetAppVersionExact *string) *GetAPIV2ChartReleasesParams {
	o.SetTargetAppVersionExact(targetAppVersionExact)
	return o
}

// SetTargetAppVersionExact adds the targetAppVersionExact to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) SetTargetAppVersionExact(targetAppVersionExact *string) {
	o.TargetAppVersionExact = targetAppVersionExact
}

// WithTargetAppVersionUse adds the targetAppVersionUse to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) WithTargetAppVersionUse(targetAppVersionUse *string) *GetAPIV2ChartReleasesParams {
	o.SetTargetAppVersionUse(targetAppVersionUse)
	return o
}

// SetTargetAppVersionUse adds the targetAppVersionUse to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) SetTargetAppVersionUse(targetAppVersionUse *string) {
	o.TargetAppVersionUse = targetAppVersionUse
}

// WithTargetChartVersionExact adds the targetChartVersionExact to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) WithTargetChartVersionExact(targetChartVersionExact *string) *GetAPIV2ChartReleasesParams {
	o.SetTargetChartVersionExact(targetChartVersionExact)
	return o
}

// SetTargetChartVersionExact adds the targetChartVersionExact to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) SetTargetChartVersionExact(targetChartVersionExact *string) {
	o.TargetChartVersionExact = targetChartVersionExact
}

// WithTargetChartVersionUse adds the targetChartVersionUse to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) WithTargetChartVersionUse(targetChartVersionUse *string) *GetAPIV2ChartReleasesParams {
	o.SetTargetChartVersionUse(targetChartVersionUse)
	return o
}

// SetTargetChartVersionUse adds the targetChartVersionUse to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) SetTargetChartVersionUse(targetChartVersionUse *string) {
	o.TargetChartVersionUse = targetChartVersionUse
}

// WithThelmaMode adds the thelmaMode to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) WithThelmaMode(thelmaMode *string) *GetAPIV2ChartReleasesParams {
	o.SetThelmaMode(thelmaMode)
	return o
}

// SetThelmaMode adds the thelmaMode to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) SetThelmaMode(thelmaMode *string) {
	o.ThelmaMode = thelmaMode
}

// WithUpdatedAt adds the updatedAt to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) WithUpdatedAt(updatedAt *string) *GetAPIV2ChartReleasesParams {
	o.SetUpdatedAt(updatedAt)
	return o
}

// SetUpdatedAt adds the updatedAt to the get API v2 chart releases params
func (o *GetAPIV2ChartReleasesParams) SetUpdatedAt(updatedAt *string) {
	o.UpdatedAt = updatedAt
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV2ChartReleasesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Chart != nil {

		// query param chart
		var qrChart string

		if o.Chart != nil {
			qrChart = *o.Chart
		}
		qChart := qrChart
		if qChart != "" {

			if err := r.SetQueryParam("chart", qChart); err != nil {
				return err
			}
		}
	}

	if o.Cluster != nil {

		// query param cluster
		var qrCluster string

		if o.Cluster != nil {
			qrCluster = *o.Cluster
		}
		qCluster := qrCluster
		if qCluster != "" {

			if err := r.SetQueryParam("cluster", qCluster); err != nil {
				return err
			}
		}
	}

	if o.CreatedAt != nil {

		// query param createdAt
		var qrCreatedAt string

		if o.CreatedAt != nil {
			qrCreatedAt = *o.CreatedAt
		}
		qCreatedAt := qrCreatedAt
		if qCreatedAt != "" {

			if err := r.SetQueryParam("createdAt", qCreatedAt); err != nil {
				return err
			}
		}
	}

	if o.CurrentAppVersionExact != nil {

		// query param currentAppVersionExact
		var qrCurrentAppVersionExact string

		if o.CurrentAppVersionExact != nil {
			qrCurrentAppVersionExact = *o.CurrentAppVersionExact
		}
		qCurrentAppVersionExact := qrCurrentAppVersionExact
		if qCurrentAppVersionExact != "" {

			if err := r.SetQueryParam("currentAppVersionExact", qCurrentAppVersionExact); err != nil {
				return err
			}
		}
	}

	if o.CurrentChartVersionExact != nil {

		// query param currentChartVersionExact
		var qrCurrentChartVersionExact string

		if o.CurrentChartVersionExact != nil {
			qrCurrentChartVersionExact = *o.CurrentChartVersionExact
		}
		qCurrentChartVersionExact := qrCurrentChartVersionExact
		if qCurrentChartVersionExact != "" {

			if err := r.SetQueryParam("currentChartVersionExact", qCurrentChartVersionExact); err != nil {
				return err
			}
		}
	}

	if o.DestinationType != nil {

		// query param destinationType
		var qrDestinationType string

		if o.DestinationType != nil {
			qrDestinationType = *o.DestinationType
		}
		qDestinationType := qrDestinationType
		if qDestinationType != "" {

			if err := r.SetQueryParam("destinationType", qDestinationType); err != nil {
				return err
			}
		}
	}

	if o.Environment != nil {

		// query param environment
		var qrEnvironment string

		if o.Environment != nil {
			qrEnvironment = *o.Environment
		}
		qEnvironment := qrEnvironment
		if qEnvironment != "" {

			if err := r.SetQueryParam("environment", qEnvironment); err != nil {
				return err
			}
		}
	}

	if o.HelmfileRef != nil {

		// query param helmfileRef
		var qrHelmfileRef string

		if o.HelmfileRef != nil {
			qrHelmfileRef = *o.HelmfileRef
		}
		qHelmfileRef := qrHelmfileRef
		if qHelmfileRef != "" {

			if err := r.SetQueryParam("helmfileRef", qHelmfileRef); err != nil {
				return err
			}
		}
	}

	if o.ID != nil {

		// query param id
		var qrID int64

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := swag.FormatInt64(qrID)
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.Namespace != nil {

		// query param namespace
		var qrNamespace string

		if o.Namespace != nil {
			qrNamespace = *o.Namespace
		}
		qNamespace := qrNamespace
		if qNamespace != "" {

			if err := r.SetQueryParam("namespace", qNamespace); err != nil {
				return err
			}
		}
	}

	if o.TargetAppVersionBranch != nil {

		// query param targetAppVersionBranch
		var qrTargetAppVersionBranch string

		if o.TargetAppVersionBranch != nil {
			qrTargetAppVersionBranch = *o.TargetAppVersionBranch
		}
		qTargetAppVersionBranch := qrTargetAppVersionBranch
		if qTargetAppVersionBranch != "" {

			if err := r.SetQueryParam("targetAppVersionBranch", qTargetAppVersionBranch); err != nil {
				return err
			}
		}
	}

	if o.TargetAppVersionCommit != nil {

		// query param targetAppVersionCommit
		var qrTargetAppVersionCommit string

		if o.TargetAppVersionCommit != nil {
			qrTargetAppVersionCommit = *o.TargetAppVersionCommit
		}
		qTargetAppVersionCommit := qrTargetAppVersionCommit
		if qTargetAppVersionCommit != "" {

			if err := r.SetQueryParam("targetAppVersionCommit", qTargetAppVersionCommit); err != nil {
				return err
			}
		}
	}

	if o.TargetAppVersionExact != nil {

		// query param targetAppVersionExact
		var qrTargetAppVersionExact string

		if o.TargetAppVersionExact != nil {
			qrTargetAppVersionExact = *o.TargetAppVersionExact
		}
		qTargetAppVersionExact := qrTargetAppVersionExact
		if qTargetAppVersionExact != "" {

			if err := r.SetQueryParam("targetAppVersionExact", qTargetAppVersionExact); err != nil {
				return err
			}
		}
	}

	if o.TargetAppVersionUse != nil {

		// query param targetAppVersionUse
		var qrTargetAppVersionUse string

		if o.TargetAppVersionUse != nil {
			qrTargetAppVersionUse = *o.TargetAppVersionUse
		}
		qTargetAppVersionUse := qrTargetAppVersionUse
		if qTargetAppVersionUse != "" {

			if err := r.SetQueryParam("targetAppVersionUse", qTargetAppVersionUse); err != nil {
				return err
			}
		}
	}

	if o.TargetChartVersionExact != nil {

		// query param targetChartVersionExact
		var qrTargetChartVersionExact string

		if o.TargetChartVersionExact != nil {
			qrTargetChartVersionExact = *o.TargetChartVersionExact
		}
		qTargetChartVersionExact := qrTargetChartVersionExact
		if qTargetChartVersionExact != "" {

			if err := r.SetQueryParam("targetChartVersionExact", qTargetChartVersionExact); err != nil {
				return err
			}
		}
	}

	if o.TargetChartVersionUse != nil {

		// query param targetChartVersionUse
		var qrTargetChartVersionUse string

		if o.TargetChartVersionUse != nil {
			qrTargetChartVersionUse = *o.TargetChartVersionUse
		}
		qTargetChartVersionUse := qrTargetChartVersionUse
		if qTargetChartVersionUse != "" {

			if err := r.SetQueryParam("targetChartVersionUse", qTargetChartVersionUse); err != nil {
				return err
			}
		}
	}

	if o.ThelmaMode != nil {

		// query param thelmaMode
		var qrThelmaMode string

		if o.ThelmaMode != nil {
			qrThelmaMode = *o.ThelmaMode
		}
		qThelmaMode := qrThelmaMode
		if qThelmaMode != "" {

			if err := r.SetQueryParam("thelmaMode", qThelmaMode); err != nil {
				return err
			}
		}
	}

	if o.UpdatedAt != nil {

		// query param updatedAt
		var qrUpdatedAt string

		if o.UpdatedAt != nil {
			qrUpdatedAt = *o.UpdatedAt
		}
		qUpdatedAt := qrUpdatedAt
		if qUpdatedAt != "" {

			if err := r.SetQueryParam("updatedAt", qUpdatedAt); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
