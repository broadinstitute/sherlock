# sherlock_python_client.MiscApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**connection_check_get**](MiscApi.md#connection_check_get) | **GET** /connection-check | Test the client&#39;s connection to Sherlock
[**status_get**](MiscApi.md#status_get) | **GET** /status | Get Sherlock&#39;s current status
[**version_get**](MiscApi.md#version_get) | **GET** /version | Get Sherlock&#39;s own current version


# **connection_check_get**
> MiscConnectionCheckResponse connection_check_get()

Test the client's connection to Sherlock

Get a static response from Sherlock to verify connection through proxies like IAP.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.misc_connection_check_response import MiscConnectionCheckResponse
from sherlock_python_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://sherlock.dsp-devops-prod.broadinstitute.org
# See configuration.py for a list of all supported configuration parameters.
configuration = sherlock_python_client.Configuration(
    host = "https://sherlock.dsp-devops-prod.broadinstitute.org"
)


# Enter a context with an instance of the API client
with sherlock_python_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = sherlock_python_client.MiscApi(api_client)

    try:
        # Test the client's connection to Sherlock
        api_response = api_instance.connection_check_get()
        print("The response of MiscApi->connection_check_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MiscApi->connection_check_get: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**MiscConnectionCheckResponse**](MiscConnectionCheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **status_get**
> MiscStatusResponse status_get()

Get Sherlock's current status

Get Sherlock's current status. Right now, this endpoint always returned OK (if the server is online). This endpoint is acceptable to use for a readiness check.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.misc_status_response import MiscStatusResponse
from sherlock_python_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://sherlock.dsp-devops-prod.broadinstitute.org
# See configuration.py for a list of all supported configuration parameters.
configuration = sherlock_python_client.Configuration(
    host = "https://sherlock.dsp-devops-prod.broadinstitute.org"
)


# Enter a context with an instance of the API client
with sherlock_python_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = sherlock_python_client.MiscApi(api_client)

    try:
        # Get Sherlock's current status
        api_response = api_instance.status_get()
        print("The response of MiscApi->status_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MiscApi->status_get: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**MiscStatusResponse**](MiscStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **version_get**
> MiscVersionResponse version_get()

Get Sherlock's own current version

Get the build version of this Sherlock instance.

### Example


```python
import sherlock_python_client
from sherlock_python_client.models.misc_version_response import MiscVersionResponse
from sherlock_python_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://sherlock.dsp-devops-prod.broadinstitute.org
# See configuration.py for a list of all supported configuration parameters.
configuration = sherlock_python_client.Configuration(
    host = "https://sherlock.dsp-devops-prod.broadinstitute.org"
)


# Enter a context with an instance of the API client
with sherlock_python_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = sherlock_python_client.MiscApi(api_client)

    try:
        # Get Sherlock's own current version
        api_response = api_instance.version_get()
        print("The response of MiscApi->version_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MiscApi->version_get: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**MiscVersionResponse**](MiscVersionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

