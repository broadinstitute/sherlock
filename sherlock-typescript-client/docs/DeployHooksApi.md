# DeployHooksApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**apiDeployHooksGithubActionsProceduresV3TestSelectorPost**](DeployHooksApi.md#apideployhooksgithubactionsproceduresv3testselectorpost) | **POST** /api/deploy-hooks/github-actions/procedures/v3/test/{selector} | Test a GithubActionsDeployHook |
| [**apiDeployHooksGithubActionsV3Get**](DeployHooksApi.md#apideployhooksgithubactionsv3get) | **GET** /api/deploy-hooks/github-actions/v3 | List GithubActionsDeployHooks matching a filter |
| [**apiDeployHooksGithubActionsV3Post**](DeployHooksApi.md#apideployhooksgithubactionsv3post) | **POST** /api/deploy-hooks/github-actions/v3 | Create a GithubActionsDeployHook |
| [**apiDeployHooksGithubActionsV3SelectorDelete**](DeployHooksApi.md#apideployhooksgithubactionsv3selectordelete) | **DELETE** /api/deploy-hooks/github-actions/v3/{selector} | Delete an individual GithubActionsDeployHook |
| [**apiDeployHooksGithubActionsV3SelectorGet**](DeployHooksApi.md#apideployhooksgithubactionsv3selectorget) | **GET** /api/deploy-hooks/github-actions/v3/{selector} | Get an individual GithubActionsDeployHook |
| [**apiDeployHooksGithubActionsV3SelectorPatch**](DeployHooksApi.md#apideployhooksgithubactionsv3selectorpatch) | **PATCH** /api/deploy-hooks/github-actions/v3/{selector} | Edit an individual GithubActionsDeployHook |
| [**apiDeployHooksSlackProceduresV3TestSelectorPost**](DeployHooksApi.md#apideployhooksslackproceduresv3testselectorpost) | **POST** /api/deploy-hooks/slack/procedures/v3/test/{selector} | Test a SlackDeployHook |
| [**apiDeployHooksSlackV3Get**](DeployHooksApi.md#apideployhooksslackv3get) | **GET** /api/deploy-hooks/slack/v3 | List SlackDeployHooks matching a filter |
| [**apiDeployHooksSlackV3Post**](DeployHooksApi.md#apideployhooksslackv3post) | **POST** /api/deploy-hooks/slack/v3 | Create a SlackDeployHook |
| [**apiDeployHooksSlackV3SelectorDelete**](DeployHooksApi.md#apideployhooksslackv3selectordelete) | **DELETE** /api/deploy-hooks/slack/v3/{selector} | Delete an individual SlackDeployHook |
| [**apiDeployHooksSlackV3SelectorGet**](DeployHooksApi.md#apideployhooksslackv3selectorget) | **GET** /api/deploy-hooks/slack/v3/{selector} | Get an individual SlackDeployHook |
| [**apiDeployHooksSlackV3SelectorPatch**](DeployHooksApi.md#apideployhooksslackv3selectorpatch) | **PATCH** /api/deploy-hooks/slack/v3/{selector} | Edit an individual SlackDeployHook |



## apiDeployHooksGithubActionsProceduresV3TestSelectorPost

> SherlockGithubActionsDeployHookTestRunResponse apiDeployHooksGithubActionsProceduresV3TestSelectorPost(selector, request)

Test a GithubActionsDeployHook

Run a GitHub Action to simulate a GithubActionsDeployHook

### Example

```ts
import {
  Configuration,
  DeployHooksApi,
} from '@sherlock-js-client/sherlock';
import type { ApiDeployHooksGithubActionsProceduresV3TestSelectorPostRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new DeployHooksApi();

  const body = {
    // string | The ID of the GithubActionsDeployHook
    selector: selector_example,
    // SherlockGithubActionsDeployHookTestRunRequest | Whether to fully execute the hook (JSON body helps with CSRF protection)
    request: ...,
  } satisfies ApiDeployHooksGithubActionsProceduresV3TestSelectorPostRequest;

  try {
    const data = await api.apiDeployHooksGithubActionsProceduresV3TestSelectorPost(body);
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
| **selector** | `string` | The ID of the GithubActionsDeployHook | [Defaults to `undefined`] |
| **request** | [SherlockGithubActionsDeployHookTestRunRequest](SherlockGithubActionsDeployHookTestRunRequest.md) | Whether to fully execute the hook (JSON body helps with CSRF protection) | |

### Return type

[**SherlockGithubActionsDeployHookTestRunResponse**](SherlockGithubActionsDeployHookTestRunResponse.md)

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


## apiDeployHooksGithubActionsV3Get

> Array&lt;SherlockGithubActionsDeployHookV3&gt; apiDeployHooksGithubActionsV3Get(createdAt, githubActionsDefaultRef, githubActionsOwner, githubActionsRefBehavior, githubActionsRepo, githubActionsWorkflowPath, id, onChartRelease, onEnvironment, onFailure, onSuccess, updatedAt, limit, offset)

List GithubActionsDeployHooks matching a filter

List GithubActionsDeployHooks matching a filter.

### Example

```ts
import {
  Configuration,
  DeployHooksApi,
} from '@sherlock-js-client/sherlock';
import type { ApiDeployHooksGithubActionsV3GetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new DeployHooksApi();

  const body = {
    // Date (optional)
    createdAt: 2013-10-20T19:20:30+01:00,
    // string (optional)
    githubActionsDefaultRef: githubActionsDefaultRef_example,
    // string (optional)
    githubActionsOwner: githubActionsOwner_example,
    // 'always-use-default-ref' | 'use-app-version-as-ref' | 'use-app-version-commit-as-ref' | This field determines what git ref the workflow will be run on. The default of always-use-default-ref always uses the default ref; use-app-version-as-ref will use the app version (when available) as the ref, useful when versions are always commit hashes or tags; use-app-version-commit-as-ref will use the app version\'s commit (when available) as the ref, useful when the repo is configured to fully report app versions to Sherlock. (optional)
    githubActionsRefBehavior: githubActionsRefBehavior_example,
    // string (optional)
    githubActionsRepo: githubActionsRepo_example,
    // string (optional)
    githubActionsWorkflowPath: githubActionsWorkflowPath_example,
    // number (optional)
    id: 56,
    // string (optional)
    onChartRelease: onChartRelease_example,
    // string (optional)
    onEnvironment: onEnvironment_example,
    // boolean (optional)
    onFailure: true,
    // boolean (optional)
    onSuccess: true,
    // Date (optional)
    updatedAt: 2013-10-20T19:20:30+01:00,
    // number | Control how many GithubActionsDeployHooks are returned (default 100) (optional)
    limit: 56,
    // number | Control the offset for the returned GithubActionsDeployHooks (default 0) (optional)
    offset: 56,
  } satisfies ApiDeployHooksGithubActionsV3GetRequest;

  try {
    const data = await api.apiDeployHooksGithubActionsV3Get(body);
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
| **githubActionsDefaultRef** | `string` |  | [Optional] [Defaults to `undefined`] |
| **githubActionsOwner** | `string` |  | [Optional] [Defaults to `undefined`] |
| **githubActionsRefBehavior** | `always-use-default-ref`, `use-app-version-as-ref`, `use-app-version-commit-as-ref` | This field determines what git ref the workflow will be run on. The default of always-use-default-ref always uses the default ref; use-app-version-as-ref will use the app version (when available) as the ref, useful when versions are always commit hashes or tags; use-app-version-commit-as-ref will use the app version\&#39;s commit (when available) as the ref, useful when the repo is configured to fully report app versions to Sherlock. | [Optional] [Defaults to `&#39;always-use-default-ref&#39;`] [Enum: always-use-default-ref, use-app-version-as-ref, use-app-version-commit-as-ref] |
| **githubActionsRepo** | `string` |  | [Optional] [Defaults to `undefined`] |
| **githubActionsWorkflowPath** | `string` |  | [Optional] [Defaults to `undefined`] |
| **id** | `number` |  | [Optional] [Defaults to `undefined`] |
| **onChartRelease** | `string` |  | [Optional] [Defaults to `undefined`] |
| **onEnvironment** | `string` |  | [Optional] [Defaults to `undefined`] |
| **onFailure** | `boolean` |  | [Optional] [Defaults to `undefined`] |
| **onSuccess** | `boolean` |  | [Optional] [Defaults to `undefined`] |
| **updatedAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **limit** | `number` | Control how many GithubActionsDeployHooks are returned (default 100) | [Optional] [Defaults to `undefined`] |
| **offset** | `number` | Control the offset for the returned GithubActionsDeployHooks (default 0) | [Optional] [Defaults to `undefined`] |

### Return type

[**Array&lt;SherlockGithubActionsDeployHookV3&gt;**](SherlockGithubActionsDeployHookV3.md)

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


## apiDeployHooksGithubActionsV3Post

> SherlockGithubActionsDeployHookV3 apiDeployHooksGithubActionsV3Post(githubActionsDeployHook)

Create a GithubActionsDeployHook

Create a GithubActionsDeployHook.

### Example

```ts
import {
  Configuration,
  DeployHooksApi,
} from '@sherlock-js-client/sherlock';
import type { ApiDeployHooksGithubActionsV3PostRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new DeployHooksApi();

  const body = {
    // SherlockGithubActionsDeployHookV3Create | The GithubActionsDeployHook to create
    githubActionsDeployHook: ...,
  } satisfies ApiDeployHooksGithubActionsV3PostRequest;

  try {
    const data = await api.apiDeployHooksGithubActionsV3Post(body);
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
| **githubActionsDeployHook** | [SherlockGithubActionsDeployHookV3Create](SherlockGithubActionsDeployHookV3Create.md) | The GithubActionsDeployHook to create | |

### Return type

[**SherlockGithubActionsDeployHookV3**](SherlockGithubActionsDeployHookV3.md)

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


## apiDeployHooksGithubActionsV3SelectorDelete

> SherlockGithubActionsDeployHookV3 apiDeployHooksGithubActionsV3SelectorDelete(selector)

Delete an individual GithubActionsDeployHook

Delete an individual GithubActionsDeployHook by its ID.

### Example

```ts
import {
  Configuration,
  DeployHooksApi,
} from '@sherlock-js-client/sherlock';
import type { ApiDeployHooksGithubActionsV3SelectorDeleteRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new DeployHooksApi();

  const body = {
    // string | The ID of the GithubActionsDeployHook
    selector: selector_example,
  } satisfies ApiDeployHooksGithubActionsV3SelectorDeleteRequest;

  try {
    const data = await api.apiDeployHooksGithubActionsV3SelectorDelete(body);
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
| **selector** | `string` | The ID of the GithubActionsDeployHook | [Defaults to `undefined`] |

### Return type

[**SherlockGithubActionsDeployHookV3**](SherlockGithubActionsDeployHookV3.md)

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


## apiDeployHooksGithubActionsV3SelectorGet

> SherlockGithubActionsDeployHookV3 apiDeployHooksGithubActionsV3SelectorGet(selector)

Get an individual GithubActionsDeployHook

Get an individual GithubActionsDeployHook by its ID.

### Example

```ts
import {
  Configuration,
  DeployHooksApi,
} from '@sherlock-js-client/sherlock';
import type { ApiDeployHooksGithubActionsV3SelectorGetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new DeployHooksApi();

  const body = {
    // string | The ID of the GithubActionsDeployHook
    selector: selector_example,
  } satisfies ApiDeployHooksGithubActionsV3SelectorGetRequest;

  try {
    const data = await api.apiDeployHooksGithubActionsV3SelectorGet(body);
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
| **selector** | `string` | The ID of the GithubActionsDeployHook | [Defaults to `undefined`] |

### Return type

[**SherlockGithubActionsDeployHookV3**](SherlockGithubActionsDeployHookV3.md)

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


## apiDeployHooksGithubActionsV3SelectorPatch

> SherlockGithubActionsDeployHookV3 apiDeployHooksGithubActionsV3SelectorPatch(selector, githubActionsDeployHook)

Edit an individual GithubActionsDeployHook

Edit an individual GithubActionsDeployHook by its ID.

### Example

```ts
import {
  Configuration,
  DeployHooksApi,
} from '@sherlock-js-client/sherlock';
import type { ApiDeployHooksGithubActionsV3SelectorPatchRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new DeployHooksApi();

  const body = {
    // string | The ID of the GithubActionsDeployHook to edit
    selector: selector_example,
    // SherlockGithubActionsDeployHookV3Edit | The edits to make to the GithubActionsDeployHook
    githubActionsDeployHook: ...,
  } satisfies ApiDeployHooksGithubActionsV3SelectorPatchRequest;

  try {
    const data = await api.apiDeployHooksGithubActionsV3SelectorPatch(body);
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
| **selector** | `string` | The ID of the GithubActionsDeployHook to edit | [Defaults to `undefined`] |
| **githubActionsDeployHook** | [SherlockGithubActionsDeployHookV3Edit](SherlockGithubActionsDeployHookV3Edit.md) | The edits to make to the GithubActionsDeployHook | |

### Return type

[**SherlockGithubActionsDeployHookV3**](SherlockGithubActionsDeployHookV3.md)

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


## apiDeployHooksSlackProceduresV3TestSelectorPost

> SherlockSlackDeployHookTestRunResponse apiDeployHooksSlackProceduresV3TestSelectorPost(selector, request)

Test a SlackDeployHook

Send a Slack message to simulate a SlackDeployHook

### Example

```ts
import {
  Configuration,
  DeployHooksApi,
} from '@sherlock-js-client/sherlock';
import type { ApiDeployHooksSlackProceduresV3TestSelectorPostRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new DeployHooksApi();

  const body = {
    // string | The ID of the SlackDeployHook to test
    selector: selector_example,
    // SherlockSlackDeployHookTestRunRequest | Whether to fully execute the hook (JSON body helps with CSRF protection)
    request: ...,
  } satisfies ApiDeployHooksSlackProceduresV3TestSelectorPostRequest;

  try {
    const data = await api.apiDeployHooksSlackProceduresV3TestSelectorPost(body);
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
| **selector** | `string` | The ID of the SlackDeployHook to test | [Defaults to `undefined`] |
| **request** | [SherlockSlackDeployHookTestRunRequest](SherlockSlackDeployHookTestRunRequest.md) | Whether to fully execute the hook (JSON body helps with CSRF protection) | |

### Return type

[**SherlockSlackDeployHookTestRunResponse**](SherlockSlackDeployHookTestRunResponse.md)

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


## apiDeployHooksSlackV3Get

> Array&lt;SherlockSlackDeployHookV3&gt; apiDeployHooksSlackV3Get(createdAt, id, mentionPeople, onChartRelease, onEnvironment, onFailure, onSuccess, slackChannel, updatedAt, limit, offset)

List SlackDeployHooks matching a filter

List SlackDeployHooks matching a filter.

### Example

```ts
import {
  Configuration,
  DeployHooksApi,
} from '@sherlock-js-client/sherlock';
import type { ApiDeployHooksSlackV3GetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new DeployHooksApi();

  const body = {
    // Date (optional)
    createdAt: 2013-10-20T19:20:30+01:00,
    // number (optional)
    id: 56,
    // boolean (optional)
    mentionPeople: true,
    // string (optional)
    onChartRelease: onChartRelease_example,
    // string (optional)
    onEnvironment: onEnvironment_example,
    // boolean (optional)
    onFailure: true,
    // boolean (optional)
    onSuccess: true,
    // string (optional)
    slackChannel: slackChannel_example,
    // Date (optional)
    updatedAt: 2013-10-20T19:20:30+01:00,
    // number | Control how many SlackDeployHooks are returned (default 100) (optional)
    limit: 56,
    // number | Control the offset for the returned SlackDeployHooks (default 0) (optional)
    offset: 56,
  } satisfies ApiDeployHooksSlackV3GetRequest;

  try {
    const data = await api.apiDeployHooksSlackV3Get(body);
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
| **mentionPeople** | `boolean` |  | [Optional] [Defaults to `undefined`] |
| **onChartRelease** | `string` |  | [Optional] [Defaults to `undefined`] |
| **onEnvironment** | `string` |  | [Optional] [Defaults to `undefined`] |
| **onFailure** | `boolean` |  | [Optional] [Defaults to `undefined`] |
| **onSuccess** | `boolean` |  | [Optional] [Defaults to `undefined`] |
| **slackChannel** | `string` |  | [Optional] [Defaults to `undefined`] |
| **updatedAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **limit** | `number` | Control how many SlackDeployHooks are returned (default 100) | [Optional] [Defaults to `undefined`] |
| **offset** | `number` | Control the offset for the returned SlackDeployHooks (default 0) | [Optional] [Defaults to `undefined`] |

### Return type

[**Array&lt;SherlockSlackDeployHookV3&gt;**](SherlockSlackDeployHookV3.md)

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


## apiDeployHooksSlackV3Post

> SherlockSlackDeployHookV3 apiDeployHooksSlackV3Post(slackDeployHook)

Create a SlackDeployHook

Create a SlackDeployHook.

### Example

```ts
import {
  Configuration,
  DeployHooksApi,
} from '@sherlock-js-client/sherlock';
import type { ApiDeployHooksSlackV3PostRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new DeployHooksApi();

  const body = {
    // SherlockSlackDeployHookV3Create | The SlackDeployHook to create
    slackDeployHook: ...,
  } satisfies ApiDeployHooksSlackV3PostRequest;

  try {
    const data = await api.apiDeployHooksSlackV3Post(body);
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
| **slackDeployHook** | [SherlockSlackDeployHookV3Create](SherlockSlackDeployHookV3Create.md) | The SlackDeployHook to create | |

### Return type

[**SherlockSlackDeployHookV3**](SherlockSlackDeployHookV3.md)

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


## apiDeployHooksSlackV3SelectorDelete

> SherlockSlackDeployHookV3 apiDeployHooksSlackV3SelectorDelete(selector)

Delete an individual SlackDeployHook

Delete an individual SlackDeployHook by its ID.

### Example

```ts
import {
  Configuration,
  DeployHooksApi,
} from '@sherlock-js-client/sherlock';
import type { ApiDeployHooksSlackV3SelectorDeleteRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new DeployHooksApi();

  const body = {
    // string | The ID of the SlackDeployHook
    selector: selector_example,
  } satisfies ApiDeployHooksSlackV3SelectorDeleteRequest;

  try {
    const data = await api.apiDeployHooksSlackV3SelectorDelete(body);
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
| **selector** | `string` | The ID of the SlackDeployHook | [Defaults to `undefined`] |

### Return type

[**SherlockSlackDeployHookV3**](SherlockSlackDeployHookV3.md)

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


## apiDeployHooksSlackV3SelectorGet

> SherlockSlackDeployHookV3 apiDeployHooksSlackV3SelectorGet(selector)

Get an individual SlackDeployHook

Get an individual SlackDeployHook by its ID.

### Example

```ts
import {
  Configuration,
  DeployHooksApi,
} from '@sherlock-js-client/sherlock';
import type { ApiDeployHooksSlackV3SelectorGetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new DeployHooksApi();

  const body = {
    // string | The ID of the SlackDeployHook
    selector: selector_example,
  } satisfies ApiDeployHooksSlackV3SelectorGetRequest;

  try {
    const data = await api.apiDeployHooksSlackV3SelectorGet(body);
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
| **selector** | `string` | The ID of the SlackDeployHook | [Defaults to `undefined`] |

### Return type

[**SherlockSlackDeployHookV3**](SherlockSlackDeployHookV3.md)

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


## apiDeployHooksSlackV3SelectorPatch

> SherlockSlackDeployHookV3 apiDeployHooksSlackV3SelectorPatch(selector, slackDeployHook)

Edit an individual SlackDeployHook

Edit an individual SlackDeployHook by its ID.

### Example

```ts
import {
  Configuration,
  DeployHooksApi,
} from '@sherlock-js-client/sherlock';
import type { ApiDeployHooksSlackV3SelectorPatchRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new DeployHooksApi();

  const body = {
    // string | The ID of the SlackDeployHook to edit
    selector: selector_example,
    // SherlockSlackDeployHookV3Edit | The edits to make to the SlackDeployHook
    slackDeployHook: ...,
  } satisfies ApiDeployHooksSlackV3SelectorPatchRequest;

  try {
    const data = await api.apiDeployHooksSlackV3SelectorPatch(body);
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
| **selector** | `string` | The ID of the SlackDeployHook to edit | [Defaults to `undefined`] |
| **slackDeployHook** | [SherlockSlackDeployHookV3Edit](SherlockSlackDeployHookV3Edit.md) | The edits to make to the SlackDeployHook | |

### Return type

[**SherlockSlackDeployHookV3**](SherlockSlackDeployHookV3.md)

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

