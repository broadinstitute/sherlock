# PagerdutySendAlertResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message** | **str** |  | [optional] 
**status** | **str** |  | [optional] 

## Example

```python
from sherlock_python_client.models.pagerduty_send_alert_response import PagerdutySendAlertResponse

# TODO update the JSON string below
json = "{}"
# create an instance of PagerdutySendAlertResponse from a JSON string
pagerduty_send_alert_response_instance = PagerdutySendAlertResponse.from_json(json)
# print the JSON string representation of the object
print(PagerdutySendAlertResponse.to_json())

# convert the object into a dict
pagerduty_send_alert_response_dict = pagerduty_send_alert_response_instance.to_dict()
# create an instance of PagerdutySendAlertResponse from a dict
pagerduty_send_alert_response_from_dict = PagerdutySendAlertResponse.from_dict(pagerduty_send_alert_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


