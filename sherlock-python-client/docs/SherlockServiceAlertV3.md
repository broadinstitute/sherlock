# SherlockServiceAlertV3


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**created_at** | **datetime** |  | [optional] 
**created_by** | **str** |  | [optional] 
**deleted_by** | **str** |  | [optional] 
**id** | **int** |  | [optional] 
**link** | **str** |  | [optional] 
**message** | **str** |  | [optional] 
**on_environment** | **str** |  | [optional] 
**severity** | **str** |  | [optional] 
**title** | **str** |  | [optional] 
**updated_at** | **datetime** |  | [optional] 
**updated_by** | **str** |  | [optional] 
**uuid** | **str** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_service_alert_v3 import SherlockServiceAlertV3

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockServiceAlertV3 from a JSON string
sherlock_service_alert_v3_instance = SherlockServiceAlertV3.from_json(json)
# print the JSON string representation of the object
print(SherlockServiceAlertV3.to_json())

# convert the object into a dict
sherlock_service_alert_v3_dict = sherlock_service_alert_v3_instance.to_dict()
# create an instance of SherlockServiceAlertV3 from a dict
sherlock_service_alert_v3_from_dict = SherlockServiceAlertV3.from_dict(sherlock_service_alert_v3_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


