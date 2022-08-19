# V2controllersEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base** | Pointer to **string** | Required when creating | [optional] 
**ChartReleasesFromTemplate** | Pointer to **bool** | Upon creation of a dynamic environment, if this is true the template&#39;s chart releases will be copied to the new environment | [optional] [default to true]
**CreatedAt** | Pointer to **string** |  | [optional] 
**DefaultCluster** | Pointer to **string** |  | [optional] 
**DefaultClusterInfo** | Pointer to [**V2controllersCluster**](V2controllersCluster.md) |  | [optional] 
**DefaultNamespace** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Lifecycle** | Pointer to **string** |  | [optional] [default to "dynamic"]
**Name** | Pointer to **string** | When creating, will be calculated if dynamic, required otherwise | [optional] 
**Owner** | Pointer to **string** | When creating, will be set to your email | [optional] 
**RequiresSuitability** | Pointer to **bool** |  | [optional] [default to false]
**TemplateEnvironment** | Pointer to **string** | Required for dynamic environments | [optional] 
**TemplateEnvironmentInfo** | Pointer to **map[string]interface{}** | Single-layer recursive; provides info of the template environment if this environment has one | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**ValuesName** | Pointer to **string** |  | [optional] 

## Methods

### NewV2controllersEnvironment

`func NewV2controllersEnvironment() *V2controllersEnvironment`

NewV2controllersEnvironment instantiates a new V2controllersEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2controllersEnvironmentWithDefaults

`func NewV2controllersEnvironmentWithDefaults() *V2controllersEnvironment`

NewV2controllersEnvironmentWithDefaults instantiates a new V2controllersEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase

`func (o *V2controllersEnvironment) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *V2controllersEnvironment) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *V2controllersEnvironment) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *V2controllersEnvironment) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetChartReleasesFromTemplate

`func (o *V2controllersEnvironment) GetChartReleasesFromTemplate() bool`

GetChartReleasesFromTemplate returns the ChartReleasesFromTemplate field if non-nil, zero value otherwise.

### GetChartReleasesFromTemplateOk

`func (o *V2controllersEnvironment) GetChartReleasesFromTemplateOk() (*bool, bool)`

GetChartReleasesFromTemplateOk returns a tuple with the ChartReleasesFromTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartReleasesFromTemplate

`func (o *V2controllersEnvironment) SetChartReleasesFromTemplate(v bool)`

SetChartReleasesFromTemplate sets ChartReleasesFromTemplate field to given value.

### HasChartReleasesFromTemplate

`func (o *V2controllersEnvironment) HasChartReleasesFromTemplate() bool`

HasChartReleasesFromTemplate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *V2controllersEnvironment) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V2controllersEnvironment) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V2controllersEnvironment) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *V2controllersEnvironment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefaultCluster

`func (o *V2controllersEnvironment) GetDefaultCluster() string`

GetDefaultCluster returns the DefaultCluster field if non-nil, zero value otherwise.

### GetDefaultClusterOk

`func (o *V2controllersEnvironment) GetDefaultClusterOk() (*string, bool)`

GetDefaultClusterOk returns a tuple with the DefaultCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCluster

`func (o *V2controllersEnvironment) SetDefaultCluster(v string)`

SetDefaultCluster sets DefaultCluster field to given value.

### HasDefaultCluster

`func (o *V2controllersEnvironment) HasDefaultCluster() bool`

HasDefaultCluster returns a boolean if a field has been set.

### GetDefaultClusterInfo

`func (o *V2controllersEnvironment) GetDefaultClusterInfo() V2controllersCluster`

GetDefaultClusterInfo returns the DefaultClusterInfo field if non-nil, zero value otherwise.

### GetDefaultClusterInfoOk

`func (o *V2controllersEnvironment) GetDefaultClusterInfoOk() (*V2controllersCluster, bool)`

GetDefaultClusterInfoOk returns a tuple with the DefaultClusterInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultClusterInfo

`func (o *V2controllersEnvironment) SetDefaultClusterInfo(v V2controllersCluster)`

SetDefaultClusterInfo sets DefaultClusterInfo field to given value.

### HasDefaultClusterInfo

`func (o *V2controllersEnvironment) HasDefaultClusterInfo() bool`

HasDefaultClusterInfo returns a boolean if a field has been set.

### GetDefaultNamespace

`func (o *V2controllersEnvironment) GetDefaultNamespace() string`

GetDefaultNamespace returns the DefaultNamespace field if non-nil, zero value otherwise.

### GetDefaultNamespaceOk

`func (o *V2controllersEnvironment) GetDefaultNamespaceOk() (*string, bool)`

GetDefaultNamespaceOk returns a tuple with the DefaultNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNamespace

`func (o *V2controllersEnvironment) SetDefaultNamespace(v string)`

SetDefaultNamespace sets DefaultNamespace field to given value.

### HasDefaultNamespace

`func (o *V2controllersEnvironment) HasDefaultNamespace() bool`

HasDefaultNamespace returns a boolean if a field has been set.

### GetId

`func (o *V2controllersEnvironment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2controllersEnvironment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2controllersEnvironment) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V2controllersEnvironment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLifecycle

