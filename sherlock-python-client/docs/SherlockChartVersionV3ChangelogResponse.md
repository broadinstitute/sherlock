# SherlockChartVersionV3ChangelogResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**changelog** | [**List[SherlockChartVersionV3]**](SherlockChartVersionV3.md) |  | [optional] 
**complete** | **bool** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_chart_version_v3_changelog_response import SherlockChartVersionV3ChangelogResponse

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockChartVersionV3ChangelogResponse from a JSON string
sherlock_chart_version_v3_changelog_response_instance = SherlockChartVersionV3ChangelogResponse.from_json(json)
# print the JSON string representation of the object
print(SherlockChartVersionV3ChangelogResponse.to_json())

# convert the object into a dict
sherlock_chart_version_v3_changelog_response_dict = sherlock_chart_version_v3_changelog_response_instance.to_dict()
# create an instance of SherlockChartVersionV3ChangelogResponse from a dict
sherlock_chart_version_v3_changelog_response_from_dict = SherlockChartVersionV3ChangelogResponse.from_dict(sherlock_chart_version_v3_changelog_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


