# SherlockChangesetV3PlanRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**chart_releases** | [**List[SherlockChangesetV3PlanRequestChartReleaseEntry]**](SherlockChangesetV3PlanRequestChartReleaseEntry.md) |  | [optional] 
**environments** | [**List[SherlockChangesetV3PlanRequestEnvironmentEntry]**](SherlockChangesetV3PlanRequestEnvironmentEntry.md) |  | [optional] 
**recreate_changesets** | **List[int]** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_changeset_v3_plan_request import SherlockChangesetV3PlanRequest

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockChangesetV3PlanRequest from a JSON string
sherlock_changeset_v3_plan_request_instance = SherlockChangesetV3PlanRequest.from_json(json)
# print the JSON string representation of the object
print(SherlockChangesetV3PlanRequest.to_json())

# convert the object into a dict
sherlock_changeset_v3_plan_request_dict = sherlock_changeset_v3_plan_request_instance.to_dict()
# create an instance of SherlockChangesetV3PlanRequest from a dict
sherlock_changeset_v3_plan_request_from_dict = SherlockChangesetV3PlanRequest.from_dict(sherlock_changeset_v3_plan_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


