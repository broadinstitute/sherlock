# sherlock_python_client.PagerdutyIntegrationsApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_pagerduty_integrations_procedures_v3_trigger_incident_selector_post**](PagerdutyIntegrationsApi.md#api_pagerduty_integrations_procedures_v3_trigger_incident_selector_post) | **POST** /api/pagerduty-integrations/procedures/v3/trigger-incident/{selector} | Get an individual PagerdutyIntegration
[**api_pagerduty_integrations_v3_get**](PagerdutyIntegrationsApi.md#api_pagerduty_integrations_v3_get) | **GET** /api/pagerduty-integrations/v3 | List PagerdutyIntegrations matching a filter
[**api_pagerduty_integrations_v3_post**](PagerdutyIntegrationsApi.md#api_pagerduty_integrations_v3_post) | **POST** /api/pagerduty-integrations/v3 | Create a PagerdutyIntegration
[**api_pagerduty_integrations_v3_selector_delete**](PagerdutyIntegrationsApi.md#api_pagerduty_integrations_v3_selector_delete) | **DELETE** /api/pagerduty-integrations/v3/{selector} | Delete an individual PagerdutyIntegration
[**api_pagerduty_integrations_v3_selector_get**](PagerdutyIntegrationsApi.md#api_pagerduty_integrations_v3_selector_get) | **GET** /api/pagerduty-integrations/v3/{selector} | Get an individual PagerdutyIntegration
[**api_pagerduty_integrations_v3_selector_patch**](PagerdutyIntegrationsApi.md#api_pagerduty_integrations_v3_selector_patch) | **PATCH** /api/pagerduty-integrations/v3/{selector} | Edit an individual PagerdutyIntegration


# **api_pagerduty_integrations_procedures_v3_trigger_incident_selector_post**
> PagerdutySendAlertResponse api_pagerduty_integrations_procedures_v3_trigger_incident_selector_post(selector, summary)

Get an individual PagerdutyIntegration

Get an individual PagerdutyIntegration.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.pagerduty_alert_summary import PagerdutyAlertSummary
from sherlock_python_client.models.pagerduty_send_alert_response import PagerdutySendAlertResponse
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
    api_instance = sherlock_python_client.PagerdutyIntegrationsApi(api_client)
    selector = 'selector_example' # str | The selector of the PagerdutyIntegration, which can be either a numeric ID or pd-id/<pagerduty-id>.
    summary = sherlock_python_client.PagerdutyAlertSummary() # PagerdutyAlertSummary | Summary of the incident

    try:
        # Get an individual PagerdutyIntegration
        api_response = api_instance.api_pagerduty_integrations_procedures_v3_trigger_incident_selector_post(selector, summary)
        print("The response of PagerdutyIntegrationsApi->api_pagerduty_integrations_procedures_v3_trigger_incident_selector_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PagerdutyIntegrationsApi->api_pagerduty_integrations_procedures_v3_trigger_incident_selector_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the PagerdutyIntegration, which can be either a numeric ID or pd-id/&lt;pagerduty-id&gt;. | 
 **summary** | [**PagerdutyAlertSummary**](PagerdutyAlertSummary.md)| Summary of the incident | 

### Return type

[**PagerdutySendAlertResponse**](PagerdutySendAlertResponse.md)

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

# **api_pagerduty_integrations_v3_get**
> List[SherlockPagerdutyIntegrationV3] api_pagerduty_integrations_v3_get(created_at=created_at, id=id, name=name, pagerduty_id=pagerduty_id, type=type, updated_at=updated_at, limit=limit, offset=offset)

List PagerdutyIntegrations matching a filter

List PagerdutyIntegrations matching a filter.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_pagerduty_integration_v3 import SherlockPagerdutyIntegrationV3
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
    api_instance = sherlock_python_client.PagerdutyIntegrationsApi(api_client)
    created_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    id = 56 # int |  (optional)
    name = 'name_example' # str |  (optional)
    pagerduty_id = 'pagerduty_id_example' # str |  (optional)
    type = 'type_example' # str |  (optional)
    updated_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    limit = 56 # int | Control how many PagerdutyIntegrations are returned (default 0, meaning all) (optional)
    offset = 56 # int | Control the offset for the returned PagerdutyIntegrations (default 0) (optional)

    try:
        # List PagerdutyIntegrations matching a filter
        api_response = api_instance.api_pagerduty_integrations_v3_get(created_at=created_at, id=id, name=name, pagerduty_id=pagerduty_id, type=type, updated_at=updated_at, limit=limit, offset=offset)
        print("The response of PagerdutyIntegrationsApi->api_pagerduty_integrations_v3_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PagerdutyIntegrationsApi->api_pagerduty_integrations_v3_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **created_at** | **datetime**|  | [optional] 
 **id** | **int**|  | [optional] 
 **name** | **str**|  | [optional] 
 **pagerduty_id** | **str**|  | [optional] 
 **type** | **str**|  | [optional] 
 **updated_at** | **datetime**|  | [optional] 
 **limit** | **int**| Control how many PagerdutyIntegrations are returned (default 0, meaning all) | [optional] 
 **offset** | **int**| Control the offset for the returned PagerdutyIntegrations (default 0) | [optional] 

### Return type

[**List[SherlockPagerdutyIntegrationV3]**](SherlockPagerdutyIntegrationV3.md)

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

# **api_pagerduty_integrations_v3_post**
> SherlockPagerdutyIntegrationV3 api_pagerduty_integrations_v3_post(pagerduty_integration)

Create a PagerdutyIntegration

Create a PagerdutyIntegration. Duplicate Pagerduty IDs will be gracefully handled by editing the existing entry. This is partially opaque because some fields are writable but not readable.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_pagerduty_integration_v3 import SherlockPagerdutyIntegrationV3
from sherlock_python_client.models.sherlock_pagerduty_integration_v3_create import SherlockPagerdutyIntegrationV3Create
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
    api_instance = sherlock_python_client.PagerdutyIntegrationsApi(api_client)
    pagerduty_integration = sherlock_python_client.SherlockPagerdutyIntegrationV3Create() # SherlockPagerdutyIntegrationV3Create | The PagerdutyIntegration to create

    try:
        # Create a PagerdutyIntegration
        api_response = api_instance.api_pagerduty_integrations_v3_post(pagerduty_integration)
        print("The response of PagerdutyIntegrationsApi->api_pagerduty_integrations_v3_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PagerdutyIntegrationsApi->api_pagerduty_integrations_v3_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pagerduty_integration** | [**SherlockPagerdutyIntegrationV3Create**](SherlockPagerdutyIntegrationV3Create.md)| The PagerdutyIntegration to create | 

### Return type

[**SherlockPagerdutyIntegrationV3**](SherlockPagerdutyIntegrationV3.md)

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

# **api_pagerduty_integrations_v3_selector_delete**
> SherlockPagerdutyIntegrationV3 api_pagerduty_integrations_v3_selector_delete(selector)

Delete an individual PagerdutyIntegration

Delete an individual PagerdutyIntegration by its ID.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_pagerduty_integration_v3 import SherlockPagerdutyIntegrationV3
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
    api_instance = sherlock_python_client.PagerdutyIntegrationsApi(api_client)
    selector = 'selector_example' # str | The selector of the PagerdutyIntegration, which can be either a numeric ID or pd-id/<pagerduty-id>.

    try:
        # Delete an individual PagerdutyIntegration
        api_response = api_instance.api_pagerduty_integrations_v3_selector_delete(selector)
        print("The response of PagerdutyIntegrationsApi->api_pagerduty_integrations_v3_selector_delete:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PagerdutyIntegrationsApi->api_pagerduty_integrations_v3_selector_delete: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the PagerdutyIntegration, which can be either a numeric ID or pd-id/&lt;pagerduty-id&gt;. | 

### Return type

[**SherlockPagerdutyIntegrationV3**](SherlockPagerdutyIntegrationV3.md)

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

# **api_pagerduty_integrations_v3_selector_get**
> SherlockPagerdutyIntegrationV3 api_pagerduty_integrations_v3_selector_get(selector)

Get an individual PagerdutyIntegration

Get an individual PagerdutyIntegration.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_pagerduty_integration_v3 import SherlockPagerdutyIntegrationV3
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
    api_instance = sherlock_python_client.PagerdutyIntegrationsApi(api_client)
    selector = 'selector_example' # str | The selector of the PagerdutyIntegration, which can be either a numeric ID or pd-id/<pagerduty-id>.

    try:
        # Get an individual PagerdutyIntegration
        api_response = api_instance.api_pagerduty_integrations_v3_selector_get(selector)
        print("The response of PagerdutyIntegrationsApi->api_pagerduty_integrations_v3_selector_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PagerdutyIntegrationsApi->api_pagerduty_integrations_v3_selector_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the PagerdutyIntegration, which can be either a numeric ID or pd-id/&lt;pagerduty-id&gt;. | 

### Return type

[**SherlockPagerdutyIntegrationV3**](SherlockPagerdutyIntegrationV3.md)

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

# **api_pagerduty_integrations_v3_selector_patch**
> SherlockPagerdutyIntegrationV3 api_pagerduty_integrations_v3_selector_patch(selector, pagerduty_integration)

Edit an individual PagerdutyIntegration

Edit an individual PagerdutyIntegration.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_pagerduty_integration_v3 import SherlockPagerdutyIntegrationV3
from sherlock_python_client.models.sherlock_pagerduty_integration_v3_edit import SherlockPagerdutyIntegrationV3Edit
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
    api_instance = sherlock_python_client.PagerdutyIntegrationsApi(api_client)
    selector = 'selector_example' # str | The selector of the PagerdutyIntegration, which can be either a numeric ID or pd-id/<pagerduty-id>.
    pagerduty_integration = sherlock_python_client.SherlockPagerdutyIntegrationV3Edit() # SherlockPagerdutyIntegrationV3Edit | The edits to make to the PagerdutyIntegration

    try:
        # Edit an individual PagerdutyIntegration
        api_response = api_instance.api_pagerduty_integrations_v3_selector_patch(selector, pagerduty_integration)
        print("The response of PagerdutyIntegrationsApi->api_pagerduty_integrations_v3_selector_patch:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PagerdutyIntegrationsApi->api_pagerduty_integrations_v3_selector_patch: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the PagerdutyIntegration, which can be either a numeric ID or pd-id/&lt;pagerduty-id&gt;. | 
 **pagerduty_integration** | [**SherlockPagerdutyIntegrationV3Edit**](SherlockPagerdutyIntegrationV3Edit.md)| The edits to make to the PagerdutyIntegration | 

### Return type

[**SherlockPagerdutyIntegrationV3**](SherlockPagerdutyIntegrationV3.md)

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

