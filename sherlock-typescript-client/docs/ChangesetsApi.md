# ChangesetsApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**apiChangesetsProceduresV3ApplyPost**](ChangesetsApi.md#apichangesetsproceduresv3applypost) | **POST** /api/changesets/procedures/v3/apply | Apply previously planned version changes to Chart Releases |
| [**apiChangesetsProceduresV3ChartReleaseHistoryChartReleaseGet**](ChangesetsApi.md#apichangesetsproceduresv3chartreleasehistorychartreleaseget) | **GET** /api/changesets/procedures/v3/chart-release-history/{chart-release} | List applied Changesets for a Chart Release |
| [**apiChangesetsProceduresV3PlanAndApplyPost**](ChangesetsApi.md#apichangesetsproceduresv3planandapplypost) | **POST** /api/changesets/procedures/v3/plan-and-apply | Plan and apply version changes in one step |
| [**apiChangesetsProceduresV3PlanPost**](ChangesetsApi.md#apichangesetsproceduresv3planpost) | **POST** /api/changesets/procedures/v3/plan | Plan--but do not apply--version changes to Chart Releases |
| [**apiChangesetsProceduresV3VersionHistoryVersionTypeChartVersionGet**](ChangesetsApi.md#apichangesetsproceduresv3versionhistoryversiontypechartversionget) | **GET** /api/changesets/procedures/v3/version-history/{version-type}/{chart}/{version} | List applied Changesets for an App or Chart Version |
| [**apiChangesetsV3Get**](ChangesetsApi.md#apichangesetsv3get) | **GET** /api/changesets/v3 | List Changesets matching a filter |
| [**apiChangesetsV3IdGet**](ChangesetsApi.md#apichangesetsv3idget) | **GET** /api/changesets/v3/{id} | Get an individual Changeset |



## apiChangesetsProceduresV3ApplyPost

> Array&lt;SherlockChangesetV3&gt; apiChangesetsProceduresV3ApplyPost(applyRequest, verboseOutput)

Apply previously planned version changes to Chart Releases

Looks up and applies previously-planned version diffs given by the ID. Other stored plans against the same Chart Releases are marked as superseded. Multiple Changesets can be specified simply by passing multiple IDs in the list.

### Example

```ts
import {
  Configuration,
  ChangesetsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiChangesetsProceduresV3ApplyPostRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ChangesetsApi();

  const body = {
    // Array<string> | String IDs of the Changesets to apply
    applyRequest: ...,
    // boolean | If full information about the changesets should be returned, defaults to true. If false, only the IDs will be returned. (optional)
    verboseOutput: true,
  } satisfies ApiChangesetsProceduresV3ApplyPostRequest;

  try {
    const data = await api.apiChangesetsProceduresV3ApplyPost(body);
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
| **applyRequest** | `Array<string>` | String IDs of the Changesets to apply | |
| **verboseOutput** | `boolean` | If full information about the changesets should be returned, defaults to true. If false, only the IDs will be returned. | [Optional] [Defaults to `undefined`] |

### Return type

[**Array&lt;SherlockChangesetV3&gt;**](SherlockChangesetV3.md)

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


## apiChangesetsProceduresV3ChartReleaseHistoryChartReleaseGet

> Array&lt;SherlockChangesetV3&gt; apiChangesetsProceduresV3ChartReleaseHistoryChartReleaseGet(chartRelease, offset, limit)

List applied Changesets for a Chart Release

List existing applied Changesets for a particular Chart Release, ordered by most recently applied.

### Example

```ts
import {
  Configuration,
  ChangesetsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiChangesetsProceduresV3ChartReleaseHistoryChartReleaseGetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ChangesetsApi();

  const body = {
    // string | Selector of the Chart Release to find applied Changesets for
    chartRelease: chartRelease_example,
    // number | An optional offset to skip a number of latest Changesets (optional)
    offset: 56,
    // number | An optional limit to the number of entries returned (optional)
    limit: 56,
  } satisfies ApiChangesetsProceduresV3ChartReleaseHistoryChartReleaseGetRequest;

  try {
    const data = await api.apiChangesetsProceduresV3ChartReleaseHistoryChartReleaseGet(body);
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
| **chartRelease** | `string` | Selector of the Chart Release to find applied Changesets for | [Defaults to `undefined`] |
| **offset** | `number` | An optional offset to skip a number of latest Changesets | [Optional] [Defaults to `undefined`] |
| **limit** | `number` | An optional limit to the number of entries returned | [Optional] [Defaults to `undefined`] |

### Return type

[**Array&lt;SherlockChangesetV3&gt;**](SherlockChangesetV3.md)

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


## apiChangesetsProceduresV3PlanAndApplyPost

> Array&lt;SherlockChangesetV3&gt; apiChangesetsProceduresV3PlanAndApplyPost(changesetPlanRequest, verboseOutput)

Plan and apply version changes in one step

Like calling the plan procedure immediately followed by the apply procedure. See those endpoints for more information.

### Example

```ts
import {
  Configuration,
  ChangesetsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiChangesetsProceduresV3PlanAndApplyPostRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ChangesetsApi();

  const body = {
    // SherlockChangesetV3PlanRequest | Info on what version changes or refreshes to apply.
    changesetPlanRequest: ...,
    // boolean | If full information about the changesets should be returned, defaults to true. If false, only the IDs will be returned. (optional)
    verboseOutput: true,
  } satisfies ApiChangesetsProceduresV3PlanAndApplyPostRequest;

  try {
    const data = await api.apiChangesetsProceduresV3PlanAndApplyPost(body);
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
| **changesetPlanRequest** | [SherlockChangesetV3PlanRequest](SherlockChangesetV3PlanRequest.md) | Info on what version changes or refreshes to apply. | |
| **verboseOutput** | `boolean` | If full information about the changesets should be returned, defaults to true. If false, only the IDs will be returned. | [Optional] [Defaults to `undefined`] |

### Return type

[**Array&lt;SherlockChangesetV3&gt;**](SherlockChangesetV3.md)

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


## apiChangesetsProceduresV3PlanPost

> Array&lt;SherlockChangesetV3&gt; apiChangesetsProceduresV3PlanPost(changesetPlanRequest, verboseOutput)

Plan--but do not apply--version changes to Chart Releases

Refreshes and calculates version diffs for Chart Releases. If there\&#39;s a diff, the plan is stored and returned so it can be applied later.

### Example

```ts
import {
  Configuration,
  ChangesetsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiChangesetsProceduresV3PlanPostRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ChangesetsApi();

  const body = {
    // SherlockChangesetV3PlanRequest | Info on what version changes or refreshes to plan
    changesetPlanRequest: ...,
    // boolean | If full information about the changesets should be returned, defaults to true. If false, only the IDs will be returned. (optional)
    verboseOutput: true,
  } satisfies ApiChangesetsProceduresV3PlanPostRequest;

  try {
    const data = await api.apiChangesetsProceduresV3PlanPost(body);
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
| **changesetPlanRequest** | [SherlockChangesetV3PlanRequest](SherlockChangesetV3PlanRequest.md) | Info on what version changes or refreshes to plan | |
| **verboseOutput** | `boolean` | If full information about the changesets should be returned, defaults to true. If false, only the IDs will be returned. | [Optional] [Defaults to `undefined`] |

### Return type

[**Array&lt;SherlockChangesetV3&gt;**](SherlockChangesetV3.md)

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


## apiChangesetsProceduresV3VersionHistoryVersionTypeChartVersionGet

> Array&lt;SherlockChangesetV3&gt; apiChangesetsProceduresV3VersionHistoryVersionTypeChartVersionGet(versionType, chart, version)

List applied Changesets for an App or Chart Version

List existing applied Changesets that newly deployed a given App Version or Chart Version, ordered by most recently applied.

### Example

```ts
import {
  Configuration,
  ChangesetsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiChangesetsProceduresV3VersionHistoryVersionTypeChartVersionGetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ChangesetsApi();

  const body = {
    // 'app' | 'chart' | The type of the version, either \'app\' or \'chart\'
    versionType: versionType_example,
    // string | The chart the version belongs to
    chart: chart_example,
    // string | The version to look for
    version: version_example,
  } satisfies ApiChangesetsProceduresV3VersionHistoryVersionTypeChartVersionGetRequest;

  try {
    const data = await api.apiChangesetsProceduresV3VersionHistoryVersionTypeChartVersionGet(body);
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
| **versionType** | `app`, `chart` | The type of the version, either \&#39;app\&#39; or \&#39;chart\&#39; | [Defaults to `undefined`] [Enum: app, chart] |
| **chart** | `string` | The chart the version belongs to | [Defaults to `undefined`] |
| **version** | `string` | The version to look for | [Defaults to `undefined`] |

### Return type

[**Array&lt;SherlockChangesetV3&gt;**](SherlockChangesetV3.md)

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


## apiChangesetsV3Get

> Array&lt;SherlockChangesetV3&gt; apiChangesetsV3Get(appliedAt, appliedBy, chartRelease, fromAppVersionBranch, fromAppVersionCommit, fromAppVersionExact, fromAppVersionFollowChartRelease, fromAppVersionReference, fromAppVersionResolver, fromChartVersionExact, fromChartVersionFollowChartRelease, fromChartVersionReference, fromChartVersionResolver, fromHelmfileRef, fromHelmfileRefEnabled, fromResolvedAt, plannedBy, supersededAt, toAppVersionBranch, toAppVersionCommit, toAppVersionExact, toAppVersionFollowChartRelease, toAppVersionReference, toAppVersionResolver, toChartVersionExact, toChartVersionFollowChartRelease, toChartVersionReference, toChartVersionResolver, toHelmfileRef, toHelmfileRefEnabled, toResolvedAt, id, limit, offset)

List Changesets matching a filter

List Changesets matching a filter.

### Example

```ts
import {
  Configuration,
  ChangesetsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiChangesetsV3GetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ChangesetsApi();

  const body = {
    // Date (optional)
    appliedAt: 2013-10-20T19:20:30+01:00,
    // string (optional)
    appliedBy: appliedBy_example,
    // string (optional)
    chartRelease: chartRelease_example,
    // string (optional)
    fromAppVersionBranch: fromAppVersionBranch_example,
    // string (optional)
    fromAppVersionCommit: fromAppVersionCommit_example,
    // string (optional)
    fromAppVersionExact: fromAppVersionExact_example,
    // string (optional)
    fromAppVersionFollowChartRelease: fromAppVersionFollowChartRelease_example,
    // string (optional)
    fromAppVersionReference: fromAppVersionReference_example,
    // string (optional)
    fromAppVersionResolver: fromAppVersionResolver_example,
    // string (optional)
    fromChartVersionExact: fromChartVersionExact_example,
    // string (optional)
    fromChartVersionFollowChartRelease: fromChartVersionFollowChartRelease_example,
    // string (optional)
    fromChartVersionReference: fromChartVersionReference_example,
    // string (optional)
    fromChartVersionResolver: fromChartVersionResolver_example,
    // string (optional)
    fromHelmfileRef: fromHelmfileRef_example,
    // boolean (optional)
    fromHelmfileRefEnabled: true,
    // Date (optional)
    fromResolvedAt: 2013-10-20T19:20:30+01:00,
    // string (optional)
    plannedBy: plannedBy_example,
    // Date (optional)
    supersededAt: 2013-10-20T19:20:30+01:00,
    // string (optional)
    toAppVersionBranch: toAppVersionBranch_example,
    // string (optional)
    toAppVersionCommit: toAppVersionCommit_example,
    // string (optional)
    toAppVersionExact: toAppVersionExact_example,
    // string (optional)
    toAppVersionFollowChartRelease: toAppVersionFollowChartRelease_example,
    // string (optional)
    toAppVersionReference: toAppVersionReference_example,
    // string (optional)
    toAppVersionResolver: toAppVersionResolver_example,
    // string (optional)
    toChartVersionExact: toChartVersionExact_example,
    // string (optional)
    toChartVersionFollowChartRelease: toChartVersionFollowChartRelease_example,
    // string (optional)
    toChartVersionReference: toChartVersionReference_example,
    // string (optional)
    toChartVersionResolver: toChartVersionResolver_example,
    // string (optional)
    toHelmfileRef: toHelmfileRef_example,
    // boolean (optional)
    toHelmfileRefEnabled: true,
    // Date (optional)
    toResolvedAt: 2013-10-20T19:20:30+01:00,
    // Array<number> | Get specific changesets by their IDs, can be passed multiple times and/or be comma-separated (optional)
    id: ...,
    // number | Control how many Changesets are returned (default 100), ignored if specific IDs are passed (optional)
    limit: 56,
    // number | Control the offset for the returned Changesets (default 0), ignored if specific IDs are passed (optional)
    offset: 56,
  } satisfies ApiChangesetsV3GetRequest;

  try {
    const data = await api.apiChangesetsV3Get(body);
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
| **appliedAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **appliedBy** | `string` |  | [Optional] [Defaults to `undefined`] |
| **chartRelease** | `string` |  | [Optional] [Defaults to `undefined`] |
| **fromAppVersionBranch** | `string` |  | [Optional] [Defaults to `undefined`] |
| **fromAppVersionCommit** | `string` |  | [Optional] [Defaults to `undefined`] |
| **fromAppVersionExact** | `string` |  | [Optional] [Defaults to `undefined`] |
| **fromAppVersionFollowChartRelease** | `string` |  | [Optional] [Defaults to `undefined`] |
| **fromAppVersionReference** | `string` |  | [Optional] [Defaults to `undefined`] |
| **fromAppVersionResolver** | `string` |  | [Optional] [Defaults to `undefined`] |
| **fromChartVersionExact** | `string` |  | [Optional] [Defaults to `undefined`] |
| **fromChartVersionFollowChartRelease** | `string` |  | [Optional] [Defaults to `undefined`] |
| **fromChartVersionReference** | `string` |  | [Optional] [Defaults to `undefined`] |
| **fromChartVersionResolver** | `string` |  | [Optional] [Defaults to `undefined`] |
| **fromHelmfileRef** | `string` |  | [Optional] [Defaults to `undefined`] |
| **fromHelmfileRefEnabled** | `boolean` |  | [Optional] [Defaults to `undefined`] |
| **fromResolvedAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **plannedBy** | `string` |  | [Optional] [Defaults to `undefined`] |
| **supersededAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **toAppVersionBranch** | `string` |  | [Optional] [Defaults to `undefined`] |
| **toAppVersionCommit** | `string` |  | [Optional] [Defaults to `undefined`] |
| **toAppVersionExact** | `string` |  | [Optional] [Defaults to `undefined`] |
| **toAppVersionFollowChartRelease** | `string` |  | [Optional] [Defaults to `undefined`] |
| **toAppVersionReference** | `string` |  | [Optional] [Defaults to `undefined`] |
| **toAppVersionResolver** | `string` |  | [Optional] [Defaults to `undefined`] |
| **toChartVersionExact** | `string` |  | [Optional] [Defaults to `undefined`] |
| **toChartVersionFollowChartRelease** | `string` |  | [Optional] [Defaults to `undefined`] |
| **toChartVersionReference** | `string` |  | [Optional] [Defaults to `undefined`] |
| **toChartVersionResolver** | `string` |  | [Optional] [Defaults to `undefined`] |
| **toHelmfileRef** | `string` |  | [Optional] [Defaults to `undefined`] |
| **toHelmfileRefEnabled** | `boolean` |  | [Optional] [Defaults to `undefined`] |
| **toResolvedAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **id** | `Array<number>` | Get specific changesets by their IDs, can be passed multiple times and/or be comma-separated | [Optional] |
| **limit** | `number` | Control how many Changesets are returned (default 100), ignored if specific IDs are passed | [Optional] [Defaults to `undefined`] |
| **offset** | `number` | Control the offset for the returned Changesets (default 0), ignored if specific IDs are passed | [Optional] [Defaults to `undefined`] |

### Return type

[**Array&lt;SherlockChangesetV3&gt;**](SherlockChangesetV3.md)

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


## apiChangesetsV3IdGet

> SherlockChangesetV3 apiChangesetsV3IdGet(id)

Get an individual Changeset

Get an individual Changeset.

### Example

```ts
import {
  Configuration,
  ChangesetsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiChangesetsV3IdGetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new ChangesetsApi();

  const body = {
    // number | The numeric ID of the changeset
    id: 56,
  } satisfies ApiChangesetsV3IdGetRequest;

  try {
    const data = await api.apiChangesetsV3IdGet(body);
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
| **id** | `number` | The numeric ID of the changeset | [Defaults to `undefined`] |

### Return type

[**SherlockChangesetV3**](SherlockChangesetV3.md)

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

