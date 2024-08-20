# SherlockAppVersionV3Create


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**app_version** | **str** | Required when creating | [optional] 
**chart** | **str** | Required when creating | [optional] 
**description** | **str** | Generally the Git commit message | [optional] 
**git_branch** | **str** |  | [optional] 
**git_commit** | **str** |  | [optional] 
**parent_app_version** | **str** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_app_version_v3_create import SherlockAppVersionV3Create

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockAppVersionV3Create from a JSON string
sherlock_app_version_v3_create_instance = SherlockAppVersionV3Create.from_json(json)
# print the JSON string representation of the object
print(SherlockAppVersionV3Create.to_json())

# convert the object into a dict
sherlock_app_version_v3_create_dict = sherlock_app_version_v3_create_instance.to_dict()
# create an instance of SherlockAppVersionV3Create from a dict
sherlock_app_version_v3_create_from_dict = SherlockAppVersionV3Create.from_dict(sherlock_app_version_v3_create_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


