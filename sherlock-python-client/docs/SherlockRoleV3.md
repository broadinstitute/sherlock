# SherlockRoleV3


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**assignments** | [**List[SherlockRoleAssignmentV3]**](SherlockRoleAssignmentV3.md) |  | [optional] 
**auto_assign_all_users** | **bool** | When true, Sherlock will automatically assign all users to this role who do not already have a role assignment | [optional] 
**can_be_glass_broken_by_role** | **int** |  | [optional] 
**can_be_glass_broken_by_role_info** | **object** |  | [optional] 
**created_at** | **datetime** |  | [optional] 
**default_glass_break_duration** | **str** |  | [optional] 
**grants_broad_institute_group** | **str** |  | [optional] 
**grants_dev_azure_account** | **bool** |  | [optional] 
**grants_dev_azure_directory_roles** | **bool** |  | [optional] 
**grants_dev_azure_group** | **str** |  | [optional] 
**grants_dev_firecloud_folder_owner** | **str** |  | [optional] 
**grants_dev_firecloud_group** | **str** |  | [optional] 
**grants_prod_azure_account** | **bool** |  | [optional] 
**grants_prod_azure_directory_roles** | **bool** |  | [optional] 
**grants_prod_azure_group** | **str** |  | [optional] 
**grants_prod_firecloud_folder_owner** | **str** |  | [optional] 
**grants_prod_firecloud_group** | **str** |  | [optional] 
**grants_qa_firecloud_folder_owner** | **str** |  | [optional] 
**grants_qa_firecloud_group** | **str** |  | [optional] 
**grants_sherlock_super_admin** | **bool** |  | [optional] 
**id** | **int** |  | [optional] 
**name** | **str** |  | [optional] 
**suspend_non_suitable_users** | **bool** | When true, the \&quot;suspended\&quot; field on role assignments will be computed by Sherlock based on suitability instead of being a mutable API field | [optional] 
**updated_at** | **datetime** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_role_v3 import SherlockRoleV3

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockRoleV3 from a JSON string
sherlock_role_v3_instance = SherlockRoleV3.from_json(json)
# print the JSON string representation of the object
print(SherlockRoleV3.to_json())

# convert the object into a dict
sherlock_role_v3_dict = sherlock_role_v3_instance.to_dict()
# create an instance of SherlockRoleV3 from a dict
sherlock_role_v3_from_dict = SherlockRoleV3.from_dict(sherlock_role_v3_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


