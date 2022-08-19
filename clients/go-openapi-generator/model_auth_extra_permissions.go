/*
Sherlock

The Data Science Platform's source-of-truth service

API version: development
Contact: dsp-devops@broadinstitute.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sherlock

import (
	"encoding/json"
)

// AuthExtraPermissions struct for AuthExtraPermissions
type AuthExtraPermissions struct {
	Suitable *bool `json:"suitable,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthExtraPermissions AuthExtraPermissions

// NewAuthExtraPermissions instantiates a new AuthExtraPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthExtraPermissions() *AuthExtraPermissions {
	this := AuthExtraPermissions{}
	return &this
}

// NewAuthExtraPermissionsWithDefaults instantiates a new AuthExtraPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthExtraPermissionsWithDefaults() *AuthExtraPermissions {
	this := AuthExtraPermissions{}
	return &this
}

// GetSuitable returns the Suitable field value if set, zero value otherwise.
func (o *AuthExtraPermissions) GetSuitable() bool {
	if o == nil || o.Suitable == nil {
		var ret bool
		return ret
	}
	return *o.Suitable
}

// GetSuitableOk returns a tuple with the Suitable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthExtraPermissions) GetSuitableOk() (*bool, bool) {
	if o == nil || o.Suitable == nil {
		return nil, false
	}
	return o.Suitable, true
}

// HasSuitable returns a boolean if a field has been set.
func (o *AuthExtraPermissions) HasSuitable() bool {
	if o != nil && o.Suitable != nil {
		return true
	}

	return false
}

// SetSuitable gets a reference to the given bool and assigns it to the Suitable field.
func (o *AuthExtraPermissions) SetSuitable(v bool) {
	o.Suitable = &v
}

func (o AuthExtraPermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Suitable != nil {
		toSerialize["suitable"] = o.Suitable
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthExtraPermissions) UnmarshalJSON(bytes []byte) (err error) {
	varAuthExtraPermissions := _AuthExtraPermissions{}

	if err = json.Unmarshal(bytes, &varAuthExtraPermissions); err == nil {
		*o = AuthExtraPermissions(varAuthExtraPermissions)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "suitable")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthExtraPermissions struct {
	value *AuthExtraPermissions
	isSet bool
}

func (v NullableAuthExtraPermissions) Get() *AuthExtraPermissions {
	return v.value
}

func (v *NullableAuthExtraPermissions) Set(val *AuthExtraPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthExtraPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthExtraPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthExtraPermissions(val *AuthExtraPermissions) *NullableAuthExtraPermissions {
	return &NullableAuthExtraPermissions{value: val, isSet: true}
}

func (v NullableAuthExtraPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthExtraPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


