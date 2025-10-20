# ChartsApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**apiChartsV3Get**](ChartsApi.md#apichartsv3get) | **GET** /api/charts/v3 | List Charts matching a filter |
| [**apiChartsV3Post**](ChartsApi.md#apichartsv3post) | **POST** /api/charts/v3 | Create a Chart |
| [**apiChartsV3SelectorDelete**](ChartsApi.md#apichartsv3selectordelete) | **DELETE** /api/charts/v3/{selector} | Delete an individual Chart |
| [**apiChartsV3SelectorGet**](ChartsApi.md#apichartsv3selectorget) | **GET** /api/charts/v3/{selector} | Get an individual Chart |
| [**apiChartsV3SelectorPatch**](ChartsApi.md#apichartsv3selectorpatch) | **PATCH** /api/charts/v3/{selector} | Edit an individual Chart |



## apiChartsV3Get

> Array&lt;SherlockChartV3&gt; apiChartsV3Get(appImageGitMainBranch, appImageGitRepo, chartExposesEndpoint, chartRepo, createdAt, defaultPort, defaultProtocol, defaultSubdomain, description, id, name, pactParticipant, playbookURL, updatedAt, limit, offset)

List Charts matching a filter

List Charts matching a filter.

### Example

```ts
import {
  Configuration,
  ChartsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiChartsV3GetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ChartsApi();

  const body = {
    // string (optional)
    appImageGitMainBranch: appImageGitMainBranch_example,
    // string (optional)
    appImageGitRepo: appImageGitRepo_example,
    // boolean | Indicates if the default subdomain, protocol, and port fields are relevant for this chart (optional)
    chartExposesEndpoint: true,
    // string (optional)
    chartRepo: chartRepo_example,
    // Date (optional)
    createdAt: 2013-10-20T19:20:30+01:00,
    // number (optional)
    defaultPort: 56,
    // string (optional)
    defaultProtocol: defaultProtocol_example,
    // string | When creating, will default to the name of the chart (optional)
    defaultSubdomain: defaultSubdomain_example,
    // string (optional)
    description: description_example,
    // number (optional)
    id: 56,
    // string | Required when creating (optional)
    name: name_example,
    // boolean (optional)
    pactParticipant: true,
    // string (optional)
    playbookURL: playbookURL_example,
    // Date (optional)
    updatedAt: 2013-10-20T19:20:30+01:00,
    // number | Control how many Charts are returned (default 0, meaning all) (optional)
    limit: 56,
    // number | Control the offset for the returned Charts (default 0) (optional)
    offset: 56,
  } satisfies ApiChartsV3GetRequest;

  try {
    const data = await api.apiChartsV3Get(body);
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
| **appImageGitMainBranch** | `string` |  | [Optional] [Defaults to `undefined`] |
| **appImageGitRepo** | `string` |  | [Optional] [Defaults to `undefined`] |
| **chartExposesEndpoint** | `boolean` | Indicates if the default subdomain, protocol, and port fields are relevant for this chart | [Optional] [Defaults to `false`] |
| **chartRepo** | `string` |  | [Optional] [Defaults to `&#39;terra-helm&#39;`] |
| **createdAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **defaultPort** | `number` |  | [Optional] [Defaults to `443`] |
| **defaultProtocol** | `string` |  | [Optional] [Defaults to `&#39;https&#39;`] |
| **defaultSubdomain** | `string` | When creating, will default to the name of the chart | [Optional] [Defaults to `undefined`] |
| **description** | `string` |  | [Optional] [Defaults to `undefined`] |
| **id** | `number` |  | [Optional] [Defaults to `undefined`] |
| **name** | `string` | Required when creating | [Optional] [Defaults to `undefined`] |
| **pactParticipant** | `boolean` |  | [Optional] [Defaults to `false`] |
| **playbookURL** | `string` |  | [Optional] [Defaults to `undefined`] |
| **updatedAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **limit** | `number` | Control how many Charts are returned (default 0, meaning all) | [Optional] [Defaults to `undefined`] |
| **offset** | `number` | Control the offset for the returned Charts (default 0) | [Optional] [Defaults to `undefined`] |

### Return type

[**Array&lt;SherlockChartV3&gt;**](SherlockChartV3.md)

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


## apiChartsV3Post

> SherlockChartV3 apiChartsV3Post(chart)

Create a Chart

Create a Chart.

### Example

```ts
import {
  Configuration,
  ChartsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiChartsV3PostRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ChartsApi();

  const body = {
    // SherlockChartV3Create | The Chart to create
    chart: ...,
  } satisfies ApiChartsV3PostRequest;

  try {
    const data = await api.apiChartsV3Post(body);
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
| **chart** | [SherlockChartV3Create](SherlockChartV3Create.md) | The Chart to create | |

### Return type

[**SherlockChartV3**](SherlockChartV3.md)

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


## apiChartsV3SelectorDelete

> SherlockChartV3 apiChartsV3SelectorDelete(selector)

Delete an individual Chart

Delete an individual Chart by its ID.

### Example

```ts
import {
  Configuration,
  ChartsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiChartsV3SelectorDeleteRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ChartsApi();

  const body = {
    // string | The selector of the Chart, which can be either a numeric ID or the name.
    selector: selector_example,
  } satisfies ApiChartsV3SelectorDeleteRequest;

  try {
    const data = await api.apiChartsV3SelectorDelete(body);
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
| **selector** | `string` | The selector of the Chart, which can be either a numeric ID or the name. | [Defaults to `undefined`] |

### Return type

[**SherlockChartV3**](SherlockChartV3.md)

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


## apiChartsV3SelectorGet

> SherlockChartV3 apiChartsV3SelectorGet(selector)

Get an individual Chart

Get an individual Chart.

### Example

```ts
import {
  Configuration,
  ChartsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiChartsV3SelectorGetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ChartsApi();

  const body = {
    // string | The selector of the Chart, which can be either a numeric ID or the name.
    selector: selector_example,
  } satisfies ApiChartsV3SelectorGetRequest;

  try {
    const data = await api.apiChartsV3SelectorGet(body);
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
| **selector** | `string` | The selector of the Chart, which can be either a numeric ID or the name. | [Defaults to `undefined`] |

### Return type

[**SherlockChartV3**](SherlockChartV3.md)

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


## apiChartsV3SelectorPatch

> SherlockChartV3 apiChartsV3SelectorPatch(selector, chart)

Edit an individual Chart

Edit an individual Chart.

### Example

```ts
import {
  Configuration,
  ChartsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiChartsV3SelectorPatchRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ChartsApi();

  const body = {
    // string | The selector of the Chart, which can be either a numeric ID or the name.
    selector: selector_example,
    // SherlockChartV3Edit | The edits to make to the Chart
    chart: ...,
  } satisfies ApiChartsV3SelectorPatchRequest;

  try {
    const data = await api.apiChartsV3SelectorPatch(body);
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
| **selector** | `string` | The selector of the Chart, which can be either a numeric ID or the name. | [Defaults to `undefined`] |
| **chart** | [SherlockChartV3Edit](SherlockChartV3Edit.md) | The edits to make to the Chart | |

### Return type

[**SherlockChartV3**](SherlockChartV3.md)

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

