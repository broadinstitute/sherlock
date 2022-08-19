# \ChartReleasesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ChartReleasesGet**](ChartReleasesApi.md#ApiV2ChartReleasesGet) | **Get** /api/v2/chart-releases | List ChartRelease entries
[**ApiV2ChartReleasesPost**](ChartReleasesApi.md#ApiV2ChartReleasesPost) | **Post** /api/v2/chart-releases | Create a new ChartRelease entry
[**ApiV2ChartReleasesSelectorDelete**](ChartReleasesApi.md#ApiV2ChartReleasesSelectorDelete) | **Delete** /api/v2/chart-releases/{selector} | Delete a ChartRelease entry
[**ApiV2ChartReleasesSelectorGet**](ChartReleasesApi.md#ApiV2ChartReleasesSelectorGet) | **Get** /api/v2/chart-releases/{selector} | Get a ChartRelease entry
[**ApiV2ChartReleasesSelectorPatch**](ChartReleasesApi.md#ApiV2ChartReleasesSelectorPatch) | **Patch** /api/v2/chart-releases/{selector} | Edit a ChartRelease entry
[**ApiV2SelectorsChartReleasesSelectorGet**](ChartReleasesApi.md#ApiV2SelectorsChartReleasesSelectorGet) | **Get** /api/v2/selectors/chart-releases/{selector} | List ChartRelease selectors



## ApiV2ChartReleasesGet

> []V2controllersChartRelease ApiV2ChartReleasesGet(ctx).Chart(chart).Cluster(cluster).CreatedAt(createdAt).CurrentAppVersionExact(currentAppVersionExact).CurrentChartVersionExact(currentChartVersionExact).DestinationType(destinationType).Environment(environment).HelmfileRef(helmfileRef).Id(id).Name(name).Namespace(namespace).TargetAppVersionBranch(targetAppVersionBranch).TargetAppVersionCommit(targetAppVersionCommit).TargetAppVersionExact(targetAppVersionExact).TargetAppVersionUse(targetAppVersionUse).TargetChartVersionExact(targetChartVersionExact).TargetChartVersionUse(targetChartVersionUse).ThelmaMode(thelmaMode).UpdatedAt(updatedAt).Limit(limit).Execute()

List ChartRelease entries



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
    chart := "chart_example" // string | Required when creating (optional)
    cluster := "cluster_example" // string | When creating, will default the environment's default cluster, if provided. Either this or environment must be provided. (optional)
    createdAt := "createdAt_example" // string |  (optional)
    currentAppVersionExact := "currentAppVersionExact_example" // string |  (optional)
    currentChartVersionExact := "currentChartVersionExact_example" // string |  (optional)
    destinationType := "destinationType_example" // string | Calculated field (optional)
    environment := "environment_example" // string | Either this or cluster must be provided. (optional)
    helmfileRef := "helmfileRef_example" // string |  (optional) (default to "HEAD")
    id := int32(56) // int32 |  (optional)
    name := "name_example" // string | When creating, will be calculated if left empty (optional)
    namespace := "namespace_example" // string | When creating, will default to the environment's default namespace, if provided (optional)
    targetAppVersionBranch := "targetAppVersionBranch_example" // string | When creating, will default to the app's main branch if it has one recorded (optional)
    targetAppVersionCommit := "targetAppVersionCommit_example" // string |  (optional)
    targetAppVersionExact := "targetAppVersionExact_example" // string |  (optional)
    targetAppVersionUse := "targetAppVersionUse_example" // string | When creating, will default to referencing any provided target app version field (exact, then commit, then branch) (optional)
    targetChartVersionExact := "targetChartVersionExact_example" // string |  (optional)
    targetChartVersionUse := "targetChartVersionUse_example" // string | When creating, will default to latest unless an exact target chart version is provided (optional)
    thelmaMode := "thelmaMode_example" // string |  (optional)
    updatedAt := "updatedAt_example" // string |  (optional)
    limit := int32(56) // int32 | An optional limit to the number of entries returned (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartReleasesApi.ApiV2ChartReleasesGet(context.Background()).Chart(chart).Cluster(cluster).CreatedAt(createdAt).CurrentAppVersionExact(currentAppVersionExact).CurrentChartVersionExact(currentChartVersionExact).DestinationType(destinationType).Environment(environment).HelmfileRef(helmfileRef).Id(id).Name(name).Namespace(namespace).TargetAppVersionBranch(targetAppVersionBranch).TargetAppVersionCommit(targetAppVersionCommit).TargetAppVersionExact(targetAppVersionExact).TargetAppVersionUse(targetAppVersionUse).TargetChartVersionExact(targetChartVersionExact).TargetChartVersionUse(targetChartVersionUse).ThelmaMode(thelmaMode).UpdatedAt(updatedAt).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartReleasesApi.ApiV2ChartReleasesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ChartReleasesGet`: []V2controllersChartRelease
    fmt.Fprintf(os.Stdout, "Response from `ChartReleasesApi.ApiV2ChartReleasesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ChartReleasesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chart** | **string** | Required when creating | 
 **cluster** | **string** | When creating, will default the environment&#39;s default cluster, if provided. Either this or environment must be provided. | 
 **createdAt** | **string** |  | 
 **currentAppVersionExact** | **string** |  | 
 **currentChartVersionExact** | **string** |  | 
 **destinationType** | **string** | Calculated field | 
 **environment** | **string** | Either this or cluster must be provided. | 
 **helmfileRef** | **string** |  | [default to &quot;HEAD&quot;]
 **id** | **int32** |  | 
 **name** | **string** | When creating, will be calculated if left empty | 
 **namespace** | **string** | When creating, will default to the environment&#39;s default namespace, if provided | 
 **targetAppVersionBranch** | **string** | When creating, will default to the app&#39;s main branch if it has one recorded | 
 **targetAppVersionCommit** | **string** |  | 
 **targetAppVersionExact** | **string** |  | 
 **targetAppVersionUse** | **string** | When creating, will default to referencing any provided target app version field (exact, then commit, then branch) | 
 **targetChartVersionExact** | **string** |  | 
 **targetChartVersionUse** | **string** | When creating, will default to latest unless an exact target chart version is provided | 
 **thelmaMode** | **string** |  | 
 **updatedAt** | **string** |  | 
 **limit** | **int32** | An optional limit to the number of entries returned | 

### Return type

[**[]V2controllersChartRelease**](V2controllersChartRelease.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ChartReleasesPost

> V2controllersChartRelease ApiV2ChartReleasesPost(ctx).ChartRelease(chartRelease).Execute()

Create a new ChartRelease entry



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
    chartRelease := *openapiclient.NewV2controllersCreatableChartRelease() // V2controllersCreatableChartRelease | The ChartRelease to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartReleasesApi.ApiV2ChartReleasesPost(context.Background()).ChartRelease(chartRelease).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartReleasesApi.ApiV2ChartReleasesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ChartReleasesPost`: V2controllersChartRelease
    fmt.Fprintf(os.Stdout, "Response from `ChartReleasesApi.ApiV2ChartReleasesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ChartReleasesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chartRelease** | [**V2controllersCreatableChartRelease**](V2controllersCreatableChartRelease.md) | The ChartRelease to create | 

### Return type

[**V2controllersChartRelease**](V2controllersChartRelease.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ChartReleasesSelectorDelete

> V2controllersChartRelease ApiV2ChartReleasesSelectorDelete(ctx, selector).Execute()

Delete a ChartRelease entry



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
    selector := "selector_example" // string | The ChartRelease to delete's selector: name or numeric ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartReleasesApi.ApiV2ChartReleasesSelectorDelete(context.Background(), selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartReleasesApi.ApiV2ChartReleasesSelectorDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ChartReleasesSelectorDelete`: V2controllersChartRelease
    fmt.Fprintf(os.Stdout, "Response from `ChartReleasesApi.ApiV2ChartReleasesSelectorDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selector** | **string** | The ChartRelease to delete&#39;s selector: name or numeric ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ChartReleasesSelectorDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2controllersChartRelease**](V2controllersChartRelease.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ChartReleasesSelectorGet

> V2controllersChartRelease ApiV2ChartReleasesSelectorGet(ctx, selector).Execute()

Get a ChartRelease entry



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
    selector := "selector_example" // string | The ChartRelease to get's selector: name or numeric ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartReleasesApi.ApiV2ChartReleasesSelectorGet(context.Background(), selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartReleasesApi.ApiV2ChartReleasesSelectorGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ChartReleasesSelectorGet`: V2controllersChartRelease
    fmt.Fprintf(os.Stdout, "Response from `ChartReleasesApi.ApiV2ChartReleasesSelectorGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selector** | **string** | The ChartRelease to get&#39;s selector: name or numeric ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ChartReleasesSelectorGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2controllersChartRelease**](V2controllersChartRelease.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ChartReleasesSelectorPatch

> V2controllersChartRelease ApiV2ChartReleasesSelectorPatch(ctx, selector).ChartRelease(chartRelease).Execute()

Edit a ChartRelease entry



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
    selector := "selector_example" // string | The ChartRelease to edit's selector: name or numeric ID
    chartRelease := *openapiclient.NewV2controllersEditableChartRelease() // V2controllersEditableChartRelease | The edits to make to the ChartRelease

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartReleasesApi.ApiV2ChartReleasesSelectorPatch(context.Background(), selector).ChartRelease(chartRelease).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartReleasesApi.ApiV2ChartReleasesSelectorPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ChartReleasesSelectorPatch`: V2controllersChartRelease
    fmt.Fprintf(os.Stdout, "Response from `ChartReleasesApi.ApiV2ChartReleasesSelectorPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selector** | **string** | The ChartRelease to edit&#39;s selector: name or numeric ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ChartReleasesSelectorPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chartRelease** | [**V2controllersEditableChartRelease**](V2controllersEditableChartRelease.md) | The edits to make to the ChartRelease | 

### Return type

[**V2controllersChartRelease**](V2controllersChartRelease.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SelectorsChartReleasesSelectorGet

> []string ApiV2SelectorsChartReleasesSelectorGet(ctx, selector).Execute()

List ChartRelease selectors



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
    selector := "selector_example" // string | The selector of the ChartRelease to list other selectors for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartReleasesApi.ApiV2SelectorsChartReleasesSelectorGet(context.Background(), selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartReleasesApi.ApiV2SelectorsChartReleasesSelectorGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2SelectorsChartReleasesSelectorGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `ChartReleasesApi.ApiV2SelectorsChartReleasesSelectorGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selector** | **string** | The selector of the ChartRelease to list other selectors for | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SelectorsChartReleasesSelectorGetRequest struct via the builder pattern


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

