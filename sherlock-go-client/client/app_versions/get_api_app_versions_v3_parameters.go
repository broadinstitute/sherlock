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
	"github.com/go-openapi/swag"
)

// NewGetAPIAppVersionsV3Params creates a new GetAPIAppVersionsV3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIAppVersionsV3Params() *GetAPIAppVersionsV3Params {
	return &GetAPIAppVersionsV3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIAppVersionsV3ParamsWithTimeout creates a new GetAPIAppVersionsV3Params object
// with the ability to set a timeout on a request.
func NewGetAPIAppVersionsV3ParamsWithTimeout(timeout time.Duration) *GetAPIAppVersionsV3Params {
	return &GetAPIAppVersionsV3Params{
		timeout: timeout,
	}
}

// NewGetAPIAppVersionsV3ParamsWithContext creates a new GetAPIAppVersionsV3Params object
// with the ability to set a context for a request.
func NewGetAPIAppVersionsV3ParamsWithContext(ctx context.Context) *GetAPIAppVersionsV3Params {
	return &GetAPIAppVersionsV3Params{
		Context: ctx,
	}
}

// NewGetAPIAppVersionsV3ParamsWithHTTPClient creates a new GetAPIAppVersionsV3Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIAppVersionsV3ParamsWithHTTPClient(client *http.Client) *GetAPIAppVersionsV3Params {
	return &GetAPIAppVersionsV3Params{
		HTTPClient: client,
	}
}

/* GetAPIAppVersionsV3Params contains all the parameters to send to the API endpoint
   for the get API app versions v3 operation.

   Typically these are written to a http.Request.
*/
type GetAPIAppVersionsV3Params struct {

	/* AppVersion.

	   Required when creating
	*/
	AppVersion *string

	/* Chart.

	   Required when creating
	*/
	Chart *string

	// CreatedAt.
	//
	// Format: date-time
	CreatedAt *strfmt.DateTime

	/* Description.

	   Generally the Git commit message
	*/
	Description *string

	// GitBranch.
	GitBranch *string

	// GitCommit.
	GitCommit *string

	// ID.
	ID *int64

	/* Limit.

	   Control how many AppVersions are returned (default 100)
	*/
	Limit *int64

	/* Offset.

	   Control the offset for the returned AppVersions (default 0)
	*/
	Offset *int64

	// ParentAppVersion.
	ParentAppVersion *string

	// UpdatedAt.
	//
	// Format: date-time
	UpdatedAt *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API app versions v3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIAppVersionsV3Params) WithDefaults() *GetAPIAppVersionsV3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API app versions v3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIAppVersionsV3Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API app versions v3 params
