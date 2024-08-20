# sherlock_python_client.DatabaseInstancesApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_database_instances_v3_get**](DatabaseInstancesApi.md#api_database_instances_v3_get) | **GET** /api/database-instances/v3 | List DatabaseInstances matching a filter
[**api_database_instances_v3_post**](DatabaseInstancesApi.md#api_database_instances_v3_post) | **POST** /api/database-instances/v3 | Create a DatabaseInstance
[**api_database_instances_v3_put**](DatabaseInstancesApi.md#api_database_instances_v3_put) | **PUT** /api/database-instances/v3 | Create or edit a DatabaseInstance
[**api_database_instances_v3_selector_delete**](DatabaseInstancesApi.md#api_database_instances_v3_selector_delete) | **DELETE** /api/database-instances/v3/{selector} | Delete an individual DatabaseInstance
[**api_database_instances_v3_selector_get**](DatabaseInstancesApi.md#api_database_instances_v3_selector_get) | **GET** /api/database-instances/v3/{selector} | Get an individual DatabaseInstance
[**api_database_instances_v3_selector_patch**](DatabaseInstancesApi.md#api_database_instances_v3_selector_patch) | **PATCH** /api/database-instances/v3/{selector} | Edit an individual DatabaseInstance


# **api_database_instances_v3_get**
> List[SherlockDatabaseInstanceV3] api_database_instances_v3_get(chart_release=chart_release, created_at=created_at, default_database=default_database, google_project=google_project, id=id, instance_name=instance_name, platform=platform, updated_at=updated_at, limit=limit, offset=offset)

List DatabaseInstances matching a filter

List DatabaseInstances matching a filter.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_database_instance_v3 import SherlockDatabaseInstanceV3
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
    api_instance = sherlock_python_client.DatabaseInstancesApi(api_client)
    chart_release = 'chart_release_example' # str | Required when creating (optional)
    created_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    default_database = 'default_database_example' # str | When creating, defaults to the chart name (optional)
    google_project = 'google_project_example' # str | Required if platform is 'google' (optional)
    id = 56 # int |  (optional)
    instance_name = 'instance_name_example' # str | Required if platform is 'google' or 'azure' (optional)
    platform = 'kubernetes' # str | 'google', 'azure', or default 'kubernetes' (optional) (default to 'kubernetes')
    updated_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    limit = 56 # int | Control how many DatabaseInstances are returned (default 0, meaning all) (optional)
    offset = 56 # int | Control the offset for the returned DatabaseInstances (default 0) (optional)

    try:
        # List DatabaseInstances matching a filter
        api_response = api_instance.api_database_instances_v3_get(chart_release=chart_release, created_at=created_at, default_database=default_database, google_project=google_project, id=id, instance_name=instance_name, platform=platform, updated_at=updated_at, limit=limit, offset=offset)
        print("The response of DatabaseInstancesApi->api_database_instances_v3_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DatabaseInstancesApi->api_database_instances_v3_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chart_release** | **str**| Required when creating | [optional] 
 **created_at** | **datetime**|  | [optional] 
 **default_database** | **str**| When creating, defaults to the chart name | [optional] 
 **google_project** | **str**| Required if platform is &#39;google&#39; | [optional] 
 **id** | **int**|  | [optional] 
 **instance_name** | **str**| Required if platform is &#39;google&#39; or &#39;azure&#39; | [optional] 
 **platform** | **str**| &#39;google&#39;, &#39;azure&#39;, or default &#39;kubernetes&#39; | [optional] [default to &#39;kubernetes&#39;]
 **updated_at** | **datetime**|  | [optional] 
 **limit** | **int**| Control how many DatabaseInstances are returned (default 0, meaning all) | [optional] 
 **offset** | **int**| Control the offset for the returned DatabaseInstances (default 0) | [optional] 

### Return type

[**List[SherlockDatabaseInstanceV3]**](SherlockDatabaseInstanceV3.md)

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

# **api_database_instances_v3_post**
> SherlockDatabaseInstanceV3 api_database_instances_v3_post(database_instance)

Create a DatabaseInstance

Create a DatabaseInstance.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_database_instance_v3 import SherlockDatabaseInstanceV3
from sherlock_python_client.models.sherlock_database_instance_v3_create import SherlockDatabaseInstanceV3Create
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
    api_instance = sherlock_python_client.DatabaseInstancesApi(api_client)
    database_instance = sherlock_python_client.SherlockDatabaseInstanceV3Create() # SherlockDatabaseInstanceV3Create | The DatabaseInstance to create

    try:
        # Create a DatabaseInstance
        api_response = api_instance.api_database_instances_v3_post(database_instance)
        print("The response of DatabaseInstancesApi->api_database_instances_v3_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DatabaseInstancesApi->api_database_instances_v3_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **database_instance** | [**SherlockDatabaseInstanceV3Create**](SherlockDatabaseInstanceV3Create.md)| The DatabaseInstance to create | 

### Return type

[**SherlockDatabaseInstanceV3**](SherlockDatabaseInstanceV3.md)

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

# **api_database_instances_v3_put**
> SherlockDatabaseInstanceV3 api_database_instances_v3_put(database_instance)

Create or edit a DatabaseInstance

Create or edit a DatabaseInstance, depending on whether one already exists for the chart release

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_database_instance_v3 import SherlockDatabaseInstanceV3
from sherlock_python_client.models.sherlock_database_instance_v3_create import SherlockDatabaseInstanceV3Create
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
    api_instance = sherlock_python_client.DatabaseInstancesApi(api_client)
    database_instance = sherlock_python_client.SherlockDatabaseInstanceV3Create() # SherlockDatabaseInstanceV3Create | The DatabaseInstance to create or edit. Defaults will only be set if creating.

    try:
        # Create or edit a DatabaseInstance
        api_response = api_instance.api_database_instances_v3_put(database_instance)
        print("The response of DatabaseInstancesApi->api_database_instances_v3_put:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DatabaseInstancesApi->api_database_instances_v3_put: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **database_instance** | [**SherlockDatabaseInstanceV3Create**](SherlockDatabaseInstanceV3Create.md)| The DatabaseInstance to create or edit. Defaults will only be set if creating. | 

### Return type

[**SherlockDatabaseInstanceV3**](SherlockDatabaseInstanceV3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**201** | Created |  -  |
**400** | Bad Request |  -  |
**403** | Forbidden |  -  |
**404** | Not Found |  -  |
**407** | Proxy Authentication Required |  -  |
**409** | Conflict |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **api_database_instances_v3_selector_delete**
> SherlockDatabaseInstanceV3 api_database_instances_v3_selector_delete(selector)

Delete an individual DatabaseInstance

Delete an individual DatabaseInstance by its selector.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_database_instance_v3 import SherlockDatabaseInstanceV3
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
    api_instance = sherlock_python_client.DatabaseInstancesApi(api_client)
    selector = 'selector_example' # str | The selector of the DatabaseInstance, which can be either a numeric ID or 'chart-release/' followed by a chart release selector.

    try:
        # Delete an individual DatabaseInstance
        api_response = api_instance.api_database_instances_v3_selector_delete(selector)
        print("The response of DatabaseInstancesApi->api_database_instances_v3_selector_delete:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DatabaseInstancesApi->api_database_instances_v3_selector_delete: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the DatabaseInstance, which can be either a numeric ID or &#39;chart-release/&#39; followed by a chart release selector. | 

### Return type

[**SherlockDatabaseInstanceV3**](SherlockDatabaseInstanceV3.md)

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

# **api_database_instances_v3_selector_get**
> SherlockDatabaseInstanceV3 api_database_instances_v3_selector_get(selector)

Get an individual DatabaseInstance

Get an individual DatabaseInstance by its selector.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_database_instance_v3 import SherlockDatabaseInstanceV3
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
    api_instance = sherlock_python_client.DatabaseInstancesApi(api_client)
    selector = 'selector_example' # str | The selector of the DatabaseInstance, which can be either a numeric ID or 'chart-release/' followed by a chart release selector.

    try:
        # Get an individual DatabaseInstance
        api_response = api_instance.api_database_instances_v3_selector_get(selector)
        print("The response of DatabaseInstancesApi->api_database_instances_v3_selector_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DatabaseInstancesApi->api_database_instances_v3_selector_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the DatabaseInstance, which can be either a numeric ID or &#39;chart-release/&#39; followed by a chart release selector. | 

### Return type

[**SherlockDatabaseInstanceV3**](SherlockDatabaseInstanceV3.md)

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

# **api_database_instances_v3_selector_patch**
> SherlockDatabaseInstanceV3 api_database_instances_v3_selector_patch(selector, database_instance)

Edit an individual DatabaseInstance

Edit an individual DatabaseInstance by its selector.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_database_instance_v3 import SherlockDatabaseInstanceV3
from sherlock_python_client.models.sherlock_database_instance_v3_edit import SherlockDatabaseInstanceV3Edit
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
    api_instance = sherlock_python_client.DatabaseInstancesApi(api_client)
    selector = 'selector_example' # str | The selector of the DatabaseInstance, which can be either a numeric ID or 'chart-release/' followed by a chart release selector.
    database_instance = sherlock_python_client.SherlockDatabaseInstanceV3Edit() # SherlockDatabaseInstanceV3Edit | The edits to make to the DatabaseInstance

    try:
        # Edit an individual DatabaseInstance
        api_response = api_instance.api_database_instances_v3_selector_patch(selector, database_instance)
        print("The response of DatabaseInstancesApi->api_database_instances_v3_selector_patch:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DatabaseInstancesApi->api_database_instances_v3_selector_patch: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the DatabaseInstance, which can be either a numeric ID or &#39;chart-release/&#39; followed by a chart release selector. | 
 **database_instance** | [**SherlockDatabaseInstanceV3Edit**](SherlockDatabaseInstanceV3Edit.md)| The edits to make to the DatabaseInstance | 

### Return type

[**SherlockDatabaseInstanceV3**](SherlockDatabaseInstanceV3.md)

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

