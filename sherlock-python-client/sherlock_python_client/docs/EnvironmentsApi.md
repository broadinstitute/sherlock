# sherlock_python_client.EnvironmentsApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_environments_v3_get**](EnvironmentsApi.md#api_environments_v3_get) | **GET** /api/environments/v3 | List Environments matching a filter
[**api_environments_v3_post**](EnvironmentsApi.md#api_environments_v3_post) | **POST** /api/environments/v3 | Create a Environment
[**api_environments_v3_selector_delete**](EnvironmentsApi.md#api_environments_v3_selector_delete) | **DELETE** /api/environments/v3/{selector} | Delete an individual Environment
[**api_environments_v3_selector_get**](EnvironmentsApi.md#api_environments_v3_selector_get) | **GET** /api/environments/v3/{selector} | Get an individual Environment
[**api_environments_v3_selector_patch**](EnvironmentsApi.md#api_environments_v3_selector_patch) | **PATCH** /api/environments/v3/{selector} | Edit an individual Environment


# **api_environments_v3_get**
> List[SherlockEnvironmentV3] api_environments_v3_get(auto_populate_chart_releases=auto_populate_chart_releases, base=base, base_domain=base_domain, created_at=created_at, default_cluster=default_cluster, default_namespace=default_namespace, delete_after=delete_after, description=description, enable_janitor=enable_janitor, helmfile_ref=helmfile_ref, id=id, lifecycle=lifecycle, name=name, name_prefixes_domain=name_prefixes_domain, offline=offline, offline_schedule_begin_enabled=offline_schedule_begin_enabled, offline_schedule_begin_time=offline_schedule_begin_time, offline_schedule_end_enabled=offline_schedule_end_enabled, offline_schedule_end_time=offline_schedule_end_time, offline_schedule_end_weekends=offline_schedule_end_weekends, owner=owner, pact_identifier=pact_identifier, pagerduty_integration=pagerduty_integration, prevent_deletion=prevent_deletion, required_role=required_role, requires_suitability=requires_suitability, template_environment=template_environment, unique_resource_prefix=unique_resource_prefix, updated_at=updated_at, values_name=values_name, limit=limit, offset=offset)

List Environments matching a filter

List Environments matching a filter.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_environment_v3 import SherlockEnvironmentV3
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
    api_instance = sherlock_python_client.EnvironmentsApi(api_client)
    auto_populate_chart_releases = True # bool | If true when creating, dynamic environments copy from template and template environments get the honeycomb chart (optional) (default to True)
    base = 'base_example' # str | Required when creating (optional)
    base_domain = 'bee.envs-terra.bio' # str |  (optional) (default to 'bee.envs-terra.bio')
    created_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    default_cluster = 'default_cluster_example' # str |  (optional)
    default_namespace = 'default_namespace_example' # str | When creating, will be calculated if left empty (optional)
    delete_after = '2013-10-20T19:20:30+01:00' # datetime | If set, the BEE will be automatically deleted after this time. Can be set to \"\" or Go's zero time value to clear the field. (optional)
    description = 'description_example' # str |  (optional)
    enable_janitor = True # bool | If true, janitor resource cleanup will be enabled for this environment. BEEs default to template's value, templates default to true, and static/live environments default to false. (optional)
    helmfile_ref = 'HEAD' # str |  (optional) (default to 'HEAD')
    id = 56 # int |  (optional)
    lifecycle = 'dynamic' # str |  (optional) (default to 'dynamic')
    name = 'name_example' # str | When creating, will be calculated if dynamic, required otherwise (optional)
    name_prefixes_domain = True # bool |  (optional) (default to True)
    offline = False # bool | Applicable for BEEs only, whether Thelma should render the BEE as \"offline\" zero replicas (this field is a target state, not a status) (optional) (default to False)
    offline_schedule_begin_enabled = True # bool | When enabled, the BEE will be slated to go offline around the begin time each day (optional)
    offline_schedule_begin_time = '2013-10-20T19:20:30+01:00' # datetime | Stored with timezone to determine day of the week (optional)
    offline_schedule_end_enabled = True # bool | When enabled, the BEE will be slated to come online around the end time each weekday (each day if weekends enabled) (optional)
    offline_schedule_end_time = '2013-10-20T19:20:30+01:00' # datetime | Stored with timezone to determine day of the week (optional)
    offline_schedule_end_weekends = True # bool |  (optional)
    owner = 'owner_example' # str | When creating, will default to you (optional)
    pact_identifier = 'pact_identifier_example' # str |  (optional)
    pagerduty_integration = 'pagerduty_integration_example' # str |  (optional)
    prevent_deletion = False # bool | Used to protect specific BEEs from deletion (thelma checks this field) (optional) (default to False)
    required_role = 'required_role_example' # str | If present, requires membership in the given role for mutations. Set to an empty string to clear. (optional)
    requires_suitability = True # bool |  (optional)
    template_environment = 'template_environment_example' # str | Required for dynamic environments (optional)
    unique_resource_prefix = 'unique_resource_prefix_example' # str | When creating, will be calculated if left empty (optional)
    updated_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    values_name = 'values_name_example' # str | When creating, defaults to template name or environment name (optional)
    limit = 56 # int | Control how many Environments are returned (default 0, meaning all) (optional)
    offset = 56 # int | Control the offset for the returned Environments (default 0) (optional)

    try:
        # List Environments matching a filter
        api_response = api_instance.api_environments_v3_get(auto_populate_chart_releases=auto_populate_chart_releases, base=base, base_domain=base_domain, created_at=created_at, default_cluster=default_cluster, default_namespace=default_namespace, delete_after=delete_after, description=description, enable_janitor=enable_janitor, helmfile_ref=helmfile_ref, id=id, lifecycle=lifecycle, name=name, name_prefixes_domain=name_prefixes_domain, offline=offline, offline_schedule_begin_enabled=offline_schedule_begin_enabled, offline_schedule_begin_time=offline_schedule_begin_time, offline_schedule_end_enabled=offline_schedule_end_enabled, offline_schedule_end_time=offline_schedule_end_time, offline_schedule_end_weekends=offline_schedule_end_weekends, owner=owner, pact_identifier=pact_identifier, pagerduty_integration=pagerduty_integration, prevent_deletion=prevent_deletion, required_role=required_role, requires_suitability=requires_suitability, template_environment=template_environment, unique_resource_prefix=unique_resource_prefix, updated_at=updated_at, values_name=values_name, limit=limit, offset=offset)
        print("The response of EnvironmentsApi->api_environments_v3_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling EnvironmentsApi->api_environments_v3_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **auto_populate_chart_releases** | **bool**| If true when creating, dynamic environments copy from template and template environments get the honeycomb chart | [optional] [default to True]
 **base** | **str**| Required when creating | [optional] 
 **base_domain** | **str**|  | [optional] [default to &#39;bee.envs-terra.bio&#39;]
 **created_at** | **datetime**|  | [optional] 
 **default_cluster** | **str**|  | [optional] 
 **default_namespace** | **str**| When creating, will be calculated if left empty | [optional] 
 **delete_after** | **datetime**| If set, the BEE will be automatically deleted after this time. Can be set to \&quot;\&quot; or Go&#39;s zero time value to clear the field. | [optional] 
 **description** | **str**|  | [optional] 
 **enable_janitor** | **bool**| If true, janitor resource cleanup will be enabled for this environment. BEEs default to template&#39;s value, templates default to true, and static/live environments default to false. | [optional] 
 **helmfile_ref** | **str**|  | [optional] [default to &#39;HEAD&#39;]
 **id** | **int**|  | [optional] 
 **lifecycle** | **str**|  | [optional] [default to &#39;dynamic&#39;]
 **name** | **str**| When creating, will be calculated if dynamic, required otherwise | [optional] 
 **name_prefixes_domain** | **bool**|  | [optional] [default to True]
 **offline** | **bool**| Applicable for BEEs only, whether Thelma should render the BEE as \&quot;offline\&quot; zero replicas (this field is a target state, not a status) | [optional] [default to False]
 **offline_schedule_begin_enabled** | **bool**| When enabled, the BEE will be slated to go offline around the begin time each day | [optional] 
 **offline_schedule_begin_time** | **datetime**| Stored with timezone to determine day of the week | [optional] 
 **offline_schedule_end_enabled** | **bool**| When enabled, the BEE will be slated to come online around the end time each weekday (each day if weekends enabled) | [optional] 
 **offline_schedule_end_time** | **datetime**| Stored with timezone to determine day of the week | [optional] 
 **offline_schedule_end_weekends** | **bool**|  | [optional] 
 **owner** | **str**| When creating, will default to you | [optional] 
 **pact_identifier** | **str**|  | [optional] 
 **pagerduty_integration** | **str**|  | [optional] 
 **prevent_deletion** | **bool**| Used to protect specific BEEs from deletion (thelma checks this field) | [optional] [default to False]
 **required_role** | **str**| If present, requires membership in the given role for mutations. Set to an empty string to clear. | [optional] 
 **requires_suitability** | **bool**|  | [optional] 
 **template_environment** | **str**| Required for dynamic environments | [optional] 
 **unique_resource_prefix** | **str**| When creating, will be calculated if left empty | [optional] 
 **updated_at** | **datetime**|  | [optional] 
 **values_name** | **str**| When creating, defaults to template name or environment name | [optional] 
 **limit** | **int**| Control how many Environments are returned (default 0, meaning all) | [optional] 
 **offset** | **int**| Control the offset for the returned Environments (default 0) | [optional] 

### Return type

[**List[SherlockEnvironmentV3]**](SherlockEnvironmentV3.md)

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

# **api_environments_v3_post**
> SherlockEnvironmentV3 api_environments_v3_post(environment)

Create a Environment

Create a Environment.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_environment_v3 import SherlockEnvironmentV3
from sherlock_python_client.models.sherlock_environment_v3_create import SherlockEnvironmentV3Create
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
    api_instance = sherlock_python_client.EnvironmentsApi(api_client)
    environment = sherlock_python_client.SherlockEnvironmentV3Create() # SherlockEnvironmentV3Create | The Environment to create

    try:
        # Create a Environment
        api_response = api_instance.api_environments_v3_post(environment)
        print("The response of EnvironmentsApi->api_environments_v3_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling EnvironmentsApi->api_environments_v3_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | [**SherlockEnvironmentV3Create**](SherlockEnvironmentV3Create.md)| The Environment to create | 

### Return type

[**SherlockEnvironmentV3**](SherlockEnvironmentV3.md)

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

# **api_environments_v3_selector_delete**
> SherlockEnvironmentV3 api_environments_v3_selector_delete(selector)

Delete an individual Environment

Delete an individual Environment by its ID.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_environment_v3 import SherlockEnvironmentV3
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
    api_instance = sherlock_python_client.EnvironmentsApi(api_client)
    selector = 'selector_example' # str | The selector of the Environment, which can be either a numeric ID, the name, or 'resource-prefix' + / + the unique resource prefix.

    try:
        # Delete an individual Environment
        api_response = api_instance.api_environments_v3_selector_delete(selector)
        print("The response of EnvironmentsApi->api_environments_v3_selector_delete:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling EnvironmentsApi->api_environments_v3_selector_delete: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the Environment, which can be either a numeric ID, the name, or &#39;resource-prefix&#39; + / + the unique resource prefix. | 

### Return type

[**SherlockEnvironmentV3**](SherlockEnvironmentV3.md)

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

# **api_environments_v3_selector_get**
> SherlockEnvironmentV3 api_environments_v3_selector_get(selector)

Get an individual Environment

Get an individual Environment.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_environment_v3 import SherlockEnvironmentV3
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
    api_instance = sherlock_python_client.EnvironmentsApi(api_client)
    selector = 'selector_example' # str | The selector of the Environment, which can be either a numeric ID, the name, or 'resource-prefix' + / + the unique resource prefix.

    try:
        # Get an individual Environment
        api_response = api_instance.api_environments_v3_selector_get(selector)
        print("The response of EnvironmentsApi->api_environments_v3_selector_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling EnvironmentsApi->api_environments_v3_selector_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the Environment, which can be either a numeric ID, the name, or &#39;resource-prefix&#39; + / + the unique resource prefix. | 

### Return type

[**SherlockEnvironmentV3**](SherlockEnvironmentV3.md)

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

# **api_environments_v3_selector_patch**
> SherlockEnvironmentV3 api_environments_v3_selector_patch(selector, environment)

Edit an individual Environment

Edit an individual Environment.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_environment_v3 import SherlockEnvironmentV3
from sherlock_python_client.models.sherlock_environment_v3_edit import SherlockEnvironmentV3Edit
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
    api_instance = sherlock_python_client.EnvironmentsApi(api_client)
    selector = 'selector_example' # str | The selector of the Environment, which can be either a numeric ID, the name, or 'resource-prefix' + / + the unique resource prefix.
    environment = sherlock_python_client.SherlockEnvironmentV3Edit() # SherlockEnvironmentV3Edit | The edits to make to the Environment

    try:
        # Edit an individual Environment
        api_response = api_instance.api_environments_v3_selector_patch(selector, environment)
        print("The response of EnvironmentsApi->api_environments_v3_selector_patch:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling EnvironmentsApi->api_environments_v3_selector_patch: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the Environment, which can be either a numeric ID, the name, or &#39;resource-prefix&#39; + / + the unique resource prefix. | 
 **environment** | [**SherlockEnvironmentV3Edit**](SherlockEnvironmentV3Edit.md)| The edits to make to the Environment | 

### Return type

[**SherlockEnvironmentV3**](SherlockEnvironmentV3.md)

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

