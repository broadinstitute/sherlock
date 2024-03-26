// Code generated by go-swagger; DO NOT EDIT.

package github_actions_jobs

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

// NewGetAPIGithubActionsJobsV3Params creates a new GetAPIGithubActionsJobsV3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIGithubActionsJobsV3Params() *GetAPIGithubActionsJobsV3Params {
	return &GetAPIGithubActionsJobsV3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIGithubActionsJobsV3ParamsWithTimeout creates a new GetAPIGithubActionsJobsV3Params object
// with the ability to set a timeout on a request.
func NewGetAPIGithubActionsJobsV3ParamsWithTimeout(timeout time.Duration) *GetAPIGithubActionsJobsV3Params {
	return &GetAPIGithubActionsJobsV3Params{
		timeout: timeout,
	}
}

// NewGetAPIGithubActionsJobsV3ParamsWithContext creates a new GetAPIGithubActionsJobsV3Params object
// with the ability to set a context for a request.
func NewGetAPIGithubActionsJobsV3ParamsWithContext(ctx context.Context) *GetAPIGithubActionsJobsV3Params {
	return &GetAPIGithubActionsJobsV3Params{
		Context: ctx,
	}
}

// NewGetAPIGithubActionsJobsV3ParamsWithHTTPClient creates a new GetAPIGithubActionsJobsV3Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIGithubActionsJobsV3ParamsWithHTTPClient(client *http.Client) *GetAPIGithubActionsJobsV3Params {
	return &GetAPIGithubActionsJobsV3Params{
		HTTPClient: client,
	}
}

/* GetAPIGithubActionsJobsV3Params contains all the parameters to send to the API endpoint
   for the get API github actions jobs v3 operation.

   Typically these are written to a http.Request.
*/
type GetAPIGithubActionsJobsV3Params struct {

	// CreatedAt.
	//
	// Format: date-time
	CreatedAt *strfmt.DateTime

	// GithubActionsAttemptNumber.
	GithubActionsAttemptNumber *int64

	// GithubActionsJobID.
	GithubActionsJobID *int64

	// GithubActionsOwner.
	GithubActionsOwner *string

	// GithubActionsRepo.
	GithubActionsRepo *string

	// GithubActionsRunID.
	GithubActionsRunID *int64

	// ID.
	ID *int64

	// JobCreatedAt.
	//
	// Format: date-time
	JobCreatedAt *strfmt.DateTime

	// JobStartedAt.
	//
	// Format: date-time
	JobStartedAt *strfmt.DateTime

	// JobTerminalAt.
	//
	// Format: date-time
	JobTerminalAt *strfmt.DateTime

	/* Limit.

	   Control how many GithubActionsJobs are returned (default 100)
	*/
	Limit *int64

	/* Offset.

	   Control the offset for the returned GithubActionsJobs (default 0)
	*/
	Offset *int64

	// Status.
	Status *string

	// UpdatedAt.
	//
	// Format: date-time
	UpdatedAt *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API github actions jobs v3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIGithubActionsJobsV3Params) WithDefaults() *GetAPIGithubActionsJobsV3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API github actions jobs v3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIGithubActionsJobsV3Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) WithTimeout(timeout time.Duration) *GetAPIGithubActionsJobsV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) WithContext(ctx context.Context) *GetAPIGithubActionsJobsV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) WithHTTPClient(client *http.Client) *GetAPIGithubActionsJobsV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreatedAt adds the createdAt to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) WithCreatedAt(createdAt *strfmt.DateTime) *GetAPIGithubActionsJobsV3Params {
	o.SetCreatedAt(createdAt)
	return o
}

// SetCreatedAt adds the createdAt to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) SetCreatedAt(createdAt *strfmt.DateTime) {
	o.CreatedAt = createdAt
}

// WithGithubActionsAttemptNumber adds the githubActionsAttemptNumber to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) WithGithubActionsAttemptNumber(githubActionsAttemptNumber *int64) *GetAPIGithubActionsJobsV3Params {
	o.SetGithubActionsAttemptNumber(githubActionsAttemptNumber)
	return o
}

