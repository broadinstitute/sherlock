# sherlock_python_client.DeployHooksApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_deploy_hooks_github_actions_procedures_v3_test_selector_post**](DeployHooksApi.md#api_deploy_hooks_github_actions_procedures_v3_test_selector_post) | **POST** /api/deploy-hooks/github-actions/procedures/v3/test/{selector} | Test a GithubActionsDeployHook
[**api_deploy_hooks_github_actions_v3_get**](DeployHooksApi.md#api_deploy_hooks_github_actions_v3_get) | **GET** /api/deploy-hooks/github-actions/v3 | List GithubActionsDeployHooks matching a filter
[**api_deploy_hooks_github_actions_v3_post**](DeployHooksApi.md#api_deploy_hooks_github_actions_v3_post) | **POST** /api/deploy-hooks/github-actions/v3 | Create a GithubActionsDeployHook
[**api_deploy_hooks_github_actions_v3_selector_delete**](DeployHooksApi.md#api_deploy_hooks_github_actions_v3_selector_delete) | **DELETE** /api/deploy-hooks/github-actions/v3/{selector} | Delete an individual GithubActionsDeployHook
[**api_deploy_hooks_github_actions_v3_selector_get**](DeployHooksApi.md#api_deploy_hooks_github_actions_v3_selector_get) | **GET** /api/deploy-hooks/github-actions/v3/{selector} | Get an individual GithubActionsDeployHook
[**api_deploy_hooks_github_actions_v3_selector_patch**](DeployHooksApi.md#api_deploy_hooks_github_actions_v3_selector_patch) | **PATCH** /api/deploy-hooks/github-actions/v3/{selector} | Edit an individual GithubActionsDeployHook
[**api_deploy_hooks_slack_procedures_v3_test_selector_post**](DeployHooksApi.md#api_deploy_hooks_slack_procedures_v3_test_selector_post) | **POST** /api/deploy-hooks/slack/procedures/v3/test/{selector} | Test a SlackDeployHook
[**api_deploy_hooks_slack_v3_get**](DeployHooksApi.md#api_deploy_hooks_slack_v3_get) | **GET** /api/deploy-hooks/slack/v3 | List SlackDeployHooks matching a filter
[**api_deploy_hooks_slack_v3_post**](DeployHooksApi.md#api_deploy_hooks_slack_v3_post) | **POST** /api/deploy-hooks/slack/v3 | Create a SlackDeployHook
[**api_deploy_hooks_slack_v3_selector_delete**](DeployHooksApi.md#api_deploy_hooks_slack_v3_selector_delete) | **DELETE** /api/deploy-hooks/slack/v3/{selector} | Delete an individual SlackDeployHook
[**api_deploy_hooks_slack_v3_selector_get**](DeployHooksApi.md#api_deploy_hooks_slack_v3_selector_get) | **GET** /api/deploy-hooks/slack/v3/{selector} | Get an individual SlackDeployHook
[**api_deploy_hooks_slack_v3_selector_patch**](DeployHooksApi.md#api_deploy_hooks_slack_v3_selector_patch) | **PATCH** /api/deploy-hooks/slack/v3/{selector} | Edit an individual SlackDeployHook


# **api_deploy_hooks_github_actions_procedures_v3_test_selector_post**
> SherlockGithubActionsDeployHookTestRunResponse api_deploy_hooks_github_actions_procedures_v3_test_selector_post(selector, request)

Test a GithubActionsDeployHook

Run a GitHub Action to simulate a GithubActionsDeployHook

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_github_actions_deploy_hook_test_run_request import SherlockGithubActionsDeployHookTestRunRequest
from sherlock_python_client.models.sherlock_github_actions_deploy_hook_test_run_response import SherlockGithubActionsDeployHookTestRunResponse
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
    api_instance = sherlock_python_client.DeployHooksApi(api_client)
    selector = 'selector_example' # str | The ID of the GithubActionsDeployHook
    request = sherlock_python_client.SherlockGithubActionsDeployHookTestRunRequest() # SherlockGithubActionsDeployHookTestRunRequest | Whether to fully execute the hook (JSON body helps with CSRF protection)

    try:
        # Test a GithubActionsDeployHook
        api_response = api_instance.api_deploy_hooks_github_actions_procedures_v3_test_selector_post(selector, request)
        print("The response of DeployHooksApi->api_deploy_hooks_github_actions_procedures_v3_test_selector_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeployHooksApi->api_deploy_hooks_github_actions_procedures_v3_test_selector_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The ID of the GithubActionsDeployHook | 
 **request** | [**SherlockGithubActionsDeployHookTestRunRequest**](SherlockGithubActionsDeployHookTestRunRequest.md)| Whether to fully execute the hook (JSON body helps with CSRF protection) | 

### Return type

[**SherlockGithubActionsDeployHookTestRunResponse**](SherlockGithubActionsDeployHookTestRunResponse.md)

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

# **api_deploy_hooks_github_actions_v3_get**
> List[SherlockGithubActionsDeployHookV3] api_deploy_hooks_github_actions_v3_get(created_at=created_at, github_actions_default_ref=github_actions_default_ref, github_actions_owner=github_actions_owner, github_actions_ref_behavior=github_actions_ref_behavior, github_actions_repo=github_actions_repo, github_actions_workflow_path=github_actions_workflow_path, id=id, on_chart_release=on_chart_release, on_environment=on_environment, on_failure=on_failure, on_success=on_success, updated_at=updated_at, limit=limit, offset=offset)

List GithubActionsDeployHooks matching a filter

List GithubActionsDeployHooks matching a filter.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_github_actions_deploy_hook_v3 import SherlockGithubActionsDeployHookV3
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
    api_instance = sherlock_python_client.DeployHooksApi(api_client)
    created_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    github_actions_default_ref = 'github_actions_default_ref_example' # str |  (optional)
    github_actions_owner = 'github_actions_owner_example' # str |  (optional)
    github_actions_ref_behavior = always-use-default-ref # str | This field determines what git ref the workflow will be run on. The default of always-use-default-ref always uses the default ref; use-app-version-as-ref will use the app version (when available) as the ref, useful when versions are always commit hashes or tags; use-app-version-commit-as-ref will use the app version's commit (when available) as the ref, useful when the repo is configured to fully report app versions to Sherlock. (optional) (default to always-use-default-ref)
    github_actions_repo = 'github_actions_repo_example' # str |  (optional)
    github_actions_workflow_path = 'github_actions_workflow_path_example' # str |  (optional)
    id = 56 # int |  (optional)
    on_chart_release = 'on_chart_release_example' # str |  (optional)
    on_environment = 'on_environment_example' # str |  (optional)
    on_failure = True # bool |  (optional)
    on_success = True # bool |  (optional)
    updated_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    limit = 56 # int | Control how many GithubActionsDeployHooks are returned (default 100) (optional)
    offset = 56 # int | Control the offset for the returned GithubActionsDeployHooks (default 0) (optional)

    try:
        # List GithubActionsDeployHooks matching a filter
        api_response = api_instance.api_deploy_hooks_github_actions_v3_get(created_at=created_at, github_actions_default_ref=github_actions_default_ref, github_actions_owner=github_actions_owner, github_actions_ref_behavior=github_actions_ref_behavior, github_actions_repo=github_actions_repo, github_actions_workflow_path=github_actions_workflow_path, id=id, on_chart_release=on_chart_release, on_environment=on_environment, on_failure=on_failure, on_success=on_success, updated_at=updated_at, limit=limit, offset=offset)
        print("The response of DeployHooksApi->api_deploy_hooks_github_actions_v3_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeployHooksApi->api_deploy_hooks_github_actions_v3_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **created_at** | **datetime**|  | [optional] 
 **github_actions_default_ref** | **str**|  | [optional] 
 **github_actions_owner** | **str**|  | [optional] 
 **github_actions_ref_behavior** | **str**| This field determines what git ref the workflow will be run on. The default of always-use-default-ref always uses the default ref; use-app-version-as-ref will use the app version (when available) as the ref, useful when versions are always commit hashes or tags; use-app-version-commit-as-ref will use the app version&#39;s commit (when available) as the ref, useful when the repo is configured to fully report app versions to Sherlock. | [optional] [default to always-use-default-ref]
 **github_actions_repo** | **str**|  | [optional] 
 **github_actions_workflow_path** | **str**|  | [optional] 
 **id** | **int**|  | [optional] 
 **on_chart_release** | **str**|  | [optional] 
 **on_environment** | **str**|  | [optional] 
 **on_failure** | **bool**|  | [optional] 
 **on_success** | **bool**|  | [optional] 
 **updated_at** | **datetime**|  | [optional] 
 **limit** | **int**| Control how many GithubActionsDeployHooks are returned (default 100) | [optional] 
 **offset** | **int**| Control the offset for the returned GithubActionsDeployHooks (default 0) | [optional] 

### Return type

[**List[SherlockGithubActionsDeployHookV3]**](SherlockGithubActionsDeployHookV3.md)

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

# **api_deploy_hooks_github_actions_v3_post**
> SherlockGithubActionsDeployHookV3 api_deploy_hooks_github_actions_v3_post(github_actions_deploy_hook)

Create a GithubActionsDeployHook

Create a GithubActionsDeployHook.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_github_actions_deploy_hook_v3 import SherlockGithubActionsDeployHookV3
from sherlock_python_client.models.sherlock_github_actions_deploy_hook_v3_create import SherlockGithubActionsDeployHookV3Create
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
    api_instance = sherlock_python_client.DeployHooksApi(api_client)
    github_actions_deploy_hook = sherlock_python_client.SherlockGithubActionsDeployHookV3Create() # SherlockGithubActionsDeployHookV3Create | The GithubActionsDeployHook to create

    try:
        # Create a GithubActionsDeployHook
        api_response = api_instance.api_deploy_hooks_github_actions_v3_post(github_actions_deploy_hook)
        print("The response of DeployHooksApi->api_deploy_hooks_github_actions_v3_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeployHooksApi->api_deploy_hooks_github_actions_v3_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **github_actions_deploy_hook** | [**SherlockGithubActionsDeployHookV3Create**](SherlockGithubActionsDeployHookV3Create.md)| The GithubActionsDeployHook to create | 

### Return type

[**SherlockGithubActionsDeployHookV3**](SherlockGithubActionsDeployHookV3.md)

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

# **api_deploy_hooks_github_actions_v3_selector_delete**
> SherlockGithubActionsDeployHookV3 api_deploy_hooks_github_actions_v3_selector_delete(selector)

Delete an individual GithubActionsDeployHook

Delete an individual GithubActionsDeployHook by its ID.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_github_actions_deploy_hook_v3 import SherlockGithubActionsDeployHookV3
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
    api_instance = sherlock_python_client.DeployHooksApi(api_client)
    selector = 'selector_example' # str | The ID of the GithubActionsDeployHook

    try:
        # Delete an individual GithubActionsDeployHook
        api_response = api_instance.api_deploy_hooks_github_actions_v3_selector_delete(selector)
        print("The response of DeployHooksApi->api_deploy_hooks_github_actions_v3_selector_delete:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeployHooksApi->api_deploy_hooks_github_actions_v3_selector_delete: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The ID of the GithubActionsDeployHook | 

### Return type

[**SherlockGithubActionsDeployHookV3**](SherlockGithubActionsDeployHookV3.md)

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

# **api_deploy_hooks_github_actions_v3_selector_get**
> SherlockGithubActionsDeployHookV3 api_deploy_hooks_github_actions_v3_selector_get(selector)

Get an individual GithubActionsDeployHook

Get an individual GithubActionsDeployHook by its ID.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_github_actions_deploy_hook_v3 import SherlockGithubActionsDeployHookV3
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
    api_instance = sherlock_python_client.DeployHooksApi(api_client)
    selector = 'selector_example' # str | The ID of the GithubActionsDeployHook

    try:
        # Get an individual GithubActionsDeployHook
        api_response = api_instance.api_deploy_hooks_github_actions_v3_selector_get(selector)
        print("The response of DeployHooksApi->api_deploy_hooks_github_actions_v3_selector_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeployHooksApi->api_deploy_hooks_github_actions_v3_selector_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The ID of the GithubActionsDeployHook | 

### Return type

[**SherlockGithubActionsDeployHookV3**](SherlockGithubActionsDeployHookV3.md)

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

# **api_deploy_hooks_github_actions_v3_selector_patch**
> SherlockGithubActionsDeployHookV3 api_deploy_hooks_github_actions_v3_selector_patch(selector, github_actions_deploy_hook)

Edit an individual GithubActionsDeployHook

Edit an individual GithubActionsDeployHook by its ID.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_github_actions_deploy_hook_v3 import SherlockGithubActionsDeployHookV3
from sherlock_python_client.models.sherlock_github_actions_deploy_hook_v3_edit import SherlockGithubActionsDeployHookV3Edit
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
    api_instance = sherlock_python_client.DeployHooksApi(api_client)
    selector = 'selector_example' # str | The ID of the GithubActionsDeployHook to edit
    github_actions_deploy_hook = sherlock_python_client.SherlockGithubActionsDeployHookV3Edit() # SherlockGithubActionsDeployHookV3Edit | The edits to make to the GithubActionsDeployHook

    try:
        # Edit an individual GithubActionsDeployHook
        api_response = api_instance.api_deploy_hooks_github_actions_v3_selector_patch(selector, github_actions_deploy_hook)
        print("The response of DeployHooksApi->api_deploy_hooks_github_actions_v3_selector_patch:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeployHooksApi->api_deploy_hooks_github_actions_v3_selector_patch: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The ID of the GithubActionsDeployHook to edit | 
 **github_actions_deploy_hook** | [**SherlockGithubActionsDeployHookV3Edit**](SherlockGithubActionsDeployHookV3Edit.md)| The edits to make to the GithubActionsDeployHook | 

### Return type

[**SherlockGithubActionsDeployHookV3**](SherlockGithubActionsDeployHookV3.md)

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

# **api_deploy_hooks_slack_procedures_v3_test_selector_post**
> SherlockSlackDeployHookTestRunResponse api_deploy_hooks_slack_procedures_v3_test_selector_post(selector, request)

Test a SlackDeployHook

Send a Slack message to simulate a SlackDeployHook

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_slack_deploy_hook_test_run_request import SherlockSlackDeployHookTestRunRequest
from sherlock_python_client.models.sherlock_slack_deploy_hook_test_run_response import SherlockSlackDeployHookTestRunResponse
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
    api_instance = sherlock_python_client.DeployHooksApi(api_client)
    selector = 'selector_example' # str | The ID of the SlackDeployHook to test
    request = sherlock_python_client.SherlockSlackDeployHookTestRunRequest() # SherlockSlackDeployHookTestRunRequest | Whether to fully execute the hook (JSON body helps with CSRF protection)

    try:
        # Test a SlackDeployHook
        api_response = api_instance.api_deploy_hooks_slack_procedures_v3_test_selector_post(selector, request)
        print("The response of DeployHooksApi->api_deploy_hooks_slack_procedures_v3_test_selector_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeployHooksApi->api_deploy_hooks_slack_procedures_v3_test_selector_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The ID of the SlackDeployHook to test | 
 **request** | [**SherlockSlackDeployHookTestRunRequest**](SherlockSlackDeployHookTestRunRequest.md)| Whether to fully execute the hook (JSON body helps with CSRF protection) | 

### Return type

[**SherlockSlackDeployHookTestRunResponse**](SherlockSlackDeployHookTestRunResponse.md)

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

# **api_deploy_hooks_slack_v3_get**
> List[SherlockSlackDeployHookV3] api_deploy_hooks_slack_v3_get(created_at=created_at, id=id, mention_people=mention_people, on_chart_release=on_chart_release, on_environment=on_environment, on_failure=on_failure, on_success=on_success, slack_channel=slack_channel, updated_at=updated_at, limit=limit, offset=offset)

List SlackDeployHooks matching a filter

List SlackDeployHooks matching a filter.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_slack_deploy_hook_v3 import SherlockSlackDeployHookV3
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
    api_instance = sherlock_python_client.DeployHooksApi(api_client)
    created_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    id = 56 # int |  (optional)
    mention_people = True # bool |  (optional)
    on_chart_release = 'on_chart_release_example' # str |  (optional)
    on_environment = 'on_environment_example' # str |  (optional)
    on_failure = True # bool |  (optional)
    on_success = True # bool |  (optional)
    slack_channel = 'slack_channel_example' # str |  (optional)
    updated_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    limit = 56 # int | Control how many SlackDeployHooks are returned (default 100) (optional)
    offset = 56 # int | Control the offset for the returned SlackDeployHooks (default 0) (optional)

    try:
        # List SlackDeployHooks matching a filter
        api_response = api_instance.api_deploy_hooks_slack_v3_get(created_at=created_at, id=id, mention_people=mention_people, on_chart_release=on_chart_release, on_environment=on_environment, on_failure=on_failure, on_success=on_success, slack_channel=slack_channel, updated_at=updated_at, limit=limit, offset=offset)
        print("The response of DeployHooksApi->api_deploy_hooks_slack_v3_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeployHooksApi->api_deploy_hooks_slack_v3_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **created_at** | **datetime**|  | [optional] 
 **id** | **int**|  | [optional] 
 **mention_people** | **bool**|  | [optional] 
 **on_chart_release** | **str**|  | [optional] 
 **on_environment** | **str**|  | [optional] 
 **on_failure** | **bool**|  | [optional] 
 **on_success** | **bool**|  | [optional] 
 **slack_channel** | **str**|  | [optional] 
 **updated_at** | **datetime**|  | [optional] 
 **limit** | **int**| Control how many SlackDeployHooks are returned (default 100) | [optional] 
 **offset** | **int**| Control the offset for the returned SlackDeployHooks (default 0) | [optional] 

### Return type

[**List[SherlockSlackDeployHookV3]**](SherlockSlackDeployHookV3.md)

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

# **api_deploy_hooks_slack_v3_post**
> SherlockSlackDeployHookV3 api_deploy_hooks_slack_v3_post(slack_deploy_hook)

Create a SlackDeployHook

Create a SlackDeployHook.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_slack_deploy_hook_v3 import SherlockSlackDeployHookV3
from sherlock_python_client.models.sherlock_slack_deploy_hook_v3_create import SherlockSlackDeployHookV3Create
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
    api_instance = sherlock_python_client.DeployHooksApi(api_client)
    slack_deploy_hook = sherlock_python_client.SherlockSlackDeployHookV3Create() # SherlockSlackDeployHookV3Create | The SlackDeployHook to create

    try:
        # Create a SlackDeployHook
        api_response = api_instance.api_deploy_hooks_slack_v3_post(slack_deploy_hook)
        print("The response of DeployHooksApi->api_deploy_hooks_slack_v3_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeployHooksApi->api_deploy_hooks_slack_v3_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **slack_deploy_hook** | [**SherlockSlackDeployHookV3Create**](SherlockSlackDeployHookV3Create.md)| The SlackDeployHook to create | 

### Return type

[**SherlockSlackDeployHookV3**](SherlockSlackDeployHookV3.md)

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

# **api_deploy_hooks_slack_v3_selector_delete**
> SherlockSlackDeployHookV3 api_deploy_hooks_slack_v3_selector_delete(selector)

Delete an individual SlackDeployHook

Delete an individual SlackDeployHook by its ID.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_slack_deploy_hook_v3 import SherlockSlackDeployHookV3
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
    api_instance = sherlock_python_client.DeployHooksApi(api_client)
    selector = 'selector_example' # str | The ID of the SlackDeployHook

    try:
        # Delete an individual SlackDeployHook
        api_response = api_instance.api_deploy_hooks_slack_v3_selector_delete(selector)
        print("The response of DeployHooksApi->api_deploy_hooks_slack_v3_selector_delete:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeployHooksApi->api_deploy_hooks_slack_v3_selector_delete: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The ID of the SlackDeployHook | 

### Return type

[**SherlockSlackDeployHookV3**](SherlockSlackDeployHookV3.md)

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

# **api_deploy_hooks_slack_v3_selector_get**
> SherlockSlackDeployHookV3 api_deploy_hooks_slack_v3_selector_get(selector)

Get an individual SlackDeployHook

Get an individual SlackDeployHook by its ID.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_slack_deploy_hook_v3 import SherlockSlackDeployHookV3
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
    api_instance = sherlock_python_client.DeployHooksApi(api_client)
    selector = 'selector_example' # str | The ID of the SlackDeployHook

    try:
        # Get an individual SlackDeployHook
        api_response = api_instance.api_deploy_hooks_slack_v3_selector_get(selector)
        print("The response of DeployHooksApi->api_deploy_hooks_slack_v3_selector_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeployHooksApi->api_deploy_hooks_slack_v3_selector_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The ID of the SlackDeployHook | 

### Return type

[**SherlockSlackDeployHookV3**](SherlockSlackDeployHookV3.md)

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

# **api_deploy_hooks_slack_v3_selector_patch**
> SherlockSlackDeployHookV3 api_deploy_hooks_slack_v3_selector_patch(selector, slack_deploy_hook)

Edit an individual SlackDeployHook

Edit an individual SlackDeployHook by its ID.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_slack_deploy_hook_v3 import SherlockSlackDeployHookV3
from sherlock_python_client.models.sherlock_slack_deploy_hook_v3_edit import SherlockSlackDeployHookV3Edit
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
    api_instance = sherlock_python_client.DeployHooksApi(api_client)
    selector = 'selector_example' # str | The ID of the SlackDeployHook to edit
    slack_deploy_hook = sherlock_python_client.SherlockSlackDeployHookV3Edit() # SherlockSlackDeployHookV3Edit | The edits to make to the SlackDeployHook

    try:
        # Edit an individual SlackDeployHook
        api_response = api_instance.api_deploy_hooks_slack_v3_selector_patch(selector, slack_deploy_hook)
        print("The response of DeployHooksApi->api_deploy_hooks_slack_v3_selector_patch:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeployHooksApi->api_deploy_hooks_slack_v3_selector_patch: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The ID of the SlackDeployHook to edit | 
 **slack_deploy_hook** | [**SherlockSlackDeployHookV3Edit**](SherlockSlackDeployHookV3Edit.md)| The edits to make to the SlackDeployHook | 

### Return type

[**SherlockSlackDeployHookV3**](SherlockSlackDeployHookV3.md)

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

