# SherlockChangesetV3


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**applied_at** | **datetime** |  | [optional] 
**applied_by** | **str** |  | [optional] 
**applied_by_info** | [**SherlockUserV3**](SherlockUserV3.md) |  | [optional] 
**chart_release** | **str** |  | [optional] 
**chart_release_info** | [**SherlockChartReleaseV3**](SherlockChartReleaseV3.md) |  | [optional] 
**ci_identifier** | [**SherlockCiIdentifierV3**](SherlockCiIdentifierV3.md) |  | [optional] 
**created_at** | **datetime** |  | [optional] 
**from_app_version_branch** | **str** |  | [optional] 
**from_app_version_commit** | **str** |  | [optional] 
**from_app_version_exact** | **str** |  | [optional] 
**from_app_version_follow_chart_release** | **str** |  | [optional] 
**from_app_version_reference** | **str** |  | [optional] 
**from_app_version_resolver** | **str** |  | [optional] 
**from_chart_version_exact** | **str** |  | [optional] 
**from_chart_version_follow_chart_release** | **str** |  | [optional] 
**from_chart_version_reference** | **str** |  | [optional] 
**from_chart_version_resolver** | **str** |  | [optional] 
**from_helmfile_ref** | **str** |  | [optional] 
**from_helmfile_ref_enabled** | **bool** |  | [optional] 
**from_resolved_at** | **datetime** |  | [optional] 
**id** | **int** |  | [optional] 
**new_app_versions** | [**List[SherlockAppVersionV3]**](SherlockAppVersionV3.md) |  | [optional] 
**new_chart_versions** | [**List[SherlockChartVersionV3]**](SherlockChartVersionV3.md) |  | [optional] 
**planned_by** | **str** |  | [optional] 
**planned_by_info** | [**SherlockUserV3**](SherlockUserV3.md) |  | [optional] 
**superseded_at** | **datetime** |  | [optional] 
**to_app_version_branch** | **str** |  | [optional] 
**to_app_version_commit** | **str** |  | [optional] 
**to_app_version_exact** | **str** |  | [optional] 
**to_app_version_follow_chart_release** | **str** |  | [optional] 
**to_app_version_reference** | **str** |  | [optional] 
**to_app_version_resolver** | **str** |  | [optional] 
**to_chart_version_exact** | **str** |  | [optional] 
**to_chart_version_follow_chart_release** | **str** |  | [optional] 
**to_chart_version_reference** | **str** |  | [optional] 
**to_chart_version_resolver** | **str** |  | [optional] 
**to_helmfile_ref** | **str** |  | [optional] 
**to_helmfile_ref_enabled** | **bool** |  | [optional] 
**to_resolved_at** | **datetime** |  | [optional] 
**updated_at** | **datetime** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_changeset_v3 import SherlockChangesetV3

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockChangesetV3 from a JSON string
sherlock_changeset_v3_instance = SherlockChangesetV3.from_json(json)
# print the JSON string representation of the object
print(SherlockChangesetV3.to_json())

# convert the object into a dict
sherlock_changeset_v3_dict = sherlock_changeset_v3_instance.to_dict()
# create an instance of SherlockChangesetV3 from a dict
sherlock_changeset_v3_from_dict = SherlockChangesetV3.from_dict(sherlock_changeset_v3_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