func (o *GetAPIAppVersionsV3Params) WithTimeout(timeout time.Duration) *GetAPIAppVersionsV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API app versions v3 params
func (o *GetAPIAppVersionsV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API app versions v3 params
func (o *GetAPIAppVersionsV3Params) WithContext(ctx context.Context) *GetAPIAppVersionsV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API app versions v3 params
func (o *GetAPIAppVersionsV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API app versions v3 params
func (o *GetAPIAppVersionsV3Params) WithHTTPClient(client *http.Client) *GetAPIAppVersionsV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API app versions v3 params
func (o *GetAPIAppVersionsV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppVersion adds the appVersion to the get API app versions v3 params
func (o *GetAPIAppVersionsV3Params) WithAppVersion(appVersion *string) *GetAPIAppVersionsV3Params {
	o.SetAppVersion(appVersion)
	return o
}

// SetAppVersion adds the appVersion to the get API app versions v3 params
func (o *GetAPIAppVersionsV3Params) SetAppVersion(appVersion *string) {
	o.AppVersion = appVersion
}

// WithChart adds the chart to the get API app versions v3 params
func (o *GetAPIAppVersionsV3Params) WithChart(chart *string) *GetAPIAppVersionsV3Params {
	o.SetChart(chart)
	return o
}

// SetChart adds the chart to the get API app versions v3 params
func (o *GetAPIAppVersionsV3Params) SetChart(chart *string) {
	o.Chart = chart
}

// WithCreatedAt adds the createdAt to the get API app versions v3 params
func (o *GetAPIAppVersionsV3Params) WithCreatedAt(createdAt *strfmt.DateTime) *GetAPIAppVersionsV3Params {
	o.SetCreatedAt(createdAt)
	return o
}

// SetCreatedAt adds the createdAt to the get API app versions v3 params
func (o *GetAPIAppVersionsV3Params) SetCreatedAt(createdAt *strfmt.DateTime) {
	o.CreatedAt = createdAt
}

// WithDescription adds the description to the get API app versions v3 params
func (o *GetAPIAppVersionsV3Params) WithDescription(description *string) *GetAPIAppVersionsV3Params {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the get API app versions v3 params
func (o *GetAPIAppVersionsV3Params) SetDescription(description *string) {
	o.Description = description
}

// WithGitBranch adds the gitBranch to the get API app versions v3 params
func (o *GetAPIAppVersionsV3Params) WithGitBranch(gitBranch *string) *GetAPIAppVersionsV3Params {
	o.SetGitBranch(gitBranch)
	return o
}

// SetGitBranch adds the gitBranch to the get API app versions v3 params
func (o *GetAPIAppVersionsV3Params) SetGitBranch(gitBranch *string) {
	o.GitBranch = gitBranch
}

// WithGitCommit adds the gitCommit to the get API app versions v3 params
func (o *GetAPIAppVersionsV3Params) WithGitCommit(gitCommit *string) *GetAPIAppVersionsV3Params {
	o.SetGitCommit(gitCommit)
	return o
}

// SetGitCommit adds the gitCommit to the get API app versions v3 params
func (o *GetAPIAppVersionsV3Params) SetGitCommit(gitCommit *string) {
	o.GitCommit = gitCommit
}

// WithID adds the id to the get API app versions v3 params
func (o *GetAPIAppVersionsV3Params) WithID(id *int64) *GetAPIAppVersionsV3Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the get API app versions v3 params
func (o *GetAPIAppVersionsV3Params) SetID(id *int64) {
	o.ID = id
}

// WithLimit adds the limit to the get API app versions v3 params
func (o *GetAPIAppVersionsV3Params) WithLimit(limit *int64) *GetAPIAppVersionsV3Params {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get API app versions v3 params
func (o *GetAPIAppVersionsV3Params) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get API app versions v3 params
func (o *GetAPIAppVersionsV3Params) WithOffset(offset *int64) *GetAPIAppVersionsV3Params {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get API app versions v3 params
func (o *GetAPIAppVersionsV3Params) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithParentAppVersion adds the parentAppVersion to the get API app versions v3 params
func (o *GetAPIAppVersionsV3Params) WithParentAppVersion(parentAppVersion *string) *GetAPIAppVersionsV3Params {
	o.SetParentAppVersion(parentAppVersion)
	return o
}

// SetParentAppVersion adds the parentAppVersion to the get API app versions v3 params
func (o *GetAPIAppVersionsV3Params) SetParentAppVersion(parentAppVersion *string) {
	o.ParentAppVersion = parentAppVersion
}

// WithUpdatedAt adds the updatedAt to the get API app versions v3 params
func (o *GetAPIAppVersionsV3Params) WithUpdatedAt(updatedAt *strfmt.DateTime) *GetAPIAppVersionsV3Params {
	o.SetUpdatedAt(updatedAt)
	return o
}

// SetUpdatedAt adds the updatedAt to the get API app versions v3 params
func (o *GetAPIAppVersionsV3Params) SetUpdatedAt(updatedAt *strfmt.DateTime) {
	o.UpdatedAt = updatedAt
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIAppVersionsV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AppVersion != nil {

		// query param appVersion
		var qrAppVersion string

		if o.AppVersion != nil {
			qrAppVersion = *o.AppVersion
		}
		qAppVersion := qrAppVersion
		if qAppVersion != "" {

			if err := r.SetQueryParam("appVersion", qAppVersion); err != nil {
				return err
			}
		}
	}

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

	if o.Description != nil {

		// query param description
		var qrDescription string

		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}
	}

	if o.GitBranch != nil {

		// query param gitBranch
		var qrGitBranch string

		if o.GitBranch != nil {
			qrGitBranch = *o.GitBranch
		}
		qGitBranch := qrGitBranch
		if qGitBranch != "" {

			if err := r.SetQueryParam("gitBranch", qGitBranch); err != nil {
				return err
			}
		}
	}

	if o.GitCommit != nil {

		// query param gitCommit
		var qrGitCommit string

		if o.GitCommit != nil {
			qrGitCommit = *o.GitCommit
		}
		qGitCommit := qrGitCommit
		if qGitCommit != "" {

			if err := r.SetQueryParam("gitCommit", qGitCommit); err != nil {
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

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.ParentAppVersion != nil {

		// query param parentAppVersion
		var qrParentAppVersion string

		if o.ParentAppVersion != nil {
			qrParentAppVersion = *o.ParentAppVersion
		}
		qParentAppVersion := qrParentAppVersion
		if qParentAppVersion != "" {

			if err := r.SetQueryParam("parentAppVersion", qParentAppVersion); err != nil {
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
