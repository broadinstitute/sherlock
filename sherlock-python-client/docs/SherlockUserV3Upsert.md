# SherlockUserV3Upsert


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**github_access_token** | **str** | An access token for the GitHub account to associate with the calling user. The access token isn&#39;t stored. The design here ensures that an association is only built when someone controls both accounts (Google via IAP and GitHub via this access token). | [optional] 
**name** | **str** |  | [optional] 
**name_from** | **str** |  | [optional] 

## Example

```python
from sherlock_python_client.models.sherlock_user_v3_upsert import SherlockUserV3Upsert

# TODO update the JSON string below
json = "{}"
# create an instance of SherlockUserV3Upsert from a JSON string
sherlock_user_v3_upsert_instance = SherlockUserV3Upsert.from_json(json)
# print the JSON string representation of the object
print(SherlockUserV3Upsert.to_json())

# convert the object into a dict
sherlock_user_v3_upsert_dict = sherlock_user_v3_upsert_instance.to_dict()
# create an instance of SherlockUserV3Upsert from a dict
sherlock_user_v3_upsert_from_dict = SherlockUserV3Upsert.from_dict(sherlock_user_v3_upsert_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


