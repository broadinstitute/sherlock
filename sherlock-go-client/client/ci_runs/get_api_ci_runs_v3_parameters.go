// Code generated by go-swagger; DO NOT EDIT.

package ci_runs

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

// NewGetAPICiRunsV3Params creates a new GetAPICiRunsV3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPICiRunsV3Params() *GetAPICiRunsV3Params {
	return &GetAPICiRunsV3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPICiRunsV3ParamsWithTimeout creates a new GetAPICiRunsV3Params object
// with the ability to set a timeout on a request.
func NewGetAPICiRunsV3ParamsWithTimeout(timeout time.Duration) *GetAPICiRunsV3Params {
	return &GetAPICiRunsV3Params{
		timeout: timeout,
	}
}

// NewGetAPICiRunsV3ParamsWithContext creates a new GetAPICiRunsV3Params object
// with the ability to set a context for a request.
func NewGetAPICiRunsV3ParamsWithContext(ctx context.Context) *GetAPICiRunsV3Params {
	return &GetAPICiRunsV3Params{
		Context: ctx,
	}
}

// NewGetAPICiRunsV3ParamsWithHTTPClient creates a new GetAPICiRunsV3Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPICiRunsV3ParamsWithHTTPClient(client *http.Client) *GetAPICiRunsV3Params {
	return &GetAPICiRunsV3Params{
		HTTPClient: client,
	}
}

/* GetAPICiRunsV3Params contains all the parameters to send to the API endpoint
   for the get API ci runs v3 operation.

   Typically these are written to a http.Request.
*/
type GetAPICiRunsV3Params struct {

	// ArgoWorkflowsName.
	ArgoWorkflowsName *string

	// ArgoWorkflowsNamespace.
	ArgoWorkflowsNamespace *string

	// ArgoWorkflowsTemplate.
	ArgoWorkflowsTemplate *string

	// CreatedAt.
	//
	// Format: date-time
	CreatedAt *strfmt.DateTime

	// GithubActionsAttemptNumber.
	GithubActionsAttemptNumber *int64

	// GithubActionsOwner.
	GithubActionsOwner *string

	// GithubActionsRepo.
	GithubActionsRepo *string

	// GithubActionsRunID.
	GithubActionsRunID *int64

	// GithubActionsWorkflowPath.
	GithubActionsWorkflowPath *string

	// ID.
	ID *int64

	/* Limit.

	   Control how many CiRuns are returned (default 100)
	*/
	Limit *int64

	/* Offset.

	   Control the offset for the returned CiRuns (default 0)
	*/
	Offset *int64

	// Platform.
	Platform *string

	// StartedAt.
	StartedAt *string

	// Status.
	Status *string

	// TerminalAt.
	TerminalAt *string

	// TerminationHooksDispatchedAt.
	//
	// Format: date-time
	TerminationHooksDispatchedAt *strfmt.DateTime

	// UpdatedAt.
	//
	// Format: date-time
	UpdatedAt *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API ci runs v3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPICiRunsV3Params) WithDefaults() *GetAPICiRunsV3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API ci runs v3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPICiRunsV3Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) WithTimeout(timeout time.Duration) *GetAPICiRunsV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) WithContext(ctx context.Context) *GetAPICiRunsV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) WithHTTPClient(client *http.Client) *GetAPICiRunsV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArgoWorkflowsName adds the argoWorkflowsName to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) WithArgoWorkflowsName(argoWorkflowsName *string) *GetAPICiRunsV3Params {
	o.SetArgoWorkflowsName(argoWorkflowsName)
	return o
}

// SetArgoWorkflowsName adds the argoWorkflowsName to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) SetArgoWorkflowsName(argoWorkflowsName *string) {
	o.ArgoWorkflowsName = argoWorkflowsName
}

// WithArgoWorkflowsNamespace adds the argoWorkflowsNamespace to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) WithArgoWorkflowsNamespace(argoWorkflowsNamespace *string) *GetAPICiRunsV3Params {
	o.SetArgoWorkflowsNamespace(argoWorkflowsNamespace)
	return o
}

// SetArgoWorkflowsNamespace adds the argoWorkflowsNamespace to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) SetArgoWorkflowsNamespace(argoWorkflowsNamespace *string) {
	o.ArgoWorkflowsNamespace = argoWorkflowsNamespace
}

// WithArgoWorkflowsTemplate adds the argoWorkflowsTemplate to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) WithArgoWorkflowsTemplate(argoWorkflowsTemplate *string) *GetAPICiRunsV3Params {
	o.SetArgoWorkflowsTemplate(argoWorkflowsTemplate)
	return o
}

