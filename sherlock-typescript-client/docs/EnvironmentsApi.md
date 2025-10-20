# EnvironmentsApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**apiEnvironmentsV3Get**](EnvironmentsApi.md#apienvironmentsv3get) | **GET** /api/environments/v3 | List Environments matching a filter |
| [**apiEnvironmentsV3Post**](EnvironmentsApi.md#apienvironmentsv3post) | **POST** /api/environments/v3 | Create a Environment |
| [**apiEnvironmentsV3SelectorDelete**](EnvironmentsApi.md#apienvironmentsv3selectordelete) | **DELETE** /api/environments/v3/{selector} | Delete an individual Environment |
| [**apiEnvironmentsV3SelectorGet**](EnvironmentsApi.md#apienvironmentsv3selectorget) | **GET** /api/environments/v3/{selector} | Get an individual Environment |
| [**apiEnvironmentsV3SelectorPatch**](EnvironmentsApi.md#apienvironmentsv3selectorpatch) | **PATCH** /api/environments/v3/{selector} | Edit an individual Environment |



## apiEnvironmentsV3Get

> Array&lt;SherlockEnvironmentV3&gt; apiEnvironmentsV3Get(autoPopulateChartReleases, base, baseDomain, createdAt, defaultCluster, defaultNamespace, deleteAfter, description, enableJanitor, helmfileRef, id, lifecycle, name, namePrefixesDomain, offline, offlineScheduleBeginEnabled, offlineScheduleBeginTime, offlineScheduleEndEnabled, offlineScheduleEndTime, offlineScheduleEndWeekends, owner, pactIdentifier, pagerdutyIntegration, preventDeletion, requiredRole, requiresSuitability, serviceBannerBucket, templateEnvironment, uniqueResourcePrefix, updatedAt, valuesName, limit, offset)

List Environments matching a filter

List Environments matching a filter.

### Example

```ts
import {
  Configuration,
  EnvironmentsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiEnvironmentsV3GetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new EnvironmentsApi();

  const body = {
    // boolean | If true when creating, dynamic environments copy from template and template environments get the honeycomb chart (optional)
    autoPopulateChartReleases: true,
    // string | Required when creating (optional)
    base: base_example,
    // string (optional)
    baseDomain: baseDomain_example,
    // Date (optional)
    createdAt: 2013-10-20T19:20:30+01:00,
    // string (optional)
    defaultCluster: defaultCluster_example,
    // string | When creating, will be calculated if left empty (optional)
    defaultNamespace: defaultNamespace_example,
    // Date | If set, the BEE will be automatically deleted after this time. Can be set to \"\" or Go\'s zero time value to clear the field. (optional)
    deleteAfter: 2013-10-20T19:20:30+01:00,
    // string (optional)
    description: description_example,
    // boolean | If true, janitor resource cleanup will be enabled for this environment. BEEs default to template\'s value, templates default to true, and static/live environments default to false. (optional)
    enableJanitor: true,
    // string (optional)
    helmfileRef: helmfileRef_example,
    // number (optional)
    id: 56,
    // string (optional)
    lifecycle: lifecycle_example,
    // string | When creating, will be calculated if dynamic, required otherwise (optional)
    name: name_example,
    // boolean (optional)
    namePrefixesDomain: true,
    // boolean | Applicable for BEEs only, whether Thelma should render the BEE as \"offline\" zero replicas (this field is a target state, not a status) (optional)
    offline: true,
    // boolean | When enabled, the BEE will be slated to go offline around the begin time each day (optional)
    offlineScheduleBeginEnabled: true,
    // Date | Stored with timezone to determine day of the week (optional)
    offlineScheduleBeginTime: 2013-10-20T19:20:30+01:00,
    // boolean | When enabled, the BEE will be slated to come online around the end time each weekday (each day if weekends enabled) (optional)
    offlineScheduleEndEnabled: true,
    // Date | Stored with timezone to determine day of the week (optional)
    offlineScheduleEndTime: 2013-10-20T19:20:30+01:00,
    // boolean (optional)
    offlineScheduleEndWeekends: true,
    // string | When creating, will default to you (optional)
    owner: owner_example,
    // string (optional)
    pactIdentifier: pactIdentifier_example,
    // string (optional)
    pagerdutyIntegration: pagerdutyIntegration_example,
    // boolean | Used to protect specific BEEs from deletion (thelma checks this field) (optional)
    preventDeletion: true,
    // string | If present, requires membership in the given role for mutations. Set to an empty string to clear. (optional)
    requiredRole: requiredRole_example,
    // boolean (optional)
    requiresSuitability: true,
    // string (optional)
    serviceBannerBucket: serviceBannerBucket_example,
    // string | Required for dynamic environments (optional)
    templateEnvironment: templateEnvironment_example,
    // string | When creating, will be calculated if left empty (optional)
    uniqueResourcePrefix: uniqueResourcePrefix_example,
    // Date (optional)
    updatedAt: 2013-10-20T19:20:30+01:00,
    // string | When creating, defaults to template name or environment name (optional)
    valuesName: valuesName_example,
    // number | Control how many Environments are returned (default 0, meaning all) (optional)
    limit: 56,
    // number | Control the offset for the returned Environments (default 0) (optional)
    offset: 56,
  } satisfies ApiEnvironmentsV3GetRequest;

  try {
    const data = await api.apiEnvironmentsV3Get(body);
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
| **autoPopulateChartReleases** | `boolean` | If true when creating, dynamic environments copy from template and template environments get the honeycomb chart | [Optional] [Defaults to `true`] |
| **base** | `string` | Required when creating | [Optional] [Defaults to `undefined`] |
| **baseDomain** | `string` |  | [Optional] [Defaults to `&#39;bee.envs-terra.bio&#39;`] |
| **createdAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **defaultCluster** | `string` |  | [Optional] [Defaults to `undefined`] |
| **defaultNamespace** | `string` | When creating, will be calculated if left empty | [Optional] [Defaults to `undefined`] |
| **deleteAfter** | `Date` | If set, the BEE will be automatically deleted after this time. Can be set to \&quot;\&quot; or Go\&#39;s zero time value to clear the field. | [Optional] [Defaults to `undefined`] |
| **description** | `string` |  | [Optional] [Defaults to `undefined`] |
| **enableJanitor** | `boolean` | If true, janitor resource cleanup will be enabled for this environment. BEEs default to template\&#39;s value, templates default to true, and static/live environments default to false. | [Optional] [Defaults to `undefined`] |
| **helmfileRef** | `string` |  | [Optional] [Defaults to `&#39;HEAD&#39;`] |
| **id** | `number` |  | [Optional] [Defaults to `undefined`] |
| **lifecycle** | `string` |  | [Optional] [Defaults to `&#39;dynamic&#39;`] |
| **name** | `string` | When creating, will be calculated if dynamic, required otherwise | [Optional] [Defaults to `undefined`] |
| **namePrefixesDomain** | `boolean` |  | [Optional] [Defaults to `true`] |
| **offline** | `boolean` | Applicable for BEEs only, whether Thelma should render the BEE as \&quot;offline\&quot; zero replicas (this field is a target state, not a status) | [Optional] [Defaults to `false`] |
| **offlineScheduleBeginEnabled** | `boolean` | When enabled, the BEE will be slated to go offline around the begin time each day | [Optional] [Defaults to `undefined`] |
| **offlineScheduleBeginTime** | `Date` | Stored with timezone to determine day of the week | [Optional] [Defaults to `undefined`] |
| **offlineScheduleEndEnabled** | `boolean` | When enabled, the BEE will be slated to come online around the end time each weekday (each day if weekends enabled) | [Optional] [Defaults to `undefined`] |
| **offlineScheduleEndTime** | `Date` | Stored with timezone to determine day of the week | [Optional] [Defaults to `undefined`] |
| **offlineScheduleEndWeekends** | `boolean` |  | [Optional] [Defaults to `undefined`] |
| **owner** | `string` | When creating, will default to you | [Optional] [Defaults to `undefined`] |
| **pactIdentifier** | `string` |  | [Optional] [Defaults to `undefined`] |
| **pagerdutyIntegration** | `string` |  | [Optional] [Defaults to `undefined`] |
| **preventDeletion** | `boolean` | Used to protect specific BEEs from deletion (thelma checks this field) | [Optional] [Defaults to `false`] |
| **requiredRole** | `string` | If present, requires membership in the given role for mutations. Set to an empty string to clear. | [Optional] [Defaults to `undefined`] |
| **requiresSuitability** | `boolean` |  | [Optional] [Defaults to `undefined`] |
| **serviceBannerBucket** | `string` |  | [Optional] [Defaults to `undefined`] |
| **templateEnvironment** | `string` | Required for dynamic environments | [Optional] [Defaults to `undefined`] |
| **uniqueResourcePrefix** | `string` | When creating, will be calculated if left empty | [Optional] [Defaults to `undefined`] |
| **updatedAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **valuesName** | `string` | When creating, defaults to template name or environment name | [Optional] [Defaults to `undefined`] |
| **limit** | `number` | Control how many Environments are returned (default 0, meaning all) | [Optional] [Defaults to `undefined`] |
| **offset** | `number` | Control the offset for the returned Environments (default 0) | [Optional] [Defaults to `undefined`] |

### Return type

[**Array&lt;SherlockEnvironmentV3&gt;**](SherlockEnvironmentV3.md)

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


## apiEnvironmentsV3Post

> SherlockEnvironmentV3 apiEnvironmentsV3Post(environment)

Create a Environment

Create a Environment.

### Example

```ts
import {
  Configuration,
  EnvironmentsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiEnvironmentsV3PostRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new EnvironmentsApi();

  const body = {
    // SherlockEnvironmentV3Create | The Environment to create
    environment: ...,
  } satisfies ApiEnvironmentsV3PostRequest;

  try {
    const data = await api.apiEnvironmentsV3Post(body);
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
| **environment** | [SherlockEnvironmentV3Create](SherlockEnvironmentV3Create.md) | The Environment to create | |

### Return type

[**SherlockEnvironmentV3**](SherlockEnvironmentV3.md)

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


## apiEnvironmentsV3SelectorDelete

> SherlockEnvironmentV3 apiEnvironmentsV3SelectorDelete(selector)

Delete an individual Environment

Delete an individual Environment by its ID.

### Example

```ts
import {
  Configuration,
  EnvironmentsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiEnvironmentsV3SelectorDeleteRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new EnvironmentsApi();

  const body = {
    // string | The selector of the Environment, which can be either a numeric ID, the name, or \'resource-prefix\' + / + the unique resource prefix.
    selector: selector_example,
  } satisfies ApiEnvironmentsV3SelectorDeleteRequest;

  try {
    const data = await api.apiEnvironmentsV3SelectorDelete(body);
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
| **selector** | `string` | The selector of the Environment, which can be either a numeric ID, the name, or \&#39;resource-prefix\&#39; + / + the unique resource prefix. | [Defaults to `undefined`] |

### Return type

[**SherlockEnvironmentV3**](SherlockEnvironmentV3.md)

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


## apiEnvironmentsV3SelectorGet

> SherlockEnvironmentV3 apiEnvironmentsV3SelectorGet(selector)

Get an individual Environment

Get an individual Environment.

### Example

```ts
import {
  Configuration,
  EnvironmentsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiEnvironmentsV3SelectorGetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new EnvironmentsApi();

  const body = {
    // string | The selector of the Environment, which can be either a numeric ID, the name, or \'resource-prefix\' + / + the unique resource prefix.
    selector: selector_example,
  } satisfies ApiEnvironmentsV3SelectorGetRequest;

  try {
    const data = await api.apiEnvironmentsV3SelectorGet(body);
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
| **selector** | `string` | The selector of the Environment, which can be either a numeric ID, the name, or \&#39;resource-prefix\&#39; + / + the unique resource prefix. | [Defaults to `undefined`] |

### Return type

[**SherlockEnvironmentV3**](SherlockEnvironmentV3.md)

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


## apiEnvironmentsV3SelectorPatch

> SherlockEnvironmentV3 apiEnvironmentsV3SelectorPatch(selector, environment)

Edit an individual Environment

Edit an individual Environment.

### Example

```ts
import {
  Configuration,
  EnvironmentsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiEnvironmentsV3SelectorPatchRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new EnvironmentsApi();

  const body = {
    // string | The selector of the Environment, which can be either a numeric ID, the name, or \'resource-prefix\' + / + the unique resource prefix.
    selector: selector_example,
    // SherlockEnvironmentV3Edit | The edits to make to the Environment
    environment: ...,
  } satisfies ApiEnvironmentsV3SelectorPatchRequest;

  try {
    const data = await api.apiEnvironmentsV3SelectorPatch(body);
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
| **selector** | `string` | The selector of the Environment, which can be either a numeric ID, the name, or \&#39;resource-prefix\&#39; + / + the unique resource prefix. | [Defaults to `undefined`] |
| **environment** | [SherlockEnvironmentV3Edit](SherlockEnvironmentV3Edit.md) | The edits to make to the Environment | |

### Return type

[**SherlockEnvironmentV3**](SherlockEnvironmentV3.md)

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

