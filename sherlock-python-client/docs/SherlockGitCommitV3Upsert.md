# SherlockGitCommitV3Upsert


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**committed_at** | **str** |  | [optional] 
**git_branch** | **str** |  | [optional] 
**git_commit** | **str** |  | [optional] 
**git_repo** | **str** |  | [optional] 
**is_main_branch** | **bool** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_git_commit_v3_upsert import SherlockGitCommitV3Upsert

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockGitCommitV3Upsert from a JSON string
sherlock_git_commit_v3_upsert_instance = SherlockGitCommitV3Upsert.from_json(json)
# print the JSON string representation of the object
print(SherlockGitCommitV3Upsert.to_json())

# convert the object into a dict
sherlock_git_commit_v3_upsert_dict = sherlock_git_commit_v3_upsert_instance.to_dict()
# create an instance of SherlockGitCommitV3Upsert from a dict
sherlock_git_commit_v3_upsert_from_dict = SherlockGitCommitV3Upsert.from_dict(sherlock_git_commit_v3_upsert_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


