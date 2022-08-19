# \ChartsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ChartsGet**](ChartsApi.md#ApiV2ChartsGet) | **Get** /api/v2/charts | List Chart entries
[**ApiV2ChartsPost**](ChartsApi.md#ApiV2ChartsPost) | **Post** /api/v2/charts | Create a new Chart entry
[**ApiV2ChartsSelectorDelete**](ChartsApi.md#ApiV2ChartsSelectorDelete) | **Delete** /api/v2/charts/{selector} | Delete a Chart entry
[**ApiV2ChartsSelectorGet**](ChartsApi.md#ApiV2ChartsSelectorGet) | **Get** /api/v2/charts/{selector} | Get a Chart entry
[**ApiV2ChartsSelectorPatch**](ChartsApi.md#ApiV2ChartsSelectorPatch) | **Patch** /api/v2/charts/{selector} | Edit a Chart entry
[**ApiV2SelectorsChartsSelectorGet**](ChartsApi.md#ApiV2SelectorsChartsSelectorGet) | **Get** /api/v2/selectors/charts/{selector} | List Chart selectors



## ApiV2ChartsGet

> []V2controllersChart ApiV2ChartsGet(ctx).AppImageGitMainBranch(appImageGitMainBranch).AppImageGitRepo(appImageGitRepo).ChartRepo(chartRepo).CreatedAt(createdAt).Id(id).Name(name).UpdatedAt(updatedAt).Limit(limit).Execute()

List Chart entries



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
    appImageGitMainBranch := "appImageGitMainBranch_example" // string |  (optional)
    appImageGitRepo := "appImageGitRepo_example" // string |  (optional)
    chartRepo := "chartRepo_example" // string |  (optional) (default to "terra-helm")
    createdAt := "createdAt_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)
    name := "name_example" // string | Required when creating (optional)
    updatedAt := "updatedAt_example" // string |  (optional)
    limit := int32(56) // int32 | An optional limit to the number of entries returned (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartsApi.ApiV2ChartsGet(context.Background()).AppImageGitMainBranch(appImageGitMainBranch).AppImageGitRepo(appImageGitRepo).ChartRepo(chartRepo).CreatedAt(createdAt).Id(id).Name(name).UpdatedAt(updatedAt).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartsApi.ApiV2ChartsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ChartsGet`: []V2controllersChart
    fmt.Fprintf(os.Stdout, "Response from `ChartsApi.ApiV2ChartsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ChartsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appImageGitMainBranch** | **string** |  | 
 **appImageGitRepo** | **string** |  | 
 **chartRepo** | **string** |  | [default to &quot;terra-helm&quot;]
 **createdAt** | **string** |  | 
 **id** | **int32** |  | 
 **name** | **string** | Required when creating | 
 **updatedAt** | **string** |  | 
 **limit** | **int32** | An optional limit to the number of entries returned | 

### Return type

[**[]V2controllersChart**](V2controllersChart.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ChartsPost

> V2controllersChart ApiV2ChartsPost(ctx).Chart(chart).Execute()

Create a new Chart entry



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
    chart := *openapiclient.NewV2controllersCreatableChart() // V2controllersCreatableChart | The Chart to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartsApi.ApiV2ChartsPost(context.Background()).Chart(chart).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartsApi.ApiV2ChartsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ChartsPost`: V2controllersChart
    fmt.Fprintf(os.Stdout, "Response from `ChartsApi.ApiV2ChartsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ChartsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chart** | [**V2controllersCreatableChart**](V2controllersCreatableChart.md) | The Chart to create | 

### Return type

[**V2controllersChart**](V2controllersChart.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ChartsSelectorDelete

> V2controllersChart ApiV2ChartsSelectorDelete(ctx, selector).Execute()

Delete a Chart entry



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
    selector := "selector_example" // string | The Chart to delete's selector: name or numeric ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartsApi.ApiV2ChartsSelectorDelete(context.Background(), selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartsApi.ApiV2ChartsSelectorDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ChartsSelectorDelete`: V2controllersChart
    fmt.Fprintf(os.Stdout, "Response from `ChartsApi.ApiV2ChartsSelectorDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selector** | **string** | The Chart to delete&#39;s selector: name or numeric ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ChartsSelectorDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2controllersChart**](V2controllersChart.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ChartsSelectorGet

> V2controllersChart ApiV2ChartsSelectorGet(ctx, selector).Execute()

Get a Chart entry



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
    selector := "selector_example" // string | The Chart to get's selector: name or numeric ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartsApi.ApiV2ChartsSelectorGet(context.Background(), selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartsApi.ApiV2ChartsSelectorGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ChartsSelectorGet`: V2controllersChart
    fmt.Fprintf(os.Stdout, "Response from `ChartsApi.ApiV2ChartsSelectorGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selector** | **string** | The Chart to get&#39;s selector: name or numeric ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ChartsSelectorGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2controllersChart**](V2controllersChart.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ChartsSelectorPatch

> V2controllersChart ApiV2ChartsSelectorPatch(ctx, selector).Chart(chart).Execute()

Edit a Chart entry



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
    selector := "selector_example" // string | The Chart to edit's selector: name or numeric ID
    chart := *openapiclient.NewV2controllersEditableChart() // V2controllersEditableChart | The edits to make to the Chart

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartsApi.ApiV2ChartsSelectorPatch(context.Background(), selector).Chart(chart).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartsApi.ApiV2ChartsSelectorPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ChartsSelectorPatch`: V2controllersChart
    fmt.Fprintf(os.Stdout, "Response from `ChartsApi.ApiV2ChartsSelectorPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selector** | **string** | The Chart to edit&#39;s selector: name or numeric ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ChartsSelectorPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chart** | [**V2controllersEditableChart**](V2controllersEditableChart.md) | The edits to make to the Chart | 

### Return type

[**V2controllersChart**](V2controllersChart.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SelectorsChartsSelectorGet

> []string ApiV2SelectorsChartsSelectorGet(ctx, selector).Execute()

List Chart selectors



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
    selector := "selector_example" // string | The selector of the Chart to list other selectors for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartsApi.ApiV2SelectorsChartsSelectorGet(context.Background(), selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartsApi.ApiV2SelectorsChartsSelectorGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2SelectorsChartsSelectorGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `ChartsApi.ApiV2SelectorsChartsSelectorGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selector** | **string** | The selector of the Chart to list other selectors for | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SelectorsChartsSelectorGetRequest struct via the builder pattern


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

