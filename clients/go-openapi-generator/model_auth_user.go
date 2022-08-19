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

// AuthUser struct for AuthUser
type AuthUser struct {
	AuthenticatedEmail *string `json:"authenticatedEmail,omitempty"`
	MatchedExtraPermissions *AuthExtraPermissions `json:"matchedExtraPermissions,omitempty"`
	MatchedFirecloudAccount *AuthFirecloudAccount `json:"matchedFirecloudAccount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthUser AuthUser

// NewAuthUser instantiates a new AuthUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthUser() *AuthUser {
	this := AuthUser{}
	return &this
}

// NewAuthUserWithDefaults instantiates a new AuthUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthUserWithDefaults() *AuthUser {
	this := AuthUser{}
	return &this
}

// GetAuthenticatedEmail returns the AuthenticatedEmail field value if set, zero value otherwise.
func (o *AuthUser) GetAuthenticatedEmail() string {
	if o == nil || o.AuthenticatedEmail == nil {
		var ret string
		return ret
	}
	return *o.AuthenticatedEmail
}

// GetAuthenticatedEmailOk returns a tuple with the AuthenticatedEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthUser) GetAuthenticatedEmailOk() (*string, bool) {
	if o == nil || o.AuthenticatedEmail == nil {
		return nil, false
	}
	return o.AuthenticatedEmail, true
}

// HasAuthenticatedEmail returns a boolean if a field has been set.
func (o *AuthUser) HasAuthenticatedEmail() bool {
	if o != nil && o.AuthenticatedEmail != nil {
		return true
	}

	return false
}

// SetAuthenticatedEmail gets a reference to the given string and assigns it to the AuthenticatedEmail field.
func (o *AuthUser) SetAuthenticatedEmail(v string) {
	o.AuthenticatedEmail = &v
}

// GetMatchedExtraPermissions returns the MatchedExtraPermissions field value if set, zero value otherwise.
func (o *AuthUser) GetMatchedExtraPermissions() AuthExtraPermissions {
	if o == nil || o.MatchedExtraPermissions == nil {
		var ret AuthExtraPermissions
		return ret
	}
	return *o.MatchedExtraPermissions
}

// GetMatchedExtraPermissionsOk returns a tuple with the MatchedExtraPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthUser) GetMatchedExtraPermissionsOk() (*AuthExtraPermissions, bool) {
	if o == nil || o.MatchedExtraPermissions == nil {
		return nil, false
	}
	return o.MatchedExtraPermissions, true
}

// HasMatchedExtraPermissions returns a boolean if a field has been set.
func (o *AuthUser) HasMatchedExtraPermissions() bool {
	if o != nil && o.MatchedExtraPermissions != nil {
		return true
	}

	return false
}

// SetMatchedExtraPermissions gets a reference to the given AuthExtraPermissions and assigns it to the MatchedExtraPermissions field.
func (o *AuthUser) SetMatchedExtraPermissions(v AuthExtraPermissions) {
	o.MatchedExtraPermissions = &v
}

// GetMatchedFirecloudAccount returns the MatchedFirecloudAccount field value if set, zero value otherwise.
func (o *AuthUser) GetMatchedFirecloudAccount() AuthFirecloudAccount {
	if o == nil || o.MatchedFirecloudAccount == nil {
		var ret AuthFirecloudAccount
		return ret
	}
	return *o.MatchedFirecloudAccount
}

// GetMatchedFirecloudAccountOk returns a tuple with the MatchedFirecloudAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthUser) GetMatchedFirecloudAccountOk() (*AuthFirecloudAccount, bool) {
	if o == nil || o.MatchedFirecloudAccount == nil {
		return nil, false
	}
	return o.MatchedFirecloudAccount, true
}

// HasMatchedFirecloudAccount returns a boolean if a field has been set.
func (o *AuthUser) HasMatchedFirecloudAccount() bool {
	if o != nil && o.MatchedFirecloudAccount != nil {
		return true
	}

	return false
}

// SetMatchedFirecloudAccount gets a reference to the given AuthFirecloudAccount and assigns it to the MatchedFirecloudAccount field.
func (o *AuthUser) SetMatchedFirecloudAccount(v AuthFirecloudAccount) {
	o.MatchedFirecloudAccount = &v
}

func (o AuthUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthenticatedEmail != nil {
		toSerialize["authenticatedEmail"] = o.AuthenticatedEmail
	}
	if o.MatchedExtraPermissions != nil {
		toSerialize["matchedExtraPermissions"] = o.MatchedExtraPermissions
	}
	if o.MatchedFirecloudAccount != nil {
		toSerialize["matchedFirecloudAccount"] = o.MatchedFirecloudAccount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthUser) UnmarshalJSON(bytes []byte) (err error) {
	varAuthUser := _AuthUser{}

	if err = json.Unmarshal(bytes, &varAuthUser); err == nil {
		*o = AuthUser(varAuthUser)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "authenticatedEmail")
		delete(additionalProperties, "matchedExtraPermissions")
		delete(additionalProperties, "matchedFirecloudAccount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthUser struct {
	value *AuthUser
	isSet bool
}

func (v NullableAuthUser) Get() *AuthUser {
	return v.value
}

func (v *NullableAuthUser) Set(val *AuthUser) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthUser) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthUser(val *AuthUser) *NullableAuthUser {
	return &NullableAuthUser{value: val, isSet: true}
}

func (v NullableAuthUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


