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

// V2controllersCreatableEnvironment v2controllers creatable environment
//
// swagger:model v2controllers.CreatableEnvironment
type V2controllersCreatableEnvironment struct {

	// Used to schedule automatic deletion of BEEs
	AutoDelete struct {
		EnvironmentAutoDelete
	} `json:"autoDelete,omitempty"`

	// Required when creating
	Base string `json:"base,omitempty"`

	// base domain
	BaseDomain *string `json:"baseDomain,omitempty"`

	// Upon creation of a dynamic environment, if this is true the template's chart releases will be copied to the new environment
	ChartReleasesFromTemplate *bool `json:"chartReleasesFromTemplate,omitempty"`

	// default cluster
	DefaultCluster string `json:"defaultCluster,omitempty"`

	// should be the environment branch for live envs. Is usually dev for template/dynamic but not necessarily
	DefaultFirecloudDevelopRef *string `json:"defaultFirecloudDevelopRef,omitempty"`

	// When creating, will be calculated if left empty
	DefaultNamespace string `json:"defaultNamespace,omitempty"`

	// helmfile ref
	HelmfileRef *string `json:"helmfileRef,omitempty"`

	// lifecycle
	Lifecycle *string `json:"lifecycle,omitempty"`

	// When creating, will be calculated if dynamic, required otherwise
	Name string `json:"name,omitempty"`

	// Used for dynamic environment name generation only, to override using the owner email handle and template name
	NamePrefix string `json:"namePrefix,omitempty"`

	// name prefixes domain
	NamePrefixesDomain *bool `json:"namePrefixesDomain,omitempty"`

	// When creating, will be set to your email
	Owner string `json:"owner,omitempty"`

	// Used to protect specific BEEs from deletion (thelma checks this field)
	PreventDeletion *bool `json:"preventDeletion,omitempty"`

	// requires suitability
	RequiresSuitability *bool `json:"requiresSuitability,omitempty"`

	// Required for dynamic environments
	TemplateEnvironment string `json:"templateEnvironment,omitempty"`

	// When creating, will be calculated if left empty
	UniqueResourcePrefix string `json:"uniqueResourcePrefix,omitempty"`
}

// Validate validates this v2controllers creatable environment
func (m *V2controllersCreatableEnvironment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoDelete(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2controllersCreatableEnvironment) validateAutoDelete(formats strfmt.Registry) error {
	if swag.IsZero(m.AutoDelete) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this v2controllers creatable environment based on the context it is used
func (m *V2controllersCreatableEnvironment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAutoDelete(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2controllersCreatableEnvironment) contextValidateAutoDelete(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *V2controllersCreatableEnvironment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2controllersCreatableEnvironment) UnmarshalBinary(b []byte) error {
	var res V2controllersCreatableEnvironment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
