# SherlockGithubActionsDeployHookV3Edit


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**github_actions_default_ref** | **str** |  | [optional] 
**github_actions_owner** | **str** |  | [optional] 
**github_actions_ref_behavior** | **str** | This field determines what git ref the workflow will be run on. The default of always-use-default-ref always uses the default ref; use-app-version-as-ref will use the app version (when available) as the ref, useful when versions are always commit hashes or tags; use-app-version-commit-as-ref will use the app version&#39;s commit (when available) as the ref, useful when the repo is configured to fully report app versions to Sherlock. | [optional] [default to 'always-use-default-ref']
**github_actions_repo** | **str** |  | [optional] 
**github_actions_workflow_inputs** | **object** | These workflow inputs will be passed statically as-is to GitHub&#39;s workflow dispatch API (https://docs.github.com/en/rest/actions/workflows#create-a-workflow-dispatch-event) as the &#x60;inputs&#x60; parameter object. | [optional] 
**github_actions_workflow_path** | **str** |  | [optional] 
**on_failure** | **bool** |  | [optional] 
**on_success** | **bool** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_github_actions_deploy_hook_v3_edit import SherlockGithubActionsDeployHookV3Edit

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockGithubActionsDeployHookV3Edit from a JSON string
sherlock_github_actions_deploy_hook_v3_edit_instance = SherlockGithubActionsDeployHookV3Edit.from_json(json)
# print the JSON string representation of the object
print(SherlockGithubActionsDeployHookV3Edit.to_json())

# convert the object into a dict
sherlock_github_actions_deploy_hook_v3_edit_dict = sherlock_github_actions_deploy_hook_v3_edit_instance.to_dict()
# create an instance of SherlockGithubActionsDeployHookV3Edit from a dict
sherlock_github_actions_deploy_hook_v3_edit_from_dict = SherlockGithubActionsDeployHookV3Edit.from_dict(sherlock_github_actions_deploy_hook_v3_edit_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


