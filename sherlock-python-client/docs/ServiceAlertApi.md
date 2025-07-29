# sherlock_python_client.ServiceAlertApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_service_alerts_procedures_v3_sync_post**](ServiceAlertApi.md#api_service_alerts_procedures_v3_sync_post) | **POST** /api/service-alerts/procedures/v3/sync | Sync service alerts
[**api_service_alerts_v3_get**](ServiceAlertApi.md#api_service_alerts_v3_get) | **GET** /api/service-alerts/v3 | List ServiceAlerts matching a filter
[**api_service_alerts_v3_post**](ServiceAlertApi.md#api_service_alerts_v3_post) | **POST** /api/service-alerts/v3 | Create a service alert
[**api_service_alerts_v3_selector_delete**](ServiceAlertApi.md#api_service_alerts_v3_selector_delete) | **DELETE** /api/service-alerts/v3/{selector} | Delete a ServiceAlert
[**api_service_alerts_v3_selector_get**](ServiceAlertApi.md#api_service_alerts_v3_selector_get) | **GET** /api/service-alerts/v3/{selector} | Get a Service Alert
[**api_service_alerts_v3_selector_patch**](ServiceAlertApi.md#api_service_alerts_v3_selector_patch) | **PATCH** /api/service-alerts/v3/{selector} | Edit a service alert


# **api_service_alerts_procedures_v3_sync_post**
> List[SherlockServiceAlertV3] api_service_alerts_procedures_v3_sync_post(environment)

Sync service alerts

Method to get all currently active service alerts from Sherlock's DB and ensure that the service alert json files placed in Google Buckets for Terra match.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_service_alert_v3 import SherlockServiceAlertV3
from sherlock_python_client.models.sherlock_service_alert_v3_sync_request import SherlockServiceAlertV3SyncRequest
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
    api_instance = sherlock_python_client.ServiceAlertApi(api_client)
    environment = sherlock_python_client.SherlockServiceAlertV3SyncRequest() # SherlockServiceAlertV3SyncRequest | Information on Service Alert environment

    try:
        # Sync service alerts
        api_response = api_instance.api_service_alerts_procedures_v3_sync_post(environment)
        print("The response of ServiceAlertApi->api_service_alerts_procedures_v3_sync_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ServiceAlertApi->api_service_alerts_procedures_v3_sync_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | [**SherlockServiceAlertV3SyncRequest**](SherlockServiceAlertV3SyncRequest.md)| Information on Service Alert environment | 

### Return type

[**List[SherlockServiceAlertV3]**](SherlockServiceAlertV3.md)

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

# **api_service_alerts_v3_get**
> List[SherlockServiceAlertV3] api_service_alerts_v3_get(created_at=created_at, created_by=created_by, delted_by=delted_by, id=id, link=link, message=message, on_environment=on_environment, severity=severity, title=title, updated_at=updated_at, updated_by=updated_by, uuid=uuid, limit=limit, offset=offset, include_deleted=include_deleted)

List ServiceAlerts matching a filter

List ServiceAlerts matching a filter.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_service_alert_v3 import SherlockServiceAlertV3
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
    api_instance = sherlock_python_client.ServiceAlertApi(api_client)
    created_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    created_by = 'created_by_example' # str |  (optional)
    delted_by = 'delted_by_example' # str |  (optional)
    id = 56 # int |  (optional)
    link = 'link_example' # str |  (optional)
    message = 'message_example' # str |  (optional)
    on_environment = 'on_environment_example' # str |  (optional)
    severity = 'severity_example' # str |  (optional)
    title = 'title_example' # str |  (optional)
    updated_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    updated_by = 'updated_by_example' # str |  (optional)
    uuid = 'uuid_example' # str |  (optional)
    limit = 56 # int | Control how many Service Alerts are returned (default 0, no limit) (optional)
    offset = 56 # int | Control the offset for the returned Service Alerts (default 0) (optional)
    include_deleted = True # bool | Control if only active Service Alerts are returned, set to true to return deleted Alerts (default false) (optional)

    try:
        # List ServiceAlerts matching a filter
        api_response = api_instance.api_service_alerts_v3_get(created_at=created_at, created_by=created_by, delted_by=delted_by, id=id, link=link, message=message, on_environment=on_environment, severity=severity, title=title, updated_at=updated_at, updated_by=updated_by, uuid=uuid, limit=limit, offset=offset, include_deleted=include_deleted)
        print("The response of ServiceAlertApi->api_service_alerts_v3_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ServiceAlertApi->api_service_alerts_v3_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **created_at** | **datetime**|  | [optional] 
 **created_by** | **str**|  | [optional] 
 **delted_by** | **str**|  | [optional] 
 **id** | **int**|  | [optional] 
 **link** | **str**|  | [optional] 
 **message** | **str**|  | [optional] 
 **on_environment** | **str**|  | [optional] 
 **severity** | **str**|  | [optional] 
 **title** | **str**|  | [optional] 
 **updated_at** | **datetime**|  | [optional] 
 **updated_by** | **str**|  | [optional] 
 **uuid** | **str**|  | [optional] 
 **limit** | **int**| Control how many Service Alerts are returned (default 0, no limit) | [optional] 
 **offset** | **int**| Control the offset for the returned Service Alerts (default 0) | [optional] 
 **include_deleted** | **bool**| Control if only active Service Alerts are returned, set to true to return deleted Alerts (default false) | [optional] 

### Return type

[**List[SherlockServiceAlertV3]**](SherlockServiceAlertV3.md)

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

# **api_service_alerts_v3_post**
> SherlockServiceAlertV3 api_service_alerts_v3_post(service_alert)

Create a service alert

Create a service alert to be displayed within terra.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_service_alert_v3 import SherlockServiceAlertV3
from sherlock_python_client.models.sherlock_service_alert_v3_create import SherlockServiceAlertV3Create
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
    api_instance = sherlock_python_client.ServiceAlertApi(api_client)
    service_alert = sherlock_python_client.SherlockServiceAlertV3Create() # SherlockServiceAlertV3Create | The initial fields the ServiceAlert should have set

    try:
        # Create a service alert
        api_response = api_instance.api_service_alerts_v3_post(service_alert)
        print("The response of ServiceAlertApi->api_service_alerts_v3_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ServiceAlertApi->api_service_alerts_v3_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service_alert** | [**SherlockServiceAlertV3Create**](SherlockServiceAlertV3Create.md)| The initial fields the ServiceAlert should have set | 

### Return type

[**SherlockServiceAlertV3**](SherlockServiceAlertV3.md)

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

# **api_service_alerts_v3_selector_delete**
> SherlockServiceAlertV3 api_service_alerts_v3_selector_delete(selector)

Delete a ServiceAlert

Delete an individual ServiceAlert.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_service_alert_v3 import SherlockServiceAlertV3
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
    api_instance = sherlock_python_client.ServiceAlertApi(api_client)
    selector = 'selector_example' # str | The selector of the ServiceAlert, ServiceAlert, which is the guid for a given alert

    try:
        # Delete a ServiceAlert
        api_response = api_instance.api_service_alerts_v3_selector_delete(selector)
        print("The response of ServiceAlertApi->api_service_alerts_v3_selector_delete:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ServiceAlertApi->api_service_alerts_v3_selector_delete: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the ServiceAlert, ServiceAlert, which is the guid for a given alert | 

### Return type

[**SherlockServiceAlertV3**](SherlockServiceAlertV3.md)

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

# **api_service_alerts_v3_selector_get**
> SherlockServiceAlertV3 api_service_alerts_v3_selector_get(selector, include_deleted=include_deleted)

Get a Service Alert

Get an individual Service Alert and it's metadata.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_service_alert_v3 import SherlockServiceAlertV3
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
    api_instance = sherlock_python_client.ServiceAlertApi(api_client)
    selector = 'selector_example' # str | The selector of the ServiceAlert, which is the guid for a given alert
    include_deleted = True # bool | Control if only active Service Alerts are returned, set to true to return deleted Alerts (default false) (optional)

    try:
        # Get a Service Alert
        api_response = api_instance.api_service_alerts_v3_selector_get(selector, include_deleted=include_deleted)
        print("The response of ServiceAlertApi->api_service_alerts_v3_selector_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ServiceAlertApi->api_service_alerts_v3_selector_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the ServiceAlert, which is the guid for a given alert | 
 **include_deleted** | **bool**| Control if only active Service Alerts are returned, set to true to return deleted Alerts (default false) | [optional] 

### Return type

[**SherlockServiceAlertV3**](SherlockServiceAlertV3.md)

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

# **api_service_alerts_v3_selector_patch**
> SherlockServiceAlertV3 api_service_alerts_v3_selector_patch(selector, service_alert)

Edit a service alert

Update a service alert with new information.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_service_alert_v3 import SherlockServiceAlertV3
from sherlock_python_client.models.sherlock_service_alert_v3_editable_fields import SherlockServiceAlertV3EditableFields
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
    api_instance = sherlock_python_client.ServiceAlertApi(api_client)
    selector = 'selector_example' # str | The selector of the ServiceAlert, which is the guid for a given alert
    service_alert = sherlock_python_client.SherlockServiceAlertV3EditableFields() # SherlockServiceAlertV3EditableFields | The edits to make to the ServiceAlert

    try:
        # Edit a service alert
        api_response = api_instance.api_service_alerts_v3_selector_patch(selector, service_alert)
        print("The response of ServiceAlertApi->api_service_alerts_v3_selector_patch:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ServiceAlertApi->api_service_alerts_v3_selector_patch: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the ServiceAlert, which is the guid for a given alert | 
 **service_alert** | [**SherlockServiceAlertV3EditableFields**](SherlockServiceAlertV3EditableFields.md)| The edits to make to the ServiceAlert | 

### Return type

[**SherlockServiceAlertV3**](SherlockServiceAlertV3.md)

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

