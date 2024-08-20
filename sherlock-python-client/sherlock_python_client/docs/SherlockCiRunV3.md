# SherlockCiRunV3


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**argo_workflows_name** | **str** |  | [optional] 
**argo_workflows_namespace** | **str** |  | [optional] 
**argo_workflows_template** | **str** |  | [optional] 
**created_at** | **datetime** |  | [optional] 
**github_actions_attempt_number** | **int** |  | [optional] 
**github_actions_owner** | **str** |  | [optional] 
**github_actions_repo** | **str** |  | [optional] 
**github_actions_run_id** | **int** |  | [optional] 
**github_actions_workflow_path** | **str** |  | [optional] 
**id** | **int** |  | [optional] 
**notify_slack_channels_upon_failure** | **List[str]** | Slack channels to notify if this CiRun fails. This field is always appended to when mutated. | [optional] 
**notify_slack_channels_upon_retry** | **List[str]** | Slack channels to notify if this CiRun is retried. This field is always appended to when mutated. It will de-dupe with the other notify fields. | [optional] 
**notify_slack_channels_upon_success** | **List[str]** | Slack channels to notify if this CiRun succeeds. This field is always appended to when mutated. | [optional] 
**notify_slack_custom_icon** | **str** | Icon to use for success or failure Slack notifications. Can be given either as a URL to an image or as a Slack emoji (using colon shortcodes, like :smiley:). An empty string is ignored to facilitate calling from GitHub Actions (where it&#39;s easier to pass an empty string than not send the field at all). | [optional] 
**platform** | **str** |  | [optional] 
**related_resources** | [**List[SherlockCiIdentifierV3]**](SherlockCiIdentifierV3.md) |  | [optional] 
**resource_status** | **str** | Available only when querying a CiRun via a CiIdentifier, indicates the status of the run for that resource | [optional] 
**started_at** | **str** |  | [optional] 
**status** | **str** |  | [optional] 
**terminal_at** | **str** |  | [optional] 
**termination_hooks_dispatched_at** | **datetime** |  | [optional] 
**updated_at** | **datetime** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_ci_run_v3 import SherlockCiRunV3

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockCiRunV3 from a JSON string
sherlock_ci_run_v3_instance = SherlockCiRunV3.from_json(json)
# print the JSON string representation of the object
print(SherlockCiRunV3.to_json())

# convert the object into a dict
sherlock_ci_run_v3_dict = sherlock_ci_run_v3_instance.to_dict()
# create an instance of SherlockCiRunV3 from a dict
sherlock_ci_run_v3_from_dict = SherlockCiRunV3.from_dict(sherlock_ci_run_v3_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


