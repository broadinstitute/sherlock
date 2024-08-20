# SherlockChangesetV3PlanRequestEnvironmentEntry


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**environment** | **str** |  | [optional] 
**exclude_charts** | **List[str]** |  | [optional] 
**filter_to_matching_branches** | **bool** | If true, chart releases app versions will only be updated if doing so wouldn&#39;t change the detected Git branch. This flag has no effect if the updated chart release has no app version branch. | [optional] 
**follow_versions_from_other_environment** | **str** |  | [optional] 
**include_charts** | **List[str]** | If omitted, will include all chart releases that haven&#39;t opted out of bulk updates | [optional] 
**use_exact_versions_from_other_environment** | **str** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_changeset_v3_plan_request_environment_entry import SherlockChangesetV3PlanRequestEnvironmentEntry

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockChangesetV3PlanRequestEnvironmentEntry from a JSON string
sherlock_changeset_v3_plan_request_environment_entry_instance = SherlockChangesetV3PlanRequestEnvironmentEntry.from_json(json)
# print the JSON string representation of the object
print(SherlockChangesetV3PlanRequestEnvironmentEntry.to_json())

# convert the object into a dict
sherlock_changeset_v3_plan_request_environment_entry_dict = sherlock_changeset_v3_plan_request_environment_entry_instance.to_dict()
# create an instance of SherlockChangesetV3PlanRequestEnvironmentEntry from a dict
sherlock_changeset_v3_plan_request_environment_entry_from_dict = SherlockChangesetV3PlanRequestEnvironmentEntry.from_dict(sherlock_changeset_v3_plan_request_environment_entry_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


