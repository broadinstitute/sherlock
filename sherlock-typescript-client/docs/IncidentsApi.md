# IncidentsApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**apiIncidentsV3Get**](IncidentsApi.md#apiincidentsv3get) | **GET** /api/incidents/v3 | List Incidents matching a filter |
| [**apiIncidentsV3Post**](IncidentsApi.md#apiincidentsv3post) | **POST** /api/incidents/v3 | Create a Incident |
| [**apiIncidentsV3SelectorDelete**](IncidentsApi.md#apiincidentsv3selectordelete) | **DELETE** /api/incidents/v3/{selector} | Delete an individual Incident |
| [**apiIncidentsV3SelectorGet**](IncidentsApi.md#apiincidentsv3selectorget) | **GET** /api/incidents/v3/{selector} | Get an individual Incident |
| [**apiIncidentsV3SelectorPatch**](IncidentsApi.md#apiincidentsv3selectorpatch) | **PATCH** /api/incidents/v3/{selector} | Edit an individual Incident |



## apiIncidentsV3Get

> Array&lt;SherlockIncidentV3&gt; apiIncidentsV3Get(createdAt, description, id, remediatedAt, reviewCompletedAt, startedAt, ticket, updatedAt, limit, offset)

List Incidents matching a filter

List Incidents matching a filter.

### Example

```ts
import {
  Configuration,
  IncidentsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiIncidentsV3GetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new IncidentsApi();

  const body = {
    // Date (optional)
    createdAt: 2013-10-20T19:20:30+01:00,
    // string (optional)
    description: description_example,
    // number (optional)
    id: 56,
    // string (optional)
    remediatedAt: remediatedAt_example,
    // string (optional)
    reviewCompletedAt: reviewCompletedAt_example,
    // string (optional)
    startedAt: startedAt_example,
    // string (optional)
    ticket: ticket_example,
    // Date (optional)
    updatedAt: 2013-10-20T19:20:30+01:00,
    // number | Control how many Incidents are returned (default 0, meaning all) (optional)
    limit: 56,
    // number | Control the offset for the returned Incidents (default 0) (optional)
    offset: 56,
  } satisfies ApiIncidentsV3GetRequest;

  try {
    const data = await api.apiIncidentsV3Get(body);
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
| **description** | `string` |  | [Optional] [Defaults to `undefined`] |
| **id** | `number` |  | [Optional] [Defaults to `undefined`] |
| **remediatedAt** | `string` |  | [Optional] [Defaults to `undefined`] |
| **reviewCompletedAt** | `string` |  | [Optional] [Defaults to `undefined`] |
| **startedAt** | `string` |  | [Optional] [Defaults to `undefined`] |
| **ticket** | `string` |  | [Optional] [Defaults to `undefined`] |
| **updatedAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **limit** | `number` | Control how many Incidents are returned (default 0, meaning all) | [Optional] [Defaults to `undefined`] |
| **offset** | `number` | Control the offset for the returned Incidents (default 0) | [Optional] [Defaults to `undefined`] |

### Return type

[**Array&lt;SherlockIncidentV3&gt;**](SherlockIncidentV3.md)

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


## apiIncidentsV3Post

> SherlockIncidentV3 apiIncidentsV3Post(incident)

Create a Incident

Create a Incident.

### Example

```ts
import {
  Configuration,
  IncidentsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiIncidentsV3PostRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new IncidentsApi();

  const body = {
    // SherlockIncidentV3Create | The Incident to create
    incident: ...,
  } satisfies ApiIncidentsV3PostRequest;

  try {
    const data = await api.apiIncidentsV3Post(body);
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
| **incident** | [SherlockIncidentV3Create](SherlockIncidentV3Create.md) | The Incident to create | |

### Return type

[**SherlockIncidentV3**](SherlockIncidentV3.md)

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


## apiIncidentsV3SelectorDelete

> SherlockIncidentV3 apiIncidentsV3SelectorDelete(selector)

Delete an individual Incident

Delete an individual Incident by its ID.

### Example

```ts
import {
  Configuration,
  IncidentsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiIncidentsV3SelectorDeleteRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new IncidentsApi();

  const body = {
    // string | The ID of the Incident
    selector: selector_example,
  } satisfies ApiIncidentsV3SelectorDeleteRequest;

  try {
    const data = await api.apiIncidentsV3SelectorDelete(body);
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
| **selector** | `string` | The ID of the Incident | [Defaults to `undefined`] |

### Return type

[**SherlockIncidentV3**](SherlockIncidentV3.md)

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


## apiIncidentsV3SelectorGet

> SherlockIncidentV3 apiIncidentsV3SelectorGet(selector)

Get an individual Incident

Get an individual Incident.

### Example

```ts
import {
  Configuration,
  IncidentsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiIncidentsV3SelectorGetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new IncidentsApi();

  const body = {
    // string | The ID of the Incident
    selector: selector_example,
  } satisfies ApiIncidentsV3SelectorGetRequest;

  try {
    const data = await api.apiIncidentsV3SelectorGet(body);
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
| **selector** | `string` | The ID of the Incident | [Defaults to `undefined`] |

### Return type

[**SherlockIncidentV3**](SherlockIncidentV3.md)

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


## apiIncidentsV3SelectorPatch

> SherlockIncidentV3 apiIncidentsV3SelectorPatch(selector, incident)

Edit an individual Incident

Edit an individual Incident.

### Example

```ts
import {
  Configuration,
  IncidentsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiIncidentsV3SelectorPatchRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new IncidentsApi();

  const body = {
    // string | The ID of the Incident
    selector: selector_example,
    // SherlockIncidentV3Edit | The edits to make to the Incident
    incident: ...,
  } satisfies ApiIncidentsV3SelectorPatchRequest;

  try {
    const data = await api.apiIncidentsV3SelectorPatch(body);
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
| **selector** | `string` | The ID of the Incident | [Defaults to `undefined`] |
| **incident** | [SherlockIncidentV3Edit](SherlockIncidentV3Edit.md) | The edits to make to the Incident | |

### Return type

[**SherlockIncidentV3**](SherlockIncidentV3.md)

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

