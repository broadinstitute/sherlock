# SherlockClusterV3Create


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**address** | **str** | Required when creating | [optional] 
**azure_subscription** | **str** | Required when creating if provider is &#39;azure&#39; | [optional] 
**base** | **str** | Required when creating | [optional] 
**google_project** | **str** | Required when creating if provider is &#39;google&#39; | [optional] 
**helmfile_ref** | **str** |  | [optional] [default to 'HEAD']
**location** | **str** |  | [optional] [default to 'us-central1-a']
**name** | **str** | Required when creating | [optional] 
**provider** | **str** |  | [optional] [default to 'google']
**required_role** | **str** | If present, requires membership in the given role for mutations. Set to an empty string to clear. | [optional] 
**requires_suitability** | **bool** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_cluster_v3_create import SherlockClusterV3Create

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockClusterV3Create from a JSON string
sherlock_cluster_v3_create_instance = SherlockClusterV3Create.from_json(json)
# print the JSON string representation of the object
print(SherlockClusterV3Create.to_json())

# convert the object into a dict
sherlock_cluster_v3_create_dict = sherlock_cluster_v3_create_instance.to_dict()
# create an instance of SherlockClusterV3Create from a dict
sherlock_cluster_v3_create_from_dict = SherlockClusterV3Create.from_dict(sherlock_cluster_v3_create_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


