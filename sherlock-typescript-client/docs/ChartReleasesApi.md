# ChartReleasesApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**apiChartReleasesV3Get**](ChartReleasesApi.md#apichartreleasesv3get) | **GET** /api/chart-releases/v3 | List ChartReleases matching a filter |
| [**apiChartReleasesV3Post**](ChartReleasesApi.md#apichartreleasesv3post) | **POST** /api/chart-releases/v3 | Create a ChartRelease |
| [**apiChartReleasesV3SelectorDelete**](ChartReleasesApi.md#apichartreleasesv3selectordelete) | **DELETE** /api/chart-releases/v3/{selector} | Delete an individual ChartRelease |
| [**apiChartReleasesV3SelectorGet**](ChartReleasesApi.md#apichartreleasesv3selectorget) | **GET** /api/chart-releases/v3/{selector} | Get an individual ChartRelease |
| [**apiChartReleasesV3SelectorPatch**](ChartReleasesApi.md#apichartreleasesv3selectorpatch) | **PATCH** /api/chart-releases/v3/{selector} | Edit an individual ChartRelease |



## apiChartReleasesV3Get

> Array&lt;SherlockChartReleaseV3&gt; apiChartReleasesV3Get(appVersionBranch, appVersionCommit, appVersionExact, appVersionFollowChartRelease, appVersionReference, appVersionResolver, chart, chartVersionExact, chartVersionFollowChartRelease, chartVersionReference, chartVersionResolver, cluster, createdAt, destinationType, environment, helmfileRef, helmfileRefEnabled, id, includedInBulkChangesets, name, namespace, pagerdutyIntegration, port, protocol, resolvedAt, subdomain, updatedAt, limit, offset)

List ChartReleases matching a filter

List ChartReleases matching a filter.

### Example

```ts
import {
  Configuration,
  ChartReleasesApi,
} from '@sherlock-js-client/sherlock';
import type { ApiChartReleasesV3GetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ChartReleasesApi();

  const body = {
    // string | When creating, will default to the app\'s mainline branch if no other app version info is present (optional)
    appVersionBranch: appVersionBranch_example,
    // string (optional)
    appVersionCommit: appVersionCommit_example,
    // string (optional)
    appVersionExact: appVersionExact_example,
    // string (optional)
    appVersionFollowChartRelease: appVersionFollowChartRelease_example,
    // string (optional)
    appVersionReference: appVersionReference_example,
    // 'branch' | 'commit' | 'exact' | 'follow' | 'none' | // When creating, will default to automatically reference any provided app version fields (optional)
    appVersionResolver: appVersionResolver_example,
    // string | Required when creating (optional)
    chart: chart_example,
    // string (optional)
    chartVersionExact: chartVersionExact_example,
    // string (optional)
    chartVersionFollowChartRelease: chartVersionFollowChartRelease_example,
    // string (optional)
    chartVersionReference: chartVersionReference_example,
    // 'latest' | 'exact' | 'follow' | When creating, will default to automatically reference any provided chart version (optional)
    chartVersionResolver: chartVersionResolver_example,
    // string | When creating, will default the environment\'s default cluster, if provided. Either this or environment must be provided. (optional)
    cluster: cluster_example,
    // Date (optional)
    createdAt: 2013-10-20T19:20:30+01:00,
    // string | Calculated field (optional)
    destinationType: destinationType_example,
    // string | Either this or cluster must be provided. (optional)
    environment: environment_example,
    // string (optional)
    helmfileRef: helmfileRef_example,
    // boolean (optional)
    helmfileRefEnabled: true,
    // number (optional)
    id: 56,
    // boolean (optional)
    includedInBulkChangesets: true,
    // string | When creating, will be calculated if left empty (optional)
    name: name_example,
    // string | When creating, will default to the environment\'s default namespace, if provided (optional)
    namespace: namespace_example,
    // string (optional)
    pagerdutyIntegration: pagerdutyIntegration_example,
    // number | When creating, will use the chart\'s default if left empty (optional)
    port: 56,
    // string | When creating, will use the chart\'s default if left empty (optional)
    protocol: protocol_example,
    // Date (optional)
    resolvedAt: 2013-10-20T19:20:30+01:00,
    // string | When creating, will use the chart\'s default if left empty (optional)
    subdomain: subdomain_example,
    // Date (optional)
    updatedAt: 2013-10-20T19:20:30+01:00,
    // number | Control how many ChartReleases are returned (default 0, meaning all) (optional)
    limit: 56,
    // number | Control the offset for the returned ChartReleases (default 0) (optional)
    offset: 56,
  } satisfies ApiChartReleasesV3GetRequest;

  try {
    const data = await api.apiChartReleasesV3Get(body);
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
| **appVersionBranch** | `string` | When creating, will default to the app\&#39;s mainline branch if no other app version info is present | [Optional] [Defaults to `undefined`] |
| **appVersionCommit** | `string` |  | [Optional] [Defaults to `undefined`] |
| **appVersionExact** | `string` |  | [Optional] [Defaults to `undefined`] |
| **appVersionFollowChartRelease** | `string` |  | [Optional] [Defaults to `undefined`] |
| **appVersionReference** | `string` |  | [Optional] [Defaults to `undefined`] |
| **appVersionResolver** | `branch`, `commit`, `exact`, `follow`, `none` | // When creating, will default to automatically reference any provided app version fields | [Optional] [Defaults to `undefined`] [Enum: branch, commit, exact, follow, none] |
| **chart** | `string` | Required when creating | [Optional] [Defaults to `undefined`] |
| **chartVersionExact** | `string` |  | [Optional] [Defaults to `undefined`] |
| **chartVersionFollowChartRelease** | `string` |  | [Optional] [Defaults to `undefined`] |
| **chartVersionReference** | `string` |  | [Optional] [Defaults to `undefined`] |
| **chartVersionResolver** | `latest`, `exact`, `follow` | When creating, will default to automatically reference any provided chart version | [Optional] [Defaults to `undefined`] [Enum: latest, exact, follow] |
| **cluster** | `string` | When creating, will default the environment\&#39;s default cluster, if provided. Either this or environment must be provided. | [Optional] [Defaults to `undefined`] |
| **createdAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **destinationType** | `string` | Calculated field | [Optional] [Defaults to `undefined`] |
| **environment** | `string` | Either this or cluster must be provided. | [Optional] [Defaults to `undefined`] |
| **helmfileRef** | `string` |  | [Optional] [Defaults to `&#39;HEAD&#39;`] |
| **helmfileRefEnabled** | `boolean` |  | [Optional] [Defaults to `false`] |
| **id** | `number` |  | [Optional] [Defaults to `undefined`] |
| **includedInBulkChangesets** | `boolean` |  | [Optional] [Defaults to `true`] |
| **name** | `string` | When creating, will be calculated if left empty | [Optional] [Defaults to `undefined`] |
| **namespace** | `string` | When creating, will default to the environment\&#39;s default namespace, if provided | [Optional] [Defaults to `undefined`] |
| **pagerdutyIntegration** | `string` |  | [Optional] [Defaults to `undefined`] |
| **port** | `number` | When creating, will use the chart\&#39;s default if left empty | [Optional] [Defaults to `undefined`] |
| **protocol** | `string` | When creating, will use the chart\&#39;s default if left empty | [Optional] [Defaults to `undefined`] |
| **resolvedAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **subdomain** | `string` | When creating, will use the chart\&#39;s default if left empty | [Optional] [Defaults to `undefined`] |
| **updatedAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **limit** | `number` | Control how many ChartReleases are returned (default 0, meaning all) | [Optional] [Defaults to `undefined`] |
| **offset** | `number` | Control the offset for the returned ChartReleases (default 0) | [Optional] [Defaults to `undefined`] |

### Return type

[**Array&lt;SherlockChartReleaseV3&gt;**](SherlockChartReleaseV3.md)

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


## apiChartReleasesV3Post

> SherlockChartReleaseV3 apiChartReleasesV3Post(chartRelease)

Create a ChartRelease

Create a ChartRelease.

### Example

```ts
import {
  Configuration,
  ChartReleasesApi,
} from '@sherlock-js-client/sherlock';
import type { ApiChartReleasesV3PostRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ChartReleasesApi();

  const body = {
    // SherlockChartReleaseV3Create | The ChartRelease to create
    chartRelease: ...,
  } satisfies ApiChartReleasesV3PostRequest;

  try {
    const data = await api.apiChartReleasesV3Post(body);
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
| **chartRelease** | [SherlockChartReleaseV3Create](SherlockChartReleaseV3Create.md) | The ChartRelease to create | |

### Return type

[**SherlockChartReleaseV3**](SherlockChartReleaseV3.md)

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


## apiChartReleasesV3SelectorDelete

> SherlockChartReleaseV3 apiChartReleasesV3SelectorDelete(selector)

Delete an individual ChartRelease

Delete an individual ChartRelease by its ID.

### Example

```ts
import {
  Configuration,
  ChartReleasesApi,
} from '@sherlock-js-client/sherlock';
import type { ApiChartReleasesV3SelectorDeleteRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ChartReleasesApi();

  const body = {
    // string | The selector of the ChartRelease, which can be either a numeric ID, the name, environment + \'/\' + chart, or cluster + \'/\' + namespace + \'/\' + chart.
    selector: selector_example,
  } satisfies ApiChartReleasesV3SelectorDeleteRequest;

  try {
    const data = await api.apiChartReleasesV3SelectorDelete(body);
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
| **selector** | `string` | The selector of the ChartRelease, which can be either a numeric ID, the name, environment + \&#39;/\&#39; + chart, or cluster + \&#39;/\&#39; + namespace + \&#39;/\&#39; + chart. | [Defaults to `undefined`] |

### Return type

[**SherlockChartReleaseV3**](SherlockChartReleaseV3.md)

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


## apiChartReleasesV3SelectorGet

> SherlockChartReleaseV3 apiChartReleasesV3SelectorGet(selector)

Get an individual ChartRelease

Get an individual ChartRelease.

### Example

```ts
import {
  Configuration,
  ChartReleasesApi,
} from '@sherlock-js-client/sherlock';
import type { ApiChartReleasesV3SelectorGetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ChartReleasesApi();

  const body = {
    // string | The selector of the ChartRelease, which can be either a numeric ID, the name, environment + \'/\' + chart, or cluster + \'/\' + namespace + \'/\' + chart.
    selector: selector_example,
  } satisfies ApiChartReleasesV3SelectorGetRequest;

  try {
    const data = await api.apiChartReleasesV3SelectorGet(body);
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
| **selector** | `string` | The selector of the ChartRelease, which can be either a numeric ID, the name, environment + \&#39;/\&#39; + chart, or cluster + \&#39;/\&#39; + namespace + \&#39;/\&#39; + chart. | [Defaults to `undefined`] |

### Return type

[**SherlockChartReleaseV3**](SherlockChartReleaseV3.md)

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


## apiChartReleasesV3SelectorPatch

> SherlockChartReleaseV3 apiChartReleasesV3SelectorPatch(selector, chartRelease)

Edit an individual ChartRelease

Edit an individual ChartRelease.

### Example

```ts
import {
  Configuration,
  ChartReleasesApi,
} from '@sherlock-js-client/sherlock';
import type { ApiChartReleasesV3SelectorPatchRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ChartReleasesApi();

  const body = {
    // string | The selector of the ChartRelease, which can be either a numeric ID, the name, environment + \'/\' + chart, or cluster + \'/\' + namespace + \'/\' + chart.
    selector: selector_example,
    // SherlockChartReleaseV3Edit | The edits to make to the ChartRelease
    chartRelease: ...,
  } satisfies ApiChartReleasesV3SelectorPatchRequest;

  try {
    const data = await api.apiChartReleasesV3SelectorPatch(body);
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
| **selector** | `string` | The selector of the ChartRelease, which can be either a numeric ID, the name, environment + \&#39;/\&#39; + chart, or cluster + \&#39;/\&#39; + namespace + \&#39;/\&#39; + chart. | [Defaults to `undefined`] |
| **chartRelease** | [SherlockChartReleaseV3Edit](SherlockChartReleaseV3Edit.md) | The edits to make to the ChartRelease | |

### Return type

[**SherlockChartReleaseV3**](SherlockChartReleaseV3.md)

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

