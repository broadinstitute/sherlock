# MiscConnectionCheckResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ok** | **bool** | Always true | [optional] 

## Example

```python
from sherlock_python_client.models.misc_connection_check_response import MiscConnectionCheckResponse

# TODO update the JSON string below
json = "{}"
# create an instance of MiscConnectionCheckResponse from a JSON string
misc_connection_check_response_instance = MiscConnectionCheckResponse.from_json(json)
# print the JSON string representation of the object
print(MiscConnectionCheckResponse.to_json())

# convert the object into a dict
misc_connection_check_response_dict = misc_connection_check_response_instance.to_dict()
# create an instance of MiscConnectionCheckResponse from a dict
misc_connection_check_response_from_dict = MiscConnectionCheckResponse.from_dict(misc_connection_check_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


