# \ClustersApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ClustersGet**](ClustersApi.md#ApiV2ClustersGet) | **Get** /api/v2/clusters | List Cluster entries
[**ApiV2ClustersPost**](ClustersApi.md#ApiV2ClustersPost) | **Post** /api/v2/clusters | Create a new Cluster entry
[**ApiV2ClustersSelectorDelete**](ClustersApi.md#ApiV2ClustersSelectorDelete) | **Delete** /api/v2/clusters/{selector} | Delete a Cluster entry
[**ApiV2ClustersSelectorGet**](ClustersApi.md#ApiV2ClustersSelectorGet) | **Get** /api/v2/clusters/{selector} | Get a Cluster entry
[**ApiV2ClustersSelectorPatch**](ClustersApi.md#ApiV2ClustersSelectorPatch) | **Patch** /api/v2/clusters/{selector} | Edit a Cluster entry
[**ApiV2SelectorsClustersSelectorGet**](ClustersApi.md#ApiV2SelectorsClustersSelectorGet) | **Get** /api/v2/selectors/clusters/{selector} | List Cluster selectors


# **ApiV2ClustersGet**
> []V2controllersCluster ApiV2ClustersGet(ctx, name, optional)
List Cluster entries

List existing Cluster entries, ordered by most recently updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Required when creating | 
 **optional** | ***ClustersApiApiV2ClustersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiApiV2ClustersGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **address** | **optional.String**| Required when creating | 
 **azureSubscription** | **optional.String**| Required when creating if providers is &#39;azure&#39; | 
 **base** | **optional.String**| Required when creating | 
 **createdAt** | **optional.String**|  | 
 **googleProject** | **optional.String**| Required when creating if provider is &#39;google&#39; | 
 **id** | **optional.Int32**|  | 
 **provider** | **optional.String**|  | [default to google]
 **requiresSuitability** | **optional.Bool**|  | [default to false]
 **updatedAt** | **optional.String**|  | 
 **limit** | **optional.Int32**| An optional limit to the number of entries returned | 

### Return type

[**[]V2controllersCluster**](v2controllers.Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ClustersPost**
> V2controllersCluster ApiV2ClustersPost(ctx, cluster)
Create a new Cluster entry

Create a new Cluster entry. Note that some fields are immutable after creation; /edit lists mutable fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cluster** | [**V2controllersCreatableCluster**](V2controllersCreatableCluster.md)| The Cluster to create | 

### Return type

[**V2controllersCluster**](v2controllers.Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ClustersSelectorDelete**
> V2controllersCluster ApiV2ClustersSelectorDelete(ctx, selector)
Delete a Cluster entry

Delete an existing Cluster entry via one of its \"selectors\": name or numeric ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selector** | **string**| The Cluster to delete&#39;s selector: name or numeric ID | 

### Return type

[**V2controllersCluster**](v2controllers.Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ClustersSelectorGet**
> V2controllersCluster ApiV2ClustersSelectorGet(ctx, selector)
Get a Cluster entry

Get an existing Cluster entry via one of its \"selectors\": name or numeric ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selector** | **string**| The Cluster to get&#39;s selector: name or numeric ID | 

### Return type

[**V2controllersCluster**](v2controllers.Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ClustersSelectorPatch**
> V2controllersCluster ApiV2ClustersSelectorPatch(ctx, selector, cluster)
Edit a Cluster entry

Edit an existing Cluster entry via one of its \"selectors\": name or numeric ID. Note that only mutable fields are available here, immutable fields can only be set using /create.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selector** | **string**| The Cluster to edit&#39;s selector: name or numeric ID | 
  **cluster** | [**V2controllersEditableCluster**](V2controllersEditableCluster.md)| The edits to make to the Cluster | 

### Return type

[**V2controllersCluster**](v2controllers.Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2SelectorsClustersSelectorGet**
> []string ApiV2SelectorsClustersSelectorGet(ctx, selector)
List Cluster selectors

Validate a given Cluster selector and provide any other selectors that would match the same Cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selector** | **string**| The selector of the Cluster to list other selectors for | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

