# \ClustersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ClustersGet**](ClustersApi.md#ApiV2ClustersGet) | **Get** /api/v2/clusters | List Cluster entries
[**ApiV2ClustersPost**](ClustersApi.md#ApiV2ClustersPost) | **Post** /api/v2/clusters | Create a new Cluster entry
[**ApiV2ClustersSelectorDelete**](ClustersApi.md#ApiV2ClustersSelectorDelete) | **Delete** /api/v2/clusters/{selector} | Delete a Cluster entry
[**ApiV2ClustersSelectorGet**](ClustersApi.md#ApiV2ClustersSelectorGet) | **Get** /api/v2/clusters/{selector} | Get a Cluster entry
[**ApiV2ClustersSelectorPatch**](ClustersApi.md#ApiV2ClustersSelectorPatch) | **Patch** /api/v2/clusters/{selector} | Edit a Cluster entry
[**ApiV2SelectorsClustersSelectorGet**](ClustersApi.md#ApiV2SelectorsClustersSelectorGet) | **Get** /api/v2/selectors/clusters/{selector} | List Cluster selectors



## ApiV2ClustersGet

> []V2controllersCluster ApiV2ClustersGet(ctx).Name(name).Address(address).AzureSubscription(azureSubscription).Base(base).CreatedAt(createdAt).GoogleProject(googleProject).Id(id).Provider(provider).RequiresSuitability(requiresSuitability).UpdatedAt(updatedAt).Limit(limit).Execute()

List Cluster entries



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Required when creating
    address := "address_example" // string | Required when creating (optional)
    azureSubscription := "azureSubscription_example" // string | Required when creating if providers is 'azure' (optional)
    base := "base_example" // string | Required when creating (optional)
    createdAt := "createdAt_example" // string |  (optional)
    googleProject := "googleProject_example" // string | Required when creating if provider is 'google' (optional)
    id := int32(56) // int32 |  (optional)
    provider := "provider_example" // string |  (optional) (default to "google")
    requiresSuitability := true // bool |  (optional) (default to false)
    updatedAt := "updatedAt_example" // string |  (optional)
    limit := int32(56) // int32 | An optional limit to the number of entries returned (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.ApiV2ClustersGet(context.Background()).Name(name).Address(address).AzureSubscription(azureSubscription).Base(base).CreatedAt(createdAt).GoogleProject(googleProject).Id(id).Provider(provider).RequiresSuitability(requiresSuitability).UpdatedAt(updatedAt).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ApiV2ClustersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ClustersGet`: []V2controllersCluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ApiV2ClustersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ClustersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Required when creating | 
 **address** | **string** | Required when creating | 
 **azureSubscription** | **string** | Required when creating if providers is &#39;azure&#39; | 
 **base** | **string** | Required when creating | 
 **createdAt** | **string** |  | 
 **googleProject** | **string** | Required when creating if provider is &#39;google&#39; | 
 **id** | **int32** |  | 
 **provider** | **string** |  | [default to &quot;google&quot;]
 **requiresSuitability** | **bool** |  | [default to false]
 **updatedAt** | **string** |  | 
 **limit** | **int32** | An optional limit to the number of entries returned | 

### Return type

[**[]V2controllersCluster**](V2controllersCluster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ClustersPost

> V2controllersCluster ApiV2ClustersPost(ctx).Cluster(cluster).Execute()

Create a new Cluster entry



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cluster := *openapiclient.NewV2controllersCreatableCluster("Name_example") // V2controllersCreatableCluster | The Cluster to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.ApiV2ClustersPost(context.Background()).Cluster(cluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ApiV2ClustersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ClustersPost`: V2controllersCluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ApiV2ClustersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ClustersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | [**V2controllersCreatableCluster**](V2controllersCreatableCluster.md) | The Cluster to create | 

### Return type

[**V2controllersCluster**](V2controllersCluster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ClustersSelectorDelete

> V2controllersCluster ApiV2ClustersSelectorDelete(ctx, selector).Execute()

Delete a Cluster entry



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    selector := "selector_example" // string | The Cluster to delete's selector: name or numeric ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.ApiV2ClustersSelectorDelete(context.Background(), selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ApiV2ClustersSelectorDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ClustersSelectorDelete`: V2controllersCluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ApiV2ClustersSelectorDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selector** | **string** | The Cluster to delete&#39;s selector: name or numeric ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ClustersSelectorDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2controllersCluster**](V2controllersCluster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ClustersSelectorGet

> V2controllersCluster ApiV2ClustersSelectorGet(ctx, selector).Execute()

Get a Cluster entry



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    selector := "selector_example" // string | The Cluster to get's selector: name or numeric ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.ApiV2ClustersSelectorGet(context.Background(), selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ApiV2ClustersSelectorGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ClustersSelectorGet`: V2controllersCluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ApiV2ClustersSelectorGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selector** | **string** | The Cluster to get&#39;s selector: name or numeric ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ClustersSelectorGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2controllersCluster**](V2controllersCluster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ClustersSelectorPatch

> V2controllersCluster ApiV2ClustersSelectorPatch(ctx, selector).Cluster(cluster).Execute()

Edit a Cluster entry



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    selector := "selector_example" // string | The Cluster to edit's selector: name or numeric ID
    cluster := *openapiclient.NewV2controllersEditableCluster() // V2controllersEditableCluster | The edits to make to the Cluster

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.ApiV2ClustersSelectorPatch(context.Background(), selector).Cluster(cluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ApiV2ClustersSelectorPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ClustersSelectorPatch`: V2controllersCluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ApiV2ClustersSelectorPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selector** | **string** | The Cluster to edit&#39;s selector: name or numeric ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ClustersSelectorPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cluster** | [**V2controllersEditableCluster**](V2controllersEditableCluster.md) | The edits to make to the Cluster | 

### Return type

[**V2controllersCluster**](V2controllersCluster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SelectorsClustersSelectorGet

> []string ApiV2SelectorsClustersSelectorGet(ctx, selector).Execute()

List Cluster selectors



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    selector := "selector_example" // string | The selector of the Cluster to list other selectors for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.ApiV2SelectorsClustersSelectorGet(context.Background(), selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ApiV2SelectorsClustersSelectorGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2SelectorsClustersSelectorGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ApiV2SelectorsClustersSelectorGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selector** | **string** | The selector of the Cluster to list other selectors for | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SelectorsClustersSelectorGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

