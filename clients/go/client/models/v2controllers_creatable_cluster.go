// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V2controllersCreatableCluster The subset of Cluster fields that can be set upon creation
//
// swagger:model v2controllers.CreatableCluster
type V2controllersCreatableCluster struct {

	// Required when creating
	Address string `json:"address,omitempty"`

	// Required when creating if providers is 'azure'
	AzureSubscription string `json:"azureSubscription,omitempty"`

	// Required when creating
	Base string `json:"base,omitempty"`

	// Required when creating if provider is 'google'
	GoogleProject string `json:"googleProject,omitempty"`

	// helmfile ref
	HelmfileRef *string `json:"helmfileRef,omitempty"`

	// location
	Location *string `json:"location,omitempty"`

	// Required when creating
	Name string `json:"name,omitempty"`

	// provider
	// Enum: [google azure]
	Provider *string `json:"provider,omitempty"`

	// requires suitability
	RequiresSuitability *bool `json:"requiresSuitability,omitempty"`
}

// Validate validates this v2controllers creatable cluster
func (m *V2controllersCreatableCluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v2controllersCreatableClusterTypeProviderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["google","azure"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v2controllersCreatableClusterTypeProviderPropEnum = append(v2controllersCreatableClusterTypeProviderPropEnum, v)
	}
}

const (

	// V2controllersCreatableClusterProviderGoogle captures enum value "google"
	V2controllersCreatableClusterProviderGoogle string = "google"

	// V2controllersCreatableClusterProviderAzure captures enum value "azure"
	V2controllersCreatableClusterProviderAzure string = "azure"
)

// prop value enum
func (m *V2controllersCreatableCluster) validateProviderEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v2controllersCreatableClusterTypeProviderPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V2controllersCreatableCluster) validateProvider(formats strfmt.Registry) error {
	if swag.IsZero(m.Provider) { // not required
		return nil
	}

	// value enum
	if err := m.validateProviderEnum("provider", "body", *m.Provider); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v2controllers creatable cluster based on context it is used
func (m *V2controllersCreatableCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V2controllersCreatableCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2controllersCreatableCluster) UnmarshalBinary(b []byte) error {
	var res V2controllersCreatableCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
