# \MiscApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MyUserGet**](MiscApi.md#MyUserGet) | **Get** /my-user | Get information about the calling user
[**StatusGet**](MiscApi.md#StatusGet) | **Get** /status | Get Sherlock&#39;s current status
[**VersionGet**](MiscApi.md#VersionGet) | **Get** /version | Get Sherlock&#39;s own current version


# **MyUserGet**
> MiscMyUserResponse MyUserGet(ctx, )
Get information about the calling user

Get Sherlock's understanding of the calling user based on IAP and the Firecloud.org Google Workspace organization.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MiscMyUserResponse**](misc.MyUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatusGet**
> MiscStatusResponse StatusGet(ctx, )
Get Sherlock's current status

Get Sherlock's current status. Right now, this endpoint always returned OK (if the server is online). This endpoint is acceptable to use for a readiness check.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MiscStatusResponse**](misc.StatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VersionGet**
> MiscVersionResponse VersionGet(ctx, )
Get Sherlock's own current version

Get the build version of this Sherlock instance.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MiscVersionResponse**](misc.VersionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

