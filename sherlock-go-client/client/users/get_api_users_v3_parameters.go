// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewGetAPIUsersV3Params creates a new GetAPIUsersV3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIUsersV3Params() *GetAPIUsersV3Params {
	return &GetAPIUsersV3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIUsersV3ParamsWithTimeout creates a new GetAPIUsersV3Params object
// with the ability to set a timeout on a request.
func NewGetAPIUsersV3ParamsWithTimeout(timeout time.Duration) *GetAPIUsersV3Params {
	return &GetAPIUsersV3Params{
		timeout: timeout,
	}
}

// NewGetAPIUsersV3ParamsWithContext creates a new GetAPIUsersV3Params object
// with the ability to set a context for a request.
func NewGetAPIUsersV3ParamsWithContext(ctx context.Context) *GetAPIUsersV3Params {
	return &GetAPIUsersV3Params{
		Context: ctx,
	}
}

// NewGetAPIUsersV3ParamsWithHTTPClient creates a new GetAPIUsersV3Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIUsersV3ParamsWithHTTPClient(client *http.Client) *GetAPIUsersV3Params {
	return &GetAPIUsersV3Params{
		HTTPClient: client,
	}
}

/* GetAPIUsersV3Params contains all the parameters to send to the API endpoint
   for the get API users v3 operation.

   Typically these are written to a http.Request.
*/
type GetAPIUsersV3Params struct {

	// CreatedAt.
	//
	// Format: date-time
	CreatedAt *strfmt.DateTime

	// Email.
	Email *string

	// GithubID.
	GithubID *string

	// GithubUsername.
	GithubUsername *string

	// GoogleID.
	GoogleID *string

	// ID.
	ID *int64

	/* Limit.

	   Control how many Users are returned (default 0, no limit)
	*/
	Limit *int64

	// Name.
	Name *string

	// NameFrom.
	NameFrom *string

	/* NameInferredFromGithub.

	     Controls whether Sherlock should automatically update the user's name based on a connected GitHub identity.
	Will be set to true if the user account has no name and a GitHub account is linked.
	*/
	NameInferredFromGithub *bool

	/* Offset.

	   Control the offset for the returned Users (default 0)
	*/
	Offset *int64

	// SlackID.
	SlackID *string

	// SlackUsername.
	SlackUsername *string

	/* SuitabilityDescription.

	   Available only in responses; describes the user's production-suitability
	*/
	SuitabilityDescription *string

	/* Suitable.

	   Available only in responses; indicates whether the user is production-suitable
	*/
	Suitable *bool

	// UpdatedAt.
	//
	// Format: date-time
	UpdatedAt *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API users v3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIUsersV3Params) WithDefaults() *GetAPIUsersV3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API users v3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIUsersV3Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API users v3 params
