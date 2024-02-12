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
	"github.com/go-openapi/swag"
)

// NewGetAPIChangesetsV3Params creates a new GetAPIChangesetsV3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIChangesetsV3Params() *GetAPIChangesetsV3Params {
	return &GetAPIChangesetsV3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIChangesetsV3ParamsWithTimeout creates a new GetAPIChangesetsV3Params object
// with the ability to set a timeout on a request.
func NewGetAPIChangesetsV3ParamsWithTimeout(timeout time.Duration) *GetAPIChangesetsV3Params {
	return &GetAPIChangesetsV3Params{
		timeout: timeout,
	}
}

// NewGetAPIChangesetsV3ParamsWithContext creates a new GetAPIChangesetsV3Params object
// with the ability to set a context for a request.
func NewGetAPIChangesetsV3ParamsWithContext(ctx context.Context) *GetAPIChangesetsV3Params {
	return &GetAPIChangesetsV3Params{
		Context: ctx,
	}
}

// NewGetAPIChangesetsV3ParamsWithHTTPClient creates a new GetAPIChangesetsV3Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIChangesetsV3ParamsWithHTTPClient(client *http.Client) *GetAPIChangesetsV3Params {
	return &GetAPIChangesetsV3Params{
		HTTPClient: client,
	}
}

/* GetAPIChangesetsV3Params contains all the parameters to send to the API endpoint
   for the get API changesets v3 operation.

   Typically these are written to a http.Request.
*/
type GetAPIChangesetsV3Params struct {

	// AppliedAt.
	//
	// Format: date-time
	AppliedAt *strfmt.DateTime

	// AppliedBy.
	AppliedBy *string

	// ChartRelease.
	ChartRelease *string

	// FromAppVersionBranch.
	FromAppVersionBranch *string

	// FromAppVersionCommit.
	FromAppVersionCommit *string

	// FromAppVersionExact.
	FromAppVersionExact *string

	// FromAppVersionFollowChartRelease.
	FromAppVersionFollowChartRelease *string

	// FromAppVersionReference.
	FromAppVersionReference *string

	// FromAppVersionResolver.
	FromAppVersionResolver *string

	// FromChartVersionExact.
	FromChartVersionExact *string

	// FromChartVersionFollowChartRelease.
	FromChartVersionFollowChartRelease *string

	// FromChartVersionReference.
	FromChartVersionReference *string

	// FromChartVersionResolver.
	FromChartVersionResolver *string

	// FromHelmfileRef.
	FromHelmfileRef *string

	// FromHelmfileRefEnabled.
	FromHelmfileRefEnabled *bool

	// FromResolvedAt.
	//
	// Format: date-time
	FromResolvedAt *strfmt.DateTime

	/* ID.

	   Get specific changesets by their IDs, can be passed multiple times and/or be comma-separated
	*/
	ID []int64

	/* Limit.

	   Control how many Changesets are returned (default 100), ignored if specific IDs are passed
	*/
	Limit *int64

	/* Offset.

	   Control the offset for the returned Changesets (default 0), ignored if specific IDs are passed
	*/
	Offset *int64

	// PlannedBy.
	PlannedBy *string

	// SupersededAt.
	//
	// Format: date-time
	SupersededAt *strfmt.DateTime

	// ToAppVersionBranch.
	ToAppVersionBranch *string

	// ToAppVersionCommit.
	ToAppVersionCommit *string

	// ToAppVersionExact.
	ToAppVersionExact *string

	// ToAppVersionFollowChartRelease.
	ToAppVersionFollowChartRelease *string

	// ToAppVersionReference.
	ToAppVersionReference *string

	// ToAppVersionResolver.
	ToAppVersionResolver *string

	// ToChartVersionExact.
	ToChartVersionExact *string

	// ToChartVersionFollowChartRelease.
	ToChartVersionFollowChartRelease *string

	// ToChartVersionReference.
	ToChartVersionReference *string

	// ToChartVersionResolver.
	ToChartVersionResolver *string

	// ToHelmfileRef.
	ToHelmfileRef *string

	// ToHelmfileRefEnabled.
	ToHelmfileRefEnabled *bool

	// ToResolvedAt.
	//
	// Format: date-time
	ToResolvedAt *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API changesets v3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIChangesetsV3Params) WithDefaults() *GetAPIChangesetsV3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API changesets v3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIChangesetsV3Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithTimeout(timeout time.Duration) *GetAPIChangesetsV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithContext(ctx context.Context) *GetAPIChangesetsV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithHTTPClient(client *http.Client) *GetAPIChangesetsV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppliedAt adds the appliedAt to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithAppliedAt(appliedAt *strfmt.DateTime) *GetAPIChangesetsV3Params {
	o.SetAppliedAt(appliedAt)
	return o
}

