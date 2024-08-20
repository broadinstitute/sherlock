# sherlock_python_client.ChartReleasesApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_chart_releases_v3_get**](ChartReleasesApi.md#api_chart_releases_v3_get) | **GET** /api/chart-releases/v3 | List ChartReleases matching a filter
[**api_chart_releases_v3_post**](ChartReleasesApi.md#api_chart_releases_v3_post) | **POST** /api/chart-releases/v3 | Create a ChartRelease
[**api_chart_releases_v3_selector_delete**](ChartReleasesApi.md#api_chart_releases_v3_selector_delete) | **DELETE** /api/chart-releases/v3/{selector} | Delete an individual ChartRelease
[**api_chart_releases_v3_selector_get**](ChartReleasesApi.md#api_chart_releases_v3_selector_get) | **GET** /api/chart-releases/v3/{selector} | Get an individual ChartRelease
[**api_chart_releases_v3_selector_patch**](ChartReleasesApi.md#api_chart_releases_v3_selector_patch) | **PATCH** /api/chart-releases/v3/{selector} | Edit an individual ChartRelease


# **api_chart_releases_v3_get**
> List[SherlockChartReleaseV3] api_chart_releases_v3_get(app_version_branch=app_version_branch, app_version_commit=app_version_commit, app_version_exact=app_version_exact, app_version_follow_chart_release=app_version_follow_chart_release, app_version_reference=app_version_reference, app_version_resolver=app_version_resolver, chart=chart, chart_version_exact=chart_version_exact, chart_version_follow_chart_release=chart_version_follow_chart_release, chart_version_reference=chart_version_reference, chart_version_resolver=chart_version_resolver, cluster=cluster, created_at=created_at, destination_type=destination_type, environment=environment, helmfile_ref=helmfile_ref, helmfile_ref_enabled=helmfile_ref_enabled, id=id, included_in_bulk_changesets=included_in_bulk_changesets, name=name, namespace=namespace, pagerduty_integration=pagerduty_integration, port=port, protocol=protocol, resolved_at=resolved_at, subdomain=subdomain, updated_at=updated_at, limit=limit, offset=offset)

List ChartReleases matching a filter

List ChartReleases matching a filter.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_chart_release_v3 import SherlockChartReleaseV3
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
    api_instance = sherlock_python_client.ChartReleasesApi(api_client)
    app_version_branch = 'app_version_branch_example' # str | When creating, will default to the app's mainline branch if no other app version info is present (optional)
    app_version_commit = 'app_version_commit_example' # str |  (optional)
    app_version_exact = 'app_version_exact_example' # str |  (optional)
    app_version_follow_chart_release = 'app_version_follow_chart_release_example' # str |  (optional)
    app_version_reference = 'app_version_reference_example' # str |  (optional)
    app_version_resolver = 'app_version_resolver_example' # str | // When creating, will default to automatically reference any provided app version fields (optional)
    chart = 'chart_example' # str | Required when creating (optional)
    chart_version_exact = 'chart_version_exact_example' # str |  (optional)
    chart_version_follow_chart_release = 'chart_version_follow_chart_release_example' # str |  (optional)
    chart_version_reference = 'chart_version_reference_example' # str |  (optional)
    chart_version_resolver = 'chart_version_resolver_example' # str | When creating, will default to automatically reference any provided chart version (optional)
    cluster = 'cluster_example' # str | When creating, will default the environment's default cluster, if provided. Either this or environment must be provided. (optional)
    created_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    destination_type = 'destination_type_example' # str | Calculated field (optional)
    environment = 'environment_example' # str | Either this or cluster must be provided. (optional)
    helmfile_ref = 'HEAD' # str |  (optional) (default to 'HEAD')
    helmfile_ref_enabled = False # bool |  (optional) (default to False)
    id = 56 # int |  (optional)
    included_in_bulk_changesets = True # bool |  (optional) (default to True)
    name = 'name_example' # str | When creating, will be calculated if left empty (optional)
    namespace = 'namespace_example' # str | When creating, will default to the environment's default namespace, if provided (optional)
    pagerduty_integration = 'pagerduty_integration_example' # str |  (optional)
    port = 56 # int | When creating, will use the chart's default if left empty (optional)
    protocol = 'protocol_example' # str | When creating, will use the chart's default if left empty (optional)
    resolved_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    subdomain = 'subdomain_example' # str | When creating, will use the chart's default if left empty (optional)
    updated_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    limit = 56 # int | Control how many ChartReleases are returned (default 0, meaning all) (optional)
    offset = 56 # int | Control the offset for the returned ChartReleases (default 0) (optional)

    try:
        # List ChartReleases matching a filter
        api_response = api_instance.api_chart_releases_v3_get(app_version_branch=app_version_branch, app_version_commit=app_version_commit, app_version_exact=app_version_exact, app_version_follow_chart_release=app_version_follow_chart_release, app_version_reference=app_version_reference, app_version_resolver=app_version_resolver, chart=chart, chart_version_exact=chart_version_exact, chart_version_follow_chart_release=chart_version_follow_chart_release, chart_version_reference=chart_version_reference, chart_version_resolver=chart_version_resolver, cluster=cluster, created_at=created_at, destination_type=destination_type, environment=environment, helmfile_ref=helmfile_ref, helmfile_ref_enabled=helmfile_ref_enabled, id=id, included_in_bulk_changesets=included_in_bulk_changesets, name=name, namespace=namespace, pagerduty_integration=pagerduty_integration, port=port, protocol=protocol, resolved_at=resolved_at, subdomain=subdomain, updated_at=updated_at, limit=limit, offset=offset)
        print("The response of ChartReleasesApi->api_chart_releases_v3_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ChartReleasesApi->api_chart_releases_v3_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_version_branch** | **str**| When creating, will default to the app&#39;s mainline branch if no other app version info is present | [optional] 
 **app_version_commit** | **str**|  | [optional] 
 **app_version_exact** | **str**|  | [optional] 
 **app_version_follow_chart_release** | **str**|  | [optional] 
 **app_version_reference** | **str**|  | [optional] 
 **app_version_resolver** | **str**| // When creating, will default to automatically reference any provided app version fields | [optional] 
 **chart** | **str**| Required when creating | [optional] 
 **chart_version_exact** | **str**|  | [optional] 
 **chart_version_follow_chart_release** | **str**|  | [optional] 
 **chart_version_reference** | **str**|  | [optional] 
 **chart_version_resolver** | **str**| When creating, will default to automatically reference any provided chart version | [optional] 
 **cluster** | **str**| When creating, will default the environment&#39;s default cluster, if provided. Either this or environment must be provided. | [optional] 
 **created_at** | **datetime**|  | [optional] 
 **destination_type** | **str**| Calculated field | [optional] 
 **environment** | **str**| Either this or cluster must be provided. | [optional] 
 **helmfile_ref** | **str**|  | [optional] [default to &#39;HEAD&#39;]
 **helmfile_ref_enabled** | **bool**|  | [optional] [default to False]
 **id** | **int**|  | [optional] 
 **included_in_bulk_changesets** | **bool**|  | [optional] [default to True]
 **name** | **str**| When creating, will be calculated if left empty | [optional] 
 **namespace** | **str**| When creating, will default to the environment&#39;s default namespace, if provided | [optional] 
 **pagerduty_integration** | **str**|  | [optional] 
 **port** | **int**| When creating, will use the chart&#39;s default if left empty | [optional] 
 **protocol** | **str**| When creating, will use the chart&#39;s default if left empty | [optional] 
 **resolved_at** | **datetime**|  | [optional] 
 **subdomain** | **str**| When creating, will use the chart&#39;s default if left empty | [optional] 
 **updated_at** | **datetime**|  | [optional] 
 **limit** | **int**| Control how many ChartReleases are returned (default 0, meaning all) | [optional] 
 **offset** | **int**| Control the offset for the returned ChartReleases (default 0) | [optional] 

### Return type

[**List[SherlockChartReleaseV3]**](SherlockChartReleaseV3.md)

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

# **api_chart_releases_v3_post**
> SherlockChartReleaseV3 api_chart_releases_v3_post(chart_release)

Create a ChartRelease

Create a ChartRelease.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_chart_release_v3 import SherlockChartReleaseV3
from sherlock_python_client.models.sherlock_chart_release_v3_create import SherlockChartReleaseV3Create
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
    api_instance = sherlock_python_client.ChartReleasesApi(api_client)
    chart_release = sherlock_python_client.SherlockChartReleaseV3Create() # SherlockChartReleaseV3Create | The ChartRelease to create

    try:
        # Create a ChartRelease
        api_response = api_instance.api_chart_releases_v3_post(chart_release)
        print("The response of ChartReleasesApi->api_chart_releases_v3_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ChartReleasesApi->api_chart_releases_v3_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chart_release** | [**SherlockChartReleaseV3Create**](SherlockChartReleaseV3Create.md)| The ChartRelease to create | 

### Return type

[**SherlockChartReleaseV3**](SherlockChartReleaseV3.md)

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

# **api_chart_releases_v3_selector_delete**
> SherlockChartReleaseV3 api_chart_releases_v3_selector_delete(selector)

Delete an individual ChartRelease

Delete an individual ChartRelease by its ID.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_chart_release_v3 import SherlockChartReleaseV3
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
    api_instance = sherlock_python_client.ChartReleasesApi(api_client)
    selector = 'selector_example' # str | The selector of the ChartRelease, which can be either a numeric ID, the name, environment + '/' + chart, or cluster + '/' + namespace + '/' + chart.

    try:
        # Delete an individual ChartRelease
        api_response = api_instance.api_chart_releases_v3_selector_delete(selector)
        print("The response of ChartReleasesApi->api_chart_releases_v3_selector_delete:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ChartReleasesApi->api_chart_releases_v3_selector_delete: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the ChartRelease, which can be either a numeric ID, the name, environment + &#39;/&#39; + chart, or cluster + &#39;/&#39; + namespace + &#39;/&#39; + chart. | 

### Return type

[**SherlockChartReleaseV3**](SherlockChartReleaseV3.md)

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

# **api_chart_releases_v3_selector_get**
> SherlockChartReleaseV3 api_chart_releases_v3_selector_get(selector)

Get an individual ChartRelease

Get an individual ChartRelease.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_chart_release_v3 import SherlockChartReleaseV3
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
    api_instance = sherlock_python_client.ChartReleasesApi(api_client)
    selector = 'selector_example' # str | The selector of the ChartRelease, which can be either a numeric ID, the name, environment + '/' + chart, or cluster + '/' + namespace + '/' + chart.

    try:
        # Get an individual ChartRelease
        api_response = api_instance.api_chart_releases_v3_selector_get(selector)
        print("The response of ChartReleasesApi->api_chart_releases_v3_selector_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ChartReleasesApi->api_chart_releases_v3_selector_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the ChartRelease, which can be either a numeric ID, the name, environment + &#39;/&#39; + chart, or cluster + &#39;/&#39; + namespace + &#39;/&#39; + chart. | 

### Return type

[**SherlockChartReleaseV3**](SherlockChartReleaseV3.md)

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

# **api_chart_releases_v3_selector_patch**
> SherlockChartReleaseV3 api_chart_releases_v3_selector_patch(selector, chart_release)

Edit an individual ChartRelease

Edit an individual ChartRelease.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_chart_release_v3 import SherlockChartReleaseV3
from sherlock_python_client.models.sherlock_chart_release_v3_edit import SherlockChartReleaseV3Edit
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
    api_instance = sherlock_python_client.ChartReleasesApi(api_client)
    selector = 'selector_example' # str | The selector of the ChartRelease, which can be either a numeric ID, the name, environment + '/' + chart, or cluster + '/' + namespace + '/' + chart.
    chart_release = sherlock_python_client.SherlockChartReleaseV3Edit() # SherlockChartReleaseV3Edit | The edits to make to the ChartRelease

    try:
        # Edit an individual ChartRelease
        api_response = api_instance.api_chart_releases_v3_selector_patch(selector, chart_release)
        print("The response of ChartReleasesApi->api_chart_releases_v3_selector_patch:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ChartReleasesApi->api_chart_releases_v3_selector_patch: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the ChartRelease, which can be either a numeric ID, the name, environment + &#39;/&#39; + chart, or cluster + &#39;/&#39; + namespace + &#39;/&#39; + chart. | 
 **chart_release** | [**SherlockChartReleaseV3Edit**](SherlockChartReleaseV3Edit.md)| The edits to make to the ChartRelease | 

### Return type

[**SherlockChartReleaseV3**](SherlockChartReleaseV3.md)

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

