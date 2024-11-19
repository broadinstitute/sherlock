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
	"github.com/go-openapi/swag"
)

// NewGetAPIEnvironmentsV3Params creates a new GetAPIEnvironmentsV3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIEnvironmentsV3Params() *GetAPIEnvironmentsV3Params {
	return &GetAPIEnvironmentsV3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIEnvironmentsV3ParamsWithTimeout creates a new GetAPIEnvironmentsV3Params object
// with the ability to set a timeout on a request.
func NewGetAPIEnvironmentsV3ParamsWithTimeout(timeout time.Duration) *GetAPIEnvironmentsV3Params {
	return &GetAPIEnvironmentsV3Params{
		timeout: timeout,
	}
}

// NewGetAPIEnvironmentsV3ParamsWithContext creates a new GetAPIEnvironmentsV3Params object
// with the ability to set a context for a request.
func NewGetAPIEnvironmentsV3ParamsWithContext(ctx context.Context) *GetAPIEnvironmentsV3Params {
	return &GetAPIEnvironmentsV3Params{
		Context: ctx,
	}
}

// NewGetAPIEnvironmentsV3ParamsWithHTTPClient creates a new GetAPIEnvironmentsV3Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIEnvironmentsV3ParamsWithHTTPClient(client *http.Client) *GetAPIEnvironmentsV3Params {
	return &GetAPIEnvironmentsV3Params{
		HTTPClient: client,
	}
}