// SetGithubActionsAttemptNumber adds the githubActionsAttemptNumber to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) SetGithubActionsAttemptNumber(githubActionsAttemptNumber *int64) {
	o.GithubActionsAttemptNumber = githubActionsAttemptNumber
}

// WithGithubActionsJobID adds the githubActionsJobID to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) WithGithubActionsJobID(githubActionsJobID *int64) *GetAPIGithubActionsJobsV3Params {
	o.SetGithubActionsJobID(githubActionsJobID)
	return o
}

// SetGithubActionsJobID adds the githubActionsJobId to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) SetGithubActionsJobID(githubActionsJobID *int64) {
	o.GithubActionsJobID = githubActionsJobID
}

// WithGithubActionsOwner adds the githubActionsOwner to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) WithGithubActionsOwner(githubActionsOwner *string) *GetAPIGithubActionsJobsV3Params {
	o.SetGithubActionsOwner(githubActionsOwner)
	return o
}

// SetGithubActionsOwner adds the githubActionsOwner to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) SetGithubActionsOwner(githubActionsOwner *string) {
	o.GithubActionsOwner = githubActionsOwner
}

// WithGithubActionsRepo adds the githubActionsRepo to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) WithGithubActionsRepo(githubActionsRepo *string) *GetAPIGithubActionsJobsV3Params {
	o.SetGithubActionsRepo(githubActionsRepo)
	return o
}

// SetGithubActionsRepo adds the githubActionsRepo to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) SetGithubActionsRepo(githubActionsRepo *string) {
	o.GithubActionsRepo = githubActionsRepo
}

// WithGithubActionsRunID adds the githubActionsRunID to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) WithGithubActionsRunID(githubActionsRunID *int64) *GetAPIGithubActionsJobsV3Params {
	o.SetGithubActionsRunID(githubActionsRunID)
	return o
}

// SetGithubActionsRunID adds the githubActionsRunId to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) SetGithubActionsRunID(githubActionsRunID *int64) {
	o.GithubActionsRunID = githubActionsRunID
}

// WithID adds the id to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) WithID(id *int64) *GetAPIGithubActionsJobsV3Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) SetID(id *int64) {
	o.ID = id
}

// WithJobCreatedAt adds the jobCreatedAt to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) WithJobCreatedAt(jobCreatedAt *strfmt.DateTime) *GetAPIGithubActionsJobsV3Params {
	o.SetJobCreatedAt(jobCreatedAt)
	return o
}

// SetJobCreatedAt adds the jobCreatedAt to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) SetJobCreatedAt(jobCreatedAt *strfmt.DateTime) {
	o.JobCreatedAt = jobCreatedAt
}

// WithJobStartedAt adds the jobStartedAt to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) WithJobStartedAt(jobStartedAt *strfmt.DateTime) *GetAPIGithubActionsJobsV3Params {
	o.SetJobStartedAt(jobStartedAt)
	return o
}

// SetJobStartedAt adds the jobStartedAt to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) SetJobStartedAt(jobStartedAt *strfmt.DateTime) {
	o.JobStartedAt = jobStartedAt
}

// WithJobTerminalAt adds the jobTerminalAt to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) WithJobTerminalAt(jobTerminalAt *strfmt.DateTime) *GetAPIGithubActionsJobsV3Params {
	o.SetJobTerminalAt(jobTerminalAt)
	return o
}

// SetJobTerminalAt adds the jobTerminalAt to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) SetJobTerminalAt(jobTerminalAt *strfmt.DateTime) {
	o.JobTerminalAt = jobTerminalAt
}

// WithLimit adds the limit to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) WithLimit(limit *int64) *GetAPIGithubActionsJobsV3Params {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) WithOffset(offset *int64) *GetAPIGithubActionsJobsV3Params {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithStatus adds the status to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) WithStatus(status *string) *GetAPIGithubActionsJobsV3Params {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) SetStatus(status *string) {
	o.Status = status
}

// WithUpdatedAt adds the updatedAt to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) WithUpdatedAt(updatedAt *strfmt.DateTime) *GetAPIGithubActionsJobsV3Params {
	o.SetUpdatedAt(updatedAt)
	return o
}

