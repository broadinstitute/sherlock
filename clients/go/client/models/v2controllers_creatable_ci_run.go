// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V2controllersCreatableCiRun v2controllers creatable ci run
//
// swagger:model v2controllers.CreatableCiRun
type V2controllersCreatableCiRun struct {

	// Always appends; will eliminate duplicates.
	AppVersions []string `json:"appVersions"`

	// argo workflows name
	ArgoWorkflowsName string `json:"argoWorkflowsName,omitempty"`

	// argo workflows namespace
	ArgoWorkflowsNamespace string `json:"argoWorkflowsNamespace,omitempty"`

	// argo workflows template
	ArgoWorkflowsTemplate string `json:"argoWorkflowsTemplate,omitempty"`

	// Always appends; will eliminate duplicates. Spreads to associated chart releases (and environments and clusters) and new app/chart versions.
	Changesets []string `json:"changesets"`

	// Always appends; will eliminate duplicates. Spreads to associated environments and clusters.
	ChartReleases []string `json:"chartReleases"`

	// Always appends; will eliminate duplicates.
	ChartVersions []string `json:"chartVersions"`

	// Always appends; will eliminate duplicates.
	Charts []string `json:"charts"`

	// Always appends; will eliminate duplicates.
	Clusters []string `json:"clusters"`

	// Always appends; will eliminate duplicates.
	Environments []string `json:"environments"`

	// github actions attempt number
	GithubActionsAttemptNumber int64 `json:"githubActionsAttemptNumber,omitempty"`

	// github actions owner
	GithubActionsOwner string `json:"githubActionsOwner,omitempty"`

	// github actions repo
	GithubActionsRepo string `json:"githubActionsRepo,omitempty"`

	// github actions run ID
	GithubActionsRunID int64 `json:"githubActionsRunID,omitempty"`

	// github actions workflow path
	GithubActionsWorkflowPath string `json:"githubActionsWorkflowPath,omitempty"`

	// platform
	Platform string `json:"platform,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// terminal at
	TerminalAt string `json:"terminalAt,omitempty"`
}

// Validate validates this v2controllers creatable ci run
func (m *V2controllersCreatableCiRun) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v2controllers creatable ci run based on context it is used
func (m *V2controllersCreatableCiRun) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V2controllersCreatableCiRun) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2controllersCreatableCiRun) UnmarshalBinary(b []byte) error {
	var res V2controllersCreatableCiRun
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
