# V2controllersChartRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chart** | Pointer to **string** | Required when creating | [optional] 
**ChartInfo** | Pointer to [**V2controllersChart**](V2controllersChart.md) |  | [optional] 
**Cluster** | Pointer to **string** | When creating, will default the environment&#39;s default cluster, if provided. Either this or environment must be provided. | [optional] 
**ClusterInfo** | Pointer to [**V2controllersCluster**](V2controllersCluster.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CurrentAppVersionExact** | Pointer to **string** |  | [optional] 
**CurrentChartVersionExact** | Pointer to **string** |  | [optional] 
**DestinationType** | Pointer to **string** | Calculated field | [optional] 
**Environment** | Pointer to **string** | Either this or cluster must be provided. | [optional] 
**EnvironmentInfo** | Pointer to [**V2controllersEnvironment**](V2controllersEnvironment.md) |  | [optional] 
**HelmfileRef** | Pointer to **string** |  | [optional] [default to "HEAD"]
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** | When creating, will be calculated if left empty | [optional] 
**Namespace** | Pointer to **string** | When creating, will default to the environment&#39;s default namespace, if provided | [optional] 
**TargetAppVersionBranch** | Pointer to **string** | When creating, will default to the app&#39;s main branch if it has one recorded | [optional] 
**TargetAppVersionCommit** | Pointer to **string** |  | [optional] 
**TargetAppVersionExact** | Pointer to **string** |  | [optional] 
**TargetAppVersionUse** | Pointer to **string** | When creating, will default to referencing any provided target app version field (exact, then commit, then branch) | [optional] 
**TargetChartVersionExact** | Pointer to **string** |  | [optional] 
**TargetChartVersionUse** | Pointer to **string** | When creating, will default to latest unless an exact target chart version is provided | [optional] 
**ThelmaMode** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewV2controllersChartRelease

`func NewV2controllersChartRelease() *V2controllersChartRelease`

NewV2controllersChartRelease instantiates a new V2controllersChartRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2controllersChartReleaseWithDefaults

`func NewV2controllersChartReleaseWithDefaults() *V2controllersChartRelease`

NewV2controllersChartReleaseWithDefaults instantiates a new V2controllersChartRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChart

`func (o *V2controllersChartRelease) GetChart() string`

GetChart returns the Chart field if non-nil, zero value otherwise.

### GetChartOk

`func (o *V2controllersChartRelease) GetChartOk() (*string, bool)`

GetChartOk returns a tuple with the Chart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChart

`func (o *V2controllersChartRelease) SetChart(v string)`

SetChart sets Chart field to given value.

### HasChart

`func (o *V2controllersChartRelease) HasChart() bool`

HasChart returns a boolean if a field has been set.

### GetChartInfo

`func (o *V2controllersChartRelease) GetChartInfo() V2controllersChart`

GetChartInfo returns the ChartInfo field if non-nil, zero value otherwise.

### GetChartInfoOk

`func (o *V2controllersChartRelease) GetChartInfoOk() (*V2controllersChart, bool)`

GetChartInfoOk returns a tuple with the ChartInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartInfo

`func (o *V2controllersChartRelease) SetChartInfo(v V2controllersChart)`

SetChartInfo sets ChartInfo field to given value.

### HasChartInfo

`func (o *V2controllersChartRelease) HasChartInfo() bool`

HasChartInfo returns a boolean if a field has been set.

### GetCluster

`func (o *V2controllersChartRelease) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V2controllersChartRelease) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V2controllersChartRelease) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *V2controllersChartRelease) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterInfo

`func (o *V2controllersChartRelease) GetClusterInfo() V2controllersCluster`

GetClusterInfo returns the ClusterInfo field if non-nil, zero value otherwise.

### GetClusterInfoOk

`func (o *V2controllersChartRelease) GetClusterInfoOk() (*V2controllersCluster, bool)`

GetClusterInfoOk returns a tuple with the ClusterInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterInfo

`func (o *V2controllersChartRelease) SetClusterInfo(v V2controllersCluster)`

SetClusterInfo sets ClusterInfo field to given value.

### HasClusterInfo

`func (o *V2controllersChartRelease) HasClusterInfo() bool`

HasClusterInfo returns a boolean if a field has been set.

### GetCreatedAt

`func (o *V2controllersChartRelease) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V2controllersChartRelease) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V2controllersChartRelease) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *V2controllersChartRelease) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrentAppVersionExact

`func (o *V2controllersChartRelease) GetCurrentAppVersionExact() string`

GetCurrentAppVersionExact returns the CurrentAppVersionExact field if non-nil, zero value otherwise.

### GetCurrentAppVersionExactOk

