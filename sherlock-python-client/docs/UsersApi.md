# sherlock_python_client.UsersApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_users_v3_get**](UsersApi.md#api_users_v3_get) | **GET** /api/users/v3 | List Users matching a filter
[**api_users_v3_put**](UsersApi.md#api_users_v3_put) | **PUT** /api/users/v3 | Update the calling User&#39;s information
[**api_users_v3_selector_get**](UsersApi.md#api_users_v3_selector_get) | **GET** /api/users/v3/{selector} | Get an individual User


# **api_users_v3_get**
> List[SherlockUserV3] api_users_v3_get(created_at=created_at, deactivated_at=deactivated_at, email=email, github_id=github_id, github_username=github_username, google_id=google_id, id=id, name=name, name_from=name_from, slack_id=slack_id, slack_username=slack_username, suitability_description=suitability_description, suitable=suitable, updated_at=updated_at, limit=limit, offset=offset, include_deactivated=include_deactivated)

List Users matching a filter

List Users matching a filter. The results will include suitability and other information. Note that the suitability info can't directly be filtered for at this time.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_user_v3 import SherlockUserV3
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
    api_instance = sherlock_python_client.UsersApi(api_client)
    created_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    deactivated_at = 'deactivated_at_example' # str | If set, indicates that the user is currently deactivated (optional)
    email = 'email_example' # str |  (optional)
    github_id = 'github_id_example' # str |  (optional)
    github_username = 'github_username_example' # str |  (optional)
    google_id = 'google_id_example' # str |  (optional)
    id = 56 # int |  (optional)
    name = 'name_example' # str |  (optional)
    name_from = 'name_from_example' # str |  (optional)
    slack_id = 'slack_id_example' # str |  (optional)
    slack_username = 'slack_username_example' # str |  (optional)
    suitability_description = 'suitability_description_example' # str | Available only in responses; describes the user's production-suitability (optional)
    suitable = True # bool | Available only in responses; indicates whether the user is production-suitable (optional)
    updated_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    limit = 56 # int | Control how many Users are returned (default 0, no limit) (optional)
    offset = 56 # int | Control the offset for the returned Users (default 0) (optional)
    include_deactivated = True # bool | Include deactivated users in the results (default false) (optional)

    try:
        # List Users matching a filter
        api_response = api_instance.api_users_v3_get(created_at=created_at, deactivated_at=deactivated_at, email=email, github_id=github_id, github_username=github_username, google_id=google_id, id=id, name=name, name_from=name_from, slack_id=slack_id, slack_username=slack_username, suitability_description=suitability_description, suitable=suitable, updated_at=updated_at, limit=limit, offset=offset, include_deactivated=include_deactivated)
        print("The response of UsersApi->api_users_v3_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UsersApi->api_users_v3_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **created_at** | **datetime**|  | [optional] 
 **deactivated_at** | **str**| If set, indicates that the user is currently deactivated | [optional] 
 **email** | **str**|  | [optional] 
 **github_id** | **str**|  | [optional] 
 **github_username** | **str**|  | [optional] 
 **google_id** | **str**|  | [optional] 
 **id** | **int**|  | [optional] 
 **name** | **str**|  | [optional] 
 **name_from** | **str**|  | [optional] 
 **slack_id** | **str**|  | [optional] 
 **slack_username** | **str**|  | [optional] 
 **suitability_description** | **str**| Available only in responses; describes the user&#39;s production-suitability | [optional] 
 **suitable** | **bool**| Available only in responses; indicates whether the user is production-suitable | [optional] 
 **updated_at** | **datetime**|  | [optional] 
 **limit** | **int**| Control how many Users are returned (default 0, no limit) | [optional] 
 **offset** | **int**| Control the offset for the returned Users (default 0) | [optional] 
 **include_deactivated** | **bool**| Include deactivated users in the results (default false) | [optional] 

### Return type

[**List[SherlockUserV3]**](SherlockUserV3.md)

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

# **api_users_v3_put**
> SherlockUserV3 api_users_v3_put(user=user)

Update the calling User's information

Update the calling User's information. As with all authenticated Sherlock endpoints, newly-observed callers will have a User record added, meaning that this endpoint behaves like an upsert.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_user_v3 import SherlockUserV3
from sherlock_python_client.models.sherlock_user_v3_upsert import SherlockUserV3Upsert
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
    api_instance = sherlock_python_client.UsersApi(api_client)
    user = sherlock_python_client.SherlockUserV3Upsert() # SherlockUserV3Upsert | The User data to update (optional)

    try:
        # Update the calling User's information
        api_response = api_instance.api_users_v3_put(user=user)
        print("The response of UsersApi->api_users_v3_put:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UsersApi->api_users_v3_put: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**SherlockUserV3Upsert**](SherlockUserV3Upsert.md)| The User data to update | [optional] 

### Return type

[**SherlockUserV3**](SherlockUserV3.md)

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

# **api_users_v3_selector_get**
> SherlockUserV3 api_users_v3_selector_get(selector)

Get an individual User

Get an individual User. As a special case, \"me\" or \"self\" can be passed as the selector to get the current user.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_user_v3 import SherlockUserV3
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
    api_instance = sherlock_python_client.UsersApi(api_client)
    selector = 'selector_example' # str | The selector of the User, which can be either a numeric ID, the email, 'google-id/{google subject ID}', 'github/{github username}', or 'github-id/{github numeric ID}'. As a special case, 'me' or 'self' can be passed to get the calling user.

    try:
        # Get an individual User
        api_response = api_instance.api_users_v3_selector_get(selector)
        print("The response of UsersApi->api_users_v3_selector_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UsersApi->api_users_v3_selector_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the User, which can be either a numeric ID, the email, &#39;google-id/{google subject ID}&#39;, &#39;github/{github username}&#39;, or &#39;github-id/{github numeric ID}&#39;. As a special case, &#39;me&#39; or &#39;self&#39; can be passed to get the calling user. | 

### Return type

[**SherlockUserV3**](SherlockUserV3.md)

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

