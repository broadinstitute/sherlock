# SherlockChartReleaseV3Edit


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**included_in_bulk_changesets** | **bool** |  | [optional] [default to True]
**pagerduty_integration** | **str** |  | [optional] 
**port** | **int** | When creating, will use the chart&#39;s default if left empty | [optional] 
**protocol** | **str** | When creating, will use the chart&#39;s default if left empty | [optional] 
**subdomain** | **str** | When creating, will use the chart&#39;s default if left empty | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_chart_release_v3_edit import SherlockChartReleaseV3Edit

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockChartReleaseV3Edit from a JSON string
sherlock_chart_release_v3_edit_instance = SherlockChartReleaseV3Edit.from_json(json)
# print the JSON string representation of the object
print(SherlockChartReleaseV3Edit.to_json())

# convert the object into a dict
sherlock_chart_release_v3_edit_dict = sherlock_chart_release_v3_edit_instance.to_dict()
# create an instance of SherlockChartReleaseV3Edit from a dict
sherlock_chart_release_v3_edit_from_dict = SherlockChartReleaseV3Edit.from_dict(sherlock_chart_release_v3_edit_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