// SetAppliedAt adds the appliedAt to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetAppliedAt(appliedAt *strfmt.DateTime) {
	o.AppliedAt = appliedAt
}

// WithAppliedBy adds the appliedBy to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithAppliedBy(appliedBy *string) *GetAPIChangesetsV3Params {
	o.SetAppliedBy(appliedBy)
	return o
}

// SetAppliedBy adds the appliedBy to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetAppliedBy(appliedBy *string) {
	o.AppliedBy = appliedBy
}

// WithChartRelease adds the chartRelease to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithChartRelease(chartRelease *string) *GetAPIChangesetsV3Params {
	o.SetChartRelease(chartRelease)
	return o
}

// SetChartRelease adds the chartRelease to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetChartRelease(chartRelease *string) {
	o.ChartRelease = chartRelease
}

// WithFromAppVersionBranch adds the fromAppVersionBranch to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithFromAppVersionBranch(fromAppVersionBranch *string) *GetAPIChangesetsV3Params {
	o.SetFromAppVersionBranch(fromAppVersionBranch)
	return o
}

// SetFromAppVersionBranch adds the fromAppVersionBranch to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetFromAppVersionBranch(fromAppVersionBranch *string) {
	o.FromAppVersionBranch = fromAppVersionBranch
}

// WithFromAppVersionCommit adds the fromAppVersionCommit to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithFromAppVersionCommit(fromAppVersionCommit *string) *GetAPIChangesetsV3Params {
	o.SetFromAppVersionCommit(fromAppVersionCommit)
	return o
}

// SetFromAppVersionCommit adds the fromAppVersionCommit to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetFromAppVersionCommit(fromAppVersionCommit *string) {
	o.FromAppVersionCommit = fromAppVersionCommit
}

// WithFromAppVersionExact adds the fromAppVersionExact to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithFromAppVersionExact(fromAppVersionExact *string) *GetAPIChangesetsV3Params {
	o.SetFromAppVersionExact(fromAppVersionExact)
	return o
}

// SetFromAppVersionExact adds the fromAppVersionExact to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetFromAppVersionExact(fromAppVersionExact *string) {
	o.FromAppVersionExact = fromAppVersionExact
}

// WithFromAppVersionFollowChartRelease adds the fromAppVersionFollowChartRelease to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithFromAppVersionFollowChartRelease(fromAppVersionFollowChartRelease *string) *GetAPIChangesetsV3Params {
	o.SetFromAppVersionFollowChartRelease(fromAppVersionFollowChartRelease)
	return o
}

// SetFromAppVersionFollowChartRelease adds the fromAppVersionFollowChartRelease to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetFromAppVersionFollowChartRelease(fromAppVersionFollowChartRelease *string) {
	o.FromAppVersionFollowChartRelease = fromAppVersionFollowChartRelease
}

// WithFromAppVersionReference adds the fromAppVersionReference to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithFromAppVersionReference(fromAppVersionReference *string) *GetAPIChangesetsV3Params {
	o.SetFromAppVersionReference(fromAppVersionReference)
	return o
}

// SetFromAppVersionReference adds the fromAppVersionReference to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetFromAppVersionReference(fromAppVersionReference *string) {
	o.FromAppVersionReference = fromAppVersionReference
}

// WithFromAppVersionResolver adds the fromAppVersionResolver to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithFromAppVersionResolver(fromAppVersionResolver *string) *GetAPIChangesetsV3Params {
	o.SetFromAppVersionResolver(fromAppVersionResolver)
	return o
}

// SetFromAppVersionResolver adds the fromAppVersionResolver to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetFromAppVersionResolver(fromAppVersionResolver *string) {
	o.FromAppVersionResolver = fromAppVersionResolver
}

// WithFromChartVersionExact adds the fromChartVersionExact to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithFromChartVersionExact(fromChartVersionExact *string) *GetAPIChangesetsV3Params {
	o.SetFromChartVersionExact(fromChartVersionExact)
	return o
}

