# V2controllersCreatableChartRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chart** | Pointer to **string** | Required when creating | [optional] 
**Cluster** | Pointer to **string** | When creating, will default the environment&#39;s default cluster, if provided. Either this or environment must be provided. | [optional] 
**CurrentAppVersionExact** | Pointer to **string** |  | [optional] 
**CurrentChartVersionExact** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to **string** | Either this or cluster must be provided. | [optional] 
**HelmfileRef** | Pointer to **string** |  | [optional] [default to "HEAD"]
**Name** | Pointer to **string** | When creating, will be calculated if left empty | [optional] 
**Namespace** | Pointer to **string** | When creating, will default to the environment&#39;s default namespace, if provided | [optional] 
**TargetAppVersionBranch** | Pointer to **string** | When creating, will default to the app&#39;s main branch if it has one recorded | [optional] 
**TargetAppVersionCommit** | Pointer to **string** |  | [optional] 
**TargetAppVersionExact** | Pointer to **string** |  | [optional] 
**TargetAppVersionUse** | Pointer to **string** | When creating, will default to referencing any provided target app version field (exact, then commit, then branch) | [optional] 
**TargetChartVersionExact** | Pointer to **string** |  | [optional] 
**TargetChartVersionUse** | Pointer to **string** | When creating, will default to latest unless an exact target chart version is provided | [optional] 
**ThelmaMode** | Pointer to **string** |  | [optional] 

## Methods

### NewV2controllersCreatableChartRelease

`func NewV2controllersCreatableChartRelease() *V2controllersCreatableChartRelease`

NewV2controllersCreatableChartRelease instantiates a new V2controllersCreatableChartRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2controllersCreatableChartReleaseWithDefaults

`func NewV2controllersCreatableChartReleaseWithDefaults() *V2controllersCreatableChartRelease`

NewV2controllersCreatableChartReleaseWithDefaults instantiates a new V2controllersCreatableChartRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChart

`func (o *V2controllersCreatableChartRelease) GetChart() string`

GetChart returns the Chart field if non-nil, zero value otherwise.

### GetChartOk

`func (o *V2controllersCreatableChartRelease) GetChartOk() (*string, bool)`

GetChartOk returns a tuple with the Chart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChart

`func (o *V2controllersCreatableChartRelease) SetChart(v string)`

SetChart sets Chart field to given value.

### HasChart

`func (o *V2controllersCreatableChartRelease) HasChart() bool`

HasChart returns a boolean if a field has been set.

### GetCluster

`func (o *V2controllersCreatableChartRelease) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V2controllersCreatableChartRelease) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V2controllersCreatableChartRelease) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *V2controllersCreatableChartRelease) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCurrentAppVersionExact

`func (o *V2controllersCreatableChartRelease) GetCurrentAppVersionExact() string`

GetCurrentAppVersionExact returns the CurrentAppVersionExact field if non-nil, zero value otherwise.

### GetCurrentAppVersionExactOk

`func (o *V2controllersCreatableChartRelease) GetCurrentAppVersionExactOk() (*string, bool)`

GetCurrentAppVersionExactOk returns a tuple with the CurrentAppVersionExact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAppVersionExact

`func (o *V2controllersCreatableChartRelease) SetCurrentAppVersionExact(v string)`

SetCurrentAppVersionExact sets CurrentAppVersionExact field to given value.

### HasCurrentAppVersionExact

`func (o *V2controllersCreatableChartRelease) HasCurrentAppVersionExact() bool`

HasCurrentAppVersionExact returns a boolean if a field has been set.

### GetCurrentChartVersionExact

`func (o *V2controllersCreatableChartRelease) GetCurrentChartVersionExact() string`

GetCurrentChartVersionExact returns the CurrentChartVersionExact field if non-nil, zero value otherwise.

### GetCurrentChartVersionExactOk

`func (o *V2controllersCreatableChartRelease) GetCurrentChartVersionExactOk() (*string, bool)`

GetCurrentChartVersionExactOk returns a tuple with the CurrentChartVersionExact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentChartVersionExact

`func (o *V2controllersCreatableChartRelease) SetCurrentChartVersionExact(v string)`

SetCurrentChartVersionExact sets CurrentChartVersionExact field to given value.

### HasCurrentChartVersionExact

`func (o *V2controllersCreatableChartRelease) HasCurrentChartVersionExact() bool`

HasCurrentChartVersionExact returns a boolean if a field has been set.

### GetEnvironment

`func (o *V2controllersCreatableChartRelease) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *V2controllersCreatableChartRelease) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *V2controllersCreatableChartRelease) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *V2controllersCreatableChartRelease) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetHelmfileRef

`func (o *V2controllersCreatableChartRelease) GetHelmfileRef() string`

GetHelmfileRef returns the HelmfileRef field if non-nil, zero value otherwise.

### GetHelmfileRefOk

`func (o *V2controllersCreatableChartRelease) GetHelmfileRefOk() (*string, bool)`

GetHelmfileRefOk returns a tuple with the HelmfileRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmfileRef

`func (o *V2controllersCreatableChartRelease) SetHelmfileRef(v string)`

SetHelmfileRef sets HelmfileRef field to given value.

### HasHelmfileRef

`func (o *V2controllersCreatableChartRelease) HasHelmfileRef() bool`

HasHelmfileRef returns a boolean if a field has been set.

### GetName

