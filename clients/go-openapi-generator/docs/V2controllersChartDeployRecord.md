# V2controllersChartDeployRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartRelease** | Pointer to **string** | Required when creating | [optional] 
**ChartReleaseInfo** | Pointer to [**V2controllersChartRelease**](V2controllersChartRelease.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**ExactAppVersion** | Pointer to **string** | When creating, will default to the value currently held by the chart release | [optional] 
**ExactChartVersion** | Pointer to **string** | When creating, will default to the value currently held by the chart release | [optional] 
**HelmfileRef** | Pointer to **string** | When creating, will default to the value currently held by the chart release | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewV2controllersChartDeployRecord

`func NewV2controllersChartDeployRecord() *V2controllersChartDeployRecord`

NewV2controllersChartDeployRecord instantiates a new V2controllersChartDeployRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2controllersChartDeployRecordWithDefaults

`func NewV2controllersChartDeployRecordWithDefaults() *V2controllersChartDeployRecord`

NewV2controllersChartDeployRecordWithDefaults instantiates a new V2controllersChartDeployRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChartRelease

`func (o *V2controllersChartDeployRecord) GetChartRelease() string`

GetChartRelease returns the ChartRelease field if non-nil, zero value otherwise.

### GetChartReleaseOk

`func (o *V2controllersChartDeployRecord) GetChartReleaseOk() (*string, bool)`

GetChartReleaseOk returns a tuple with the ChartRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartRelease

`func (o *V2controllersChartDeployRecord) SetChartRelease(v string)`

SetChartRelease sets ChartRelease field to given value.

### HasChartRelease

`func (o *V2controllersChartDeployRecord) HasChartRelease() bool`

HasChartRelease returns a boolean if a field has been set.

### GetChartReleaseInfo

`func (o *V2controllersChartDeployRecord) GetChartReleaseInfo() V2controllersChartRelease`

GetChartReleaseInfo returns the ChartReleaseInfo field if non-nil, zero value otherwise.

### GetChartReleaseInfoOk

`func (o *V2controllersChartDeployRecord) GetChartReleaseInfoOk() (*V2controllersChartRelease, bool)`

GetChartReleaseInfoOk returns a tuple with the ChartReleaseInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartReleaseInfo

`func (o *V2controllersChartDeployRecord) SetChartReleaseInfo(v V2controllersChartRelease)`

SetChartReleaseInfo sets ChartReleaseInfo field to given value.

### HasChartReleaseInfo

`func (o *V2controllersChartDeployRecord) HasChartReleaseInfo() bool`

HasChartReleaseInfo returns a boolean if a field has been set.

### GetCreatedAt

`func (o *V2controllersChartDeployRecord) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V2controllersChartDeployRecord) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V2controllersChartDeployRecord) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *V2controllersChartDeployRecord) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExactAppVersion

`func (o *V2controllersChartDeployRecord) GetExactAppVersion() string`

GetExactAppVersion returns the ExactAppVersion field if non-nil, zero value otherwise.

### GetExactAppVersionOk

`func (o *V2controllersChartDeployRecord) GetExactAppVersionOk() (*string, bool)`

GetExactAppVersionOk returns a tuple with the ExactAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactAppVersion

`func (o *V2controllersChartDeployRecord) SetExactAppVersion(v string)`

SetExactAppVersion sets ExactAppVersion field to given value.

### HasExactAppVersion

`func (o *V2controllersChartDeployRecord) HasExactAppVersion() bool`

HasExactAppVersion returns a boolean if a field has been set.

### GetExactChartVersion

`func (o *V2controllersChartDeployRecord) GetExactChartVersion() string`

GetExactChartVersion returns the ExactChartVersion field if non-nil, zero value otherwise.

### GetExactChartVersionOk

`func (o *V2controllersChartDeployRecord) GetExactChartVersionOk() (*string, bool)`

GetExactChartVersionOk returns a tuple with the ExactChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactChartVersion

`func (o *V2controllersChartDeployRecord) SetExactChartVersion(v string)`

SetExactChartVersion sets ExactChartVersion field to given value.

### HasExactChartVersion

`func (o *V2controllersChartDeployRecord) HasExactChartVersion() bool`

HasExactChartVersion returns a boolean if a field has been set.

### GetHelmfileRef

`func (o *V2controllersChartDeployRecord) GetHelmfileRef() string`

GetHelmfileRef returns the HelmfileRef field if non-nil, zero value otherwise.

### GetHelmfileRefOk

`func (o *V2controllersChartDeployRecord) GetHelmfileRefOk() (*string, bool)`

GetHelmfileRefOk returns a tuple with the HelmfileRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmfileRef

`func (o *V2controllersChartDeployRecord) SetHelmfileRef(v string)`

SetHelmfileRef sets HelmfileRef field to given value.

### HasHelmfileRef

`func (o *V2controllersChartDeployRecord) HasHelmfileRef() bool`

HasHelmfileRef returns a boolean if a field has been set.

### GetId

`func (o *V2controllersChartDeployRecord) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2controllersChartDeployRecord) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2controllersChartDeployRecord) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V2controllersChartDeployRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *V2controllersChartDeployRecord) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *V2controllersChartDeployRecord) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *V2controllersChartDeployRecord) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *V2controllersChartDeployRecord) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


