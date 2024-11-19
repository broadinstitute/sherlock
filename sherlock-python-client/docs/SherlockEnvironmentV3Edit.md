# SherlockEnvironmentV3Edit


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**base_domain** | **str** |  | [optional] [default to 'bee.envs-terra.bio']
**default_cluster** | **str** |  | [optional] 
**delete_after** | **datetime** | If set, the BEE will be automatically deleted after this time. Can be set to \&quot;\&quot; or Go&#39;s zero time value to clear the field. | [optional] 
**description** | **str** |  | [optional] 
**enable_janitor** | **bool** | If true, janitor resource cleanup will be enabled for this environment. BEEs default to template&#39;s value, templates default to true, and static/live environments default to false. | [optional] 
**helmfile_ref** | **str** |  | [optional] [default to 'HEAD']
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

## Example

```python
from sherlock_python_client.models.sherlock_environment_v3_edit import SherlockEnvironmentV3Edit

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockEnvironmentV3Edit from a JSON string
sherlock_environment_v3_edit_instance = SherlockEnvironmentV3Edit.from_json(json)
# print the JSON string representation of the object
print(SherlockEnvironmentV3Edit.to_json())

# convert the object into a dict
sherlock_environment_v3_edit_dict = sherlock_environment_v3_edit_instance.to_dict()
# create an instance of SherlockEnvironmentV3Edit from a dict
sherlock_environment_v3_edit_from_dict = SherlockEnvironmentV3Edit.from_dict(sherlock_environment_v3_edit_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


