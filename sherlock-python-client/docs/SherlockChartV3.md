# SherlockChartV3


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**app_image_git_main_branch** | **str** |  | [optional] 
**app_image_git_repo** | **str** |  | [optional] 
**chart_exposes_endpoint** | **bool** | Indicates if the default subdomain, protocol, and port fields are relevant for this chart | [optional] [default to False]
**chart_repo** | **str** |  | [optional] [default to 'terra-helm']
**ci_identifier** | [**SherlockCiIdentifierV3**](SherlockCiIdentifierV3.md) |  | [optional] 
**created_at** | **datetime** |  | [optional] 
**default_port** | **int** |  | [optional] 
**default_protocol** | **str** |  | [optional] [default to 'https']
**default_subdomain** | **str** | When creating, will default to the name of the chart | [optional] 
**description** | **str** |  | [optional] 
**id** | **int** |  | [optional] 
**name** | **str** | Required when creating | [optional] 
**pact_participant** | **bool** |  | [optional] [default to False]
**playbook_url** | **str** |  | [optional] 
**updated_at** | **datetime** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_chart_v3 import SherlockChartV3

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockChartV3 from a JSON string
sherlock_chart_v3_instance = SherlockChartV3.from_json(json)
# print the JSON string representation of the object
print(SherlockChartV3.to_json())

# convert the object into a dict
sherlock_chart_v3_dict = sherlock_chart_v3_instance.to_dict()
# create an instance of SherlockChartV3 from a dict
sherlock_chart_v3_from_dict = SherlockChartV3.from_dict(sherlock_chart_v3_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


