# SherlockAppVersionV3


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**app_version** | **str** | Required when creating | [optional] 
**authored_by** | **str** |  | [optional] 
**authored_by_info** | [**SherlockUserV3**](SherlockUserV3.md) |  | [optional] 
**chart** | **str** | Required when creating | [optional] 
**chart_info** | [**SherlockChartV3**](SherlockChartV3.md) |  | [optional] 
**ci_identifier** | [**SherlockCiIdentifierV3**](SherlockCiIdentifierV3.md) |  | [optional] 
**created_at** | **datetime** |  | [optional] 
**description** | **str** | Generally the Git commit message | [optional] 
**git_branch** | **str** |  | [optional] 
**git_commit** | **str** |  | [optional] 
**id** | **int** |  | [optional] 
**parent_app_version** | **str** |  | [optional] 
**parent_app_version_info** | **object** |  | [optional] 
**updated_at** | **datetime** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_app_version_v3 import SherlockAppVersionV3

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockAppVersionV3 from a JSON string
sherlock_app_version_v3_instance = SherlockAppVersionV3.from_json(json)
# print the JSON string representation of the object
print(SherlockAppVersionV3.to_json())

# convert the object into a dict
sherlock_app_version_v3_dict = sherlock_app_version_v3_instance.to_dict()
# create an instance of SherlockAppVersionV3 from a dict
sherlock_app_version_v3_from_dict = SherlockAppVersionV3.from_dict(sherlock_app_version_v3_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


