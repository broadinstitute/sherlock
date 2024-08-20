# SherlockUserV3


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**assignments** | [**List[SherlockRoleAssignmentV3]**](SherlockRoleAssignmentV3.md) |  | [optional] 
**created_at** | **datetime** |  | [optional] 
**email** | **str** |  | [optional] 
**github_id** | **str** |  | [optional] 
**github_username** | **str** |  | [optional] 
**google_id** | **str** |  | [optional] 
**id** | **int** |  | [optional] 
**name** | **str** |  | [optional] 
**name_from** | **str** |  | [optional] 
**name_inferred_from_github** | **bool** | Controls whether Sherlock should automatically update the user&#39;s name based on a connected GitHub identity. Will be set to true if the user account has no name and a GitHub account is linked. | [optional] 
**slack_id** | **str** |  | [optional] 
**slack_username** | **str** |  | [optional] 
**suitability_description** | **str** | Available only in responses; describes the user&#39;s production-suitability | [optional] 
**suitable** | **bool** | Available only in responses; indicates whether the user is production-suitable | [optional] 
**updated_at** | **datetime** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_user_v3 import SherlockUserV3

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockUserV3 from a JSON string
sherlock_user_v3_instance = SherlockUserV3.from_json(json)
# print the JSON string representation of the object
print(SherlockUserV3.to_json())

# convert the object into a dict
sherlock_user_v3_dict = sherlock_user_v3_instance.to_dict()
# create an instance of SherlockUserV3 from a dict
sherlock_user_v3_from_dict = SherlockUserV3.from_dict(sherlock_user_v3_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


