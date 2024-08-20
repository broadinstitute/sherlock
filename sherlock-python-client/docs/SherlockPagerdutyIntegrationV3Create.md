# SherlockPagerdutyIntegrationV3Create


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**key** | **str** |  | [optional] 
**name** | **str** |  | [optional] 
**pagerduty_id** | **str** |  | [optional] 
**type** | **str** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_pagerduty_integration_v3_create import SherlockPagerdutyIntegrationV3Create

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockPagerdutyIntegrationV3Create from a JSON string
sherlock_pagerduty_integration_v3_create_instance = SherlockPagerdutyIntegrationV3Create.from_json(json)
# print the JSON string representation of the object
print(SherlockPagerdutyIntegrationV3Create.to_json())

# convert the object into a dict
sherlock_pagerduty_integration_v3_create_dict = sherlock_pagerduty_integration_v3_create_instance.to_dict()
# create an instance of SherlockPagerdutyIntegrationV3Create from a dict
sherlock_pagerduty_integration_v3_create_from_dict = SherlockPagerdutyIntegrationV3Create.from_dict(sherlock_pagerduty_integration_v3_create_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


