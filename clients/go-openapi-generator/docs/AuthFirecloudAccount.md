# AuthFirecloudAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedGoogleTerms** | Pointer to **bool** |  | [optional] 
**Archived** | Pointer to **bool** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**EnrolledIn2Fa** | Pointer to **bool** |  | [optional] 
**Groups** | Pointer to [**AuthFirecloudGroupMembership**](AuthFirecloudGroupMembership.md) |  | [optional] 
**Suspended** | Pointer to **bool** |  | [optional] 
**SuspensionReason** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthFirecloudAccount

`func NewAuthFirecloudAccount() *AuthFirecloudAccount`

NewAuthFirecloudAccount instantiates a new AuthFirecloudAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthFirecloudAccountWithDefaults

`func NewAuthFirecloudAccountWithDefaults() *AuthFirecloudAccount`

NewAuthFirecloudAccountWithDefaults instantiates a new AuthFirecloudAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedGoogleTerms

`func (o *AuthFirecloudAccount) GetAcceptedGoogleTerms() bool`

GetAcceptedGoogleTerms returns the AcceptedGoogleTerms field if non-nil, zero value otherwise.

### GetAcceptedGoogleTermsOk

`func (o *AuthFirecloudAccount) GetAcceptedGoogleTermsOk() (*bool, bool)`

GetAcceptedGoogleTermsOk returns a tuple with the AcceptedGoogleTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedGoogleTerms

`func (o *AuthFirecloudAccount) SetAcceptedGoogleTerms(v bool)`

SetAcceptedGoogleTerms sets AcceptedGoogleTerms field to given value.

### HasAcceptedGoogleTerms

`func (o *AuthFirecloudAccount) HasAcceptedGoogleTerms() bool`

HasAcceptedGoogleTerms returns a boolean if a field has been set.

### GetArchived

`func (o *AuthFirecloudAccount) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *AuthFirecloudAccount) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *AuthFirecloudAccount) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *AuthFirecloudAccount) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetEmail

`func (o *AuthFirecloudAccount) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AuthFirecloudAccount) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AuthFirecloudAccount) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AuthFirecloudAccount) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnrolledIn2Fa

`func (o *AuthFirecloudAccount) GetEnrolledIn2Fa() bool`

GetEnrolledIn2Fa returns the EnrolledIn2Fa field if non-nil, zero value otherwise.

### GetEnrolledIn2FaOk

`func (o *AuthFirecloudAccount) GetEnrolledIn2FaOk() (*bool, bool)`

GetEnrolledIn2FaOk returns a tuple with the EnrolledIn2Fa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolledIn2Fa

`func (o *AuthFirecloudAccount) SetEnrolledIn2Fa(v bool)`

SetEnrolledIn2Fa sets EnrolledIn2Fa field to given value.

### HasEnrolledIn2Fa

`func (o *AuthFirecloudAccount) HasEnrolledIn2Fa() bool`

HasEnrolledIn2Fa returns a boolean if a field has been set.

### GetGroups

`func (o *AuthFirecloudAccount) GetGroups() AuthFirecloudGroupMembership`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *AuthFirecloudAccount) GetGroupsOk() (*AuthFirecloudGroupMembership, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *AuthFirecloudAccount) SetGroups(v AuthFirecloudGroupMembership)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *AuthFirecloudAccount) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetSuspended

`func (o *AuthFirecloudAccount) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *AuthFirecloudAccount) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *AuthFirecloudAccount) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *AuthFirecloudAccount) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetSuspensionReason

`func (o *AuthFirecloudAccount) GetSuspensionReason() string`

GetSuspensionReason returns the SuspensionReason field if non-nil, zero value otherwise.

### GetSuspensionReasonOk

`func (o *AuthFirecloudAccount) GetSuspensionReasonOk() (*string, bool)`

GetSuspensionReasonOk returns a tuple with the SuspensionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspensionReason

`func (o *AuthFirecloudAccount) SetSuspensionReason(v string)`

SetSuspensionReason sets SuspensionReason field to given value.

### HasSuspensionReason

`func (o *AuthFirecloudAccount) HasSuspensionReason() bool`

HasSuspensionReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


