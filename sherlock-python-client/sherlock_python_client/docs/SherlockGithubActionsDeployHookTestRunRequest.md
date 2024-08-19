# SherlockGithubActionsDeployHookTestRunRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**execute** | **bool** | Required, whether to fully run the GHA | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_github_actions_deploy_hook_test_run_request import SherlockGithubActionsDeployHookTestRunRequest

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockGithubActionsDeployHookTestRunRequest from a JSON string
sherlock_github_actions_deploy_hook_test_run_request_instance = SherlockGithubActionsDeployHookTestRunRequest.from_json(json)
# print the JSON string representation of the object
print(SherlockGithubActionsDeployHookTestRunRequest.to_json())

# convert the object into a dict
sherlock_github_actions_deploy_hook_test_run_request_dict = sherlock_github_actions_deploy_hook_test_run_request_instance.to_dict()
# create an instance of SherlockGithubActionsDeployHookTestRunRequest from a dict
sherlock_github_actions_deploy_hook_test_run_request_from_dict = SherlockGithubActionsDeployHookTestRunRequest.from_dict(sherlock_github_actions_deploy_hook_test_run_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


