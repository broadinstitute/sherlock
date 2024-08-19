# SherlockDatabaseInstanceV3Edit


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**default_database** | **str** | When creating, defaults to the chart name | [optional] 
**google_project** | **str** | Required if platform is &#39;google&#39; | [optional] 
**instance_name** | **str** | Required if platform is &#39;google&#39; or &#39;azure&#39; | [optional] 
**platform** | **str** | &#39;google&#39;, &#39;azure&#39;, or default &#39;kubernetes&#39; | [optional] [default to 'kubernetes']

## Example

```python
from sherlock_python_client.models.sherlock_database_instance_v3_edit import SherlockDatabaseInstanceV3Edit

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockDatabaseInstanceV3Edit from a JSON string
sherlock_database_instance_v3_edit_instance = SherlockDatabaseInstanceV3Edit.from_json(json)
# print the JSON string representation of the object
print(SherlockDatabaseInstanceV3Edit.to_json())

# convert the object into a dict
sherlock_database_instance_v3_edit_dict = sherlock_database_instance_v3_edit_instance.to_dict()
# create an instance of SherlockDatabaseInstanceV3Edit from a dict
sherlock_database_instance_v3_edit_from_dict = SherlockDatabaseInstanceV3Edit.from_dict(sherlock_database_instance_v3_edit_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


