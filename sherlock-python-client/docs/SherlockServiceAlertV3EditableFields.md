# SherlockServiceAlertV3EditableFields


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**link** | **str** |  | [optional] 
**message** | **str** |  | [optional] 
**severity** | **str** |  | [optional] 
**title** | **str** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_service_alert_v3_editable_fields import SherlockServiceAlertV3EditableFields

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockServiceAlertV3EditableFields from a JSON string
sherlock_service_alert_v3_editable_fields_instance = SherlockServiceAlertV3EditableFields.from_json(json)
# print the JSON string representation of the object
print(SherlockServiceAlertV3EditableFields.to_json())

# convert the object into a dict
sherlock_service_alert_v3_editable_fields_dict = sherlock_service_alert_v3_editable_fields_instance.to_dict()
# create an instance of SherlockServiceAlertV3EditableFields from a dict
sherlock_service_alert_v3_editable_fields_from_dict = SherlockServiceAlertV3EditableFields.from_dict(sherlock_service_alert_v3_editable_fields_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


