# SherlockChartVersionV3


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**authored_by** | **str** |  | [optional] 
**authored_by_info** | [**SherlockUserV3**](SherlockUserV3.md) |  | [optional] 
**chart** | **str** | Required when creating | [optional] 
**chart_info** | [**SherlockChartV3**](SherlockChartV3.md) |  | [optional] 
**chart_version** | **str** | Required when creating | [optional] 
**ci_identifier** | [**SherlockCiIdentifierV3**](SherlockCiIdentifierV3.md) |  | [optional] 
**created_at** | **datetime** |  | [optional] 
**description** | **str** | Generally the Git commit message | [optional] 
**id** | **int** |  | [optional] 
**parent_chart_version** | **str** |  | [optional] 
**parent_chart_version_info** | **object** |  | [optional] 
**updated_at** | **datetime** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_chart_version_v3 import SherlockChartVersionV3

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockChartVersionV3 from a JSON string
sherlock_chart_version_v3_instance = SherlockChartVersionV3.from_json(json)
# print the JSON string representation of the object
print(SherlockChartVersionV3.to_json())

# convert the object into a dict
sherlock_chart_version_v3_dict = sherlock_chart_version_v3_instance.to_dict()
# create an instance of SherlockChartVersionV3 from a dict
sherlock_chart_version_v3_from_dict = SherlockChartVersionV3.from_dict(sherlock_chart_version_v3_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


