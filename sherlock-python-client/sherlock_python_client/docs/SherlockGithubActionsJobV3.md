# SherlockGithubActionsJobV3


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**created_at** | **datetime** |  | [optional] 
**github_actions_attempt_number** | **int** |  | [optional] 
**github_actions_job_id** | **int** |  | [optional] 
**github_actions_owner** | **str** |  | [optional] 
**github_actions_repo** | **str** |  | [optional] 
**github_actions_run_id** | **int** |  | [optional] 
**id** | **int** |  | [optional] 
**job_created_at** | **datetime** |  | [optional] 
**job_started_at** | **datetime** |  | [optional] 
**job_terminal_at** | **datetime** |  | [optional] 
**status** | **str** |  | [optional] 
**updated_at** | **datetime** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_github_actions_job_v3 import SherlockGithubActionsJobV3

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockGithubActionsJobV3 from a JSON string
sherlock_github_actions_job_v3_instance = SherlockGithubActionsJobV3.from_json(json)
# print the JSON string representation of the object
print(SherlockGithubActionsJobV3.to_json())

# convert the object into a dict
sherlock_github_actions_job_v3_dict = sherlock_github_actions_job_v3_instance.to_dict()
# create an instance of SherlockGithubActionsJobV3 from a dict
sherlock_github_actions_job_v3_from_dict = SherlockGithubActionsJobV3.from_dict(sherlock_github_actions_job_v3_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