// SetFromChartVersionExact adds the fromChartVersionExact to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetFromChartVersionExact(fromChartVersionExact *string) {
	o.FromChartVersionExact = fromChartVersionExact
}

// WithFromChartVersionFollowChartRelease adds the fromChartVersionFollowChartRelease to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithFromChartVersionFollowChartRelease(fromChartVersionFollowChartRelease *string) *GetAPIChangesetsV3Params {
	o.SetFromChartVersionFollowChartRelease(fromChartVersionFollowChartRelease)
	return o
}

// SetFromChartVersionFollowChartRelease adds the fromChartVersionFollowChartRelease to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetFromChartVersionFollowChartRelease(fromChartVersionFollowChartRelease *string) {
	o.FromChartVersionFollowChartRelease = fromChartVersionFollowChartRelease
}

// WithFromChartVersionReference adds the fromChartVersionReference to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithFromChartVersionReference(fromChartVersionReference *string) *GetAPIChangesetsV3Params {
	o.SetFromChartVersionReference(fromChartVersionReference)
	return o
}

// SetFromChartVersionReference adds the fromChartVersionReference to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetFromChartVersionReference(fromChartVersionReference *string) {
	o.FromChartVersionReference = fromChartVersionReference
}

// WithFromChartVersionResolver adds the fromChartVersionResolver to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithFromChartVersionResolver(fromChartVersionResolver *string) *GetAPIChangesetsV3Params {
	o.SetFromChartVersionResolver(fromChartVersionResolver)
	return o
}

// SetFromChartVersionResolver adds the fromChartVersionResolver to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetFromChartVersionResolver(fromChartVersionResolver *string) {
	o.FromChartVersionResolver = fromChartVersionResolver
}

// WithFromHelmfileRef adds the fromHelmfileRef to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithFromHelmfileRef(fromHelmfileRef *string) *GetAPIChangesetsV3Params {
	o.SetFromHelmfileRef(fromHelmfileRef)
	return o
}

// SetFromHelmfileRef adds the fromHelmfileRef to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetFromHelmfileRef(fromHelmfileRef *string) {
	o.FromHelmfileRef = fromHelmfileRef
}

// WithFromHelmfileRefEnabled adds the fromHelmfileRefEnabled to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithFromHelmfileRefEnabled(fromHelmfileRefEnabled *bool) *GetAPIChangesetsV3Params {
	o.SetFromHelmfileRefEnabled(fromHelmfileRefEnabled)
	return o
}

// SetFromHelmfileRefEnabled adds the fromHelmfileRefEnabled to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetFromHelmfileRefEnabled(fromHelmfileRefEnabled *bool) {
	o.FromHelmfileRefEnabled = fromHelmfileRefEnabled
}

// WithFromResolvedAt adds the fromResolvedAt to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithFromResolvedAt(fromResolvedAt *strfmt.DateTime) *GetAPIChangesetsV3Params {
	o.SetFromResolvedAt(fromResolvedAt)
	return o
}

// SetFromResolvedAt adds the fromResolvedAt to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetFromResolvedAt(fromResolvedAt *strfmt.DateTime) {
	o.FromResolvedAt = fromResolvedAt
}

// WithID adds the id to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithID(id []int64) *GetAPIChangesetsV3Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetID(id []int64) {
	o.ID = id
}

// WithLimit adds the limit to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithLimit(limit *int64) *GetAPIChangesetsV3Params {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithOffset(offset *int64) *GetAPIChangesetsV3Params {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPlannedBy adds the plannedBy to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithPlannedBy(plannedBy *string) *GetAPIChangesetsV3Params {
	o.SetPlannedBy(plannedBy)
	return o
}

// SetPlannedBy adds the plannedBy to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetPlannedBy(plannedBy *string) {
	o.PlannedBy = plannedBy
}

// WithSupersededAt adds the supersededAt to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithSupersededAt(supersededAt *strfmt.DateTime) *GetAPIChangesetsV3Params {
	o.SetSupersededAt(supersededAt)
	return o
}

// SetSupersededAt adds the supersededAt to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetSupersededAt(supersededAt *strfmt.DateTime) {
	o.SupersededAt = supersededAt
}

// WithToAppVersionBranch adds the toAppVersionBranch to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithToAppVersionBranch(toAppVersionBranch *string) *GetAPIChangesetsV3Params {
	o.SetToAppVersionBranch(toAppVersionBranch)
	return o
}

