# SherlockDatabaseInstanceV3


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**chart_release** | **str** | Required when creating | [optional] 
**chart_release_info** | [**SherlockChartReleaseV3**](SherlockChartReleaseV3.md) |  | [optional] 
**created_at** | **datetime** |  | [optional] 
**default_database** | **str** | When creating, defaults to the chart name | [optional] 
**google_project** | **str** | Required if platform is &#39;google&#39; | [optional] 
**id** | **int** |  | [optional] 
**instance_name** | **str** | Required if platform is &#39;google&#39; or &#39;azure&#39; | [optional] 
**platform** | **str** | &#39;google&#39;, &#39;azure&#39;, or default &#39;kubernetes&#39; | [optional] [default to 'kubernetes']
**updated_at** | **datetime** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_database_instance_v3 import SherlockDatabaseInstanceV3

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockDatabaseInstanceV3 from a JSON string
sherlock_database_instance_v3_instance = SherlockDatabaseInstanceV3.from_json(json)
# print the JSON string representation of the object
print(SherlockDatabaseInstanceV3.to_json())

# convert the object into a dict
sherlock_database_instance_v3_dict = sherlock_database_instance_v3_instance.to_dict()
# create an instance of SherlockDatabaseInstanceV3 from a dict
sherlock_database_instance_v3_from_dict = SherlockDatabaseInstanceV3.from_dict(sherlock_database_instance_v3_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