/* GetAPIEnvironmentsV3Params contains all the parameters to send to the API endpoint
   for the get API environments v3 operation.

   Typically these are written to a http.Request.
*/
type GetAPIEnvironmentsV3Params struct {

	// PactIdentifier.
	PactIdentifier *string

	/* AutoPopulateChartReleases.

	   If true when creating, dynamic environments copy from template and template environments get the honeycomb chart

	   Default: true
	*/
	AutoPopulateChartReleases *bool

	/* Base.

	   Required when creating
	*/
	Base *string

	// BaseDomain.
	//
	// Default: "bee.envs-terra.bio"
	BaseDomain *string

	// CreatedAt.
	//
	// Format: date-time
	CreatedAt *strfmt.DateTime

	// DefaultCluster.
	DefaultCluster *string

	/* DefaultNamespace.

	   When creating, will be calculated if left empty
	*/
	DefaultNamespace *string

	/* DeleteAfter.

	   If set, the BEE will be automatically deleted after this time. Can be set to "" or Go's zero time value to clear the field.

	   Format: date-time
	*/
	DeleteAfter *strfmt.DateTime

	// Description.
	Description *string

	/* EnableJanitor.

	   If true, janitor resource cleanup will be enabled for this environment. BEEs default to template's value, templates default to true, and static/live environments default to false.
	*/
	EnableJanitor *bool

	// HelmfileRef.
	//
	// Default: "HEAD"
	HelmfileRef *string

	// ID.
	ID *int64

	// Lifecycle.
	//
	// Default: "dynamic"
	Lifecycle *string

	/* Limit.

	   Control how many Environments are returned (default 0, meaning all)
	*/
	Limit *int64

	/* Name.

	   When creating, will be calculated if dynamic, required otherwise
	*/
	Name *string

	// NamePrefixesDomain.
	//
	// Default: true
	NamePrefixesDomain *bool

	/* Offline.

	   Applicable for BEEs only, whether Thelma should render the BEE as "offline" zero replicas (this field is a target state, not a status)
	*/
	Offline *bool

	/* OfflineScheduleBeginEnabled.

	   When enabled, the BEE will be slated to go offline around the begin time each day
	*/
	OfflineScheduleBeginEnabled *bool

	/* OfflineScheduleBeginTime.

	   Stored with timezone to determine day of the week

	   Format: date-time
	*/
	OfflineScheduleBeginTime *strfmt.DateTime

	/* OfflineScheduleEndEnabled.

	   When enabled, the BEE will be slated to come online around the end time each weekday (each day if weekends enabled)
	*/
	OfflineScheduleEndEnabled *bool

	/* OfflineScheduleEndTime.

	   Stored with timezone to determine day of the week

	   Format: date-time
	*/
	OfflineScheduleEndTime *strfmt.DateTime

	// OfflineScheduleEndWeekends.
	OfflineScheduleEndWeekends *bool

	/* Offset.

	   Control the offset for the returned Environments (default 0)
	*/
	Offset *int64

	/* Owner.

	   When creating, will default to you
	*/
	Owner *string

	// PagerdutyIntegration.
	PagerdutyIntegration *string

	/* PreventDeletion.

	   Used to protect specific BEEs from deletion (thelma checks this field)
	*/
	PreventDeletion *bool

	/* RequiredRole.

	   If present, requires membership in the given role for mutations. Set to an empty string to clear.
	*/
	RequiredRole *string

	// RequiresSuitability.
	RequiresSuitability *bool

	// ServiceBannerBucket.
	ServiceBannerBucket *string

	/* TemplateEnvironment.

	   Required for dynamic environments
	*/
	TemplateEnvironment *string

	/* UniqueResourcePrefix.

	   When creating, will be calculated if left empty
	*/
	UniqueResourcePrefix *string

	// UpdatedAt.
	//
	// Format: date-time
	UpdatedAt *strfmt.DateTime

	/* ValuesName.

	   When creating, defaults to template name or environment name
	*/
	ValuesName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API environments v3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIEnvironmentsV3Params) WithDefaults() *GetAPIEnvironmentsV3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API environments v3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIEnvironmentsV3Params) SetDefaults() {
	var (
		autoPopulateChartReleasesDefault = bool(true)

		baseDomainDefault = string("bee.envs-terra.bio")

		helmfileRefDefault = string("HEAD")

		lifecycleDefault = string("dynamic")

		namePrefixesDomainDefault = bool(true)

		offlineDefault = bool(false)

		preventDeletionDefault = bool(false)
	)

	val := GetAPIEnvironmentsV3Params{
		AutoPopulateChartReleases: &autoPopulateChartReleasesDefault,
		BaseDomain:                &baseDomainDefault,
		HelmfileRef:               &helmfileRefDefault,
		Lifecycle:                 &lifecycleDefault,
		NamePrefixesDomain:        &namePrefixesDomainDefault,
		Offline:                   &offlineDefault,
		PreventDeletion:           &preventDeletionDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithTimeout(timeout time.Duration) *GetAPIEnvironmentsV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithContext(ctx context.Context) *GetAPIEnvironmentsV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithHTTPClient(client *http.Client) *GetAPIEnvironmentsV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPactIdentifier adds the pactIdentifier to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithPactIdentifier(pactIdentifier *string) *GetAPIEnvironmentsV3Params {
	o.SetPactIdentifier(pactIdentifier)
	return o
}

// SetPactIdentifier adds the pactIdentifier to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetPactIdentifier(pactIdentifier *string) {
	o.PactIdentifier = pactIdentifier
}

// WithAutoPopulateChartReleases adds the autoPopulateChartReleases to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithAutoPopulateChartReleases(autoPopulateChartReleases *bool) *GetAPIEnvironmentsV3Params {
	o.SetAutoPopulateChartReleases(autoPopulateChartReleases)
	return o
}

// SetAutoPopulateChartReleases adds the autoPopulateChartReleases to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetAutoPopulateChartReleases(autoPopulateChartReleases *bool) {
	o.AutoPopulateChartReleases = autoPopulateChartReleases
}

// WithBase adds the base to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithBase(base *string) *GetAPIEnvironmentsV3Params {
	o.SetBase(base)
	return o
}

// SetBase adds the base to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetBase(base *string) {
	o.Base = base
}

// WithBaseDomain adds the baseDomain to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithBaseDomain(baseDomain *string) *GetAPIEnvironmentsV3Params {
	o.SetBaseDomain(baseDomain)
	return o
}

// SetBaseDomain adds the baseDomain to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetBaseDomain(baseDomain *string) {
	o.BaseDomain = baseDomain
}

// WithCreatedAt adds the createdAt to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithCreatedAt(createdAt *strfmt.DateTime) *GetAPIEnvironmentsV3Params {
	o.SetCreatedAt(createdAt)
	return o
}

// SetCreatedAt adds the createdAt to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetCreatedAt(createdAt *strfmt.DateTime) {
	o.CreatedAt = createdAt
}

// WithDefaultCluster adds the defaultCluster to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithDefaultCluster(defaultCluster *string) *GetAPIEnvironmentsV3Params {
	o.SetDefaultCluster(defaultCluster)
	return o
}

