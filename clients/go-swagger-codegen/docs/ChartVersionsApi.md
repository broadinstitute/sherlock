# \ChartVersionsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ChartVersionsGet**](ChartVersionsApi.md#ApiV2ChartVersionsGet) | **Get** /api/v2/chart-versions | List ChartVersion entries
[**ApiV2ChartVersionsPost**](ChartVersionsApi.md#ApiV2ChartVersionsPost) | **Post** /api/v2/chart-versions | Create a new ChartVersion entry
[**ApiV2ChartVersionsSelectorGet**](ChartVersionsApi.md#ApiV2ChartVersionsSelectorGet) | **Get** /api/v2/chart-versions/{selector} | Get a ChartVersion entry
[**ApiV2SelectorsChartVersionsSelectorGet**](ChartVersionsApi.md#ApiV2SelectorsChartVersionsSelectorGet) | **Get** /api/v2/selectors/chart-versions/{selector} | List ChartVersion selectors


# **ApiV2ChartVersionsGet**
> []V2controllersChartVersion ApiV2ChartVersionsGet(ctx, optional)
List ChartVersion entries

List existing ChartVersion entries, ordered by most recently updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChartVersionsApiApiV2ChartVersionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChartVersionsApiApiV2ChartVersionsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chart** | **optional.String**| Required when creating | 
 **chartVersion** | **optional.String**| Required when creating | 
 **createdAt** | **optional.String**|  | 
 **id** | **optional.Int32**|  | 
 **updatedAt** | **optional.String**|  | 
 **limit** | **optional.Int32**| An optional limit to the number of entries returned | 

### Return type

[**[]V2controllersChartVersion**](v2controllers.ChartVersion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ChartVersionsPost**
> V2controllersChartVersion ApiV2ChartVersionsPost(ctx, chartVersion)
Create a new ChartVersion entry

Create a new ChartVersion entry. Note that fields are immutable after creation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **chartVersion** | [**V2controllersCreatableChartVersion**](V2controllersCreatableChartVersion.md)| The ChartVersion to create | 

### Return type

[**V2controllersChartVersion**](v2controllers.ChartVersion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ChartVersionsSelectorGet**
> V2controllersChartVersion ApiV2ChartVersionsSelectorGet(ctx, selector)
Get a ChartVersion entry

Get an existing ChartVersion entry via one its \"selector\"--its numeric ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selector** | **string**| The ChartVersion to get&#39;s selector: name or numeric ID | 

### Return type

[**V2controllersChartVersion**](v2controllers.ChartVersion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2SelectorsChartVersionsSelectorGet**
> []string ApiV2SelectorsChartVersionsSelectorGet(ctx, selector)
List ChartVersion selectors

Validate a given ChartVersion selector and provide any other selectors that would match the same ChartVersion.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selector** | **string**| The selector of the ChartVersion to list other selectors for | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

