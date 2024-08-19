# sherlock_python_client.IncidentsApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_incidents_v3_get**](IncidentsApi.md#api_incidents_v3_get) | **GET** /api/incidents/v3 | List Incidents matching a filter
[**api_incidents_v3_post**](IncidentsApi.md#api_incidents_v3_post) | **POST** /api/incidents/v3 | Create a Incident
[**api_incidents_v3_selector_delete**](IncidentsApi.md#api_incidents_v3_selector_delete) | **DELETE** /api/incidents/v3/{selector} | Delete an individual Incident
[**api_incidents_v3_selector_get**](IncidentsApi.md#api_incidents_v3_selector_get) | **GET** /api/incidents/v3/{selector} | Get an individual Incident
[**api_incidents_v3_selector_patch**](IncidentsApi.md#api_incidents_v3_selector_patch) | **PATCH** /api/incidents/v3/{selector} | Edit an individual Incident


# **api_incidents_v3_get**
> List[SherlockIncidentV3] api_incidents_v3_get(created_at=created_at, description=description, id=id, remediated_at=remediated_at, review_completed_at=review_completed_at, started_at=started_at, ticket=ticket, updated_at=updated_at, limit=limit, offset=offset)

List Incidents matching a filter

List Incidents matching a filter.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_incident_v3 import SherlockIncidentV3
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
    api_instance = sherlock_python_client.IncidentsApi(api_client)
    created_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    description = 'description_example' # str |  (optional)
    id = 56 # int |  (optional)
    remediated_at = 'remediated_at_example' # str |  (optional)
    review_completed_at = 'review_completed_at_example' # str |  (optional)
    started_at = 'started_at_example' # str |  (optional)
    ticket = 'ticket_example' # str |  (optional)
    updated_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    limit = 56 # int | Control how many Incidents are returned (default 0, meaning all) (optional)
    offset = 56 # int | Control the offset for the returned Incidents (default 0) (optional)

    try:
        # List Incidents matching a filter
        api_response = api_instance.api_incidents_v3_get(created_at=created_at, description=description, id=id, remediated_at=remediated_at, review_completed_at=review_completed_at, started_at=started_at, ticket=ticket, updated_at=updated_at, limit=limit, offset=offset)
        print("The response of IncidentsApi->api_incidents_v3_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling IncidentsApi->api_incidents_v3_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **created_at** | **datetime**|  | [optional] 
 **description** | **str**|  | [optional] 
 **id** | **int**|  | [optional] 
 **remediated_at** | **str**|  | [optional] 
 **review_completed_at** | **str**|  | [optional] 
 **started_at** | **str**|  | [optional] 
 **ticket** | **str**|  | [optional] 
 **updated_at** | **datetime**|  | [optional] 
 **limit** | **int**| Control how many Incidents are returned (default 0, meaning all) | [optional] 
 **offset** | **int**| Control the offset for the returned Incidents (default 0) | [optional] 

### Return type

[**List[SherlockIncidentV3]**](SherlockIncidentV3.md)

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

# **api_incidents_v3_post**
> SherlockIncidentV3 api_incidents_v3_post(incident)

Create a Incident

Create a Incident.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_incident_v3 import SherlockIncidentV3
from sherlock_python_client.models.sherlock_incident_v3_create import SherlockIncidentV3Create
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
    api_instance = sherlock_python_client.IncidentsApi(api_client)
    incident = sherlock_python_client.SherlockIncidentV3Create() # SherlockIncidentV3Create | The Incident to create

    try:
        # Create a Incident
        api_response = api_instance.api_incidents_v3_post(incident)
        print("The response of IncidentsApi->api_incidents_v3_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling IncidentsApi->api_incidents_v3_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **incident** | [**SherlockIncidentV3Create**](SherlockIncidentV3Create.md)| The Incident to create | 

### Return type

[**SherlockIncidentV3**](SherlockIncidentV3.md)

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

# **api_incidents_v3_selector_delete**
> SherlockIncidentV3 api_incidents_v3_selector_delete(selector)

Delete an individual Incident

Delete an individual Incident by its ID.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_incident_v3 import SherlockIncidentV3
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
    api_instance = sherlock_python_client.IncidentsApi(api_client)
    selector = 'selector_example' # str | The ID of the Incident

    try:
        # Delete an individual Incident
        api_response = api_instance.api_incidents_v3_selector_delete(selector)
        print("The response of IncidentsApi->api_incidents_v3_selector_delete:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling IncidentsApi->api_incidents_v3_selector_delete: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The ID of the Incident | 

### Return type

[**SherlockIncidentV3**](SherlockIncidentV3.md)

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

# **api_incidents_v3_selector_get**
> SherlockIncidentV3 api_incidents_v3_selector_get(selector)

Get an individual Incident

Get an individual Incident.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_incident_v3 import SherlockIncidentV3
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
    api_instance = sherlock_python_client.IncidentsApi(api_client)
    selector = 'selector_example' # str | The ID of the Incident

    try:
        # Get an individual Incident
        api_response = api_instance.api_incidents_v3_selector_get(selector)
        print("The response of IncidentsApi->api_incidents_v3_selector_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling IncidentsApi->api_incidents_v3_selector_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The ID of the Incident | 

### Return type

[**SherlockIncidentV3**](SherlockIncidentV3.md)

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

# **api_incidents_v3_selector_patch**
> SherlockIncidentV3 api_incidents_v3_selector_patch(selector, incident)

Edit an individual Incident

Edit an individual Incident.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_incident_v3 import SherlockIncidentV3
from sherlock_python_client.models.sherlock_incident_v3_edit import SherlockIncidentV3Edit
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
    api_instance = sherlock_python_client.IncidentsApi(api_client)
    selector = 'selector_example' # str | The ID of the Incident
    incident = sherlock_python_client.SherlockIncidentV3Edit() # SherlockIncidentV3Edit | The edits to make to the Incident

    try:
        # Edit an individual Incident
        api_response = api_instance.api_incidents_v3_selector_patch(selector, incident)
        print("The response of IncidentsApi->api_incidents_v3_selector_patch:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling IncidentsApi->api_incidents_v3_selector_patch: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The ID of the Incident | 
 **incident** | [**SherlockIncidentV3Edit**](SherlockIncidentV3Edit.md)| The edits to make to the Incident | 

### Return type

[**SherlockIncidentV3**](SherlockIncidentV3.md)

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

