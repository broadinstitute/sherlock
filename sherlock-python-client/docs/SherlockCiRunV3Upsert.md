# SherlockCiRunV3Upsert


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**app_versions** | **List[str]** | Always appends; will eliminate duplicates. | [optional] 
**argo_workflows_name** | **str** |  | [optional] 
**argo_workflows_namespace** | **str** |  | [optional] 
**argo_workflows_template** | **str** |  | [optional] 
**changesets** | **List[str]** | Always appends; will eliminate duplicates. Spreads to associated chart releases, environments, and clusters. | [optional] 
**chart_release_statuses** | **Dict[str, str]** | Keys treated like chartReleases. Values set resource-specific statuses for chart releases and associated changesets, new app versions, and new chart versions. | [optional] 
**chart_releases** | **List[str]** | Always appends; will eliminate duplicates. Spreads to associated environments and clusters. | [optional] 
**chart_versions** | **List[str]** | Always appends; will eliminate duplicates. | [optional] 
**charts** | **List[str]** | Always appends; will eliminate duplicates. | [optional] 
**clusters** | **List[str]** | Always appends; will eliminate duplicates. Spreads to contained chart releases and their environments. | [optional] 
**environments** | **List[str]** | Always appends; will eliminate duplicates. Spreads to contained chart releases and their clusters. | [optional] 
**github_actions_attempt_number** | **int** |  | [optional] 
**github_actions_owner** | **str** |  | [optional] 
**github_actions_repo** | **str** |  | [optional] 
**github_actions_run_id** | **int** |  | [optional] 
**github_actions_workflow_path** | **str** |  | [optional] 
**ignore_bad_selectors** | **bool** | If set to true, errors handling selectors for relations should be ignored. Normally, passing an unknown chart, cluster, etc. will abort the request, but they won&#39;t if this is true. | [optional] [default to False]
**notify_slack_channels_upon_failure** | **List[str]** | Slack channels to notify if this CiRun fails. This field is always appended to when mutated. | [optional] 
**notify_slack_channels_upon_retry** | **List[str]** | Slack channels to notify if this CiRun is retried. This field is always appended to when mutated. It will de-dupe with the other notify fields. | [optional] 
**notify_slack_channels_upon_success** | **List[str]** | Slack channels to notify if this CiRun succeeds. This field is always appended to when mutated. | [optional] 
**notify_slack_custom_icon** | **str** | Icon to use for success or failure Slack notifications. Can be given either as a URL to an image or as a Slack emoji (using colon shortcodes, like :smiley:). An empty string is ignored to facilitate calling from GitHub Actions (where it&#39;s easier to pass an empty string than not send the field at all). | [optional] 
**platform** | **str** |  | [optional] 
**relate_to_changeset_new_versions** | **str** | Makes entries in the changesets field also spread to new app versions and chart versions deployed by the changeset. &#39;when-static&#39; is the default and does this spreading only when the chart release is in a static environment. | [optional] [default to 'when-static']
**started_at** | **str** |  | [optional] 
**status** | **str** |  | [optional] 
**terminal_at** | **str** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_ci_run_v3_upsert import SherlockCiRunV3Upsert

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockCiRunV3Upsert from a JSON string
sherlock_ci_run_v3_upsert_instance = SherlockCiRunV3Upsert.from_json(json)
# print the JSON string representation of the object
print(SherlockCiRunV3Upsert.to_json())

# convert the object into a dict
sherlock_ci_run_v3_upsert_dict = sherlock_ci_run_v3_upsert_instance.to_dict()
# create an instance of SherlockCiRunV3Upsert from a dict
sherlock_ci_run_v3_upsert_from_dict = SherlockCiRunV3Upsert.from_dict(sherlock_ci_run_v3_upsert_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


