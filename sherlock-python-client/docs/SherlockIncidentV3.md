# SherlockIncidentV3


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**created_at** | **datetime** |  | [optional] 
**description** | **str** |  | [optional] 
**id** | **int** |  | [optional] 
**remediated_at** | **str** |  | [optional] 
**review_completed_at** | **str** |  | [optional] 
**started_at** | **str** |  | [optional] 
**ticket** | **str** |  | [optional] 
**updated_at** | **datetime** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_incident_v3 import SherlockIncidentV3

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockIncidentV3 from a JSON string
sherlock_incident_v3_instance = SherlockIncidentV3.from_json(json)
# print the JSON string representation of the object
print(SherlockIncidentV3.to_json())

# convert the object into a dict
sherlock_incident_v3_dict = sherlock_incident_v3_instance.to_dict()
# create an instance of SherlockIncidentV3 from a dict
sherlock_incident_v3_from_dict = SherlockIncidentV3.from_dict(sherlock_incident_v3_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


