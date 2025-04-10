# sherlock_python_client.RoleAssignmentsApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_role_assignments_v3_get**](RoleAssignmentsApi.md#api_role_assignments_v3_get) | **GET** /api/role-assignments/v3 | List RoleAssignments matching a filter
[**api_role_assignments_v3_role_selector_user_selector_delete**](RoleAssignmentsApi.md#api_role_assignments_v3_role_selector_user_selector_delete) | **DELETE** /api/role-assignments/v3/{role-selector}/{user-selector} | Delete a RoleAssignment
[**api_role_assignments_v3_role_selector_user_selector_get**](RoleAssignmentsApi.md#api_role_assignments_v3_role_selector_user_selector_get) | **GET** /api/role-assignments/v3/{role-selector}/{user-selector} | Get a RoleAssignment
[**api_role_assignments_v3_role_selector_user_selector_patch**](RoleAssignmentsApi.md#api_role_assignments_v3_role_selector_user_selector_patch) | **PATCH** /api/role-assignments/v3/{role-selector}/{user-selector} | Edit a RoleAssignment
[**api_role_assignments_v3_role_selector_user_selector_post**](RoleAssignmentsApi.md#api_role_assignments_v3_role_selector_user_selector_post) | **POST** /api/role-assignments/v3/{role-selector}/{user-selector} | Create a RoleAssignment


# **api_role_assignments_v3_get**
> List[SherlockRoleAssignmentV3] api_role_assignments_v3_get(expires_at=expires_at, expires_in=expires_in, suspended=suspended, limit=limit, offset=offset)

List RoleAssignments matching a filter

List RoleAssignments matching a filter. The correct way to list RoleAssignments for a particular Role or User is to get that Role or User specifically, not to use this endpoint.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_role_assignment_v3 import SherlockRoleAssignmentV3
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
    api_instance = sherlock_python_client.RoleAssignmentsApi(api_client)
    expires_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    expires_in = 'expires_in_example' # str | A Go time.Duration string that will be added to the current time to attempt to set expiresAt (this may be more convenient than setting expiresAt directly) (optional)
    suspended = True # bool | If the assignment should be active. This field is only mutable through the API if the role doesn't automatically suspend non-suitable users (optional)
    limit = 56 # int | Control how many RoleAssignments are returned (default 0, no limit) (optional)
    offset = 56 # int | Control the offset for the returned RoleAssignments (default 0) (optional)

    try:
        # List RoleAssignments matching a filter
        api_response = api_instance.api_role_assignments_v3_get(expires_at=expires_at, expires_in=expires_in, suspended=suspended, limit=limit, offset=offset)
        print("The response of RoleAssignmentsApi->api_role_assignments_v3_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RoleAssignmentsApi->api_role_assignments_v3_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expires_at** | **datetime**|  | [optional] 
 **expires_in** | **str**| A Go time.Duration string that will be added to the current time to attempt to set expiresAt (this may be more convenient than setting expiresAt directly) | [optional] 
 **suspended** | **bool**| If the assignment should be active. This field is only mutable through the API if the role doesn&#39;t automatically suspend non-suitable users | [optional] 
 **limit** | **int**| Control how many RoleAssignments are returned (default 0, no limit) | [optional] 
 **offset** | **int**| Control the offset for the returned RoleAssignments (default 0) | [optional] 

### Return type

[**List[SherlockRoleAssignmentV3]**](SherlockRoleAssignmentV3.md)

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

# **api_role_assignments_v3_role_selector_user_selector_delete**
> api_role_assignments_v3_role_selector_user_selector_delete(role_selector, user_selector)

Delete a RoleAssignment

Delete the RoleAssignment between a given Role and User.
Non-super-admins may only mutate RoleAssignments for themselves, only for roles they can break-glass into, and only with an expiry no further than the role's default break-glass duration in the future.
Propagation will be triggered after this operation.

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
    api_instance = sherlock_python_client.RoleAssignmentsApi(api_client)
    role_selector = 'role_selector_example' # str | The selector of the Role, which can be either the numeric ID or the name
    user_selector = 'user_selector_example' # str | The selector of the User, which can be either a numeric ID, the email, 'google-id/{google subject ID}', 'github/{github username}', or 'github-id/{github numeric ID}'.

    try:
        # Delete a RoleAssignment
        api_instance.api_role_assignments_v3_role_selector_user_selector_delete(role_selector, user_selector)
    except Exception as e:
        print("Exception when calling RoleAssignmentsApi->api_role_assignments_v3_role_selector_user_selector_delete: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role_selector** | **str**| The selector of the Role, which can be either the numeric ID or the name | 
 **user_selector** | **str**| The selector of the User, which can be either a numeric ID, the email, &#39;google-id/{google subject ID}&#39;, &#39;github/{github username}&#39;, or &#39;github-id/{github numeric ID}&#39;. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**400** | Bad Request |  -  |
**403** | Forbidden |  -  |
**404** | Not Found |  -  |
**407** | Proxy Authentication Required |  -  |
**409** | Conflict |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **api_role_assignments_v3_role_selector_user_selector_get**
> SherlockRoleAssignmentV3 api_role_assignments_v3_role_selector_user_selector_get(role_selector, user_selector)

Get a RoleAssignment

Get the RoleAssignment between a given Role and User.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_role_assignment_v3 import SherlockRoleAssignmentV3
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
    api_instance = sherlock_python_client.RoleAssignmentsApi(api_client)
    role_selector = 'role_selector_example' # str | The selector of the Role, which can be either the numeric ID or the name
    user_selector = 'user_selector_example' # str | The selector of the User, which can be either a numeric ID, the email, 'google-id/{google subject ID}', 'github/{github username}', or 'github-id/{github numeric ID}'.

    try:
        # Get a RoleAssignment
        api_response = api_instance.api_role_assignments_v3_role_selector_user_selector_get(role_selector, user_selector)
        print("The response of RoleAssignmentsApi->api_role_assignments_v3_role_selector_user_selector_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RoleAssignmentsApi->api_role_assignments_v3_role_selector_user_selector_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role_selector** | **str**| The selector of the Role, which can be either the numeric ID or the name | 
 **user_selector** | **str**| The selector of the User, which can be either a numeric ID, the email, &#39;google-id/{google subject ID}&#39;, &#39;github/{github username}&#39;, or &#39;github-id/{github numeric ID}&#39;. | 

### Return type

[**SherlockRoleAssignmentV3**](SherlockRoleAssignmentV3.md)

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

# **api_role_assignments_v3_role_selector_user_selector_patch**
> SherlockRoleAssignmentV3 api_role_assignments_v3_role_selector_user_selector_patch(role_selector, user_selector, role_assignment)

Edit a RoleAssignment

Edit the RoleAssignment between a given Role and User.
Non-super-admins may only mutate RoleAssignments for themselves, only for roles they can break-glass into, and only with an expiry no further than the role's default break-glass duration in the future.
Propagation will be triggered after this operation.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_role_assignment_v3 import SherlockRoleAssignmentV3
from sherlock_python_client.models.sherlock_role_assignment_v3_edit import SherlockRoleAssignmentV3Edit
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
    api_instance = sherlock_python_client.RoleAssignmentsApi(api_client)
    role_selector = 'role_selector_example' # str | The selector of the Role, which can be either the numeric ID or the name
    user_selector = 'user_selector_example' # str | The selector of the User, which can be either a numeric ID, the email, 'google-id/{google subject ID}', 'github/{github username}', or 'github-id/{github numeric ID}'.
    role_assignment = sherlock_python_client.SherlockRoleAssignmentV3Edit() # SherlockRoleAssignmentV3Edit | The edits to make to the RoleAssignment

    try:
        # Edit a RoleAssignment
        api_response = api_instance.api_role_assignments_v3_role_selector_user_selector_patch(role_selector, user_selector, role_assignment)
        print("The response of RoleAssignmentsApi->api_role_assignments_v3_role_selector_user_selector_patch:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RoleAssignmentsApi->api_role_assignments_v3_role_selector_user_selector_patch: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role_selector** | **str**| The selector of the Role, which can be either the numeric ID or the name | 
 **user_selector** | **str**| The selector of the User, which can be either a numeric ID, the email, &#39;google-id/{google subject ID}&#39;, &#39;github/{github username}&#39;, or &#39;github-id/{github numeric ID}&#39;. | 
 **role_assignment** | [**SherlockRoleAssignmentV3Edit**](SherlockRoleAssignmentV3Edit.md)| The edits to make to the RoleAssignment | 

### Return type

[**SherlockRoleAssignmentV3**](SherlockRoleAssignmentV3.md)

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

# **api_role_assignments_v3_role_selector_user_selector_post**
> SherlockRoleAssignmentV3 api_role_assignments_v3_role_selector_user_selector_post(role_selector, user_selector, role_assignment)

Create a RoleAssignment

Create the RoleAssignment between a given Role and User.
Non-super-admins may only mutate RoleAssignments for themselves, only for roles they can break-glass into, and only with an expiry no further than the role's default break-glass duration in the future.
Propagation will be triggered after this operation.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_role_assignment_v3 import SherlockRoleAssignmentV3
from sherlock_python_client.models.sherlock_role_assignment_v3_edit import SherlockRoleAssignmentV3Edit
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
    api_instance = sherlock_python_client.RoleAssignmentsApi(api_client)
    role_selector = 'role_selector_example' # str | The selector of the Role, which can be either the numeric ID or the name
    user_selector = 'user_selector_example' # str | The selector of the User, which can be either a numeric ID, the email, 'google-id/{google subject ID}', 'github/{github username}', or 'github-id/{github numeric ID}'.
    role_assignment = sherlock_python_client.SherlockRoleAssignmentV3Edit() # SherlockRoleAssignmentV3Edit | The initial fields to set for the new RoleAssignment

    try:
        # Create a RoleAssignment
        api_response = api_instance.api_role_assignments_v3_role_selector_user_selector_post(role_selector, user_selector, role_assignment)
        print("The response of RoleAssignmentsApi->api_role_assignments_v3_role_selector_user_selector_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RoleAssignmentsApi->api_role_assignments_v3_role_selector_user_selector_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role_selector** | **str**| The selector of the Role, which can be either the numeric ID or the name | 
 **user_selector** | **str**| The selector of the User, which can be either a numeric ID, the email, &#39;google-id/{google subject ID}&#39;, &#39;github/{github username}&#39;, or &#39;github-id/{github numeric ID}&#39;. | 
 **role_assignment** | [**SherlockRoleAssignmentV3Edit**](SherlockRoleAssignmentV3Edit.md)| The initial fields to set for the new RoleAssignment | 

### Return type

[**SherlockRoleAssignmentV3**](SherlockRoleAssignmentV3.md)

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

