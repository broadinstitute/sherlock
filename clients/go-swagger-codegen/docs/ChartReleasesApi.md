# \ChartReleasesApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ChartReleasesGet**](ChartReleasesApi.md#ApiV2ChartReleasesGet) | **Get** /api/v2/chart-releases | List ChartRelease entries
[**ApiV2ChartReleasesPost**](ChartReleasesApi.md#ApiV2ChartReleasesPost) | **Post** /api/v2/chart-releases | Create a new ChartRelease entry
[**ApiV2ChartReleasesSelectorDelete**](ChartReleasesApi.md#ApiV2ChartReleasesSelectorDelete) | **Delete** /api/v2/chart-releases/{selector} | Delete a ChartRelease entry
[**ApiV2ChartReleasesSelectorGet**](ChartReleasesApi.md#ApiV2ChartReleasesSelectorGet) | **Get** /api/v2/chart-releases/{selector} | Get a ChartRelease entry
[**ApiV2ChartReleasesSelectorPatch**](ChartReleasesApi.md#ApiV2ChartReleasesSelectorPatch) | **Patch** /api/v2/chart-releases/{selector} | Edit a ChartRelease entry
[**ApiV2SelectorsChartReleasesSelectorGet**](ChartReleasesApi.md#ApiV2SelectorsChartReleasesSelectorGet) | **Get** /api/v2/selectors/chart-releases/{selector} | List ChartRelease selectors


# **ApiV2ChartReleasesGet**
> []V2controllersChartRelease ApiV2ChartReleasesGet(ctx, optional)
List ChartRelease entries

List existing ChartRelease entries, ordered by most recently updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChartReleasesApiApiV2ChartReleasesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChartReleasesApiApiV2ChartReleasesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chart** | **optional.String**| Required when creating | 
 **cluster** | **optional.String**| When creating, will default the environment&#39;s default cluster, if provided. Either this or environment must be provided. | 
 **createdAt** | **optional.String**|  | 
 **currentAppVersionExact** | **optional.String**|  | 
 **currentChartVersionExact** | **optional.String**|  | 
 **destinationType** | **optional.String**| Calculated field | 
 **environment** | **optional.String**| Either this or cluster must be provided. | 
 **helmfileRef** | **optional.String**|  | [default to HEAD]
 **id** | **optional.Int32**|  | 
 **name** | **optional.String**| When creating, will be calculated if left empty | 
 **namespace** | **optional.String**| When creating, will default to the environment&#39;s default namespace, if provided | 
 **targetAppVersionBranch** | **optional.String**| When creating, will default to the app&#39;s main branch if it has one recorded | 
 **targetAppVersionCommit** | **optional.String**|  | 
 **targetAppVersionExact** | **optional.String**|  | 
 **targetAppVersionUse** | **optional.String**| When creating, will default to referencing any provided target app version field (exact, then commit, then branch) | 
 **targetChartVersionExact** | **optional.String**|  | 
 **targetChartVersionUse** | **optional.String**| When creating, will default to latest unless an exact target chart version is provided | 
 **thelmaMode** | **optional.String**|  | 
 **updatedAt** | **optional.String**|  | 
 **limit** | **optional.Int32**| An optional limit to the number of entries returned | 

### Return type

[**[]V2controllersChartRelease**](v2controllers.ChartRelease.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ChartReleasesPost**
> V2controllersChartRelease ApiV2ChartReleasesPost(ctx, chartRelease)
Create a new ChartRelease entry

Create a new ChartRelease entry. Note that some fields are immutable after creation; /edit lists mutable fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **chartRelease** | [**V2controllersCreatableChartRelease**](V2controllersCreatableChartRelease.md)| The ChartRelease to create | 

### Return type

[**V2controllersChartRelease**](v2controllers.ChartRelease.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ChartReleasesSelectorDelete**
> V2controllersChartRelease ApiV2ChartReleasesSelectorDelete(ctx, selector)
Delete a ChartRelease entry

Delete an existing ChartRelease entry via one of its \"selectors\": name, numeric ID, environment/chart, or cluster/namespace/chart.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selector** | **string**| The ChartRelease to delete&#39;s selector: name or numeric ID | 

### Return type

[**V2controllersChartRelease**](v2controllers.ChartRelease.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ChartReleasesSelectorGet**
> V2controllersChartRelease ApiV2ChartReleasesSelectorGet(ctx, selector)
Get a ChartRelease entry

Get an existing ChartRelease entry via one of its \"selectors\": name, numeric ID, environment/chart, or cluster/namespace/chart.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selector** | **string**| The ChartRelease to get&#39;s selector: name or numeric ID | 

### Return type

[**V2controllersChartRelease**](v2controllers.ChartRelease.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ChartReleasesSelectorPatch**
> V2controllersChartRelease ApiV2ChartReleasesSelectorPatch(ctx, selector, chartRelease)
Edit a ChartRelease entry

Edit an existing ChartRelease entry via one of its \"selectors\": name, numeric ID, environment/chart, or cluster/namespace/chart. Note that only mutable fields are available here, immutable fields can only be set using /create.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selector** | **string**| The ChartRelease to edit&#39;s selector: name or numeric ID | 
  **chartRelease** | [**V2controllersEditableChartRelease**](V2controllersEditableChartRelease.md)| The edits to make to the ChartRelease | 

### Return type

[**V2controllersChartRelease**](v2controllers.ChartRelease.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2SelectorsChartReleasesSelectorGet**
> []string ApiV2SelectorsChartReleasesSelectorGet(ctx, selector)
List ChartRelease selectors

Validate a given ChartRelease selector and provide any other selectors that would match the same ChartRelease.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selector** | **string**| The selector of the ChartRelease to list other selectors for | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

