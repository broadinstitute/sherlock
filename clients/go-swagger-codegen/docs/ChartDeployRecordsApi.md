# \ChartDeployRecordsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ChartDeployRecordsGet**](ChartDeployRecordsApi.md#ApiV2ChartDeployRecordsGet) | **Get** /api/v2/chart-deploy-records | List ChartDeployRecord entries
[**ApiV2ChartDeployRecordsPost**](ChartDeployRecordsApi.md#ApiV2ChartDeployRecordsPost) | **Post** /api/v2/chart-deploy-records | Create a new ChartDeployRecord entry
[**ApiV2ChartDeployRecordsSelectorGet**](ChartDeployRecordsApi.md#ApiV2ChartDeployRecordsSelectorGet) | **Get** /api/v2/chart-deploy-records/{selector} | Get a ChartDeployRecord entry
[**ApiV2SelectorsChartDeployRecordsSelectorGet**](ChartDeployRecordsApi.md#ApiV2SelectorsChartDeployRecordsSelectorGet) | **Get** /api/v2/selectors/chart-deploy-records/{selector} | List ChartDeployRecord selectors


# **ApiV2ChartDeployRecordsGet**
> []V2controllersChartDeployRecord ApiV2ChartDeployRecordsGet(ctx, optional)
List ChartDeployRecord entries

List existing ChartDeployRecord entries, ordered by most recently updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChartDeployRecordsApiApiV2ChartDeployRecordsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChartDeployRecordsApiApiV2ChartDeployRecordsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chartRelease** | **optional.String**| Required when creating | 
 **createdAt** | **optional.String**|  | 
 **exactAppVersion** | **optional.String**| When creating, will default to the value currently held by the chart release | 
 **exactChartVersion** | **optional.String**| When creating, will default to the value currently held by the chart release | 
 **helmfileRef** | **optional.String**| When creating, will default to the value currently held by the chart release | 
 **id** | **optional.Int32**|  | 
 **updatedAt** | **optional.String**|  | 
 **limit** | **optional.Int32**| An optional limit to the number of entries returned | 

### Return type

[**[]V2controllersChartDeployRecord**](v2controllers.ChartDeployRecord.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ChartDeployRecordsPost**
> V2controllersChartDeployRecord ApiV2ChartDeployRecordsPost(ctx, chartDeployRecord)
Create a new ChartDeployRecord entry

Create a new ChartDeployRecord entry. Note that fields are immutable after creation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **chartDeployRecord** | [**V2controllersCreatableChartDeployRecord**](V2controllersCreatableChartDeployRecord.md)| The ChartDeployRecord to create | 

### Return type

[**V2controllersChartDeployRecord**](v2controllers.ChartDeployRecord.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ChartDeployRecordsSelectorGet**
> V2controllersChartDeployRecord ApiV2ChartDeployRecordsSelectorGet(ctx, selector)
Get a ChartDeployRecord entry

Get an existing ChartDeployRecord entry via one its \"selector\"--its numeric ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selector** | **string**| The ChartDeployRecord to get&#39;s selector: name or numeric ID | 

### Return type

[**V2controllersChartDeployRecord**](v2controllers.ChartDeployRecord.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2SelectorsChartDeployRecordsSelectorGet**
> []string ApiV2SelectorsChartDeployRecordsSelectorGet(ctx, selector)
List ChartDeployRecord selectors

Validate a given ChartDeployRecord selector and provide any other selectors that would match the same ChartDeployRecord.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selector** | **string**| The selector of the ChartDeployRecord to list other selectors for | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

