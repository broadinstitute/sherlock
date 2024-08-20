# SherlockIncidentV3Edit


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**description** | **str** |  | [optional] 
**remediated_at** | **str** |  | [optional] 
**review_completed_at** | **str** |  | [optional] 
**started_at** | **str** |  | [optional] 
**ticket** | **str** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_incident_v3_edit import SherlockIncidentV3Edit

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockIncidentV3Edit from a JSON string
sherlock_incident_v3_edit_instance = SherlockIncidentV3Edit.from_json(json)
# print the JSON string representation of the object
print(SherlockIncidentV3Edit.to_json())

# convert the object into a dict
sherlock_incident_v3_edit_dict = sherlock_incident_v3_edit_instance.to_dict()
# create an instance of SherlockIncidentV3Edit from a dict
sherlock_incident_v3_edit_from_dict = SherlockIncidentV3Edit.from_dict(sherlock_incident_v3_edit_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


