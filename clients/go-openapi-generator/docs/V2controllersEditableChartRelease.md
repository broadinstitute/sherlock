# V2controllersEditableChartRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentAppVersionExact** | Pointer to **string** |  | [optional] 
**CurrentChartVersionExact** | Pointer to **string** |  | [optional] 
**HelmfileRef** | Pointer to **string** |  | [optional] [default to "HEAD"]
**TargetAppVersionBranch** | Pointer to **string** | When creating, will default to the app&#39;s main branch if it has one recorded | [optional] 
**TargetAppVersionCommit** | Pointer to **string** |  | [optional] 
**TargetAppVersionExact** | Pointer to **string** |  | [optional] 
**TargetAppVersionUse** | Pointer to **string** | When creating, will default to referencing any provided target app version field (exact, then commit, then branch) | [optional] 
**TargetChartVersionExact** | Pointer to **string** |  | [optional] 
**TargetChartVersionUse** | Pointer to **string** | When creating, will default to latest unless an exact target chart version is provided | [optional] 
**ThelmaMode** | Pointer to **string** |  | [optional] 

## Methods

### NewV2controllersEditableChartRelease

`func NewV2controllersEditableChartRelease() *V2controllersEditableChartRelease`

NewV2controllersEditableChartRelease instantiates a new V2controllersEditableChartRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2controllersEditableChartReleaseWithDefaults

`func NewV2controllersEditableChartReleaseWithDefaults() *V2controllersEditableChartRelease`

NewV2controllersEditableChartReleaseWithDefaults instantiates a new V2controllersEditableChartRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentAppVersionExact

`func (o *V2controllersEditableChartRelease) GetCurrentAppVersionExact() string`

GetCurrentAppVersionExact returns the CurrentAppVersionExact field if non-nil, zero value otherwise.

### GetCurrentAppVersionExactOk

`func (o *V2controllersEditableChartRelease) GetCurrentAppVersionExactOk() (*string, bool)`

GetCurrentAppVersionExactOk returns a tuple with the CurrentAppVersionExact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAppVersionExact

`func (o *V2controllersEditableChartRelease) SetCurrentAppVersionExact(v string)`

SetCurrentAppVersionExact sets CurrentAppVersionExact field to given value.

### HasCurrentAppVersionExact

`func (o *V2controllersEditableChartRelease) HasCurrentAppVersionExact() bool`

HasCurrentAppVersionExact returns a boolean if a field has been set.

### GetCurrentChartVersionExact

`func (o *V2controllersEditableChartRelease) GetCurrentChartVersionExact() string`

GetCurrentChartVersionExact returns the CurrentChartVersionExact field if non-nil, zero value otherwise.

### GetCurrentChartVersionExactOk

`func (o *V2controllersEditableChartRelease) GetCurrentChartVersionExactOk() (*string, bool)`

GetCurrentChartVersionExactOk returns a tuple with the CurrentChartVersionExact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentChartVersionExact

`func (o *V2controllersEditableChartRelease) SetCurrentChartVersionExact(v string)`

SetCurrentChartVersionExact sets CurrentChartVersionExact field to given value.

### HasCurrentChartVersionExact

`func (o *V2controllersEditableChartRelease) HasCurrentChartVersionExact() bool`

HasCurrentChartVersionExact returns a boolean if a field has been set.

### GetHelmfileRef

`func (o *V2controllersEditableChartRelease) GetHelmfileRef() string`

GetHelmfileRef returns the HelmfileRef field if non-nil, zero value otherwise.

### GetHelmfileRefOk

`func (o *V2controllersEditableChartRelease) GetHelmfileRefOk() (*string, bool)`

GetHelmfileRefOk returns a tuple with the HelmfileRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmfileRef

`func (o *V2controllersEditableChartRelease) SetHelmfileRef(v string)`

SetHelmfileRef sets HelmfileRef field to given value.

### HasHelmfileRef

`func (o *V2controllersEditableChartRelease) HasHelmfileRef() bool`

HasHelmfileRef returns a boolean if a field has been set.

### GetTargetAppVersionBranch

`func (o *V2controllersEditableChartRelease) GetTargetAppVersionBranch() string`

GetTargetAppVersionBranch returns the TargetAppVersionBranch field if non-nil, zero value otherwise.

### GetTargetAppVersionBranchOk

`func (o *V2controllersEditableChartRelease) GetTargetAppVersionBranchOk() (*string, bool)`

GetTargetAppVersionBranchOk returns a tuple with the TargetAppVersionBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAppVersionBranch

`func (o *V2controllersEditableChartRelease) SetTargetAppVersionBranch(v string)`

SetTargetAppVersionBranch sets TargetAppVersionBranch field to given value.

### HasTargetAppVersionBranch

`func (o *V2controllersEditableChartRelease) HasTargetAppVersionBranch() bool`

HasTargetAppVersionBranch returns a boolean if a field has been set.

