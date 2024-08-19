# SherlockChartV3Create


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**app_image_git_main_branch** | **str** |  | [optional] 
**app_image_git_repo** | **str** |  | [optional] 
**chart_exposes_endpoint** | **bool** | Indicates if the default subdomain, protocol, and port fields are relevant for this chart | [optional] [default to False]
**chart_repo** | **str** |  | [optional] [default to 'terra-helm']
**default_port** | **int** |  | [optional] 
**default_protocol** | **str** |  | [optional] [default to 'https']
**default_subdomain** | **str** | When creating, will default to the name of the chart | [optional] 
**description** | **str** |  | [optional] 
**name** | **str** | Required when creating | [optional] 
**pact_participant** | **bool** |  | [optional] [default to False]
**playbook_url** | **str** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_chart_v3_create import SherlockChartV3Create

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockChartV3Create from a JSON string
sherlock_chart_v3_create_instance = SherlockChartV3Create.from_json(json)
# print the JSON string representation of the object
print(SherlockChartV3Create.to_json())

# convert the object into a dict
sherlock_chart_v3_create_dict = sherlock_chart_v3_create_instance.to_dict()
# create an instance of SherlockChartV3Create from a dict
sherlock_chart_v3_create_from_dict = SherlockChartV3Create.from_dict(sherlock_chart_v3_create_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


