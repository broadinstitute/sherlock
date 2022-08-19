# \ChartDeployRecordsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ChartDeployRecordsGet**](ChartDeployRecordsApi.md#ApiV2ChartDeployRecordsGet) | **Get** /api/v2/chart-deploy-records | List ChartDeployRecord entries
[**ApiV2ChartDeployRecordsPost**](ChartDeployRecordsApi.md#ApiV2ChartDeployRecordsPost) | **Post** /api/v2/chart-deploy-records | Create a new ChartDeployRecord entry
[**ApiV2ChartDeployRecordsSelectorGet**](ChartDeployRecordsApi.md#ApiV2ChartDeployRecordsSelectorGet) | **Get** /api/v2/chart-deploy-records/{selector} | Get a ChartDeployRecord entry
[**ApiV2SelectorsChartDeployRecordsSelectorGet**](ChartDeployRecordsApi.md#ApiV2SelectorsChartDeployRecordsSelectorGet) | **Get** /api/v2/selectors/chart-deploy-records/{selector} | List ChartDeployRecord selectors



## ApiV2ChartDeployRecordsGet

> []V2controllersChartDeployRecord ApiV2ChartDeployRecordsGet(ctx).ChartRelease(chartRelease).CreatedAt(createdAt).ExactAppVersion(exactAppVersion).ExactChartVersion(exactChartVersion).HelmfileRef(helmfileRef).Id(id).UpdatedAt(updatedAt).Limit(limit).Execute()

List ChartDeployRecord entries



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
    chartRelease := "chartRelease_example" // string | Required when creating (optional)
    createdAt := "createdAt_example" // string |  (optional)
    exactAppVersion := "exactAppVersion_example" // string | When creating, will default to the value currently held by the chart release (optional)
    exactChartVersion := "exactChartVersion_example" // string | When creating, will default to the value currently held by the chart release (optional)
    helmfileRef := "helmfileRef_example" // string | When creating, will default to the value currently held by the chart release (optional)
    id := int32(56) // int32 |  (optional)
    updatedAt := "updatedAt_example" // string |  (optional)
    limit := int32(56) // int32 | An optional limit to the number of entries returned (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartDeployRecordsApi.ApiV2ChartDeployRecordsGet(context.Background()).ChartRelease(chartRelease).CreatedAt(createdAt).ExactAppVersion(exactAppVersion).ExactChartVersion(exactChartVersion).HelmfileRef(helmfileRef).Id(id).UpdatedAt(updatedAt).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartDeployRecordsApi.ApiV2ChartDeployRecordsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ChartDeployRecordsGet`: []V2controllersChartDeployRecord
    fmt.Fprintf(os.Stdout, "Response from `ChartDeployRecordsApi.ApiV2ChartDeployRecordsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ChartDeployRecordsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chartRelease** | **string** | Required when creating | 
 **createdAt** | **string** |  | 
 **exactAppVersion** | **string** | When creating, will default to the value currently held by the chart release | 
 **exactChartVersion** | **string** | When creating, will default to the value currently held by the chart release | 
 **helmfileRef** | **string** | When creating, will default to the value currently held by the chart release | 
 **id** | **int32** |  | 
 **updatedAt** | **string** |  | 
 **limit** | **int32** | An optional limit to the number of entries returned | 

### Return type

[**[]V2controllersChartDeployRecord**](V2controllersChartDeployRecord.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ChartDeployRecordsPost

> V2controllersChartDeployRecord ApiV2ChartDeployRecordsPost(ctx).ChartDeployRecord(chartDeployRecord).Execute()

Create a new ChartDeployRecord entry



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
    chartDeployRecord := *openapiclient.NewV2controllersCreatableChartDeployRecord() // V2controllersCreatableChartDeployRecord | The ChartDeployRecord to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartDeployRecordsApi.ApiV2ChartDeployRecordsPost(context.Background()).ChartDeployRecord(chartDeployRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartDeployRecordsApi.ApiV2ChartDeployRecordsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ChartDeployRecordsPost`: V2controllersChartDeployRecord
    fmt.Fprintf(os.Stdout, "Response from `ChartDeployRecordsApi.ApiV2ChartDeployRecordsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ChartDeployRecordsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chartDeployRecord** | [**V2controllersCreatableChartDeployRecord**](V2controllersCreatableChartDeployRecord.md) | The ChartDeployRecord to create | 

### Return type

[**V2controllersChartDeployRecord**](V2controllersChartDeployRecord.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ChartDeployRecordsSelectorGet

> V2controllersChartDeployRecord ApiV2ChartDeployRecordsSelectorGet(ctx, selector).Execute()

Get a ChartDeployRecord entry



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
    selector := "selector_example" // string | The ChartDeployRecord to get's selector: name or numeric ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartDeployRecordsApi.ApiV2ChartDeployRecordsSelectorGet(context.Background(), selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartDeployRecordsApi.ApiV2ChartDeployRecordsSelectorGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ChartDeployRecordsSelectorGet`: V2controllersChartDeployRecord
    fmt.Fprintf(os.Stdout, "Response from `ChartDeployRecordsApi.ApiV2ChartDeployRecordsSelectorGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selector** | **string** | The ChartDeployRecord to get&#39;s selector: name or numeric ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ChartDeployRecordsSelectorGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2controllersChartDeployRecord**](V2controllersChartDeployRecord.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SelectorsChartDeployRecordsSelectorGet

> []string ApiV2SelectorsChartDeployRecordsSelectorGet(ctx, selector).Execute()

List ChartDeployRecord selectors



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
    selector := "selector_example" // string | The selector of the ChartDeployRecord to list other selectors for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartDeployRecordsApi.ApiV2SelectorsChartDeployRecordsSelectorGet(context.Background(), selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartDeployRecordsApi.ApiV2SelectorsChartDeployRecordsSelectorGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2SelectorsChartDeployRecordsSelectorGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `ChartDeployRecordsApi.ApiV2SelectorsChartDeployRecordsSelectorGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selector** | **string** | The selector of the ChartDeployRecord to list other selectors for | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SelectorsChartDeployRecordsSelectorGetRequest struct via the builder pattern


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