### GetTargetAppVersionCommit

`func (o *V2controllersEditableChartRelease) GetTargetAppVersionCommit() string`

GetTargetAppVersionCommit returns the TargetAppVersionCommit field if non-nil, zero value otherwise.

### GetTargetAppVersionCommitOk

`func (o *V2controllersEditableChartRelease) GetTargetAppVersionCommitOk() (*string, bool)`

GetTargetAppVersionCommitOk returns a tuple with the TargetAppVersionCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAppVersionCommit

`func (o *V2controllersEditableChartRelease) SetTargetAppVersionCommit(v string)`

SetTargetAppVersionCommit sets TargetAppVersionCommit field to given value.

### HasTargetAppVersionCommit

`func (o *V2controllersEditableChartRelease) HasTargetAppVersionCommit() bool`

HasTargetAppVersionCommit returns a boolean if a field has been set.

### GetTargetAppVersionExact

`func (o *V2controllersEditableChartRelease) GetTargetAppVersionExact() string`

GetTargetAppVersionExact returns the TargetAppVersionExact field if non-nil, zero value otherwise.

### GetTargetAppVersionExactOk

`func (o *V2controllersEditableChartRelease) GetTargetAppVersionExactOk() (*string, bool)`

GetTargetAppVersionExactOk returns a tuple with the TargetAppVersionExact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAppVersionExact

`func (o *V2controllersEditableChartRelease) SetTargetAppVersionExact(v string)`

SetTargetAppVersionExact sets TargetAppVersionExact field to given value.

### HasTargetAppVersionExact

`func (o *V2controllersEditableChartRelease) HasTargetAppVersionExact() bool`

HasTargetAppVersionExact returns a boolean if a field has been set.

### GetTargetAppVersionUse

`func (o *V2controllersEditableChartRelease) GetTargetAppVersionUse() string`

GetTargetAppVersionUse returns the TargetAppVersionUse field if non-nil, zero value otherwise.

### GetTargetAppVersionUseOk

`func (o *V2controllersEditableChartRelease) GetTargetAppVersionUseOk() (*string, bool)`

GetTargetAppVersionUseOk returns a tuple with the TargetAppVersionUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAppVersionUse

`func (o *V2controllersEditableChartRelease) SetTargetAppVersionUse(v string)`

SetTargetAppVersionUse sets TargetAppVersionUse field to given value.

### HasTargetAppVersionUse

`func (o *V2controllersEditableChartRelease) HasTargetAppVersionUse() bool`

HasTargetAppVersionUse returns a boolean if a field has been set.

### GetTargetChartVersionExact

`func (o *V2controllersEditableChartRelease) GetTargetChartVersionExact() string`

GetTargetChartVersionExact returns the TargetChartVersionExact field if non-nil, zero value otherwise.

### GetTargetChartVersionExactOk

`func (o *V2controllersEditableChartRelease) GetTargetChartVersionExactOk() (*string, bool)`

GetTargetChartVersionExactOk returns a tuple with the TargetChartVersionExact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetChartVersionExact

`func (o *V2controllersEditableChartRelease) SetTargetChartVersionExact(v string)`

SetTargetChartVersionExact sets TargetChartVersionExact field to given value.

### HasTargetChartVersionExact

`func (o *V2controllersEditableChartRelease) HasTargetChartVersionExact() bool`

HasTargetChartVersionExact returns a boolean if a field has been set.

### GetTargetChartVersionUse

`func (o *V2controllersEditableChartRelease) GetTargetChartVersionUse() string`

GetTargetChartVersionUse returns the TargetChartVersionUse field if non-nil, zero value otherwise.

### GetTargetChartVersionUseOk

`func (o *V2controllersEditableChartRelease) GetTargetChartVersionUseOk() (*string, bool)`

GetTargetChartVersionUseOk returns a tuple with the TargetChartVersionUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetChartVersionUse

`func (o *V2controllersEditableChartRelease) SetTargetChartVersionUse(v string)`

SetTargetChartVersionUse sets TargetChartVersionUse field to given value.

### HasTargetChartVersionUse

`func (o *V2controllersEditableChartRelease) HasTargetChartVersionUse() bool`

HasTargetChartVersionUse returns a boolean if a field has been set.

### GetThelmaMode

`func (o *V2controllersEditableChartRelease) GetThelmaMode() string`

GetThelmaMode returns the ThelmaMode field if non-nil, zero value otherwise.

### GetThelmaModeOk

`func (o *V2controllersEditableChartRelease) GetThelmaModeOk() (*string, bool)`

GetThelmaModeOk returns a tuple with the ThelmaMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThelmaMode

`func (o *V2controllersEditableChartRelease) SetThelmaMode(v string)`

SetThelmaMode sets ThelmaMode field to given value.

### HasThelmaMode

`func (o *V2controllersEditableChartRelease) HasThelmaMode() bool`

HasThelmaMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


