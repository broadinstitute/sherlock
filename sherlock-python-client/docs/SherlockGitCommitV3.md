# SherlockGitCommitV3


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**committed_at** | **str** |  | [optional] 
**created_at** | **datetime** |  | [optional] 
**git_branch** | **str** |  | [optional] 
**git_commit** | **str** |  | [optional] 
**git_repo** | **str** |  | [optional] 
**id** | **int** |  | [optional] 
**is_main_branch** | **bool** |  | [optional] 
**sec_since_prev** | **int** |  | [optional] 
**updated_at** | **datetime** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_git_commit_v3 import SherlockGitCommitV3

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockGitCommitV3 from a JSON string
sherlock_git_commit_v3_instance = SherlockGitCommitV3.from_json(json)
# print the JSON string representation of the object
print(SherlockGitCommitV3.to_json())

# convert the object into a dict
sherlock_git_commit_v3_dict = sherlock_git_commit_v3_instance.to_dict()
# create an instance of SherlockGitCommitV3 from a dict
sherlock_git_commit_v3_from_dict = SherlockGitCommitV3.from_dict(sherlock_git_commit_v3_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


