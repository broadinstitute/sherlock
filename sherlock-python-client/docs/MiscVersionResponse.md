# MiscVersionResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**build_info** | **Dict[str, str]** |  | [optional] 
**go_version** | **str** |  | [optional] 
**version** | **str** |  | [optional] 

## Example

```python
from sherlock_python_client.models.misc_version_response import MiscVersionResponse

# TODO update the JSON string below
json = "{}"
# create an instance of MiscVersionResponse from a JSON string
misc_version_response_instance = MiscVersionResponse.from_json(json)
# print the JSON string representation of the object
print(MiscVersionResponse.to_json())

# convert the object into a dict
misc_version_response_dict = misc_version_response_instance.to_dict()
# create an instance of MiscVersionResponse from a dict
misc_version_response_from_dict = MiscVersionResponse.from_dict(misc_version_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