`func (o *V2controllersChartRelease) GetCurrentAppVersionExactOk() (*string, bool)`

GetCurrentAppVersionExactOk returns a tuple with the CurrentAppVersionExact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAppVersionExact

`func (o *V2controllersChartRelease) SetCurrentAppVersionExact(v string)`

SetCurrentAppVersionExact sets CurrentAppVersionExact field to given value.

### HasCurrentAppVersionExact

`func (o *V2controllersChartRelease) HasCurrentAppVersionExact() bool`

HasCurrentAppVersionExact returns a boolean if a field has been set.

### GetCurrentChartVersionExact

`func (o *V2controllersChartRelease) GetCurrentChartVersionExact() string`

GetCurrentChartVersionExact returns the CurrentChartVersionExact field if non-nil, zero value otherwise.

### GetCurrentChartVersionExactOk

`func (o *V2controllersChartRelease) GetCurrentChartVersionExactOk() (*string, bool)`

GetCurrentChartVersionExactOk returns a tuple with the CurrentChartVersionExact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentChartVersionExact

`func (o *V2controllersChartRelease) SetCurrentChartVersionExact(v string)`

SetCurrentChartVersionExact sets CurrentChartVersionExact field to given value.

### HasCurrentChartVersionExact

`func (o *V2controllersChartRelease) HasCurrentChartVersionExact() bool`

HasCurrentChartVersionExact returns a boolean if a field has been set.

### GetDestinationType

`func (o *V2controllersChartRelease) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *V2controllersChartRelease) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *V2controllersChartRelease) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *V2controllersChartRelease) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetEnvironment

`func (o *V2controllersChartRelease) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *V2controllersChartRelease) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *V2controllersChartRelease) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *V2controllersChartRelease) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetEnvironmentInfo

`func (o *V2controllersChartRelease) GetEnvironmentInfo() V2controllersEnvironment`

GetEnvironmentInfo returns the EnvironmentInfo field if non-nil, zero value otherwise.

### GetEnvironmentInfoOk

`func (o *V2controllersChartRelease) GetEnvironmentInfoOk() (*V2controllersEnvironment, bool)`

GetEnvironmentInfoOk returns a tuple with the EnvironmentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentInfo

`func (o *V2controllersChartRelease) SetEnvironmentInfo(v V2controllersEnvironment)`

SetEnvironmentInfo sets EnvironmentInfo field to given value.

### HasEnvironmentInfo

`func (o *V2controllersChartRelease) HasEnvironmentInfo() bool`

HasEnvironmentInfo returns a boolean if a field has been set.

### GetHelmfileRef

`func (o *V2controllersChartRelease) GetHelmfileRef() string`

GetHelmfileRef returns the HelmfileRef field if non-nil, zero value otherwise.

### GetHelmfileRefOk

`func (o *V2controllersChartRelease) GetHelmfileRefOk() (*string, bool)`

GetHelmfileRefOk returns a tuple with the HelmfileRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmfileRef

`func (o *V2controllersChartRelease) SetHelmfileRef(v string)`

SetHelmfileRef sets HelmfileRef field to given value.

### HasHelmfileRef

`func (o *V2controllersChartRelease) HasHelmfileRef() bool`

HasHelmfileRef returns a boolean if a field has been set.

### GetId

