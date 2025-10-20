# PagerdutyIntegrationsApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**apiPagerdutyIntegrationsProceduresV3TriggerIncidentSelectorPost**](PagerdutyIntegrationsApi.md#apipagerdutyintegrationsproceduresv3triggerincidentselectorpost) | **POST** /api/pagerduty-integrations/procedures/v3/trigger-incident/{selector} | Get an individual PagerdutyIntegration |
| [**apiPagerdutyIntegrationsV3Get**](PagerdutyIntegrationsApi.md#apipagerdutyintegrationsv3get) | **GET** /api/pagerduty-integrations/v3 | List PagerdutyIntegrations matching a filter |
| [**apiPagerdutyIntegrationsV3Post**](PagerdutyIntegrationsApi.md#apipagerdutyintegrationsv3post) | **POST** /api/pagerduty-integrations/v3 | Create a PagerdutyIntegration |
| [**apiPagerdutyIntegrationsV3SelectorDelete**](PagerdutyIntegrationsApi.md#apipagerdutyintegrationsv3selectordelete) | **DELETE** /api/pagerduty-integrations/v3/{selector} | Delete an individual PagerdutyIntegration |
| [**apiPagerdutyIntegrationsV3SelectorGet**](PagerdutyIntegrationsApi.md#apipagerdutyintegrationsv3selectorget) | **GET** /api/pagerduty-integrations/v3/{selector} | Get an individual PagerdutyIntegration |
| [**apiPagerdutyIntegrationsV3SelectorPatch**](PagerdutyIntegrationsApi.md#apipagerdutyintegrationsv3selectorpatch) | **PATCH** /api/pagerduty-integrations/v3/{selector} | Edit an individual PagerdutyIntegration |



## apiPagerdutyIntegrationsProceduresV3TriggerIncidentSelectorPost

> PagerdutySendAlertResponse apiPagerdutyIntegrationsProceduresV3TriggerIncidentSelectorPost(selector, summary)

Get an individual PagerdutyIntegration

Get an individual PagerdutyIntegration.

### Example

```ts
import {
  Configuration,
  PagerdutyIntegrationsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiPagerdutyIntegrationsProceduresV3TriggerIncidentSelectorPostRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new PagerdutyIntegrationsApi();

  const body = {
    // string | The selector of the PagerdutyIntegration, which can be either a numeric ID or pd-id/<pagerduty-id>.
    selector: selector_example,
    // PagerdutyAlertSummary | Summary of the incident
    summary: ...,
  } satisfies ApiPagerdutyIntegrationsProceduresV3TriggerIncidentSelectorPostRequest;

  try {
    const data = await api.apiPagerdutyIntegrationsProceduresV3TriggerIncidentSelectorPost(body);
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
| **selector** | `string` | The selector of the PagerdutyIntegration, which can be either a numeric ID or pd-id/&lt;pagerduty-id&gt;. | [Defaults to `undefined`] |
| **summary** | [PagerdutyAlertSummary](PagerdutyAlertSummary.md) | Summary of the incident | |

### Return type

[**PagerdutySendAlertResponse**](PagerdutySendAlertResponse.md)

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


## apiPagerdutyIntegrationsV3Get

> Array&lt;SherlockPagerdutyIntegrationV3&gt; apiPagerdutyIntegrationsV3Get(createdAt, id, name, pagerdutyID, type, updatedAt, limit, offset)

List PagerdutyIntegrations matching a filter

List PagerdutyIntegrations matching a filter.

### Example

```ts
import {
  Configuration,
  PagerdutyIntegrationsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiPagerdutyIntegrationsV3GetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new PagerdutyIntegrationsApi();

  const body = {
    // Date (optional)
    createdAt: 2013-10-20T19:20:30+01:00,
    // number (optional)
    id: 56,
    // string (optional)
    name: name_example,
    // string (optional)
    pagerdutyID: pagerdutyID_example,
    // string (optional)
    type: type_example,
    // Date (optional)
    updatedAt: 2013-10-20T19:20:30+01:00,
    // number | Control how many PagerdutyIntegrations are returned (default 0, meaning all) (optional)
    limit: 56,
    // number | Control the offset for the returned PagerdutyIntegrations (default 0) (optional)
    offset: 56,
  } satisfies ApiPagerdutyIntegrationsV3GetRequest;

  try {
    const data = await api.apiPagerdutyIntegrationsV3Get(body);
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
| **name** | `string` |  | [Optional] [Defaults to `undefined`] |
| **pagerdutyID** | `string` |  | [Optional] [Defaults to `undefined`] |
| **type** | `string` |  | [Optional] [Defaults to `undefined`] |
| **updatedAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **limit** | `number` | Control how many PagerdutyIntegrations are returned (default 0, meaning all) | [Optional] [Defaults to `undefined`] |
| **offset** | `number` | Control the offset for the returned PagerdutyIntegrations (default 0) | [Optional] [Defaults to `undefined`] |

### Return type

[**Array&lt;SherlockPagerdutyIntegrationV3&gt;**](SherlockPagerdutyIntegrationV3.md)

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


## apiPagerdutyIntegrationsV3Post

> SherlockPagerdutyIntegrationV3 apiPagerdutyIntegrationsV3Post(pagerdutyIntegration)

Create a PagerdutyIntegration

Create a PagerdutyIntegration. Duplicate Pagerduty IDs will be gracefully handled by editing the existing entry. This is partially opaque because some fields are writable but not readable.

### Example

```ts
import {
  Configuration,
  PagerdutyIntegrationsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiPagerdutyIntegrationsV3PostRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new PagerdutyIntegrationsApi();

  const body = {
    // SherlockPagerdutyIntegrationV3Create | The PagerdutyIntegration to create
    pagerdutyIntegration: ...,
  } satisfies ApiPagerdutyIntegrationsV3PostRequest;

  try {
    const data = await api.apiPagerdutyIntegrationsV3Post(body);
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
| **pagerdutyIntegration** | [SherlockPagerdutyIntegrationV3Create](SherlockPagerdutyIntegrationV3Create.md) | The PagerdutyIntegration to create | |

### Return type

[**SherlockPagerdutyIntegrationV3**](SherlockPagerdutyIntegrationV3.md)

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


## apiPagerdutyIntegrationsV3SelectorDelete

> SherlockPagerdutyIntegrationV3 apiPagerdutyIntegrationsV3SelectorDelete(selector)

Delete an individual PagerdutyIntegration

Delete an individual PagerdutyIntegration by its ID.

### Example

```ts
import {
  Configuration,
  PagerdutyIntegrationsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiPagerdutyIntegrationsV3SelectorDeleteRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new PagerdutyIntegrationsApi();

  const body = {
    // string | The selector of the PagerdutyIntegration, which can be either a numeric ID or pd-id/<pagerduty-id>.
    selector: selector_example,
  } satisfies ApiPagerdutyIntegrationsV3SelectorDeleteRequest;

  try {
    const data = await api.apiPagerdutyIntegrationsV3SelectorDelete(body);
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
| **selector** | `string` | The selector of the PagerdutyIntegration, which can be either a numeric ID or pd-id/&lt;pagerduty-id&gt;. | [Defaults to `undefined`] |

### Return type

[**SherlockPagerdutyIntegrationV3**](SherlockPagerdutyIntegrationV3.md)

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


## apiPagerdutyIntegrationsV3SelectorGet

> SherlockPagerdutyIntegrationV3 apiPagerdutyIntegrationsV3SelectorGet(selector)

Get an individual PagerdutyIntegration

Get an individual PagerdutyIntegration.

### Example

```ts
import {
  Configuration,
  PagerdutyIntegrationsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiPagerdutyIntegrationsV3SelectorGetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new PagerdutyIntegrationsApi();

  const body = {
    // string | The selector of the PagerdutyIntegration, which can be either a numeric ID or pd-id/<pagerduty-id>.
    selector: selector_example,
  } satisfies ApiPagerdutyIntegrationsV3SelectorGetRequest;

  try {
    const data = await api.apiPagerdutyIntegrationsV3SelectorGet(body);
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
| **selector** | `string` | The selector of the PagerdutyIntegration, which can be either a numeric ID or pd-id/&lt;pagerduty-id&gt;. | [Defaults to `undefined`] |

### Return type

[**SherlockPagerdutyIntegrationV3**](SherlockPagerdutyIntegrationV3.md)

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


## apiPagerdutyIntegrationsV3SelectorPatch

> SherlockPagerdutyIntegrationV3 apiPagerdutyIntegrationsV3SelectorPatch(selector, pagerdutyIntegration)

Edit an individual PagerdutyIntegration

Edit an individual PagerdutyIntegration.

### Example

```ts
import {
  Configuration,
  PagerdutyIntegrationsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiPagerdutyIntegrationsV3SelectorPatchRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new PagerdutyIntegrationsApi();

  const body = {
    // string | The selector of the PagerdutyIntegration, which can be either a numeric ID or pd-id/<pagerduty-id>.
    selector: selector_example,
    // SherlockPagerdutyIntegrationV3Edit | The edits to make to the PagerdutyIntegration
    pagerdutyIntegration: ...,
  } satisfies ApiPagerdutyIntegrationsV3SelectorPatchRequest;

  try {
    const data = await api.apiPagerdutyIntegrationsV3SelectorPatch(body);
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
| **selector** | `string` | The selector of the PagerdutyIntegration, which can be either a numeric ID or pd-id/&lt;pagerduty-id&gt;. | [Defaults to `undefined`] |
| **pagerdutyIntegration** | [SherlockPagerdutyIntegrationV3Edit](SherlockPagerdutyIntegrationV3Edit.md) | The edits to make to the PagerdutyIntegration | |

### Return type

[**SherlockPagerdutyIntegrationV3**](SherlockPagerdutyIntegrationV3.md)

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

