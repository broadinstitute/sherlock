# SherlockSlackDeployHookV3


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**created_at** | **datetime** |  | [optional] 
**id** | **int** |  | [optional] 
**mention_people** | **bool** |  | [optional] 
**on_chart_release** | **str** |  | [optional] 
**on_environment** | **str** |  | [optional] 
**on_failure** | **bool** |  | [optional] 
**on_success** | **bool** |  | [optional] 
**slack_channel** | **str** |  | [optional] 
**updated_at** | **datetime** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_slack_deploy_hook_v3 import SherlockSlackDeployHookV3

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockSlackDeployHookV3 from a JSON string
sherlock_slack_deploy_hook_v3_instance = SherlockSlackDeployHookV3.from_json(json)
# print the JSON string representation of the object
print(SherlockSlackDeployHookV3.to_json())

# convert the object into a dict
sherlock_slack_deploy_hook_v3_dict = sherlock_slack_deploy_hook_v3_instance.to_dict()
# create an instance of SherlockSlackDeployHookV3 from a dict
sherlock_slack_deploy_hook_v3_from_dict = SherlockSlackDeployHookV3.from_dict(sherlock_slack_deploy_hook_v3_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


