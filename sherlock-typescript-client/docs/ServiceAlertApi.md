# ServiceAlertApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**apiServiceAlertsProceduresV3SyncPost**](ServiceAlertApi.md#apiservicealertsproceduresv3syncpost) | **POST** /api/service-alerts/procedures/v3/sync | Sync service alerts |
| [**apiServiceAlertsV3Get**](ServiceAlertApi.md#apiservicealertsv3get) | **GET** /api/service-alerts/v3 | List ServiceAlerts matching a filter |
| [**apiServiceAlertsV3Post**](ServiceAlertApi.md#apiservicealertsv3post) | **POST** /api/service-alerts/v3 | Create a service alert |
| [**apiServiceAlertsV3SelectorDelete**](ServiceAlertApi.md#apiservicealertsv3selectordelete) | **DELETE** /api/service-alerts/v3/{selector} | Delete a ServiceAlert |
| [**apiServiceAlertsV3SelectorGet**](ServiceAlertApi.md#apiservicealertsv3selectorget) | **GET** /api/service-alerts/v3/{selector} | Get a Service Alert |
| [**apiServiceAlertsV3SelectorPatch**](ServiceAlertApi.md#apiservicealertsv3selectorpatch) | **PATCH** /api/service-alerts/v3/{selector} | Edit a service alert |



## apiServiceAlertsProceduresV3SyncPost

> Array&lt;SherlockServiceAlertV3&gt; apiServiceAlertsProceduresV3SyncPost(environment)

Sync service alerts

Method to get all currently active service alerts from Sherlock\&#39;s DB and ensure that the service alert json files placed in Google Buckets for Terra match.

### Example

```ts
import {
  Configuration,
  ServiceAlertApi,
} from '@sherlock-js-client/sherlock';
import type { ApiServiceAlertsProceduresV3SyncPostRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ServiceAlertApi();

  const body = {
    // SherlockServiceAlertV3SyncRequest | Information on Service Alert environment
    environment: ...,
  } satisfies ApiServiceAlertsProceduresV3SyncPostRequest;

  try {
    const data = await api.apiServiceAlertsProceduresV3SyncPost(body);
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
| **environment** | [SherlockServiceAlertV3SyncRequest](SherlockServiceAlertV3SyncRequest.md) | Information on Service Alert environment | |

### Return type

[**Array&lt;SherlockServiceAlertV3&gt;**](SherlockServiceAlertV3.md)

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


## apiServiceAlertsV3Get

> Array&lt;SherlockServiceAlertV3&gt; apiServiceAlertsV3Get(createdAt, createdBy, deltedBy, id, link, message, onEnvironment, severity, title, updatedAt, updatedBy, uuid, limit, offset, includeDeleted)

List ServiceAlerts matching a filter

List ServiceAlerts matching a filter.

### Example

```ts
import {
  Configuration,
  ServiceAlertApi,
} from '@sherlock-js-client/sherlock';
import type { ApiServiceAlertsV3GetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ServiceAlertApi();

  const body = {
    // Date (optional)
    createdAt: 2013-10-20T19:20:30+01:00,
    // string (optional)
    createdBy: createdBy_example,
    // string (optional)
    deltedBy: deltedBy_example,
    // number (optional)
    id: 56,
    // string (optional)
    link: link_example,
    // string (optional)
    message: message_example,
    // string (optional)
    onEnvironment: onEnvironment_example,
    // 'blocker' | ' critical' | ' minor' (optional)
    severity: severity_example,
    // string (optional)
    title: title_example,
    // Date (optional)
    updatedAt: 2013-10-20T19:20:30+01:00,
    // string (optional)
    updatedBy: updatedBy_example,
    // string (optional)
    uuid: uuid_example,
    // number | Control how many Service Alerts are returned (default 0, no limit) (optional)
    limit: 56,
    // number | Control the offset for the returned Service Alerts (default 0) (optional)
    offset: 56,
    // boolean | Control if only active Service Alerts are returned, set to true to return deleted Alerts (default false) (optional)
    includeDeleted: true,
  } satisfies ApiServiceAlertsV3GetRequest;

  try {
    const data = await api.apiServiceAlertsV3Get(body);
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
| **createdBy** | `string` |  | [Optional] [Defaults to `undefined`] |
| **deltedBy** | `string` |  | [Optional] [Defaults to `undefined`] |
| **id** | `number` |  | [Optional] [Defaults to `undefined`] |
| **link** | `string` |  | [Optional] [Defaults to `undefined`] |
| **message** | `string` |  | [Optional] [Defaults to `undefined`] |
| **onEnvironment** | `string` |  | [Optional] [Defaults to `undefined`] |
| **severity** | `blocker`, ` critical`, ` minor` |  | [Optional] [Defaults to `undefined`] [Enum: blocker,  critical,  minor] |
| **title** | `string` |  | [Optional] [Defaults to `undefined`] |
| **updatedAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **updatedBy** | `string` |  | [Optional] [Defaults to `undefined`] |
| **uuid** | `string` |  | [Optional] [Defaults to `undefined`] |
| **limit** | `number` | Control how many Service Alerts are returned (default 0, no limit) | [Optional] [Defaults to `undefined`] |
| **offset** | `number` | Control the offset for the returned Service Alerts (default 0) | [Optional] [Defaults to `undefined`] |
| **includeDeleted** | `boolean` | Control if only active Service Alerts are returned, set to true to return deleted Alerts (default false) | [Optional] [Defaults to `undefined`] |

### Return type

[**Array&lt;SherlockServiceAlertV3&gt;**](SherlockServiceAlertV3.md)

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


## apiServiceAlertsV3Post

> SherlockServiceAlertV3 apiServiceAlertsV3Post(serviceAlert)

Create a service alert

Create a service alert to be displayed within terra.

### Example

```ts
import {
  Configuration,
  ServiceAlertApi,
} from '@sherlock-js-client/sherlock';
import type { ApiServiceAlertsV3PostRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ServiceAlertApi();

  const body = {
    // SherlockServiceAlertV3Create | The initial fields the ServiceAlert should have set
    serviceAlert: ...,
  } satisfies ApiServiceAlertsV3PostRequest;

  try {
    const data = await api.apiServiceAlertsV3Post(body);
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
| **serviceAlert** | [SherlockServiceAlertV3Create](SherlockServiceAlertV3Create.md) | The initial fields the ServiceAlert should have set | |

### Return type

[**SherlockServiceAlertV3**](SherlockServiceAlertV3.md)

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


## apiServiceAlertsV3SelectorDelete

> SherlockServiceAlertV3 apiServiceAlertsV3SelectorDelete(selector)

Delete a ServiceAlert

Delete an individual ServiceAlert.

### Example

```ts
import {
  Configuration,
  ServiceAlertApi,
} from '@sherlock-js-client/sherlock';
import type { ApiServiceAlertsV3SelectorDeleteRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ServiceAlertApi();

  const body = {
    // string | The selector of the ServiceAlert, ServiceAlert, which is the guid for a given alert
    selector: selector_example,
  } satisfies ApiServiceAlertsV3SelectorDeleteRequest;

  try {
    const data = await api.apiServiceAlertsV3SelectorDelete(body);
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
| **selector** | `string` | The selector of the ServiceAlert, ServiceAlert, which is the guid for a given alert | [Defaults to `undefined`] |

### Return type

[**SherlockServiceAlertV3**](SherlockServiceAlertV3.md)

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


## apiServiceAlertsV3SelectorGet

> SherlockServiceAlertV3 apiServiceAlertsV3SelectorGet(selector, includeDeleted)

Get a Service Alert

Get an individual Service Alert and it\&#39;s metadata.

### Example

```ts
import {
  Configuration,
  ServiceAlertApi,
} from '@sherlock-js-client/sherlock';
import type { ApiServiceAlertsV3SelectorGetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ServiceAlertApi();

  const body = {
    // string | The selector of the ServiceAlert, which is the guid for a given alert
    selector: selector_example,
    // boolean | Control if only active Service Alerts are returned, set to true to return deleted Alerts (default false) (optional)
    includeDeleted: true,
  } satisfies ApiServiceAlertsV3SelectorGetRequest;

  try {
    const data = await api.apiServiceAlertsV3SelectorGet(body);
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
| **selector** | `string` | The selector of the ServiceAlert, which is the guid for a given alert | [Defaults to `undefined`] |
| **includeDeleted** | `boolean` | Control if only active Service Alerts are returned, set to true to return deleted Alerts (default false) | [Optional] [Defaults to `undefined`] |

### Return type

[**SherlockServiceAlertV3**](SherlockServiceAlertV3.md)

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


## apiServiceAlertsV3SelectorPatch

> SherlockServiceAlertV3 apiServiceAlertsV3SelectorPatch(selector, serviceAlert)

Edit a service alert

Update a service alert with new information.

### Example

```ts
import {
  Configuration,
  ServiceAlertApi,
} from '@sherlock-js-client/sherlock';
import type { ApiServiceAlertsV3SelectorPatchRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ServiceAlertApi();

  const body = {
    // string | The selector of the ServiceAlert, which is the guid for a given alert
    selector: selector_example,
    // SherlockServiceAlertV3EditableFields | The edits to make to the ServiceAlert
    serviceAlert: ...,
  } satisfies ApiServiceAlertsV3SelectorPatchRequest;

  try {
    const data = await api.apiServiceAlertsV3SelectorPatch(body);
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
| **selector** | `string` | The selector of the ServiceAlert, which is the guid for a given alert | [Defaults to `undefined`] |
| **serviceAlert** | [SherlockServiceAlertV3EditableFields](SherlockServiceAlertV3EditableFields.md) | The edits to make to the ServiceAlert | |

### Return type

[**SherlockServiceAlertV3**](SherlockServiceAlertV3.md)

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

