# sherlock_python_client.GitCommitsApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_git_commits_v3_put**](GitCommitsApi.md#api_git_commits_v3_put) | **PUT** /api/git-commits/v3 | Upsert a GitCommit


# **api_git_commits_v3_put**
> SherlockGitCommitV3 api_git_commits_v3_put(git_commit)

Upsert a GitCommit

Upsert a GitCommit.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.sherlock_git_commit_v3 import SherlockGitCommitV3
from sherlock_python_client.models.sherlock_git_commit_v3_upsert import SherlockGitCommitV3Upsert
from sherlock_python_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://sherlock.dsp-devops-prod.broadinstitute.org
# See configuration.py for a list of all supported configuration parameters.
configuration = sherlock_python_client.Configuration(
    host = "https://sherlock.dsp-devops-prod.broadinstitute.org"
)


# Enter a context with an instance of the API client
with sherlock_python_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = sherlock_python_client.GitCommitsApi(api_client)
    git_commit = sherlock_python_client.SherlockGitCommitV3Upsert() # SherlockGitCommitV3Upsert | The GitCommit to upsert

    try:
        # Upsert a GitCommit
        api_response = api_instance.api_git_commits_v3_put(git_commit)
        print("The response of GitCommitsApi->api_git_commits_v3_put:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling GitCommitsApi->api_git_commits_v3_put: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **git_commit** | [**SherlockGitCommitV3Upsert**](SherlockGitCommitV3Upsert.md)| The GitCommit to upsert | 

### Return type

[**SherlockGitCommitV3**](SherlockGitCommitV3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Created |  -  |
**400** | Bad Request |  -  |
**403** | Forbidden |  -  |
**404** | Not Found |  -  |
**407** | Proxy Authentication Required |  -  |
**409** | Conflict |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

