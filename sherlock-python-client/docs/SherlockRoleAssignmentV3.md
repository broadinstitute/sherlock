# SherlockRoleAssignmentV3


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**expires_at** | **datetime** |  | [optional] 
**expires_in** | **str** | A Go time.Duration string that will be added to the current time to attempt to set expiresAt (this may be more convenient than setting expiresAt directly) | [optional] 
**role_info** | **object** |  | [optional] 
**suspended** | **bool** | If the assignment should be active. This field is only mutable through the API if the role doesn&#39;t automatically suspend non-suitable users | [optional] 
**user_info** | **object** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_role_assignment_v3 import SherlockRoleAssignmentV3

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockRoleAssignmentV3 from a JSON string
sherlock_role_assignment_v3_instance = SherlockRoleAssignmentV3.from_json(json)
# print the JSON string representation of the object
print(SherlockRoleAssignmentV3.to_json())

# convert the object into a dict
sherlock_role_assignment_v3_dict = sherlock_role_assignment_v3_instance.to_dict()
# create an instance of SherlockRoleAssignmentV3 from a dict
sherlock_role_assignment_v3_from_dict = SherlockRoleAssignmentV3.from_dict(sherlock_role_assignment_v3_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


