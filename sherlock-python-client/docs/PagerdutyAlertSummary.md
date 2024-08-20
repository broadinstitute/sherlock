# PagerdutyAlertSummary


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**source_link** | **str** |  | [optional] 
**summary** | **str** |  | [optional] 

## Example

```python
from sherlock_python_client.models.pagerduty_alert_summary import PagerdutyAlertSummary

# TODO update the JSON string below
json = "{}"
# create an instance of PagerdutyAlertSummary from a JSON string
pagerduty_alert_summary_instance = PagerdutyAlertSummary.from_json(json)
# print the JSON string representation of the object
print(PagerdutyAlertSummary.to_json())

# convert the object into a dict
pagerduty_alert_summary_dict = pagerduty_alert_summary_instance.to_dict()
# create an instance of PagerdutyAlertSummary from a dict
pagerduty_alert_summary_from_dict = PagerdutyAlertSummary.from_dict(pagerduty_alert_summary_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


