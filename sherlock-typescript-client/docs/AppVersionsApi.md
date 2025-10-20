# AppVersionsApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**apiAppVersionsProceduresV3ChangelogGet**](AppVersionsApi.md#apiappversionsproceduresv3changelogget) | **GET** /api/app-versions/procedures/v3/changelog | Get a changelog between two AppVersions |
| [**apiAppVersionsV3Get**](AppVersionsApi.md#apiappversionsv3get) | **GET** /api/app-versions/v3 | List AppVersions matching a filter |
| [**apiAppVersionsV3Put**](AppVersionsApi.md#apiappversionsv3put) | **PUT** /api/app-versions/v3 | Upsert a AppVersion |
| [**apiAppVersionsV3SelectorGet**](AppVersionsApi.md#apiappversionsv3selectorget) | **GET** /api/app-versions/v3/{selector} | Get an individual AppVersion |
| [**apiAppVersionsV3SelectorPatch**](AppVersionsApi.md#apiappversionsv3selectorpatch) | **PATCH** /api/app-versions/v3/{selector} | Edit an individual AppVersion |



## apiAppVersionsProceduresV3ChangelogGet

> SherlockAppVersionV3ChangelogResponse apiAppVersionsProceduresV3ChangelogGet(child, parent)

Get a changelog between two AppVersions

Get the path through parent references from a child AppVersion (inclusive) to a parent AppVersion (exclusive), if possible. Because parent references point from newer children to older parents, the newer AppVersion should be the child. The result will always exclude the parent.

### Example

```ts
import {
  Configuration,
  AppVersionsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiAppVersionsProceduresV3ChangelogGetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new AppVersionsApi();

  const body = {
    // string | The selector of the newer AppVersion for the changelog
    child: child_example,
    // string | The selector of the older AppVersion for the changelog
    parent: parent_example,
  } satisfies ApiAppVersionsProceduresV3ChangelogGetRequest;

  try {
    const data = await api.apiAppVersionsProceduresV3ChangelogGet(body);
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
| **child** | `string` | The selector of the newer AppVersion for the changelog | [Defaults to `undefined`] |
| **parent** | `string` | The selector of the older AppVersion for the changelog | [Defaults to `undefined`] |

### Return type

[**SherlockAppVersionV3ChangelogResponse**](SherlockAppVersionV3ChangelogResponse.md)

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


## apiAppVersionsV3Get

> Array&lt;SherlockAppVersionV3&gt; apiAppVersionsV3Get(appVersion, authoredBy, chart, createdAt, description, gitBranch, gitCommit, id, parentAppVersion, updatedAt, limit, offset)

List AppVersions matching a filter

List AppVersions matching a filter.

### Example

```ts
import {
  Configuration,
  AppVersionsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiAppVersionsV3GetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new AppVersionsApi();

  const body = {
    // string | Required when creating (optional)
    appVersion: appVersion_example,
    // string (optional)
    authoredBy: authoredBy_example,
    // string | Required when creating (optional)
    chart: chart_example,
    // Date (optional)
    createdAt: 2013-10-20T19:20:30+01:00,
    // string | Generally the Git commit message (optional)
    description: description_example,
    // string (optional)
    gitBranch: gitBranch_example,
    // string (optional)
    gitCommit: gitCommit_example,
    // number (optional)
    id: 56,
    // string (optional)
    parentAppVersion: parentAppVersion_example,
    // Date (optional)
    updatedAt: 2013-10-20T19:20:30+01:00,
    // number | Control how many AppVersions are returned (default 100) (optional)
    limit: 56,
    // number | Control the offset for the returned AppVersions (default 0) (optional)
    offset: 56,
  } satisfies ApiAppVersionsV3GetRequest;

  try {
    const data = await api.apiAppVersionsV3Get(body);
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
| **appVersion** | `string` | Required when creating | [Optional] [Defaults to `undefined`] |
| **authoredBy** | `string` |  | [Optional] [Defaults to `undefined`] |
| **chart** | `string` | Required when creating | [Optional] [Defaults to `undefined`] |
| **createdAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **description** | `string` | Generally the Git commit message | [Optional] [Defaults to `undefined`] |
| **gitBranch** | `string` |  | [Optional] [Defaults to `undefined`] |
| **gitCommit** | `string` |  | [Optional] [Defaults to `undefined`] |
| **id** | `number` |  | [Optional] [Defaults to `undefined`] |
| **parentAppVersion** | `string` |  | [Optional] [Defaults to `undefined`] |
| **updatedAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **limit** | `number` | Control how many AppVersions are returned (default 100) | [Optional] [Defaults to `undefined`] |
| **offset** | `number` | Control the offset for the returned AppVersions (default 0) | [Optional] [Defaults to `undefined`] |

### Return type

[**Array&lt;SherlockAppVersionV3&gt;**](SherlockAppVersionV3.md)

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


## apiAppVersionsV3Put

> SherlockAppVersionV3 apiAppVersionsV3Put(appVersion)

Upsert a AppVersion

Upsert a AppVersion.

### Example

```ts
import {
  Configuration,
  AppVersionsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiAppVersionsV3PutRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new AppVersionsApi();

  const body = {
    // SherlockAppVersionV3Create | The AppVersion to upsert
    appVersion: ...,
  } satisfies ApiAppVersionsV3PutRequest;

  try {
    const data = await api.apiAppVersionsV3Put(body);
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
| **appVersion** | [SherlockAppVersionV3Create](SherlockAppVersionV3Create.md) | The AppVersion to upsert | |

### Return type

[**SherlockAppVersionV3**](SherlockAppVersionV3.md)

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


## apiAppVersionsV3SelectorGet

> SherlockAppVersionV3 apiAppVersionsV3SelectorGet(selector)

Get an individual AppVersion

Get an individual AppVersion.

### Example

```ts
import {
  Configuration,
  AppVersionsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiAppVersionsV3SelectorGetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new AppVersionsApi();

  const body = {
    // string | The selector of the AppVersion, which can be either a numeric ID or chart/version.
    selector: selector_example,
  } satisfies ApiAppVersionsV3SelectorGetRequest;

  try {
    const data = await api.apiAppVersionsV3SelectorGet(body);
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
| **selector** | `string` | The selector of the AppVersion, which can be either a numeric ID or chart/version. | [Defaults to `undefined`] |

### Return type

[**SherlockAppVersionV3**](SherlockAppVersionV3.md)

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


## apiAppVersionsV3SelectorPatch

> SherlockAppVersionV3 apiAppVersionsV3SelectorPatch(selector, appVersion)

Edit an individual AppVersion

Edit an individual AppVersion.

### Example

```ts
import {
  Configuration,
  AppVersionsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiAppVersionsV3SelectorPatchRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new AppVersionsApi();

  const body = {
    // string | The selector of the AppVersion, which can be either a numeric ID or chart/version.
    selector: selector_example,
    // SherlockAppVersionV3Edit | The edits to make to the AppVersion
    appVersion: ...,
  } satisfies ApiAppVersionsV3SelectorPatchRequest;

  try {
    const data = await api.apiAppVersionsV3SelectorPatch(body);
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
| **selector** | `string` | The selector of the AppVersion, which can be either a numeric ID or chart/version. | [Defaults to `undefined`] |
| **appVersion** | [SherlockAppVersionV3Edit](SherlockAppVersionV3Edit.md) | The edits to make to the AppVersion | |

### Return type

[**SherlockAppVersionV3**](SherlockAppVersionV3.md)

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

