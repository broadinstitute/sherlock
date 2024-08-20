# SherlockAppVersionV3ChangelogResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**changelog** | [**List[SherlockAppVersionV3]**](SherlockAppVersionV3.md) |  | [optional] 
**complete** | **bool** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_app_version_v3_changelog_response import SherlockAppVersionV3ChangelogResponse

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockAppVersionV3ChangelogResponse from a JSON string
sherlock_app_version_v3_changelog_response_instance = SherlockAppVersionV3ChangelogResponse.from_json(json)
# print the JSON string representation of the object
print(SherlockAppVersionV3ChangelogResponse.to_json())

# convert the object into a dict
sherlock_app_version_v3_changelog_response_dict = sherlock_app_version_v3_changelog_response_instance.to_dict()
# create an instance of SherlockAppVersionV3ChangelogResponse from a dict
sherlock_app_version_v3_changelog_response_from_dict = SherlockAppVersionV3ChangelogResponse.from_dict(sherlock_app_version_v3_changelog_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


