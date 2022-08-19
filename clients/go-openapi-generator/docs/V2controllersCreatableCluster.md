# V2controllersCreatableCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Required when creating | [optional] 
**AzureSubscription** | Pointer to **string** | Required when creating if providers is &#39;azure&#39; | [optional] 
**Base** | Pointer to **string** | Required when creating | [optional] 
**GoogleProject** | Pointer to **string** | Required when creating if provider is &#39;google&#39; | [optional] 
**Name** | **string** | Required when creating | 
**Provider** | Pointer to **string** |  | [optional] [default to "google"]
**RequiresSuitability** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewV2controllersCreatableCluster

`func NewV2controllersCreatableCluster(name string, ) *V2controllersCreatableCluster`

NewV2controllersCreatableCluster instantiates a new V2controllersCreatableCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2controllersCreatableClusterWithDefaults

`func NewV2controllersCreatableClusterWithDefaults() *V2controllersCreatableCluster`

NewV2controllersCreatableClusterWithDefaults instantiates a new V2controllersCreatableCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *V2controllersCreatableCluster) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *V2controllersCreatableCluster) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *V2controllersCreatableCluster) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *V2controllersCreatableCluster) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAzureSubscription

`func (o *V2controllersCreatableCluster) GetAzureSubscription() string`

GetAzureSubscription returns the AzureSubscription field if non-nil, zero value otherwise.

### GetAzureSubscriptionOk

`func (o *V2controllersCreatableCluster) GetAzureSubscriptionOk() (*string, bool)`

GetAzureSubscriptionOk returns a tuple with the AzureSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscription

`func (o *V2controllersCreatableCluster) SetAzureSubscription(v string)`

SetAzureSubscription sets AzureSubscription field to given value.

### HasAzureSubscription

`func (o *V2controllersCreatableCluster) HasAzureSubscription() bool`

HasAzureSubscription returns a boolean if a field has been set.

### GetBase

`func (o *V2controllersCreatableCluster) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *V2controllersCreatableCluster) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *V2controllersCreatableCluster) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *V2controllersCreatableCluster) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetGoogleProject

`func (o *V2controllersCreatableCluster) GetGoogleProject() string`

GetGoogleProject returns the GoogleProject field if non-nil, zero value otherwise.

### GetGoogleProjectOk

`func (o *V2controllersCreatableCluster) GetGoogleProjectOk() (*string, bool)`

GetGoogleProjectOk returns a tuple with the GoogleProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleProject

`func (o *V2controllersCreatableCluster) SetGoogleProject(v string)`

SetGoogleProject sets GoogleProject field to given value.

### HasGoogleProject

`func (o *V2controllersCreatableCluster) HasGoogleProject() bool`

HasGoogleProject returns a boolean if a field has been set.

### GetName

`func (o *V2controllersCreatableCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2controllersCreatableCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2controllersCreatableCluster) SetName(v string)`

SetName sets Name field to given value.


### GetProvider

`func (o *V2controllersCreatableCluster) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *V2controllersCreatableCluster) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *V2controllersCreatableCluster) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *V2controllersCreatableCluster) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRequiresSuitability

`func (o *V2controllersCreatableCluster) GetRequiresSuitability() bool`

GetRequiresSuitability returns the RequiresSuitability field if non-nil, zero value otherwise.

### GetRequiresSuitabilityOk

`func (o *V2controllersCreatableCluster) GetRequiresSuitabilityOk() (*bool, bool)`

GetRequiresSuitabilityOk returns a tuple with the RequiresSuitability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresSuitability

`func (o *V2controllersCreatableCluster) SetRequiresSuitability(v bool)`

SetRequiresSuitability sets RequiresSuitability field to given value.

### HasRequiresSuitability

`func (o *V2controllersCreatableCluster) HasRequiresSuitability() bool`

HasRequiresSuitability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


