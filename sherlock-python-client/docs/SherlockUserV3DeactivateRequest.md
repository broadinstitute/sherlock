# SherlockUserV3DeactivateRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**suspend_email_handles_across_google_workspace_domains** | **List[str]** |  | [optional] 
**user_email_home_domain** | **str** | Domain of UserEmails that can be swapped out for the domains in SuspendEmailHandlesAcrossGoogleWorkspaceDomains | [optional] [default to 'broadinstitute.org']
**user_emails** | **List[str]** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_user_v3_deactivate_request import SherlockUserV3DeactivateRequest

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockUserV3DeactivateRequest from a JSON string
sherlock_user_v3_deactivate_request_instance = SherlockUserV3DeactivateRequest.from_json(json)
# print the JSON string representation of the object
print(SherlockUserV3DeactivateRequest.to_json())

# convert the object into a dict
sherlock_user_v3_deactivate_request_dict = sherlock_user_v3_deactivate_request_instance.to_dict()
# create an instance of SherlockUserV3DeactivateRequest from a dict
sherlock_user_v3_deactivate_request_from_dict = SherlockUserV3DeactivateRequest.from_dict(sherlock_user_v3_deactivate_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


