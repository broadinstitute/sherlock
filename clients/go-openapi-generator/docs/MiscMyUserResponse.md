# MiscMyUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**RawInfo** | Pointer to [**AuthUser**](AuthUser.md) |  | [optional] 
**Suitability** | Pointer to **string** |  | [optional] 

## Methods

### NewMiscMyUserResponse

`func NewMiscMyUserResponse() *MiscMyUserResponse`

NewMiscMyUserResponse instantiates a new MiscMyUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMiscMyUserResponseWithDefaults

`func NewMiscMyUserResponseWithDefaults() *MiscMyUserResponse`

NewMiscMyUserResponseWithDefaults instantiates a new MiscMyUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *MiscMyUserResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MiscMyUserResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MiscMyUserResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MiscMyUserResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetRawInfo

`func (o *MiscMyUserResponse) GetRawInfo() AuthUser`

GetRawInfo returns the RawInfo field if non-nil, zero value otherwise.

### GetRawInfoOk

`func (o *MiscMyUserResponse) GetRawInfoOk() (*AuthUser, bool)`

GetRawInfoOk returns a tuple with the RawInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawInfo

`func (o *MiscMyUserResponse) SetRawInfo(v AuthUser)`

SetRawInfo sets RawInfo field to given value.

### HasRawInfo

`func (o *MiscMyUserResponse) HasRawInfo() bool`

HasRawInfo returns a boolean if a field has been set.

### GetSuitability

`func (o *MiscMyUserResponse) GetSuitability() string`

GetSuitability returns the Suitability field if non-nil, zero value otherwise.

### GetSuitabilityOk

`func (o *MiscMyUserResponse) GetSuitabilityOk() (*string, bool)`

GetSuitabilityOk returns a tuple with the Suitability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuitability

`func (o *MiscMyUserResponse) SetSuitability(v string)`

SetSuitability sets Suitability field to given value.

### HasSuitability

`func (o *MiscMyUserResponse) HasSuitability() bool`

HasSuitability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


