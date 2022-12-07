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

// NewGetAPIV2EnvironmentsParams creates a new GetAPIV2EnvironmentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIV2EnvironmentsParams() *GetAPIV2EnvironmentsParams {
	return &GetAPIV2EnvironmentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV2EnvironmentsParamsWithTimeout creates a new GetAPIV2EnvironmentsParams object
// with the ability to set a timeout on a request.
func NewGetAPIV2EnvironmentsParamsWithTimeout(timeout time.Duration) *GetAPIV2EnvironmentsParams {
	return &GetAPIV2EnvironmentsParams{
		timeout: timeout,
	}
}

// NewGetAPIV2EnvironmentsParamsWithContext creates a new GetAPIV2EnvironmentsParams object
// with the ability to set a context for a request.
func NewGetAPIV2EnvironmentsParamsWithContext(ctx context.Context) *GetAPIV2EnvironmentsParams {
	return &GetAPIV2EnvironmentsParams{
		Context: ctx,
	}
}

// NewGetAPIV2EnvironmentsParamsWithHTTPClient creates a new GetAPIV2EnvironmentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIV2EnvironmentsParamsWithHTTPClient(client *http.Client) *GetAPIV2EnvironmentsParams {
	return &GetAPIV2EnvironmentsParams{
		HTTPClient: client,
	}
}

/*
GetAPIV2EnvironmentsParams contains all the parameters to send to the API endpoint

	for the get API v2 environments operation.

	Typically these are written to a http.Request.
*/
type GetAPIV2EnvironmentsParams struct {

	/* Base.

	   Required when creating
	*/
	Base *string

	// BaseDomain.
	//
	// Default: "bee.envs-terra.bio"
	BaseDomain *string

	/* ChartReleasesFromTemplate.

	   Upon creation of a dynamic environment, if this is true the template's chart releases will be copied to the new environment

	   Default: true
	*/
	ChartReleasesFromTemplate *bool

	// CreatedAt.
	//
	// Format: date-time
	CreatedAt *strfmt.DateTime

	// DefaultCluster.
	DefaultCluster *string

	/* DefaultFirecloudDevelopRef.

	   should be the environment branch for live envs. Is usually dev for template/dynamic but not necessarily

	   Default: "dev"
	*/
	DefaultFirecloudDevelopRef *string

	/* DefaultNamespace.

	   When creating, will be calculated if left empty
	*/
	DefaultNamespace *string

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

	   An optional limit to the number of entries returned
	*/
	Limit *int64

	/* Name.

	   When creating, will be calculated if dynamic, required otherwise
	*/
	Name *string

	/* NamePrefix.

	   Used for dynamic environment name generation only, to override using the owner email handle and template name
	*/
	NamePrefix *string

	// NamePrefixesDomain.
	//
	// Default: true
	NamePrefixesDomain *bool

	/* Owner.

	   When creating, will be set to your email
	*/
	Owner *string

	/* PreventDeletion.

	   Used to protect specific BEEs from deletion (thelma checks this field)
	*/
	PreventDeletion *bool

	// RequiresSuitability.
	RequiresSuitability *bool

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

	// ValuesName.
	ValuesName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API v2 environments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2EnvironmentsParams) WithDefaults() *GetAPIV2EnvironmentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API v2 environments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2EnvironmentsParams) SetDefaults() {
	var (
		baseDomainDefault = string("bee.envs-terra.bio")

		chartReleasesFromTemplateDefault = bool(true)

		defaultFirecloudDevelopRefDefault = string("dev")

		helmfileRefDefault = string("HEAD")

		lifecycleDefault = string("dynamic")

		namePrefixesDomainDefault = bool(true)

		preventDeletionDefault = bool(false)

		requiresSuitabilityDefault = bool(false)
	)

	val := GetAPIV2EnvironmentsParams{
		BaseDomain:                 &baseDomainDefault,
		ChartReleasesFromTemplate:  &chartReleasesFromTemplateDefault,
		DefaultFirecloudDevelopRef: &defaultFirecloudDevelopRefDefault,
		HelmfileRef:                &helmfileRefDefault,
		Lifecycle:                  &lifecycleDefault,
		NamePrefixesDomain:         &namePrefixesDomainDefault,
		PreventDeletion:            &preventDeletionDefault,
		RequiresSuitability:        &requiresSuitabilityDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) WithTimeout(timeout time.Duration) *GetAPIV2EnvironmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) WithContext(ctx context.Context) *GetAPIV2EnvironmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) WithHTTPClient(client *http.Client) *GetAPIV2EnvironmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBase adds the base to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) WithBase(base *string) *GetAPIV2EnvironmentsParams {
	o.SetBase(base)
	return o
}

