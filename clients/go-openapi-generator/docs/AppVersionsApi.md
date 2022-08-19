# \AppVersionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2AppVersionsGet**](AppVersionsApi.md#ApiV2AppVersionsGet) | **Get** /api/v2/app-versions | List AppVersion entries
[**ApiV2AppVersionsPost**](AppVersionsApi.md#ApiV2AppVersionsPost) | **Post** /api/v2/app-versions | Create a new AppVersion entry
[**ApiV2AppVersionsSelectorGet**](AppVersionsApi.md#ApiV2AppVersionsSelectorGet) | **Get** /api/v2/app-versions/{selector} | Get a AppVersion entry
[**ApiV2SelectorsAppVersionsSelectorGet**](AppVersionsApi.md#ApiV2SelectorsAppVersionsSelectorGet) | **Get** /api/v2/selectors/app-versions/{selector} | List AppVersion selectors



## ApiV2AppVersionsGet

> []V2controllersAppVersion ApiV2AppVersionsGet(ctx).AppVersion(appVersion).Chart(chart).CreatedAt(createdAt).GitBranch(gitBranch).GitCommit(gitCommit).Id(id).UpdatedAt(updatedAt).Limit(limit).Execute()

List AppVersion entries



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
    appVersion := "appVersion_example" // string | Required when creating (optional)
    chart := "chart_example" // string | Required when creating (optional)
    createdAt := "createdAt_example" // string |  (optional)
    gitBranch := "gitBranch_example" // string |  (optional)
    gitCommit := "gitCommit_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)
    updatedAt := "updatedAt_example" // string |  (optional)
    limit := int32(56) // int32 | An optional limit to the number of entries returned (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVersionsApi.ApiV2AppVersionsGet(context.Background()).AppVersion(appVersion).Chart(chart).CreatedAt(createdAt).GitBranch(gitBranch).GitCommit(gitCommit).Id(id).UpdatedAt(updatedAt).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVersionsApi.ApiV2AppVersionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2AppVersionsGet`: []V2controllersAppVersion
    fmt.Fprintf(os.Stdout, "Response from `AppVersionsApi.ApiV2AppVersionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AppVersionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appVersion** | **string** | Required when creating | 
 **chart** | **string** | Required when creating | 
 **createdAt** | **string** |  | 
 **gitBranch** | **string** |  | 
 **gitCommit** | **string** |  | 
 **id** | **int32** |  | 
 **updatedAt** | **string** |  | 
 **limit** | **int32** | An optional limit to the number of entries returned | 

### Return type

[**[]V2controllersAppVersion**](V2controllersAppVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2AppVersionsPost

> V2controllersAppVersion ApiV2AppVersionsPost(ctx).AppVersion(appVersion).Execute()

Create a new AppVersion entry



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
    appVersion := *openapiclient.NewV2controllersCreatableAppVersion() // V2controllersCreatableAppVersion | The AppVersion to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVersionsApi.ApiV2AppVersionsPost(context.Background()).AppVersion(appVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVersionsApi.ApiV2AppVersionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2AppVersionsPost`: V2controllersAppVersion
    fmt.Fprintf(os.Stdout, "Response from `AppVersionsApi.ApiV2AppVersionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AppVersionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appVersion** | [**V2controllersCreatableAppVersion**](V2controllersCreatableAppVersion.md) | The AppVersion to create | 

### Return type

[**V2controllersAppVersion**](V2controllersAppVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2AppVersionsSelectorGet

> V2controllersAppVersion ApiV2AppVersionsSelectorGet(ctx, selector).Execute()

Get a AppVersion entry



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
    selector := "selector_example" // string | The AppVersion to get's selector: name or numeric ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVersionsApi.ApiV2AppVersionsSelectorGet(context.Background(), selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVersionsApi.ApiV2AppVersionsSelectorGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2AppVersionsSelectorGet`: V2controllersAppVersion
    fmt.Fprintf(os.Stdout, "Response from `AppVersionsApi.ApiV2AppVersionsSelectorGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selector** | **string** | The AppVersion to get&#39;s selector: name or numeric ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AppVersionsSelectorGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2controllersAppVersion**](V2controllersAppVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SelectorsAppVersionsSelectorGet

> []string ApiV2SelectorsAppVersionsSelectorGet(ctx, selector).Execute()

List AppVersion selectors



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
    selector := "selector_example" // string | The selector of the AppVersion to list other selectors for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVersionsApi.ApiV2SelectorsAppVersionsSelectorGet(context.Background(), selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVersionsApi.ApiV2SelectorsAppVersionsSelectorGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2SelectorsAppVersionsSelectorGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `AppVersionsApi.ApiV2SelectorsAppVersionsSelectorGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selector** | **string** | The selector of the AppVersion to list other selectors for | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SelectorsAppVersionsSelectorGetRequest struct via the builder pattern


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

