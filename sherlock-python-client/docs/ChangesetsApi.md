# sherlock_python_client.ChangesetsApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_changesets_procedures_v3_apply_post**](ChangesetsApi.md#api_changesets_procedures_v3_apply_post) | **POST** /api/changesets/procedures/v3/apply | Apply previously planned version changes to Chart Releases
[**api_changesets_procedures_v3_chart_release_history_chart_release_get**](ChangesetsApi.md#api_changesets_procedures_v3_chart_release_history_chart_release_get) | **GET** /api/changesets/procedures/v3/chart-release-history/{chart-release} | List applied Changesets for a Chart Release
[**api_changesets_procedures_v3_plan_and_apply_post**](ChangesetsApi.md#api_changesets_procedures_v3_plan_and_apply_post) | **POST** /api/changesets/procedures/v3/plan-and-apply | Plan and apply version changes in one step
[**api_changesets_procedures_v3_plan_post**](ChangesetsApi.md#api_changesets_procedures_v3_plan_post) | **POST** /api/changesets/procedures/v3/plan | Plan--but do not apply--version changes to Chart Releases
[**api_changesets_procedures_v3_version_history_version_type_chart_version_get**](ChangesetsApi.md#api_changesets_procedures_v3_version_history_version_type_chart_version_get) | **GET** /api/changesets/procedures/v3/version-history/{version-type}/{chart}/{version} | List applied Changesets for an App or Chart Version
[**api_changesets_v3_get**](ChangesetsApi.md#api_changesets_v3_get) | **GET** /api/changesets/v3 | List Changesets matching a filter
[**api_changesets_v3_id_get**](ChangesetsApi.md#api_changesets_v3_id_get) | **GET** /api/changesets/v3/{id} | Get an individual Changeset


# **api_changesets_procedures_v3_apply_post**
> List[SherlockChangesetV3] api_changesets_procedures_v3_apply_post(apply_request, verbose_output=verbose_output)

Apply previously planned version changes to Chart Releases

Looks up and applies previously-planned version diffs given by the ID. Other stored plans against the same Chart Releases are marked as superseded.
Multiple Changesets can be specified simply by passing multiple IDs in the list.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_changeset_v3 import SherlockChangesetV3
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
    api_instance = sherlock_python_client.ChangesetsApi(api_client)
    apply_request = ['apply_request_example'] # List[str] | String IDs of the Changesets to apply
    verbose_output = True # bool | If full information about the changesets should be returned, defaults to true. If false, only the IDs will be returned. (optional)

    try:
        # Apply previously planned version changes to Chart Releases
        api_response = api_instance.api_changesets_procedures_v3_apply_post(apply_request, verbose_output=verbose_output)
        print("The response of ChangesetsApi->api_changesets_procedures_v3_apply_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ChangesetsApi->api_changesets_procedures_v3_apply_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apply_request** | [**List[str]**](str.md)| String IDs of the Changesets to apply | 
 **verbose_output** | **bool**| If full information about the changesets should be returned, defaults to true. If false, only the IDs will be returned. | [optional] 

### Return type

[**List[SherlockChangesetV3]**](SherlockChangesetV3.md)

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

# **api_changesets_procedures_v3_chart_release_history_chart_release_get**
> List[SherlockChangesetV3] api_changesets_procedures_v3_chart_release_history_chart_release_get(chart_release, offset=offset, limit=limit)

List applied Changesets for a Chart Release

List existing applied Changesets for a particular Chart Release, ordered by most recently applied.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_changeset_v3 import SherlockChangesetV3
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
    api_instance = sherlock_python_client.ChangesetsApi(api_client)
    chart_release = 'chart_release_example' # str | Selector of the Chart Release to find applied Changesets for
    offset = 56 # int | An optional offset to skip a number of latest Changesets (optional)
    limit = 56 # int | An optional limit to the number of entries returned (optional)

    try:
        # List applied Changesets for a Chart Release
        api_response = api_instance.api_changesets_procedures_v3_chart_release_history_chart_release_get(chart_release, offset=offset, limit=limit)
        print("The response of ChangesetsApi->api_changesets_procedures_v3_chart_release_history_chart_release_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ChangesetsApi->api_changesets_procedures_v3_chart_release_history_chart_release_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chart_release** | **str**| Selector of the Chart Release to find applied Changesets for | 
 **offset** | **int**| An optional offset to skip a number of latest Changesets | [optional] 
 **limit** | **int**| An optional limit to the number of entries returned | [optional] 

### Return type

[**List[SherlockChangesetV3]**](SherlockChangesetV3.md)

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

# **api_changesets_procedures_v3_plan_and_apply_post**
> List[SherlockChangesetV3] api_changesets_procedures_v3_plan_and_apply_post(changeset_plan_request, verbose_output=verbose_output)

Plan and apply version changes in one step

Like calling the plan procedure immediately followed by the apply procedure. See those endpoints for more information.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_changeset_v3 import SherlockChangesetV3
from sherlock_python_client.models.sherlock_changeset_v3_plan_request import SherlockChangesetV3PlanRequest
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
    api_instance = sherlock_python_client.ChangesetsApi(api_client)
    changeset_plan_request = sherlock_python_client.SherlockChangesetV3PlanRequest() # SherlockChangesetV3PlanRequest | Info on what version changes or refreshes to apply.
    verbose_output = True # bool | If full information about the changesets should be returned, defaults to true. If false, only the IDs will be returned. (optional)

    try:
        # Plan and apply version changes in one step
        api_response = api_instance.api_changesets_procedures_v3_plan_and_apply_post(changeset_plan_request, verbose_output=verbose_output)
        print("The response of ChangesetsApi->api_changesets_procedures_v3_plan_and_apply_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ChangesetsApi->api_changesets_procedures_v3_plan_and_apply_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changeset_plan_request** | [**SherlockChangesetV3PlanRequest**](SherlockChangesetV3PlanRequest.md)| Info on what version changes or refreshes to apply. | 
 **verbose_output** | **bool**| If full information about the changesets should be returned, defaults to true. If false, only the IDs will be returned. | [optional] 

### Return type

[**List[SherlockChangesetV3]**](SherlockChangesetV3.md)

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

# **api_changesets_procedures_v3_plan_post**
> List[SherlockChangesetV3] api_changesets_procedures_v3_plan_post(changeset_plan_request, verbose_output=verbose_output)

Plan--but do not apply--version changes to Chart Releases

Refreshes and calculates version diffs for Chart Releases. If there's a diff, the plan is stored and returned so it can be applied later.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_changeset_v3 import SherlockChangesetV3
from sherlock_python_client.models.sherlock_changeset_v3_plan_request import SherlockChangesetV3PlanRequest
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
    api_instance = sherlock_python_client.ChangesetsApi(api_client)
    changeset_plan_request = sherlock_python_client.SherlockChangesetV3PlanRequest() # SherlockChangesetV3PlanRequest | Info on what version changes or refreshes to plan
    verbose_output = True # bool | If full information about the changesets should be returned, defaults to true. If false, only the IDs will be returned. (optional)

    try:
        # Plan--but do not apply--version changes to Chart Releases
        api_response = api_instance.api_changesets_procedures_v3_plan_post(changeset_plan_request, verbose_output=verbose_output)
        print("The response of ChangesetsApi->api_changesets_procedures_v3_plan_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ChangesetsApi->api_changesets_procedures_v3_plan_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changeset_plan_request** | [**SherlockChangesetV3PlanRequest**](SherlockChangesetV3PlanRequest.md)| Info on what version changes or refreshes to plan | 
 **verbose_output** | **bool**| If full information about the changesets should be returned, defaults to true. If false, only the IDs will be returned. | [optional] 

### Return type

[**List[SherlockChangesetV3]**](SherlockChangesetV3.md)

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

# **api_changesets_procedures_v3_version_history_version_type_chart_version_get**
> List[SherlockChangesetV3] api_changesets_procedures_v3_version_history_version_type_chart_version_get(version_type, chart, version)

List applied Changesets for an App or Chart Version

List existing applied Changesets that newly deployed a given App Version or Chart Version, ordered by most recently applied.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_changeset_v3 import SherlockChangesetV3
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
    api_instance = sherlock_python_client.ChangesetsApi(api_client)
    version_type = 'version_type_example' # str | The type of the version, either 'app' or 'chart'
    chart = 'chart_example' # str | The chart the version belongs to
    version = 'version_example' # str | The version to look for

    try:
        # List applied Changesets for an App or Chart Version
        api_response = api_instance.api_changesets_procedures_v3_version_history_version_type_chart_version_get(version_type, chart, version)
        print("The response of ChangesetsApi->api_changesets_procedures_v3_version_history_version_type_chart_version_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ChangesetsApi->api_changesets_procedures_v3_version_history_version_type_chart_version_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version_type** | **str**| The type of the version, either &#39;app&#39; or &#39;chart&#39; | 
 **chart** | **str**| The chart the version belongs to | 
 **version** | **str**| The version to look for | 

### Return type

[**List[SherlockChangesetV3]**](SherlockChangesetV3.md)

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

# **api_changesets_v3_get**
> List[SherlockChangesetV3] api_changesets_v3_get(applied_at=applied_at, applied_by=applied_by, chart_release=chart_release, from_app_version_branch=from_app_version_branch, from_app_version_commit=from_app_version_commit, from_app_version_exact=from_app_version_exact, from_app_version_follow_chart_release=from_app_version_follow_chart_release, from_app_version_reference=from_app_version_reference, from_app_version_resolver=from_app_version_resolver, from_chart_version_exact=from_chart_version_exact, from_chart_version_follow_chart_release=from_chart_version_follow_chart_release, from_chart_version_reference=from_chart_version_reference, from_chart_version_resolver=from_chart_version_resolver, from_helmfile_ref=from_helmfile_ref, from_helmfile_ref_enabled=from_helmfile_ref_enabled, from_resolved_at=from_resolved_at, planned_by=planned_by, superseded_at=superseded_at, to_app_version_branch=to_app_version_branch, to_app_version_commit=to_app_version_commit, to_app_version_exact=to_app_version_exact, to_app_version_follow_chart_release=to_app_version_follow_chart_release, to_app_version_reference=to_app_version_reference, to_app_version_resolver=to_app_version_resolver, to_chart_version_exact=to_chart_version_exact, to_chart_version_follow_chart_release=to_chart_version_follow_chart_release, to_chart_version_reference=to_chart_version_reference, to_chart_version_resolver=to_chart_version_resolver, to_helmfile_ref=to_helmfile_ref, to_helmfile_ref_enabled=to_helmfile_ref_enabled, to_resolved_at=to_resolved_at, id=id, limit=limit, offset=offset)

List Changesets matching a filter

List Changesets matching a filter.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_changeset_v3 import SherlockChangesetV3
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
    api_instance = sherlock_python_client.ChangesetsApi(api_client)
    applied_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    applied_by = 'applied_by_example' # str |  (optional)
    chart_release = 'chart_release_example' # str |  (optional)
    from_app_version_branch = 'from_app_version_branch_example' # str |  (optional)
    from_app_version_commit = 'from_app_version_commit_example' # str |  (optional)
    from_app_version_exact = 'from_app_version_exact_example' # str |  (optional)
    from_app_version_follow_chart_release = 'from_app_version_follow_chart_release_example' # str |  (optional)
    from_app_version_reference = 'from_app_version_reference_example' # str |  (optional)
    from_app_version_resolver = 'from_app_version_resolver_example' # str |  (optional)
    from_chart_version_exact = 'from_chart_version_exact_example' # str |  (optional)
    from_chart_version_follow_chart_release = 'from_chart_version_follow_chart_release_example' # str |  (optional)
    from_chart_version_reference = 'from_chart_version_reference_example' # str |  (optional)
    from_chart_version_resolver = 'from_chart_version_resolver_example' # str |  (optional)
    from_helmfile_ref = 'from_helmfile_ref_example' # str |  (optional)
    from_helmfile_ref_enabled = True # bool |  (optional)
    from_resolved_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    planned_by = 'planned_by_example' # str |  (optional)
    superseded_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    to_app_version_branch = 'to_app_version_branch_example' # str |  (optional)
    to_app_version_commit = 'to_app_version_commit_example' # str |  (optional)
    to_app_version_exact = 'to_app_version_exact_example' # str |  (optional)
    to_app_version_follow_chart_release = 'to_app_version_follow_chart_release_example' # str |  (optional)
    to_app_version_reference = 'to_app_version_reference_example' # str |  (optional)
    to_app_version_resolver = 'to_app_version_resolver_example' # str |  (optional)
    to_chart_version_exact = 'to_chart_version_exact_example' # str |  (optional)
    to_chart_version_follow_chart_release = 'to_chart_version_follow_chart_release_example' # str |  (optional)
    to_chart_version_reference = 'to_chart_version_reference_example' # str |  (optional)
    to_chart_version_resolver = 'to_chart_version_resolver_example' # str |  (optional)
    to_helmfile_ref = 'to_helmfile_ref_example' # str |  (optional)
    to_helmfile_ref_enabled = True # bool |  (optional)
    to_resolved_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    id = [56] # List[int] | Get specific changesets by their IDs, can be passed multiple times and/or be comma-separated (optional)
    limit = 56 # int | Control how many Changesets are returned (default 100), ignored if specific IDs are passed (optional)
    offset = 56 # int | Control the offset for the returned Changesets (default 0), ignored if specific IDs are passed (optional)

    try:
        # List Changesets matching a filter
        api_response = api_instance.api_changesets_v3_get(applied_at=applied_at, applied_by=applied_by, chart_release=chart_release, from_app_version_branch=from_app_version_branch, from_app_version_commit=from_app_version_commit, from_app_version_exact=from_app_version_exact, from_app_version_follow_chart_release=from_app_version_follow_chart_release, from_app_version_reference=from_app_version_reference, from_app_version_resolver=from_app_version_resolver, from_chart_version_exact=from_chart_version_exact, from_chart_version_follow_chart_release=from_chart_version_follow_chart_release, from_chart_version_reference=from_chart_version_reference, from_chart_version_resolver=from_chart_version_resolver, from_helmfile_ref=from_helmfile_ref, from_helmfile_ref_enabled=from_helmfile_ref_enabled, from_resolved_at=from_resolved_at, planned_by=planned_by, superseded_at=superseded_at, to_app_version_branch=to_app_version_branch, to_app_version_commit=to_app_version_commit, to_app_version_exact=to_app_version_exact, to_app_version_follow_chart_release=to_app_version_follow_chart_release, to_app_version_reference=to_app_version_reference, to_app_version_resolver=to_app_version_resolver, to_chart_version_exact=to_chart_version_exact, to_chart_version_follow_chart_release=to_chart_version_follow_chart_release, to_chart_version_reference=to_chart_version_reference, to_chart_version_resolver=to_chart_version_resolver, to_helmfile_ref=to_helmfile_ref, to_helmfile_ref_enabled=to_helmfile_ref_enabled, to_resolved_at=to_resolved_at, id=id, limit=limit, offset=offset)
        print("The response of ChangesetsApi->api_changesets_v3_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ChangesetsApi->api_changesets_v3_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applied_at** | **datetime**|  | [optional] 
 **applied_by** | **str**|  | [optional] 
 **chart_release** | **str**|  | [optional] 
 **from_app_version_branch** | **str**|  | [optional] 
 **from_app_version_commit** | **str**|  | [optional] 
 **from_app_version_exact** | **str**|  | [optional] 
 **from_app_version_follow_chart_release** | **str**|  | [optional] 
 **from_app_version_reference** | **str**|  | [optional] 
 **from_app_version_resolver** | **str**|  | [optional] 
 **from_chart_version_exact** | **str**|  | [optional] 
 **from_chart_version_follow_chart_release** | **str**|  | [optional] 
 **from_chart_version_reference** | **str**|  | [optional] 
 **from_chart_version_resolver** | **str**|  | [optional] 
 **from_helmfile_ref** | **str**|  | [optional] 
 **from_helmfile_ref_enabled** | **bool**|  | [optional] 
 **from_resolved_at** | **datetime**|  | [optional] 
 **planned_by** | **str**|  | [optional] 
 **superseded_at** | **datetime**|  | [optional] 
 **to_app_version_branch** | **str**|  | [optional] 
 **to_app_version_commit** | **str**|  | [optional] 
 **to_app_version_exact** | **str**|  | [optional] 
 **to_app_version_follow_chart_release** | **str**|  | [optional] 
 **to_app_version_reference** | **str**|  | [optional] 
 **to_app_version_resolver** | **str**|  | [optional] 
 **to_chart_version_exact** | **str**|  | [optional] 
 **to_chart_version_follow_chart_release** | **str**|  | [optional] 
 **to_chart_version_reference** | **str**|  | [optional] 
 **to_chart_version_resolver** | **str**|  | [optional] 
 **to_helmfile_ref** | **str**|  | [optional] 
 **to_helmfile_ref_enabled** | **bool**|  | [optional] 
 **to_resolved_at** | **datetime**|  | [optional] 
 **id** | [**List[int]**](int.md)| Get specific changesets by their IDs, can be passed multiple times and/or be comma-separated | [optional] 
 **limit** | **int**| Control how many Changesets are returned (default 100), ignored if specific IDs are passed | [optional] 
 **offset** | **int**| Control the offset for the returned Changesets (default 0), ignored if specific IDs are passed | [optional] 

### Return type

[**List[SherlockChangesetV3]**](SherlockChangesetV3.md)

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

# **api_changesets_v3_id_get**
> SherlockChangesetV3 api_changesets_v3_id_get(id)

Get an individual Changeset

Get an individual Changeset.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_changeset_v3 import SherlockChangesetV3
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
    api_instance = sherlock_python_client.ChangesetsApi(api_client)
    id = 56 # int | The numeric ID of the changeset

    try:
        # Get an individual Changeset
        api_response = api_instance.api_changesets_v3_id_get(id)
        print("The response of ChangesetsApi->api_changesets_v3_id_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ChangesetsApi->api_changesets_v3_id_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| The numeric ID of the changeset | 

### Return type

[**SherlockChangesetV3**](SherlockChangesetV3.md)

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