// SetBase adds the base to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) SetBase(base *string) {
	o.Base = base
}

// WithBaseDomain adds the baseDomain to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) WithBaseDomain(baseDomain *string) *GetAPIV2EnvironmentsParams {
	o.SetBaseDomain(baseDomain)
	return o
}

// SetBaseDomain adds the baseDomain to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) SetBaseDomain(baseDomain *string) {
	o.BaseDomain = baseDomain
}

// WithChartReleasesFromTemplate adds the chartReleasesFromTemplate to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) WithChartReleasesFromTemplate(chartReleasesFromTemplate *bool) *GetAPIV2EnvironmentsParams {
	o.SetChartReleasesFromTemplate(chartReleasesFromTemplate)
	return o
}

// SetChartReleasesFromTemplate adds the chartReleasesFromTemplate to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) SetChartReleasesFromTemplate(chartReleasesFromTemplate *bool) {
	o.ChartReleasesFromTemplate = chartReleasesFromTemplate
}

// WithCreatedAt adds the createdAt to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) WithCreatedAt(createdAt *strfmt.DateTime) *GetAPIV2EnvironmentsParams {
	o.SetCreatedAt(createdAt)
	return o
}

// SetCreatedAt adds the createdAt to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) SetCreatedAt(createdAt *strfmt.DateTime) {
	o.CreatedAt = createdAt
}

// WithDefaultCluster adds the defaultCluster to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) WithDefaultCluster(defaultCluster *string) *GetAPIV2EnvironmentsParams {
	o.SetDefaultCluster(defaultCluster)
	return o
}

// SetDefaultCluster adds the defaultCluster to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) SetDefaultCluster(defaultCluster *string) {
	o.DefaultCluster = defaultCluster
}

// WithDefaultFirecloudDevelopRef adds the defaultFirecloudDevelopRef to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) WithDefaultFirecloudDevelopRef(defaultFirecloudDevelopRef *string) *GetAPIV2EnvironmentsParams {
	o.SetDefaultFirecloudDevelopRef(defaultFirecloudDevelopRef)
	return o
}

// SetDefaultFirecloudDevelopRef adds the defaultFirecloudDevelopRef to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) SetDefaultFirecloudDevelopRef(defaultFirecloudDevelopRef *string) {
	o.DefaultFirecloudDevelopRef = defaultFirecloudDevelopRef
}

// WithDefaultNamespace adds the defaultNamespace to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) WithDefaultNamespace(defaultNamespace *string) *GetAPIV2EnvironmentsParams {
	o.SetDefaultNamespace(defaultNamespace)
	return o
}

// SetDefaultNamespace adds the defaultNamespace to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) SetDefaultNamespace(defaultNamespace *string) {
	o.DefaultNamespace = defaultNamespace
}

// WithHelmfileRef adds the helmfileRef to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) WithHelmfileRef(helmfileRef *string) *GetAPIV2EnvironmentsParams {
	o.SetHelmfileRef(helmfileRef)
	return o
}

// SetHelmfileRef adds the helmfileRef to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) SetHelmfileRef(helmfileRef *string) {
	o.HelmfileRef = helmfileRef
}

// WithID adds the id to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) WithID(id *int64) *GetAPIV2EnvironmentsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) SetID(id *int64) {
	o.ID = id
}

// WithLifecycle adds the lifecycle to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) WithLifecycle(lifecycle *string) *GetAPIV2EnvironmentsParams {
	o.SetLifecycle(lifecycle)
	return o
}

// SetLifecycle adds the lifecycle to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) SetLifecycle(lifecycle *string) {
	o.Lifecycle = lifecycle
}

// WithLimit adds the limit to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) WithLimit(limit *int64) *GetAPIV2EnvironmentsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) WithName(name *string) *GetAPIV2EnvironmentsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) SetName(name *string) {
	o.Name = name
}

// WithNamePrefix adds the namePrefix to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) WithNamePrefix(namePrefix *string) *GetAPIV2EnvironmentsParams {
	o.SetNamePrefix(namePrefix)
	return o
}

// SetNamePrefix adds the namePrefix to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) SetNamePrefix(namePrefix *string) {
	o.NamePrefix = namePrefix
}

// WithNamePrefixesDomain adds the namePrefixesDomain to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) WithNamePrefixesDomain(namePrefixesDomain *bool) *GetAPIV2EnvironmentsParams {
	o.SetNamePrefixesDomain(namePrefixesDomain)
	return o
}

