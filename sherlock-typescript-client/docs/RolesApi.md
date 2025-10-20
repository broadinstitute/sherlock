# RolesApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**apiRolesV3Get**](RolesApi.md#apirolesv3get) | **GET** /api/roles/v3 | List Roles matching a filter |
| [**apiRolesV3Post**](RolesApi.md#apirolesv3post) | **POST** /api/roles/v3 | Create a Role |
| [**apiRolesV3SelectorDelete**](RolesApi.md#apirolesv3selectordelete) | **DELETE** /api/roles/v3/{selector} | Delete a Role |
| [**apiRolesV3SelectorGet**](RolesApi.md#apirolesv3selectorget) | **GET** /api/roles/v3/{selector} | Get a Role |
| [**apiRolesV3SelectorPatch**](RolesApi.md#apirolesv3selectorpatch) | **PATCH** /api/roles/v3/{selector} | Edit a Role |



## apiRolesV3Get

> Array&lt;SherlockRoleV3&gt; apiRolesV3Get(autoAssignAllUsers, canBeGlassBrokenByRole, createdAt, defaultGlassBreakDuration, grantsBroadInstituteGroup, grantsDevAzureAccount, grantsDevAzureDirectoryRoles, grantsDevAzureGroup, grantsDevFirecloudFolderOwner, grantsDevFirecloudGroup, grantsProdAzureAccount, grantsProdAzureDirectoryRoles, grantsProdAzureGroup, grantsProdFirecloudFolderOwner, grantsProdFirecloudGroup, grantsQaFirecloudFolderOwner, grantsQaFirecloudGroup, grantsSherlockSuperAdmin, id, name, suspendNonSuitableUsers, updatedAt, limit, offset)

List Roles matching a filter

List Roles matching a filter.

### Example

```ts
import {
  Configuration,
  RolesApi,
} from '@sherlock-js-client/sherlock';
import type { ApiRolesV3GetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new RolesApi();

  const body = {
    // boolean | When true, Sherlock will automatically assign all users to this role who do not already have a role assignment (optional)
    autoAssignAllUsers: true,
    // number (optional)
    canBeGlassBrokenByRole: 56,
    // Date (optional)
    createdAt: 2013-10-20T19:20:30+01:00,
    // string (optional)
    defaultGlassBreakDuration: defaultGlassBreakDuration_example,
    // string (optional)
    grantsBroadInstituteGroup: grantsBroadInstituteGroup_example,
    // boolean (optional)
    grantsDevAzureAccount: true,
    // boolean (optional)
    grantsDevAzureDirectoryRoles: true,
    // string (optional)
    grantsDevAzureGroup: grantsDevAzureGroup_example,
    // string (optional)
    grantsDevFirecloudFolderOwner: grantsDevFirecloudFolderOwner_example,
    // string (optional)
    grantsDevFirecloudGroup: grantsDevFirecloudGroup_example,
    // boolean (optional)
    grantsProdAzureAccount: true,
    // boolean (optional)
    grantsProdAzureDirectoryRoles: true,
    // string (optional)
    grantsProdAzureGroup: grantsProdAzureGroup_example,
    // string (optional)
    grantsProdFirecloudFolderOwner: grantsProdFirecloudFolderOwner_example,
    // string (optional)
    grantsProdFirecloudGroup: grantsProdFirecloudGroup_example,
    // string (optional)
    grantsQaFirecloudFolderOwner: grantsQaFirecloudFolderOwner_example,
    // string (optional)
    grantsQaFirecloudGroup: grantsQaFirecloudGroup_example,
    // boolean (optional)
    grantsSherlockSuperAdmin: true,
    // number (optional)
    id: 56,
    // string (optional)
    name: name_example,
    // boolean | When true, the \"suspended\" field on role assignments will be computed by Sherlock based on suitability instead of being a mutable API field (optional)
    suspendNonSuitableUsers: true,
    // Date (optional)
    updatedAt: 2013-10-20T19:20:30+01:00,
    // number | Control how many Roles are returned (default 0, no limit) (optional)
    limit: 56,
    // number | Control the offset for the returned Roles (default 0) (optional)
    offset: 56,
  } satisfies ApiRolesV3GetRequest;

  try {
    const data = await api.apiRolesV3Get(body);
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
| **autoAssignAllUsers** | `boolean` | When true, Sherlock will automatically assign all users to this role who do not already have a role assignment | [Optional] [Defaults to `undefined`] |
| **canBeGlassBrokenByRole** | `number` |  | [Optional] [Defaults to `undefined`] |
| **createdAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **defaultGlassBreakDuration** | `string` |  | [Optional] [Defaults to `undefined`] |
| **grantsBroadInstituteGroup** | `string` |  | [Optional] [Defaults to `undefined`] |
| **grantsDevAzureAccount** | `boolean` |  | [Optional] [Defaults to `undefined`] |
| **grantsDevAzureDirectoryRoles** | `boolean` |  | [Optional] [Defaults to `undefined`] |
| **grantsDevAzureGroup** | `string` |  | [Optional] [Defaults to `undefined`] |
| **grantsDevFirecloudFolderOwner** | `string` |  | [Optional] [Defaults to `undefined`] |
| **grantsDevFirecloudGroup** | `string` |  | [Optional] [Defaults to `undefined`] |
| **grantsProdAzureAccount** | `boolean` |  | [Optional] [Defaults to `undefined`] |
| **grantsProdAzureDirectoryRoles** | `boolean` |  | [Optional] [Defaults to `undefined`] |
| **grantsProdAzureGroup** | `string` |  | [Optional] [Defaults to `undefined`] |
| **grantsProdFirecloudFolderOwner** | `string` |  | [Optional] [Defaults to `undefined`] |
| **grantsProdFirecloudGroup** | `string` |  | [Optional] [Defaults to `undefined`] |
| **grantsQaFirecloudFolderOwner** | `string` |  | [Optional] [Defaults to `undefined`] |
| **grantsQaFirecloudGroup** | `string` |  | [Optional] [Defaults to `undefined`] |
| **grantsSherlockSuperAdmin** | `boolean` |  | [Optional] [Defaults to `undefined`] |
| **id** | `number` |  | [Optional] [Defaults to `undefined`] |
| **name** | `string` |  | [Optional] [Defaults to `undefined`] |
| **suspendNonSuitableUsers** | `boolean` | When true, the \&quot;suspended\&quot; field on role assignments will be computed by Sherlock based on suitability instead of being a mutable API field | [Optional] [Defaults to `undefined`] |
| **updatedAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **limit** | `number` | Control how many Roles are returned (default 0, no limit) | [Optional] [Defaults to `undefined`] |
| **offset** | `number` | Control the offset for the returned Roles (default 0) | [Optional] [Defaults to `undefined`] |

### Return type

[**Array&lt;SherlockRoleV3&gt;**](SherlockRoleV3.md)

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


## apiRolesV3Post

> SherlockRoleV3 apiRolesV3Post(role)

Create a Role

Create an individual Role with no one assigned to it. Only super-admins may mutate Roles. Propagation will be triggered after this operation.

### Example

```ts
import {
  Configuration,
  RolesApi,
} from '@sherlock-js-client/sherlock';
import type { ApiRolesV3PostRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new RolesApi();

  const body = {
    // SherlockRoleV3Edit | The initial fields the Role should have set
    role: ...,
  } satisfies ApiRolesV3PostRequest;

  try {
    const data = await api.apiRolesV3Post(body);
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
| **role** | [SherlockRoleV3Edit](SherlockRoleV3Edit.md) | The initial fields the Role should have set | |

### Return type

[**SherlockRoleV3**](SherlockRoleV3.md)

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


## apiRolesV3SelectorDelete

> SherlockRoleV3 apiRolesV3SelectorDelete(selector)

Delete a Role

Delete an individual Role. Only super-admins may mutate Roles. Propagation will NOT be triggered after this operation -- the grants will become un-managed by Sherlock and left as-is. Remove role assignments first to remove users from grants.

### Example

```ts
import {
  Configuration,
  RolesApi,
} from '@sherlock-js-client/sherlock';
import type { ApiRolesV3SelectorDeleteRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new RolesApi();

  const body = {
    // string | The selector of the Role, which can be either the numeric ID or the name
    selector: selector_example,
  } satisfies ApiRolesV3SelectorDeleteRequest;

  try {
    const data = await api.apiRolesV3SelectorDelete(body);
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
| **selector** | `string` | The selector of the Role, which can be either the numeric ID or the name | [Defaults to `undefined`] |

### Return type

[**SherlockRoleV3**](SherlockRoleV3.md)

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


## apiRolesV3SelectorGet

> SherlockRoleV3 apiRolesV3SelectorGet(selector)

Get a Role

Get an individual Role and the Users assigned to it.

### Example

```ts
import {
  Configuration,
  RolesApi,
} from '@sherlock-js-client/sherlock';
import type { ApiRolesV3SelectorGetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new RolesApi();

  const body = {
    // string | The selector of the Role, which can be either the numeric ID or the name
    selector: selector_example,
  } satisfies ApiRolesV3SelectorGetRequest;

  try {
    const data = await api.apiRolesV3SelectorGet(body);
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
| **selector** | `string` | The selector of the Role, which can be either the numeric ID or the name | [Defaults to `undefined`] |

### Return type

[**SherlockRoleV3**](SherlockRoleV3.md)

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


## apiRolesV3SelectorPatch

> SherlockRoleV3 apiRolesV3SelectorPatch(selector, role)

Edit a Role

Edit an individual Role. Only super-admins may mutate Roles. Propagation will be triggered after this operation.

### Example

```ts
import {
  Configuration,
  RolesApi,
} from '@sherlock-js-client/sherlock';
import type { ApiRolesV3SelectorPatchRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new RolesApi();

  const body = {
    // string | The selector of the Role, which can be either the numeric ID or the name
    selector: selector_example,
    // SherlockRoleV3Edit | The edits to make to the Role
    role: ...,
  } satisfies ApiRolesV3SelectorPatchRequest;

  try {
    const data = await api.apiRolesV3SelectorPatch(body);
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
| **selector** | `string` | The selector of the Role, which can be either the numeric ID or the name | [Defaults to `undefined`] |
| **role** | [SherlockRoleV3Edit](SherlockRoleV3Edit.md) | The edits to make to the Role | |

### Return type

[**SherlockRoleV3**](SherlockRoleV3.md)

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

