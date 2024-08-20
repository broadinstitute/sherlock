# SherlockChartVersionV3Create


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**chart** | **str** | Required when creating | [optional] 
**chart_version** | **str** | Required when creating | [optional] 
**description** | **str** | Generally the Git commit message | [optional] 
**parent_chart_version** | **str** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_chart_version_v3_create import SherlockChartVersionV3Create

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockChartVersionV3Create from a JSON string
sherlock_chart_version_v3_create_instance = SherlockChartVersionV3Create.from_json(json)
# print the JSON string representation of the object
print(SherlockChartVersionV3Create.to_json())

# convert the object into a dict
sherlock_chart_version_v3_create_dict = sherlock_chart_version_v3_create_instance.to_dict()
# create an instance of SherlockChartVersionV3Create from a dict
sherlock_chart_version_v3_create_from_dict = SherlockChartVersionV3Create.from_dict(sherlock_chart_version_v3_create_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


