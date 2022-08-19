# \ChartsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ChartsGet**](ChartsApi.md#ApiV2ChartsGet) | **Get** /api/v2/charts | List Chart entries
[**ApiV2ChartsPost**](ChartsApi.md#ApiV2ChartsPost) | **Post** /api/v2/charts | Create a new Chart entry
[**ApiV2ChartsSelectorDelete**](ChartsApi.md#ApiV2ChartsSelectorDelete) | **Delete** /api/v2/charts/{selector} | Delete a Chart entry
[**ApiV2ChartsSelectorGet**](ChartsApi.md#ApiV2ChartsSelectorGet) | **Get** /api/v2/charts/{selector} | Get a Chart entry
[**ApiV2ChartsSelectorPatch**](ChartsApi.md#ApiV2ChartsSelectorPatch) | **Patch** /api/v2/charts/{selector} | Edit a Chart entry
[**ApiV2SelectorsChartsSelectorGet**](ChartsApi.md#ApiV2SelectorsChartsSelectorGet) | **Get** /api/v2/selectors/charts/{selector} | List Chart selectors


# **ApiV2ChartsGet**
> []V2controllersChart ApiV2ChartsGet(ctx, optional)
List Chart entries

List existing Chart entries, ordered by most recently updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChartsApiApiV2ChartsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChartsApiApiV2ChartsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appImageGitMainBranch** | **optional.String**|  | 
 **appImageGitRepo** | **optional.String**|  | 
 **chartRepo** | **optional.String**|  | [default to terra-helm]
 **createdAt** | **optional.String**|  | 
 **id** | **optional.Int32**|  | 
 **name** | **optional.String**| Required when creating | 
 **updatedAt** | **optional.String**|  | 
 **limit** | **optional.Int32**| An optional limit to the number of entries returned | 

### Return type

[**[]V2controllersChart**](v2controllers.Chart.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ChartsPost**
> V2controllersChart ApiV2ChartsPost(ctx, chart)
Create a new Chart entry

Create a new Chart entry. Note that some fields are immutable after creation; /edit lists mutable fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **chart** | [**V2controllersCreatableChart**](V2controllersCreatableChart.md)| The Chart to create | 

### Return type

[**V2controllersChart**](v2controllers.Chart.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ChartsSelectorDelete**
> V2controllersChart ApiV2ChartsSelectorDelete(ctx, selector)
Delete a Chart entry

Delete an existing Chart entry via one of its \"selectors\": name or numeric ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selector** | **string**| The Chart to delete&#39;s selector: name or numeric ID | 

### Return type

[**V2controllersChart**](v2controllers.Chart.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ChartsSelectorGet**
> V2controllersChart ApiV2ChartsSelectorGet(ctx, selector)
Get a Chart entry

Get an existing Chart entry via one of its \"selectors\": name or numeric ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selector** | **string**| The Chart to get&#39;s selector: name or numeric ID | 

### Return type

[**V2controllersChart**](v2controllers.Chart.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ChartsSelectorPatch**
> V2controllersChart ApiV2ChartsSelectorPatch(ctx, selector, chart)
Edit a Chart entry

Edit an existing Chart entry via one of its \"selectors\": name or numeric ID. Note that only mutable fields are available here, immutable fields can only be set using /create.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selector** | **string**| The Chart to edit&#39;s selector: name or numeric ID | 
  **chart** | [**V2controllersEditableChart**](V2controllersEditableChart.md)| The edits to make to the Chart | 

### Return type

[**V2controllersChart**](v2controllers.Chart.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2SelectorsChartsSelectorGet**
> []string ApiV2SelectorsChartsSelectorGet(ctx, selector)
List Chart selectors

Validate a given Chart selector and provide any other selectors that would match the same Chart.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selector** | **string**| The selector of the Chart to list other selectors for | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