// SetToAppVersionBranch adds the toAppVersionBranch to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetToAppVersionBranch(toAppVersionBranch *string) {
	o.ToAppVersionBranch = toAppVersionBranch
}

// WithToAppVersionCommit adds the toAppVersionCommit to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithToAppVersionCommit(toAppVersionCommit *string) *GetAPIChangesetsV3Params {
	o.SetToAppVersionCommit(toAppVersionCommit)
	return o
}

// SetToAppVersionCommit adds the toAppVersionCommit to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetToAppVersionCommit(toAppVersionCommit *string) {
	o.ToAppVersionCommit = toAppVersionCommit
}

// WithToAppVersionExact adds the toAppVersionExact to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithToAppVersionExact(toAppVersionExact *string) *GetAPIChangesetsV3Params {
	o.SetToAppVersionExact(toAppVersionExact)
	return o
}

// SetToAppVersionExact adds the toAppVersionExact to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetToAppVersionExact(toAppVersionExact *string) {
	o.ToAppVersionExact = toAppVersionExact
}

// WithToAppVersionFollowChartRelease adds the toAppVersionFollowChartRelease to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithToAppVersionFollowChartRelease(toAppVersionFollowChartRelease *string) *GetAPIChangesetsV3Params {
	o.SetToAppVersionFollowChartRelease(toAppVersionFollowChartRelease)
	return o
}

// SetToAppVersionFollowChartRelease adds the toAppVersionFollowChartRelease to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetToAppVersionFollowChartRelease(toAppVersionFollowChartRelease *string) {
	o.ToAppVersionFollowChartRelease = toAppVersionFollowChartRelease
}

// WithToAppVersionReference adds the toAppVersionReference to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithToAppVersionReference(toAppVersionReference *string) *GetAPIChangesetsV3Params {
	o.SetToAppVersionReference(toAppVersionReference)
	return o
}

// SetToAppVersionReference adds the toAppVersionReference to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetToAppVersionReference(toAppVersionReference *string) {
	o.ToAppVersionReference = toAppVersionReference
}

// WithToAppVersionResolver adds the toAppVersionResolver to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithToAppVersionResolver(toAppVersionResolver *string) *GetAPIChangesetsV3Params {
	o.SetToAppVersionResolver(toAppVersionResolver)
	return o
}

// SetToAppVersionResolver adds the toAppVersionResolver to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetToAppVersionResolver(toAppVersionResolver *string) {
	o.ToAppVersionResolver = toAppVersionResolver
}

// WithToChartVersionExact adds the toChartVersionExact to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithToChartVersionExact(toChartVersionExact *string) *GetAPIChangesetsV3Params {
	o.SetToChartVersionExact(toChartVersionExact)
	return o
}

// SetToChartVersionExact adds the toChartVersionExact to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetToChartVersionExact(toChartVersionExact *string) {
	o.ToChartVersionExact = toChartVersionExact
}

// WithToChartVersionFollowChartRelease adds the toChartVersionFollowChartRelease to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithToChartVersionFollowChartRelease(toChartVersionFollowChartRelease *string) *GetAPIChangesetsV3Params {
	o.SetToChartVersionFollowChartRelease(toChartVersionFollowChartRelease)
	return o
}

// SetToChartVersionFollowChartRelease adds the toChartVersionFollowChartRelease to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetToChartVersionFollowChartRelease(toChartVersionFollowChartRelease *string) {
	o.ToChartVersionFollowChartRelease = toChartVersionFollowChartRelease
}

// WithToChartVersionReference adds the toChartVersionReference to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithToChartVersionReference(toChartVersionReference *string) *GetAPIChangesetsV3Params {
	o.SetToChartVersionReference(toChartVersionReference)
	return o
}

// SetToChartVersionReference adds the toChartVersionReference to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetToChartVersionReference(toChartVersionReference *string) {
	o.ToChartVersionReference = toChartVersionReference
}

// WithToChartVersionResolver adds the toChartVersionResolver to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithToChartVersionResolver(toChartVersionResolver *string) *GetAPIChangesetsV3Params {
	o.SetToChartVersionResolver(toChartVersionResolver)
	return o
}

// SetToChartVersionResolver adds the toChartVersionResolver to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetToChartVersionResolver(toChartVersionResolver *string) {
	o.ToChartVersionResolver = toChartVersionResolver
}

