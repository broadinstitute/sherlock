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

// V2controllersCluster The full set of Cluster fields that can be read or used for filtering queries
//
// swagger:model v2controllers.Cluster
type V2controllersCluster struct {

	// Required when creating
	Address string `json:"address,omitempty"`

	// Required when creating if providers is 'azure'
	AzureSubscription string `json:"azureSubscription,omitempty"`

	// Required when creating
	Base string `json:"base,omitempty"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// Required when creating if provider is 'google'
	GoogleProject string `json:"googleProject,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// Required when creating
	// Required: true
	Name *string `json:"name"`

	// provider
	// Enum: [google azure]
	Provider *string `json:"provider,omitempty"`

	// requires suitability
	RequiresSuitability *bool `json:"requiresSuitability,omitempty"`

	// updated at
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this v2controllers cluster
func (m *V2controllersCluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2controllersCluster) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var v2controllersClusterTypeProviderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["google","azure"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v2controllersClusterTypeProviderPropEnum = append(v2controllersClusterTypeProviderPropEnum, v)
	}
}

const (

	// V2controllersClusterProviderGoogle captures enum value "google"
	V2controllersClusterProviderGoogle string = "google"

	// V2controllersClusterProviderAzure captures enum value "azure"
	V2controllersClusterProviderAzure string = "azure"
)

// prop value enum
func (m *V2controllersCluster) validateProviderEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v2controllersClusterTypeProviderPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V2controllersCluster) validateProvider(formats strfmt.Registry) error {
	if swag.IsZero(m.Provider) { // not required
		return nil
	}

	// value enum
	if err := m.validateProviderEnum("provider", "body", *m.Provider); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v2controllers cluster based on context it is used
func (m *V2controllersCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V2controllersCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2controllersCluster) UnmarshalBinary(b []byte) error {
	var res V2controllersCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
