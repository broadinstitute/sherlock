# RoleAssignmentsApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**apiRoleAssignmentsV3Get**](RoleAssignmentsApi.md#apiroleassignmentsv3get) | **GET** /api/role-assignments/v3 | List RoleAssignments matching a filter |
| [**apiRoleAssignmentsV3RoleSelectorUserSelectorDelete**](RoleAssignmentsApi.md#apiroleassignmentsv3roleselectoruserselectordelete) | **DELETE** /api/role-assignments/v3/{role-selector}/{user-selector} | Delete a RoleAssignment |
| [**apiRoleAssignmentsV3RoleSelectorUserSelectorGet**](RoleAssignmentsApi.md#apiroleassignmentsv3roleselectoruserselectorget) | **GET** /api/role-assignments/v3/{role-selector}/{user-selector} | Get a RoleAssignment |
| [**apiRoleAssignmentsV3RoleSelectorUserSelectorPatch**](RoleAssignmentsApi.md#apiroleassignmentsv3roleselectoruserselectorpatch) | **PATCH** /api/role-assignments/v3/{role-selector}/{user-selector} | Edit a RoleAssignment |
| [**apiRoleAssignmentsV3RoleSelectorUserSelectorPost**](RoleAssignmentsApi.md#apiroleassignmentsv3roleselectoruserselectorpost) | **POST** /api/role-assignments/v3/{role-selector}/{user-selector} | Create a RoleAssignment |



## apiRoleAssignmentsV3Get

> Array&lt;SherlockRoleAssignmentV3&gt; apiRoleAssignmentsV3Get(expiresAt, expiresIn, suspended, limit, offset)

List RoleAssignments matching a filter

List RoleAssignments matching a filter. The correct way to list RoleAssignments for a particular Role or User is to get that Role or User specifically, not to use this endpoint.

### Example

```ts
import {
  Configuration,
  RoleAssignmentsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiRoleAssignmentsV3GetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new RoleAssignmentsApi();

  const body = {
    // Date (optional)
    expiresAt: 2013-10-20T19:20:30+01:00,
    // string | A Go time.Duration string that will be added to the current time to attempt to set expiresAt (this may be more convenient than setting expiresAt directly) (optional)
    expiresIn: expiresIn_example,
    // boolean | If the assignment should be active. This field is only mutable through the API if the role doesn\'t automatically suspend non-suitable users (optional)
    suspended: true,
    // number | Control how many RoleAssignments are returned (default 0, no limit) (optional)
    limit: 56,
    // number | Control the offset for the returned RoleAssignments (default 0) (optional)
    offset: 56,
  } satisfies ApiRoleAssignmentsV3GetRequest;

  try {
    const data = await api.apiRoleAssignmentsV3Get(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **expiresAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **expiresIn** | `string` | A Go time.Duration string that will be added to the current time to attempt to set expiresAt (this may be more convenient than setting expiresAt directly) | [Optional] [Defaults to `undefined`] |
| **suspended** | `boolean` | If the assignment should be active. This field is only mutable through the API if the role doesn\&#39;t automatically suspend non-suitable users | [Optional] [Defaults to `undefined`] |
| **limit** | `number` | Control how many RoleAssignments are returned (default 0, no limit) | [Optional] [Defaults to `undefined`] |
| **offset** | `number` | Control the offset for the returned RoleAssignments (default 0) | [Optional] [Defaults to `undefined`] |

### Return type

[**Array&lt;SherlockRoleAssignmentV3&gt;**](SherlockRoleAssignmentV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Bad Request |  -  |
| **403** | Forbidden |  -  |
| **404** | Not Found |  -  |
| **407** | Proxy Authentication Required |  -  |
| **409** | Conflict |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## apiRoleAssignmentsV3RoleSelectorUserSelectorDelete

> apiRoleAssignmentsV3RoleSelectorUserSelectorDelete(roleSelector, userSelector)

Delete a RoleAssignment

Delete the RoleAssignment between a given Role and User. Non-super-admins may only mutate RoleAssignments for themselves, only for roles they can break-glass into, and only with an expiry no further than the role\&#39;s default break-glass duration in the future. Propagation will be triggered after this operation.

### Example

```ts
import {
  Configuration,
  RoleAssignmentsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiRoleAssignmentsV3RoleSelectorUserSelectorDeleteRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new RoleAssignmentsApi();

  const body = {
    // string | The selector of the Role, which can be either the numeric ID or the name
    roleSelector: roleSelector_example,
    // string | The selector of the User, which can be either a numeric ID, the email, \'google-id/{google subject ID}\', \'github/{github username}\', or \'github-id/{github numeric ID}\'.
    userSelector: userSelector_example,
  } satisfies ApiRoleAssignmentsV3RoleSelectorUserSelectorDeleteRequest;

  try {
    const data = await api.apiRoleAssignmentsV3RoleSelectorUserSelectorDelete(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **roleSelector** | `string` | The selector of the Role, which can be either the numeric ID or the name | [Defaults to `undefined`] |
| **userSelector** | `string` | The selector of the User, which can be either a numeric ID, the email, \&#39;google-id/{google subject ID}\&#39;, \&#39;github/{github username}\&#39;, or \&#39;github-id/{github numeric ID}\&#39;. | [Defaults to `undefined`] |

### Return type

`void` (Empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **400** | Bad Request |  -  |
| **403** | Forbidden |  -  |
| **404** | Not Found |  -  |
| **407** | Proxy Authentication Required |  -  |
| **409** | Conflict |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## apiRoleAssignmentsV3RoleSelectorUserSelectorGet

> SherlockRoleAssignmentV3 apiRoleAssignmentsV3RoleSelectorUserSelectorGet(roleSelector, userSelector)

Get a RoleAssignment

Get the RoleAssignment between a given Role and User.

### Example

```ts
import {
  Configuration,
  RoleAssignmentsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiRoleAssignmentsV3RoleSelectorUserSelectorGetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new RoleAssignmentsApi();

  const body = {
    // string | The selector of the Role, which can be either the numeric ID or the name
    roleSelector: roleSelector_example,
    // string | The selector of the User, which can be either a numeric ID, the email, \'google-id/{google subject ID}\', \'github/{github username}\', or \'github-id/{github numeric ID}\'.
    userSelector: userSelector_example,
  } satisfies ApiRoleAssignmentsV3RoleSelectorUserSelectorGetRequest;

  try {
    const data = await api.apiRoleAssignmentsV3RoleSelectorUserSelectorGet(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **roleSelector** | `string` | The selector of the Role, which can be either the numeric ID or the name | [Defaults to `undefined`] |
| **userSelector** | `string` | The selector of the User, which can be either a numeric ID, the email, \&#39;google-id/{google subject ID}\&#39;, \&#39;github/{github username}\&#39;, or \&#39;github-id/{github numeric ID}\&#39;. | [Defaults to `undefined`] |

### Return type

[**SherlockRoleAssignmentV3**](SherlockRoleAssignmentV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Bad Request |  -  |
| **403** | Forbidden |  -  |
| **404** | Not Found |  -  |
| **407** | Proxy Authentication Required |  -  |
| **409** | Conflict |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## apiRoleAssignmentsV3RoleSelectorUserSelectorPatch

> SherlockRoleAssignmentV3 apiRoleAssignmentsV3RoleSelectorUserSelectorPatch(roleSelector, userSelector, roleAssignment)

Edit a RoleAssignment

Edit the RoleAssignment between a given Role and User. Non-super-admins may only mutate RoleAssignments for themselves, only for roles they can break-glass into, and only with an expiry no further than the role\&#39;s default break-glass duration in the future. Propagation will be triggered after this operation.

### Example

```ts
import {
  Configuration,
  RoleAssignmentsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiRoleAssignmentsV3RoleSelectorUserSelectorPatchRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new RoleAssignmentsApi();

  const body = {
    // string | The selector of the Role, which can be either the numeric ID or the name
    roleSelector: roleSelector_example,
    // string | The selector of the User, which can be either a numeric ID, the email, \'google-id/{google subject ID}\', \'github/{github username}\', or \'github-id/{github numeric ID}\'.
    userSelector: userSelector_example,
    // SherlockRoleAssignmentV3Edit | The edits to make to the RoleAssignment
    roleAssignment: ...,
  } satisfies ApiRoleAssignmentsV3RoleSelectorUserSelectorPatchRequest;

  try {
    const data = await api.apiRoleAssignmentsV3RoleSelectorUserSelectorPatch(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **roleSelector** | `string` | The selector of the Role, which can be either the numeric ID or the name | [Defaults to `undefined`] |
| **userSelector** | `string` | The selector of the User, which can be either a numeric ID, the email, \&#39;google-id/{google subject ID}\&#39;, \&#39;github/{github username}\&#39;, or \&#39;github-id/{github numeric ID}\&#39;. | [Defaults to `undefined`] |
| **roleAssignment** | [SherlockRoleAssignmentV3Edit](SherlockRoleAssignmentV3Edit.md) | The edits to make to the RoleAssignment | |

### Return type

[**SherlockRoleAssignmentV3**](SherlockRoleAssignmentV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Bad Request |  -  |
| **403** | Forbidden |  -  |
| **404** | Not Found |  -  |
| **407** | Proxy Authentication Required |  -  |
| **409** | Conflict |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## apiRoleAssignmentsV3RoleSelectorUserSelectorPost

> SherlockRoleAssignmentV3 apiRoleAssignmentsV3RoleSelectorUserSelectorPost(roleSelector, userSelector, roleAssignment)

Create a RoleAssignment

Create the RoleAssignment between a given Role and User. Non-super-admins may only mutate RoleAssignments for themselves, only for roles they can break-glass into, and only with an expiry no further than the role\&#39;s default break-glass duration in the future. Propagation will be triggered after this operation.

### Example

```ts
import {
  Configuration,
  RoleAssignmentsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiRoleAssignmentsV3RoleSelectorUserSelectorPostRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new RoleAssignmentsApi();

  const body = {
    // string | The selector of the Role, which can be either the numeric ID or the name
    roleSelector: roleSelector_example,
    // string | The selector of the User, which can be either a numeric ID, the email, \'google-id/{google subject ID}\', \'github/{github username}\', or \'github-id/{github numeric ID}\'.
    userSelector: userSelector_example,
    // SherlockRoleAssignmentV3Edit | The initial fields to set for the new RoleAssignment
    roleAssignment: ...,
  } satisfies ApiRoleAssignmentsV3RoleSelectorUserSelectorPostRequest;

  try {
    const data = await api.apiRoleAssignmentsV3RoleSelectorUserSelectorPost(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **roleSelector** | `string` | The selector of the Role, which can be either the numeric ID or the name | [Defaults to `undefined`] |
| **userSelector** | `string` | The selector of the User, which can be either a numeric ID, the email, \&#39;google-id/{google subject ID}\&#39;, \&#39;github/{github username}\&#39;, or \&#39;github-id/{github numeric ID}\&#39;. | [Defaults to `undefined`] |
| **roleAssignment** | [SherlockRoleAssignmentV3Edit](SherlockRoleAssignmentV3Edit.md) | The initial fields to set for the new RoleAssignment | |

### Return type

[**SherlockRoleAssignmentV3**](SherlockRoleAssignmentV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **201** | Created |  -  |
| **400** | Bad Request |  -  |
| **403** | Forbidden |  -  |
| **404** | Not Found |  -  |
| **407** | Proxy Authentication Required |  -  |
| **409** | Conflict |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

