# \EnvironmentsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2EnvironmentsGet**](EnvironmentsApi.md#ApiV2EnvironmentsGet) | **Get** /api/v2/environments | List Environment entries
[**ApiV2EnvironmentsPost**](EnvironmentsApi.md#ApiV2EnvironmentsPost) | **Post** /api/v2/environments | Create a new Environment entry
[**ApiV2EnvironmentsSelectorDelete**](EnvironmentsApi.md#ApiV2EnvironmentsSelectorDelete) | **Delete** /api/v2/environments/{selector} | Delete a Environment entry
[**ApiV2EnvironmentsSelectorGet**](EnvironmentsApi.md#ApiV2EnvironmentsSelectorGet) | **Get** /api/v2/environments/{selector} | Get a Environment entry
[**ApiV2EnvironmentsSelectorPatch**](EnvironmentsApi.md#ApiV2EnvironmentsSelectorPatch) | **Patch** /api/v2/environments/{selector} | Edit a Environment entry
[**ApiV2SelectorsEnvironmentsSelectorGet**](EnvironmentsApi.md#ApiV2SelectorsEnvironmentsSelectorGet) | **Get** /api/v2/selectors/environments/{selector} | List Environment selectors



## ApiV2EnvironmentsGet

> []V2controllersEnvironment ApiV2EnvironmentsGet(ctx).Base(base).ChartReleasesFromTemplate(chartReleasesFromTemplate).CreatedAt(createdAt).DefaultCluster(defaultCluster).DefaultNamespace(defaultNamespace).Id(id).Lifecycle(lifecycle).Name(name).Owner(owner).RequiresSuitability(requiresSuitability).TemplateEnvironment(templateEnvironment).UpdatedAt(updatedAt).ValuesName(valuesName).Limit(limit).Execute()

List Environment entries



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
    base := "base_example" // string | Required when creating (optional)
    chartReleasesFromTemplate := true // bool | Upon creation of a dynamic environment, if this is true the template's chart releases will be copied to the new environment (optional) (default to true)
    createdAt := "createdAt_example" // string |  (optional)
    defaultCluster := "defaultCluster_example" // string |  (optional)
    defaultNamespace := "defaultNamespace_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)
    lifecycle := "lifecycle_example" // string |  (optional) (default to "dynamic")
    name := "name_example" // string | When creating, will be calculated if dynamic, required otherwise (optional)
    owner := "owner_example" // string | When creating, will be set to your email (optional)
    requiresSuitability := true // bool |  (optional) (default to false)
    templateEnvironment := "templateEnvironment_example" // string | Required for dynamic environments (optional)
    updatedAt := "updatedAt_example" // string |  (optional)
    valuesName := "valuesName_example" // string |  (optional)
    limit := int32(56) // int32 | An optional limit to the number of entries returned (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.ApiV2EnvironmentsGet(context.Background()).Base(base).ChartReleasesFromTemplate(chartReleasesFromTemplate).CreatedAt(createdAt).DefaultCluster(defaultCluster).DefaultNamespace(defaultNamespace).Id(id).Lifecycle(lifecycle).Name(name).Owner(owner).RequiresSuitability(requiresSuitability).TemplateEnvironment(templateEnvironment).UpdatedAt(updatedAt).ValuesName(valuesName).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.ApiV2EnvironmentsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2EnvironmentsGet`: []V2controllersEnvironment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.ApiV2EnvironmentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2EnvironmentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **base** | **string** | Required when creating | 
 **chartReleasesFromTemplate** | **bool** | Upon creation of a dynamic environment, if this is true the template&#39;s chart releases will be copied to the new environment | [default to true]
 **createdAt** | **string** |  | 
 **defaultCluster** | **string** |  | 
 **defaultNamespace** | **string** |  | 
 **id** | **int32** |  | 
 **lifecycle** | **string** |  | [default to &quot;dynamic&quot;]
 **name** | **string** | When creating, will be calculated if dynamic, required otherwise | 
 **owner** | **string** | When creating, will be set to your email | 
 **requiresSuitability** | **bool** |  | [default to false]
 **templateEnvironment** | **string** | Required for dynamic environments | 
 **updatedAt** | **string** |  | 
 **valuesName** | **string** |  | 
 **limit** | **int32** | An optional limit to the number of entries returned | 

### Return type

[**[]V2controllersEnvironment**](V2controllersEnvironment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2EnvironmentsPost

> V2controllersEnvironment ApiV2EnvironmentsPost(ctx).Environment(environment).Execute()

Create a new Environment entry



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
    environment := *openapiclient.NewV2controllersCreatableEnvironment() // V2controllersCreatableEnvironment | The Environment to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.ApiV2EnvironmentsPost(context.Background()).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.ApiV2EnvironmentsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2EnvironmentsPost`: V2controllersEnvironment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.ApiV2EnvironmentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2EnvironmentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | [**V2controllersCreatableEnvironment**](V2controllersCreatableEnvironment.md) | The Environment to create | 

### Return type

[**V2controllersEnvironment**](V2controllersEnvironment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2EnvironmentsSelectorDelete

> V2controllersEnvironment ApiV2EnvironmentsSelectorDelete(ctx, selector).Execute()

Delete a Environment entry



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
    selector := "selector_example" // string | The Environment to delete's selector: name or numeric ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.ApiV2EnvironmentsSelectorDelete(context.Background(), selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.ApiV2EnvironmentsSelectorDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2EnvironmentsSelectorDelete`: V2controllersEnvironment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.ApiV2EnvironmentsSelectorDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selector** | **string** | The Environment to delete&#39;s selector: name or numeric ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2EnvironmentsSelectorDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2controllersEnvironment**](V2controllersEnvironment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2EnvironmentsSelectorGet

> V2controllersEnvironment ApiV2EnvironmentsSelectorGet(ctx, selector).Execute()

Get a Environment entry



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
    selector := "selector_example" // string | The Environment to get's selector: name or numeric ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.ApiV2EnvironmentsSelectorGet(context.Background(), selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.ApiV2EnvironmentsSelectorGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2EnvironmentsSelectorGet`: V2controllersEnvironment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.ApiV2EnvironmentsSelectorGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selector** | **string** | The Environment to get&#39;s selector: name or numeric ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2EnvironmentsSelectorGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2controllersEnvironment**](V2controllersEnvironment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2EnvironmentsSelectorPatch

> V2controllersEnvironment ApiV2EnvironmentsSelectorPatch(ctx, selector).Environment(environment).Execute()

Edit a Environment entry



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
    selector := "selector_example" // string | The Environment to edit's selector: name or numeric ID
    environment := *openapiclient.NewV2controllersEditableEnvironment() // V2controllersEditableEnvironment | The edits to make to the Environment

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.ApiV2EnvironmentsSelectorPatch(context.Background(), selector).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.ApiV2EnvironmentsSelectorPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2EnvironmentsSelectorPatch`: V2controllersEnvironment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.ApiV2EnvironmentsSelectorPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selector** | **string** | The Environment to edit&#39;s selector: name or numeric ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2EnvironmentsSelectorPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environment** | [**V2controllersEditableEnvironment**](V2controllersEditableEnvironment.md) | The edits to make to the Environment | 

### Return type

[**V2controllersEnvironment**](V2controllersEnvironment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SelectorsEnvironmentsSelectorGet

> []string ApiV2SelectorsEnvironmentsSelectorGet(ctx, selector).Execute()

List Environment selectors



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
    selector := "selector_example" // string | The selector of the Environment to list other selectors for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.ApiV2SelectorsEnvironmentsSelectorGet(context.Background(), selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.ApiV2SelectorsEnvironmentsSelectorGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2SelectorsEnvironmentsSelectorGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.ApiV2SelectorsEnvironmentsSelectorGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selector** | **string** | The selector of the Environment to list other selectors for | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SelectorsEnvironmentsSelectorGetRequest struct via the builder pattern


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