// SetUpdatedAt adds the updatedAt to the get API github actions jobs v3 params
func (o *GetAPIGithubActionsJobsV3Params) SetUpdatedAt(updatedAt *strfmt.DateTime) {
	o.UpdatedAt = updatedAt
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIGithubActionsJobsV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.GithubActionsAttemptNumber != nil {

		// query param githubActionsAttemptNumber
		var qrGithubActionsAttemptNumber int64

		if o.GithubActionsAttemptNumber != nil {
			qrGithubActionsAttemptNumber = *o.GithubActionsAttemptNumber
		}
		qGithubActionsAttemptNumber := swag.FormatInt64(qrGithubActionsAttemptNumber)
		if qGithubActionsAttemptNumber != "" {

			if err := r.SetQueryParam("githubActionsAttemptNumber", qGithubActionsAttemptNumber); err != nil {
				return err
			}
		}
	}

	if o.GithubActionsJobID != nil {

		// query param githubActionsJobID
		var qrGithubActionsJobID int64

		if o.GithubActionsJobID != nil {
			qrGithubActionsJobID = *o.GithubActionsJobID
		}
		qGithubActionsJobID := swag.FormatInt64(qrGithubActionsJobID)
		if qGithubActionsJobID != "" {

			if err := r.SetQueryParam("githubActionsJobID", qGithubActionsJobID); err != nil {
				return err
			}
		}
	}

	if o.GithubActionsOwner != nil {

		// query param githubActionsOwner
		var qrGithubActionsOwner string

		if o.GithubActionsOwner != nil {
			qrGithubActionsOwner = *o.GithubActionsOwner
		}
		qGithubActionsOwner := qrGithubActionsOwner
		if qGithubActionsOwner != "" {

			if err := r.SetQueryParam("githubActionsOwner", qGithubActionsOwner); err != nil {
				return err
			}
		}
	}

	if o.GithubActionsRepo != nil {

		// query param githubActionsRepo
		var qrGithubActionsRepo string

		if o.GithubActionsRepo != nil {
			qrGithubActionsRepo = *o.GithubActionsRepo
		}
		qGithubActionsRepo := qrGithubActionsRepo
		if qGithubActionsRepo != "" {

			if err := r.SetQueryParam("githubActionsRepo", qGithubActionsRepo); err != nil {
				return err
			}
		}
	}

	if o.GithubActionsRunID != nil {

		// query param githubActionsRunID
		var qrGithubActionsRunID int64

		if o.GithubActionsRunID != nil {
			qrGithubActionsRunID = *o.GithubActionsRunID
		}
		qGithubActionsRunID := swag.FormatInt64(qrGithubActionsRunID)
		if qGithubActionsRunID != "" {

			if err := r.SetQueryParam("githubActionsRunID", qGithubActionsRunID); err != nil {
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

	if o.JobCreatedAt != nil {

		// query param jobCreatedAt
		var qrJobCreatedAt strfmt.DateTime

		if o.JobCreatedAt != nil {
			qrJobCreatedAt = *o.JobCreatedAt
		}
		qJobCreatedAt := qrJobCreatedAt.String()
		if qJobCreatedAt != "" {

			if err := r.SetQueryParam("jobCreatedAt", qJobCreatedAt); err != nil {
				return err
			}
		}
	}

	if o.JobStartedAt != nil {

		// query param jobStartedAt
		var qrJobStartedAt strfmt.DateTime

		if o.JobStartedAt != nil {
			qrJobStartedAt = *o.JobStartedAt
		}
		qJobStartedAt := qrJobStartedAt.String()
		if qJobStartedAt != "" {

			if err := r.SetQueryParam("jobStartedAt", qJobStartedAt); err != nil {
				return err
			}
		}
	}

	if o.JobTerminalAt != nil {

		// query param jobTerminalAt
		var qrJobTerminalAt strfmt.DateTime

		if o.JobTerminalAt != nil {
			qrJobTerminalAt = *o.JobTerminalAt
		}
		qJobTerminalAt := qrJobTerminalAt.String()
		if qJobTerminalAt != "" {

			if err := r.SetQueryParam("jobTerminalAt", qJobTerminalAt); err != nil {
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

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
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
