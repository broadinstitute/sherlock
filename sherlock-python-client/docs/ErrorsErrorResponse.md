# ErrorsErrorResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message** | **str** |  | [optional] 
**to_blame** | **str** |  | [optional] 
**type** | **str** |  | [optional] 

## Example

```python
from sherlock_python_client.models.errors_error_response import ErrorsErrorResponse

# TODO update the JSON string below
json = "{}"
# create an instance of ErrorsErrorResponse from a JSON string
errors_error_response_instance = ErrorsErrorResponse.from_json(json)
# print the JSON string representation of the object
print(ErrorsErrorResponse.to_json())

# convert the object into a dict
errors_error_response_dict = errors_error_response_instance.to_dict()
# create an instance of ErrorsErrorResponse from a dict
errors_error_response_from_dict = ErrorsErrorResponse.from_dict(errors_error_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


