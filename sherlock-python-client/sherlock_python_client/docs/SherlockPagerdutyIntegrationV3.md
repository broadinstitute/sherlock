# SherlockPagerdutyIntegrationV3


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**created_at** | **datetime** |  | [optional] 
**id** | **int** |  | [optional] 
**name** | **str** |  | [optional] 
**pagerduty_id** | **str** |  | [optional] 
**type** | **str** |  | [optional] 
**updated_at** | **datetime** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_pagerduty_integration_v3 import SherlockPagerdutyIntegrationV3

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockPagerdutyIntegrationV3 from a JSON string
sherlock_pagerduty_integration_v3_instance = SherlockPagerdutyIntegrationV3.from_json(json)
# print the JSON string representation of the object
print(SherlockPagerdutyIntegrationV3.to_json())

# convert the object into a dict
sherlock_pagerduty_integration_v3_dict = sherlock_pagerduty_integration_v3_instance.to_dict()
# create an instance of SherlockPagerdutyIntegrationV3 from a dict
sherlock_pagerduty_integration_v3_from_dict = SherlockPagerdutyIntegrationV3.from_dict(sherlock_pagerduty_integration_v3_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