`func (o *V2controllersCreatableChartRelease) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2controllersCreatableChartRelease) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2controllersCreatableChartRelease) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2controllersCreatableChartRelease) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *V2controllersCreatableChartRelease) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *V2controllersCreatableChartRelease) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *V2controllersCreatableChartRelease) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *V2controllersCreatableChartRelease) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetTargetAppVersionBranch

`func (o *V2controllersCreatableChartRelease) GetTargetAppVersionBranch() string`

GetTargetAppVersionBranch returns the TargetAppVersionBranch field if non-nil, zero value otherwise.

### GetTargetAppVersionBranchOk

`func (o *V2controllersCreatableChartRelease) GetTargetAppVersionBranchOk() (*string, bool)`

GetTargetAppVersionBranchOk returns a tuple with the TargetAppVersionBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAppVersionBranch

`func (o *V2controllersCreatableChartRelease) SetTargetAppVersionBranch(v string)`

SetTargetAppVersionBranch sets TargetAppVersionBranch field to given value.

### HasTargetAppVersionBranch

`func (o *V2controllersCreatableChartRelease) HasTargetAppVersionBranch() bool`

HasTargetAppVersionBranch returns a boolean if a field has been set.

### GetTargetAppVersionCommit

`func (o *V2controllersCreatableChartRelease) GetTargetAppVersionCommit() string`

GetTargetAppVersionCommit returns the TargetAppVersionCommit field if non-nil, zero value otherwise.

### GetTargetAppVersionCommitOk

`func (o *V2controllersCreatableChartRelease) GetTargetAppVersionCommitOk() (*string, bool)`

GetTargetAppVersionCommitOk returns a tuple with the TargetAppVersionCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAppVersionCommit

`func (o *V2controllersCreatableChartRelease) SetTargetAppVersionCommit(v string)`

SetTargetAppVersionCommit sets TargetAppVersionCommit field to given value.

### HasTargetAppVersionCommit

`func (o *V2controllersCreatableChartRelease) HasTargetAppVersionCommit() bool`

HasTargetAppVersionCommit returns a boolean if a field has been set.

### GetTargetAppVersionExact

`func (o *V2controllersCreatableChartRelease) GetTargetAppVersionExact() string`

GetTargetAppVersionExact returns the TargetAppVersionExact field if non-nil, zero value otherwise.

### GetTargetAppVersionExactOk

`func (o *V2controllersCreatableChartRelease) GetTargetAppVersionExactOk() (*string, bool)`

GetTargetAppVersionExactOk returns a tuple with the TargetAppVersionExact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAppVersionExact

`func (o *V2controllersCreatableChartRelease) SetTargetAppVersionExact(v string)`

SetTargetAppVersionExact sets TargetAppVersionExact field to given value.

### HasTargetAppVersionExact

`func (o *V2controllersCreatableChartRelease) HasTargetAppVersionExact() bool`

HasTargetAppVersionExact returns a boolean if a field has been set.

### GetTargetAppVersionUse

`func (o *V2controllersCreatableChartRelease) GetTargetAppVersionUse() string`

GetTargetAppVersionUse returns the TargetAppVersionUse field if non-nil, zero value otherwise.

### GetTargetAppVersionUseOk

`func (o *V2controllersCreatableChartRelease) GetTargetAppVersionUseOk() (*string, bool)`

GetTargetAppVersionUseOk returns a tuple with the TargetAppVersionUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAppVersionUse

`func (o *V2controllersCreatableChartRelease) SetTargetAppVersionUse(v string)`

SetTargetAppVersionUse sets TargetAppVersionUse field to given value.

### HasTargetAppVersionUse

`func (o *V2controllersCreatableChartRelease) HasTargetAppVersionUse() bool`

HasTargetAppVersionUse returns a boolean if a field has been set.

### GetTargetChartVersionExact

`func (o *V2controllersCreatableChartRelease) GetTargetChartVersionExact() string`

GetTargetChartVersionExact returns the TargetChartVersionExact field if non-nil, zero value otherwise.

### GetTargetChartVersionExactOk

`func (o *V2controllersCreatableChartRelease) GetTargetChartVersionExactOk() (*string, bool)`

GetTargetChartVersionExactOk returns a tuple with the TargetChartVersionExact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetChartVersionExact

`func (o *V2controllersCreatableChartRelease) SetTargetChartVersionExact(v string)`

SetTargetChartVersionExact sets TargetChartVersionExact field to given value.

### HasTargetChartVersionExact

`func (o *V2controllersCreatableChartRelease) HasTargetChartVersionExact() bool`

HasTargetChartVersionExact returns a boolean if a field has been set.

### GetTargetChartVersionUse

`func (o *V2controllersCreatableChartRelease) GetTargetChartVersionUse() string`

GetTargetChartVersionUse returns the TargetChartVersionUse field if non-nil, zero value otherwise.

### GetTargetChartVersionUseOk

`func (o *V2controllersCreatableChartRelease) GetTargetChartVersionUseOk() (*string, bool)`

GetTargetChartVersionUseOk returns a tuple with the TargetChartVersionUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetChartVersionUse

`func (o *V2controllersCreatableChartRelease) SetTargetChartVersionUse(v string)`

SetTargetChartVersionUse sets TargetChartVersionUse field to given value.

### HasTargetChartVersionUse

`func (o *V2controllersCreatableChartRelease) HasTargetChartVersionUse() bool`

HasTargetChartVersionUse returns a boolean if a field has been set.

### GetThelmaMode

`func (o *V2controllersCreatableChartRelease) GetThelmaMode() string`

GetThelmaMode returns the ThelmaMode field if non-nil, zero value otherwise.

### GetThelmaModeOk

`func (o *V2controllersCreatableChartRelease) GetThelmaModeOk() (*string, bool)`

GetThelmaModeOk returns a tuple with the ThelmaMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThelmaMode

`func (o *V2controllersCreatableChartRelease) SetThelmaMode(v string)`

SetThelmaMode sets ThelmaMode field to given value.

### HasThelmaMode

`func (o *V2controllersCreatableChartRelease) HasThelmaMode() bool`

HasThelmaMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


