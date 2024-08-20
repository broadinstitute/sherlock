# sherlock_python_client.GithubActionsJobsApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_github_actions_jobs_v3_get**](GithubActionsJobsApi.md#api_github_actions_jobs_v3_get) | **GET** /api/github-actions-jobs/v3 | List GithubActionsJobs matching a filter
[**api_github_actions_jobs_v3_put**](GithubActionsJobsApi.md#api_github_actions_jobs_v3_put) | **PUT** /api/github-actions-jobs/v3 | Upsert GithubActionsJob
[**api_github_actions_jobs_v3_selector_get**](GithubActionsJobsApi.md#api_github_actions_jobs_v3_selector_get) | **GET** /api/github-actions-jobs/v3/{selector} | Get an individual GithubActionsJob


# **api_github_actions_jobs_v3_get**
> List[SherlockGithubActionsJobV3] api_github_actions_jobs_v3_get(created_at=created_at, github_actions_attempt_number=github_actions_attempt_number, github_actions_job_id=github_actions_job_id, github_actions_owner=github_actions_owner, github_actions_repo=github_actions_repo, github_actions_run_id=github_actions_run_id, id=id, job_created_at=job_created_at, job_started_at=job_started_at, job_terminal_at=job_terminal_at, status=status, updated_at=updated_at, limit=limit, offset=offset)

List GithubActionsJobs matching a filter

List GithubActionsJobs matching a filter. Results are ordered by start time, starting at most recent.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_github_actions_job_v3 import SherlockGithubActionsJobV3
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
    api_instance = sherlock_python_client.GithubActionsJobsApi(api_client)
    created_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    github_actions_attempt_number = 56 # int |  (optional)
    github_actions_job_id = 56 # int |  (optional)
    github_actions_owner = 'github_actions_owner_example' # str |  (optional)
    github_actions_repo = 'github_actions_repo_example' # str |  (optional)
    github_actions_run_id = 56 # int |  (optional)
    id = 56 # int |  (optional)
    job_created_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    job_started_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    job_terminal_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    status = 'status_example' # str |  (optional)
    updated_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    limit = 56 # int | Control how many GithubActionsJobs are returned (default 100) (optional)
    offset = 56 # int | Control the offset for the returned GithubActionsJobs (default 0) (optional)

    try:
        # List GithubActionsJobs matching a filter
        api_response = api_instance.api_github_actions_jobs_v3_get(created_at=created_at, github_actions_attempt_number=github_actions_attempt_number, github_actions_job_id=github_actions_job_id, github_actions_owner=github_actions_owner, github_actions_repo=github_actions_repo, github_actions_run_id=github_actions_run_id, id=id, job_created_at=job_created_at, job_started_at=job_started_at, job_terminal_at=job_terminal_at, status=status, updated_at=updated_at, limit=limit, offset=offset)
        print("The response of GithubActionsJobsApi->api_github_actions_jobs_v3_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling GithubActionsJobsApi->api_github_actions_jobs_v3_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **created_at** | **datetime**|  | [optional] 
 **github_actions_attempt_number** | **int**|  | [optional] 
 **github_actions_job_id** | **int**|  | [optional] 
 **github_actions_owner** | **str**|  | [optional] 
 **github_actions_repo** | **str**|  | [optional] 
 **github_actions_run_id** | **int**|  | [optional] 
 **id** | **int**|  | [optional] 
 **job_created_at** | **datetime**|  | [optional] 
 **job_started_at** | **datetime**|  | [optional] 
 **job_terminal_at** | **datetime**|  | [optional] 
 **status** | **str**|  | [optional] 
 **updated_at** | **datetime**|  | [optional] 
 **limit** | **int**| Control how many GithubActionsJobs are returned (default 100) | [optional] 
 **offset** | **int**| Control the offset for the returned GithubActionsJobs (default 0) | [optional] 

### Return type

[**List[SherlockGithubActionsJobV3]**](SherlockGithubActionsJobV3.md)

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

# **api_github_actions_jobs_v3_put**
> SherlockGithubActionsJobV3 api_github_actions_jobs_v3_put(github_actions_job)

Upsert GithubActionsJob

Upsert GithubActionsJob.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_github_actions_job_v3 import SherlockGithubActionsJobV3
from sherlock_python_client.models.sherlock_github_actions_job_v3_create import SherlockGithubActionsJobV3Create
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
    api_instance = sherlock_python_client.GithubActionsJobsApi(api_client)
    github_actions_job = sherlock_python_client.SherlockGithubActionsJobV3Create() # SherlockGithubActionsJobV3Create | The GithubActionsJob to upsert

    try:
        # Upsert GithubActionsJob
        api_response = api_instance.api_github_actions_jobs_v3_put(github_actions_job)
        print("The response of GithubActionsJobsApi->api_github_actions_jobs_v3_put:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling GithubActionsJobsApi->api_github_actions_jobs_v3_put: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **github_actions_job** | [**SherlockGithubActionsJobV3Create**](SherlockGithubActionsJobV3Create.md)| The GithubActionsJob to upsert | 

### Return type

[**SherlockGithubActionsJobV3**](SherlockGithubActionsJobV3.md)

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

# **api_github_actions_jobs_v3_selector_get**
> SherlockGithubActionsJobV3 api_github_actions_jobs_v3_selector_get(selector)

Get an individual GithubActionsJob

Get an individual GithubActionsJob.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_github_actions_job_v3 import SherlockGithubActionsJobV3
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
    api_instance = sherlock_python_client.GithubActionsJobsApi(api_client)
    selector = 'selector_example' # str | The selector of the GithubActionsJob, either Sherlock ID or '{owner}/{repo}/{job ID}'

    try:
        # Get an individual GithubActionsJob
        api_response = api_instance.api_github_actions_jobs_v3_selector_get(selector)
        print("The response of GithubActionsJobsApi->api_github_actions_jobs_v3_selector_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling GithubActionsJobsApi->api_github_actions_jobs_v3_selector_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the GithubActionsJob, either Sherlock ID or &#39;{owner}/{repo}/{job ID}&#39; | 

### Return type

[**SherlockGithubActionsJobV3**](SherlockGithubActionsJobV3.md)

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

