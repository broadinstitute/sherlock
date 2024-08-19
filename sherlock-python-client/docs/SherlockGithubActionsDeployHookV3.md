# SherlockGithubActionsDeployHookV3


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**created_at** | **datetime** |  | [optional] 
**github_actions_default_ref** | **str** |  | [optional] 
**github_actions_owner** | **str** |  | [optional] 
**github_actions_ref_behavior** | **str** | This field determines what git ref the workflow will be run on. The default of always-use-default-ref always uses the default ref; use-app-version-as-ref will use the app version (when available) as the ref, useful when versions are always commit hashes or tags; use-app-version-commit-as-ref will use the app version&#39;s commit (when available) as the ref, useful when the repo is configured to fully report app versions to Sherlock. | [optional] [default to 'always-use-default-ref']
**github_actions_repo** | **str** |  | [optional] 
**github_actions_workflow_inputs** | **object** | These workflow inputs will be passed statically as-is to GitHub&#39;s workflow dispatch API (https://docs.github.com/en/rest/actions/workflows#create-a-workflow-dispatch-event) as the &#x60;inputs&#x60; parameter object. | [optional] 
**github_actions_workflow_path** | **str** |  | [optional] 
**id** | **int** |  | [optional] 
**on_chart_release** | **str** |  | [optional] 
**on_environment** | **str** |  | [optional] 
**on_failure** | **bool** |  | [optional] 
**on_success** | **bool** |  | [optional] 
**updated_at** | **datetime** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_github_actions_deploy_hook_v3 import SherlockGithubActionsDeployHookV3

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockGithubActionsDeployHookV3 from a JSON string
sherlock_github_actions_deploy_hook_v3_instance = SherlockGithubActionsDeployHookV3.from_json(json)
# print the JSON string representation of the object
print(SherlockGithubActionsDeployHookV3.to_json())

# convert the object into a dict
sherlock_github_actions_deploy_hook_v3_dict = sherlock_github_actions_deploy_hook_v3_instance.to_dict()
# create an instance of SherlockGithubActionsDeployHookV3 from a dict
sherlock_github_actions_deploy_hook_v3_from_dict = SherlockGithubActionsDeployHookV3.from_dict(sherlock_github_actions_deploy_hook_v3_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


