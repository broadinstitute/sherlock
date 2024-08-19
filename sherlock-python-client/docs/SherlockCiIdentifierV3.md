# SherlockCiIdentifierV3


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ci_runs** | [**List[SherlockCiRunV3]**](SherlockCiRunV3.md) |  | [optional] 
**created_at** | **datetime** |  | [optional] 
**id** | **int** |  | [optional] 
**resource_id** | **int** |  | [optional] 
**resource_status** | **str** | Available only when querying a CiIdentifier via a CiRun, indicates the status of the run for that resource | [optional] 
**resource_type** | **str** |  | [optional] 
**updated_at** | **datetime** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_ci_identifier_v3 import SherlockCiIdentifierV3

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockCiIdentifierV3 from a JSON string
sherlock_ci_identifier_v3_instance = SherlockCiIdentifierV3.from_json(json)
# print the JSON string representation of the object
print(SherlockCiIdentifierV3.to_json())

# convert the object into a dict
sherlock_ci_identifier_v3_dict = sherlock_ci_identifier_v3_instance.to_dict()
# create an instance of SherlockCiIdentifierV3 from a dict
sherlock_ci_identifier_v3_from_dict = SherlockCiIdentifierV3.from_dict(sherlock_ci_identifier_v3_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


