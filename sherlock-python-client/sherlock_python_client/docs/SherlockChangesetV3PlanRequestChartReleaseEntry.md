# SherlockChangesetV3PlanRequestChartReleaseEntry


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**chart_release** | **str** |  | [optional] 
**follow_versions_from_other_chart_release** | **str** |  | [optional] 
**to_app_version_branch** | **str** |  | [optional] 
**to_app_version_commit** | **str** |  | [optional] 
**to_app_version_exact** | **str** |  | [optional] 
**to_app_version_follow_chart_release** | **str** |  | [optional] 
**to_app_version_resolver** | **str** |  | [optional] 
**to_chart_version_exact** | **str** |  | [optional] 
**to_chart_version_follow_chart_release** | **str** |  | [optional] 
**to_chart_version_resolver** | **str** |  | [optional] 
**to_helmfile_ref** | **str** |  | [optional] 
**to_helmfile_ref_enabled** | **bool** |  | [optional] 
**use_exact_versions_from_other_chart_release** | **str** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_changeset_v3_plan_request_chart_release_entry import SherlockChangesetV3PlanRequestChartReleaseEntry

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockChangesetV3PlanRequestChartReleaseEntry from a JSON string
sherlock_changeset_v3_plan_request_chart_release_entry_instance = SherlockChangesetV3PlanRequestChartReleaseEntry.from_json(json)
# print the JSON string representation of the object
print(SherlockChangesetV3PlanRequestChartReleaseEntry.to_json())

# convert the object into a dict
sherlock_changeset_v3_plan_request_chart_release_entry_dict = sherlock_changeset_v3_plan_request_chart_release_entry_instance.to_dict()
# create an instance of SherlockChangesetV3PlanRequestChartReleaseEntry from a dict
sherlock_changeset_v3_plan_request_chart_release_entry_from_dict = SherlockChangesetV3PlanRequestChartReleaseEntry.from_dict(sherlock_changeset_v3_plan_request_chart_release_entry_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


