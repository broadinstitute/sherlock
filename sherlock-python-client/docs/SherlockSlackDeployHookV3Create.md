# SherlockSlackDeployHookV3Create


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**mention_people** | **bool** |  | [optional] 
**on_chart_release** | **str** |  | [optional] 
**on_environment** | **str** |  | [optional] 
**on_failure** | **bool** |  | [optional] 
**on_success** | **bool** |  | [optional] 
**slack_channel** | **str** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_slack_deploy_hook_v3_create import SherlockSlackDeployHookV3Create

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockSlackDeployHookV3Create from a JSON string
sherlock_slack_deploy_hook_v3_create_instance = SherlockSlackDeployHookV3Create.from_json(json)
# print the JSON string representation of the object
print(SherlockSlackDeployHookV3Create.to_json())

# convert the object into a dict
sherlock_slack_deploy_hook_v3_create_dict = sherlock_slack_deploy_hook_v3_create_instance.to_dict()
# create an instance of SherlockSlackDeployHookV3Create from a dict
sherlock_slack_deploy_hook_v3_create_from_dict = SherlockSlackDeployHookV3Create.from_dict(sherlock_slack_deploy_hook_v3_create_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


