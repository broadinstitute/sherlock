# \ChartVersionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ChartVersionsGet**](ChartVersionsApi.md#ApiV2ChartVersionsGet) | **Get** /api/v2/chart-versions | List ChartVersion entries
[**ApiV2ChartVersionsPost**](ChartVersionsApi.md#ApiV2ChartVersionsPost) | **Post** /api/v2/chart-versions | Create a new ChartVersion entry
[**ApiV2ChartVersionsSelectorGet**](ChartVersionsApi.md#ApiV2ChartVersionsSelectorGet) | **Get** /api/v2/chart-versions/{selector} | Get a ChartVersion entry
[**ApiV2SelectorsChartVersionsSelectorGet**](ChartVersionsApi.md#ApiV2SelectorsChartVersionsSelectorGet) | **Get** /api/v2/selectors/chart-versions/{selector} | List ChartVersion selectors



## ApiV2ChartVersionsGet

> []V2controllersChartVersion ApiV2ChartVersionsGet(ctx).Chart(chart).ChartVersion(chartVersion).CreatedAt(createdAt).Id(id).UpdatedAt(updatedAt).Limit(limit).Execute()

List ChartVersion entries



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
    chartVersion := "chartVersion_example" // string | Required when creating (optional)
    createdAt := "createdAt_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)
    updatedAt := "updatedAt_example" // string |  (optional)
    limit := int32(56) // int32 | An optional limit to the number of entries returned (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartVersionsApi.ApiV2ChartVersionsGet(context.Background()).Chart(chart).ChartVersion(chartVersion).CreatedAt(createdAt).Id(id).UpdatedAt(updatedAt).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartVersionsApi.ApiV2ChartVersionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ChartVersionsGet`: []V2controllersChartVersion
    fmt.Fprintf(os.Stdout, "Response from `ChartVersionsApi.ApiV2ChartVersionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ChartVersionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chart** | **string** | Required when creating | 
 **chartVersion** | **string** | Required when creating | 
 **createdAt** | **string** |  | 
 **id** | **int32** |  | 
 **updatedAt** | **string** |  | 
 **limit** | **int32** | An optional limit to the number of entries returned | 

### Return type

[**[]V2controllersChartVersion**](V2controllersChartVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ChartVersionsPost

> V2controllersChartVersion ApiV2ChartVersionsPost(ctx).ChartVersion(chartVersion).Execute()

Create a new ChartVersion entry



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
    chartVersion := *openapiclient.NewV2controllersCreatableChartVersion() // V2controllersCreatableChartVersion | The ChartVersion to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartVersionsApi.ApiV2ChartVersionsPost(context.Background()).ChartVersion(chartVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartVersionsApi.ApiV2ChartVersionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ChartVersionsPost`: V2controllersChartVersion
    fmt.Fprintf(os.Stdout, "Response from `ChartVersionsApi.ApiV2ChartVersionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ChartVersionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chartVersion** | [**V2controllersCreatableChartVersion**](V2controllersCreatableChartVersion.md) | The ChartVersion to create | 

### Return type

[**V2controllersChartVersion**](V2controllersChartVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ChartVersionsSelectorGet

> V2controllersChartVersion ApiV2ChartVersionsSelectorGet(ctx, selector).Execute()

Get a ChartVersion entry



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
    selector := "selector_example" // string | The ChartVersion to get's selector: name or numeric ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartVersionsApi.ApiV2ChartVersionsSelectorGet(context.Background(), selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartVersionsApi.ApiV2ChartVersionsSelectorGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ChartVersionsSelectorGet`: V2controllersChartVersion
    fmt.Fprintf(os.Stdout, "Response from `ChartVersionsApi.ApiV2ChartVersionsSelectorGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selector** | **string** | The ChartVersion to get&#39;s selector: name or numeric ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ChartVersionsSelectorGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2controllersChartVersion**](V2controllersChartVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SelectorsChartVersionsSelectorGet

> []string ApiV2SelectorsChartVersionsSelectorGet(ctx, selector).Execute()

List ChartVersion selectors



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
    selector := "selector_example" // string | The selector of the ChartVersion to list other selectors for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartVersionsApi.ApiV2SelectorsChartVersionsSelectorGet(context.Background(), selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartVersionsApi.ApiV2SelectorsChartVersionsSelectorGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2SelectorsChartVersionsSelectorGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `ChartVersionsApi.ApiV2SelectorsChartVersionsSelectorGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selector** | **string** | The selector of the ChartVersion to list other selectors for | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SelectorsChartVersionsSelectorGetRequest struct via the builder pattern


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

