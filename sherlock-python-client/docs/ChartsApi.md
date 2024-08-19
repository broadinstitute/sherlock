# sherlock_python_client.ChartsApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_charts_v3_get**](ChartsApi.md#api_charts_v3_get) | **GET** /api/charts/v3 | List Charts matching a filter
[**api_charts_v3_post**](ChartsApi.md#api_charts_v3_post) | **POST** /api/charts/v3 | Create a Chart
[**api_charts_v3_selector_delete**](ChartsApi.md#api_charts_v3_selector_delete) | **DELETE** /api/charts/v3/{selector} | Delete an individual Chart
[**api_charts_v3_selector_get**](ChartsApi.md#api_charts_v3_selector_get) | **GET** /api/charts/v3/{selector} | Get an individual Chart
[**api_charts_v3_selector_patch**](ChartsApi.md#api_charts_v3_selector_patch) | **PATCH** /api/charts/v3/{selector} | Edit an individual Chart


# **api_charts_v3_get**
> List[SherlockChartV3] api_charts_v3_get(app_image_git_main_branch=app_image_git_main_branch, app_image_git_repo=app_image_git_repo, chart_exposes_endpoint=chart_exposes_endpoint, chart_repo=chart_repo, created_at=created_at, default_port=default_port, default_protocol=default_protocol, default_subdomain=default_subdomain, description=description, id=id, name=name, pact_participant=pact_participant, playbook_url=playbook_url, updated_at=updated_at, limit=limit, offset=offset)

List Charts matching a filter

List Charts matching a filter.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_chart_v3 import SherlockChartV3
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
    api_instance = sherlock_python_client.ChartsApi(api_client)
    app_image_git_main_branch = 'app_image_git_main_branch_example' # str |  (optional)
    app_image_git_repo = 'app_image_git_repo_example' # str |  (optional)
    chart_exposes_endpoint = False # bool | Indicates if the default subdomain, protocol, and port fields are relevant for this chart (optional) (default to False)
    chart_repo = 'terra-helm' # str |  (optional) (default to 'terra-helm')
    created_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    default_port = 443 # int |  (optional) (default to 443)
    default_protocol = 'https' # str |  (optional) (default to 'https')
    default_subdomain = 'default_subdomain_example' # str | When creating, will default to the name of the chart (optional)
    description = 'description_example' # str |  (optional)
    id = 56 # int |  (optional)
    name = 'name_example' # str | Required when creating (optional)
    pact_participant = False # bool |  (optional) (default to False)
    playbook_url = 'playbook_url_example' # str |  (optional)
    updated_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    limit = 56 # int | Control how many Charts are returned (default 0, meaning all) (optional)
    offset = 56 # int | Control the offset for the returned Charts (default 0) (optional)

    try:
        # List Charts matching a filter
        api_response = api_instance.api_charts_v3_get(app_image_git_main_branch=app_image_git_main_branch, app_image_git_repo=app_image_git_repo, chart_exposes_endpoint=chart_exposes_endpoint, chart_repo=chart_repo, created_at=created_at, default_port=default_port, default_protocol=default_protocol, default_subdomain=default_subdomain, description=description, id=id, name=name, pact_participant=pact_participant, playbook_url=playbook_url, updated_at=updated_at, limit=limit, offset=offset)
        print("The response of ChartsApi->api_charts_v3_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ChartsApi->api_charts_v3_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_image_git_main_branch** | **str**|  | [optional] 
 **app_image_git_repo** | **str**|  | [optional] 
 **chart_exposes_endpoint** | **bool**| Indicates if the default subdomain, protocol, and port fields are relevant for this chart | [optional] [default to False]
 **chart_repo** | **str**|  | [optional] [default to &#39;terra-helm&#39;]
 **created_at** | **datetime**|  | [optional] 
 **default_port** | **int**|  | [optional] [default to 443]
 **default_protocol** | **str**|  | [optional] [default to &#39;https&#39;]
 **default_subdomain** | **str**| When creating, will default to the name of the chart | [optional] 
 **description** | **str**|  | [optional] 
 **id** | **int**|  | [optional] 
 **name** | **str**| Required when creating | [optional] 
 **pact_participant** | **bool**|  | [optional] [default to False]
 **playbook_url** | **str**|  | [optional] 
 **updated_at** | **datetime**|  | [optional] 
 **limit** | **int**| Control how many Charts are returned (default 0, meaning all) | [optional] 
 **offset** | **int**| Control the offset for the returned Charts (default 0) | [optional] 

### Return type

[**List[SherlockChartV3]**](SherlockChartV3.md)

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

# **api_charts_v3_post**
> SherlockChartV3 api_charts_v3_post(chart)

Create a Chart

Create a Chart.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_chart_v3 import SherlockChartV3
from sherlock_python_client.models.sherlock_chart_v3_create import SherlockChartV3Create
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
    api_instance = sherlock_python_client.ChartsApi(api_client)
    chart = sherlock_python_client.SherlockChartV3Create() # SherlockChartV3Create | The Chart to create

    try:
        # Create a Chart
        api_response = api_instance.api_charts_v3_post(chart)
        print("The response of ChartsApi->api_charts_v3_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ChartsApi->api_charts_v3_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chart** | [**SherlockChartV3Create**](SherlockChartV3Create.md)| The Chart to create | 

### Return type

[**SherlockChartV3**](SherlockChartV3.md)

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

# **api_charts_v3_selector_delete**
> SherlockChartV3 api_charts_v3_selector_delete(selector)

Delete an individual Chart

Delete an individual Chart by its ID.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_chart_v3 import SherlockChartV3
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
    api_instance = sherlock_python_client.ChartsApi(api_client)
    selector = 'selector_example' # str | The selector of the Chart, which can be either a numeric ID or the name.

    try:
        # Delete an individual Chart
        api_response = api_instance.api_charts_v3_selector_delete(selector)
        print("The response of ChartsApi->api_charts_v3_selector_delete:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ChartsApi->api_charts_v3_selector_delete: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the Chart, which can be either a numeric ID or the name. | 

### Return type

[**SherlockChartV3**](SherlockChartV3.md)

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

# **api_charts_v3_selector_get**
> SherlockChartV3 api_charts_v3_selector_get(selector)

Get an individual Chart

Get an individual Chart.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_chart_v3 import SherlockChartV3
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
    api_instance = sherlock_python_client.ChartsApi(api_client)
    selector = 'selector_example' # str | The selector of the Chart, which can be either a numeric ID or the name.

    try:
        # Get an individual Chart
        api_response = api_instance.api_charts_v3_selector_get(selector)
        print("The response of ChartsApi->api_charts_v3_selector_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ChartsApi->api_charts_v3_selector_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the Chart, which can be either a numeric ID or the name. | 

### Return type

[**SherlockChartV3**](SherlockChartV3.md)

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

# **api_charts_v3_selector_patch**
> SherlockChartV3 api_charts_v3_selector_patch(selector, chart)

Edit an individual Chart

Edit an individual Chart.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_chart_v3 import SherlockChartV3
from sherlock_python_client.models.sherlock_chart_v3_edit import SherlockChartV3Edit
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
    api_instance = sherlock_python_client.ChartsApi(api_client)
    selector = 'selector_example' # str | The selector of the Chart, which can be either a numeric ID or the name.
    chart = sherlock_python_client.SherlockChartV3Edit() # SherlockChartV3Edit | The edits to make to the Chart

    try:
        # Edit an individual Chart
        api_response = api_instance.api_charts_v3_selector_patch(selector, chart)
        print("The response of ChartsApi->api_charts_v3_selector_patch:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ChartsApi->api_charts_v3_selector_patch: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the Chart, which can be either a numeric ID or the name. | 
 **chart** | [**SherlockChartV3Edit**](SherlockChartV3Edit.md)| The edits to make to the Chart | 

### Return type

[**SherlockChartV3**](SherlockChartV3.md)

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

