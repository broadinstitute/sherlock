# sherlock_python_client.ChartVersionsApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_chart_versions_procedures_v3_changelog_get**](ChartVersionsApi.md#api_chart_versions_procedures_v3_changelog_get) | **GET** /api/chart-versions/procedures/v3/changelog | Get a changelog between two ChartVersions
[**api_chart_versions_v3_get**](ChartVersionsApi.md#api_chart_versions_v3_get) | **GET** /api/chart-versions/v3 | List ChartVersions matching a filter
[**api_chart_versions_v3_put**](ChartVersionsApi.md#api_chart_versions_v3_put) | **PUT** /api/chart-versions/v3 | Upsert a ChartVersion
[**api_chart_versions_v3_selector_get**](ChartVersionsApi.md#api_chart_versions_v3_selector_get) | **GET** /api/chart-versions/v3/{selector} | Get an individual ChartVersion
[**api_chart_versions_v3_selector_patch**](ChartVersionsApi.md#api_chart_versions_v3_selector_patch) | **PATCH** /api/chart-versions/v3/{selector} | Edit an individual ChartVersion


# **api_chart_versions_procedures_v3_changelog_get**
> SherlockChartVersionV3ChangelogResponse api_chart_versions_procedures_v3_changelog_get(child, parent)

Get a changelog between two ChartVersions

Get the path through parent references from a child ChartVersion (inclusive) to a parent ChartVersion (exclusive), if possible. Because parent references point from newer children to older parents, the newer ChartVersion should be the child. The result will always exclude the parent.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_chart_version_v3_changelog_response import SherlockChartVersionV3ChangelogResponse
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
    api_instance = sherlock_python_client.ChartVersionsApi(api_client)
    child = 'child_example' # str | The selector of the newer ChartVersion for the changelog
    parent = 'parent_example' # str | The selector of the older ChartVersion for the changelog

    try:
        # Get a changelog between two ChartVersions
        api_response = api_instance.api_chart_versions_procedures_v3_changelog_get(child, parent)
        print("The response of ChartVersionsApi->api_chart_versions_procedures_v3_changelog_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ChartVersionsApi->api_chart_versions_procedures_v3_changelog_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **child** | **str**| The selector of the newer ChartVersion for the changelog | 
 **parent** | **str**| The selector of the older ChartVersion for the changelog | 

### Return type

[**SherlockChartVersionV3ChangelogResponse**](SherlockChartVersionV3ChangelogResponse.md)

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

# **api_chart_versions_v3_get**
> List[SherlockChartVersionV3] api_chart_versions_v3_get(authored_by=authored_by, chart=chart, chart_version=chart_version, created_at=created_at, description=description, id=id, parent_chart_version=parent_chart_version, updated_at=updated_at, limit=limit, offset=offset)

List ChartVersions matching a filter

List ChartVersions matching a filter.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_chart_version_v3 import SherlockChartVersionV3
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
    api_instance = sherlock_python_client.ChartVersionsApi(api_client)
    authored_by = 'authored_by_example' # str |  (optional)
    chart = 'chart_example' # str | Required when creating (optional)
    chart_version = 'chart_version_example' # str | Required when creating (optional)
    created_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    description = 'description_example' # str | Generally the Git commit message (optional)
    id = 56 # int |  (optional)
    parent_chart_version = 'parent_chart_version_example' # str |  (optional)
    updated_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    limit = 56 # int | Control how many ChartVersions are returned (default 100) (optional)
    offset = 56 # int | Control the offset for the returned ChartVersions (default 0) (optional)

    try:
        # List ChartVersions matching a filter
        api_response = api_instance.api_chart_versions_v3_get(authored_by=authored_by, chart=chart, chart_version=chart_version, created_at=created_at, description=description, id=id, parent_chart_version=parent_chart_version, updated_at=updated_at, limit=limit, offset=offset)
        print("The response of ChartVersionsApi->api_chart_versions_v3_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ChartVersionsApi->api_chart_versions_v3_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authored_by** | **str**|  | [optional] 
 **chart** | **str**| Required when creating | [optional] 
 **chart_version** | **str**| Required when creating | [optional] 
 **created_at** | **datetime**|  | [optional] 
 **description** | **str**| Generally the Git commit message | [optional] 
 **id** | **int**|  | [optional] 
 **parent_chart_version** | **str**|  | [optional] 
 **updated_at** | **datetime**|  | [optional] 
 **limit** | **int**| Control how many ChartVersions are returned (default 100) | [optional] 
 **offset** | **int**| Control the offset for the returned ChartVersions (default 0) | [optional] 

### Return type

[**List[SherlockChartVersionV3]**](SherlockChartVersionV3.md)

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

# **api_chart_versions_v3_put**
> SherlockChartVersionV3 api_chart_versions_v3_put(chart_version)

Upsert a ChartVersion

Upsert a ChartVersion.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_chart_version_v3 import SherlockChartVersionV3
from sherlock_python_client.models.sherlock_chart_version_v3_create import SherlockChartVersionV3Create
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
    api_instance = sherlock_python_client.ChartVersionsApi(api_client)
    chart_version = sherlock_python_client.SherlockChartVersionV3Create() # SherlockChartVersionV3Create | The ChartVersion to upsert

    try:
        # Upsert a ChartVersion
        api_response = api_instance.api_chart_versions_v3_put(chart_version)
        print("The response of ChartVersionsApi->api_chart_versions_v3_put:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ChartVersionsApi->api_chart_versions_v3_put: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chart_version** | [**SherlockChartVersionV3Create**](SherlockChartVersionV3Create.md)| The ChartVersion to upsert | 

### Return type

[**SherlockChartVersionV3**](SherlockChartVersionV3.md)

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

# **api_chart_versions_v3_selector_get**
> SherlockChartVersionV3 api_chart_versions_v3_selector_get(selector)

Get an individual ChartVersion

Get an individual ChartVersion.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_chart_version_v3 import SherlockChartVersionV3
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
    api_instance = sherlock_python_client.ChartVersionsApi(api_client)
    selector = 'selector_example' # str | The selector of the ChartVersion, which can be either a numeric ID or chart/version.

    try:
        # Get an individual ChartVersion
        api_response = api_instance.api_chart_versions_v3_selector_get(selector)
        print("The response of ChartVersionsApi->api_chart_versions_v3_selector_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ChartVersionsApi->api_chart_versions_v3_selector_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the ChartVersion, which can be either a numeric ID or chart/version. | 

### Return type

[**SherlockChartVersionV3**](SherlockChartVersionV3.md)

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

# **api_chart_versions_v3_selector_patch**
> SherlockChartVersionV3 api_chart_versions_v3_selector_patch(selector, chart_version)

Edit an individual ChartVersion

Edit an individual ChartVersion.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_chart_version_v3 import SherlockChartVersionV3
from sherlock_python_client.models.sherlock_chart_version_v3_edit import SherlockChartVersionV3Edit
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
    api_instance = sherlock_python_client.ChartVersionsApi(api_client)
    selector = 'selector_example' # str | The selector of the ChartVersion, which can be either a numeric ID or chart/version.
    chart_version = sherlock_python_client.SherlockChartVersionV3Edit() # SherlockChartVersionV3Edit | The edits to make to the ChartVersion

    try:
        # Edit an individual ChartVersion
        api_response = api_instance.api_chart_versions_v3_selector_patch(selector, chart_version)
        print("The response of ChartVersionsApi->api_chart_versions_v3_selector_patch:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ChartVersionsApi->api_chart_versions_v3_selector_patch: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the ChartVersion, which can be either a numeric ID or chart/version. | 
 **chart_version** | [**SherlockChartVersionV3Edit**](SherlockChartVersionV3Edit.md)| The edits to make to the ChartVersion | 

### Return type

[**SherlockChartVersionV3**](SherlockChartVersionV3.md)

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

