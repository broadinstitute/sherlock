# GithubActionsJobsApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**apiGithubActionsJobsV3Get**](GithubActionsJobsApi.md#apigithubactionsjobsv3get) | **GET** /api/github-actions-jobs/v3 | List GithubActionsJobs matching a filter |
| [**apiGithubActionsJobsV3Put**](GithubActionsJobsApi.md#apigithubactionsjobsv3put) | **PUT** /api/github-actions-jobs/v3 | Upsert GithubActionsJob |
| [**apiGithubActionsJobsV3SelectorGet**](GithubActionsJobsApi.md#apigithubactionsjobsv3selectorget) | **GET** /api/github-actions-jobs/v3/{selector} | Get an individual GithubActionsJob |



## apiGithubActionsJobsV3Get

> Array&lt;SherlockGithubActionsJobV3&gt; apiGithubActionsJobsV3Get(createdAt, githubActionsAttemptNumber, githubActionsJobID, githubActionsOwner, githubActionsRepo, githubActionsRunID, id, jobCreatedAt, jobStartedAt, jobTerminalAt, status, updatedAt, limit, offset)

List GithubActionsJobs matching a filter

List GithubActionsJobs matching a filter. Results are ordered by start time, starting at most recent.

### Example

```ts
import {
  Configuration,
  GithubActionsJobsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiGithubActionsJobsV3GetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new GithubActionsJobsApi();

  const body = {
    // Date (optional)
    createdAt: 2013-10-20T19:20:30+01:00,
    // number (optional)
    githubActionsAttemptNumber: 56,
    // number (optional)
    githubActionsJobID: 56,
    // string (optional)
    githubActionsOwner: githubActionsOwner_example,
    // string (optional)
    githubActionsRepo: githubActionsRepo_example,
    // number (optional)
    githubActionsRunID: 56,
    // number (optional)
    id: 56,
    // Date (optional)
    jobCreatedAt: 2013-10-20T19:20:30+01:00,
    // Date (optional)
    jobStartedAt: 2013-10-20T19:20:30+01:00,
    // Date (optional)
    jobTerminalAt: 2013-10-20T19:20:30+01:00,
    // string (optional)
    status: status_example,
    // Date (optional)
    updatedAt: 2013-10-20T19:20:30+01:00,
    // number | Control how many GithubActionsJobs are returned (default 100) (optional)
    limit: 56,
    // number | Control the offset for the returned GithubActionsJobs (default 0) (optional)
    offset: 56,
  } satisfies ApiGithubActionsJobsV3GetRequest;

  try {
    const data = await api.apiGithubActionsJobsV3Get(body);
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
| **githubActionsAttemptNumber** | `number` |  | [Optional] [Defaults to `undefined`] |
| **githubActionsJobID** | `number` |  | [Optional] [Defaults to `undefined`] |
| **githubActionsOwner** | `string` |  | [Optional] [Defaults to `undefined`] |
| **githubActionsRepo** | `string` |  | [Optional] [Defaults to `undefined`] |
| **githubActionsRunID** | `number` |  | [Optional] [Defaults to `undefined`] |
| **id** | `number` |  | [Optional] [Defaults to `undefined`] |
| **jobCreatedAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **jobStartedAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **jobTerminalAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **status** | `string` |  | [Optional] [Defaults to `undefined`] |
| **updatedAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **limit** | `number` | Control how many GithubActionsJobs are returned (default 100) | [Optional] [Defaults to `undefined`] |
| **offset** | `number` | Control the offset for the returned GithubActionsJobs (default 0) | [Optional] [Defaults to `undefined`] |

### Return type

[**Array&lt;SherlockGithubActionsJobV3&gt;**](SherlockGithubActionsJobV3.md)

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


## apiGithubActionsJobsV3Put

> SherlockGithubActionsJobV3 apiGithubActionsJobsV3Put(githubActionsJob)

Upsert GithubActionsJob

Upsert GithubActionsJob.

### Example

```ts
import {
  Configuration,
  GithubActionsJobsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiGithubActionsJobsV3PutRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new GithubActionsJobsApi();

  const body = {
    // SherlockGithubActionsJobV3Create | The GithubActionsJob to upsert
    githubActionsJob: ...,
  } satisfies ApiGithubActionsJobsV3PutRequest;

  try {
    const data = await api.apiGithubActionsJobsV3Put(body);
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
| **githubActionsJob** | [SherlockGithubActionsJobV3Create](SherlockGithubActionsJobV3Create.md) | The GithubActionsJob to upsert | |

### Return type

[**SherlockGithubActionsJobV3**](SherlockGithubActionsJobV3.md)

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


## apiGithubActionsJobsV3SelectorGet

> SherlockGithubActionsJobV3 apiGithubActionsJobsV3SelectorGet(selector)

Get an individual GithubActionsJob

Get an individual GithubActionsJob.

### Example

```ts
import {
  Configuration,
  GithubActionsJobsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiGithubActionsJobsV3SelectorGetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new GithubActionsJobsApi();

  const body = {
    // string | The selector of the GithubActionsJob, either Sherlock ID or \'{owner}/{repo}/{job ID}\'
    selector: selector_example,
  } satisfies ApiGithubActionsJobsV3SelectorGetRequest;

  try {
    const data = await api.apiGithubActionsJobsV3SelectorGet(body);
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
| **selector** | `string` | The selector of the GithubActionsJob, either Sherlock ID or \&#39;{owner}/{repo}/{job ID}\&#39; | [Defaults to `undefined`] |

### Return type

[**SherlockGithubActionsJobV3**](SherlockGithubActionsJobV3.md)

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

