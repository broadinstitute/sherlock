# SherlockRoleV3Edit


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**auto_assign_all_users** | **bool** | When true, Sherlock will automatically assign all users to this role who do not already have a role assignment | [optional] 
**can_be_glass_broken_by_role** | **int** |  | [optional] 
**default_glass_break_duration** | **str** |  | [optional] 
**grants_broad_institute_group** | **str** |  | [optional] 
**grants_dev_azure_group** | **str** |  | [optional] 
**grants_dev_firecloud_folder_owner** | **str** |  | [optional] 
**grants_dev_firecloud_group** | **str** |  | [optional] 
**grants_prod_azure_group** | **str** |  | [optional] 
**grants_prod_firecloud_folder_owner** | **str** |  | [optional] 
**grants_prod_firecloud_group** | **str** |  | [optional] 
**grants_qa_firecloud_folder_owner** | **str** |  | [optional] 
**grants_qa_firecloud_group** | **str** |  | [optional] 
**grants_sherlock_super_admin** | **bool** |  | [optional] 
**name** | **str** |  | [optional] 
**suspend_non_suitable_users** | **bool** | When true, the \&quot;suspended\&quot; field on role assignments will be computed by Sherlock based on suitability instead of being a mutable API field | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_role_v3_edit import SherlockRoleV3Edit

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockRoleV3Edit from a JSON string
sherlock_role_v3_edit_instance = SherlockRoleV3Edit.from_json(json)
# print the JSON string representation of the object
print(SherlockRoleV3Edit.to_json())

# convert the object into a dict
sherlock_role_v3_edit_dict = sherlock_role_v3_edit_instance.to_dict()
# create an instance of SherlockRoleV3Edit from a dict
sherlock_role_v3_edit_from_dict = SherlockRoleV3Edit.from_dict(sherlock_role_v3_edit_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


