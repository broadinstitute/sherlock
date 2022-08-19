# MiscVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildInfo** | Pointer to **map[string]string** |  | [optional] 
**GoVersion** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewMiscVersionResponse

`func NewMiscVersionResponse() *MiscVersionResponse`

NewMiscVersionResponse instantiates a new MiscVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMiscVersionResponseWithDefaults

`func NewMiscVersionResponseWithDefaults() *MiscVersionResponse`

NewMiscVersionResponseWithDefaults instantiates a new MiscVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildInfo

`func (o *MiscVersionResponse) GetBuildInfo() map[string]string`

GetBuildInfo returns the BuildInfo field if non-nil, zero value otherwise.

### GetBuildInfoOk

`func (o *MiscVersionResponse) GetBuildInfoOk() (*map[string]string, bool)`

GetBuildInfoOk returns a tuple with the BuildInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildInfo

`func (o *MiscVersionResponse) SetBuildInfo(v map[string]string)`

SetBuildInfo sets BuildInfo field to given value.

### HasBuildInfo

`func (o *MiscVersionResponse) HasBuildInfo() bool`

HasBuildInfo returns a boolean if a field has been set.

### GetGoVersion

`func (o *MiscVersionResponse) GetGoVersion() string`

GetGoVersion returns the GoVersion field if non-nil, zero value otherwise.

### GetGoVersionOk

`func (o *MiscVersionResponse) GetGoVersionOk() (*string, bool)`

GetGoVersionOk returns a tuple with the GoVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoVersion

`func (o *MiscVersionResponse) SetGoVersion(v string)`

SetGoVersion sets GoVersion field to given value.

### HasGoVersion

`func (o *MiscVersionResponse) HasGoVersion() bool`

HasGoVersion returns a boolean if a field has been set.

### GetVersion

`func (o *MiscVersionResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MiscVersionResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MiscVersionResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MiscVersionResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


