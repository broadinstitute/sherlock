# sherlock_python_client.AppVersionsApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_app_versions_procedures_v3_changelog_get**](AppVersionsApi.md#api_app_versions_procedures_v3_changelog_get) | **GET** /api/app-versions/procedures/v3/changelog | Get a changelog between two AppVersions
[**api_app_versions_v3_get**](AppVersionsApi.md#api_app_versions_v3_get) | **GET** /api/app-versions/v3 | List AppVersions matching a filter
[**api_app_versions_v3_put**](AppVersionsApi.md#api_app_versions_v3_put) | **PUT** /api/app-versions/v3 | Upsert a AppVersion
[**api_app_versions_v3_selector_get**](AppVersionsApi.md#api_app_versions_v3_selector_get) | **GET** /api/app-versions/v3/{selector} | Get an individual AppVersion
[**api_app_versions_v3_selector_patch**](AppVersionsApi.md#api_app_versions_v3_selector_patch) | **PATCH** /api/app-versions/v3/{selector} | Edit an individual AppVersion


# **api_app_versions_procedures_v3_changelog_get**
> SherlockAppVersionV3ChangelogResponse api_app_versions_procedures_v3_changelog_get(child, parent)

Get a changelog between two AppVersions

Get the path through parent references from a child AppVersion (inclusive) to a parent AppVersion (exclusive), if possible. Because parent references point from newer children to older parents, the newer AppVersion should be the child. The result will always exclude the parent.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_app_version_v3_changelog_response import SherlockAppVersionV3ChangelogResponse
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
    api_instance = sherlock_python_client.AppVersionsApi(api_client)
    child = 'child_example' # str | The selector of the newer AppVersion for the changelog
    parent = 'parent_example' # str | The selector of the older AppVersion for the changelog

    try:
        # Get a changelog between two AppVersions
        api_response = api_instance.api_app_versions_procedures_v3_changelog_get(child, parent)
        print("The response of AppVersionsApi->api_app_versions_procedures_v3_changelog_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling AppVersionsApi->api_app_versions_procedures_v3_changelog_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **child** | **str**| The selector of the newer AppVersion for the changelog | 
 **parent** | **str**| The selector of the older AppVersion for the changelog | 

### Return type

[**SherlockAppVersionV3ChangelogResponse**](SherlockAppVersionV3ChangelogResponse.md)

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

# **api_app_versions_v3_get**
> List[SherlockAppVersionV3] api_app_versions_v3_get(app_version=app_version, authored_by=authored_by, chart=chart, created_at=created_at, description=description, git_branch=git_branch, git_commit=git_commit, id=id, parent_app_version=parent_app_version, updated_at=updated_at, limit=limit, offset=offset)

List AppVersions matching a filter

List AppVersions matching a filter.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_app_version_v3 import SherlockAppVersionV3
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
    api_instance = sherlock_python_client.AppVersionsApi(api_client)
    app_version = 'app_version_example' # str | Required when creating (optional)
    authored_by = 'authored_by_example' # str |  (optional)
    chart = 'chart_example' # str | Required when creating (optional)
    created_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    description = 'description_example' # str | Generally the Git commit message (optional)
    git_branch = 'git_branch_example' # str |  (optional)
    git_commit = 'git_commit_example' # str |  (optional)
    id = 56 # int |  (optional)
    parent_app_version = 'parent_app_version_example' # str |  (optional)
    updated_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    limit = 56 # int | Control how many AppVersions are returned (default 100) (optional)
    offset = 56 # int | Control the offset for the returned AppVersions (default 0) (optional)

    try:
        # List AppVersions matching a filter
        api_response = api_instance.api_app_versions_v3_get(app_version=app_version, authored_by=authored_by, chart=chart, created_at=created_at, description=description, git_branch=git_branch, git_commit=git_commit, id=id, parent_app_version=parent_app_version, updated_at=updated_at, limit=limit, offset=offset)
        print("The response of AppVersionsApi->api_app_versions_v3_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling AppVersionsApi->api_app_versions_v3_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_version** | **str**| Required when creating | [optional] 
 **authored_by** | **str**|  | [optional] 
 **chart** | **str**| Required when creating | [optional] 
 **created_at** | **datetime**|  | [optional] 
 **description** | **str**| Generally the Git commit message | [optional] 
 **git_branch** | **str**|  | [optional] 
 **git_commit** | **str**|  | [optional] 
 **id** | **int**|  | [optional] 
 **parent_app_version** | **str**|  | [optional] 
 **updated_at** | **datetime**|  | [optional] 
 **limit** | **int**| Control how many AppVersions are returned (default 100) | [optional] 
 **offset** | **int**| Control the offset for the returned AppVersions (default 0) | [optional] 

### Return type

[**List[SherlockAppVersionV3]**](SherlockAppVersionV3.md)

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

# **api_app_versions_v3_put**
> SherlockAppVersionV3 api_app_versions_v3_put(app_version)

Upsert a AppVersion

Upsert a AppVersion.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_app_version_v3 import SherlockAppVersionV3
from sherlock_python_client.models.sherlock_app_version_v3_create import SherlockAppVersionV3Create
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
    api_instance = sherlock_python_client.AppVersionsApi(api_client)
    app_version = sherlock_python_client.SherlockAppVersionV3Create() # SherlockAppVersionV3Create | The AppVersion to upsert

    try:
        # Upsert a AppVersion
        api_response = api_instance.api_app_versions_v3_put(app_version)
        print("The response of AppVersionsApi->api_app_versions_v3_put:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling AppVersionsApi->api_app_versions_v3_put: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_version** | [**SherlockAppVersionV3Create**](SherlockAppVersionV3Create.md)| The AppVersion to upsert | 

### Return type

[**SherlockAppVersionV3**](SherlockAppVersionV3.md)

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

# **api_app_versions_v3_selector_get**
> SherlockAppVersionV3 api_app_versions_v3_selector_get(selector)

Get an individual AppVersion

Get an individual AppVersion.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_app_version_v3 import SherlockAppVersionV3
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
    api_instance = sherlock_python_client.AppVersionsApi(api_client)
    selector = 'selector_example' # str | The selector of the AppVersion, which can be either a numeric ID or chart/version.

    try:
        # Get an individual AppVersion
        api_response = api_instance.api_app_versions_v3_selector_get(selector)
        print("The response of AppVersionsApi->api_app_versions_v3_selector_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling AppVersionsApi->api_app_versions_v3_selector_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the AppVersion, which can be either a numeric ID or chart/version. | 

### Return type

[**SherlockAppVersionV3**](SherlockAppVersionV3.md)

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

# **api_app_versions_v3_selector_patch**
> SherlockAppVersionV3 api_app_versions_v3_selector_patch(selector, app_version)

Edit an individual AppVersion

Edit an individual AppVersion.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_app_version_v3 import SherlockAppVersionV3
from sherlock_python_client.models.sherlock_app_version_v3_edit import SherlockAppVersionV3Edit
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
    api_instance = sherlock_python_client.AppVersionsApi(api_client)
    selector = 'selector_example' # str | The selector of the AppVersion, which can be either a numeric ID or chart/version.
    app_version = sherlock_python_client.SherlockAppVersionV3Edit() # SherlockAppVersionV3Edit | The edits to make to the AppVersion

    try:
        # Edit an individual AppVersion
        api_response = api_instance.api_app_versions_v3_selector_patch(selector, app_version)
        print("The response of AppVersionsApi->api_app_versions_v3_selector_patch:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling AppVersionsApi->api_app_versions_v3_selector_patch: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the AppVersion, which can be either a numeric ID or chart/version. | 
 **app_version** | [**SherlockAppVersionV3Edit**](SherlockAppVersionV3Edit.md)| The edits to make to the AppVersion | 

### Return type

[**SherlockAppVersionV3**](SherlockAppVersionV3.md)

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

