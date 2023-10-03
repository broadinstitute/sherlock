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

// SherlockClusterV3Create sherlock cluster v3 create
//
// swagger:model sherlock.ClusterV3Create
type SherlockClusterV3Create struct {

	// Required when creating
	Address string `json:"address,omitempty"`

	// Required when creating if provider is 'azure'
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

// Validate validates this sherlock cluster v3 create
func (m *SherlockClusterV3Create) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var sherlockClusterV3CreateTypeProviderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["google","azure"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sherlockClusterV3CreateTypeProviderPropEnum = append(sherlockClusterV3CreateTypeProviderPropEnum, v)
	}
}

const (

	// SherlockClusterV3CreateProviderGoogle captures enum value "google"
	SherlockClusterV3CreateProviderGoogle string = "google"

	// SherlockClusterV3CreateProviderAzure captures enum value "azure"
	SherlockClusterV3CreateProviderAzure string = "azure"
)

// prop value enum
func (m *SherlockClusterV3Create) validateProviderEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sherlockClusterV3CreateTypeProviderPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SherlockClusterV3Create) validateProvider(formats strfmt.Registry) error {
	if swag.IsZero(m.Provider) { // not required
		return nil
	}

	// value enum
	if err := m.validateProviderEnum("provider", "body", *m.Provider); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sherlock cluster v3 create based on context it is used
func (m *SherlockClusterV3Create) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SherlockClusterV3Create) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SherlockClusterV3Create) UnmarshalBinary(b []byte) error {
	var res SherlockClusterV3Create
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
