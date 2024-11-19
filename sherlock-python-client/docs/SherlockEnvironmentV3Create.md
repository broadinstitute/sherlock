# SherlockEnvironmentV3Create


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**auto_populate_chart_releases** | **bool** | If true when creating, dynamic environments copy from template and template environments get the honeycomb chart | [optional] [default to True]
**base** | **str** | Required when creating | [optional] 
**base_domain** | **str** |  | [optional] [default to 'bee.envs-terra.bio']
**default_cluster** | **str** |  | [optional] 
**default_namespace** | **str** | When creating, will be calculated if left empty | [optional] 
**delete_after** | **datetime** | If set, the BEE will be automatically deleted after this time. Can be set to \&quot;\&quot; or Go&#39;s zero time value to clear the field. | [optional] 
**description** | **str** |  | [optional] 
**enable_janitor** | **bool** | If true, janitor resource cleanup will be enabled for this environment. BEEs default to template&#39;s value, templates default to true, and static/live environments default to false. | [optional] 
**helmfile_ref** | **str** |  | [optional] [default to 'HEAD']
**lifecycle** | **str** |  | [optional] [default to 'dynamic']
**name** | **str** | When creating, will be calculated if dynamic, required otherwise | [optional] 
**name_prefixes_domain** | **bool** |  | [optional] [default to True]
**offline** | **bool** | Applicable for BEEs only, whether Thelma should render the BEE as \&quot;offline\&quot; zero replicas (this field is a target state, not a status) | [optional] [default to False]
**offline_schedule_begin_enabled** | **bool** | When enabled, the BEE will be slated to go offline around the begin time each day | [optional] 
**offline_schedule_begin_time** | **datetime** | Stored with timezone to determine day of the week | [optional] 
**offline_schedule_end_enabled** | **bool** | When enabled, the BEE will be slated to come online around the end time each weekday (each day if weekends enabled) | [optional] 
**offline_schedule_end_time** | **datetime** | Stored with timezone to determine day of the week | [optional] 
**offline_schedule_end_weekends** | **bool** |  | [optional] 
**owner** | **str** | When creating, will default to you | [optional] 
**pact_identifier** | **str** |  | [optional] 
**pagerduty_integration** | **str** |  | [optional] 
**prevent_deletion** | **bool** | Used to protect specific BEEs from deletion (thelma checks this field) | [optional] [default to False]
**required_role** | **str** | If present, requires membership in the given role for mutations. Set to an empty string to clear. | [optional] 
**requires_suitability** | **bool** |  | [optional] 
**service_banner_bucket** | **str** |  | [optional] 
**template_environment** | **str** | Required for dynamic environments | [optional] 
**unique_resource_prefix** | **str** | When creating, will be calculated if left empty | [optional] 
**values_name** | **str** | When creating, defaults to template name or environment name | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_environment_v3_create import SherlockEnvironmentV3Create

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockEnvironmentV3Create from a JSON string
sherlock_environment_v3_create_instance = SherlockEnvironmentV3Create.from_json(json)
# print the JSON string representation of the object
print(SherlockEnvironmentV3Create.to_json())

# convert the object into a dict
sherlock_environment_v3_create_dict = sherlock_environment_v3_create_instance.to_dict()
# create an instance of SherlockEnvironmentV3Create from a dict
sherlock_environment_v3_create_from_dict = SherlockEnvironmentV3Create.from_dict(sherlock_environment_v3_create_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