// SetNamePrefixesDomain adds the namePrefixesDomain to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) SetNamePrefixesDomain(namePrefixesDomain *bool) {
	o.NamePrefixesDomain = namePrefixesDomain
}

// WithOwner adds the owner to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) WithOwner(owner *string) *GetAPIV2EnvironmentsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) SetOwner(owner *string) {
	o.Owner = owner
}

// WithPreventDeletion adds the preventDeletion to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) WithPreventDeletion(preventDeletion *bool) *GetAPIV2EnvironmentsParams {
	o.SetPreventDeletion(preventDeletion)
	return o
}

// SetPreventDeletion adds the preventDeletion to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) SetPreventDeletion(preventDeletion *bool) {
	o.PreventDeletion = preventDeletion
}

// WithRequiresSuitability adds the requiresSuitability to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) WithRequiresSuitability(requiresSuitability *bool) *GetAPIV2EnvironmentsParams {
	o.SetRequiresSuitability(requiresSuitability)
	return o
}

// SetRequiresSuitability adds the requiresSuitability to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) SetRequiresSuitability(requiresSuitability *bool) {
	o.RequiresSuitability = requiresSuitability
}

// WithTemplateEnvironment adds the templateEnvironment to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) WithTemplateEnvironment(templateEnvironment *string) *GetAPIV2EnvironmentsParams {
	o.SetTemplateEnvironment(templateEnvironment)
	return o
}

// SetTemplateEnvironment adds the templateEnvironment to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) SetTemplateEnvironment(templateEnvironment *string) {
	o.TemplateEnvironment = templateEnvironment
}

// WithUniqueResourcePrefix adds the uniqueResourcePrefix to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) WithUniqueResourcePrefix(uniqueResourcePrefix *string) *GetAPIV2EnvironmentsParams {
	o.SetUniqueResourcePrefix(uniqueResourcePrefix)
	return o
}

// SetUniqueResourcePrefix adds the uniqueResourcePrefix to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) SetUniqueResourcePrefix(uniqueResourcePrefix *string) {
	o.UniqueResourcePrefix = uniqueResourcePrefix
}

// WithUpdatedAt adds the updatedAt to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) WithUpdatedAt(updatedAt *strfmt.DateTime) *GetAPIV2EnvironmentsParams {
	o.SetUpdatedAt(updatedAt)
	return o
}

// SetUpdatedAt adds the updatedAt to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) SetUpdatedAt(updatedAt *strfmt.DateTime) {
	o.UpdatedAt = updatedAt
}

// WithValuesName adds the valuesName to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) WithValuesName(valuesName *string) *GetAPIV2EnvironmentsParams {
	o.SetValuesName(valuesName)
	return o
}

// SetValuesName adds the valuesName to the get API v2 environments params
func (o *GetAPIV2EnvironmentsParams) SetValuesName(valuesName *string) {
	o.ValuesName = valuesName
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV2EnvironmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.ChartReleasesFromTemplate != nil {

		// query param chartReleasesFromTemplate
		var qrChartReleasesFromTemplate bool

		if o.ChartReleasesFromTemplate != nil {
			qrChartReleasesFromTemplate = *o.ChartReleasesFromTemplate
		}
		qChartReleasesFromTemplate := swag.FormatBool(qrChartReleasesFromTemplate)
		if qChartReleasesFromTemplate != "" {

			if err := r.SetQueryParam("chartReleasesFromTemplate", qChartReleasesFromTemplate); err != nil {
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

	if o.DefaultFirecloudDevelopRef != nil {

		// query param defaultFirecloudDevelopRef
		var qrDefaultFirecloudDevelopRef string

		if o.DefaultFirecloudDevelopRef != nil {
			qrDefaultFirecloudDevelopRef = *o.DefaultFirecloudDevelopRef
		}
		qDefaultFirecloudDevelopRef := qrDefaultFirecloudDevelopRef
		if qDefaultFirecloudDevelopRef != "" {

			if err := r.SetQueryParam("defaultFirecloudDevelopRef", qDefaultFirecloudDevelopRef); err != nil {
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

	if o.NamePrefix != nil {

		// query param namePrefix
		var qrNamePrefix string

		if o.NamePrefix != nil {
			qrNamePrefix = *o.NamePrefix
		}
		qNamePrefix := qrNamePrefix
		if qNamePrefix != "" {

			if err := r.SetQueryParam("namePrefix", qNamePrefix); err != nil {
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
