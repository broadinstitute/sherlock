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

// NewGetAPIV2CiRunsParams creates a new GetAPIV2CiRunsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIV2CiRunsParams() *GetAPIV2CiRunsParams {
	return &GetAPIV2CiRunsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV2CiRunsParamsWithTimeout creates a new GetAPIV2CiRunsParams object
// with the ability to set a timeout on a request.
func NewGetAPIV2CiRunsParamsWithTimeout(timeout time.Duration) *GetAPIV2CiRunsParams {
	return &GetAPIV2CiRunsParams{
		timeout: timeout,
	}
}

// NewGetAPIV2CiRunsParamsWithContext creates a new GetAPIV2CiRunsParams object
// with the ability to set a context for a request.
func NewGetAPIV2CiRunsParamsWithContext(ctx context.Context) *GetAPIV2CiRunsParams {
	return &GetAPIV2CiRunsParams{
		Context: ctx,
	}
}

// NewGetAPIV2CiRunsParamsWithHTTPClient creates a new GetAPIV2CiRunsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIV2CiRunsParamsWithHTTPClient(client *http.Client) *GetAPIV2CiRunsParams {
	return &GetAPIV2CiRunsParams{
		HTTPClient: client,
	}
}

/* GetAPIV2CiRunsParams contains all the parameters to send to the API endpoint
   for the get API v2 ci runs operation.

   Typically these are written to a http.Request.
*/
type GetAPIV2CiRunsParams struct {

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

	   An optional limit to the number of entries returned
	*/
	Limit *int64

	// Platform.
	Platform *string

	// Status.
	Status *string

	// TerminalAt.
	TerminalAt *string

	// UpdatedAt.
	//
	// Format: date-time
	UpdatedAt *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API v2 ci runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2CiRunsParams) WithDefaults() *GetAPIV2CiRunsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API v2 ci runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2CiRunsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) WithTimeout(timeout time.Duration) *GetAPIV2CiRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) WithContext(ctx context.Context) *GetAPIV2CiRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) WithHTTPClient(client *http.Client) *GetAPIV2CiRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArgoWorkflowsName adds the argoWorkflowsName to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) WithArgoWorkflowsName(argoWorkflowsName *string) *GetAPIV2CiRunsParams {
	o.SetArgoWorkflowsName(argoWorkflowsName)
	return o
}

// SetArgoWorkflowsName adds the argoWorkflowsName to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) SetArgoWorkflowsName(argoWorkflowsName *string) {
	o.ArgoWorkflowsName = argoWorkflowsName
}

// WithArgoWorkflowsNamespace adds the argoWorkflowsNamespace to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) WithArgoWorkflowsNamespace(argoWorkflowsNamespace *string) *GetAPIV2CiRunsParams {
	o.SetArgoWorkflowsNamespace(argoWorkflowsNamespace)
	return o
}

// SetArgoWorkflowsNamespace adds the argoWorkflowsNamespace to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) SetArgoWorkflowsNamespace(argoWorkflowsNamespace *string) {
	o.ArgoWorkflowsNamespace = argoWorkflowsNamespace
}

// WithArgoWorkflowsTemplate adds the argoWorkflowsTemplate to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) WithArgoWorkflowsTemplate(argoWorkflowsTemplate *string) *GetAPIV2CiRunsParams {
	o.SetArgoWorkflowsTemplate(argoWorkflowsTemplate)
	return o
}

// SetArgoWorkflowsTemplate adds the argoWorkflowsTemplate to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) SetArgoWorkflowsTemplate(argoWorkflowsTemplate *string) {
	o.ArgoWorkflowsTemplate = argoWorkflowsTemplate
}

// WithCreatedAt adds the createdAt to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) WithCreatedAt(createdAt *strfmt.DateTime) *GetAPIV2CiRunsParams {
	o.SetCreatedAt(createdAt)
	return o
}

// SetCreatedAt adds the createdAt to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) SetCreatedAt(createdAt *strfmt.DateTime) {
	o.CreatedAt = createdAt
}

