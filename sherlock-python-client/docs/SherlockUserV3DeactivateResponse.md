# SherlockUserV3DeactivateResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**already_deactivated_emails** | **List[str]** |  | [optional] 
**newly_deactivated_emails** | **List[str]** |  | [optional] 
**not_found_emails** | **List[str]** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_user_v3_deactivate_response import SherlockUserV3DeactivateResponse

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockUserV3DeactivateResponse from a JSON string
sherlock_user_v3_deactivate_response_instance = SherlockUserV3DeactivateResponse.from_json(json)
# print the JSON string representation of the object
print(SherlockUserV3DeactivateResponse.to_json())

# convert the object into a dict
sherlock_user_v3_deactivate_response_dict = sherlock_user_v3_deactivate_response_instance.to_dict()
# create an instance of SherlockUserV3DeactivateResponse from a dict
sherlock_user_v3_deactivate_response_from_dict = SherlockUserV3DeactivateResponse.from_dict(sherlock_user_v3_deactivate_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