func (o *GetAPIUsersV3Params) WithTimeout(timeout time.Duration) *GetAPIUsersV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API users v3 params
func (o *GetAPIUsersV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API users v3 params
func (o *GetAPIUsersV3Params) WithContext(ctx context.Context) *GetAPIUsersV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API users v3 params
func (o *GetAPIUsersV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API users v3 params
func (o *GetAPIUsersV3Params) WithHTTPClient(client *http.Client) *GetAPIUsersV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API users v3 params
func (o *GetAPIUsersV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreatedAt adds the createdAt to the get API users v3 params
func (o *GetAPIUsersV3Params) WithCreatedAt(createdAt *strfmt.DateTime) *GetAPIUsersV3Params {
	o.SetCreatedAt(createdAt)
	return o
}

// SetCreatedAt adds the createdAt to the get API users v3 params
func (o *GetAPIUsersV3Params) SetCreatedAt(createdAt *strfmt.DateTime) {
	o.CreatedAt = createdAt
}

// WithEmail adds the email to the get API users v3 params
func (o *GetAPIUsersV3Params) WithEmail(email *string) *GetAPIUsersV3Params {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the get API users v3 params
func (o *GetAPIUsersV3Params) SetEmail(email *string) {
	o.Email = email
}

// WithGithubID adds the githubID to the get API users v3 params
func (o *GetAPIUsersV3Params) WithGithubID(githubID *string) *GetAPIUsersV3Params {
	o.SetGithubID(githubID)
	return o
}

// SetGithubID adds the githubId to the get API users v3 params
func (o *GetAPIUsersV3Params) SetGithubID(githubID *string) {
	o.GithubID = githubID
}

// WithGithubUsername adds the githubUsername to the get API users v3 params
func (o *GetAPIUsersV3Params) WithGithubUsername(githubUsername *string) *GetAPIUsersV3Params {
	o.SetGithubUsername(githubUsername)
	return o
}

// SetGithubUsername adds the githubUsername to the get API users v3 params
func (o *GetAPIUsersV3Params) SetGithubUsername(githubUsername *string) {
	o.GithubUsername = githubUsername
}

// WithGoogleID adds the googleID to the get API users v3 params
func (o *GetAPIUsersV3Params) WithGoogleID(googleID *string) *GetAPIUsersV3Params {
	o.SetGoogleID(googleID)
	return o
}

// SetGoogleID adds the googleId to the get API users v3 params
func (o *GetAPIUsersV3Params) SetGoogleID(googleID *string) {
	o.GoogleID = googleID
}

// WithID adds the id to the get API users v3 params
func (o *GetAPIUsersV3Params) WithID(id *int64) *GetAPIUsersV3Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the get API users v3 params
func (o *GetAPIUsersV3Params) SetID(id *int64) {
	o.ID = id
}

// WithLimit adds the limit to the get API users v3 params
func (o *GetAPIUsersV3Params) WithLimit(limit *int64) *GetAPIUsersV3Params {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get API users v3 params
func (o *GetAPIUsersV3Params) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the get API users v3 params
func (o *GetAPIUsersV3Params) WithName(name *string) *GetAPIUsersV3Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the get API users v3 params
func (o *GetAPIUsersV3Params) SetName(name *string) {
	o.Name = name
}

// WithNameFrom adds the nameFrom to the get API users v3 params
func (o *GetAPIUsersV3Params) WithNameFrom(nameFrom *string) *GetAPIUsersV3Params {
	o.SetNameFrom(nameFrom)
	return o
}

// SetNameFrom adds the nameFrom to the get API users v3 params
func (o *GetAPIUsersV3Params) SetNameFrom(nameFrom *string) {
	o.NameFrom = nameFrom
}

// WithNameInferredFromGithub adds the nameInferredFromGithub to the get API users v3 params
func (o *GetAPIUsersV3Params) WithNameInferredFromGithub(nameInferredFromGithub *bool) *GetAPIUsersV3Params {
	o.SetNameInferredFromGithub(nameInferredFromGithub)
	return o
}

// SetNameInferredFromGithub adds the nameInferredFromGithub to the get API users v3 params
func (o *GetAPIUsersV3Params) SetNameInferredFromGithub(nameInferredFromGithub *bool) {
	o.NameInferredFromGithub = nameInferredFromGithub
}

// WithOffset adds the offset to the get API users v3 params
func (o *GetAPIUsersV3Params) WithOffset(offset *int64) *GetAPIUsersV3Params {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get API users v3 params
func (o *GetAPIUsersV3Params) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSlackID adds the slackID to the get API users v3 params
func (o *GetAPIUsersV3Params) WithSlackID(slackID *string) *GetAPIUsersV3Params {
	o.SetSlackID(slackID)
	return o
}

// SetSlackID adds the slackId to the get API users v3 params
func (o *GetAPIUsersV3Params) SetSlackID(slackID *string) {
	o.SlackID = slackID
}

// WithSlackUsername adds the slackUsername to the get API users v3 params
func (o *GetAPIUsersV3Params) WithSlackUsername(slackUsername *string) *GetAPIUsersV3Params {
	o.SetSlackUsername(slackUsername)
	return o
}

// SetSlackUsername adds the slackUsername to the get API users v3 params
func (o *GetAPIUsersV3Params) SetSlackUsername(slackUsername *string) {
	o.SlackUsername = slackUsername
}

// WithSuitabilityDescription adds the suitabilityDescription to the get API users v3 params
func (o *GetAPIUsersV3Params) WithSuitabilityDescription(suitabilityDescription *string) *GetAPIUsersV3Params {
	o.SetSuitabilityDescription(suitabilityDescription)
	return o
}

// SetSuitabilityDescription adds the suitabilityDescription to the get API users v3 params
func (o *GetAPIUsersV3Params) SetSuitabilityDescription(suitabilityDescription *string) {
	o.SuitabilityDescription = suitabilityDescription
}

// WithSuitable adds the suitable to the get API users v3 params
func (o *GetAPIUsersV3Params) WithSuitable(suitable *bool) *GetAPIUsersV3Params {
	o.SetSuitable(suitable)
	return o
}

// SetSuitable adds the suitable to the get API users v3 params
func (o *GetAPIUsersV3Params) SetSuitable(suitable *bool) {
	o.Suitable = suitable
}

// WithUpdatedAt adds the updatedAt to the get API users v3 params
func (o *GetAPIUsersV3Params) WithUpdatedAt(updatedAt *strfmt.DateTime) *GetAPIUsersV3Params {
	o.SetUpdatedAt(updatedAt)
	return o
}

// SetUpdatedAt adds the updatedAt to the get API users v3 params
func (o *GetAPIUsersV3Params) SetUpdatedAt(updatedAt *strfmt.DateTime) {
	o.UpdatedAt = updatedAt
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIUsersV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Email != nil {

		// query param email
		var qrEmail string

		if o.Email != nil {
			qrEmail = *o.Email
		}
		qEmail := qrEmail
		if qEmail != "" {

			if err := r.SetQueryParam("email", qEmail); err != nil {
				return err
			}
		}
	}

	if o.GithubID != nil {

		// query param githubID
		var qrGithubID string

		if o.GithubID != nil {
			qrGithubID = *o.GithubID
		}
		qGithubID := qrGithubID
		if qGithubID != "" {

			if err := r.SetQueryParam("githubID", qGithubID); err != nil {
				return err
			}
		}
	}

	if o.GithubUsername != nil {

		// query param githubUsername
		var qrGithubUsername string

		if o.GithubUsername != nil {
			qrGithubUsername = *o.GithubUsername
		}
		qGithubUsername := qrGithubUsername
		if qGithubUsername != "" {

			if err := r.SetQueryParam("githubUsername", qGithubUsername); err != nil {
				return err
			}
		}
	}

	if o.GoogleID != nil {

		// query param googleID
		var qrGoogleID string

		if o.GoogleID != nil {
			qrGoogleID = *o.GoogleID
		}
		qGoogleID := qrGoogleID
		if qGoogleID != "" {

			if err := r.SetQueryParam("googleID", qGoogleID); err != nil {
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

	if o.NameFrom != nil {

		// query param nameFrom
		var qrNameFrom string

		if o.NameFrom != nil {
			qrNameFrom = *o.NameFrom
		}
		qNameFrom := qrNameFrom
		if qNameFrom != "" {

			if err := r.SetQueryParam("nameFrom", qNameFrom); err != nil {
				return err
			}
		}
	}

	if o.NameInferredFromGithub != nil {

		// query param nameInferredFromGithub
		var qrNameInferredFromGithub bool

		if o.NameInferredFromGithub != nil {
			qrNameInferredFromGithub = *o.NameInferredFromGithub
		}
		qNameInferredFromGithub := swag.FormatBool(qrNameInferredFromGithub)
		if qNameInferredFromGithub != "" {

			if err := r.SetQueryParam("nameInferredFromGithub", qNameInferredFromGithub); err != nil {
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

	if o.SlackID != nil {

		// query param slackID
		var qrSlackID string

		if o.SlackID != nil {
			qrSlackID = *o.SlackID
		}
		qSlackID := qrSlackID
		if qSlackID != "" {

			if err := r.SetQueryParam("slackID", qSlackID); err != nil {
				return err
			}
		}
	}

	if o.SlackUsername != nil {

		// query param slackUsername
		var qrSlackUsername string

		if o.SlackUsername != nil {
			qrSlackUsername = *o.SlackUsername
		}
		qSlackUsername := qrSlackUsername
		if qSlackUsername != "" {

			if err := r.SetQueryParam("slackUsername", qSlackUsername); err != nil {
				return err
			}
		}
	}

	if o.SuitabilityDescription != nil {

		// query param suitabilityDescription
		var qrSuitabilityDescription string

		if o.SuitabilityDescription != nil {
			qrSuitabilityDescription = *o.SuitabilityDescription
		}
		qSuitabilityDescription := qrSuitabilityDescription
		if qSuitabilityDescription != "" {

			if err := r.SetQueryParam("suitabilityDescription", qSuitabilityDescription); err != nil {
				return err
			}
		}
	}

	if o.Suitable != nil {

		// query param suitable
		var qrSuitable bool

		if o.Suitable != nil {
			qrSuitable = *o.Suitable
		}
		qSuitable := swag.FormatBool(qrSuitable)
		if qSuitable != "" {

			if err := r.SetQueryParam("suitable", qSuitable); err != nil {
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
