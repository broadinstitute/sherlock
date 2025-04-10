# sherlock_python_client.CiRunsApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_ci_runs_procedures_v3_github_info_get**](CiRunsApi.md#api_ci_runs_procedures_v3_github_info_get) | **GET** /api/ci-runs/procedures/v3/github-info | List GitHub info gleaned from CiRuns
[**api_ci_runs_v3_get**](CiRunsApi.md#api_ci_runs_v3_get) | **GET** /api/ci-runs/v3 | List CiRuns matching a filter
[**api_ci_runs_v3_put**](CiRunsApi.md#api_ci_runs_v3_put) | **PUT** /api/ci-runs/v3 | Create or update a CiRun
[**api_ci_runs_v3_selector_get**](CiRunsApi.md#api_ci_runs_v3_selector_get) | **GET** /api/ci-runs/v3/{selector} | Get a CiRun, including CiIdentifiers for related resources


# **api_ci_runs_procedures_v3_github_info_get**
> Dict[str, Dict[str, List[str]]] api_ci_runs_procedures_v3_github_info_get()

List GitHub info gleaned from CiRuns

List info about GitHub repos and their workflow files as determined by CiRuns from the past 90 days.
This is a useful proxy for figuring out what repos Sherlock probably has access to: workflows listed
here can probably successfully called by a GitHub Actions deploy hook.

### Example


```python
import sherlock_python_client
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
    api_instance = sherlock_python_client.CiRunsApi(api_client)

    try:
        # List GitHub info gleaned from CiRuns
        api_response = api_instance.api_ci_runs_procedures_v3_github_info_get()
        print("The response of CiRunsApi->api_ci_runs_procedures_v3_github_info_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CiRunsApi->api_ci_runs_procedures_v3_github_info_get: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

**Dict[str, Dict[str, List[str]]]**

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

# **api_ci_runs_v3_get**
> List[SherlockCiRunV3] api_ci_runs_v3_get(argo_workflows_name=argo_workflows_name, argo_workflows_namespace=argo_workflows_namespace, argo_workflows_template=argo_workflows_template, created_at=created_at, github_actions_attempt_number=github_actions_attempt_number, github_actions_owner=github_actions_owner, github_actions_repo=github_actions_repo, github_actions_run_id=github_actions_run_id, github_actions_workflow_path=github_actions_workflow_path, id=id, notify_slack_channels_upon_failure=notify_slack_channels_upon_failure, notify_slack_channels_upon_retry=notify_slack_channels_upon_retry, notify_slack_channels_upon_success=notify_slack_channels_upon_success, notify_slack_custom_icon=notify_slack_custom_icon, platform=platform, resource_status=resource_status, started_at=started_at, status=status, terminal_at=terminal_at, termination_hooks_dispatched_at=termination_hooks_dispatched_at, updated_at=updated_at, limit=limit, offset=offset)

List CiRuns matching a filter

List CiRuns matching a filter. The CiRuns would have to re-queried directly to load any related resources.
Results are ordered by start time, starting at most recent.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_ci_run_v3 import SherlockCiRunV3
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
    api_instance = sherlock_python_client.CiRunsApi(api_client)
    argo_workflows_name = 'argo_workflows_name_example' # str |  (optional)
    argo_workflows_namespace = 'argo_workflows_namespace_example' # str |  (optional)
    argo_workflows_template = 'argo_workflows_template_example' # str |  (optional)
    created_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    github_actions_attempt_number = 56 # int |  (optional)
    github_actions_owner = 'github_actions_owner_example' # str |  (optional)
    github_actions_repo = 'github_actions_repo_example' # str |  (optional)
    github_actions_run_id = 56 # int |  (optional)
    github_actions_workflow_path = 'github_actions_workflow_path_example' # str |  (optional)
    id = 56 # int |  (optional)
    notify_slack_channels_upon_failure = ['notify_slack_channels_upon_failure_example'] # List[str] | Slack channels to notify if this CiRun fails. This field is always appended to when mutated. (optional)
    notify_slack_channels_upon_retry = ['notify_slack_channels_upon_retry_example'] # List[str] | Slack channels to notify if this CiRun is retried. This field is always appended to when mutated. It will de-dupe with the other notify fields. (optional)
    notify_slack_channels_upon_success = ['notify_slack_channels_upon_success_example'] # List[str] | Slack channels to notify if this CiRun succeeds. This field is always appended to when mutated. (optional)
    notify_slack_custom_icon = 'notify_slack_custom_icon_example' # str | Icon to use for success or failure Slack notifications. Can be given either as a URL to an image or as a Slack emoji (using colon shortcodes, like :smiley:). An empty string is ignored to facilitate calling from GitHub Actions (where it's easier to pass an empty string than not send the field at all). (optional)
    platform = 'platform_example' # str |  (optional)
    resource_status = 'resource_status_example' # str | Available only when querying a CiRun via a CiIdentifier, indicates the status of the run for that resource (optional)
    started_at = 'started_at_example' # str |  (optional)
    status = 'status_example' # str |  (optional)
    terminal_at = 'terminal_at_example' # str |  (optional)
    termination_hooks_dispatched_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    updated_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    limit = 56 # int | Control how many CiRuns are returned (default 100) (optional)
    offset = 56 # int | Control the offset for the returned CiRuns (default 0) (optional)

    try:
        # List CiRuns matching a filter
        api_response = api_instance.api_ci_runs_v3_get(argo_workflows_name=argo_workflows_name, argo_workflows_namespace=argo_workflows_namespace, argo_workflows_template=argo_workflows_template, created_at=created_at, github_actions_attempt_number=github_actions_attempt_number, github_actions_owner=github_actions_owner, github_actions_repo=github_actions_repo, github_actions_run_id=github_actions_run_id, github_actions_workflow_path=github_actions_workflow_path, id=id, notify_slack_channels_upon_failure=notify_slack_channels_upon_failure, notify_slack_channels_upon_retry=notify_slack_channels_upon_retry, notify_slack_channels_upon_success=notify_slack_channels_upon_success, notify_slack_custom_icon=notify_slack_custom_icon, platform=platform, resource_status=resource_status, started_at=started_at, status=status, terminal_at=terminal_at, termination_hooks_dispatched_at=termination_hooks_dispatched_at, updated_at=updated_at, limit=limit, offset=offset)
        print("The response of CiRunsApi->api_ci_runs_v3_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CiRunsApi->api_ci_runs_v3_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **argo_workflows_name** | **str**|  | [optional] 
 **argo_workflows_namespace** | **str**|  | [optional] 
 **argo_workflows_template** | **str**|  | [optional] 
 **created_at** | **datetime**|  | [optional] 
 **github_actions_attempt_number** | **int**|  | [optional] 
 **github_actions_owner** | **str**|  | [optional] 
 **github_actions_repo** | **str**|  | [optional] 
 **github_actions_run_id** | **int**|  | [optional] 
 **github_actions_workflow_path** | **str**|  | [optional] 
 **id** | **int**|  | [optional] 
 **notify_slack_channels_upon_failure** | [**List[str]**](str.md)| Slack channels to notify if this CiRun fails. This field is always appended to when mutated. | [optional] 
 **notify_slack_channels_upon_retry** | [**List[str]**](str.md)| Slack channels to notify if this CiRun is retried. This field is always appended to when mutated. It will de-dupe with the other notify fields. | [optional] 
 **notify_slack_channels_upon_success** | [**List[str]**](str.md)| Slack channels to notify if this CiRun succeeds. This field is always appended to when mutated. | [optional] 
 **notify_slack_custom_icon** | **str**| Icon to use for success or failure Slack notifications. Can be given either as a URL to an image or as a Slack emoji (using colon shortcodes, like :smiley:). An empty string is ignored to facilitate calling from GitHub Actions (where it&#39;s easier to pass an empty string than not send the field at all). | [optional] 
 **platform** | **str**|  | [optional] 
 **resource_status** | **str**| Available only when querying a CiRun via a CiIdentifier, indicates the status of the run for that resource | [optional] 
 **started_at** | **str**|  | [optional] 
 **status** | **str**|  | [optional] 
 **terminal_at** | **str**|  | [optional] 
 **termination_hooks_dispatched_at** | **datetime**|  | [optional] 
 **updated_at** | **datetime**|  | [optional] 
 **limit** | **int**| Control how many CiRuns are returned (default 100) | [optional] 
 **offset** | **int**| Control the offset for the returned CiRuns (default 0) | [optional] 

### Return type

[**List[SherlockCiRunV3]**](SherlockCiRunV3.md)

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

# **api_ci_runs_v3_put**
> SherlockCiRunV3 api_ci_runs_v3_put(ci_run)

Create or update a CiRun

Create or update a CiRun with timing, status, and related resource information. This endpoint is idempotent.
The fields for clusters, charts, chart releases, environments, etc. all accept selectors, and they will
be smart about "spreading" to indirect relations. More info is available on the CiRunV3Upsert data type,
but the gist is that specifying a changeset implies its chart release (and optionally app/chart versions),
specifying or implying a chart release implies its environment/cluster, and specifying an environment/cluster
implies all chart releases they contain.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_ci_run_v3 import SherlockCiRunV3
from sherlock_python_client.models.sherlock_ci_run_v3_upsert import SherlockCiRunV3Upsert
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
    api_instance = sherlock_python_client.CiRunsApi(api_client)
    ci_run = sherlock_python_client.SherlockCiRunV3Upsert() # SherlockCiRunV3Upsert | The CiRun to upsert

    try:
        # Create or update a CiRun
        api_response = api_instance.api_ci_runs_v3_put(ci_run)
        print("The response of CiRunsApi->api_ci_runs_v3_put:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CiRunsApi->api_ci_runs_v3_put: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ci_run** | [**SherlockCiRunV3Upsert**](SherlockCiRunV3Upsert.md)| The CiRun to upsert | 

### Return type

[**SherlockCiRunV3**](SherlockCiRunV3.md)

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

# **api_ci_runs_v3_selector_get**
> SherlockCiRunV3 api_ci_runs_v3_selector_get(selector)

Get a CiRun, including CiIdentifiers for related resources

Get a CiRun, including CiIdentifiers representing related resources or resources it affected.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_ci_run_v3 import SherlockCiRunV3
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
    api_instance = sherlock_python_client.CiRunsApi(api_client)
    selector = 'selector_example' # str | The selector of the CiRun, which can be either its numeric ID, 'github-actions/{owner}/{repo}/{run ID}/{attempt}', or 'argo-workflows/{namespace}/{name}'

    try:
        # Get a CiRun, including CiIdentifiers for related resources
        api_response = api_instance.api_ci_runs_v3_selector_get(selector)
        print("The response of CiRunsApi->api_ci_runs_v3_selector_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CiRunsApi->api_ci_runs_v3_selector_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the CiRun, which can be either its numeric ID, &#39;github-actions/{owner}/{repo}/{run ID}/{attempt}&#39;, or &#39;argo-workflows/{namespace}/{name}&#39; | 

### Return type

[**SherlockCiRunV3**](SherlockCiRunV3.md)

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

