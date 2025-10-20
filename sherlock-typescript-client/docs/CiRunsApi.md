# CiRunsApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**apiCiRunsProceduresV3GithubInfoGet**](CiRunsApi.md#apicirunsproceduresv3githubinfoget) | **GET** /api/ci-runs/procedures/v3/github-info | List GitHub info gleaned from CiRuns |
| [**apiCiRunsV3Get**](CiRunsApi.md#apicirunsv3get) | **GET** /api/ci-runs/v3 | List CiRuns matching a filter |
| [**apiCiRunsV3Put**](CiRunsApi.md#apicirunsv3put) | **PUT** /api/ci-runs/v3 | Create or update a CiRun |
| [**apiCiRunsV3SelectorGet**](CiRunsApi.md#apicirunsv3selectorget) | **GET** /api/ci-runs/v3/{selector} | Get a CiRun, including CiIdentifiers for related resources |



## apiCiRunsProceduresV3GithubInfoGet

> { [key: string]: { [key: string]: Array&lt;string&gt;; }; } apiCiRunsProceduresV3GithubInfoGet()

List GitHub info gleaned from CiRuns

List info about GitHub repos and their workflow files as determined by CiRuns from the past 90 days. This is a useful proxy for figuring out what repos Sherlock probably has access to: workflows listed here can probably successfully called by a GitHub Actions deploy hook.

### Example

```ts
import {
  Configuration,
  CiRunsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiCiRunsProceduresV3GithubInfoGetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new CiRunsApi();

  try {
    const data = await api.apiCiRunsProceduresV3GithubInfoGet();
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters

This endpoint does not need any parameter.

### Return type

**{ [key: string]: { [key: string]: Array<string>; }; }**

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


## apiCiRunsV3Get

> Array&lt;SherlockCiRunV3&gt; apiCiRunsV3Get(argoWorkflowsName, argoWorkflowsNamespace, argoWorkflowsTemplate, createdAt, githubActionsAttemptNumber, githubActionsOwner, githubActionsRepo, githubActionsRunID, githubActionsWorkflowPath, id, notifySlackChannelsUponFailure, notifySlackChannelsUponRetry, notifySlackChannelsUponSuccess, notifySlackCustomIcon, platform, resourceStatus, startedAt, status, terminalAt, terminationHooksDispatchedAt, updatedAt, limit, offset)

List CiRuns matching a filter

List CiRuns matching a filter. The CiRuns would have to re-queried directly to load any related resources. Results are ordered by start time, starting at most recent.

### Example

```ts
import {
  Configuration,
  CiRunsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiCiRunsV3GetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new CiRunsApi();

  const body = {
    // string (optional)
    argoWorkflowsName: argoWorkflowsName_example,
    // string (optional)
    argoWorkflowsNamespace: argoWorkflowsNamespace_example,
    // string (optional)
    argoWorkflowsTemplate: argoWorkflowsTemplate_example,
    // Date (optional)
    createdAt: 2013-10-20T19:20:30+01:00,
    // number (optional)
    githubActionsAttemptNumber: 56,
    // string (optional)
    githubActionsOwner: githubActionsOwner_example,
    // string (optional)
    githubActionsRepo: githubActionsRepo_example,
    // number (optional)
    githubActionsRunID: 56,
    // string (optional)
    githubActionsWorkflowPath: githubActionsWorkflowPath_example,
    // number (optional)
    id: 56,
    // Array<string> | Slack channels to notify if this CiRun fails. This field is always appended to when mutated. (optional)
    notifySlackChannelsUponFailure: ...,
    // Array<string> | Slack channels to notify if this CiRun is retried. This field is always appended to when mutated. It will de-dupe with the other notify fields. (optional)
    notifySlackChannelsUponRetry: ...,
    // Array<string> | Slack channels to notify if this CiRun succeeds. This field is always appended to when mutated. (optional)
    notifySlackChannelsUponSuccess: ...,
    // string | Icon to use for success or failure Slack notifications. Can be given either as a URL to an image or as a Slack emoji (using colon shortcodes, like :smiley:). An empty string is ignored to facilitate calling from GitHub Actions (where it\'s easier to pass an empty string than not send the field at all). (optional)
    notifySlackCustomIcon: notifySlackCustomIcon_example,
    // string (optional)
    platform: platform_example,
    // string | Available only when querying a CiRun via a CiIdentifier, indicates the status of the run for that resource (optional)
    resourceStatus: resourceStatus_example,
    // string (optional)
    startedAt: startedAt_example,
    // string (optional)
    status: status_example,
    // string (optional)
    terminalAt: terminalAt_example,
    // Date (optional)
    terminationHooksDispatchedAt: 2013-10-20T19:20:30+01:00,
    // Date (optional)
    updatedAt: 2013-10-20T19:20:30+01:00,
    // number | Control how many CiRuns are returned (default 100) (optional)
    limit: 56,
    // number | Control the offset for the returned CiRuns (default 0) (optional)
    offset: 56,
  } satisfies ApiCiRunsV3GetRequest;

  try {
    const data = await api.apiCiRunsV3Get(body);
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
| **argoWorkflowsName** | `string` |  | [Optional] [Defaults to `undefined`] |
| **argoWorkflowsNamespace** | `string` |  | [Optional] [Defaults to `undefined`] |
| **argoWorkflowsTemplate** | `string` |  | [Optional] [Defaults to `undefined`] |
| **createdAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **githubActionsAttemptNumber** | `number` |  | [Optional] [Defaults to `undefined`] |
| **githubActionsOwner** | `string` |  | [Optional] [Defaults to `undefined`] |
| **githubActionsRepo** | `string` |  | [Optional] [Defaults to `undefined`] |
| **githubActionsRunID** | `number` |  | [Optional] [Defaults to `undefined`] |
| **githubActionsWorkflowPath** | `string` |  | [Optional] [Defaults to `undefined`] |
| **id** | `number` |  | [Optional] [Defaults to `undefined`] |
| **notifySlackChannelsUponFailure** | `Array<string>` | Slack channels to notify if this CiRun fails. This field is always appended to when mutated. | [Optional] |
| **notifySlackChannelsUponRetry** | `Array<string>` | Slack channels to notify if this CiRun is retried. This field is always appended to when mutated. It will de-dupe with the other notify fields. | [Optional] |
| **notifySlackChannelsUponSuccess** | `Array<string>` | Slack channels to notify if this CiRun succeeds. This field is always appended to when mutated. | [Optional] |
| **notifySlackCustomIcon** | `string` | Icon to use for success or failure Slack notifications. Can be given either as a URL to an image or as a Slack emoji (using colon shortcodes, like :smiley:). An empty string is ignored to facilitate calling from GitHub Actions (where it\&#39;s easier to pass an empty string than not send the field at all). | [Optional] [Defaults to `undefined`] |
| **platform** | `string` |  | [Optional] [Defaults to `undefined`] |
| **resourceStatus** | `string` | Available only when querying a CiRun via a CiIdentifier, indicates the status of the run for that resource | [Optional] [Defaults to `undefined`] |
| **startedAt** | `string` |  | [Optional] [Defaults to `undefined`] |
| **status** | `string` |  | [Optional] [Defaults to `undefined`] |
| **terminalAt** | `string` |  | [Optional] [Defaults to `undefined`] |
| **terminationHooksDispatchedAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **updatedAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **limit** | `number` | Control how many CiRuns are returned (default 100) | [Optional] [Defaults to `undefined`] |
| **offset** | `number` | Control the offset for the returned CiRuns (default 0) | [Optional] [Defaults to `undefined`] |

### Return type

[**Array&lt;SherlockCiRunV3&gt;**](SherlockCiRunV3.md)

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


## apiCiRunsV3Put

> SherlockCiRunV3 apiCiRunsV3Put(ciRun)

Create or update a CiRun

Create or update a CiRun with timing, status, and related resource information. This endpoint is idempotent. The fields for clusters, charts, chart releases, environments, etc. all accept selectors, and they will be smart about \&quot;spreading\&quot; to indirect relations. More info is available on the CiRunV3Upsert data type, but the gist is that specifying a changeset implies its chart release (and optionally app/chart versions), specifying or implying a chart release implies its environment/cluster, and specifying an environment/cluster implies all chart releases they contain.

### Example

```ts
import {
  Configuration,
  CiRunsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiCiRunsV3PutRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new CiRunsApi();

  const body = {
    // SherlockCiRunV3Upsert | The CiRun to upsert
    ciRun: ...,
  } satisfies ApiCiRunsV3PutRequest;

  try {
    const data = await api.apiCiRunsV3Put(body);
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
| **ciRun** | [SherlockCiRunV3Upsert](SherlockCiRunV3Upsert.md) | The CiRun to upsert | |

### Return type

[**SherlockCiRunV3**](SherlockCiRunV3.md)

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


## apiCiRunsV3SelectorGet

> SherlockCiRunV3 apiCiRunsV3SelectorGet(selector)

Get a CiRun, including CiIdentifiers for related resources

Get a CiRun, including CiIdentifiers representing related resources or resources it affected.

### Example

```ts
import {
  Configuration,
  CiRunsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiCiRunsV3SelectorGetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new CiRunsApi();

  const body = {
    // string | The selector of the CiRun, which can be either its numeric ID, \'github-actions/{owner}/{repo}/{run ID}/{attempt}\', or \'argo-workflows/{namespace}/{name}\'
    selector: selector_example,
  } satisfies ApiCiRunsV3SelectorGetRequest;

  try {
    const data = await api.apiCiRunsV3SelectorGet(body);
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
| **selector** | `string` | The selector of the CiRun, which can be either its numeric ID, \&#39;github-actions/{owner}/{repo}/{run ID}/{attempt}\&#39;, or \&#39;argo-workflows/{namespace}/{name}\&#39; | [Defaults to `undefined`] |

### Return type

[**SherlockCiRunV3**](SherlockCiRunV3.md)

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