// WithGithubActionsAttemptNumber adds the githubActionsAttemptNumber to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) WithGithubActionsAttemptNumber(githubActionsAttemptNumber *int64) *GetAPIV2CiRunsParams {
	o.SetGithubActionsAttemptNumber(githubActionsAttemptNumber)
	return o
}

// SetGithubActionsAttemptNumber adds the githubActionsAttemptNumber to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) SetGithubActionsAttemptNumber(githubActionsAttemptNumber *int64) {
	o.GithubActionsAttemptNumber = githubActionsAttemptNumber
}

// WithGithubActionsOwner adds the githubActionsOwner to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) WithGithubActionsOwner(githubActionsOwner *string) *GetAPIV2CiRunsParams {
	o.SetGithubActionsOwner(githubActionsOwner)
	return o
}

// SetGithubActionsOwner adds the githubActionsOwner to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) SetGithubActionsOwner(githubActionsOwner *string) {
	o.GithubActionsOwner = githubActionsOwner
}

// WithGithubActionsRepo adds the githubActionsRepo to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) WithGithubActionsRepo(githubActionsRepo *string) *GetAPIV2CiRunsParams {
	o.SetGithubActionsRepo(githubActionsRepo)
	return o
}

// SetGithubActionsRepo adds the githubActionsRepo to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) SetGithubActionsRepo(githubActionsRepo *string) {
	o.GithubActionsRepo = githubActionsRepo
}

// WithGithubActionsRunID adds the githubActionsRunID to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) WithGithubActionsRunID(githubActionsRunID *int64) *GetAPIV2CiRunsParams {
	o.SetGithubActionsRunID(githubActionsRunID)
	return o
}

// SetGithubActionsRunID adds the githubActionsRunId to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) SetGithubActionsRunID(githubActionsRunID *int64) {
	o.GithubActionsRunID = githubActionsRunID
}

// WithGithubActionsWorkflowPath adds the githubActionsWorkflowPath to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) WithGithubActionsWorkflowPath(githubActionsWorkflowPath *string) *GetAPIV2CiRunsParams {
	o.SetGithubActionsWorkflowPath(githubActionsWorkflowPath)
	return o
}

// SetGithubActionsWorkflowPath adds the githubActionsWorkflowPath to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) SetGithubActionsWorkflowPath(githubActionsWorkflowPath *string) {
	o.GithubActionsWorkflowPath = githubActionsWorkflowPath
}

// WithID adds the id to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) WithID(id *int64) *GetAPIV2CiRunsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) SetID(id *int64) {
	o.ID = id
}

// WithLimit adds the limit to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) WithLimit(limit *int64) *GetAPIV2CiRunsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPlatform adds the platform to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) WithPlatform(platform *string) *GetAPIV2CiRunsParams {
	o.SetPlatform(platform)
	return o
}

// SetPlatform adds the platform to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) SetPlatform(platform *string) {
	o.Platform = platform
}

// WithStatus adds the status to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) WithStatus(status *string) *GetAPIV2CiRunsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) SetStatus(status *string) {
	o.Status = status
}

// WithTerminalAt adds the terminalAt to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) WithTerminalAt(terminalAt *string) *GetAPIV2CiRunsParams {
	o.SetTerminalAt(terminalAt)
	return o
}

// SetTerminalAt adds the terminalAt to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) SetTerminalAt(terminalAt *string) {
	o.TerminalAt = terminalAt
}

// WithUpdatedAt adds the updatedAt to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) WithUpdatedAt(updatedAt *strfmt.DateTime) *GetAPIV2CiRunsParams {
	o.SetUpdatedAt(updatedAt)
	return o
}

// SetUpdatedAt adds the updatedAt to the get API v2 ci runs params
func (o *GetAPIV2CiRunsParams) SetUpdatedAt(updatedAt *strfmt.DateTime) {
	o.UpdatedAt = updatedAt
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV2CiRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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