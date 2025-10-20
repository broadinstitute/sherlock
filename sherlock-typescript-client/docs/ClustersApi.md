# ClustersApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**apiClustersV3Get**](ClustersApi.md#apiclustersv3get) | **GET** /api/clusters/v3 | List Clusters matching a filter |
| [**apiClustersV3Post**](ClustersApi.md#apiclustersv3post) | **POST** /api/clusters/v3 | Create a Cluster |
| [**apiClustersV3SelectorDelete**](ClustersApi.md#apiclustersv3selectordelete) | **DELETE** /api/clusters/v3/{selector} | Delete an individual Cluster |
| [**apiClustersV3SelectorGet**](ClustersApi.md#apiclustersv3selectorget) | **GET** /api/clusters/v3/{selector} | Get an individual Cluster |
| [**apiClustersV3SelectorPatch**](ClustersApi.md#apiclustersv3selectorpatch) | **PATCH** /api/clusters/v3/{selector} | Edit an individual Cluster |



## apiClustersV3Get

> Array&lt;SherlockClusterV3&gt; apiClustersV3Get(address, azureSubscription, base, createdAt, googleProject, helmfileRef, id, location, name, provider, requiredRole, requiresSuitability, updatedAt, limit, offset)

List Clusters matching a filter

List Clusters matching a filter.

### Example

```ts
import {
  Configuration,
  ClustersApi,
} from '@sherlock-js-client/sherlock';
import type { ApiClustersV3GetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ClustersApi();

  const body = {
    // string | Required when creating (optional)
    address: address_example,
    // string | Required when creating if provider is \'azure\' (optional)
    azureSubscription: azureSubscription_example,
    // string | Required when creating (optional)
    base: base_example,
    // Date (optional)
    createdAt: 2013-10-20T19:20:30+01:00,
    // string | Required when creating if provider is \'google\' (optional)
    googleProject: googleProject_example,
    // string (optional)
    helmfileRef: helmfileRef_example,
    // number (optional)
    id: 56,
    // string (optional)
    location: location_example,
    // string | Required when creating (optional)
    name: name_example,
    // 'google' | 'azure' (optional)
    provider: provider_example,
    // string | If present, requires membership in the given role for mutations. Set to an empty string to clear. (optional)
    requiredRole: requiredRole_example,
    // boolean (optional)
    requiresSuitability: true,
    // Date (optional)
    updatedAt: 2013-10-20T19:20:30+01:00,
    // number | Control how many Clusters are returned (default 0, meaning all) (optional)
    limit: 56,
    // number | Control the offset for the returned Clusters (default 0) (optional)
    offset: 56,
  } satisfies ApiClustersV3GetRequest;

  try {
    const data = await api.apiClustersV3Get(body);
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
| **address** | `string` | Required when creating | [Optional] [Defaults to `undefined`] |
| **azureSubscription** | `string` | Required when creating if provider is \&#39;azure\&#39; | [Optional] [Defaults to `undefined`] |
| **base** | `string` | Required when creating | [Optional] [Defaults to `undefined`] |
| **createdAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **googleProject** | `string` | Required when creating if provider is \&#39;google\&#39; | [Optional] [Defaults to `undefined`] |
| **helmfileRef** | `string` |  | [Optional] [Defaults to `&#39;HEAD&#39;`] |
| **id** | `number` |  | [Optional] [Defaults to `undefined`] |
| **location** | `string` |  | [Optional] [Defaults to `&#39;us-central1-a&#39;`] |
| **name** | `string` | Required when creating | [Optional] [Defaults to `undefined`] |
| **provider** | `google`, `azure` |  | [Optional] [Defaults to `&#39;google&#39;`] [Enum: google, azure] |
| **requiredRole** | `string` | If present, requires membership in the given role for mutations. Set to an empty string to clear. | [Optional] [Defaults to `undefined`] |
| **requiresSuitability** | `boolean` |  | [Optional] [Defaults to `undefined`] |
| **updatedAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **limit** | `number` | Control how many Clusters are returned (default 0, meaning all) | [Optional] [Defaults to `undefined`] |
| **offset** | `number` | Control the offset for the returned Clusters (default 0) | [Optional] [Defaults to `undefined`] |

### Return type

[**Array&lt;SherlockClusterV3&gt;**](SherlockClusterV3.md)

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


## apiClustersV3Post

> SherlockClusterV3 apiClustersV3Post(cluster)

Create a Cluster

Create a Cluster.

### Example

```ts
import {
  Configuration,
  ClustersApi,
} from '@sherlock-js-client/sherlock';
import type { ApiClustersV3PostRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ClustersApi();

  const body = {
    // SherlockClusterV3Create | The Cluster to create
    cluster: ...,
  } satisfies ApiClustersV3PostRequest;

  try {
    const data = await api.apiClustersV3Post(body);
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
| **cluster** | [SherlockClusterV3Create](SherlockClusterV3Create.md) | The Cluster to create | |

### Return type

[**SherlockClusterV3**](SherlockClusterV3.md)

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


## apiClustersV3SelectorDelete

> SherlockClusterV3 apiClustersV3SelectorDelete(selector)

Delete an individual Cluster

Delete an individual Cluster by its ID.

### Example

```ts
import {
  Configuration,
  ClustersApi,
} from '@sherlock-js-client/sherlock';
import type { ApiClustersV3SelectorDeleteRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ClustersApi();

  const body = {
    // string | The selector of the Cluster, which can be either a numeric ID or the name.
    selector: selector_example,
  } satisfies ApiClustersV3SelectorDeleteRequest;

  try {
    const data = await api.apiClustersV3SelectorDelete(body);
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
| **selector** | `string` | The selector of the Cluster, which can be either a numeric ID or the name. | [Defaults to `undefined`] |

### Return type

[**SherlockClusterV3**](SherlockClusterV3.md)

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


## apiClustersV3SelectorGet

> SherlockClusterV3 apiClustersV3SelectorGet(selector)

Get an individual Cluster

Get an individual Cluster.

### Example

```ts
import {
  Configuration,
  ClustersApi,
} from '@sherlock-js-client/sherlock';
import type { ApiClustersV3SelectorGetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ClustersApi();

  const body = {
    // string | The selector of the Cluster, which can be either a numeric ID or the name.
    selector: selector_example,
  } satisfies ApiClustersV3SelectorGetRequest;

  try {
    const data = await api.apiClustersV3SelectorGet(body);
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
| **selector** | `string` | The selector of the Cluster, which can be either a numeric ID or the name. | [Defaults to `undefined`] |

### Return type

[**SherlockClusterV3**](SherlockClusterV3.md)

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


## apiClustersV3SelectorPatch

> SherlockClusterV3 apiClustersV3SelectorPatch(selector, cluster)

Edit an individual Cluster

Edit an individual Cluster.

### Example

```ts
import {
  Configuration,
  ClustersApi,
} from '@sherlock-js-client/sherlock';
import type { ApiClustersV3SelectorPatchRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ClustersApi();

  const body = {
    // string | The selector of the Cluster, which can be either a numeric ID or the name.
    selector: selector_example,
    // SherlockClusterV3Edit | The edits to make to the Cluster
    cluster: ...,
  } satisfies ApiClustersV3SelectorPatchRequest;

  try {
    const data = await api.apiClustersV3SelectorPatch(body);
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
| **selector** | `string` | The selector of the Cluster, which can be either a numeric ID or the name. | [Defaults to `undefined`] |
| **cluster** | [SherlockClusterV3Edit](SherlockClusterV3Edit.md) | The edits to make to the Cluster | |

### Return type

[**SherlockClusterV3**](SherlockClusterV3.md)

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

