# SherlockChartReleaseV3


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**app_version_branch** | **str** | When creating, will default to the app&#39;s mainline branch if no other app version info is present | [optional] 
**app_version_commit** | **str** |  | [optional] 
**app_version_exact** | **str** |  | [optional] 
**app_version_follow_chart_release** | **str** |  | [optional] 
**app_version_info** | [**SherlockAppVersionV3**](SherlockAppVersionV3.md) |  | [optional] 
**app_version_reference** | **str** |  | [optional] 
**app_version_resolver** | **str** | // When creating, will default to automatically reference any provided app version fields | [optional] 
**chart** | **str** | Required when creating | [optional] 
**chart_info** | [**SherlockChartV3**](SherlockChartV3.md) |  | [optional] 
**chart_version_exact** | **str** |  | [optional] 
**chart_version_follow_chart_release** | **str** |  | [optional] 
**chart_version_info** | [**SherlockChartVersionV3**](SherlockChartVersionV3.md) |  | [optional] 
**chart_version_reference** | **str** |  | [optional] 
**chart_version_resolver** | **str** | When creating, will default to automatically reference any provided chart version | [optional] 
**ci_identifier** | [**SherlockCiIdentifierV3**](SherlockCiIdentifierV3.md) |  | [optional] 
**cluster** | **str** | When creating, will default the environment&#39;s default cluster, if provided. Either this or environment must be provided. | [optional] 
**cluster_info** | [**SherlockClusterV3**](SherlockClusterV3.md) |  | [optional] 
**created_at** | **datetime** |  | [optional] 
**destination_type** | **str** | Calculated field | [optional] 
**environment** | **str** | Either this or cluster must be provided. | [optional] 
**environment_info** | [**SherlockEnvironmentV3**](SherlockEnvironmentV3.md) |  | [optional] 
**helmfile_ref** | **str** |  | [optional] [default to 'HEAD']
**helmfile_ref_enabled** | **bool** |  | [optional] [default to False]
**id** | **int** |  | [optional] 
**included_in_bulk_changesets** | **bool** |  | [optional] [default to True]
**name** | **str** | When creating, will be calculated if left empty | [optional] 
**namespace** | **str** | When creating, will default to the environment&#39;s default namespace, if provided | [optional] 
**pagerduty_integration** | **str** |  | [optional] 
**pagerduty_integration_info** | [**SherlockPagerdutyIntegrationV3**](SherlockPagerdutyIntegrationV3.md) |  | [optional] 
**port** | **int** | When creating, will use the chart&#39;s default if left empty | [optional] 
**protocol** | **str** | When creating, will use the chart&#39;s default if left empty | [optional] 
**resolved_at** | **datetime** |  | [optional] 
**subdomain** | **str** | When creating, will use the chart&#39;s default if left empty | [optional] 
**updated_at** | **datetime** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_chart_release_v3 import SherlockChartReleaseV3

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockChartReleaseV3 from a JSON string
sherlock_chart_release_v3_instance = SherlockChartReleaseV3.from_json(json)
# print the JSON string representation of the object
print(SherlockChartReleaseV3.to_json())

# convert the object into a dict
sherlock_chart_release_v3_dict = sherlock_chart_release_v3_instance.to_dict()
# create an instance of SherlockChartReleaseV3 from a dict
sherlock_chart_release_v3_from_dict = SherlockChartReleaseV3.from_dict(sherlock_chart_release_v3_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


