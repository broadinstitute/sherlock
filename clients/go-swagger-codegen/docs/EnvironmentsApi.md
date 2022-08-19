# \EnvironmentsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2EnvironmentsGet**](EnvironmentsApi.md#ApiV2EnvironmentsGet) | **Get** /api/v2/environments | List Environment entries
[**ApiV2EnvironmentsPost**](EnvironmentsApi.md#ApiV2EnvironmentsPost) | **Post** /api/v2/environments | Create a new Environment entry
[**ApiV2EnvironmentsSelectorDelete**](EnvironmentsApi.md#ApiV2EnvironmentsSelectorDelete) | **Delete** /api/v2/environments/{selector} | Delete a Environment entry
[**ApiV2EnvironmentsSelectorGet**](EnvironmentsApi.md#ApiV2EnvironmentsSelectorGet) | **Get** /api/v2/environments/{selector} | Get a Environment entry
[**ApiV2EnvironmentsSelectorPatch**](EnvironmentsApi.md#ApiV2EnvironmentsSelectorPatch) | **Patch** /api/v2/environments/{selector} | Edit a Environment entry
[**ApiV2SelectorsEnvironmentsSelectorGet**](EnvironmentsApi.md#ApiV2SelectorsEnvironmentsSelectorGet) | **Get** /api/v2/selectors/environments/{selector} | List Environment selectors


# **ApiV2EnvironmentsGet**
> []V2controllersEnvironment ApiV2EnvironmentsGet(ctx, optional)
List Environment entries

List existing Environment entries, ordered by most recently updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EnvironmentsApiApiV2EnvironmentsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnvironmentsApiApiV2EnvironmentsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **base** | **optional.String**| Required when creating | 
 **chartReleasesFromTemplate** | **optional.Bool**| Upon creation of a dynamic environment, if this is true the template&#39;s chart releases will be copied to the new environment | [default to true]
 **createdAt** | **optional.String**|  | 
 **defaultCluster** | **optional.String**|  | 
 **defaultNamespace** | **optional.String**|  | 
 **id** | **optional.Int32**|  | 
 **lifecycle** | **optional.String**|  | [default to dynamic]
 **name** | **optional.String**| When creating, will be calculated if dynamic, required otherwise | 
 **owner** | **optional.String**| When creating, will be set to your email | 
 **requiresSuitability** | **optional.Bool**|  | [default to false]
 **templateEnvironment** | **optional.String**| Required for dynamic environments | 
 **updatedAt** | **optional.String**|  | 
 **valuesName** | **optional.String**|  | 
 **limit** | **optional.Int32**| An optional limit to the number of entries returned | 

### Return type

[**[]V2controllersEnvironment**](v2controllers.Environment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2EnvironmentsPost**
> V2controllersEnvironment ApiV2EnvironmentsPost(ctx, environment)
Create a new Environment entry

Create a new Environment entry. Note that some fields are immutable after creation; /edit lists mutable fields. Creating a dynamic environment based on a template will also copy ChartReleases from the template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environment** | [**V2controllersCreatableEnvironment**](V2controllersCreatableEnvironment.md)| The Environment to create | 

### Return type

[**V2controllersEnvironment**](v2controllers.Environment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2EnvironmentsSelectorDelete**
> V2controllersEnvironment ApiV2EnvironmentsSelectorDelete(ctx, selector)
Delete a Environment entry

Delete an existing Environment entry via one of its \"selectors\": name or numeric ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selector** | **string**| The Environment to delete&#39;s selector: name or numeric ID | 

### Return type

[**V2controllersEnvironment**](v2controllers.Environment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2EnvironmentsSelectorGet**
> V2controllersEnvironment ApiV2EnvironmentsSelectorGet(ctx, selector)
Get a Environment entry

Get an existing Environment entry via one of its \"selectors\": name or numeric ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selector** | **string**| The Environment to get&#39;s selector: name or numeric ID | 

### Return type

[**V2controllersEnvironment**](v2controllers.Environment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2EnvironmentsSelectorPatch**
> V2controllersEnvironment ApiV2EnvironmentsSelectorPatch(ctx, selector, environment)
Edit a Environment entry

Edit an existing Environment entry via one of its \"selectors\": name or numeric ID. Note that only mutable fields are available here, immutable fields can only be set using /create.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selector** | **string**| The Environment to edit&#39;s selector: name or numeric ID | 
  **environment** | [**V2controllersEditableEnvironment**](V2controllersEditableEnvironment.md)| The edits to make to the Environment | 

### Return type

[**V2controllersEnvironment**](v2controllers.Environment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2SelectorsEnvironmentsSelectorGet**
> []string ApiV2SelectorsEnvironmentsSelectorGet(ctx, selector)
List Environment selectors

Validate a given Environment selector and provide any other selectors that would match the same Environment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selector** | **string**| The selector of the Environment to list other selectors for | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

