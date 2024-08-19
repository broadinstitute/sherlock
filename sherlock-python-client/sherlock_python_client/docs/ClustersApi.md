# sherlock_python_client.ClustersApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_clusters_v3_get**](ClustersApi.md#api_clusters_v3_get) | **GET** /api/clusters/v3 | List Clusters matching a filter
[**api_clusters_v3_post**](ClustersApi.md#api_clusters_v3_post) | **POST** /api/clusters/v3 | Create a Cluster
[**api_clusters_v3_selector_delete**](ClustersApi.md#api_clusters_v3_selector_delete) | **DELETE** /api/clusters/v3/{selector} | Delete an individual Cluster
[**api_clusters_v3_selector_get**](ClustersApi.md#api_clusters_v3_selector_get) | **GET** /api/clusters/v3/{selector} | Get an individual Cluster
[**api_clusters_v3_selector_patch**](ClustersApi.md#api_clusters_v3_selector_patch) | **PATCH** /api/clusters/v3/{selector} | Edit an individual Cluster


# **api_clusters_v3_get**
> List[SherlockClusterV3] api_clusters_v3_get(address=address, azure_subscription=azure_subscription, base=base, created_at=created_at, google_project=google_project, helmfile_ref=helmfile_ref, id=id, location=location, name=name, provider=provider, required_role=required_role, requires_suitability=requires_suitability, updated_at=updated_at, limit=limit, offset=offset)

List Clusters matching a filter

List Clusters matching a filter.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_cluster_v3 import SherlockClusterV3
from sherlock_python_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://sherlock.dsp-devops-prod.broadinstitute.org
# See configuration.py for a list of all supported configuration parameters.
configuration = sherlock_python_client.Configuration(
    host = "https://sherlock.dsp-devops-prod.broadinstitute.org"
)


# Enter a context with an instance of the API client
with sherlock_python_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = sherlock_python_client.ClustersApi(api_client)
    address = 'address_example' # str | Required when creating (optional)
    azure_subscription = 'azure_subscription_example' # str | Required when creating if provider is 'azure' (optional)
    base = 'base_example' # str | Required when creating (optional)
    created_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    google_project = 'google_project_example' # str | Required when creating if provider is 'google' (optional)
    helmfile_ref = 'HEAD' # str |  (optional) (default to 'HEAD')
    id = 56 # int |  (optional)
    location = 'us-central1-a' # str |  (optional) (default to 'us-central1-a')
    name = 'name_example' # str | Required when creating (optional)
    provider = google # str |  (optional) (default to google)
    required_role = 'required_role_example' # str | If present, requires membership in the given role for mutations. Set to an empty string to clear. (optional)
    requires_suitability = True # bool |  (optional)
    updated_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    limit = 56 # int | Control how many Clusters are returned (default 0, meaning all) (optional)
    offset = 56 # int | Control the offset for the returned Clusters (default 0) (optional)

    try:
        # List Clusters matching a filter
        api_response = api_instance.api_clusters_v3_get(address=address, azure_subscription=azure_subscription, base=base, created_at=created_at, google_project=google_project, helmfile_ref=helmfile_ref, id=id, location=location, name=name, provider=provider, required_role=required_role, requires_suitability=requires_suitability, updated_at=updated_at, limit=limit, offset=offset)
        print("The response of ClustersApi->api_clusters_v3_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ClustersApi->api_clusters_v3_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address** | **str**| Required when creating | [optional] 
 **azure_subscription** | **str**| Required when creating if provider is &#39;azure&#39; | [optional] 
 **base** | **str**| Required when creating | [optional] 
 **created_at** | **datetime**|  | [optional] 
 **google_project** | **str**| Required when creating if provider is &#39;google&#39; | [optional] 
 **helmfile_ref** | **str**|  | [optional] [default to &#39;HEAD&#39;]
 **id** | **int**|  | [optional] 
 **location** | **str**|  | [optional] [default to &#39;us-central1-a&#39;]
 **name** | **str**| Required when creating | [optional] 
 **provider** | **str**|  | [optional] [default to google]
 **required_role** | **str**| If present, requires membership in the given role for mutations. Set to an empty string to clear. | [optional] 
 **requires_suitability** | **bool**|  | [optional] 
 **updated_at** | **datetime**|  | [optional] 
 **limit** | **int**| Control how many Clusters are returned (default 0, meaning all) | [optional] 
 **offset** | **int**| Control the offset for the returned Clusters (default 0) | [optional] 

### Return type

[**List[SherlockClusterV3]**](SherlockClusterV3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request |  -  |
**403** | Forbidden |  -  |
**404** | Not Found |  -  |
**407** | Proxy Authentication Required |  -  |
**409** | Conflict |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **api_clusters_v3_post**
> SherlockClusterV3 api_clusters_v3_post(cluster)

Create a Cluster

Create a Cluster.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_cluster_v3 import SherlockClusterV3
from sherlock_python_client.models.sherlock_cluster_v3_create import SherlockClusterV3Create
from sherlock_python_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://sherlock.dsp-devops-prod.broadinstitute.org
# See configuration.py for a list of all supported configuration parameters.
configuration = sherlock_python_client.Configuration(
    host = "https://sherlock.dsp-devops-prod.broadinstitute.org"
)


# Enter a context with an instance of the API client
with sherlock_python_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = sherlock_python_client.ClustersApi(api_client)
    cluster = sherlock_python_client.SherlockClusterV3Create() # SherlockClusterV3Create | The Cluster to create

    try:
        # Create a Cluster
        api_response = api_instance.api_clusters_v3_post(cluster)
        print("The response of ClustersApi->api_clusters_v3_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ClustersApi->api_clusters_v3_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | [**SherlockClusterV3Create**](SherlockClusterV3Create.md)| The Cluster to create | 

### Return type

[**SherlockClusterV3**](SherlockClusterV3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Created |  -  |
**400** | Bad Request |  -  |
**403** | Forbidden |  -  |
**404** | Not Found |  -  |
**407** | Proxy Authentication Required |  -  |
**409** | Conflict |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **api_clusters_v3_selector_delete**
> SherlockClusterV3 api_clusters_v3_selector_delete(selector)

Delete an individual Cluster

Delete an individual Cluster by its ID.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_cluster_v3 import SherlockClusterV3
from sherlock_python_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://sherlock.dsp-devops-prod.broadinstitute.org
# See configuration.py for a list of all supported configuration parameters.
configuration = sherlock_python_client.Configuration(
    host = "https://sherlock.dsp-devops-prod.broadinstitute.org"
)


# Enter a context with an instance of the API client
with sherlock_python_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = sherlock_python_client.ClustersApi(api_client)
    selector = 'selector_example' # str | The selector of the Cluster, which can be either a numeric ID or the name.

    try:
        # Delete an individual Cluster
        api_response = api_instance.api_clusters_v3_selector_delete(selector)
        print("The response of ClustersApi->api_clusters_v3_selector_delete:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ClustersApi->api_clusters_v3_selector_delete: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the Cluster, which can be either a numeric ID or the name. | 

### Return type

[**SherlockClusterV3**](SherlockClusterV3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request |  -  |
**403** | Forbidden |  -  |
**404** | Not Found |  -  |
**407** | Proxy Authentication Required |  -  |
**409** | Conflict |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **api_clusters_v3_selector_get**
> SherlockClusterV3 api_clusters_v3_selector_get(selector)

Get an individual Cluster

Get an individual Cluster.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_cluster_v3 import SherlockClusterV3
from sherlock_python_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://sherlock.dsp-devops-prod.broadinstitute.org
# See configuration.py for a list of all supported configuration parameters.
configuration = sherlock_python_client.Configuration(
    host = "https://sherlock.dsp-devops-prod.broadinstitute.org"
)


# Enter a context with an instance of the API client
with sherlock_python_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = sherlock_python_client.ClustersApi(api_client)
    selector = 'selector_example' # str | The selector of the Cluster, which can be either a numeric ID or the name.

    try:
        # Get an individual Cluster
        api_response = api_instance.api_clusters_v3_selector_get(selector)
        print("The response of ClustersApi->api_clusters_v3_selector_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ClustersApi->api_clusters_v3_selector_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the Cluster, which can be either a numeric ID or the name. | 

### Return type

[**SherlockClusterV3**](SherlockClusterV3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request |  -  |
**403** | Forbidden |  -  |
**404** | Not Found |  -  |
**407** | Proxy Authentication Required |  -  |
**409** | Conflict |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **api_clusters_v3_selector_patch**
> SherlockClusterV3 api_clusters_v3_selector_patch(selector, cluster)

Edit an individual Cluster

Edit an individual Cluster.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_cluster_v3 import SherlockClusterV3
from sherlock_python_client.models.sherlock_cluster_v3_edit import SherlockClusterV3Edit
from sherlock_python_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://sherlock.dsp-devops-prod.broadinstitute.org
# See configuration.py for a list of all supported configuration parameters.
configuration = sherlock_python_client.Configuration(
    host = "https://sherlock.dsp-devops-prod.broadinstitute.org"
)


# Enter a context with an instance of the API client
with sherlock_python_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = sherlock_python_client.ClustersApi(api_client)
    selector = 'selector_example' # str | The selector of the Cluster, which can be either a numeric ID or the name.
    cluster = sherlock_python_client.SherlockClusterV3Edit() # SherlockClusterV3Edit | The edits to make to the Cluster

    try:
        # Edit an individual Cluster
        api_response = api_instance.api_clusters_v3_selector_patch(selector, cluster)
        print("The response of ClustersApi->api_clusters_v3_selector_patch:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ClustersApi->api_clusters_v3_selector_patch: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the Cluster, which can be either a numeric ID or the name. | 
 **cluster** | [**SherlockClusterV3Edit**](SherlockClusterV3Edit.md)| The edits to make to the Cluster | 

### Return type

[**SherlockClusterV3**](SherlockClusterV3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request |  -  |
**403** | Forbidden |  -  |
**404** | Not Found |  -  |
**407** | Proxy Authentication Required |  -  |
**409** | Conflict |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

