# SherlockGithubActionsJobV3Create


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**github_actions_attempt_number** | **int** |  | [optional] 
**github_actions_job_id** | **int** |  | [optional] 
**github_actions_owner** | **str** |  | [optional] 
**github_actions_repo** | **str** |  | [optional] 
**github_actions_run_id** | **int** |  | [optional] 
**job_created_at** | **datetime** |  | [optional] 
**job_started_at** | **datetime** |  | [optional] 
**job_terminal_at** | **datetime** |  | [optional] 
**status** | **str** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_github_actions_job_v3_create import SherlockGithubActionsJobV3Create

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockGithubActionsJobV3Create from a JSON string
sherlock_github_actions_job_v3_create_instance = SherlockGithubActionsJobV3Create.from_json(json)
# print the JSON string representation of the object
print(SherlockGithubActionsJobV3Create.to_json())

# convert the object into a dict
sherlock_github_actions_job_v3_create_dict = sherlock_github_actions_job_v3_create_instance.to_dict()
# create an instance of SherlockGithubActionsJobV3Create from a dict
sherlock_github_actions_job_v3_create_from_dict = SherlockGithubActionsJobV3Create.from_dict(sherlock_github_actions_job_v3_create_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


