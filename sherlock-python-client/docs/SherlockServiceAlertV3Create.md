# SherlockServiceAlertV3Create


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**link** | **str** |  | [optional] 
**message** | **str** |  | [optional] 
**on_environment** | **str** |  | [optional] 
**severity** | **str** |  | [optional] 
**title** | **str** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_service_alert_v3_create import SherlockServiceAlertV3Create

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockServiceAlertV3Create from a JSON string
sherlock_service_alert_v3_create_instance = SherlockServiceAlertV3Create.from_json(json)
# print the JSON string representation of the object
print(SherlockServiceAlertV3Create.to_json())

# convert the object into a dict
sherlock_service_alert_v3_create_dict = sherlock_service_alert_v3_create_instance.to_dict()
# create an instance of SherlockServiceAlertV3Create from a dict
sherlock_service_alert_v3_create_from_dict = SherlockServiceAlertV3Create.from_dict(sherlock_service_alert_v3_create_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


