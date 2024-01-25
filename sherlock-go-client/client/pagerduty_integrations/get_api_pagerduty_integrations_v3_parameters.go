// Code generated by go-swagger; DO NOT EDIT.

package pagerduty_integrations

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

// NewGetAPIPagerdutyIntegrationsV3Params creates a new GetAPIPagerdutyIntegrationsV3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIPagerdutyIntegrationsV3Params() *GetAPIPagerdutyIntegrationsV3Params {
	return &GetAPIPagerdutyIntegrationsV3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIPagerdutyIntegrationsV3ParamsWithTimeout creates a new GetAPIPagerdutyIntegrationsV3Params object
// with the ability to set a timeout on a request.
func NewGetAPIPagerdutyIntegrationsV3ParamsWithTimeout(timeout time.Duration) *GetAPIPagerdutyIntegrationsV3Params {
	return &GetAPIPagerdutyIntegrationsV3Params{
		timeout: timeout,
	}
}

// NewGetAPIPagerdutyIntegrationsV3ParamsWithContext creates a new GetAPIPagerdutyIntegrationsV3Params object
// with the ability to set a context for a request.
func NewGetAPIPagerdutyIntegrationsV3ParamsWithContext(ctx context.Context) *GetAPIPagerdutyIntegrationsV3Params {
	return &GetAPIPagerdutyIntegrationsV3Params{
		Context: ctx,
	}
}

// NewGetAPIPagerdutyIntegrationsV3ParamsWithHTTPClient creates a new GetAPIPagerdutyIntegrationsV3Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIPagerdutyIntegrationsV3ParamsWithHTTPClient(client *http.Client) *GetAPIPagerdutyIntegrationsV3Params {
	return &GetAPIPagerdutyIntegrationsV3Params{
		HTTPClient: client,
	}
}

/* GetAPIPagerdutyIntegrationsV3Params contains all the parameters to send to the API endpoint
   for the get API pagerduty integrations v3 operation.

   Typically these are written to a http.Request.
*/
type GetAPIPagerdutyIntegrationsV3Params struct {

	// CreatedAt.
	//
	// Format: date-time
	CreatedAt *strfmt.DateTime

	// ID.
	ID *int64

	/* Limit.

	   Control how many PagerdutyIntegrations are returned (default 0, meaning all)
	*/
	Limit *int64

	// Name.
	Name *string

	/* Offset.

	   Control the offset for the returned PagerdutyIntegrations (default 0)
	*/
	Offset *int64

	// PagerdutyID.
	PagerdutyID *string

	// Type.
	Type *string

	// UpdatedAt.
	//
	// Format: date-time
	UpdatedAt *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API pagerduty integrations v3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIPagerdutyIntegrationsV3Params) WithDefaults() *GetAPIPagerdutyIntegrationsV3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API pagerduty integrations v3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIPagerdutyIntegrationsV3Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API pagerduty integrations v3 params
func (o *GetAPIPagerdutyIntegrationsV3Params) WithTimeout(timeout time.Duration) *GetAPIPagerdutyIntegrationsV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API pagerduty integrations v3 params
func (o *GetAPIPagerdutyIntegrationsV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API pagerduty integrations v3 params
func (o *GetAPIPagerdutyIntegrationsV3Params) WithContext(ctx context.Context) *GetAPIPagerdutyIntegrationsV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API pagerduty integrations v3 params
func (o *GetAPIPagerdutyIntegrationsV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API pagerduty integrations v3 params
func (o *GetAPIPagerdutyIntegrationsV3Params) WithHTTPClient(client *http.Client) *GetAPIPagerdutyIntegrationsV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API pagerduty integrations v3 params
func (o *GetAPIPagerdutyIntegrationsV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreatedAt adds the createdAt to the get API pagerduty integrations v3 params
func (o *GetAPIPagerdutyIntegrationsV3Params) WithCreatedAt(createdAt *strfmt.DateTime) *GetAPIPagerdutyIntegrationsV3Params {
	o.SetCreatedAt(createdAt)
	return o
}

// SetCreatedAt adds the createdAt to the get API pagerduty integrations v3 params
func (o *GetAPIPagerdutyIntegrationsV3Params) SetCreatedAt(createdAt *strfmt.DateTime) {
	o.CreatedAt = createdAt
}

// WithID adds the id to the get API pagerduty integrations v3 params
func (o *GetAPIPagerdutyIntegrationsV3Params) WithID(id *int64) *GetAPIPagerdutyIntegrationsV3Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the get API pagerduty integrations v3 params
func (o *GetAPIPagerdutyIntegrationsV3Params) SetID(id *int64) {
	o.ID = id
}

// WithLimit adds the limit to the get API pagerduty integrations v3 params
func (o *GetAPIPagerdutyIntegrationsV3Params) WithLimit(limit *int64) *GetAPIPagerdutyIntegrationsV3Params {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get API pagerduty integrations v3 params
func (o *GetAPIPagerdutyIntegrationsV3Params) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the get API pagerduty integrations v3 params
func (o *GetAPIPagerdutyIntegrationsV3Params) WithName(name *string) *GetAPIPagerdutyIntegrationsV3Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the get API pagerduty integrations v3 params
func (o *GetAPIPagerdutyIntegrationsV3Params) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the get API pagerduty integrations v3 params
func (o *GetAPIPagerdutyIntegrationsV3Params) WithOffset(offset *int64) *GetAPIPagerdutyIntegrationsV3Params {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get API pagerduty integrations v3 params
func (o *GetAPIPagerdutyIntegrationsV3Params) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPagerdutyID adds the pagerdutyID to the get API pagerduty integrations v3 params
func (o *GetAPIPagerdutyIntegrationsV3Params) WithPagerdutyID(pagerdutyID *string) *GetAPIPagerdutyIntegrationsV3Params {
	o.SetPagerdutyID(pagerdutyID)
	return o
}

// SetPagerdutyID adds the pagerdutyId to the get API pagerduty integrations v3 params
func (o *GetAPIPagerdutyIntegrationsV3Params) SetPagerdutyID(pagerdutyID *string) {
	o.PagerdutyID = pagerdutyID
}

// WithType adds the typeVar to the get API pagerduty integrations v3 params
func (o *GetAPIPagerdutyIntegrationsV3Params) WithType(typeVar *string) *GetAPIPagerdutyIntegrationsV3Params {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get API pagerduty integrations v3 params
func (o *GetAPIPagerdutyIntegrationsV3Params) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithUpdatedAt adds the updatedAt to the get API pagerduty integrations v3 params
func (o *GetAPIPagerdutyIntegrationsV3Params) WithUpdatedAt(updatedAt *strfmt.DateTime) *GetAPIPagerdutyIntegrationsV3Params {
	o.SetUpdatedAt(updatedAt)
	return o
}

// SetUpdatedAt adds the updatedAt to the get API pagerduty integrations v3 params
func (o *GetAPIPagerdutyIntegrationsV3Params) SetUpdatedAt(updatedAt *strfmt.DateTime) {
	o.UpdatedAt = updatedAt
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIPagerdutyIntegrationsV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.PagerdutyID != nil {

		// query param pagerdutyID
		var qrPagerdutyID string

		if o.PagerdutyID != nil {
			qrPagerdutyID = *o.PagerdutyID
		}
		qPagerdutyID := qrPagerdutyID
		if qPagerdutyID != "" {

			if err := r.SetQueryParam("pagerdutyID", qPagerdutyID); err != nil {
				return err
			}
		}
	}

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
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