// SetArgoWorkflowsTemplate adds the argoWorkflowsTemplate to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) SetArgoWorkflowsTemplate(argoWorkflowsTemplate *string) {
	o.ArgoWorkflowsTemplate = argoWorkflowsTemplate
}

// WithCreatedAt adds the createdAt to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) WithCreatedAt(createdAt *strfmt.DateTime) *GetAPICiRunsV3Params {
	o.SetCreatedAt(createdAt)
	return o
}

// SetCreatedAt adds the createdAt to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) SetCreatedAt(createdAt *strfmt.DateTime) {
	o.CreatedAt = createdAt
}

// WithGithubActionsAttemptNumber adds the githubActionsAttemptNumber to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) WithGithubActionsAttemptNumber(githubActionsAttemptNumber *int64) *GetAPICiRunsV3Params {
	o.SetGithubActionsAttemptNumber(githubActionsAttemptNumber)
	return o
}

// SetGithubActionsAttemptNumber adds the githubActionsAttemptNumber to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) SetGithubActionsAttemptNumber(githubActionsAttemptNumber *int64) {
	o.GithubActionsAttemptNumber = githubActionsAttemptNumber
}

// WithGithubActionsOwner adds the githubActionsOwner to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) WithGithubActionsOwner(githubActionsOwner *string) *GetAPICiRunsV3Params {
	o.SetGithubActionsOwner(githubActionsOwner)
	return o
}

// SetGithubActionsOwner adds the githubActionsOwner to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) SetGithubActionsOwner(githubActionsOwner *string) {
	o.GithubActionsOwner = githubActionsOwner
}

// WithGithubActionsRepo adds the githubActionsRepo to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) WithGithubActionsRepo(githubActionsRepo *string) *GetAPICiRunsV3Params {
	o.SetGithubActionsRepo(githubActionsRepo)
	return o
}

// SetGithubActionsRepo adds the githubActionsRepo to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) SetGithubActionsRepo(githubActionsRepo *string) {
	o.GithubActionsRepo = githubActionsRepo
}

// WithGithubActionsRunID adds the githubActionsRunID to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) WithGithubActionsRunID(githubActionsRunID *int64) *GetAPICiRunsV3Params {
	o.SetGithubActionsRunID(githubActionsRunID)
	return o
}

// SetGithubActionsRunID adds the githubActionsRunId to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) SetGithubActionsRunID(githubActionsRunID *int64) {
	o.GithubActionsRunID = githubActionsRunID
}

// WithGithubActionsWorkflowPath adds the githubActionsWorkflowPath to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) WithGithubActionsWorkflowPath(githubActionsWorkflowPath *string) *GetAPICiRunsV3Params {
	o.SetGithubActionsWorkflowPath(githubActionsWorkflowPath)
	return o
}

// SetGithubActionsWorkflowPath adds the githubActionsWorkflowPath to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) SetGithubActionsWorkflowPath(githubActionsWorkflowPath *string) {
	o.GithubActionsWorkflowPath = githubActionsWorkflowPath
}

// WithID adds the id to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) WithID(id *int64) *GetAPICiRunsV3Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) SetID(id *int64) {
	o.ID = id
}

// WithLimit adds the limit to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) WithLimit(limit *int64) *GetAPICiRunsV3Params {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) WithOffset(offset *int64) *GetAPICiRunsV3Params {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPlatform adds the platform to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) WithPlatform(platform *string) *GetAPICiRunsV3Params {
	o.SetPlatform(platform)
	return o
}

// SetPlatform adds the platform to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) SetPlatform(platform *string) {
	o.Platform = platform
}

// WithStartedAt adds the startedAt to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) WithStartedAt(startedAt *string) *GetAPICiRunsV3Params {
	o.SetStartedAt(startedAt)
	return o
}

// SetStartedAt adds the startedAt to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) SetStartedAt(startedAt *string) {
	o.StartedAt = startedAt
}

// WithStatus adds the status to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) WithStatus(status *string) *GetAPICiRunsV3Params {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) SetStatus(status *string) {
	o.Status = status
}

// WithTerminalAt adds the terminalAt to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) WithTerminalAt(terminalAt *string) *GetAPICiRunsV3Params {
	o.SetTerminalAt(terminalAt)
	return o
}

// SetTerminalAt adds the terminalAt to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) SetTerminalAt(terminalAt *string) {
	o.TerminalAt = terminalAt
}

// WithTerminationHooksDispatchedAt adds the terminationHooksDispatchedAt to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) WithTerminationHooksDispatchedAt(terminationHooksDispatchedAt *strfmt.DateTime) *GetAPICiRunsV3Params {
	o.SetTerminationHooksDispatchedAt(terminationHooksDispatchedAt)
	return o
}

