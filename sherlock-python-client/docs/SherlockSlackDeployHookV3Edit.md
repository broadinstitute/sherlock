# SherlockSlackDeployHookV3Edit


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**mention_people** | **bool** |  | [optional] 
**on_failure** | **bool** |  | [optional] 
**on_success** | **bool** |  | [optional] 
**slack_channel** | **str** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_slack_deploy_hook_v3_edit import SherlockSlackDeployHookV3Edit

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockSlackDeployHookV3Edit from a JSON string
sherlock_slack_deploy_hook_v3_edit_instance = SherlockSlackDeployHookV3Edit.from_json(json)
# print the JSON string representation of the object
print(SherlockSlackDeployHookV3Edit.to_json())

# convert the object into a dict
sherlock_slack_deploy_hook_v3_edit_dict = sherlock_slack_deploy_hook_v3_edit_instance.to_dict()
# create an instance of SherlockSlackDeployHookV3Edit from a dict
sherlock_slack_deploy_hook_v3_edit_from_dict = SherlockSlackDeployHookV3Edit.from_dict(sherlock_slack_deploy_hook_v3_edit_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