// SetDefaultCluster adds the defaultCluster to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetDefaultCluster(defaultCluster *string) {
	o.DefaultCluster = defaultCluster
}

// WithDefaultNamespace adds the defaultNamespace to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithDefaultNamespace(defaultNamespace *string) *GetAPIEnvironmentsV3Params {
	o.SetDefaultNamespace(defaultNamespace)
	return o
}

// SetDefaultNamespace adds the defaultNamespace to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetDefaultNamespace(defaultNamespace *string) {
	o.DefaultNamespace = defaultNamespace
}

// WithDeleteAfter adds the deleteAfter to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithDeleteAfter(deleteAfter *strfmt.DateTime) *GetAPIEnvironmentsV3Params {
	o.SetDeleteAfter(deleteAfter)
	return o
}

// SetDeleteAfter adds the deleteAfter to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetDeleteAfter(deleteAfter *strfmt.DateTime) {
	o.DeleteAfter = deleteAfter
}

// WithDescription adds the description to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithDescription(description *string) *GetAPIEnvironmentsV3Params {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetDescription(description *string) {
	o.Description = description
}

// WithEnableJanitor adds the enableJanitor to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithEnableJanitor(enableJanitor *bool) *GetAPIEnvironmentsV3Params {
	o.SetEnableJanitor(enableJanitor)
	return o
}

// SetEnableJanitor adds the enableJanitor to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetEnableJanitor(enableJanitor *bool) {
	o.EnableJanitor = enableJanitor
}

// WithHelmfileRef adds the helmfileRef to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithHelmfileRef(helmfileRef *string) *GetAPIEnvironmentsV3Params {
	o.SetHelmfileRef(helmfileRef)
	return o
}

// SetHelmfileRef adds the helmfileRef to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetHelmfileRef(helmfileRef *string) {
	o.HelmfileRef = helmfileRef
}

// WithID adds the id to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithID(id *int64) *GetAPIEnvironmentsV3Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetID(id *int64) {
	o.ID = id
}

// WithLifecycle adds the lifecycle to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithLifecycle(lifecycle *string) *GetAPIEnvironmentsV3Params {
	o.SetLifecycle(lifecycle)
	return o
}

// SetLifecycle adds the lifecycle to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetLifecycle(lifecycle *string) {
	o.Lifecycle = lifecycle
}

// WithLimit adds the limit to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithLimit(limit *int64) *GetAPIEnvironmentsV3Params {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithName(name *string) *GetAPIEnvironmentsV3Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetName(name *string) {
	o.Name = name
}

// WithNamePrefixesDomain adds the namePrefixesDomain to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithNamePrefixesDomain(namePrefixesDomain *bool) *GetAPIEnvironmentsV3Params {
	o.SetNamePrefixesDomain(namePrefixesDomain)
	return o
}

// SetNamePrefixesDomain adds the namePrefixesDomain to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetNamePrefixesDomain(namePrefixesDomain *bool) {
	o.NamePrefixesDomain = namePrefixesDomain
}

// WithOffline adds the offline to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithOffline(offline *bool) *GetAPIEnvironmentsV3Params {
	o.SetOffline(offline)
	return o
}

// SetOffline adds the offline to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetOffline(offline *bool) {
	o.Offline = offline
}

// WithOfflineScheduleBeginEnabled adds the offlineScheduleBeginEnabled to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithOfflineScheduleBeginEnabled(offlineScheduleBeginEnabled *bool) *GetAPIEnvironmentsV3Params {
	o.SetOfflineScheduleBeginEnabled(offlineScheduleBeginEnabled)
	return o
}

// SetOfflineScheduleBeginEnabled adds the offlineScheduleBeginEnabled to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetOfflineScheduleBeginEnabled(offlineScheduleBeginEnabled *bool) {
	o.OfflineScheduleBeginEnabled = offlineScheduleBeginEnabled
}

// WithOfflineScheduleBeginTime adds the offlineScheduleBeginTime to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithOfflineScheduleBeginTime(offlineScheduleBeginTime *strfmt.DateTime) *GetAPIEnvironmentsV3Params {
	o.SetOfflineScheduleBeginTime(offlineScheduleBeginTime)
	return o
}

// SetOfflineScheduleBeginTime adds the offlineScheduleBeginTime to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetOfflineScheduleBeginTime(offlineScheduleBeginTime *strfmt.DateTime) {
	o.OfflineScheduleBeginTime = offlineScheduleBeginTime
}

