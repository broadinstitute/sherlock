# AuthUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticatedEmail** | Pointer to **string** |  | [optional] 
**MatchedExtraPermissions** | Pointer to [**AuthExtraPermissions**](AuthExtraPermissions.md) |  | [optional] 
**MatchedFirecloudAccount** | Pointer to [**AuthFirecloudAccount**](AuthFirecloudAccount.md) |  | [optional] 

## Methods

### NewAuthUser

`func NewAuthUser() *AuthUser`

NewAuthUser instantiates a new AuthUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthUserWithDefaults

`func NewAuthUserWithDefaults() *AuthUser`

NewAuthUserWithDefaults instantiates a new AuthUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticatedEmail

`func (o *AuthUser) GetAuthenticatedEmail() string`

GetAuthenticatedEmail returns the AuthenticatedEmail field if non-nil, zero value otherwise.

### GetAuthenticatedEmailOk

`func (o *AuthUser) GetAuthenticatedEmailOk() (*string, bool)`

GetAuthenticatedEmailOk returns a tuple with the AuthenticatedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatedEmail

`func (o *AuthUser) SetAuthenticatedEmail(v string)`

SetAuthenticatedEmail sets AuthenticatedEmail field to given value.

### HasAuthenticatedEmail

`func (o *AuthUser) HasAuthenticatedEmail() bool`

HasAuthenticatedEmail returns a boolean if a field has been set.

### GetMatchedExtraPermissions

`func (o *AuthUser) GetMatchedExtraPermissions() AuthExtraPermissions`

GetMatchedExtraPermissions returns the MatchedExtraPermissions field if non-nil, zero value otherwise.

### GetMatchedExtraPermissionsOk

`func (o *AuthUser) GetMatchedExtraPermissionsOk() (*AuthExtraPermissions, bool)`

GetMatchedExtraPermissionsOk returns a tuple with the MatchedExtraPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedExtraPermissions

`func (o *AuthUser) SetMatchedExtraPermissions(v AuthExtraPermissions)`

SetMatchedExtraPermissions sets MatchedExtraPermissions field to given value.

### HasMatchedExtraPermissions

`func (o *AuthUser) HasMatchedExtraPermissions() bool`

HasMatchedExtraPermissions returns a boolean if a field has been set.

### GetMatchedFirecloudAccount

`func (o *AuthUser) GetMatchedFirecloudAccount() AuthFirecloudAccount`

GetMatchedFirecloudAccount returns the MatchedFirecloudAccount field if non-nil, zero value otherwise.

### GetMatchedFirecloudAccountOk

`func (o *AuthUser) GetMatchedFirecloudAccountOk() (*AuthFirecloudAccount, bool)`

GetMatchedFirecloudAccountOk returns a tuple with the MatchedFirecloudAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedFirecloudAccount

`func (o *AuthUser) SetMatchedFirecloudAccount(v AuthFirecloudAccount)`

SetMatchedFirecloudAccount sets MatchedFirecloudAccount field to given value.

### HasMatchedFirecloudAccount

`func (o *AuthUser) HasMatchedFirecloudAccount() bool`

HasMatchedFirecloudAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


