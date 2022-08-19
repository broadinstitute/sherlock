# V2controllersCreatableEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base** | Pointer to **string** | Required when creating | [optional] 
**ChartReleasesFromTemplate** | Pointer to **bool** | Upon creation of a dynamic environment, if this is true the template&#39;s chart releases will be copied to the new environment | [optional] [default to true]
**DefaultCluster** | Pointer to **string** |  | [optional] 
**DefaultNamespace** | Pointer to **string** |  | [optional] 
**Lifecycle** | Pointer to **string** |  | [optional] [default to "dynamic"]
**Name** | Pointer to **string** | When creating, will be calculated if dynamic, required otherwise | [optional] 
**Owner** | Pointer to **string** | When creating, will be set to your email | [optional] 
**RequiresSuitability** | Pointer to **bool** |  | [optional] [default to false]
**TemplateEnvironment** | Pointer to **string** | Required for dynamic environments | [optional] 

## Methods

### NewV2controllersCreatableEnvironment

`func NewV2controllersCreatableEnvironment() *V2controllersCreatableEnvironment`

NewV2controllersCreatableEnvironment instantiates a new V2controllersCreatableEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2controllersCreatableEnvironmentWithDefaults

`func NewV2controllersCreatableEnvironmentWithDefaults() *V2controllersCreatableEnvironment`

NewV2controllersCreatableEnvironmentWithDefaults instantiates a new V2controllersCreatableEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase

`func (o *V2controllersCreatableEnvironment) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *V2controllersCreatableEnvironment) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *V2controllersCreatableEnvironment) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *V2controllersCreatableEnvironment) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetChartReleasesFromTemplate

`func (o *V2controllersCreatableEnvironment) GetChartReleasesFromTemplate() bool`

GetChartReleasesFromTemplate returns the ChartReleasesFromTemplate field if non-nil, zero value otherwise.

### GetChartReleasesFromTemplateOk

`func (o *V2controllersCreatableEnvironment) GetChartReleasesFromTemplateOk() (*bool, bool)`

GetChartReleasesFromTemplateOk returns a tuple with the ChartReleasesFromTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartReleasesFromTemplate

`func (o *V2controllersCreatableEnvironment) SetChartReleasesFromTemplate(v bool)`

SetChartReleasesFromTemplate sets ChartReleasesFromTemplate field to given value.

### HasChartReleasesFromTemplate

`func (o *V2controllersCreatableEnvironment) HasChartReleasesFromTemplate() bool`

HasChartReleasesFromTemplate returns a boolean if a field has been set.

### GetDefaultCluster

`func (o *V2controllersCreatableEnvironment) GetDefaultCluster() string`

GetDefaultCluster returns the DefaultCluster field if non-nil, zero value otherwise.

### GetDefaultClusterOk

`func (o *V2controllersCreatableEnvironment) GetDefaultClusterOk() (*string, bool)`

GetDefaultClusterOk returns a tuple with the DefaultCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCluster

`func (o *V2controllersCreatableEnvironment) SetDefaultCluster(v string)`

SetDefaultCluster sets DefaultCluster field to given value.

### HasDefaultCluster

`func (o *V2controllersCreatableEnvironment) HasDefaultCluster() bool`

HasDefaultCluster returns a boolean if a field has been set.

### GetDefaultNamespace

`func (o *V2controllersCreatableEnvironment) GetDefaultNamespace() string`

GetDefaultNamespace returns the DefaultNamespace field if non-nil, zero value otherwise.

### GetDefaultNamespaceOk

`func (o *V2controllersCreatableEnvironment) GetDefaultNamespaceOk() (*string, bool)`

GetDefaultNamespaceOk returns a tuple with the DefaultNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNamespace

`func (o *V2controllersCreatableEnvironment) SetDefaultNamespace(v string)`

SetDefaultNamespace sets DefaultNamespace field to given value.

### HasDefaultNamespace

`func (o *V2controllersCreatableEnvironment) HasDefaultNamespace() bool`

HasDefaultNamespace returns a boolean if a field has been set.

### GetLifecycle

`func (o *V2controllersCreatableEnvironment) GetLifecycle() string`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *V2controllersCreatableEnvironment) GetLifecycleOk() (*string, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *V2controllersCreatableEnvironment) SetLifecycle(v string)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *V2controllersCreatableEnvironment) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetName

`func (o *V2controllersCreatableEnvironment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2controllersCreatableEnvironment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2controllersCreatableEnvironment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2controllersCreatableEnvironment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *V2controllersCreatableEnvironment) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *V2controllersCreatableEnvironment) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *V2controllersCreatableEnvironment) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *V2controllersCreatableEnvironment) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRequiresSuitability

`func (o *V2controllersCreatableEnvironment) GetRequiresSuitability() bool`

GetRequiresSuitability returns the RequiresSuitability field if non-nil, zero value otherwise.

### GetRequiresSuitabilityOk

`func (o *V2controllersCreatableEnvironment) GetRequiresSuitabilityOk() (*bool, bool)`

GetRequiresSuitabilityOk returns a tuple with the RequiresSuitability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresSuitability

`func (o *V2controllersCreatableEnvironment) SetRequiresSuitability(v bool)`

SetRequiresSuitability sets RequiresSuitability field to given value.

### HasRequiresSuitability

`func (o *V2controllersCreatableEnvironment) HasRequiresSuitability() bool`

HasRequiresSuitability returns a boolean if a field has been set.

### GetTemplateEnvironment

`func (o *V2controllersCreatableEnvironment) GetTemplateEnvironment() string`

GetTemplateEnvironment returns the TemplateEnvironment field if non-nil, zero value otherwise.

### GetTemplateEnvironmentOk

`func (o *V2controllersCreatableEnvironment) GetTemplateEnvironmentOk() (*string, bool)`

GetTemplateEnvironmentOk returns a tuple with the TemplateEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateEnvironment

`func (o *V2controllersCreatableEnvironment) SetTemplateEnvironment(v string)`

SetTemplateEnvironment sets TemplateEnvironment field to given value.

### HasTemplateEnvironment

`func (o *V2controllersCreatableEnvironment) HasTemplateEnvironment() bool`

HasTemplateEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