// SetTerminationHooksDispatchedAt adds the terminationHooksDispatchedAt to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) SetTerminationHooksDispatchedAt(terminationHooksDispatchedAt *strfmt.DateTime) {
	o.TerminationHooksDispatchedAt = terminationHooksDispatchedAt
}

// WithUpdatedAt adds the updatedAt to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) WithUpdatedAt(updatedAt *strfmt.DateTime) *GetAPICiRunsV3Params {
	o.SetUpdatedAt(updatedAt)
	return o
}

// SetUpdatedAt adds the updatedAt to the get API ci runs v3 params
func (o *GetAPICiRunsV3Params) SetUpdatedAt(updatedAt *strfmt.DateTime) {
	o.UpdatedAt = updatedAt
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPICiRunsV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ArgoWorkflowsName != nil {

		// query param argoWorkflowsName
		var qrArgoWorkflowsName string

		if o.ArgoWorkflowsName != nil {
			qrArgoWorkflowsName = *o.ArgoWorkflowsName
		}
		qArgoWorkflowsName := qrArgoWorkflowsName
		if qArgoWorkflowsName != "" {

			if err := r.SetQueryParam("argoWorkflowsName", qArgoWorkflowsName); err != nil {
				return err
			}
		}
	}

	if o.ArgoWorkflowsNamespace != nil {

		// query param argoWorkflowsNamespace
		var qrArgoWorkflowsNamespace string

		if o.ArgoWorkflowsNamespace != nil {
			qrArgoWorkflowsNamespace = *o.ArgoWorkflowsNamespace
		}
		qArgoWorkflowsNamespace := qrArgoWorkflowsNamespace
		if qArgoWorkflowsNamespace != "" {

			if err := r.SetQueryParam("argoWorkflowsNamespace", qArgoWorkflowsNamespace); err != nil {
				return err
			}
		}
	}

	if o.ArgoWorkflowsTemplate != nil {

		// query param argoWorkflowsTemplate
		var qrArgoWorkflowsTemplate string

		if o.ArgoWorkflowsTemplate != nil {
			qrArgoWorkflowsTemplate = *o.ArgoWorkflowsTemplate
		}
		qArgoWorkflowsTemplate := qrArgoWorkflowsTemplate
		if qArgoWorkflowsTemplate != "" {

			if err := r.SetQueryParam("argoWorkflowsTemplate", qArgoWorkflowsTemplate); err != nil {
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

	if o.GithubActionsWorkflowPath != nil {

		// query param githubActionsWorkflowPath
		var qrGithubActionsWorkflowPath string

		if o.GithubActionsWorkflowPath != nil {
			qrGithubActionsWorkflowPath = *o.GithubActionsWorkflowPath
		}
		qGithubActionsWorkflowPath := qrGithubActionsWorkflowPath
		if qGithubActionsWorkflowPath != "" {

			if err := r.SetQueryParam("githubActionsWorkflowPath", qGithubActionsWorkflowPath); err != nil {
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

	if o.Platform != nil {

		// query param platform
		var qrPlatform string

		if o.Platform != nil {
			qrPlatform = *o.Platform
		}
		qPlatform := qrPlatform
		if qPlatform != "" {

			if err := r.SetQueryParam("platform", qPlatform); err != nil {
				return err
			}
		}
	}

	if o.StartedAt != nil {

		// query param startedAt
		var qrStartedAt string

		if o.StartedAt != nil {
			qrStartedAt = *o.StartedAt
		}
		qStartedAt := qrStartedAt
		if qStartedAt != "" {

			if err := r.SetQueryParam("startedAt", qStartedAt); err != nil {
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

	if o.TerminalAt != nil {

		// query param terminalAt
		var qrTerminalAt string

		if o.TerminalAt != nil {
			qrTerminalAt = *o.TerminalAt
		}
		qTerminalAt := qrTerminalAt
		if qTerminalAt != "" {

			if err := r.SetQueryParam("terminalAt", qTerminalAt); err != nil {
				return err
			}
		}
	}

	if o.TerminationHooksDispatchedAt != nil {

		// query param terminationHooksDispatchedAt
		var qrTerminationHooksDispatchedAt strfmt.DateTime

		if o.TerminationHooksDispatchedAt != nil {
			qrTerminationHooksDispatchedAt = *o.TerminationHooksDispatchedAt
		}
		qTerminationHooksDispatchedAt := qrTerminationHooksDispatchedAt.String()
		if qTerminationHooksDispatchedAt != "" {

			if err := r.SetQueryParam("terminationHooksDispatchedAt", qTerminationHooksDispatchedAt); err != nil {
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
