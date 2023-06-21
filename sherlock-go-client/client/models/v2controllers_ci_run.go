// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V2controllersCiRun v2controllers ci run
//
// swagger:model v2controllers.CiRun
type V2controllersCiRun struct {

	// argo workflows name
	ArgoWorkflowsName string `json:"argoWorkflowsName,omitempty"`

	// argo workflows namespace
	ArgoWorkflowsNamespace string `json:"argoWorkflowsNamespace,omitempty"`

	// argo workflows template
	ArgoWorkflowsTemplate string `json:"argoWorkflowsTemplate,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

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

	// id
	ID int64 `json:"id,omitempty"`

	// platform
	Platform string `json:"platform,omitempty"`

	// related resources
	RelatedResources []*V2controllersCiIdentifier `json:"relatedResources"`

	// started at
	StartedAt string `json:"startedAt,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// terminal at
	TerminalAt string `json:"terminalAt,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this v2controllers ci run
func (m *V2controllersCiRun) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelatedResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2controllersCiRun) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V2controllersCiRun) validateRelatedResources(formats strfmt.Registry) error {
	if swag.IsZero(m.RelatedResources) { // not required
		return nil
	}

	for i := 0; i < len(m.RelatedResources); i++ {
		if swag.IsZero(m.RelatedResources[i]) { // not required
			continue
		}

		if m.RelatedResources[i] != nil {
			if err := m.RelatedResources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relatedResources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("relatedResources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V2controllersCiRun) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v2controllers ci run based on the context it is used
func (m *V2controllersCiRun) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRelatedResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2controllersCiRun) contextValidateRelatedResources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RelatedResources); i++ {

		if m.RelatedResources[i] != nil {
			if err := m.RelatedResources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relatedResources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("relatedResources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V2controllersCiRun) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2controllersCiRun) UnmarshalBinary(b []byte) error {
	var res V2controllersCiRun
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}