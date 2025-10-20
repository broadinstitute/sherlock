# CiIdentifiersApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**apiCiIdentifiersV3Get**](CiIdentifiersApi.md#apiciidentifiersv3get) | **GET** /api/ci-identifiers/v3 | List CiIdentifiers matching a filter |
| [**apiCiIdentifiersV3SelectorGet**](CiIdentifiersApi.md#apiciidentifiersv3selectorget) | **GET** /api/ci-identifiers/v3/{selector} | Get CiRuns for a resource by its CiIdentifier |



## apiCiIdentifiersV3Get

> Array&lt;SherlockCiIdentifierV3&gt; apiCiIdentifiersV3Get(createdAt, id, resourceID, resourceStatus, resourceType, updatedAt, limit, offset)

List CiIdentifiers matching a filter

List CiIdentifiers matching a filter. The CiRuns would have to re-queried directly to load the CiRuns. This is mainly helpful for debugging and directly querying the existence of a CiIdentifier. Results are ordered by creation date, starting at most recent.

### Example

```ts
import {
  Configuration,
  CiIdentifiersApi,
} from '@sherlock-js-client/sherlock';
import type { ApiCiIdentifiersV3GetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new CiIdentifiersApi();

  const body = {
    // Date (optional)
    createdAt: 2013-10-20T19:20:30+01:00,
    // number (optional)
    id: 56,
    // number (optional)
    resourceID: 56,
    // string | Available only when querying a CiIdentifier via a CiRun, indicates the status of the run for that resource (optional)
    resourceStatus: resourceStatus_example,
    // string (optional)
    resourceType: resourceType_example,
    // Date (optional)
    updatedAt: 2013-10-20T19:20:30+01:00,
    // number | Control how many CiIdentifiers are returned (default 100) (optional)
    limit: 56,
    // number | Control the offset for the returned CiIdentifiers (default 0) (optional)
    offset: 56,
  } satisfies ApiCiIdentifiersV3GetRequest;

  try {
    const data = await api.apiCiIdentifiersV3Get(body);
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
| **createdAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **id** | `number` |  | [Optional] [Defaults to `undefined`] |
| **resourceID** | `number` |  | [Optional] [Defaults to `undefined`] |
| **resourceStatus** | `string` | Available only when querying a CiIdentifier via a CiRun, indicates the status of the run for that resource | [Optional] [Defaults to `undefined`] |
| **resourceType** | `string` |  | [Optional] [Defaults to `undefined`] |
| **updatedAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **limit** | `number` | Control how many CiIdentifiers are returned (default 100) | [Optional] [Defaults to `undefined`] |
| **offset** | `number` | Control the offset for the returned CiIdentifiers (default 0) | [Optional] [Defaults to `undefined`] |

### Return type

[**Array&lt;SherlockCiIdentifierV3&gt;**](SherlockCiIdentifierV3.md)

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


## apiCiIdentifiersV3SelectorGet

> SherlockCiIdentifierV3 apiCiIdentifiersV3SelectorGet(selector, limitCiRuns, offsetCiRuns, allowStubCiRuns)

Get CiRuns for a resource by its CiIdentifier

Get CiRuns for a resource by its CiIdentifier, which can be referenced by \&#39;{type}/{selector...}\&#39;.

### Example

```ts
import {
  Configuration,
  CiIdentifiersApi,
} from '@sherlock-js-client/sherlock';
import type { ApiCiIdentifiersV3SelectorGetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new CiIdentifiersApi();

  const body = {
    // string | The selector of CiIdentifier, which can be referenced either by numeric ID or indirectly by \'{type}/{selector...}\'
    selector: selector_example,
    // number | Control how many CiRuns are returned (default 10) (optional)
    limitCiRuns: 56,
    // number | Control the offset for the returned CiRuns (default 0) (optional)
    offsetCiRuns: 56,
    // boolean | Allow stub CiRuns potentially lacking fields like status or startedAt to be returned (default false) (optional)
    allowStubCiRuns: true,
  } satisfies ApiCiIdentifiersV3SelectorGetRequest;

  try {
    const data = await api.apiCiIdentifiersV3SelectorGet(body);
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
| **selector** | `string` | The selector of CiIdentifier, which can be referenced either by numeric ID or indirectly by \&#39;{type}/{selector...}\&#39; | [Defaults to `undefined`] |
| **limitCiRuns** | `number` | Control how many CiRuns are returned (default 10) | [Optional] [Defaults to `undefined`] |
| **offsetCiRuns** | `number` | Control the offset for the returned CiRuns (default 0) | [Optional] [Defaults to `undefined`] |
| **allowStubCiRuns** | `boolean` | Allow stub CiRuns potentially lacking fields like status or startedAt to be returned (default false) | [Optional] [Defaults to `undefined`] |

### Return type

[**SherlockCiIdentifierV3**](SherlockCiIdentifierV3.md)

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

