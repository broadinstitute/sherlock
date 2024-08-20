# SherlockClusterV3Edit


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**address** | **str** | Required when creating | [optional] 
**base** | **str** | Required when creating | [optional] 
**helmfile_ref** | **str** |  | [optional] [default to 'HEAD']
**required_role** | **str** | If present, requires membership in the given role for mutations. Set to an empty string to clear. | [optional] 
**requires_suitability** | **bool** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_cluster_v3_edit import SherlockClusterV3Edit

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockClusterV3Edit from a JSON string
sherlock_cluster_v3_edit_instance = SherlockClusterV3Edit.from_json(json)
# print the JSON string representation of the object
print(SherlockClusterV3Edit.to_json())

# convert the object into a dict
sherlock_cluster_v3_edit_dict = sherlock_cluster_v3_edit_instance.to_dict()
# create an instance of SherlockClusterV3Edit from a dict
sherlock_cluster_v3_edit_from_dict = SherlockClusterV3Edit.from_dict(sherlock_cluster_v3_edit_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