// WithOfflineScheduleEndEnabled adds the offlineScheduleEndEnabled to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithOfflineScheduleEndEnabled(offlineScheduleEndEnabled *bool) *GetAPIEnvironmentsV3Params {
	o.SetOfflineScheduleEndEnabled(offlineScheduleEndEnabled)
	return o
}

// SetOfflineScheduleEndEnabled adds the offlineScheduleEndEnabled to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetOfflineScheduleEndEnabled(offlineScheduleEndEnabled *bool) {
	o.OfflineScheduleEndEnabled = offlineScheduleEndEnabled
}

// WithOfflineScheduleEndTime adds the offlineScheduleEndTime to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithOfflineScheduleEndTime(offlineScheduleEndTime *strfmt.DateTime) *GetAPIEnvironmentsV3Params {
	o.SetOfflineScheduleEndTime(offlineScheduleEndTime)
	return o
}

// SetOfflineScheduleEndTime adds the offlineScheduleEndTime to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetOfflineScheduleEndTime(offlineScheduleEndTime *strfmt.DateTime) {
	o.OfflineScheduleEndTime = offlineScheduleEndTime
}

// WithOfflineScheduleEndWeekends adds the offlineScheduleEndWeekends to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithOfflineScheduleEndWeekends(offlineScheduleEndWeekends *bool) *GetAPIEnvironmentsV3Params {
	o.SetOfflineScheduleEndWeekends(offlineScheduleEndWeekends)
	return o
}

// SetOfflineScheduleEndWeekends adds the offlineScheduleEndWeekends to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetOfflineScheduleEndWeekends(offlineScheduleEndWeekends *bool) {
	o.OfflineScheduleEndWeekends = offlineScheduleEndWeekends
}

// WithOffset adds the offset to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithOffset(offset *int64) *GetAPIEnvironmentsV3Params {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOwner adds the owner to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithOwner(owner *string) *GetAPIEnvironmentsV3Params {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetOwner(owner *string) {
	o.Owner = owner
}

// WithPagerdutyIntegration adds the pagerdutyIntegration to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithPagerdutyIntegration(pagerdutyIntegration *string) *GetAPIEnvironmentsV3Params {
	o.SetPagerdutyIntegration(pagerdutyIntegration)
	return o
}

// SetPagerdutyIntegration adds the pagerdutyIntegration to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetPagerdutyIntegration(pagerdutyIntegration *string) {
	o.PagerdutyIntegration = pagerdutyIntegration
}

// WithPreventDeletion adds the preventDeletion to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithPreventDeletion(preventDeletion *bool) *GetAPIEnvironmentsV3Params {
	o.SetPreventDeletion(preventDeletion)
	return o
}

// SetPreventDeletion adds the preventDeletion to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetPreventDeletion(preventDeletion *bool) {
	o.PreventDeletion = preventDeletion
}

// WithRequiredRole adds the requiredRole to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithRequiredRole(requiredRole *string) *GetAPIEnvironmentsV3Params {
	o.SetRequiredRole(requiredRole)
	return o
}

// SetRequiredRole adds the requiredRole to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetRequiredRole(requiredRole *string) {
	o.RequiredRole = requiredRole
}

// WithRequiresSuitability adds the requiresSuitability to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithRequiresSuitability(requiresSuitability *bool) *GetAPIEnvironmentsV3Params {
	o.SetRequiresSuitability(requiresSuitability)
	return o
}

// SetRequiresSuitability adds the requiresSuitability to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetRequiresSuitability(requiresSuitability *bool) {
	o.RequiresSuitability = requiresSuitability
}

// WithServiceBannerBucket adds the serviceBannerBucket to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithServiceBannerBucket(serviceBannerBucket *string) *GetAPIEnvironmentsV3Params {
	o.SetServiceBannerBucket(serviceBannerBucket)
	return o
}

// SetServiceBannerBucket adds the serviceBannerBucket to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetServiceBannerBucket(serviceBannerBucket *string) {
	o.ServiceBannerBucket = serviceBannerBucket
}

// WithTemplateEnvironment adds the templateEnvironment to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithTemplateEnvironment(templateEnvironment *string) *GetAPIEnvironmentsV3Params {
	o.SetTemplateEnvironment(templateEnvironment)
	return o
}

