// Code generated by go-swagger; DO NOT EDIT.

package charts

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

// NewGetAPIV2ChartsParams creates a new GetAPIV2ChartsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIV2ChartsParams() *GetAPIV2ChartsParams {
	return &GetAPIV2ChartsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV2ChartsParamsWithTimeout creates a new GetAPIV2ChartsParams object
// with the ability to set a timeout on a request.
func NewGetAPIV2ChartsParamsWithTimeout(timeout time.Duration) *GetAPIV2ChartsParams {
	return &GetAPIV2ChartsParams{
		timeout: timeout,
	}
}

// NewGetAPIV2ChartsParamsWithContext creates a new GetAPIV2ChartsParams object
// with the ability to set a context for a request.
func NewGetAPIV2ChartsParamsWithContext(ctx context.Context) *GetAPIV2ChartsParams {
	return &GetAPIV2ChartsParams{
		Context: ctx,
	}
}

// NewGetAPIV2ChartsParamsWithHTTPClient creates a new GetAPIV2ChartsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIV2ChartsParamsWithHTTPClient(client *http.Client) *GetAPIV2ChartsParams {
	return &GetAPIV2ChartsParams{
		HTTPClient: client,
	}
}

/* GetAPIV2ChartsParams contains all the parameters to send to the API endpoint
   for the get API v2 charts operation.

   Typically these are written to a http.Request.
*/
type GetAPIV2ChartsParams struct {

	// AppImageGitMainBranch.
	AppImageGitMainBranch *string

	// AppImageGitRepo.
	AppImageGitRepo *string

	/* ChartExposesEndpoint.

	   Indicates if the default subdomain, protocol, and port fields are relevant for this chart
	*/
	ChartExposesEndpoint *bool

	// ChartRepo.
	//
	// Default: "terra-helm"
	ChartRepo *string

	// CreatedAt.
	//
	// Format: date-time
	CreatedAt *strfmt.DateTime

	// DefaultPort.
	//
	// Default: 443
	DefaultPort *int64

	// DefaultProtocol.
	//
	// Default: "https"
	DefaultProtocol *string

	/* DefaultSubdomain.

	   When creating, will default to the name of the chart
	*/
	DefaultSubdomain *string

	// ID.
	ID *int64

	/* LegacyConfigsEnbled.

	   Indicates whether a chart requires config rendering from firecloud-develop
	*/
	LegacyConfigsEnbled *bool

	/* Limit.

	   An optional limit to the number of entries returned
	*/
	Limit *int64

	/* Name.

	   Required when creating
	*/
	Name *string

	// UpdatedAt.
	//
	// Format: date-time
	UpdatedAt *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API v2 charts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2ChartsParams) WithDefaults() *GetAPIV2ChartsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API v2 charts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2ChartsParams) SetDefaults() {
	var (
		chartExposesEndpointDefault = bool(false)

		chartRepoDefault = string("terra-helm")

		defaultPortDefault = int64(443)

		defaultProtocolDefault = string("https")

		legacyConfigsEnbledDefault = bool(false)
	)

	val := GetAPIV2ChartsParams{
		ChartExposesEndpoint: &chartExposesEndpointDefault,
		ChartRepo:            &chartRepoDefault,
		DefaultPort:          &defaultPortDefault,
		DefaultProtocol:      &defaultProtocolDefault,
		LegacyConfigsEnbled:  &legacyConfigsEnbledDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get API v2 charts params
func (o *GetAPIV2ChartsParams) WithTimeout(timeout time.Duration) *GetAPIV2ChartsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v2 charts params
func (o *GetAPIV2ChartsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v2 charts params
func (o *GetAPIV2ChartsParams) WithContext(ctx context.Context) *GetAPIV2ChartsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v2 charts params
func (o *GetAPIV2ChartsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v2 charts params
func (o *GetAPIV2ChartsParams) WithHTTPClient(client *http.Client) *GetAPIV2ChartsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v2 charts params
func (o *GetAPIV2ChartsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppImageGitMainBranch adds the appImageGitMainBranch to the get API v2 charts params
func (o *GetAPIV2ChartsParams) WithAppImageGitMainBranch(appImageGitMainBranch *string) *GetAPIV2ChartsParams {
	o.SetAppImageGitMainBranch(appImageGitMainBranch)
	return o
}

// SetAppImageGitMainBranch adds the appImageGitMainBranch to the get API v2 charts params
func (o *GetAPIV2ChartsParams) SetAppImageGitMainBranch(appImageGitMainBranch *string) {
	o.AppImageGitMainBranch = appImageGitMainBranch
}

// WithAppImageGitRepo adds the appImageGitRepo to the get API v2 charts params
func (o *GetAPIV2ChartsParams) WithAppImageGitRepo(appImageGitRepo *string) *GetAPIV2ChartsParams {
	o.SetAppImageGitRepo(appImageGitRepo)
	return o
}

// SetAppImageGitRepo adds the appImageGitRepo to the get API v2 charts params
func (o *GetAPIV2ChartsParams) SetAppImageGitRepo(appImageGitRepo *string) {
	o.AppImageGitRepo = appImageGitRepo
}

// WithChartExposesEndpoint adds the chartExposesEndpoint to the get API v2 charts params
func (o *GetAPIV2ChartsParams) WithChartExposesEndpoint(chartExposesEndpoint *bool) *GetAPIV2ChartsParams {
	o.SetChartExposesEndpoint(chartExposesEndpoint)
	return o
}

// SetChartExposesEndpoint adds the chartExposesEndpoint to the get API v2 charts params
func (o *GetAPIV2ChartsParams) SetChartExposesEndpoint(chartExposesEndpoint *bool) {
	o.ChartExposesEndpoint = chartExposesEndpoint
}

// WithChartRepo adds the chartRepo to the get API v2 charts params
func (o *GetAPIV2ChartsParams) WithChartRepo(chartRepo *string) *GetAPIV2ChartsParams {
	o.SetChartRepo(chartRepo)
	return o
}

// SetChartRepo adds the chartRepo to the get API v2 charts params
func (o *GetAPIV2ChartsParams) SetChartRepo(chartRepo *string) {
	o.ChartRepo = chartRepo
}

// WithCreatedAt adds the createdAt to the get API v2 charts params
func (o *GetAPIV2ChartsParams) WithCreatedAt(createdAt *strfmt.DateTime) *GetAPIV2ChartsParams {
	o.SetCreatedAt(createdAt)
	return o
}

// SetCreatedAt adds the createdAt to the get API v2 charts params
func (o *GetAPIV2ChartsParams) SetCreatedAt(createdAt *strfmt.DateTime) {
	o.CreatedAt = createdAt
}

// WithDefaultPort adds the defaultPort to the get API v2 charts params
func (o *GetAPIV2ChartsParams) WithDefaultPort(defaultPort *int64) *GetAPIV2ChartsParams {
	o.SetDefaultPort(defaultPort)
	return o
}

// SetDefaultPort adds the defaultPort to the get API v2 charts params
func (o *GetAPIV2ChartsParams) SetDefaultPort(defaultPort *int64) {
	o.DefaultPort = defaultPort
}

// WithDefaultProtocol adds the defaultProtocol to the get API v2 charts params
func (o *GetAPIV2ChartsParams) WithDefaultProtocol(defaultProtocol *string) *GetAPIV2ChartsParams {
	o.SetDefaultProtocol(defaultProtocol)
	return o
}

// SetDefaultProtocol adds the defaultProtocol to the get API v2 charts params
func (o *GetAPIV2ChartsParams) SetDefaultProtocol(defaultProtocol *string) {
	o.DefaultProtocol = defaultProtocol
}

// WithDefaultSubdomain adds the defaultSubdomain to the get API v2 charts params
func (o *GetAPIV2ChartsParams) WithDefaultSubdomain(defaultSubdomain *string) *GetAPIV2ChartsParams {
	o.SetDefaultSubdomain(defaultSubdomain)
	return o
}

// SetDefaultSubdomain adds the defaultSubdomain to the get API v2 charts params
func (o *GetAPIV2ChartsParams) SetDefaultSubdomain(defaultSubdomain *string) {
	o.DefaultSubdomain = defaultSubdomain
}

// WithID adds the id to the get API v2 charts params
func (o *GetAPIV2ChartsParams) WithID(id *int64) *GetAPIV2ChartsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get API v2 charts params
func (o *GetAPIV2ChartsParams) SetID(id *int64) {
	o.ID = id
}

// WithLegacyConfigsEnbled adds the legacyConfigsEnbled to the get API v2 charts params
func (o *GetAPIV2ChartsParams) WithLegacyConfigsEnbled(legacyConfigsEnbled *bool) *GetAPIV2ChartsParams {
	o.SetLegacyConfigsEnbled(legacyConfigsEnbled)
	return o
}

// SetLegacyConfigsEnbled adds the legacyConfigsEnbled to the get API v2 charts params
func (o *GetAPIV2ChartsParams) SetLegacyConfigsEnbled(legacyConfigsEnbled *bool) {
	o.LegacyConfigsEnbled = legacyConfigsEnbled
}

// WithLimit adds the limit to the get API v2 charts params
func (o *GetAPIV2ChartsParams) WithLimit(limit *int64) *GetAPIV2ChartsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get API v2 charts params
func (o *GetAPIV2ChartsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the get API v2 charts params
func (o *GetAPIV2ChartsParams) WithName(name *string) *GetAPIV2ChartsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get API v2 charts params
func (o *GetAPIV2ChartsParams) SetName(name *string) {
	o.Name = name
}

// WithUpdatedAt adds the updatedAt to the get API v2 charts params
func (o *GetAPIV2ChartsParams) WithUpdatedAt(updatedAt *strfmt.DateTime) *GetAPIV2ChartsParams {
	o.SetUpdatedAt(updatedAt)
	return o
}

// SetUpdatedAt adds the updatedAt to the get API v2 charts params
func (o *GetAPIV2ChartsParams) SetUpdatedAt(updatedAt *strfmt.DateTime) {
	o.UpdatedAt = updatedAt
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV2ChartsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AppImageGitMainBranch != nil {

		// query param appImageGitMainBranch
		var qrAppImageGitMainBranch string

		if o.AppImageGitMainBranch != nil {
			qrAppImageGitMainBranch = *o.AppImageGitMainBranch
		}
		qAppImageGitMainBranch := qrAppImageGitMainBranch
		if qAppImageGitMainBranch != "" {

			if err := r.SetQueryParam("appImageGitMainBranch", qAppImageGitMainBranch); err != nil {
				return err
			}
		}
	}

	if o.AppImageGitRepo != nil {

		// query param appImageGitRepo
		var qrAppImageGitRepo string

		if o.AppImageGitRepo != nil {
			qrAppImageGitRepo = *o.AppImageGitRepo
		}
		qAppImageGitRepo := qrAppImageGitRepo
		if qAppImageGitRepo != "" {

			if err := r.SetQueryParam("appImageGitRepo", qAppImageGitRepo); err != nil {
				return err
			}
		}
	}

	if o.ChartExposesEndpoint != nil {

		// query param chartExposesEndpoint
		var qrChartExposesEndpoint bool

		if o.ChartExposesEndpoint != nil {
			qrChartExposesEndpoint = *o.ChartExposesEndpoint
		}
		qChartExposesEndpoint := swag.FormatBool(qrChartExposesEndpoint)
		if qChartExposesEndpoint != "" {

			if err := r.SetQueryParam("chartExposesEndpoint", qChartExposesEndpoint); err != nil {
				return err
			}
		}
	}

	if o.ChartRepo != nil {

		// query param chartRepo
		var qrChartRepo string

		if o.ChartRepo != nil {
			qrChartRepo = *o.ChartRepo
		}
		qChartRepo := qrChartRepo
		if qChartRepo != "" {

			if err := r.SetQueryParam("chartRepo", qChartRepo); err != nil {
				return err
			}
		}
	}

	if o.CreatedAt != nil {

		// query param createdAt
		var qrCreatedAt strfmt.DateTime

		if o.CreatedAt != nil {
			qrCreatedAt = *o.CreatedAt
		}
		qCreatedAt := qrCreatedAt.String()
		if qCreatedAt != "" {

			if err := r.SetQueryParam("createdAt", qCreatedAt); err != nil {
				return err
			}
		}
	}

	if o.DefaultPort != nil {

		// query param defaultPort
		var qrDefaultPort int64

		if o.DefaultPort != nil {
			qrDefaultPort = *o.DefaultPort
		}
		qDefaultPort := swag.FormatInt64(qrDefaultPort)
		if qDefaultPort != "" {

			if err := r.SetQueryParam("defaultPort", qDefaultPort); err != nil {
				return err
			}
		}
	}

	if o.DefaultProtocol != nil {

		// query param defaultProtocol
		var qrDefaultProtocol string

		if o.DefaultProtocol != nil {
			qrDefaultProtocol = *o.DefaultProtocol
		}
		qDefaultProtocol := qrDefaultProtocol
		if qDefaultProtocol != "" {

			if err := r.SetQueryParam("defaultProtocol", qDefaultProtocol); err != nil {
				return err
			}
		}
	}

	if o.DefaultSubdomain != nil {

		// query param defaultSubdomain
		var qrDefaultSubdomain string

		if o.DefaultSubdomain != nil {
			qrDefaultSubdomain = *o.DefaultSubdomain
		}
		qDefaultSubdomain := qrDefaultSubdomain
		if qDefaultSubdomain != "" {

			if err := r.SetQueryParam("defaultSubdomain", qDefaultSubdomain); err != nil {
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

	if o.LegacyConfigsEnbled != nil {

		// query param legacyConfigsEnbled
		var qrLegacyConfigsEnbled bool

		if o.LegacyConfigsEnbled != nil {
			qrLegacyConfigsEnbled = *o.LegacyConfigsEnbled
		}
		qLegacyConfigsEnbled := swag.FormatBool(qrLegacyConfigsEnbled)
		if qLegacyConfigsEnbled != "" {

			if err := r.SetQueryParam("legacyConfigsEnbled", qLegacyConfigsEnbled); err != nil {
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

	if o.UpdatedAt != nil {

		// query param updatedAt
		var qrUpdatedAt strfmt.DateTime

		if o.UpdatedAt != nil {
			qrUpdatedAt = *o.UpdatedAt
		}
		qUpdatedAt := qrUpdatedAt.String()
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
