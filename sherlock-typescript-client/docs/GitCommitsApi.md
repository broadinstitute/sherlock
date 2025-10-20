# GitCommitsApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**apiGitCommitsV3Put**](GitCommitsApi.md#apigitcommitsv3put) | **PUT** /api/git-commits/v3 | Upsert a GitCommit |



## apiGitCommitsV3Put

> SherlockGitCommitV3 apiGitCommitsV3Put(gitCommit)

Upsert a GitCommit

Upsert a GitCommit.

### Example

```ts
import {
  Configuration,
  GitCommitsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiGitCommitsV3PutRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new GitCommitsApi();

  const body = {
    // SherlockGitCommitV3Upsert | The GitCommit to upsert
    gitCommit: ...,
  } satisfies ApiGitCommitsV3PutRequest;

  try {
    const data = await api.apiGitCommitsV3Put(body);
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
| **gitCommit** | [SherlockGitCommitV3Upsert](SherlockGitCommitV3Upsert.md) | The GitCommit to upsert | |

### Return type

[**SherlockGitCommitV3**](SherlockGitCommitV3.md)

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

