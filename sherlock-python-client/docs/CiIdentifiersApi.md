# sherlock_python_client.CiIdentifiersApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_ci_identifiers_v3_get**](CiIdentifiersApi.md#api_ci_identifiers_v3_get) | **GET** /api/ci-identifiers/v3 | List CiIdentifiers matching a filter
[**api_ci_identifiers_v3_selector_get**](CiIdentifiersApi.md#api_ci_identifiers_v3_selector_get) | **GET** /api/ci-identifiers/v3/{selector} | Get CiRuns for a resource by its CiIdentifier


# **api_ci_identifiers_v3_get**
> List[SherlockCiIdentifierV3] api_ci_identifiers_v3_get(created_at=created_at, id=id, resource_id=resource_id, resource_status=resource_status, resource_type=resource_type, updated_at=updated_at, limit=limit, offset=offset)

List CiIdentifiers matching a filter

List CiIdentifiers matching a filter. The CiRuns would have to re-queried directly to load the CiRuns. This is mainly helpful for debugging and directly querying the existence of a CiIdentifier. Results are ordered by creation date, starting at most recent.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_ci_identifier_v3 import SherlockCiIdentifierV3
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
    api_instance = sherlock_python_client.CiIdentifiersApi(api_client)
    created_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    id = 56 # int |  (optional)
    resource_id = 56 # int |  (optional)
    resource_status = 'resource_status_example' # str | Available only when querying a CiIdentifier via a CiRun, indicates the status of the run for that resource (optional)
    resource_type = 'resource_type_example' # str |  (optional)
    updated_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    limit = 56 # int | Control how many CiIdentifiers are returned (default 100) (optional)
    offset = 56 # int | Control the offset for the returned CiIdentifiers (default 0) (optional)

    try:
        # List CiIdentifiers matching a filter
        api_response = api_instance.api_ci_identifiers_v3_get(created_at=created_at, id=id, resource_id=resource_id, resource_status=resource_status, resource_type=resource_type, updated_at=updated_at, limit=limit, offset=offset)
        print("The response of CiIdentifiersApi->api_ci_identifiers_v3_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CiIdentifiersApi->api_ci_identifiers_v3_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **created_at** | **datetime**|  | [optional] 
 **id** | **int**|  | [optional] 
 **resource_id** | **int**|  | [optional] 
 **resource_status** | **str**| Available only when querying a CiIdentifier via a CiRun, indicates the status of the run for that resource | [optional] 
 **resource_type** | **str**|  | [optional] 
 **updated_at** | **datetime**|  | [optional] 
 **limit** | **int**| Control how many CiIdentifiers are returned (default 100) | [optional] 
 **offset** | **int**| Control the offset for the returned CiIdentifiers (default 0) | [optional] 

### Return type

[**List[SherlockCiIdentifierV3]**](SherlockCiIdentifierV3.md)

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

# **api_ci_identifiers_v3_selector_get**
> SherlockCiIdentifierV3 api_ci_identifiers_v3_selector_get(selector, limit_ci_runs=limit_ci_runs, offset_ci_runs=offset_ci_runs, allow_stub_ci_runs=allow_stub_ci_runs)

Get CiRuns for a resource by its CiIdentifier

Get CiRuns for a resource by its CiIdentifier, which can be referenced by '{type}/{selector...}'.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_ci_identifier_v3 import SherlockCiIdentifierV3
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
    api_instance = sherlock_python_client.CiIdentifiersApi(api_client)
    selector = 'selector_example' # str | The selector of CiIdentifier, which can be referenced either by numeric ID or indirectly by '{type}/{selector...}'
    limit_ci_runs = 56 # int | Control how many CiRuns are returned (default 10) (optional)
    offset_ci_runs = 56 # int | Control the offset for the returned CiRuns (default 0) (optional)
    allow_stub_ci_runs = True # bool | Allow stub CiRuns potentially lacking fields like status or startedAt to be returned (default false) (optional)

    try:
        # Get CiRuns for a resource by its CiIdentifier
        api_response = api_instance.api_ci_identifiers_v3_selector_get(selector, limit_ci_runs=limit_ci_runs, offset_ci_runs=offset_ci_runs, allow_stub_ci_runs=allow_stub_ci_runs)
        print("The response of CiIdentifiersApi->api_ci_identifiers_v3_selector_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CiIdentifiersApi->api_ci_identifiers_v3_selector_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of CiIdentifier, which can be referenced either by numeric ID or indirectly by &#39;{type}/{selector...}&#39; | 
 **limit_ci_runs** | **int**| Control how many CiRuns are returned (default 10) | [optional] 
 **offset_ci_runs** | **int**| Control the offset for the returned CiRuns (default 0) | [optional] 
 **allow_stub_ci_runs** | **bool**| Allow stub CiRuns potentially lacking fields like status or startedAt to be returned (default false) | [optional] 

### Return type

[**SherlockCiIdentifierV3**](SherlockCiIdentifierV3.md)

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

