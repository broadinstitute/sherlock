# ChartVersionsApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**apiChartVersionsProceduresV3ChangelogGet**](ChartVersionsApi.md#apichartversionsproceduresv3changelogget) | **GET** /api/chart-versions/procedures/v3/changelog | Get a changelog between two ChartVersions |
| [**apiChartVersionsV3Get**](ChartVersionsApi.md#apichartversionsv3get) | **GET** /api/chart-versions/v3 | List ChartVersions matching a filter |
| [**apiChartVersionsV3Put**](ChartVersionsApi.md#apichartversionsv3put) | **PUT** /api/chart-versions/v3 | Upsert a ChartVersion |
| [**apiChartVersionsV3SelectorGet**](ChartVersionsApi.md#apichartversionsv3selectorget) | **GET** /api/chart-versions/v3/{selector} | Get an individual ChartVersion |
| [**apiChartVersionsV3SelectorPatch**](ChartVersionsApi.md#apichartversionsv3selectorpatch) | **PATCH** /api/chart-versions/v3/{selector} | Edit an individual ChartVersion |



## apiChartVersionsProceduresV3ChangelogGet

> SherlockChartVersionV3ChangelogResponse apiChartVersionsProceduresV3ChangelogGet(child, parent)

Get a changelog between two ChartVersions

Get the path through parent references from a child ChartVersion (inclusive) to a parent ChartVersion (exclusive), if possible. Because parent references point from newer children to older parents, the newer ChartVersion should be the child. The result will always exclude the parent.

### Example

```ts
import {
  Configuration,
  ChartVersionsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiChartVersionsProceduresV3ChangelogGetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ChartVersionsApi();

  const body = {
    // string | The selector of the newer ChartVersion for the changelog
    child: child_example,
    // string | The selector of the older ChartVersion for the changelog
    parent: parent_example,
  } satisfies ApiChartVersionsProceduresV3ChangelogGetRequest;

  try {
    const data = await api.apiChartVersionsProceduresV3ChangelogGet(body);
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
| **child** | `string` | The selector of the newer ChartVersion for the changelog | [Defaults to `undefined`] |
| **parent** | `string` | The selector of the older ChartVersion for the changelog | [Defaults to `undefined`] |

### Return type

[**SherlockChartVersionV3ChangelogResponse**](SherlockChartVersionV3ChangelogResponse.md)

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


## apiChartVersionsV3Get

> Array&lt;SherlockChartVersionV3&gt; apiChartVersionsV3Get(authoredBy, chart, chartVersion, createdAt, description, id, parentChartVersion, updatedAt, limit, offset)

List ChartVersions matching a filter

List ChartVersions matching a filter.

### Example

```ts
import {
  Configuration,
  ChartVersionsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiChartVersionsV3GetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ChartVersionsApi();

  const body = {
    // string (optional)
    authoredBy: authoredBy_example,
    // string | Required when creating (optional)
    chart: chart_example,
    // string | Required when creating (optional)
    chartVersion: chartVersion_example,
    // Date (optional)
    createdAt: 2013-10-20T19:20:30+01:00,
    // string | Generally the Git commit message (optional)
    description: description_example,
    // number (optional)
    id: 56,
    // string (optional)
    parentChartVersion: parentChartVersion_example,
    // Date (optional)
    updatedAt: 2013-10-20T19:20:30+01:00,
    // number | Control how many ChartVersions are returned (default 100) (optional)
    limit: 56,
    // number | Control the offset for the returned ChartVersions (default 0) (optional)
    offset: 56,
  } satisfies ApiChartVersionsV3GetRequest;

  try {
    const data = await api.apiChartVersionsV3Get(body);
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
| **authoredBy** | `string` |  | [Optional] [Defaults to `undefined`] |
| **chart** | `string` | Required when creating | [Optional] [Defaults to `undefined`] |
| **chartVersion** | `string` | Required when creating | [Optional] [Defaults to `undefined`] |
| **createdAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **description** | `string` | Generally the Git commit message | [Optional] [Defaults to `undefined`] |
| **id** | `number` |  | [Optional] [Defaults to `undefined`] |
| **parentChartVersion** | `string` |  | [Optional] [Defaults to `undefined`] |
| **updatedAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **limit** | `number` | Control how many ChartVersions are returned (default 100) | [Optional] [Defaults to `undefined`] |
| **offset** | `number` | Control the offset for the returned ChartVersions (default 0) | [Optional] [Defaults to `undefined`] |

### Return type

[**Array&lt;SherlockChartVersionV3&gt;**](SherlockChartVersionV3.md)

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


## apiChartVersionsV3Put

> SherlockChartVersionV3 apiChartVersionsV3Put(chartVersion)

Upsert a ChartVersion

Upsert a ChartVersion.

### Example

```ts
import {
  Configuration,
  ChartVersionsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiChartVersionsV3PutRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ChartVersionsApi();

  const body = {
    // SherlockChartVersionV3Create | The ChartVersion to upsert
    chartVersion: ...,
  } satisfies ApiChartVersionsV3PutRequest;

  try {
    const data = await api.apiChartVersionsV3Put(body);
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
| **chartVersion** | [SherlockChartVersionV3Create](SherlockChartVersionV3Create.md) | The ChartVersion to upsert | |

### Return type

[**SherlockChartVersionV3**](SherlockChartVersionV3.md)

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


## apiChartVersionsV3SelectorGet

> SherlockChartVersionV3 apiChartVersionsV3SelectorGet(selector)

Get an individual ChartVersion

Get an individual ChartVersion.

### Example

```ts
import {
  Configuration,
  ChartVersionsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiChartVersionsV3SelectorGetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ChartVersionsApi();

  const body = {
    // string | The selector of the ChartVersion, which can be either a numeric ID or chart/version.
    selector: selector_example,
  } satisfies ApiChartVersionsV3SelectorGetRequest;

  try {
    const data = await api.apiChartVersionsV3SelectorGet(body);
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
| **selector** | `string` | The selector of the ChartVersion, which can be either a numeric ID or chart/version. | [Defaults to `undefined`] |

### Return type

[**SherlockChartVersionV3**](SherlockChartVersionV3.md)

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


## apiChartVersionsV3SelectorPatch

> SherlockChartVersionV3 apiChartVersionsV3SelectorPatch(selector, chartVersion)

Edit an individual ChartVersion

Edit an individual ChartVersion.

### Example

```ts
import {
  Configuration,
  ChartVersionsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiChartVersionsV3SelectorPatchRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ChartVersionsApi();

  const body = {
    // string | The selector of the ChartVersion, which can be either a numeric ID or chart/version.
    selector: selector_example,
    // SherlockChartVersionV3Edit | The edits to make to the ChartVersion
    chartVersion: ...,
  } satisfies ApiChartVersionsV3SelectorPatchRequest;

  try {
    const data = await api.apiChartVersionsV3SelectorPatch(body);
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
| **selector** | `string` | The selector of the ChartVersion, which can be either a numeric ID or chart/version. | [Defaults to `undefined`] |
| **chartVersion** | [SherlockChartVersionV3Edit](SherlockChartVersionV3Edit.md) | The edits to make to the ChartVersion | |

### Return type

[**SherlockChartVersionV3**](SherlockChartVersionV3.md)

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