// WithToHelmfileRef adds the toHelmfileRef to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithToHelmfileRef(toHelmfileRef *string) *GetAPIChangesetsV3Params {
	o.SetToHelmfileRef(toHelmfileRef)
	return o
}

// SetToHelmfileRef adds the toHelmfileRef to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetToHelmfileRef(toHelmfileRef *string) {
	o.ToHelmfileRef = toHelmfileRef
}

// WithToHelmfileRefEnabled adds the toHelmfileRefEnabled to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithToHelmfileRefEnabled(toHelmfileRefEnabled *bool) *GetAPIChangesetsV3Params {
	o.SetToHelmfileRefEnabled(toHelmfileRefEnabled)
	return o
}

// SetToHelmfileRefEnabled adds the toHelmfileRefEnabled to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetToHelmfileRefEnabled(toHelmfileRefEnabled *bool) {
	o.ToHelmfileRefEnabled = toHelmfileRefEnabled
}

// WithToResolvedAt adds the toResolvedAt to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) WithToResolvedAt(toResolvedAt *strfmt.DateTime) *GetAPIChangesetsV3Params {
	o.SetToResolvedAt(toResolvedAt)
	return o
}

// SetToResolvedAt adds the toResolvedAt to the get API changesets v3 params
func (o *GetAPIChangesetsV3Params) SetToResolvedAt(toResolvedAt *strfmt.DateTime) {
	o.ToResolvedAt = toResolvedAt
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIChangesetsV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AppliedAt != nil {

		// query param appliedAt
		var qrAppliedAt strfmt.DateTime

		if o.AppliedAt != nil {
			qrAppliedAt = *o.AppliedAt
		}
		qAppliedAt := qrAppliedAt.String()
		if qAppliedAt != "" {

			if err := r.SetQueryParam("appliedAt", qAppliedAt); err != nil {
				return err
			}
		}
	}

	if o.AppliedBy != nil {

		// query param appliedBy
		var qrAppliedBy string

		if o.AppliedBy != nil {
			qrAppliedBy = *o.AppliedBy
		}
		qAppliedBy := qrAppliedBy
		if qAppliedBy != "" {

			if err := r.SetQueryParam("appliedBy", qAppliedBy); err != nil {
				return err
			}
		}
	}

	if o.ChartRelease != nil {

		// query param chartRelease
		var qrChartRelease string

		if o.ChartRelease != nil {
			qrChartRelease = *o.ChartRelease
		}
		qChartRelease := qrChartRelease
		if qChartRelease != "" {

			if err := r.SetQueryParam("chartRelease", qChartRelease); err != nil {
				return err
			}
		}
	}

	if o.FromAppVersionBranch != nil {

		// query param fromAppVersionBranch
		var qrFromAppVersionBranch string

		if o.FromAppVersionBranch != nil {
			qrFromAppVersionBranch = *o.FromAppVersionBranch
		}
		qFromAppVersionBranch := qrFromAppVersionBranch
		if qFromAppVersionBranch != "" {

			if err := r.SetQueryParam("fromAppVersionBranch", qFromAppVersionBranch); err != nil {
				return err
			}
		}
	}

	if o.FromAppVersionCommit != nil {

		// query param fromAppVersionCommit
		var qrFromAppVersionCommit string

		if o.FromAppVersionCommit != nil {
			qrFromAppVersionCommit = *o.FromAppVersionCommit
		}
		qFromAppVersionCommit := qrFromAppVersionCommit
		if qFromAppVersionCommit != "" {

			if err := r.SetQueryParam("fromAppVersionCommit", qFromAppVersionCommit); err != nil {
				return err
			}
		}
	}

	if o.FromAppVersionExact != nil {

		// query param fromAppVersionExact
		var qrFromAppVersionExact string

		if o.FromAppVersionExact != nil {
			qrFromAppVersionExact = *o.FromAppVersionExact
		}
		qFromAppVersionExact := qrFromAppVersionExact
		if qFromAppVersionExact != "" {

			if err := r.SetQueryParam("fromAppVersionExact", qFromAppVersionExact); err != nil {
				return err
			}
		}
	}

	if o.FromAppVersionFollowChartRelease != nil {

		// query param fromAppVersionFollowChartRelease
		var qrFromAppVersionFollowChartRelease string

		if o.FromAppVersionFollowChartRelease != nil {
			qrFromAppVersionFollowChartRelease = *o.FromAppVersionFollowChartRelease
		}
		qFromAppVersionFollowChartRelease := qrFromAppVersionFollowChartRelease
		if qFromAppVersionFollowChartRelease != "" {

			if err := r.SetQueryParam("fromAppVersionFollowChartRelease", qFromAppVersionFollowChartRelease); err != nil {
				return err
			}
		}
	}

	if o.FromAppVersionReference != nil {

		// query param fromAppVersionReference
		var qrFromAppVersionReference string

		if o.FromAppVersionReference != nil {
			qrFromAppVersionReference = *o.FromAppVersionReference
		}
		qFromAppVersionReference := qrFromAppVersionReference
		if qFromAppVersionReference != "" {

			if err := r.SetQueryParam("fromAppVersionReference", qFromAppVersionReference); err != nil {
				return err
			}
		}
	}

	if o.FromAppVersionResolver != nil {

		// query param fromAppVersionResolver
		var qrFromAppVersionResolver string

		if o.FromAppVersionResolver != nil {
			qrFromAppVersionResolver = *o.FromAppVersionResolver
		}
		qFromAppVersionResolver := qrFromAppVersionResolver
		if qFromAppVersionResolver != "" {

			if err := r.SetQueryParam("fromAppVersionResolver", qFromAppVersionResolver); err != nil {
				return err
			}
		}
	}

	if o.FromChartVersionExact != nil {

		// query param fromChartVersionExact
		var qrFromChartVersionExact string

		if o.FromChartVersionExact != nil {
			qrFromChartVersionExact = *o.FromChartVersionExact
		}
		qFromChartVersionExact := qrFromChartVersionExact
		if qFromChartVersionExact != "" {

			if err := r.SetQueryParam("fromChartVersionExact", qFromChartVersionExact); err != nil {
				return err
			}
		}
	}

	if o.FromChartVersionFollowChartRelease != nil {

		// query param fromChartVersionFollowChartRelease
		var qrFromChartVersionFollowChartRelease string

		if o.FromChartVersionFollowChartRelease != nil {
			qrFromChartVersionFollowChartRelease = *o.FromChartVersionFollowChartRelease
		}
		qFromChartVersionFollowChartRelease := qrFromChartVersionFollowChartRelease
		if qFromChartVersionFollowChartRelease != "" {

			if err := r.SetQueryParam("fromChartVersionFollowChartRelease", qFromChartVersionFollowChartRelease); err != nil {
				return err
			}
		}
	}

	if o.FromChartVersionReference != nil {

		// query param fromChartVersionReference
		var qrFromChartVersionReference string

		if o.FromChartVersionReference != nil {
			qrFromChartVersionReference = *o.FromChartVersionReference
		}
		qFromChartVersionReference := qrFromChartVersionReference
		if qFromChartVersionReference != "" {

			if err := r.SetQueryParam("fromChartVersionReference", qFromChartVersionReference); err != nil {
				return err
			}
		}
	}

	if o.FromChartVersionResolver != nil {

		// query param fromChartVersionResolver
		var qrFromChartVersionResolver string

		if o.FromChartVersionResolver != nil {
			qrFromChartVersionResolver = *o.FromChartVersionResolver
		}
		qFromChartVersionResolver := qrFromChartVersionResolver
		if qFromChartVersionResolver != "" {

			if err := r.SetQueryParam("fromChartVersionResolver", qFromChartVersionResolver); err != nil {
				return err
			}
		}
	}

	if o.FromHelmfileRef != nil {

		// query param fromHelmfileRef
		var qrFromHelmfileRef string

		if o.FromHelmfileRef != nil {
			qrFromHelmfileRef = *o.FromHelmfileRef
		}
		qFromHelmfileRef := qrFromHelmfileRef
		if qFromHelmfileRef != "" {

			if err := r.SetQueryParam("fromHelmfileRef", qFromHelmfileRef); err != nil {
				return err
			}
		}
	}

	if o.FromHelmfileRefEnabled != nil {

		// query param fromHelmfileRefEnabled
		var qrFromHelmfileRefEnabled bool

		if o.FromHelmfileRefEnabled != nil {
			qrFromHelmfileRefEnabled = *o.FromHelmfileRefEnabled
		}
		qFromHelmfileRefEnabled := swag.FormatBool(qrFromHelmfileRefEnabled)
		if qFromHelmfileRefEnabled != "" {

			if err := r.SetQueryParam("fromHelmfileRefEnabled", qFromHelmfileRefEnabled); err != nil {
				return err
			}
		}
	}

	if o.FromResolvedAt != nil {

		// query param fromResolvedAt
		var qrFromResolvedAt strfmt.DateTime

		if o.FromResolvedAt != nil {
			qrFromResolvedAt = *o.FromResolvedAt
		}
		qFromResolvedAt := qrFromResolvedAt.String()
		if qFromResolvedAt != "" {

			if err := r.SetQueryParam("fromResolvedAt", qFromResolvedAt); err != nil {
				return err
			}
		}
	}

	if o.ID != nil {

		// binding items for id
		joinedID := o.bindParamID(reg)

		// query array param id
		if err := r.SetQueryParam("id", joinedID...); err != nil {
			return err
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

	if o.PlannedBy != nil {

		// query param plannedBy
		var qrPlannedBy string

		if o.PlannedBy != nil {
			qrPlannedBy = *o.PlannedBy
		}
		qPlannedBy := qrPlannedBy
		if qPlannedBy != "" {

			if err := r.SetQueryParam("plannedBy", qPlannedBy); err != nil {
				return err
			}
		}
	}

	if o.SupersededAt != nil {

		// query param supersededAt
		var qrSupersededAt strfmt.DateTime

		if o.SupersededAt != nil {
			qrSupersededAt = *o.SupersededAt
		}
		qSupersededAt := qrSupersededAt.String()
		if qSupersededAt != "" {

			if err := r.SetQueryParam("supersededAt", qSupersededAt); err != nil {
				return err
			}
		}
	}

	if o.ToAppVersionBranch != nil {

		// query param toAppVersionBranch
		var qrToAppVersionBranch string

		if o.ToAppVersionBranch != nil {
			qrToAppVersionBranch = *o.ToAppVersionBranch
		}
		qToAppVersionBranch := qrToAppVersionBranch
		if qToAppVersionBranch != "" {

			if err := r.SetQueryParam("toAppVersionBranch", qToAppVersionBranch); err != nil {
				return err
			}
		}
	}

	if o.ToAppVersionCommit != nil {

		// query param toAppVersionCommit
		var qrToAppVersionCommit string

		if o.ToAppVersionCommit != nil {
			qrToAppVersionCommit = *o.ToAppVersionCommit
		}
		qToAppVersionCommit := qrToAppVersionCommit
		if qToAppVersionCommit != "" {

			if err := r.SetQueryParam("toAppVersionCommit", qToAppVersionCommit); err != nil {
				return err
			}
		}
	}

	if o.ToAppVersionExact != nil {

		// query param toAppVersionExact
		var qrToAppVersionExact string

		if o.ToAppVersionExact != nil {
			qrToAppVersionExact = *o.ToAppVersionExact
		}
		qToAppVersionExact := qrToAppVersionExact
		if qToAppVersionExact != "" {

			if err := r.SetQueryParam("toAppVersionExact", qToAppVersionExact); err != nil {
				return err
			}
		}
	}

	if o.ToAppVersionFollowChartRelease != nil {

		// query param toAppVersionFollowChartRelease
		var qrToAppVersionFollowChartRelease string

		if o.ToAppVersionFollowChartRelease != nil {
			qrToAppVersionFollowChartRelease = *o.ToAppVersionFollowChartRelease
		}
		qToAppVersionFollowChartRelease := qrToAppVersionFollowChartRelease
		if qToAppVersionFollowChartRelease != "" {

			if err := r.SetQueryParam("toAppVersionFollowChartRelease", qToAppVersionFollowChartRelease); err != nil {
				return err
			}
		}
	}

	if o.ToAppVersionReference != nil {

		// query param toAppVersionReference
		var qrToAppVersionReference string

		if o.ToAppVersionReference != nil {
			qrToAppVersionReference = *o.ToAppVersionReference
		}
		qToAppVersionReference := qrToAppVersionReference
		if qToAppVersionReference != "" {

			if err := r.SetQueryParam("toAppVersionReference", qToAppVersionReference); err != nil {
				return err
			}
		}
	}

	if o.ToAppVersionResolver != nil {

		// query param toAppVersionResolver
		var qrToAppVersionResolver string

		if o.ToAppVersionResolver != nil {
			qrToAppVersionResolver = *o.ToAppVersionResolver
		}
		qToAppVersionResolver := qrToAppVersionResolver
		if qToAppVersionResolver != "" {

			if err := r.SetQueryParam("toAppVersionResolver", qToAppVersionResolver); err != nil {
				return err
			}
		}
	}

	if o.ToChartVersionExact != nil {

		// query param toChartVersionExact
		var qrToChartVersionExact string

		if o.ToChartVersionExact != nil {
			qrToChartVersionExact = *o.ToChartVersionExact
		}
		qToChartVersionExact := qrToChartVersionExact
		if qToChartVersionExact != "" {

			if err := r.SetQueryParam("toChartVersionExact", qToChartVersionExact); err != nil {
				return err
			}
		}
	}

	if o.ToChartVersionFollowChartRelease != nil {

		// query param toChartVersionFollowChartRelease
		var qrToChartVersionFollowChartRelease string

		if o.ToChartVersionFollowChartRelease != nil {
			qrToChartVersionFollowChartRelease = *o.ToChartVersionFollowChartRelease
		}
		qToChartVersionFollowChartRelease := qrToChartVersionFollowChartRelease
		if qToChartVersionFollowChartRelease != "" {

			if err := r.SetQueryParam("toChartVersionFollowChartRelease", qToChartVersionFollowChartRelease); err != nil {
				return err
			}
		}
	}

	if o.ToChartVersionReference != nil {

		// query param toChartVersionReference
		var qrToChartVersionReference string

		if o.ToChartVersionReference != nil {
			qrToChartVersionReference = *o.ToChartVersionReference
		}
		qToChartVersionReference := qrToChartVersionReference
		if qToChartVersionReference != "" {

			if err := r.SetQueryParam("toChartVersionReference", qToChartVersionReference); err != nil {
				return err
			}
		}
	}

	if o.ToChartVersionResolver != nil {

		// query param toChartVersionResolver
		var qrToChartVersionResolver string

		if o.ToChartVersionResolver != nil {
			qrToChartVersionResolver = *o.ToChartVersionResolver
		}
		qToChartVersionResolver := qrToChartVersionResolver
		if qToChartVersionResolver != "" {

			if err := r.SetQueryParam("toChartVersionResolver", qToChartVersionResolver); err != nil {
				return err
			}
		}
	}

	if o.ToHelmfileRef != nil {

		// query param toHelmfileRef
		var qrToHelmfileRef string

		if o.ToHelmfileRef != nil {
			qrToHelmfileRef = *o.ToHelmfileRef
		}
		qToHelmfileRef := qrToHelmfileRef
		if qToHelmfileRef != "" {

			if err := r.SetQueryParam("toHelmfileRef", qToHelmfileRef); err != nil {
				return err
			}
		}
	}

	if o.ToHelmfileRefEnabled != nil {

		// query param toHelmfileRefEnabled
		var qrToHelmfileRefEnabled bool

		if o.ToHelmfileRefEnabled != nil {
			qrToHelmfileRefEnabled = *o.ToHelmfileRefEnabled
		}
		qToHelmfileRefEnabled := swag.FormatBool(qrToHelmfileRefEnabled)
		if qToHelmfileRefEnabled != "" {

			if err := r.SetQueryParam("toHelmfileRefEnabled", qToHelmfileRefEnabled); err != nil {
				return err
			}
		}
	}

	if o.ToResolvedAt != nil {

		// query param toResolvedAt
		var qrToResolvedAt strfmt.DateTime

		if o.ToResolvedAt != nil {
			qrToResolvedAt = *o.ToResolvedAt
		}
		qToResolvedAt := qrToResolvedAt.String()
		if qToResolvedAt != "" {

			if err := r.SetQueryParam("toResolvedAt", qToResolvedAt); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetAPIChangesetsV3 binds the parameter id
func (o *GetAPIChangesetsV3Params) bindParamID(formats strfmt.Registry) []string {
	iDIR := o.ID

	var iDIC []string
	for _, iDIIR := range iDIR { // explode []int64

		iDIIV := swag.FormatInt64(iDIIR) // int64 as string
		iDIC = append(iDIC, iDIIV)
	}

	// items.CollectionFormat: "csv"
	iDIS := swag.JoinByFormat(iDIC, "csv")

	return iDIS
}
