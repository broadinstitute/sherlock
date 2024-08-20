# SherlockIncidentV3Create


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
from sherlock_python_client.models.sherlock_incident_v3_create import SherlockIncidentV3Create

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockIncidentV3Create from a JSON string
sherlock_incident_v3_create_instance = SherlockIncidentV3Create.from_json(json)
# print the JSON string representation of the object
print(SherlockIncidentV3Create.to_json())

# convert the object into a dict
sherlock_incident_v3_create_dict = sherlock_incident_v3_create_instance.to_dict()
# create an instance of SherlockIncidentV3Create from a dict
sherlock_incident_v3_create_from_dict = SherlockIncidentV3Create.from_dict(sherlock_incident_v3_create_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


