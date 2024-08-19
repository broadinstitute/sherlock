# SherlockClusterV3


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**address** | **str** | Required when creating | [optional] 
**azure_subscription** | **str** | Required when creating if provider is &#39;azure&#39; | [optional] 
**base** | **str** | Required when creating | [optional] 
**ci_identifier** | [**SherlockCiIdentifierV3**](SherlockCiIdentifierV3.md) |  | [optional] 
**created_at** | **datetime** |  | [optional] 
**google_project** | **str** | Required when creating if provider is &#39;google&#39; | [optional] 
**helmfile_ref** | **str** |  | [optional] [default to 'HEAD']
**id** | **int** |  | [optional] 
**location** | **str** |  | [optional] [default to 'us-central1-a']
**name** | **str** | Required when creating | [optional] 
**provider** | **str** |  | [optional] [default to 'google']
**required_role** | **str** | If present, requires membership in the given role for mutations. Set to an empty string to clear. | [optional] 
**required_role_info** | [**SherlockRoleV3**](SherlockRoleV3.md) |  | [optional] 
**requires_suitability** | **bool** |  | [optional] 
**updated_at** | **datetime** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_cluster_v3 import SherlockClusterV3

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockClusterV3 from a JSON string
sherlock_cluster_v3_instance = SherlockClusterV3.from_json(json)
# print the JSON string representation of the object
print(SherlockClusterV3.to_json())

# convert the object into a dict
sherlock_cluster_v3_dict = sherlock_cluster_v3_instance.to_dict()
# create an instance of SherlockClusterV3 from a dict
sherlock_cluster_v3_from_dict = SherlockClusterV3.from_dict(sherlock_cluster_v3_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


