# \AppVersionsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2AppVersionsGet**](AppVersionsApi.md#ApiV2AppVersionsGet) | **Get** /api/v2/app-versions | List AppVersion entries
[**ApiV2AppVersionsPost**](AppVersionsApi.md#ApiV2AppVersionsPost) | **Post** /api/v2/app-versions | Create a new AppVersion entry
[**ApiV2AppVersionsSelectorGet**](AppVersionsApi.md#ApiV2AppVersionsSelectorGet) | **Get** /api/v2/app-versions/{selector} | Get a AppVersion entry
[**ApiV2SelectorsAppVersionsSelectorGet**](AppVersionsApi.md#ApiV2SelectorsAppVersionsSelectorGet) | **Get** /api/v2/selectors/app-versions/{selector} | List AppVersion selectors


# **ApiV2AppVersionsGet**
> []V2controllersAppVersion ApiV2AppVersionsGet(ctx, optional)
List AppVersion entries

List existing AppVersion entries, ordered by most recently updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppVersionsApiApiV2AppVersionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppVersionsApiApiV2AppVersionsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appVersion** | **optional.String**| Required when creating | 
 **chart** | **optional.String**| Required when creating | 
 **createdAt** | **optional.String**|  | 
 **gitBranch** | **optional.String**|  | 
 **gitCommit** | **optional.String**|  | 
 **id** | **optional.Int32**|  | 
 **updatedAt** | **optional.String**|  | 
 **limit** | **optional.Int32**| An optional limit to the number of entries returned | 

### Return type

[**[]V2controllersAppVersion**](v2controllers.AppVersion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2AppVersionsPost**
> V2controllersAppVersion ApiV2AppVersionsPost(ctx, appVersion)
Create a new AppVersion entry

Create a new AppVersion entry. Note that fields are immutable after creation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appVersion** | [**V2controllersCreatableAppVersion**](V2controllersCreatableAppVersion.md)| The AppVersion to create | 

### Return type

[**V2controllersAppVersion**](v2controllers.AppVersion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2AppVersionsSelectorGet**
> V2controllersAppVersion ApiV2AppVersionsSelectorGet(ctx, selector)
Get a AppVersion entry

Get an existing AppVersion entry via one its \"selector\"--its numeric ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selector** | **string**| The AppVersion to get&#39;s selector: name or numeric ID | 

### Return type

[**V2controllersAppVersion**](v2controllers.AppVersion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2SelectorsAppVersionsSelectorGet**
> []string ApiV2SelectorsAppVersionsSelectorGet(ctx, selector)
List AppVersion selectors

Validate a given AppVersion selector and provide any other selectors that would match the same AppVersion.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selector** | **string**| The selector of the AppVersion to list other selectors for | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