// SetTemplateEnvironment adds the templateEnvironment to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetTemplateEnvironment(templateEnvironment *string) {
	o.TemplateEnvironment = templateEnvironment
}

// WithUniqueResourcePrefix adds the uniqueResourcePrefix to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithUniqueResourcePrefix(uniqueResourcePrefix *string) *GetAPIEnvironmentsV3Params {
	o.SetUniqueResourcePrefix(uniqueResourcePrefix)
	return o
}

// SetUniqueResourcePrefix adds the uniqueResourcePrefix to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetUniqueResourcePrefix(uniqueResourcePrefix *string) {
	o.UniqueResourcePrefix = uniqueResourcePrefix
}

// WithUpdatedAt adds the updatedAt to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithUpdatedAt(updatedAt *strfmt.DateTime) *GetAPIEnvironmentsV3Params {
	o.SetUpdatedAt(updatedAt)
	return o
}

// SetUpdatedAt adds the updatedAt to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetUpdatedAt(updatedAt *strfmt.DateTime) {
	o.UpdatedAt = updatedAt
}

// WithValuesName adds the valuesName to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) WithValuesName(valuesName *string) *GetAPIEnvironmentsV3Params {
	o.SetValuesName(valuesName)
	return o
}

// SetValuesName adds the valuesName to the get API environments v3 params
func (o *GetAPIEnvironmentsV3Params) SetValuesName(valuesName *string) {
	o.ValuesName = valuesName
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIEnvironmentsV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PactIdentifier != nil {

		// query param PactIdentifier
		var qrPactIdentifier string

		if o.PactIdentifier != nil {
			qrPactIdentifier = *o.PactIdentifier
		}
		qPactIdentifier := qrPactIdentifier
		if qPactIdentifier != "" {

			if err := r.SetQueryParam("PactIdentifier", qPactIdentifier); err != nil {
				return err
			}
		}
	}

	if o.AutoPopulateChartReleases != nil {

		// query param autoPopulateChartReleases
		var qrAutoPopulateChartReleases bool

		if o.AutoPopulateChartReleases != nil {
			qrAutoPopulateChartReleases = *o.AutoPopulateChartReleases
		}
		qAutoPopulateChartReleases := swag.FormatBool(qrAutoPopulateChartReleases)
		if qAutoPopulateChartReleases != "" {

			if err := r.SetQueryParam("autoPopulateChartReleases", qAutoPopulateChartReleases); err != nil {
				return err
			}
		}
	}

	if o.Base != nil {

		// query param base
		var qrBase string

		if o.Base != nil {
			qrBase = *o.Base
		}
		qBase := qrBase
		if qBase != "" {

			if err := r.SetQueryParam("base", qBase); err != nil {
				return err
			}
		}
	}

	if o.BaseDomain != nil {

		// query param baseDomain
		var qrBaseDomain string

		if o.BaseDomain != nil {
			qrBaseDomain = *o.BaseDomain
		}
		qBaseDomain := qrBaseDomain
		if qBaseDomain != "" {

			if err := r.SetQueryParam("baseDomain", qBaseDomain); err != nil {
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

	if o.DefaultCluster != nil {

		// query param defaultCluster
		var qrDefaultCluster string

		if o.DefaultCluster != nil {
			qrDefaultCluster = *o.DefaultCluster
		}
		qDefaultCluster := qrDefaultCluster
		if qDefaultCluster != "" {

			if err := r.SetQueryParam("defaultCluster", qDefaultCluster); err != nil {
				return err
			}
		}
	}

	if o.DefaultNamespace != nil {

		// query param defaultNamespace
		var qrDefaultNamespace string

		if o.DefaultNamespace != nil {
			qrDefaultNamespace = *o.DefaultNamespace
		}
		qDefaultNamespace := qrDefaultNamespace
		if qDefaultNamespace != "" {

			if err := r.SetQueryParam("defaultNamespace", qDefaultNamespace); err != nil {
				return err
			}
		}
	}

	if o.DeleteAfter != nil {

		// query param deleteAfter
		var qrDeleteAfter strfmt.DateTime

		if o.DeleteAfter != nil {
			qrDeleteAfter = *o.DeleteAfter
		}
		qDeleteAfter := qrDeleteAfter.String()
		if qDeleteAfter != "" {

			if err := r.SetQueryParam("deleteAfter", qDeleteAfter); err != nil {
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

	if o.EnableJanitor != nil {

		// query param enableJanitor
		var qrEnableJanitor bool

		if o.EnableJanitor != nil {
			qrEnableJanitor = *o.EnableJanitor
		}
		qEnableJanitor := swag.FormatBool(qrEnableJanitor)
		if qEnableJanitor != "" {

			if err := r.SetQueryParam("enableJanitor", qEnableJanitor); err != nil {
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

	if o.Lifecycle != nil {

		// query param lifecycle
		var qrLifecycle string

		if o.Lifecycle != nil {
			qrLifecycle = *o.Lifecycle
		}
		qLifecycle := qrLifecycle
		if qLifecycle != "" {

			if err := r.SetQueryParam("lifecycle", qLifecycle); err != nil {
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

	if o.NamePrefixesDomain != nil {

		// query param namePrefixesDomain
		var qrNamePrefixesDomain bool

		if o.NamePrefixesDomain != nil {
			qrNamePrefixesDomain = *o.NamePrefixesDomain
		}
		qNamePrefixesDomain := swag.FormatBool(qrNamePrefixesDomain)
		if qNamePrefixesDomain != "" {

			if err := r.SetQueryParam("namePrefixesDomain", qNamePrefixesDomain); err != nil {
				return err
			}
		}
	}

	if o.Offline != nil {

		// query param offline
		var qrOffline bool

		if o.Offline != nil {
			qrOffline = *o.Offline
		}
		qOffline := swag.FormatBool(qrOffline)
		if qOffline != "" {

			if err := r.SetQueryParam("offline", qOffline); err != nil {
				return err
			}
		}
	}

	if o.OfflineScheduleBeginEnabled != nil {

		// query param offlineScheduleBeginEnabled
		var qrOfflineScheduleBeginEnabled bool

		if o.OfflineScheduleBeginEnabled != nil {
			qrOfflineScheduleBeginEnabled = *o.OfflineScheduleBeginEnabled
		}
		qOfflineScheduleBeginEnabled := swag.FormatBool(qrOfflineScheduleBeginEnabled)
		if qOfflineScheduleBeginEnabled != "" {

			if err := r.SetQueryParam("offlineScheduleBeginEnabled", qOfflineScheduleBeginEnabled); err != nil {
				return err
			}
		}
	}

	if o.OfflineScheduleBeginTime != nil {

		// query param offlineScheduleBeginTime
		var qrOfflineScheduleBeginTime strfmt.DateTime

		if o.OfflineScheduleBeginTime != nil {
			qrOfflineScheduleBeginTime = *o.OfflineScheduleBeginTime
		}
		qOfflineScheduleBeginTime := qrOfflineScheduleBeginTime.String()
		if qOfflineScheduleBeginTime != "" {

			if err := r.SetQueryParam("offlineScheduleBeginTime", qOfflineScheduleBeginTime); err != nil {
				return err
			}
		}
	}

	if o.OfflineScheduleEndEnabled != nil {

		// query param offlineScheduleEndEnabled
		var qrOfflineScheduleEndEnabled bool

		if o.OfflineScheduleEndEnabled != nil {
			qrOfflineScheduleEndEnabled = *o.OfflineScheduleEndEnabled
		}
		qOfflineScheduleEndEnabled := swag.FormatBool(qrOfflineScheduleEndEnabled)
		if qOfflineScheduleEndEnabled != "" {

			if err := r.SetQueryParam("offlineScheduleEndEnabled", qOfflineScheduleEndEnabled); err != nil {
				return err
			}
		}
	}

	if o.OfflineScheduleEndTime != nil {

		// query param offlineScheduleEndTime
		var qrOfflineScheduleEndTime strfmt.DateTime

		if o.OfflineScheduleEndTime != nil {
			qrOfflineScheduleEndTime = *o.OfflineScheduleEndTime
		}
		qOfflineScheduleEndTime := qrOfflineScheduleEndTime.String()
		if qOfflineScheduleEndTime != "" {

			if err := r.SetQueryParam("offlineScheduleEndTime", qOfflineScheduleEndTime); err != nil {
				return err
			}
		}
	}

	if o.OfflineScheduleEndWeekends != nil {

		// query param offlineScheduleEndWeekends
		var qrOfflineScheduleEndWeekends bool

		if o.OfflineScheduleEndWeekends != nil {
			qrOfflineScheduleEndWeekends = *o.OfflineScheduleEndWeekends
		}
		qOfflineScheduleEndWeekends := swag.FormatBool(qrOfflineScheduleEndWeekends)
		if qOfflineScheduleEndWeekends != "" {

			if err := r.SetQueryParam("offlineScheduleEndWeekends", qOfflineScheduleEndWeekends); err != nil {
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

	if o.Owner != nil {

		// query param owner
		var qrOwner string

		if o.Owner != nil {
			qrOwner = *o.Owner
		}
		qOwner := qrOwner
		if qOwner != "" {

			if err := r.SetQueryParam("owner", qOwner); err != nil {
				return err
			}
		}
	}

	if o.PagerdutyIntegration != nil {

		// query param pagerdutyIntegration
		var qrPagerdutyIntegration string

		if o.PagerdutyIntegration != nil {
			qrPagerdutyIntegration = *o.PagerdutyIntegration
		}
		qPagerdutyIntegration := qrPagerdutyIntegration
		if qPagerdutyIntegration != "" {

			if err := r.SetQueryParam("pagerdutyIntegration", qPagerdutyIntegration); err != nil {
				return err
			}
		}
	}

	if o.PreventDeletion != nil {

		// query param preventDeletion
		var qrPreventDeletion bool

		if o.PreventDeletion != nil {
			qrPreventDeletion = *o.PreventDeletion
		}
		qPreventDeletion := swag.FormatBool(qrPreventDeletion)
		if qPreventDeletion != "" {

			if err := r.SetQueryParam("preventDeletion", qPreventDeletion); err != nil {
				return err
			}
		}
	}

	if o.RequiredRole != nil {

		// query param requiredRole
		var qrRequiredRole string

		if o.RequiredRole != nil {
			qrRequiredRole = *o.RequiredRole
		}
		qRequiredRole := qrRequiredRole
		if qRequiredRole != "" {

			if err := r.SetQueryParam("requiredRole", qRequiredRole); err != nil {
				return err
			}
		}
	}

	if o.RequiresSuitability != nil {

		// query param requiresSuitability
		var qrRequiresSuitability bool

		if o.RequiresSuitability != nil {
			qrRequiresSuitability = *o.RequiresSuitability
		}
		qRequiresSuitability := swag.FormatBool(qrRequiresSuitability)
		if qRequiresSuitability != "" {

			if err := r.SetQueryParam("requiresSuitability", qRequiresSuitability); err != nil {
				return err
			}
		}
	}

	if o.ServiceBannerBucket != nil {

		// query param serviceBannerBucket
		var qrServiceBannerBucket string

		if o.ServiceBannerBucket != nil {
			qrServiceBannerBucket = *o.ServiceBannerBucket
		}
		qServiceBannerBucket := qrServiceBannerBucket
		if qServiceBannerBucket != "" {

			if err := r.SetQueryParam("serviceBannerBucket", qServiceBannerBucket); err != nil {
				return err
			}
		}
	}

	if o.TemplateEnvironment != nil {

		// query param templateEnvironment
		var qrTemplateEnvironment string

		if o.TemplateEnvironment != nil {
			qrTemplateEnvironment = *o.TemplateEnvironment
		}
		qTemplateEnvironment := qrTemplateEnvironment
		if qTemplateEnvironment != "" {

			if err := r.SetQueryParam("templateEnvironment", qTemplateEnvironment); err != nil {
				return err
			}
		}
	}

	if o.UniqueResourcePrefix != nil {

		// query param uniqueResourcePrefix
		var qrUniqueResourcePrefix string

		if o.UniqueResourcePrefix != nil {
			qrUniqueResourcePrefix = *o.UniqueResourcePrefix
		}
		qUniqueResourcePrefix := qrUniqueResourcePrefix
		if qUniqueResourcePrefix != "" {

			if err := r.SetQueryParam("uniqueResourcePrefix", qUniqueResourcePrefix); err != nil {
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

	if o.ValuesName != nil {

		// query param valuesName
		var qrValuesName string

		if o.ValuesName != nil {
			qrValuesName = *o.ValuesName
		}
		qValuesName := qrValuesName
		if qValuesName != "" {

			if err := r.SetQueryParam("valuesName", qValuesName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
