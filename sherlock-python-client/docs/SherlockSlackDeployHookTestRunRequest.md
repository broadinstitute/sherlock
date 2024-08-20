# SherlockSlackDeployHookTestRunRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**execute** | **bool** | Required, whether to actually send the Slack message | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_slack_deploy_hook_test_run_request import SherlockSlackDeployHookTestRunRequest

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockSlackDeployHookTestRunRequest from a JSON string
sherlock_slack_deploy_hook_test_run_request_instance = SherlockSlackDeployHookTestRunRequest.from_json(json)
# print the JSON string representation of the object
print(SherlockSlackDeployHookTestRunRequest.to_json())

# convert the object into a dict
sherlock_slack_deploy_hook_test_run_request_dict = sherlock_slack_deploy_hook_test_run_request_instance.to_dict()
# create an instance of SherlockSlackDeployHookTestRunRequest from a dict
sherlock_slack_deploy_hook_test_run_request_from_dict = SherlockSlackDeployHookTestRunRequest.from_dict(sherlock_slack_deploy_hook_test_run_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