`func (o *V2controllersEnvironment) GetLifecycle() string`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *V2controllersEnvironment) GetLifecycleOk() (*string, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *V2controllersEnvironment) SetLifecycle(v string)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *V2controllersEnvironment) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetName

`func (o *V2controllersEnvironment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2controllersEnvironment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2controllersEnvironment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2controllersEnvironment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *V2controllersEnvironment) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *V2controllersEnvironment) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *V2controllersEnvironment) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *V2controllersEnvironment) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRequiresSuitability

`func (o *V2controllersEnvironment) GetRequiresSuitability() bool`

GetRequiresSuitability returns the RequiresSuitability field if non-nil, zero value otherwise.

### GetRequiresSuitabilityOk

`func (o *V2controllersEnvironment) GetRequiresSuitabilityOk() (*bool, bool)`

GetRequiresSuitabilityOk returns a tuple with the RequiresSuitability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresSuitability

`func (o *V2controllersEnvironment) SetRequiresSuitability(v bool)`

SetRequiresSuitability sets RequiresSuitability field to given value.

### HasRequiresSuitability

`func (o *V2controllersEnvironment) HasRequiresSuitability() bool`

HasRequiresSuitability returns a boolean if a field has been set.

### GetTemplateEnvironment

`func (o *V2controllersEnvironment) GetTemplateEnvironment() string`

GetTemplateEnvironment returns the TemplateEnvironment field if non-nil, zero value otherwise.

### GetTemplateEnvironmentOk

`func (o *V2controllersEnvironment) GetTemplateEnvironmentOk() (*string, bool)`

GetTemplateEnvironmentOk returns a tuple with the TemplateEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateEnvironment

`func (o *V2controllersEnvironment) SetTemplateEnvironment(v string)`

SetTemplateEnvironment sets TemplateEnvironment field to given value.

### HasTemplateEnvironment

`func (o *V2controllersEnvironment) HasTemplateEnvironment() bool`

HasTemplateEnvironment returns a boolean if a field has been set.

### GetTemplateEnvironmentInfo

`func (o *V2controllersEnvironment) GetTemplateEnvironmentInfo() map[string]interface{}`

GetTemplateEnvironmentInfo returns the TemplateEnvironmentInfo field if non-nil, zero value otherwise.

### GetTemplateEnvironmentInfoOk

`func (o *V2controllersEnvironment) GetTemplateEnvironmentInfoOk() (*map[string]interface{}, bool)`

GetTemplateEnvironmentInfoOk returns a tuple with the TemplateEnvironmentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateEnvironmentInfo

`func (o *V2controllersEnvironment) SetTemplateEnvironmentInfo(v map[string]interface{})`

SetTemplateEnvironmentInfo sets TemplateEnvironmentInfo field to given value.

### HasTemplateEnvironmentInfo

`func (o *V2controllersEnvironment) HasTemplateEnvironmentInfo() bool`

HasTemplateEnvironmentInfo returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *V2controllersEnvironment) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *V2controllersEnvironment) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *V2controllersEnvironment) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *V2controllersEnvironment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetValuesName

`func (o *V2controllersEnvironment) GetValuesName() string`

GetValuesName returns the ValuesName field if non-nil, zero value otherwise.

### GetValuesNameOk

`func (o *V2controllersEnvironment) GetValuesNameOk() (*string, bool)`

GetValuesNameOk returns a tuple with the ValuesName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesName

`func (o *V2controllersEnvironment) SetValuesName(v string)`

SetValuesName sets ValuesName field to given value.

### HasValuesName

`func (o *V2controllersEnvironment) HasValuesName() bool`

HasValuesName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