`func (o *V2controllersChartRelease) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2controllersChartRelease) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2controllersChartRelease) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V2controllersChartRelease) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V2controllersChartRelease) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2controllersChartRelease) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2controllersChartRelease) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2controllersChartRelease) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *V2controllersChartRelease) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *V2controllersChartRelease) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *V2controllersChartRelease) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *V2controllersChartRelease) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetTargetAppVersionBranch

`func (o *V2controllersChartRelease) GetTargetAppVersionBranch() string`

GetTargetAppVersionBranch returns the TargetAppVersionBranch field if non-nil, zero value otherwise.

### GetTargetAppVersionBranchOk

`func (o *V2controllersChartRelease) GetTargetAppVersionBranchOk() (*string, bool)`

GetTargetAppVersionBranchOk returns a tuple with the TargetAppVersionBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAppVersionBranch

`func (o *V2controllersChartRelease) SetTargetAppVersionBranch(v string)`

SetTargetAppVersionBranch sets TargetAppVersionBranch field to given value.

### HasTargetAppVersionBranch

`func (o *V2controllersChartRelease) HasTargetAppVersionBranch() bool`

HasTargetAppVersionBranch returns a boolean if a field has been set.

### GetTargetAppVersionCommit

`func (o *V2controllersChartRelease) GetTargetAppVersionCommit() string`

GetTargetAppVersionCommit returns the TargetAppVersionCommit field if non-nil, zero value otherwise.

### GetTargetAppVersionCommitOk

`func (o *V2controllersChartRelease) GetTargetAppVersionCommitOk() (*string, bool)`

GetTargetAppVersionCommitOk returns a tuple with the TargetAppVersionCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAppVersionCommit

`func (o *V2controllersChartRelease) SetTargetAppVersionCommit(v string)`

SetTargetAppVersionCommit sets TargetAppVersionCommit field to given value.

### HasTargetAppVersionCommit

`func (o *V2controllersChartRelease) HasTargetAppVersionCommit() bool`

HasTargetAppVersionCommit returns a boolean if a field has been set.

### GetTargetAppVersionExact

`func (o *V2controllersChartRelease) GetTargetAppVersionExact() string`

GetTargetAppVersionExact returns the TargetAppVersionExact field if non-nil, zero value otherwise.

### GetTargetAppVersionExactOk

`func (o *V2controllersChartRelease) GetTargetAppVersionExactOk() (*string, bool)`

GetTargetAppVersionExactOk returns a tuple with the TargetAppVersionExact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAppVersionExact

`func (o *V2controllersChartRelease) SetTargetAppVersionExact(v string)`

SetTargetAppVersionExact sets TargetAppVersionExact field to given value.

### HasTargetAppVersionExact

`func (o *V2controllersChartRelease) HasTargetAppVersionExact() bool`

HasTargetAppVersionExact returns a boolean if a field has been set.

### GetTargetAppVersionUse

`func (o *V2controllersChartRelease) GetTargetAppVersionUse() string`

GetTargetAppVersionUse returns the TargetAppVersionUse field if non-nil, zero value otherwise.

### GetTargetAppVersionUseOk

`func (o *V2controllersChartRelease) GetTargetAppVersionUseOk() (*string, bool)`

GetTargetAppVersionUseOk returns a tuple with the TargetAppVersionUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAppVersionUse

`func (o *V2controllersChartRelease) SetTargetAppVersionUse(v string)`

SetTargetAppVersionUse sets TargetAppVersionUse field to given value.

### HasTargetAppVersionUse

`func (o *V2controllersChartRelease) HasTargetAppVersionUse() bool`

HasTargetAppVersionUse returns a boolean if a field has been set.

### GetTargetChartVersionExact

`func (o *V2controllersChartRelease) GetTargetChartVersionExact() string`

GetTargetChartVersionExact returns the TargetChartVersionExact field if non-nil, zero value otherwise.

### GetTargetChartVersionExactOk

`func (o *V2controllersChartRelease) GetTargetChartVersionExactOk() (*string, bool)`

GetTargetChartVersionExactOk returns a tuple with the TargetChartVersionExact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetChartVersionExact

`func (o *V2controllersChartRelease) SetTargetChartVersionExact(v string)`

SetTargetChartVersionExact sets TargetChartVersionExact field to given value.

### HasTargetChartVersionExact

`func (o *V2controllersChartRelease) HasTargetChartVersionExact() bool`

HasTargetChartVersionExact returns a boolean if a field has been set.

### GetTargetChartVersionUse

`func (o *V2controllersChartRelease) GetTargetChartVersionUse() string`

GetTargetChartVersionUse returns the TargetChartVersionUse field if non-nil, zero value otherwise.

### GetTargetChartVersionUseOk

`func (o *V2controllersChartRelease) GetTargetChartVersionUseOk() (*string, bool)`

GetTargetChartVersionUseOk returns a tuple with the TargetChartVersionUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetChartVersionUse

`func (o *V2controllersChartRelease) SetTargetChartVersionUse(v string)`

SetTargetChartVersionUse sets TargetChartVersionUse field to given value.

### HasTargetChartVersionUse

`func (o *V2controllersChartRelease) HasTargetChartVersionUse() bool`

HasTargetChartVersionUse returns a boolean if a field has been set.

### GetThelmaMode

`func (o *V2controllersChartRelease) GetThelmaMode() string`

GetThelmaMode returns the ThelmaMode field if non-nil, zero value otherwise.

### GetThelmaModeOk

`func (o *V2controllersChartRelease) GetThelmaModeOk() (*string, bool)`

GetThelmaModeOk returns a tuple with the ThelmaMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThelmaMode

`func (o *V2controllersChartRelease) SetThelmaMode(v string)`

SetThelmaMode sets ThelmaMode field to given value.

### HasThelmaMode

`func (o *V2controllersChartRelease) HasThelmaMode() bool`

HasThelmaMode returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *V2controllersChartRelease) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *V2controllersChartRelease) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *V2controllersChartRelease) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *V2controllersChartRelease) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


