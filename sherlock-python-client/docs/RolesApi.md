# sherlock_python_client.RolesApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_roles_v3_get**](RolesApi.md#api_roles_v3_get) | **GET** /api/roles/v3 | List Roles matching a filter
[**api_roles_v3_post**](RolesApi.md#api_roles_v3_post) | **POST** /api/roles/v3 | Create a Role
[**api_roles_v3_selector_delete**](RolesApi.md#api_roles_v3_selector_delete) | **DELETE** /api/roles/v3/{selector} | Delete a Role
[**api_roles_v3_selector_get**](RolesApi.md#api_roles_v3_selector_get) | **GET** /api/roles/v3/{selector} | Get a Role
[**api_roles_v3_selector_patch**](RolesApi.md#api_roles_v3_selector_patch) | **PATCH** /api/roles/v3/{selector} | Edit a Role


# **api_roles_v3_get**
> List[SherlockRoleV3] api_roles_v3_get(auto_assign_all_users=auto_assign_all_users, can_be_glass_broken_by_role=can_be_glass_broken_by_role, created_at=created_at, default_glass_break_duration=default_glass_break_duration, grants_broad_institute_group=grants_broad_institute_group, grants_dev_azure_group=grants_dev_azure_group, grants_dev_firecloud_folder_owner=grants_dev_firecloud_folder_owner, grants_dev_firecloud_group=grants_dev_firecloud_group, grants_prod_azure_group=grants_prod_azure_group, grants_prod_firecloud_folder_owner=grants_prod_firecloud_folder_owner, grants_prod_firecloud_group=grants_prod_firecloud_group, grants_qa_firecloud_folder_owner=grants_qa_firecloud_folder_owner, grants_qa_firecloud_group=grants_qa_firecloud_group, grants_sherlock_super_admin=grants_sherlock_super_admin, id=id, name=name, suspend_non_suitable_users=suspend_non_suitable_users, updated_at=updated_at, limit=limit, offset=offset)

List Roles matching a filter

List Roles matching a filter.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_role_v3 import SherlockRoleV3
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
    api_instance = sherlock_python_client.RolesApi(api_client)
    auto_assign_all_users = True # bool | When true, Sherlock will automatically assign all users to this role who do not already have a role assignment (optional)
    can_be_glass_broken_by_role = 56 # int |  (optional)
    created_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    default_glass_break_duration = 'default_glass_break_duration_example' # str |  (optional)
    grants_broad_institute_group = 'grants_broad_institute_group_example' # str |  (optional)
    grants_dev_azure_group = 'grants_dev_azure_group_example' # str |  (optional)
    grants_dev_firecloud_folder_owner = 'grants_dev_firecloud_folder_owner_example' # str |  (optional)
    grants_dev_firecloud_group = 'grants_dev_firecloud_group_example' # str |  (optional)
    grants_prod_azure_group = 'grants_prod_azure_group_example' # str |  (optional)
    grants_prod_firecloud_folder_owner = 'grants_prod_firecloud_folder_owner_example' # str |  (optional)
    grants_prod_firecloud_group = 'grants_prod_firecloud_group_example' # str |  (optional)
    grants_qa_firecloud_folder_owner = 'grants_qa_firecloud_folder_owner_example' # str |  (optional)
    grants_qa_firecloud_group = 'grants_qa_firecloud_group_example' # str |  (optional)
    grants_sherlock_super_admin = True # bool |  (optional)
    id = 56 # int |  (optional)
    name = 'name_example' # str |  (optional)
    suspend_non_suitable_users = True # bool | When true, the \"suspended\" field on role assignments will be computed by Sherlock based on suitability instead of being a mutable API field (optional)
    updated_at = '2013-10-20T19:20:30+01:00' # datetime |  (optional)
    limit = 56 # int | Control how many Roles are returned (default 0, no limit) (optional)
    offset = 56 # int | Control the offset for the returned Roles (default 0) (optional)

    try:
        # List Roles matching a filter
        api_response = api_instance.api_roles_v3_get(auto_assign_all_users=auto_assign_all_users, can_be_glass_broken_by_role=can_be_glass_broken_by_role, created_at=created_at, default_glass_break_duration=default_glass_break_duration, grants_broad_institute_group=grants_broad_institute_group, grants_dev_azure_group=grants_dev_azure_group, grants_dev_firecloud_folder_owner=grants_dev_firecloud_folder_owner, grants_dev_firecloud_group=grants_dev_firecloud_group, grants_prod_azure_group=grants_prod_azure_group, grants_prod_firecloud_folder_owner=grants_prod_firecloud_folder_owner, grants_prod_firecloud_group=grants_prod_firecloud_group, grants_qa_firecloud_folder_owner=grants_qa_firecloud_folder_owner, grants_qa_firecloud_group=grants_qa_firecloud_group, grants_sherlock_super_admin=grants_sherlock_super_admin, id=id, name=name, suspend_non_suitable_users=suspend_non_suitable_users, updated_at=updated_at, limit=limit, offset=offset)
        print("The response of RolesApi->api_roles_v3_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RolesApi->api_roles_v3_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **auto_assign_all_users** | **bool**| When true, Sherlock will automatically assign all users to this role who do not already have a role assignment | [optional] 
 **can_be_glass_broken_by_role** | **int**|  | [optional] 
 **created_at** | **datetime**|  | [optional] 
 **default_glass_break_duration** | **str**|  | [optional] 
 **grants_broad_institute_group** | **str**|  | [optional] 
 **grants_dev_azure_group** | **str**|  | [optional] 
 **grants_dev_firecloud_folder_owner** | **str**|  | [optional] 
 **grants_dev_firecloud_group** | **str**|  | [optional] 
 **grants_prod_azure_group** | **str**|  | [optional] 
 **grants_prod_firecloud_folder_owner** | **str**|  | [optional] 
 **grants_prod_firecloud_group** | **str**|  | [optional] 
 **grants_qa_firecloud_folder_owner** | **str**|  | [optional] 
 **grants_qa_firecloud_group** | **str**|  | [optional] 
 **grants_sherlock_super_admin** | **bool**|  | [optional] 
 **id** | **int**|  | [optional] 
 **name** | **str**|  | [optional] 
 **suspend_non_suitable_users** | **bool**| When true, the \&quot;suspended\&quot; field on role assignments will be computed by Sherlock based on suitability instead of being a mutable API field | [optional] 
 **updated_at** | **datetime**|  | [optional] 
 **limit** | **int**| Control how many Roles are returned (default 0, no limit) | [optional] 
 **offset** | **int**| Control the offset for the returned Roles (default 0) | [optional] 

### Return type

[**List[SherlockRoleV3]**](SherlockRoleV3.md)

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

# **api_roles_v3_post**
> SherlockRoleV3 api_roles_v3_post(role)

Create a Role

Create an individual Role with no one assigned to it. Only super-admins may mutate Roles. Propagation will be triggered after this operation.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_role_v3 import SherlockRoleV3
from sherlock_python_client.models.sherlock_role_v3_edit import SherlockRoleV3Edit
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
    api_instance = sherlock_python_client.RolesApi(api_client)
    role = sherlock_python_client.SherlockRoleV3Edit() # SherlockRoleV3Edit | The initial fields the Role should have set

    try:
        # Create a Role
        api_response = api_instance.api_roles_v3_post(role)
        print("The response of RolesApi->api_roles_v3_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RolesApi->api_roles_v3_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role** | [**SherlockRoleV3Edit**](SherlockRoleV3Edit.md)| The initial fields the Role should have set | 

### Return type

[**SherlockRoleV3**](SherlockRoleV3.md)

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

# **api_roles_v3_selector_delete**
> SherlockRoleV3 api_roles_v3_selector_delete(selector)

Delete a Role

Delete an individual Role. Only super-admins may mutate Roles. Propagation will NOT be triggered after this operation -- the grants will become un-managed by Sherlock and left as-is. Remove role assignments first to remove users from grants.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_role_v3 import SherlockRoleV3
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
    api_instance = sherlock_python_client.RolesApi(api_client)
    selector = 'selector_example' # str | The selector of the Role, which can be either the numeric ID or the name

    try:
        # Delete a Role
        api_response = api_instance.api_roles_v3_selector_delete(selector)
        print("The response of RolesApi->api_roles_v3_selector_delete:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RolesApi->api_roles_v3_selector_delete: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the Role, which can be either the numeric ID or the name | 

### Return type

[**SherlockRoleV3**](SherlockRoleV3.md)

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

# **api_roles_v3_selector_get**
> SherlockRoleV3 api_roles_v3_selector_get(selector)

Get a Role

Get an individual Role and the Users assigned to it.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_role_v3 import SherlockRoleV3
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
    api_instance = sherlock_python_client.RolesApi(api_client)
    selector = 'selector_example' # str | The selector of the Role, which can be either the numeric ID or the name

    try:
        # Get a Role
        api_response = api_instance.api_roles_v3_selector_get(selector)
        print("The response of RolesApi->api_roles_v3_selector_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RolesApi->api_roles_v3_selector_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the Role, which can be either the numeric ID or the name | 

### Return type

[**SherlockRoleV3**](SherlockRoleV3.md)

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

# **api_roles_v3_selector_patch**
> SherlockRoleV3 api_roles_v3_selector_patch(selector, role)

Edit a Role

Edit an individual Role. Only super-admins may mutate Roles. Propagation will be triggered after this operation.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_role_v3 import SherlockRoleV3
from sherlock_python_client.models.sherlock_role_v3_edit import SherlockRoleV3Edit
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
    api_instance = sherlock_python_client.RolesApi(api_client)
    selector = 'selector_example' # str | The selector of the Role, which can be either the numeric ID or the name
    role = sherlock_python_client.SherlockRoleV3Edit() # SherlockRoleV3Edit | The edits to make to the Role

    try:
        # Edit a Role
        api_response = api_instance.api_roles_v3_selector_patch(selector, role)
        print("The response of RolesApi->api_roles_v3_selector_patch:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RolesApi->api_roles_v3_selector_patch: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selector** | **str**| The selector of the Role, which can be either the numeric ID or the name | 
 **role** | [**SherlockRoleV3Edit**](SherlockRoleV3Edit.md)| The edits to make to the Role | 

### Return type

[**SherlockRoleV3**](SherlockRoleV3.md)

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

