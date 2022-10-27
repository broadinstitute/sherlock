// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V2controllersEnvironment v2controllers environment
//
// swagger:model v2controllers.Environment
type V2controllersEnvironment struct {

	// Required when creating
	Base string `json:"base,omitempty"`

	// base domain
	BaseDomain *string `json:"baseDomain,omitempty"`

	// Upon creation of a dynamic environment, if this is true the template's chart releases will be copied to the new environment
	ChartReleasesFromTemplate *bool `json:"chartReleasesFromTemplate,omitempty"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// default cluster
	DefaultCluster string `json:"defaultCluster,omitempty"`

	// default cluster info
	DefaultClusterInfo *V2controllersCluster `json:"defaultClusterInfo,omitempty"`

	// should be the environment branch for live envs. Is usually dev for template/dynamic but not necessarily
	DefaultFirecloudDevelopRef *string `json:"defaultFirecloudDevelopRef,omitempty"`

	// default namespace
	DefaultNamespace string `json:"defaultNamespace,omitempty"`

	// helmfile ref
	HelmfileRef *string `json:"helmfileRef,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// lifecycle
	Lifecycle *string `json:"lifecycle,omitempty"`

	// When creating, will be calculated if dynamic, required otherwise
	Name string `json:"name,omitempty"`

	// name prefixes domain
	NamePrefixesDomain *bool `json:"namePrefixesDomain,omitempty"`

	// When creating, will be set to your email
	Owner string `json:"owner,omitempty"`

	// requires suitability
	RequiresSuitability *bool `json:"requiresSuitability,omitempty"`

	// Required for dynamic environments
	TemplateEnvironment string `json:"templateEnvironment,omitempty"`

	// Single-layer recursive; provides info of the template environment if this environment has one
	TemplateEnvironmentInfo interface{} `json:"templateEnvironmentInfo,omitempty"`

	// When creating, will be calculated if left empty
	UniqueResourcePrefix string `json:"uniqueResourcePrefix,omitempty"`

	// updated at
	UpdatedAt string `json:"updatedAt,omitempty"`

	// values name
	ValuesName string `json:"valuesName,omitempty"`
}

// Validate validates this v2controllers environment
func (m *V2controllersEnvironment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultClusterInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2controllersEnvironment) validateDefaultClusterInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultClusterInfo) { // not required
		return nil
	}

	if m.DefaultClusterInfo != nil {
		if err := m.DefaultClusterInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultClusterInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("defaultClusterInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v2controllers environment based on the context it is used
func (m *V2controllersEnvironment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDefaultClusterInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2controllersEnvironment) contextValidateDefaultClusterInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultClusterInfo != nil {
		if err := m.DefaultClusterInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultClusterInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("defaultClusterInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V2controllersEnvironment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2controllersEnvironment) UnmarshalBinary(b []byte) error {
	var res V2controllersEnvironment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
