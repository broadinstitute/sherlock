# V2controllersCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Required when creating | [optional] 
**AzureSubscription** | Pointer to **string** | Required when creating if providers is &#39;azure&#39; | [optional] 
**Base** | Pointer to **string** | Required when creating | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**GoogleProject** | Pointer to **string** | Required when creating if provider is &#39;google&#39; | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** | Required when creating | 
**Provider** | Pointer to **string** |  | [optional] [default to "google"]
**RequiresSuitability** | Pointer to **bool** |  | [optional] [default to false]
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewV2controllersCluster

`func NewV2controllersCluster(name string, ) *V2controllersCluster`

NewV2controllersCluster instantiates a new V2controllersCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2controllersClusterWithDefaults

`func NewV2controllersClusterWithDefaults() *V2controllersCluster`

NewV2controllersClusterWithDefaults instantiates a new V2controllersCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *V2controllersCluster) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *V2controllersCluster) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *V2controllersCluster) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *V2controllersCluster) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAzureSubscription

`func (o *V2controllersCluster) GetAzureSubscription() string`

GetAzureSubscription returns the AzureSubscription field if non-nil, zero value otherwise.

### GetAzureSubscriptionOk

`func (o *V2controllersCluster) GetAzureSubscriptionOk() (*string, bool)`

GetAzureSubscriptionOk returns a tuple with the AzureSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscription

`func (o *V2controllersCluster) SetAzureSubscription(v string)`

SetAzureSubscription sets AzureSubscription field to given value.

### HasAzureSubscription

`func (o *V2controllersCluster) HasAzureSubscription() bool`

HasAzureSubscription returns a boolean if a field has been set.

### GetBase

`func (o *V2controllersCluster) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *V2controllersCluster) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *V2controllersCluster) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *V2controllersCluster) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetCreatedAt

`func (o *V2controllersCluster) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V2controllersCluster) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V2controllersCluster) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *V2controllersCluster) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGoogleProject

`func (o *V2controllersCluster) GetGoogleProject() string`

GetGoogleProject returns the GoogleProject field if non-nil, zero value otherwise.

### GetGoogleProjectOk

`func (o *V2controllersCluster) GetGoogleProjectOk() (*string, bool)`

GetGoogleProjectOk returns a tuple with the GoogleProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleProject

`func (o *V2controllersCluster) SetGoogleProject(v string)`

SetGoogleProject sets GoogleProject field to given value.

### HasGoogleProject

`func (o *V2controllersCluster) HasGoogleProject() bool`

HasGoogleProject returns a boolean if a field has been set.

### GetId

`func (o *V2controllersCluster) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2controllersCluster) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2controllersCluster) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V2controllersCluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V2controllersCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2controllersCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2controllersCluster) SetName(v string)`

SetName sets Name field to given value.


### GetProvider

`func (o *V2controllersCluster) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *V2controllersCluster) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *V2controllersCluster) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *V2controllersCluster) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRequiresSuitability

`func (o *V2controllersCluster) GetRequiresSuitability() bool`

GetRequiresSuitability returns the RequiresSuitability field if non-nil, zero value otherwise.

### GetRequiresSuitabilityOk

`func (o *V2controllersCluster) GetRequiresSuitabilityOk() (*bool, bool)`

GetRequiresSuitabilityOk returns a tuple with the RequiresSuitability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresSuitability

`func (o *V2controllersCluster) SetRequiresSuitability(v bool)`

SetRequiresSuitability sets RequiresSuitability field to given value.

### HasRequiresSuitability

`func (o *V2controllersCluster) HasRequiresSuitability() bool`

HasRequiresSuitability returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *V2controllersCluster) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *V2controllersCluster) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *V2controllersCluster) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *V2controllersCluster) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


