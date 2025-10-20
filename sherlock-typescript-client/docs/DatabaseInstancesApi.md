# DatabaseInstancesApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**apiDatabaseInstancesV3Get**](DatabaseInstancesApi.md#apidatabaseinstancesv3get) | **GET** /api/database-instances/v3 | List DatabaseInstances matching a filter |
| [**apiDatabaseInstancesV3Post**](DatabaseInstancesApi.md#apidatabaseinstancesv3post) | **POST** /api/database-instances/v3 | Create a DatabaseInstance |
| [**apiDatabaseInstancesV3Put**](DatabaseInstancesApi.md#apidatabaseinstancesv3put) | **PUT** /api/database-instances/v3 | Create or edit a DatabaseInstance |
| [**apiDatabaseInstancesV3SelectorDelete**](DatabaseInstancesApi.md#apidatabaseinstancesv3selectordelete) | **DELETE** /api/database-instances/v3/{selector} | Delete an individual DatabaseInstance |
| [**apiDatabaseInstancesV3SelectorGet**](DatabaseInstancesApi.md#apidatabaseinstancesv3selectorget) | **GET** /api/database-instances/v3/{selector} | Get an individual DatabaseInstance |
| [**apiDatabaseInstancesV3SelectorPatch**](DatabaseInstancesApi.md#apidatabaseinstancesv3selectorpatch) | **PATCH** /api/database-instances/v3/{selector} | Edit an individual DatabaseInstance |



## apiDatabaseInstancesV3Get

> Array&lt;SherlockDatabaseInstanceV3&gt; apiDatabaseInstancesV3Get(chartRelease, createdAt, defaultDatabase, googleProject, id, instanceName, platform, updatedAt, limit, offset)

List DatabaseInstances matching a filter

List DatabaseInstances matching a filter.

### Example

```ts
import {
  Configuration,
  DatabaseInstancesApi,
} from '@sherlock-js-client/sherlock';
import type { ApiDatabaseInstancesV3GetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new DatabaseInstancesApi();

  const body = {
    // string | Required when creating (optional)
    chartRelease: chartRelease_example,
    // Date (optional)
    createdAt: 2013-10-20T19:20:30+01:00,
    // string | When creating, defaults to the chart name (optional)
    defaultDatabase: defaultDatabase_example,
    // string | Required if platform is \'google\' (optional)
    googleProject: googleProject_example,
    // number (optional)
    id: 56,
    // string | Required if platform is \'google\' or \'azure\' (optional)
    instanceName: instanceName_example,
    // string | \'google\', \'azure\', or default \'kubernetes\' (optional)
    platform: platform_example,
    // Date (optional)
    updatedAt: 2013-10-20T19:20:30+01:00,
    // number | Control how many DatabaseInstances are returned (default 0, meaning all) (optional)
    limit: 56,
    // number | Control the offset for the returned DatabaseInstances (default 0) (optional)
    offset: 56,
  } satisfies ApiDatabaseInstancesV3GetRequest;

  try {
    const data = await api.apiDatabaseInstancesV3Get(body);
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
| **chartRelease** | `string` | Required when creating | [Optional] [Defaults to `undefined`] |
| **createdAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **defaultDatabase** | `string` | When creating, defaults to the chart name | [Optional] [Defaults to `undefined`] |
| **googleProject** | `string` | Required if platform is \&#39;google\&#39; | [Optional] [Defaults to `undefined`] |
| **id** | `number` |  | [Optional] [Defaults to `undefined`] |
| **instanceName** | `string` | Required if platform is \&#39;google\&#39; or \&#39;azure\&#39; | [Optional] [Defaults to `undefined`] |
| **platform** | `string` | \&#39;google\&#39;, \&#39;azure\&#39;, or default \&#39;kubernetes\&#39; | [Optional] [Defaults to `&#39;kubernetes&#39;`] |
| **updatedAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **limit** | `number` | Control how many DatabaseInstances are returned (default 0, meaning all) | [Optional] [Defaults to `undefined`] |
| **offset** | `number` | Control the offset for the returned DatabaseInstances (default 0) | [Optional] [Defaults to `undefined`] |

### Return type

[**Array&lt;SherlockDatabaseInstanceV3&gt;**](SherlockDatabaseInstanceV3.md)

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


## apiDatabaseInstancesV3Post

> SherlockDatabaseInstanceV3 apiDatabaseInstancesV3Post(databaseInstance)

Create a DatabaseInstance

Create a DatabaseInstance.

### Example

```ts
import {
  Configuration,
  DatabaseInstancesApi,
} from '@sherlock-js-client/sherlock';
import type { ApiDatabaseInstancesV3PostRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new DatabaseInstancesApi();

  const body = {
    // SherlockDatabaseInstanceV3Create | The DatabaseInstance to create
    databaseInstance: ...,
  } satisfies ApiDatabaseInstancesV3PostRequest;

  try {
    const data = await api.apiDatabaseInstancesV3Post(body);
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
| **databaseInstance** | [SherlockDatabaseInstanceV3Create](SherlockDatabaseInstanceV3Create.md) | The DatabaseInstance to create | |

### Return type

[**SherlockDatabaseInstanceV3**](SherlockDatabaseInstanceV3.md)

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


## apiDatabaseInstancesV3Put

> SherlockDatabaseInstanceV3 apiDatabaseInstancesV3Put(databaseInstance)

Create or edit a DatabaseInstance

Create or edit a DatabaseInstance, depending on whether one already exists for the chart release

### Example

```ts
import {
  Configuration,
  DatabaseInstancesApi,
} from '@sherlock-js-client/sherlock';
import type { ApiDatabaseInstancesV3PutRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new DatabaseInstancesApi();

  const body = {
    // SherlockDatabaseInstanceV3Create | The DatabaseInstance to create or edit. Defaults will only be set if creating.
    databaseInstance: ...,
  } satisfies ApiDatabaseInstancesV3PutRequest;

  try {
    const data = await api.apiDatabaseInstancesV3Put(body);
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
| **databaseInstance** | [SherlockDatabaseInstanceV3Create](SherlockDatabaseInstanceV3Create.md) | The DatabaseInstance to create or edit. Defaults will only be set if creating. | |

### Return type

[**SherlockDatabaseInstanceV3**](SherlockDatabaseInstanceV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **201** | Created |  -  |
| **400** | Bad Request |  -  |
| **403** | Forbidden |  -  |
| **404** | Not Found |  -  |
| **407** | Proxy Authentication Required |  -  |
| **409** | Conflict |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## apiDatabaseInstancesV3SelectorDelete

> SherlockDatabaseInstanceV3 apiDatabaseInstancesV3SelectorDelete(selector)

Delete an individual DatabaseInstance

Delete an individual DatabaseInstance by its selector.

### Example

```ts
import {
  Configuration,
  DatabaseInstancesApi,
} from '@sherlock-js-client/sherlock';
import type { ApiDatabaseInstancesV3SelectorDeleteRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new DatabaseInstancesApi();

  const body = {
    // string | The selector of the DatabaseInstance, which can be either a numeric ID or \'chart-release/\' followed by a chart release selector.
    selector: selector_example,
  } satisfies ApiDatabaseInstancesV3SelectorDeleteRequest;

  try {
    const data = await api.apiDatabaseInstancesV3SelectorDelete(body);
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
| **selector** | `string` | The selector of the DatabaseInstance, which can be either a numeric ID or \&#39;chart-release/\&#39; followed by a chart release selector. | [Defaults to `undefined`] |

### Return type

[**SherlockDatabaseInstanceV3**](SherlockDatabaseInstanceV3.md)

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


## apiDatabaseInstancesV3SelectorGet

> SherlockDatabaseInstanceV3 apiDatabaseInstancesV3SelectorGet(selector)

Get an individual DatabaseInstance

Get an individual DatabaseInstance by its selector.

### Example

```ts
import {
  Configuration,
  DatabaseInstancesApi,
} from '@sherlock-js-client/sherlock';
import type { ApiDatabaseInstancesV3SelectorGetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new DatabaseInstancesApi();

  const body = {
    // string | The selector of the DatabaseInstance, which can be either a numeric ID or \'chart-release/\' followed by a chart release selector.
    selector: selector_example,
  } satisfies ApiDatabaseInstancesV3SelectorGetRequest;

  try {
    const data = await api.apiDatabaseInstancesV3SelectorGet(body);
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
| **selector** | `string` | The selector of the DatabaseInstance, which can be either a numeric ID or \&#39;chart-release/\&#39; followed by a chart release selector. | [Defaults to `undefined`] |

### Return type

[**SherlockDatabaseInstanceV3**](SherlockDatabaseInstanceV3.md)

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


## apiDatabaseInstancesV3SelectorPatch

> SherlockDatabaseInstanceV3 apiDatabaseInstancesV3SelectorPatch(selector, databaseInstance)

Edit an individual DatabaseInstance

Edit an individual DatabaseInstance by its selector.

### Example

```ts
import {
  Configuration,
  DatabaseInstancesApi,
} from '@sherlock-js-client/sherlock';
import type { ApiDatabaseInstancesV3SelectorPatchRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new DatabaseInstancesApi();

  const body = {
    // string | The selector of the DatabaseInstance, which can be either a numeric ID or \'chart-release/\' followed by a chart release selector.
    selector: selector_example,
    // SherlockDatabaseInstanceV3Edit | The edits to make to the DatabaseInstance
    databaseInstance: ...,
  } satisfies ApiDatabaseInstancesV3SelectorPatchRequest;

  try {
    const data = await api.apiDatabaseInstancesV3SelectorPatch(body);
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
| **selector** | `string` | The selector of the DatabaseInstance, which can be either a numeric ID or \&#39;chart-release/\&#39; followed by a chart release selector. | [Defaults to `undefined`] |
| **databaseInstance** | [SherlockDatabaseInstanceV3Edit](SherlockDatabaseInstanceV3Edit.md) | The edits to make to the DatabaseInstance | |

### Return type

[**SherlockDatabaseInstanceV3**](SherlockDatabaseInstanceV3.md)

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

